package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/keeper"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgIssueDenom    = "op_weight_msg_issue_denom"
	OpWeightMsgMintNFT       = "op_weight_msg_issue_nft"
	OpWeightMsgEditNFT       = "op_weight_msg_edit_nft_tokenData"
	OpWeightMsgTransferNFT   = "op_weight_msg_transfer_nft"
	OpWeightMsgBurnNFT       = "op_weight_msg_transfer_burn_nft"
	OpWeightMsgTransferDenom = "op_weight_msg_transfer_denom"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simulation.WeightedOperations {

	var weightIssueDenom, weightMint, weightEdit, weightBurn, weightTransfer, weightTransferDenom int
	appParams.GetOrGenerate(
		cdc, OpWeightMsgIssueDenom, &weightIssueDenom, nil,
		func(_ *rand.Rand) {
			weightIssueDenom = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgMintNFT, &weightMint, nil,
		func(_ *rand.Rand) {
			weightMint = 100
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgEditNFT, &weightEdit, nil,
		func(_ *rand.Rand) {
			weightEdit = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgTransferNFT, &weightTransfer, nil,
		func(_ *rand.Rand) {
			weightTransfer = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgBurnNFT, &weightBurn, nil,
		func(_ *rand.Rand) {
			weightBurn = 10
		},
	)
	appParams.GetOrGenerate(
		cdc, OpWeightMsgTransferDenom, &weightTransferDenom, nil,
		func(_ *rand.Rand) {
			weightTransferDenom = 10
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightIssueDenom,
			nil,
		),
		simulation.NewWeightedOperation(
			weightMint,
			nil,
		),
		simulation.NewWeightedOperation(
			weightEdit,
			nil,
		),
		simulation.NewWeightedOperation(
			weightTransfer,
			nil,
		),
		simulation.NewWeightedOperation(
			weightBurn,
			nil,
		),
		simulation.NewWeightedOperation(
			weightTransferDenom,
			nil,
		),
	}
}
