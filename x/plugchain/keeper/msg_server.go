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
		Symbol:         types.GetSymbol(msg.Symbol),
		OriginalSymbol: msg.OriginalSymbol,
		Description:    msg.Description,
		WholeName:      msg.WholeName,
		TotalSupply:    msg.TotalSupply,
		Owner:          msg.Owner,
		Founder:        msg.Owner,
		Decimal:        msg.Decimal,
		Mintable:       msg.Mintable,
	}
	if err := k.Keeper.ValildateCreateToken(ctx, tokenMsg); err != nil {
		return &types.MsgCreateTokenResponse{}, err
	}
	err := k.Keeper.NewMintToken(ctx, tokenMsg)
	if err != nil {
		return &types.MsgCreateTokenResponse{}, err
	}
	k.Keeper.SetToken(ctx, tokenMsg, true)
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

func (k msgServer) BurnToken(goCtx context.Context, msg *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	burnMsg := types.MsgBurnToken{
		Symbol:  types.GetSymbol(msg.Symbol),
		Address: msg.Address,
		Account: msg.Account,
	}
	//查询币信息
	tokenMsg := k.Keeper.GetToken(ctx, types.GetTokenKey(burnMsg.Symbol))
	if tokenMsg.Owner != burnMsg.Address {
		return &types.MsgBurnTokenResponse{}, types.ReturnErrAccAddressNotPermission(burnMsg.Address)
	}
	if !tokenMsg.Mintable {
		return &types.MsgBurnTokenResponse{}, types.ErrBurnTokenNotPermission
	}
	burnAddr, err := sdk.AccAddressFromBech32(burnMsg.Address)
	if err != nil {
		panic(err)
	}

	//执行造币
	mintCoin := sdk.NewCoin(burnMsg.Symbol, *burnMsg.Account)
	mintCoins := sdk.NewCoins(mintCoin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return &types.MsgBurnTokenResponse{}, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, burnAddr, mintCoins); err != nil {
		return &types.MsgBurnTokenResponse{}, err
	}
	newTotal := sdk.NewIntFromBigInt(tokenMsg.TotalSupply.BigInt().Add(tokenMsg.TotalSupply.BigInt(), burnMsg.Account.BigInt()))
	tokenMsg.TotalSupply = &newTotal

	k.SetToken(ctx, tokenMsg, false)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeBurnToken,
			sdk.NewAttribute(types.AttributeValueTokenSymbol, burnMsg.Symbol),
			sdk.NewAttribute(types.AttributeValueTokenOwner, burnMsg.Address),
			sdk.NewAttribute(types.AttributeValueTokenBurnAccount, burnMsg.Account.String()),
		),
	})
	return &types.MsgBurnTokenResponse{}, nil
}
