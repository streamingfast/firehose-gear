package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/mostynb/go-grpc-compression/zstd"
	"google.golang.org/grpc"

	"github.com/streamingfast/cli/sflags"

	"github.com/streamingfast/firehose-core/firehose/client"

	"github.com/streamingfast/firehose-gear/block"

	"github.com/spf13/cobra"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewDecoderCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode-block",
		Short: "Consume blocks from firehose and decode them",
		RunE:  decodeRunE(logger, tracer),
	}
	cmd.Flags().Int64("first-streamable-block", 0, "first block to decode")
	cmd.Flags().BoolP("plaintext", "p", false, "Use plaintext connection to Firehose")
	cmd.Flags().BoolP("insecure", "k", false, "Use SSL connection to Firehose but skip SSL certificate validation")

	return cmd
}

func decodeRunE(logger *zap.Logger, tracer logging.Tracer) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		firehoseEndpoint := args[0]
		plaintext := sflags.MustGetBool(cmd, "plaintext")
		insecure := sflags.MustGetBool(cmd, "insecure")
		firstStreamableBlock := sflags.MustGetInt64(cmd, "first-streamable-block")

		decoder := block.NewDecoder(logger)
		jwt := os.Getenv("FIREHOSE_API_TOKEN")

		firehoseClient, connClose, grpcCallOpts, err := client.NewFirehoseClient(firehoseEndpoint, jwt, "", insecure, plaintext)
		if err != nil {
			return fmt.Errorf("creating firehose client: %w", err)
		}
		defer connClose()

		grpcCallOpts = append(grpcCallOpts, grpc.UseCompressor(zstd.Name))

		//todo: load cursor here

		request := &pbfirehose.Request{
			StartBlockNum:   firstStreamableBlock,
			StopBlockNum:    0,
			FinalBlocksOnly: false,
		}

		stream, err := firehoseClient.Blocks(ctx, request, grpcCallOpts...)
		if err != nil {
			return fmt.Errorf("unable to start blocks stream: %w", err)
		}
		logger.Info("starting blocks stream")
		lastBlockNum := uint64(firstStreamableBlock)
		for {
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("stream error while receiving: %w", err)
			}

			logger.Info("got response")
			block := &pbgear.Block{}
			if err := anypb.UnmarshalTo(response.Block, block, proto.UnmarshalOptions{}); err != nil {
				return fmt.Errorf("unmarshalling anypb: %w", err)
			}
			logger.Info("got block", zap.Uint64("block_num", block.Number))

			if len(block.Header.UpdatedMetadata) > 0 {
				logger.Info("Updating metadata", zap.Uint64("block_num", block.Number), zap.Uint32("version", block.Header.SpecVersion))
				md := LoadMetadata(block.Header.UpdatedMetadata)
				factory := registry.NewFactory()
				callRegistry, err := factory.CreateCallRegistry(md)
				if err != nil {
					return fmt.Errorf("creating call registry: %w", err)
				}
				decoder.SetCallRegistry(callRegistry)
				//todo: save has last metadata seen and reload it at startup
			}

			if lastBlockNum != 0 && (block.Number-lastBlockNum) != 1 {
				panic(fmt.Sprintf("incorrect number of blocks sequence. Expected  %d got %d", lastBlockNum+1, block.Number))
			}
			lastBlockNum = block.Number

			decodedBlock, err := decoder.Decoded(block)
			if err != nil {
				return fmt.Errorf("decoding block %d %s: %w", block.Number, hex.EncodeToString(block.Hash), err)
			}

			j, err := json.Marshal(decodedBlock)
			if err != nil {
				return fmt.Errorf("marshalling block %d: %w", block.Number, err)
			}
			fmt.Println(string(j))
			//todo: print DM log
			//todo: save Cursor
		}
		return nil
	}
}

func LoadMetadata(data []byte) *types.Metadata {
	metadata := &types.Metadata{}
	err := codec.Decode(data, metadata)
	if err != nil {
		panic(err)
	}
	return metadata
}
