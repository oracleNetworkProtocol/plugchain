package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultStringValue = "[do-not-modify]"

	MinDenomLen  = 6
	MaxDenomLen  = 64
	MaxNFTURLLen = 256
)

var (
	RegexAlphaNumeric = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	RegexAlphaTop     = regexp.MustCompile(`^[a-z].*`).MatchString

	keyWords         = strings.Join([]string{"ibc", "plug"}, "|")
	regexpKeywordFmt = fmt.Sprintf("^(%s).*", keyWords)
	regexpKeyword    = regexp.MustCompile(regexpKeywordFmt).MatchString

	URLMatchWords = strings.Join([]string{"http://", "https://"}, "|")
	regexURLFmt   = fmt.Sprintf("^(%s).*", URLMatchWords)
	regexpURL     = regexp.MustCompile(regexURLFmt).MatchString
)

// ValidateDenomID verifies whether the  parameters are legal
func ValidateDenomID(denomID string) error {
	if len(denomID) < MinDenomLen || len(denomID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidDenom, "the length of denom(%s) only accepts value [%d, %d]", denomID, MinDenomLen, MaxDenomLen)
	}
	if !RegexAlphaNumeric(denomID) || !RegexAlphaTop(denomID) {
		return sdkerrors.Wrapf(ErrInvalidDenom, "the denom(%s) only accepts alphanumeric characters, and begin with an english letter", denomID)
	}
	return ValidateKeywords(denomID)
}

// ValidateKeywords checks if the given denomId begins with `DenomKeywords`
func ValidateKeywords(denomId string) error {
	if regexpKeyword(denomId) {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denomId: %s, can not begin with keyword: (%s)", denomId, keyWords)
	}
	return nil
}

//ValidateNFTID verify that the nftID is legal
func ValidateNFTID(nftID string) error {
	if len(nftID) < MinDenomLen || len(nftID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidNFTID, "the length of nft id(%s) only accepts value [%d, %d]", nftID, MinDenomLen, MaxDenomLen)
	}
	if !RegexAlphaNumeric(nftID) || !RegexAlphaTop(nftID) {
		return sdkerrors.Wrapf(ErrInvalidNFTID, "nft id(%s) only accepts alphanumeric characters, and begin with an english letter", nftID)
	}
	return nil
}

func ValidateNFTURL(url string) error {
	if len(url) > MaxNFTURLLen {
		return sdkerrors.Wrapf(ErrInvalidNFTURL, "the length of nft url(%s) only accepts value [0, %d]", url, MaxNFTURLLen)
	}
	if !regexpURL(url) {
		return sdkerrors.Wrapf(ErrInvalidNFTURL, "url begin with: (%s) ", URLMatchWords)
	}
	return nil
}

// Modified returns whether the field is modified
func Modified(item string) bool {
	return item != DefaultStringValue
}
