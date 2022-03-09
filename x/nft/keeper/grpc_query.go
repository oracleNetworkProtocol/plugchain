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

func (q Keeper) Class(c context.Context, req *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	data, ok := q.GetClassByID(ctx, req.ClassId)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrInvalidClass, "class ID %s not exists", req.ClassId)
	}

	return &types.QueryClassResponse{Class: &data}, nil
}

func (q Keeper) Classes(c context.Context, req *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var denoms []types.Class
	store := ctx.KVStore(q.storeKey)
	denomStore := prefix.NewStore(store, types.GetKeyClassID(""))
	pageRes, err := query.Paginate(denomStore, req.Pagination, func(key, value []byte) error {
		var denom types.Class
		q.cdc.MustUnmarshal(value, &denom)
		denoms = append(denoms, denom)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
	}
	return &types.QueryClassesResponse{
		Classes:    denoms,
		Pagination: pageRes,
	}, nil
}

func (q Keeper) NFT(c context.Context, req *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	nft, err := q.GetNFT(ctx, req.ClassId, req.NftId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrKnownNFT, "invalid NFT %s from collection %s", req.NftId, req.ClassId)
	}
	nfti, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrKnownNFT, "invalid type NFT %s from collection %s", req.NftId, req.ClassId)
	}
	return &types.QueryNFTResponse{
		Nft: &nfti,
	}, nil
}

func (q Keeper) NFTs(c context.Context, req *types.QueryNFTsRequest) (*types.QueryNFTsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	var err error
	if len(req.ClassId) > 0 {
		if err := types.ValidateClassID(req.ClassId); err != nil {
			return nil, err
		}
	}
	var owner sdk.AccAddress
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, err
		}
	}
	var (
		nfts    []*types.NFT
		pageRes *query.PageResponse
	)
	ctx := sdk.UnwrapSDKContext(c)

	switch {
	case len(req.ClassId) > 0 && len(req.Owner) > 0:
		if pageRes, err = query.Paginate(q.getStoreByOwnerClass(ctx, owner, req.ClassId), req.Pagination, func(key []byte, _ []byte) error {
			nft, err := q.GetNFT(ctx, req.ClassId, string(key))
			if err != nil {
				return err
			}
			nfti, ok := nft.(types.NFT)
			if ok {
				nfts = append(nfts, &nfti)

			}
			return nil

		}); err != nil {
			return nil, err
		}
	case len(req.ClassId) > 0 && len(req.Owner) == 0:
		nftStore := q.getStoreByClass(ctx, req.ClassId)
		if pageRes, err = query.Paginate(nftStore, req.Pagination, func(key, value []byte) error {
			var nft types.NFT
			if err := q.cdc.Unmarshal(value, &nft); err != nil {
				return err
			}
			nfts = append(nfts, &nft)
			return nil
		}); err != nil {
			return nil, err
		}
	case len(req.ClassId) == 0 && len(req.Owner) > 0:
		ownerStore := q.getStoreByOwner(ctx, owner)
		if pageRes, err = query.Paginate(ownerStore, req.Pagination, func(key, value []byte) error {
			classID, nftID, _ := types.SplitKeyDenom(key)
			nft, err := q.GetNFT(ctx, classID, nftID)
			if err != nil {
				return err
			}
			nfti, ok := nft.(types.NFT)
			if ok {
				nfts = append(nfts, &nfti)

			}
			return nil
		}); err != nil {
			return nil, err
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest.Wrap("must provide at least one of classID or owner")
	}

	return &types.QueryNFTsResponse{
		Nfts:       nfts,
		Pagination: pageRes,
	}, nil
}

func (q Keeper) Supply(c context.Context, req *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, found := q.GetClassByID(ctx, req.ClassId)
	if !found {
		return nil, status.Errorf(codes.InvalidArgument, "ClassID %s not existed", req.ClassId)
	}

	var supply = q.GetTotalSupply(ctx, req.ClassId)
	return &types.QuerySupplyResponse{Amount: supply}, nil
}

func (q Keeper) Owner(c context.Context, req *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	ownerAdr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", req.Address)
	}
	owner := types.Owner{
		Address:       ownerAdr.String(),
		CollectionIDs: []types.CollectionID{},
	}
	idsMap := make(map[string][]string)

	store := ctx.KVStore(q.storeKey)
	nftStore := prefix.NewStore(store, types.GetKeyOwner(ownerAdr, req.ClassId, ""))

	pageRes, err := query.Paginate(nftStore, req.Pagination, func(key, value []byte) error {
		classID := req.ClassId
		nftID := string(key)
		if len(classID) == 0 {
			classID, nftID, _ = types.SplitKeyDenom(key)
		}

		//Injection classID to nftIDs
		if ids, ok := idsMap[classID]; ok {
			idsMap[classID] = append(ids, nftID)
		} else {
			idsMap[classID] = []string{nftID}
			owner.CollectionIDs = append(owner.CollectionIDs, types.CollectionID{ClassID: classID})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	//assignment value
	idsCount := len(owner.CollectionIDs)
	for i := 0; i < idsCount; i++ {
		owner.CollectionIDs[i].NFTIDs = idsMap[owner.CollectionIDs[i].ClassID]
	}
	return &types.QueryOwnerResponse{Owner: &owner, Pagination: pageRes}, nil
}
