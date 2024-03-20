package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
	"github.com/spf13/cobra"
)

// CmdClaim queries a claim by hash
func CmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim",
		Short: "query claim by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			hash := args[0]
			params := &types.QueryClaimRequest{ClaimHash: hash}

			res, err := queryClient.Claim(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdAllClaims queries all clams
func CmdAllClaims() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-claims",
		Short: "query all claims",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllClaimsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllClaims(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
