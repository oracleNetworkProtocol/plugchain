package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyTokenTaxRate         = []byte("TokenTaxRate")
	KeyIssueTokenBaseFee    = []byte("IssueTokenBaseFee")
	KeyOperateTokenFeeRatio = []byte("OperateTokenFeeRatio")
)

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIssueTokenBaseFee, &p.IssueTokenBaseFee, validateIssueTokenBaseFee),
		paramtypes.NewParamSetPair(KeyOperateTokenFeeRatio, &p.OperateTokenFeeRatio, validateOperateTokenFeeRatio),
	}
}

func validateIssueTokenBaseFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("base fee for issuing token should not be negative")
	}
	return nil
}
func validateOperateTokenFeeRatio(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.GT(sdk.NewDec(1)) || v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fee ratio for minting tokens [%s] should be between [0, 1]", v.String())
	}
	return nil
}

// ValidateParams validates the given params
func ValidateParams(p Params) error {
	if err := validateOperateTokenFeeRatio(p.OperateTokenFeeRatio); err != nil {
		return err
	}
	if err := validateIssueTokenBaseFee(p.IssueTokenBaseFee); err != nil {
		return err
	}

	return nil
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// DefaultParams return the default params
func DefaultParams() Params {
	localToken = GetLocalToken()
	return Params{
		IssueTokenBaseFee:    sdk.NewCoin(localToken.Symbol, sdk.NewInt(50000)),
		OperateTokenFeeRatio: sdk.NewDecWithPrec(1, 1), // 0.1 (10%)
	}
}

// ParamKeyTable returns the TypeTable for the token module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}
