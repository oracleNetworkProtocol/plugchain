package types

import (
	fmt "fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	keywords          = strings.Join([]string{"ibc"}, "|")
	regexpKeywordsFmt = fmt.Sprintf("^(%s).*", keywords)
	regexpKeyword     = regexp.MustCompile(regexpKeywordsFmt).MatchString

	regexpSymbolFmt = fmt.Sprintf("^[a-z][a-z0-9]{%d,%d}$", MinSymbolLen-1, MaxSymbolLen-1)
	regexpSymbol    = regexp.MustCompile(regexpSymbolFmt).MatchString

	regexpMinUintFmt = fmt.Sprintf("^[a-z][a-z0-9]{%d,%d}$", MinMinUnitLen-1, MaxMinUnitLen-1)
	regexpMinUint    = regexp.MustCompile(regexpMinUintFmt).MatchString
)

// ValidateAmount
func ValidateAmount(amount uint64) error {
	if amount == 0 {
		return sdkerrors.Wrapf(ErrInvalidMaxSupply, "invalid token amount %d, only accepts value (0, %d]", amount, uint64(MaximumMaxSupply))
	}
	return nil
}

// ValidateInitialSupply verifies whether the initial supply is legal
func ValidateInitialSupply(initialSupply uint64) error {
	if initialSupply > MaximumInitSupply {
		return sdkerrors.Wrapf(ErrInvalidInitSupply, "invalid token initial supply %d, only accepts value [0, %d]", initialSupply, MaximumInitSupply)
	}
	return nil
}

// ValidateSymbol checks if the given symbol is valid
func ValidateSymbol(symbol string) error {
	if !regexpSymbol(symbol) {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid symbol: %s, only accepts english lowercase letters and numbers, length [%d, %d], and begin with an english letter, regexp: %s", symbol, MinSymbolLen, MaxSymbolLen, regexpSymbolFmt)
	}
	return ValidateKeywords(symbol)
}

// ValidateKeywords checks if the given denom begins with `TokenKeywords`
func ValidateKeywords(denom string) error {
	if regexpKeyword(denom) {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid token: %s, can not begin with keyword: (%s)", denom, keywords)
	}
	return nil
}

// ValidateMinUnit checks if the given min unit is valid
func ValidateMinUnit(minUnit string) error {
	if !regexpMinUint(minUnit) {
		return sdkerrors.Wrapf(ErrInvalidMinUnit, "invalid minUnit: %s, only accepts english lowercase letters and numbers, length [%d, %d], and begin with an english letter, regexp: %s", minUnit, MinMinUnitLen, MaxMinUnitLen, regexpMinUintFmt)
	}
	return ValidateKeywords(minUnit)
}

// ValidateName verifies whether the given name is legal
func ValidateName(name string) error {
	if len(name) == 0 || len(name) > MaxNameLen {
		return sdkerrors.Wrapf(ErrInvalidName, "invalid token name %s, only accepts length (0, %d]", name, MaxNameLen)
	}
	return nil
}

// ValidateScale verifies whether the given scale is legal
func ValidateScale(scale uint32) error {
	if scale > MaximumScale {
		return sdkerrors.Wrapf(ErrInvalidScale, "invalid token scale %d, only accepts value [0, %d]", scale, MaximumScale)
	}
	return nil
}
