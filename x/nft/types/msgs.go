package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgIssueClass    = "issue_class"
	TypeMsgIssueNFT      = "issue_nft"
	TypeMsgEditNFT       = "edit_nft"
	TypeMsgBurnNFT       = "burn_nft"
	TypeMsgTransferNFT   = "transfer_nft"
	TypeMsgTransferClass = "transfer_class"
)

var (
	_ sdk.Msg = &MsgIssueClass{}
	_ sdk.Msg = &MsgIssueNFT{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgTransferClass{}
)

// NewMsgIssueClass is a constructor function for MsgSetName
func NewMsgIssueClass(classID, className, schema, sender, symbol string, mintRestricted, editRestricted bool) *MsgIssueClass {
	return &MsgIssueClass{
		ID:             classID,
		Name:           className,
		Schema:         schema,
		Symbol:         symbol,
		Owner:          sender,
		MintRestricted: mintRestricted,
		EditRestricted: editRestricted,
	}
}

func (mid MsgIssueClass) Route() string { return RouterKey }
func (mid MsgIssueClass) Type() string  { return TypeMsgIssueClass }
func (mid MsgIssueClass) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(mid.Owner); err != nil {
		return err
	}

	if err := ValidateClassID(mid.ID); err != nil {
		return err
	}

	return nil
}

func (mid MsgIssueClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&mid)
	return sdk.MustSortJSON(bz)
}

func (mid MsgIssueClass) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mid.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgIssueNFT is a constructor function for MsgIssueNFT
func NewMsgIssueNFT(nftID, classID, className, uri, schema, owner, recipient string) *MsgIssueNFT {
	return &MsgIssueNFT{
		ID:        nftID,
		ClassID:   classID,
		URI:       uri,
		Data:      schema,
		Name:      className,
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

	if err := ValidateClassID(min.ClassID); err != nil {
		return err
	}

	if err := ValidateNFTID(min.ID); err != nil {
		return err
	}

	return ValidateNFTURI(min.URI)
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
func NewMsgEditNFT(nftID, classID, className, uri, schema, owner string) *MsgEditNFT {
	return &MsgEditNFT{
		ID:      nftID,
		ClassID: classID,
		URI:     uri,
		Data:    schema,
		Name:    className,
		Owner:   owner,
	}
}

func (med MsgEditNFT) Route() string { return RouterKey }
func (med MsgEditNFT) Type() string  { return TypeMsgEditNFT }
func (med MsgEditNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(med.Owner); err != nil {
		return err
	}
	if err := ValidateClassID(med.ClassID); err != nil {
		return err
	}

	if err := ValidateNFTID(med.ID); err != nil {
		return err
	}

	return ValidateNFTURI(med.URI)
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
func NewMsgBurnNFT(nftID, classID, owner string) *MsgBurnNFT {
	return &MsgBurnNFT{
		ID:      nftID,
		ClassID: classID,
		Owner:   owner,
	}
}

func (mbn MsgBurnNFT) Route() string { return RouterKey }
func (mbn MsgBurnNFT) Type() string  { return TypeMsgBurnNFT }
func (mbn MsgBurnNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(mbn.Owner); err != nil {
		return err
	}
	if err := ValidateClassID(mbn.ClassID); err != nil {
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

// NewMsgTransferNFT is a constructor function for MsgTransferNFT
func NewMsgTransferNFT(nftID, classID, owner, recipient string) *MsgTransferNFT {
	return &MsgTransferNFT{
		ID:        nftID,
		ClassID:   classID,
		Recipient: recipient,
		Owner:     owner,
	}
}

func (mtn MsgTransferNFT) Route() string { return RouterKey }
func (mtn MsgTransferNFT) Type() string  { return TypeMsgTransferNFT }
func (mtn MsgTransferNFT) ValidateBasic() error {
	if err := ValidateClassID(mtn.ClassID); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(mtn.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(mtn.Owner); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return ValidateNFTID(mtn.ID)
}

func (mtn MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&mtn)
	return sdk.MustSortJSON(bz)
}

func (mtn MsgTransferNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mtn.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgTransferClass is a constructor function for MsgTransferClass
func NewMsgTransferClass(ID, owner, recipient string) *MsgTransferClass {
	return &MsgTransferClass{
		ID:        ID,
		Recipient: recipient,
		Owner:     owner,
	}
}

func (mtd MsgTransferClass) Route() string { return RouterKey }
func (mtd MsgTransferClass) Type() string  { return TypeMsgTransferClass }
func (mtd MsgTransferClass) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(mtd.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(mtd.Owner); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return ValidateClassID(mtd.ID)
}

func (mtd MsgTransferClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&mtd)
	return sdk.MustSortJSON(bz)
}

func (mtd MsgTransferClass) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mtd.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
