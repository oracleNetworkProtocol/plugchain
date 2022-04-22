package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) HasNFTByID(ctx sdk.Context, classID, nftID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetKeyNFT(classID, nftID))
}

func (k Keeper) setNFT(ctx sdk.Context, classID string, nft types.NFT) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&nft)
	store.Set(types.GetKeyNFT(classID, nft.ID), bz)
}

func (k Keeper) GetNFTs(ctx sdk.Context, classID string) (nfts []types.NFTI) {

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetKeyNFT(classID, ""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var nft types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &nft)
		nfts = append(nfts, nft)
	}
	return nfts
}

func (k Keeper) GetNFT(ctx sdk.Context, classID, ID string) (types.NFTI, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetKeyNFT(classID, ID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found NFT: %s", ID)
	}
	var nft types.NFT
	k.cdc.MustUnmarshal(bz, &nft)
	return nft, nil
}

//Authorize checks if the sender is the owner of the given NFT
func (k Keeper) Authorize(ctx sdk.Context, classID, ID string, owner sdk.AccAddress) (types.NFT, error) {
	nft, err := k.GetNFT(ctx, classID, ID)
	if err != nil {
		return types.NFT{}, err
	}

	if !owner.Equals(nft.GetOwner()) {
		return types.NFT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return nft.(types.NFT), nil
}

func (k Keeper) deleteNFT(ctx sdk.Context, classID, ID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyNFT(classID, ID))
}

func (k Keeper) getStoreByOwnerClass(ctx sdk.Context, owner sdk.AccAddress, classID string) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOwner(owner, classID, "")
	return prefix.NewStore(store, key)
}

func (k Keeper) getStoreByClass(ctx sdk.Context, classID string) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetKeyClassID(classID))
}

func (k Keeper) getStoreByOwner(ctx sdk.Context, owner sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetKeyOwner(owner, "", ""))
}
