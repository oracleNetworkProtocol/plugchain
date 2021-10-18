package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) setOwner(ctx sdk.Context, denomID, nftID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalNFTID(k.cdc, nftID)
	store.Set(types.GetKeyOwner(owner, denomID, nftID), bz)
}

func (k Keeper) delOwner(ctx sdk.Context, denomID, nftID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyOwner(owner, denomID, nftID))
}

func (k Keeper) swapOwner(ctx sdk.Context, denomID, nftID string, owner, recipient sdk.AccAddress) {
	k.delOwner(ctx, denomID, nftID, owner)

	k.setOwner(ctx, denomID, nftID, recipient)
}
