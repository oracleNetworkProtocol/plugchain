package keeper

import (
	"fmt"
	golog "log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/oracleNetworkProtocol/plugchain/x/token/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		moduleAddrs   map[string]bool
		paramSpace    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	moduleAddr map[string]bool,
) *Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	golog.Println("token module address:", accountKeeper.GetModuleAddress(types.ModuleName).String())
	golog.Printf("module address >>>>>>: %+v", moduleAddr)

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		paramSpace:    paramSpace,
		moduleAddrs:   moduleAddr,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IssueToken(
	ctx sdk.Context,
	symbol, name, minUnit string,
	scale uint32,
	initialSupply uint64,
	maxSupply uint64,
	mintable bool,
	owner sdk.AccAddress,
) error {
	token := types.NewToken(
		symbol, name, minUnit, scale, initialSupply,
		maxSupply, mintable, owner,
	)
	if err := k.AddToken(ctx, token); err != nil {
		return err
	}

	precision := sdk.NewIntWithDecimal(1, int(token.Scale))
	initialCoin := sdk.NewCoin(
		token.MinUnit, sdk.NewIntFromUint64(token.InitialSupply).Mul(precision),
	)
	mintCoins := sdk.NewCoins(initialCoin)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, mintCoins)
}

func (k Keeper) MintToken(ctx sdk.Context, symbol string, amount uint64, to, owner sdk.AccAddress) error {

	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}
	if !token.Mintable {
		return sdkerrors.Wrapf(types.ErrNotMintable, "%s", symbol)
	}

	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", token.Owner, symbol)
	}
	supply := k.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(token.MinUnit)
	precision := sdk.NewIntWithDecimal(1, int(token.Scale))
	mintableAmt := sdk.NewIntFromUint64(token.MaxSupply).Mul(precision).Sub(supply)
	mintableMainAmt := mintableAmt.Quo(precision).Uint64()

	if amount > mintableMainAmt {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			"the amount exceeds the mintable token amount; expected (0, %d], got %d",
			mintableMainAmt, amount,
		)
	}

	mintCoins := sdk.NewCoins(
		sdk.NewCoin(token.MinUnit, sdk.NewIntFromUint64(amount).Mul(precision)),
	)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}
	if to.Empty() {
		to = owner
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, mintCoins)
}

func (k Keeper) EditToken(ctx sdk.Context, symbol, name string, maxSupply uint64, owner sdk.AccAddress) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}
	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", token.Owner, symbol)
	}
	if maxSupply > 0 {
		issueSupply := k.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(token.MinUnit)
		issueBaseSupply := uint64(issueSupply.Quo(sdk.NewIntWithDecimal(1, int(token.Scale))).Int64())

		if maxSupply < issueBaseSupply {
			return sdkerrors.Wrapf(types.ErrInvalidMaxSupply, "max supply must not be less than %d", issueBaseSupply)
		}

		token.MaxSupply = maxSupply
	}
	if name != types.DoNotModify {
		token.Name = name
		metadata := k.bankKeeper.GetDenomMetaData(ctx, token.MinUnit)
		metadata.Description = name
		k.bankKeeper.SetDenomMetaData(ctx, metadata)
	}
	k.setToken(ctx, token)

	return nil
}

func (k Keeper) BurnToken(ctx sdk.Context, symbol string, amount uint64, owner sdk.AccAddress) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}
	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", token.Owner, symbol)
	}
	precision := sdk.NewIntWithDecimal(1, int(token.Scale))
	burnCoin := sdk.NewCoin(token.MinUnit, sdk.NewIntFromUint64(amount).Mul(precision))
	burnCoins := sdk.NewCoins(burnCoin)

	addrTotal := k.bankKeeper.GetBalance(ctx, owner, symbol)

	if !addrTotal.Amount.GT(burnCoin.Amount) {
		return sdkerrors.Wrapf(types.ErrInvalidAmount, "the amount exceeds the account token amount; expected (0, %d], got %d", addrTotal.Amount.String(), burnCoin.Amount.String())
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, burnCoins); err != nil {
		return err
	}

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
}

func (k Keeper) TransferOwnerToken(ctx sdk.Context, symbol string, owner, to sdk.AccAddress) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}
	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", owner, symbol)
	}
	token.Owner = to.String()

	k.setToken(ctx, token)

	//重置地址查询索引
	k.resetStoreKeyForQueryToken(ctx, token.Symbol, owner, to)

	return nil
}

// SetParamSet sets token params to the global param store
func (k Keeper) SetParamSet(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
