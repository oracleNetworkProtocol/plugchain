package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type NFTI interface {
	GetOwner() sdk.AccAddress
	GetData() string
}

var _ NFTI = NFT{}

func NewNFT(id, name, url, data string, owner sdk.AccAddress) NFT {
	return NFT{
		ID:    id,
		Name:  name,
		URL:   url,
		Data:  data,
		Owner: owner.String(),
	}
}

func (nf NFT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(nf.Owner)
	return owner
}

func (nf NFT) GetData() string {
	return nf.Data
}
