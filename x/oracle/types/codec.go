package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the
// evidence module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.Claim)(nil), nil)
	cdc.RegisterConcrete(&MsgVote{}, "oracle/MsgVote", nil)
	cdc.RegisterConcrete(&MsgPrevote{}, "oracle/MsgPrevote", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "oracle/MsgDelegateFeedConsent", nil)
	cdc.RegisterConcrete(&TestClaim{}, "oracle/TestClaim", nil)
}

// RegisterInterfaces registers interfaces
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgVote{},
		&MsgPrevote{},
		&MsgDelegate{},
	)
	registry.RegisterInterface(
		"relevantcommunity.oracle.oracle.Claim",
		(*exported.Claim)(nil),
		&TestClaim{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/oracle module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/oracle and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
