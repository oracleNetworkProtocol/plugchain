package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) GetClassByID(ctx sdk.Context, id string) (denom types.Class, ok bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetKeyClassID(id))
	if len(bz) == 0 {
		return denom, false
	}
	k.cdc.MustUnmarshal(bz, &denom)
	return denom, true
}

func (k Keeper) GetClasses(ctx sdk.Context) (denoms []types.Class) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetKeyClassID(""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var denom types.Class
		k.cdc.MustUnmarshal(iterator.Value(), &denom)
		denoms = append(denoms, denom)
	}
	return
}

func (k Keeper) SetClass(ctx sdk.Context, denom types.Class) error {

	if k.HasClassByID(ctx, denom.ID) {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "classID %s has already exists", denom.ID)
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&denom)
	store.Set(types.GetKeyClassID(denom.ID), bz)

	return nil
}

func (k Keeper) HasClassByID(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetKeyClassID(id))
}
func (k Keeper) UpdateDenom(ctx sdk.Context, denom types.Class) error {
	if !k.HasClassByID(ctx, denom.ID) {
		return sdkerrors.Wrapf(types.ErrInvalidClass, "classID %s not exists", denom.ID)
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&denom)
	store.Set(types.GetKeyClassID(denom.ID), bz)

	return nil
}
