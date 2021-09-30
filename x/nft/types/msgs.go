package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssueDenom = "issue_denom"
	TypeMsgIssueNFT   = "issue_nft"
	TypeMsgEditNFT    = "edit_nft"
	TypeMsgBurnNFT    = "burn_nft"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
	_ sdk.Msg = &MsgIssueNFT{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
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

// NewMsgEditNFT is a constructor function for MsgEditNFT
func NewMsgEditNFT(nftID, denomID, denomName, url, schema, owner string) *MsgEditNFT {
	return &MsgEditNFT{
		ID:      nftID,
		DenomID: denomID,
		URL:     url,
		Data:    schema,
		Name:    denomName,
		Owner:   owner,
	}
}

func (med MsgEditNFT) Route() string { return RouterKey }
func (med MsgEditNFT) Type() string  { return TypeMsgEditNFT }
func (med MsgEditNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(med.Owner); err != nil {
		return err
	}
	if err := ValidateDenomID(med.DenomID); err != nil {
		return err
	}

	if err := ValidateNFTID(med.ID); err != nil {
		return err
	}

	return ValidateNFTURL(med.URL)
}

func (med MsgEditNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&med)
	return sdk.MustSortJSON(bz)
}

func (med MsgEditNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(med.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(nftID, denomID, owner string) *MsgBurnNFT {
	return &MsgBurnNFT{
		ID:      nftID,
		DenomID: denomID,
		Owner:   owner,
	}
}

func (mbn MsgBurnNFT) Route() string { return RouterKey }
func (mbn MsgBurnNFT) Type() string  { return TypeMsgBurnNFT }
func (mbn MsgBurnNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(mbn.Owner); err != nil {
		return err
	}
	if err := ValidateDenomID(mbn.DenomID); err != nil {
		return err
	}

	return ValidateNFTID(mbn.ID)
}

func (mbn MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&mbn)
	return sdk.MustSortJSON(bz)
}

func (mbn MsgBurnNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mbn.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
