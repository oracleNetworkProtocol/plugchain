package cli

import (
	"github.com/spf13/pflag"
)

const (
	DefaultStringValue = "[do-not-modify]"

	FlagNFTName = "nft-name"
	FlagNFTURL  = "nft-url"
	FlagNFTData = "nft-data"
)

var (
	FsEditNFT = pflag.NewFlagSet("", pflag.ContinueOnError)
)

func init() {
	FsEditNFT.String(FlagNFTName, DefaultStringValue, "The name of the nft")
	FsEditNFT.String(FlagNFTURL, DefaultStringValue, "URL for the supplemental off-chain nft data (should return a JSON object)")
	FsEditNFT.String(FlagNFTData, DefaultStringValue, "The nft data of the nft")
}
