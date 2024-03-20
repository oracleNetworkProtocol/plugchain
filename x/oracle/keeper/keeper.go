package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
)

type (
	// Keeper of the oracle store
	Keeper struct {
		cdc           codec.Codec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		StakingKeeper types.StakingKeeper
		paramspace    types.ParamSubspace
	}
)

// NewKeeper instatiates the oracle keeper
func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey, stakingKeeper types.StakingKeeper, paramspace types.ParamSubspace) *Keeper {

	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		StakingKeeper: stakingKeeper,
		paramspace:    paramspace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
