package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
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
	return &types.QueryDenomsResponse{}, nil
}
