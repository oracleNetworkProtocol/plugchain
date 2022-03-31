package keeper

import (
	"context"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gogo/protobuf/proto"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/oracleNetworkProtocol/plugchain/x/prc10/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) Token(c context.Context, req *types.QueryTokenRequest) (*types.QueryTokenResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	token, err := q.GetToken(ctx, req.Denom)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "token %s not found", req.Denom)
	}
	msg, ok := token.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "can't protomarshal %T", token)
	}
	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryTokenResponse{Token: any}, nil
}

func (q Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParamSet(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q Keeper) Tokens(c context.Context, req *types.QueryTokensRequest) (*types.QueryTokensResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	var (
		owner sdk.AccAddress
		err   error
	)
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid owner address (%s)", err))
		}
	}
	var (
		tokens  []types.TokenI
		pageRes *query.PageResponse
	)
	store := ctx.KVStore(q.storeKey)

	if owner == nil {
		tokenStore := prefix.NewStore(store, types.PrefixTokenForSymbol)
		pageRes, err = query.Paginate(tokenStore, req.Pagination, func(key, value []byte) error {
			var token types.Token
			q.cdc.MustUnmarshal(value, &token)
			tokens = append(tokens, &token)
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
		}
	} else {
		tokenStore := prefix.NewStore(store, types.KeyTokens(owner, ""))
		pageRes, err = query.Paginate(tokenStore, req.Pagination, func(key, value []byte) error {
			var symbol gogotypes.StringValue
			q.cdc.MustUnmarshal(value, &symbol)
			token, err := q.GetToken(ctx, symbol.Value)
			if err == nil {
				tokens = append(tokens, token)
			}
			return err
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate:%v", err)
		}
	}

	result := make([]*codectypes.Any, len(tokens))
	for i, v := range tokens {
		msg, ok := v.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "%T does not implement proto.Message", v)
		}
		var err error
		if result[i], err = codectypes.NewAnyWithValue(msg); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &types.QueryTokensResponse{Tokens: result, Pagination: pageRes}, nil
}

func (q Keeper) Fees(c context.Context, req *types.QueryFeesRequest) (*types.QueryFeesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	if err := types.ValidateSymbol(req.Symbol); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	issueFee, err := q.GetIssueTokenFee(ctx, req.Symbol)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	operateFee, err := q.GetOperateTokenFee(ctx, req.Symbol)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &types.QueryFeesResponse{
		Exist:      q.HasSymbol(ctx, req.Symbol),
		IssueFee:   issueFee,
		OperateFee: operateFee,
	}, nil
}
