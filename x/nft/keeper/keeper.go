package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: ak,
		bankKeeper:    bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IssueClass(ctx sdk.Context, id, name, schema, symbol string, owner sdk.AccAddress, mintRestricted, editRestricted bool) error {
	denom := types.NewClass(id, name, schema, symbol, owner, mintRestricted, editRestricted)
	return k.SetClass(ctx, denom)
}

func (k Keeper) IssueNFT(ctx sdk.Context, classID, ID, name, url, data string, owner sdk.AccAddress) error {
	denom, ok := k.GetClassByID(ctx, classID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "class ID (%s) not exists", classID)
	}
	if denom.MintRestricted && denom.Owner != owner.String() {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to issue NFT of denom %s", denom.Owner, classID)
	}
	if k.HasNFTByID(ctx, classID, ID) {
		return sdkerrors.Wrapf(types.ErrNFTAreadyExists, "NFT %s already exists in collection %s", ID, classID)
	}
	k.setNFT(ctx, classID,
		types.NewNFT(
			ID, name, url, data, owner,
		))
	// count++
	k.increaseSupply(ctx, classID)
	k.setOwner(ctx, classID, ID, owner)
	return nil
}

func (k Keeper) EditNFT(ctx sdk.Context, classID, ID, name, url, data string, owner sdk.AccAddress) error {
	denom, ok := k.GetClassByID(ctx, classID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "denom ID (%s) not exists", classID)
	}

	if denom.EditRestricted {
		// if true , nobody can edit the NFT under this denom
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "nobody can edit the NFT under this denom %s", denom.ID)
	}

	nft, err := k.Authorize(ctx, classID, ID, owner)
	if err != nil {
		return err
	}

	if types.Modified(name) {
		nft.Name = name
	}

	if types.Modified(url) {
		nft.URL = url
	}

	if types.Modified(data) {
		nft.Data = data
	}

	k.setNFT(ctx, classID, nft)

	return nil
}

func (k Keeper) BurnNFT(ctx sdk.Context, classID, ID string, owner sdk.AccAddress) error {
	if !k.HasClassByID(ctx, classID) {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "class ID %s not exists", classID)
	}
	_, err := k.Authorize(ctx, classID, ID, owner)
	if err != nil {
		return err
	}

	k.deleteNFT(ctx, classID, ID)
	k.decreaseSupply(ctx, classID)
	k.delOwner(ctx, classID, ID, owner)

	return nil
}

func (k Keeper) TransferNFTToOwner(ctx sdk.Context, owner, recipient sdk.AccAddress, nftID, classID string) error {
	_, ok := k.GetClassByID(ctx, classID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "class ID %s not exists", classID)
	}

	nft, err := k.Authorize(ctx, classID, nftID, owner)
	if err != nil {
		return err
	}

	nft.Owner = recipient.String()

	k.setNFT(ctx, classID, nft)
	k.swapOwner(ctx, classID, nftID, owner, recipient)
	return nil
}
func (k Keeper) TransferClassToOwner(ctx sdk.Context, owner, recipient sdk.AccAddress, classID string) error {
	denom, ok := k.GetClassByID(ctx, classID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "class ID %s not exists", classID)
	}
	if owner.String() != denom.Owner {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer class %s", owner.String(), classID)
	}
	denom.Owner = recipient.String()
	err := k.UpdateDenom(ctx, denom)
	if err != nil {
		return err
	}
	return nil
}
