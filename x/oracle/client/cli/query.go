package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group oracle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdAllClaims())
	cmd.AddCommand(CmdClaim())
	cmd.AddCommand(CmdParams())
	cmd.AddCommand(CmdRound())
	cmd.AddCommand(CmdPendingRounds())
	cmd.AddCommand(CmdAllRounds())
	cmd.AddCommand(CmdValidatorAddress())
	cmd.AddCommand(CmdDelegeateAddress())

	return cmd
}

// CmdParams implements a command to fetch oracle parameters.
func CmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "query the current oracle parameters",
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

			params := &types.QueryParamsRequest{}
			res, err := queryClient.Params(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdRound implements a command to fetch an oracle round.
func CmdRound() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "round",
		Short: "query round by claimType and roundID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			claimType := args[0]
			roundID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryRoundRequest{ClaimType: claimType, RoundId: roundID}

			res, err := queryClient.Round(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdPendingRounds implements a command to fetch all pending oracle rounds.
func CmdPendingRounds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-rounds",
		Short: "query pending rounds",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			claimType := args[0]
			params := &types.QueryPendingRoundsRequest{ClaimType: claimType}

			res, err := queryClient.PendingRounds(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdAllRounds implements a command to fetch all rounds.
func CmdAllRounds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-rounds",
		Short: "query all rounds",
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
			params := &types.QueryAllRoundsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllRounds(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdDelegeateAddress query delegate address from the chain given validators address
func CmdDelegeateAddress() *cobra.Command {
	return &cobra.Command{
		Use:     "delegate-address [validator-address]",
		Aliases: []string{"del"},
		Args:    cobra.ExactArgs(1),
		Short:   "query delegate address from the chain given validators address",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryDelegeateAddressRequest{Validator: args[0]}

			res, err := queryClient.QueryDelegeateAddress(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
}

// CmdValidatorAddress query validator address from the chain given the address that validator delegated to
func CmdValidatorAddress() *cobra.Command {
	return &cobra.Command{
		Use:     "validator-address [delegate-address]",
		Aliases: []string{"val"},
		Args:    cobra.ExactArgs(1),
		Short:   "query validator address from the chain given the address that validator delegated to",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryValidatorAddressRequest{Delegate: args[0]}

			res, err := queryClient.QueryValidatorAddress(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
}
