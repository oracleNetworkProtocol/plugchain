package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/oracleNetworkProtocol/plugchain/x/token/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group token queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryToken(),
		GetCmdQueryParams(),
		GetCmdQueryTokens(),
		GetCmdQueryTokenFee(),
	)

	return cmd
}

// GetCmdQueryToken implements the query token command.
func GetCmdQueryToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info [denom]",
		Short:   "Query a token detail by symbol or min unit.",
		Example: fmt.Sprintf("$ %s query token info <denom>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Token(context.Background(), &types.QueryTokenRequest{
				Denom: args[0],
			})

			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res.Token)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Query values set as token parameters",
		Example: fmt.Sprintf("$ %s q token params", version.AppName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(&res.Params)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tokens [owner]",
		Short:   "Query tokens by the owner",
		Example: fmt.Sprintf("$ %s q token tokens <owner>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var owner sdk.AccAddress
			if len(args) > 0 {
				owner, err = sdk.AccAddressFromBech32(args[0])
				if err != nil {
					return err
				}
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.Tokens(context.Background(), &types.QueryTokensRequest{
				Owner:      owner.String(),
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			tokens := make([]types.Token, 0, len(res.Tokens))
			for _, v := range res.Tokens {
				var token types.TokenI
				if err = clientCtx.InterfaceRegistry.UnpackAny(v, &token); err != nil {
					return err
				}
				t := token.(*types.Token)
				tokens = append(tokens, *t)
			}
			return clientCtx.PrintObjectLegacy(tokens)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all tokens")

	return cmd
}

func GetCmdQueryTokenFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fee [symbol]",
		Short:   "Query the token related fees",
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("$ %s q token fee <symbol>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			symbol := args[0]
			if err := types.ValidateSymbol(symbol); err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Fees(context.Background(), &types.QueryFeesRequest{Symbol: symbol})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
