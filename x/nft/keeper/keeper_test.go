package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	// "github.com/cosmos/cosmos-sdk/codec"
	// codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	// "github.com/cosmos/cosmos-sdk/store"
	// "github.com/oracleNetworkProtocol/plugchain/x/nft/types"
	// "github.com/stretchr/testify/require"
	// "github.com/tendermint/tendermint/libs/log"
	// tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	// tmdb "github.com/tendermint/tm-db"
)

type KeeperSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper Keeper
	app    *simapp.SimApp
}

func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
	// storeKey := sdk.NewKVStoreKey(types.StoreKey)

	// db := tmdb.NewMemDB()
	// stateStore := store.NewCommitMultiStore(db)
	// stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	// require.NoError(t, stateStore.LoadLatestVersion())

	// registry := codectypes.NewInterfaceRegistry()
	// keeper := NewKeeper(codec.NewProtoCodec(registry), storeKey,)

	// ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	// return keeper, ctx
	return nil, sdk.Context{}
}
