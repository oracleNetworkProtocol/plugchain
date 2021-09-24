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
	"github.com/cosmos/cosmos-sdk/version"

	// "github.com/cosmos/cosmos-sdk/client/flags"
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
	)
	return cmd
}

// GetCmdIssueDenom is the CLI command for an IssueDenom transaction
func GetCmdIssueDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [denom-id] [denom-name] [denom-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json]",
		Short: "Issue a new denom.",
		Args:  cobra.ExactArgs(7),
		Long: strings.TrimSpace(
			fmt.Sprintf(`issue Denom.
Example:
$ %s tx %s issue "ID66666" "first-denom" "shui" true true ./schema-ID66666.json --from mykey --chain-id plugchain --fees 500plug
This example creates a denom of id ID666666 and name first-denom .
[denom-id]: The name of the collection
[denom-name]: The name of the denom
[mint-restricted]: mint restricted of nft under denom
[edit-restricted]: update restricted of nft under denom
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
