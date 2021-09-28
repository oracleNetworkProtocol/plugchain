package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewDenom(id, name, schema, symbol string, owner sdk.AccAddress, mintRestricted, editRestricted bool) Denom {
	return Denom{
		ID:             id,
		Name:           name,
		Schema:         schema,
		Owner:          owner.String(),
		Symbol:         symbol,
		MintRestricted: mintRestricted,
		EditRestricted: editRestricted,
	}
}
