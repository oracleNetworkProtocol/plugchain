package token

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/token/keeper"
	"github.com/oracleNetworkProtocol/plugchain/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err.Error())
	}
	k.SetParamSet(ctx, genState.Params)

	// init token
	for _, token := range genState.Tokens {
		if err := k.AddToken(ctx, token); err != nil {
			panic(err)
		}

	}
	// assert the symbol exists
	if !k.HasToken(ctx, genState.Params.IssueTokenBaseFee.Denom) {

		panic(fmt.Sprintf("Token %s does not exist", genState.Params.IssueTokenBaseFee.Denom))
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var tokens []types.Token
	for _, v := range k.GetTokens(ctx, nil) {
		t := v.(types.Token)
		tokens = append(tokens, t)
	}
	return &types.GenesisState{
		Params: k.GetParamSet(ctx),
		Tokens: tokens,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *types.GenesisState {
	return &types.GenesisState{
		Params: types.DefaultParams(),
		Tokens: []types.Token{types.GetLocalToken()},
	}
}
