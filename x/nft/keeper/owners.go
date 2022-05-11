package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) setOwner(ctx sdk.Context, classID, nftID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalNFTID(k.cdc, nftID)
	store.Set(types.GetKeyOwner(owner, classID, nftID), bz)
}

func (k Keeper) delOwner(ctx sdk.Context, classID, nftID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyOwner(owner, classID, nftID))
}

func (k Keeper) swapOwner(ctx sdk.Context, classID, nftID string, owner, recipient sdk.AccAddress) {
	k.delOwner(ctx, classID, nftID, owner)

	k.setOwner(ctx, classID, nftID, recipient)
}
