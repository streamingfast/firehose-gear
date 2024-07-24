package gen_types

import (
    "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
    pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)



func to_Balances_Address20(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
}

func to_repeated_uint32(fields []*registry.DecodedField) []uint32 {
	out := make([]uint32, len(fields))
	for _, f := range fields {
		out = append(out, f.Value.(uint32))
	}
	return out
}  
 

func to_Balances_Address32(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
} 
 

func to_Balances_CompactString(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_CompactString{}
        out.Value = fields[0].Value.(string)
    return out
}  
 

func to_Balances_Dest(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Dest{}
        

    //===========
    switch {
    
    case matchFields(fields, []int64{0}):
    out.Value = to_Balances_Id(fields) //Passthrough fields...
    
    case matchFields(fields, []int64{95}):
    out.Value = to_Balances_Index(fields) //Passthrough fields...
    
    case matchFields(fields, []int64{13}):
    out.Value = to_Balances_Raw(fields) //Passthrough fields...
    
    case matchFields(fields, []int64{1}):
    out.Value = to_Balances_Address32(fields) //Passthrough fields...
    
    case matchFields(fields, []int64{96}):
    out.Value = to_Balances_Address20(fields) //Passthrough fields...
    }
    //============
    return out
}

func to_oneof_Balances_Value(fields []*registry.DecodedField) any {IdIndexRawAddress32Address20
    var out *&pbgear.Balances_Value
    out.Value = to_oneof_Balances_Value(fields) //Passthrough fields...
	return out
}  
 

func to_Balances_Id(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Id{}
        out.Value_0 = to_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
 

func to_Balances_Index(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Index{}
        out.Value_0 = to_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
    return out
}  
 

func to_Balances_Raw(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
} 
 

func to_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) any {
    out := &pbgear.Balances_TransferKeepAliveCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
        out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
    return out
}  
  
 

func to_CompactTupleNull(fields []*registry.DecodedField) any {
    out := &pbgear.CompactTupleNull{}
        out.Value = to_TupleNull(fields[0].Value.([]*registry.DecodedField))
    return out
}  
 

func to_SpCoreCryptoAccountId32(fields []*registry.DecodedField) any {
    out := &pbgear.SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
} 
 

func to_TupleNull(fields []*registry.DecodedField) any {
    out := &pbgear.TupleNull{}
    return out
}  



func matchFields(fields []*registry.DecodedField, ids []int64) bool {
	if len(fields) != len(ids) {
		return false
	}
	for i, f := range fields {
		if ids[i] != f.LookupIndex{
			return false
		}
	}
	return true
}