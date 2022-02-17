package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	gogotypes "github.com/gogo/protobuf/types"
)

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueClass{},
		&MsgIssueNFT{},
		&MsgEditNFT{},
		&MsgBurnNFT{},
		&MsgTransferNFT{},
		&MsgTransferClass{},
	)

	registry.RegisterInterface(
		"plugchain.nft.NFTI",
		(*NFTI)(nil),
		&NFT{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.Codec, supply uint64) []byte {
	supplyByte := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyByte)
}

// return th supply
func MustUnMarshalSupply(cdc codec.Codec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

func MustMarshalNFTID(cdc codec.Codec, nftID string) []byte {
	nftIDWrap := gogotypes.StringValue{Value: nftID}
	return cdc.MustMarshal(&nftIDWrap)
}

func MustUnMarshalNFTID(cdc codec.Codec, value []byte) string {
	var nftIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &nftIDWrap)
	return nftIDWrap.Value
}
