package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

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
		GetQueryDenomsCmd(),
		GetQueryNFTCmd(),
		GetQueryCollectionCmd(),
		GetQuerySupplyCmd(),
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

func GetQueryNFTCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nft [denom-id] [nft-id]",
		Long:    "Query a single NFT from a collection",
		Example: fmt.Sprintf("$ %s q nft nft <denom-id> <nft-id>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.NFT(context.Background(), &types.QueryNFTRequest{DenomId: args[0], NftId: args[1]})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetQueryCollectionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "collection [denom-id]",
		Long:    "Get all the NFTs from a given collection.",
		Example: fmt.Sprintf("$ %s q nft collection <denom-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
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
			resp, err := queryClient.Collection(
				context.Background(),
				&types.QueryCollectionRequest{
					DenomId:    args[0],
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nfts")
	return cmd
}

func GetQuerySupplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "supply [denomID]",
		Long:    "total supply of a collection or owner of NFTs",
		Example: fmt.Sprintf("$ %s q nft supply <denom-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Supply(context.Background(), &types.QuerySupplyRequest{
				DenomId: args[0],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}
	cmd.Flags().AddFlagSet(FsQuerySupply)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
