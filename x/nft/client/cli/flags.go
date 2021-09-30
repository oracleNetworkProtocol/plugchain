package cli

import (
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
	"github.com/spf13/pflag"
)

const (
	FlagNFTName = "nft-name"
	FlagNFTURL  = "nft-url"
	FlagNFTData = "nft-data"
)

var (
	FsEditNFT = pflag.NewFlagSet("", pflag.ContinueOnError)
)

func init() {
	FsEditNFT.String(FlagNFTName, types.DefaultStringValue, "The name of the nft")
	FsEditNFT.String(FlagNFTURL, types.DefaultStringValue, "URL for the supplemental off-chain nft data (should return a JSON object)")
	FsEditNFT.String(FlagNFTData, types.DefaultStringValue, "The nft data of the nft")
}
