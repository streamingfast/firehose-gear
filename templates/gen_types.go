package gen_types

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func To_AllowedSlots(fields []*registry.DecodedField) *pbgear.AllowedSlots {
	out := &pbgear.AllowedSlots{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.AllowedSlots_PrimarySlots{
			PrimarySlots: To_PrimarySlots(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_BTreeSet(fields []*registry.DecodedField) *pbgear.BTreeSet {
	out := &pbgear.BTreeSet{}
	out.Value0 = To_repeated_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: true
// optional primitive: value0
func To_repeated_GprimitivesActorId(fields []*registry.DecodedField) []*pbgear.GprimitivesActorId {
	out := make([]*pbgear.GprimitivesActorId, len(fields))
	for _, f := range fields {
		out = append(out, To_GprimitivesActorId(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Babe_BabeTrieNodesList(fields []*registry.DecodedField) *pbgear.Babe_BabeTrieNodesList {
	out := &pbgear.Babe_BabeTrieNodesList{}
	out.TrieNodes = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "trie_nodes" optional: false repeated: false

func To_Babe_CompactUint32(fields []*registry.DecodedField) *pbgear.Babe_CompactUint32 {
	out := &pbgear.Babe_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Babe_Config(fields []*registry.DecodedField) *pbgear.Babe_Config {
	out := &pbgear.Babe_Config{}

	switch {
	case matchFields(fields, []int64{79, 80}):
		out.Value = &pbgear.Babe_Config_V1{
			V1: To_Babe_V1(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Babe_Consensus(fields []*registry.DecodedField) *pbgear.Babe_Consensus {
	out := &pbgear.Babe_Consensus{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_Logs(fields []*registry.DecodedField) *pbgear.Babe_Logs {
	out := &pbgear.Babe_Logs{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Babe_Logs_PreRuntime{
			PreRuntime: To_Babe_PreRuntime(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Babe_Other(fields []*registry.DecodedField) *pbgear.Babe_Other {
	out := &pbgear.Babe_Other{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_PlanConfigChangeCall(fields []*registry.DecodedField) *pbgear.Babe_PlanConfigChangeCall {
	out := &pbgear.Babe_PlanConfigChangeCall{}
	out.Config = To_Babe_Config(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "config" optional: false repeated: false

func To_Babe_PreRuntime(fields []*registry.DecodedField) *pbgear.Babe_PreRuntime {
	out := &pbgear.Babe_PreRuntime{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
	out := &pbgear.Babe_PrimitiveTypesH256{}
	out.ParentHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationCall {
	out := &pbgear.Babe_ReportEquivocationCall{}
	out.EquivocationProof = To_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField))
	out.KeyOwnerProof = To_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "equivocation_proof" optional: false repeated: false
// not seen "key_owner_proof" optional: false repeated: false

func To_Babe_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationUnsignedCall {
	out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
	out.EquivocationProof = To_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField))
	out.KeyOwnerProof = To_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_RuntimeEnvironmentUpdated(fields []*registry.DecodedField) *pbgear.Babe_RuntimeEnvironmentUpdated {
	out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
	return out
}

func To_Babe_Seal(fields []*registry.DecodedField) *pbgear.Babe_Seal {
	out := &pbgear.Babe_Seal{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
	out := &pbgear.Babe_SpConsensusBabeAppPublic{}
	out.Offender = To_Babe_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "offender" optional: false repeated: false

func To_Babe_SpConsensusSlotsEquivocationProof(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
	out := &pbgear.Babe_SpConsensusSlotsEquivocationProof{}
	out.Offender = To_Babe_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField))
	out.Slot = To_Babe_SpConsensusSlotsSlot(fields[1].Value.([]*registry.DecodedField))
	out.FirstHeader = To_Babe_SpRuntimeGenericHeaderHeader(fields[2].Value.([]*registry.DecodedField))
	out.SecondHeader = To_Babe_SpRuntimeGenericHeaderHeader(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "offender" optional: false repeated: false
// not seen "slot" optional: false repeated: false
// not seen "first_header" optional: false repeated: false

func To_Babe_SpConsensusSlotsSlot(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
	out := &pbgear.Babe_SpConsensusSlotsSlot{}
	out.Slot = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "slot" optional: false repeated: false

func To_Babe_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
	out := &pbgear.Babe_SpCoreSr25519Public{}
	out.Offender = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_SpRuntimeGenericDigestDigest(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
	out := &pbgear.Babe_SpRuntimeGenericDigestDigest{}
	out.Logs = To_repeated_Babe_SpRuntimeGenericDigestDigestItem(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "logs" optional: false repeated: true
// optional primitive: logs
func To_repeated_Babe_SpRuntimeGenericDigestDigestItem(fields []*registry.DecodedField) []*pbgear.Babe_SpRuntimeGenericDigestDigestItem {
	out := make([]*pbgear.Babe_SpRuntimeGenericDigestDigestItem, len(fields))
	for _, f := range fields {
		out = append(out, To_Babe_SpRuntimeGenericDigestDigestItem(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Babe_SpRuntimeGenericDigestDigestItem(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigestItem {
	out := &pbgear.Babe_SpRuntimeGenericDigestDigestItem{}
	out.Logs = To_Babe_Logs(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "logs" optional: false repeated: false

func To_Babe_SpRuntimeGenericHeaderHeader(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
	out := &pbgear.Babe_SpRuntimeGenericHeaderHeader{}
	out.ParentHash = To_Babe_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	out.Number = To_Babe_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	out.StateRoot = To_Babe_PrimitiveTypesH256(fields[2].Value.([]*registry.DecodedField))
	out.ExtrinsicsRoot = To_Babe_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField))
	out.Digest = To_Babe_SpRuntimeGenericDigestDigest(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "parent_hash" optional: false repeated: false
// not seen "number" optional: false repeated: false
// not seen "digest" optional: false repeated: false

func To_Babe_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
	out := &pbgear.Babe_SpSessionMembershipProof{}
	out.Session = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.TrieNodes = To_repeated_Babe_BabeTrieNodesList(fields[1].Value.([]*registry.DecodedField))
	out.ValidatorCount = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "trie_nodes" optional: false repeated: true
// optional primitive: trie_nodes
func To_repeated_Babe_BabeTrieNodesList(fields []*registry.DecodedField) []*pbgear.Babe_BabeTrieNodesList {
	out := make([]*pbgear.Babe_BabeTrieNodesList, len(fields))
	for _, f := range fields {
		out = append(out, To_Babe_BabeTrieNodesList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Babe_TupleUint64Uint64(fields []*registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
	out := &pbgear.Babe_TupleUint64Uint64{}
	out.Value0 = To_uint64(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_uint64(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Babe_V1(fields []*registry.DecodedField) *pbgear.Babe_V1 {
	out := &pbgear.Babe_V1{}
	out.C = To_Babe_TupleUint64Uint64(fields[0].Value.([]*registry.DecodedField))
	out.AllowedSlots = To_AllowedSlots(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "c" optional: false repeated: false
// not seen "allowed_slots" optional: false repeated: false

func To_BagsList_Address20(fields []*registry.DecodedField) *pbgear.BagsList_Address20 {
	out := &pbgear.BagsList_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_BagsList_Address32(fields []*registry.DecodedField) *pbgear.BagsList_Address32 {
	out := &pbgear.BagsList_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_BagsList_CompactTupleNull(fields []*registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
	out := &pbgear.BagsList_CompactTupleNull{}
	out.Value = To_BagsList_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_BagsList_Dislocated(fields []*registry.DecodedField) *pbgear.BagsList_Dislocated {
	out := &pbgear.BagsList_Dislocated{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.BagsList_Dislocated_Id{
			Id: To_BagsList_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.BagsList_Dislocated_Index{
			Index: To_BagsList_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.BagsList_Dislocated_Raw{
			Raw: To_BagsList_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.BagsList_Dislocated_Address32{
			Address32: To_BagsList_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.BagsList_Dislocated_Address20{
			Address20: To_BagsList_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_BagsList_Heavier(fields []*registry.DecodedField) *pbgear.BagsList_Heavier {
	out := &pbgear.BagsList_Heavier{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.BagsList_Heavier_Id{
			Id: To_BagsList_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.BagsList_Heavier_Index{
			Index: To_BagsList_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.BagsList_Heavier_Raw{
			Raw: To_BagsList_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.BagsList_Heavier_Address32{
			Address32: To_BagsList_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.BagsList_Heavier_Address20{
			Address20: To_BagsList_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_BagsList_Id(fields []*registry.DecodedField) *pbgear.BagsList_Id {
	out := &pbgear.BagsList_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_BagsList_Index(fields []*registry.DecodedField) *pbgear.BagsList_Index {
	out := &pbgear.BagsList_Index{}
	out.Value0 = To_BagsList_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_BagsList_Lighter(fields []*registry.DecodedField) *pbgear.BagsList_Lighter {
	out := &pbgear.BagsList_Lighter{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.BagsList_Lighter_Id{
			Id: To_BagsList_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.BagsList_Lighter_Index{
			Index: To_BagsList_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.BagsList_Lighter_Raw{
			Raw: To_BagsList_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.BagsList_Lighter_Address32{
			Address32: To_BagsList_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.BagsList_Lighter_Address20{
			Address20: To_BagsList_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_BagsList_PutInFrontOfCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfCall {
	out := &pbgear.BagsList_PutInFrontOfCall{}
	out.Lighter = To_BagsList_Lighter(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "lighter" optional: false repeated: false

func To_BagsList_PutInFrontOfOtherCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfOtherCall {
	out := &pbgear.BagsList_PutInFrontOfOtherCall{}
	out.Heavier = To_BagsList_Heavier(fields[0].Value.([]*registry.DecodedField))
	out.Lighter = To_BagsList_Lighter(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "heavier" optional: false repeated: false

func To_BagsList_Raw(fields []*registry.DecodedField) *pbgear.BagsList_Raw {
	out := &pbgear.BagsList_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_BagsList_RebagCall(fields []*registry.DecodedField) *pbgear.BagsList_RebagCall {
	out := &pbgear.BagsList_RebagCall{}
	out.Dislocated = To_BagsList_Dislocated(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "dislocated" optional: false repeated: false

func To_BagsList_TupleNull(fields []*registry.DecodedField) *pbgear.BagsList_TupleNull {
	out := &pbgear.BagsList_TupleNull{}
	return out
}

func To_Balances_Address20(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
	out := &pbgear.Balances_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_Address32(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
	out := &pbgear.Balances_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_CompactString(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
	out := &pbgear.Balances_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Balances_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
	out := &pbgear.Balances_CompactTupleNull{}
	out.Value = To_Balances_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Balances_Dest(fields []*registry.DecodedField) *pbgear.Balances_Dest {
	out := &pbgear.Balances_Dest{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Balances_Dest_Id{
			Id: To_Balances_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Balances_Dest_Index{
			Index: To_Balances_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Balances_Dest_Raw{
			Raw: To_Balances_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Balances_Dest_Address32{
			Address32: To_Balances_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Balances_Dest_Address20{
			Address20: To_Balances_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Balances_ForceSetBalanceCall(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
	out := &pbgear.Balances_ForceSetBalanceCall{}
	out.Who = To_Balances_Who(fields[0].Value.([]*registry.DecodedField))
	out.NewFree = To_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "who" optional: false repeated: false
// not seen "new_free" optional: false repeated: false

func To_Balances_ForceTransferCall(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
	out := &pbgear.Balances_ForceTransferCall{}
	out.Source = To_Balances_Source(fields[0].Value.([]*registry.DecodedField))
	out.Dest = To_Balances_Dest(fields[1].Value.([]*registry.DecodedField))
	out.Value = To_Balances_CompactString(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "source" optional: false repeated: false
// not seen "dest" optional: false repeated: false

func To_Balances_ForceUnreserveCall(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
	out := &pbgear.Balances_ForceUnreserveCall{}
	out.Who = To_Balances_Who(fields[0].Value.([]*registry.DecodedField))
	out.Amount = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_Id(fields []*registry.DecodedField) *pbgear.Balances_Id {
	out := &pbgear.Balances_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_Index(fields []*registry.DecodedField) *pbgear.Balances_Index {
	out := &pbgear.Balances_Index{}
	out.Value0 = To_Balances_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Balances_Raw(fields []*registry.DecodedField) *pbgear.Balances_Raw {
	out := &pbgear.Balances_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_Source(fields []*registry.DecodedField) *pbgear.Balances_Source {
	out := &pbgear.Balances_Source{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Balances_Source_Id{
			Id: To_Balances_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Balances_Source_Index{
			Index: To_Balances_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Balances_Source_Raw{
			Raw: To_Balances_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Balances_Source_Address32{
			Address32: To_Balances_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Balances_Source_Address20{
			Address20: To_Balances_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Balances_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
	out := &pbgear.Balances_SpCoreCryptoAccountId32{}
	out.Who = To_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "who" optional: false repeated: false

func To_Balances_TransferAllCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
	out := &pbgear.Balances_TransferAllCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "keep_alive" optional: false repeated: false

func To_Balances_TransferAllowDeathCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
	out := &pbgear.Balances_TransferAllowDeathCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
	out.Value = To_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
	out := &pbgear.Balances_TransferKeepAliveCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.([]*registry.DecodedField))
	out.Value = To_Balances_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_TupleNull(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
	out := &pbgear.Balances_TupleNull{}
	return out
}

func To_Balances_UpgradeAccountsCall(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
	out := &pbgear.Balances_UpgradeAccountsCall{}
	out.Who = To_repeated_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Balances_Who(fields []*registry.DecodedField) *pbgear.Balances_Who {
	out := &pbgear.Balances_Who{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Balances_Who_Id{
			Id: To_Balances_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Balances_Who_Index{
			Index: To_Balances_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Balances_Who_Raw{
			Raw: To_Balances_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Balances_Who_Address32{
			Address32: To_Balances_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Balances_Who_Address20{
			Address20: To_Balances_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	out := &pbgear.BoundedCollectionsBoundedVecBoundedVec{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_AcceptCuratorCall {
	out := &pbgear.Bounties_AcceptCuratorCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "bounty_id" optional: false repeated: false

func To_Bounties_Address20(fields []*registry.DecodedField) *pbgear.Bounties_Address20 {
	out := &pbgear.Bounties_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_Address32(fields []*registry.DecodedField) *pbgear.Bounties_Address32 {
	out := &pbgear.Bounties_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_ApproveBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ApproveBountyCall {
	out := &pbgear.Bounties_ApproveBountyCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_AwardBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_AwardBountyCall {
	out := &pbgear.Bounties_AwardBountyCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Beneficiary = To_Bounties_Beneficiary(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "beneficiary" optional: false repeated: false

func To_Bounties_Beneficiary(fields []*registry.DecodedField) *pbgear.Bounties_Beneficiary {
	out := &pbgear.Bounties_Beneficiary{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Bounties_Beneficiary_Id{
			Id: To_Bounties_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Bounties_Beneficiary_Index{
			Index: To_Bounties_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Bounties_Beneficiary_Raw{
			Raw: To_Bounties_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Bounties_Beneficiary_Address32{
			Address32: To_Bounties_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Bounties_Beneficiary_Address20{
			Address20: To_Bounties_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Bounties_ClaimBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ClaimBountyCall {
	out := &pbgear.Bounties_ClaimBountyCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_CloseBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_CloseBountyCall {
	out := &pbgear.Bounties_CloseBountyCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_CompactString(fields []*registry.DecodedField) *pbgear.Bounties_CompactString {
	out := &pbgear.Bounties_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
	out := &pbgear.Bounties_CompactTupleNull{}
	out.Value = To_Bounties_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Bounties_CompactUint32(fields []*registry.DecodedField) *pbgear.Bounties_CompactUint32 {
	out := &pbgear.Bounties_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_Curator(fields []*registry.DecodedField) *pbgear.Bounties_Curator {
	out := &pbgear.Bounties_Curator{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Bounties_Curator_Id{
			Id: To_Bounties_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Bounties_Curator_Index{
			Index: To_Bounties_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Bounties_Curator_Raw{
			Raw: To_Bounties_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Bounties_Curator_Address32{
			Address32: To_Bounties_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Bounties_Curator_Address20{
			Address20: To_Bounties_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Bounties_ExtendBountyExpiryCall(fields []*registry.DecodedField) *pbgear.Bounties_ExtendBountyExpiryCall {
	out := &pbgear.Bounties_ExtendBountyExpiryCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Remark = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_Id(fields []*registry.DecodedField) *pbgear.Bounties_Id {
	out := &pbgear.Bounties_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_Index(fields []*registry.DecodedField) *pbgear.Bounties_Index {
	out := &pbgear.Bounties_Index{}
	out.Value0 = To_Bounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Bounties_ProposeBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeBountyCall {
	out := &pbgear.Bounties_ProposeBountyCall{}
	out.Value = To_Bounties_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Description = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Bounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeCuratorCall {
	out := &pbgear.Bounties_ProposeCuratorCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Curator = To_Bounties_Curator(fields[1].Value.([]*registry.DecodedField))
	out.Fee = To_Bounties_CompactString(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "curator" optional: false repeated: false

func To_Bounties_Raw(fields []*registry.DecodedField) *pbgear.Bounties_Raw {
	out := &pbgear.Bounties_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Bounties_TupleNull(fields []*registry.DecodedField) *pbgear.Bounties_TupleNull {
	out := &pbgear.Bounties_TupleNull{}
	return out
}

func To_Bounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_UnassignCuratorCall {
	out := &pbgear.Bounties_UnassignCuratorCall{}
	out.BountyId = To_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AcceptCuratorCall {
	out := &pbgear.ChildBounties_AcceptCuratorCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "parent_bounty_id" optional: false repeated: false

func To_ChildBounties_AddChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AddChildBountyCall {
	out := &pbgear.ChildBounties_AddChildBountyCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value = To_ChildBounties_CompactString(fields[1].Value.([]*registry.DecodedField))
	out.Description = To_bytes(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_ChildBounties_Address20(fields []*registry.DecodedField) *pbgear.ChildBounties_Address20 {
	out := &pbgear.ChildBounties_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_Address32(fields []*registry.DecodedField) *pbgear.ChildBounties_Address32 {
	out := &pbgear.ChildBounties_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_AwardChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AwardChildBountyCall {
	out := &pbgear.ChildBounties_AwardChildBountyCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	out.Beneficiary = To_ChildBounties_Beneficiary(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "beneficiary" optional: false repeated: false

func To_ChildBounties_Beneficiary(fields []*registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
	out := &pbgear.ChildBounties_Beneficiary{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.ChildBounties_Beneficiary_Id{
			Id: To_ChildBounties_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.ChildBounties_Beneficiary_Index{
			Index: To_ChildBounties_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.ChildBounties_Beneficiary_Raw{
			Raw: To_ChildBounties_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.ChildBounties_Beneficiary_Address32{
			Address32: To_ChildBounties_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.ChildBounties_Beneficiary_Address20{
			Address20: To_ChildBounties_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_ChildBounties_ClaimChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ClaimChildBountyCall {
	out := &pbgear.ChildBounties_ClaimChildBountyCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_CloseChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_CloseChildBountyCall {
	out := &pbgear.ChildBounties_CloseChildBountyCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_CompactString(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactString {
	out := &pbgear.ChildBounties_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
	out := &pbgear.ChildBounties_CompactTupleNull{}
	out.Value = To_ChildBounties_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_ChildBounties_CompactUint32(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
	out := &pbgear.ChildBounties_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_Curator(fields []*registry.DecodedField) *pbgear.ChildBounties_Curator {
	out := &pbgear.ChildBounties_Curator{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.ChildBounties_Curator_Id{
			Id: To_ChildBounties_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.ChildBounties_Curator_Index{
			Index: To_ChildBounties_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.ChildBounties_Curator_Raw{
			Raw: To_ChildBounties_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.ChildBounties_Curator_Address32{
			Address32: To_ChildBounties_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.ChildBounties_Curator_Address20{
			Address20: To_ChildBounties_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ChildBounties_Id(fields []*registry.DecodedField) *pbgear.ChildBounties_Id {
	out := &pbgear.ChildBounties_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_Index(fields []*registry.DecodedField) *pbgear.ChildBounties_Index {
	out := &pbgear.ChildBounties_Index{}
	out.Value0 = To_ChildBounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_ChildBounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ProposeCuratorCall {
	out := &pbgear.ChildBounties_ProposeCuratorCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	out.Curator = To_ChildBounties_Curator(fields[2].Value.([]*registry.DecodedField))
	out.Fee = To_ChildBounties_CompactString(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "curator" optional: false repeated: false

func To_ChildBounties_Raw(fields []*registry.DecodedField) *pbgear.ChildBounties_Raw {
	out := &pbgear.ChildBounties_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ChildBounties_TupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_TupleNull {
	out := &pbgear.ChildBounties_TupleNull{}
	return out
}

func To_ChildBounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_UnassignCuratorCall {
	out := &pbgear.ChildBounties_UnassignCuratorCall{}
	out.ParentBountyId = To_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.ChildBountyId = To_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Address20(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address20 {
	out := &pbgear.ConvictionVoting_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Address32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address32 {
	out := &pbgear.ConvictionVoting_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
	out := &pbgear.ConvictionVoting_CompactTupleNull{}
	out.Value = To_ConvictionVoting_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_ConvictionVoting_CompactUint32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
	out := &pbgear.ConvictionVoting_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Conviction(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
	out := &pbgear.ConvictionVoting_Conviction{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.ConvictionVoting_Conviction_None{
			None: To_ConvictionVoting_None(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_ConvictionVoting_DelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_DelegateCall {
	out := &pbgear.ConvictionVoting_DelegateCall{}
	out.Class = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.To = To_ConvictionVoting_To(fields[1].Value.([]*registry.DecodedField))
	out.Conviction = To_ConvictionVoting_Conviction(fields[2].Value.([]*registry.DecodedField))
	out.Balance = To_string(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "to" optional: false repeated: false
// not seen "conviction" optional: false repeated: false

func To_ConvictionVoting_Id(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Id {
	out := &pbgear.ConvictionVoting_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Index(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Index {
	out := &pbgear.ConvictionVoting_Index{}
	out.Value0 = To_ConvictionVoting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_ConvictionVoting_Locked1X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked1X {
	out := &pbgear.ConvictionVoting_Locked1X{}
	return out
}

func To_ConvictionVoting_Locked2X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked2X {
	out := &pbgear.ConvictionVoting_Locked2X{}
	return out
}

func To_ConvictionVoting_Locked3X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked3X {
	out := &pbgear.ConvictionVoting_Locked3X{}
	return out
}

func To_ConvictionVoting_Locked4X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked4X {
	out := &pbgear.ConvictionVoting_Locked4X{}
	return out
}

func To_ConvictionVoting_Locked5X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked5X {
	out := &pbgear.ConvictionVoting_Locked5X{}
	return out
}

func To_ConvictionVoting_Locked6X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked6X {
	out := &pbgear.ConvictionVoting_Locked6X{}
	return out
}

func To_ConvictionVoting_None(fields []*registry.DecodedField) *pbgear.ConvictionVoting_None {
	out := &pbgear.ConvictionVoting_None{}
	return out
}

func To_ConvictionVoting_Raw(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Raw {
	out := &pbgear.ConvictionVoting_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_RemoveOtherVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
	out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
	out.Target = To_ConvictionVoting_Target(fields[0].Value.([]*registry.DecodedField))
	out.Class = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "target" optional: false repeated: false

func To_ConvictionVoting_RemoveVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveVoteCall {
	out := &pbgear.ConvictionVoting_RemoveVoteCall{}
	out.Class = To_optional_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Split(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Split {
	out := &pbgear.ConvictionVoting_Split{}
	out.Aye = To_string(fields[0].Value.([]*registry.DecodedField))
	out.Nay = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_SplitAbstain(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SplitAbstain {
	out := &pbgear.ConvictionVoting_SplitAbstain{}
	out.Aye = To_string(fields[0].Value.([]*registry.DecodedField))
	out.Nay = To_string(fields[1].Value.([]*registry.DecodedField))
	out.Abstain = To_string(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Standard(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Standard {
	out := &pbgear.ConvictionVoting_Standard{}
	out.Vote = To_conviction_voting_PalletConvictionVotingVoteVote(fields[0].Value.([]*registry.DecodedField))
	out.Balance = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "vote" optional: false repeated: false

func To_ConvictionVoting_Target(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Target {
	out := &pbgear.ConvictionVoting_Target{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.ConvictionVoting_Target_Id{
			Id: To_ConvictionVoting_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.ConvictionVoting_Target_Index{
			Index: To_ConvictionVoting_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.ConvictionVoting_Target_Raw{
			Raw: To_ConvictionVoting_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.ConvictionVoting_Target_Address32{
			Address32: To_ConvictionVoting_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.ConvictionVoting_Target_Address20{
			Address20: To_ConvictionVoting_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ConvictionVoting_To(fields []*registry.DecodedField) *pbgear.ConvictionVoting_To {
	out := &pbgear.ConvictionVoting_To{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.ConvictionVoting_To_Id{
			Id: To_ConvictionVoting_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.ConvictionVoting_To_Index{
			Index: To_ConvictionVoting_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.ConvictionVoting_To_Raw{
			Raw: To_ConvictionVoting_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.ConvictionVoting_To_Address32{
			Address32: To_ConvictionVoting_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.ConvictionVoting_To_Address20{
			Address20: To_ConvictionVoting_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ConvictionVoting_TupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
	out := &pbgear.ConvictionVoting_TupleNull{}
	return out
}

func To_ConvictionVoting_UndelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UndelegateCall {
	out := &pbgear.ConvictionVoting_UndelegateCall{}
	out.Class = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_UnlockCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UnlockCall {
	out := &pbgear.ConvictionVoting_UnlockCall{}
	out.Class = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Target = To_ConvictionVoting_Target(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ConvictionVoting_Vote(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Vote {
	out := &pbgear.ConvictionVoting_Vote{}

	switch {
	case matchFields(fields, []int64{126, 6}):
		out.Value = &pbgear.ConvictionVoting_Vote_Standard{
			Standard: To_ConvictionVoting_Standard(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{6, 6}):
		out.Value = &pbgear.ConvictionVoting_Vote_Split{
			Split: To_ConvictionVoting_Split(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{6, 6, 6}):
		out.Value = &pbgear.ConvictionVoting_Vote_SplitAbstain{
			SplitAbstain: To_ConvictionVoting_SplitAbstain(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ConvictionVoting_VoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_VoteCall {
	out := &pbgear.ConvictionVoting_VoteCall{}
	out.PollIndex = To_ConvictionVoting_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Vote = To_ConvictionVoting_Vote(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "poll_index" optional: false repeated: false
// not seen "vote" optional: false repeated: false

func To_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
	out := &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{}
	out.Value = To_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_ElectionProviderMultiPhase_CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_GovernanceFallbackCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
	out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
	out.MaybeMaxVoters = To_optional_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaybeMaxTargets = To_optional_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
	out.Solution = To_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields[0].Value.([]*registry.DecodedField))
	out.Score = To_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[1].Value.([]*registry.DecodedField))
	out.Round = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "solution" optional: false repeated: false
// not seen "score" optional: false repeated: false

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
	out.Voters = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Targets = To_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "voters" optional: false repeated: false

func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
	out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
	out.Supports = To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "supports" optional: false repeated: true
// optional primitive: supports
func To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
	out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
	out.MaybeNextScore = To_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
	out := &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
	out := &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
	out := &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{}
	out.MinimalStake = To_string(fields[0].Value.([]*registry.DecodedField))
	out.SumStake = To_string(fields[1].Value.([]*registry.DecodedField))
	out.SumStakeSquared = To_string(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
	out := &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{}
	out.Total = To_string(fields[0].Value.([]*registry.DecodedField))
	out.Voters = To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "voters" optional: false repeated: true
// optional primitive: voters
func To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_ElectionProviderMultiPhase_SubmitCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "raw_solution" optional: false repeated: false

func To_ElectionProviderMultiPhase_SubmitUnsignedCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField))
	out.Witness = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "witness" optional: false repeated: false

func To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value1" optional: false repeated: false

func To_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "value1" optional: false repeated: true
// optional primitive: value1
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{}
	out.Value0 = To_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField))
	out.Value2 = To_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
	out.Value0 = To_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_ElectionProviderMultiPhase_SpNposElectionsSupport(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false
// not seen "value1" optional: false repeated: false

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
	out.Value0 = To_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
	out := &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{}
	out.Votes1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Votes2 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields[1].Value.([]*registry.DecodedField))
	out.Votes3 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields[2].Value.([]*registry.DecodedField))
	out.Votes4 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields[3].Value.([]*registry.DecodedField))
	out.Votes5 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields[4].Value.([]*registry.DecodedField))
	out.Votes6 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields[5].Value.([]*registry.DecodedField))
	out.Votes7 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields[6].Value.([]*registry.DecodedField))
	out.Votes8 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields[7].Value.([]*registry.DecodedField))
	out.Votes9 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields[8].Value.([]*registry.DecodedField))
	out.Votes10 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields[9].Value.([]*registry.DecodedField))
	out.Votes11 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields[10].Value.([]*registry.DecodedField))
	out.Votes12 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields[11].Value.([]*registry.DecodedField))
	out.Votes13 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields[12].Value.([]*registry.DecodedField))
	out.Votes14 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields[13].Value.([]*registry.DecodedField))
	out.Votes15 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields[14].Value.([]*registry.DecodedField))
	out.Votes16 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields[15].Value.([]*registry.DecodedField))
	return out
}

// not seen "votes1" optional: false repeated: true
// optional primitive: votes1
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes2" optional: false repeated: true
// optional primitive: votes2
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes3" optional: false repeated: true
// optional primitive: votes3
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes4" optional: false repeated: true
// optional primitive: votes4
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes5" optional: false repeated: true
// optional primitive: votes5
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes6" optional: false repeated: true
// optional primitive: votes6
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes7" optional: false repeated: true
// optional primitive: votes7
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes8" optional: false repeated: true
// optional primitive: votes8
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes9" optional: false repeated: true
// optional primitive: votes9
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes10" optional: false repeated: true
// optional primitive: votes10
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes11" optional: false repeated: true
// optional primitive: votes11
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes12" optional: false repeated: true
// optional primitive: votes12
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes13" optional: false repeated: true
// optional primitive: votes13
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes14" optional: false repeated: true
// optional primitive: votes14
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes15" optional: false repeated: true
// optional primitive: votes15
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "votes16" optional: false repeated: true
// optional primitive: votes16
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields []*registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_FellowshipCollective_AddMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_AddMemberCall {
	out := &pbgear.FellowshipCollective_AddMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "who" optional: false repeated: false

func To_FellowshipCollective_Address20(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address20 {
	out := &pbgear.FellowshipCollective_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_Address32(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address32 {
	out := &pbgear.FellowshipCollective_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_CleanupPollCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CleanupPollCall {
	out := &pbgear.FellowshipCollective_CleanupPollCall{}
	out.PollIndex = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Max = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_CompactTupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
	out := &pbgear.FellowshipCollective_CompactTupleNull{}
	out.Value = To_FellowshipCollective_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_FellowshipCollective_DemoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_DemoteMemberCall {
	out := &pbgear.FellowshipCollective_DemoteMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_Id(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Id {
	out := &pbgear.FellowshipCollective_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_Index(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Index {
	out := &pbgear.FellowshipCollective_Index{}
	out.Value0 = To_FellowshipCollective_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_FellowshipCollective_PromoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_PromoteMemberCall {
	out := &pbgear.FellowshipCollective_PromoteMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_Raw(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Raw {
	out := &pbgear.FellowshipCollective_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_RemoveMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_RemoveMemberCall {
	out := &pbgear.FellowshipCollective_RemoveMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField))
	out.MinRank = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_TupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
	out := &pbgear.FellowshipCollective_TupleNull{}
	return out
}

func To_FellowshipCollective_VoteCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_VoteCall {
	out := &pbgear.FellowshipCollective_VoteCall{}
	out.Poll = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Aye = To_bool(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipCollective_Who(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Who {
	out := &pbgear.FellowshipCollective_Who{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.FellowshipCollective_Who_Id{
			Id: To_FellowshipCollective_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.FellowshipCollective_Who_Index{
			Index: To_FellowshipCollective_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.FellowshipCollective_Who_Raw{
			Raw: To_FellowshipCollective_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.FellowshipCollective_Who_Address32{
			Address32: To_FellowshipCollective_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.FellowshipCollective_Who_Address20{
			Address20: To_FellowshipCollective_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_FellowshipReferenda_After(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_After {
	out := &pbgear.FellowshipReferenda_After{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_At(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_At {
	out := &pbgear.FellowshipReferenda_At{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_CancelCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_CancelCall {
	out := &pbgear.FellowshipReferenda_CancelCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
	out := &pbgear.FellowshipReferenda_EnactmentMoment{}

	switch {
	case matchFields(fields, []int64{4}):
		out.Value = &pbgear.FellowshipReferenda_EnactmentMoment_At{
			At: To_FellowshipReferenda_At(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{4}):
		out.Value = &pbgear.FellowshipReferenda_EnactmentMoment_After{
			After: To_FellowshipReferenda_After(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_FellowshipReferenda_Inline(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Inline {
	out := &pbgear.FellowshipReferenda_Inline{}
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_FellowshipReferenda_KillCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_KillCall {
	out := &pbgear.FellowshipReferenda_KillCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_Legacy(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Legacy {
	out := &pbgear.FellowshipReferenda_Legacy{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "hash" optional: false repeated: false

func To_FellowshipReferenda_Lookup(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Lookup {
	out := &pbgear.FellowshipReferenda_Lookup{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	out.Len = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_NudgeReferendumCall {
	out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
	out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
	out.Track = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_Origins(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Origins {
	out := &pbgear.FellowshipReferenda_Origins{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_FellowshipReferenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
	out := &pbgear.FellowshipReferenda_PrimitiveTypesH256{}
	out.MaybeHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_Proposal(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
	out := &pbgear.FellowshipReferenda_Proposal{}

	switch {
	case matchFields(fields, []int64{12}):
		out.Value = &pbgear.FellowshipReferenda_Proposal_Legacy{
			Legacy: To_FellowshipReferenda_Legacy(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{279}):
		out.Value = &pbgear.FellowshipReferenda_Proposal_Inline{
			Inline: To_FellowshipReferenda_Inline(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{12, 4}):
		out.Value = &pbgear.FellowshipReferenda_Proposal_Lookup{
			Lookup: To_FellowshipReferenda_Lookup(fields), //Passthrough fields...
		}
	}
	return out
}

func To_FellowshipReferenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
	out := &pbgear.FellowshipReferenda_ProposalOrigin{}

	switch {
	case matchFields(fields, []int64{121}):
		out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_System{
			System: To_FellowshipReferenda_System(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{122}):
		out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_Origins{
			Origins: To_FellowshipReferenda_Origins(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{123}):
		out.Value = &pbgear.FellowshipReferenda_ProposalOrigin_Void{
			Void: To_FellowshipReferenda_Void(fields), //Passthrough fields...
		}
	}
	return out
}

func To_FellowshipReferenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SetMetadataCall {
	out := &pbgear.FellowshipReferenda_SetMetadataCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaybeHash = To_optional_FellowshipReferenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "maybe_hash" optional: true repeated: false
// optional field: maybe_hash

func To_FellowshipReferenda_SubmitCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SubmitCall {
	out := &pbgear.FellowshipReferenda_SubmitCall{}
	out.ProposalOrigin = To_FellowshipReferenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField))
	out.Proposal = To_FellowshipReferenda_Proposal(fields[1].Value.([]*registry.DecodedField))
	out.EnactmentMoment = To_FellowshipReferenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "proposal_origin" optional: false repeated: false
// not seen "proposal" optional: false repeated: false
// not seen "enactment_moment" optional: false repeated: false

func To_FellowshipReferenda_System(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_System {
	out := &pbgear.FellowshipReferenda_System{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FellowshipReferenda_Void(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Void {
	out := &pbgear.FellowshipReferenda_Void{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_FinalityGrandpaEquivocation(fields []*registry.DecodedField) *pbgear.FinalityGrandpaEquivocation {
	out := &pbgear.FinalityGrandpaEquivocation{}
	out.RoundNumber = To_uint64(fields[0].Value.([]*registry.DecodedField))
	out.Identity = To_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField))
	out.First = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[2].Value.([]*registry.DecodedField))
	out.Second = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "identity" optional: false repeated: false
// not seen "first" optional: false repeated: false

func To_FinalityGrandpaPrevote(fields []*registry.DecodedField) *pbgear.FinalityGrandpaPrevote {
	out := &pbgear.FinalityGrandpaPrevote{}
	out.TargetHash = To_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	out.TargetNumber = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_AppendPrograms(fields []*registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
	out := &pbgear.GearVoucher_AppendPrograms{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.GearVoucher_AppendPrograms_None{
			None: To_GearVoucher_None(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_GearVoucher_BTreeSet(fields []*registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
	out := &pbgear.GearVoucher_BTreeSet{}
	out.Programs = To_repeated_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "programs" optional: false repeated: true
// optional primitive: programs
func To_repeated_GearVoucher_GprimitivesActorId(fields []*registry.DecodedField) []*pbgear.GearVoucher_GprimitivesActorId {
	out := make([]*pbgear.GearVoucher_GprimitivesActorId, len(fields))
	for _, f := range fields {
		out = append(out, To_GearVoucher_GprimitivesActorId(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_GearVoucher_Call(fields []*registry.DecodedField) *pbgear.GearVoucher_Call {
	out := &pbgear.GearVoucher_Call{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.GearVoucher_Call_SendMessage{
			SendMessage: To_GearVoucher_SendMessage(fields), //Passthrough fields...
		}
	}
	return out
}

func To_GearVoucher_CallCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallCall {
	out := &pbgear.GearVoucher_CallCall{}
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField))
	out.Call = To_GearVoucher_Call(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "voucher_id" optional: false repeated: false
// not seen "call" optional: false repeated: false

func To_GearVoucher_CallDeprecatedCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallDeprecatedCall {
	out := &pbgear.GearVoucher_CallDeprecatedCall{}
	out.Call = To_GearVoucher_Call(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_DeclineCall(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineCall {
	out := &pbgear.GearVoucher_DeclineCall{}
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_DeclineVoucher(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineVoucher {
	out := &pbgear.GearVoucher_DeclineVoucher{}
	return out
}

func To_GearVoucher_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
	out := &pbgear.GearVoucher_GprimitivesActorId{}
	out.Programs = To_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_IssueCall(fields []*registry.DecodedField) *pbgear.GearVoucher_IssueCall {
	out := &pbgear.GearVoucher_IssueCall{}
	out.Spender = To_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.Balance = To_string(fields[1].Value.([]*registry.DecodedField))
	out.Programs = To_optional_GearVoucher_BTreeSet(fields[2].Value.([]*registry.DecodedField))
	out.CodeUploading = To_bool(fields[3].Value.([]*registry.DecodedField))
	out.Duration = To_uint32(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "spender" optional: false repeated: false
// not seen "programs" optional: true repeated: false
// optional field: programs

func To_GearVoucher_None(fields []*registry.DecodedField) *pbgear.GearVoucher_None {
	out := &pbgear.GearVoucher_None{}
	return out
}

func To_GearVoucher_PalletGearVoucherInternalVoucherId(fields []*registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
	out.VoucherId = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_RevokeCall(fields []*registry.DecodedField) *pbgear.GearVoucher_RevokeCall {
	out := &pbgear.GearVoucher_RevokeCall{}
	out.Spender = To_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_SendMessage(fields []*registry.DecodedField) *pbgear.GearVoucher_SendMessage {
	out := &pbgear.GearVoucher_SendMessage{}
	out.Destination = To_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField))
	out.Payload = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[2].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[3].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[4].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_SendReply(fields []*registry.DecodedField) *pbgear.GearVoucher_SendReply {
	out := &pbgear.GearVoucher_SendReply{}
	out.ReplyToId = To_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField))
	out.Payload = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[2].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[3].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "reply_to_id" optional: false repeated: false

func To_GearVoucher_Some(fields []*registry.DecodedField) *pbgear.GearVoucher_Some {
	out := &pbgear.GearVoucher_Some{}
	out.Value0 = To_BTreeSet(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_GearVoucher_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
	out := &pbgear.GearVoucher_SpCoreCryptoAccountId32{}
	out.Spender = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GearVoucher_UpdateCall(fields []*registry.DecodedField) *pbgear.GearVoucher_UpdateCall {
	out := &pbgear.GearVoucher_UpdateCall{}
	out.Spender = To_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField))
	out.MoveOwnership = To_optional_GearVoucher_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField))
	out.BalanceTopUp = To_optional_string(fields[3].Value.([]*registry.DecodedField))
	out.AppendPrograms = To_optional_GearVoucher_AppendPrograms(fields[4].Value.([]*registry.DecodedField))
	out.CodeUploading = To_optional_bool(fields[5].Value.([]*registry.DecodedField))
	out.ProlongDuration = To_optional_uint32(fields[6].Value.([]*registry.DecodedField))
	return out
}

// not seen "append_programs" optional: true repeated: false
// optional field: append_programs

func To_GearVoucher_UploadCode(fields []*registry.DecodedField) *pbgear.GearVoucher_UploadCode {
	out := &pbgear.GearVoucher_UploadCode{}
	out.Code = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_ClaimValueCall(fields []*registry.DecodedField) *pbgear.Gear_ClaimValueCall {
	out := &pbgear.Gear_ClaimValueCall{}
	out.MessageId = To_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "message_id" optional: false repeated: false

func To_Gear_CreateProgramCall(fields []*registry.DecodedField) *pbgear.Gear_CreateProgramCall {
	out := &pbgear.Gear_CreateProgramCall{}
	out.CodeId = To_Gear_GprimitivesCodeId(fields[0].Value.([]*registry.DecodedField))
	out.Salt = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.InitPayload = To_bytes(fields[2].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[3].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[4].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[5].Value.([]*registry.DecodedField))
	return out
}

// not seen "code_id" optional: false repeated: false

func To_Gear_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
	out := &pbgear.Gear_GprimitivesActorId{}
	out.Destination = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_GprimitivesCodeId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
	out := &pbgear.Gear_GprimitivesCodeId{}
	out.CodeId = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
	out := &pbgear.Gear_GprimitivesMessageId{}
	out.ReplyToId = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_RunCall(fields []*registry.DecodedField) *pbgear.Gear_RunCall {
	out := &pbgear.Gear_RunCall{}
	out.MaxGas = To_optional_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_SendMessageCall(fields []*registry.DecodedField) *pbgear.Gear_SendMessageCall {
	out := &pbgear.Gear_SendMessageCall{}
	out.Destination = To_Gear_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField))
	out.Payload = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[2].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[3].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "destination" optional: false repeated: false

func To_Gear_SendReplyCall(fields []*registry.DecodedField) *pbgear.Gear_SendReplyCall {
	out := &pbgear.Gear_SendReplyCall{}
	out.ReplyToId = To_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField))
	out.Payload = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[2].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[3].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[4].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_SetExecuteInherentCall(fields []*registry.DecodedField) *pbgear.Gear_SetExecuteInherentCall {
	out := &pbgear.Gear_SetExecuteInherentCall{}
	out.Value = To_bool(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_UploadCodeCall(fields []*registry.DecodedField) *pbgear.Gear_UploadCodeCall {
	out := &pbgear.Gear_UploadCodeCall{}
	out.Code = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Gear_UploadProgramCall(fields []*registry.DecodedField) *pbgear.Gear_UploadProgramCall {
	out := &pbgear.Gear_UploadProgramCall{}
	out.Code = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Salt = To_bytes(fields[1].Value.([]*registry.DecodedField))
	out.InitPayload = To_bytes(fields[2].Value.([]*registry.DecodedField))
	out.GasLimit = To_uint64(fields[3].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[4].Value.([]*registry.DecodedField))
	out.KeepAlive = To_bool(fields[5].Value.([]*registry.DecodedField))
	return out
}

func To_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.GprimitivesActorId {
	out := &pbgear.GprimitivesActorId{}
	out.Destination = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.GprimitivesMessageId {
	out := &pbgear.GprimitivesMessageId{}
	out.ReplyToId = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Grandpa_Equivocation(fields []*registry.DecodedField) *pbgear.Grandpa_Equivocation {
	out := &pbgear.Grandpa_Equivocation{}

	switch {
	case matchFields(fields, []int64{84}):
		out.Value = &pbgear.Grandpa_Equivocation_Prevote{
			Prevote: To_Grandpa_Prevote(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{90}):
		out.Value = &pbgear.Grandpa_Equivocation_Precommit{
			Precommit: To_Grandpa_Precommit(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Grandpa_GrandpaTrieNodesList(fields []*registry.DecodedField) *pbgear.Grandpa_GrandpaTrieNodesList {
	out := &pbgear.Grandpa_GrandpaTrieNodesList{}
	out.TrieNodes = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Grandpa_NoteStalledCall(fields []*registry.DecodedField) *pbgear.Grandpa_NoteStalledCall {
	out := &pbgear.Grandpa_NoteStalledCall{}
	out.Delay = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.BestFinalizedBlockNumber = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Grandpa_Precommit(fields []*registry.DecodedField) *pbgear.Grandpa_Precommit {
	out := &pbgear.Grandpa_Precommit{}
	out.Value0 = To_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Grandpa_Prevote(fields []*registry.DecodedField) *pbgear.Grandpa_Prevote {
	out := &pbgear.Grandpa_Prevote{}
	out.Value0 = To_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Grandpa_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationCall {
	out := &pbgear.Grandpa_ReportEquivocationCall{}
	out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField))
	out.KeyOwnerProof = To_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "equivocation_proof" optional: false repeated: false
// not seen "key_owner_proof" optional: false repeated: false

func To_Grandpa_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
	out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
	out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField))
	out.KeyOwnerProof = To_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Grandpa_SpConsensusGrandpaEquivocationProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
	out := &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{}
	out.SetId = To_uint64(fields[0].Value.([]*registry.DecodedField))
	out.Equivocation = To_Grandpa_Equivocation(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "equivocation" optional: false repeated: false

func To_Grandpa_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
	out := &pbgear.Grandpa_SpSessionMembershipProof{}
	out.Session = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.TrieNodes = To_repeated_Grandpa_GrandpaTrieNodesList(fields[1].Value.([]*registry.DecodedField))
	out.ValidatorCount = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "trie_nodes" optional: false repeated: true
// optional primitive: trie_nodes
func To_repeated_Grandpa_GrandpaTrieNodesList(fields []*registry.DecodedField) []*pbgear.Grandpa_GrandpaTrieNodesList {
	out := make([]*pbgear.Grandpa_GrandpaTrieNodesList, len(fields))
	for _, f := range fields {
		out = append(out, To_Grandpa_GrandpaTrieNodesList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Identity_Account(fields []*registry.DecodedField) *pbgear.Identity_Account {
	out := &pbgear.Identity_Account{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Identity_Account_Id{
			Id: To_Identity_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Identity_Account_Index{
			Index: To_Identity_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Identity_Account_Raw{
			Raw: To_Identity_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Identity_Account_Address32{
			Address32: To_Identity_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Identity_Account_Address20{
			Address20: To_Identity_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Identity_AddRegistrarCall(fields []*registry.DecodedField) *pbgear.Identity_AddRegistrarCall {
	out := &pbgear.Identity_AddRegistrarCall{}
	out.Account = To_Identity_Account(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "account" optional: false repeated: false

func To_Identity_AddSubCall(fields []*registry.DecodedField) *pbgear.Identity_AddSubCall {
	out := &pbgear.Identity_AddSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.([]*registry.DecodedField))
	out.Data = To_Identity_Data(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "sub" optional: false repeated: false
// not seen "data" optional: false repeated: false

func To_Identity_Address20(fields []*registry.DecodedField) *pbgear.Identity_Address20 {
	out := &pbgear.Identity_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Address32(fields []*registry.DecodedField) *pbgear.Identity_Address32 {
	out := &pbgear.Identity_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_BlakeTwo256(fields []*registry.DecodedField) *pbgear.Identity_BlakeTwo256 {
	out := &pbgear.Identity_BlakeTwo256{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
	out := &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{}
	out.Additional = To_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "additional" optional: false repeated: true
// optional primitive: additional
func To_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields []*registry.DecodedField) []*pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	out := make([]*pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData, len(fields))
	for _, f := range fields {
		out = append(out, To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Identity_CancelRequestCall(fields []*registry.DecodedField) *pbgear.Identity_CancelRequestCall {
	out := &pbgear.Identity_CancelRequestCall{}
	out.RegIndex = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_ClearIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_ClearIdentityCall {
	out := &pbgear.Identity_ClearIdentityCall{}
	return out
}

func To_Identity_CompactString(fields []*registry.DecodedField) *pbgear.Identity_CompactString {
	out := &pbgear.Identity_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Identity_CompactTupleNull {
	out := &pbgear.Identity_CompactTupleNull{}
	out.Value = To_Identity_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Identity_CompactUint32(fields []*registry.DecodedField) *pbgear.Identity_CompactUint32 {
	out := &pbgear.Identity_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Data(fields []*registry.DecodedField) *pbgear.Identity_Data {
	out := &pbgear.Identity_Data{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Data_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Display(fields []*registry.DecodedField) *pbgear.Identity_Display {
	out := &pbgear.Identity_Display{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Display_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Email(fields []*registry.DecodedField) *pbgear.Identity_Email {
	out := &pbgear.Identity_Email{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Email_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Erroneous(fields []*registry.DecodedField) *pbgear.Identity_Erroneous {
	out := &pbgear.Identity_Erroneous{}
	return out
}

func To_Identity_FeePaid(fields []*registry.DecodedField) *pbgear.Identity_FeePaid {
	out := &pbgear.Identity_FeePaid{}
	out.Value0 = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Id(fields []*registry.DecodedField) *pbgear.Identity_Id {
	out := &pbgear.Identity_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Image(fields []*registry.DecodedField) *pbgear.Identity_Image {
	out := &pbgear.Identity_Image{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Image_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Index(fields []*registry.DecodedField) *pbgear.Identity_Index {
	out := &pbgear.Identity_Index{}
	out.Value0 = To_Identity_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Identity_Judgement(fields []*registry.DecodedField) *pbgear.Identity_Judgement {
	out := &pbgear.Identity_Judgement{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Judgement_Unknown{
			Unknown: To_Identity_Unknown(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Keccak256(fields []*registry.DecodedField) *pbgear.Identity_Keccak256 {
	out := &pbgear.Identity_Keccak256{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_KillIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_KillIdentityCall {
	out := &pbgear.Identity_KillIdentityCall{}
	out.Target = To_Identity_Target(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "target" optional: false repeated: false

func To_Identity_KnownGood(fields []*registry.DecodedField) *pbgear.Identity_KnownGood {
	out := &pbgear.Identity_KnownGood{}
	return out
}

func To_Identity_Legal(fields []*registry.DecodedField) *pbgear.Identity_Legal {
	out := &pbgear.Identity_Legal{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Legal_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_LowQuality(fields []*registry.DecodedField) *pbgear.Identity_LowQuality {
	out := &pbgear.Identity_LowQuality{}
	return out
}

func To_Identity_New(fields []*registry.DecodedField) *pbgear.Identity_New {
	out := &pbgear.Identity_New{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Identity_New_Id{
			Id: To_Identity_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Identity_New_Index{
			Index: To_Identity_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Identity_New_Raw{
			Raw: To_Identity_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Identity_New_Address32{
			Address32: To_Identity_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Identity_New_Address20{
			Address20: To_Identity_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_None(fields []*registry.DecodedField) *pbgear.Identity_None {
	out := &pbgear.Identity_None{}
	return out
}

func To_Identity_OutOfDate(fields []*registry.DecodedField) *pbgear.Identity_OutOfDate {
	out := &pbgear.Identity_OutOfDate{}
	return out
}

func To_Identity_PalletIdentitySimpleIdentityInfo(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
	out.Additional = To_Identity_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField))
	out.Display = To_Identity_Display(fields[1].Value.([]*registry.DecodedField))
	out.Legal = To_Identity_Legal(fields[2].Value.([]*registry.DecodedField))
	out.Web = To_Identity_Web(fields[3].Value.([]*registry.DecodedField))
	out.Riot = To_Identity_Riot(fields[4].Value.([]*registry.DecodedField))
	out.Email = To_Identity_Email(fields[5].Value.([]*registry.DecodedField))
	out.PgpFingerprint = To_bytes(fields[6].Value.([]*registry.DecodedField))
	out.Image = To_Identity_Image(fields[7].Value.([]*registry.DecodedField))
	out.Twitter = To_Identity_Twitter(fields[8].Value.([]*registry.DecodedField))
	return out
}

// not seen "additional" optional: false repeated: false
// not seen "display" optional: false repeated: false
// not seen "legal" optional: false repeated: false
// not seen "web" optional: false repeated: false
// not seen "riot" optional: false repeated: false
// not seen "email" optional: false repeated: false
// not seen "image" optional: false repeated: false
// not seen "twitter" optional: false repeated: false

func To_Identity_PalletIdentityTypesBitFlags(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
	out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
	out.Fields = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
	out := &pbgear.Identity_PrimitiveTypesH256{}
	out.Identity = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_ProvideJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_ProvideJudgementCall {
	out := &pbgear.Identity_ProvideJudgementCall{}
	out.RegIndex = To_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Target = To_Identity_Target(fields[1].Value.([]*registry.DecodedField))
	out.Judgement = To_Identity_Judgement(fields[2].Value.([]*registry.DecodedField))
	out.Identity = To_Identity_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "reg_index" optional: false repeated: false
// not seen "judgement" optional: false repeated: false
// not seen "identity" optional: false repeated: false

func To_Identity_QuitSubCall(fields []*registry.DecodedField) *pbgear.Identity_QuitSubCall {
	out := &pbgear.Identity_QuitSubCall{}
	return out
}

func To_Identity_Raw(fields []*registry.DecodedField) *pbgear.Identity_Raw {
	out := &pbgear.Identity_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw0(fields []*registry.DecodedField) *pbgear.Identity_Raw0 {
	out := &pbgear.Identity_Raw0{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw1(fields []*registry.DecodedField) *pbgear.Identity_Raw1 {
	out := &pbgear.Identity_Raw1{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw10(fields []*registry.DecodedField) *pbgear.Identity_Raw10 {
	out := &pbgear.Identity_Raw10{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw11(fields []*registry.DecodedField) *pbgear.Identity_Raw11 {
	out := &pbgear.Identity_Raw11{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw12(fields []*registry.DecodedField) *pbgear.Identity_Raw12 {
	out := &pbgear.Identity_Raw12{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw13(fields []*registry.DecodedField) *pbgear.Identity_Raw13 {
	out := &pbgear.Identity_Raw13{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw14(fields []*registry.DecodedField) *pbgear.Identity_Raw14 {
	out := &pbgear.Identity_Raw14{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw15(fields []*registry.DecodedField) *pbgear.Identity_Raw15 {
	out := &pbgear.Identity_Raw15{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw16(fields []*registry.DecodedField) *pbgear.Identity_Raw16 {
	out := &pbgear.Identity_Raw16{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw17(fields []*registry.DecodedField) *pbgear.Identity_Raw17 {
	out := &pbgear.Identity_Raw17{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw18(fields []*registry.DecodedField) *pbgear.Identity_Raw18 {
	out := &pbgear.Identity_Raw18{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw19(fields []*registry.DecodedField) *pbgear.Identity_Raw19 {
	out := &pbgear.Identity_Raw19{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw2(fields []*registry.DecodedField) *pbgear.Identity_Raw2 {
	out := &pbgear.Identity_Raw2{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw20(fields []*registry.DecodedField) *pbgear.Identity_Raw20 {
	out := &pbgear.Identity_Raw20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw21(fields []*registry.DecodedField) *pbgear.Identity_Raw21 {
	out := &pbgear.Identity_Raw21{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw22(fields []*registry.DecodedField) *pbgear.Identity_Raw22 {
	out := &pbgear.Identity_Raw22{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw23(fields []*registry.DecodedField) *pbgear.Identity_Raw23 {
	out := &pbgear.Identity_Raw23{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw24(fields []*registry.DecodedField) *pbgear.Identity_Raw24 {
	out := &pbgear.Identity_Raw24{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw25(fields []*registry.DecodedField) *pbgear.Identity_Raw25 {
	out := &pbgear.Identity_Raw25{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw26(fields []*registry.DecodedField) *pbgear.Identity_Raw26 {
	out := &pbgear.Identity_Raw26{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw27(fields []*registry.DecodedField) *pbgear.Identity_Raw27 {
	out := &pbgear.Identity_Raw27{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw28(fields []*registry.DecodedField) *pbgear.Identity_Raw28 {
	out := &pbgear.Identity_Raw28{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw29(fields []*registry.DecodedField) *pbgear.Identity_Raw29 {
	out := &pbgear.Identity_Raw29{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw3(fields []*registry.DecodedField) *pbgear.Identity_Raw3 {
	out := &pbgear.Identity_Raw3{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw30(fields []*registry.DecodedField) *pbgear.Identity_Raw30 {
	out := &pbgear.Identity_Raw30{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw31(fields []*registry.DecodedField) *pbgear.Identity_Raw31 {
	out := &pbgear.Identity_Raw31{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw32(fields []*registry.DecodedField) *pbgear.Identity_Raw32 {
	out := &pbgear.Identity_Raw32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw4(fields []*registry.DecodedField) *pbgear.Identity_Raw4 {
	out := &pbgear.Identity_Raw4{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw5(fields []*registry.DecodedField) *pbgear.Identity_Raw5 {
	out := &pbgear.Identity_Raw5{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw6(fields []*registry.DecodedField) *pbgear.Identity_Raw6 {
	out := &pbgear.Identity_Raw6{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw7(fields []*registry.DecodedField) *pbgear.Identity_Raw7 {
	out := &pbgear.Identity_Raw7{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw8(fields []*registry.DecodedField) *pbgear.Identity_Raw8 {
	out := &pbgear.Identity_Raw8{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Raw9(fields []*registry.DecodedField) *pbgear.Identity_Raw9 {
	out := &pbgear.Identity_Raw9{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Reasonable(fields []*registry.DecodedField) *pbgear.Identity_Reasonable {
	out := &pbgear.Identity_Reasonable{}
	return out
}

func To_Identity_RemoveSubCall(fields []*registry.DecodedField) *pbgear.Identity_RemoveSubCall {
	out := &pbgear.Identity_RemoveSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_RenameSubCall(fields []*registry.DecodedField) *pbgear.Identity_RenameSubCall {
	out := &pbgear.Identity_RenameSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.([]*registry.DecodedField))
	out.Data = To_Identity_Data(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_RequestJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_RequestJudgementCall {
	out := &pbgear.Identity_RequestJudgementCall{}
	out.RegIndex = To_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.MaxFee = To_Identity_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "max_fee" optional: false repeated: false

func To_Identity_Riot(fields []*registry.DecodedField) *pbgear.Identity_Riot {
	out := &pbgear.Identity_Riot{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Riot_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_SetAccountIdCall(fields []*registry.DecodedField) *pbgear.Identity_SetAccountIdCall {
	out := &pbgear.Identity_SetAccountIdCall{}
	out.Index = To_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.New = To_Identity_New(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "new" optional: false repeated: false

func To_Identity_SetFeeCall(fields []*registry.DecodedField) *pbgear.Identity_SetFeeCall {
	out := &pbgear.Identity_SetFeeCall{}
	out.Index = To_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Fee = To_Identity_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_SetFieldsCall(fields []*registry.DecodedField) *pbgear.Identity_SetFieldsCall {
	out := &pbgear.Identity_SetFieldsCall{}
	out.Index = To_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	out.Fields = To_Identity_PalletIdentityTypesBitFlags(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "fields" optional: false repeated: false

func To_Identity_SetIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_SetIdentityCall {
	out := &pbgear.Identity_SetIdentityCall{}
	out.Info = To_Identity_PalletIdentitySimpleIdentityInfo(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "info" optional: false repeated: false

func To_Identity_SetSubsCall(fields []*registry.DecodedField) *pbgear.Identity_SetSubsCall {
	out := &pbgear.Identity_SetSubsCall{}
	out.Subs = To_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "subs" optional: false repeated: true
// optional primitive: subs
func To_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields []*registry.DecodedField) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	out := make([]*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData, len(fields))
	for _, f := range fields {
		out = append(out, To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Identity_Sha256(fields []*registry.DecodedField) *pbgear.Identity_Sha256 {
	out := &pbgear.Identity_Sha256{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_ShaThree256(fields []*registry.DecodedField) *pbgear.Identity_ShaThree256 {
	out := &pbgear.Identity_ShaThree256{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
	out := &pbgear.Identity_SpCoreCryptoAccountId32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Identity_Sub(fields []*registry.DecodedField) *pbgear.Identity_Sub {
	out := &pbgear.Identity_Sub{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Identity_Sub_Id{
			Id: To_Identity_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Identity_Sub_Index{
			Index: To_Identity_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Identity_Sub_Raw{
			Raw: To_Identity_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Identity_Sub_Address32{
			Address32: To_Identity_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Identity_Sub_Address20{
			Address20: To_Identity_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Target(fields []*registry.DecodedField) *pbgear.Identity_Target {
	out := &pbgear.Identity_Target{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Identity_Target_Id{
			Id: To_Identity_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Identity_Target_Index{
			Index: To_Identity_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Identity_Target_Raw{
			Raw: To_Identity_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Identity_Target_Address32{
			Address32: To_Identity_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Identity_Target_Address20{
			Address20: To_Identity_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_TupleNull(fields []*registry.DecodedField) *pbgear.Identity_TupleNull {
	out := &pbgear.Identity_TupleNull{}
	return out
}

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
	out.Value0 = To_Identity_Value0(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_Identity_Value1(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false
// not seen "value1" optional: false repeated: false

func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
	out.Value0 = To_Identity_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_Identity_Value1(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Identity_Twitter(fields []*registry.DecodedField) *pbgear.Identity_Twitter {
	out := &pbgear.Identity_Twitter{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Twitter_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Unknown(fields []*registry.DecodedField) *pbgear.Identity_Unknown {
	out := &pbgear.Identity_Unknown{}
	return out
}

func To_Identity_Value0(fields []*registry.DecodedField) *pbgear.Identity_Value0 {
	out := &pbgear.Identity_Value0{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Value0_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Value1(fields []*registry.DecodedField) *pbgear.Identity_Value1 {
	out := &pbgear.Identity_Value1{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Value1_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Web(fields []*registry.DecodedField) *pbgear.Identity_Web {
	out := &pbgear.Identity_Web{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Web_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ImOnline_HeartbeatCall(fields []*registry.DecodedField) *pbgear.ImOnline_HeartbeatCall {
	out := &pbgear.ImOnline_HeartbeatCall{}
	out.Heartbeat = To_ImOnline_PalletImOnlineHeartbeat(fields[0].Value.([]*registry.DecodedField))
	out.Signature = To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "heartbeat" optional: false repeated: false
// not seen "signature" optional: false repeated: false

func To_ImOnline_PalletImOnlineHeartbeat(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
	out.BlockNumber = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.SessionIndex = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.AuthorityIndex = To_uint32(fields[2].Value.([]*registry.DecodedField))
	out.ValidatorsLen = To_uint32(fields[3].Value.([]*registry.DecodedField))
	return out
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
	out.Signature = To_ImOnline_SpCoreSr25519Signature(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "signature" optional: false repeated: false

func To_ImOnline_SpCoreSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
	out := &pbgear.ImOnline_SpCoreSr25519Signature{}
	out.Signature = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Multisig_ApproveAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_ApproveAsMultiCall {
	out := &pbgear.Multisig_ApproveAsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.OtherSignatories = To_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	out.MaybeTimepoint = To_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField))
	out.CallHash = To_bytes(fields[3].Value.([]*registry.DecodedField))
	out.MaxWeight = To_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "other_signatories" optional: false repeated: true
// optional primitive: other_signatories
func To_repeated_Multisig_SpCoreCryptoAccountId32(fields []*registry.DecodedField) []*pbgear.Multisig_SpCoreCryptoAccountId32 {
	out := make([]*pbgear.Multisig_SpCoreCryptoAccountId32, len(fields))
	for _, f := range fields {
		out = append(out, To_Multisig_SpCoreCryptoAccountId32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

// not seen "maybe_timepoint" optional: true repeated: false
// optional field: maybe_timepoint
// not seen "max_weight" optional: false repeated: false

func To_Multisig_AsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiCall {
	out := &pbgear.Multisig_AsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.OtherSignatories = To_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	out.MaybeTimepoint = To_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Multisig_AsMultiCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Multisig_AsMultiCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Multisig_AsMultiCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Multisig_AsMultiCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Multisig_AsMultiCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Multisig_AsMultiCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Multisig_AsMultiCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Multisig_AsMultiCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Multisig_AsMultiCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Multisig_AsMultiCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Multisig_AsMultiCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Multisig_AsMultiCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Multisig_AsMultiCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Multisig_AsMultiCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Multisig_AsMultiCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Multisig_AsMultiCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	out.MaxWeight = To_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "call" optional: false repeated: false

func To_Multisig_AsMultiThreshold1Call(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiThreshold1Call {
	out := &pbgear.Multisig_AsMultiThreshold1Call{}
	out.OtherSignatories = To_repeated_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Multisig_AsMultiThreshold1Call_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Multisig_CancelAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_CancelAsMultiCall {
	out := &pbgear.Multisig_CancelAsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.OtherSignatories = To_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	out.Timepoint = To_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField))
	out.CallHash = To_bytes(fields[3].Value.([]*registry.DecodedField))
	return out
}

func To_Multisig_CompactUint64(fields []*registry.DecodedField) *pbgear.Multisig_CompactUint64 {
	out := &pbgear.Multisig_CompactUint64{}
	out.Value = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Multisig_PalletMultisigTimepoint(fields []*registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
	out := &pbgear.Multisig_PalletMultisigTimepoint{}
	out.Height = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Multisig_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Multisig_SpCoreCryptoAccountId32 {
	out := &pbgear.Multisig_SpCoreCryptoAccountId32{}
	out.OtherSignatories = To_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Multisig_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
	out := &pbgear.Multisig_SpWeightsWeightV2Weight{}
	out.RefTime = To_Multisig_CompactUint64(fields[0].Value.([]*registry.DecodedField))
	out.ProofSize = To_Multisig_CompactUint64(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "ref_time" optional: false repeated: false

func To_NominationPools_Address20(fields []*registry.DecodedField) *pbgear.NominationPools_Address20 {
	out := &pbgear.NominationPools_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Address32(fields []*registry.DecodedField) *pbgear.NominationPools_Address32 {
	out := &pbgear.NominationPools_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_AdjustPoolDepositCall(fields []*registry.DecodedField) *pbgear.NominationPools_AdjustPoolDepositCall {
	out := &pbgear.NominationPools_AdjustPoolDepositCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Blocked(fields []*registry.DecodedField) *pbgear.NominationPools_Blocked {
	out := &pbgear.NominationPools_Blocked{}
	return out
}

func To_NominationPools_BondExtraCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraCall {
	out := &pbgear.NominationPools_BondExtraCall{}
	out.Extra = To_NominationPools_Extra(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "extra" optional: false repeated: false

func To_NominationPools_BondExtraOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraOtherCall {
	out := &pbgear.NominationPools_BondExtraOtherCall{}
	out.Member = To_NominationPools_Member(fields[0].Value.([]*registry.DecodedField))
	out.Extra = To_NominationPools_Extra(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "member" optional: false repeated: false

func To_NominationPools_Bouncer(fields []*registry.DecodedField) *pbgear.NominationPools_Bouncer {
	out := &pbgear.NominationPools_Bouncer{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.NominationPools_Bouncer_Id{
			Id: To_NominationPools_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.NominationPools_Bouncer_Index{
			Index: To_NominationPools_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.NominationPools_Bouncer_Raw{
			Raw: To_NominationPools_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.NominationPools_Bouncer_Address32{
			Address32: To_NominationPools_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.NominationPools_Bouncer_Address20{
			Address20: To_NominationPools_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_NominationPools_ChillCall(fields []*registry.DecodedField) *pbgear.NominationPools_ChillCall {
	out := &pbgear.NominationPools_ChillCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_ClaimCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimCommissionCall {
	out := &pbgear.NominationPools_ClaimCommissionCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_ClaimPayoutCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutCall {
	out := &pbgear.NominationPools_ClaimPayoutCall{}
	return out
}

func To_NominationPools_ClaimPayoutOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutOtherCall {
	out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
	out.Other = To_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "other" optional: false repeated: false

func To_NominationPools_CompactString(fields []*registry.DecodedField) *pbgear.NominationPools_CompactString {
	out := &pbgear.NominationPools_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_CompactTupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
	out := &pbgear.NominationPools_CompactTupleNull{}
	out.Value = To_NominationPools_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_NominationPools_CreateCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateCall {
	out := &pbgear.NominationPools_CreateCall{}
	out.Amount = To_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Root = To_NominationPools_Root(fields[1].Value.([]*registry.DecodedField))
	out.Nominator = To_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField))
	out.Bouncer = To_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "amount" optional: false repeated: false
// not seen "root" optional: false repeated: false
// not seen "nominator" optional: false repeated: false
// not seen "bouncer" optional: false repeated: false

func To_NominationPools_CreateWithPoolIdCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateWithPoolIdCall {
	out := &pbgear.NominationPools_CreateWithPoolIdCall{}
	out.Amount = To_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Root = To_NominationPools_Root(fields[1].Value.([]*registry.DecodedField))
	out.Nominator = To_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField))
	out.Bouncer = To_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField))
	out.PoolId = To_uint32(fields[4].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Destroying(fields []*registry.DecodedField) *pbgear.NominationPools_Destroying {
	out := &pbgear.NominationPools_Destroying{}
	return out
}

func To_NominationPools_Extra(fields []*registry.DecodedField) *pbgear.NominationPools_Extra {
	out := &pbgear.NominationPools_Extra{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_Extra_FreeBalance{
			FreeBalance: To_NominationPools_FreeBalance(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_FreeBalance(fields []*registry.DecodedField) *pbgear.NominationPools_FreeBalance {
	out := &pbgear.NominationPools_FreeBalance{}
	out.Value0 = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_GlobalMaxCommission(fields []*registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
	out := &pbgear.NominationPools_GlobalMaxCommission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_GlobalMaxCommission_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Id(fields []*registry.DecodedField) *pbgear.NominationPools_Id {
	out := &pbgear.NominationPools_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Index(fields []*registry.DecodedField) *pbgear.NominationPools_Index {
	out := &pbgear.NominationPools_Index{}
	out.Value0 = To_NominationPools_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_NominationPools_JoinCall(fields []*registry.DecodedField) *pbgear.NominationPools_JoinCall {
	out := &pbgear.NominationPools_JoinCall{}
	out.Amount = To_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.PoolId = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_MaxMembers(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembers {
	out := &pbgear.NominationPools_MaxMembers{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxMembers_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MaxMembersPerPool(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
	out := &pbgear.NominationPools_MaxMembersPerPool{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxMembersPerPool_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MaxPools(fields []*registry.DecodedField) *pbgear.NominationPools_MaxPools {
	out := &pbgear.NominationPools_MaxPools{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxPools_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Member(fields []*registry.DecodedField) *pbgear.NominationPools_Member {
	out := &pbgear.NominationPools_Member{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.NominationPools_Member_Id{
			Id: To_NominationPools_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.NominationPools_Member_Index{
			Index: To_NominationPools_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.NominationPools_Member_Raw{
			Raw: To_NominationPools_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.NominationPools_Member_Address32{
			Address32: To_NominationPools_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.NominationPools_Member_Address20{
			Address20: To_NominationPools_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MemberAccount(fields []*registry.DecodedField) *pbgear.NominationPools_MemberAccount {
	out := &pbgear.NominationPools_MemberAccount{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.NominationPools_MemberAccount_Id{
			Id: To_NominationPools_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.NominationPools_MemberAccount_Index{
			Index: To_NominationPools_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.NominationPools_MemberAccount_Raw{
			Raw: To_NominationPools_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.NominationPools_MemberAccount_Address32{
			Address32: To_NominationPools_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.NominationPools_MemberAccount_Address20{
			Address20: To_NominationPools_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MinCreateBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
	out := &pbgear.NominationPools_MinCreateBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MinCreateBond_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MinJoinBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
	out := &pbgear.NominationPools_MinJoinBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MinJoinBond_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewBouncer(fields []*registry.DecodedField) *pbgear.NominationPools_NewBouncer {
	out := &pbgear.NominationPools_NewBouncer{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewBouncer_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewNominator(fields []*registry.DecodedField) *pbgear.NominationPools_NewNominator {
	out := &pbgear.NominationPools_NewNominator{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewNominator_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewRoot(fields []*registry.DecodedField) *pbgear.NominationPools_NewRoot {
	out := &pbgear.NominationPools_NewRoot{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewRoot_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NominateCall(fields []*registry.DecodedField) *pbgear.NominationPools_NominateCall {
	out := &pbgear.NominationPools_NominateCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Validators = To_repeated_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Nominator(fields []*registry.DecodedField) *pbgear.NominationPools_Nominator {
	out := &pbgear.NominationPools_Nominator{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.NominationPools_Nominator_Id{
			Id: To_NominationPools_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.NominationPools_Nominator_Index{
			Index: To_NominationPools_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.NominationPools_Nominator_Raw{
			Raw: To_NominationPools_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.NominationPools_Nominator_Address32{
			Address32: To_NominationPools_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.NominationPools_Nominator_Address20{
			Address20: To_NominationPools_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Noop(fields []*registry.DecodedField) *pbgear.NominationPools_Noop {
	out := &pbgear.NominationPools_Noop{}
	return out
}

func To_NominationPools_Open(fields []*registry.DecodedField) *pbgear.NominationPools_Open {
	out := &pbgear.NominationPools_Open{}
	return out
}

func To_NominationPools_PalletNominationPoolsCommissionChangeRate(fields []*registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
	out.MaxIncrease = To_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField))
	out.MinDelay = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "max_increase" optional: false repeated: false

func To_NominationPools_Permission(fields []*registry.DecodedField) *pbgear.NominationPools_Permission {
	out := &pbgear.NominationPools_Permission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_Permission_Permissioned{
			Permissioned: To_NominationPools_Permissioned(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Permissioned(fields []*registry.DecodedField) *pbgear.NominationPools_Permissioned {
	out := &pbgear.NominationPools_Permissioned{}
	return out
}

func To_NominationPools_PermissionlessAll(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessAll {
	out := &pbgear.NominationPools_PermissionlessAll{}
	return out
}

func To_NominationPools_PermissionlessCompound(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessCompound {
	out := &pbgear.NominationPools_PermissionlessCompound{}
	return out
}

func To_NominationPools_PermissionlessWithdraw(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessWithdraw {
	out := &pbgear.NominationPools_PermissionlessWithdraw{}
	return out
}

func To_NominationPools_PoolWithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
	out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.NumSlashingSpans = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Raw(fields []*registry.DecodedField) *pbgear.NominationPools_Raw {
	out := &pbgear.NominationPools_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_Remove(fields []*registry.DecodedField) *pbgear.NominationPools_Remove {
	out := &pbgear.NominationPools_Remove{}
	return out
}

func To_NominationPools_Rewards(fields []*registry.DecodedField) *pbgear.NominationPools_Rewards {
	out := &pbgear.NominationPools_Rewards{}
	return out
}

func To_NominationPools_Root(fields []*registry.DecodedField) *pbgear.NominationPools_Root {
	out := &pbgear.NominationPools_Root{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.NominationPools_Root_Id{
			Id: To_NominationPools_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.NominationPools_Root_Index{
			Index: To_NominationPools_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.NominationPools_Root_Raw{
			Raw: To_NominationPools_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.NominationPools_Root_Address32{
			Address32: To_NominationPools_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.NominationPools_Root_Address20{
			Address20: To_NominationPools_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Set(fields []*registry.DecodedField) *pbgear.NominationPools_Set {
	out := &pbgear.NominationPools_Set{}
	out.Value0 = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_SetClaimPermissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetClaimPermissionCall {
	out := &pbgear.NominationPools_SetClaimPermissionCall{}
	out.Permission = To_NominationPools_Permission(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "permission" optional: false repeated: false

func To_NominationPools_SetCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionCall {
	out := &pbgear.NominationPools_SetCommissionCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.NewCommission = To_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "new_commission" optional: true repeated: false
// optional field: new_commission

func To_NominationPools_SetCommissionChangeRateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionChangeRateCall {
	out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.ChangeRate = To_NominationPools_PalletNominationPoolsCommissionChangeRate(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "change_rate" optional: false repeated: false

func To_NominationPools_SetCommissionMaxCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionMaxCall {
	out := &pbgear.NominationPools_SetCommissionMaxCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaxCommission = To_NominationPools_SpArithmeticPerThingsPerbill(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_SetConfigsCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetConfigsCall {
	out := &pbgear.NominationPools_SetConfigsCall{}
	out.MinJoinBond = To_NominationPools_MinJoinBond(fields[0].Value.([]*registry.DecodedField))
	out.MinCreateBond = To_NominationPools_MinCreateBond(fields[1].Value.([]*registry.DecodedField))
	out.MaxPools = To_NominationPools_MaxPools(fields[2].Value.([]*registry.DecodedField))
	out.MaxMembers = To_NominationPools_MaxMembers(fields[3].Value.([]*registry.DecodedField))
	out.MaxMembersPerPool = To_NominationPools_MaxMembersPerPool(fields[4].Value.([]*registry.DecodedField))
	out.GlobalMaxCommission = To_NominationPools_GlobalMaxCommission(fields[5].Value.([]*registry.DecodedField))
	return out
}

// not seen "min_join_bond" optional: false repeated: false
// not seen "min_create_bond" optional: false repeated: false
// not seen "max_pools" optional: false repeated: false
// not seen "max_members" optional: false repeated: false
// not seen "max_members_per_pool" optional: false repeated: false
// not seen "global_max_commission" optional: false repeated: false

func To_NominationPools_SetMetadataCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetMetadataCall {
	out := &pbgear.NominationPools_SetMetadataCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Metadata = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_SetStateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetStateCall {
	out := &pbgear.NominationPools_SetStateCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.State = To_NominationPools_State(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "state" optional: false repeated: false

func To_NominationPools_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
	out := &pbgear.NominationPools_SpArithmeticPerThingsPerbill{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_SpCoreCryptoAccountId32 {
	out := &pbgear.NominationPools_SpCoreCryptoAccountId32{}
	out.Validators = To_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_State(fields []*registry.DecodedField) *pbgear.NominationPools_State {
	out := &pbgear.NominationPools_State{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_State_Open{
			Open: To_NominationPools_Open(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_TupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_TupleNull {
	out := &pbgear.NominationPools_TupleNull{}
	return out
}

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
	out.Value0 = To_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_NominationPools_UnbondCall(fields []*registry.DecodedField) *pbgear.NominationPools_UnbondCall {
	out := &pbgear.NominationPools_UnbondCall{}
	out.MemberAccount = To_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField))
	out.UnbondingPoints = To_NominationPools_CompactString(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "member_account" optional: false repeated: false

func To_NominationPools_UpdateRolesCall(fields []*registry.DecodedField) *pbgear.NominationPools_UpdateRolesCall {
	out := &pbgear.NominationPools_UpdateRolesCall{}
	out.PoolId = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.NewRoot = To_NominationPools_NewRoot(fields[1].Value.([]*registry.DecodedField))
	out.NewNominator = To_NominationPools_NewNominator(fields[2].Value.([]*registry.DecodedField))
	out.NewBouncer = To_NominationPools_NewBouncer(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "new_root" optional: false repeated: false
// not seen "new_nominator" optional: false repeated: false
// not seen "new_bouncer" optional: false repeated: false

func To_NominationPools_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_WithdrawUnbondedCall {
	out := &pbgear.NominationPools_WithdrawUnbondedCall{}
	out.MemberAccount = To_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField))
	out.NumSlashingSpans = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_None(fields []*registry.DecodedField) *pbgear.None {
	out := &pbgear.None{}
	return out
}

func To_Preimage_EnsureUpdatedCall(fields []*registry.DecodedField) *pbgear.Preimage_EnsureUpdatedCall {
	out := &pbgear.Preimage_EnsureUpdatedCall{}
	out.Hashes = To_repeated_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "hashes" optional: false repeated: true
// optional primitive: hashes
func To_repeated_Preimage_PrimitiveTypesH256(fields []*registry.DecodedField) []*pbgear.Preimage_PrimitiveTypesH256 {
	out := make([]*pbgear.Preimage_PrimitiveTypesH256, len(fields))
	for _, f := range fields {
		out = append(out, To_Preimage_PrimitiveTypesH256(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Preimage_NotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_NotePreimageCall {
	out := &pbgear.Preimage_NotePreimageCall{}
	out.Bytes = fields[0].Value.(uint8)
	return out
}

func To_Preimage_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
	out := &pbgear.Preimage_PrimitiveTypesH256{}
	out.Hash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Preimage_RequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_RequestPreimageCall {
	out := &pbgear.Preimage_RequestPreimageCall{}
	out.Hash = To_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Preimage_UnnotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnnotePreimageCall {
	out := &pbgear.Preimage_UnnotePreimageCall{}
	out.Hash = To_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Preimage_UnrequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnrequestPreimageCall {
	out := &pbgear.Preimage_UnrequestPreimageCall{}
	out.Hash = To_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_PrimaryAndSecondaryPlainSlots(fields []*registry.DecodedField) *pbgear.PrimaryAndSecondaryPlainSlots {
	out := &pbgear.PrimaryAndSecondaryPlainSlots{}
	return out
}

func To_PrimaryAndSecondaryVRFSlots(fields []*registry.DecodedField) *pbgear.PrimaryAndSecondaryVRFSlots {
	out := &pbgear.PrimaryAndSecondaryVRFSlots{}
	return out
}

func To_PrimarySlots(fields []*registry.DecodedField) *pbgear.PrimarySlots {
	out := &pbgear.PrimarySlots{}
	return out
}

func To_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.PrimitiveTypesH256 {
	out := &pbgear.PrimitiveTypesH256{}
	out.TargetHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_AddProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_AddProxyCall {
	out := &pbgear.Proxy_AddProxyCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField))
	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField))
	out.Delay = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "delegate" optional: false repeated: false
// not seen "proxy_type" optional: false repeated: false

func To_Proxy_Address20(fields []*registry.DecodedField) *pbgear.Proxy_Address20 {
	out := &pbgear.Proxy_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_Address32(fields []*registry.DecodedField) *pbgear.Proxy_Address32 {
	out := &pbgear.Proxy_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_AnnounceCall(fields []*registry.DecodedField) *pbgear.Proxy_AnnounceCall {
	out := &pbgear.Proxy_AnnounceCall{}
	out.Real = To_Proxy_Real(fields[0].Value.([]*registry.DecodedField))
	out.CallHash = To_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "real" optional: false repeated: false
// not seen "call_hash" optional: false repeated: false

func To_Proxy_Any(fields []*registry.DecodedField) *pbgear.Proxy_Any {
	out := &pbgear.Proxy_Any{}
	return out
}

func To_Proxy_CancelProxy(fields []*registry.DecodedField) *pbgear.Proxy_CancelProxy {
	out := &pbgear.Proxy_CancelProxy{}
	return out
}

func To_Proxy_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
	out := &pbgear.Proxy_CompactTupleNull{}
	out.Value = To_Proxy_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Proxy_CompactUint32(fields []*registry.DecodedField) *pbgear.Proxy_CompactUint32 {
	out := &pbgear.Proxy_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_CreatePureCall(fields []*registry.DecodedField) *pbgear.Proxy_CreatePureCall {
	out := &pbgear.Proxy_CreatePureCall{}
	out.ProxyType = To_Proxy_ProxyType(fields[0].Value.([]*registry.DecodedField))
	out.Delay = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_Delegate(fields []*registry.DecodedField) *pbgear.Proxy_Delegate {
	out := &pbgear.Proxy_Delegate{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Proxy_Delegate_Id{
			Id: To_Proxy_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Proxy_Delegate_Index{
			Index: To_Proxy_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Proxy_Delegate_Raw{
			Raw: To_Proxy_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Proxy_Delegate_Address32{
			Address32: To_Proxy_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Proxy_Delegate_Address20{
			Address20: To_Proxy_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Proxy_ForceProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ForceProxyType {
	out := &pbgear.Proxy_ForceProxyType{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Proxy_ForceProxyType_Any{
			Any: To_Proxy_Any(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_Governance(fields []*registry.DecodedField) *pbgear.Proxy_Governance {
	out := &pbgear.Proxy_Governance{}
	return out
}

func To_Proxy_Id(fields []*registry.DecodedField) *pbgear.Proxy_Id {
	out := &pbgear.Proxy_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_IdentityJudgement(fields []*registry.DecodedField) *pbgear.Proxy_IdentityJudgement {
	out := &pbgear.Proxy_IdentityJudgement{}
	return out
}

func To_Proxy_Index(fields []*registry.DecodedField) *pbgear.Proxy_Index {
	out := &pbgear.Proxy_Index{}
	out.Value0 = To_Proxy_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Proxy_KillPureCall(fields []*registry.DecodedField) *pbgear.Proxy_KillPureCall {
	out := &pbgear.Proxy_KillPureCall{}
	out.Spawner = To_Proxy_Spawner(fields[0].Value.([]*registry.DecodedField))
	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[2].Value.([]*registry.DecodedField))
	out.Height = To_Proxy_CompactUint32(fields[3].Value.([]*registry.DecodedField))
	out.ExtIndex = To_Proxy_CompactUint32(fields[4].Value.([]*registry.DecodedField))
	return out
}

// not seen "spawner" optional: false repeated: false
// not seen "height" optional: false repeated: false

func To_Proxy_NonTransfer(fields []*registry.DecodedField) *pbgear.Proxy_NonTransfer {
	out := &pbgear.Proxy_NonTransfer{}
	return out
}

func To_Proxy_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
	out := &pbgear.Proxy_PrimitiveTypesH256{}
	out.CallHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_ProxyAnnouncedCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyAnnouncedCall {
	out := &pbgear.Proxy_ProxyAnnouncedCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField))
	out.Real = To_Proxy_Real(fields[1].Value.([]*registry.DecodedField))
	out.ForceProxyType = To_optional_Proxy_ForceProxyType(fields[2].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Proxy_ProxyAnnouncedCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "force_proxy_type" optional: true repeated: false
// optional field: force_proxy_type
// not seen "call" optional: false repeated: false

func To_Proxy_ProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyCall {
	out := &pbgear.Proxy_ProxyCall{}
	out.Real = To_Proxy_Real(fields[0].Value.([]*registry.DecodedField))
	out.ForceProxyType = To_optional_Proxy_ForceProxyType(fields[1].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Proxy_ProxyCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Proxy_ProxyCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Proxy_ProxyCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Proxy_ProxyCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Proxy_ProxyCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Proxy_ProxyCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Proxy_ProxyCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Proxy_ProxyCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Proxy_ProxyCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Proxy_ProxyCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Proxy_ProxyCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Proxy_ProxyCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Proxy_ProxyCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Proxy_ProxyCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Proxy_ProxyCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Proxy_ProxyCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_ProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ProxyType {
	out := &pbgear.Proxy_ProxyType{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Proxy_ProxyType_Any{
			Any: To_Proxy_Any(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_Raw(fields []*registry.DecodedField) *pbgear.Proxy_Raw {
	out := &pbgear.Proxy_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_Real(fields []*registry.DecodedField) *pbgear.Proxy_Real {
	out := &pbgear.Proxy_Real{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Proxy_Real_Id{
			Id: To_Proxy_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Proxy_Real_Index{
			Index: To_Proxy_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Proxy_Real_Raw{
			Raw: To_Proxy_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Proxy_Real_Address32{
			Address32: To_Proxy_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Proxy_Real_Address20{
			Address20: To_Proxy_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_RejectAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RejectAnnouncementCall {
	out := &pbgear.Proxy_RejectAnnouncementCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField))
	out.CallHash = To_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_RemoveAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveAnnouncementCall {
	out := &pbgear.Proxy_RemoveAnnouncementCall{}
	out.Real = To_Proxy_Real(fields[0].Value.([]*registry.DecodedField))
	out.CallHash = To_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_RemoveProxiesCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxiesCall {
	out := &pbgear.Proxy_RemoveProxiesCall{}
	return out
}

func To_Proxy_RemoveProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxyCall {
	out := &pbgear.Proxy_RemoveProxyCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField))
	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField))
	out.Delay = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_Proxy_Spawner(fields []*registry.DecodedField) *pbgear.Proxy_Spawner {
	out := &pbgear.Proxy_Spawner{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Proxy_Spawner_Id{
			Id: To_Proxy_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Proxy_Spawner_Index{
			Index: To_Proxy_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Proxy_Spawner_Raw{
			Raw: To_Proxy_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Proxy_Spawner_Address32{
			Address32: To_Proxy_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Proxy_Spawner_Address20{
			Address20: To_Proxy_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_Staking(fields []*registry.DecodedField) *pbgear.Proxy_Staking {
	out := &pbgear.Proxy_Staking{}
	return out
}

func To_Proxy_TupleNull(fields []*registry.DecodedField) *pbgear.Proxy_TupleNull {
	out := &pbgear.Proxy_TupleNull{}
	return out
}

func To_Referenda_After(fields []*registry.DecodedField) *pbgear.Referenda_After {
	out := &pbgear.Referenda_After{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_At(fields []*registry.DecodedField) *pbgear.Referenda_At {
	out := &pbgear.Referenda_At{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_CancelCall(fields []*registry.DecodedField) *pbgear.Referenda_CancelCall {
	out := &pbgear.Referenda_CancelCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
	out := &pbgear.Referenda_EnactmentMoment{}

	switch {
	case matchFields(fields, []int64{4}):
		out.Value = &pbgear.Referenda_EnactmentMoment_At{
			At: To_Referenda_At(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{4}):
		out.Value = &pbgear.Referenda_EnactmentMoment_After{
			After: To_Referenda_After(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Referenda_Inline(fields []*registry.DecodedField) *pbgear.Referenda_Inline {
	out := &pbgear.Referenda_Inline{}
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_KillCall(fields []*registry.DecodedField) *pbgear.Referenda_KillCall {
	out := &pbgear.Referenda_KillCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_Legacy(fields []*registry.DecodedField) *pbgear.Referenda_Legacy {
	out := &pbgear.Referenda_Legacy{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_Lookup(fields []*registry.DecodedField) *pbgear.Referenda_Lookup {
	out := &pbgear.Referenda_Lookup{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	out.Len = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.Referenda_NudgeReferendumCall {
	out := &pbgear.Referenda_NudgeReferendumCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.Referenda_OneFewerDecidingCall {
	out := &pbgear.Referenda_OneFewerDecidingCall{}
	out.Track = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_Origins(fields []*registry.DecodedField) *pbgear.Referenda_Origins {
	out := &pbgear.Referenda_Origins{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_PlaceDecisionDepositCall {
	out := &pbgear.Referenda_PlaceDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
	out := &pbgear.Referenda_PrimitiveTypesH256{}
	out.MaybeHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_Proposal(fields []*registry.DecodedField) *pbgear.Referenda_Proposal {
	out := &pbgear.Referenda_Proposal{}

	switch {
	case matchFields(fields, []int64{12}):
		out.Value = &pbgear.Referenda_Proposal_Legacy{
			Legacy: To_Referenda_Legacy(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{279}):
		out.Value = &pbgear.Referenda_Proposal_Inline{
			Inline: To_Referenda_Inline(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{12, 4}):
		out.Value = &pbgear.Referenda_Proposal_Lookup{
			Lookup: To_Referenda_Lookup(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Referenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
	out := &pbgear.Referenda_ProposalOrigin{}

	switch {
	case matchFields(fields, []int64{121}):
		out.Value = &pbgear.Referenda_ProposalOrigin_System{
			System: To_Referenda_System(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{122}):
		out.Value = &pbgear.Referenda_ProposalOrigin_Origins{
			Origins: To_Referenda_Origins(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{123}):
		out.Value = &pbgear.Referenda_ProposalOrigin_Void{
			Void: To_Referenda_Void(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Referenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundDecisionDepositCall {
	out := &pbgear.Referenda_RefundDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundSubmissionDepositCall {
	out := &pbgear.Referenda_RefundSubmissionDepositCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.Referenda_SetMetadataCall {
	out := &pbgear.Referenda_SetMetadataCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaybeHash = To_optional_Referenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "maybe_hash" optional: true repeated: false
// optional field: maybe_hash

func To_Referenda_SubmitCall(fields []*registry.DecodedField) *pbgear.Referenda_SubmitCall {
	out := &pbgear.Referenda_SubmitCall{}
	out.ProposalOrigin = To_Referenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField))
	out.Proposal = To_Referenda_Proposal(fields[1].Value.([]*registry.DecodedField))
	out.EnactmentMoment = To_Referenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "proposal_origin" optional: false repeated: false
// not seen "proposal" optional: false repeated: false
// not seen "enactment_moment" optional: false repeated: false

func To_Referenda_System(fields []*registry.DecodedField) *pbgear.Referenda_System {
	out := &pbgear.Referenda_System{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Referenda_Void(fields []*registry.DecodedField) *pbgear.Referenda_Void {
	out := &pbgear.Referenda_Void{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Root(fields []*registry.DecodedField) *pbgear.Root {
	out := &pbgear.Root{}
	return out
}

func To_Scheduler_CancelCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelCall {
	out := &pbgear.Scheduler_CancelCall{}
	out.When = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Index = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Scheduler_CancelNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelNamedCall {
	out := &pbgear.Scheduler_CancelNamedCall{}
	out.Id = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Scheduler_ScheduleAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleAfterCall {
	out := &pbgear.Scheduler_ScheduleAfterCall{}
	out.After = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField))
	out.Priority = To_uint32(fields[2].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Scheduler_ScheduleAfterCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "maybe_periodic" optional: true repeated: false
// optional field: maybe_periodic
// not seen "call" optional: false repeated: false

func To_Scheduler_ScheduleCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleCall {
	out := &pbgear.Scheduler_ScheduleCall{}
	out.When = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField))
	out.Priority = To_uint32(fields[2].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Scheduler_ScheduleCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Scheduler_ScheduleCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Scheduler_ScheduleCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Scheduler_ScheduleCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Scheduler_ScheduleCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Scheduler_ScheduleCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Scheduler_ScheduleCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Scheduler_ScheduleCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Scheduler_ScheduleCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Scheduler_ScheduleCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Scheduler_ScheduleCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Scheduler_ScheduleNamedAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedAfterCall {
	out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
	out.Id = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.After = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField))
	out.Priority = To_uint32(fields[3].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Scheduler_ScheduleNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedCall {
	out := &pbgear.Scheduler_ScheduleNamedCall{}
	out.Id = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.When = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField))
	out.Priority = To_uint32(fields[3].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Scheduler_ScheduleNamedCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Scheduler_TupleUint32Uint32(fields []*registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
	out := &pbgear.Scheduler_TupleUint32Uint32{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Session_PalletImOnlineSr25519AppSr25519Public(fields []*registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
	out := &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{}
	out.ImOnline = To_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "im_online" optional: false repeated: false

func To_Session_PurgeKeysCall(fields []*registry.DecodedField) *pbgear.Session_PurgeKeysCall {
	out := &pbgear.Session_PurgeKeysCall{}
	return out
}

func To_Session_SetKeysCall(fields []*registry.DecodedField) *pbgear.Session_SetKeysCall {
	out := &pbgear.Session_SetKeysCall{}
	out.Keys = To_Session_VaraRuntimeSessionKeys(fields[0].Value.([]*registry.DecodedField))
	out.Proof = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "keys" optional: false repeated: false

func To_Session_SpAuthorityDiscoveryAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
	out := &pbgear.Session_SpAuthorityDiscoveryAppPublic{}
	out.AuthorityDiscovery = To_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Session_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
	out := &pbgear.Session_SpConsensusBabeAppPublic{}
	out.Babe = To_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Session_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
	out := &pbgear.Session_SpConsensusGrandpaAppPublic{}
	out.Grandpa = To_Session_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "grandpa" optional: false repeated: false

func To_Session_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
	out := &pbgear.Session_SpCoreEd25519Public{}
	out.Grandpa = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Session_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
	out := &pbgear.Session_SpCoreSr25519Public{}
	out.Babe = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Session_VaraRuntimeSessionKeys(fields []*registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
	out := &pbgear.Session_VaraRuntimeSessionKeys{}
	out.Babe = To_Session_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField))
	out.Grandpa = To_Session_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField))
	out.ImOnline = To_Session_PalletImOnlineSr25519AppSr25519Public(fields[2].Value.([]*registry.DecodedField))
	out.AuthorityDiscovery = To_Session_SpAuthorityDiscoveryAppPublic(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "babe" optional: false repeated: false
// not seen "grandpa" optional: false repeated: false
// not seen "im_online" optional: false repeated: false
// not seen "authority_discovery" optional: false repeated: false

func To_Signed(fields []*registry.DecodedField) *pbgear.Signed {
	out := &pbgear.Signed{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.SpArithmeticPerThingsPerbill {
	out := &pbgear.SpArithmeticPerThingsPerbill{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_SpArithmeticPerThingsPercent(fields []*registry.DecodedField) *pbgear.SpArithmeticPerThingsPercent {
	out := &pbgear.SpArithmeticPerThingsPercent{}
	out.Value0 = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.SpConsensusGrandpaAppPublic {
	out := &pbgear.SpConsensusGrandpaAppPublic{}
	out.Identity = To_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "identity" optional: false repeated: false

func To_SpConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.SpConsensusGrandpaAppSignature {
	out := &pbgear.SpConsensusGrandpaAppSignature{}
	out.Value1 = To_SpCoreEd25519Signature(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value1" optional: false repeated: false

func To_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.SpCoreCryptoAccountId32 {
	out := &pbgear.SpCoreCryptoAccountId32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.SpCoreEd25519Public {
	out := &pbgear.SpCoreEd25519Public{}
	out.Identity = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_SpCoreEd25519Signature(fields []*registry.DecodedField) *pbgear.SpCoreEd25519Signature {
	out := &pbgear.SpCoreEd25519Signature{}
	out.Value1 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_Address20(fields []*registry.DecodedField) *pbgear.StakingRewards_Address20 {
	out := &pbgear.StakingRewards_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_Address32(fields []*registry.DecodedField) *pbgear.StakingRewards_Address32 {
	out := &pbgear.StakingRewards_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_AlignSupplyCall(fields []*registry.DecodedField) *pbgear.StakingRewards_AlignSupplyCall {
	out := &pbgear.StakingRewards_AlignSupplyCall{}
	out.Target = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_CompactTupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
	out := &pbgear.StakingRewards_CompactTupleNull{}
	out.Value = To_StakingRewards_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_StakingRewards_ForceRefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_ForceRefillCall {
	out := &pbgear.StakingRewards_ForceRefillCall{}
	out.From = To_StakingRewards_From(fields[0].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "from" optional: false repeated: false

func To_StakingRewards_From(fields []*registry.DecodedField) *pbgear.StakingRewards_From {
	out := &pbgear.StakingRewards_From{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.StakingRewards_From_Id{
			Id: To_StakingRewards_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.StakingRewards_From_Index{
			Index: To_StakingRewards_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.StakingRewards_From_Raw{
			Raw: To_StakingRewards_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.StakingRewards_From_Address32{
			Address32: To_StakingRewards_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.StakingRewards_From_Address20{
			Address20: To_StakingRewards_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_StakingRewards_Id(fields []*registry.DecodedField) *pbgear.StakingRewards_Id {
	out := &pbgear.StakingRewards_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_Index(fields []*registry.DecodedField) *pbgear.StakingRewards_Index {
	out := &pbgear.StakingRewards_Index{}
	out.Value0 = To_StakingRewards_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_StakingRewards_Raw(fields []*registry.DecodedField) *pbgear.StakingRewards_Raw {
	out := &pbgear.StakingRewards_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_RefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_RefillCall {
	out := &pbgear.StakingRewards_RefillCall{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_StakingRewards_To(fields []*registry.DecodedField) *pbgear.StakingRewards_To {
	out := &pbgear.StakingRewards_To{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.StakingRewards_To_Id{
			Id: To_StakingRewards_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.StakingRewards_To_Index{
			Index: To_StakingRewards_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.StakingRewards_To_Raw{
			Raw: To_StakingRewards_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.StakingRewards_To_Address32{
			Address32: To_StakingRewards_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.StakingRewards_To_Address20{
			Address20: To_StakingRewards_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_StakingRewards_TupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_TupleNull {
	out := &pbgear.StakingRewards_TupleNull{}
	return out
}

func To_StakingRewards_WithdrawCall(fields []*registry.DecodedField) *pbgear.StakingRewards_WithdrawCall {
	out := &pbgear.StakingRewards_WithdrawCall{}
	out.To = To_StakingRewards_To(fields[0].Value.([]*registry.DecodedField))
	out.Value = To_string(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "to" optional: false repeated: false

func To_Staking_Account(fields []*registry.DecodedField) *pbgear.Staking_Account {
	out := &pbgear.Staking_Account{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Address20(fields []*registry.DecodedField) *pbgear.Staking_Address20 {
	out := &pbgear.Staking_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Address32(fields []*registry.DecodedField) *pbgear.Staking_Address32 {
	out := &pbgear.Staking_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_BondCall(fields []*registry.DecodedField) *pbgear.Staking_BondCall {
	out := &pbgear.Staking_BondCall{}
	out.Value = To_Staking_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Payee = To_Staking_Payee(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false
// not seen "payee" optional: false repeated: false

func To_Staking_BondExtraCall(fields []*registry.DecodedField) *pbgear.Staking_BondExtraCall {
	out := &pbgear.Staking_BondExtraCall{}
	out.MaxAdditional = To_Staking_CompactString(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_CancelDeferredSlashCall(fields []*registry.DecodedField) *pbgear.Staking_CancelDeferredSlashCall {
	out := &pbgear.Staking_CancelDeferredSlashCall{}
	out.Era = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.SlashIndices = To_repeated_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_ChillCall(fields []*registry.DecodedField) *pbgear.Staking_ChillCall {
	out := &pbgear.Staking_ChillCall{}
	return out
}

func To_Staking_ChillOtherCall(fields []*registry.DecodedField) *pbgear.Staking_ChillOtherCall {
	out := &pbgear.Staking_ChillOtherCall{}
	out.Controller = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "controller" optional: false repeated: false

func To_Staking_ChillThreshold(fields []*registry.DecodedField) *pbgear.Staking_ChillThreshold {
	out := &pbgear.Staking_ChillThreshold{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_ChillThreshold_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Staking_CompactSpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
	out := &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{}
	out.Value = To_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Staking_CompactString(fields []*registry.DecodedField) *pbgear.Staking_CompactString {
	out := &pbgear.Staking_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Staking_CompactTupleNull {
	out := &pbgear.Staking_CompactTupleNull{}
	out.Value = To_Staking_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Staking_CompactUint32(fields []*registry.DecodedField) *pbgear.Staking_CompactUint32 {
	out := &pbgear.Staking_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Controller(fields []*registry.DecodedField) *pbgear.Staking_Controller {
	out := &pbgear.Staking_Controller{}
	return out
}

func To_Staking_ForceApplyMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_ForceApplyMinCommissionCall {
	out := &pbgear.Staking_ForceApplyMinCommissionCall{}
	out.ValidatorStash = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_ForceNewEraAlwaysCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraAlwaysCall {
	out := &pbgear.Staking_ForceNewEraAlwaysCall{}
	return out
}

func To_Staking_ForceNewEraCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraCall {
	out := &pbgear.Staking_ForceNewEraCall{}
	return out
}

func To_Staking_ForceNoErasCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNoErasCall {
	out := &pbgear.Staking_ForceNoErasCall{}
	return out
}

func To_Staking_ForceUnstakeCall(fields []*registry.DecodedField) *pbgear.Staking_ForceUnstakeCall {
	out := &pbgear.Staking_ForceUnstakeCall{}
	out.Stash = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.NumSlashingSpans = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Id(fields []*registry.DecodedField) *pbgear.Staking_Id {
	out := &pbgear.Staking_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_IncreaseValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_IncreaseValidatorCountCall {
	out := &pbgear.Staking_IncreaseValidatorCountCall{}
	out.Additional = To_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "additional" optional: false repeated: false

func To_Staking_Index(fields []*registry.DecodedField) *pbgear.Staking_Index {
	out := &pbgear.Staking_Index{}
	out.Value0 = To_Staking_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Staking_KickCall(fields []*registry.DecodedField) *pbgear.Staking_KickCall {
	out := &pbgear.Staking_KickCall{}
	out.Who = To_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "who" optional: false repeated: true
// optional primitive: who
func To_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields []*registry.DecodedField) []*pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
	out := make([]*pbgear.Staking_SpRuntimeMultiaddressMultiAddress, len(fields))
	for _, f := range fields {
		out = append(out, To_Staking_SpRuntimeMultiaddressMultiAddress(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Staking_MaxNominatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
	out := &pbgear.Staking_MaxNominatorCount{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MaxNominatorCount_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MaxValidatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
	out := &pbgear.Staking_MaxValidatorCount{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MaxValidatorCount_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinCommission(fields []*registry.DecodedField) *pbgear.Staking_MinCommission {
	out := &pbgear.Staking_MinCommission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinCommission_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinNominatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinNominatorBond {
	out := &pbgear.Staking_MinNominatorBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinNominatorBond_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinValidatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinValidatorBond {
	out := &pbgear.Staking_MinValidatorBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinValidatorBond_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_NominateCall(fields []*registry.DecodedField) *pbgear.Staking_NominateCall {
	out := &pbgear.Staking_NominateCall{}
	out.Targets = To_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_None(fields []*registry.DecodedField) *pbgear.Staking_None {
	out := &pbgear.Staking_None{}
	return out
}

func To_Staking_Noop(fields []*registry.DecodedField) *pbgear.Staking_Noop {
	out := &pbgear.Staking_Noop{}
	return out
}

func To_Staking_PalletStakingValidatorPrefs(fields []*registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
	out := &pbgear.Staking_PalletStakingValidatorPrefs{}
	out.Commission = To_Staking_CompactSpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField))
	out.Blocked = To_bool(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "commission" optional: false repeated: false

func To_Staking_Payee(fields []*registry.DecodedField) *pbgear.Staking_Payee {
	out := &pbgear.Staking_Payee{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_Payee_Staked{
			Staked: To_Staking_Staked(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_PayoutStakersCall(fields []*registry.DecodedField) *pbgear.Staking_PayoutStakersCall {
	out := &pbgear.Staking_PayoutStakersCall{}
	out.ValidatorStash = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.Era = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Raw(fields []*registry.DecodedField) *pbgear.Staking_Raw {
	out := &pbgear.Staking_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_ReapStashCall(fields []*registry.DecodedField) *pbgear.Staking_ReapStashCall {
	out := &pbgear.Staking_ReapStashCall{}
	out.Stash = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	out.NumSlashingSpans = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_RebondCall(fields []*registry.DecodedField) *pbgear.Staking_RebondCall {
	out := &pbgear.Staking_RebondCall{}
	out.Value = To_Staking_CompactString(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_Remove(fields []*registry.DecodedField) *pbgear.Staking_Remove {
	out := &pbgear.Staking_Remove{}
	return out
}

func To_Staking_ScaleValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_ScaleValidatorCountCall {
	out := &pbgear.Staking_ScaleValidatorCountCall{}
	out.Factor = To_Staking_SpArithmeticPerThingsPercent(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "factor" optional: false repeated: false

func To_Staking_Set(fields []*registry.DecodedField) *pbgear.Staking_Set {
	out := &pbgear.Staking_Set{}
	out.Value0 = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SetControllerCall(fields []*registry.DecodedField) *pbgear.Staking_SetControllerCall {
	out := &pbgear.Staking_SetControllerCall{}
	return out
}

func To_Staking_SetInvulnerablesCall(fields []*registry.DecodedField) *pbgear.Staking_SetInvulnerablesCall {
	out := &pbgear.Staking_SetInvulnerablesCall{}
	out.Invulnerables = To_repeated_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SetMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_SetMinCommissionCall {
	out := &pbgear.Staking_SetMinCommissionCall{}
	out.New = To_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SetPayeeCall(fields []*registry.DecodedField) *pbgear.Staking_SetPayeeCall {
	out := &pbgear.Staking_SetPayeeCall{}
	out.Payee = To_Staking_Payee(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SetStakingConfigsCall(fields []*registry.DecodedField) *pbgear.Staking_SetStakingConfigsCall {
	out := &pbgear.Staking_SetStakingConfigsCall{}
	out.MinNominatorBond = To_Staking_MinNominatorBond(fields[0].Value.([]*registry.DecodedField))
	out.MinValidatorBond = To_Staking_MinValidatorBond(fields[1].Value.([]*registry.DecodedField))
	out.MaxNominatorCount = To_Staking_MaxNominatorCount(fields[2].Value.([]*registry.DecodedField))
	out.MaxValidatorCount = To_Staking_MaxValidatorCount(fields[3].Value.([]*registry.DecodedField))
	out.ChillThreshold = To_Staking_ChillThreshold(fields[4].Value.([]*registry.DecodedField))
	out.MinCommission = To_Staking_MinCommission(fields[5].Value.([]*registry.DecodedField))
	return out
}

// not seen "min_nominator_bond" optional: false repeated: false
// not seen "min_validator_bond" optional: false repeated: false
// not seen "max_nominator_count" optional: false repeated: false
// not seen "max_validator_count" optional: false repeated: false
// not seen "chill_threshold" optional: false repeated: false
// not seen "min_commission" optional: false repeated: false

func To_Staking_SetValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_SetValidatorCountCall {
	out := &pbgear.Staking_SetValidatorCountCall{}
	out.New = To_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
	out := &pbgear.Staking_SpArithmeticPerThingsPerbill{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SpArithmeticPerThingsPercent(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
	out := &pbgear.Staking_SpArithmeticPerThingsPercent{}
	out.Factor = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
	out := &pbgear.Staking_SpCoreCryptoAccountId32{}
	out.Invulnerables = To_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_SpRuntimeMultiaddressMultiAddress(fields []*registry.DecodedField) *pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
	out := &pbgear.Staking_SpRuntimeMultiaddressMultiAddress{}
	out.Targets = To_Staking_Targets(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "targets" optional: false repeated: false

func To_Staking_Staked(fields []*registry.DecodedField) *pbgear.Staking_Staked {
	out := &pbgear.Staking_Staked{}
	return out
}

func To_Staking_Stash(fields []*registry.DecodedField) *pbgear.Staking_Stash {
	out := &pbgear.Staking_Stash{}
	return out
}

func To_Staking_Targets(fields []*registry.DecodedField) *pbgear.Staking_Targets {
	out := &pbgear.Staking_Targets{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Staking_Targets_Id{
			Id: To_Staking_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Staking_Targets_Index{
			Index: To_Staking_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Staking_Targets_Raw{
			Raw: To_Staking_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Staking_Targets_Address32{
			Address32: To_Staking_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Staking_Targets_Address20{
			Address20: To_Staking_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_TupleNull(fields []*registry.DecodedField) *pbgear.Staking_TupleNull {
	out := &pbgear.Staking_TupleNull{}
	return out
}

func To_Staking_UnbondCall(fields []*registry.DecodedField) *pbgear.Staking_UnbondCall {
	out := &pbgear.Staking_UnbondCall{}
	out.Value = To_Staking_CompactString(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Staking_ValidateCall(fields []*registry.DecodedField) *pbgear.Staking_ValidateCall {
	out := &pbgear.Staking_ValidateCall{}
	out.Prefs = To_Staking_PalletStakingValidatorPrefs(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "prefs" optional: false repeated: false

func To_Staking_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.Staking_WithdrawUnbondedCall {
	out := &pbgear.Staking_WithdrawUnbondedCall{}
	out.NumSlashingSpans = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_KillPrefixCall(fields []*registry.DecodedField) *pbgear.System_KillPrefixCall {
	out := &pbgear.System_KillPrefixCall{}
	out.Prefix = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Subkeys = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_System_KillStorageCall(fields []*registry.DecodedField) *pbgear.System_KillStorageCall {
	out := &pbgear.System_KillStorageCall{}
	out.Keys = To_repeated_System_SystemKeysList(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "keys" optional: false repeated: true
// optional primitive: keys
func To_repeated_System_SystemKeysList(fields []*registry.DecodedField) []*pbgear.System_SystemKeysList {
	out := make([]*pbgear.System_SystemKeysList, len(fields))
	for _, f := range fields {
		out = append(out, To_System_SystemKeysList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_System_RemarkCall(fields []*registry.DecodedField) *pbgear.System_RemarkCall {
	out := &pbgear.System_RemarkCall{}
	out.Remark = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_RemarkWithEventCall(fields []*registry.DecodedField) *pbgear.System_RemarkWithEventCall {
	out := &pbgear.System_RemarkWithEventCall{}
	out.Remark = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_SetCodeCall(fields []*registry.DecodedField) *pbgear.System_SetCodeCall {
	out := &pbgear.System_SetCodeCall{}
	out.Code = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_SetCodeWithoutChecksCall(fields []*registry.DecodedField) *pbgear.System_SetCodeWithoutChecksCall {
	out := &pbgear.System_SetCodeWithoutChecksCall{}
	out.Code = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_SetHeapPagesCall(fields []*registry.DecodedField) *pbgear.System_SetHeapPagesCall {
	out := &pbgear.System_SetHeapPagesCall{}
	out.Pages = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_SetStorageCall(fields []*registry.DecodedField) *pbgear.System_SetStorageCall {
	out := &pbgear.System_SetStorageCall{}
	out.Items = To_repeated_System_TupleSystemItemsListSystemItemsList(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "items" optional: false repeated: true
// optional primitive: items
func To_repeated_System_TupleSystemItemsListSystemItemsList(fields []*registry.DecodedField) []*pbgear.System_TupleSystemItemsListSystemItemsList {
	out := make([]*pbgear.System_TupleSystemItemsListSystemItemsList, len(fields))
	for _, f := range fields {
		out = append(out, To_System_TupleSystemItemsListSystemItemsList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_System_SystemKeysList(fields []*registry.DecodedField) *pbgear.System_SystemKeysList {
	out := &pbgear.System_SystemKeysList{}
	out.Keys = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_System_TupleSystemItemsListSystemItemsList(fields []*registry.DecodedField) *pbgear.System_TupleSystemItemsListSystemItemsList {
	out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_bytes(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Timestamp_CompactUint64(fields []*registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
	out := &pbgear.Timestamp_CompactUint64{}
	out.Value = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Timestamp_SetCall(fields []*registry.DecodedField) *pbgear.Timestamp_SetCall {
	out := &pbgear.Timestamp_SetCall{}
	out.Now = To_Timestamp_CompactUint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "now" optional: false repeated: false

func To_Treasury_Address20(fields []*registry.DecodedField) *pbgear.Treasury_Address20 {
	out := &pbgear.Treasury_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_Address32(fields []*registry.DecodedField) *pbgear.Treasury_Address32 {
	out := &pbgear.Treasury_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_ApproveProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_ApproveProposalCall {
	out := &pbgear.Treasury_ApproveProposalCall{}
	out.ProposalId = To_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "proposal_id" optional: false repeated: false

func To_Treasury_Beneficiary(fields []*registry.DecodedField) *pbgear.Treasury_Beneficiary {
	out := &pbgear.Treasury_Beneficiary{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Treasury_Beneficiary_Id{
			Id: To_Treasury_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Treasury_Beneficiary_Index{
			Index: To_Treasury_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Treasury_Beneficiary_Raw{
			Raw: To_Treasury_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Treasury_Beneficiary_Address32{
			Address32: To_Treasury_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Treasury_Beneficiary_Address20{
			Address20: To_Treasury_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Treasury_CheckStatusCall(fields []*registry.DecodedField) *pbgear.Treasury_CheckStatusCall {
	out := &pbgear.Treasury_CheckStatusCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_CompactString(fields []*registry.DecodedField) *pbgear.Treasury_CompactString {
	out := &pbgear.Treasury_CompactString{}
	out.Value = To_string(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
	out := &pbgear.Treasury_CompactTupleNull{}
	out.Value = To_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Treasury_CompactUint32(fields []*registry.DecodedField) *pbgear.Treasury_CompactUint32 {
	out := &pbgear.Treasury_CompactUint32{}
	out.Value = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_Id(fields []*registry.DecodedField) *pbgear.Treasury_Id {
	out := &pbgear.Treasury_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_Index(fields []*registry.DecodedField) *pbgear.Treasury_Index {
	out := &pbgear.Treasury_Index{}
	out.Value0 = To_Treasury_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Treasury_PayoutCall(fields []*registry.DecodedField) *pbgear.Treasury_PayoutCall {
	out := &pbgear.Treasury_PayoutCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_ProposeSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_ProposeSpendCall {
	out := &pbgear.Treasury_ProposeSpendCall{}
	out.Value = To_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Beneficiary = To_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false
// not seen "beneficiary" optional: false repeated: false

func To_Treasury_Raw(fields []*registry.DecodedField) *pbgear.Treasury_Raw {
	out := &pbgear.Treasury_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_RejectProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_RejectProposalCall {
	out := &pbgear.Treasury_RejectProposalCall{}
	out.ProposalId = To_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_RemoveApprovalCall(fields []*registry.DecodedField) *pbgear.Treasury_RemoveApprovalCall {
	out := &pbgear.Treasury_RemoveApprovalCall{}
	out.ProposalId = To_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
	out := &pbgear.Treasury_SpCoreCryptoAccountId32{}
	out.Beneficiary = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_SpendCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendCall {
	out := &pbgear.Treasury_SpendCall{}
	out.AssetKind = To_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField))
	out.Amount = To_Treasury_CompactString(fields[1].Value.([]*registry.DecodedField))
	out.Beneficiary = To_Treasury_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField))
	out.ValidFrom = To_optional_uint32(fields[3].Value.([]*registry.DecodedField))
	return out
}

// not seen "beneficiary" optional: false repeated: false

func To_Treasury_SpendLocalCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendLocalCall {
	out := &pbgear.Treasury_SpendLocalCall{}
	out.Amount = To_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField))
	out.Beneficiary = To_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Treasury_TupleNull(fields []*registry.DecodedField) *pbgear.Treasury_TupleNull {
	out := &pbgear.Treasury_TupleNull{}
	return out
}

func To_Treasury_VoidSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_VoidSpendCall {
	out := &pbgear.Treasury_VoidSpendCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	out := &pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
	out.Value0 = To_FinalityGrandpaPrevote(fields[0].Value.([]*registry.DecodedField))
	out.Value1 = To_SpConsensusGrandpaAppSignature(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false
// not seen "value1" optional: false repeated: false

func To_Utility_AsDerivativeCall(fields []*registry.DecodedField) *pbgear.Utility_AsDerivativeCall {
	out := &pbgear.Utility_AsDerivativeCall{}
	out.Index = To_uint32(fields[0].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Utility_AsDerivativeCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Utility_AsDerivativeCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Utility_AsDerivativeCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Utility_AsDerivativeCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Utility_AsDerivativeCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Utility_AsDerivativeCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Utility_AsDerivativeCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Utility_AsDerivativeCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Utility_AsDerivativeCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Utility_AsDerivativeCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Utility_AsDerivativeCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "call" optional: false repeated: false

func To_Utility_AsOrigin(fields []*registry.DecodedField) *pbgear.Utility_AsOrigin {
	out := &pbgear.Utility_AsOrigin{}

	switch {
	case matchFields(fields, []int64{121}):
		out.Value = &pbgear.Utility_AsOrigin_System{
			System: To_Utility_System(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{122}):
		out.Value = &pbgear.Utility_AsOrigin_Origins{
			Origins: To_Utility_Origins(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{123}):
		out.Value = &pbgear.Utility_AsOrigin_Void{
			Void: To_Utility_Void(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Utility_BatchAllCall(fields []*registry.DecodedField) *pbgear.Utility_BatchAllCall {
	out := &pbgear.Utility_BatchAllCall{}
	out.Calls = To_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "calls" optional: false repeated: true
// optional primitive: calls
func To_repeated_Utility_VaraRuntimeRuntimeCall(fields []*registry.DecodedField) []*pbgear.Utility_VaraRuntimeRuntimeCall {
	out := make([]*pbgear.Utility_VaraRuntimeRuntimeCall, len(fields))
	for _, f := range fields {
		out = append(out, To_Utility_VaraRuntimeRuntimeCall(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Utility_BatchCall(fields []*registry.DecodedField) *pbgear.Utility_BatchCall {
	out := &pbgear.Utility_BatchCall{}
	out.Calls = To_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_CompactUint64(fields []*registry.DecodedField) *pbgear.Utility_CompactUint64 {
	out := &pbgear.Utility_CompactUint64{}
	out.Value = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_DispatchAsCall(fields []*registry.DecodedField) *pbgear.Utility_DispatchAsCall {
	out := &pbgear.Utility_DispatchAsCall{}
	out.AsOrigin = To_Utility_AsOrigin(fields[0].Value.([]*registry.DecodedField))

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Utility_DispatchAsCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Utility_DispatchAsCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Utility_DispatchAsCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Utility_DispatchAsCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Utility_DispatchAsCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Utility_DispatchAsCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Utility_DispatchAsCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Utility_DispatchAsCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Utility_DispatchAsCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Utility_DispatchAsCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Utility_DispatchAsCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Utility_DispatchAsCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Utility_DispatchAsCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Utility_DispatchAsCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Utility_DispatchAsCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Utility_DispatchAsCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "as_origin" optional: false repeated: false

func To_Utility_ForceBatchCall(fields []*registry.DecodedField) *pbgear.Utility_ForceBatchCall {
	out := &pbgear.Utility_ForceBatchCall{}
	out.Calls = To_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_Origins(fields []*registry.DecodedField) *pbgear.Utility_Origins {
	out := &pbgear.Utility_Origins{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
	out := &pbgear.Utility_SpWeightsWeightV2Weight{}
	out.RefTime = To_Utility_CompactUint64(fields[0].Value.([]*registry.DecodedField))
	out.ProofSize = To_Utility_CompactUint64(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "ref_time" optional: false repeated: false

func To_Utility_System(fields []*registry.DecodedField) *pbgear.Utility_System {
	out := &pbgear.Utility_System{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_VaraRuntimeRuntimeCall(fields []*registry.DecodedField) *pbgear.Utility_VaraRuntimeRuntimeCall {
	out := &pbgear.Utility_VaraRuntimeRuntimeCall{}

	switch {
	case matchFields(fields, []int64{66}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Calls = &pbgear.Utility_VaraRuntimeRuntimeCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "calls" optional: false repeated: false

func To_Utility_Void(fields []*registry.DecodedField) *pbgear.Utility_Void {
	out := &pbgear.Utility_Void{}
	out.Value0 = To_Value0(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Utility_WithWeightCall(fields []*registry.DecodedField) *pbgear.Utility_WithWeightCall {
	out := &pbgear.Utility_WithWeightCall{}

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Utility_WithWeightCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Utility_WithWeightCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Utility_WithWeightCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Utility_WithWeightCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Utility_WithWeightCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Utility_WithWeightCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Utility_WithWeightCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Utility_WithWeightCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Utility_WithWeightCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Utility_WithWeightCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Utility_WithWeightCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Utility_WithWeightCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Utility_WithWeightCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Utility_WithWeightCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Utility_WithWeightCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Utility_WithWeightCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	out.Weight = To_Utility_SpWeightsWeightV2Weight(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "weight" optional: false repeated: false

func To_Value0(fields []*registry.DecodedField) *pbgear.Value0 {
	out := &pbgear.Value0{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Value0_Root{
			Root: To_Root(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Vesting_Address20(fields []*registry.DecodedField) *pbgear.Vesting_Address20 {
	out := &pbgear.Vesting_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_Address32(fields []*registry.DecodedField) *pbgear.Vesting_Address32 {
	out := &pbgear.Vesting_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
	out := &pbgear.Vesting_CompactTupleNull{}
	out.Value = To_Vesting_TupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value" optional: false repeated: false

func To_Vesting_ForceVestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_ForceVestedTransferCall {
	out := &pbgear.Vesting_ForceVestedTransferCall{}
	out.Source = To_Vesting_Source(fields[0].Value.([]*registry.DecodedField))
	out.Target = To_Vesting_Target(fields[1].Value.([]*registry.DecodedField))
	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "source" optional: false repeated: false
// not seen "target" optional: false repeated: false
// not seen "schedule" optional: false repeated: false

func To_Vesting_Id(fields []*registry.DecodedField) *pbgear.Vesting_Id {
	out := &pbgear.Vesting_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_Index(fields []*registry.DecodedField) *pbgear.Vesting_Index {
	out := &pbgear.Vesting_Index{}
	out.Value0 = To_Vesting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField))
	return out
}

// not seen "value0" optional: false repeated: false

func To_Vesting_MergeSchedulesCall(fields []*registry.DecodedField) *pbgear.Vesting_MergeSchedulesCall {
	out := &pbgear.Vesting_MergeSchedulesCall{}
	out.Schedule1Index = To_uint32(fields[0].Value.([]*registry.DecodedField))
	out.Schedule2Index = To_uint32(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_PalletVestingVestingInfoVestingInfo(fields []*registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
	out.Locked = To_string(fields[0].Value.([]*registry.DecodedField))
	out.PerBlock = To_string(fields[1].Value.([]*registry.DecodedField))
	out.StartingBlock = To_uint32(fields[2].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_Raw(fields []*registry.DecodedField) *pbgear.Vesting_Raw {
	out := &pbgear.Vesting_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_Source(fields []*registry.DecodedField) *pbgear.Vesting_Source {
	out := &pbgear.Vesting_Source{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Vesting_Source_Id{
			Id: To_Vesting_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Vesting_Source_Index{
			Index: To_Vesting_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Vesting_Source_Raw{
			Raw: To_Vesting_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Vesting_Source_Address32{
			Address32: To_Vesting_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Vesting_Source_Address20{
			Address20: To_Vesting_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "value" optional: false repeated: false

func To_Vesting_Target(fields []*registry.DecodedField) *pbgear.Vesting_Target {
	out := &pbgear.Vesting_Target{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Vesting_Target_Id{
			Id: To_Vesting_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Vesting_Target_Index{
			Index: To_Vesting_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Vesting_Target_Raw{
			Raw: To_Vesting_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Vesting_Target_Address32{
			Address32: To_Vesting_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Vesting_Target_Address20{
			Address20: To_Vesting_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Vesting_TupleNull(fields []*registry.DecodedField) *pbgear.Vesting_TupleNull {
	out := &pbgear.Vesting_TupleNull{}
	return out
}

func To_Vesting_VestCall(fields []*registry.DecodedField) *pbgear.Vesting_VestCall {
	out := &pbgear.Vesting_VestCall{}
	return out
}

func To_Vesting_VestOtherCall(fields []*registry.DecodedField) *pbgear.Vesting_VestOtherCall {
	out := &pbgear.Vesting_VestOtherCall{}
	out.Target = To_Vesting_Target(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Vesting_VestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_VestedTransferCall {
	out := &pbgear.Vesting_VestedTransferCall{}
	out.Target = To_Vesting_Target(fields[0].Value.([]*registry.DecodedField))
	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(fields[1].Value.([]*registry.DecodedField))
	return out
}

func To_Whitelist_CompactUint64(fields []*registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
	out := &pbgear.Whitelist_CompactUint64{}
	out.Value = To_uint64(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Whitelist_DispatchWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
	out.CallHash = To_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	out.CallEncodedLen = To_uint32(fields[1].Value.([]*registry.DecodedField))
	out.CallWeightWitness = To_Whitelist_SpWeightsWeightV2Weight(fields[2].Value.([]*registry.DecodedField))
	return out
}

// not seen "call_hash" optional: false repeated: false
// not seen "call_weight_witness" optional: false repeated: false

func To_Whitelist_DispatchWhitelistedCallWithPreimageCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{}

	switch {
	case matchFields(fields, []int64{66}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Call = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

// not seen "call" optional: false repeated: false

func To_Whitelist_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
	out := &pbgear.Whitelist_PrimitiveTypesH256{}
	out.CallHash = To_bytes(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Whitelist_RemoveWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_RemoveWhitelistedCallCall {
	out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
	out.CallHash = To_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_Whitelist_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
	out := &pbgear.Whitelist_SpWeightsWeightV2Weight{}
	out.RefTime = To_Whitelist_CompactUint64(fields[0].Value.([]*registry.DecodedField))
	out.ProofSize = To_Whitelist_CompactUint64(fields[1].Value.([]*registry.DecodedField))
	return out
}

// not seen "ref_time" optional: false repeated: false

func To_Whitelist_WhitelistCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_WhitelistCallCall {
	out := &pbgear.Whitelist_WhitelistCallCall{}
	out.CallHash = To_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField))
	return out
}

func To_conviction_voting_PalletConvictionVotingVoteVote(fields []*registry.DecodedField) *pbgear.conviction_voting_PalletConvictionVotingVoteVote {
	out := &pbgear.conviction_voting_PalletConvictionVotingVoteVote{}
	out.Vote = To_uint32(fields[0].Value.([]*registry.DecodedField))
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

func To_bytes(fields []*registry.DecodedField) []byte {
	var out []byte

	data := fields[0].Value.([]interface{})
	for _, f := range data {
		u := uint8(f.(types.U8))
		out = append(out, byte(u))
	}
	return out
}

func To_string(fields []*registry.DecodedField) string {
	return string(To_bytes(fields))
}

func To_uint32(fields []*registry.DecodedField) uint32 {
	return fields[0].Value.(uint32)
}

func To_uint64(fields []*registry.DecodedField) uint64 {
	return fields[0].Value.(uint64)
}

func To_bool(fields []*registry.DecodedField) bool {
	return fields[0].Value.(bool)
}

func To_optional_string(fields []*registry.DecodedField) *string {
	return string(To_bytes(fields))
}

func To_optional_uint32(fields []*registry.DecodedField) *uint32 {
	return fields[0].Value.(*uint32)
}

func To_optional_uint64(fields []*registry.DecodedField) *uint64 {
	return fields[0].Value.(*uint64)
}

func To_optional_bool(fields []*registry.DecodedField) *bool {
	return fields[0].Value.(*bool)
}
