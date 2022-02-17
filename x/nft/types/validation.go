package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultStringValue = "[do-not-modify]"

	MinClassLen  = 3
	MaxClassLen  = 64
	MaxNFTURILen = 256
)

var (
	RegexAlphaNumeric = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	RegexAlphaTop     = regexp.MustCompile(`^[a-z].*`).MatchString

	keyWords         = strings.Join([]string{"ibc", "plug"}, "|")
	regexpKeywordFmt = fmt.Sprintf("^(%s).*", keyWords)
	regexpKeyword    = regexp.MustCompile(regexpKeywordFmt).MatchString

	URIMatchWords = strings.Join([]string{"http://", "https://"}, "|")
	regexURIFmt   = fmt.Sprintf("^(%s).*", URIMatchWords)
	regexpURI     = regexp.MustCompile(regexURIFmt).MatchString
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

func ValidateNFTURI(uri string) error {
	if len(uri) > MaxNFTURILen {
		return sdkerrors.Wrapf(ErrInvalidNFTURI, "the length of nft uri(%s) only accepts value [0, %d]", uri, MaxNFTURILen)
	}
	if !regexpURI(uri) {
		return sdkerrors.Wrapf(ErrInvalidNFTURI, "uri begin with: (%s) ", URIMatchWords)
	}
	return nil
}

// Modified returns whether the field is modified
func Modified(item string) bool {
	return item != DefaultStringValue
}
