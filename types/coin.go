package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// BaseNativeDenom defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the PoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	BaseNativeDenom string = "uplugcn"

	//DisplayNativeDenom defines the denomination displayed to users in client applications.
	DisplayNativeDenom string = "pc"

	// BaseDenomUnit defines the base denomination unit for plugcns.
	// 1 pc = 1x10^{BaseDenomUnit} uplugcn
	BaseDenomUnit = 6

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
// var PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewSocketCoin is a utility function that returns an "uplugcn" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewSocketCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(BaseNativeDenom, amount)
}

// NewSocketDecCoin is a utility function that returns an "uplugcn" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewSocketDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(BaseNativeDenom, amount)
}

// NewSocketCoinInt64 is a utility function that returns an "uplugcn" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewSocketCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(BaseNativeDenom, amount)
}
