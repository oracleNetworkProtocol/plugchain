package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
)

// SetValidatorDelegateAddress sets a new address that will have the power to send data on behalf of the validator
// we store the mapping in both directions for easy lookup
func (k Keeper) SetValidatorDelegateAddress(ctx sdk.Context, val, del sdk.AccAddress) {
	if val.String() == del.String() {
		ctx.KVStore(k.storeKey).Delete(types.GetDelValKey(del))
		ctx.KVStore(k.storeKey).Delete(types.GetValDelKey(val))
	} else {
		ctx.KVStore(k.storeKey).Set(types.GetDelValKey(del), val.Bytes())
		ctx.KVStore(k.storeKey).Set(types.GetValDelKey(val), del.Bytes())
	}
}

// GetValidatorAddressFromDelegate returns the delegate address for a given validator
func (k Keeper) GetValidatorAddressFromDelegate(ctx sdk.Context, del sdk.AccAddress) sdk.AccAddress {
	return sdk.AccAddress(ctx.KVStore(k.storeKey).Get(types.GetDelValKey(del)))
}

// GetDelegateAddressFromValidator returns the validator address for a given delegate
func (k Keeper) GetDelegateAddressFromValidator(ctx sdk.Context, val sdk.AccAddress) sdk.AccAddress {
	return sdk.AccAddress(ctx.KVStore(k.storeKey).Get(types.GetValDelKey(val)))
}

// IsDelegateAddress returns true if the validator has delegated their feed to an address
func (k Keeper) IsDelegateAddress(ctx sdk.Context, del sdk.AccAddress) bool {
	return ctx.KVStore(k.storeKey).Has(types.GetDelValKey(del))
}

// GetAllDelegations return all delegations
func (k Keeper) GetAllDelegations(ctx sdk.Context) []types.MsgDelegate {
	delegations := []types.MsgDelegate{}
	k.IterateDelegations(ctx, func(del sdk.AccAddress, val sdk.AccAddress) (stop bool) {
		delegations = append(delegations, *types.NewMsgDelegate(val, del))
		return
	})
	return delegations
}

// IterateDelegations iterates over all delegate address pairs in the store
func (k Keeper) IterateDelegations(ctx sdk.Context, handler func(del, val sdk.AccAddress) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.DelValKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		del := sdk.AccAddress(bytes.TrimPrefix(iter.Key(), types.DelValKey))
		val := sdk.AccAddress(iter.Value())
		if handler(del, val) {
			break
		}
	}
}
