package rpc

import (
	"bytes"
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	evmtypes "github.com/oracleNetworkProtocol/ethermint/x/evm/types"
)

// TxLogsFromEvents parses ethereum logs from cosmos events
func TxLogsFromEvents(events []abci.Event) ([]*evmtypes.Log, error) {
	logs := make([]*evmtypes.Log, 0)
	for _, event := range events {
		if event.Type != evmtypes.EventTypeTxLog {
			continue
		}

		for _, attr := range event.Attributes {
			if !bytes.Equal(attr.Key, []byte(evmtypes.AttributeKeyTxLog)) {
				continue
			}

			var log evmtypes.Log
			if err := json.Unmarshal(attr.Value, &log); err != nil {
				return nil, err
			}

			logs = append(logs, &log)
		}
	}
	return logs, nil
}
