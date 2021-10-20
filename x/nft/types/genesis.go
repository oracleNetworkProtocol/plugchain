package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, item := range gs.Collections {
		if err := ValidateDenomID(item.Denom.ID); err != nil {
			return err
		}
		for _, nft := range item.NFTs {
			if err := ValidateNFTID(nft.ID); err != nil {
				return err
			}
			if err := ValidateNFTURL(nft.URL); err != nil {
				return err
			}
			if nft.GetOwner().Empty() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner")
			}
		}

	}
	return nil
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(collections []Collection) *GenesisState {
	return &GenesisState{
		Collections: collections,
	}
}
