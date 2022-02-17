package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewClass(id, name, schema, symbol string, owner sdk.AccAddress, mintRestricted, editRestricted bool) Class {
	return Class{
		ID:             id,
		Name:           name,
		Schema:         schema,
		Owner:          owner.String(),
		Symbol:         symbol,
		MintRestricted: mintRestricted,
		EditRestricted: editRestricted,
	}
}
