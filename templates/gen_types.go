package gen_types

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func to_Babe_AllowedSlots(fields []*registry.DecodedField) *pbgear.Babe_AllowedSlots {
	return &pbgear.Babe_AllowedSlots{
		Value: to_oneof_Babe_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Babe_Value(in *registry.DecodedField) *pbgear.Babe_Value {
	return nil
}

func to_Babe_BabeTrieNodesList(fields []*registry.DecodedField) *pbgear.Babe_BabeTrieNodesList {
	return &pbgear.Babe_BabeTrieNodesList{
		TrieNodes: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_repeated_uint32(in []*registry.DecodedField) []uint32 {
	out := make([]uint32, len(in))
	for i := range in {
		out = append(out, i.Value.(uint32))
	}
	return out
}

func to_Babe_CompactUint32(fields []*registry.DecodedField) *pbgear.Babe_CompactUint32 {
	return &pbgear.Babe_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Babe_Config(fields []*registry.DecodedField) *pbgear.Babe_Config {
	return &pbgear.Babe_Config{
		Value: to_oneof_Babe_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_Consensus(fields []*registry.DecodedField) *pbgear.Babe_Consensus {
	return &pbgear.Babe_Consensus{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_Logs(fields []*registry.DecodedField) *pbgear.Babe_Logs {
	return &pbgear.Babe_Logs{
		Value: to_oneof_Babe_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_Other(fields []*registry.DecodedField) *pbgear.Babe_Other {
	return &pbgear.Babe_Other{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_PlanConfigChangeCall(fields []*registry.DecodedField) *pbgear.Babe_PlanConfigChangeCall {
	return &pbgear.Babe_PlanConfigChangeCall{
		Config: to_Babe_Config(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_PreRuntime(fields []*registry.DecodedField) *pbgear.Babe_PreRuntime {
	return &pbgear.Babe_PreRuntime{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_PrimaryAndSecondaryPlainSlots(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
	return &pbgear.Babe_PrimaryAndSecondaryPlainSlots{}
}
func to_Babe_PrimaryAndSecondaryVRFSlots(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryVRFSlots {
	return &pbgear.Babe_PrimaryAndSecondaryVRFSlots{}
}
func to_Babe_PrimarySlots(fields []*registry.DecodedField) *pbgear.Babe_PrimarySlots {
	return &pbgear.Babe_PrimarySlots{}
}
func to_Babe_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
	return &pbgear.Babe_PrimitiveTypesH256{
		ParentHash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationCall {
	return &pbgear.Babe_ReportEquivocationCall{
		EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
		KeyOwnerProof:     to_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationUnsignedCall {
	return &pbgear.Babe_ReportEquivocationUnsignedCall{
		EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
		KeyOwnerProof:     to_Babe_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_RuntimeEnvironmentUpdated(fields []*registry.DecodedField) *pbgear.Babe_RuntimeEnvironmentUpdated {
	return &pbgear.Babe_RuntimeEnvironmentUpdated{}
}
func to_Babe_Seal(fields []*registry.DecodedField) *pbgear.Babe_Seal {
	return &pbgear.Babe_Seal{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
	return &pbgear.Babe_SpConsensusBabeAppPublic{
		Offender: to_Babe_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpConsensusSlotsEquivocationProof(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
	return &pbgear.Babe_SpConsensusSlotsEquivocationProof{
		Offender:     to_Babe_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField)),
		Slot:         to_Babe_SpConsensusSlotsSlot(fields[1].Value.([]*registry.DecodedField)),
		FirstHeader:  to_Babe_SpRuntimeGenericHeaderHeader(fields[2].Value.([]*registry.DecodedField)),
		SecondHeader: to_Babe_SpRuntimeGenericHeaderHeader(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpConsensusSlotsSlot(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
	return &pbgear.Babe_SpConsensusSlotsSlot{
		Slot: fields[0].Value.(uint64),
	}
}

func to_Babe_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
	return &pbgear.Babe_SpCoreSr25519Public{
		Offender: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpRuntimeGenericDigestDigest(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
	return &pbgear.Babe_SpRuntimeGenericDigestDigest{
		Logs: to_repeated_Babe_SpRuntimeGenericDigestDigestItem(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpRuntimeGenericDigestDigestItem(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigestItem {
	return &pbgear.Babe_SpRuntimeGenericDigestDigestItem{
		Logs: to_Babe_Logs(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpRuntimeGenericHeaderHeader(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
	return &pbgear.Babe_SpRuntimeGenericHeaderHeader{
		ParentHash:     to_Babe_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
		Number:         to_Babe_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
		StateRoot:      to_Babe_PrimitiveTypesH256(fields[2].Value.([]*registry.DecodedField)),
		ExtrinsicsRoot: to_Babe_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField)),
		Digest:         to_Babe_SpRuntimeGenericDigestDigest(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_Babe_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
	return &pbgear.Babe_SpSessionMembershipProof{
		Session:        fields[0].Value.(uint32),
		TrieNodes:      to_repeated_Babe_BabeTrieNodesList(fields[1].Value.([]*registry.DecodedField)),
		ValidatorCount: fields[2].Value.(uint32),
	}
}

func to_Babe_TupleUint64Uint64(fields []*registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
	return &pbgear.Babe_TupleUint64Uint64{
		Value_0: fields[0].Value.(uint64),
		Value_1: fields[1].Value.(uint64),
	}
}

func to_Babe_V1(fields []*registry.DecodedField) *pbgear.Babe_V1 {
	return &pbgear.Babe_V1{
		C:            to_Babe_TupleUint64Uint64(fields[0].Value.([]*registry.DecodedField)),
		AllowedSlots: to_Babe_AllowedSlots(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Address20(fields []*registry.DecodedField) *pbgear.BagsList_Address20 {
	return &pbgear.BagsList_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Address32(fields []*registry.DecodedField) *pbgear.BagsList_Address32 {
	return &pbgear.BagsList_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_CompactTupleNull(fields []*registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
	return &pbgear.BagsList_CompactTupleNull{
		Value: to_BagsList_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Dislocated(fields []*registry.DecodedField) *pbgear.BagsList_Dislocated {
	return &pbgear.BagsList_Dislocated{
		Value: to_oneof_BagsList_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_BagsList_Value(in *registry.DecodedField) *pbgear.BagsList_Value {
	return nil
}

func to_BagsList_Heavier(fields []*registry.DecodedField) *pbgear.BagsList_Heavier {
	return &pbgear.BagsList_Heavier{
		Value: to_oneof_BagsList_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Id(fields []*registry.DecodedField) *pbgear.BagsList_Id {
	return &pbgear.BagsList_Id{
		Value_0: to_BagsList_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Index(fields []*registry.DecodedField) *pbgear.BagsList_Index {
	return &pbgear.BagsList_Index{
		Value_0: to_BagsList_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Lighter(fields []*registry.DecodedField) *pbgear.BagsList_Lighter {
	return &pbgear.BagsList_Lighter{
		Value: to_oneof_BagsList_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_PutInFrontOfCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfCall {
	return &pbgear.BagsList_PutInFrontOfCall{
		Lighter: to_BagsList_Lighter(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_PutInFrontOfOtherCall(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfOtherCall {
	return &pbgear.BagsList_PutInFrontOfOtherCall{
		Heavier: to_BagsList_Heavier(fields[0].Value.([]*registry.DecodedField)),
		Lighter: to_BagsList_Lighter(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_Raw(fields []*registry.DecodedField) *pbgear.BagsList_Raw {
	return &pbgear.BagsList_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_RebagCall(fields []*registry.DecodedField) *pbgear.BagsList_RebagCall {
	return &pbgear.BagsList_RebagCall{
		Dislocated: to_BagsList_Dislocated(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.BagsList_SpCoreCryptoAccountId32 {
	return &pbgear.BagsList_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_BagsList_TupleNull(fields []*registry.DecodedField) *pbgear.BagsList_TupleNull {
	return &pbgear.BagsList_TupleNull{}
}
func to_Balances_Address20(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
	return &pbgear.Balances_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Address32(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
	return &pbgear.Balances_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_CompactString(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
	return &pbgear.Balances_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_Balances_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
	return &pbgear.Balances_CompactTupleNull{
		Value: to_Balances_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Dest(fields []*registry.DecodedField) *pbgear.Balances_Dest {
	return &pbgear.Balances_Dest{
		Value: to_oneof_Balances_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Balances_Value(in *registry.DecodedField) *pbgear.Balances_Value {
	return nil
}

func to_Balances_ForceSetBalanceCall(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
	return &pbgear.Balances_ForceSetBalanceCall{
		Who:     to_Balances_Who(fields[0].Value.([]*registry.DecodedField)),
		NewFree: to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_ForceTransferCall(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
	return &pbgear.Balances_ForceTransferCall{
		Source: to_Balances_Source(fields[0].Value.([]*registry.DecodedField)),
		Dest:   to_Balances_Dest(fields[1].Value.([]*registry.DecodedField)),
		Value:  to_Balances_CompactString(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_ForceUnreserveCall(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
	return &pbgear.Balances_ForceUnreserveCall{
		Who:    to_Balances_Who(fields[0].Value.([]*registry.DecodedField)),
		Amount: fields[1].Value.(string),
	}
}

func to_Balances_Id(fields []*registry.DecodedField) *pbgear.Balances_Id {
	return &pbgear.Balances_Id{
		Value_0: to_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Index(fields []*registry.DecodedField) *pbgear.Balances_Index {
	return &pbgear.Balances_Index{
		Value_0: to_Balances_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Raw(fields []*registry.DecodedField) *pbgear.Balances_Raw {
	return &pbgear.Balances_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Source(fields []*registry.DecodedField) *pbgear.Balances_Source {
	return &pbgear.Balances_Source{
		Value: to_oneof_Balances_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
	return &pbgear.Balances_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_TransferAllCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
	return &pbgear.Balances_TransferAllCall{
		Dest:      to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
		KeepAlive: fields[1].Value.(bool),
	}
}

func to_Balances_TransferAllowDeathCall(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
	return &pbgear.Balances_TransferAllowDeathCall{
		Dest:  to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
		Value: to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_TransferKeepAliveCall(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
	return &pbgear.Balances_TransferKeepAliveCall{
		Dest:  to_Balances_Dest(fields[0].Value.([]*registry.DecodedField)),
		Value: to_Balances_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_TupleNull(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
	return &pbgear.Balances_TupleNull{}
}
func to_Balances_UpgradeAccountsCall(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
	return &pbgear.Balances_UpgradeAccountsCall{
		Who: to_repeated_Balances_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Balances_Who(fields []*registry.DecodedField) *pbgear.Balances_Who {
	return &pbgear.Balances_Who{
		Value: to_oneof_Balances_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_AcceptCuratorCall {
	return &pbgear.Bounties_AcceptCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Address20(fields []*registry.DecodedField) *pbgear.Bounties_Address20 {
	return &pbgear.Bounties_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Address32(fields []*registry.DecodedField) *pbgear.Bounties_Address32 {
	return &pbgear.Bounties_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_ApproveBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ApproveBountyCall {
	return &pbgear.Bounties_ApproveBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_AwardBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_AwardBountyCall {
	return &pbgear.Bounties_AwardBountyCall{
		BountyId:    to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Beneficiary: to_Bounties_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Beneficiary(fields []*registry.DecodedField) *pbgear.Bounties_Beneficiary {
	return &pbgear.Bounties_Beneficiary{
		Value: to_oneof_Bounties_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Bounties_Value(in *registry.DecodedField) *pbgear.Bounties_Value {
	return nil
}

func to_Bounties_ClaimBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ClaimBountyCall {
	return &pbgear.Bounties_ClaimBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_CloseBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_CloseBountyCall {
	return &pbgear.Bounties_CloseBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_CompactString(fields []*registry.DecodedField) *pbgear.Bounties_CompactString {
	return &pbgear.Bounties_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_Bounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
	return &pbgear.Bounties_CompactTupleNull{
		Value: to_Bounties_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_CompactUint32(fields []*registry.DecodedField) *pbgear.Bounties_CompactUint32 {
	return &pbgear.Bounties_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Bounties_Curator(fields []*registry.DecodedField) *pbgear.Bounties_Curator {
	return &pbgear.Bounties_Curator{
		Value: to_oneof_Bounties_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_ExtendBountyExpiryCall(fields []*registry.DecodedField) *pbgear.Bounties_ExtendBountyExpiryCall {
	return &pbgear.Bounties_ExtendBountyExpiryCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Remark:   to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Id(fields []*registry.DecodedField) *pbgear.Bounties_Id {
	return &pbgear.Bounties_Id{
		Value_0: to_Bounties_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Index(fields []*registry.DecodedField) *pbgear.Bounties_Index {
	return &pbgear.Bounties_Index{
		Value_0: to_Bounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_ProposeBountyCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeBountyCall {
	return &pbgear.Bounties_ProposeBountyCall{
		Value:       to_Bounties_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Description: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_ProposeCuratorCall {
	return &pbgear.Bounties_ProposeCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Curator:  to_Bounties_Curator(fields[1].Value.([]*registry.DecodedField)),
		Fee:      to_Bounties_CompactString(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_Raw(fields []*registry.DecodedField) *pbgear.Bounties_Raw {
	return &pbgear.Bounties_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Bounties_SpCoreCryptoAccountId32 {
	return &pbgear.Bounties_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Bounties_TupleNull(fields []*registry.DecodedField) *pbgear.Bounties_TupleNull {
	return &pbgear.Bounties_TupleNull{}
}
func to_Bounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.Bounties_UnassignCuratorCall {
	return &pbgear.Bounties_UnassignCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_AcceptCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AcceptCuratorCall {
	return &pbgear.ChildBounties_AcceptCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_AddChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AddChildBountyCall {
	return &pbgear.ChildBounties_AddChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value:          to_ChildBounties_CompactString(fields[1].Value.([]*registry.DecodedField)),
		Description:    to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Address20(fields []*registry.DecodedField) *pbgear.ChildBounties_Address20 {
	return &pbgear.ChildBounties_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Address32(fields []*registry.DecodedField) *pbgear.ChildBounties_Address32 {
	return &pbgear.ChildBounties_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_AwardChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_AwardChildBountyCall {
	return &pbgear.ChildBounties_AwardChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
		Beneficiary:    to_ChildBounties_Beneficiary(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Beneficiary(fields []*registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
	return &pbgear.ChildBounties_Beneficiary{
		Value: to_oneof_ChildBounties_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_ChildBounties_Value(in *registry.DecodedField) *pbgear.ChildBounties_Value {
	return nil
}

func to_ChildBounties_ClaimChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ClaimChildBountyCall {
	return &pbgear.ChildBounties_ClaimChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_CloseChildBountyCall(fields []*registry.DecodedField) *pbgear.ChildBounties_CloseChildBountyCall {
	return &pbgear.ChildBounties_CloseChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_CompactString(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactString {
	return &pbgear.ChildBounties_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_ChildBounties_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
	return &pbgear.ChildBounties_CompactTupleNull{
		Value: to_ChildBounties_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_CompactUint32(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
	return &pbgear.ChildBounties_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_ChildBounties_Curator(fields []*registry.DecodedField) *pbgear.ChildBounties_Curator {
	return &pbgear.ChildBounties_Curator{
		Value: to_oneof_ChildBounties_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Id(fields []*registry.DecodedField) *pbgear.ChildBounties_Id {
	return &pbgear.ChildBounties_Id{
		Value_0: to_ChildBounties_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Index(fields []*registry.DecodedField) *pbgear.ChildBounties_Index {
	return &pbgear.ChildBounties_Index{
		Value_0: to_ChildBounties_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_ProposeCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_ProposeCuratorCall {
	return &pbgear.ChildBounties_ProposeCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
		Curator:        to_ChildBounties_Curator(fields[2].Value.([]*registry.DecodedField)),
		Fee:            to_ChildBounties_CompactString(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_Raw(fields []*registry.DecodedField) *pbgear.ChildBounties_Raw {
	return &pbgear.ChildBounties_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ChildBounties_SpCoreCryptoAccountId32 {
	return &pbgear.ChildBounties_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ChildBounties_TupleNull(fields []*registry.DecodedField) *pbgear.ChildBounties_TupleNull {
	return &pbgear.ChildBounties_TupleNull{}
}
func to_ChildBounties_UnassignCuratorCall(fields []*registry.DecodedField) *pbgear.ChildBounties_UnassignCuratorCall {
	return &pbgear.ChildBounties_UnassignCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Address20(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address20 {
	return &pbgear.ConvictionVoting_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Address32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address32 {
	return &pbgear.ConvictionVoting_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
	return &pbgear.ConvictionVoting_CompactTupleNull{
		Value: to_ConvictionVoting_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_CompactUint32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
	return &pbgear.ConvictionVoting_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_ConvictionVoting_Conviction(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
	return &pbgear.ConvictionVoting_Conviction{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_ConvictionVoting_Value(in *registry.DecodedField) *pbgear.ConvictionVoting_Value {
	return nil
}

func to_ConvictionVoting_DelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_DelegateCall {
	return &pbgear.ConvictionVoting_DelegateCall{
		Class:      fields[0].Value.(uint32),
		To:         to_ConvictionVoting_To(fields[1].Value.([]*registry.DecodedField)),
		Conviction: to_ConvictionVoting_Conviction(fields[2].Value.([]*registry.DecodedField)),
		Balance:    fields[3].Value.(string),
	}
}

func to_ConvictionVoting_Id(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Id {
	return &pbgear.ConvictionVoting_Id{
		Value_0: to_ConvictionVoting_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Index(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Index {
	return &pbgear.ConvictionVoting_Index{
		Value_0: to_ConvictionVoting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Locked1X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked1X {
	return &pbgear.ConvictionVoting_Locked1X{}
}
func to_ConvictionVoting_Locked2X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked2X {
	return &pbgear.ConvictionVoting_Locked2X{}
}
func to_ConvictionVoting_Locked3X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked3X {
	return &pbgear.ConvictionVoting_Locked3X{}
}
func to_ConvictionVoting_Locked4X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked4X {
	return &pbgear.ConvictionVoting_Locked4X{}
}
func to_ConvictionVoting_Locked5X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked5X {
	return &pbgear.ConvictionVoting_Locked5X{}
}
func to_ConvictionVoting_Locked6X(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked6X {
	return &pbgear.ConvictionVoting_Locked6X{}
}
func to_ConvictionVoting_None(fields []*registry.DecodedField) *pbgear.ConvictionVoting_None {
	return &pbgear.ConvictionVoting_None{}
}
func to_ConvictionVoting_PalletConvictionVotingVoteVote(fields []*registry.DecodedField) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
	return &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{
		Vote: fields[0].Value.(uint32),
	}
}

func to_ConvictionVoting_Raw(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Raw {
	return &pbgear.ConvictionVoting_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_RemoveOtherVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
	return &pbgear.ConvictionVoting_RemoveOtherVoteCall{
		Target: to_ConvictionVoting_Target(fields[0].Value.([]*registry.DecodedField)),
		Class:  fields[1].Value.(uint32),
		Index:  fields[2].Value.(uint32),
	}
}

func to_ConvictionVoting_RemoveVoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveVoteCall {
	return &pbgear.ConvictionVoting_RemoveVoteCall{
		Class: fields[0].Value.(uint32),
		Index: fields[1].Value.(uint32),
	}
}

func to_ConvictionVoting_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SpCoreCryptoAccountId32 {
	return &pbgear.ConvictionVoting_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Split(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Split {
	return &pbgear.ConvictionVoting_Split{
		Aye: fields[0].Value.(string),
		Nay: fields[1].Value.(string),
	}
}

func to_ConvictionVoting_SplitAbstain(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SplitAbstain {
	return &pbgear.ConvictionVoting_SplitAbstain{
		Aye:     fields[0].Value.(string),
		Nay:     fields[1].Value.(string),
		Abstain: fields[2].Value.(string),
	}
}

func to_ConvictionVoting_Standard(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Standard {
	return &pbgear.ConvictionVoting_Standard{
		Vote:    to_ConvictionVoting_PalletConvictionVotingVoteVote(fields[0].Value.([]*registry.DecodedField)),
		Balance: fields[1].Value.(string),
	}
}

func to_ConvictionVoting_Target(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Target {
	return &pbgear.ConvictionVoting_Target{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_To(fields []*registry.DecodedField) *pbgear.ConvictionVoting_To {
	return &pbgear.ConvictionVoting_To{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_TupleNull(fields []*registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
	return &pbgear.ConvictionVoting_TupleNull{}
}
func to_ConvictionVoting_UndelegateCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UndelegateCall {
	return &pbgear.ConvictionVoting_UndelegateCall{
		Class: fields[0].Value.(uint32),
	}
}

func to_ConvictionVoting_UnlockCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UnlockCall {
	return &pbgear.ConvictionVoting_UnlockCall{
		Class:  fields[0].Value.(uint32),
		Target: to_ConvictionVoting_Target(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_Vote(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Vote {
	return &pbgear.ConvictionVoting_Vote{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ConvictionVoting_VoteCall(fields []*registry.DecodedField) *pbgear.ConvictionVoting_VoteCall {
	return &pbgear.ConvictionVoting_VoteCall{
		PollIndex: to_ConvictionVoting_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Vote:      to_ConvictionVoting_Vote(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{
		Value: to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_ElectionProviderMultiPhase_GovernanceFallbackCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
	return &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{
		MaybeMaxVoters:  fields[0].Value.(uint32),
		MaybeMaxTargets: fields[1].Value.(uint32),
	}
}

func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{
		Solution: to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields[0].Value.([]*registry.DecodedField)),
		Score:    to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[1].Value.([]*registry.DecodedField)),
		Round:    fields[2].Value.(uint32),
	}
}

func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{
		Voters:  to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Targets: to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
	return &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{
		Supports: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
	return &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{
		MaybeNextScore: to_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{
		Value: fields[0].Value.(uint32),
	}
}

func to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
	return &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
	return &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{
		MinimalStake:    fields[0].Value.(string),
		SumStake:        fields[1].Value.(string),
		SumStakeSquared: fields[2].Value.(string),
	}
}

func to_ElectionProviderMultiPhase_SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
	return &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{
		Total:  fields[0].Value.(string),
		Voters: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SubmitCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitCall {
	return &pbgear.ElectionProviderMultiPhase_SubmitCall{
		RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_SubmitUnsignedCall(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
	return &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{
		RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.([]*registry.DecodedField)),
		Witness:     to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.([]*registry.DecodedField)),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{
		Value_0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_ElectionProviderMultiPhase_SpNposElectionsSupport(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{
		Value_0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: fields[1].Value.(string),
	}
}

func to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
	return &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{
		Votes_1:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Votes_2:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields[1].Value.([]*registry.DecodedField)),
		Votes_3:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields[2].Value.([]*registry.DecodedField)),
		Votes_4:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields[3].Value.([]*registry.DecodedField)),
		Votes_5:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields[4].Value.([]*registry.DecodedField)),
		Votes_6:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields[5].Value.([]*registry.DecodedField)),
		Votes_7:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields[6].Value.([]*registry.DecodedField)),
		Votes_8:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields[7].Value.([]*registry.DecodedField)),
		Votes_9:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields[8].Value.([]*registry.DecodedField)),
		Votes_10: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields[9].Value.([]*registry.DecodedField)),
		Votes_11: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields[10].Value.([]*registry.DecodedField)),
		Votes_12: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields[11].Value.([]*registry.DecodedField)),
		Votes_13: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields[12].Value.([]*registry.DecodedField)),
		Votes_14: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields[13].Value.([]*registry.DecodedField)),
		Votes_15: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields[14].Value.([]*registry.DecodedField)),
		Votes_16: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields[15].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_AddMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_AddMemberCall {
	return &pbgear.FellowshipCollective_AddMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_Address20(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address20 {
	return &pbgear.FellowshipCollective_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_Address32(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address32 {
	return &pbgear.FellowshipCollective_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_CleanupPollCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CleanupPollCall {
	return &pbgear.FellowshipCollective_CleanupPollCall{
		PollIndex: fields[0].Value.(uint32),
		Max:       fields[1].Value.(uint32),
	}
}

func to_FellowshipCollective_CompactTupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
	return &pbgear.FellowshipCollective_CompactTupleNull{
		Value: to_FellowshipCollective_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_DemoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_DemoteMemberCall {
	return &pbgear.FellowshipCollective_DemoteMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_Id(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Id {
	return &pbgear.FellowshipCollective_Id{
		Value_0: to_FellowshipCollective_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_Index(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Index {
	return &pbgear.FellowshipCollective_Index{
		Value_0: to_FellowshipCollective_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_PromoteMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_PromoteMemberCall {
	return &pbgear.FellowshipCollective_PromoteMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_Raw(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Raw {
	return &pbgear.FellowshipCollective_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_RemoveMemberCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_RemoveMemberCall {
	return &pbgear.FellowshipCollective_RemoveMemberCall{
		Who:     to_FellowshipCollective_Who(fields[0].Value.([]*registry.DecodedField)),
		MinRank: fields[1].Value.(uint32),
	}
}

func to_FellowshipCollective_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.FellowshipCollective_SpCoreCryptoAccountId32 {
	return &pbgear.FellowshipCollective_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipCollective_TupleNull(fields []*registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
	return &pbgear.FellowshipCollective_TupleNull{}
}
func to_FellowshipCollective_VoteCall(fields []*registry.DecodedField) *pbgear.FellowshipCollective_VoteCall {
	return &pbgear.FellowshipCollective_VoteCall{
		Poll: fields[0].Value.(uint32),
		Aye:  fields[1].Value.(bool),
	}
}

func to_FellowshipCollective_Who(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Who {
	return &pbgear.FellowshipCollective_Who{
		Value: to_oneof_FellowshipCollective_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_FellowshipCollective_Value(in *registry.DecodedField) *pbgear.FellowshipCollective_Value {
	return nil
}

func to_FellowshipReferenda_After(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_After {
	return &pbgear.FellowshipReferenda_After{
		Value_0: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_At(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_At {
	return &pbgear.FellowshipReferenda_At{
		Value_0: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_CancelCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_CancelCall {
	return &pbgear.FellowshipReferenda_CancelCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
	return &pbgear.FellowshipReferenda_EnactmentMoment{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_FellowshipReferenda_Value(in *registry.DecodedField) *pbgear.FellowshipReferenda_Value {
	return nil
}

func to_FellowshipReferenda_Inline(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Inline {
	return &pbgear.FellowshipReferenda_Inline{
		Value_0: to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_KillCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_KillCall {
	return &pbgear.FellowshipReferenda_KillCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_Legacy(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Legacy {
	return &pbgear.FellowshipReferenda_Legacy{
		Hash: to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_Lookup(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Lookup {
	return &pbgear.FellowshipReferenda_Lookup{
		Hash: to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
		Len:  fields[1].Value.(uint32),
	}
}

func to_FellowshipReferenda_None(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_None {
	return &pbgear.FellowshipReferenda_None{}
}
func to_FellowshipReferenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_NudgeReferendumCall {
	return &pbgear.FellowshipReferenda_NudgeReferendumCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
	return &pbgear.FellowshipReferenda_OneFewerDecidingCall{
		Track: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_Origins(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Origins {
	return &pbgear.FellowshipReferenda_Origins{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
	return &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
	return &pbgear.FellowshipReferenda_PrimitiveTypesH256{
		Hash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_Proposal(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
	return &pbgear.FellowshipReferenda_Proposal{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
	return &pbgear.FellowshipReferenda_ProposalOrigin{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
	return &pbgear.FellowshipReferenda_RefundDecisionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
	return &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_FellowshipReferenda_Root(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Root {
	return &pbgear.FellowshipReferenda_Root{}
}
func to_FellowshipReferenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SetMetadataCall {
	return &pbgear.FellowshipReferenda_SetMetadataCall{
		Index:     fields[0].Value.(uint32),
		MaybeHash: to_optional_FellowshipReferenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_Signed(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Signed {
	return &pbgear.FellowshipReferenda_Signed{
		Value_0: to_FellowshipReferenda_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SpCoreCryptoAccountId32 {
	return &pbgear.FellowshipReferenda_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_SubmitCall(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SubmitCall {
	return &pbgear.FellowshipReferenda_SubmitCall{
		ProposalOrigin:  to_FellowshipReferenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField)),
		Proposal:        to_FellowshipReferenda_Proposal(fields[1].Value.([]*registry.DecodedField)),
		EnactmentMoment: to_FellowshipReferenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_System(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_System {
	return &pbgear.FellowshipReferenda_System{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_Value0(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Value0 {
	return &pbgear.FellowshipReferenda_Value0{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_FellowshipReferenda_Void(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Void {
	return &pbgear.FellowshipReferenda_Void{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_AppendPrograms(fields []*registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
	return &pbgear.GearVoucher_AppendPrograms{
		Value: to_oneof_GearVoucher_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_GearVoucher_Value(in *registry.DecodedField) *pbgear.GearVoucher_Value {
	return nil
}

func to_GearVoucher_BTreeSet(fields []*registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
	return &pbgear.GearVoucher_BTreeSet{
		Programs: to_repeated_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_Call(fields []*registry.DecodedField) *pbgear.GearVoucher_Call {
	return &pbgear.GearVoucher_Call{
		Value: to_oneof_GearVoucher_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_CallCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallCall {
	return &pbgear.GearVoucher_CallCall{
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField)),
		Call:      to_GearVoucher_Call(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_CallDeprecatedCall(fields []*registry.DecodedField) *pbgear.GearVoucher_CallDeprecatedCall {
	return &pbgear.GearVoucher_CallDeprecatedCall{
		Call: to_GearVoucher_Call(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_DeclineCall(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineCall {
	return &pbgear.GearVoucher_DeclineCall{
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_DeclineVoucher(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineVoucher {
	return &pbgear.GearVoucher_DeclineVoucher{}
}
func to_GearVoucher_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
	return &pbgear.GearVoucher_GprimitivesActorId{
		Programs: to_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesMessageId {
	return &pbgear.GearVoucher_GprimitivesMessageId{
		ReplyToId: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_IssueCall(fields []*registry.DecodedField) *pbgear.GearVoucher_IssueCall {
	return &pbgear.GearVoucher_IssueCall{
		Spender:       to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Balance:       fields[1].Value.(string),
		Programs:      to_optional_GearVoucher_BTreeSet(fields[2].Value.([]*registry.DecodedField)),
		CodeUploading: fields[3].Value.(bool),
		Duration:      fields[4].Value.(uint32),
	}
}

func to_GearVoucher_None(fields []*registry.DecodedField) *pbgear.GearVoucher_None {
	return &pbgear.GearVoucher_None{}
}
func to_GearVoucher_PalletGearVoucherInternalVoucherId(fields []*registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{
		VoucherId: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_RevokeCall(fields []*registry.DecodedField) *pbgear.GearVoucher_RevokeCall {
	return &pbgear.GearVoucher_RevokeCall{
		Spender:   to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_SendMessage(fields []*registry.DecodedField) *pbgear.GearVoucher_SendMessage {
	return &pbgear.GearVoucher_SendMessage{
		Destination: to_GearVoucher_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
		Payload:     to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		GasLimit:    fields[2].Value.(uint64),
		Value:       fields[3].Value.(string),
		KeepAlive:   fields[4].Value.(bool),
	}
}

func to_GearVoucher_SendReply(fields []*registry.DecodedField) *pbgear.GearVoucher_SendReply {
	return &pbgear.GearVoucher_SendReply{
		ReplyToId: to_GearVoucher_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
		Payload:   to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		GasLimit:  fields[2].Value.(uint64),
		Value:     fields[3].Value.(string),
		KeepAlive: fields[4].Value.(bool),
	}
}

func to_GearVoucher_Some(fields []*registry.DecodedField) *pbgear.GearVoucher_Some {
	return &pbgear.GearVoucher_Some{
		Value_0: to_GearVoucher_BTreeSet(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
	return &pbgear.GearVoucher_SpCoreCryptoAccountId32{
		Spender: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_GearVoucher_UpdateCall(fields []*registry.DecodedField) *pbgear.GearVoucher_UpdateCall {
	return &pbgear.GearVoucher_UpdateCall{
		Spender:         to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		VoucherId:       to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.([]*registry.DecodedField)),
		MoveOwnership:   to_optional_GearVoucher_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField)),
		BalanceTopUp:    fields[3].Value.(string),
		AppendPrograms:  to_optional_GearVoucher_AppendPrograms(fields[4].Value.([]*registry.DecodedField)),
		CodeUploading:   fields[5].Value.(bool),
		ProlongDuration: fields[6].Value.(uint32),
	}
}

func to_GearVoucher_UploadCode(fields []*registry.DecodedField) *pbgear.GearVoucher_UploadCode {
	return &pbgear.GearVoucher_UploadCode{
		Code: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_ClaimValueCall(fields []*registry.DecodedField) *pbgear.Gear_ClaimValueCall {
	return &pbgear.Gear_ClaimValueCall{
		MessageId: to_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_CreateProgramCall(fields []*registry.DecodedField) *pbgear.Gear_CreateProgramCall {
	return &pbgear.Gear_CreateProgramCall{
		CodeId:      to_Gear_GprimitivesCodeId(fields[0].Value.([]*registry.DecodedField)),
		Salt:        to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		InitPayload: to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
		GasLimit:    fields[3].Value.(uint64),
		Value:       fields[4].Value.(string),
		KeepAlive:   fields[5].Value.(bool),
	}
}

func to_Gear_GprimitivesActorId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
	return &pbgear.Gear_GprimitivesActorId{
		Destination: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_GprimitivesCodeId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
	return &pbgear.Gear_GprimitivesCodeId{
		CodeId: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_GprimitivesMessageId(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
	return &pbgear.Gear_GprimitivesMessageId{
		ReplyToId: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_RunCall(fields []*registry.DecodedField) *pbgear.Gear_RunCall {
	return &pbgear.Gear_RunCall{
		MaxGas: fields[0].Value.(uint64),
	}
}

func to_Gear_SendMessageCall(fields []*registry.DecodedField) *pbgear.Gear_SendMessageCall {
	return &pbgear.Gear_SendMessageCall{
		Destination: to_Gear_GprimitivesActorId(fields[0].Value.([]*registry.DecodedField)),
		Payload:     to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		GasLimit:    fields[2].Value.(uint64),
		Value:       fields[3].Value.(string),
		KeepAlive:   fields[4].Value.(bool),
	}
}

func to_Gear_SendReplyCall(fields []*registry.DecodedField) *pbgear.Gear_SendReplyCall {
	return &pbgear.Gear_SendReplyCall{
		ReplyToId: to_Gear_GprimitivesMessageId(fields[0].Value.([]*registry.DecodedField)),
		Payload:   to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		GasLimit:  fields[2].Value.(uint64),
		Value:     fields[3].Value.(string),
		KeepAlive: fields[4].Value.(bool),
	}
}

func to_Gear_SetExecuteInherentCall(fields []*registry.DecodedField) *pbgear.Gear_SetExecuteInherentCall {
	return &pbgear.Gear_SetExecuteInherentCall{
		Value: fields[0].Value.(bool),
	}
}

func to_Gear_UploadCodeCall(fields []*registry.DecodedField) *pbgear.Gear_UploadCodeCall {
	return &pbgear.Gear_UploadCodeCall{
		Code: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Gear_UploadProgramCall(fields []*registry.DecodedField) *pbgear.Gear_UploadProgramCall {
	return &pbgear.Gear_UploadProgramCall{
		Code:        to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Salt:        to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
		InitPayload: to_repeated_uint32(fields[2].Value.([]*registry.DecodedField)),
		GasLimit:    fields[3].Value.(uint64),
		Value:       fields[4].Value.(string),
		KeepAlive:   fields[5].Value.(bool),
	}
}

func to_Grandpa_Equivocation(fields []*registry.DecodedField) *pbgear.Grandpa_Equivocation {
	return &pbgear.Grandpa_Equivocation{
		Value: to_oneof_Grandpa_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Grandpa_Value(in *registry.DecodedField) *pbgear.Grandpa_Value {
	return nil
}

func to_Grandpa_FinalityGrandpaEquivocation(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaEquivocation {
	return &pbgear.Grandpa_FinalityGrandpaEquivocation{
		RoundNumber: fields[0].Value.(uint64),
		Identity:    to_Grandpa_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField)),
		First:       to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[2].Value.([]*registry.DecodedField)),
		Second:      to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_FinalityGrandpaPrevote(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaPrevote {
	return &pbgear.Grandpa_FinalityGrandpaPrevote{
		TargetHash:   to_Grandpa_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
		TargetNumber: fields[1].Value.(uint32),
	}
}

func to_Grandpa_GrandpaTrieNodesList(fields []*registry.DecodedField) *pbgear.Grandpa_GrandpaTrieNodesList {
	return &pbgear.Grandpa_GrandpaTrieNodesList{
		TrieNodes: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_NoteStalledCall(fields []*registry.DecodedField) *pbgear.Grandpa_NoteStalledCall {
	return &pbgear.Grandpa_NoteStalledCall{
		Delay:                    fields[0].Value.(uint32),
		BestFinalizedBlockNumber: fields[1].Value.(uint32),
	}
}

func to_Grandpa_Precommit(fields []*registry.DecodedField) *pbgear.Grandpa_Precommit {
	return &pbgear.Grandpa_Precommit{
		Value_0: to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_Prevote(fields []*registry.DecodedField) *pbgear.Grandpa_Prevote {
	return &pbgear.Grandpa_Prevote{
		Value_0: to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Grandpa_PrimitiveTypesH256 {
	return &pbgear.Grandpa_PrimitiveTypesH256{
		TargetHash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_ReportEquivocationCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationCall {
	return &pbgear.Grandpa_ReportEquivocationCall{
		EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
		KeyOwnerProof:     to_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_ReportEquivocationUnsignedCall(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
	return &pbgear.Grandpa_ReportEquivocationUnsignedCall{
		EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value.([]*registry.DecodedField)),
		KeyOwnerProof:     to_Grandpa_SpSessionMembershipProof(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
	return &pbgear.Grandpa_SpConsensusGrandpaAppPublic{
		Identity: to_Grandpa_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
	return &pbgear.Grandpa_SpConsensusGrandpaAppSignature{
		Value_1: to_Grandpa_SpCoreEd25519Signature(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpConsensusGrandpaEquivocationProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
	return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{
		SetId:        fields[0].Value.(uint64),
		Equivocation: to_Grandpa_Equivocation(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Public {
	return &pbgear.Grandpa_SpCoreEd25519Public{
		Identity: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpCoreEd25519Signature(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Signature {
	return &pbgear.Grandpa_SpCoreEd25519Signature{
		Value_1: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Grandpa_SpSessionMembershipProof(fields []*registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
	return &pbgear.Grandpa_SpSessionMembershipProof{
		Session:        fields[0].Value.(uint32),
		TrieNodes:      to_repeated_Grandpa_GrandpaTrieNodesList(fields[1].Value.([]*registry.DecodedField)),
		ValidatorCount: fields[2].Value.(uint32),
	}
}

func to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields []*registry.DecodedField) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	return &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{
		Value_0: to_Grandpa_FinalityGrandpaPrevote(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_Grandpa_SpConsensusGrandpaAppSignature(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Account(fields []*registry.DecodedField) *pbgear.Identity_Account {
	return &pbgear.Identity_Account{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Identity_Value(in *registry.DecodedField) *pbgear.Identity_Value {
	return nil
}

func to_Identity_AddRegistrarCall(fields []*registry.DecodedField) *pbgear.Identity_AddRegistrarCall {
	return &pbgear.Identity_AddRegistrarCall{
		Account: to_Identity_Account(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_AddSubCall(fields []*registry.DecodedField) *pbgear.Identity_AddSubCall {
	return &pbgear.Identity_AddSubCall{
		Sub:  to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
		Data: to_Identity_Data(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Address20(fields []*registry.DecodedField) *pbgear.Identity_Address20 {
	return &pbgear.Identity_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Address32(fields []*registry.DecodedField) *pbgear.Identity_Address32 {
	return &pbgear.Identity_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_BlakeTwo256(fields []*registry.DecodedField) *pbgear.Identity_BlakeTwo256 {
	return &pbgear.Identity_BlakeTwo256{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{
		Additional: to_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_CancelRequestCall(fields []*registry.DecodedField) *pbgear.Identity_CancelRequestCall {
	return &pbgear.Identity_CancelRequestCall{
		RegIndex: fields[0].Value.(uint32),
	}
}

func to_Identity_ClearIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_ClearIdentityCall {
	return &pbgear.Identity_ClearIdentityCall{}
}
func to_Identity_CompactString(fields []*registry.DecodedField) *pbgear.Identity_CompactString {
	return &pbgear.Identity_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_Identity_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Identity_CompactTupleNull {
	return &pbgear.Identity_CompactTupleNull{
		Value: to_Identity_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_CompactUint32(fields []*registry.DecodedField) *pbgear.Identity_CompactUint32 {
	return &pbgear.Identity_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Identity_Data(fields []*registry.DecodedField) *pbgear.Identity_Data {
	return &pbgear.Identity_Data{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Display(fields []*registry.DecodedField) *pbgear.Identity_Display {
	return &pbgear.Identity_Display{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Email(fields []*registry.DecodedField) *pbgear.Identity_Email {
	return &pbgear.Identity_Email{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Erroneous(fields []*registry.DecodedField) *pbgear.Identity_Erroneous {
	return &pbgear.Identity_Erroneous{}
}
func to_Identity_FeePaid(fields []*registry.DecodedField) *pbgear.Identity_FeePaid {
	return &pbgear.Identity_FeePaid{
		Value_0: fields[0].Value.(string),
	}
}

func to_Identity_Id(fields []*registry.DecodedField) *pbgear.Identity_Id {
	return &pbgear.Identity_Id{
		Value_0: to_Identity_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Image(fields []*registry.DecodedField) *pbgear.Identity_Image {
	return &pbgear.Identity_Image{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Index(fields []*registry.DecodedField) *pbgear.Identity_Index {
	return &pbgear.Identity_Index{
		Value_0: to_Identity_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Judgement(fields []*registry.DecodedField) *pbgear.Identity_Judgement {
	return &pbgear.Identity_Judgement{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Keccak256(fields []*registry.DecodedField) *pbgear.Identity_Keccak256 {
	return &pbgear.Identity_Keccak256{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_KillIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_KillIdentityCall {
	return &pbgear.Identity_KillIdentityCall{
		Target: to_Identity_Target(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_KnownGood(fields []*registry.DecodedField) *pbgear.Identity_KnownGood {
	return &pbgear.Identity_KnownGood{}
}
func to_Identity_Legal(fields []*registry.DecodedField) *pbgear.Identity_Legal {
	return &pbgear.Identity_Legal{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_LowQuality(fields []*registry.DecodedField) *pbgear.Identity_LowQuality {
	return &pbgear.Identity_LowQuality{}
}
func to_Identity_New(fields []*registry.DecodedField) *pbgear.Identity_New {
	return &pbgear.Identity_New{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_None(fields []*registry.DecodedField) *pbgear.Identity_None {
	return &pbgear.Identity_None{}
}
func to_Identity_OutOfDate(fields []*registry.DecodedField) *pbgear.Identity_OutOfDate {
	return &pbgear.Identity_OutOfDate{}
}
func to_Identity_PalletIdentitySimpleIdentityInfo(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	return &pbgear.Identity_PalletIdentitySimpleIdentityInfo{
		Additional:     to_Identity_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
		Display:        to_Identity_Display(fields[1].Value.([]*registry.DecodedField)),
		Legal:          to_Identity_Legal(fields[2].Value.([]*registry.DecodedField)),
		Web:            to_Identity_Web(fields[3].Value.([]*registry.DecodedField)),
		Riot:           to_Identity_Riot(fields[4].Value.([]*registry.DecodedField)),
		Email:          to_Identity_Email(fields[5].Value.([]*registry.DecodedField)),
		PgpFingerprint: to_repeated_uint32(fields[6].Value.([]*registry.DecodedField)),
		Image:          to_Identity_Image(fields[7].Value.([]*registry.DecodedField)),
		Twitter:        to_Identity_Twitter(fields[8].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_PalletIdentityTypesBitFlags(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
	return &pbgear.Identity_PalletIdentityTypesBitFlags{
		Fields: fields[0].Value.(uint64),
	}
}

func to_Identity_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
	return &pbgear.Identity_PrimitiveTypesH256{
		Identity: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_ProvideJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_ProvideJudgementCall {
	return &pbgear.Identity_ProvideJudgementCall{
		RegIndex:  to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Target:    to_Identity_Target(fields[1].Value.([]*registry.DecodedField)),
		Judgement: to_Identity_Judgement(fields[2].Value.([]*registry.DecodedField)),
		Identity:  to_Identity_PrimitiveTypesH256(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_QuitSubCall(fields []*registry.DecodedField) *pbgear.Identity_QuitSubCall {
	return &pbgear.Identity_QuitSubCall{}
}
func to_Identity_Raw(fields []*registry.DecodedField) *pbgear.Identity_Raw {
	return &pbgear.Identity_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw0(fields []*registry.DecodedField) *pbgear.Identity_Raw0 {
	return &pbgear.Identity_Raw0{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw1(fields []*registry.DecodedField) *pbgear.Identity_Raw1 {
	return &pbgear.Identity_Raw1{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw10(fields []*registry.DecodedField) *pbgear.Identity_Raw10 {
	return &pbgear.Identity_Raw10{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw11(fields []*registry.DecodedField) *pbgear.Identity_Raw11 {
	return &pbgear.Identity_Raw11{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw12(fields []*registry.DecodedField) *pbgear.Identity_Raw12 {
	return &pbgear.Identity_Raw12{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw13(fields []*registry.DecodedField) *pbgear.Identity_Raw13 {
	return &pbgear.Identity_Raw13{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw14(fields []*registry.DecodedField) *pbgear.Identity_Raw14 {
	return &pbgear.Identity_Raw14{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw15(fields []*registry.DecodedField) *pbgear.Identity_Raw15 {
	return &pbgear.Identity_Raw15{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw16(fields []*registry.DecodedField) *pbgear.Identity_Raw16 {
	return &pbgear.Identity_Raw16{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw17(fields []*registry.DecodedField) *pbgear.Identity_Raw17 {
	return &pbgear.Identity_Raw17{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw18(fields []*registry.DecodedField) *pbgear.Identity_Raw18 {
	return &pbgear.Identity_Raw18{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw19(fields []*registry.DecodedField) *pbgear.Identity_Raw19 {
	return &pbgear.Identity_Raw19{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw2(fields []*registry.DecodedField) *pbgear.Identity_Raw2 {
	return &pbgear.Identity_Raw2{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw20(fields []*registry.DecodedField) *pbgear.Identity_Raw20 {
	return &pbgear.Identity_Raw20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw21(fields []*registry.DecodedField) *pbgear.Identity_Raw21 {
	return &pbgear.Identity_Raw21{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw22(fields []*registry.DecodedField) *pbgear.Identity_Raw22 {
	return &pbgear.Identity_Raw22{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw23(fields []*registry.DecodedField) *pbgear.Identity_Raw23 {
	return &pbgear.Identity_Raw23{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw24(fields []*registry.DecodedField) *pbgear.Identity_Raw24 {
	return &pbgear.Identity_Raw24{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw25(fields []*registry.DecodedField) *pbgear.Identity_Raw25 {
	return &pbgear.Identity_Raw25{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw26(fields []*registry.DecodedField) *pbgear.Identity_Raw26 {
	return &pbgear.Identity_Raw26{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw27(fields []*registry.DecodedField) *pbgear.Identity_Raw27 {
	return &pbgear.Identity_Raw27{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw28(fields []*registry.DecodedField) *pbgear.Identity_Raw28 {
	return &pbgear.Identity_Raw28{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw29(fields []*registry.DecodedField) *pbgear.Identity_Raw29 {
	return &pbgear.Identity_Raw29{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw3(fields []*registry.DecodedField) *pbgear.Identity_Raw3 {
	return &pbgear.Identity_Raw3{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw30(fields []*registry.DecodedField) *pbgear.Identity_Raw30 {
	return &pbgear.Identity_Raw30{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw31(fields []*registry.DecodedField) *pbgear.Identity_Raw31 {
	return &pbgear.Identity_Raw31{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw32(fields []*registry.DecodedField) *pbgear.Identity_Raw32 {
	return &pbgear.Identity_Raw32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw4(fields []*registry.DecodedField) *pbgear.Identity_Raw4 {
	return &pbgear.Identity_Raw4{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw5(fields []*registry.DecodedField) *pbgear.Identity_Raw5 {
	return &pbgear.Identity_Raw5{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw6(fields []*registry.DecodedField) *pbgear.Identity_Raw6 {
	return &pbgear.Identity_Raw6{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw7(fields []*registry.DecodedField) *pbgear.Identity_Raw7 {
	return &pbgear.Identity_Raw7{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw8(fields []*registry.DecodedField) *pbgear.Identity_Raw8 {
	return &pbgear.Identity_Raw8{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Raw9(fields []*registry.DecodedField) *pbgear.Identity_Raw9 {
	return &pbgear.Identity_Raw9{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Reasonable(fields []*registry.DecodedField) *pbgear.Identity_Reasonable {
	return &pbgear.Identity_Reasonable{}
}
func to_Identity_RemoveSubCall(fields []*registry.DecodedField) *pbgear.Identity_RemoveSubCall {
	return &pbgear.Identity_RemoveSubCall{
		Sub: to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_RenameSubCall(fields []*registry.DecodedField) *pbgear.Identity_RenameSubCall {
	return &pbgear.Identity_RenameSubCall{
		Sub:  to_Identity_Sub(fields[0].Value.([]*registry.DecodedField)),
		Data: to_Identity_Data(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_RequestJudgementCall(fields []*registry.DecodedField) *pbgear.Identity_RequestJudgementCall {
	return &pbgear.Identity_RequestJudgementCall{
		RegIndex: to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		MaxFee:   to_Identity_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Riot(fields []*registry.DecodedField) *pbgear.Identity_Riot {
	return &pbgear.Identity_Riot{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SetAccountIdCall(fields []*registry.DecodedField) *pbgear.Identity_SetAccountIdCall {
	return &pbgear.Identity_SetAccountIdCall{
		Index: to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		New:   to_Identity_New(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SetFeeCall(fields []*registry.DecodedField) *pbgear.Identity_SetFeeCall {
	return &pbgear.Identity_SetFeeCall{
		Index: to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Fee:   to_Identity_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SetFieldsCall(fields []*registry.DecodedField) *pbgear.Identity_SetFieldsCall {
	return &pbgear.Identity_SetFieldsCall{
		Index:  to_Identity_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
		Fields: to_Identity_PalletIdentityTypesBitFlags(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SetIdentityCall(fields []*registry.DecodedField) *pbgear.Identity_SetIdentityCall {
	return &pbgear.Identity_SetIdentityCall{
		Info: to_Identity_PalletIdentitySimpleIdentityInfo(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SetSubsCall(fields []*registry.DecodedField) *pbgear.Identity_SetSubsCall {
	return &pbgear.Identity_SetSubsCall{
		Subs: to_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Sha256(fields []*registry.DecodedField) *pbgear.Identity_Sha256 {
	return &pbgear.Identity_Sha256{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_ShaThree256(fields []*registry.DecodedField) *pbgear.Identity_ShaThree256 {
	return &pbgear.Identity_ShaThree256{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
	return &pbgear.Identity_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Sub(fields []*registry.DecodedField) *pbgear.Identity_Sub {
	return &pbgear.Identity_Sub{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Target(fields []*registry.DecodedField) *pbgear.Identity_Target {
	return &pbgear.Identity_Target{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_TupleNull(fields []*registry.DecodedField) *pbgear.Identity_TupleNull {
	return &pbgear.Identity_TupleNull{}
}
func to_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{
		Value_0: to_Identity_Value0(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_Identity_Value1(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields []*registry.DecodedField) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{
		Value_0: to_Identity_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_Identity_Value1(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Twitter(fields []*registry.DecodedField) *pbgear.Identity_Twitter {
	return &pbgear.Identity_Twitter{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Unknown(fields []*registry.DecodedField) *pbgear.Identity_Unknown {
	return &pbgear.Identity_Unknown{}
}
func to_Identity_Value0(fields []*registry.DecodedField) *pbgear.Identity_Value0 {
	return &pbgear.Identity_Value0{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Value1(fields []*registry.DecodedField) *pbgear.Identity_Value1 {
	return &pbgear.Identity_Value1{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Identity_Web(fields []*registry.DecodedField) *pbgear.Identity_Web {
	return &pbgear.Identity_Web{
		Value: to_oneof_Identity_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ImOnline_HeartbeatCall(fields []*registry.DecodedField) *pbgear.ImOnline_HeartbeatCall {
	return &pbgear.ImOnline_HeartbeatCall{
		Heartbeat: to_ImOnline_PalletImOnlineHeartbeat(fields[0].Value.([]*registry.DecodedField)),
		Signature: to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_ImOnline_PalletImOnlineHeartbeat(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	return &pbgear.ImOnline_PalletImOnlineHeartbeat{
		BlockNumber:    fields[0].Value.(uint32),
		SessionIndex:   fields[1].Value.(uint32),
		AuthorityIndex: fields[2].Value.(uint32),
		ValidatorsLen:  fields[3].Value.(uint32),
	}
}

func to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	return &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{
		Signature: to_ImOnline_SpCoreSr25519Signature(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_ImOnline_SpCoreSr25519Signature(fields []*registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
	return &pbgear.ImOnline_SpCoreSr25519Signature{
		Signature: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Multisig_ApproveAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_ApproveAsMultiCall {
	return &pbgear.Multisig_ApproveAsMultiCall{
		Threshold:        fields[0].Value.(uint32),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
		MaybeTimepoint:   to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
		CallHash:         to_repeated_uint32(fields[3].Value.([]*registry.DecodedField)),
		MaxWeight:        to_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_Multisig_AsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiCall {
	return &pbgear.Multisig_AsMultiCall{
		Threshold:        fields[0].Value.(uint32),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
		MaybeTimepoint:   to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
		Call:             to_oneof_Multisig_Call(fields[3].Value.([]*registry.DecodedField)),
		MaxWeight:        to_Multisig_SpWeightsWeightV2Weight(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Multisig_Call(in *registry.DecodedField) *pbgear.Multisig_Call {
	return nil
}

func to_Multisig_AsMultiThreshold1Call(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiThreshold1Call {
	return &pbgear.Multisig_AsMultiThreshold1Call{
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Call:             to_oneof_Multisig_Call(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Multisig_CancelAsMultiCall(fields []*registry.DecodedField) *pbgear.Multisig_CancelAsMultiCall {
	return &pbgear.Multisig_CancelAsMultiCall{
		Threshold:        fields[0].Value.(uint32),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
		Timepoint:        to_Multisig_PalletMultisigTimepoint(fields[2].Value.([]*registry.DecodedField)),
		CallHash:         to_repeated_uint32(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_Multisig_CompactUint64(fields []*registry.DecodedField) *pbgear.Multisig_CompactUint64 {
	return &pbgear.Multisig_CompactUint64{
		Value: fields[0].Value.(uint64),
	}
}

func to_Multisig_PalletMultisigTimepoint(fields []*registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
	return &pbgear.Multisig_PalletMultisigTimepoint{
		Height: fields[0].Value.(uint32),
		Index:  fields[1].Value.(uint32),
	}
}

func to_Multisig_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Multisig_SpCoreCryptoAccountId32 {
	return &pbgear.Multisig_SpCoreCryptoAccountId32{
		OtherSignatories: to_Multisig_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Multisig_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
	return &pbgear.Multisig_SpWeightsWeightV2Weight{
		RefTime:   to_Multisig_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
		ProofSize: to_Multisig_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Address20(fields []*registry.DecodedField) *pbgear.NominationPools_Address20 {
	return &pbgear.NominationPools_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Address32(fields []*registry.DecodedField) *pbgear.NominationPools_Address32 {
	return &pbgear.NominationPools_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_AdjustPoolDepositCall(fields []*registry.DecodedField) *pbgear.NominationPools_AdjustPoolDepositCall {
	return &pbgear.NominationPools_AdjustPoolDepositCall{
		PoolId: fields[0].Value.(uint32),
	}
}

func to_NominationPools_Blocked(fields []*registry.DecodedField) *pbgear.NominationPools_Blocked {
	return &pbgear.NominationPools_Blocked{}
}
func to_NominationPools_BondExtraCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraCall {
	return &pbgear.NominationPools_BondExtraCall{
		Extra: to_NominationPools_Extra(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_BondExtraOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraOtherCall {
	return &pbgear.NominationPools_BondExtraOtherCall{
		Member: to_NominationPools_Member(fields[0].Value.([]*registry.DecodedField)),
		Extra:  to_NominationPools_Extra(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Bouncer(fields []*registry.DecodedField) *pbgear.NominationPools_Bouncer {
	return &pbgear.NominationPools_Bouncer{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_NominationPools_Value(in *registry.DecodedField) *pbgear.NominationPools_Value {
	return nil
}

func to_NominationPools_ChillCall(fields []*registry.DecodedField) *pbgear.NominationPools_ChillCall {
	return &pbgear.NominationPools_ChillCall{
		PoolId: fields[0].Value.(uint32),
	}
}

func to_NominationPools_ClaimCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimCommissionCall {
	return &pbgear.NominationPools_ClaimCommissionCall{
		PoolId: fields[0].Value.(uint32),
	}
}

func to_NominationPools_ClaimPayoutCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutCall {
	return &pbgear.NominationPools_ClaimPayoutCall{}
}
func to_NominationPools_ClaimPayoutOtherCall(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutOtherCall {
	return &pbgear.NominationPools_ClaimPayoutOtherCall{
		Other: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_CompactString(fields []*registry.DecodedField) *pbgear.NominationPools_CompactString {
	return &pbgear.NominationPools_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_NominationPools_CompactTupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
	return &pbgear.NominationPools_CompactTupleNull{
		Value: to_NominationPools_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_CreateCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateCall {
	return &pbgear.NominationPools_CreateCall{
		Amount:    to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Root:      to_NominationPools_Root(fields[1].Value.([]*registry.DecodedField)),
		Nominator: to_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField)),
		Bouncer:   to_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_CreateWithPoolIdCall(fields []*registry.DecodedField) *pbgear.NominationPools_CreateWithPoolIdCall {
	return &pbgear.NominationPools_CreateWithPoolIdCall{
		Amount:    to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Root:      to_NominationPools_Root(fields[1].Value.([]*registry.DecodedField)),
		Nominator: to_NominationPools_Nominator(fields[2].Value.([]*registry.DecodedField)),
		Bouncer:   to_NominationPools_Bouncer(fields[3].Value.([]*registry.DecodedField)),
		PoolId:    fields[4].Value.(uint32),
	}
}

func to_NominationPools_Destroying(fields []*registry.DecodedField) *pbgear.NominationPools_Destroying {
	return &pbgear.NominationPools_Destroying{}
}
func to_NominationPools_Extra(fields []*registry.DecodedField) *pbgear.NominationPools_Extra {
	return &pbgear.NominationPools_Extra{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_FreeBalance(fields []*registry.DecodedField) *pbgear.NominationPools_FreeBalance {
	return &pbgear.NominationPools_FreeBalance{
		Value_0: fields[0].Value.(string),
	}
}

func to_NominationPools_GlobalMaxCommission(fields []*registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
	return &pbgear.NominationPools_GlobalMaxCommission{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Id(fields []*registry.DecodedField) *pbgear.NominationPools_Id {
	return &pbgear.NominationPools_Id{
		Value_0: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Index(fields []*registry.DecodedField) *pbgear.NominationPools_Index {
	return &pbgear.NominationPools_Index{
		Value_0: to_NominationPools_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_JoinCall(fields []*registry.DecodedField) *pbgear.NominationPools_JoinCall {
	return &pbgear.NominationPools_JoinCall{
		Amount: to_NominationPools_CompactString(fields[0].Value.([]*registry.DecodedField)),
		PoolId: fields[1].Value.(uint32),
	}
}

func to_NominationPools_MaxMembers(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembers {
	return &pbgear.NominationPools_MaxMembers{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_MaxMembersPerPool(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
	return &pbgear.NominationPools_MaxMembersPerPool{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_MaxPools(fields []*registry.DecodedField) *pbgear.NominationPools_MaxPools {
	return &pbgear.NominationPools_MaxPools{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Member(fields []*registry.DecodedField) *pbgear.NominationPools_Member {
	return &pbgear.NominationPools_Member{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_MemberAccount(fields []*registry.DecodedField) *pbgear.NominationPools_MemberAccount {
	return &pbgear.NominationPools_MemberAccount{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_MinCreateBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
	return &pbgear.NominationPools_MinCreateBond{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_MinJoinBond(fields []*registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
	return &pbgear.NominationPools_MinJoinBond{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_NewBouncer(fields []*registry.DecodedField) *pbgear.NominationPools_NewBouncer {
	return &pbgear.NominationPools_NewBouncer{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_NewNominator(fields []*registry.DecodedField) *pbgear.NominationPools_NewNominator {
	return &pbgear.NominationPools_NewNominator{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_NewRoot(fields []*registry.DecodedField) *pbgear.NominationPools_NewRoot {
	return &pbgear.NominationPools_NewRoot{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_NominateCall(fields []*registry.DecodedField) *pbgear.NominationPools_NominateCall {
	return &pbgear.NominationPools_NominateCall{
		PoolId:     fields[0].Value.(uint32),
		Validators: to_repeated_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Nominator(fields []*registry.DecodedField) *pbgear.NominationPools_Nominator {
	return &pbgear.NominationPools_Nominator{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Noop(fields []*registry.DecodedField) *pbgear.NominationPools_Noop {
	return &pbgear.NominationPools_Noop{}
}
func to_NominationPools_Open(fields []*registry.DecodedField) *pbgear.NominationPools_Open {
	return &pbgear.NominationPools_Open{}
}
func to_NominationPools_PalletNominationPoolsCommissionChangeRate(fields []*registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	return &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{
		MaxIncrease: to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
		MinDelay:    fields[1].Value.(uint32),
	}
}

func to_NominationPools_Permission(fields []*registry.DecodedField) *pbgear.NominationPools_Permission {
	return &pbgear.NominationPools_Permission{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Permissioned(fields []*registry.DecodedField) *pbgear.NominationPools_Permissioned {
	return &pbgear.NominationPools_Permissioned{}
}
func to_NominationPools_PermissionlessAll(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessAll {
	return &pbgear.NominationPools_PermissionlessAll{}
}
func to_NominationPools_PermissionlessCompound(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessCompound {
	return &pbgear.NominationPools_PermissionlessCompound{}
}
func to_NominationPools_PermissionlessWithdraw(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessWithdraw {
	return &pbgear.NominationPools_PermissionlessWithdraw{}
}
func to_NominationPools_PoolWithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
	return &pbgear.NominationPools_PoolWithdrawUnbondedCall{
		PoolId:           fields[0].Value.(uint32),
		NumSlashingSpans: fields[1].Value.(uint32),
	}
}

func to_NominationPools_Raw(fields []*registry.DecodedField) *pbgear.NominationPools_Raw {
	return &pbgear.NominationPools_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Remove(fields []*registry.DecodedField) *pbgear.NominationPools_Remove {
	return &pbgear.NominationPools_Remove{}
}
func to_NominationPools_Rewards(fields []*registry.DecodedField) *pbgear.NominationPools_Rewards {
	return &pbgear.NominationPools_Rewards{}
}
func to_NominationPools_Root(fields []*registry.DecodedField) *pbgear.NominationPools_Root {
	return &pbgear.NominationPools_Root{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_Set(fields []*registry.DecodedField) *pbgear.NominationPools_Set {
	return &pbgear.NominationPools_Set{
		Value_0: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetClaimPermissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetClaimPermissionCall {
	return &pbgear.NominationPools_SetClaimPermissionCall{
		Permission: to_NominationPools_Permission(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetCommissionCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionCall {
	return &pbgear.NominationPools_SetCommissionCall{
		PoolId:        fields[0].Value.(uint32),
		NewCommission: to_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetCommissionChangeRateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionChangeRateCall {
	return &pbgear.NominationPools_SetCommissionChangeRateCall{
		PoolId:     fields[0].Value.(uint32),
		ChangeRate: to_NominationPools_PalletNominationPoolsCommissionChangeRate(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetCommissionMaxCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionMaxCall {
	return &pbgear.NominationPools_SetCommissionMaxCall{
		PoolId:        fields[0].Value.(uint32),
		MaxCommission: to_NominationPools_SpArithmeticPerThingsPerbill(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetConfigsCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetConfigsCall {
	return &pbgear.NominationPools_SetConfigsCall{
		MinJoinBond:         to_NominationPools_MinJoinBond(fields[0].Value.([]*registry.DecodedField)),
		MinCreateBond:       to_NominationPools_MinCreateBond(fields[1].Value.([]*registry.DecodedField)),
		MaxPools:            to_NominationPools_MaxPools(fields[2].Value.([]*registry.DecodedField)),
		MaxMembers:          to_NominationPools_MaxMembers(fields[3].Value.([]*registry.DecodedField)),
		MaxMembersPerPool:   to_NominationPools_MaxMembersPerPool(fields[4].Value.([]*registry.DecodedField)),
		GlobalMaxCommission: to_NominationPools_GlobalMaxCommission(fields[5].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetMetadataCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetMetadataCall {
	return &pbgear.NominationPools_SetMetadataCall{
		PoolId:   fields[0].Value.(uint32),
		Metadata: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SetStateCall(fields []*registry.DecodedField) *pbgear.NominationPools_SetStateCall {
	return &pbgear.NominationPools_SetStateCall{
		PoolId: fields[0].Value.(uint32),
		State:  to_NominationPools_State(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
	return &pbgear.NominationPools_SpArithmeticPerThingsPerbill{
		Value_0: fields[0].Value.(uint32),
	}
}

func to_NominationPools_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_SpCoreCryptoAccountId32 {
	return &pbgear.NominationPools_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_State(fields []*registry.DecodedField) *pbgear.NominationPools_State {
	return &pbgear.NominationPools_State{
		Value: to_oneof_NominationPools_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_TupleNull(fields []*registry.DecodedField) *pbgear.NominationPools_TupleNull {
	return &pbgear.NominationPools_TupleNull{}
}
func to_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	return &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{
		Value_0: to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_NominationPools_SpCoreCryptoAccountId32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_UnbondCall(fields []*registry.DecodedField) *pbgear.NominationPools_UnbondCall {
	return &pbgear.NominationPools_UnbondCall{
		MemberAccount:   to_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField)),
		UnbondingPoints: to_NominationPools_CompactString(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_UpdateRolesCall(fields []*registry.DecodedField) *pbgear.NominationPools_UpdateRolesCall {
	return &pbgear.NominationPools_UpdateRolesCall{
		PoolId:       fields[0].Value.(uint32),
		NewRoot:      to_NominationPools_NewRoot(fields[1].Value.([]*registry.DecodedField)),
		NewNominator: to_NominationPools_NewNominator(fields[2].Value.([]*registry.DecodedField)),
		NewBouncer:   to_NominationPools_NewBouncer(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_NominationPools_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.NominationPools_WithdrawUnbondedCall {
	return &pbgear.NominationPools_WithdrawUnbondedCall{
		MemberAccount:    to_NominationPools_MemberAccount(fields[0].Value.([]*registry.DecodedField)),
		NumSlashingSpans: fields[1].Value.(uint32),
	}
}

func to_Preimage_EnsureUpdatedCall(fields []*registry.DecodedField) *pbgear.Preimage_EnsureUpdatedCall {
	return &pbgear.Preimage_EnsureUpdatedCall{
		Hashes: to_repeated_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Preimage_NotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_NotePreimageCall {
	return &pbgear.Preimage_NotePreimageCall{
		Bytes: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Preimage_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
	return &pbgear.Preimage_PrimitiveTypesH256{
		Hash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Preimage_RequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_RequestPreimageCall {
	return &pbgear.Preimage_RequestPreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Preimage_UnnotePreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnnotePreimageCall {
	return &pbgear.Preimage_UnnotePreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Preimage_UnrequestPreimageCall(fields []*registry.DecodedField) *pbgear.Preimage_UnrequestPreimageCall {
	return &pbgear.Preimage_UnrequestPreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_AddProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_AddProxyCall {
	return &pbgear.Proxy_AddProxyCall{
		Delegate:  to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
		ProxyType: to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
		Delay:     fields[2].Value.(uint32),
	}
}

func to_Proxy_Address20(fields []*registry.DecodedField) *pbgear.Proxy_Address20 {
	return &pbgear.Proxy_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Address32(fields []*registry.DecodedField) *pbgear.Proxy_Address32 {
	return &pbgear.Proxy_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_AnnounceCall(fields []*registry.DecodedField) *pbgear.Proxy_AnnounceCall {
	return &pbgear.Proxy_AnnounceCall{
		Real:     to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Any(fields []*registry.DecodedField) *pbgear.Proxy_Any {
	return &pbgear.Proxy_Any{}
}
func to_Proxy_CancelProxy(fields []*registry.DecodedField) *pbgear.Proxy_CancelProxy {
	return &pbgear.Proxy_CancelProxy{}
}
func to_Proxy_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
	return &pbgear.Proxy_CompactTupleNull{
		Value: to_Proxy_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_CompactUint32(fields []*registry.DecodedField) *pbgear.Proxy_CompactUint32 {
	return &pbgear.Proxy_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Proxy_CreatePureCall(fields []*registry.DecodedField) *pbgear.Proxy_CreatePureCall {
	return &pbgear.Proxy_CreatePureCall{
		ProxyType: to_Proxy_ProxyType(fields[0].Value.([]*registry.DecodedField)),
		Delay:     fields[1].Value.(uint32),
		Index:     fields[2].Value.(uint32),
	}
}

func to_Proxy_Delegate(fields []*registry.DecodedField) *pbgear.Proxy_Delegate {
	return &pbgear.Proxy_Delegate{
		Value: to_oneof_Proxy_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Proxy_Value(in *registry.DecodedField) *pbgear.Proxy_Value {
	return nil
}

func to_Proxy_ForceProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ForceProxyType {
	return &pbgear.Proxy_ForceProxyType{
		Value: to_oneof_Proxy_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Governance(fields []*registry.DecodedField) *pbgear.Proxy_Governance {
	return &pbgear.Proxy_Governance{}
}
func to_Proxy_Id(fields []*registry.DecodedField) *pbgear.Proxy_Id {
	return &pbgear.Proxy_Id{
		Value_0: to_Proxy_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_IdentityJudgement(fields []*registry.DecodedField) *pbgear.Proxy_IdentityJudgement {
	return &pbgear.Proxy_IdentityJudgement{}
}
func to_Proxy_Index(fields []*registry.DecodedField) *pbgear.Proxy_Index {
	return &pbgear.Proxy_Index{
		Value_0: to_Proxy_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_KillPureCall(fields []*registry.DecodedField) *pbgear.Proxy_KillPureCall {
	return &pbgear.Proxy_KillPureCall{
		Spawner:   to_Proxy_Spawner(fields[0].Value.([]*registry.DecodedField)),
		ProxyType: to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
		Index:     fields[2].Value.(uint32),
		Height:    to_Proxy_CompactUint32(fields[3].Value.([]*registry.DecodedField)),
		ExtIndex:  to_Proxy_CompactUint32(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_NonTransfer(fields []*registry.DecodedField) *pbgear.Proxy_NonTransfer {
	return &pbgear.Proxy_NonTransfer{}
}
func to_Proxy_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
	return &pbgear.Proxy_PrimitiveTypesH256{
		CallHash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_ProxyAnnouncedCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyAnnouncedCall {
	return &pbgear.Proxy_ProxyAnnouncedCall{
		Delegate:       to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
		Real:           to_Proxy_Real(fields[1].Value.([]*registry.DecodedField)),
		ForceProxyType: to_optional_Proxy_ForceProxyType(fields[2].Value.([]*registry.DecodedField)),
		Call:           to_oneof_Proxy_Call(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Proxy_Call(in *registry.DecodedField) *pbgear.Proxy_Call {
	return nil
}

func to_Proxy_ProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_ProxyCall {
	return &pbgear.Proxy_ProxyCall{
		Real:           to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
		ForceProxyType: to_optional_Proxy_ForceProxyType(fields[1].Value.([]*registry.DecodedField)),
		Call:           to_oneof_Proxy_Call(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_ProxyType(fields []*registry.DecodedField) *pbgear.Proxy_ProxyType {
	return &pbgear.Proxy_ProxyType{
		Value: to_oneof_Proxy_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Raw(fields []*registry.DecodedField) *pbgear.Proxy_Raw {
	return &pbgear.Proxy_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Real(fields []*registry.DecodedField) *pbgear.Proxy_Real {
	return &pbgear.Proxy_Real{
		Value: to_oneof_Proxy_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_RejectAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RejectAnnouncementCall {
	return &pbgear.Proxy_RejectAnnouncementCall{
		Delegate: to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_RemoveAnnouncementCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveAnnouncementCall {
	return &pbgear.Proxy_RemoveAnnouncementCall{
		Real:     to_Proxy_Real(fields[0].Value.([]*registry.DecodedField)),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_RemoveProxiesCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxiesCall {
	return &pbgear.Proxy_RemoveProxiesCall{}
}
func to_Proxy_RemoveProxyCall(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxyCall {
	return &pbgear.Proxy_RemoveProxyCall{
		Delegate:  to_Proxy_Delegate(fields[0].Value.([]*registry.DecodedField)),
		ProxyType: to_Proxy_ProxyType(fields[1].Value.([]*registry.DecodedField)),
		Delay:     fields[2].Value.(uint32),
	}
}

func to_Proxy_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Proxy_SpCoreCryptoAccountId32 {
	return &pbgear.Proxy_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Spawner(fields []*registry.DecodedField) *pbgear.Proxy_Spawner {
	return &pbgear.Proxy_Spawner{
		Value: to_oneof_Proxy_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Proxy_Staking(fields []*registry.DecodedField) *pbgear.Proxy_Staking {
	return &pbgear.Proxy_Staking{}
}
func to_Proxy_TupleNull(fields []*registry.DecodedField) *pbgear.Proxy_TupleNull {
	return &pbgear.Proxy_TupleNull{}
}
func to_Referenda_After(fields []*registry.DecodedField) *pbgear.Referenda_After {
	return &pbgear.Referenda_After{
		Value_0: fields[0].Value.(uint32),
	}
}

func to_Referenda_At(fields []*registry.DecodedField) *pbgear.Referenda_At {
	return &pbgear.Referenda_At{
		Value_0: fields[0].Value.(uint32),
	}
}

func to_Referenda_BoundedCollectionsBoundedVecBoundedVec(fields []*registry.DecodedField) *pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_CancelCall(fields []*registry.DecodedField) *pbgear.Referenda_CancelCall {
	return &pbgear.Referenda_CancelCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_EnactmentMoment(fields []*registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
	return &pbgear.Referenda_EnactmentMoment{
		Value: to_oneof_Referenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Referenda_Value(in *registry.DecodedField) *pbgear.Referenda_Value {
	return nil
}

func to_Referenda_Inline(fields []*registry.DecodedField) *pbgear.Referenda_Inline {
	return &pbgear.Referenda_Inline{
		Value_0: to_Referenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_KillCall(fields []*registry.DecodedField) *pbgear.Referenda_KillCall {
	return &pbgear.Referenda_KillCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_Legacy(fields []*registry.DecodedField) *pbgear.Referenda_Legacy {
	return &pbgear.Referenda_Legacy{
		Hash: to_Referenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_Lookup(fields []*registry.DecodedField) *pbgear.Referenda_Lookup {
	return &pbgear.Referenda_Lookup{
		Hash: to_Referenda_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
		Len:  fields[1].Value.(uint32),
	}
}

func to_Referenda_None(fields []*registry.DecodedField) *pbgear.Referenda_None {
	return &pbgear.Referenda_None{}
}
func to_Referenda_NudgeReferendumCall(fields []*registry.DecodedField) *pbgear.Referenda_NudgeReferendumCall {
	return &pbgear.Referenda_NudgeReferendumCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_OneFewerDecidingCall(fields []*registry.DecodedField) *pbgear.Referenda_OneFewerDecidingCall {
	return &pbgear.Referenda_OneFewerDecidingCall{
		Track: fields[0].Value.(uint32),
	}
}

func to_Referenda_Origins(fields []*registry.DecodedField) *pbgear.Referenda_Origins {
	return &pbgear.Referenda_Origins{
		Value_0: to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_PlaceDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_PlaceDecisionDepositCall {
	return &pbgear.Referenda_PlaceDecisionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
	return &pbgear.Referenda_PrimitiveTypesH256{
		Hash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_Proposal(fields []*registry.DecodedField) *pbgear.Referenda_Proposal {
	return &pbgear.Referenda_Proposal{
		Value: to_oneof_Referenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_ProposalOrigin(fields []*registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
	return &pbgear.Referenda_ProposalOrigin{
		Value: to_oneof_Referenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_RefundDecisionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundDecisionDepositCall {
	return &pbgear.Referenda_RefundDecisionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_RefundSubmissionDepositCall(fields []*registry.DecodedField) *pbgear.Referenda_RefundSubmissionDepositCall {
	return &pbgear.Referenda_RefundSubmissionDepositCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Referenda_Root(fields []*registry.DecodedField) *pbgear.Referenda_Root {
	return &pbgear.Referenda_Root{}
}
func to_Referenda_SetMetadataCall(fields []*registry.DecodedField) *pbgear.Referenda_SetMetadataCall {
	return &pbgear.Referenda_SetMetadataCall{
		Index:     fields[0].Value.(uint32),
		MaybeHash: to_optional_Referenda_PrimitiveTypesH256(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_Signed(fields []*registry.DecodedField) *pbgear.Referenda_Signed {
	return &pbgear.Referenda_Signed{
		Value_0: to_Referenda_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Referenda_SpCoreCryptoAccountId32 {
	return &pbgear.Referenda_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_SubmitCall(fields []*registry.DecodedField) *pbgear.Referenda_SubmitCall {
	return &pbgear.Referenda_SubmitCall{
		ProposalOrigin:  to_Referenda_ProposalOrigin(fields[0].Value.([]*registry.DecodedField)),
		Proposal:        to_Referenda_Proposal(fields[1].Value.([]*registry.DecodedField)),
		EnactmentMoment: to_Referenda_EnactmentMoment(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_System(fields []*registry.DecodedField) *pbgear.Referenda_System {
	return &pbgear.Referenda_System{
		Value_0: to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_Value0(fields []*registry.DecodedField) *pbgear.Referenda_Value0 {
	return &pbgear.Referenda_Value0{
		Value: to_oneof_Referenda_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Referenda_Void(fields []*registry.DecodedField) *pbgear.Referenda_Void {
	return &pbgear.Referenda_Void{
		Value_0: to_Referenda_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Scheduler_CancelCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelCall {
	return &pbgear.Scheduler_CancelCall{
		When:  fields[0].Value.(uint32),
		Index: fields[1].Value.(uint32),
	}
}

func to_Scheduler_CancelNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_CancelNamedCall {
	return &pbgear.Scheduler_CancelNamedCall{
		Id: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Scheduler_ScheduleAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleAfterCall {
	return &pbgear.Scheduler_ScheduleAfterCall{
		After:         fields[0].Value.(uint32),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField)),
		Priority:      fields[2].Value.(uint32),
		Call:          to_oneof_Scheduler_Call(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Scheduler_Call(in *registry.DecodedField) *pbgear.Scheduler_Call {
	return nil
}

func to_Scheduler_ScheduleCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleCall {
	return &pbgear.Scheduler_ScheduleCall{
		When:          fields[0].Value.(uint32),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[1].Value.([]*registry.DecodedField)),
		Priority:      fields[2].Value.(uint32),
		Call:          to_oneof_Scheduler_Call(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_Scheduler_ScheduleNamedAfterCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedAfterCall {
	return &pbgear.Scheduler_ScheduleNamedAfterCall{
		Id:            to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		After:         fields[1].Value.(uint32),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField)),
		Priority:      fields[3].Value.(uint32),
		Call:          to_oneof_Scheduler_Call(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_Scheduler_ScheduleNamedCall(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedCall {
	return &pbgear.Scheduler_ScheduleNamedCall{
		Id:            to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		When:          fields[1].Value.(uint32),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[2].Value.([]*registry.DecodedField)),
		Priority:      fields[3].Value.(uint32),
		Call:          to_oneof_Scheduler_Call(fields[4].Value.([]*registry.DecodedField)),
	}
}

func to_Scheduler_TupleUint32Uint32(fields []*registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
	return &pbgear.Scheduler_TupleUint32Uint32{
		Value_0: fields[0].Value.(uint32),
		Value_1: fields[1].Value.(uint32),
	}
}

func to_Session_PalletImOnlineSr25519AppSr25519Public(fields []*registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
	return &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{
		ImOnline: to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_PurgeKeysCall(fields []*registry.DecodedField) *pbgear.Session_PurgeKeysCall {
	return &pbgear.Session_PurgeKeysCall{}
}
func to_Session_SetKeysCall(fields []*registry.DecodedField) *pbgear.Session_SetKeysCall {
	return &pbgear.Session_SetKeysCall{
		Keys:  to_Session_VaraRuntimeSessionKeys(fields[0].Value.([]*registry.DecodedField)),
		Proof: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Session_SpAuthorityDiscoveryAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
	return &pbgear.Session_SpAuthorityDiscoveryAppPublic{
		AuthorityDiscovery: to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_SpConsensusBabeAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
	return &pbgear.Session_SpConsensusBabeAppPublic{
		Babe: to_Session_SpCoreSr25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_SpConsensusGrandpaAppPublic(fields []*registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
	return &pbgear.Session_SpConsensusGrandpaAppPublic{
		Grandpa: to_Session_SpCoreEd25519Public(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_SpCoreEd25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
	return &pbgear.Session_SpCoreEd25519Public{
		Grandpa: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_SpCoreSr25519Public(fields []*registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
	return &pbgear.Session_SpCoreSr25519Public{
		Babe: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Session_VaraRuntimeSessionKeys(fields []*registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
	return &pbgear.Session_VaraRuntimeSessionKeys{
		Babe:               to_Session_SpConsensusBabeAppPublic(fields[0].Value.([]*registry.DecodedField)),
		Grandpa:            to_Session_SpConsensusGrandpaAppPublic(fields[1].Value.([]*registry.DecodedField)),
		ImOnline:           to_Session_PalletImOnlineSr25519AppSr25519Public(fields[2].Value.([]*registry.DecodedField)),
		AuthorityDiscovery: to_Session_SpAuthorityDiscoveryAppPublic(fields[3].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_Address20(fields []*registry.DecodedField) *pbgear.StakingRewards_Address20 {
	return &pbgear.StakingRewards_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_Address32(fields []*registry.DecodedField) *pbgear.StakingRewards_Address32 {
	return &pbgear.StakingRewards_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_AlignSupplyCall(fields []*registry.DecodedField) *pbgear.StakingRewards_AlignSupplyCall {
	return &pbgear.StakingRewards_AlignSupplyCall{
		Target: fields[0].Value.(string),
	}
}

func to_StakingRewards_CompactTupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
	return &pbgear.StakingRewards_CompactTupleNull{
		Value: to_StakingRewards_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_ForceRefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_ForceRefillCall {
	return &pbgear.StakingRewards_ForceRefillCall{
		From:  to_StakingRewards_From(fields[0].Value.([]*registry.DecodedField)),
		Value: fields[1].Value.(string),
	}
}

func to_StakingRewards_From(fields []*registry.DecodedField) *pbgear.StakingRewards_From {
	return &pbgear.StakingRewards_From{
		Value: to_oneof_StakingRewards_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_StakingRewards_Value(in *registry.DecodedField) *pbgear.StakingRewards_Value {
	return nil
}

func to_StakingRewards_Id(fields []*registry.DecodedField) *pbgear.StakingRewards_Id {
	return &pbgear.StakingRewards_Id{
		Value_0: to_StakingRewards_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_Index(fields []*registry.DecodedField) *pbgear.StakingRewards_Index {
	return &pbgear.StakingRewards_Index{
		Value_0: to_StakingRewards_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_Raw(fields []*registry.DecodedField) *pbgear.StakingRewards_Raw {
	return &pbgear.StakingRewards_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_RefillCall(fields []*registry.DecodedField) *pbgear.StakingRewards_RefillCall {
	return &pbgear.StakingRewards_RefillCall{
		Value: fields[0].Value.(string),
	}
}

func to_StakingRewards_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.StakingRewards_SpCoreCryptoAccountId32 {
	return &pbgear.StakingRewards_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_To(fields []*registry.DecodedField) *pbgear.StakingRewards_To {
	return &pbgear.StakingRewards_To{
		Value: to_oneof_StakingRewards_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_StakingRewards_TupleNull(fields []*registry.DecodedField) *pbgear.StakingRewards_TupleNull {
	return &pbgear.StakingRewards_TupleNull{}
}
func to_StakingRewards_WithdrawCall(fields []*registry.DecodedField) *pbgear.StakingRewards_WithdrawCall {
	return &pbgear.StakingRewards_WithdrawCall{
		To:    to_StakingRewards_To(fields[0].Value.([]*registry.DecodedField)),
		Value: fields[1].Value.(string),
	}
}

func to_Staking_Account(fields []*registry.DecodedField) *pbgear.Staking_Account {
	return &pbgear.Staking_Account{
		Value_0: to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Address20(fields []*registry.DecodedField) *pbgear.Staking_Address20 {
	return &pbgear.Staking_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Address32(fields []*registry.DecodedField) *pbgear.Staking_Address32 {
	return &pbgear.Staking_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_BondCall(fields []*registry.DecodedField) *pbgear.Staking_BondCall {
	return &pbgear.Staking_BondCall{
		Value: to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Payee: to_Staking_Payee(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_BondExtraCall(fields []*registry.DecodedField) *pbgear.Staking_BondExtraCall {
	return &pbgear.Staking_BondExtraCall{
		MaxAdditional: to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_CancelDeferredSlashCall(fields []*registry.DecodedField) *pbgear.Staking_CancelDeferredSlashCall {
	return &pbgear.Staking_CancelDeferredSlashCall{
		Era:          fields[0].Value.(uint32),
		SlashIndices: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_ChillCall(fields []*registry.DecodedField) *pbgear.Staking_ChillCall {
	return &pbgear.Staking_ChillCall{}
}
func to_Staking_ChillOtherCall(fields []*registry.DecodedField) *pbgear.Staking_ChillOtherCall {
	return &pbgear.Staking_ChillOtherCall{
		Controller: to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_ChillThreshold(fields []*registry.DecodedField) *pbgear.Staking_ChillThreshold {
	return &pbgear.Staking_ChillThreshold{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Staking_Value(in *registry.DecodedField) *pbgear.Staking_Value {
	return nil
}

func to_Staking_CompactSpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
	return &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{
		Value: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_CompactString(fields []*registry.DecodedField) *pbgear.Staking_CompactString {
	return &pbgear.Staking_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_Staking_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Staking_CompactTupleNull {
	return &pbgear.Staking_CompactTupleNull{
		Value: to_Staking_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_CompactUint32(fields []*registry.DecodedField) *pbgear.Staking_CompactUint32 {
	return &pbgear.Staking_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Staking_Controller(fields []*registry.DecodedField) *pbgear.Staking_Controller {
	return &pbgear.Staking_Controller{}
}
func to_Staking_ForceApplyMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_ForceApplyMinCommissionCall {
	return &pbgear.Staking_ForceApplyMinCommissionCall{
		ValidatorStash: to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_ForceNewEraAlwaysCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraAlwaysCall {
	return &pbgear.Staking_ForceNewEraAlwaysCall{}
}
func to_Staking_ForceNewEraCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraCall {
	return &pbgear.Staking_ForceNewEraCall{}
}
func to_Staking_ForceNoErasCall(fields []*registry.DecodedField) *pbgear.Staking_ForceNoErasCall {
	return &pbgear.Staking_ForceNoErasCall{}
}
func to_Staking_ForceUnstakeCall(fields []*registry.DecodedField) *pbgear.Staking_ForceUnstakeCall {
	return &pbgear.Staking_ForceUnstakeCall{
		Stash:            to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		NumSlashingSpans: fields[1].Value.(uint32),
	}
}

func to_Staking_Id(fields []*registry.DecodedField) *pbgear.Staking_Id {
	return &pbgear.Staking_Id{
		Value_0: to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_IncreaseValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_IncreaseValidatorCountCall {
	return &pbgear.Staking_IncreaseValidatorCountCall{
		Additional: to_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Index(fields []*registry.DecodedField) *pbgear.Staking_Index {
	return &pbgear.Staking_Index{
		Value_0: to_Staking_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_KickCall(fields []*registry.DecodedField) *pbgear.Staking_KickCall {
	return &pbgear.Staking_KickCall{
		Who: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_MaxNominatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
	return &pbgear.Staking_MaxNominatorCount{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_MaxValidatorCount(fields []*registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
	return &pbgear.Staking_MaxValidatorCount{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_MinCommission(fields []*registry.DecodedField) *pbgear.Staking_MinCommission {
	return &pbgear.Staking_MinCommission{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_MinNominatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinNominatorBond {
	return &pbgear.Staking_MinNominatorBond{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_MinValidatorBond(fields []*registry.DecodedField) *pbgear.Staking_MinValidatorBond {
	return &pbgear.Staking_MinValidatorBond{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_NominateCall(fields []*registry.DecodedField) *pbgear.Staking_NominateCall {
	return &pbgear.Staking_NominateCall{
		Targets: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_None(fields []*registry.DecodedField) *pbgear.Staking_None {
	return &pbgear.Staking_None{}
}
func to_Staking_Noop(fields []*registry.DecodedField) *pbgear.Staking_Noop {
	return &pbgear.Staking_Noop{}
}
func to_Staking_PalletStakingValidatorPrefs(fields []*registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
	return &pbgear.Staking_PalletStakingValidatorPrefs{
		Commission: to_Staking_CompactSpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
		Blocked:    fields[1].Value.(bool),
	}
}

func to_Staking_Payee(fields []*registry.DecodedField) *pbgear.Staking_Payee {
	return &pbgear.Staking_Payee{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_PayoutStakersCall(fields []*registry.DecodedField) *pbgear.Staking_PayoutStakersCall {
	return &pbgear.Staking_PayoutStakersCall{
		ValidatorStash: to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		Era:            fields[1].Value.(uint32),
	}
}

func to_Staking_Raw(fields []*registry.DecodedField) *pbgear.Staking_Raw {
	return &pbgear.Staking_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_ReapStashCall(fields []*registry.DecodedField) *pbgear.Staking_ReapStashCall {
	return &pbgear.Staking_ReapStashCall{
		Stash:            to_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
		NumSlashingSpans: fields[1].Value.(uint32),
	}
}

func to_Staking_RebondCall(fields []*registry.DecodedField) *pbgear.Staking_RebondCall {
	return &pbgear.Staking_RebondCall{
		Value: to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Remove(fields []*registry.DecodedField) *pbgear.Staking_Remove {
	return &pbgear.Staking_Remove{}
}
func to_Staking_ScaleValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_ScaleValidatorCountCall {
	return &pbgear.Staking_ScaleValidatorCountCall{
		Factor: to_Staking_SpArithmeticPerThingsPercent(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Set(fields []*registry.DecodedField) *pbgear.Staking_Set {
	return &pbgear.Staking_Set{
		Value_0: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SetControllerCall(fields []*registry.DecodedField) *pbgear.Staking_SetControllerCall {
	return &pbgear.Staking_SetControllerCall{}
}
func to_Staking_SetInvulnerablesCall(fields []*registry.DecodedField) *pbgear.Staking_SetInvulnerablesCall {
	return &pbgear.Staking_SetInvulnerablesCall{
		Invulnerables: to_repeated_Staking_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SetMinCommissionCall(fields []*registry.DecodedField) *pbgear.Staking_SetMinCommissionCall {
	return &pbgear.Staking_SetMinCommissionCall{
		New: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SetPayeeCall(fields []*registry.DecodedField) *pbgear.Staking_SetPayeeCall {
	return &pbgear.Staking_SetPayeeCall{
		Payee: to_Staking_Payee(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SetStakingConfigsCall(fields []*registry.DecodedField) *pbgear.Staking_SetStakingConfigsCall {
	return &pbgear.Staking_SetStakingConfigsCall{
		MinNominatorBond:  to_Staking_MinNominatorBond(fields[0].Value.([]*registry.DecodedField)),
		MinValidatorBond:  to_Staking_MinValidatorBond(fields[1].Value.([]*registry.DecodedField)),
		MaxNominatorCount: to_Staking_MaxNominatorCount(fields[2].Value.([]*registry.DecodedField)),
		MaxValidatorCount: to_Staking_MaxValidatorCount(fields[3].Value.([]*registry.DecodedField)),
		ChillThreshold:    to_Staking_ChillThreshold(fields[4].Value.([]*registry.DecodedField)),
		MinCommission:     to_Staking_MinCommission(fields[5].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SetValidatorCountCall(fields []*registry.DecodedField) *pbgear.Staking_SetValidatorCountCall {
	return &pbgear.Staking_SetValidatorCountCall{
		New: to_Staking_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SpArithmeticPerThingsPerbill(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
	return &pbgear.Staking_SpArithmeticPerThingsPerbill{
		Value: fields[0].Value.(uint32),
	}
}

func to_Staking_SpArithmeticPerThingsPercent(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
	return &pbgear.Staking_SpArithmeticPerThingsPercent{
		Factor: fields[0].Value.(uint32),
	}
}

func to_Staking_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
	return &pbgear.Staking_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_SpRuntimeMultiaddressMultiAddress(fields []*registry.DecodedField) *pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
	return &pbgear.Staking_SpRuntimeMultiaddressMultiAddress{
		Targets: to_Staking_Targets(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_Staked(fields []*registry.DecodedField) *pbgear.Staking_Staked {
	return &pbgear.Staking_Staked{}
}
func to_Staking_Stash(fields []*registry.DecodedField) *pbgear.Staking_Stash {
	return &pbgear.Staking_Stash{}
}
func to_Staking_Targets(fields []*registry.DecodedField) *pbgear.Staking_Targets {
	return &pbgear.Staking_Targets{
		Value: to_oneof_Staking_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_TupleNull(fields []*registry.DecodedField) *pbgear.Staking_TupleNull {
	return &pbgear.Staking_TupleNull{}
}
func to_Staking_UnbondCall(fields []*registry.DecodedField) *pbgear.Staking_UnbondCall {
	return &pbgear.Staking_UnbondCall{
		Value: to_Staking_CompactString(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_ValidateCall(fields []*registry.DecodedField) *pbgear.Staking_ValidateCall {
	return &pbgear.Staking_ValidateCall{
		Prefs: to_Staking_PalletStakingValidatorPrefs(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Staking_WithdrawUnbondedCall(fields []*registry.DecodedField) *pbgear.Staking_WithdrawUnbondedCall {
	return &pbgear.Staking_WithdrawUnbondedCall{
		NumSlashingSpans: fields[0].Value.(uint32),
	}
}

func to_System_KillPrefixCall(fields []*registry.DecodedField) *pbgear.System_KillPrefixCall {
	return &pbgear.System_KillPrefixCall{
		Prefix:  to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Subkeys: fields[1].Value.(uint32),
	}
}

func to_System_KillStorageCall(fields []*registry.DecodedField) *pbgear.System_KillStorageCall {
	return &pbgear.System_KillStorageCall{
		Keys: to_repeated_System_SystemKeysList(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_RemarkCall(fields []*registry.DecodedField) *pbgear.System_RemarkCall {
	return &pbgear.System_RemarkCall{
		Remark: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_RemarkWithEventCall(fields []*registry.DecodedField) *pbgear.System_RemarkWithEventCall {
	return &pbgear.System_RemarkWithEventCall{
		Remark: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_SetCodeCall(fields []*registry.DecodedField) *pbgear.System_SetCodeCall {
	return &pbgear.System_SetCodeCall{
		Code: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_SetCodeWithoutChecksCall(fields []*registry.DecodedField) *pbgear.System_SetCodeWithoutChecksCall {
	return &pbgear.System_SetCodeWithoutChecksCall{
		Code: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_SetHeapPagesCall(fields []*registry.DecodedField) *pbgear.System_SetHeapPagesCall {
	return &pbgear.System_SetHeapPagesCall{
		Pages: fields[0].Value.(uint64),
	}
}

func to_System_SetStorageCall(fields []*registry.DecodedField) *pbgear.System_SetStorageCall {
	return &pbgear.System_SetStorageCall{
		Items: to_repeated_System_TupleSystemItemsListSystemItemsList(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_SystemKeysList(fields []*registry.DecodedField) *pbgear.System_SystemKeysList {
	return &pbgear.System_SystemKeysList{
		Keys: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_System_TupleSystemItemsListSystemItemsList(fields []*registry.DecodedField) *pbgear.System_TupleSystemItemsListSystemItemsList {
	return &pbgear.System_TupleSystemItemsListSystemItemsList{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
		Value_1: to_repeated_uint32(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Timestamp_CompactUint64(fields []*registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
	return &pbgear.Timestamp_CompactUint64{
		Value: fields[0].Value.(uint64),
	}
}

func to_Timestamp_SetCall(fields []*registry.DecodedField) *pbgear.Timestamp_SetCall {
	return &pbgear.Timestamp_SetCall{
		Now: to_Timestamp_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_Address20(fields []*registry.DecodedField) *pbgear.Treasury_Address20 {
	return &pbgear.Treasury_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_Address32(fields []*registry.DecodedField) *pbgear.Treasury_Address32 {
	return &pbgear.Treasury_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_ApproveProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_ApproveProposalCall {
	return &pbgear.Treasury_ApproveProposalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_Beneficiary(fields []*registry.DecodedField) *pbgear.Treasury_Beneficiary {
	return &pbgear.Treasury_Beneficiary{
		Value: to_oneof_Treasury_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Treasury_Value(in *registry.DecodedField) *pbgear.Treasury_Value {
	return nil
}

func to_Treasury_CheckStatusCall(fields []*registry.DecodedField) *pbgear.Treasury_CheckStatusCall {
	return &pbgear.Treasury_CheckStatusCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Treasury_CompactString(fields []*registry.DecodedField) *pbgear.Treasury_CompactString {
	return &pbgear.Treasury_CompactString{
		Value: fields[0].Value.(string),
	}
}

func to_Treasury_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
	return &pbgear.Treasury_CompactTupleNull{
		Value: to_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_CompactUint32(fields []*registry.DecodedField) *pbgear.Treasury_CompactUint32 {
	return &pbgear.Treasury_CompactUint32{
		Value: fields[0].Value.(uint32),
	}
}

func to_Treasury_Id(fields []*registry.DecodedField) *pbgear.Treasury_Id {
	return &pbgear.Treasury_Id{
		Value_0: to_Treasury_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_Index(fields []*registry.DecodedField) *pbgear.Treasury_Index {
	return &pbgear.Treasury_Index{
		Value_0: to_Treasury_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_PayoutCall(fields []*registry.DecodedField) *pbgear.Treasury_PayoutCall {
	return &pbgear.Treasury_PayoutCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Treasury_ProposeSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_ProposeSpendCall {
	return &pbgear.Treasury_ProposeSpendCall{
		Value:       to_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Beneficiary: to_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_Raw(fields []*registry.DecodedField) *pbgear.Treasury_Raw {
	return &pbgear.Treasury_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_RejectProposalCall(fields []*registry.DecodedField) *pbgear.Treasury_RejectProposalCall {
	return &pbgear.Treasury_RejectProposalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_RemoveApprovalCall(fields []*registry.DecodedField) *pbgear.Treasury_RemoveApprovalCall {
	return &pbgear.Treasury_RemoveApprovalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
	return &pbgear.Treasury_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_SpendCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendCall {
	return &pbgear.Treasury_SpendCall{
		AssetKind:   to_Treasury_TupleNull(fields[0].Value.([]*registry.DecodedField)),
		Amount:      to_Treasury_CompactString(fields[1].Value.([]*registry.DecodedField)),
		Beneficiary: to_Treasury_SpCoreCryptoAccountId32(fields[2].Value.([]*registry.DecodedField)),
		ValidFrom:   fields[3].Value.(uint32),
	}
}

func to_Treasury_SpendLocalCall(fields []*registry.DecodedField) *pbgear.Treasury_SpendLocalCall {
	return &pbgear.Treasury_SpendLocalCall{
		Amount:      to_Treasury_CompactString(fields[0].Value.([]*registry.DecodedField)),
		Beneficiary: to_Treasury_Beneficiary(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Treasury_TupleNull(fields []*registry.DecodedField) *pbgear.Treasury_TupleNull {
	return &pbgear.Treasury_TupleNull{}
}
func to_Treasury_VoidSpendCall(fields []*registry.DecodedField) *pbgear.Treasury_VoidSpendCall {
	return &pbgear.Treasury_VoidSpendCall{
		Index: fields[0].Value.(uint32),
	}
}

func to_Utility_AsDerivativeCall(fields []*registry.DecodedField) *pbgear.Utility_AsDerivativeCall {
	return &pbgear.Utility_AsDerivativeCall{
		Index: fields[0].Value.(uint32),
		Call:  to_oneof_Utility_Call(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Utility_Call(in *registry.DecodedField) *pbgear.Utility_Call {
	return nil
}

func to_Utility_AsOrigin(fields []*registry.DecodedField) *pbgear.Utility_AsOrigin {
	return &pbgear.Utility_AsOrigin{
		Value: to_oneof_Utility_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Utility_Value(in *registry.DecodedField) *pbgear.Utility_Value {
	return nil
}

func to_Utility_BatchAllCall(fields []*registry.DecodedField) *pbgear.Utility_BatchAllCall {
	return &pbgear.Utility_BatchAllCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_BatchCall(fields []*registry.DecodedField) *pbgear.Utility_BatchCall {
	return &pbgear.Utility_BatchCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_CompactUint64(fields []*registry.DecodedField) *pbgear.Utility_CompactUint64 {
	return &pbgear.Utility_CompactUint64{
		Value: fields[0].Value.(uint64),
	}
}

func to_Utility_DispatchAsCall(fields []*registry.DecodedField) *pbgear.Utility_DispatchAsCall {
	return &pbgear.Utility_DispatchAsCall{
		AsOrigin: to_Utility_AsOrigin(fields[0].Value.([]*registry.DecodedField)),
		Call:     to_oneof_Utility_Call(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_ForceBatchCall(fields []*registry.DecodedField) *pbgear.Utility_ForceBatchCall {
	return &pbgear.Utility_ForceBatchCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_None(fields []*registry.DecodedField) *pbgear.Utility_None {
	return &pbgear.Utility_None{}
}
func to_Utility_Origins(fields []*registry.DecodedField) *pbgear.Utility_Origins {
	return &pbgear.Utility_Origins{
		Value_0: to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_Root(fields []*registry.DecodedField) *pbgear.Utility_Root {
	return &pbgear.Utility_Root{}
}
func to_Utility_Signed(fields []*registry.DecodedField) *pbgear.Utility_Signed {
	return &pbgear.Utility_Signed{
		Value_0: to_Utility_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Utility_SpCoreCryptoAccountId32 {
	return &pbgear.Utility_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
	return &pbgear.Utility_SpWeightsWeightV2Weight{
		RefTime:   to_Utility_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
		ProofSize: to_Utility_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_System(fields []*registry.DecodedField) *pbgear.Utility_System {
	return &pbgear.Utility_System{
		Value_0: to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_Value0(fields []*registry.DecodedField) *pbgear.Utility_Value0 {
	return &pbgear.Utility_Value0{
		Value: to_oneof_Utility_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_VaraRuntimeRuntimeCall(fields []*registry.DecodedField) *pbgear.Utility_VaraRuntimeRuntimeCall {
	return &pbgear.Utility_VaraRuntimeRuntimeCall{
		Calls: to_oneof_Utility_Calls(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Utility_Calls(in *registry.DecodedField) *pbgear.Utility_Calls {
	return nil
}

func to_Utility_Void(fields []*registry.DecodedField) *pbgear.Utility_Void {
	return &pbgear.Utility_Void{
		Value_0: to_Utility_Value0(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Utility_WithWeightCall(fields []*registry.DecodedField) *pbgear.Utility_WithWeightCall {
	return &pbgear.Utility_WithWeightCall{
		Call:   to_oneof_Utility_Call(fields[0].Value.([]*registry.DecodedField)),
		Weight: to_Utility_SpWeightsWeightV2Weight(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Address20(fields []*registry.DecodedField) *pbgear.Vesting_Address20 {
	return &pbgear.Vesting_Address20{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Address32(fields []*registry.DecodedField) *pbgear.Vesting_Address32 {
	return &pbgear.Vesting_Address32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_CompactTupleNull(fields []*registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
	return &pbgear.Vesting_CompactTupleNull{
		Value: to_Vesting_TupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_ForceVestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_ForceVestedTransferCall {
	return &pbgear.Vesting_ForceVestedTransferCall{
		Source:   to_Vesting_Source(fields[0].Value.([]*registry.DecodedField)),
		Target:   to_Vesting_Target(fields[1].Value.([]*registry.DecodedField)),
		Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Id(fields []*registry.DecodedField) *pbgear.Vesting_Id {
	return &pbgear.Vesting_Id{
		Value_0: to_Vesting_SpCoreCryptoAccountId32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Index(fields []*registry.DecodedField) *pbgear.Vesting_Index {
	return &pbgear.Vesting_Index{
		Value_0: to_Vesting_CompactTupleNull(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_MergeSchedulesCall(fields []*registry.DecodedField) *pbgear.Vesting_MergeSchedulesCall {
	return &pbgear.Vesting_MergeSchedulesCall{
		Schedule_1Index: fields[0].Value.(uint32),
		Schedule_2Index: fields[1].Value.(uint32),
	}
}

func to_Vesting_PalletVestingVestingInfoVestingInfo(fields []*registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	return &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{
		Locked:        fields[0].Value.(string),
		PerBlock:      fields[1].Value.(string),
		StartingBlock: fields[2].Value.(uint32),
	}
}

func to_Vesting_Raw(fields []*registry.DecodedField) *pbgear.Vesting_Raw {
	return &pbgear.Vesting_Raw{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Source(fields []*registry.DecodedField) *pbgear.Vesting_Source {
	return &pbgear.Vesting_Source{
		Value: to_oneof_Vesting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Vesting_Value(in *registry.DecodedField) *pbgear.Vesting_Value {
	return nil
}

func to_Vesting_SpCoreCryptoAccountId32(fields []*registry.DecodedField) *pbgear.Vesting_SpCoreCryptoAccountId32 {
	return &pbgear.Vesting_SpCoreCryptoAccountId32{
		Value_0: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_Target(fields []*registry.DecodedField) *pbgear.Vesting_Target {
	return &pbgear.Vesting_Target{
		Value: to_oneof_Vesting_Value(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_TupleNull(fields []*registry.DecodedField) *pbgear.Vesting_TupleNull {
	return &pbgear.Vesting_TupleNull{}
}
func to_Vesting_VestCall(fields []*registry.DecodedField) *pbgear.Vesting_VestCall {
	return &pbgear.Vesting_VestCall{}
}
func to_Vesting_VestOtherCall(fields []*registry.DecodedField) *pbgear.Vesting_VestOtherCall {
	return &pbgear.Vesting_VestOtherCall{
		Target: to_Vesting_Target(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Vesting_VestedTransferCall(fields []*registry.DecodedField) *pbgear.Vesting_VestedTransferCall {
	return &pbgear.Vesting_VestedTransferCall{
		Target:   to_Vesting_Target(fields[0].Value.([]*registry.DecodedField)),
		Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Whitelist_CompactUint64(fields []*registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
	return &pbgear.Whitelist_CompactUint64{
		Value: fields[0].Value.(uint64),
	}
}

func to_Whitelist_DispatchWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallCall {
	return &pbgear.Whitelist_DispatchWhitelistedCallCall{
		CallHash:          to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
		CallEncodedLen:    fields[1].Value.(uint32),
		CallWeightWitness: to_Whitelist_SpWeightsWeightV2Weight(fields[2].Value.([]*registry.DecodedField)),
	}
}

func to_Whitelist_DispatchWhitelistedCallWithPreimageCall(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
	return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{
		Call: to_oneof_Whitelist_Call(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_oneof_Whitelist_Call(in *registry.DecodedField) *pbgear.Whitelist_Call {
	return nil
}

func to_Whitelist_PrimitiveTypesH256(fields []*registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
	return &pbgear.Whitelist_PrimitiveTypesH256{
		CallHash: to_repeated_uint32(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Whitelist_RemoveWhitelistedCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_RemoveWhitelistedCallCall {
	return &pbgear.Whitelist_RemoveWhitelistedCallCall{
		CallHash: to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}

func to_Whitelist_SpWeightsWeightV2Weight(fields []*registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
	return &pbgear.Whitelist_SpWeightsWeightV2Weight{
		RefTime:   to_Whitelist_CompactUint64(fields[0].Value.([]*registry.DecodedField)),
		ProofSize: to_Whitelist_CompactUint64(fields[1].Value.([]*registry.DecodedField)),
	}
}

func to_Whitelist_WhitelistCallCall(fields []*registry.DecodedField) *pbgear.Whitelist_WhitelistCallCall {
	return &pbgear.Whitelist_WhitelistCallCall{
		CallHash: to_Whitelist_PrimitiveTypesH256(fields[0].Value.([]*registry.DecodedField)),
	}
}
