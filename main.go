package main

//import (
//	"encoding/json"
//	"fmt"
//	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
//	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
//	exec "github.com/centrifuge/go-substrate-rpc-client/v4/registry/exec"
//	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
//	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
//	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
//	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
//	"github.com/gobeam/stringy"
//	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
//	"log"
//	"strings"
//)
//
//func main() {
//	blockHash, err := types.NewHashFromHexString("0xad1ce65fb2dde8ee826f2b25999f873f92815224300cd3b747031123f4cb6611")
//	if err != nil {
//		log.Fatalf("Failed to create block hash: %v", err)
//	}
//
//	url := "https://vara-mainnet.public.blastapi.io" // Replace with the actual URL of your Gear Tech node
//	api, err := gsrpc.NewSubstrateAPI(url)
//	if err != nil {
//		log.Fatal("Error creating API instance")
//	}
//
//	block, err := api.RPC.Chain.GetBlock(blockHash)
//	if err != nil {
//		log.Fatalf("Failed to get block: %v", err)
//	}
//
//	j, err := json.MarshalIndent(block, "", "  ")
//	if err != nil {
//		log.Fatalf("Failed to marshal block: %v", err)
//	}
//	_ = j
//	// fmt.Println("block: ", string(j))
//
//	chain := generic.NewDefaultChain(api.Client)
//	factory := registry.NewFactory()
//
//	extrinsicParser := parser.NewDefaultExtrinsicParser()
//	extrinsicRetriever, err := retriever.NewDefaultExtrinsicRetriever(
//		extrinsicParser,
//		chain,
//		api.RPC.State,
//		factory,
//		exec.NewRetryableExecutor[*generic.DefaultGenericSignedBlock](),
//		exec.NewRetryableExecutor[[]*parser.DefaultExtrinsic](),
//	)
//
//	if err != nil {
//		log.Fatalf("Error creating extrinsic retriever: %v", err)
//	}
//
//	extrinsics, err := extrinsicRetriever.GetExtrinsics(blockHash)
//	if err != nil {
//		log.Fatalf("Failed to get extrinsics: %v", err)
//	}
//
//	metadata, err := api.RPC.State.GetMetadata(blockHash)
//	//if err != nil {
//	//	log.Fatalf("Failed to get metadata: %v", err)
//	//}
//	//
//	for _, extrinsic := range extrinsics {
//		parts := strings.Split(extrinsic.Name, ".")
//		pallet := parts[0]
//		call := parts[1]
//		call = stringy.New(call).PascalCase().Get()
//		structName := pallet + "_" + call + "_Call"
//
//		switch structName {
//		case "Timestamp_Set_Call":
//			call := from_Timestamp_Set_Call(extrinsic.CallFields)
//			fmt.Println(call)
//		case "Balances_TransferKeepAlive_Call":
//			call := from_Balances_Transfer_Keep_Alive_Call(extrinsic.CallFields)
//			fmt.Println(call)
//		case "Gear_Run_Call":
//			from_Gear_Run_Call(extrinsic.CallFields)
//		}
//	}
//}
//
//func from_Balances_Transfer_Keep_Alive_Call(callFields []*registry.DecodedField) *pbgear.Balances_TransferKeepAlive_Call {
//	return &pbgear.Balances_TransferKeepAlive_Call{
//		Dest:  from_balance_dest(callFields[0]),
//		Value: nil,
//	}
//}
//
//func from_balance_dest(in *registry.DecodedField) *pbgear.BalancesDest {
//	out := &pbgear.BalancesDest{}
//	switch in.Name {
//	case "sp_runtime.multiaddress.MultiAddress.dest":
//		out.Value = from_BalancesDest_Balances_Id(in.Value.(registry.DecodedFields)[0])
//	}
//
//	return out
//}
//
//func from_BalancesDest_Balances_Id(in *registry.DecodedField) *pbgear.BalancesDest_Balances_Id {
//	return &pbgear.BalancesDest_Balances_Id{
//		Balances_Id: &pbgear.Balances_Id{
//			Value_0: from_sp_runtime_multi_address_multi_address(in.Value.(registry.DecodedFields)[0]),
//		},
//	}
//}
//
//func from_sp_runtime_multi_address_multi_address(in *registry.DecodedField) *pbgear.SpCoreCrypto_AccountId32 {
//	return &pbgear.SpCoreCrypto_AccountId32{Value_0: from_array_u8(in)}
//}
//
//func from_array_u8(in *registry.DecodedField) []uint32 {
//	out := make([]uint32, 0)
//	for _, value := range in.Value.([]interface{}) {
//		v := value.(types.U8)
//		out = append(out, uint32(v))
//	}
//	return out
//}
//
//func from_Timestamp_Set_Call(callFields []*registry.DecodedField) *pbgear.Timestamp_Set_Call {
//	return &pbgear.Timestamp_Set_Call{
//		Now: from_Compact_uint64(callFields[0]),
//	}
//}
//
//func from_Compact_uint64(in *registry.DecodedField) *pbgear.CompactUint64 {
//	v := in.Value.(types.UCompact)
//	return &pbgear.CompactUint64{
//		Value: from_uint64(v),
//	}
//}
//
//func from_uint64(in types.UCompact) uint64 {
//	return uint64(in.Int64())
//}
//
//func from_Gear_Run_Call(callFields []*registry.DecodedField) *pbgear.Gear_Run_Call {
//	return &pbgear.Gear_Run_Call{
//		MaxGas: from_optional_uint64(callFields[0]),
//	}
//}
//
//func from_optional_uint64(field *registry.DecodedField) *uint64 {
//	var out *uint64
//
//	if v, ok := field.Value.(uint8); ok {
//		if v != 0 {
//			panic("expected none")
//		}
//	} else {
//		out = field.Value.(*uint64)
//	}
//	return out
//}
