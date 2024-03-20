package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
)

// Message types for the oracle module
const (
	TypeMsgVote = "vote"
)

var (
	_ sdk.Msg                       = &MsgVote{}
	_ types.UnpackInterfacesMessage = MsgVote{}
	_ exported.MsgVoteI             = &MsgVote{}
)

// NewMsgVote returns a new MsgVote with a signer/submitter.
func NewMsgVote(s sdk.AccAddress, claim exported.Claim, salt string) (*MsgVote, error) {
	msg, ok := claim.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("cannot proto marshal %T", claim)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	return &MsgVote{Signer: s.String(), Claim: any, Salt: salt}, nil
}

// Route get msg route
func (msg *MsgVote) Route() string {
	return RouterKey
}

// Type get msg type
func (msg *MsgVote) Type() string {
	return TypeMsgVote
}

// GetSigners get msg signers
func (msg *MsgVote) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{accAddr}
}

// GetSignBytes get msg get getsingbytes
func (msg *MsgVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validation
func (msg *MsgVote) ValidateBasic() error {
	if msg.Signer == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	claim := msg.GetClaim()
	if claim == nil {
		return sdkerrors.Wrap(ErrInvalidClaim, "missing claim")
	}
	if err := claim.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetClaim get the claim
func (msg MsgVote) GetClaim() exported.Claim {
	claim, ok := msg.Claim.GetCachedValue().(exported.Claim)
	if !ok {
		return nil
	}
	return claim
}

// GetSigner get the submitter
func (msg MsgVote) GetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil
	}
	return accAddr
}

// MustGetSigner returns submitter
func (msg MsgVote) MustGetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return accAddr
}

// UnpackInterfaces unpack
func (msg MsgVote) UnpackInterfaces(ctx types.AnyUnpacker) error {
	var claim exported.Claim
	return ctx.UnpackAny(msg.Claim, &claim)
}
