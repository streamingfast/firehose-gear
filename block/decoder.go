package block

import (
	"bytes"
	"fmt"
	"reflect"

	// "github.com/streamingfast/substreams-gear/generated/convert100"
	// "github.com/streamingfast/substreams-gear/generated/convert1000"
	// "github.com/streamingfast/substreams-gear/generated/convert1010"
	// "github.com/streamingfast/substreams-gear/generated/convert1020"
	// "github.com/streamingfast/substreams-gear/generated/convert1030"
	// "github.com/streamingfast/substreams-gear/generated/convert1040"
	// "github.com/streamingfast/substreams-gear/generated/convert1050"
	// "github.com/streamingfast/substreams-gear/generated/convert1110"
	// "github.com/streamingfast/substreams-gear/generated/convert120"
	// "github.com/streamingfast/substreams-gear/generated/convert1200"
	// "github.com/streamingfast/substreams-gear/generated/convert1210"
	// "github.com/streamingfast/substreams-gear/generated/convert130"
	// "github.com/streamingfast/substreams-gear/generated/convert1300"
	// "github.com/streamingfast/substreams-gear/generated/convert1310"
	// "github.com/streamingfast/substreams-gear/generated/convert140"
	// "github.com/streamingfast/substreams-gear/generated/convert1400"
	// "github.com/streamingfast/substreams-gear/generated/convert1410"
	// "github.com/streamingfast/substreams-gear/generated/convert1420"
	// "github.com/streamingfast/substreams-gear/generated/convert210"
	// "github.com/streamingfast/substreams-gear/generated/convert310"
	// "github.com/streamingfast/substreams-gear/generated/convert320"
	// "github.com/streamingfast/substreams-gear/generated/convert330"
	// "github.com/streamingfast/substreams-gear/generated/convert340"
	// "github.com/streamingfast/substreams-gear/generated/convert350"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"

	// v1 "github.com/streamingfast/substreams-gear/pb/sf/substreams/gear/type/v1"
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

func (d *Decoder) Decoded(block *pbgear.Block) (*pbgear.Block, error) {
	// func (d *Decoder) Decoded(block *pbgear.Block) (*v1.Block, error) {
	// decodedExtrinsics, err := decodeExtrinsics("", d.callRegistry, block.Extrinsics, d.logger)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to decode extrinsics: %w", err)
	// }

	// _ = decodedExtrinsics

	return nil, nil

	// return &v1.Block{
	// 	Number:        block.Number,
	// 	Hash:          block.Hash,
	// 	Header:        block.Header,
	// 	DigestItems:   block.DigestItems,
	// 	Justification: block.Justification,
	// 	Extrinsics:    decodedExtrinsics,
	// 	Events:        nil,
	// }, nil
}

var versionFuncMap map[string]map[string]reflect.Value

func init() {
	versionFuncMap = map[string]map[string]reflect.Value{
		// "1420": convert1420.FuncMap,
		// "1410": convert1410.FuncMap,
		// "1400": convert1400.FuncMap,
		// "1310": convert1310.FuncMap,
		// "1300": convert1300.FuncMap,
		// "1210": convert1210.FuncMap,
		// "1200": convert1200.FuncMap,
		// "1110": convert1110.FuncMap,
		// "1050": convert1050.FuncMap,
		// "1040": convert1040.FuncMap,
		// "1030": convert1030.FuncMap,
		// "1020": convert1020.FuncMap,
		// "1010": convert1010.FuncMap,
		// "1000": convert1000.FuncMap,
		// "350":  convert350.FuncMap,
		// "340":  convert340.FuncMap,
		// "330":  convert330.FuncMap,
		// "320":  convert320.FuncMap,
		// "310":  convert310.FuncMap,
		// "210":  convert210.FuncMap,
		// "140":  convert140.FuncMap,
		// "130":  convert130.FuncMap,
		// "120":  convert120.FuncMap,
		// "100":  convert100.FuncMap,
	}
}

// func decodeExtrinsics(version string, callRegistry registry.CallRegistry, extrinsics []*pbgear.Extrinsic, logger *zap.Logger) ([]*v1.Extrinsic, error) {
// 	var decodedExtrinsics []*v1.Extrinsic
// 	for i, extrinsic := range extrinsics {
// 		logger.Info("decoding extrinsic", zap.Int("index", i), zap.Uint32("section", extrinsic.Method.CallIndex.SectionIndex), zap.Uint32("method", extrinsic.Method.CallIndex.MethodIndex))
// 		callName, decodedFields, err := decodeCallExtrinsics(callRegistry, extrinsic)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to decode extrinsic: %w", err)
// 		}
// 		_ = decodedFields

// 		parts := strings.Split(callName, ".")
// 		pallet := parts[0]
// 		call := parts[1]
// 		call = stringy.New(call).PascalCase().Get()
// 		structName := pallet + "_" + call + "Call"
// 		funcName := "To_" + structName
// 		_ = funcName
// 		// funcMap := versionFuncMap[]
// 		// if fn, found := funcMap[funcName]; found {
// 		// 	o := fn.Call([]reflect.Value{reflect.ValueOf(decodedFields)})

// 		// 	e := o[0].Interface().(*v1.Extrinsic)
// 		// 	decodedExtrinsics = append(decodedExtrinsics, e)

// 		// } else {
// 		// 	panic(fmt.Sprintf("unknown extrinsic call: %s", callName))
// 		// }
// 	}
// 	return decodedExtrinsics, nil
// }

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
