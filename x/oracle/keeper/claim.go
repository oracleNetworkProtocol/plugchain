package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// CreateClaim sets Claim by hash in the module's KVStore.
func (k Keeper) CreateClaim(ctx sdk.Context, claim exported.Claim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ClaimKey)
	store.Set(claim.Hash(), k.MustMarshalClaim(claim))
}

// GetClaim retrieves Claim by hash if it exists. If no Claim exists for
// the given hash, (nil, false) is returned.
func (k Keeper) GetClaim(ctx sdk.Context, hash tmbytes.HexBytes) exported.Claim {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ClaimKey)

	bz := store.Get(hash)
	if len(bz) == 0 {
		return nil
	}

	return k.MustUnmarshalClaim(bz)
}

// GetAllClaims returns all claims
func (k Keeper) GetAllClaims(ctx sdk.Context) (msgs []exported.Claim) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ClaimKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		claim := k.MustUnmarshalClaim(iterator.Value())
		msgs = append(msgs, claim)
	}

	return
}

// DeleteClaim deletes claim by hash
func (k Keeper) DeleteClaim(ctx sdk.Context, hash tmbytes.HexBytes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ClaimKey)
	store.Delete(hash)
}

// MustUnmarshalClaim attempts to decode and return an Claim object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalClaim(bz []byte) exported.Claim {
	Claim, err := k.UnmarshalClaim(bz)
	if err != nil {
		panic(fmt.Errorf("failed to decode Claim: %w", err))
	}

	return Claim
}

// MustMarshalClaim attempts to encode an Claim object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalClaim(Claim exported.Claim) []byte {
	bz, err := k.MarshalClaim(Claim)
	if err != nil {
		panic(fmt.Errorf("failed to encode Claim: %w", err))
	}

	return bz
}

// MarshalClaim marshals a Claim interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalClaim(claimI exported.Claim) ([]byte, error) {
	return k.cdc.MarshalInterface(claimI)
}

// UnmarshalClaim returns a Claim interface from raw encoded Claim
// bytes of a Proto-based Claim type. An error is returned upon decoding
// failure.
func (k Keeper) UnmarshalClaim(bz []byte) (exported.Claim, error) {
	var claim exported.Claim
	if err := k.cdc.UnmarshalInterface(bz, &claim); err != nil {
		return nil, err
	}

	return claim, nil
}
