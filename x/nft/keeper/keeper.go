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

func (k Keeper) IssueDenom(ctx sdk.Context, id, name, schema, symbol string, owner sdk.AccAddress, mintRestricted, editRestricted bool) error {
	denom := types.NewDenom(id, name, schema, symbol, owner, mintRestricted, editRestricted)
	return k.SetDenom(ctx, denom)
}

func (k Keeper) IssueNFT(ctx sdk.Context, denomID, ID, name, url, data string, owner sdk.AccAddress) error {
	denom, ok := k.GetDenomByID(ctx, denomID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID (%s) not exists", denomID)
	}
	if denom.MintRestricted && denom.Owner != owner.String() {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to issue NFT of denom %s", denom.Owner, denomID)
	}
	if k.HasNFTByID(ctx, denomID, ID) {
		return sdkerrors.Wrapf(types.ErrNFTAreadyExists, "NFT %s already exists in collection %s", ID, denomID)
	}
	k.setNFT(ctx, denomID,
		types.NewNFT(
			ID, name, url, data, owner,
		))

	return nil
}

func (k Keeper) EditNFT(ctx sdk.Context, denomID, ID, name, url, data string, owner sdk.AccAddress) error {
	denom, ok := k.GetDenomByID(ctx, denomID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID (%s) not exists", denomID)
	}

	if denom.EditRestricted {
		// if true , nobody can edit the NFT under this denom
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "nobody can edit the NFT under this denom %s", denom.ID)
	}

	nft, err := k.Authorize(ctx, denomID, ID, owner)
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

	k.setNFT(ctx, denomID, nft)

	return nil
}
