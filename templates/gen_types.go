package gen_types

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func To_AllowedSlots(in registry.DecodedFields) *pbgear.AllowedSlots {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.AllowedSlots{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.AllowedSlots_PrimarySlots{
			PrimarySlots: To_PrimarySlots(fields), //Passthrough fields...
		}
	}
	return out
}

func To_BTreeSet(in registry.DecodedFields) *pbgear.BTreeSet {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BTreeSet{}
	out.Value = To_repeated_GprimitivesActorId(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_GprimitivesActorId(f registry.DecodedFields) []*pbgear.GprimitivesActorId {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.GprimitivesActorId, len(fields))
	for _, f := range fields {
		out = append(out, To_GprimitivesActorId(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Babe_BabeTrieNodesList(in registry.DecodedFields) *pbgear.Babe_BabeTrieNodesList {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_BabeTrieNodesList{}
	out.TrieNodes = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Babe_Config(in registry.DecodedFields) *pbgear.Babe_Config {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_Config{}

	switch {
	case matchFields(fields, []int64{79, 80}):
		out.Value = &pbgear.Babe_Config_V1{
			V1: To_Babe_V1(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Babe_Consensus(in registry.DecodedFields) *pbgear.Babe_Consensus {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_Consensus{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	out.Value1 = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Babe_Logs(in registry.DecodedFields) *pbgear.Babe_Logs {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_Logs{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Babe_Logs_PreRuntime{
			PreRuntime: To_Babe_PreRuntime(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Babe_Other(in registry.DecodedFields) *pbgear.Babe_Other {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_Other{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Babe_PlanConfigChangeCall(in registry.DecodedFields) *pbgear.Babe_PlanConfigChangeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_PlanConfigChangeCall{}
	out.Config = To_Babe_Config(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Babe_PreRuntime(in registry.DecodedFields) *pbgear.Babe_PreRuntime {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_PreRuntime{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	out.Value1 = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Babe_ReportEquivocationCall(in registry.DecodedFields) *pbgear.Babe_ReportEquivocationCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_ReportEquivocationCall{}
	out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(fields[0].Value.(registry.DecodedFields))

	out.KeyOwnerProof = To_SpSessionMembershipProof(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Babe_ReportEquivocationUnsignedCall(in registry.DecodedFields) *pbgear.Babe_ReportEquivocationUnsignedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
	out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(fields[0].Value.(registry.DecodedFields))

	out.KeyOwnerProof = To_SpSessionMembershipProof(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Babe_RuntimeEnvironmentUpdated(in registry.DecodedFields) *pbgear.Babe_RuntimeEnvironmentUpdated {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
	return out
}

func To_Babe_Seal(in registry.DecodedFields) *pbgear.Babe_Seal {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_Seal{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	out.Value1 = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Babe_TupleUint64Uint64(in registry.DecodedFields) *pbgear.Babe_TupleUint64Uint64 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_TupleUint64Uint64{}
	out.Value0 = To_uint64(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_uint64(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Babe_V1(in registry.DecodedFields) *pbgear.Babe_V1 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Babe_V1{}
	out.C = To_Babe_TupleUint64Uint64(fields[0].Value.(registry.DecodedFields))

	out.AllowedSlots = To_AllowedSlots(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_BagsList_Address20(in registry.DecodedFields) *pbgear.BagsList_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_BagsList_Address32(in registry.DecodedFields) *pbgear.BagsList_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_BagsList_CompactTupleNull(in types.UCompact) *pbgear.BagsList_CompactTupleNull {

	out := &pbgear.BagsList_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_BagsList_Dislocated(in registry.DecodedFields) *pbgear.BagsList_Dislocated {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_BagsList_Heavier(in registry.DecodedFields) *pbgear.BagsList_Heavier {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_BagsList_Id(in registry.DecodedFields) *pbgear.BagsList_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_BagsList_Index(in registry.DecodedFields) *pbgear.BagsList_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_Index{}
	out.Value0 = To_BagsList_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_BagsList_Lighter(in registry.DecodedFields) *pbgear.BagsList_Lighter {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_BagsList_PutInFrontOfCall(in registry.DecodedFields) *pbgear.BagsList_PutInFrontOfCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_PutInFrontOfCall{}
	out.Lighter = To_BagsList_Lighter(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_BagsList_PutInFrontOfOtherCall(in registry.DecodedFields) *pbgear.BagsList_PutInFrontOfOtherCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_PutInFrontOfOtherCall{}
	out.Heavier = To_BagsList_Heavier(fields[0].Value.(registry.DecodedFields))

	out.Lighter = To_BagsList_Lighter(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_BagsList_Raw(in registry.DecodedFields) *pbgear.BagsList_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_BagsList_RebagCall(in registry.DecodedFields) *pbgear.BagsList_RebagCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BagsList_RebagCall{}
	out.Dislocated = To_BagsList_Dislocated(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Balances_Address20(in registry.DecodedFields) *pbgear.Balances_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Balances_Address32(in registry.DecodedFields) *pbgear.Balances_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Balances_CompactTupleNull(in types.UCompact) *pbgear.Balances_CompactTupleNull {

	out := &pbgear.Balances_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Balances_Dest(in registry.DecodedFields) *pbgear.Balances_Dest {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Balances_ForceSetBalanceCall(in registry.DecodedFields) *pbgear.Balances_ForceSetBalanceCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_ForceSetBalanceCall{}
	out.Who = To_Balances_Who(fields[0].Value.(registry.DecodedFields))

	out.NewFree = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_Balances_ForceTransferCall(in registry.DecodedFields) *pbgear.Balances_ForceTransferCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_ForceTransferCall{}
	out.Source = To_Balances_Source(fields[0].Value.(registry.DecodedFields))

	out.Dest = To_Balances_Dest(fields[1].Value.(registry.DecodedFields))

	out.Value = To_CompactString(fields[2].Value.(types.UCompact))

	return out
}

func To_Balances_ForceUnreserveCall(in registry.DecodedFields) *pbgear.Balances_ForceUnreserveCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_ForceUnreserveCall{}
	out.Who = To_Balances_Who(fields[0].Value.(registry.DecodedFields))

	out.Amount = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Balances_Id(in registry.DecodedFields) *pbgear.Balances_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Balances_Index(in registry.DecodedFields) *pbgear.Balances_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_Index{}
	out.Value0 = To_Balances_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Balances_Raw(in registry.DecodedFields) *pbgear.Balances_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Balances_Source(in registry.DecodedFields) *pbgear.Balances_Source {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Balances_TransferAllCall(in registry.DecodedFields) *pbgear.Balances_TransferAllCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_TransferAllCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Balances_TransferAllowDeathCall(in registry.DecodedFields) *pbgear.Balances_TransferAllowDeathCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_TransferAllowDeathCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.(registry.DecodedFields))

	out.Value = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_Balances_TransferKeepAliveCall(in registry.DecodedFields) *pbgear.Balances_TransferKeepAliveCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_TransferKeepAliveCall{}
	out.Dest = To_Balances_Dest(fields[0].Value.(registry.DecodedFields))

	out.Value = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_Balances_UpgradeAccountsCall(in registry.DecodedFields) *pbgear.Balances_UpgradeAccountsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Balances_UpgradeAccountsCall{}
	out.Who = To_repeated_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_SpCoreCryptoAccountId32(f registry.DecodedFields) []*pbgear.SpCoreCryptoAccountId32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.SpCoreCryptoAccountId32, len(fields))
	for _, f := range fields {
		out = append(out, To_SpCoreCryptoAccountId32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Balances_Who(in registry.DecodedFields) *pbgear.Balances_Who {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_BoundedCollectionsBoundedVecBoundedVec(in registry.DecodedFields) *pbgear.BoundedCollectionsBoundedVecBoundedVec {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.BoundedCollectionsBoundedVecBoundedVec{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Bounties_AcceptCuratorCall(in registry.DecodedFields) *pbgear.Bounties_AcceptCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_AcceptCuratorCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Bounties_Address20(in registry.DecodedFields) *pbgear.Bounties_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Bounties_Address32(in registry.DecodedFields) *pbgear.Bounties_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Bounties_ApproveBountyCall(in registry.DecodedFields) *pbgear.Bounties_ApproveBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_ApproveBountyCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Bounties_AwardBountyCall(in registry.DecodedFields) *pbgear.Bounties_AwardBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_AwardBountyCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Beneficiary = To_Bounties_Beneficiary(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Bounties_Beneficiary(in registry.DecodedFields) *pbgear.Bounties_Beneficiary {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Bounties_ClaimBountyCall(in registry.DecodedFields) *pbgear.Bounties_ClaimBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_ClaimBountyCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Bounties_CloseBountyCall(in registry.DecodedFields) *pbgear.Bounties_CloseBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_CloseBountyCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Bounties_CompactTupleNull(in types.UCompact) *pbgear.Bounties_CompactTupleNull {

	out := &pbgear.Bounties_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Bounties_Curator(in registry.DecodedFields) *pbgear.Bounties_Curator {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Bounties_ExtendBountyExpiryCall(in registry.DecodedFields) *pbgear.Bounties_ExtendBountyExpiryCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_ExtendBountyExpiryCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Remark = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Bounties_Id(in registry.DecodedFields) *pbgear.Bounties_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Bounties_Index(in registry.DecodedFields) *pbgear.Bounties_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_Index{}
	out.Value0 = To_Bounties_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Bounties_ProposeBountyCall(in registry.DecodedFields) *pbgear.Bounties_ProposeBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_ProposeBountyCall{}
	out.Value = To_CompactString(fields[0].Value.(types.UCompact))

	out.Description = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Bounties_ProposeCuratorCall(in registry.DecodedFields) *pbgear.Bounties_ProposeCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_ProposeCuratorCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Curator = To_Bounties_Curator(fields[1].Value.(registry.DecodedFields))

	out.Fee = To_CompactString(fields[2].Value.(types.UCompact))

	return out
}

func To_Bounties_Raw(in registry.DecodedFields) *pbgear.Bounties_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Bounties_UnassignCuratorCall(in registry.DecodedFields) *pbgear.Bounties_UnassignCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bounties_UnassignCuratorCall{}
	out.BountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Bytes(in registry.DecodedFields) *pbgear.Bytes {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Bytes{}
	out.Remark = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ChildBounties_AcceptCuratorCall(in registry.DecodedFields) *pbgear.ChildBounties_AcceptCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_AcceptCuratorCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_ChildBounties_AddChildBountyCall(in registry.DecodedFields) *pbgear.ChildBounties_AddChildBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_AddChildBountyCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value = To_CompactString(fields[1].Value.(types.UCompact))

	out.Description = To_bytes(fields[2].Value.([]interface{}))
	return out
}

func To_ChildBounties_Address20(in registry.DecodedFields) *pbgear.ChildBounties_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ChildBounties_Address32(in registry.DecodedFields) *pbgear.ChildBounties_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ChildBounties_AwardChildBountyCall(in registry.DecodedFields) *pbgear.ChildBounties_AwardChildBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_AwardChildBountyCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	out.Beneficiary = To_ChildBounties_Beneficiary(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_ChildBounties_Beneficiary(in registry.DecodedFields) *pbgear.ChildBounties_Beneficiary {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_ChildBounties_ClaimChildBountyCall(in registry.DecodedFields) *pbgear.ChildBounties_ClaimChildBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_ClaimChildBountyCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_ChildBounties_CloseChildBountyCall(in registry.DecodedFields) *pbgear.ChildBounties_CloseChildBountyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_CloseChildBountyCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_ChildBounties_CompactTupleNull(in types.UCompact) *pbgear.ChildBounties_CompactTupleNull {

	out := &pbgear.ChildBounties_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_ChildBounties_Curator(in registry.DecodedFields) *pbgear.ChildBounties_Curator {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_ChildBounties_Id(in registry.DecodedFields) *pbgear.ChildBounties_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ChildBounties_Index(in registry.DecodedFields) *pbgear.ChildBounties_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_Index{}
	out.Value0 = To_ChildBounties_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_ChildBounties_ProposeCuratorCall(in registry.DecodedFields) *pbgear.ChildBounties_ProposeCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_ProposeCuratorCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	out.Curator = To_ChildBounties_Curator(fields[2].Value.(registry.DecodedFields))

	out.Fee = To_CompactString(fields[3].Value.(types.UCompact))

	return out
}

func To_ChildBounties_Raw(in registry.DecodedFields) *pbgear.ChildBounties_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ChildBounties_UnassignCuratorCall(in registry.DecodedFields) *pbgear.ChildBounties_UnassignCuratorCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ChildBounties_UnassignCuratorCall{}
	out.ParentBountyId = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.ChildBountyId = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_CompactSpArithmeticPerThingsPerU16(in types.UCompact) *pbgear.CompactSpArithmeticPerThingsPerU16 {

	out := &pbgear.CompactSpArithmeticPerThingsPerU16{}
	out.Value = To_SpArithmeticPerThingsPerU16(in)
	return out
}

func To_CompactSpArithmeticPerThingsPerbill(in types.UCompact) *pbgear.CompactSpArithmeticPerThingsPerbill {

	out := &pbgear.CompactSpArithmeticPerThingsPerbill{}
	out.Value = To_SpArithmeticPerThingsPerbill(in)
	return out
}

func To_CompactString(in types.UCompact) *pbgear.CompactString {

	//&pbgear.CompactString
	return nil
}

func To_CompactUint32(in types.UCompact) *pbgear.CompactUint32 {

	//&pbgear.CompactUint32
	return nil
}

func To_CompactUint64(in types.UCompact) *pbgear.CompactUint64 {

	//&pbgear.CompactUint64
	return nil
}

func To_ConvictionVoting_Address20(in registry.DecodedFields) *pbgear.ConvictionVoting_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ConvictionVoting_Address32(in registry.DecodedFields) *pbgear.ConvictionVoting_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ConvictionVoting_CompactTupleNull(in types.UCompact) *pbgear.ConvictionVoting_CompactTupleNull {

	out := &pbgear.ConvictionVoting_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_ConvictionVoting_Conviction(in registry.DecodedFields) *pbgear.ConvictionVoting_Conviction {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Conviction{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.ConvictionVoting_Conviction_None{
			None: To_ConvictionVoting_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ConvictionVoting_DelegateCall(in registry.DecodedFields) *pbgear.ConvictionVoting_DelegateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_DelegateCall{}
	out.Class = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.To = To_ConvictionVoting_To(fields[1].Value.(registry.DecodedFields))

	out.Conviction = To_ConvictionVoting_Conviction(fields[2].Value.(registry.DecodedFields))

	out.Balance = To_string(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Id(in registry.DecodedFields) *pbgear.ConvictionVoting_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Index(in registry.DecodedFields) *pbgear.ConvictionVoting_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Index{}
	out.Value0 = To_ConvictionVoting_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_ConvictionVoting_Locked1X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked1X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked1X{}
	return out
}

func To_ConvictionVoting_Locked2X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked2X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked2X{}
	return out
}

func To_ConvictionVoting_Locked3X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked3X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked3X{}
	return out
}

func To_ConvictionVoting_Locked4X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked4X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked4X{}
	return out
}

func To_ConvictionVoting_Locked5X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked5X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked5X{}
	return out
}

func To_ConvictionVoting_Locked6X(in registry.DecodedFields) *pbgear.ConvictionVoting_Locked6X {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Locked6X{}
	return out
}

func To_ConvictionVoting_None(in registry.DecodedFields) *pbgear.ConvictionVoting_None {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_None{}
	return out
}

func To_ConvictionVoting_PalletConvictionVotingVoteVote(in registry.DecodedFields) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{}
	out.Value = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Raw(in registry.DecodedFields) *pbgear.ConvictionVoting_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_ConvictionVoting_RemoveOtherVoteCall(in registry.DecodedFields) *pbgear.ConvictionVoting_RemoveOtherVoteCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
	out.Target = To_ConvictionVoting_Target(fields[0].Value.(registry.DecodedFields))

	out.Class = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_RemoveVoteCall(in registry.DecodedFields) *pbgear.ConvictionVoting_RemoveVoteCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_RemoveVoteCall{}
	out.Class = To_optional_uint32(fields[0].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Split(in registry.DecodedFields) *pbgear.ConvictionVoting_Split {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Split{}
	out.Aye = To_string(fields[0].Value.(registry.DecodedFields))

	out.Nay = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_SplitAbstain(in registry.DecodedFields) *pbgear.ConvictionVoting_SplitAbstain {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_SplitAbstain{}
	out.Aye = To_string(fields[0].Value.(registry.DecodedFields))

	out.Nay = To_string(fields[1].Value.(registry.DecodedFields))

	out.Abstain = To_string(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Standard(in registry.DecodedFields) *pbgear.ConvictionVoting_Standard {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_Standard{}
	out.Vote = To_ConvictionVoting_PalletConvictionVotingVoteVote(fields[0].Value.(registry.DecodedFields))

	out.Balance = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Target(in registry.DecodedFields) *pbgear.ConvictionVoting_Target {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_ConvictionVoting_To(in registry.DecodedFields) *pbgear.ConvictionVoting_To {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_ConvictionVoting_UndelegateCall(in registry.DecodedFields) *pbgear.ConvictionVoting_UndelegateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_UndelegateCall{}
	out.Class = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_UnlockCall(in registry.DecodedFields) *pbgear.ConvictionVoting_UnlockCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_UnlockCall{}
	out.Class = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Target = To_ConvictionVoting_Target(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ConvictionVoting_Vote(in registry.DecodedFields) *pbgear.ConvictionVoting_Vote {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_ConvictionVoting_VoteCall(in registry.DecodedFields) *pbgear.ConvictionVoting_VoteCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ConvictionVoting_VoteCall{}
	out.PollIndex = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Vote = To_ConvictionVoting_Vote(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_GovernanceFallbackCall(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
	out.MaybeMaxVoters = To_optional_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaybeMaxTargets = To_optional_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
	out.Solution = To_VaraRuntimeNposSolution16(fields[0].Value.(registry.DecodedFields))

	out.Score = To_SpNposElectionsElectionScore(fields[1].Value.(registry.DecodedFields))

	out.Round = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
	out.Voters = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Targets = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
	out.Supports = To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
	out.MaybeNextScore = To_optional_SpNposElectionsElectionScore(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_optional_SpNposElectionsElectionScore(f registry.DecodedFields) *pbgear.SpNposElectionsElectionScore {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_SpNposElectionsElectionScore(fields)
	return out
}

func To_ElectionProviderMultiPhase_SubmitCall(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_SubmitCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_SubmitUnsignedCall(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value.(registry.DecodedFields))

	out.Witness = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_CompactSpArithmeticPerThingsPerU16(fields[1].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_CompactUint32(fields[1].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{}
	out.Value0 = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Value1 = To_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value.(registry.DecodedFields))

	out.Value2 = To_CompactUint32(fields[2].Value.(types.UCompact))

	return out
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_SpNposElectionsSupport(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(in registry.DecodedFields) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_AddMemberCall(in registry.DecodedFields) *pbgear.FellowshipCollective_AddMemberCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_AddMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_Address20(in registry.DecodedFields) *pbgear.FellowshipCollective_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_FellowshipCollective_Address32(in registry.DecodedFields) *pbgear.FellowshipCollective_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_FellowshipCollective_CleanupPollCall(in registry.DecodedFields) *pbgear.FellowshipCollective_CleanupPollCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_CleanupPollCall{}
	out.PollIndex = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Max = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_CompactTupleNull(in types.UCompact) *pbgear.FellowshipCollective_CompactTupleNull {

	out := &pbgear.FellowshipCollective_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_FellowshipCollective_DemoteMemberCall(in registry.DecodedFields) *pbgear.FellowshipCollective_DemoteMemberCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_DemoteMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_Id(in registry.DecodedFields) *pbgear.FellowshipCollective_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_Index(in registry.DecodedFields) *pbgear.FellowshipCollective_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_Index{}
	out.Value0 = To_FellowshipCollective_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_FellowshipCollective_PromoteMemberCall(in registry.DecodedFields) *pbgear.FellowshipCollective_PromoteMemberCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_PromoteMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_Raw(in registry.DecodedFields) *pbgear.FellowshipCollective_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_FellowshipCollective_RemoveMemberCall(in registry.DecodedFields) *pbgear.FellowshipCollective_RemoveMemberCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_RemoveMemberCall{}
	out.Who = To_FellowshipCollective_Who(fields[0].Value.(registry.DecodedFields))

	out.MinRank = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_VoteCall(in registry.DecodedFields) *pbgear.FellowshipCollective_VoteCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipCollective_VoteCall{}
	out.Poll = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Aye = To_bool(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipCollective_Who(in registry.DecodedFields) *pbgear.FellowshipCollective_Who {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_FellowshipReferenda_After(in registry.DecodedFields) *pbgear.FellowshipReferenda_After {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_After{}
	out.Value0 = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_At(in registry.DecodedFields) *pbgear.FellowshipReferenda_At {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_At{}
	out.Value0 = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_CancelCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_CancelCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_CancelCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_EnactmentMoment(in registry.DecodedFields) *pbgear.FellowshipReferenda_EnactmentMoment {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_FellowshipReferenda_Inline(in registry.DecodedFields) *pbgear.FellowshipReferenda_Inline {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_Inline{}
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_KillCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_KillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_KillCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_Legacy(in registry.DecodedFields) *pbgear.FellowshipReferenda_Legacy {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_Legacy{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_Lookup(in registry.DecodedFields) *pbgear.FellowshipReferenda_Lookup {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_Lookup{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.Len = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_NudgeReferendumCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_NudgeReferendumCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_OneFewerDecidingCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_OneFewerDecidingCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
	out.Track = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_Origins(in registry.DecodedFields) *pbgear.FellowshipReferenda_Origins {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_Origins{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_PlaceDecisionDepositCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_Proposal(in registry.DecodedFields) *pbgear.FellowshipReferenda_Proposal {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_FellowshipReferenda_ProposalOrigin(in registry.DecodedFields) *pbgear.FellowshipReferenda_ProposalOrigin {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_FellowshipReferenda_RefundDecisionDepositCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_RefundSubmissionDepositCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_SetMetadataCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_SetMetadataCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_SetMetadataCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaybeHash = To_optional_PrimitiveTypesH256(fields[1].Value.(registry.DecodedFields))

	return out
}
func To_optional_PrimitiveTypesH256(f registry.DecodedFields) *pbgear.PrimitiveTypesH256 {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_PrimitiveTypesH256(fields)
	return out
}

func To_FellowshipReferenda_SubmitCall(in registry.DecodedFields) *pbgear.FellowshipReferenda_SubmitCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_SubmitCall{}
	out.ProposalOrigin = To_FellowshipReferenda_ProposalOrigin(fields[0].Value.(registry.DecodedFields))

	out.Proposal = To_FellowshipReferenda_Proposal(fields[1].Value.(registry.DecodedFields))

	out.EnactmentMoment = To_FellowshipReferenda_EnactmentMoment(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_System(in registry.DecodedFields) *pbgear.FellowshipReferenda_System {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_System{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FellowshipReferenda_Void(in registry.DecodedFields) *pbgear.FellowshipReferenda_Void {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FellowshipReferenda_Void{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_FinalityGrandpaEquivocation(in registry.DecodedFields) *pbgear.FinalityGrandpaEquivocation {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FinalityGrandpaEquivocation{}
	out.RoundNumber = To_uint64(fields[0].Value.(registry.DecodedFields))

	out.Identity = To_SpConsensusGrandpaAppPublic(fields[1].Value.(registry.DecodedFields))

	out.First = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[2].Value.(registry.DecodedFields))

	out.Second = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_FinalityGrandpaPrecommit(in registry.DecodedFields) *pbgear.FinalityGrandpaPrecommit {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FinalityGrandpaPrecommit{}
	out.TargetHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.TargetNumber = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_FinalityGrandpaPrevote(in registry.DecodedFields) *pbgear.FinalityGrandpaPrevote {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.FinalityGrandpaPrevote{}
	out.TargetHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.TargetNumber = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_AppendPrograms(in registry.DecodedFields) *pbgear.GearVoucher_AppendPrograms {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_AppendPrograms{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.GearVoucher_AppendPrograms_None{
			None: To_GearVoucher_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_GearVoucher_Call(in registry.DecodedFields) *pbgear.GearVoucher_Call {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_Call{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.GearVoucher_Call_SendMessage{
			SendMessage: To_GearVoucher_SendMessage(fields), //Passthrough fields...
		}
	}
	return out
}

func To_GearVoucher_CallCall(in registry.DecodedFields) *pbgear.GearVoucher_CallCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_CallCall{}
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.(registry.DecodedFields))

	out.Call = To_GearVoucher_Call(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_CallDeprecatedCall(in registry.DecodedFields) *pbgear.GearVoucher_CallDeprecatedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_CallDeprecatedCall{}
	out.Call = To_GearVoucher_Call(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_DeclineCall(in registry.DecodedFields) *pbgear.GearVoucher_DeclineCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_DeclineCall{}
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_DeclineVoucher(in registry.DecodedFields) *pbgear.GearVoucher_DeclineVoucher {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_DeclineVoucher{}
	return out
}

func To_GearVoucher_IssueCall(in registry.DecodedFields) *pbgear.GearVoucher_IssueCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_IssueCall{}
	out.Spender = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.Balance = To_string(fields[1].Value.(registry.DecodedFields))

	out.Programs = To_optional_BTreeSet(fields[2].Value.(registry.DecodedFields))

	out.CodeUploading = To_bool(fields[3].Value.(registry.DecodedFields))

	out.Duration = To_uint32(fields[4].Value.(registry.DecodedFields))

	return out
}
func To_optional_BTreeSet(f registry.DecodedFields) *pbgear.BTreeSet {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_BTreeSet(fields)
	return out
}

func To_GearVoucher_None(in registry.DecodedFields) *pbgear.GearVoucher_None {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_None{}
	return out
}

func To_GearVoucher_PalletGearVoucherInternalVoucherId(in registry.DecodedFields) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_GearVoucher_RevokeCall(in registry.DecodedFields) *pbgear.GearVoucher_RevokeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_RevokeCall{}
	out.Spender = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_SendMessage(in registry.DecodedFields) *pbgear.GearVoucher_SendMessage {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_SendMessage{}
	out.Destination = To_GprimitivesActorId(fields[0].Value.(registry.DecodedFields))

	out.Payload = To_bytes(fields[1].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[2].Value.(registry.DecodedFields))

	out.Value = To_string(fields[3].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_SendReply(in registry.DecodedFields) *pbgear.GearVoucher_SendReply {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_SendReply{}
	out.ReplyToId = To_GprimitivesMessageId(fields[0].Value.(registry.DecodedFields))

	out.Payload = To_bytes(fields[1].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[2].Value.(registry.DecodedFields))

	out.Value = To_string(fields[3].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_Some(in registry.DecodedFields) *pbgear.GearVoucher_Some {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_Some{}
	out.Value0 = To_BTreeSet(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_GearVoucher_UpdateCall(in registry.DecodedFields) *pbgear.GearVoucher_UpdateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_UpdateCall{}
	out.Spender = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value.(registry.DecodedFields))

	out.MoveOwnership = To_optional_SpCoreCryptoAccountId32(fields[2].Value.(registry.DecodedFields))

	out.BalanceTopUp = To_optional_string(fields[3].Value.(registry.DecodedFields))

	out.AppendPrograms = To_optional_GearVoucher_AppendPrograms(fields[4].Value.(registry.DecodedFields))

	out.CodeUploading = To_optional_bool(fields[5].Value.(registry.DecodedFields))

	out.ProlongDuration = To_optional_uint32(fields[6].Value.(registry.DecodedFields))

	return out
}
func To_optional_SpCoreCryptoAccountId32(f registry.DecodedFields) *pbgear.SpCoreCryptoAccountId32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_SpCoreCryptoAccountId32(fields)
	return out
}
func To_optional_GearVoucher_AppendPrograms(f registry.DecodedFields) *pbgear.GearVoucher_AppendPrograms {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_GearVoucher_AppendPrograms(fields)
	return out
}

func To_GearVoucher_UploadCode(in registry.DecodedFields) *pbgear.GearVoucher_UploadCode {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GearVoucher_UploadCode{}
	out.Code = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Gear_ClaimValueCall(in registry.DecodedFields) *pbgear.Gear_ClaimValueCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_ClaimValueCall{}
	out.MessageId = To_GprimitivesMessageId(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Gear_CreateProgramCall(in registry.DecodedFields) *pbgear.Gear_CreateProgramCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_CreateProgramCall{}
	out.CodeId = To_GprimitivesCodeId(fields[0].Value.(registry.DecodedFields))

	out.Salt = To_bytes(fields[1].Value.([]interface{}))
	out.InitPayload = To_bytes(fields[2].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[3].Value.(registry.DecodedFields))

	out.Value = To_string(fields[4].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[5].Value.(registry.DecodedFields))

	return out
}

func To_Gear_RunCall(in registry.DecodedFields) *pbgear.Gear_RunCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_RunCall{}
	out.MaxGas = To_optional_uint64(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Gear_SendMessageCall(in registry.DecodedFields) *pbgear.Gear_SendMessageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_SendMessageCall{}
	out.Destination = To_GprimitivesActorId(fields[0].Value.(registry.DecodedFields))

	out.Payload = To_bytes(fields[1].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[2].Value.(registry.DecodedFields))

	out.Value = To_string(fields[3].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_Gear_SendReplyCall(in registry.DecodedFields) *pbgear.Gear_SendReplyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_SendReplyCall{}
	out.ReplyToId = To_GprimitivesMessageId(fields[0].Value.(registry.DecodedFields))

	out.Payload = To_bytes(fields[1].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[2].Value.(registry.DecodedFields))

	out.Value = To_string(fields[3].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_Gear_SetExecuteInherentCall(in registry.DecodedFields) *pbgear.Gear_SetExecuteInherentCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_SetExecuteInherentCall{}
	out.Value = To_bool(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Gear_UploadCodeCall(in registry.DecodedFields) *pbgear.Gear_UploadCodeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_UploadCodeCall{}
	out.Code = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Gear_UploadProgramCall(in registry.DecodedFields) *pbgear.Gear_UploadProgramCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Gear_UploadProgramCall{}
	out.Code = To_bytes(fields[0].Value.([]interface{}))
	out.Salt = To_bytes(fields[1].Value.([]interface{}))
	out.InitPayload = To_bytes(fields[2].Value.([]interface{}))
	out.GasLimit = To_uint64(fields[3].Value.(registry.DecodedFields))

	out.Value = To_string(fields[4].Value.(registry.DecodedFields))

	out.KeepAlive = To_bool(fields[5].Value.(registry.DecodedFields))

	return out
}

func To_GprimitivesActorId(in registry.DecodedFields) *pbgear.GprimitivesActorId {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GprimitivesActorId{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_GprimitivesCodeId(in registry.DecodedFields) *pbgear.GprimitivesCodeId {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GprimitivesCodeId{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_GprimitivesMessageId(in registry.DecodedFields) *pbgear.GprimitivesMessageId {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.GprimitivesMessageId{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Grandpa_Equivocation(in registry.DecodedFields) *pbgear.Grandpa_Equivocation {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Grandpa_GrandpaTrieNodesList(in registry.DecodedFields) *pbgear.Grandpa_GrandpaTrieNodesList {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_GrandpaTrieNodesList{}
	out.TrieNodes = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Grandpa_NoteStalledCall(in registry.DecodedFields) *pbgear.Grandpa_NoteStalledCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_NoteStalledCall{}
	out.Delay = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.BestFinalizedBlockNumber = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Grandpa_Precommit(in registry.DecodedFields) *pbgear.Grandpa_Precommit {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_Precommit{}
	out.Value0 = To_FinalityGrandpaEquivocation(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Grandpa_Prevote(in registry.DecodedFields) *pbgear.Grandpa_Prevote {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_Prevote{}
	out.Value0 = To_FinalityGrandpaEquivocation(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Grandpa_ReportEquivocationCall(in registry.DecodedFields) *pbgear.Grandpa_ReportEquivocationCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_ReportEquivocationCall{}
	out.EquivocationProof = To_SpConsensusGrandpaEquivocationProof(fields[0].Value.(registry.DecodedFields))

	out.KeyOwnerProof = To_SpSessionMembershipProof(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Grandpa_ReportEquivocationUnsignedCall(in registry.DecodedFields) *pbgear.Grandpa_ReportEquivocationUnsignedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
	out.EquivocationProof = To_SpConsensusGrandpaEquivocationProof(fields[0].Value.(registry.DecodedFields))

	out.KeyOwnerProof = To_SpSessionMembershipProof(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_Account(in registry.DecodedFields) *pbgear.Identity_Account {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Identity_AddRegistrarCall(in registry.DecodedFields) *pbgear.Identity_AddRegistrarCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_AddRegistrarCall{}
	out.Account = To_Identity_Account(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_AddSubCall(in registry.DecodedFields) *pbgear.Identity_AddSubCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_AddSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.(registry.DecodedFields))

	out.Data = To_Identity_Data(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_Address20(in registry.DecodedFields) *pbgear.Identity_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Address32(in registry.DecodedFields) *pbgear.Identity_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_BlakeTwo256(in registry.DecodedFields) *pbgear.Identity_BlakeTwo256 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_BlakeTwo256{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_CancelRequestCall(in registry.DecodedFields) *pbgear.Identity_CancelRequestCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_CancelRequestCall{}
	out.RegIndex = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_ClearIdentityCall(in registry.DecodedFields) *pbgear.Identity_ClearIdentityCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_ClearIdentityCall{}
	return out
}

func To_Identity_CompactTupleNull(in types.UCompact) *pbgear.Identity_CompactTupleNull {

	out := &pbgear.Identity_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Identity_Data(in registry.DecodedFields) *pbgear.Identity_Data {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Data{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Data_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Display(in registry.DecodedFields) *pbgear.Identity_Display {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Display{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Display_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Email(in registry.DecodedFields) *pbgear.Identity_Email {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Email{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Email_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Erroneous(in registry.DecodedFields) *pbgear.Identity_Erroneous {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Erroneous{}
	return out
}

func To_Identity_FeePaid(in registry.DecodedFields) *pbgear.Identity_FeePaid {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_FeePaid{}
	out.Value0 = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_Id(in registry.DecodedFields) *pbgear.Identity_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_Image(in registry.DecodedFields) *pbgear.Identity_Image {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Image{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Image_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Index(in registry.DecodedFields) *pbgear.Identity_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Index{}
	out.Value0 = To_Identity_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Identity_Judgement(in registry.DecodedFields) *pbgear.Identity_Judgement {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Judgement{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Judgement_Unknown{
			Unknown: To_Identity_Unknown(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Keccak256(in registry.DecodedFields) *pbgear.Identity_Keccak256 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Keccak256{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_KillIdentityCall(in registry.DecodedFields) *pbgear.Identity_KillIdentityCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_KillIdentityCall{}
	out.Target = To_Identity_Target(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_KnownGood(in registry.DecodedFields) *pbgear.Identity_KnownGood {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_KnownGood{}
	return out
}

func To_Identity_Legal(in registry.DecodedFields) *pbgear.Identity_Legal {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Legal{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Legal_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_LowQuality(in registry.DecodedFields) *pbgear.Identity_LowQuality {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_LowQuality{}
	return out
}

func To_Identity_New(in registry.DecodedFields) *pbgear.Identity_New {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Identity_None(in registry.DecodedFields) *pbgear.Identity_None {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_None{}
	return out
}

func To_Identity_OutOfDate(in registry.DecodedFields) *pbgear.Identity_OutOfDate {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_OutOfDate{}
	return out
}

func To_Identity_PalletIdentitySimpleIdentityInfo(in registry.DecodedFields) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
	out.Additional = To_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.(registry.DecodedFields))

	out.Display = To_Identity_Display(fields[1].Value.(registry.DecodedFields))

	out.Legal = To_Identity_Legal(fields[2].Value.(registry.DecodedFields))

	out.Web = To_Identity_Web(fields[3].Value.(registry.DecodedFields))

	out.Riot = To_Identity_Riot(fields[4].Value.(registry.DecodedFields))

	out.Email = To_Identity_Email(fields[5].Value.(registry.DecodedFields))

	out.PgpFingerprint = To_bytes(fields[6].Value.([]interface{}))
	out.Image = To_Identity_Image(fields[7].Value.(registry.DecodedFields))

	out.Twitter = To_Identity_Twitter(fields[8].Value.(registry.DecodedFields))

	return out
}

func To_Identity_PalletIdentityTypesBitFlags(in registry.DecodedFields) *pbgear.Identity_PalletIdentityTypesBitFlags {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
	out.Value = To_uint64(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_ProvideJudgementCall(in registry.DecodedFields) *pbgear.Identity_ProvideJudgementCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_ProvideJudgementCall{}
	out.RegIndex = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Target = To_Identity_Target(fields[1].Value.(registry.DecodedFields))

	out.Judgement = To_Identity_Judgement(fields[2].Value.(registry.DecodedFields))

	out.Identity = To_PrimitiveTypesH256(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_Identity_QuitSubCall(in registry.DecodedFields) *pbgear.Identity_QuitSubCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_QuitSubCall{}
	return out
}

func To_Identity_Raw(in registry.DecodedFields) *pbgear.Identity_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw0(in registry.DecodedFields) *pbgear.Identity_Raw0 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw0{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw1(in registry.DecodedFields) *pbgear.Identity_Raw1 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw1{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw10(in registry.DecodedFields) *pbgear.Identity_Raw10 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw10{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw11(in registry.DecodedFields) *pbgear.Identity_Raw11 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw11{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw12(in registry.DecodedFields) *pbgear.Identity_Raw12 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw12{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw13(in registry.DecodedFields) *pbgear.Identity_Raw13 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw13{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw14(in registry.DecodedFields) *pbgear.Identity_Raw14 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw14{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw15(in registry.DecodedFields) *pbgear.Identity_Raw15 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw15{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw16(in registry.DecodedFields) *pbgear.Identity_Raw16 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw16{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw17(in registry.DecodedFields) *pbgear.Identity_Raw17 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw17{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw18(in registry.DecodedFields) *pbgear.Identity_Raw18 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw18{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw19(in registry.DecodedFields) *pbgear.Identity_Raw19 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw19{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw2(in registry.DecodedFields) *pbgear.Identity_Raw2 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw2{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw20(in registry.DecodedFields) *pbgear.Identity_Raw20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw21(in registry.DecodedFields) *pbgear.Identity_Raw21 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw21{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw22(in registry.DecodedFields) *pbgear.Identity_Raw22 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw22{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw23(in registry.DecodedFields) *pbgear.Identity_Raw23 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw23{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw24(in registry.DecodedFields) *pbgear.Identity_Raw24 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw24{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw25(in registry.DecodedFields) *pbgear.Identity_Raw25 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw25{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw26(in registry.DecodedFields) *pbgear.Identity_Raw26 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw26{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw27(in registry.DecodedFields) *pbgear.Identity_Raw27 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw27{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw28(in registry.DecodedFields) *pbgear.Identity_Raw28 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw28{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw29(in registry.DecodedFields) *pbgear.Identity_Raw29 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw29{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw3(in registry.DecodedFields) *pbgear.Identity_Raw3 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw3{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw30(in registry.DecodedFields) *pbgear.Identity_Raw30 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw30{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw31(in registry.DecodedFields) *pbgear.Identity_Raw31 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw31{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw32(in registry.DecodedFields) *pbgear.Identity_Raw32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw4(in registry.DecodedFields) *pbgear.Identity_Raw4 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw4{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw5(in registry.DecodedFields) *pbgear.Identity_Raw5 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw5{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw6(in registry.DecodedFields) *pbgear.Identity_Raw6 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw6{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw7(in registry.DecodedFields) *pbgear.Identity_Raw7 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw7{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw8(in registry.DecodedFields) *pbgear.Identity_Raw8 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw8{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Raw9(in registry.DecodedFields) *pbgear.Identity_Raw9 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Raw9{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Reasonable(in registry.DecodedFields) *pbgear.Identity_Reasonable {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Reasonable{}
	return out
}

func To_Identity_RemoveSubCall(in registry.DecodedFields) *pbgear.Identity_RemoveSubCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_RemoveSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_RenameSubCall(in registry.DecodedFields) *pbgear.Identity_RenameSubCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_RenameSubCall{}
	out.Sub = To_Identity_Sub(fields[0].Value.(registry.DecodedFields))

	out.Data = To_Identity_Data(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_RequestJudgementCall(in registry.DecodedFields) *pbgear.Identity_RequestJudgementCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_RequestJudgementCall{}
	out.RegIndex = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.MaxFee = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_Identity_Riot(in registry.DecodedFields) *pbgear.Identity_Riot {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Riot{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Riot_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_SetAccountIdCall(in registry.DecodedFields) *pbgear.Identity_SetAccountIdCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_SetAccountIdCall{}
	out.Index = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.New = To_Identity_New(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_SetFeeCall(in registry.DecodedFields) *pbgear.Identity_SetFeeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_SetFeeCall{}
	out.Index = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Fee = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_Identity_SetFieldsCall(in registry.DecodedFields) *pbgear.Identity_SetFieldsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_SetFieldsCall{}
	out.Index = To_CompactUint32(fields[0].Value.(types.UCompact))

	out.Fields = To_Identity_PalletIdentityTypesBitFlags(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_SetIdentityCall(in registry.DecodedFields) *pbgear.Identity_SetIdentityCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_SetIdentityCall{}
	out.Info = To_Identity_PalletIdentitySimpleIdentityInfo(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Identity_SetSubsCall(in registry.DecodedFields) *pbgear.Identity_SetSubsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_SetSubsCall{}
	out.Subs = To_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(f registry.DecodedFields) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData, len(fields))
	for _, f := range fields {
		out = append(out, To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Identity_Sha256(in registry.DecodedFields) *pbgear.Identity_Sha256 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Sha256{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_ShaThree256(in registry.DecodedFields) *pbgear.Identity_ShaThree256 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_ShaThree256{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Identity_Sub(in registry.DecodedFields) *pbgear.Identity_Sub {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Identity_Target(in registry.DecodedFields) *pbgear.Identity_Target {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(in registry.DecodedFields) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
	out.Value0 = To_Identity_Value0(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_Identity_Value1(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(in registry.DecodedFields) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_Identity_Value1(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Identity_Twitter(in registry.DecodedFields) *pbgear.Identity_Twitter {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Twitter{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Twitter_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Unknown(in registry.DecodedFields) *pbgear.Identity_Unknown {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Unknown{}
	return out
}

func To_Identity_Value0(in registry.DecodedFields) *pbgear.Identity_Value0 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Value0{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Value0_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Value1(in registry.DecodedFields) *pbgear.Identity_Value1 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Value1{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Value1_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Identity_Web(in registry.DecodedFields) *pbgear.Identity_Web {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Identity_Web{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Identity_Web_None{
			None: To_Identity_None(fields), //Passthrough fields...
		}
	}
	return out
}

func To_ImOnline_HeartbeatCall(in registry.DecodedFields) *pbgear.ImOnline_HeartbeatCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ImOnline_HeartbeatCall{}
	out.Heartbeat = To_ImOnline_PalletImOnlineHeartbeat(fields[0].Value.(registry.DecodedFields))

	out.Signature = To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_ImOnline_PalletImOnlineHeartbeat(in registry.DecodedFields) *pbgear.ImOnline_PalletImOnlineHeartbeat {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
	out.BlockNumber = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.SessionIndex = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.AuthorityIndex = To_uint32(fields[2].Value.(registry.DecodedFields))

	out.ValidatorsLen = To_uint32(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Public(in registry.DecodedFields) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public{}
	out.Value = To_SpCoreSr25519Public(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(in registry.DecodedFields) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
	out.Value = To_SpCoreSr25519Signature(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Multisig_ApproveAsMultiCall(in registry.DecodedFields) *pbgear.Multisig_ApproveAsMultiCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Multisig_ApproveAsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.OtherSignatories = To_repeated_SpCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	out.MaybeTimepoint = To_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.(registry.DecodedFields))

	out.CallHash = To_bytes(fields[3].Value.([]interface{}))
	out.MaxWeight = To_SpWeightsWeightV2Weight(fields[4].Value.(registry.DecodedFields))

	return out
}
func To_optional_Multisig_PalletMultisigTimepoint(f registry.DecodedFields) *pbgear.Multisig_PalletMultisigTimepoint {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_Multisig_PalletMultisigTimepoint(fields)
	return out
}

func To_Multisig_AsMultiCall(in registry.DecodedFields) *pbgear.Multisig_AsMultiCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Multisig_AsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.OtherSignatories = To_repeated_SpCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	out.MaybeTimepoint = To_optional_Multisig_PalletMultisigTimepoint(fields[2].Value.(registry.DecodedFields))

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
	out.MaxWeight = To_SpWeightsWeightV2Weight(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_Multisig_AsMultiThreshold1Call(in registry.DecodedFields) *pbgear.Multisig_AsMultiThreshold1Call {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Multisig_AsMultiThreshold1Call{}
	out.OtherSignatories = To_repeated_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

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

func To_Multisig_CancelAsMultiCall(in registry.DecodedFields) *pbgear.Multisig_CancelAsMultiCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Multisig_CancelAsMultiCall{}
	out.Threshold = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.OtherSignatories = To_repeated_SpCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	out.Timepoint = To_Multisig_PalletMultisigTimepoint(fields[2].Value.(registry.DecodedFields))

	out.CallHash = To_bytes(fields[3].Value.([]interface{}))
	return out
}

func To_Multisig_PalletMultisigTimepoint(in registry.DecodedFields) *pbgear.Multisig_PalletMultisigTimepoint {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Multisig_PalletMultisigTimepoint{}
	out.Height = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Address20(in registry.DecodedFields) *pbgear.NominationPools_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_NominationPools_Address32(in registry.DecodedFields) *pbgear.NominationPools_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_NominationPools_AdjustPoolDepositCall(in registry.DecodedFields) *pbgear.NominationPools_AdjustPoolDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_AdjustPoolDepositCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Blocked(in registry.DecodedFields) *pbgear.NominationPools_Blocked {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Blocked{}
	return out
}

func To_NominationPools_BondExtraCall(in registry.DecodedFields) *pbgear.NominationPools_BondExtraCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_BondExtraCall{}
	out.Extra = To_NominationPools_Extra(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_BondExtraOtherCall(in registry.DecodedFields) *pbgear.NominationPools_BondExtraOtherCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_BondExtraOtherCall{}
	out.Member = To_NominationPools_Member(fields[0].Value.(registry.DecodedFields))

	out.Extra = To_NominationPools_Extra(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Bouncer(in registry.DecodedFields) *pbgear.NominationPools_Bouncer {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_NominationPools_ChillCall(in registry.DecodedFields) *pbgear.NominationPools_ChillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_ChillCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_ClaimCommissionCall(in registry.DecodedFields) *pbgear.NominationPools_ClaimCommissionCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_ClaimCommissionCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_ClaimPayoutCall(in registry.DecodedFields) *pbgear.NominationPools_ClaimPayoutCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_ClaimPayoutCall{}
	return out
}

func To_NominationPools_ClaimPayoutOtherCall(in registry.DecodedFields) *pbgear.NominationPools_ClaimPayoutOtherCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
	out.Other = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_CompactTupleNull(in types.UCompact) *pbgear.NominationPools_CompactTupleNull {

	out := &pbgear.NominationPools_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_NominationPools_CreateCall(in registry.DecodedFields) *pbgear.NominationPools_CreateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_CreateCall{}
	out.Amount = To_CompactString(fields[0].Value.(types.UCompact))

	out.Root = To_NominationPools_Root(fields[1].Value.(registry.DecodedFields))

	out.Nominator = To_NominationPools_Nominator(fields[2].Value.(registry.DecodedFields))

	out.Bouncer = To_NominationPools_Bouncer(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_CreateWithPoolIdCall(in registry.DecodedFields) *pbgear.NominationPools_CreateWithPoolIdCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_CreateWithPoolIdCall{}
	out.Amount = To_CompactString(fields[0].Value.(types.UCompact))

	out.Root = To_NominationPools_Root(fields[1].Value.(registry.DecodedFields))

	out.Nominator = To_NominationPools_Nominator(fields[2].Value.(registry.DecodedFields))

	out.Bouncer = To_NominationPools_Bouncer(fields[3].Value.(registry.DecodedFields))

	out.PoolId = To_uint32(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Destroying(in registry.DecodedFields) *pbgear.NominationPools_Destroying {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Destroying{}
	return out
}

func To_NominationPools_Extra(in registry.DecodedFields) *pbgear.NominationPools_Extra {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Extra{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_Extra_FreeBalance{
			FreeBalance: To_NominationPools_FreeBalance(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_FreeBalance(in registry.DecodedFields) *pbgear.NominationPools_FreeBalance {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_FreeBalance{}
	out.Value0 = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_GlobalMaxCommission(in registry.DecodedFields) *pbgear.NominationPools_GlobalMaxCommission {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_GlobalMaxCommission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_GlobalMaxCommission_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Id(in registry.DecodedFields) *pbgear.NominationPools_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Index(in registry.DecodedFields) *pbgear.NominationPools_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Index{}
	out.Value0 = To_NominationPools_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_NominationPools_JoinCall(in registry.DecodedFields) *pbgear.NominationPools_JoinCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_JoinCall{}
	out.Amount = To_CompactString(fields[0].Value.(types.UCompact))

	out.PoolId = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_MaxMembers(in registry.DecodedFields) *pbgear.NominationPools_MaxMembers {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_MaxMembers{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxMembers_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MaxMembersPerPool(in registry.DecodedFields) *pbgear.NominationPools_MaxMembersPerPool {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_MaxMembersPerPool{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxMembersPerPool_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MaxPools(in registry.DecodedFields) *pbgear.NominationPools_MaxPools {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_MaxPools{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MaxPools_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Member(in registry.DecodedFields) *pbgear.NominationPools_Member {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_NominationPools_MemberAccount(in registry.DecodedFields) *pbgear.NominationPools_MemberAccount {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_NominationPools_MinCreateBond(in registry.DecodedFields) *pbgear.NominationPools_MinCreateBond {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_MinCreateBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MinCreateBond_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_MinJoinBond(in registry.DecodedFields) *pbgear.NominationPools_MinJoinBond {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_MinJoinBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_MinJoinBond_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewBouncer(in registry.DecodedFields) *pbgear.NominationPools_NewBouncer {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_NewBouncer{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewBouncer_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewNominator(in registry.DecodedFields) *pbgear.NominationPools_NewNominator {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_NewNominator{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewNominator_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NewRoot(in registry.DecodedFields) *pbgear.NominationPools_NewRoot {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_NewRoot{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_NewRoot_Noop{
			Noop: To_NominationPools_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_NominateCall(in registry.DecodedFields) *pbgear.NominationPools_NominateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_NominateCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Validators = To_repeated_SpCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Nominator(in registry.DecodedFields) *pbgear.NominationPools_Nominator {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_NominationPools_Noop(in registry.DecodedFields) *pbgear.NominationPools_Noop {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Noop{}
	return out
}

func To_NominationPools_Open(in registry.DecodedFields) *pbgear.NominationPools_Open {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Open{}
	return out
}

func To_NominationPools_PalletNominationPoolsCommissionChangeRate(in registry.DecodedFields) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
	out.MaxIncrease = To_SpArithmeticPerThingsPerbill(fields[0].Value.(registry.DecodedFields))

	out.MinDelay = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Permission(in registry.DecodedFields) *pbgear.NominationPools_Permission {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Permission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_Permission_Permissioned{
			Permissioned: To_NominationPools_Permissioned(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_Permissioned(in registry.DecodedFields) *pbgear.NominationPools_Permissioned {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Permissioned{}
	return out
}

func To_NominationPools_PermissionlessAll(in registry.DecodedFields) *pbgear.NominationPools_PermissionlessAll {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_PermissionlessAll{}
	return out
}

func To_NominationPools_PermissionlessCompound(in registry.DecodedFields) *pbgear.NominationPools_PermissionlessCompound {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_PermissionlessCompound{}
	return out
}

func To_NominationPools_PermissionlessWithdraw(in registry.DecodedFields) *pbgear.NominationPools_PermissionlessWithdraw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_PermissionlessWithdraw{}
	return out
}

func To_NominationPools_PoolWithdrawUnbondedCall(in registry.DecodedFields) *pbgear.NominationPools_PoolWithdrawUnbondedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.NumSlashingSpans = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_Raw(in registry.DecodedFields) *pbgear.NominationPools_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_NominationPools_Remove(in registry.DecodedFields) *pbgear.NominationPools_Remove {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Remove{}
	return out
}

func To_NominationPools_Rewards(in registry.DecodedFields) *pbgear.NominationPools_Rewards {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Rewards{}
	return out
}

func To_NominationPools_Root(in registry.DecodedFields) *pbgear.NominationPools_Root {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_NominationPools_Set(in registry.DecodedFields) *pbgear.NominationPools_Set {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_Set{}
	out.Value0 = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_SetClaimPermissionCall(in registry.DecodedFields) *pbgear.NominationPools_SetClaimPermissionCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetClaimPermissionCall{}
	out.Permission = To_NominationPools_Permission(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_SetCommissionCall(in registry.DecodedFields) *pbgear.NominationPools_SetCommissionCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetCommissionCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.NewCommission = To_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	return out
}
func To_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(f registry.DecodedFields) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields)
	return out
}

func To_NominationPools_SetCommissionChangeRateCall(in registry.DecodedFields) *pbgear.NominationPools_SetCommissionChangeRateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.ChangeRate = To_NominationPools_PalletNominationPoolsCommissionChangeRate(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_SetCommissionMaxCall(in registry.DecodedFields) *pbgear.NominationPools_SetCommissionMaxCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetCommissionMaxCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaxCommission = To_SpArithmeticPerThingsPerbill(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_SetConfigsCall(in registry.DecodedFields) *pbgear.NominationPools_SetConfigsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetConfigsCall{}
	out.MinJoinBond = To_NominationPools_MinJoinBond(fields[0].Value.(registry.DecodedFields))

	out.MinCreateBond = To_NominationPools_MinCreateBond(fields[1].Value.(registry.DecodedFields))

	out.MaxPools = To_NominationPools_MaxPools(fields[2].Value.(registry.DecodedFields))

	out.MaxMembers = To_NominationPools_MaxMembers(fields[3].Value.(registry.DecodedFields))

	out.MaxMembersPerPool = To_NominationPools_MaxMembersPerPool(fields[4].Value.(registry.DecodedFields))

	out.GlobalMaxCommission = To_NominationPools_GlobalMaxCommission(fields[5].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_SetMetadataCall(in registry.DecodedFields) *pbgear.NominationPools_SetMetadataCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetMetadataCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Metadata = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_NominationPools_SetStateCall(in registry.DecodedFields) *pbgear.NominationPools_SetStateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_SetStateCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.State = To_NominationPools_State(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_State(in registry.DecodedFields) *pbgear.NominationPools_State {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_State{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.NominationPools_State_Open{
			Open: To_NominationPools_Open(fields), //Passthrough fields...
		}
	}
	return out
}

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in registry.DecodedFields) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
	out.Value0 = To_SpArithmeticPerThingsPerbill(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_SpCoreCryptoAccountId32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_UnbondCall(in registry.DecodedFields) *pbgear.NominationPools_UnbondCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_UnbondCall{}
	out.MemberAccount = To_NominationPools_MemberAccount(fields[0].Value.(registry.DecodedFields))

	out.UnbondingPoints = To_CompactString(fields[1].Value.(types.UCompact))

	return out
}

func To_NominationPools_UpdateRolesCall(in registry.DecodedFields) *pbgear.NominationPools_UpdateRolesCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_UpdateRolesCall{}
	out.PoolId = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.NewRoot = To_NominationPools_NewRoot(fields[1].Value.(registry.DecodedFields))

	out.NewNominator = To_NominationPools_NewNominator(fields[2].Value.(registry.DecodedFields))

	out.NewBouncer = To_NominationPools_NewBouncer(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_NominationPools_WithdrawUnbondedCall(in registry.DecodedFields) *pbgear.NominationPools_WithdrawUnbondedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.NominationPools_WithdrawUnbondedCall{}
	out.MemberAccount = To_NominationPools_MemberAccount(fields[0].Value.(registry.DecodedFields))

	out.NumSlashingSpans = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_None(in registry.DecodedFields) *pbgear.None {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.None{}
	return out
}

func To_Preimage_EnsureUpdatedCall(in registry.DecodedFields) *pbgear.Preimage_EnsureUpdatedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Preimage_EnsureUpdatedCall{}
	out.Hashes = To_repeated_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_PrimitiveTypesH256(f registry.DecodedFields) []*pbgear.PrimitiveTypesH256 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.PrimitiveTypesH256, len(fields))
	for _, f := range fields {
		out = append(out, To_PrimitiveTypesH256(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Preimage_NotePreimageCall(in registry.DecodedFields) *pbgear.Preimage_NotePreimageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Preimage_NotePreimageCall{}
	out.Bytes = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Preimage_RequestPreimageCall(in registry.DecodedFields) *pbgear.Preimage_RequestPreimageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Preimage_RequestPreimageCall{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Preimage_UnnotePreimageCall(in registry.DecodedFields) *pbgear.Preimage_UnnotePreimageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Preimage_UnnotePreimageCall{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Preimage_UnrequestPreimageCall(in registry.DecodedFields) *pbgear.Preimage_UnrequestPreimageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Preimage_UnrequestPreimageCall{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_PrimaryAndSecondaryPlainSlots(in registry.DecodedFields) *pbgear.PrimaryAndSecondaryPlainSlots {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.PrimaryAndSecondaryPlainSlots{}
	return out
}

func To_PrimaryAndSecondaryVRFSlots(in registry.DecodedFields) *pbgear.PrimaryAndSecondaryVRFSlots {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.PrimaryAndSecondaryVRFSlots{}
	return out
}

func To_PrimarySlots(in registry.DecodedFields) *pbgear.PrimarySlots {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.PrimarySlots{}
	return out
}

func To_PrimitiveTypesH256(in registry.DecodedFields) *pbgear.PrimitiveTypesH256 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.PrimitiveTypesH256{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Proxy_AddProxyCall(in registry.DecodedFields) *pbgear.Proxy_AddProxyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_AddProxyCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.(registry.DecodedFields))

	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.(registry.DecodedFields))

	out.Delay = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_Address20(in registry.DecodedFields) *pbgear.Proxy_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Proxy_Address32(in registry.DecodedFields) *pbgear.Proxy_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Proxy_AnnounceCall(in registry.DecodedFields) *pbgear.Proxy_AnnounceCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_AnnounceCall{}
	out.Real = To_Proxy_Real(fields[0].Value.(registry.DecodedFields))

	out.CallHash = To_PrimitiveTypesH256(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_Any(in registry.DecodedFields) *pbgear.Proxy_Any {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Any{}
	return out
}

func To_Proxy_CancelProxy(in registry.DecodedFields) *pbgear.Proxy_CancelProxy {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_CancelProxy{}
	return out
}

func To_Proxy_CompactTupleNull(in types.UCompact) *pbgear.Proxy_CompactTupleNull {

	out := &pbgear.Proxy_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Proxy_CreatePureCall(in registry.DecodedFields) *pbgear.Proxy_CreatePureCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_CreatePureCall{}
	out.ProxyType = To_Proxy_ProxyType(fields[0].Value.(registry.DecodedFields))

	out.Delay = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_Delegate(in registry.DecodedFields) *pbgear.Proxy_Delegate {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Proxy_ForceProxyType(in registry.DecodedFields) *pbgear.Proxy_ForceProxyType {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_ForceProxyType{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Proxy_ForceProxyType_Any{
			Any: To_Proxy_Any(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_Governance(in registry.DecodedFields) *pbgear.Proxy_Governance {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Governance{}
	return out
}

func To_Proxy_Id(in registry.DecodedFields) *pbgear.Proxy_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_IdentityJudgement(in registry.DecodedFields) *pbgear.Proxy_IdentityJudgement {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_IdentityJudgement{}
	return out
}

func To_Proxy_Index(in registry.DecodedFields) *pbgear.Proxy_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Index{}
	out.Value0 = To_Proxy_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Proxy_KillPureCall(in registry.DecodedFields) *pbgear.Proxy_KillPureCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_KillPureCall{}
	out.Spawner = To_Proxy_Spawner(fields[0].Value.(registry.DecodedFields))

	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[2].Value.(registry.DecodedFields))

	out.Height = To_CompactUint32(fields[3].Value.(types.UCompact))

	out.ExtIndex = To_CompactUint32(fields[4].Value.(types.UCompact))

	return out
}

func To_Proxy_NonTransfer(in registry.DecodedFields) *pbgear.Proxy_NonTransfer {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_NonTransfer{}
	return out
}

func To_Proxy_ProxyAnnouncedCall(in registry.DecodedFields) *pbgear.Proxy_ProxyAnnouncedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_ProxyAnnouncedCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.(registry.DecodedFields))

	out.Real = To_Proxy_Real(fields[1].Value.(registry.DecodedFields))

	out.ForceProxyType = To_optional_Proxy_ForceProxyType(fields[2].Value.(registry.DecodedFields))

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
func To_optional_Proxy_ForceProxyType(f registry.DecodedFields) *pbgear.Proxy_ForceProxyType {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_Proxy_ForceProxyType(fields)
	return out
}

func To_Proxy_ProxyCall(in registry.DecodedFields) *pbgear.Proxy_ProxyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_ProxyCall{}
	out.Real = To_Proxy_Real(fields[0].Value.(registry.DecodedFields))

	out.ForceProxyType = To_optional_Proxy_ForceProxyType(fields[1].Value.(registry.DecodedFields))

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

func To_Proxy_ProxyType(in registry.DecodedFields) *pbgear.Proxy_ProxyType {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_ProxyType{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Proxy_ProxyType_Any{
			Any: To_Proxy_Any(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Proxy_Raw(in registry.DecodedFields) *pbgear.Proxy_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Proxy_Real(in registry.DecodedFields) *pbgear.Proxy_Real {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Proxy_RejectAnnouncementCall(in registry.DecodedFields) *pbgear.Proxy_RejectAnnouncementCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_RejectAnnouncementCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.(registry.DecodedFields))

	out.CallHash = To_PrimitiveTypesH256(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_RemoveAnnouncementCall(in registry.DecodedFields) *pbgear.Proxy_RemoveAnnouncementCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_RemoveAnnouncementCall{}
	out.Real = To_Proxy_Real(fields[0].Value.(registry.DecodedFields))

	out.CallHash = To_PrimitiveTypesH256(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_RemoveProxiesCall(in registry.DecodedFields) *pbgear.Proxy_RemoveProxiesCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_RemoveProxiesCall{}
	return out
}

func To_Proxy_RemoveProxyCall(in registry.DecodedFields) *pbgear.Proxy_RemoveProxyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_RemoveProxyCall{}
	out.Delegate = To_Proxy_Delegate(fields[0].Value.(registry.DecodedFields))

	out.ProxyType = To_Proxy_ProxyType(fields[1].Value.(registry.DecodedFields))

	out.Delay = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Proxy_Spawner(in registry.DecodedFields) *pbgear.Proxy_Spawner {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Proxy_Staking(in registry.DecodedFields) *pbgear.Proxy_Staking {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Proxy_Staking{}
	return out
}

func To_Referenda_After(in registry.DecodedFields) *pbgear.Referenda_After {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_After{}
	out.Value0 = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_At(in registry.DecodedFields) *pbgear.Referenda_At {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_At{}
	out.Value0 = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_CancelCall(in registry.DecodedFields) *pbgear.Referenda_CancelCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_CancelCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_EnactmentMoment(in registry.DecodedFields) *pbgear.Referenda_EnactmentMoment {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Referenda_Inline(in registry.DecodedFields) *pbgear.Referenda_Inline {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_Inline{}
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_KillCall(in registry.DecodedFields) *pbgear.Referenda_KillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_KillCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_Legacy(in registry.DecodedFields) *pbgear.Referenda_Legacy {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_Legacy{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_Lookup(in registry.DecodedFields) *pbgear.Referenda_Lookup {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_Lookup{}
	out.Hash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.Len = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_NudgeReferendumCall(in registry.DecodedFields) *pbgear.Referenda_NudgeReferendumCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_NudgeReferendumCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_OneFewerDecidingCall(in registry.DecodedFields) *pbgear.Referenda_OneFewerDecidingCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_OneFewerDecidingCall{}
	out.Track = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_Origins(in registry.DecodedFields) *pbgear.Referenda_Origins {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_Origins{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_PlaceDecisionDepositCall(in registry.DecodedFields) *pbgear.Referenda_PlaceDecisionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_PlaceDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_Proposal(in registry.DecodedFields) *pbgear.Referenda_Proposal {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Referenda_ProposalOrigin(in registry.DecodedFields) *pbgear.Referenda_ProposalOrigin {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Referenda_RefundDecisionDepositCall(in registry.DecodedFields) *pbgear.Referenda_RefundDecisionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_RefundDecisionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_RefundSubmissionDepositCall(in registry.DecodedFields) *pbgear.Referenda_RefundSubmissionDepositCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_RefundSubmissionDepositCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_SetMetadataCall(in registry.DecodedFields) *pbgear.Referenda_SetMetadataCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_SetMetadataCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaybeHash = To_optional_PrimitiveTypesH256(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_SubmitCall(in registry.DecodedFields) *pbgear.Referenda_SubmitCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_SubmitCall{}
	out.ProposalOrigin = To_Referenda_ProposalOrigin(fields[0].Value.(registry.DecodedFields))

	out.Proposal = To_Referenda_Proposal(fields[1].Value.(registry.DecodedFields))

	out.EnactmentMoment = To_Referenda_EnactmentMoment(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_System(in registry.DecodedFields) *pbgear.Referenda_System {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_System{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Referenda_Void(in registry.DecodedFields) *pbgear.Referenda_Void {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Referenda_Void{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Root(in registry.DecodedFields) *pbgear.Root {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Root{}
	return out
}

func To_Scheduler_CancelCall(in registry.DecodedFields) *pbgear.Scheduler_CancelCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_CancelCall{}
	out.When = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Index = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Scheduler_CancelNamedCall(in registry.DecodedFields) *pbgear.Scheduler_CancelNamedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_CancelNamedCall{}
	out.Id = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Scheduler_ScheduleAfterCall(in registry.DecodedFields) *pbgear.Scheduler_ScheduleAfterCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_ScheduleAfterCall{}
	out.After = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[1].Value.(registry.DecodedFields))

	out.Priority = To_uint32(fields[2].Value.(registry.DecodedFields))

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
func To_optional_Scheduler_TupleUint32Uint32(f registry.DecodedFields) *pbgear.Scheduler_TupleUint32Uint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	var out = To_Scheduler_TupleUint32Uint32(fields)
	return out
}

func To_Scheduler_ScheduleCall(in registry.DecodedFields) *pbgear.Scheduler_ScheduleCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_ScheduleCall{}
	out.When = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[1].Value.(registry.DecodedFields))

	out.Priority = To_uint32(fields[2].Value.(registry.DecodedFields))

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

func To_Scheduler_ScheduleNamedAfterCall(in registry.DecodedFields) *pbgear.Scheduler_ScheduleNamedAfterCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
	out.Id = To_bytes(fields[0].Value.([]interface{}))
	out.After = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[2].Value.(registry.DecodedFields))

	out.Priority = To_uint32(fields[3].Value.(registry.DecodedFields))

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

func To_Scheduler_ScheduleNamedCall(in registry.DecodedFields) *pbgear.Scheduler_ScheduleNamedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_ScheduleNamedCall{}
	out.Id = To_bytes(fields[0].Value.([]interface{}))
	out.When = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.MaybePeriodic = To_optional_Scheduler_TupleUint32Uint32(fields[2].Value.(registry.DecodedFields))

	out.Priority = To_uint32(fields[3].Value.(registry.DecodedFields))

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

func To_Scheduler_TupleUint32Uint32(in registry.DecodedFields) *pbgear.Scheduler_TupleUint32Uint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Scheduler_TupleUint32Uint32{}
	out.Value0 = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Session_PurgeKeysCall(in registry.DecodedFields) *pbgear.Session_PurgeKeysCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Session_PurgeKeysCall{}
	return out
}

func To_Session_SetKeysCall(in registry.DecodedFields) *pbgear.Session_SetKeysCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Session_SetKeysCall{}
	out.Keys = To_VaraRuntimeSessionKeys(fields[0].Value.(registry.DecodedFields))

	out.Proof = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Signed(in registry.DecodedFields) *pbgear.Signed {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Signed{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpArithmeticPerThingsPerU16(in registry.DecodedFields) *pbgear.SpArithmeticPerThingsPerU16 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpArithmeticPerThingsPerU16{}
	out.Value = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpArithmeticPerThingsPerbill(in registry.DecodedFields) *pbgear.SpArithmeticPerThingsPerbill {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpArithmeticPerThingsPerbill{}
	out.Value = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpArithmeticPerThingsPercent(in registry.DecodedFields) *pbgear.SpArithmeticPerThingsPercent {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpArithmeticPerThingsPercent{}
	out.Value = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpAuthorityDiscoveryAppPublic(in registry.DecodedFields) *pbgear.SpAuthorityDiscoveryAppPublic {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpAuthorityDiscoveryAppPublic{}
	out.Value = To_SpCoreSr25519Public(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusBabeAppPublic(in registry.DecodedFields) *pbgear.SpConsensusBabeAppPublic {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusBabeAppPublic{}
	out.Value = To_SpCoreSr25519Public(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusGrandpaAppPublic(in registry.DecodedFields) *pbgear.SpConsensusGrandpaAppPublic {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusGrandpaAppPublic{}
	out.Value = To_SpCoreEd25519Public(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusGrandpaAppSignature(in registry.DecodedFields) *pbgear.SpConsensusGrandpaAppSignature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusGrandpaAppSignature{}
	out.Value = To_SpCoreEd25519Signature(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusGrandpaEquivocationProof(in registry.DecodedFields) *pbgear.SpConsensusGrandpaEquivocationProof {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusGrandpaEquivocationProof{}
	out.SetId = To_uint64(fields[0].Value.(registry.DecodedFields))

	out.Equivocation = To_Grandpa_Equivocation(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusSlotsEquivocationProof(in registry.DecodedFields) *pbgear.SpConsensusSlotsEquivocationProof {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusSlotsEquivocationProof{}
	out.Offender = To_SpConsensusBabeAppPublic(fields[0].Value.(registry.DecodedFields))

	out.Slot = To_SpConsensusSlotsSlot(fields[1].Value.(registry.DecodedFields))

	out.FirstHeader = To_SpRuntimeGenericHeaderHeader(fields[2].Value.(registry.DecodedFields))

	out.SecondHeader = To_SpRuntimeGenericHeaderHeader(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_SpConsensusSlotsSlot(in registry.DecodedFields) *pbgear.SpConsensusSlotsSlot {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpConsensusSlotsSlot{}
	out.Value = To_uint64(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpCoreCryptoAccountId32(in registry.DecodedFields) *pbgear.SpCoreCryptoAccountId32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpCoreCryptoAccountId32{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_SpCoreEd25519Public(in registry.DecodedFields) *pbgear.SpCoreEd25519Public {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpCoreEd25519Public{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_SpCoreEd25519Signature(in registry.DecodedFields) *pbgear.SpCoreEd25519Signature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpCoreEd25519Signature{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_SpCoreSr25519Public(in registry.DecodedFields) *pbgear.SpCoreSr25519Public {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpCoreSr25519Public{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_SpCoreSr25519Signature(in registry.DecodedFields) *pbgear.SpCoreSr25519Signature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpCoreSr25519Signature{}
	out.Value = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_SpNposElectionsElectionScore(in registry.DecodedFields) *pbgear.SpNposElectionsElectionScore {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpNposElectionsElectionScore{}
	out.MinimalStake = To_string(fields[0].Value.(registry.DecodedFields))

	out.SumStake = To_string(fields[1].Value.(registry.DecodedFields))

	out.SumStakeSquared = To_string(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_SpNposElectionsSupport(in registry.DecodedFields) *pbgear.SpNposElectionsSupport {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpNposElectionsSupport{}
	out.Total = To_string(fields[0].Value.(registry.DecodedFields))

	out.Voters = To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields[1].Value.(registry.DecodedFields))

	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_SpRuntimeGenericDigestDigest(in registry.DecodedFields) *pbgear.SpRuntimeGenericDigestDigest {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpRuntimeGenericDigestDigest{}
	out.Logs = To_repeated_SpRuntimeGenericDigestDigestItem(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_SpRuntimeGenericDigestDigestItem(f registry.DecodedFields) []*pbgear.SpRuntimeGenericDigestDigestItem {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.SpRuntimeGenericDigestDigestItem, len(fields))
	for _, f := range fields {
		out = append(out, To_SpRuntimeGenericDigestDigestItem(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_SpRuntimeGenericDigestDigestItem(in registry.DecodedFields) *pbgear.SpRuntimeGenericDigestDigestItem {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpRuntimeGenericDigestDigestItem{}
	out.Logs = To_Babe_Logs(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpRuntimeGenericHeaderHeader(in registry.DecodedFields) *pbgear.SpRuntimeGenericHeaderHeader {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpRuntimeGenericHeaderHeader{}
	out.ParentHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.Number = To_CompactUint32(fields[1].Value.(types.UCompact))

	out.StateRoot = To_PrimitiveTypesH256(fields[2].Value.(registry.DecodedFields))

	out.ExtrinsicsRoot = To_PrimitiveTypesH256(fields[3].Value.(registry.DecodedFields))

	out.Digest = To_SpRuntimeGenericDigestDigest(fields[4].Value.(registry.DecodedFields))

	return out
}

func To_SpRuntimeMultiaddressMultiAddress(in registry.DecodedFields) *pbgear.SpRuntimeMultiaddressMultiAddress {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpRuntimeMultiaddressMultiAddress{}
	out.Targets = To_Staking_Targets(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_SpSessionMembershipProof(in registry.DecodedFields) *pbgear.SpSessionMembershipProof {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpSessionMembershipProof{}
	out.Session = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.TrieNodes = To_repeated_Babe_BabeTrieNodesList(fields[1].Value.(registry.DecodedFields))

	out.ValidatorCount = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}
func To_repeated_Babe_BabeTrieNodesList(f registry.DecodedFields) []*pbgear.Babe_BabeTrieNodesList {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.Babe_BabeTrieNodesList, len(fields))
	for _, f := range fields {
		out = append(out, To_Babe_BabeTrieNodesList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_SpWeightsWeightV2Weight(in registry.DecodedFields) *pbgear.SpWeightsWeightV2Weight {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.SpWeightsWeightV2Weight{}
	out.RefTime = To_CompactUint64(fields[0].Value.(types.UCompact))

	out.ProofSize = To_CompactUint64(fields[1].Value.(types.UCompact))

	return out
}

func To_StakingRewards_Address20(in registry.DecodedFields) *pbgear.StakingRewards_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_StakingRewards_Address32(in registry.DecodedFields) *pbgear.StakingRewards_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_StakingRewards_AlignSupplyCall(in registry.DecodedFields) *pbgear.StakingRewards_AlignSupplyCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_AlignSupplyCall{}
	out.Target = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_StakingRewards_CompactTupleNull(in types.UCompact) *pbgear.StakingRewards_CompactTupleNull {

	out := &pbgear.StakingRewards_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_StakingRewards_ForceRefillCall(in registry.DecodedFields) *pbgear.StakingRewards_ForceRefillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_ForceRefillCall{}
	out.From = To_StakingRewards_From(fields[0].Value.(registry.DecodedFields))

	out.Value = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_StakingRewards_From(in registry.DecodedFields) *pbgear.StakingRewards_From {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_StakingRewards_Id(in registry.DecodedFields) *pbgear.StakingRewards_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_StakingRewards_Index(in registry.DecodedFields) *pbgear.StakingRewards_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_Index{}
	out.Value0 = To_StakingRewards_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_StakingRewards_Raw(in registry.DecodedFields) *pbgear.StakingRewards_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_StakingRewards_RefillCall(in registry.DecodedFields) *pbgear.StakingRewards_RefillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_RefillCall{}
	out.Value = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_StakingRewards_To(in registry.DecodedFields) *pbgear.StakingRewards_To {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_StakingRewards_WithdrawCall(in registry.DecodedFields) *pbgear.StakingRewards_WithdrawCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.StakingRewards_WithdrawCall{}
	out.To = To_StakingRewards_To(fields[0].Value.(registry.DecodedFields))

	out.Value = To_string(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Account(in registry.DecodedFields) *pbgear.Staking_Account {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Account{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Address20(in registry.DecodedFields) *pbgear.Staking_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Staking_Address32(in registry.DecodedFields) *pbgear.Staking_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Staking_BondCall(in registry.DecodedFields) *pbgear.Staking_BondCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_BondCall{}
	out.Value = To_CompactString(fields[0].Value.(types.UCompact))

	out.Payee = To_Staking_Payee(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_BondExtraCall(in registry.DecodedFields) *pbgear.Staking_BondExtraCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_BondExtraCall{}
	out.MaxAdditional = To_CompactString(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_CancelDeferredSlashCall(in registry.DecodedFields) *pbgear.Staking_CancelDeferredSlashCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_CancelDeferredSlashCall{}
	out.Era = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.SlashIndices = To_repeated_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_ChillCall(in registry.DecodedFields) *pbgear.Staking_ChillCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ChillCall{}
	return out
}

func To_Staking_ChillOtherCall(in registry.DecodedFields) *pbgear.Staking_ChillOtherCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ChillOtherCall{}
	out.Controller = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_ChillThreshold(in registry.DecodedFields) *pbgear.Staking_ChillThreshold {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ChillThreshold{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_ChillThreshold_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_CompactTupleNull(in types.UCompact) *pbgear.Staking_CompactTupleNull {

	out := &pbgear.Staking_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Staking_Controller(in registry.DecodedFields) *pbgear.Staking_Controller {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Controller{}
	return out
}

func To_Staking_ForceApplyMinCommissionCall(in registry.DecodedFields) *pbgear.Staking_ForceApplyMinCommissionCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ForceApplyMinCommissionCall{}
	out.ValidatorStash = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_ForceNewEraAlwaysCall(in registry.DecodedFields) *pbgear.Staking_ForceNewEraAlwaysCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ForceNewEraAlwaysCall{}
	return out
}

func To_Staking_ForceNewEraCall(in registry.DecodedFields) *pbgear.Staking_ForceNewEraCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ForceNewEraCall{}
	return out
}

func To_Staking_ForceNoErasCall(in registry.DecodedFields) *pbgear.Staking_ForceNoErasCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ForceNoErasCall{}
	return out
}

func To_Staking_ForceUnstakeCall(in registry.DecodedFields) *pbgear.Staking_ForceUnstakeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ForceUnstakeCall{}
	out.Stash = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.NumSlashingSpans = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Id(in registry.DecodedFields) *pbgear.Staking_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_IncreaseValidatorCountCall(in registry.DecodedFields) *pbgear.Staking_IncreaseValidatorCountCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_IncreaseValidatorCountCall{}
	out.Additional = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_Index(in registry.DecodedFields) *pbgear.Staking_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Index{}
	out.Value0 = To_Staking_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_KickCall(in registry.DecodedFields) *pbgear.Staking_KickCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_KickCall{}
	out.Who = To_repeated_SpRuntimeMultiaddressMultiAddress(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_SpRuntimeMultiaddressMultiAddress(f registry.DecodedFields) []*pbgear.SpRuntimeMultiaddressMultiAddress {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.SpRuntimeMultiaddressMultiAddress, len(fields))
	for _, f := range fields {
		out = append(out, To_SpRuntimeMultiaddressMultiAddress(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Staking_MaxNominatorCount(in registry.DecodedFields) *pbgear.Staking_MaxNominatorCount {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_MaxNominatorCount{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MaxNominatorCount_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MaxValidatorCount(in registry.DecodedFields) *pbgear.Staking_MaxValidatorCount {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_MaxValidatorCount{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MaxValidatorCount_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinCommission(in registry.DecodedFields) *pbgear.Staking_MinCommission {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_MinCommission{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinCommission_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinNominatorBond(in registry.DecodedFields) *pbgear.Staking_MinNominatorBond {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_MinNominatorBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinNominatorBond_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_MinValidatorBond(in registry.DecodedFields) *pbgear.Staking_MinValidatorBond {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_MinValidatorBond{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_MinValidatorBond_Noop{
			Noop: To_Staking_Noop(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_NominateCall(in registry.DecodedFields) *pbgear.Staking_NominateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_NominateCall{}
	out.Targets = To_repeated_SpRuntimeMultiaddressMultiAddress(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_None(in registry.DecodedFields) *pbgear.Staking_None {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_None{}
	return out
}

func To_Staking_Noop(in registry.DecodedFields) *pbgear.Staking_Noop {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Noop{}
	return out
}

func To_Staking_PalletStakingValidatorPrefs(in registry.DecodedFields) *pbgear.Staking_PalletStakingValidatorPrefs {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_PalletStakingValidatorPrefs{}
	out.Commission = To_CompactSpArithmeticPerThingsPerbill(fields[0].Value.(types.UCompact))

	out.Blocked = To_bool(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Payee(in registry.DecodedFields) *pbgear.Staking_Payee {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Payee{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Staking_Payee_Staked{
			Staked: To_Staking_Staked(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_PayoutStakersCall(in registry.DecodedFields) *pbgear.Staking_PayoutStakersCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_PayoutStakersCall{}
	out.ValidatorStash = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.Era = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Raw(in registry.DecodedFields) *pbgear.Staking_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Staking_ReapStashCall(in registry.DecodedFields) *pbgear.Staking_ReapStashCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ReapStashCall{}
	out.Stash = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	out.NumSlashingSpans = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Staking_RebondCall(in registry.DecodedFields) *pbgear.Staking_RebondCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_RebondCall{}
	out.Value = To_CompactString(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_Remove(in registry.DecodedFields) *pbgear.Staking_Remove {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Remove{}
	return out
}

func To_Staking_ScaleValidatorCountCall(in registry.DecodedFields) *pbgear.Staking_ScaleValidatorCountCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ScaleValidatorCountCall{}
	out.Factor = To_SpArithmeticPerThingsPercent(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Set(in registry.DecodedFields) *pbgear.Staking_Set {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Set{}
	out.Value0 = To_string(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_SetControllerCall(in registry.DecodedFields) *pbgear.Staking_SetControllerCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetControllerCall{}
	return out
}

func To_Staking_SetInvulnerablesCall(in registry.DecodedFields) *pbgear.Staking_SetInvulnerablesCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetInvulnerablesCall{}
	out.Invulnerables = To_repeated_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_SetMinCommissionCall(in registry.DecodedFields) *pbgear.Staking_SetMinCommissionCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetMinCommissionCall{}
	out.New = To_SpArithmeticPerThingsPerbill(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_SetPayeeCall(in registry.DecodedFields) *pbgear.Staking_SetPayeeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetPayeeCall{}
	out.Payee = To_Staking_Payee(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_SetStakingConfigsCall(in registry.DecodedFields) *pbgear.Staking_SetStakingConfigsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetStakingConfigsCall{}
	out.MinNominatorBond = To_Staking_MinNominatorBond(fields[0].Value.(registry.DecodedFields))

	out.MinValidatorBond = To_Staking_MinValidatorBond(fields[1].Value.(registry.DecodedFields))

	out.MaxNominatorCount = To_Staking_MaxNominatorCount(fields[2].Value.(registry.DecodedFields))

	out.MaxValidatorCount = To_Staking_MaxValidatorCount(fields[3].Value.(registry.DecodedFields))

	out.ChillThreshold = To_Staking_ChillThreshold(fields[4].Value.(registry.DecodedFields))

	out.MinCommission = To_Staking_MinCommission(fields[5].Value.(registry.DecodedFields))

	return out
}

func To_Staking_SetValidatorCountCall(in registry.DecodedFields) *pbgear.Staking_SetValidatorCountCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_SetValidatorCountCall{}
	out.New = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_Staked(in registry.DecodedFields) *pbgear.Staking_Staked {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Staked{}
	return out
}

func To_Staking_Stash(in registry.DecodedFields) *pbgear.Staking_Stash {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Stash{}
	return out
}

func To_Staking_Targets(in registry.DecodedFields) *pbgear.Staking_Targets {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Staking_UnbondCall(in registry.DecodedFields) *pbgear.Staking_UnbondCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_UnbondCall{}
	out.Value = To_CompactString(fields[0].Value.(types.UCompact))

	return out
}

func To_Staking_ValidateCall(in registry.DecodedFields) *pbgear.Staking_ValidateCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_ValidateCall{}
	out.Prefs = To_Staking_PalletStakingValidatorPrefs(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Staking_Who(in registry.DecodedFields) *pbgear.Staking_Who {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_Who{}

	switch {
	case matchFields(fields, []int64{0}):
		out.Value = &pbgear.Staking_Who_Id{
			Id: To_Staking_Id(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{95}):
		out.Value = &pbgear.Staking_Who_Index{
			Index: To_Staking_Index(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{13}):
		out.Value = &pbgear.Staking_Who_Raw{
			Raw: To_Staking_Raw(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{1}):
		out.Value = &pbgear.Staking_Who_Address32{
			Address32: To_Staking_Address32(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{96}):
		out.Value = &pbgear.Staking_Who_Address20{
			Address20: To_Staking_Address20(fields), //Passthrough fields...
		}
	}
	return out
}

func To_Staking_WithdrawUnbondedCall(in registry.DecodedFields) *pbgear.Staking_WithdrawUnbondedCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Staking_WithdrawUnbondedCall{}
	out.NumSlashingSpans = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_System_KillPrefixCall(in registry.DecodedFields) *pbgear.System_KillPrefixCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_KillPrefixCall{}
	out.Prefix = To_bytes(fields[0].Value.([]interface{}))
	out.Subkeys = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_System_KillStorageCall(in registry.DecodedFields) *pbgear.System_KillStorageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_KillStorageCall{}
	out.Keys = To_repeated_System_SystemKeysList(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_System_SystemKeysList(f registry.DecodedFields) []*pbgear.System_SystemKeysList {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.System_SystemKeysList, len(fields))
	for _, f := range fields {
		out = append(out, To_System_SystemKeysList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_System_RemarkCall(in registry.DecodedFields) *pbgear.System_RemarkCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_RemarkCall{}
	out.Remark = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_System_RemarkWithEventCall(in registry.DecodedFields) *pbgear.System_RemarkWithEventCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_RemarkWithEventCall{}
	out.Remark = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_System_SetCodeCall(in registry.DecodedFields) *pbgear.System_SetCodeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_SetCodeCall{}
	out.Code = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_System_SetCodeWithoutChecksCall(in registry.DecodedFields) *pbgear.System_SetCodeWithoutChecksCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_SetCodeWithoutChecksCall{}
	out.Code = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_System_SetHeapPagesCall(in registry.DecodedFields) *pbgear.System_SetHeapPagesCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_SetHeapPagesCall{}
	out.Pages = To_uint64(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_System_SetStorageCall(in registry.DecodedFields) *pbgear.System_SetStorageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_SetStorageCall{}
	out.Items = To_repeated_System_TupleSystemItemsListSystemItemsList(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_System_TupleSystemItemsListSystemItemsList(f registry.DecodedFields) []*pbgear.System_TupleSystemItemsListSystemItemsList {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.System_TupleSystemItemsListSystemItemsList, len(fields))
	for _, f := range fields {
		out = append(out, To_System_TupleSystemItemsListSystemItemsList(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_System_SystemKeysList(in registry.DecodedFields) *pbgear.System_SystemKeysList {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_SystemKeysList{}
	out.Keys = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_System_TupleSystemItemsListSystemItemsList(in registry.DecodedFields) *pbgear.System_TupleSystemItemsListSystemItemsList {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	out.Value1 = To_bytes(fields[1].Value.([]interface{}))
	return out
}

func To_Timestamp_SetCall(in registry.DecodedFields) *pbgear.Timestamp_SetCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Timestamp_SetCall{}
	out.Now = To_CompactUint64(fields[0].Value.(types.UCompact))

	return out
}

func To_Treasury_Address20(in registry.DecodedFields) *pbgear.Treasury_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Treasury_Address32(in registry.DecodedFields) *pbgear.Treasury_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Treasury_ApproveProposalCall(in registry.DecodedFields) *pbgear.Treasury_ApproveProposalCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_ApproveProposalCall{}
	out.ProposalId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Treasury_Beneficiary(in registry.DecodedFields) *pbgear.Treasury_Beneficiary {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Treasury_CheckStatusCall(in registry.DecodedFields) *pbgear.Treasury_CheckStatusCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_CheckStatusCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_CompactTupleNull(in types.UCompact) *pbgear.Treasury_CompactTupleNull {

	out := &pbgear.Treasury_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Treasury_Id(in registry.DecodedFields) *pbgear.Treasury_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_Index(in registry.DecodedFields) *pbgear.Treasury_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_Index{}
	out.Value0 = To_Treasury_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Treasury_PayoutCall(in registry.DecodedFields) *pbgear.Treasury_PayoutCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_PayoutCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_ProposeSpendCall(in registry.DecodedFields) *pbgear.Treasury_ProposeSpendCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_ProposeSpendCall{}
	out.Value = To_CompactString(fields[0].Value.(types.UCompact))

	out.Beneficiary = To_Treasury_Beneficiary(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_Raw(in registry.DecodedFields) *pbgear.Treasury_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Treasury_RejectProposalCall(in registry.DecodedFields) *pbgear.Treasury_RejectProposalCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_RejectProposalCall{}
	out.ProposalId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Treasury_RemoveApprovalCall(in registry.DecodedFields) *pbgear.Treasury_RemoveApprovalCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_RemoveApprovalCall{}
	out.ProposalId = To_CompactUint32(fields[0].Value.(types.UCompact))

	return out
}

func To_Treasury_SpendCall(in registry.DecodedFields) *pbgear.Treasury_SpendCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_SpendCall{}
	out.AssetKind = To_Treasury_TupleNull(fields[0].Value.(registry.DecodedFields))

	out.Amount = To_CompactString(fields[1].Value.(types.UCompact))

	out.Beneficiary = To_SpCoreCryptoAccountId32(fields[2].Value.(registry.DecodedFields))

	out.ValidFrom = To_optional_uint32(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_SpendLocalCall(in registry.DecodedFields) *pbgear.Treasury_SpendLocalCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_SpendLocalCall{}
	out.Amount = To_CompactString(fields[0].Value.(types.UCompact))

	out.Beneficiary = To_Treasury_Beneficiary(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Treasury_TupleNull(in registry.DecodedFields) *pbgear.Treasury_TupleNull {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_TupleNull{}
	return out
}

func To_Treasury_VoidSpendCall(in registry.DecodedFields) *pbgear.Treasury_VoidSpendCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Treasury_VoidSpendCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature(in registry.DecodedFields) *pbgear.TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature{}
	out.Value0 = To_FinalityGrandpaPrecommit(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_SpConsensusGrandpaAppSignature(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(in registry.DecodedFields) *pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
	out.Value0 = To_FinalityGrandpaPrevote(fields[0].Value.(registry.DecodedFields))

	out.Value1 = To_SpConsensusGrandpaAppSignature(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_TupleNull() *pbgear.TupleNull {

	out := &pbgear.TupleNull{}
	return out
}

func To_Uint32(in registry.DecodedFields) *pbgear.Uint32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Uint32{}
	out.SlashIndices = To_uint32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_AsDerivativeCall(in registry.DecodedFields) *pbgear.Utility_AsDerivativeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_AsDerivativeCall{}
	out.Index = To_uint32(fields[0].Value.(registry.DecodedFields))

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

func To_Utility_AsOrigin(in registry.DecodedFields) *pbgear.Utility_AsOrigin {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Utility_BatchAllCall(in registry.DecodedFields) *pbgear.Utility_BatchAllCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_BatchAllCall{}
	out.Calls = To_repeated_VaraRuntimeRuntimeCall(fields[0].Value.(registry.DecodedFields))

	return out
}
func To_repeated_VaraRuntimeRuntimeCall(f registry.DecodedFields) []*pbgear.VaraRuntimeRuntimeCall {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.VaraRuntimeRuntimeCall, len(fields))
	for _, f := range fields {
		out = append(out, To_VaraRuntimeRuntimeCall(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_Utility_BatchCall(in registry.DecodedFields) *pbgear.Utility_BatchCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_BatchCall{}
	out.Calls = To_repeated_VaraRuntimeRuntimeCall(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_DispatchAsCall(in registry.DecodedFields) *pbgear.Utility_DispatchAsCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_DispatchAsCall{}
	out.AsOrigin = To_Utility_AsOrigin(fields[0].Value.(registry.DecodedFields))

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

func To_Utility_ForceBatchCall(in registry.DecodedFields) *pbgear.Utility_ForceBatchCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_ForceBatchCall{}
	out.Calls = To_repeated_VaraRuntimeRuntimeCall(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_Origins(in registry.DecodedFields) *pbgear.Utility_Origins {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_Origins{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_System(in registry.DecodedFields) *pbgear.Utility_System {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_System{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_Void(in registry.DecodedFields) *pbgear.Utility_Void {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Utility_Void{}
	out.Value0 = To_Value0(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Utility_WithWeightCall(in registry.DecodedFields) *pbgear.Utility_WithWeightCall {

	fields := []*registry.DecodedField(in)
	_ = fields
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
	out.Weight = To_SpWeightsWeightV2Weight(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Value0(in registry.DecodedFields) *pbgear.Value0 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Value0{}

	switch {
	case matchFields(fields, []int64{}):
		out.Value = &pbgear.Value0_Root{
			Root: To_Root(fields), //Passthrough fields...
		}
	}
	return out
}

func To_VaraRuntimeNposSolution16(in registry.DecodedFields) *pbgear.VaraRuntimeNposSolution16 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.VaraRuntimeNposSolution16{}
	out.Votes1 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields[0].Value.(registry.DecodedFields))

	out.Votes2 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields[1].Value.(registry.DecodedFields))

	out.Votes3 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields[2].Value.(registry.DecodedFields))

	out.Votes4 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields[3].Value.(registry.DecodedFields))

	out.Votes5 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields[4].Value.(registry.DecodedFields))

	out.Votes6 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields[5].Value.(registry.DecodedFields))

	out.Votes7 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields[6].Value.(registry.DecodedFields))

	out.Votes8 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields[7].Value.(registry.DecodedFields))

	out.Votes9 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields[8].Value.(registry.DecodedFields))

	out.Votes10 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields[9].Value.(registry.DecodedFields))

	out.Votes11 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields[10].Value.(registry.DecodedFields))

	out.Votes12 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields[11].Value.(registry.DecodedFields))

	out.Votes13 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields[12].Value.(registry.DecodedFields))

	out.Votes14 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields[13].Value.(registry.DecodedFields))

	out.Votes15 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields[14].Value.(registry.DecodedFields))

	out.Votes16 = To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields[15].Value.(registry.DecodedFields))

	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}
func To_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(f registry.DecodedFields) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
	fields := []*registry.DecodedField(f)
	_ = fields

	out := make([]*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32, len(fields))
	for _, f := range fields {
		out = append(out, To_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(f.Value.([]*registry.DecodedField)))
	}
	return out
}

func To_VaraRuntimeRuntimeCall(in registry.DecodedFields) *pbgear.VaraRuntimeRuntimeCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.VaraRuntimeRuntimeCall{}

	switch {
	case matchFields(fields, []int64{66}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemRemark{
			SystemRemark: To_System_RemarkCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{70}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemSetHeapPages{
			SystemSetHeapPages: To_System_SetHeapPagesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{71}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemSetCode{
			SystemSetCode: To_System_SetCodeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{81}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemSetCodeWithoutChecks{
			SystemSetCodeWithoutChecks: To_System_SetCodeWithoutChecksCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{93}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemSetStorage{
			SystemSetStorage: To_System_SetStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{98}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemKillStorage{
			SystemKillStorage: To_System_KillStorageCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{100}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemKillPrefix{
			SystemKillPrefix: To_System_KillPrefixCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{101}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_SystemRemarkWithEvent{
			SystemRemarkWithEvent: To_System_RemarkWithEventCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{105}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_TimestampSet{
			TimestampSet: To_Timestamp_SetCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{113}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BabeReportEquivocation{
			BabeReportEquivocation: To_Babe_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{116}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BabeReportEquivocationUnsigned{
			BabeReportEquivocationUnsigned: To_Babe_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{118}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BabePlanConfigChange{
			BabePlanConfigChange: To_Babe_PlanConfigChangeCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{124}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_GrandpaReportEquivocation{
			GrandpaReportEquivocation: To_Grandpa_ReportEquivocationCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{129}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_GrandpaReportEquivocationUnsigned{
			GrandpaReportEquivocationUnsigned: To_Grandpa_ReportEquivocationUnsignedCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{132}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_GrandpaNoteStalled{
			GrandpaNoteStalled: To_Grandpa_NoteStalledCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{133}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesTransferAllowDeath{
			BalancesTransferAllowDeath: To_Balances_TransferAllowDeathCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{134}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesForceTransfer{
			BalancesForceTransfer: To_Balances_ForceTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{135}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesTransferKeepAlive{
			BalancesTransferKeepAlive: To_Balances_TransferKeepAliveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{138}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesTransferAll{
			BalancesTransferAll: To_Balances_TransferAllCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{140}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesForceUnreserve{
			BalancesForceUnreserve: To_Balances_ForceUnreserveCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{182}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesUpgradeAccounts{
			BalancesUpgradeAccounts: To_Balances_UpgradeAccountsCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{185}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BalancesForceSetBalance{
			BalancesForceSetBalance: To_Balances_ForceSetBalanceCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{188}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_VestingVest{
			VestingVest: To_Vesting_VestCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{249}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_VestingVestOther{
			VestingVestOther: To_Vesting_VestOtherCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{250}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_VestingVestedTransfer{
			VestingVestedTransfer: To_Vesting_VestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{251}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_VestingForceVestedTransfer{
			VestingForceVestedTransfer: To_Vesting_ForceVestedTransferCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{262}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_VestingMergeSchedules{
			VestingMergeSchedules: To_Vesting_MergeSchedulesCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{267}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BagsListRebag{
			BagsListRebag: To_BagsList_RebagCall(fields), //Passthrough fields...
		}
	case matchFields(fields, []int64{268}):
		out.Calls = &pbgear.VaraRuntimeRuntimeCall_BagsListPutInFrontOf{
			BagsListPutInFrontOf: To_BagsList_PutInFrontOfCall(fields), //Passthrough fields...
		}
	}
	return out
}

func To_VaraRuntimeSessionKeys(in registry.DecodedFields) *pbgear.VaraRuntimeSessionKeys {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.VaraRuntimeSessionKeys{}
	out.Babe = To_SpConsensusBabeAppPublic(fields[0].Value.(registry.DecodedFields))

	out.Grandpa = To_SpConsensusGrandpaAppPublic(fields[1].Value.(registry.DecodedFields))

	out.ImOnline = To_ImOnline_PalletImOnlineSr25519AppSr25519Public(fields[2].Value.(registry.DecodedFields))

	out.AuthorityDiscovery = To_SpAuthorityDiscoveryAppPublic(fields[3].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_Address20(in registry.DecodedFields) *pbgear.Vesting_Address20 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_Address20{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Vesting_Address32(in registry.DecodedFields) *pbgear.Vesting_Address32 {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_Address32{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Vesting_CompactTupleNull(in types.UCompact) *pbgear.Vesting_CompactTupleNull {

	out := &pbgear.Vesting_CompactTupleNull{}
	out.Value = To_TupleNull()
	return out
}

func To_Vesting_ForceVestedTransferCall(in registry.DecodedFields) *pbgear.Vesting_ForceVestedTransferCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_ForceVestedTransferCall{}
	out.Source = To_Vesting_Source(fields[0].Value.(registry.DecodedFields))

	out.Target = To_Vesting_Target(fields[1].Value.(registry.DecodedFields))

	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_Id(in registry.DecodedFields) *pbgear.Vesting_Id {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_Id{}
	out.Value0 = To_SpCoreCryptoAccountId32(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_Index(in registry.DecodedFields) *pbgear.Vesting_Index {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_Index{}
	out.Value0 = To_Vesting_CompactTupleNull(fields[0].Value.(types.UCompact))

	return out
}

func To_Vesting_MergeSchedulesCall(in registry.DecodedFields) *pbgear.Vesting_MergeSchedulesCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_MergeSchedulesCall{}
	out.Schedule1Index = To_uint32(fields[0].Value.(registry.DecodedFields))

	out.Schedule2Index = To_uint32(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_PalletVestingVestingInfoVestingInfo(in registry.DecodedFields) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
	out.Locked = To_string(fields[0].Value.(registry.DecodedFields))

	out.PerBlock = To_string(fields[1].Value.(registry.DecodedFields))

	out.StartingBlock = To_uint32(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_Raw(in registry.DecodedFields) *pbgear.Vesting_Raw {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_Raw{}
	out.Value0 = To_bytes(fields[0].Value.([]interface{}))
	return out
}

func To_Vesting_Source(in registry.DecodedFields) *pbgear.Vesting_Source {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Vesting_Target(in registry.DecodedFields) *pbgear.Vesting_Target {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Vesting_VestCall(in registry.DecodedFields) *pbgear.Vesting_VestCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_VestCall{}
	return out
}

func To_Vesting_VestOtherCall(in registry.DecodedFields) *pbgear.Vesting_VestOtherCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_VestOtherCall{}
	out.Target = To_Vesting_Target(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Vesting_VestedTransferCall(in registry.DecodedFields) *pbgear.Vesting_VestedTransferCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Vesting_VestedTransferCall{}
	out.Target = To_Vesting_Target(fields[0].Value.(registry.DecodedFields))

	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(fields[1].Value.(registry.DecodedFields))

	return out
}

func To_Whitelist_DispatchWhitelistedCallCall(in registry.DecodedFields) *pbgear.Whitelist_DispatchWhitelistedCallCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
	out.CallHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	out.CallEncodedLen = To_uint32(fields[1].Value.(registry.DecodedFields))

	out.CallWeightWitness = To_SpWeightsWeightV2Weight(fields[2].Value.(registry.DecodedFields))

	return out
}

func To_Whitelist_DispatchWhitelistedCallWithPreimageCall(in registry.DecodedFields) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {

	fields := []*registry.DecodedField(in)
	_ = fields
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

func To_Whitelist_RemoveWhitelistedCallCall(in registry.DecodedFields) *pbgear.Whitelist_RemoveWhitelistedCallCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
	out.CallHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func To_Whitelist_WhitelistCallCall(in registry.DecodedFields) *pbgear.Whitelist_WhitelistCallCall {

	fields := []*registry.DecodedField(in)
	_ = fields
	out := &pbgear.Whitelist_WhitelistCallCall{}
	out.CallHash = To_PrimitiveTypesH256(fields[0].Value.(registry.DecodedFields))

	return out
}

func matchFields(f registry.DecodedFields, ids []int64) bool {
	fields := []*registry.DecodedField(f)

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

func To_bytes(f []interface{}) []byte {
	var out []byte
	for _, f := range f {
		u := uint8(f.(types.U8))
		out = append(out, byte(u))
	}
	return out
}

func To_string(f registry.DecodedFields) string {
	fields := []*registry.DecodedField(f)

	return string(To_bytes(fields[0].Value.([]interface{})))
}

func To_uint32(f registry.DecodedFields) uint32 {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(uint32)
}

func To_uint64(f registry.DecodedFields) uint64 {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(uint64)
}

func To_bool(f registry.DecodedFields) bool {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(bool)
}

func To_optional_string(f registry.DecodedFields) *string {
	fields := []*registry.DecodedField(f)

	s := string(To_bytes(fields[0].Value.([]interface{})))
	return &s
}

func To_optional_uint32(f registry.DecodedFields) *uint32 {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(*uint32)
}

func To_optional_uint64(f registry.DecodedFields) *uint64 {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(*uint64)
}

func To_optional_bool(f registry.DecodedFields) *bool {
	fields := []*registry.DecodedField(f)

	return fields[0].Value.(*bool)
}

func To_repeated_uint32(f registry.DecodedFields) []uint32 {
	fields := []*registry.DecodedField(f)

	var out []uint32
	for _, f := range fields {
		out = append(out, f.Value.(uint32))
	}
	return out
}
