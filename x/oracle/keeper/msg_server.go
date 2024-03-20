package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claim := msg.GetClaim()
	signer := msg.MustGetSigner()
	valAddr := getValidatorAddr(ctx, k, signer)

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, valAddr)
	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, valAddr.String())
	}

	prevoteHash, err := k.isCorrectRound(ctx, msg, signer)
	if err != nil {
		return nil, err
	}

	// store the validator vote
	k.CreateVote(ctx, claim, valAddr)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeVote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyClaimHash, claim.Hash().String()),
		),
	)

	if len(prevoteHash) != 0 {
		k.DeletePrevote(ctx, prevoteHash)
	}

	return &types.MsgVoteResponse{
		Hash: claim.Hash(),
	}, nil
}

// Delegate implements types.MsgServer
func (k msgServer) Delegate(c context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	val, del := msg.MustGetValidator(), msg.MustGetDelegate()

	if k.Keeper.StakingKeeper.Validator(ctx, sdk.ValAddress(val)) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}

	k.SetValidatorDelegateAddress(ctx, val, del)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDelegate),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyDelegate, msg.Delegate),
		),
	)

	return &types.MsgDelegateResponse{}, nil
}

func (k msgServer) Prevote(goCtx context.Context, msg *types.MsgPrevote) (*types.MsgPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := msg.MustGetSigner()

	valAddr := getValidatorAddr(ctx, k, signer)

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, valAddr)
	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, valAddr.String())
	}

	k.CreatePrevote(ctx, msg.Hash)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypePrevote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyPrevoteHash, fmt.Sprintf("%x", msg.Hash)),
		),
	)

	return &types.MsgPrevoteResponse{}, nil
}

// HELPERS

func (k msgServer) isCorrectRound(ctx sdk.Context, msg *types.MsgVote, signer sdk.AccAddress) ([]byte, error) {
	claim := msg.GetClaim()
	claimType := claim.Type()

	// will return empty struct if doesn't exist
	claimParams := k.ClaimParamsForType(ctx, claimType)
	if claimParams.ClaimType != claimType {
		return nil, sdkerrors.Wrap(types.ErrNoClaimTypeExists, claimType)
	}

	claimRoundID := claim.GetRoundID()
	lastFinalizedRound := k.GetLastFinalizedRound(ctx, claimType)

	// RoundID should be greater than the LastFinalizedRound
	if claimRoundID <= lastFinalizedRound {
		return []byte{}, sdkerrors.Wrap(
			types.ErrIncorrectClaimRound,
			fmt.Sprintf("expected current round %d, to be greater than last finalized round %d", claimRoundID, lastFinalizedRound),
		)
	}

	// if no prevote we are done
	if claimParams.Prevote != true {
		return []byte{}, nil
	}

	// when using prevote claims must be submited only after the prevote round
	// claim.RoundID + VotePeriod >= currentRound
	currentRound := k.GetCurrentRound(ctx, claimType)
	if claimRoundID+claimParams.VotePeriod < currentRound {
		return []byte{}, sdkerrors.Wrap(types.ErrIncorrectClaimRound, fmt.Sprintf("expected %d, got %d", currentRound-claimParams.VotePeriod, claimRoundID))
	}

	prevoteHash := types.VoteHash(msg.Salt, claim.Hash().String(), signer)
	hasPrevote := k.HasPrevote(ctx, prevoteHash)
	if hasPrevote == false {
		return []byte{}, sdkerrors.Wrap(types.ErrNoPrevote, claim.Hash().String())
	}

	return prevoteHash, nil
}

func getValidatorAddr(ctx sdk.Context, k msgServer, signer sdk.AccAddress) sdk.ValAddress {
	// get delegator's validator
	valAddr := sdk.ValAddress(k.GetValidatorAddressFromDelegate(ctx, signer))

	// if there is no delegation it must be the validator
	if valAddr == nil {
		valAddr = sdk.ValAddress(signer)
	}

	return valAddr
}
