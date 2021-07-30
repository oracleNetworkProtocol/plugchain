package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

//验证创建token 信息
//币种缩写长度不大于4 ，缩写纯大写
//小数位数最小4位
func (k Keeper) ValildateCreateToken(ctx sdk.Context, msg types.MsgCreateToken) error {
	if len(msg.Symbol) > len(types.TokenNamePrefix)+types.SymbolMaxLen {
		return types.ReturnErrSymbolLenNotValid(types.SymbolMaxLen)
	}
	if len(msg.Description) > types.DescriptionMaxLen {
		return types.ReturnErrDescriptionNotValid(types.DescriptionMaxLen)
	}

	if msg.Decimal < types.MinDecimal {
		return types.ReturnErrDecimalNotValid(types.MinDecimal)
	}

	//判断最小发行量
	if msg.TotalSupply.Uint64() < 1 {
		return types.ReturnErrTotalSupplyNotValid(types.MinTotalSupply)
	}
	return nil
}

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
