package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/app"
	"github.com/oracleNetworkProtocol/plugchain/cmd/plugchaind/cmd"
)

func main() {
	setupConfig()

	// app.RegisterDenoms()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	//set the address prefixes
	config := sdk.GetConfig()
	app.SetBech32Prefixes(config)
	app.SetBip44CoinType(config)
	config.Seal()
}
