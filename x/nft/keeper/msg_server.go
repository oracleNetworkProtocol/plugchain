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

func (m msgServer) IssueClass(c context.Context, in *types.MsgIssueClass) (*types.MsgIssueClassResponse, error) {

	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.IssueClass(ctx, in.ID, in.Name, in.Schema, in.Symbol, owner, in.MintRestricted, in.EditRestricted); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueClass,
			sdk.NewAttribute(types.AttributeKeyClassID, in.ID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgIssueClassResponse{}, nil
}

func (m msgServer) IssueNFT(c context.Context, in *types.MsgIssueNFT) (*types.MsgIssueNFTResponse, error) {
	recipient, err := sdk.AccAddressFromBech32(in.Recipient)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.IssueNFT(ctx, in.ClassID, in.ID, in.Name, in.URI, in.Data, recipient); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueNFT,
			sdk.NewAttribute(types.AttributeKeyClassID, in.ClassID),
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

	if err := m.Keeper.EditNFT(ctx, in.ClassID, in.ID, in.Name, in.URI, in.Data, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFT,
			sdk.NewAttribute(types.AttributeKeyClassID, in.ClassID),
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

func (m msgServer) BurnNFT(c context.Context, in *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {

	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.BurnNFT(ctx, in.ClassID, in.ID, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnNFT,
			sdk.NewAttribute(types.AttributeKeyNFTID, in.ID),
			sdk.NewAttribute(types.AttributeKeyClassID, in.ClassID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgBurnNFTResponse{}, nil
}

func (m msgServer) TransferNFT(c context.Context, in *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(in.Recipient)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.TransferNFTToOwner(ctx, owner, recipient, in.ID, in.ClassID); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferNFT,
			sdk.NewAttribute(types.AttributeKeyNFTID, in.ID),
			sdk.NewAttribute(types.AttributeKeyClassID, in.ClassID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
			sdk.NewAttribute(types.AttributeKeyRecipient, in.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgTransferNFTResponse{}, nil
}

func (m msgServer) TransferClass(c context.Context, in *types.MsgTransferClass) (*types.MsgTransferClassResponse, error) {
	owner, err := sdk.AccAddressFromBech32(in.Owner)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(in.Recipient)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	if err := m.Keeper.TransferClassToOwner(ctx, owner, recipient, in.ID); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferClass,
			sdk.NewAttribute(types.AttributeKeyClassID, in.ID),
			sdk.NewAttribute(types.AttributeKeyOwner, in.Owner),
			sdk.NewAttribute(types.AttributeKeyRecipient, in.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, in.Owner),
		),
	})

	return &types.MsgTransferClassResponse{}, nil
}
