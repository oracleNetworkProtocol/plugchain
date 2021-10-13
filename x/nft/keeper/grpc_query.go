package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) Denom(c context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	data, ok := q.GetDenomByID(ctx, req.DenomId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", req.DenomId)
	}

	return &types.QueryDenomResponse{Denom: &data}, nil
}

func (q Keeper) Denoms(c context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var denoms []types.Denom
	store := ctx.KVStore(q.storeKey)
	denomStore := prefix.NewStore(store, types.GetKeyDenomID(""))
	pageRes, err := query.Paginate(denomStore, req.Pagination, func(key, value []byte) error {
		var denom types.Denom
		q.cdc.MustUnmarshalBinaryBare(value, &denom)
		denoms = append(denoms, denom)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
	}
	return &types.QueryDenomsResponse{
		Denoms:     denoms,
		Pagination: pageRes,
	}, nil
}

func (q Keeper) NFT(c context.Context, req *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	nft, err := q.GetNFT(ctx, req.DenomId, req.NftId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrKnownNFT, "invalid NFT %s from collection %s", req.NftId, req.DenomId)
	}
	nfti, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrKnownNFT, "invalid type NFT %s from collection %s", req.NftId, req.DenomId)
	}
	return &types.QueryNFTResponse{
		Nft: &nfti,
	}, nil
}

func (q Keeper) Collection(c context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	collection, pageRes, err := q.GetPaginateCollection(ctx, req, req.DenomId)
	if err != nil {
		return nil, err
	}
	return &types.QueryCollectionResponse{
		Collection: &collection,
		Pagination: pageRes,
	}, nil
}

func (q Keeper) Supply(c context.Context, req *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var supply = q.GetTotalSupply(ctx, req.DenomId)
	return &types.QuerySupplyResponse{Amount: supply}, nil
}
