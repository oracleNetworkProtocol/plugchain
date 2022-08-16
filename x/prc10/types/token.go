package types

import (
	math "math"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
)

var (
	_ proto.Message = &Token{}
)

const (
	MinMinUnitLen     = 2
	MaxMinUnitLen     = 8
	MinSymbolLen      = 2
	MaxSymbolLen      = 8
	MaximumMaxSupply  = math.MaxUint64
	MaximumInitSupply = uint64(100000000000)
	MaximumScale      = uint32(8)

	MaxNameLen = 32
)

type TokenI interface {
	GetOwner() sdk.AccAddress
	ReturnBaseCoin(coin sdk.Coin) (sdk.DecCoin, error)
	ReturnMinUnitCoin(coin sdk.DecCoin) (newCoin sdk.Coin, err error)
}

func NewToken(
	symbol string,
	name string,
	minUnit string,
	scale uint32,
	initialSupply,
	maxSupply uint64,
	mintable bool,
	owner sdk.AccAddress,
) Token {
	if maxSupply == 0 {
		if mintable {
			maxSupply = MaximumMaxSupply
		} else {
			maxSupply = initialSupply
		}
	}

	return Token{
		Symbol:        symbol,
		Name:          name,
		MinUnit:       minUnit,
		Scale:         scale,
		InitialSupply: initialSupply,
		MaxSupply:     maxSupply,
		Mintable:      mintable,
		Owner:         owner.String(),
	}
}

func (t Token) String() string {
	bz, _ := yaml.Marshal(t)
	return string(bz)
}

// ValidateToken checks if the given token is valid
func ValidateToken(token Token) error {
	if len(token.Owner) > 0 {
		if _, err := sdk.AccAddressFromBech32(token.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}

	if err := ValidateSymbol(token.Symbol); err != nil {
		return err
	}

	if err := ValidateMinUnit(token.MinUnit); err != nil {
		return err
	}
	if err := ValidateScale(token.Scale); err != nil {
		return err
	}

	if err := ValidateName(token.Name); err != nil {
		return err
	}

	if token.MaxSupply < token.InitialSupply {
		return sdkerrors.Wrapf(ErrInvalidMaxSupply, "invalid token max supply %d, only accepts value [%d, %d]", token.MaxSupply, token.InitialSupply, uint64(MaximumMaxSupply))
	}

	return ValidateInitialSupply(token.InitialSupply)
}

// GetOwner implements exported.TokenI
func (t Token) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(t.Owner)
	return owner
}

// ReturnBaseCoin returns the main denom coin from args
func (t Token) ReturnBaseCoin(coin sdk.Coin) (sdk.DecCoin, error) {
	if t.Symbol != coin.Denom && t.MinUnit != coin.Denom {
		return sdk.NewDecCoinFromDec(coin.Denom, sdk.ZeroDec()), sdkerrors.Wrapf(ErrTokenNotExists, "token not match")
	}

	if t.Symbol == coin.Denom {
		return sdk.NewDecCoin(coin.Denom, coin.Amount), nil
	}

	precision := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(t.Scale)), nil)
	// dest amount = src amount / 10^(scale)
	amount := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromBigInt(precision))
	return sdk.NewDecCoinFromDec(t.Symbol, amount), nil
}

// ReturnMinUnitCoin returns the min denom coin from args
func (t Token) ReturnMinUnitCoin(coin sdk.DecCoin) (newCoin sdk.Coin, err error) {
	if t.Symbol != coin.Denom && t.MinUnit != coin.Denom {
		return sdk.NewCoin(coin.Denom, sdk.ZeroInt()), sdkerrors.Wrapf(ErrTokenNotExists, "token not match")
	}

	if t.MinUnit == coin.Denom {
		return sdk.NewCoin(coin.Denom, coin.Amount.TruncateInt()), nil
	}

	precision := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(t.Scale)), nil)
	// dest amount = src amount * 10^(dest scale)
	amount := coin.Amount.Mul(sdk.NewDecFromBigInt(precision))
	return sdk.NewCoin(t.MinUnit, amount.TruncateInt()), nil
}
