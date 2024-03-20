package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	exportedOracle "github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"relevantcommunity.oracle.oracle.Claim",
		(*exportedOracle.Claim)(nil),
		&AtomUsd{},
	)
}
