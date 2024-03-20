package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message types for the oracle module
const (
	TypeMsgPrevote = "prevote"
)

var (
	_ sdk.Msg = &MsgPrevote{}
)

// NewMsgPrevote returns a new MsgPrevotePrevote with a signer.
func NewMsgPrevote(s sdk.AccAddress, hash []byte) *MsgPrevote {
	return &MsgPrevote{Signer: s.String(), Hash: hash}
}

// Route get msg route
func (msg *MsgPrevote) Route() string {
	return RouterKey
}

// Type get msg type
func (msg *MsgPrevote) Type() string {
	return TypeMsgPrevote
}

// GetSigners implements sdk.Msg
func (msg *MsgPrevote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MustGetSigner()}
}

// GetSignBytes get msg get getsingbytes
func (msg *MsgPrevote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validation
func (msg *MsgPrevote) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.Hash) == 0 {
		return fmt.Errorf("prevote hash cannot be empty")
	}

	return nil
}

// MustGetSigner returns submitter
func (msg MsgPrevote) MustGetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return accAddr
}
