package rpc

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type EthLog struct {
	*ethtypes.Log
}

func (l EthLog) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address     sdk.AccAddress `json:"address" gencodec:"required"`
		Topics      []common.Hash  `json:"topics" gencodec:"required"`
		Data        hexutil.Bytes  `json:"data" gencodec:"required"`
		BlockNumber uint64         `json:"blockNumber"`
		TxHash      common.Hash    `json:"transactionHash" gencodec:"required"`
		TxIndex     uint           `json:"transactionIndex"`
		BlockHash   common.Hash    `json:"blockHash"`
		Index       uint           `json:"logIndex"`
		Removed     bool           `json:"removed"`
	}
	var enc Log
	enc.Address = sdk.AccAddress(l.Address[:])
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = l.BlockNumber
	enc.TxHash = l.TxHash
	enc.TxIndex = l.TxIndex
	enc.BlockHash = l.BlockHash
	enc.Index = l.Index
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}
