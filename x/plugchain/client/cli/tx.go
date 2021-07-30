package cli

import (
	"log"
	"time"

	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
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

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		CmdCreateToken(),
		CmdBurnToken(),
	)

	return cmd
}
func CmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [symbol] [description] [originalSymbol] [decimal] [wholeName] [totalSupply] [mintable]",
		Short: "Broadcast message create-token",
		Args:  cobra.ExactArgs(7),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create Token.

Example:
$ %s tx %s create-token ABC "This is First token" "apple boy cat" 100000 "apple boy cat" 10000000000000 true --from mykey

This example creates a token of symbol ABC and totalSupply 10000000000000 .

[symbol]: The symbol is the abbreviation of the currency name
[decimal]: Token decimal places
[totalSupply]: Total token issuance
[mintable]: Is there an additional issuance authority
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			argsSymbol := cast.ToString(args[0])
			argsDescription := cast.ToString(args[1])
			argsOriginalSymbol := cast.ToString(args[2])
			argsDecimal := cast.ToUint64(args[3])
			argsWholeName := cast.ToString(args[4])
			argsTotalSupply := cast.ToString(args[5])
			newSupply, ok := sdk.NewIntFromString(argsTotalSupply)
			if !ok {
				return errors.New("total supply is failed")
			}
			argsMintable, _ := strconv.ParseBool(args[6])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateToken(clientCtx.GetFromAddress().String(), argsWholeName, argsOriginalSymbol, argsDescription, argsSymbol, &newSupply, argsDecimal, argsMintable)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBurnToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [symbol] [account] ",
		Short: "Broadcast message burn-token",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Burn Token.
Example:
$ %s tx %s burn ABC 10000000000000  --from mykey
This example burn a token of symbol ABC and total 10000000000000 .

[symbol]: The symbol is the abbreviation of the currency name
[account]: Total amount of additional issuance
`,
				version.AppName, types.ModuleName,
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {

			argsSymbol := cast.ToString(args[0])
			argsTotalSupply := cast.ToString(args[1])
			newSupply, ok := sdk.NewIntFromString(argsTotalSupply)
			if !ok {
				return errors.New("total amount is failed")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgBurnToken(argsSymbol, clientCtx.GetFromAddress().String(), &newSupply)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			log.Println("sadfasdfasdfasdfasfd")
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
