package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	plugchaintypes "github.com/oracleNetworkProtocol/plugchain/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	localToken              Token
	Initialized             bool
	DefaultTokenDescription = "x/token module of Plug Chain Hub token."
	NativeStakingTokenDesc  = "The native staking token of the Plug Chain Hub."
)

func GetLocalToken() Token {
	if !Initialized {
		localToken = Token{
			Symbol:        plugchaintypes.DisplayNativeDenom,
			Name:          "Plug Chain Hub token",
			Scale:         plugchaintypes.BaseDenomUnit,
			MinUnit:       plugchaintypes.BaseNativeDenom,
			InitialSupply: 15989000000,
			MaxSupply:     100000000000,
			Mintable:      true,
			Owner:         sdk.AccAddress(crypto.AddressHash([]byte(ModuleName))).String(),
		}
		Initialized = true
	}
	return localToken
}

func SetLocalToken(symbol, name, minUnit string, scale uint32, initialSupply, maxSupply uint64, mintable bool, owner sdk.AccAddress) {
	localToken = NewToken(symbol, name, minUnit, scale, initialSupply, maxSupply, mintable, owner)
	Initialized = true
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
