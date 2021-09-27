package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) HasNFTByID(ctx sdk.Context, denomID, nftID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetKeyNFT(denomID, nftID))
}

func (k Keeper) SetNFT(ctx sdk.Context, denomID string, nft types.NFT) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&nft)
	store.Set(types.GetKeyNFT(denomID, nft.ID), bz)
}
