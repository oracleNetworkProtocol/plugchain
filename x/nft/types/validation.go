package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MinDenomLen = 6
	MaxDenomLen = 64
)

var (
	RegexAlphaNumeric = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	RegexAlphaTop     = regexp.MustCompile(`^[a-z].*`).MatchString

	keyWords         = strings.Join([]string{"ibc", "plug"}, "|")
	regexpKeywordFmt = fmt.Sprintf("^(%s).*", keyWords)
	regexpKeyword    = regexp.MustCompile(regexpKeywordFmt).MatchString
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
