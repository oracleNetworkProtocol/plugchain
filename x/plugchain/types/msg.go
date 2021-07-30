package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateToken{}
var _ sdk.Msg = &MsgBurnToken{}

const (
	TypeMsgCreateToken = "create_token"
	TypeMsgBurnToken   = "burn"
)

func NewMsgCreateToken(owner, wholeName, originalSymbol, description, symbol string, totalSupply *sdk.Int, decimal uint64, mintable bool) *MsgCreateToken {
	return &MsgCreateToken{
		Symbol:         symbol,
		OriginalSymbol: originalSymbol,
		Description:    description,
		WholeName:      wholeName,
		TotalSupply:    totalSupply,
		Owner:          owner,
		Founder:        owner,
		Decimal:        decimal,
		Mintable:       mintable,
	}
}

func (msg *MsgCreateToken) Route() string {
	return RouterKey
}

func (msg *MsgCreateToken) Type() string {
	return TypeMsgCreateToken
}

func (msg *MsgCreateToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg MsgCreateToken) GetTokenCreator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return addr
}

func NewMsgBurnToken(symbol, address string, account *sdk.Int) *MsgBurnToken {
	return &MsgBurnToken{
		Symbol:  symbol,
		Address: address,
		Account: account,
	}
}

func (burn *MsgBurnToken) Route() string {
	return RouterKey
}

func (burn *MsgBurnToken) Type() string {
	return TypeMsgBurnToken
}

func (burn *MsgBurnToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(burn.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (burn *MsgBurnToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(burn)
	return sdk.MustSortJSON(bz)
}

func (burn *MsgBurnToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(burn.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
