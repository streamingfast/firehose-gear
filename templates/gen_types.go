package gen_types

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

//	type regEntry struct{
//	   fn func(fields []*registry.DecodedField) any
//	   returnType reflect.Type
//	}
//
// var reg map[int64]*regEntry
//
//	func init() {
//		reg = make(map[int64]*regEntry)
//	       reg[0] = &regEntry{
//	           fn: to_SpCoreCryptoAccountId32,
//	       }
//	}
func To_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) any {
	out := &pbgear.Balances_TransferKeepAliveCall{}
	f := fields[0].Value.(registry.DecodedFields)

	out.Dest = to_Balances_Dest([]*registry.DecodedField(f))
	//out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

func to_Balances_Dest(fields []*registry.DecodedField) *pbgear.Balances_Dest {
	out := &pbgear.Balances_Dest{}

	//===========
	switch {

	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Balances_Dest_Id{
			Id: to_Balances_Id(fields),
		} //Passthrough fields...
	}
	//============

	return out
}

func to_Balances_Id(fields []*registry.DecodedField) *pbgear.Balances_Id {
	out := &pbgear.Balances_Id{}
	out.Value_0 = to_SpCoreCryptoAccountId32(fields)
	return out
}

func to_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.SpCoreCryptoAccountId32 {
	out := &pbgear.SpCoreCryptoAccountId32{}
	f := fields[0].Value.(registry.DecodedFields)
	out.Value_0 = to_repeated_uint32(f)
	return out
}

func to_repeated_uint32(fields []*registry.DecodedField) []uint32 {
	out := make([]uint32, len(fields))

	data := fields[0].Value.([]interface{})
	for _, f := range data {
		u := uint8(f.(types.U8))
		out = append(out, uint32(u))
	}
	return out
}

func matchFields(fields []*registry.DecodedField, ids []int64) bool {
	if len(fields) != len(ids) {
		return false
	}
	for i, f := range fields {
		if ids[i] != f.LookupIndex {
			return false
		}
	}
	return true
}
