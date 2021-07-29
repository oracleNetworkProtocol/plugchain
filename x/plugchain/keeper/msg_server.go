package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenMsg := types.MsgCreateToken{
		Symbol:         msg.Symbol,
		OriginalSymbol: msg.OriginalSymbol,
		Description:    msg.Description,
		WholeName:      msg.WholeName,
		TotalSupply:    msg.TotalSupply,
		Owner:          msg.Owner,
		Founder:        msg.Owner,
		Decimal:        msg.Decimal,
		Mintable:       msg.Mintable,
	}

	err := k.Keeper.NewMintToken(ctx, tokenMsg)
	if err != nil {
		return &types.MsgCreateTokenResponse{}, err
	}
	k.Keeper.SetToken(ctx, tokenMsg)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeCreateToken,
			sdk.NewAttribute(types.AttributeValueTokenSymbol, tokenMsg.Symbol),
			sdk.NewAttribute(types.AttributeValueTokenOriginalSymbol, tokenMsg.OriginalSymbol),
			sdk.NewAttribute(types.AttributeValueTokenDescription, tokenMsg.Description),
			sdk.NewAttribute(types.AttributeValueTokenWholeName, tokenMsg.OriginalSymbol),
			sdk.NewAttribute(types.AttributeValueTokenOwner, tokenMsg.Owner),
			sdk.NewAttribute(types.AttributeValueTokenDecimal, strconv.FormatUint(tokenMsg.Decimal, 10)),
			sdk.NewAttribute(types.AttributeValueTokenTotalSupply, tokenMsg.TotalSupply.String()),
		),
	})
	k.Keeper.Logger(ctx).Info(
		"create token msg success,attr for:",
		"symbol", tokenMsg.Symbol,
		"original_symbol", tokenMsg.OriginalSymbol,
		"whole_name", tokenMsg.WholeName,
		"description", tokenMsg.Description,
		"total_supply", tokenMsg.TotalSupply,
		"decimal", tokenMsg.Decimal,
		"owner", tokenMsg.Owner,
		"mintable", tokenMsg.Mintable,
	)
	return &types.MsgCreateTokenResponse{}, nil
}
