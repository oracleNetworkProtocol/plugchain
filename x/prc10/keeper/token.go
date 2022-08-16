package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	gogotypes "github.com/gogo/protobuf/types"
	plugchaintypes "github.com/oracleNetworkProtocol/plugchain/types"
	"github.com/oracleNetworkProtocol/plugchain/x/prc10/types"
)

//GetTokens returns all existing tokens
func (k Keeper) GetTokens(ctx sdk.Context, owner sdk.AccAddress) (tokens []types.TokenI) {
	store := ctx.KVStore(k.storeKey)
	var it sdk.Iterator
	if owner == nil {
		it = sdk.KVStorePrefixIterator(store, types.PrefixTokenForSymbol)
		defer it.Close()
		for ; it.Valid(); it.Next() {
			var token types.Token
			k.cdc.MustUnmarshal(it.Value(), &token)
			tokens = append(tokens, token)
		}
		return
	}
	it = sdk.KVStorePrefixIterator(store, types.KeyTokens(owner, ""))
	defer it.Close()
	for ; it.Valid(); it.Next() {
		var symbol gogotypes.StringValue
		k.cdc.MustUnmarshal(it.Value(), &symbol)
		token, err := k.getTokenBySymbol(ctx, symbol.Value)
		if err != nil {
			continue
		}
		tokens = append(tokens, token)
	}
	return
}

//GetToken returns the token of the specified symbol or min-unit
func (k Keeper) GetToken(ctx sdk.Context, denom string) (types.TokenI, error) {
	// query token by symbol
	if token, err := k.getTokenBySymbol(ctx, denom); err == nil {
		return &token, nil
	}

	// query token by min unit
	if token, err := k.getTokenByMinUnit(ctx, denom); err == nil {
		return &token, nil
	}
	return nil, sdkerrors.Wrapf(types.ErrTokenNotExists, "token: %s does not exist", denom)
}

func (k Keeper) getTokenBySymbol(ctx sdk.Context, symbol string) (token types.Token, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeySymbol(symbol))
	if bz == nil {
		return token, sdkerrors.Wrap(types.ErrTokenNotExists, fmt.Sprintf("token symbol %s does not exist", symbol))
	}
	k.cdc.MustUnmarshal(bz, &token)
	return
}

//getTokenByMinUnit
func (k Keeper) getTokenByMinUnit(ctx sdk.Context, minUnit string) (token types.Token, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyMinUint(minUnit))

	if bz == nil {
		return token, sdkerrors.Wrap(types.ErrTokenNotExists, fmt.Sprintf("token minUnit %s does not exist", minUnit))
	}
	var symbol gogotypes.StringValue

	k.cdc.MustUnmarshal(bz, &symbol)

	return k.getTokenBySymbol(ctx, symbol.Value)
}
func (k Keeper) AddToken(ctx sdk.Context, token types.Token) error {
	if k.HasToken(ctx, token.Symbol) {
		return sdkerrors.Wrapf(types.ErrSymbolAlreadyExists, "symbol already exists: %s", token.Symbol)
	}
	if k.HasToken(ctx, token.MinUnit) {
		return sdkerrors.Wrapf(types.ErrSymbolAlreadyExists, "min-unit already exists: %s", token.MinUnit)
	}
	//set token
	k.setToken(ctx, token)

	if len(token.Owner) != 0 {
		k.setWithOwner(ctx, token.GetOwner(), token.Symbol)
	}

	//Set token to be prefixed with min_unit
	k.setWithMinUnit(ctx, token.MinUnit, token.Symbol)

	data := banktypes.Metadata{
		Name:        token.Name,
		Symbol:      strings.ToUpper(token.Symbol),
		Description: types.DefaultTokenDescription,
		Base:        token.MinUnit,
		Display:     token.Symbol,
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: token.MinUnit, Exponent: 0},
			{Denom: token.Symbol, Exponent: token.Scale},
		},
	}
	if data.Display == plugchaintypes.DisplayNativeDenom {
		data.Description = types.NativeStakingTokenDesc
	}
	k.bankKeeper.SetDenomMetaData(ctx, data)

	return nil
}

//DeleteUpgradeToken upgrade del token func
func (k Keeper) DeleteUpgradeToken(ctx sdk.Context, denom string) error {
	var (
		token types.Token
		err   error
	)
	// query token by symbol
	token, err = k.getTokenBySymbol(ctx, denom)
	if err != nil {
		// query token by min unit
		token, err = k.getTokenByMinUnit(ctx, denom)
		if err != nil {
			return err
		}
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeySymbol(token.Symbol))
	store.Delete(types.KeyMinUint(token.MinUnit))
	store.Delete(types.KeyTokens(token.GetOwner(), token.Symbol))
	return nil
}

//SetUpgradeToken upgrade set token func
func (k Keeper) SetUpgradeToken(ctx sdk.Context, token types.Token) {
	//To prevent duplicates, delete newly created symbols
	k.DeleteUpgradeToken(ctx, token.Symbol)

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&token)
	store.Set(types.KeySymbol(token.Symbol), bz)
	k.setWithOwner(ctx, token.GetOwner(), token.Symbol)
	k.setWithMinUnit(ctx, token.MinUnit, token.Symbol)
}

func (k Keeper) setWithMinUnit(ctx sdk.Context, minUnit, symbol string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: symbol})
	store.Set(types.KeyMinUint(minUnit), bz)
}
func (k Keeper) setWithOwner(ctx sdk.Context, owner sdk.AccAddress, symbol string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: symbol})
	store.Set(types.KeyTokens(owner, symbol), bz)
}

func (k Keeper) setToken(ctx sdk.Context, token types.Token) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&token)
	store.Set(types.KeySymbol(token.Symbol), bz)
}

func (k Keeper) HasToken(ctx sdk.Context, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	if k.HasSymbol(ctx, denom) {
		return true
	}

	return store.Has(types.KeyMinUint(denom))
}

func (k Keeper) HasSymbol(ctx sdk.Context, symbol string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeySymbol(symbol))
}

// GetParamSet returns token params from the global param store
func (k Keeper) GetParamSet(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

func (k Keeper) resetStoreKeyForQueryToken(ctx sdk.Context, symbol string, owner, to sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyTokens(owner, symbol))
	k.setWithOwner(ctx, to, symbol)
}
