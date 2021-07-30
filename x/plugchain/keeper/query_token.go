package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q Querier) TokenList(goCtx context.Context, req *types.QueryTokenListRequest) (*types.QueryTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(q.storeKey)
	tokenStore := prefix.NewStore(store, types.TokenKeyPrefix)
	var tokens []types.MsgCreateToken

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key, value []byte) error {
		var token types.MsgCreateToken
		err := q.cdc.UnmarshalBinaryBare(value, &token)
		if err != nil {
			return err
		}
		tokens = append(tokens, token)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(tokens) == 0 {
		return nil, status.Error(codes.NotFound, "There are no tokens present.")
	}

	return &types.QueryTokenListResponse{
		Tokens:     tokens,
		Pagination: pageRes,
	}, nil
}

func (q Querier) TokensByAccAddr(goCtx context.Context, req *types.QueryTokensByAccAddrRequest) (*types.QueryTokensByAccAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.NotFound, "invalid address")
	}
	accAdrBech32, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "address is invalid (%s)", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenKeys := q.GetAccTokens(ctx, accAdrBech32)
	var tokenList []types.MsgCreateToken
	if len(tokenKeys.Key) != 0 {
		for _, v := range tokenKeys.Key {
			val := q.GetToken(ctx, v)
			tokenList = append(tokenList, val)
		}
	}
	return &types.QueryTokensByAccAddrResponse{Token: tokenList}, nil
}

func (q Querier) TokenInfo(goCtx context.Context, req *types.QueryTokenInfoRequest) (*types.QueryTokenInfoResponse, error) {

	if req.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid symbol")
	}
	symbol := types.GetSymbol(req.Symbol)

	ctx := sdk.UnwrapSDKContext(goCtx)
	key := types.GetTokenKey(symbol)
	token := q.GetToken(ctx, key)

	return &types.QueryTokenInfoResponse{Token: token}, nil
}
