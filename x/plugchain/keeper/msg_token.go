package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

func (k Keeper) NewMintToken(ctx sdk.Context, msg types.MsgCreateToken) error {
	AccCreatorAddress := msg.GetTokenCreator()

	if k.HasTokenBySymbol(ctx, msg.Symbol) {
		return types.ReturnErrSymbolIsExists(msg.Symbol)
	}
	mintCoin := sdk.NewCoin(msg.Symbol, *msg.TotalSupply)
	mintCoins := sdk.NewCoins(mintCoin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, AccCreatorAddress, mintCoins); err != nil {
		return err
	}
	return nil
}
