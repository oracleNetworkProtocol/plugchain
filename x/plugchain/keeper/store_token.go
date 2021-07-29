package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

func (k Keeper) HasTokenBySymbol(ctx sdk.Context, symbol string) bool {
	store := ctx.KVStore(k.storeKey)
	ok := store.Has(types.GetTokenKey(symbol))
	k.Logger(ctx).Info("has token by symbol success",
		"symbol-store-key", types.GetTokenKey(symbol),
		"symbol", symbol,
		"isExists", ok,
	)
	return ok
}

func (k Keeper) GetToken(ctx sdk.Context, key []byte) types.MsgCreateToken {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	var tokenInfo types.MsgCreateToken
	if nil != bz {
		k.cdc.MustUnmarshalBinaryBare(bz, &tokenInfo)
	}
	return tokenInfo
}

func (k Keeper) SetToken(ctx sdk.Context, tokenMsg types.MsgCreateToken) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&tokenMsg)
	store.Set(types.GetTokenKey(tokenMsg.Symbol), bz)
	keyID := k.GetNextTokenIDWithInit(ctx)
	k.Logger(ctx).Info("set-token success",
		"ID", keyID,
		"symbol", tokenMsg.Symbol,
		"total-supply", tokenMsg.TotalSupply,
	)
	k.SetAccTokensWithInit(ctx, tokenMsg)
}

func (k Keeper) GetNextTokenIDWithInit(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	keyID := k.GetNextTokenID(ctx)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: keyID + 1})
	store.Set(types.TokenCountKey, bz)
	return keyID
}

func (k Keeper) GetNextTokenID(ctx sdk.Context) uint64 {
	var TokenID uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenCountKey)
	if nil == bz {
		TokenID = 0
	} else {
		val := gogotypes.UInt64Value{}
		err := k.cdc.UnmarshalBinaryBare(bz, &val)
		if err != nil {
			panic(err)
		}
		TokenID = val.GetValue()
	}
	return TokenID
}

func (k Keeper) SetAccTokensWithInit(ctx sdk.Context, msg types.MsgCreateToken) types.AccAddressTokens {
	var (
		zeroString   string
		store        = ctx.KVStore(k.storeKey)
		accAdrBech32 = msg.GetTokenCreator()
		accTokens    = k.GetAccTokens(ctx, accAdrBech32)
	)
	if zeroString == accTokens.Address {
		accTokens.Address = msg.Owner
	}

	accTokens.Count += 1
	accTokens.Key = append(accTokens.Key, types.GetTokenKey(msg.Symbol))

	bz := k.cdc.MustMarshalBinaryBare(&accTokens)
	store.Set(types.GetTokenAccAddressPrefix(accAdrBech32), bz)
	return accTokens
}
func (k Keeper) GetAccTokens(ctx sdk.Context, accAdrBech32 []byte) types.AccAddressTokens {
	var accToken types.AccAddressTokens
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTokenAccAddressPrefix(accAdrBech32))
	if nil != bz {
		err := k.cdc.UnmarshalBinaryBare(bz, &accToken)
		if err != nil {
			panic(err)
		}
	}
	return accToken
}
