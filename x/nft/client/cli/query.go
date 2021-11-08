package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		GetQueryClassCmd(),
		GetQueryClassesCmd(),
		GetQueryNFTCmd(),
		GetQueryCollectionCmd(),
		GetQuerySupplyCmd(),
		GetQueryOwnerCmd(),
	)

	return cmd
}

func GetQueryClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "class [class-id]",
		Long:    "Query the class by the specified class id.",
		Example: fmt.Sprintf("$ %s q nft class <class-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Class(
				context.Background(),
				&types.QueryClassRequest{ClassId: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp.Class)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetQueryClassesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "classes",
		Long:    "Query all denominations of all collections of NFTs.",
		Example: fmt.Sprintf("$ %s q nft classes", version.AppName),
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
			resp, err := queryClient.Classes(context.Background(), &types.QueryClassesRequest{Pagination: pageReq})
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
		Use:     "nft [class-id] [nft-id]",
		Long:    "Query a single NFT from a collection",
		Example: fmt.Sprintf("$ %s q nft nft <class-id> <nft-id>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.NFT(context.Background(), &types.QueryNFTRequest{ClassId: args[0], NftId: args[1]})
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
		Use:     "collection [class-id]",
		Long:    "Get all the NFTs from a given collection.",
		Example: fmt.Sprintf("$ %s q nft collection <class-id>", version.AppName),
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
					ClassId:    args[0],
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
		Use:     "supply [classID]",
		Long:    "total supply of a collection or owner of NFTs",
		Example: fmt.Sprintf("$ %s q nft supply <class-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Supply(context.Background(), &types.QuerySupplyRequest{
				ClassId: args[0],
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

func GetQueryOwnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "owner [address] [class-id]",
		Long:    "Get the NFTs owned by an account addr.",
		Example: fmt.Sprintf("$ %s q nft owner <address> <denomID>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Owner(context.Background(), &types.QueryOwnerRequest{
				ClassId:    args[1],
				Address:    args[0],
				Pagination: pageReq,
			})
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
