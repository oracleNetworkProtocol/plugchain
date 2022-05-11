package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, v := range collection.NFTs {
		if err := k.IssueNFT(ctx, collection.Class.ID, v.ID, v.Name, v.URI, v.Data, v.GetOwner()); err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) GetCollections(ctx sdk.Context) (list []types.Collection) {
	denoms := k.GetClasses(ctx)
	if len(denoms) == 0 {
		return
	}
	for _, v := range denoms {
		nfts := k.GetNFTs(ctx, v.ID)
		list = append(list, types.NewCollection(v, nfts))
	}
	return
}

func (k Keeper) GetPaginateCollection(ctx sdk.Context, req *types.QueryNFTsRequest, classID string) ([]types.NFTI, *query.PageResponse, error) {
	_, found := k.GetClassByID(ctx, classID)
	if !found {
		return nil, nil, sdkerrors.Wrapf(types.ErrInvalidClass, "classID %s not existed", classID)
	}
	var nfts []types.NFTI
	store := ctx.KVStore(k.storeKey)
	nftStore := prefix.NewStore(store, types.GetKeyNFT(classID, ""))
	pageRes, err := query.Paginate(nftStore, req.Pagination, func(key, value []byte) error {
		var nft types.NFT
		k.cdc.MustUnmarshal(value, &nft)
		nfts = append(nfts, nft)
		return nil
	})
	if err != nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
	}

	return nfts, pageRes, nil
}

// get denom count
func (k Keeper) GetTotalSupply(ctx sdk.Context, classID string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollectionByClassID(classID))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)

}

// supply++
func (k Keeper) increaseSupply(ctx sdk.Context, classID string) {
	supply := k.GetTotalSupply(ctx, classID)
	supply++
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollectionByClassID(classID), bz)
}

// supply--
func (k Keeper) decreaseSupply(ctx sdk.Context, classID string) {
	supply := k.GetTotalSupply(ctx, classID)
	supply--
	store := ctx.KVStore(k.storeKey)

	if supply == 0 {
		store.Delete(types.KeyCollectionByClassID(classID))
		return
	}

	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollectionByClassID(classID), bz)
}
