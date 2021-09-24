package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssueDenom = "issue_denom"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
)

// NewMsgIssueDenom is a constructor function for MsgSetName
func NewMsgIssueDenom(denomID, denomName, schema, sender, symbol string, mintRestricted, editRestricted bool) *MsgIssueDenom {
	return &MsgIssueDenom{
		ID:             denomID,
		Name:           denomName,
		Schema:         schema,
		Symbol:         symbol,
		Owner:          sender,
		MintRestricted: mintRestricted,
		EditRestricted: editRestricted,
	}
}

func (mid MsgIssueDenom) Route() string { return RouterKey }
func (mid MsgIssueDenom) Type() string  { return TypeMsgIssueDenom }
func (mid MsgIssueDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(mid.Owner); err != nil {
		return err
	}

	if err := ValidateDenomID(mid.ID); err != nil {
		return err
	}

	return nil
}

func (mid MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&mid)
	return sdk.MustSortJSON(bz)
}

func (mid MsgIssueDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mid.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
