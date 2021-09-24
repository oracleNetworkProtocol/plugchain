package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) SetDenom(ctx sdk.Context, denom types.Denom) error {
	if k.HasDenomByID(ctx, denom.ID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s has already exists", denom.ID)
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&denom)
	store.Set(types.GetKeyDenomID(denom.ID), bz)
	return nil
}

func (k Keeper) HasDenomByID(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetKeyDenomID(id))
}
