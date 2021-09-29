package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
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

func (m msgServer) IssueDenom(c context.Context, in *types.MsgIssueDenom) (*types.MsgIssueDenomResponse, error) {

	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.IssueDenom(ctx, in.ID, in.Name, in.Schema, in.Symbol, owner, in.MintRestricted, in.EditRestricted); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueDenom,
			sdk.NewAttribute(types.AttributeKeyDenomID, in.ID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgIssueDenomResponse{}, nil
}

func (m msgServer) IssueNFT(c context.Context, in *types.MsgIssueNFT) (*types.MsgIssueNFTResponse, error) {
	recipient, err := sdk.AccAddressFromBech32(in.Recipient)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.IssueNFT(ctx, in.DenomID, in.ID, in.Name, in.URL, in.Data, recipient); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueNFT,
			sdk.NewAttribute(types.AttributeKeyDenomID, in.DenomID),
			sdk.NewAttribute(types.AttributeKeyNFTID, in.ID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgIssueNFTResponse{}, nil
}

func (m msgServer) EditNFT(c context.Context, in *types.MsgEditNFT) (*types.MsgEditNFTResponse, error) {

	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.EditNFT(ctx, in.DenomID, in.ID, in.Name, in.URL, in.Data, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFT,
			sdk.NewAttribute(types.AttributeKeyDenomID, in.DenomID),
			sdk.NewAttribute(types.AttributeKeyNFTID, in.ID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})
	return &types.MsgEditNFTResponse{}, nil
}
