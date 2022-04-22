package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgEditToken          = "edit_token"
	TypeMsgIssueToken         = "issue_token"
	TypeMsgMintToken          = "mint_token"
	TypeMsgBurnToken          = "burn_token"
	TypeMsgTransferOwnerToken = "change_owner_token"
	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
)

var (
	_ sdk.Msg = &MsgIssueToken{}
	_ sdk.Msg = &MsgMintToken{}
	_ sdk.Msg = &MsgEditToken{}
	_ sdk.Msg = &MsgBurnToken{}
	_ sdk.Msg = &MsgTransferOwnerToken{}
)

// NewMsgIssueToken - construct token issue msg.
func NewMsgIssueToken(
	symbol string, minUnit string, name string,
	scale uint32, initialSupply, maxSupply uint64,
	mintable bool, owner string,
) *MsgIssueToken {
	return &MsgIssueToken{
		Symbol:        symbol,
		Name:          name,
		Scale:         scale,
		MinUnit:       minUnit,
		InitialSupply: initialSupply,
		MaxSupply:     maxSupply,
		Mintable:      mintable,
		Owner:         owner,
	}
}

// Route Implements Msg.
func (msg MsgIssueToken) Route() string { return ModuleName }

// Type Implements Msg.
func (msg MsgIssueToken) Type() string { return TypeMsgIssueToken }

// ValidateBasic Implements Msg.
func (msg MsgIssueToken) ValidateBasic() error {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	return ValidateToken(
		NewToken(
			msg.Symbol,
			msg.Name,
			msg.MinUnit,
			msg.Scale,
			msg.InitialSupply,
			msg.MaxSupply,
			msg.Mintable,
			owner,
		),
	)
}

// GetSignBytes Implements Msg.
func (msg MsgIssueToken) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgIssueToken) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgIssueToken - construct token issue msg.
func NewMsgMintToken(symbol string, amount uint64, to, owner string) *MsgMintToken {

	return &MsgMintToken{
		Symbol: symbol,
		Amount: amount,
		To:     to,
		Owner:  owner,
	}
}

// Route Implements Msg.
func (msg MsgMintToken) Route() string { return ModuleName }

// Type Implements Msg.
func (msg MsgMintToken) Type() string { return TypeMsgMintToken }

// ValidateBasic Implements Msg.
func (msg MsgMintToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintToken) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgEditToken
func NewMsgEditToken(symbol, name string, maxSupply uint64, owner string) *MsgEditToken {

	return &MsgEditToken{
		Symbol:    symbol,
		Name:      name,
		MaxSupply: maxSupply,
		Owner:     owner,
	}
}

// Route Implements Msg.
func (msg MsgEditToken) Route() string { return ModuleName }

// Type Implements Msg.
func (msg MsgEditToken) Type() string { return TypeMsgEditToken }

// ValidateBasic Implements Msg.
func (msg MsgEditToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if err := ValidateAmount(msg.MaxSupply); err != nil {
		return err
	}
	return ValidateSymbol(msg.Symbol)
}

// GetSignBytes Implements Msg.
func (msg MsgEditToken) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgEditToken) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBurnToken
func NewMsgBurnToken(symbol string, amount uint64, owner string) *MsgBurnToken {

	return &MsgBurnToken{
		Symbol: symbol,
		Amount: amount,
		Owner:  owner,
	}
}

// Route Implements Msg.
func (msg MsgBurnToken) Route() string { return ModuleName }

// Type Implements Msg.
func (msg MsgBurnToken) Type() string { return TypeMsgBurnToken }

// ValidateBasic Implements Msg.
func (msg MsgBurnToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if err := ValidateAmount(msg.Amount); err != nil {
		return err
	}
	return ValidateSymbol(msg.Symbol)
}

// GetSignBytes Implements Msg.
func (msg MsgBurnToken) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgBurnToken) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgTransferOwnerToken
func NewMsgTransferOwnerToken(symbol, src, dst string) *MsgTransferOwnerToken {

	return &MsgTransferOwnerToken{
		Symbol: symbol,
		Owner:  src,
		To:     dst,
	}
}

// Route Implements Msg.
func (msg MsgTransferOwnerToken) Route() string { return ModuleName }

// Type Implements Msg.
func (msg MsgTransferOwnerToken) Type() string { return TypeMsgTransferOwnerToken }

// ValidateBasic Implements Msg.
func (msg MsgTransferOwnerToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	return ValidateSymbol(msg.Symbol)
}

// GetSignBytes Implements Msg.
func (msg MsgTransferOwnerToken) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgTransferOwnerToken) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
