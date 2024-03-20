package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracle "github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
	oracletypes "github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

type (
	// OracleKeeper interface
	OracleKeeper interface {
		FinalizeRound(ctx sdk.Context, claimType string, roundID uint64)
		GetPendingRounds(ctx sdk.Context, roundType string) (rounds []uint64)
		TallyVotes(ctx sdk.Context, claimType string, roundID uint64) *oracletypes.RoundResult
		GetClaim(ctx sdk.Context, hash tmbytes.HexBytes) oracle.Claim
	}
)
