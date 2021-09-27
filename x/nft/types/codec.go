package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*NFTI)(nil), nil)
	cdc.RegisterConcrete(&MsgIssueDenom{}, "x/nft/issue_denom", nil)
	cdc.RegisterConcrete(&MsgIssueNFT{}, "x/nft/issue_nft", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgIssueNFT{},
	)

	registry.RegisterInterface(
		"plugchain.nft.NFTI",
		(*NFTI)(nil),
		&NFT{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
