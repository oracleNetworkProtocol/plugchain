package cli

import (
	"context"
	"fmt"
	"strings"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/oracleNetworkProtocol/plugchain/x/atom/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group atom queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdAtomUsd())
	return cmd
}

// CmdAtomUsd implements a command to fetch the AtomUsd price.
func CmdAtomUsd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atomUsd",
		Short: "query the Atom/USD price",
		Args:  cobra.NoArgs,
		Long: strings.TrimSpace(`Query genesis parameters for the oracle module:

$ <appd> query oracle params
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAtomUsdRequest{}
			res, err := queryClient.AtomUsd(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.AtomUsd)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
