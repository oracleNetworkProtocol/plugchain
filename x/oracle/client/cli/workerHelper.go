package cli

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
	"github.com/spf13/cobra"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tendermint "github.com/tendermint/tendermint/types"
)

// store claim and salt of previous round to submit after prevote
// TODO: a non-ephemeral storage option?
type prevClaim struct {
	claim exported.Claim
	salt  string
}

var prevVotes = map[string](*prevClaim){}

// WorkerHelper helps with oracle queries and tx
type WorkerHelper struct {
	cmd          *cobra.Command
	OracleParams types.Params
	QueryClient  types.QueryClient
	BlockHeight  int64
}

// NewWorkerHelper is a worker helper util
func NewWorkerHelper(cmd *cobra.Command, blockEvent ctypes.ResultEvent) (*WorkerHelper, error) {
	blockHeight :=
		blockEvent.Data.(tendermint.EventDataNewBlock).Block.Header.Height

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
	}
	// Make sure we querey the data from the right height
	clientCtx.Height = blockHeight
	queryClient := types.NewQueryClient(clientCtx)

	params := &types.QueryParamsRequest{}

	// we query the array of claim parameters for different claims
	res, err := queryClient.Params(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &WorkerHelper{
		cmd:          cmd,
		OracleParams: res.Params,
		QueryClient:  queryClient,
		BlockHeight:  blockHeight,
	}, nil
}

// SubmitWorkerTx submits the correct transactions based on oracle params
func (h WorkerHelper) SubmitWorkerTx(claim exported.Claim) {

	///////////////
	/// NO PREVOTE
	///////////////
	claimType := claim.Type()
	usePrevote := h.OracleParams.ClaimParams[claimType].Prevote

	// use the most recent data and no salt = ""
	if usePrevote == false {
		err := h.SubmitVote(claim, "")
		if err != nil {
			fmt.Println("Error submitting claim vote", err)
		}
		return
	}

	///////////////
	/// YES PREVOTE
	///////////////

	// 1. submit data for prev round
	// 2. submit prevote for current round
	prev := prevVotes[claimType]
	if prev != nil {
		err := h.SubmitVote(prev.claim, prev.salt)
		if err != nil {
			fmt.Println("Error submitting claim vote", err)
		}
		// clear the submitted claim
		prevVotes[claimType] = nil
	}
	err := h.SubmitPrevote(claim)
	if err != nil {
		fmt.Println("Error submitting prevote", err)
	}
}

// SubmitPrevote submits prevote of current claim
func (h WorkerHelper) SubmitPrevote(claim exported.Claim) error {
	clientCtx, err := client.GetClientTxContext(h.cmd)
	if err != nil {
		return err
	}

	salt := genrandstr(9)
	prevoteHash := types.VoteHash(salt, claim.Hash().String(), clientCtx.FromAddress)
	prevoteMsg := types.NewMsgPrevote(clientCtx.GetFromAddress(), prevoteHash)
	if err := prevoteMsg.ValidateBasic(); err != nil {
		return err
	}
	err = tx.GenerateOrBroadcastTxCLI(clientCtx, h.cmd.Flags(), prevoteMsg)
	if err != nil {
		return err
	}
	prevVotes[claim.Type()] = &prevClaim{
		claim: claim,
		salt:  salt,
	}
	return nil
}

// SubmitVote submits a vote of the given claim + salt
func (h WorkerHelper) SubmitVote(claim exported.Claim, salt string) error {
	clientCtx, err := client.GetClientTxContext(h.cmd)
	if err != nil {
		return err
	}

	voteMsg, err := types.NewMsgVote(clientCtx.GetFromAddress(), claim, salt)
	if err != nil {
		return err
	}
	if err := voteMsg.ValidateBasic(); err != nil {
		return err
	}

	// Submit Claim from previous round
	err = tx.GenerateOrBroadcastTxCLI(clientCtx, h.cmd.Flags(), voteMsg)

	// TODO retry if the TX fails
	if err != nil {
		return err
	}
	return nil
}

// IsRoundStart determins if the current blockHeight is the start of a round
func (h WorkerHelper) IsRoundStart(claimType string) bool {
	votePeriod := h.OracleParams.ClaimParams[claimType].VotePeriod
	// default is every round
	if votePeriod == 0 {
		return true
	}
	return h.BlockHeight%int64(h.OracleParams.ClaimParams[claimType].VotePeriod) == 0
}

func genrandstr(s int) string {
	b := make([]byte, s)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
