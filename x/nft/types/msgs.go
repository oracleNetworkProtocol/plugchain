package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssueDenom = "issue_denom"
	TypeMsgIssueNFT   = "issue_nft"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
	_ sdk.Msg = &MsgIssueNFT{}
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

// NewMsgIssueNFT is a constructor function for MsgIssueNFT
func NewMsgIssueNFT(nftID, denomID, denomName, url, schema, owner, recipient string) *MsgIssueNFT {
	return &MsgIssueNFT{
		ID:        nftID,
		DenomID:   denomID,
		URL:       url,
		Data:      schema,
		Name:      denomName,
		Owner:     owner,
		Recipient: recipient,
	}
}

func (min MsgIssueNFT) Route() string { return RouterKey }
func (min MsgIssueNFT) Type() string  { return TypeMsgIssueNFT }
func (min MsgIssueNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(min.Owner); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(min.Recipient); err != nil {
		return err
	}

	if err := ValidateDenomID(min.DenomID); err != nil {
		return err
	}

	if err := ValidateNFTID(min.ID); err != nil {
		return err
	}

	return ValidateNFTURL(min.URL)
}

func (min MsgIssueNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&min)
	return sdk.MustSortJSON(bz)
}

func (min MsgIssueNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(min.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
