package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) HasNFTByID(ctx sdk.Context, denomID, nftID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetKeyNFT(denomID, nftID))
}

func (k Keeper) setNFT(ctx sdk.Context, denomID string, nft types.NFT) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&nft)
	store.Set(types.GetKeyNFT(denomID, nft.ID), bz)
}

func (k Keeper) GetNFTs(ctx sdk.Context, denomID string) (nfts []types.NFTI) {
	return nil
}

func (k Keeper) GetNFT(ctx sdk.Context, denomID, ID string) (types.NFTI, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetKeyNFT(denomID, ID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found NFT: %s", ID)
	}
	var nft types.NFT
	k.cdc.MustUnmarshalBinaryBare(bz, &nft)
	return nft, nil
}

//Authorize checks if the sender is the owner of the given NFT
func (k Keeper) Authorize(ctx sdk.Context, denomID, ID string, owner sdk.AccAddress) (types.NFT, error) {
	nft, err := k.GetNFT(ctx, denomID, ID)
	if err != nil {
		return types.NFT{}, err
	}

	if !owner.Equals(nft.GetOwner()) {
		return types.NFT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return nft.(types.NFT), nil
}
