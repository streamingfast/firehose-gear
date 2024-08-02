package block

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/streamingfast/substreams-gear/generated/convert1420"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"
	v1 "github.com/streamingfast/substreams-gear/pb/sf/substreams/gear/type/v1"
	"go.uber.org/zap"
)

type Decoder struct {
	callRegistry registry.CallRegistry
	logger       *zap.Logger
}

func NewDecoder(logger *zap.Logger) *Decoder {
	return &Decoder{
		logger: logger,
	}
}

func (d *Decoder) Decoded(block *pbgear.Block) (*v1.Block, error) {

	decodedExtrinsics, err := decodeExtrinsics(d.callRegistry, block.Extrinsics, d.logger)
	if err != nil {
		return nil, fmt.Errorf("failed to decode extrinsics: %w", err)
	}

	return &v1.Block{
		Number:        block.Number,
		Hash:          block.Hash,
		Header:        block.Header,
		DigestItems:   block.DigestItems,
		Justification: block.Justification,
		Extrinsics:    decodedExtrinsics,
		Events:        nil,
	}, nil
}

var versionFuncMap map[string]map[string]reflect.Value

func init() {
	versionFuncMap = map[string]map[string]reflect.Value{
		"1420": convert1420.FuncMap,
	}
}

func decodeExtrinsics(version string, callRegistry registry.CallRegistry, extrinsics []*pbgear.Extrinsic, logger *zap.Logger) ([]*v1.Extrinsic, error) {
	var decodedExtrinsics []*v1.Extrinsic
	for i, extrinsic := range extrinsics {
		logger.Info("decoding extrinsic", zap.Int("index", i), zap.Uint32("section", extrinsic.Method.CallIndex.SectionIndex), zap.Uint32("method", extrinsic.Method.CallIndex.MethodIndex))
		callName, decodedFields, err := decodeCallExtrinsics(callRegistry, extrinsic)
		if err != nil {
			return nil, fmt.Errorf("failed to decode extrinsic: %w", err)
		}

		parts := strings.Split(callName, ".")
		pallet := parts[0]
		call := parts[1]
		call = stringy.New(call).PascalCase().Get()
		structName := pallet + "_" + call + "Call"
		funcName := "To_" + structName

		funcMap := versionFuncMap[]
		if fn, found := funcMap[funcName]; found {
			o := fn.Call([]reflect.Value{reflect.ValueOf(decodedFields)})

			e := o[0].Interface().(*v1.Extrinsic)
			decodedExtrinsics = append(decodedExtrinsics, e)

		} else {
			panic(fmt.Sprintf("unknown extrinsic call: %s", callName))
		}
	}
	return decodedExtrinsics, nil
}

func decodeCallExtrinsics(callRegistry registry.CallRegistry, extrinsic *pbgear.Extrinsic) (string, registry.DecodedFields, error) {
	callIndex := extrinsic.Method.CallIndex
	args := extrinsic.Method.Args

	callDecoder, found := callRegistry[toCallIndex(callIndex)]
	if !found {
		return "", nil, fmt.Errorf("failed to get call decoder for call at index %d %d", callIndex.SectionIndex, callIndex.MethodIndex)
	}

	if args != nil {
		decoder := scale.NewDecoder(bytes.NewReader(args))
		callFields, err := callDecoder.Decode(decoder)
		if err != nil {
			return "", nil, fmt.Errorf("failed to decode call: %w", err)
		}
		return callDecoder.Name, callFields, nil
	}

	return callDecoder.Name, nil, nil
}

func toCallIndex(ci *pbgear.CallIndex) types.CallIndex {
	return types.CallIndex{
		SectionIndex: uint8(ci.SectionIndex),
		MethodIndex:  uint8(ci.MethodIndex),
	}
}
