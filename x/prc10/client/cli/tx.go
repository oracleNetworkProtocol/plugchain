package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/oracleNetworkProtocol/plugchain/x/prc10/types"
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
		GetCmdIssueToken(),
		GetCmdEditToken(),
		GetCmdMintToken(),
		GetCmdBurnToken(),
		GetCmdTransferOwnerToken(),
	)
	return cmd
}

// GetCmdIssueToken implements the issue token command
func GetCmdIssueToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue",
		Short: "Issue a new token.",
		Example: fmt.Sprintf(
			"$ %s tx token issue "+
				"--name=\"Kitty Token\" "+
				"--symbol=\"kitty\" "+
				"--min-unit=\"kitty\" "+
				"--scale=0 "+
				"--initial-supply=100000000000 "+
				"--max-supply=1000000000000 "+
				"--mintable=true "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()
			symbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return err
			}
			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}
			minUnit, err := cmd.Flags().GetString(FlagMinUnit)
			if err != nil {
				return err
			}
			scale, err := cmd.Flags().GetUint32(FlagScale)
			if err != nil {
				return err
			}
			initialSupply, err := cmd.Flags().GetUint64(FlagInitialSupply)
			if err != nil {
				return err
			}
			maxSupply, err := cmd.Flags().GetUint64(FlagMaxSupply)
			if err != nil {
				return err
			}
			mintable, err := cmd.Flags().GetBool(FlagMintable)
			if err != nil {
				return err
			}

			msg := &types.MsgIssueToken{
				Symbol:        symbol,
				Name:          name,
				MinUnit:       minUnit,
				Scale:         scale,
				InitialSupply: initialSupply,
				MaxSupply:     maxSupply,
				Mintable:      mintable,
				Owner:         owner.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsIssueToken)
	_ = cmd.MarkFlagRequired(FlagSymbol)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagInitialSupply)
	_ = cmd.MarkFlagRequired(FlagScale)
	_ = cmd.MarkFlagRequired(FlagMinUnit)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
func GetCmdMintToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [symbol]",
		Short: "Mint tokens to a specified address.",
		Example: fmt.Sprintf(
			"$ %s tx token mint <symbol> "+
				"--amount=<amount> "+
				"--to=<to> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			amount, err := cmd.Flags().GetUint64(FlagAmount)
			if err != nil {
				return err
			}

			to, err := cmd.Flags().GetString(FlagTo)
			if err != nil {
				return err
			}
			if len(to) > 0 {
				if _, err = sdk.AccAddressFromBech32(to); err != nil {
					return err
				}
			}

			msg := types.NewMsgMintToken(
				strings.TrimSpace(args[0]), amount, to, owner,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsMintToken)

	_ = cmd.MarkFlagRequired(FlagAmount)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditToken implements the edit token command
func GetCmdEditToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [symbol]",
		Short: "Edit an existing token.",
		Example: fmt.Sprintf(
			"$ %s tx token edit <symbol> "+
				"--name=\"Cat Token\" "+
				"--max-supply=100000000000 "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}
			maxSupply, err := cmd.Flags().GetUint64(FlagMaxSupply)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditToken(args[0], name, maxSupply, owner)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsEditToken)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdBurnToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [symbol]",
		Short: "Burn tokens.",
		Example: fmt.Sprintf(
			"$ %s tx token burn <symbol> "+
				"--amount=<amount> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress().String()

			amount, err := cmd.Flags().GetUint64(FlagAmount)
			if err != nil {
				return err
			}

			addr, err := cmd.Flags().GetString(FlagTo)
			if err != nil {
				return err
			}
			if len(addr) > 0 {
				if _, err = sdk.AccAddressFromBech32(addr); err != nil {
					return err
				}
			}

			msg := types.NewMsgBurnToken(
				strings.TrimSpace(args[0]), amount, owner,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsMintToken)
	_ = cmd.MarkFlagRequired(FlagAmount)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTransferOwnerToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "transfer owner for token.",
		Example: fmt.Sprintf(
			"$ %s tx token transfer <symbol>"+
				"--to=<new-owner-address>"+
				"--from=<key-name>"+
				"--chain-id=<chain-id>"+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			owner := clientCtx.GetFromAddress().String()
			symbol := args[0]

			to, err := cmd.Flags().GetString(FlagTo)
			if err != nil {
				return err
			}
			if len(to) > 0 {
				if _, err := sdk.AccAddressFromBech32(to); err != nil {
					return err
				}
			}

			msg := types.NewMsgTransferOwnerToken(symbol, owner, to)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	cmd.Flags().AddFlagSet(FsTransferOwnerToken)
	_ = cmd.MarkFlagRequired(FlagTo)

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
