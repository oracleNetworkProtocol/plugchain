package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/oracleNetworkProtocol/plugchain/x/token/types"
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

func (m msgServer) IssueToken(c context.Context, in *types.MsgIssueToken) (*types.MsgIssueTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}

	if m.Keeper.moduleAddrs[in.Owner] {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is a module account", in.Owner)
	}

	ctx := sdk.UnwrapSDKContext(c)

	// deduct issue fee
	err = m.DeductIssueTokenFee(ctx, owner, in.Symbol)
	if err != nil {
		return nil, err
	}

	if err := m.Keeper.IssueToken(
		ctx, in.Symbol, in.Name, in.MinUnit, in.Scale, in.InitialSupply, in.MaxSupply, in.Mintable, owner,
	); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueToken,
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
			sdk.NewAttribute(types.AttributeKeySymbol, in.Symbol),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgIssueTokenResponse{}, nil
}

func (m msgServer) MintToken(c context.Context, in *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)
	recipient := owner
	if len(in.To) != 0 {
		recipient, err = sdk.AccAddressFromBech32(in.To)
		if err != nil {
			return nil, err
		}
	}

	if m.Keeper.moduleAddrs[recipient.String()] {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is a module account", recipient.String())
	}

	err = m.Keeper.DeductOperateTokenFee(ctx, owner, in.Symbol)
	if err != nil {
		return nil, err
	}
	err = m.Keeper.MintToken(ctx, in.Symbol, in.Amount, recipient, owner)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintTokenResponse{}, nil
}

func (m msgServer) EditToken(c context.Context, in *types.MsgEditToken) (*types.MsgEditTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	err = m.Keeper.DeductOperateTokenFee(ctx, owner, in.Symbol)
	if err != nil {
		return nil, err
	}

	if err := m.Keeper.EditToken(ctx, in.Symbol, in.Name, in.MaxSupply, owner); err != nil {
		return nil, err
	}
	//事件标注
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditToken,
			sdk.NewAttribute(types.AttributeKeySymbol, in.Symbol),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			types.TypeMsgEditToken,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgEditTokenResponse{}, nil
}

func (m msgServer) BurnToken(c context.Context, in *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	err = m.Keeper.DeductOperateTokenFee(ctx, owner, in.Symbol)
	if err != nil {
		return nil, err
	}
	err = m.Keeper.BurnToken(ctx, in.Symbol, in.Amount, owner)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnToken,
			sdk.NewAttribute(types.AttributeKeySymbol, in.Symbol),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
			sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatUint(in.Amount, 10)+in.Symbol),
		),
		sdk.NewEvent(
			types.TypeMsgBurnToken,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgBurnTokenResponse{}, nil
}

func (m msgServer) TransferOwnerToken(c context.Context, in *types.MsgTransferOwnerToken) (*types.MsgTransferOwnerTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	to, err := sdk.AccAddressFromBech32(in.To)
	if err != nil {
		return nil, err
	}
	if m.Keeper.moduleAddrs[in.To] {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is a module account", in.To)
	}

	if err := m.Keeper.DeductOperateTokenFee(ctx, owner, in.Symbol); err != nil {
		return nil, err
	}

	if err := m.Keeper.TransferOwnerToken(ctx, in.Symbol, owner, to); err != nil {
		return nil, err
	}

	return &types.MsgTransferOwnerTokenResponse{}, nil
}
