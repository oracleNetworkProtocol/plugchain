package cli

import (
	"context"
	"fmt"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group nft queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetQueryDenomCmd(),
	)

	return cmd
}

func GetQueryDenomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "denom [denom-id]",
		Long:    "Query the denom by the specified denom id.",
		Example: fmt.Sprintf("$ %s q nft denom <denom-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Denom(
				context.Background(),
				&types.QueryDenomRequest{DenomId: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp.Denom)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetQueryDenomsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "denoms",
		Long:    "Query all denominations of all collections of NFTs.",
		Example: fmt.Sprintf("$ %s q nft denoms", version.AppName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Denoms(context.Background(), &types.QueryDenomsRequest{Pagination: pageReq})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all denoms")
	return cmd
}
