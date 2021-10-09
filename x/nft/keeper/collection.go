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

func (k Keeper) GetCollection(ctx sdk.Context, denomID string) (types.Collection, error) {
	denom, ok := k.GetDenomByID(ctx, denomID)
	if !ok {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID (%s) not existed", denomID)
	}
	nfts := k.GetNFTs(ctx, denomID)
	return types.NewCollection(denom, nfts), nil
}

func (k Keeper) GetPaginateCollection(ctx sdk.Context, req *types.QueryCollectionRequest, denomID string) (types.Collection, *query.PageResponse, error) {
	denom, found := k.GetDenomByID(ctx, denomID)
	if !found {
		return types.Collection{}, nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not existed", denomID)
	}
	var nfts []types.NFTI
	store := ctx.KVStore(k.storeKey)
	nftStore := prefix.NewStore(store, types.GetKeyNFT(denomID, ""))
	pageRes, err := query.Paginate(nftStore, req.Pagination, func(key, value []byte) error {
		var nft types.NFT
		k.cdc.MustUnmarshalBinaryBare(value, &nft)
		nfts = append(nfts, nft)
		return nil
	})
	if err != nil {
		return types.Collection{}, nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
	}
	return types.NewCollection(denom, nfts), pageRes, nil
}
