package keeper

import (
	"math"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/prc10/types"
)

const (
	FeeBase = types.MinSymbolLen
	FeeExp  = 3
)

// DeductIssueTokenFee performs fee handling for issuing token
func (k Keeper) DeductIssueTokenFee(ctx sdk.Context, owner sdk.AccAddress, symbol string) error {
	newFee, err := k.GetIssueTokenFee(ctx, symbol)
	if err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, sdk.NewCoins(newFee)); err != nil {
		return err
	}
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(newFee))
}

// DeductOperateTokenFee performs fee handling for operate token
func (k Keeper) DeductOperateTokenFee(ctx sdk.Context, owner sdk.AccAddress, symbol string) error {
	newFee, err := k.GetOperateTokenFee(ctx, symbol)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, sdk.NewCoins(newFee)); err != nil {
		return err
	}
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(newFee))
}

//GetIssueTokenFee returns operate token fee
func (k Keeper) GetIssueTokenFee(ctx sdk.Context, symbol string) (sdk.Coin, error) {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)

	ratioFee := ratioFeeFactor(symbol)

	actualFee := sdk.NewDecFromInt(p.IssueTokenBaseFee.Amount).Quo(ratioFee)

	var (
		fee sdk.Coin
	)
	fee = sdk.NewCoin(p.IssueTokenBaseFee.Denom, sdk.OneInt())

	if actualFee.GT(sdk.NewDec(1)) {
		fee = sdk.NewCoin(p.IssueTokenBaseFee.Denom, actualFee.TruncateInt())
	}

	token, err := k.GetToken(ctx, fee.Denom)
	if err != nil {
		return sdk.Coin{}, err
	}
	newFee, er := token.ReturnMinUnitCoin(sdk.NewDecCoinFromCoin(fee))
	if er != nil {
		return sdk.Coin{}, er
	}

	return newFee, nil
}

//GetOperateTokenFee returns operate token fee
func (k Keeper) GetOperateTokenFee(ctx sdk.Context, symbol string) (sdk.Coin, error) {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)

	ratioFee := ratioFeeFactor(symbol)

	actualFee := sdk.NewDecFromInt(p.IssueTokenBaseFee.Amount).Quo(ratioFee).Mul(p.OperateTokenFeeRatio)

	var (
		fee sdk.Coin
	)
	fee = sdk.NewCoin(p.IssueTokenBaseFee.Denom, sdk.OneInt())

	if actualFee.GT(sdk.NewDec(1)) {
		fee = sdk.NewCoin(p.IssueTokenBaseFee.Denom, actualFee.TruncateInt())
	}

	token, err := k.GetToken(ctx, fee.Denom)
	if err != nil {
		return sdk.Coin{}, err
	}
	newFee, er := token.ReturnMinUnitCoin(sdk.NewDecCoinFromCoin(fee))
	if er != nil {
		return sdk.Coin{}, er
	}
	return newFee, nil
}

//ratioFeeFactor computes the fee factor of the given name
func ratioFeeFactor(symbol string) sdk.Dec {
	sLen := len(symbol)
	if sLen == 0 {
		panic("the length of name must be greater than 0")
	}
	divi := math.Log(FeeBase)
	divid := math.Log(float64(sLen))

	busPower := math.Pow(divid/divi, FeeExp)
	feeDec, err := sdk.NewDecFromStr(strconv.FormatFloat(busPower, 'f', 2, 64))
	if err != nil {
		panic("invalid string")
	}
	return feeDec
}
