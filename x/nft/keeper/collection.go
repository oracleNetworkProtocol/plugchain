package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

func (k Keeper) GetCollection(ctx sdk.Context, denomID string) (types.Collection, error) {
	denom, ok := k.GetDenomByID(ctx, denomID)
	if !ok {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID (%s) not existed", denomID)
	}
	nfts := k.GetNFTs(ctx, denomID)
	return types.NewCollection(denom, nfts), nil
}
