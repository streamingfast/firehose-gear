package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/firehose-core/blockpoller"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewFetchCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch <first-streamable-block>",
		Short: "Fetch data from rpc endpoint",
		RunE:  fetchRunE(logger, tracer),
	}

	cmd.Flags().StringArray("endpoints", []string{"https://vara-mainnet.public.blastapi.io"}, "List of endpoints to use to fetch different method calls")
	cmd.Flags().String("state-dir", "/data", "location to store the cursor.json")
	cmd.Flags().Duration("interval-between-fetch", 0, "interval between fetch")
	cmd.Flags().Duration("latest-block-retry-interval", time.Second, "interval between fetch when latest block is not available yet")
	cmd.Flags().Int("block-fetch-batch-size", 10, "Number of blocks to fetch in a single batch")

	return cmd
}

func fetchRunE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		rpcEndpoints := sflags.MustGetStringArray(cmd, "endpoints")
		stateDir := sflags.MustGetString(cmd, "state-dir")
		startBlock, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("unable to parse first streamable block %d: %w", startBlock, err)
		}

		fetchInterval := sflags.MustGetDuration(cmd, "interval-between-fetch")

		logger.Info(
			"launching firehose-gear poller",
			zap.Strings("rpc_endpoint", rpcEndpoints),
			zap.String("state_dir", stateDir),
			zap.Uint64("first_streamable_block", startBlock),
			zap.Duration("interval_between_fetch", fetchInterval),
			zap.Duration("latest_block_retry_interval", sflags.MustGetDuration(cmd, "latest-block-retry-interval")),
		)

		latestBlockRetryInterval := sflags.MustGetDuration(cmd, "latest-block-retry-interval")

		gearClients := firecoreRPC.NewClients[*rpc.Client]()

		for _, rpcEndpoint := range rpcEndpoints {
			gearClient := rpc.NewClient(rpcEndpoint)
			gearClients.Add(gearClient)
		}

		rpcFetcher := rpc.NewFetcher(gearClients, fetchInterval, latestBlockRetryInterval, logger)

		poller := blockpoller.New(
			rpcFetcher,
			blockpoller.NewFireBlockHandler("type.googleapis.com/sf.gear.type.v1.Block"),
			blockpoller.WithStoringState(stateDir),
			blockpoller.WithLogger(logger),
		)

		err = poller.Run(ctx, startBlock, sflags.MustGetInt(cmd, "block-fetch-batch-size"))
		if err != nil {
			return fmt.Errorf("running poller: %w", err)
		}

		return nil
	}
}
