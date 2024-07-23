package gen_types

import (
    "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
    pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)
                                         func to_Balances_Address20(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
    out := &pbgear.Balances_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Address32(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
    out := &pbgear.Balances_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_CompactString(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
    out := &pbgear.Balances_CompactString{}
        out.Value = fields[0].Value.(string)
    return out
}  
  func to_Balances_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
    out := &pbgear.Balances_CompactTupleNull{}
        out.Value = to_Balances_TupleNull(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Dest(fields []*registry.DecodedField) *pbgear.Balances_Dest {
    out := &pbgear.Balances_Dest{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Dest_Id {
            Id: to_Balances_Id(field)
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Dest_Index {
            Index: to_Balances_Index(field)
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Dest_Raw {
            Raw: to_Balances_Raw(field)
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Dest_Address32 {
            Address32: to_Balances_Address32(field)
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Dest_Address20 {
            Address20: to_Balances_Address20(field)
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  func to_Balances_ForceSetBalanceCall(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
    out := &pbgear.Balances_ForceSetBalanceCall{}
        out.Who = to_Balances_Who(fields[0].Value.([]*registry.DecodedField))
        out.NewFree = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
    return out
}  
  
  func to_Balances_ForceTransferCall(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
    out := &pbgear.Balances_ForceTransferCall{}
        out.Source = to_Balances_Source(fields[0].Value.([]*registry.DecodedField))
        out.Dest = to_Balances_Dest(fields[1].Value.([]*registry.DecodedField))
        out.Value = to_Balances_CompactString(fields[2].Value.([]*registry.DecodedField))
    return out
}  
  
  
  func to_Balances_ForceUnreserveCall(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
    out := &pbgear.Balances_ForceUnreserveCall{}
        out.Who = to_Balances_Who(fields[0].Value.([]*registry.DecodedField))
        out.Amount = fields[1].Value.(string)
    return out
}  
  
  func to_Balances_Id(fields []*registry.DecodedField) *pbgear.Balances_Id {
    out := &pbgear.Balances_Id{}
        out.Value_0 = to_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Index(fields []*registry.DecodedField) *pbgear.Balances_Index {
    out := &pbgear.Balances_Index{}
        out.Value_0 = to_Balances_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Raw(fields []*registry.DecodedField) *pbgear.Balances_Raw {
    out := &pbgear.Balances_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Source(fields []*registry.DecodedField) *pbgear.Balances_Source {
    out := &pbgear.Balances_Source{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Source_Id {
            Id: to_Balances_Id(field)
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Source_Index {
            Index: to_Balances_Index(field)
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Source_Raw {
            Raw: to_Balances_Raw(field)
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Source_Address32 {
            Address32: to_Balances_Address32(field)
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Source_Address20 {
            Address20: to_Balances_Address20(field)
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  func to_Balances_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
    out := &pbgear.Balances_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_TransferAllCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
    out := &pbgear.Balances_TransferAllCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
        out.KeepAlive = fields[1].Value.(bool)
    return out
}  
  
  func to_Balances_TransferAllowDeathCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
    out := &pbgear.Balances_TransferAllowDeathCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
        out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
    return out
}  
  
  func to_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
    out := &pbgear.Balances_TransferKeepAliveCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
        out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
    return out
}  
  
  func to_Balances_TupleNull(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
    out := &pbgear.Balances_TupleNull{}
    return out
}  func to_Balances_UpgradeAccountsCall(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
    out := &pbgear.Balances_UpgradeAccountsCall{}
        out.Who = to_repeated_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
    return out
}  
  func to_Balances_Who(fields []*registry.DecodedField) *pbgear.Balances_Who {
    out := &pbgear.Balances_Who{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Who_Id {
            Id: to_Balances_Id(field)
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Who_Index {
            Index: to_Balances_Index(field)
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Who_Raw {
            Raw: to_Balances_Raw(field)
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Who_Address32 {
            Address32: to_Balances_Address32(field)
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Who_Address20 {
            Address20: to_Balances_Address20(field)
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        

