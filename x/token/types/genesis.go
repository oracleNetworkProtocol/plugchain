package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var localToken = Token{
	Symbol:        sdk.DefaultBondDenom,
	Name:          "Plug Chain Hub",
	Scale:         0,
	MinUnit:       sdk.DefaultBondDenom,
	InitialSupply: MaximumInitSupply,
	MaxSupply:     MaximumMaxSupply,
	Mintable:      true,
	Owner:         sdk.AccAddress(crypto.AddressHash([]byte(ModuleName))).String(),
}

func GetLocalToken() Token {
	return localToken
}

func SetLocalToken(symbol, name, minUnit string, scale uint32, initialSupply, maxSupply uint64, mintable bool, owner sdk.AccAddress) {
	localToken = NewToken(symbol, name, minUnit, scale, initialSupply, maxSupply, mintable, owner)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := ValidateParams(gs.Params); err != nil {
		return err
	}

	//validate token
	for _, token := range gs.Tokens {
		if err := ValidateToken(token); err != nil {
			return err
		}
	}
	return nil
}

func NewGenesisState(p Params, ts []Token) GenesisState {
	return GenesisState{
		Params: p,
		Tokens: ts,
	}
}
