package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type NFTI interface {
	GetOwner() sdk.AccAddress
	GetData() string
}
