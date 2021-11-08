package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultStringValue = "[do-not-modify]"

	MinClassLen  = 6
	MaxClassLen  = 64
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

// ValidateClassID verifies whether the  parameters are legal
func ValidateClassID(classID string) error {
	if len(classID) < MinClassLen || len(classID) > MaxClassLen {
		return sdkerrors.Wrapf(ErrInvalidClass, "the length of Class(%s) only accepts value [%d, %d]", classID, MinClassLen, MaxClassLen)
	}
	if !RegexAlphaNumeric(classID) || !RegexAlphaTop(classID) {
		return sdkerrors.Wrapf(ErrInvalidClass, "the Class(%s) only accepts alphanumeric characters, and begin with an english letter", classID)
	}
	return ValidateKeywords(classID)
}

// ValidateKeywords checks if the given classID begins with `DenomKeywords`
func ValidateKeywords(classID string) error {
	if regexpKeyword(classID) {
		return sdkerrors.Wrapf(ErrInvalidClass, "invalid classID: %s, can not begin with keyword: (%s)", classID, keyWords)
	}
	return nil
}

//ValidateNFTID verify that the nftID is legal
func ValidateNFTID(nftID string) error {
	if len(nftID) < MinClassLen || len(nftID) > MaxClassLen {
		return sdkerrors.Wrapf(ErrInvalidNFTID, "the length of nft id(%s) only accepts value [%d, %d]", nftID, MinClassLen, MaxClassLen)
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
