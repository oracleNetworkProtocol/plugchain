package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultStringValue = "[do-not-modify]"

	MaxNFTURILen = 256
)

var (
	// reClassIDString can be 3 ~ 100 characters long and support letters, followed by either
	// a letter, a number or a slash ('/') or a colon (':') or ('-').
	reClassIDString = `[a-zA-Z][a-zA-Z0-9/:-]{2,100}`
	reClassID       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reClassIDString))
	// reNFTIDString can be 3 ~ 100 characters long and support letters, followed by either
	// a letter, a number or a slash ('/') or a colon (':') or ('-').
	reNFTID = reClassID

	URIMatchWords = strings.Join([]string{"http://", "https://"}, "|")
	regexURIFmt   = fmt.Sprintf("^(%s).*", URIMatchWords)
	regexpURI     = regexp.MustCompile(regexURIFmt).MatchString
)

// ValidateClassID verifies whether the  parameters are legal
func ValidateClassID(id string) error {
	if !reClassID.MatchString(id) {
		return sdkerrors.Wrapf(ErrInvalidClass, "invalid class id: %s", id)
	}
	return nil
}

// ValidateNFTID returns whether the nft id is valid
func ValidateNFTID(id string) error {
	if !reNFTID.MatchString(id) {
		return sdkerrors.Wrapf(ErrInvalidNFTID, "invalid nft id: %s", id)
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
