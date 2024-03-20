package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelegate{}

// Message types for the oracle module
const (
	TypeMsgDelegate = "delegate"
)

// NewMsgDelegate returns a new MsgDelegateFeedConsent
func NewMsgDelegate(val, del sdk.AccAddress) *MsgDelegate {
	return &MsgDelegate{
		Validator: val.String(),
		Delegate:  del.String(),
	}
}

// Route implements sdk.Msg
func (m *MsgDelegate) Route() string { return ModuleName }

// Type implements sdk.Msg
func (m *MsgDelegate) Type() string { return TypeMsgDelegate }

// ValidateBasic implements sdk.Msg
func (m *MsgDelegate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Validator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Delegate); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	return nil
}

// GetSignBytes implements sdk.Msg
func (m *MsgDelegate) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners implements sdk.Msg
func (m *MsgDelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.MustGetValidator()}
}

// MustGetValidator returns the sdk.AccAddress for the validator
func (m *MsgDelegate) MustGetValidator() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(m.Validator)
	if err != nil {
		panic(err)
	}
	return val
}

// MustGetDelegate returns the sdk.AccAddress for the delegate
func (m *MsgDelegate) MustGetDelegate() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(m.Delegate)
	if err != nil {
		panic(err)
	}
	return val
}
