package cli

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetCmdIssueDenom(),
		GetCmdIssueNFT(),
		GetCmdEditNFT(),
	)
	return cmd
}

// GetCmdIssueDenom is the CLI command for an IssueDenom transaction
func GetCmdIssueDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-denom [denom-id] [denom-name] [denom-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json]",
		Short: "Issue a new denom.",
		Args:  cobra.ExactArgs(7),
		Long: strings.TrimSpace(
			fmt.Sprintf(`issue Denom.
Example:
$ %s tx %s issue-denom "ID66666" "first-denom" "shui" true true ./schema-ID66666.json --from mykey --chain-id plugchain --fees 500plug
This example creates a denom of id ID666666 and name first-denom .
[denom-id]: The name of the collection
[denom-name]: The name of the denom
[mint-restricted]: MintRestricted is true means that only Denom owners can issue NFTs under this category, false means anyone can
[edit-restricted]: EditRestricted is true means that no one in this category can edit the NFT, false means that only the owner of this NFT can edit   
[schema-content or path to schema.json]: Denom data structure definition. nft metadata (metadata) can be stored directly on the chain, or the URL of its storage source outside the chain can be stored on the chain. nft metadata is organized according to a specific [JSON Schema](https://json-schema.org/)
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsDenomID := cast.ToString(args[0])
			argsDenomName := cast.ToString(args[1])
			argsDenomSymbol := cast.ToString(args[2])

			argsMintRestricted, _ := strconv.ParseBool(args[3])
			argsEditRestricted, _ := strconv.ParseBool(args[4])

			argsSchema := cast.ToString(args[5])
			optionsContent, err := ioutil.ReadFile(argsSchema)
			if err == nil {
				argsSchema = string(optionsContent)
			}

			msg := types.NewMsgIssueDenom(argsDenomID, argsDenomName, argsSchema, clientCtx.GetFromAddress().String(), argsDenomSymbol, argsMintRestricted, argsEditRestricted)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdIssueNFT is the CLI command for an IssueNFT transaction
func GetCmdIssueNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-nft [denom-id] [nft-id] [nft-name] [nft-url] [nft-data] [nft-recipient]",
		Short: "Issue a new nft.",
		Args:  cobra.ExactArgs(6),
		Long: strings.TrimSpace(
			fmt.Sprintf(`issue NFT.
Example:
$ %s tx %s issue-nft "ID66666" "nft-666" "nftshop" "https://google.com" "./nft666-schema.json" --from mykey --chain-id plugchain --fees 500plug
This example creates a nft of id nft-666 and name nftshop .
[denom-id]: The name of the collection
[nft-id]: The id of the nft
[nft-name]: The name of nft	
[nft-url]: URI of off-chain NFT data
[nft-data]:The data of the nft data [schema.json]
[nft-recipient]: Receiver of the nft. Can be empty, when empty, the source of this value --from
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsDenomID := cast.ToString(args[0])
			argsNFTID := cast.ToString(args[1])
			argsNFTName := cast.ToString(args[2])

			argsURL := cast.ToString(args[3])

			argsSchema := cast.ToString(args[4])
			optionsContent, err := ioutil.ReadFile(argsSchema)
			if err == nil {
				argsSchema = string(optionsContent)
			}
			from := clientCtx.GetFromAddress().String()
			recipient := strings.TrimSpace(cast.ToString(args[5]))
			if len(recipient) > 0 {
				if _, err = sdk.AccAddressFromBech32(recipient); err != nil {
					return err
				}
			} else {
				recipient = from
			}
			msg := types.NewMsgIssueNFT(argsNFTID, argsDenomID, argsNFTName, argsURL, argsSchema, from, recipient)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditNFT is the CLI command for an EditNFT transaction
func GetCmdEditNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-nft [denom-id] [nft-id] [nft-name] [nft-url] [nft-data] [nft-recipient]",
		Short: "edit a  nft.",
		Args:  cobra.ExactArgs(6),
		Long: strings.TrimSpace(
			fmt.Sprintf(`edit NFT.
Example:
$ %s tx %s issue-nft "ID66666" "nft-666" "nftshop" "https://google.com" "./nft666-schema.json" --from mykey --chain-id plugchain --fees 500plug
This example creates a nft of id nft-666 and name nftshop .
[denom-id]: The name of the collection
[nft-id]: The id of the nft
[nft-name]: The name of nft	
[nft-url]: URI of off-chain NFT data
[nft-data]:The data of the nft data [schema.json]
[nft-recipient]: Receiver of the nft. Can be empty, when empty, the source of this value --from
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsDenomID := cast.ToString(args[0])
			argsNFTID := cast.ToString(args[1])
			argsNFTName := cast.ToString(args[2])

			argsURL := cast.ToString(args[3])

			argsSchema := cast.ToString(args[4])
			optionsContent, err := ioutil.ReadFile(argsSchema)
			if err == nil {
				argsSchema = string(optionsContent)
			}
			from := clientCtx.GetFromAddress().String()
			recipient := strings.TrimSpace(cast.ToString(args[5]))
			if len(recipient) > 0 {
				if _, err = sdk.AccAddressFromBech32(recipient); err != nil {
					return err
				}
			} else {
				recipient = from
			}
			msg := types.NewMsgIssueNFT(argsNFTID, argsDenomID, argsNFTName, argsURL, argsSchema, from, recipient)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
