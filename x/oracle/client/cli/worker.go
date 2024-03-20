package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tendermint "github.com/tendermint/tendermint/types"
	"golang.org/x/sync/errgroup"
)

var (
	txQuery    = "tm.event='Tx'"
	blockQuery = "tm.event='NewBlock'"
	once       sync.Once
)

// BlockHandler is the code that the worker must run every block
type BlockHandler func(*cobra.Command, ctypes.ResultEvent) error

// TxHandler is the code that the worker runs on every tx event
type TxHandler func(*cobra.Command, ctypes.ResultEvent) error

// Worker is a singleton type
type Worker struct {
	handleBlock BlockHandler
	handleTx    TxHandler
}

var instance *Worker

// InitializeWorker intializes the singleton instance and sets the worker process
// This method should be called from the app's cmd/root.go file
func InitializeWorker(blockHandler BlockHandler, txHandler TxHandler) *Worker {
	instance = &Worker{
		handleBlock: blockHandler,
		handleTx:    txHandler,
	}
	return instance
}

// StartWorkerCmd starts the off-chain worker process
func StartWorkerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-worker [stopAtHeight]",
		Short: "Starts offchain worker",
		RunE: func(cmd *cobra.Command, args []string) error {

			if instance == nil {
				return fmt.Errorf("Worker process is not intialized, did you forget to initialize the worker and set handlers in root.go?")
			}

			// the process after stopAtHeight (used for testing)
			stopAtHeight := int64(0)
			var err error
			if len(args) > 0 {
				stopAtHeight, err = strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					stopAtHeight = 0
				}
			}

			goctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			eg, goctx := errgroup.WithContext(goctx)

			eg.Go(func() error {
				return loop(goctx, cancel, cmd, stopAtHeight)
			})
			eg.Go(func() error {
				return stopLoop(goctx, cancel)
			})
			if err := eg.Wait(); err != nil {
				return err
			}
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().StringP(tmcli.OutputFlag, "o", "text", "Output format (text|json)")
	return cmd
}

func loop(goctx context.Context, cancel context.CancelFunc, cmd *cobra.Command, stopAtHeight int64) error {
	var (
		blockEvent  <-chan ctypes.ResultEvent
		blockCancel context.CancelFunc

		txEvent  <-chan ctypes.ResultEvent
		txCancel context.CancelFunc
	)
	defer cancel()

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	if clientCtx.NodeURI == "" {
		return fmt.Errorf("Missing Tendermint Node URI")
	}

	rpcClient, _ := http.New(clientCtx.NodeURI, "/websocket")
	err = rpcClient.Start()

	// Retry if we don't have a connection
	for err != nil {
		fmt.Println("Could not connect to Tendermint node, retrying in 1 second.")
		time.Sleep(1 * time.Second)
		err = rpcClient.Start()
	}
	defer rpcClient.Stop()

	// subscribe to block events
	blockEvent, blockCancel = subscribe(goctx, rpcClient, blockQuery)
	defer blockCancel()

	// subscribe to tx events
	txEvent, txCancel = subscribe(goctx, rpcClient, txQuery)
	if err != nil {
		return err
	}

	defer txCancel()

	// Loop on received messages.
	for {
		select {
		case block := <-blockEvent:
			// run the custom block handler
			if err = instance.handleBlock(cmd, block); err != nil {
				fmt.Println("There was an error handling a new block", err)
			}

			// cancel loop if we reach endBlockHeight (for testing)
			if blockHeight :=
				block.Data.(tendermint.EventDataNewBlock).Block.Header.Height; stopAtHeight != 0 && stopAtHeight <= blockHeight {
				cancel()
			}

		case txEvent := <-txEvent:
			// run the custom tx handler
			if err = instance.handleTx(cmd, txEvent); err != nil {
				fmt.Println("There was an error handling a tx event", err)
				// return err
			}
		case <-goctx.Done():
			return nil
		}
	}
}

// subscribe returns channel of events from the oracle chain given a query (TX or BLOCK)
func subscribe(ctx context.Context, clnt *http.HTTP, query string) (<-chan ctypes.ResultEvent, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ch, err := clnt.Subscribe(ctx, fmt.Sprintf("oracle-feeder"), query)

	for err != nil {
		fmt.Println("Could not subscribe to Tendermint event, retrying in 1 second.")
		time.Sleep(1 * time.Second)
		ch, err = clnt.Subscribe(ctx, fmt.Sprintf("oracle-feeder"), query)
	}

	return ch, cancel
}

// stopLoop waits for a SIGINT or SIGTERM and returns an error
func stopLoop(ctx context.Context, cancel context.CancelFunc) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-sigCh:
			close(sigCh)
			cancel()
			fmt.Printf("Exiting feeder loop, recieved stop signal(%s)", sig.String())
			return nil
		case <-ctx.Done():
			close(sigCh)
			return nil
		}
	}
}
