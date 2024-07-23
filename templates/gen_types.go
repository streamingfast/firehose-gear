package gen_types

import (
    "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
    pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func to_Babe_AllowedSlots(fields []*registry.DecodedField) *pbgear.Babe_AllowedSlots {
    out := &pbgear.Babe_AllowedSlots{}
        
    switch Babe_Value {
    case Babe_PrimarySlots:
        out.Value = &pbgear.Babe_AllowedSlots_PrimarySlots {
            PrimarySlots: to_Babe_PrimarySlots(field),
        }
    case Babe_PrimaryAndSecondaryPlainSlots:
        out.Value = &pbgear.Babe_AllowedSlots_PrimaryAndSecondaryPlainSlots {
            PrimaryAndSecondaryPlainSlots: to_Babe_PrimaryAndSecondaryPlainSlots(field),
        }
    case Babe_PrimaryAndSecondaryVRFSlots:
        out.Value = &pbgear.Babe_AllowedSlots_PrimaryAndSecondaryVRFSlots {
            PrimaryAndSecondaryVRFSlots: to_Babe_PrimaryAndSecondaryVRFSlots(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Babe_BabeTrieNodesList(fields []*registry.DecodedField) *pbgear.Babe_BabeTrieNodesList {
    out := &pbgear.Babe_BabeTrieNodesList{}
        out.TrieNodes = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_CompactUint32(fields []*registry.DecodedField) *pbgear.Babe_CompactUint32 {
    out := &pbgear.Babe_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Babe_Config(fields []*registry.DecodedField) *pbgear.Babe_Config {
    out := &pbgear.Babe_Config{}
        
    switch Babe_Value {
    case Babe_V1:
        out.Value = &pbgear.Babe_Config_V1 {
            V1: to_Babe_V1(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Babe_Consensus(fields []*registry.DecodedField) *pbgear.Babe_Consensus {
    out := &pbgear.Babe_Consensus{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Babe_Logs(fields []*registry.DecodedField) *pbgear.Babe_Logs {
    out := &pbgear.Babe_Logs{}
        
    switch Babe_Value {
    case Babe_PreRuntime:
        out.Value = &pbgear.Babe_Logs_PreRuntime {
            PreRuntime: to_Babe_PreRuntime(field),
        }
    case Babe_Consensus:
        out.Value = &pbgear.Babe_Logs_Consensus {
            Consensus: to_Babe_Consensus(field),
        }
    case Babe_Seal:
        out.Value = &pbgear.Babe_Logs_Seal {
            Seal: to_Babe_Seal(field),
        }
    case Babe_Other:
        out.Value = &pbgear.Babe_Logs_Other {
            Other: to_Babe_Other(field),
        }
    case Babe_RuntimeEnvironmentUpdated:
        out.Value = &pbgear.Babe_Logs_RuntimeEnvironmentUpdated {
            RuntimeEnvironmentUpdated: to_Babe_RuntimeEnvironmentUpdated(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Babe_Other(fields []*registry.DecodedField) *pbgear.Babe_Other {
    out := &pbgear.Babe_Other{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_PlanConfigChangeCall(fields []*registry.DecodedField) *pbgear.Babe_PlanConfigChangeCall {
    out := &pbgear.Babe_PlanConfigChangeCall{}
        out.Config = to_Babe_Config(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_PreRuntime(fields []*registry.DecodedField) *pbgear.Babe_PreRuntime {
    out := &pbgear.Babe_PreRuntime{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Babe_PrimaryAndSecondaryPlainSlots(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
    out := &pbgear.Babe_PrimaryAndSecondaryPlainSlots{}
    return out
} 
func to_Babe_PrimaryAndSecondaryVRFSlots(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryVRFSlots {
    out := &pbgear.Babe_PrimaryAndSecondaryVRFSlots{}
    return out
} 
func to_Babe_PrimarySlots(fields []*registry.DecodedField) *pbgear.Babe_PrimarySlots {
    out := &pbgear.Babe_PrimarySlots{}
    return out
} 
func to_Babe_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
    out := &pbgear.Babe_PrimitiveTypesH256{}
        out.ParentHash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationCall {
    out := &pbgear.Babe_ReportEquivocationCall{}
        out.EquivocationProof = to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
        out.KeyOwnerProof = to_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Babe_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationUnsignedCall {
    out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
        out.EquivocationProof = to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
        out.KeyOwnerProof = to_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Babe_RuntimeEnvironmentUpdated(fields []*registry.DecodedField) *pbgear.Babe_RuntimeEnvironmentUpdated {
    out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
    return out
} 
func to_Babe_Seal(fields []*registry.DecodedField) *pbgear.Babe_Seal {
    out := &pbgear.Babe_Seal{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Babe_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
    out := &pbgear.Babe_SpConsensusBabeAppPublic{}
        out.Offender = to_Babe_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_SpConsensusSlotsEquivocationProof(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
    out := &pbgear.Babe_SpConsensusSlotsEquivocationProof{}
        out.Offender = to_Babe_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField)),
        out.Slot = to_Babe_SpConsensusSlotsSlot(fields[1].Value.([]*registry.DecodedField)),
        out.FirstHeader = to_Babe_SpRuntimeGenericHeaderHeader(fields[2].Value.([]*registry.DecodedField)),
        out.SecondHeader = to_Babe_SpRuntimeGenericHeaderHeader(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_Babe_SpConsensusSlotsSlot(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
    out := &pbgear.Babe_SpConsensusSlotsSlot{}
        out.Slot = fields[0].Value.(uint64),
    return out
}  
 
func to_Babe_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
    out := &pbgear.Babe_SpCoreSr25519Public{}
        out.Offender = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_SpRuntimeGenericDigestDigest(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
    out := &pbgear.Babe_SpRuntimeGenericDigestDigest{}
        out.Logs = to_repeated_Babe_SpRuntimeGenericDigestDigestItem(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_SpRuntimeGenericDigestDigestItem(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigestItem {
    out := &pbgear.Babe_SpRuntimeGenericDigestDigestItem{}
        out.Logs = to_Babe_Logs(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Babe_SpRuntimeGenericHeaderHeader(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
    out := &pbgear.Babe_SpRuntimeGenericHeaderHeader{}
        out.ParentHash = to_Babe_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
        out.Number = to_Babe_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
        out.StateRoot = to_Babe_PrimitiveTypesH256(fields[2].Value.([]*registry.DecodedField)),
        out.ExtrinsicsRoot = to_Babe_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField)),
        out.Digest = to_Babe_SpRuntimeGenericDigestDigest(fields[4].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
 
func to_Babe_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
    out := &pbgear.Babe_SpSessionMembershipProof{}
        out.Session = fields[0].Value.(uint32),
        out.TrieNodes = to_repeated_Babe_BabeTrieNodesList(fields[1].Value.([]*registry.DecodedField)),
        out.ValidatorCount = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Babe_TupleUint64Uint64(fields []*registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
    out := &pbgear.Babe_TupleUint64Uint64{}
        out.Value_0 = fields[0].Value.(uint64),
        out.Value_1 = fields[1].Value.(uint64),
    return out
}  
  
 
func to_Babe_V1(fields []*registry.DecodedField) *pbgear.Babe_V1 {
    out := &pbgear.Babe_V1{}
        out.C = to_Babe_TupleUint64Uint64(fields[0].Value.([]*registry.DecodedField)),
        out.AllowedSlots = to_Babe_AllowedSlots(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_BagsList_Address20(fields []*registry.DecodedField) *pbgear.BagsList_Address20 {
    out := &pbgear.BagsList_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_Address32(fields []*registry.DecodedField) *pbgear.BagsList_Address32 {
    out := &pbgear.BagsList_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_CompactTupleNull(fields []*registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
    out := &pbgear.BagsList_CompactTupleNull{}
        out.Value = to_BagsList_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_Dislocated(fields []*registry.DecodedField) *pbgear.BagsList_Dislocated {
    out := &pbgear.BagsList_Dislocated{}
        
    switch BagsList_Value {
    case BagsList_Id:
        out.Value = &pbgear.BagsList_Dislocated_Id {
            Id: to_BagsList_Id(field),
        }
    case BagsList_Index:
        out.Value = &pbgear.BagsList_Dislocated_Index {
            Index: to_BagsList_Index(field),
        }
    case BagsList_Raw:
        out.Value = &pbgear.BagsList_Dislocated_Raw {
            Raw: to_BagsList_Raw(field),
        }
    case BagsList_Address32:
        out.Value = &pbgear.BagsList_Dislocated_Address32 {
            Address32: to_BagsList_Address32(field),
        }
    case BagsList_Address20:
        out.Value = &pbgear.BagsList_Dislocated_Address20 {
            Address20: to_BagsList_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_BagsList_Heavier(fields []*registry.DecodedField) *pbgear.BagsList_Heavier {
    out := &pbgear.BagsList_Heavier{}
        
    switch BagsList_Value {
    case BagsList_Id:
        out.Value = &pbgear.BagsList_Heavier_Id {
            Id: to_BagsList_Id(field),
        }
    case BagsList_Index:
        out.Value = &pbgear.BagsList_Heavier_Index {
            Index: to_BagsList_Index(field),
        }
    case BagsList_Raw:
        out.Value = &pbgear.BagsList_Heavier_Raw {
            Raw: to_BagsList_Raw(field),
        }
    case BagsList_Address32:
        out.Value = &pbgear.BagsList_Heavier_Address32 {
            Address32: to_BagsList_Address32(field),
        }
    case BagsList_Address20:
        out.Value = &pbgear.BagsList_Heavier_Address20 {
            Address20: to_BagsList_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_BagsList_Id(fields []*registry.DecodedField) *pbgear.BagsList_Id {
    out := &pbgear.BagsList_Id{}
        out.Value_0 = to_BagsList_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_Index(fields []*registry.DecodedField) *pbgear.BagsList_Index {
    out := &pbgear.BagsList_Index{}
        out.Value_0 = to_BagsList_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_Lighter(fields []*registry.DecodedField) *pbgear.BagsList_Lighter {
    out := &pbgear.BagsList_Lighter{}
        
    switch BagsList_Value {
    case BagsList_Id:
        out.Value = &pbgear.BagsList_Lighter_Id {
            Id: to_BagsList_Id(field),
        }
    case BagsList_Index:
        out.Value = &pbgear.BagsList_Lighter_Index {
            Index: to_BagsList_Index(field),
        }
    case BagsList_Raw:
        out.Value = &pbgear.BagsList_Lighter_Raw {
            Raw: to_BagsList_Raw(field),
        }
    case BagsList_Address32:
        out.Value = &pbgear.BagsList_Lighter_Address32 {
            Address32: to_BagsList_Address32(field),
        }
    case BagsList_Address20:
        out.Value = &pbgear.BagsList_Lighter_Address20 {
            Address20: to_BagsList_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_BagsList_PutInFrontOfCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfCall {
    out := &pbgear.BagsList_PutInFrontOfCall{}
        out.Lighter = to_BagsList_Lighter(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_PutInFrontOfOtherCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfOtherCall {
    out := &pbgear.BagsList_PutInFrontOfOtherCall{}
        out.Heavier = to_BagsList_Heavier(fields[0].Value.([]*registry.DecodedField)),
        out.Lighter = to_BagsList_Lighter(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_BagsList_Raw(fields []*registry.DecodedField) *pbgear.BagsList_Raw {
    out := &pbgear.BagsList_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_RebagCall(fields []*registry.DecodedField) *pbgear.BagsList_RebagCall {
    out := &pbgear.BagsList_RebagCall{}
        out.Dislocated = to_BagsList_Dislocated(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.BagsList_SpCoreCryptoAccountId32 {
    out := &pbgear.BagsList_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_BagsList_TupleNull(fields []*registry.DecodedField) *pbgear.BagsList_TupleNull {
    out := &pbgear.BagsList_TupleNull{}
    return out
} 
func to_Balances_Address20(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
    out := &pbgear.Balances_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Address32(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
    out := &pbgear.Balances_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_CompactString(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
    out := &pbgear.Balances_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_Balances_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
    out := &pbgear.Balances_CompactTupleNull{}
        out.Value = to_Balances_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Dest(fields []*registry.DecodedField) *pbgear.Balances_Dest {
    out := &pbgear.Balances_Dest{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Dest_Id {
            Id: to_Balances_Id(field),
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Dest_Index {
            Index: to_Balances_Index(field),
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Dest_Raw {
            Raw: to_Balances_Raw(field),
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Dest_Address32 {
            Address32: to_Balances_Address32(field),
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Dest_Address20 {
            Address20: to_Balances_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Balances_ForceSetBalanceCall(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
    out := &pbgear.Balances_ForceSetBalanceCall{}
        out.Who = to_Balances_Who(fields[0].Value.([]*registry.DecodedField)),
        out.NewFree = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Balances_ForceTransferCall(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
    out := &pbgear.Balances_ForceTransferCall{}
        out.Source = to_Balances_Source(fields[0].Value.([]*registry.DecodedField)),
        out.Dest = to_Balances_Dest(fields[1].Value.([]*registry.DecodedField)),
        out.Value = to_Balances_CompactString(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_Balances_ForceUnreserveCall(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
    out := &pbgear.Balances_ForceUnreserveCall{}
        out.Who = to_Balances_Who(fields[0].Value.([]*registry.DecodedField)),
        out.Amount = fields[1].Value.(string),
    return out
}  
  
 
func to_Balances_Id(fields []*registry.DecodedField) *pbgear.Balances_Id {
    out := &pbgear.Balances_Id{}
        out.Value_0 = to_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Index(fields []*registry.DecodedField) *pbgear.Balances_Index {
    out := &pbgear.Balances_Index{}
        out.Value_0 = to_Balances_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Raw(fields []*registry.DecodedField) *pbgear.Balances_Raw {
    out := &pbgear.Balances_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Source(fields []*registry.DecodedField) *pbgear.Balances_Source {
    out := &pbgear.Balances_Source{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Source_Id {
            Id: to_Balances_Id(field),
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Source_Index {
            Index: to_Balances_Index(field),
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Source_Raw {
            Raw: to_Balances_Raw(field),
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Source_Address32 {
            Address32: to_Balances_Address32(field),
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Source_Address20 {
            Address20: to_Balances_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Balances_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
    out := &pbgear.Balances_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_TransferAllCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
    out := &pbgear.Balances_TransferAllCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
        out.KeepAlive = fields[1].Value.(bool),
    return out
}  
  
 
func to_Balances_TransferAllowDeathCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
    out := &pbgear.Balances_TransferAllowDeathCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
        out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
    out := &pbgear.Balances_TransferKeepAliveCall{}
        out.Dest = to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
        out.Value = to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Balances_TupleNull(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
    out := &pbgear.Balances_TupleNull{}
    return out
} 
func to_Balances_UpgradeAccountsCall(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
    out := &pbgear.Balances_UpgradeAccountsCall{}
        out.Who = to_repeated_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Balances_Who(fields []*registry.DecodedField) *pbgear.Balances_Who {
    out := &pbgear.Balances_Who{}
        
    switch Balances_Value {
    case Balances_Id:
        out.Value = &pbgear.Balances_Who_Id {
            Id: to_Balances_Id(field),
        }
    case Balances_Index:
        out.Value = &pbgear.Balances_Who_Index {
            Index: to_Balances_Index(field),
        }
    case Balances_Raw:
        out.Value = &pbgear.Balances_Who_Raw {
            Raw: to_Balances_Raw(field),
        }
    case Balances_Address32:
        out.Value = &pbgear.Balances_Who_Address32 {
            Address32: to_Balances_Address32(field),
        }
    case Balances_Address20:
        out.Value = &pbgear.Balances_Who_Address20 {
            Address20: to_Balances_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Bounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_AcceptCuratorCall {
    out := &pbgear.Bounties_AcceptCuratorCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_Address20(fields []*registry.DecodedField) *pbgear.Bounties_Address20 {
    out := &pbgear.Bounties_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_Address32(fields []*registry.DecodedField) *pbgear.Bounties_Address32 {
    out := &pbgear.Bounties_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_ApproveBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ApproveBountyCall {
    out := &pbgear.Bounties_ApproveBountyCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_AwardBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_AwardBountyCall {
    out := &pbgear.Bounties_AwardBountyCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Beneficiary = to_Bounties_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Bounties_Beneficiary(fields []*registry.DecodedField) *pbgear.Bounties_Beneficiary {
    out := &pbgear.Bounties_Beneficiary{}
        
    switch Bounties_Value {
    case Bounties_Id:
        out.Value = &pbgear.Bounties_Beneficiary_Id {
            Id: to_Bounties_Id(field),
        }
    case Bounties_Index:
        out.Value = &pbgear.Bounties_Beneficiary_Index {
            Index: to_Bounties_Index(field),
        }
    case Bounties_Raw:
        out.Value = &pbgear.Bounties_Beneficiary_Raw {
            Raw: to_Bounties_Raw(field),
        }
    case Bounties_Address32:
        out.Value = &pbgear.Bounties_Beneficiary_Address32 {
            Address32: to_Bounties_Address32(field),
        }
    case Bounties_Address20:
        out.Value = &pbgear.Bounties_Beneficiary_Address20 {
            Address20: to_Bounties_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Bounties_ClaimBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ClaimBountyCall {
    out := &pbgear.Bounties_ClaimBountyCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_CloseBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_CloseBountyCall {
    out := &pbgear.Bounties_CloseBountyCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_CompactString(fields []*registry.DecodedField) *pbgear.Bounties_CompactString {
    out := &pbgear.Bounties_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_Bounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
    out := &pbgear.Bounties_CompactTupleNull{}
        out.Value = to_Bounties_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_CompactUint32(fields []*registry.DecodedField) *pbgear.Bounties_CompactUint32 {
    out := &pbgear.Bounties_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Bounties_Curator(fields []*registry.DecodedField) *pbgear.Bounties_Curator {
    out := &pbgear.Bounties_Curator{}
        
    switch Bounties_Value {
    case Bounties_Id:
        out.Value = &pbgear.Bounties_Curator_Id {
            Id: to_Bounties_Id(field),
        }
    case Bounties_Index:
        out.Value = &pbgear.Bounties_Curator_Index {
            Index: to_Bounties_Index(field),
        }
    case Bounties_Raw:
        out.Value = &pbgear.Bounties_Curator_Raw {
            Raw: to_Bounties_Raw(field),
        }
    case Bounties_Address32:
        out.Value = &pbgear.Bounties_Curator_Address32 {
            Address32: to_Bounties_Address32(field),
        }
    case Bounties_Address20:
        out.Value = &pbgear.Bounties_Curator_Address20 {
            Address20: to_Bounties_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Bounties_ExtendBountyExpiryCall(fields []*registry.DecodedField) *pbgear.Bounties_ExtendBountyExpiryCall {
    out := &pbgear.Bounties_ExtendBountyExpiryCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Remark = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Bounties_Id(fields []*registry.DecodedField) *pbgear.Bounties_Id {
    out := &pbgear.Bounties_Id{}
        out.Value_0 = to_Bounties_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_Index(fields []*registry.DecodedField) *pbgear.Bounties_Index {
    out := &pbgear.Bounties_Index{}
        out.Value_0 = to_Bounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_ProposeBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeBountyCall {
    out := &pbgear.Bounties_ProposeBountyCall{}
        out.Value = to_Bounties_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Description = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Bounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeCuratorCall {
    out := &pbgear.Bounties_ProposeCuratorCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Curator = to_Bounties_Curator(fields[1].Value.([]*registry.DecodedField)),
        out.Fee = to_Bounties_CompactString(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_Bounties_Raw(fields []*registry.DecodedField) *pbgear.Bounties_Raw {
    out := &pbgear.Bounties_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Bounties_SpCoreCryptoAccountId32 {
    out := &pbgear.Bounties_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Bounties_TupleNull(fields []*registry.DecodedField) *pbgear.Bounties_TupleNull {
    out := &pbgear.Bounties_TupleNull{}
    return out
} 
func to_Bounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_UnassignCuratorCall {
    out := &pbgear.Bounties_UnassignCuratorCall{}
        out.BountyId = to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AcceptCuratorCall {
    out := &pbgear.ChildBounties_AcceptCuratorCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ChildBounties_AddChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AddChildBountyCall {
    out := &pbgear.ChildBounties_AddChildBountyCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value = to_ChildBounties_CompactString(fields[1].Value.([]*registry.DecodedField)),
        out.Description = to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ChildBounties_Address20(fields []*registry.DecodedField) *pbgear.ChildBounties_Address20 {
    out := &pbgear.ChildBounties_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_Address32(fields []*registry.DecodedField) *pbgear.ChildBounties_Address32 {
    out := &pbgear.ChildBounties_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_AwardChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AwardChildBountyCall {
    out := &pbgear.ChildBounties_AwardChildBountyCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
        out.Beneficiary = to_ChildBounties_Beneficiary(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ChildBounties_Beneficiary(fields []*registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
    out := &pbgear.ChildBounties_Beneficiary{}
        
    switch ChildBounties_Value {
    case ChildBounties_Id:
        out.Value = &pbgear.ChildBounties_Beneficiary_Id {
            Id: to_ChildBounties_Id(field),
        }
    case ChildBounties_Index:
        out.Value = &pbgear.ChildBounties_Beneficiary_Index {
            Index: to_ChildBounties_Index(field),
        }
    case ChildBounties_Raw:
        out.Value = &pbgear.ChildBounties_Beneficiary_Raw {
            Raw: to_ChildBounties_Raw(field),
        }
    case ChildBounties_Address32:
        out.Value = &pbgear.ChildBounties_Beneficiary_Address32 {
            Address32: to_ChildBounties_Address32(field),
        }
    case ChildBounties_Address20:
        out.Value = &pbgear.ChildBounties_Beneficiary_Address20 {
            Address20: to_ChildBounties_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ChildBounties_ClaimChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ClaimChildBountyCall {
    out := &pbgear.ChildBounties_ClaimChildBountyCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ChildBounties_CloseChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_CloseChildBountyCall {
    out := &pbgear.ChildBounties_CloseChildBountyCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ChildBounties_CompactString(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactString {
    out := &pbgear.ChildBounties_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_ChildBounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
    out := &pbgear.ChildBounties_CompactTupleNull{}
        out.Value = to_ChildBounties_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_CompactUint32(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
    out := &pbgear.ChildBounties_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_ChildBounties_Curator(fields []*registry.DecodedField) *pbgear.ChildBounties_Curator {
    out := &pbgear.ChildBounties_Curator{}
        
    switch ChildBounties_Value {
    case ChildBounties_Id:
        out.Value = &pbgear.ChildBounties_Curator_Id {
            Id: to_ChildBounties_Id(field),
        }
    case ChildBounties_Index:
        out.Value = &pbgear.ChildBounties_Curator_Index {
            Index: to_ChildBounties_Index(field),
        }
    case ChildBounties_Raw:
        out.Value = &pbgear.ChildBounties_Curator_Raw {
            Raw: to_ChildBounties_Raw(field),
        }
    case ChildBounties_Address32:
        out.Value = &pbgear.ChildBounties_Curator_Address32 {
            Address32: to_ChildBounties_Address32(field),
        }
    case ChildBounties_Address20:
        out.Value = &pbgear.ChildBounties_Curator_Address20 {
            Address20: to_ChildBounties_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ChildBounties_Id(fields []*registry.DecodedField) *pbgear.ChildBounties_Id {
    out := &pbgear.ChildBounties_Id{}
        out.Value_0 = to_ChildBounties_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_Index(fields []*registry.DecodedField) *pbgear.ChildBounties_Index {
    out := &pbgear.ChildBounties_Index{}
        out.Value_0 = to_ChildBounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ProposeCuratorCall {
    out := &pbgear.ChildBounties_ProposeCuratorCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
        out.Curator = to_ChildBounties_Curator(fields[2].Value.([]*registry.DecodedField)),
        out.Fee = to_ChildBounties_CompactString(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_ChildBounties_Raw(fields []*registry.DecodedField) *pbgear.ChildBounties_Raw {
    out := &pbgear.ChildBounties_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ChildBounties_SpCoreCryptoAccountId32 {
    out := &pbgear.ChildBounties_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ChildBounties_TupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_TupleNull {
    out := &pbgear.ChildBounties_TupleNull{}
    return out
} 
func to_ChildBounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_UnassignCuratorCall {
    out := &pbgear.ChildBounties_UnassignCuratorCall{}
        out.ParentBountyId = to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.ChildBountyId = to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ConvictionVoting_Address20(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address20 {
    out := &pbgear.ConvictionVoting_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_Address32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address32 {
    out := &pbgear.ConvictionVoting_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
    out := &pbgear.ConvictionVoting_CompactTupleNull{}
        out.Value = to_ConvictionVoting_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_CompactUint32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
    out := &pbgear.ConvictionVoting_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_ConvictionVoting_Conviction(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
    out := &pbgear.ConvictionVoting_Conviction{}
        
    switch ConvictionVoting_Value {
    case ConvictionVoting_None:
        out.Value = &pbgear.ConvictionVoting_Conviction_None {
            None: to_ConvictionVoting_None(field),
        }
    case ConvictionVoting_Locked1X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked1x {
            Locked1x: to_ConvictionVoting_Locked1X(field),
        }
    case ConvictionVoting_Locked2X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked2x {
            Locked2x: to_ConvictionVoting_Locked2X(field),
        }
    case ConvictionVoting_Locked3X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked3x {
            Locked3x: to_ConvictionVoting_Locked3X(field),
        }
    case ConvictionVoting_Locked4X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked4x {
            Locked4x: to_ConvictionVoting_Locked4X(field),
        }
    case ConvictionVoting_Locked5X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked5x {
            Locked5x: to_ConvictionVoting_Locked5X(field),
        }
    case ConvictionVoting_Locked6X:
        out.Value = &pbgear.ConvictionVoting_Conviction_Locked6x {
            Locked6x: to_ConvictionVoting_Locked6X(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ConvictionVoting_DelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_DelegateCall {
    out := &pbgear.ConvictionVoting_DelegateCall{}
        out.Class = fields[0].Value.(uint32),
        out.To = to_ConvictionVoting_To(fields[1].Value.([]*registry.DecodedField)),
        out.Conviction = to_ConvictionVoting_Conviction(fields[2].Value.([]*registry.DecodedField)),
        out.Balance = fields[3].Value.(string),
    return out
}  
  
  
  
 
func to_ConvictionVoting_Id(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Id {
    out := &pbgear.ConvictionVoting_Id{}
        out.Value_0 = to_ConvictionVoting_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_Index(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Index {
    out := &pbgear.ConvictionVoting_Index{}
        out.Value_0 = to_ConvictionVoting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_Locked1X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked1X {
    out := &pbgear.ConvictionVoting_Locked1X{}
    return out
} 
func to_ConvictionVoting_Locked2X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked2X {
    out := &pbgear.ConvictionVoting_Locked2X{}
    return out
} 
func to_ConvictionVoting_Locked3X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked3X {
    out := &pbgear.ConvictionVoting_Locked3X{}
    return out
} 
func to_ConvictionVoting_Locked4X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked4X {
    out := &pbgear.ConvictionVoting_Locked4X{}
    return out
} 
func to_ConvictionVoting_Locked5X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked5X {
    out := &pbgear.ConvictionVoting_Locked5X{}
    return out
} 
func to_ConvictionVoting_Locked6X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked6X {
    out := &pbgear.ConvictionVoting_Locked6X{}
    return out
} 
func to_ConvictionVoting_None(fields []*registry.DecodedField) *pbgear.ConvictionVoting_None {
    out := &pbgear.ConvictionVoting_None{}
    return out
} 
func to_ConvictionVoting_PalletConvictionVotingVoteVote(fields []*registry.DecodedField) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
    out := &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{}
        out.Vote = fields[0].Value.(uint32),
    return out
}  
 
func to_ConvictionVoting_Raw(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Raw {
    out := &pbgear.ConvictionVoting_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_RemoveOtherVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
    out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
        out.Target = to_ConvictionVoting_Target(fields[0].Value.([]*registry.DecodedField)),
        out.Class = fields[1].Value.(uint32),
        out.Index = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_ConvictionVoting_RemoveVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveVoteCall {
    out := &pbgear.ConvictionVoting_RemoveVoteCall{}
        out.Class = fields[0].Value.(uint32),
        out.Index = fields[1].Value.(uint32),
    return out
}  
  
 
func to_ConvictionVoting_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SpCoreCryptoAccountId32 {
    out := &pbgear.ConvictionVoting_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ConvictionVoting_Split(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Split {
    out := &pbgear.ConvictionVoting_Split{}
        out.Aye = fields[0].Value.(string),
        out.Nay = fields[1].Value.(string),
    return out
}  
  
 
func to_ConvictionVoting_SplitAbstain(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SplitAbstain {
    out := &pbgear.ConvictionVoting_SplitAbstain{}
        out.Aye = fields[0].Value.(string),
        out.Nay = fields[1].Value.(string),
        out.Abstain = fields[2].Value.(string),
    return out
}  
  
  
 
func to_ConvictionVoting_Standard(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Standard {
    out := &pbgear.ConvictionVoting_Standard{}
        out.Vote = to_ConvictionVoting_PalletConvictionVotingVoteVote(fields[0].Value.([]*registry.DecodedField)),
        out.Balance = fields[1].Value.(string),
    return out
}  
  
 
func to_ConvictionVoting_Target(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Target {
    out := &pbgear.ConvictionVoting_Target{}
        
    switch ConvictionVoting_Value {
    case ConvictionVoting_Id:
        out.Value = &pbgear.ConvictionVoting_Target_Id {
            Id: to_ConvictionVoting_Id(field),
        }
    case ConvictionVoting_Index:
        out.Value = &pbgear.ConvictionVoting_Target_Index {
            Index: to_ConvictionVoting_Index(field),
        }
    case ConvictionVoting_Raw:
        out.Value = &pbgear.ConvictionVoting_Target_Raw {
            Raw: to_ConvictionVoting_Raw(field),
        }
    case ConvictionVoting_Address32:
        out.Value = &pbgear.ConvictionVoting_Target_Address32 {
            Address32: to_ConvictionVoting_Address32(field),
        }
    case ConvictionVoting_Address20:
        out.Value = &pbgear.ConvictionVoting_Target_Address20 {
            Address20: to_ConvictionVoting_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ConvictionVoting_To(fields []*registry.DecodedField) *pbgear.ConvictionVoting_To {
    out := &pbgear.ConvictionVoting_To{}
        
    switch ConvictionVoting_Value {
    case ConvictionVoting_Id:
        out.Value = &pbgear.ConvictionVoting_To_Id {
            Id: to_ConvictionVoting_Id(field),
        }
    case ConvictionVoting_Index:
        out.Value = &pbgear.ConvictionVoting_To_Index {
            Index: to_ConvictionVoting_Index(field),
        }
    case ConvictionVoting_Raw:
        out.Value = &pbgear.ConvictionVoting_To_Raw {
            Raw: to_ConvictionVoting_Raw(field),
        }
    case ConvictionVoting_Address32:
        out.Value = &pbgear.ConvictionVoting_To_Address32 {
            Address32: to_ConvictionVoting_Address32(field),
        }
    case ConvictionVoting_Address20:
        out.Value = &pbgear.ConvictionVoting_To_Address20 {
            Address20: to_ConvictionVoting_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ConvictionVoting_TupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
    out := &pbgear.ConvictionVoting_TupleNull{}
    return out
} 
func to_ConvictionVoting_UndelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UndelegateCall {
    out := &pbgear.ConvictionVoting_UndelegateCall{}
        out.Class = fields[0].Value.(uint32),
    return out
}  
 
func to_ConvictionVoting_UnlockCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UnlockCall {
    out := &pbgear.ConvictionVoting_UnlockCall{}
        out.Class = fields[0].Value.(uint32),
        out.Target = to_ConvictionVoting_Target(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ConvictionVoting_Vote(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Vote {
    out := &pbgear.ConvictionVoting_Vote{}
        
    switch ConvictionVoting_Value {
    case ConvictionVoting_Standard:
        out.Value = &pbgear.ConvictionVoting_Vote_Standard {
            Standard: to_ConvictionVoting_Standard(field),
        }
    case ConvictionVoting_Split:
        out.Value = &pbgear.ConvictionVoting_Vote_Split {
            Split: to_ConvictionVoting_Split(field),
        }
    case ConvictionVoting_SplitAbstain:
        out.Value = &pbgear.ConvictionVoting_Vote_SplitAbstain {
            SplitAbstain: to_ConvictionVoting_SplitAbstain(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ConvictionVoting_VoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_VoteCall {
    out := &pbgear.ConvictionVoting_VoteCall{}
        out.PollIndex = to_ConvictionVoting_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Vote = to_ConvictionVoting_Vote(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
    out := &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{}
        out.Value = to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ElectionProviderMultiPhase_CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_ElectionProviderMultiPhase_GovernanceFallbackCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
    out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
        out.MaybeMaxVoters = fields[0].Value.(uint32),
        out.MaybeMaxTargets = fields[1].Value.(uint32),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
    out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
        out.Solution = to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields[0].Value.([]*registry.DecodedField)),
        out.Score = to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[1].Value.([]*registry.DecodedField)),
        out.Round = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
    out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
        out.Voters = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Targets = to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
    out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
        out.Supports = to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
    out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
        out.MaybeNextScore = to_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
    out := &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
    out := &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
    out := &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{}
        out.MinimalStake = fields[0].Value.(string),
        out.SumStake = fields[1].Value.(string),
        out.SumStakeSquared = fields[2].Value.(string),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
    out := &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{}
        out.Total = fields[0].Value.(string),
        out.Voters = to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_SubmitCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitCall {
    out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
        out.RawSolution = to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ElectionProviderMultiPhase_SubmitUnsignedCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
    out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
        out.RawSolution = to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField)),
        out.Witness = to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{}
        out.Value_0 = to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
        out.Value_2 = to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
    out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
        out.Value_0 = to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_ElectionProviderMultiPhase_SpNposElectionsSupport(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
    out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
        out.Value_0 = to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = fields[1].Value.(string),
    return out
}  
  
 
func to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
    out := &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{}
        out.Votes_1 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Votes_2 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields[1].Value.([]*registry.DecodedField)),
        out.Votes_3 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields[2].Value.([]*registry.DecodedField)),
        out.Votes_4 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields[3].Value.([]*registry.DecodedField)),
        out.Votes_5 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields[4].Value.([]*registry.DecodedField)),
        out.Votes_6 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields[5].Value.([]*registry.DecodedField)),
        out.Votes_7 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields[6].Value.([]*registry.DecodedField)),
        out.Votes_8 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields[7].Value.([]*registry.DecodedField)),
        out.Votes_9 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields[8].Value.([]*registry.DecodedField)),
        out.Votes_10 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields[9].Value.([]*registry.DecodedField)),
        out.Votes_11 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields[10].Value.([]*registry.DecodedField)),
        out.Votes_12 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields[11].Value.([]*registry.DecodedField)),
        out.Votes_13 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields[12].Value.([]*registry.DecodedField)),
        out.Votes_14 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields[13].Value.([]*registry.DecodedField)),
        out.Votes_15 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields[14].Value.([]*registry.DecodedField)),
        out.Votes_16 = to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields[15].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
 
func to_FellowshipCollective_AddMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_AddMemberCall {
    out := &pbgear.FellowshipCollective_AddMemberCall{}
        out.Who = to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_Address20(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address20 {
    out := &pbgear.FellowshipCollective_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_Address32(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address32 {
    out := &pbgear.FellowshipCollective_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_CleanupPollCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CleanupPollCall {
    out := &pbgear.FellowshipCollective_CleanupPollCall{}
        out.PollIndex = fields[0].Value.(uint32),
        out.Max = fields[1].Value.(uint32),
    return out
}  
  
 
func to_FellowshipCollective_CompactTupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
    out := &pbgear.FellowshipCollective_CompactTupleNull{}
        out.Value = to_FellowshipCollective_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_DemoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_DemoteMemberCall {
    out := &pbgear.FellowshipCollective_DemoteMemberCall{}
        out.Who = to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_Id(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Id {
    out := &pbgear.FellowshipCollective_Id{}
        out.Value_0 = to_FellowshipCollective_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_Index(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Index {
    out := &pbgear.FellowshipCollective_Index{}
        out.Value_0 = to_FellowshipCollective_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_PromoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_PromoteMemberCall {
    out := &pbgear.FellowshipCollective_PromoteMemberCall{}
        out.Who = to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_Raw(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Raw {
    out := &pbgear.FellowshipCollective_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_RemoveMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_RemoveMemberCall {
    out := &pbgear.FellowshipCollective_RemoveMemberCall{}
        out.Who = to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
        out.MinRank = fields[1].Value.(uint32),
    return out
}  
  
 
func to_FellowshipCollective_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.FellowshipCollective_SpCoreCryptoAccountId32 {
    out := &pbgear.FellowshipCollective_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipCollective_TupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
    out := &pbgear.FellowshipCollective_TupleNull{}
    return out
} 
func to_FellowshipCollective_VoteCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_VoteCall {
    out := &pbgear.FellowshipCollective_VoteCall{}
        out.Poll = fields[0].Value.(uint32),
        out.Aye = fields[1].Value.(bool),
    return out
}  
  
 
func to_FellowshipCollective_Who(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Who {
    out := &pbgear.FellowshipCollective_Who{}
        
    switch FellowshipCollective_Value {
    case FellowshipCollective_Id:
        out.Value = &pbgear.FellowshipCollective_Who_Id {
            Id: to_FellowshipCollective_Id(field),
        }
    case FellowshipCollective_Index:
        out.Value = &pbgear.FellowshipCollective_Who_Index {
            Index: to_FellowshipCollective_Index(field),
        }
    case FellowshipCollective_Raw:
        out.Value = &pbgear.FellowshipCollective_Who_Raw {
            Raw: to_FellowshipCollective_Raw(field),
        }
    case FellowshipCollective_Address32:
        out.Value = &pbgear.FellowshipCollective_Who_Address32 {
            Address32: to_FellowshipCollective_Address32(field),
        }
    case FellowshipCollective_Address20:
        out.Value = &pbgear.FellowshipCollective_Who_Address20 {
            Address20: to_FellowshipCollective_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_FellowshipReferenda_After(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_After {
    out := &pbgear.FellowshipReferenda_After{}
        out.Value_0 = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_At(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_At {
    out := &pbgear.FellowshipReferenda_At{}
        out.Value_0 = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec {
    out := &pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_CancelCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_CancelCall {
    out := &pbgear.FellowshipReferenda_CancelCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
    out := &pbgear.FellowshipReferenda_EnactmentMoment{}
        
    switch FellowshipReferenda_Value {
    case FellowshipReferenda_At:
        out.Value = &pbgear.FellowshipReferenda_EnactmentMoment_At {
            At: to_FellowshipReferenda_At(field),
        }
    case FellowshipReferenda_After:
        out.Value = &pbgear.FellowshipReferenda_EnactmentMoment_After {
            After: to_FellowshipReferenda_After(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_FellowshipReferenda_Inline(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Inline {
    out := &pbgear.FellowshipReferenda_Inline{}
        out.Value_0 = to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_KillCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_KillCall {
    out := &pbgear.FellowshipReferenda_KillCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_Legacy(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Legacy {
    out := &pbgear.FellowshipReferenda_Legacy{}
        out.Hash = to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_Lookup(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Lookup {
    out := &pbgear.FellowshipReferenda_Lookup{}
        out.Hash = to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
        out.Len = fields[1].Value.(uint32),
    return out
}  
  
 
func to_FellowshipReferenda_None(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_None {
    out := &pbgear.FellowshipReferenda_None{}
    return out
} 
func to_FellowshipReferenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_NudgeReferendumCall {
    out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
    out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
        out.Track = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_Origins(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Origins {
    out := &pbgear.FellowshipReferenda_Origins{}
        out.Value_0 = to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
    out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
    out := &pbgear.FellowshipReferenda_PrimitiveTypesH256{}
        out.Hash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_Proposal(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
    out := &pbgear.FellowshipReferenda_Proposal{}
        
    switch FellowshipReferenda_Value {
    case FellowshipReferenda_Legacy:
        out.Value = &pbgear.FellowshipReferenda_Proposal_Legacy {
            Legacy: to_FellowshipReferenda_Legacy(field),
        }
    case FellowshipReferenda_Inline:
        out.Value = &pbgear.FellowshipReferenda_Proposal_Inline {
            Inline: to_FellowshipReferenda_Inline(field),
        }
    case FellowshipReferenda_Lookup:
        out.Value = &pbgear.FellowshipReferenda_Proposal_Lookup {
            Lookup: to_FellowshipReferenda_Lookup(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_FellowshipReferenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
    out := &pbgear.FellowshipReferenda_ProposalOrigin{}
        
    switch FellowshipReferenda_Value {
    case FellowshipReferenda_System:
        out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_system {
            system: to_FellowshipReferenda_System(field),
        }
    case FellowshipReferenda_Origins:
        out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_Origins {
            Origins: to_FellowshipReferenda_Origins(field),
        }
    case FellowshipReferenda_Void:
        out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_Void {
            Void: to_FellowshipReferenda_Void(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_FellowshipReferenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
    out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
    out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_FellowshipReferenda_Root(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Root {
    out := &pbgear.FellowshipReferenda_Root{}
    return out
} 
func to_FellowshipReferenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SetMetadataCall {
    out := &pbgear.FellowshipReferenda_SetMetadataCall{}
        out.Index = fields[0].Value.(uint32),
        out.MaybeHash = to_optional_FellowshipReferenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_FellowshipReferenda_Signed(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Signed {
    out := &pbgear.FellowshipReferenda_Signed{}
        out.Value_0 = to_FellowshipReferenda_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SpCoreCryptoAccountId32 {
    out := &pbgear.FellowshipReferenda_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_SubmitCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SubmitCall {
    out := &pbgear.FellowshipReferenda_SubmitCall{}
        out.ProposalOrigin = to_FellowshipReferenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField)),
        out.Proposal = to_FellowshipReferenda_Proposal(fields[1].Value.([]*registry.DecodedField)),
        out.EnactmentMoment = to_FellowshipReferenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_FellowshipReferenda_System(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_System {
    out := &pbgear.FellowshipReferenda_System{}
        out.Value_0 = to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_FellowshipReferenda_Value0(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Value0 {
    out := &pbgear.FellowshipReferenda_Value0{}
        
    switch FellowshipReferenda_Value {
    case FellowshipReferenda_Root:
        out.Value = &pbgear.FellowshipReferenda_Value0_Root {
            Root: to_FellowshipReferenda_Root(field),
        }
    case FellowshipReferenda_Signed:
        out.Value = &pbgear.FellowshipReferenda_Value0_Signed {
            Signed: to_FellowshipReferenda_Signed(field),
        }
    case FellowshipReferenda_None:
        out.Value = &pbgear.FellowshipReferenda_Value0_None {
            None: to_FellowshipReferenda_None(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_FellowshipReferenda_Void(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Void {
    out := &pbgear.FellowshipReferenda_Void{}
        out.Value_0 = to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_AppendPrograms(fields []*registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
    out := &pbgear.GearVoucher_AppendPrograms{}
        
    switch GearVoucher_Value {
    case GearVoucher_None:
        out.Value = &pbgear.GearVoucher_AppendPrograms_None {
            None: to_GearVoucher_None(field),
        }
    case GearVoucher_Some:
        out.Value = &pbgear.GearVoucher_AppendPrograms_Some {
            Some: to_GearVoucher_Some(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_GearVoucher_BTreeSet(fields []*registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
    out := &pbgear.GearVoucher_BTreeSet{}
        out.Programs = to_repeated_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_Call(fields []*registry.DecodedField) *pbgear.GearVoucher_Call {
    out := &pbgear.GearVoucher_Call{}
        
    switch GearVoucher_Value {
    case GearVoucher_SendMessage:
        out.Value = &pbgear.GearVoucher_Call_SendMessage {
            SendMessage: to_GearVoucher_SendMessage(field),
        }
    case GearVoucher_SendReply:
        out.Value = &pbgear.GearVoucher_Call_SendReply {
            SendReply: to_GearVoucher_SendReply(field),
        }
    case GearVoucher_UploadCode:
        out.Value = &pbgear.GearVoucher_Call_UploadCode {
            UploadCode: to_GearVoucher_UploadCode(field),
        }
    case GearVoucher_DeclineVoucher:
        out.Value = &pbgear.GearVoucher_Call_DeclineVoucher {
            DeclineVoucher: to_GearVoucher_DeclineVoucher(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_GearVoucher_CallCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallCall {
    out := &pbgear.GearVoucher_CallCall{}
        out.VoucherId = to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField)),
        out.Call = to_GearVoucher_Call(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_GearVoucher_CallDeprecatedCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallDeprecatedCall {
    out := &pbgear.GearVoucher_CallDeprecatedCall{}
        out.Call = to_GearVoucher_Call(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_DeclineCall(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineCall {
    out := &pbgear.GearVoucher_DeclineCall{}
        out.VoucherId = to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_DeclineVoucher(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineVoucher {
    out := &pbgear.GearVoucher_DeclineVoucher{}
    return out
} 
func to_GearVoucher_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
    out := &pbgear.GearVoucher_GprimitivesActorId{}
        out.Programs = to_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesMessageId {
    out := &pbgear.GearVoucher_GprimitivesMessageId{}
        out.ReplyToId = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_IssueCall(fields []*registry.DecodedField) *pbgear.GearVoucher_IssueCall {
    out := &pbgear.GearVoucher_IssueCall{}
        out.Spender = to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.Balance = fields[1].Value.(string),
        out.Programs = to_optional_GearVoucher_BTreeSet(fields[2].Value.([]*registry.DecodedField)),
        out.CodeUploading = fields[3].Value.(bool),
        out.Duration = fields[4].Value.(uint32),
    return out
}  
  
  
  
  
 
func to_GearVoucher_None(fields []*registry.DecodedField) *pbgear.GearVoucher_None {
    out := &pbgear.GearVoucher_None{}
    return out
} 
func to_GearVoucher_PalletGearVoucherInternalVoucherId(fields []*registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
    out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
        out.VoucherId = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_RevokeCall(fields []*registry.DecodedField) *pbgear.GearVoucher_RevokeCall {
    out := &pbgear.GearVoucher_RevokeCall{}
        out.Spender = to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.VoucherId = to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_GearVoucher_SendMessage(fields []*registry.DecodedField) *pbgear.GearVoucher_SendMessage {
    out := &pbgear.GearVoucher_SendMessage{}
        out.Destination = to_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
        out.Payload = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[2].Value.(uint64),
        out.Value = fields[3].Value.(string),
        out.KeepAlive = fields[4].Value.(bool),
    return out
}  
  
  
  
  
 
func to_GearVoucher_SendReply(fields []*registry.DecodedField) *pbgear.GearVoucher_SendReply {
    out := &pbgear.GearVoucher_SendReply{}
        out.ReplyToId = to_GearVoucher_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
        out.Payload = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[2].Value.(uint64),
        out.Value = fields[3].Value.(string),
        out.KeepAlive = fields[4].Value.(bool),
    return out
}  
  
  
  
  
 
func to_GearVoucher_Some(fields []*registry.DecodedField) *pbgear.GearVoucher_Some {
    out := &pbgear.GearVoucher_Some{}
        out.Value_0 = to_GearVoucher_BTreeSet(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
    out := &pbgear.GearVoucher_SpCoreCryptoAccountId32{}
        out.Spender = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_GearVoucher_UpdateCall(fields []*registry.DecodedField) *pbgear.GearVoucher_UpdateCall {
    out := &pbgear.GearVoucher_UpdateCall{}
        out.Spender = to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.VoucherId = to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField)),
        out.MoveOwnership = to_optional_GearVoucher_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField)),
        out.BalanceTopUp = fields[3].Value.(string),
        out.AppendPrograms = to_optional_GearVoucher_AppendPrograms(fields[4].Value.([]*registry.DecodedField)),
        out.CodeUploading = fields[5].Value.(bool),
        out.ProlongDuration = fields[6].Value.(uint32),
    return out
}  
  
  
  
  
  
  
 
func to_GearVoucher_UploadCode(fields []*registry.DecodedField) *pbgear.GearVoucher_UploadCode {
    out := &pbgear.GearVoucher_UploadCode{}
        out.Code = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_ClaimValueCall(fields []*registry.DecodedField) *pbgear.Gear_ClaimValueCall {
    out := &pbgear.Gear_ClaimValueCall{}
        out.MessageId = to_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_CreateProgramCall(fields []*registry.DecodedField) *pbgear.Gear_CreateProgramCall {
    out := &pbgear.Gear_CreateProgramCall{}
        out.CodeId = to_Gear_GprimitivesCodeId(fields[0].Value.([]*registry.DecodedField)),
        out.Salt = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.InitPayload = to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[3].Value.(uint64),
        out.Value = fields[4].Value.(string),
        out.KeepAlive = fields[5].Value.(bool),
    return out
}  
  
  
  
  
  
 
func to_Gear_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
    out := &pbgear.Gear_GprimitivesActorId{}
        out.Destination = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_GprimitivesCodeId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
    out := &pbgear.Gear_GprimitivesCodeId{}
        out.CodeId = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
    out := &pbgear.Gear_GprimitivesMessageId{}
        out.ReplyToId = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_RunCall(fields []*registry.DecodedField) *pbgear.Gear_RunCall {
    out := &pbgear.Gear_RunCall{}
        out.MaxGas = fields[0].Value.(uint64),
    return out
}  
 
func to_Gear_SendMessageCall(fields []*registry.DecodedField) *pbgear.Gear_SendMessageCall {
    out := &pbgear.Gear_SendMessageCall{}
        out.Destination = to_Gear_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
        out.Payload = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[2].Value.(uint64),
        out.Value = fields[3].Value.(string),
        out.KeepAlive = fields[4].Value.(bool),
    return out
}  
  
  
  
  
 
func to_Gear_SendReplyCall(fields []*registry.DecodedField) *pbgear.Gear_SendReplyCall {
    out := &pbgear.Gear_SendReplyCall{}
        out.ReplyToId = to_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
        out.Payload = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[2].Value.(uint64),
        out.Value = fields[3].Value.(string),
        out.KeepAlive = fields[4].Value.(bool),
    return out
}  
  
  
  
  
 
func to_Gear_SetExecuteInherentCall(fields []*registry.DecodedField) *pbgear.Gear_SetExecuteInherentCall {
    out := &pbgear.Gear_SetExecuteInherentCall{}
        out.Value = fields[0].Value.(bool),
    return out
}  
 
func to_Gear_UploadCodeCall(fields []*registry.DecodedField) *pbgear.Gear_UploadCodeCall {
    out := &pbgear.Gear_UploadCodeCall{}
        out.Code = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Gear_UploadProgramCall(fields []*registry.DecodedField) *pbgear.Gear_UploadProgramCall {
    out := &pbgear.Gear_UploadProgramCall{}
        out.Code = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Salt = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
        out.InitPayload = to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
        out.GasLimit = fields[3].Value.(uint64),
        out.Value = fields[4].Value.(string),
        out.KeepAlive = fields[5].Value.(bool),
    return out
}  
  
  
  
  
  
 
func to_Grandpa_Equivocation(fields []*registry.DecodedField) *pbgear.Grandpa_Equivocation {
    out := &pbgear.Grandpa_Equivocation{}
        
    switch Grandpa_Value {
    case Grandpa_Prevote:
        out.Value = &pbgear.Grandpa_Equivocation_Prevote {
            Prevote: to_Grandpa_Prevote(field),
        }
    case Grandpa_Precommit:
        out.Value = &pbgear.Grandpa_Equivocation_Precommit {
            Precommit: to_Grandpa_Precommit(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Grandpa_FinalityGrandpaEquivocation(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaEquivocation {
    out := &pbgear.Grandpa_FinalityGrandpaEquivocation{}
        out.RoundNumber = fields[0].Value.(uint64),
        out.Identity = to_Grandpa_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField)),
        out.First = to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[2].Value.([]*registry.DecodedField)),
        out.Second = to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_Grandpa_FinalityGrandpaPrevote(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaPrevote {
    out := &pbgear.Grandpa_FinalityGrandpaPrevote{}
        out.TargetHash = to_Grandpa_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
        out.TargetNumber = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Grandpa_GrandpaTrieNodesList(fields []*registry.DecodedField) *pbgear.Grandpa_GrandpaTrieNodesList {
    out := &pbgear.Grandpa_GrandpaTrieNodesList{}
        out.TrieNodes = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_NoteStalledCall(fields []*registry.DecodedField) *pbgear.Grandpa_NoteStalledCall {
    out := &pbgear.Grandpa_NoteStalledCall{}
        out.Delay = fields[0].Value.(uint32),
        out.BestFinalizedBlockNumber = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Grandpa_Precommit(fields []*registry.DecodedField) *pbgear.Grandpa_Precommit {
    out := &pbgear.Grandpa_Precommit{}
        out.Value_0 = to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_Prevote(fields []*registry.DecodedField) *pbgear.Grandpa_Prevote {
    out := &pbgear.Grandpa_Prevote{}
        out.Value_0 = to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Grandpa_PrimitiveTypesH256 {
    out := &pbgear.Grandpa_PrimitiveTypesH256{}
        out.TargetHash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationCall {
    out := &pbgear.Grandpa_ReportEquivocationCall{}
        out.EquivocationProof = to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
        out.KeyOwnerProof = to_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Grandpa_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
    out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
        out.EquivocationProof = to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
        out.KeyOwnerProof = to_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Grandpa_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
    out := &pbgear.Grandpa_SpConsensusGrandpaAppPublic{}
        out.Identity = to_Grandpa_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_SpConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
    out := &pbgear.Grandpa_SpConsensusGrandpaAppSignature{}
        out.Value_1 = to_Grandpa_SpCoreEd25519Signature(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_SpConsensusGrandpaEquivocationProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
    out := &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{}
        out.SetId = fields[0].Value.(uint64),
        out.Equivocation = to_Grandpa_Equivocation(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Grandpa_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Public {
    out := &pbgear.Grandpa_SpCoreEd25519Public{}
        out.Identity = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_SpCoreEd25519Signature(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Signature {
    out := &pbgear.Grandpa_SpCoreEd25519Signature{}
        out.Value_1 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Grandpa_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
    out := &pbgear.Grandpa_SpSessionMembershipProof{}
        out.Session = fields[0].Value.(uint32),
        out.TrieNodes = to_repeated_Grandpa_GrandpaTrieNodesList(fields[1].Value.([]*registry.DecodedField)),
        out.ValidatorCount = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
    out := &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
        out.Value_0 = to_Grandpa_FinalityGrandpaPrevote(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_Grandpa_SpConsensusGrandpaAppSignature(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_Account(fields []*registry.DecodedField) *pbgear.Identity_Account {
    out := &pbgear.Identity_Account{}
        
    switch Identity_Value {
    case Identity_Id:
        out.Value = &pbgear.Identity_Account_Id {
            Id: to_Identity_Id(field),
        }
    case Identity_Index:
        out.Value = &pbgear.Identity_Account_Index {
            Index: to_Identity_Index(field),
        }
    case Identity_Raw:
        out.Value = &pbgear.Identity_Account_Raw {
            Raw: to_Identity_Raw(field),
        }
    case Identity_Address32:
        out.Value = &pbgear.Identity_Account_Address32 {
            Address32: to_Identity_Address32(field),
        }
    case Identity_Address20:
        out.Value = &pbgear.Identity_Account_Address20 {
            Address20: to_Identity_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_AddRegistrarCall(fields []*registry.DecodedField) *pbgear.Identity_AddRegistrarCall {
    out := &pbgear.Identity_AddRegistrarCall{}
        out.Account = to_Identity_Account(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_AddSubCall(fields []*registry.DecodedField) *pbgear.Identity_AddSubCall {
    out := &pbgear.Identity_AddSubCall{}
        out.Sub = to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
        out.Data = to_Identity_Data(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_Address20(fields []*registry.DecodedField) *pbgear.Identity_Address20 {
    out := &pbgear.Identity_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Address32(fields []*registry.DecodedField) *pbgear.Identity_Address32 {
    out := &pbgear.Identity_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_BlakeTwo256(fields []*registry.DecodedField) *pbgear.Identity_BlakeTwo256 {
    out := &pbgear.Identity_BlakeTwo256{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
    out := &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{}
        out.Additional = to_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_CancelRequestCall(fields []*registry.DecodedField) *pbgear.Identity_CancelRequestCall {
    out := &pbgear.Identity_CancelRequestCall{}
        out.RegIndex = fields[0].Value.(uint32),
    return out
}  
 
func to_Identity_ClearIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_ClearIdentityCall {
    out := &pbgear.Identity_ClearIdentityCall{}
    return out
} 
func to_Identity_CompactString(fields []*registry.DecodedField) *pbgear.Identity_CompactString {
    out := &pbgear.Identity_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_Identity_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Identity_CompactTupleNull {
    out := &pbgear.Identity_CompactTupleNull{}
        out.Value = to_Identity_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_CompactUint32(fields []*registry.DecodedField) *pbgear.Identity_CompactUint32 {
    out := &pbgear.Identity_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Identity_Data(fields []*registry.DecodedField) *pbgear.Identity_Data {
    out := &pbgear.Identity_Data{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Data_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Data_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Data_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Data_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Data_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Data_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Data_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Data_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Data_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Data_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Data_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Data_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Data_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Data_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Data_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Data_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Data_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Data_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Data_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Data_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Data_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Data_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Data_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Data_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Data_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Data_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Data_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Data_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Data_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Data_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Data_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Data_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Data_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Data_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Data_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Data_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Data_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Data_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Display(fields []*registry.DecodedField) *pbgear.Identity_Display {
    out := &pbgear.Identity_Display{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Display_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Display_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Display_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Display_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Display_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Display_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Display_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Display_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Display_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Display_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Display_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Display_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Display_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Display_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Display_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Display_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Display_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Display_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Display_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Display_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Display_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Display_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Display_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Display_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Display_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Display_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Display_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Display_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Display_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Display_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Display_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Display_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Display_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Display_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Display_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Display_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Display_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Display_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Email(fields []*registry.DecodedField) *pbgear.Identity_Email {
    out := &pbgear.Identity_Email{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Email_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Email_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Email_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Email_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Email_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Email_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Email_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Email_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Email_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Email_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Email_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Email_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Email_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Email_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Email_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Email_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Email_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Email_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Email_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Email_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Email_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Email_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Email_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Email_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Email_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Email_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Email_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Email_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Email_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Email_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Email_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Email_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Email_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Email_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Email_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Email_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Email_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Email_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Erroneous(fields []*registry.DecodedField) *pbgear.Identity_Erroneous {
    out := &pbgear.Identity_Erroneous{}
    return out
} 
func to_Identity_FeePaid(fields []*registry.DecodedField) *pbgear.Identity_FeePaid {
    out := &pbgear.Identity_FeePaid{}
        out.Value_0 = fields[0].Value.(string),
    return out
}  
 
func to_Identity_Id(fields []*registry.DecodedField) *pbgear.Identity_Id {
    out := &pbgear.Identity_Id{}
        out.Value_0 = to_Identity_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Image(fields []*registry.DecodedField) *pbgear.Identity_Image {
    out := &pbgear.Identity_Image{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Image_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Image_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Image_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Image_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Image_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Image_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Image_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Image_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Image_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Image_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Image_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Image_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Image_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Image_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Image_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Image_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Image_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Image_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Image_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Image_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Image_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Image_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Image_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Image_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Image_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Image_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Image_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Image_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Image_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Image_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Image_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Image_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Image_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Image_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Image_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Image_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Image_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Image_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Index(fields []*registry.DecodedField) *pbgear.Identity_Index {
    out := &pbgear.Identity_Index{}
        out.Value_0 = to_Identity_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Judgement(fields []*registry.DecodedField) *pbgear.Identity_Judgement {
    out := &pbgear.Identity_Judgement{}
        
    switch Identity_Value {
    case Identity_Unknown:
        out.Value = &pbgear.Identity_Judgement_Unknown {
            Unknown: to_Identity_Unknown(field),
        }
    case Identity_FeePaid:
        out.Value = &pbgear.Identity_Judgement_FeePaid {
            FeePaid: to_Identity_FeePaid(field),
        }
    case Identity_Reasonable:
        out.Value = &pbgear.Identity_Judgement_Reasonable {
            Reasonable: to_Identity_Reasonable(field),
        }
    case Identity_KnownGood:
        out.Value = &pbgear.Identity_Judgement_KnownGood {
            KnownGood: to_Identity_KnownGood(field),
        }
    case Identity_OutOfDate:
        out.Value = &pbgear.Identity_Judgement_OutOfDate {
            OutOfDate: to_Identity_OutOfDate(field),
        }
    case Identity_LowQuality:
        out.Value = &pbgear.Identity_Judgement_LowQuality {
            LowQuality: to_Identity_LowQuality(field),
        }
    case Identity_Erroneous:
        out.Value = &pbgear.Identity_Judgement_Erroneous {
            Erroneous: to_Identity_Erroneous(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Keccak256(fields []*registry.DecodedField) *pbgear.Identity_Keccak256 {
    out := &pbgear.Identity_Keccak256{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_KillIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_KillIdentityCall {
    out := &pbgear.Identity_KillIdentityCall{}
        out.Target = to_Identity_Target(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_KnownGood(fields []*registry.DecodedField) *pbgear.Identity_KnownGood {
    out := &pbgear.Identity_KnownGood{}
    return out
} 
func to_Identity_Legal(fields []*registry.DecodedField) *pbgear.Identity_Legal {
    out := &pbgear.Identity_Legal{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Legal_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Legal_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Legal_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Legal_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Legal_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Legal_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Legal_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Legal_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Legal_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Legal_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Legal_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Legal_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Legal_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Legal_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Legal_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Legal_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Legal_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Legal_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Legal_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Legal_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Legal_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Legal_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Legal_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Legal_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Legal_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Legal_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Legal_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Legal_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Legal_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Legal_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Legal_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Legal_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Legal_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Legal_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Legal_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Legal_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Legal_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Legal_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_LowQuality(fields []*registry.DecodedField) *pbgear.Identity_LowQuality {
    out := &pbgear.Identity_LowQuality{}
    return out
} 
func to_Identity_New(fields []*registry.DecodedField) *pbgear.Identity_New {
    out := &pbgear.Identity_New{}
        
    switch Identity_Value {
    case Identity_Id:
        out.Value = &pbgear.Identity_New_Id {
            Id: to_Identity_Id(field),
        }
    case Identity_Index:
        out.Value = &pbgear.Identity_New_Index {
            Index: to_Identity_Index(field),
        }
    case Identity_Raw:
        out.Value = &pbgear.Identity_New_Raw {
            Raw: to_Identity_Raw(field),
        }
    case Identity_Address32:
        out.Value = &pbgear.Identity_New_Address32 {
            Address32: to_Identity_Address32(field),
        }
    case Identity_Address20:
        out.Value = &pbgear.Identity_New_Address20 {
            Address20: to_Identity_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_None(fields []*registry.DecodedField) *pbgear.Identity_None {
    out := &pbgear.Identity_None{}
    return out
} 
func to_Identity_OutOfDate(fields []*registry.DecodedField) *pbgear.Identity_OutOfDate {
    out := &pbgear.Identity_OutOfDate{}
    return out
} 
func to_Identity_PalletIdentitySimpleIdentityInfo(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
    out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
        out.Additional = to_Identity_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
        out.Display = to_Identity_Display(fields[1].Value.([]*registry.DecodedField)),
        out.Legal = to_Identity_Legal(fields[2].Value.([]*registry.DecodedField)),
        out.Web = to_Identity_Web(fields[3].Value.([]*registry.DecodedField)),
        out.Riot = to_Identity_Riot(fields[4].Value.([]*registry.DecodedField)),
        out.Email = to_Identity_Email(fields[5].Value.([]*registry.DecodedField)),
        out.PgpFingerprint = to_repeated_uint32(fields[6].Value.([]*registry.DecodedField)),
        out.Image = to_Identity_Image(fields[7].Value.([]*registry.DecodedField)),
        out.Twitter = to_Identity_Twitter(fields[8].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
  
  
  
  
 
func to_Identity_PalletIdentityTypesBitFlags(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
    out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
        out.Fields = fields[0].Value.(uint64),
    return out
}  
 
func to_Identity_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
    out := &pbgear.Identity_PrimitiveTypesH256{}
        out.Identity = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_ProvideJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_ProvideJudgementCall {
    out := &pbgear.Identity_ProvideJudgementCall{}
        out.RegIndex = to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Target = to_Identity_Target(fields[1].Value.([]*registry.DecodedField)),
        out.Judgement = to_Identity_Judgement(fields[2].Value.([]*registry.DecodedField)),
        out.Identity = to_Identity_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_Identity_QuitSubCall(fields []*registry.DecodedField) *pbgear.Identity_QuitSubCall {
    out := &pbgear.Identity_QuitSubCall{}
    return out
} 
func to_Identity_Raw(fields []*registry.DecodedField) *pbgear.Identity_Raw {
    out := &pbgear.Identity_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw0(fields []*registry.DecodedField) *pbgear.Identity_Raw0 {
    out := &pbgear.Identity_Raw0{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw1(fields []*registry.DecodedField) *pbgear.Identity_Raw1 {
    out := &pbgear.Identity_Raw1{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw10(fields []*registry.DecodedField) *pbgear.Identity_Raw10 {
    out := &pbgear.Identity_Raw10{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw11(fields []*registry.DecodedField) *pbgear.Identity_Raw11 {
    out := &pbgear.Identity_Raw11{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw12(fields []*registry.DecodedField) *pbgear.Identity_Raw12 {
    out := &pbgear.Identity_Raw12{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw13(fields []*registry.DecodedField) *pbgear.Identity_Raw13 {
    out := &pbgear.Identity_Raw13{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw14(fields []*registry.DecodedField) *pbgear.Identity_Raw14 {
    out := &pbgear.Identity_Raw14{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw15(fields []*registry.DecodedField) *pbgear.Identity_Raw15 {
    out := &pbgear.Identity_Raw15{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw16(fields []*registry.DecodedField) *pbgear.Identity_Raw16 {
    out := &pbgear.Identity_Raw16{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw17(fields []*registry.DecodedField) *pbgear.Identity_Raw17 {
    out := &pbgear.Identity_Raw17{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw18(fields []*registry.DecodedField) *pbgear.Identity_Raw18 {
    out := &pbgear.Identity_Raw18{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw19(fields []*registry.DecodedField) *pbgear.Identity_Raw19 {
    out := &pbgear.Identity_Raw19{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw2(fields []*registry.DecodedField) *pbgear.Identity_Raw2 {
    out := &pbgear.Identity_Raw2{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw20(fields []*registry.DecodedField) *pbgear.Identity_Raw20 {
    out := &pbgear.Identity_Raw20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw21(fields []*registry.DecodedField) *pbgear.Identity_Raw21 {
    out := &pbgear.Identity_Raw21{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw22(fields []*registry.DecodedField) *pbgear.Identity_Raw22 {
    out := &pbgear.Identity_Raw22{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw23(fields []*registry.DecodedField) *pbgear.Identity_Raw23 {
    out := &pbgear.Identity_Raw23{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw24(fields []*registry.DecodedField) *pbgear.Identity_Raw24 {
    out := &pbgear.Identity_Raw24{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw25(fields []*registry.DecodedField) *pbgear.Identity_Raw25 {
    out := &pbgear.Identity_Raw25{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw26(fields []*registry.DecodedField) *pbgear.Identity_Raw26 {
    out := &pbgear.Identity_Raw26{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw27(fields []*registry.DecodedField) *pbgear.Identity_Raw27 {
    out := &pbgear.Identity_Raw27{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw28(fields []*registry.DecodedField) *pbgear.Identity_Raw28 {
    out := &pbgear.Identity_Raw28{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw29(fields []*registry.DecodedField) *pbgear.Identity_Raw29 {
    out := &pbgear.Identity_Raw29{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw3(fields []*registry.DecodedField) *pbgear.Identity_Raw3 {
    out := &pbgear.Identity_Raw3{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw30(fields []*registry.DecodedField) *pbgear.Identity_Raw30 {
    out := &pbgear.Identity_Raw30{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw31(fields []*registry.DecodedField) *pbgear.Identity_Raw31 {
    out := &pbgear.Identity_Raw31{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw32(fields []*registry.DecodedField) *pbgear.Identity_Raw32 {
    out := &pbgear.Identity_Raw32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw4(fields []*registry.DecodedField) *pbgear.Identity_Raw4 {
    out := &pbgear.Identity_Raw4{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw5(fields []*registry.DecodedField) *pbgear.Identity_Raw5 {
    out := &pbgear.Identity_Raw5{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw6(fields []*registry.DecodedField) *pbgear.Identity_Raw6 {
    out := &pbgear.Identity_Raw6{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw7(fields []*registry.DecodedField) *pbgear.Identity_Raw7 {
    out := &pbgear.Identity_Raw7{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw8(fields []*registry.DecodedField) *pbgear.Identity_Raw8 {
    out := &pbgear.Identity_Raw8{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Raw9(fields []*registry.DecodedField) *pbgear.Identity_Raw9 {
    out := &pbgear.Identity_Raw9{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Reasonable(fields []*registry.DecodedField) *pbgear.Identity_Reasonable {
    out := &pbgear.Identity_Reasonable{}
    return out
} 
func to_Identity_RemoveSubCall(fields []*registry.DecodedField) *pbgear.Identity_RemoveSubCall {
    out := &pbgear.Identity_RemoveSubCall{}
        out.Sub = to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_RenameSubCall(fields []*registry.DecodedField) *pbgear.Identity_RenameSubCall {
    out := &pbgear.Identity_RenameSubCall{}
        out.Sub = to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
        out.Data = to_Identity_Data(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_RequestJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_RequestJudgementCall {
    out := &pbgear.Identity_RequestJudgementCall{}
        out.RegIndex = to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.MaxFee = to_Identity_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_Riot(fields []*registry.DecodedField) *pbgear.Identity_Riot {
    out := &pbgear.Identity_Riot{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Riot_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Riot_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Riot_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Riot_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Riot_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Riot_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Riot_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Riot_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Riot_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Riot_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Riot_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Riot_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Riot_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Riot_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Riot_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Riot_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Riot_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Riot_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Riot_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Riot_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Riot_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Riot_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Riot_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Riot_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Riot_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Riot_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Riot_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Riot_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Riot_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Riot_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Riot_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Riot_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Riot_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Riot_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Riot_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Riot_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Riot_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Riot_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_SetAccountIdCall(fields []*registry.DecodedField) *pbgear.Identity_SetAccountIdCall {
    out := &pbgear.Identity_SetAccountIdCall{}
        out.Index = to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.New = to_Identity_New(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_SetFeeCall(fields []*registry.DecodedField) *pbgear.Identity_SetFeeCall {
    out := &pbgear.Identity_SetFeeCall{}
        out.Index = to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Fee = to_Identity_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_SetFieldsCall(fields []*registry.DecodedField) *pbgear.Identity_SetFieldsCall {
    out := &pbgear.Identity_SetFieldsCall{}
        out.Index = to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
        out.Fields = to_Identity_PalletIdentityTypesBitFlags(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_SetIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_SetIdentityCall {
    out := &pbgear.Identity_SetIdentityCall{}
        out.Info = to_Identity_PalletIdentitySimpleIdentityInfo(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_SetSubsCall(fields []*registry.DecodedField) *pbgear.Identity_SetSubsCall {
    out := &pbgear.Identity_SetSubsCall{}
        out.Subs = to_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Sha256(fields []*registry.DecodedField) *pbgear.Identity_Sha256 {
    out := &pbgear.Identity_Sha256{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_ShaThree256(fields []*registry.DecodedField) *pbgear.Identity_ShaThree256 {
    out := &pbgear.Identity_ShaThree256{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
    out := &pbgear.Identity_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Identity_Sub(fields []*registry.DecodedField) *pbgear.Identity_Sub {
    out := &pbgear.Identity_Sub{}
        
    switch Identity_Value {
    case Identity_Id:
        out.Value = &pbgear.Identity_Sub_Id {
            Id: to_Identity_Id(field),
        }
    case Identity_Index:
        out.Value = &pbgear.Identity_Sub_Index {
            Index: to_Identity_Index(field),
        }
    case Identity_Raw:
        out.Value = &pbgear.Identity_Sub_Raw {
            Raw: to_Identity_Raw(field),
        }
    case Identity_Address32:
        out.Value = &pbgear.Identity_Sub_Address32 {
            Address32: to_Identity_Address32(field),
        }
    case Identity_Address20:
        out.Value = &pbgear.Identity_Sub_Address20 {
            Address20: to_Identity_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Target(fields []*registry.DecodedField) *pbgear.Identity_Target {
    out := &pbgear.Identity_Target{}
        
    switch Identity_Value {
    case Identity_Id:
        out.Value = &pbgear.Identity_Target_Id {
            Id: to_Identity_Id(field),
        }
    case Identity_Index:
        out.Value = &pbgear.Identity_Target_Index {
            Index: to_Identity_Index(field),
        }
    case Identity_Raw:
        out.Value = &pbgear.Identity_Target_Raw {
            Raw: to_Identity_Raw(field),
        }
    case Identity_Address32:
        out.Value = &pbgear.Identity_Target_Address32 {
            Address32: to_Identity_Address32(field),
        }
    case Identity_Address20:
        out.Value = &pbgear.Identity_Target_Address20 {
            Address20: to_Identity_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_TupleNull(fields []*registry.DecodedField) *pbgear.Identity_TupleNull {
    out := &pbgear.Identity_TupleNull{}
    return out
} 
func to_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
    out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
        out.Value_0 = to_Identity_Value0(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_Identity_Value1(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
    out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
        out.Value_0 = to_Identity_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_Identity_Value1(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Identity_Twitter(fields []*registry.DecodedField) *pbgear.Identity_Twitter {
    out := &pbgear.Identity_Twitter{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Twitter_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Twitter_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Twitter_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Twitter_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Twitter_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Twitter_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Twitter_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Twitter_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Twitter_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Twitter_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Twitter_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Twitter_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Twitter_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Twitter_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Twitter_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Twitter_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Twitter_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Twitter_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Twitter_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Twitter_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Twitter_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Twitter_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Twitter_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Twitter_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Twitter_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Twitter_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Twitter_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Twitter_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Twitter_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Twitter_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Twitter_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Twitter_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Twitter_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Twitter_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Twitter_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Twitter_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Twitter_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Twitter_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Unknown(fields []*registry.DecodedField) *pbgear.Identity_Unknown {
    out := &pbgear.Identity_Unknown{}
    return out
} 
func to_Identity_Value0(fields []*registry.DecodedField) *pbgear.Identity_Value0 {
    out := &pbgear.Identity_Value0{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Value0_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Value0_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Value0_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Value0_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Value0_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Value0_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Value0_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Value0_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Value0_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Value0_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Value0_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Value0_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Value0_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Value0_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Value0_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Value0_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Value0_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Value0_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Value0_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Value0_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Value0_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Value0_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Value0_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Value0_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Value0_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Value0_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Value0_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Value0_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Value0_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Value0_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Value0_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Value0_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Value0_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Value0_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Value0_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Value0_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Value0_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Value0_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Value1(fields []*registry.DecodedField) *pbgear.Identity_Value1 {
    out := &pbgear.Identity_Value1{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Value1_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Value1_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Value1_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Value1_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Value1_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Value1_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Value1_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Value1_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Value1_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Value1_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Value1_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Value1_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Value1_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Value1_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Value1_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Value1_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Value1_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Value1_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Value1_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Value1_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Value1_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Value1_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Value1_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Value1_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Value1_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Value1_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Value1_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Value1_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Value1_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Value1_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Value1_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Value1_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Value1_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Value1_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Value1_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Value1_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Value1_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Value1_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Identity_Web(fields []*registry.DecodedField) *pbgear.Identity_Web {
    out := &pbgear.Identity_Web{}
        
    switch Identity_Value {
    case Identity_None:
        out.Value = &pbgear.Identity_Web_None {
            None: to_Identity_None(field),
        }
    case Identity_Raw0:
        out.Value = &pbgear.Identity_Web_Raw0 {
            Raw0: to_Identity_Raw0(field),
        }
    case Identity_Raw1:
        out.Value = &pbgear.Identity_Web_Raw1 {
            Raw1: to_Identity_Raw1(field),
        }
    case Identity_Raw2:
        out.Value = &pbgear.Identity_Web_Raw2 {
            Raw2: to_Identity_Raw2(field),
        }
    case Identity_Raw3:
        out.Value = &pbgear.Identity_Web_Raw3 {
            Raw3: to_Identity_Raw3(field),
        }
    case Identity_Raw4:
        out.Value = &pbgear.Identity_Web_Raw4 {
            Raw4: to_Identity_Raw4(field),
        }
    case Identity_Raw5:
        out.Value = &pbgear.Identity_Web_Raw5 {
            Raw5: to_Identity_Raw5(field),
        }
    case Identity_Raw6:
        out.Value = &pbgear.Identity_Web_Raw6 {
            Raw6: to_Identity_Raw6(field),
        }
    case Identity_Raw7:
        out.Value = &pbgear.Identity_Web_Raw7 {
            Raw7: to_Identity_Raw7(field),
        }
    case Identity_Raw8:
        out.Value = &pbgear.Identity_Web_Raw8 {
            Raw8: to_Identity_Raw8(field),
        }
    case Identity_Raw9:
        out.Value = &pbgear.Identity_Web_Raw9 {
            Raw9: to_Identity_Raw9(field),
        }
    case Identity_Raw10:
        out.Value = &pbgear.Identity_Web_Raw10 {
            Raw10: to_Identity_Raw10(field),
        }
    case Identity_Raw11:
        out.Value = &pbgear.Identity_Web_Raw11 {
            Raw11: to_Identity_Raw11(field),
        }
    case Identity_Raw12:
        out.Value = &pbgear.Identity_Web_Raw12 {
            Raw12: to_Identity_Raw12(field),
        }
    case Identity_Raw13:
        out.Value = &pbgear.Identity_Web_Raw13 {
            Raw13: to_Identity_Raw13(field),
        }
    case Identity_Raw14:
        out.Value = &pbgear.Identity_Web_Raw14 {
            Raw14: to_Identity_Raw14(field),
        }
    case Identity_Raw15:
        out.Value = &pbgear.Identity_Web_Raw15 {
            Raw15: to_Identity_Raw15(field),
        }
    case Identity_Raw16:
        out.Value = &pbgear.Identity_Web_Raw16 {
            Raw16: to_Identity_Raw16(field),
        }
    case Identity_Raw17:
        out.Value = &pbgear.Identity_Web_Raw17 {
            Raw17: to_Identity_Raw17(field),
        }
    case Identity_Raw18:
        out.Value = &pbgear.Identity_Web_Raw18 {
            Raw18: to_Identity_Raw18(field),
        }
    case Identity_Raw19:
        out.Value = &pbgear.Identity_Web_Raw19 {
            Raw19: to_Identity_Raw19(field),
        }
    case Identity_Raw20:
        out.Value = &pbgear.Identity_Web_Raw20 {
            Raw20: to_Identity_Raw20(field),
        }
    case Identity_Raw21:
        out.Value = &pbgear.Identity_Web_Raw21 {
            Raw21: to_Identity_Raw21(field),
        }
    case Identity_Raw22:
        out.Value = &pbgear.Identity_Web_Raw22 {
            Raw22: to_Identity_Raw22(field),
        }
    case Identity_Raw23:
        out.Value = &pbgear.Identity_Web_Raw23 {
            Raw23: to_Identity_Raw23(field),
        }
    case Identity_Raw24:
        out.Value = &pbgear.Identity_Web_Raw24 {
            Raw24: to_Identity_Raw24(field),
        }
    case Identity_Raw25:
        out.Value = &pbgear.Identity_Web_Raw25 {
            Raw25: to_Identity_Raw25(field),
        }
    case Identity_Raw26:
        out.Value = &pbgear.Identity_Web_Raw26 {
            Raw26: to_Identity_Raw26(field),
        }
    case Identity_Raw27:
        out.Value = &pbgear.Identity_Web_Raw27 {
            Raw27: to_Identity_Raw27(field),
        }
    case Identity_Raw28:
        out.Value = &pbgear.Identity_Web_Raw28 {
            Raw28: to_Identity_Raw28(field),
        }
    case Identity_Raw29:
        out.Value = &pbgear.Identity_Web_Raw29 {
            Raw29: to_Identity_Raw29(field),
        }
    case Identity_Raw30:
        out.Value = &pbgear.Identity_Web_Raw30 {
            Raw30: to_Identity_Raw30(field),
        }
    case Identity_Raw31:
        out.Value = &pbgear.Identity_Web_Raw31 {
            Raw31: to_Identity_Raw31(field),
        }
    case Identity_Raw32:
        out.Value = &pbgear.Identity_Web_Raw32 {
            Raw32: to_Identity_Raw32(field),
        }
    case Identity_BlakeTwo256:
        out.Value = &pbgear.Identity_Web_BlakeTwo256 {
            BlakeTwo256: to_Identity_BlakeTwo256(field),
        }
    case Identity_Sha256:
        out.Value = &pbgear.Identity_Web_Sha256 {
            Sha256: to_Identity_Sha256(field),
        }
    case Identity_Keccak256:
        out.Value = &pbgear.Identity_Web_Keccak256 {
            Keccak256: to_Identity_Keccak256(field),
        }
    case Identity_ShaThree256:
        out.Value = &pbgear.Identity_Web_ShaThree256 {
            ShaThree256: to_Identity_ShaThree256(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_ImOnline_HeartbeatCall(fields []*registry.DecodedField) *pbgear.ImOnline_HeartbeatCall {
    out := &pbgear.ImOnline_HeartbeatCall{}
        out.Heartbeat = to_ImOnline_PalletImOnlineHeartbeat(fields[0].Value.([]*registry.DecodedField)),
        out.Signature = to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_ImOnline_PalletImOnlineHeartbeat(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
    out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
        out.BlockNumber = fields[0].Value.(uint32),
        out.SessionIndex = fields[1].Value.(uint32),
        out.AuthorityIndex = fields[2].Value.(uint32),
        out.ValidatorsLen = fields[3].Value.(uint32),
    return out
}  
  
  
  
 
func to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
    out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
        out.Signature = to_ImOnline_SpCoreSr25519Signature(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_ImOnline_SpCoreSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
    out := &pbgear.ImOnline_SpCoreSr25519Signature{}
        out.Signature = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Multisig_ApproveAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_ApproveAsMultiCall {
    out := &pbgear.Multisig_ApproveAsMultiCall{}
        out.Threshold = fields[0].Value.(uint32),
        out.OtherSignatories = to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
        out.MaybeTimepoint = to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
        out.CallHash = to_repeated_uint32(fields[3].Value.([]*registry.DecodedField)),
        out.MaxWeight = to_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
 
func to_Multisig_AsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiCall {
    out := &pbgear.Multisig_AsMultiCall{}
        out.Threshold = fields[0].Value.(uint32),
        out.OtherSignatories = to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
        out.MaybeTimepoint = to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
        
    switch Multisig_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Multisig_AsMultiCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Multisig_AsMultiCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Multisig_AsMultiCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Multisig_AsMultiCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Multisig_AsMultiCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Multisig_AsMultiCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Multisig_AsMultiCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Multisig_AsMultiCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Multisig_AsMultiCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Multisig_AsMultiCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Multisig_AsMultiCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
        out.MaxWeight = to_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
 
func to_Multisig_AsMultiThreshold1Call(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiThreshold1Call {
    out := &pbgear.Multisig_AsMultiThreshold1Call{}
        out.OtherSignatories = to_repeated_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        
    switch Multisig_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
 
func to_Multisig_CancelAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_CancelAsMultiCall {
    out := &pbgear.Multisig_CancelAsMultiCall{}
        out.Threshold = fields[0].Value.(uint32),
        out.OtherSignatories = to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
        out.Timepoint = to_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
        out.CallHash = to_repeated_uint32(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_Multisig_CompactUint64(fields []*registry.DecodedField) *pbgear.Multisig_CompactUint64 {
    out := &pbgear.Multisig_CompactUint64{}
        out.Value = fields[0].Value.(uint64),
    return out
}  
 
func to_Multisig_PalletMultisigTimepoint(fields []*registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
    out := &pbgear.Multisig_PalletMultisigTimepoint{}
        out.Height = fields[0].Value.(uint32),
        out.Index = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Multisig_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Multisig_SpCoreCryptoAccountId32 {
    out := &pbgear.Multisig_SpCoreCryptoAccountId32{}
        out.OtherSignatories = to_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Multisig_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
    out := &pbgear.Multisig_SpWeightsWeightV2Weight{}
        out.RefTime = to_Multisig_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
        out.ProofSize = to_Multisig_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_Address20(fields []*registry.DecodedField) *pbgear.NominationPools_Address20 {
    out := &pbgear.NominationPools_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_Address32(fields []*registry.DecodedField) *pbgear.NominationPools_Address32 {
    out := &pbgear.NominationPools_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_AdjustPoolDepositCall(fields []*registry.DecodedField) *pbgear.NominationPools_AdjustPoolDepositCall {
    out := &pbgear.NominationPools_AdjustPoolDepositCall{}
        out.PoolId = fields[0].Value.(uint32),
    return out
}  
 
func to_NominationPools_Blocked(fields []*registry.DecodedField) *pbgear.NominationPools_Blocked {
    out := &pbgear.NominationPools_Blocked{}
    return out
} 
func to_NominationPools_BondExtraCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraCall {
    out := &pbgear.NominationPools_BondExtraCall{}
        out.Extra = to_NominationPools_Extra(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_BondExtraOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraOtherCall {
    out := &pbgear.NominationPools_BondExtraOtherCall{}
        out.Member = to_NominationPools_Member(fields[0].Value.([]*registry.DecodedField)),
        out.Extra = to_NominationPools_Extra(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_Bouncer(fields []*registry.DecodedField) *pbgear.NominationPools_Bouncer {
    out := &pbgear.NominationPools_Bouncer{}
        
    switch NominationPools_Value {
    case NominationPools_Id:
        out.Value = &pbgear.NominationPools_Bouncer_Id {
            Id: to_NominationPools_Id(field),
        }
    case NominationPools_Index:
        out.Value = &pbgear.NominationPools_Bouncer_Index {
            Index: to_NominationPools_Index(field),
        }
    case NominationPools_Raw:
        out.Value = &pbgear.NominationPools_Bouncer_Raw {
            Raw: to_NominationPools_Raw(field),
        }
    case NominationPools_Address32:
        out.Value = &pbgear.NominationPools_Bouncer_Address32 {
            Address32: to_NominationPools_Address32(field),
        }
    case NominationPools_Address20:
        out.Value = &pbgear.NominationPools_Bouncer_Address20 {
            Address20: to_NominationPools_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_ChillCall(fields []*registry.DecodedField) *pbgear.NominationPools_ChillCall {
    out := &pbgear.NominationPools_ChillCall{}
        out.PoolId = fields[0].Value.(uint32),
    return out
}  
 
func to_NominationPools_ClaimCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimCommissionCall {
    out := &pbgear.NominationPools_ClaimCommissionCall{}
        out.PoolId = fields[0].Value.(uint32),
    return out
}  
 
func to_NominationPools_ClaimPayoutCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutCall {
    out := &pbgear.NominationPools_ClaimPayoutCall{}
    return out
} 
func to_NominationPools_ClaimPayoutOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutOtherCall {
    out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
        out.Other = to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_CompactString(fields []*registry.DecodedField) *pbgear.NominationPools_CompactString {
    out := &pbgear.NominationPools_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_NominationPools_CompactTupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
    out := &pbgear.NominationPools_CompactTupleNull{}
        out.Value = to_NominationPools_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_CreateCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateCall {
    out := &pbgear.NominationPools_CreateCall{}
        out.Amount = to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Root = to_NominationPools_Root(fields[1].Value.([]*registry.DecodedField)),
        out.Nominator = to_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField)),
        out.Bouncer = to_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_NominationPools_CreateWithPoolIdCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateWithPoolIdCall {
    out := &pbgear.NominationPools_CreateWithPoolIdCall{}
        out.Amount = to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Root = to_NominationPools_Root(fields[1].Value.([]*registry.DecodedField)),
        out.Nominator = to_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField)),
        out.Bouncer = to_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField)),
        out.PoolId = fields[4].Value.(uint32),
    return out
}  
  
  
  
  
 
func to_NominationPools_Destroying(fields []*registry.DecodedField) *pbgear.NominationPools_Destroying {
    out := &pbgear.NominationPools_Destroying{}
    return out
} 
func to_NominationPools_Extra(fields []*registry.DecodedField) *pbgear.NominationPools_Extra {
    out := &pbgear.NominationPools_Extra{}
        
    switch NominationPools_Value {
    case NominationPools_FreeBalance:
        out.Value = &pbgear.NominationPools_Extra_FreeBalance {
            FreeBalance: to_NominationPools_FreeBalance(field),
        }
    case NominationPools_Rewards:
        out.Value = &pbgear.NominationPools_Extra_Rewards {
            Rewards: to_NominationPools_Rewards(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_FreeBalance(fields []*registry.DecodedField) *pbgear.NominationPools_FreeBalance {
    out := &pbgear.NominationPools_FreeBalance{}
        out.Value_0 = fields[0].Value.(string),
    return out
}  
 
func to_NominationPools_GlobalMaxCommission(fields []*registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
    out := &pbgear.NominationPools_GlobalMaxCommission{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_GlobalMaxCommission_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_GlobalMaxCommission_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_GlobalMaxCommission_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_Id(fields []*registry.DecodedField) *pbgear.NominationPools_Id {
    out := &pbgear.NominationPools_Id{}
        out.Value_0 = to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_Index(fields []*registry.DecodedField) *pbgear.NominationPools_Index {
    out := &pbgear.NominationPools_Index{}
        out.Value_0 = to_NominationPools_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_JoinCall(fields []*registry.DecodedField) *pbgear.NominationPools_JoinCall {
    out := &pbgear.NominationPools_JoinCall{}
        out.Amount = to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.PoolId = fields[1].Value.(uint32),
    return out
}  
  
 
func to_NominationPools_MaxMembers(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembers {
    out := &pbgear.NominationPools_MaxMembers{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_MaxMembers_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_MaxMembers_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_MaxMembers_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_MaxMembersPerPool(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
    out := &pbgear.NominationPools_MaxMembersPerPool{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_MaxMembersPerPool_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_MaxMembersPerPool_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_MaxMembersPerPool_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_MaxPools(fields []*registry.DecodedField) *pbgear.NominationPools_MaxPools {
    out := &pbgear.NominationPools_MaxPools{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_MaxPools_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_MaxPools_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_MaxPools_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_Member(fields []*registry.DecodedField) *pbgear.NominationPools_Member {
    out := &pbgear.NominationPools_Member{}
        
    switch NominationPools_Value {
    case NominationPools_Id:
        out.Value = &pbgear.NominationPools_Member_Id {
            Id: to_NominationPools_Id(field),
        }
    case NominationPools_Index:
        out.Value = &pbgear.NominationPools_Member_Index {
            Index: to_NominationPools_Index(field),
        }
    case NominationPools_Raw:
        out.Value = &pbgear.NominationPools_Member_Raw {
            Raw: to_NominationPools_Raw(field),
        }
    case NominationPools_Address32:
        out.Value = &pbgear.NominationPools_Member_Address32 {
            Address32: to_NominationPools_Address32(field),
        }
    case NominationPools_Address20:
        out.Value = &pbgear.NominationPools_Member_Address20 {
            Address20: to_NominationPools_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_MemberAccount(fields []*registry.DecodedField) *pbgear.NominationPools_MemberAccount {
    out := &pbgear.NominationPools_MemberAccount{}
        
    switch NominationPools_Value {
    case NominationPools_Id:
        out.Value = &pbgear.NominationPools_MemberAccount_Id {
            Id: to_NominationPools_Id(field),
        }
    case NominationPools_Index:
        out.Value = &pbgear.NominationPools_MemberAccount_Index {
            Index: to_NominationPools_Index(field),
        }
    case NominationPools_Raw:
        out.Value = &pbgear.NominationPools_MemberAccount_Raw {
            Raw: to_NominationPools_Raw(field),
        }
    case NominationPools_Address32:
        out.Value = &pbgear.NominationPools_MemberAccount_Address32 {
            Address32: to_NominationPools_Address32(field),
        }
    case NominationPools_Address20:
        out.Value = &pbgear.NominationPools_MemberAccount_Address20 {
            Address20: to_NominationPools_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_MinCreateBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
    out := &pbgear.NominationPools_MinCreateBond{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_MinCreateBond_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_MinCreateBond_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_MinCreateBond_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_MinJoinBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
    out := &pbgear.NominationPools_MinJoinBond{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_MinJoinBond_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_MinJoinBond_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_MinJoinBond_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_NewBouncer(fields []*registry.DecodedField) *pbgear.NominationPools_NewBouncer {
    out := &pbgear.NominationPools_NewBouncer{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_NewBouncer_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_NewBouncer_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_NewBouncer_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_NewNominator(fields []*registry.DecodedField) *pbgear.NominationPools_NewNominator {
    out := &pbgear.NominationPools_NewNominator{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_NewNominator_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_NewNominator_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_NewNominator_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_NewRoot(fields []*registry.DecodedField) *pbgear.NominationPools_NewRoot {
    out := &pbgear.NominationPools_NewRoot{}
        
    switch NominationPools_Value {
    case NominationPools_Noop:
        out.Value = &pbgear.NominationPools_NewRoot_Noop {
            Noop: to_NominationPools_Noop(field),
        }
    case NominationPools_Set:
        out.Value = &pbgear.NominationPools_NewRoot_Set {
            Set: to_NominationPools_Set(field),
        }
    case NominationPools_Remove:
        out.Value = &pbgear.NominationPools_NewRoot_Remove {
            Remove: to_NominationPools_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_NominateCall(fields []*registry.DecodedField) *pbgear.NominationPools_NominateCall {
    out := &pbgear.NominationPools_NominateCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.Validators = to_repeated_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_Nominator(fields []*registry.DecodedField) *pbgear.NominationPools_Nominator {
    out := &pbgear.NominationPools_Nominator{}
        
    switch NominationPools_Value {
    case NominationPools_Id:
        out.Value = &pbgear.NominationPools_Nominator_Id {
            Id: to_NominationPools_Id(field),
        }
    case NominationPools_Index:
        out.Value = &pbgear.NominationPools_Nominator_Index {
            Index: to_NominationPools_Index(field),
        }
    case NominationPools_Raw:
        out.Value = &pbgear.NominationPools_Nominator_Raw {
            Raw: to_NominationPools_Raw(field),
        }
    case NominationPools_Address32:
        out.Value = &pbgear.NominationPools_Nominator_Address32 {
            Address32: to_NominationPools_Address32(field),
        }
    case NominationPools_Address20:
        out.Value = &pbgear.NominationPools_Nominator_Address20 {
            Address20: to_NominationPools_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_Noop(fields []*registry.DecodedField) *pbgear.NominationPools_Noop {
    out := &pbgear.NominationPools_Noop{}
    return out
} 
func to_NominationPools_Open(fields []*registry.DecodedField) *pbgear.NominationPools_Open {
    out := &pbgear.NominationPools_Open{}
    return out
} 
func to_NominationPools_PalletNominationPoolsCommissionChangeRate(fields []*registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
    out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
        out.MaxIncrease = to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
        out.MinDelay = fields[1].Value.(uint32),
    return out
}  
  
 
func to_NominationPools_Permission(fields []*registry.DecodedField) *pbgear.NominationPools_Permission {
    out := &pbgear.NominationPools_Permission{}
        
    switch NominationPools_Value {
    case NominationPools_Permissioned:
        out.Value = &pbgear.NominationPools_Permission_Permissioned {
            Permissioned: to_NominationPools_Permissioned(field),
        }
    case NominationPools_PermissionlessCompound:
        out.Value = &pbgear.NominationPools_Permission_PermissionlessCompound {
            PermissionlessCompound: to_NominationPools_PermissionlessCompound(field),
        }
    case NominationPools_PermissionlessWithdraw:
        out.Value = &pbgear.NominationPools_Permission_PermissionlessWithdraw {
            PermissionlessWithdraw: to_NominationPools_PermissionlessWithdraw(field),
        }
    case NominationPools_PermissionlessAll:
        out.Value = &pbgear.NominationPools_Permission_PermissionlessAll {
            PermissionlessAll: to_NominationPools_PermissionlessAll(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_Permissioned(fields []*registry.DecodedField) *pbgear.NominationPools_Permissioned {
    out := &pbgear.NominationPools_Permissioned{}
    return out
} 
func to_NominationPools_PermissionlessAll(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessAll {
    out := &pbgear.NominationPools_PermissionlessAll{}
    return out
} 
func to_NominationPools_PermissionlessCompound(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessCompound {
    out := &pbgear.NominationPools_PermissionlessCompound{}
    return out
} 
func to_NominationPools_PermissionlessWithdraw(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessWithdraw {
    out := &pbgear.NominationPools_PermissionlessWithdraw{}
    return out
} 
func to_NominationPools_PoolWithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
    out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.NumSlashingSpans = fields[1].Value.(uint32),
    return out
}  
  
 
func to_NominationPools_Raw(fields []*registry.DecodedField) *pbgear.NominationPools_Raw {
    out := &pbgear.NominationPools_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_Remove(fields []*registry.DecodedField) *pbgear.NominationPools_Remove {
    out := &pbgear.NominationPools_Remove{}
    return out
} 
func to_NominationPools_Rewards(fields []*registry.DecodedField) *pbgear.NominationPools_Rewards {
    out := &pbgear.NominationPools_Rewards{}
    return out
} 
func to_NominationPools_Root(fields []*registry.DecodedField) *pbgear.NominationPools_Root {
    out := &pbgear.NominationPools_Root{}
        
    switch NominationPools_Value {
    case NominationPools_Id:
        out.Value = &pbgear.NominationPools_Root_Id {
            Id: to_NominationPools_Id(field),
        }
    case NominationPools_Index:
        out.Value = &pbgear.NominationPools_Root_Index {
            Index: to_NominationPools_Index(field),
        }
    case NominationPools_Raw:
        out.Value = &pbgear.NominationPools_Root_Raw {
            Raw: to_NominationPools_Raw(field),
        }
    case NominationPools_Address32:
        out.Value = &pbgear.NominationPools_Root_Address32 {
            Address32: to_NominationPools_Address32(field),
        }
    case NominationPools_Address20:
        out.Value = &pbgear.NominationPools_Root_Address20 {
            Address20: to_NominationPools_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_Set(fields []*registry.DecodedField) *pbgear.NominationPools_Set {
    out := &pbgear.NominationPools_Set{}
        out.Value_0 = to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_SetClaimPermissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetClaimPermissionCall {
    out := &pbgear.NominationPools_SetClaimPermissionCall{}
        out.Permission = to_NominationPools_Permission(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_SetCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionCall {
    out := &pbgear.NominationPools_SetCommissionCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.NewCommission = to_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_SetCommissionChangeRateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionChangeRateCall {
    out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.ChangeRate = to_NominationPools_PalletNominationPoolsCommissionChangeRate(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_SetCommissionMaxCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionMaxCall {
    out := &pbgear.NominationPools_SetCommissionMaxCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.MaxCommission = to_NominationPools_SpArithmeticPerThingsPerbill(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_SetConfigsCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetConfigsCall {
    out := &pbgear.NominationPools_SetConfigsCall{}
        out.MinJoinBond = to_NominationPools_MinJoinBond(fields[0].Value.([]*registry.DecodedField)),
        out.MinCreateBond = to_NominationPools_MinCreateBond(fields[1].Value.([]*registry.DecodedField)),
        out.MaxPools = to_NominationPools_MaxPools(fields[2].Value.([]*registry.DecodedField)),
        out.MaxMembers = to_NominationPools_MaxMembers(fields[3].Value.([]*registry.DecodedField)),
        out.MaxMembersPerPool = to_NominationPools_MaxMembersPerPool(fields[4].Value.([]*registry.DecodedField)),
        out.GlobalMaxCommission = to_NominationPools_GlobalMaxCommission(fields[5].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
  
 
func to_NominationPools_SetMetadataCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetMetadataCall {
    out := &pbgear.NominationPools_SetMetadataCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.Metadata = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_SetStateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetStateCall {
    out := &pbgear.NominationPools_SetStateCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.State = to_NominationPools_State(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
    out := &pbgear.NominationPools_SpArithmeticPerThingsPerbill{}
        out.Value_0 = fields[0].Value.(uint32),
    return out
}  
 
func to_NominationPools_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_SpCoreCryptoAccountId32 {
    out := &pbgear.NominationPools_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_NominationPools_State(fields []*registry.DecodedField) *pbgear.NominationPools_State {
    out := &pbgear.NominationPools_State{}
        
    switch NominationPools_Value {
    case NominationPools_Open:
        out.Value = &pbgear.NominationPools_State_Open {
            Open: to_NominationPools_Open(field),
        }
    case NominationPools_Blocked:
        out.Value = &pbgear.NominationPools_State_Blocked {
            Blocked: to_NominationPools_Blocked(field),
        }
    case NominationPools_Destroying:
        out.Value = &pbgear.NominationPools_State_Destroying {
            Destroying: to_NominationPools_Destroying(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_NominationPools_TupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_TupleNull {
    out := &pbgear.NominationPools_TupleNull{}
    return out
} 
func to_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
        out.Value_0 = to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_UnbondCall(fields []*registry.DecodedField) *pbgear.NominationPools_UnbondCall {
    out := &pbgear.NominationPools_UnbondCall{}
        out.MemberAccount = to_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField)),
        out.UnbondingPoints = to_NominationPools_CompactString(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_NominationPools_UpdateRolesCall(fields []*registry.DecodedField) *pbgear.NominationPools_UpdateRolesCall {
    out := &pbgear.NominationPools_UpdateRolesCall{}
        out.PoolId = fields[0].Value.(uint32),
        out.NewRoot = to_NominationPools_NewRoot(fields[1].Value.([]*registry.DecodedField)),
        out.NewNominator = to_NominationPools_NewNominator(fields[2].Value.([]*registry.DecodedField)),
        out.NewBouncer = to_NominationPools_NewBouncer(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_NominationPools_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_WithdrawUnbondedCall {
    out := &pbgear.NominationPools_WithdrawUnbondedCall{}
        out.MemberAccount = to_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField)),
        out.NumSlashingSpans = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Preimage_EnsureUpdatedCall(fields []*registry.DecodedField) *pbgear.Preimage_EnsureUpdatedCall {
    out := &pbgear.Preimage_EnsureUpdatedCall{}
        out.Hashes = to_repeated_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Preimage_NotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_NotePreimageCall {
    out := &pbgear.Preimage_NotePreimageCall{}
        out.Bytes = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Preimage_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
    out := &pbgear.Preimage_PrimitiveTypesH256{}
        out.Hash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Preimage_RequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_RequestPreimageCall {
    out := &pbgear.Preimage_RequestPreimageCall{}
        out.Hash = to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Preimage_UnnotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnnotePreimageCall {
    out := &pbgear.Preimage_UnnotePreimageCall{}
        out.Hash = to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Preimage_UnrequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnrequestPreimageCall {
    out := &pbgear.Preimage_UnrequestPreimageCall{}
        out.Hash = to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_AddProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_AddProxyCall {
    out := &pbgear.Proxy_AddProxyCall{}
        out.Delegate = to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
        out.ProxyType = to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
        out.Delay = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Proxy_Address20(fields []*registry.DecodedField) *pbgear.Proxy_Address20 {
    out := &pbgear.Proxy_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_Address32(fields []*registry.DecodedField) *pbgear.Proxy_Address32 {
    out := &pbgear.Proxy_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_AnnounceCall(fields []*registry.DecodedField) *pbgear.Proxy_AnnounceCall {
    out := &pbgear.Proxy_AnnounceCall{}
        out.Real = to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
        out.CallHash = to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Proxy_Any(fields []*registry.DecodedField) *pbgear.Proxy_Any {
    out := &pbgear.Proxy_Any{}
    return out
} 
func to_Proxy_CancelProxy(fields []*registry.DecodedField) *pbgear.Proxy_CancelProxy {
    out := &pbgear.Proxy_CancelProxy{}
    return out
} 
func to_Proxy_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
    out := &pbgear.Proxy_CompactTupleNull{}
        out.Value = to_Proxy_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_CompactUint32(fields []*registry.DecodedField) *pbgear.Proxy_CompactUint32 {
    out := &pbgear.Proxy_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Proxy_CreatePureCall(fields []*registry.DecodedField) *pbgear.Proxy_CreatePureCall {
    out := &pbgear.Proxy_CreatePureCall{}
        out.ProxyType = to_Proxy_ProxyType(fields[0].Value.([]*registry.DecodedField)),
        out.Delay = fields[1].Value.(uint32),
        out.Index = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Proxy_Delegate(fields []*registry.DecodedField) *pbgear.Proxy_Delegate {
    out := &pbgear.Proxy_Delegate{}
        
    switch Proxy_Value {
    case Proxy_Id:
        out.Value = &pbgear.Proxy_Delegate_Id {
            Id: to_Proxy_Id(field),
        }
    case Proxy_Index:
        out.Value = &pbgear.Proxy_Delegate_Index {
            Index: to_Proxy_Index(field),
        }
    case Proxy_Raw:
        out.Value = &pbgear.Proxy_Delegate_Raw {
            Raw: to_Proxy_Raw(field),
        }
    case Proxy_Address32:
        out.Value = &pbgear.Proxy_Delegate_Address32 {
            Address32: to_Proxy_Address32(field),
        }
    case Proxy_Address20:
        out.Value = &pbgear.Proxy_Delegate_Address20 {
            Address20: to_Proxy_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Proxy_ForceProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ForceProxyType {
    out := &pbgear.Proxy_ForceProxyType{}
        
    switch Proxy_Value {
    case Proxy_Any:
        out.Value = &pbgear.Proxy_ForceProxyType_Any {
            Any: to_Proxy_Any(field),
        }
    case Proxy_NonTransfer:
        out.Value = &pbgear.Proxy_ForceProxyType_NonTransfer {
            NonTransfer: to_Proxy_NonTransfer(field),
        }
    case Proxy_Governance:
        out.Value = &pbgear.Proxy_ForceProxyType_Governance {
            Governance: to_Proxy_Governance(field),
        }
    case Proxy_Staking:
        out.Value = &pbgear.Proxy_ForceProxyType_Staking {
            Staking: to_Proxy_Staking(field),
        }
    case Proxy_IdentityJudgement:
        out.Value = &pbgear.Proxy_ForceProxyType_IdentityJudgement {
            IdentityJudgement: to_Proxy_IdentityJudgement(field),
        }
    case Proxy_CancelProxy:
        out.Value = &pbgear.Proxy_ForceProxyType_CancelProxy {
            CancelProxy: to_Proxy_CancelProxy(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Proxy_Governance(fields []*registry.DecodedField) *pbgear.Proxy_Governance {
    out := &pbgear.Proxy_Governance{}
    return out
} 
func to_Proxy_Id(fields []*registry.DecodedField) *pbgear.Proxy_Id {
    out := &pbgear.Proxy_Id{}
        out.Value_0 = to_Proxy_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_IdentityJudgement(fields []*registry.DecodedField) *pbgear.Proxy_IdentityJudgement {
    out := &pbgear.Proxy_IdentityJudgement{}
    return out
} 
func to_Proxy_Index(fields []*registry.DecodedField) *pbgear.Proxy_Index {
    out := &pbgear.Proxy_Index{}
        out.Value_0 = to_Proxy_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_KillPureCall(fields []*registry.DecodedField) *pbgear.Proxy_KillPureCall {
    out := &pbgear.Proxy_KillPureCall{}
        out.Spawner = to_Proxy_Spawner(fields[0].Value.([]*registry.DecodedField)),
        out.ProxyType = to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
        out.Index = fields[2].Value.(uint32),
        out.Height = to_Proxy_CompactUint32(fields[3].Value.([]*registry.DecodedField)),
        out.ExtIndex = to_Proxy_CompactUint32(fields[4].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
 
func to_Proxy_NonTransfer(fields []*registry.DecodedField) *pbgear.Proxy_NonTransfer {
    out := &pbgear.Proxy_NonTransfer{}
    return out
} 
func to_Proxy_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
    out := &pbgear.Proxy_PrimitiveTypesH256{}
        out.CallHash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_ProxyAnnouncedCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyAnnouncedCall {
    out := &pbgear.Proxy_ProxyAnnouncedCall{}
        out.Delegate = to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
        out.Real = to_Proxy_Real(fields[1].Value.([]*registry.DecodedField)),
        out.ForceProxyType = to_optional_Proxy_ForceProxyType(fields[2].Value.([]*registry.DecodedField)),
        
    switch Proxy_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
  
 
func to_Proxy_ProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyCall {
    out := &pbgear.Proxy_ProxyCall{}
        out.Real = to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
        out.ForceProxyType = to_optional_Proxy_ForceProxyType(fields[1].Value.([]*registry.DecodedField)),
        
    switch Proxy_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Proxy_ProxyCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Proxy_ProxyCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Proxy_ProxyCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Proxy_ProxyCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Proxy_ProxyCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Proxy_ProxyCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Proxy_ProxyCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Proxy_ProxyCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Proxy_ProxyCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Proxy_ProxyCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Proxy_ProxyCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Proxy_ProxyCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Proxy_ProxyCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Proxy_ProxyCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Proxy_ProxyCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Proxy_ProxyCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Proxy_ProxyCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Proxy_ProxyCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Proxy_ProxyCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Proxy_ProxyCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Proxy_ProxyCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Proxy_ProxyCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Proxy_ProxyCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Proxy_ProxyCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Proxy_ProxyCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Proxy_ProxyCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Proxy_ProxyCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Proxy_ProxyCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Proxy_ProxyCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Proxy_ProxyCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Proxy_ProxyCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Proxy_ProxyCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Proxy_ProxyCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Proxy_ProxyCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Proxy_ProxyCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Proxy_ProxyCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Proxy_ProxyCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Proxy_ProxyCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Proxy_ProxyCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Proxy_ProxyCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
 
func to_Proxy_ProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ProxyType {
    out := &pbgear.Proxy_ProxyType{}
        
    switch Proxy_Value {
    case Proxy_Any:
        out.Value = &pbgear.Proxy_ProxyType_Any {
            Any: to_Proxy_Any(field),
        }
    case Proxy_NonTransfer:
        out.Value = &pbgear.Proxy_ProxyType_NonTransfer {
            NonTransfer: to_Proxy_NonTransfer(field),
        }
    case Proxy_Governance:
        out.Value = &pbgear.Proxy_ProxyType_Governance {
            Governance: to_Proxy_Governance(field),
        }
    case Proxy_Staking:
        out.Value = &pbgear.Proxy_ProxyType_Staking {
            Staking: to_Proxy_Staking(field),
        }
    case Proxy_IdentityJudgement:
        out.Value = &pbgear.Proxy_ProxyType_IdentityJudgement {
            IdentityJudgement: to_Proxy_IdentityJudgement(field),
        }
    case Proxy_CancelProxy:
        out.Value = &pbgear.Proxy_ProxyType_CancelProxy {
            CancelProxy: to_Proxy_CancelProxy(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Proxy_Raw(fields []*registry.DecodedField) *pbgear.Proxy_Raw {
    out := &pbgear.Proxy_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_Real(fields []*registry.DecodedField) *pbgear.Proxy_Real {
    out := &pbgear.Proxy_Real{}
        
    switch Proxy_Value {
    case Proxy_Id:
        out.Value = &pbgear.Proxy_Real_Id {
            Id: to_Proxy_Id(field),
        }
    case Proxy_Index:
        out.Value = &pbgear.Proxy_Real_Index {
            Index: to_Proxy_Index(field),
        }
    case Proxy_Raw:
        out.Value = &pbgear.Proxy_Real_Raw {
            Raw: to_Proxy_Raw(field),
        }
    case Proxy_Address32:
        out.Value = &pbgear.Proxy_Real_Address32 {
            Address32: to_Proxy_Address32(field),
        }
    case Proxy_Address20:
        out.Value = &pbgear.Proxy_Real_Address20 {
            Address20: to_Proxy_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Proxy_RejectAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RejectAnnouncementCall {
    out := &pbgear.Proxy_RejectAnnouncementCall{}
        out.Delegate = to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
        out.CallHash = to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Proxy_RemoveAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveAnnouncementCall {
    out := &pbgear.Proxy_RemoveAnnouncementCall{}
        out.Real = to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
        out.CallHash = to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Proxy_RemoveProxiesCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxiesCall {
    out := &pbgear.Proxy_RemoveProxiesCall{}
    return out
} 
func to_Proxy_RemoveProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxyCall {
    out := &pbgear.Proxy_RemoveProxyCall{}
        out.Delegate = to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
        out.ProxyType = to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
        out.Delay = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Proxy_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Proxy_SpCoreCryptoAccountId32 {
    out := &pbgear.Proxy_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Proxy_Spawner(fields []*registry.DecodedField) *pbgear.Proxy_Spawner {
    out := &pbgear.Proxy_Spawner{}
        
    switch Proxy_Value {
    case Proxy_Id:
        out.Value = &pbgear.Proxy_Spawner_Id {
            Id: to_Proxy_Id(field),
        }
    case Proxy_Index:
        out.Value = &pbgear.Proxy_Spawner_Index {
            Index: to_Proxy_Index(field),
        }
    case Proxy_Raw:
        out.Value = &pbgear.Proxy_Spawner_Raw {
            Raw: to_Proxy_Raw(field),
        }
    case Proxy_Address32:
        out.Value = &pbgear.Proxy_Spawner_Address32 {
            Address32: to_Proxy_Address32(field),
        }
    case Proxy_Address20:
        out.Value = &pbgear.Proxy_Spawner_Address20 {
            Address20: to_Proxy_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Proxy_Staking(fields []*registry.DecodedField) *pbgear.Proxy_Staking {
    out := &pbgear.Proxy_Staking{}
    return out
} 
func to_Proxy_TupleNull(fields []*registry.DecodedField) *pbgear.Proxy_TupleNull {
    out := &pbgear.Proxy_TupleNull{}
    return out
} 
func to_Referenda_After(fields []*registry.DecodedField) *pbgear.Referenda_After {
    out := &pbgear.Referenda_After{}
        out.Value_0 = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_At(fields []*registry.DecodedField) *pbgear.Referenda_At {
    out := &pbgear.Referenda_At{}
        out.Value_0 = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec {
    out := &pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_CancelCall(fields []*registry.DecodedField) *pbgear.Referenda_CancelCall {
    out := &pbgear.Referenda_CancelCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
    out := &pbgear.Referenda_EnactmentMoment{}
        
    switch Referenda_Value {
    case Referenda_At:
        out.Value = &pbgear.Referenda_EnactmentMoment_At {
            At: to_Referenda_At(field),
        }
    case Referenda_After:
        out.Value = &pbgear.Referenda_EnactmentMoment_After {
            After: to_Referenda_After(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Referenda_Inline(fields []*registry.DecodedField) *pbgear.Referenda_Inline {
    out := &pbgear.Referenda_Inline{}
        out.Value_0 = to_Referenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_KillCall(fields []*registry.DecodedField) *pbgear.Referenda_KillCall {
    out := &pbgear.Referenda_KillCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_Legacy(fields []*registry.DecodedField) *pbgear.Referenda_Legacy {
    out := &pbgear.Referenda_Legacy{}
        out.Hash = to_Referenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_Lookup(fields []*registry.DecodedField) *pbgear.Referenda_Lookup {
    out := &pbgear.Referenda_Lookup{}
        out.Hash = to_Referenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
        out.Len = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Referenda_None(fields []*registry.DecodedField) *pbgear.Referenda_None {
    out := &pbgear.Referenda_None{}
    return out
} 
func to_Referenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.Referenda_NudgeReferendumCall {
    out := &pbgear.Referenda_NudgeReferendumCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.Referenda_OneFewerDecidingCall {
    out := &pbgear.Referenda_OneFewerDecidingCall{}
        out.Track = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_Origins(fields []*registry.DecodedField) *pbgear.Referenda_Origins {
    out := &pbgear.Referenda_Origins{}
        out.Value_0 = to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_PlaceDecisionDepositCall {
    out := &pbgear.Referenda_PlaceDecisionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
    out := &pbgear.Referenda_PrimitiveTypesH256{}
        out.Hash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_Proposal(fields []*registry.DecodedField) *pbgear.Referenda_Proposal {
    out := &pbgear.Referenda_Proposal{}
        
    switch Referenda_Value {
    case Referenda_Legacy:
        out.Value = &pbgear.Referenda_Proposal_Legacy {
            Legacy: to_Referenda_Legacy(field),
        }
    case Referenda_Inline:
        out.Value = &pbgear.Referenda_Proposal_Inline {
            Inline: to_Referenda_Inline(field),
        }
    case Referenda_Lookup:
        out.Value = &pbgear.Referenda_Proposal_Lookup {
            Lookup: to_Referenda_Lookup(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Referenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
    out := &pbgear.Referenda_ProposalOrigin{}
        
    switch Referenda_Value {
    case Referenda_System:
        out.Value = &pbgear.Referenda_ProposalOrigin_system {
            system: to_Referenda_System(field),
        }
    case Referenda_Origins:
        out.Value = &pbgear.Referenda_ProposalOrigin_Origins {
            Origins: to_Referenda_Origins(field),
        }
    case Referenda_Void:
        out.Value = &pbgear.Referenda_ProposalOrigin_Void {
            Void: to_Referenda_Void(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Referenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundDecisionDepositCall {
    out := &pbgear.Referenda_RefundDecisionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundSubmissionDepositCall {
    out := &pbgear.Referenda_RefundSubmissionDepositCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Referenda_Root(fields []*registry.DecodedField) *pbgear.Referenda_Root {
    out := &pbgear.Referenda_Root{}
    return out
} 
func to_Referenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.Referenda_SetMetadataCall {
    out := &pbgear.Referenda_SetMetadataCall{}
        out.Index = fields[0].Value.(uint32),
        out.MaybeHash = to_optional_Referenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Referenda_Signed(fields []*registry.DecodedField) *pbgear.Referenda_Signed {
    out := &pbgear.Referenda_Signed{}
        out.Value_0 = to_Referenda_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Referenda_SpCoreCryptoAccountId32 {
    out := &pbgear.Referenda_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_SubmitCall(fields []*registry.DecodedField) *pbgear.Referenda_SubmitCall {
    out := &pbgear.Referenda_SubmitCall{}
        out.ProposalOrigin = to_Referenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField)),
        out.Proposal = to_Referenda_Proposal(fields[1].Value.([]*registry.DecodedField)),
        out.EnactmentMoment = to_Referenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_Referenda_System(fields []*registry.DecodedField) *pbgear.Referenda_System {
    out := &pbgear.Referenda_System{}
        out.Value_0 = to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Referenda_Value0(fields []*registry.DecodedField) *pbgear.Referenda_Value0 {
    out := &pbgear.Referenda_Value0{}
        
    switch Referenda_Value {
    case Referenda_Root:
        out.Value = &pbgear.Referenda_Value0_Root {
            Root: to_Referenda_Root(field),
        }
    case Referenda_Signed:
        out.Value = &pbgear.Referenda_Value0_Signed {
            Signed: to_Referenda_Signed(field),
        }
    case Referenda_None:
        out.Value = &pbgear.Referenda_Value0_None {
            None: to_Referenda_None(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Referenda_Void(fields []*registry.DecodedField) *pbgear.Referenda_Void {
    out := &pbgear.Referenda_Void{}
        out.Value_0 = to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Scheduler_CancelCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelCall {
    out := &pbgear.Scheduler_CancelCall{}
        out.When = fields[0].Value.(uint32),
        out.Index = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Scheduler_CancelNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelNamedCall {
    out := &pbgear.Scheduler_CancelNamedCall{}
        out.Id = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Scheduler_ScheduleAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleAfterCall {
    out := &pbgear.Scheduler_ScheduleAfterCall{}
        out.After = fields[0].Value.(uint32),
        out.MaybePeriodic = to_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField)),
        out.Priority = fields[2].Value.(uint32),
        
    switch Scheduler_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
  
 
func to_Scheduler_ScheduleCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleCall {
    out := &pbgear.Scheduler_ScheduleCall{}
        out.When = fields[0].Value.(uint32),
        out.MaybePeriodic = to_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField)),
        out.Priority = fields[2].Value.(uint32),
        
    switch Scheduler_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Scheduler_ScheduleCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Scheduler_ScheduleCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
  
 
func to_Scheduler_ScheduleNamedAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedAfterCall {
    out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
        out.Id = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.After = fields[1].Value.(uint32),
        out.MaybePeriodic = to_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField)),
        out.Priority = fields[3].Value.(uint32),
        
    switch Scheduler_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
  
  
 
func to_Scheduler_ScheduleNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedCall {
    out := &pbgear.Scheduler_ScheduleNamedCall{}
        out.Id = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.When = fields[1].Value.(uint32),
        out.MaybePeriodic = to_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField)),
        out.Priority = fields[3].Value.(uint32),
        
    switch Scheduler_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
  
  
  
 
func to_Scheduler_TupleUint32Uint32(fields []*registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
    out := &pbgear.Scheduler_TupleUint32Uint32{}
        out.Value_0 = fields[0].Value.(uint32),
        out.Value_1 = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Session_PalletImOnlineSr25519AppSr25519Public(fields []*registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
    out := &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{}
        out.ImOnline = to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_PurgeKeysCall(fields []*registry.DecodedField) *pbgear.Session_PurgeKeysCall {
    out := &pbgear.Session_PurgeKeysCall{}
    return out
} 
func to_Session_SetKeysCall(fields []*registry.DecodedField) *pbgear.Session_SetKeysCall {
    out := &pbgear.Session_SetKeysCall{}
        out.Keys = to_Session_VaraRuntimeSessionKeys(fields[0].Value.([]*registry.DecodedField)),
        out.Proof = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Session_SpAuthorityDiscoveryAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
    out := &pbgear.Session_SpAuthorityDiscoveryAppPublic{}
        out.AuthorityDiscovery = to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
    out := &pbgear.Session_SpConsensusBabeAppPublic{}
        out.Babe = to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
    out := &pbgear.Session_SpConsensusGrandpaAppPublic{}
        out.Grandpa = to_Session_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
    out := &pbgear.Session_SpCoreEd25519Public{}
        out.Grandpa = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
    out := &pbgear.Session_SpCoreSr25519Public{}
        out.Babe = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Session_VaraRuntimeSessionKeys(fields []*registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
    out := &pbgear.Session_VaraRuntimeSessionKeys{}
        out.Babe = to_Session_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField)),
        out.Grandpa = to_Session_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField)),
        out.ImOnline = to_Session_PalletImOnlineSr25519AppSr25519Public(fields[2].Value.([]*registry.DecodedField)),
        out.AuthorityDiscovery = to_Session_SpAuthorityDiscoveryAppPublic(fields[3].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
 
func to_StakingRewards_Address20(fields []*registry.DecodedField) *pbgear.StakingRewards_Address20 {
    out := &pbgear.StakingRewards_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_Address32(fields []*registry.DecodedField) *pbgear.StakingRewards_Address32 {
    out := &pbgear.StakingRewards_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_AlignSupplyCall(fields []*registry.DecodedField) *pbgear.StakingRewards_AlignSupplyCall {
    out := &pbgear.StakingRewards_AlignSupplyCall{}
        out.Target = fields[0].Value.(string),
    return out
}  
 
func to_StakingRewards_CompactTupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
    out := &pbgear.StakingRewards_CompactTupleNull{}
        out.Value = to_StakingRewards_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_ForceRefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_ForceRefillCall {
    out := &pbgear.StakingRewards_ForceRefillCall{}
        out.From = to_StakingRewards_From(fields[0].Value.([]*registry.DecodedField)),
        out.Value = fields[1].Value.(string),
    return out
}  
  
 
func to_StakingRewards_From(fields []*registry.DecodedField) *pbgear.StakingRewards_From {
    out := &pbgear.StakingRewards_From{}
        
    switch StakingRewards_Value {
    case StakingRewards_Id:
        out.Value = &pbgear.StakingRewards_From_Id {
            Id: to_StakingRewards_Id(field),
        }
    case StakingRewards_Index:
        out.Value = &pbgear.StakingRewards_From_Index {
            Index: to_StakingRewards_Index(field),
        }
    case StakingRewards_Raw:
        out.Value = &pbgear.StakingRewards_From_Raw {
            Raw: to_StakingRewards_Raw(field),
        }
    case StakingRewards_Address32:
        out.Value = &pbgear.StakingRewards_From_Address32 {
            Address32: to_StakingRewards_Address32(field),
        }
    case StakingRewards_Address20:
        out.Value = &pbgear.StakingRewards_From_Address20 {
            Address20: to_StakingRewards_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_StakingRewards_Id(fields []*registry.DecodedField) *pbgear.StakingRewards_Id {
    out := &pbgear.StakingRewards_Id{}
        out.Value_0 = to_StakingRewards_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_Index(fields []*registry.DecodedField) *pbgear.StakingRewards_Index {
    out := &pbgear.StakingRewards_Index{}
        out.Value_0 = to_StakingRewards_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_Raw(fields []*registry.DecodedField) *pbgear.StakingRewards_Raw {
    out := &pbgear.StakingRewards_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_RefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_RefillCall {
    out := &pbgear.StakingRewards_RefillCall{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_StakingRewards_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.StakingRewards_SpCoreCryptoAccountId32 {
    out := &pbgear.StakingRewards_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_StakingRewards_To(fields []*registry.DecodedField) *pbgear.StakingRewards_To {
    out := &pbgear.StakingRewards_To{}
        
    switch StakingRewards_Value {
    case StakingRewards_Id:
        out.Value = &pbgear.StakingRewards_To_Id {
            Id: to_StakingRewards_Id(field),
        }
    case StakingRewards_Index:
        out.Value = &pbgear.StakingRewards_To_Index {
            Index: to_StakingRewards_Index(field),
        }
    case StakingRewards_Raw:
        out.Value = &pbgear.StakingRewards_To_Raw {
            Raw: to_StakingRewards_Raw(field),
        }
    case StakingRewards_Address32:
        out.Value = &pbgear.StakingRewards_To_Address32 {
            Address32: to_StakingRewards_Address32(field),
        }
    case StakingRewards_Address20:
        out.Value = &pbgear.StakingRewards_To_Address20 {
            Address20: to_StakingRewards_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_StakingRewards_TupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_TupleNull {
    out := &pbgear.StakingRewards_TupleNull{}
    return out
} 
func to_StakingRewards_WithdrawCall(fields []*registry.DecodedField) *pbgear.StakingRewards_WithdrawCall {
    out := &pbgear.StakingRewards_WithdrawCall{}
        out.To = to_StakingRewards_To(fields[0].Value.([]*registry.DecodedField)),
        out.Value = fields[1].Value.(string),
    return out
}  
  
 
func to_Staking_Account(fields []*registry.DecodedField) *pbgear.Staking_Account {
    out := &pbgear.Staking_Account{}
        out.Value_0 = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Address20(fields []*registry.DecodedField) *pbgear.Staking_Address20 {
    out := &pbgear.Staking_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Address32(fields []*registry.DecodedField) *pbgear.Staking_Address32 {
    out := &pbgear.Staking_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_BondCall(fields []*registry.DecodedField) *pbgear.Staking_BondCall {
    out := &pbgear.Staking_BondCall{}
        out.Value = to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Payee = to_Staking_Payee(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Staking_BondExtraCall(fields []*registry.DecodedField) *pbgear.Staking_BondExtraCall {
    out := &pbgear.Staking_BondExtraCall{}
        out.MaxAdditional = to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_CancelDeferredSlashCall(fields []*registry.DecodedField) *pbgear.Staking_CancelDeferredSlashCall {
    out := &pbgear.Staking_CancelDeferredSlashCall{}
        out.Era = fields[0].Value.(uint32),
        out.SlashIndices = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Staking_ChillCall(fields []*registry.DecodedField) *pbgear.Staking_ChillCall {
    out := &pbgear.Staking_ChillCall{}
    return out
} 
func to_Staking_ChillOtherCall(fields []*registry.DecodedField) *pbgear.Staking_ChillOtherCall {
    out := &pbgear.Staking_ChillOtherCall{}
        out.Controller = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_ChillThreshold(fields []*registry.DecodedField) *pbgear.Staking_ChillThreshold {
    out := &pbgear.Staking_ChillThreshold{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_ChillThreshold_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_ChillThreshold_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_ChillThreshold_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_CompactSpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
    out := &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{}
        out.Value = to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_CompactString(fields []*registry.DecodedField) *pbgear.Staking_CompactString {
    out := &pbgear.Staking_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_Staking_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Staking_CompactTupleNull {
    out := &pbgear.Staking_CompactTupleNull{}
        out.Value = to_Staking_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_CompactUint32(fields []*registry.DecodedField) *pbgear.Staking_CompactUint32 {
    out := &pbgear.Staking_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Staking_Controller(fields []*registry.DecodedField) *pbgear.Staking_Controller {
    out := &pbgear.Staking_Controller{}
    return out
} 
func to_Staking_ForceApplyMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_ForceApplyMinCommissionCall {
    out := &pbgear.Staking_ForceApplyMinCommissionCall{}
        out.ValidatorStash = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_ForceNewEraAlwaysCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraAlwaysCall {
    out := &pbgear.Staking_ForceNewEraAlwaysCall{}
    return out
} 
func to_Staking_ForceNewEraCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraCall {
    out := &pbgear.Staking_ForceNewEraCall{}
    return out
} 
func to_Staking_ForceNoErasCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNoErasCall {
    out := &pbgear.Staking_ForceNoErasCall{}
    return out
} 
func to_Staking_ForceUnstakeCall(fields []*registry.DecodedField) *pbgear.Staking_ForceUnstakeCall {
    out := &pbgear.Staking_ForceUnstakeCall{}
        out.Stash = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.NumSlashingSpans = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Staking_Id(fields []*registry.DecodedField) *pbgear.Staking_Id {
    out := &pbgear.Staking_Id{}
        out.Value_0 = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_IncreaseValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_IncreaseValidatorCountCall {
    out := &pbgear.Staking_IncreaseValidatorCountCall{}
        out.Additional = to_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Index(fields []*registry.DecodedField) *pbgear.Staking_Index {
    out := &pbgear.Staking_Index{}
        out.Value_0 = to_Staking_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_KickCall(fields []*registry.DecodedField) *pbgear.Staking_KickCall {
    out := &pbgear.Staking_KickCall{}
        out.Who = to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_MaxNominatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
    out := &pbgear.Staking_MaxNominatorCount{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_MaxNominatorCount_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_MaxNominatorCount_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_MaxNominatorCount_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_MaxValidatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
    out := &pbgear.Staking_MaxValidatorCount{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_MaxValidatorCount_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_MaxValidatorCount_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_MaxValidatorCount_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_MinCommission(fields []*registry.DecodedField) *pbgear.Staking_MinCommission {
    out := &pbgear.Staking_MinCommission{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_MinCommission_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_MinCommission_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_MinCommission_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_MinNominatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinNominatorBond {
    out := &pbgear.Staking_MinNominatorBond{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_MinNominatorBond_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_MinNominatorBond_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_MinNominatorBond_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_MinValidatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinValidatorBond {
    out := &pbgear.Staking_MinValidatorBond{}
        
    switch Staking_Value {
    case Staking_Noop:
        out.Value = &pbgear.Staking_MinValidatorBond_Noop {
            Noop: to_Staking_Noop(field),
        }
    case Staking_Set:
        out.Value = &pbgear.Staking_MinValidatorBond_Set {
            Set: to_Staking_Set(field),
        }
    case Staking_Remove:
        out.Value = &pbgear.Staking_MinValidatorBond_Remove {
            Remove: to_Staking_Remove(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_NominateCall(fields []*registry.DecodedField) *pbgear.Staking_NominateCall {
    out := &pbgear.Staking_NominateCall{}
        out.Targets = to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_None(fields []*registry.DecodedField) *pbgear.Staking_None {
    out := &pbgear.Staking_None{}
    return out
} 
func to_Staking_Noop(fields []*registry.DecodedField) *pbgear.Staking_Noop {
    out := &pbgear.Staking_Noop{}
    return out
} 
func to_Staking_PalletStakingValidatorPrefs(fields []*registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
    out := &pbgear.Staking_PalletStakingValidatorPrefs{}
        out.Commission = to_Staking_CompactSpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
        out.Blocked = fields[1].Value.(bool),
    return out
}  
  
 
func to_Staking_Payee(fields []*registry.DecodedField) *pbgear.Staking_Payee {
    out := &pbgear.Staking_Payee{}
        
    switch Staking_Value {
    case Staking_Staked:
        out.Value = &pbgear.Staking_Payee_Staked {
            Staked: to_Staking_Staked(field),
        }
    case Staking_Stash:
        out.Value = &pbgear.Staking_Payee_Stash {
            Stash: to_Staking_Stash(field),
        }
    case Staking_Controller:
        out.Value = &pbgear.Staking_Payee_Controller {
            Controller: to_Staking_Controller(field),
        }
    case Staking_Account:
        out.Value = &pbgear.Staking_Payee_Account {
            Account: to_Staking_Account(field),
        }
    case Staking_None:
        out.Value = &pbgear.Staking_Payee_None {
            None: to_Staking_None(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_PayoutStakersCall(fields []*registry.DecodedField) *pbgear.Staking_PayoutStakersCall {
    out := &pbgear.Staking_PayoutStakersCall{}
        out.ValidatorStash = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.Era = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Staking_Raw(fields []*registry.DecodedField) *pbgear.Staking_Raw {
    out := &pbgear.Staking_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_ReapStashCall(fields []*registry.DecodedField) *pbgear.Staking_ReapStashCall {
    out := &pbgear.Staking_ReapStashCall{}
        out.Stash = to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
        out.NumSlashingSpans = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Staking_RebondCall(fields []*registry.DecodedField) *pbgear.Staking_RebondCall {
    out := &pbgear.Staking_RebondCall{}
        out.Value = to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Remove(fields []*registry.DecodedField) *pbgear.Staking_Remove {
    out := &pbgear.Staking_Remove{}
    return out
} 
func to_Staking_ScaleValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_ScaleValidatorCountCall {
    out := &pbgear.Staking_ScaleValidatorCountCall{}
        out.Factor = to_Staking_SpArithmeticPerThingsPercent(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Set(fields []*registry.DecodedField) *pbgear.Staking_Set {
    out := &pbgear.Staking_Set{}
        out.Value_0 = to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SetControllerCall(fields []*registry.DecodedField) *pbgear.Staking_SetControllerCall {
    out := &pbgear.Staking_SetControllerCall{}
    return out
} 
func to_Staking_SetInvulnerablesCall(fields []*registry.DecodedField) *pbgear.Staking_SetInvulnerablesCall {
    out := &pbgear.Staking_SetInvulnerablesCall{}
        out.Invulnerables = to_repeated_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SetMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_SetMinCommissionCall {
    out := &pbgear.Staking_SetMinCommissionCall{}
        out.New = to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SetPayeeCall(fields []*registry.DecodedField) *pbgear.Staking_SetPayeeCall {
    out := &pbgear.Staking_SetPayeeCall{}
        out.Payee = to_Staking_Payee(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SetStakingConfigsCall(fields []*registry.DecodedField) *pbgear.Staking_SetStakingConfigsCall {
    out := &pbgear.Staking_SetStakingConfigsCall{}
        out.MinNominatorBond = to_Staking_MinNominatorBond(fields[0].Value.([]*registry.DecodedField)),
        out.MinValidatorBond = to_Staking_MinValidatorBond(fields[1].Value.([]*registry.DecodedField)),
        out.MaxNominatorCount = to_Staking_MaxNominatorCount(fields[2].Value.([]*registry.DecodedField)),
        out.MaxValidatorCount = to_Staking_MaxValidatorCount(fields[3].Value.([]*registry.DecodedField)),
        out.ChillThreshold = to_Staking_ChillThreshold(fields[4].Value.([]*registry.DecodedField)),
        out.MinCommission = to_Staking_MinCommission(fields[5].Value.([]*registry.DecodedField)),
    return out
}  
  
  
  
  
  
 
func to_Staking_SetValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_SetValidatorCountCall {
    out := &pbgear.Staking_SetValidatorCountCall{}
        out.New = to_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
    out := &pbgear.Staking_SpArithmeticPerThingsPerbill{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Staking_SpArithmeticPerThingsPercent(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
    out := &pbgear.Staking_SpArithmeticPerThingsPercent{}
        out.Factor = fields[0].Value.(uint32),
    return out
}  
 
func to_Staking_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
    out := &pbgear.Staking_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_SpRuntimeMultiaddressMultiAddress(fields []*registry.DecodedField) *pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
    out := &pbgear.Staking_SpRuntimeMultiaddressMultiAddress{}
        out.Targets = to_Staking_Targets(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_Staked(fields []*registry.DecodedField) *pbgear.Staking_Staked {
    out := &pbgear.Staking_Staked{}
    return out
} 
func to_Staking_Stash(fields []*registry.DecodedField) *pbgear.Staking_Stash {
    out := &pbgear.Staking_Stash{}
    return out
} 
func to_Staking_Targets(fields []*registry.DecodedField) *pbgear.Staking_Targets {
    out := &pbgear.Staking_Targets{}
        
    switch Staking_Value {
    case Staking_Id:
        out.Value = &pbgear.Staking_Targets_Id {
            Id: to_Staking_Id(field),
        }
    case Staking_Index:
        out.Value = &pbgear.Staking_Targets_Index {
            Index: to_Staking_Index(field),
        }
    case Staking_Raw:
        out.Value = &pbgear.Staking_Targets_Raw {
            Raw: to_Staking_Raw(field),
        }
    case Staking_Address32:
        out.Value = &pbgear.Staking_Targets_Address32 {
            Address32: to_Staking_Address32(field),
        }
    case Staking_Address20:
        out.Value = &pbgear.Staking_Targets_Address20 {
            Address20: to_Staking_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Staking_TupleNull(fields []*registry.DecodedField) *pbgear.Staking_TupleNull {
    out := &pbgear.Staking_TupleNull{}
    return out
} 
func to_Staking_UnbondCall(fields []*registry.DecodedField) *pbgear.Staking_UnbondCall {
    out := &pbgear.Staking_UnbondCall{}
        out.Value = to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_ValidateCall(fields []*registry.DecodedField) *pbgear.Staking_ValidateCall {
    out := &pbgear.Staking_ValidateCall{}
        out.Prefs = to_Staking_PalletStakingValidatorPrefs(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Staking_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.Staking_WithdrawUnbondedCall {
    out := &pbgear.Staking_WithdrawUnbondedCall{}
        out.NumSlashingSpans = fields[0].Value.(uint32),
    return out
}  
 
func to_System_KillPrefixCall(fields []*registry.DecodedField) *pbgear.System_KillPrefixCall {
    out := &pbgear.System_KillPrefixCall{}
        out.Prefix = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Subkeys = fields[1].Value.(uint32),
    return out
}  
  
 
func to_System_KillStorageCall(fields []*registry.DecodedField) *pbgear.System_KillStorageCall {
    out := &pbgear.System_KillStorageCall{}
        out.Keys = to_repeated_System_SystemKeysList(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_RemarkCall(fields []*registry.DecodedField) *pbgear.System_RemarkCall {
    out := &pbgear.System_RemarkCall{}
        out.Remark = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_RemarkWithEventCall(fields []*registry.DecodedField) *pbgear.System_RemarkWithEventCall {
    out := &pbgear.System_RemarkWithEventCall{}
        out.Remark = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_SetCodeCall(fields []*registry.DecodedField) *pbgear.System_SetCodeCall {
    out := &pbgear.System_SetCodeCall{}
        out.Code = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_SetCodeWithoutChecksCall(fields []*registry.DecodedField) *pbgear.System_SetCodeWithoutChecksCall {
    out := &pbgear.System_SetCodeWithoutChecksCall{}
        out.Code = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_SetHeapPagesCall(fields []*registry.DecodedField) *pbgear.System_SetHeapPagesCall {
    out := &pbgear.System_SetHeapPagesCall{}
        out.Pages = fields[0].Value.(uint64),
    return out
}  
 
func to_System_SetStorageCall(fields []*registry.DecodedField) *pbgear.System_SetStorageCall {
    out := &pbgear.System_SetStorageCall{}
        out.Items = to_repeated_System_TupleSystemItemsListSystemItemsList(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_SystemKeysList(fields []*registry.DecodedField) *pbgear.System_SystemKeysList {
    out := &pbgear.System_SystemKeysList{}
        out.Keys = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_System_TupleSystemItemsListSystemItemsList(fields []*registry.DecodedField) *pbgear.System_TupleSystemItemsListSystemItemsList {
    out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
        out.Value_1 = to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Timestamp_CompactUint64(fields []*registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
    out := &pbgear.Timestamp_CompactUint64{}
        out.Value = fields[0].Value.(uint64),
    return out
}  
 
func to_Timestamp_SetCall(fields []*registry.DecodedField) *pbgear.Timestamp_SetCall {
    out := &pbgear.Timestamp_SetCall{}
        out.Now = to_Timestamp_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_Address20(fields []*registry.DecodedField) *pbgear.Treasury_Address20 {
    out := &pbgear.Treasury_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_Address32(fields []*registry.DecodedField) *pbgear.Treasury_Address32 {
    out := &pbgear.Treasury_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_ApproveProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_ApproveProposalCall {
    out := &pbgear.Treasury_ApproveProposalCall{}
        out.ProposalId = to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_Beneficiary(fields []*registry.DecodedField) *pbgear.Treasury_Beneficiary {
    out := &pbgear.Treasury_Beneficiary{}
        
    switch Treasury_Value {
    case Treasury_Id:
        out.Value = &pbgear.Treasury_Beneficiary_Id {
            Id: to_Treasury_Id(field),
        }
    case Treasury_Index:
        out.Value = &pbgear.Treasury_Beneficiary_Index {
            Index: to_Treasury_Index(field),
        }
    case Treasury_Raw:
        out.Value = &pbgear.Treasury_Beneficiary_Raw {
            Raw: to_Treasury_Raw(field),
        }
    case Treasury_Address32:
        out.Value = &pbgear.Treasury_Beneficiary_Address32 {
            Address32: to_Treasury_Address32(field),
        }
    case Treasury_Address20:
        out.Value = &pbgear.Treasury_Beneficiary_Address20 {
            Address20: to_Treasury_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Treasury_CheckStatusCall(fields []*registry.DecodedField) *pbgear.Treasury_CheckStatusCall {
    out := &pbgear.Treasury_CheckStatusCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Treasury_CompactString(fields []*registry.DecodedField) *pbgear.Treasury_CompactString {
    out := &pbgear.Treasury_CompactString{}
        out.Value = fields[0].Value.(string),
    return out
}  
 
func to_Treasury_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
    out := &pbgear.Treasury_CompactTupleNull{}
        out.Value = to_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_CompactUint32(fields []*registry.DecodedField) *pbgear.Treasury_CompactUint32 {
    out := &pbgear.Treasury_CompactUint32{}
        out.Value = fields[0].Value.(uint32),
    return out
}  
 
func to_Treasury_Id(fields []*registry.DecodedField) *pbgear.Treasury_Id {
    out := &pbgear.Treasury_Id{}
        out.Value_0 = to_Treasury_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_Index(fields []*registry.DecodedField) *pbgear.Treasury_Index {
    out := &pbgear.Treasury_Index{}
        out.Value_0 = to_Treasury_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_PayoutCall(fields []*registry.DecodedField) *pbgear.Treasury_PayoutCall {
    out := &pbgear.Treasury_PayoutCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Treasury_ProposeSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_ProposeSpendCall {
    out := &pbgear.Treasury_ProposeSpendCall{}
        out.Value = to_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Beneficiary = to_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Treasury_Raw(fields []*registry.DecodedField) *pbgear.Treasury_Raw {
    out := &pbgear.Treasury_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_RejectProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_RejectProposalCall {
    out := &pbgear.Treasury_RejectProposalCall{}
        out.ProposalId = to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_RemoveApprovalCall(fields []*registry.DecodedField) *pbgear.Treasury_RemoveApprovalCall {
    out := &pbgear.Treasury_RemoveApprovalCall{}
        out.ProposalId = to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
    out := &pbgear.Treasury_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Treasury_SpendCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendCall {
    out := &pbgear.Treasury_SpendCall{}
        out.AssetKind = to_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField)),
        out.Amount = to_Treasury_CompactString(fields[1].Value.([]*registry.DecodedField)),
        out.Beneficiary = to_Treasury_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField)),
        out.ValidFrom = fields[3].Value.(uint32),
    return out
}  
  
  
  
 
func to_Treasury_SpendLocalCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendLocalCall {
    out := &pbgear.Treasury_SpendLocalCall{}
        out.Amount = to_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField)),
        out.Beneficiary = to_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Treasury_TupleNull(fields []*registry.DecodedField) *pbgear.Treasury_TupleNull {
    out := &pbgear.Treasury_TupleNull{}
    return out
} 
func to_Treasury_VoidSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_VoidSpendCall {
    out := &pbgear.Treasury_VoidSpendCall{}
        out.Index = fields[0].Value.(uint32),
    return out
}  
 
func to_Utility_AsDerivativeCall(fields []*registry.DecodedField) *pbgear.Utility_AsDerivativeCall {
    out := &pbgear.Utility_AsDerivativeCall{}
        out.Index = fields[0].Value.(uint32),
        
    switch Utility_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Utility_AsDerivativeCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Utility_AsDerivativeCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
 
func to_Utility_AsOrigin(fields []*registry.DecodedField) *pbgear.Utility_AsOrigin {
    out := &pbgear.Utility_AsOrigin{}
        
    switch Utility_Value {
    case Utility_System:
        out.Value = &pbgear.Utility_AsOrigin_system {
            system: to_Utility_System(field),
        }
    case Utility_Origins:
        out.Value = &pbgear.Utility_AsOrigin_Origins {
            Origins: to_Utility_Origins(field),
        }
    case Utility_Void:
        out.Value = &pbgear.Utility_AsOrigin_Void {
            Void: to_Utility_Void(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Utility_BatchAllCall(fields []*registry.DecodedField) *pbgear.Utility_BatchAllCall {
    out := &pbgear.Utility_BatchAllCall{}
        out.Calls = to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_BatchCall(fields []*registry.DecodedField) *pbgear.Utility_BatchCall {
    out := &pbgear.Utility_BatchCall{}
        out.Calls = to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_CompactUint64(fields []*registry.DecodedField) *pbgear.Utility_CompactUint64 {
    out := &pbgear.Utility_CompactUint64{}
        out.Value = fields[0].Value.(uint64),
    return out
}  
 
func to_Utility_DispatchAsCall(fields []*registry.DecodedField) *pbgear.Utility_DispatchAsCall {
    out := &pbgear.Utility_DispatchAsCall{}
        out.AsOrigin = to_Utility_AsOrigin(fields[0].Value.([]*registry.DecodedField)),
        
    switch Utility_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Utility_DispatchAsCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Utility_DispatchAsCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Utility_DispatchAsCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Utility_DispatchAsCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Utility_DispatchAsCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Utility_DispatchAsCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Utility_DispatchAsCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Utility_DispatchAsCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Utility_DispatchAsCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Utility_DispatchAsCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Utility_DispatchAsCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Utility_DispatchAsCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
  
 
func to_Utility_ForceBatchCall(fields []*registry.DecodedField) *pbgear.Utility_ForceBatchCall {
    out := &pbgear.Utility_ForceBatchCall{}
        out.Calls = to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_None(fields []*registry.DecodedField) *pbgear.Utility_None {
    out := &pbgear.Utility_None{}
    return out
} 
func to_Utility_Origins(fields []*registry.DecodedField) *pbgear.Utility_Origins {
    out := &pbgear.Utility_Origins{}
        out.Value_0 = to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_Root(fields []*registry.DecodedField) *pbgear.Utility_Root {
    out := &pbgear.Utility_Root{}
    return out
} 
func to_Utility_Signed(fields []*registry.DecodedField) *pbgear.Utility_Signed {
    out := &pbgear.Utility_Signed{}
        out.Value_0 = to_Utility_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Utility_SpCoreCryptoAccountId32 {
    out := &pbgear.Utility_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
    out := &pbgear.Utility_SpWeightsWeightV2Weight{}
        out.RefTime = to_Utility_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
        out.ProofSize = to_Utility_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Utility_System(fields []*registry.DecodedField) *pbgear.Utility_System {
    out := &pbgear.Utility_System{}
        out.Value_0 = to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_Value0(fields []*registry.DecodedField) *pbgear.Utility_Value0 {
    out := &pbgear.Utility_Value0{}
        
    switch Utility_Value {
    case Utility_Root:
        out.Value = &pbgear.Utility_Value0_Root {
            Root: to_Utility_Root(field),
        }
    case Utility_Signed:
        out.Value = &pbgear.Utility_Value0_Signed {
            Signed: to_Utility_Signed(field),
        }
    case Utility_None:
        out.Value = &pbgear.Utility_Value0_None {
            None: to_Utility_None(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Utility_VaraRuntimeRuntimeCall(fields []*registry.DecodedField) *pbgear.Utility_VaraRuntimeRuntimeCall {
    out := &pbgear.Utility_VaraRuntimeRuntimeCall{}
        
    switch Utility_Calls {
    case System_RemarkCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Utility_Void(fields []*registry.DecodedField) *pbgear.Utility_Void {
    out := &pbgear.Utility_Void{}
        out.Value_0 = to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Utility_WithWeightCall(fields []*registry.DecodedField) *pbgear.Utility_WithWeightCall {
    out := &pbgear.Utility_WithWeightCall{}
        
    switch Utility_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Utility_WithWeightCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Utility_WithWeightCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Utility_WithWeightCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Utility_WithWeightCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Utility_WithWeightCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Utility_WithWeightCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Utility_WithWeightCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Utility_WithWeightCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Utility_WithWeightCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Utility_WithWeightCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Utility_WithWeightCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Utility_WithWeightCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Utility_WithWeightCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Utility_WithWeightCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Utility_WithWeightCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Utility_WithWeightCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Utility_WithWeightCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Utility_WithWeightCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Utility_WithWeightCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Utility_WithWeightCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Utility_WithWeightCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Utility_WithWeightCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Utility_WithWeightCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Utility_WithWeightCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Utility_WithWeightCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Utility_WithWeightCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Utility_WithWeightCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Utility_WithWeightCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Utility_WithWeightCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Utility_WithWeightCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Utility_WithWeightCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Utility_WithWeightCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Utility_WithWeightCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Utility_WithWeightCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Utility_WithWeightCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Utility_WithWeightCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Utility_WithWeightCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Utility_WithWeightCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Utility_WithWeightCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Utility_WithWeightCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Utility_WithWeightCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Utility_WithWeightCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
        out.Weight = to_Utility_SpWeightsWeightV2Weight(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Vesting_Address20(fields []*registry.DecodedField) *pbgear.Vesting_Address20 {
    out := &pbgear.Vesting_Address20{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_Address32(fields []*registry.DecodedField) *pbgear.Vesting_Address32 {
    out := &pbgear.Vesting_Address32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
    out := &pbgear.Vesting_CompactTupleNull{}
        out.Value = to_Vesting_TupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_ForceVestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_ForceVestedTransferCall {
    out := &pbgear.Vesting_ForceVestedTransferCall{}
        out.Source = to_Vesting_Source(fields[0].Value.([]*registry.DecodedField)),
        out.Target = to_Vesting_Target(fields[1].Value.([]*registry.DecodedField)),
        out.Schedule = to_Vesting_PalletVestingVestingInfoVestingInfo(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_Vesting_Id(fields []*registry.DecodedField) *pbgear.Vesting_Id {
    out := &pbgear.Vesting_Id{}
        out.Value_0 = to_Vesting_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_Index(fields []*registry.DecodedField) *pbgear.Vesting_Index {
    out := &pbgear.Vesting_Index{}
        out.Value_0 = to_Vesting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_MergeSchedulesCall(fields []*registry.DecodedField) *pbgear.Vesting_MergeSchedulesCall {
    out := &pbgear.Vesting_MergeSchedulesCall{}
        out.Schedule_1Index = fields[0].Value.(uint32),
        out.Schedule_2Index = fields[1].Value.(uint32),
    return out
}  
  
 
func to_Vesting_PalletVestingVestingInfoVestingInfo(fields []*registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
    out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
        out.Locked = fields[0].Value.(string),
        out.PerBlock = fields[1].Value.(string),
        out.StartingBlock = fields[2].Value.(uint32),
    return out
}  
  
  
 
func to_Vesting_Raw(fields []*registry.DecodedField) *pbgear.Vesting_Raw {
    out := &pbgear.Vesting_Raw{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_Source(fields []*registry.DecodedField) *pbgear.Vesting_Source {
    out := &pbgear.Vesting_Source{}
        
    switch Vesting_Value {
    case Vesting_Id:
        out.Value = &pbgear.Vesting_Source_Id {
            Id: to_Vesting_Id(field),
        }
    case Vesting_Index:
        out.Value = &pbgear.Vesting_Source_Index {
            Index: to_Vesting_Index(field),
        }
    case Vesting_Raw:
        out.Value = &pbgear.Vesting_Source_Raw {
            Raw: to_Vesting_Raw(field),
        }
    case Vesting_Address32:
        out.Value = &pbgear.Vesting_Source_Address32 {
            Address32: to_Vesting_Address32(field),
        }
    case Vesting_Address20:
        out.Value = &pbgear.Vesting_Source_Address20 {
            Address20: to_Vesting_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Vesting_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Vesting_SpCoreCryptoAccountId32 {
    out := &pbgear.Vesting_SpCoreCryptoAccountId32{}
        out.Value_0 = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_Target(fields []*registry.DecodedField) *pbgear.Vesting_Target {
    out := &pbgear.Vesting_Target{}
        
    switch Vesting_Value {
    case Vesting_Id:
        out.Value = &pbgear.Vesting_Target_Id {
            Id: to_Vesting_Id(field),
        }
    case Vesting_Index:
        out.Value = &pbgear.Vesting_Target_Index {
            Index: to_Vesting_Index(field),
        }
    case Vesting_Raw:
        out.Value = &pbgear.Vesting_Target_Raw {
            Raw: to_Vesting_Raw(field),
        }
    case Vesting_Address32:
        out.Value = &pbgear.Vesting_Target_Address32 {
            Address32: to_Vesting_Address32(field),
        }
    case Vesting_Address20:
        out.Value = &pbgear.Vesting_Target_Address20 {
            Address20: to_Vesting_Address20(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Vesting_TupleNull(fields []*registry.DecodedField) *pbgear.Vesting_TupleNull {
    out := &pbgear.Vesting_TupleNull{}
    return out
} 
func to_Vesting_VestCall(fields []*registry.DecodedField) *pbgear.Vesting_VestCall {
    out := &pbgear.Vesting_VestCall{}
    return out
} 
func to_Vesting_VestOtherCall(fields []*registry.DecodedField) *pbgear.Vesting_VestOtherCall {
    out := &pbgear.Vesting_VestOtherCall{}
        out.Target = to_Vesting_Target(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Vesting_VestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_VestedTransferCall {
    out := &pbgear.Vesting_VestedTransferCall{}
        out.Target = to_Vesting_Target(fields[0].Value.([]*registry.DecodedField)),
        out.Schedule = to_Vesting_PalletVestingVestingInfoVestingInfo(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Whitelist_CompactUint64(fields []*registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
    out := &pbgear.Whitelist_CompactUint64{}
        out.Value = fields[0].Value.(uint64),
    return out
}  
 
func to_Whitelist_DispatchWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallCall {
    out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
        out.CallHash = to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
        out.CallEncodedLen = fields[1].Value.(uint32),
        out.CallWeightWitness = to_Whitelist_SpWeightsWeightV2Weight(fields[2].Value.([]*registry.DecodedField)),
    return out
}  
  
  
 
func to_Whitelist_DispatchWhitelistedCallWithPreimageCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
    out := &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{}
        
    switch Whitelist_Call {
    case System_RemarkCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_remark {
            System_remark: to_System_RemarkCall(field),
        }
    case System_SetHeapPagesCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_set_heap_pages {
            System_set_heap_pages: to_System_SetHeapPagesCall(field),
        }
    case System_SetCodeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_set_code {
            System_set_code: to_System_SetCodeCall(field),
        }
    case System_SetCodeWithoutChecksCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_set_code_without_checks {
            System_set_code_without_checks: to_System_SetCodeWithoutChecksCall(field),
        }
    case System_SetStorageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_set_storage {
            System_set_storage: to_System_SetStorageCall(field),
        }
    case System_KillStorageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_kill_storage {
            System_kill_storage: to_System_KillStorageCall(field),
        }
    case System_KillPrefixCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_kill_prefix {
            System_kill_prefix: to_System_KillPrefixCall(field),
        }
    case System_RemarkWithEventCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System_remark_with_event {
            System_remark_with_event: to_System_RemarkWithEventCall(field),
        }
    case Timestamp_SetCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Timestamp_set {
            Timestamp_set: to_Timestamp_SetCall(field),
        }
    case Babe_ReportEquivocationCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Babe_report_equivocation {
            Babe_report_equivocation: to_Babe_ReportEquivocationCall(field),
        }
    case Babe_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Babe_report_equivocation_unsigned {
            Babe_report_equivocation_unsigned: to_Babe_ReportEquivocationUnsignedCall(field),
        }
    case Babe_PlanConfigChangeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Babe_plan_config_change {
            Babe_plan_config_change: to_Babe_PlanConfigChangeCall(field),
        }
    case Grandpa_ReportEquivocationCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Grandpa_report_equivocation {
            Grandpa_report_equivocation: to_Grandpa_ReportEquivocationCall(field),
        }
    case Grandpa_ReportEquivocationUnsignedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Grandpa_report_equivocation_unsigned {
            Grandpa_report_equivocation_unsigned: to_Grandpa_ReportEquivocationUnsignedCall(field),
        }
    case Grandpa_NoteStalledCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Grandpa_note_stalled {
            Grandpa_note_stalled: to_Grandpa_NoteStalledCall(field),
        }
    case Balances_TransferAllowDeathCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_transfer_allow_death {
            Balances_transfer_allow_death: to_Balances_TransferAllowDeathCall(field),
        }
    case Balances_ForceTransferCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_force_transfer {
            Balances_force_transfer: to_Balances_ForceTransferCall(field),
        }
    case Balances_TransferKeepAliveCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_transfer_keep_alive {
            Balances_transfer_keep_alive: to_Balances_TransferKeepAliveCall(field),
        }
    case Balances_TransferAllCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_transfer_all {
            Balances_transfer_all: to_Balances_TransferAllCall(field),
        }
    case Balances_ForceUnreserveCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_force_unreserve {
            Balances_force_unreserve: to_Balances_ForceUnreserveCall(field),
        }
    case Balances_UpgradeAccountsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_upgrade_accounts {
            Balances_upgrade_accounts: to_Balances_UpgradeAccountsCall(field),
        }
    case Balances_ForceSetBalanceCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances_force_set_balance {
            Balances_force_set_balance: to_Balances_ForceSetBalanceCall(field),
        }
    case Vesting_VestCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting_vest {
            Vesting_vest: to_Vesting_VestCall(field),
        }
    case Vesting_VestOtherCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting_vest_other {
            Vesting_vest_other: to_Vesting_VestOtherCall(field),
        }
    case Vesting_VestedTransferCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting_vested_transfer {
            Vesting_vested_transfer: to_Vesting_VestedTransferCall(field),
        }
    case Vesting_ForceVestedTransferCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting_force_vested_transfer {
            Vesting_force_vested_transfer: to_Vesting_ForceVestedTransferCall(field),
        }
    case Vesting_MergeSchedulesCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting_merge_schedules {
            Vesting_merge_schedules: to_Vesting_MergeSchedulesCall(field),
        }
    case BagsList_RebagCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsList_rebag {
            BagsList_rebag: to_BagsList_RebagCall(field),
        }
    case BagsList_PutInFrontOfCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsList_put_in_front_of {
            BagsList_put_in_front_of: to_BagsList_PutInFrontOfCall(field),
        }
    case BagsList_PutInFrontOfOtherCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsList_put_in_front_of_other {
            BagsList_put_in_front_of_other: to_BagsList_PutInFrontOfOtherCall(field),
        }
    case ImOnline_HeartbeatCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ImOnline_heartbeat {
            ImOnline_heartbeat: to_ImOnline_HeartbeatCall(field),
        }
    case Staking_BondCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_bond {
            Staking_bond: to_Staking_BondCall(field),
        }
    case Staking_BondExtraCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_bond_extra {
            Staking_bond_extra: to_Staking_BondExtraCall(field),
        }
    case Staking_UnbondCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_unbond {
            Staking_unbond: to_Staking_UnbondCall(field),
        }
    case Staking_WithdrawUnbondedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_withdraw_unbonded {
            Staking_withdraw_unbonded: to_Staking_WithdrawUnbondedCall(field),
        }
    case Staking_ValidateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_validate {
            Staking_validate: to_Staking_ValidateCall(field),
        }
    case Staking_NominateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_nominate {
            Staking_nominate: to_Staking_NominateCall(field),
        }
    case Staking_ChillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_chill {
            Staking_chill: to_Staking_ChillCall(field),
        }
    case Staking_SetPayeeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_payee {
            Staking_set_payee: to_Staking_SetPayeeCall(field),
        }
    case Staking_SetControllerCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_controller {
            Staking_set_controller: to_Staking_SetControllerCall(field),
        }
    case Staking_SetValidatorCountCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_validator_count {
            Staking_set_validator_count: to_Staking_SetValidatorCountCall(field),
        }
    case Staking_IncreaseValidatorCountCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_increase_validator_count {
            Staking_increase_validator_count: to_Staking_IncreaseValidatorCountCall(field),
        }
    case Staking_ScaleValidatorCountCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_scale_validator_count {
            Staking_scale_validator_count: to_Staking_ScaleValidatorCountCall(field),
        }
    case Staking_ForceNoErasCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_force_no_eras {
            Staking_force_no_eras: to_Staking_ForceNoErasCall(field),
        }
    case Staking_ForceNewEraCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_force_new_era {
            Staking_force_new_era: to_Staking_ForceNewEraCall(field),
        }
    case Staking_SetInvulnerablesCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_invulnerables {
            Staking_set_invulnerables: to_Staking_SetInvulnerablesCall(field),
        }
    case Staking_ForceUnstakeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_force_unstake {
            Staking_force_unstake: to_Staking_ForceUnstakeCall(field),
        }
    case Staking_ForceNewEraAlwaysCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_force_new_era_always {
            Staking_force_new_era_always: to_Staking_ForceNewEraAlwaysCall(field),
        }
    case Staking_CancelDeferredSlashCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_cancel_deferred_slash {
            Staking_cancel_deferred_slash: to_Staking_CancelDeferredSlashCall(field),
        }
    case Staking_PayoutStakersCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_payout_stakers {
            Staking_payout_stakers: to_Staking_PayoutStakersCall(field),
        }
    case Staking_RebondCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_rebond {
            Staking_rebond: to_Staking_RebondCall(field),
        }
    case Staking_ReapStashCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_reap_stash {
            Staking_reap_stash: to_Staking_ReapStashCall(field),
        }
    case Staking_KickCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_kick {
            Staking_kick: to_Staking_KickCall(field),
        }
    case Staking_SetStakingConfigsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_staking_configs {
            Staking_set_staking_configs: to_Staking_SetStakingConfigsCall(field),
        }
    case Staking_ChillOtherCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_chill_other {
            Staking_chill_other: to_Staking_ChillOtherCall(field),
        }
    case Staking_ForceApplyMinCommissionCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_force_apply_min_commission {
            Staking_force_apply_min_commission: to_Staking_ForceApplyMinCommissionCall(field),
        }
    case Staking_SetMinCommissionCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking_set_min_commission {
            Staking_set_min_commission: to_Staking_SetMinCommissionCall(field),
        }
    case Session_SetKeysCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Session_set_keys {
            Session_set_keys: to_Session_SetKeysCall(field),
        }
    case Session_PurgeKeysCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Session_purge_keys {
            Session_purge_keys: to_Session_PurgeKeysCall(field),
        }
    case Treasury_ProposeSpendCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_propose_spend {
            Treasury_propose_spend: to_Treasury_ProposeSpendCall(field),
        }
    case Treasury_RejectProposalCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_reject_proposal {
            Treasury_reject_proposal: to_Treasury_RejectProposalCall(field),
        }
    case Treasury_ApproveProposalCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_approve_proposal {
            Treasury_approve_proposal: to_Treasury_ApproveProposalCall(field),
        }
    case Treasury_SpendLocalCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_spend_local {
            Treasury_spend_local: to_Treasury_SpendLocalCall(field),
        }
    case Treasury_RemoveApprovalCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_remove_approval {
            Treasury_remove_approval: to_Treasury_RemoveApprovalCall(field),
        }
    case Treasury_SpendCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_spend {
            Treasury_spend: to_Treasury_SpendCall(field),
        }
    case Treasury_PayoutCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_payout {
            Treasury_payout: to_Treasury_PayoutCall(field),
        }
    case Treasury_CheckStatusCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_check_status {
            Treasury_check_status: to_Treasury_CheckStatusCall(field),
        }
    case Treasury_VoidSpendCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury_void_spend {
            Treasury_void_spend: to_Treasury_VoidSpendCall(field),
        }
    case Utility_BatchCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_batch {
            Utility_batch: to_Utility_BatchCall(field),
        }
    case Utility_AsDerivativeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_as_derivative {
            Utility_as_derivative: to_Utility_AsDerivativeCall(field),
        }
    case Utility_BatchAllCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_batch_all {
            Utility_batch_all: to_Utility_BatchAllCall(field),
        }
    case Utility_DispatchAsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_dispatch_as {
            Utility_dispatch_as: to_Utility_DispatchAsCall(field),
        }
    case Utility_ForceBatchCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_force_batch {
            Utility_force_batch: to_Utility_ForceBatchCall(field),
        }
    case Utility_WithWeightCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility_with_weight {
            Utility_with_weight: to_Utility_WithWeightCall(field),
        }
    case ConvictionVoting_VoteCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_vote {
            ConvictionVoting_vote: to_ConvictionVoting_VoteCall(field),
        }
    case ConvictionVoting_DelegateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_delegate {
            ConvictionVoting_delegate: to_ConvictionVoting_DelegateCall(field),
        }
    case ConvictionVoting_UndelegateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_undelegate {
            ConvictionVoting_undelegate: to_ConvictionVoting_UndelegateCall(field),
        }
    case ConvictionVoting_UnlockCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_unlock {
            ConvictionVoting_unlock: to_ConvictionVoting_UnlockCall(field),
        }
    case ConvictionVoting_RemoveVoteCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_remove_vote {
            ConvictionVoting_remove_vote: to_ConvictionVoting_RemoveVoteCall(field),
        }
    case ConvictionVoting_RemoveOtherVoteCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting_remove_other_vote {
            ConvictionVoting_remove_other_vote: to_ConvictionVoting_RemoveOtherVoteCall(field),
        }
    case Referenda_SubmitCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_submit {
            Referenda_submit: to_Referenda_SubmitCall(field),
        }
    case Referenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_place_decision_deposit {
            Referenda_place_decision_deposit: to_Referenda_PlaceDecisionDepositCall(field),
        }
    case Referenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_refund_decision_deposit {
            Referenda_refund_decision_deposit: to_Referenda_RefundDecisionDepositCall(field),
        }
    case Referenda_CancelCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_cancel {
            Referenda_cancel: to_Referenda_CancelCall(field),
        }
    case Referenda_KillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_kill {
            Referenda_kill: to_Referenda_KillCall(field),
        }
    case Referenda_NudgeReferendumCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_nudge_referendum {
            Referenda_nudge_referendum: to_Referenda_NudgeReferendumCall(field),
        }
    case Referenda_OneFewerDecidingCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_one_fewer_deciding {
            Referenda_one_fewer_deciding: to_Referenda_OneFewerDecidingCall(field),
        }
    case Referenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_refund_submission_deposit {
            Referenda_refund_submission_deposit: to_Referenda_RefundSubmissionDepositCall(field),
        }
    case Referenda_SetMetadataCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda_set_metadata {
            Referenda_set_metadata: to_Referenda_SetMetadataCall(field),
        }
    case FellowshipCollective_AddMemberCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_add_member {
            FellowshipCollective_add_member: to_FellowshipCollective_AddMemberCall(field),
        }
    case FellowshipCollective_PromoteMemberCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_promote_member {
            FellowshipCollective_promote_member: to_FellowshipCollective_PromoteMemberCall(field),
        }
    case FellowshipCollective_DemoteMemberCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_demote_member {
            FellowshipCollective_demote_member: to_FellowshipCollective_DemoteMemberCall(field),
        }
    case FellowshipCollective_RemoveMemberCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_remove_member {
            FellowshipCollective_remove_member: to_FellowshipCollective_RemoveMemberCall(field),
        }
    case FellowshipCollective_VoteCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_vote {
            FellowshipCollective_vote: to_FellowshipCollective_VoteCall(field),
        }
    case FellowshipCollective_CleanupPollCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective_cleanup_poll {
            FellowshipCollective_cleanup_poll: to_FellowshipCollective_CleanupPollCall(field),
        }
    case FellowshipReferenda_SubmitCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_submit {
            FellowshipReferenda_submit: to_FellowshipReferenda_SubmitCall(field),
        }
    case FellowshipReferenda_PlaceDecisionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_place_decision_deposit {
            FellowshipReferenda_place_decision_deposit: to_FellowshipReferenda_PlaceDecisionDepositCall(field),
        }
    case FellowshipReferenda_RefundDecisionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_refund_decision_deposit {
            FellowshipReferenda_refund_decision_deposit: to_FellowshipReferenda_RefundDecisionDepositCall(field),
        }
    case FellowshipReferenda_CancelCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_cancel {
            FellowshipReferenda_cancel: to_FellowshipReferenda_CancelCall(field),
        }
    case FellowshipReferenda_KillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_kill {
            FellowshipReferenda_kill: to_FellowshipReferenda_KillCall(field),
        }
    case FellowshipReferenda_NudgeReferendumCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_nudge_referendum {
            FellowshipReferenda_nudge_referendum: to_FellowshipReferenda_NudgeReferendumCall(field),
        }
    case FellowshipReferenda_OneFewerDecidingCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_one_fewer_deciding {
            FellowshipReferenda_one_fewer_deciding: to_FellowshipReferenda_OneFewerDecidingCall(field),
        }
    case FellowshipReferenda_RefundSubmissionDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_refund_submission_deposit {
            FellowshipReferenda_refund_submission_deposit: to_FellowshipReferenda_RefundSubmissionDepositCall(field),
        }
    case FellowshipReferenda_SetMetadataCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda_set_metadata {
            FellowshipReferenda_set_metadata: to_FellowshipReferenda_SetMetadataCall(field),
        }
    case Whitelist_WhitelistCallCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Whitelist_whitelist_call {
            Whitelist_whitelist_call: to_Whitelist_WhitelistCallCall(field),
        }
    case Whitelist_RemoveWhitelistedCallCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Whitelist_remove_whitelisted_call {
            Whitelist_remove_whitelisted_call: to_Whitelist_RemoveWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Whitelist_dispatch_whitelisted_call {
            Whitelist_dispatch_whitelisted_call: to_Whitelist_DispatchWhitelistedCallCall(field),
        }
    case Whitelist_DispatchWhitelistedCallWithPreimageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Whitelist_dispatch_whitelisted_call_with_preimage {
            Whitelist_dispatch_whitelisted_call_with_preimage: to_Whitelist_DispatchWhitelistedCallWithPreimageCall(field),
        }
    case Scheduler_ScheduleCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_schedule {
            Scheduler_schedule: to_Scheduler_ScheduleCall(field),
        }
    case Scheduler_CancelCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_cancel {
            Scheduler_cancel: to_Scheduler_CancelCall(field),
        }
    case Scheduler_ScheduleNamedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_schedule_named {
            Scheduler_schedule_named: to_Scheduler_ScheduleNamedCall(field),
        }
    case Scheduler_CancelNamedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_cancel_named {
            Scheduler_cancel_named: to_Scheduler_CancelNamedCall(field),
        }
    case Scheduler_ScheduleAfterCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_schedule_after {
            Scheduler_schedule_after: to_Scheduler_ScheduleAfterCall(field),
        }
    case Scheduler_ScheduleNamedAfterCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler_schedule_named_after {
            Scheduler_schedule_named_after: to_Scheduler_ScheduleNamedAfterCall(field),
        }
    case Preimage_NotePreimageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage_note_preimage {
            Preimage_note_preimage: to_Preimage_NotePreimageCall(field),
        }
    case Preimage_UnnotePreimageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage_unnote_preimage {
            Preimage_unnote_preimage: to_Preimage_UnnotePreimageCall(field),
        }
    case Preimage_RequestPreimageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage_request_preimage {
            Preimage_request_preimage: to_Preimage_RequestPreimageCall(field),
        }
    case Preimage_UnrequestPreimageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage_unrequest_preimage {
            Preimage_unrequest_preimage: to_Preimage_UnrequestPreimageCall(field),
        }
    case Preimage_EnsureUpdatedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage_ensure_updated {
            Preimage_ensure_updated: to_Preimage_EnsureUpdatedCall(field),
        }
    case Identity_AddRegistrarCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_add_registrar {
            Identity_add_registrar: to_Identity_AddRegistrarCall(field),
        }
    case Identity_SetIdentityCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_set_identity {
            Identity_set_identity: to_Identity_SetIdentityCall(field),
        }
    case Identity_SetSubsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_set_subs {
            Identity_set_subs: to_Identity_SetSubsCall(field),
        }
    case Identity_ClearIdentityCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_clear_identity {
            Identity_clear_identity: to_Identity_ClearIdentityCall(field),
        }
    case Identity_RequestJudgementCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_request_judgement {
            Identity_request_judgement: to_Identity_RequestJudgementCall(field),
        }
    case Identity_CancelRequestCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_cancel_request {
            Identity_cancel_request: to_Identity_CancelRequestCall(field),
        }
    case Identity_SetFeeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_set_fee {
            Identity_set_fee: to_Identity_SetFeeCall(field),
        }
    case Identity_SetAccountIdCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_set_account_id {
            Identity_set_account_id: to_Identity_SetAccountIdCall(field),
        }
    case Identity_SetFieldsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_set_fields {
            Identity_set_fields: to_Identity_SetFieldsCall(field),
        }
    case Identity_ProvideJudgementCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_provide_judgement {
            Identity_provide_judgement: to_Identity_ProvideJudgementCall(field),
        }
    case Identity_KillIdentityCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_kill_identity {
            Identity_kill_identity: to_Identity_KillIdentityCall(field),
        }
    case Identity_AddSubCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_add_sub {
            Identity_add_sub: to_Identity_AddSubCall(field),
        }
    case Identity_RenameSubCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_rename_sub {
            Identity_rename_sub: to_Identity_RenameSubCall(field),
        }
    case Identity_RemoveSubCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_remove_sub {
            Identity_remove_sub: to_Identity_RemoveSubCall(field),
        }
    case Identity_QuitSubCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity_quit_sub {
            Identity_quit_sub: to_Identity_QuitSubCall(field),
        }
    case Proxy_ProxyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_proxy {
            Proxy_proxy: to_Proxy_ProxyCall(field),
        }
    case Proxy_AddProxyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_add_proxy {
            Proxy_add_proxy: to_Proxy_AddProxyCall(field),
        }
    case Proxy_RemoveProxyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_remove_proxy {
            Proxy_remove_proxy: to_Proxy_RemoveProxyCall(field),
        }
    case Proxy_RemoveProxiesCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_remove_proxies {
            Proxy_remove_proxies: to_Proxy_RemoveProxiesCall(field),
        }
    case Proxy_CreatePureCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_create_pure {
            Proxy_create_pure: to_Proxy_CreatePureCall(field),
        }
    case Proxy_KillPureCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_kill_pure {
            Proxy_kill_pure: to_Proxy_KillPureCall(field),
        }
    case Proxy_AnnounceCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_announce {
            Proxy_announce: to_Proxy_AnnounceCall(field),
        }
    case Proxy_RemoveAnnouncementCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_remove_announcement {
            Proxy_remove_announcement: to_Proxy_RemoveAnnouncementCall(field),
        }
    case Proxy_RejectAnnouncementCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_reject_announcement {
            Proxy_reject_announcement: to_Proxy_RejectAnnouncementCall(field),
        }
    case Proxy_ProxyAnnouncedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy_proxy_announced {
            Proxy_proxy_announced: to_Proxy_ProxyAnnouncedCall(field),
        }
    case Multisig_AsMultiThreshold1Call:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Multisig_as_multi_threshold_1 {
            Multisig_as_multi_threshold_1: to_Multisig_AsMultiThreshold1Call(field),
        }
    case Multisig_AsMultiCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Multisig_as_multi {
            Multisig_as_multi: to_Multisig_AsMultiCall(field),
        }
    case Multisig_ApproveAsMultiCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Multisig_approve_as_multi {
            Multisig_approve_as_multi: to_Multisig_ApproveAsMultiCall(field),
        }
    case Multisig_CancelAsMultiCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Multisig_cancel_as_multi {
            Multisig_cancel_as_multi: to_Multisig_CancelAsMultiCall(field),
        }
    case ElectionProviderMultiPhase_SubmitUnsignedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase_submit_unsigned {
            ElectionProviderMultiPhase_submit_unsigned: to_ElectionProviderMultiPhase_SubmitUnsignedCall(field),
        }
    case ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase_set_minimum_untrusted_score {
            ElectionProviderMultiPhase_set_minimum_untrusted_score: to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(field),
        }
    case ElectionProviderMultiPhase_SetEmergencyElectionResultCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase_set_emergency_election_result {
            ElectionProviderMultiPhase_set_emergency_election_result: to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(field),
        }
    case ElectionProviderMultiPhase_SubmitCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase_submit {
            ElectionProviderMultiPhase_submit: to_ElectionProviderMultiPhase_SubmitCall(field),
        }
    case ElectionProviderMultiPhase_GovernanceFallbackCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase_governance_fallback {
            ElectionProviderMultiPhase_governance_fallback: to_ElectionProviderMultiPhase_GovernanceFallbackCall(field),
        }
    case Bounties_ProposeBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_propose_bounty {
            Bounties_propose_bounty: to_Bounties_ProposeBountyCall(field),
        }
    case Bounties_ApproveBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_approve_bounty {
            Bounties_approve_bounty: to_Bounties_ApproveBountyCall(field),
        }
    case Bounties_ProposeCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_propose_curator {
            Bounties_propose_curator: to_Bounties_ProposeCuratorCall(field),
        }
    case Bounties_UnassignCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_unassign_curator {
            Bounties_unassign_curator: to_Bounties_UnassignCuratorCall(field),
        }
    case Bounties_AcceptCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_accept_curator {
            Bounties_accept_curator: to_Bounties_AcceptCuratorCall(field),
        }
    case Bounties_AwardBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_award_bounty {
            Bounties_award_bounty: to_Bounties_AwardBountyCall(field),
        }
    case Bounties_ClaimBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_claim_bounty {
            Bounties_claim_bounty: to_Bounties_ClaimBountyCall(field),
        }
    case Bounties_CloseBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_close_bounty {
            Bounties_close_bounty: to_Bounties_CloseBountyCall(field),
        }
    case Bounties_ExtendBountyExpiryCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties_extend_bounty_expiry {
            Bounties_extend_bounty_expiry: to_Bounties_ExtendBountyExpiryCall(field),
        }
    case ChildBounties_AddChildBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_add_child_bounty {
            ChildBounties_add_child_bounty: to_ChildBounties_AddChildBountyCall(field),
        }
    case ChildBounties_ProposeCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_propose_curator {
            ChildBounties_propose_curator: to_ChildBounties_ProposeCuratorCall(field),
        }
    case ChildBounties_AcceptCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_accept_curator {
            ChildBounties_accept_curator: to_ChildBounties_AcceptCuratorCall(field),
        }
    case ChildBounties_UnassignCuratorCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_unassign_curator {
            ChildBounties_unassign_curator: to_ChildBounties_UnassignCuratorCall(field),
        }
    case ChildBounties_AwardChildBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_award_child_bounty {
            ChildBounties_award_child_bounty: to_ChildBounties_AwardChildBountyCall(field),
        }
    case ChildBounties_ClaimChildBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_claim_child_bounty {
            ChildBounties_claim_child_bounty: to_ChildBounties_ClaimChildBountyCall(field),
        }
    case ChildBounties_CloseChildBountyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties_close_child_bounty {
            ChildBounties_close_child_bounty: to_ChildBounties_CloseChildBountyCall(field),
        }
    case NominationPools_JoinCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_join {
            NominationPools_join: to_NominationPools_JoinCall(field),
        }
    case NominationPools_BondExtraCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_bond_extra {
            NominationPools_bond_extra: to_NominationPools_BondExtraCall(field),
        }
    case NominationPools_ClaimPayoutCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_claim_payout {
            NominationPools_claim_payout: to_NominationPools_ClaimPayoutCall(field),
        }
    case NominationPools_UnbondCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_unbond {
            NominationPools_unbond: to_NominationPools_UnbondCall(field),
        }
    case NominationPools_PoolWithdrawUnbondedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_pool_withdraw_unbonded {
            NominationPools_pool_withdraw_unbonded: to_NominationPools_PoolWithdrawUnbondedCall(field),
        }
    case NominationPools_WithdrawUnbondedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_withdraw_unbonded {
            NominationPools_withdraw_unbonded: to_NominationPools_WithdrawUnbondedCall(field),
        }
    case NominationPools_CreateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_create {
            NominationPools_create: to_NominationPools_CreateCall(field),
        }
    case NominationPools_CreateWithPoolIdCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_create_with_pool_id {
            NominationPools_create_with_pool_id: to_NominationPools_CreateWithPoolIdCall(field),
        }
    case NominationPools_NominateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_nominate {
            NominationPools_nominate: to_NominationPools_NominateCall(field),
        }
    case NominationPools_SetStateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_state {
            NominationPools_set_state: to_NominationPools_SetStateCall(field),
        }
    case NominationPools_SetMetadataCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_metadata {
            NominationPools_set_metadata: to_NominationPools_SetMetadataCall(field),
        }
    case NominationPools_SetConfigsCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_configs {
            NominationPools_set_configs: to_NominationPools_SetConfigsCall(field),
        }
    case NominationPools_UpdateRolesCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_update_roles {
            NominationPools_update_roles: to_NominationPools_UpdateRolesCall(field),
        }
    case NominationPools_ChillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_chill {
            NominationPools_chill: to_NominationPools_ChillCall(field),
        }
    case NominationPools_BondExtraOtherCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_bond_extra_other {
            NominationPools_bond_extra_other: to_NominationPools_BondExtraOtherCall(field),
        }
    case NominationPools_SetClaimPermissionCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_claim_permission {
            NominationPools_set_claim_permission: to_NominationPools_SetClaimPermissionCall(field),
        }
    case NominationPools_ClaimPayoutOtherCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_claim_payout_other {
            NominationPools_claim_payout_other: to_NominationPools_ClaimPayoutOtherCall(field),
        }
    case NominationPools_SetCommissionCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_commission {
            NominationPools_set_commission: to_NominationPools_SetCommissionCall(field),
        }
    case NominationPools_SetCommissionMaxCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_commission_max {
            NominationPools_set_commission_max: to_NominationPools_SetCommissionMaxCall(field),
        }
    case NominationPools_SetCommissionChangeRateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_set_commission_change_rate {
            NominationPools_set_commission_change_rate: to_NominationPools_SetCommissionChangeRateCall(field),
        }
    case NominationPools_ClaimCommissionCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_claim_commission {
            NominationPools_claim_commission: to_NominationPools_ClaimCommissionCall(field),
        }
    case NominationPools_AdjustPoolDepositCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools_adjust_pool_deposit {
            NominationPools_adjust_pool_deposit: to_NominationPools_AdjustPoolDepositCall(field),
        }
    case Gear_UploadCodeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_upload_code {
            Gear_upload_code: to_Gear_UploadCodeCall(field),
        }
    case Gear_UploadProgramCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_upload_program {
            Gear_upload_program: to_Gear_UploadProgramCall(field),
        }
    case Gear_CreateProgramCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_create_program {
            Gear_create_program: to_Gear_CreateProgramCall(field),
        }
    case Gear_SendMessageCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_send_message {
            Gear_send_message: to_Gear_SendMessageCall(field),
        }
    case Gear_SendReplyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_send_reply {
            Gear_send_reply: to_Gear_SendReplyCall(field),
        }
    case Gear_ClaimValueCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_claim_value {
            Gear_claim_value: to_Gear_ClaimValueCall(field),
        }
    case Gear_RunCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_run {
            Gear_run: to_Gear_RunCall(field),
        }
    case Gear_SetExecuteInherentCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear_set_execute_inherent {
            Gear_set_execute_inherent: to_Gear_SetExecuteInherentCall(field),
        }
    case StakingRewards_RefillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_StakingRewards_refill {
            StakingRewards_refill: to_StakingRewards_RefillCall(field),
        }
    case StakingRewards_ForceRefillCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_StakingRewards_force_refill {
            StakingRewards_force_refill: to_StakingRewards_ForceRefillCall(field),
        }
    case StakingRewards_WithdrawCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_StakingRewards_withdraw {
            StakingRewards_withdraw: to_StakingRewards_WithdrawCall(field),
        }
    case StakingRewards_AlignSupplyCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_StakingRewards_align_supply {
            StakingRewards_align_supply: to_StakingRewards_AlignSupplyCall(field),
        }
    case GearVoucher_IssueCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_issue {
            GearVoucher_issue: to_GearVoucher_IssueCall(field),
        }
    case GearVoucher_CallCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_call {
            GearVoucher_call: to_GearVoucher_CallCall(field),
        }
    case GearVoucher_RevokeCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_revoke {
            GearVoucher_revoke: to_GearVoucher_RevokeCall(field),
        }
    case GearVoucher_UpdateCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_update {
            GearVoucher_update: to_GearVoucher_UpdateCall(field),
        }
    case GearVoucher_CallDeprecatedCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_call_deprecated {
            GearVoucher_call_deprecated: to_GearVoucher_CallDeprecatedCall(field),
        }
    case GearVoucher_DeclineCall:
        out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher_decline {
            GearVoucher_decline: to_GearVoucher_DeclineCall(field),
        }
    default: 
        panic("unsupported type")
    }
    return out
}  
 
func to_Whitelist_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
    out := &pbgear.Whitelist_PrimitiveTypesH256{}
        out.CallHash = to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Whitelist_RemoveWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_RemoveWhitelistedCallCall {
    out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
        out.CallHash = to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
 
func to_Whitelist_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
    out := &pbgear.Whitelist_SpWeightsWeightV2Weight{}
        out.RefTime = to_Whitelist_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
        out.ProofSize = to_Whitelist_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
    return out
}  
  
 
func to_Whitelist_WhitelistCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_WhitelistCallCall {
    out := &pbgear.Whitelist_WhitelistCallCall{}
        out.CallHash = to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
    return out
}  
  



