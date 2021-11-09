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
		GetCmdIssueClass(),
		GetCmdIssueNFT(),
		GetCmdEditNFT(),
		GetCmdBurnNFT(),
		GetCmdTransferNFT(),
		GetCmdTransferClass(),
	)
	return cmd
}

// GetCmdIssueClass is the CLI command for an IssueClass transaction
func GetCmdIssueClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-class [class-id] [class-name] [class-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json]",
		Short: "Issue a new class.",
		Args:  cobra.ExactArgs(6),
		Long: strings.TrimSpace(
			fmt.Sprintf(`issue class.
Example:
$ %s tx %s issue-class "ID66666" "first-class" "shui" true true ./schema-ID66666.json --from mykey --chain-id plugchain --fees 500plug
This example creates a class of id ID666666 and name first-class .
[class-id]: The name of the collection
[class-name]: The name of the class
[mint-restricted]: MintRestricted is true means that only class owners can issue NFTs under this category, false means anyone can
[edit-restricted]: EditRestricted is true means that no one in this category can edit the NFT, false means that only the owner of this NFT can edit   
[schema-content or path to schema.json]: Class data structure definition. nft metadata (metadata) can be stored directly on the chain, or the URI of its storage source outside the chain can be stored on the chain. nft metadata is organized according to a specific [JSON Schema](https://json-schema.org/)
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsClassID := cast.ToString(args[0])
			argsClassName := cast.ToString(args[1])
			argsClassSymbol := cast.ToString(args[2])

			argsMintRestricted, _ := strconv.ParseBool(args[3])
			argsEditRestricted, _ := strconv.ParseBool(args[4])

			argsSchema := cast.ToString(args[5])
			optionsContent, err := ioutil.ReadFile(argsSchema)
			if err == nil {
				argsSchema = string(optionsContent)
			}

			msg := types.NewMsgIssueClass(argsClassID, argsClassName, argsSchema, clientCtx.GetFromAddress().String(), argsClassSymbol, argsMintRestricted, argsEditRestricted)
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
		Use:   "issue-nft [class-id] [nft-id] [nft-name] [nft-uri] [nft-data] [nft-recipient]",
		Short: "Issue a new nft.",
		Args:  cobra.ExactArgs(6),
		Long: strings.TrimSpace(
			fmt.Sprintf(`issue NFT.
Example:
$ %s tx %s issue-nft "ID66666" "nft-666" "nftshop" "https://google.com" "./nft666-schema.json" "" --from mykey --chain-id plugchain --fees 500plug
This example creates a nft of id nft-666 and name nftshop .
[class-id]: The name of the collection
[nft-id]: The id of the nft
[nft-name]: The name of nft	
[nft-uri]: URI of off-chain NFT data
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

			argsClassID := cast.ToString(args[0])
			argsNFTID := cast.ToString(args[1])
			argsNFTName := cast.ToString(args[2])

			argsURI := cast.ToString(args[3])

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
			msg := types.NewMsgIssueNFT(argsNFTID, argsClassID, argsNFTName, argsURI, argsSchema, from, recipient)
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
		Use:   "edit-nft [class-id] [nft-id]",
		Short: "edit a nft.",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`edit NFT.
Example:
$ %s tx %s edit-nft "ID66666" "nft-666" --nft-name="nftshop" --nft-uri="https://google.com/" --nft-data="./nft666-schema.json" --from=mykey --chain-id=plugchain --fees=500plug
This example edit a nft of id nft-666 .

[class-id]: The name of the collection
[nft-id]: The id of the nft
[nft-name]: The name of nft	
[nft-uri]: URI of off-chain NFT data
[nft-data]:The data of the nft data [schema.json]
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsClassID := cast.ToString(args[0])
			argsNFTID := cast.ToString(args[1])
			argsNFTName, err := cmd.Flags().GetString(FlagNFTName)
			if err != nil {
				return err
			}
			argsURI, err := cmd.Flags().GetString(FlagNFTURI)
			if err != nil {
				return err
			}
			argsSchema, err := cmd.Flags().GetString(FlagNFTData)
			if err != nil {
				return err
			}
			optionsContent, err := ioutil.ReadFile(argsSchema)
			if err == nil {
				argsSchema = string(optionsContent)
			}
			from := clientCtx.GetFromAddress().String()
			msg := types.NewMsgEditNFT(argsNFTID, argsClassID, argsNFTName, argsURI, argsSchema, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//GetCmdBurnNFT is the CLI command for an BurnNFT transaction
func GetCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [class-id] [nft-id]",
		Short: "burn a nft.",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`burn NFT.
Example:
$ %s tx %s burn-nft "ID66666" "nft-666" --from=mykey --chain-id=plugchain --fees=500plug
This example burning a nft of id nft-666 .

[class-id]: The name of the collection
[nft-id]: The id of the nft
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsClassID := cast.ToString(args[0])
			argsNFTID := cast.ToString(args[1])

			from := clientCtx.GetFromAddress().String()
			msg := types.NewMsgBurnNFT(argsNFTID, argsClassID, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-nft [class-id] [nft-id] [recipient]",
		Short: "transfer an NFT to a recipient.",
		Example: fmt.Sprintf(
			"$ %s tx nft transfer-nft <class-id> <nft-id> <recipient-address> "+
				"--from=myAddress "+
				"--chain-id=plugchain "+
				"--fees=200plug ",
			version.AppName,
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(args[2]); err != nil {
				return err
			}
			msg := types.NewMsgTransferNFT(
				args[1],
				args[0],
				clientCtx.GetFromAddress().String(),
				args[2],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdTransferClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-class [class-id] [recipient]",
		Short: "transfer an Class to a recipient.",
		Example: fmt.Sprintf(
			"$ %s tx nft transfer-class <class-id> <recipient-address> "+
				"--from=myAddress "+
				"--chain-id=plugchain "+
				"--fees=200plug ",
			version.AppName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(args[2]); err != nil {
				return err
			}
			msg := types.NewMsgTransferClass(
				args[0],
				clientCtx.GetFromAddress().String(),
				args[2],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
