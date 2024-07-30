package gen_types

import (
	"reflect"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func To_AllowedSlots(in any) *pbgear.AllowedSlots {
	out := &pbgear.AllowedSlots{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_AllowedSlots_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_AllowedSlots_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.AllowedSlots_PrimarySlots{
			PrimarySlots: To_PrimarySlots(fields),
		}
	case 1:
		target = &pbgear.AllowedSlots_PrimaryAndSecondaryPlainSlots{
			PrimaryAndSecondaryPlainSlots: To_PrimaryAndSecondaryPlainSlots(fields),
		}
	case 2:
		target = &pbgear.AllowedSlots_PrimaryAndSecondaryVrfSlots{
			PrimaryAndSecondaryVrfSlots: To_PrimaryAndSecondaryVrfSlots(fields),
		}
	}
}

func To_AuthorityDiscoveryPallet(in any) *pbgear.AuthorityDiscoveryPallet {
	out := &pbgear.AuthorityDiscoveryPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_AuthorshipPallet(in any) *pbgear.AuthorshipPallet {
	out := &pbgear.AuthorshipPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_BTreeSet(in any) *pbgear.BTreeSet {
	out := &pbgear.BTreeSet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Value0
	out.Value0 = To_Repeated_BTreeSet_value0(fields[0].Value)

	return out //from message
}

func To_Repeated_BTreeSet_value0(in any) []*pbgear.GprimitivesActorId {
	items := in.([]interface{})

	var out []*pbgear.GprimitivesActorId
	for _, item := range items {
		o := To_GprimitivesActorId(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_BabePallet(in any) *pbgear.BabePallet {
	out := &pbgear.BabePallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_BabePallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_BabePallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BabePallet_ReportEquivocationCall{
			ReportEquivocationCall: To_Babe_ReportEquivocationCall(fields),
		}
	case 1:
		target = &pbgear.BabePallet_ReportEquivocationUnsignedCall{
			ReportEquivocationUnsignedCall: To_Babe_ReportEquivocationUnsignedCall(fields),
		}
	case 2:
		target = &pbgear.BabePallet_PlanConfigChangeCall{
			PlanConfigChangeCall: To_Babe_PlanConfigChangeCall(fields),
		}
	}
}

func To_Babe_BabeTrieNodesList(in any) *pbgear.Babe_BabeTrieNodesList {
	out := &pbgear.Babe_BabeTrieNodesList{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field TrieNodes
	out.TrieNodes = To_bytes(fields[0].Value)

	return out //from message
}

func To_Babe_Config(in any) *pbgear.Babe_Config {
	out := &pbgear.Babe_Config{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Babe_Config_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Babe_Config_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 1:
		target = &pbgear.Babe_Config_V1{
			V1: To_Babe_V1(fields),
		}
	}
}

func To_Babe_Consensus(in any) *pbgear.Babe_Consensus {
	out := &pbgear.Babe_Consensus{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_bytes(fields[1].Value)

	return out //from message
}

func To_Babe_Logs(in any) *pbgear.Babe_Logs {
	out := &pbgear.Babe_Logs{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Babe_Logs_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Babe_Logs_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 6:
		target = &pbgear.Babe_Logs_PreRuntime{
			PreRuntime: To_Babe_PreRuntime(fields),
		}
	case 4:
		target = &pbgear.Babe_Logs_Consensus{
			Consensus: To_Babe_Consensus(fields),
		}
	case 5:
		target = &pbgear.Babe_Logs_Seal{
			Seal: To_Babe_Seal(fields),
		}
	case 0:
		target = &pbgear.Babe_Logs_Other{
			Other: To_Babe_Other(fields),
		}
	case 8:
		target = &pbgear.Babe_Logs_RuntimeEnvironmentUpdated{
			RuntimeEnvironmentUpdated: To_Babe_RuntimeEnvironmentUpdated(fields),
		}
	}
}

func To_Babe_Other(in any) *pbgear.Babe_Other {
	out := &pbgear.Babe_Other{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Babe_PlanConfigChangeCall(in any) *pbgear.Babe_PlanConfigChangeCall {
	out := &pbgear.Babe_PlanConfigChangeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Config
	out.Config = To_Babe_PlanConfigChangeCall_config(fields[0].Value)

	return out //from message
}

func To_Babe_PlanConfigChangeCall_config(in any) *pbgear.Babe_Config {
	return nil //simple field
}

func To_Babe_PreRuntime(in any) *pbgear.Babe_PreRuntime {
	out := &pbgear.Babe_PreRuntime{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_bytes(fields[1].Value)

	return out //from message
}

func To_Babe_ReportEquivocationCall(in any) *pbgear.Babe_ReportEquivocationCall {
	out := &pbgear.Babe_ReportEquivocationCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field EquivocationProof
	out.EquivocationProof = To_Babe_ReportEquivocationCall_equivocation_proof(fields[0].Value)
	// field KeyOwnerProof
	out.KeyOwnerProof = To_Babe_ReportEquivocationCall_key_owner_proof(fields[1].Value)

	return out //from message
}

func To_Babe_ReportEquivocationCall_equivocation_proof(in any) *pbgear.SpConsensusSlotsEquivocationProof {
	return nil //simple field
}
func To_Babe_ReportEquivocationCall_key_owner_proof(in any) *pbgear.SpSessionMembershipProof {
	return nil //simple field
}

func To_Babe_ReportEquivocationUnsignedCall(in any) *pbgear.Babe_ReportEquivocationUnsignedCall {
	out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field EquivocationProof
	out.EquivocationProof = To_Babe_ReportEquivocationUnsignedCall_equivocation_proof(fields[0].Value)
	// field KeyOwnerProof
	out.KeyOwnerProof = To_Babe_ReportEquivocationUnsignedCall_key_owner_proof(fields[1].Value)

	return out //from message
}

func To_Babe_ReportEquivocationUnsignedCall_equivocation_proof(in any) *pbgear.SpConsensusSlotsEquivocationProof {
	return nil //simple field
}
func To_Babe_ReportEquivocationUnsignedCall_key_owner_proof(in any) *pbgear.SpSessionMembershipProof {
	return nil //simple field
}

func To_Babe_RuntimeEnvironmentUpdated(in any) *pbgear.Babe_RuntimeEnvironmentUpdated {
	out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Babe_Seal(in any) *pbgear.Babe_Seal {
	out := &pbgear.Babe_Seal{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_bytes(fields[1].Value)

	return out //from message
}

func To_Babe_TupleUint64Uint64(in any) *pbgear.Babe_TupleUint64Uint64 {
	out := &pbgear.Babe_TupleUint64Uint64{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint64(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_uint64(fields[1].Value)

	return out //from message
}

func To_Babe_V1(in any) *pbgear.Babe_V1 {
	out := &pbgear.Babe_V1{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field C
	out.C = To_Babe_V1_c(fields[0].Value)
	// field AllowedSlots
	out.AllowedSlots = To_Babe_V1_allowed_slots(fields[1].Value)

	return out //from message
}

func To_Babe_V1_c(in any) *pbgear.Babe_TupleUint64Uint64 {
	return nil //simple field
}
func To_Babe_V1_allowed_slots(in any) *pbgear.AllowedSlots {
	return nil //simple field
}

func To_BagsListPallet(in any) *pbgear.BagsListPallet {
	out := &pbgear.BagsListPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_BagsListPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_BagsListPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BagsListPallet_RebagCall{
			RebagCall: To_BagsList_RebagCall(fields),
		}
	case 1:
		target = &pbgear.BagsListPallet_PutInFrontOfCall{
			PutInFrontOfCall: To_BagsList_PutInFrontOfCall(fields),
		}
	case 2:
		target = &pbgear.BagsListPallet_PutInFrontOfOtherCall{
			PutInFrontOfOtherCall: To_BagsList_PutInFrontOfOtherCall(fields),
		}
	}
}

func To_BagsList_Address20(in any) *pbgear.BagsList_Address20 {
	out := &pbgear.BagsList_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_BagsList_Address32(in any) *pbgear.BagsList_Address32 {
	out := &pbgear.BagsList_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_BagsList_Dislocated(in any) *pbgear.BagsList_Dislocated {
	out := &pbgear.BagsList_Dislocated{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_BagsList_Dislocated_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_BagsList_Dislocated_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BagsList_Dislocated_Id{
			Id: To_BagsList_Id(fields),
		}
	case 1:
		target = &pbgear.BagsList_Dislocated_Index{
			Index: To_BagsList_Index(fields),
		}
	case 2:
		target = &pbgear.BagsList_Dislocated_Raw{
			Raw: To_BagsList_Raw(fields),
		}
	case 3:
		target = &pbgear.BagsList_Dislocated_Address32{
			Address32: To_BagsList_Address32(fields),
		}
	case 4:
		target = &pbgear.BagsList_Dislocated_Address20{
			Address20: To_BagsList_Address20(fields),
		}
	}
}

func To_BagsList_Heavier(in any) *pbgear.BagsList_Heavier {
	out := &pbgear.BagsList_Heavier{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_BagsList_Heavier_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_BagsList_Heavier_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BagsList_Heavier_Id{
			Id: To_BagsList_Id(fields),
		}
	case 1:
		target = &pbgear.BagsList_Heavier_Index{
			Index: To_BagsList_Index(fields),
		}
	case 2:
		target = &pbgear.BagsList_Heavier_Raw{
			Raw: To_BagsList_Raw(fields),
		}
	case 3:
		target = &pbgear.BagsList_Heavier_Address32{
			Address32: To_BagsList_Address32(fields),
		}
	case 4:
		target = &pbgear.BagsList_Heavier_Address20{
			Address20: To_BagsList_Address20(fields),
		}
	}
}

func To_BagsList_Id(in any) *pbgear.BagsList_Id {
	out := &pbgear.BagsList_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_BagsList_Id_value0(fields[0].Value)

	return out //from message
}

func To_BagsList_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_BagsList_Index(in any) *pbgear.BagsList_Index {
	out := &pbgear.BagsList_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_BagsList_Index_value0(fields[0].Value)

	return out //from message
}

func To_BagsList_Index_value0(in any) *pbgear.BagsList_TupleNull {
	return nil //simple field
}

func To_BagsList_Lighter(in any) *pbgear.BagsList_Lighter {
	out := &pbgear.BagsList_Lighter{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_BagsList_Lighter_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_BagsList_Lighter_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BagsList_Lighter_Id{
			Id: To_BagsList_Id(fields),
		}
	case 1:
		target = &pbgear.BagsList_Lighter_Index{
			Index: To_BagsList_Index(fields),
		}
	case 2:
		target = &pbgear.BagsList_Lighter_Raw{
			Raw: To_BagsList_Raw(fields),
		}
	case 3:
		target = &pbgear.BagsList_Lighter_Address32{
			Address32: To_BagsList_Address32(fields),
		}
	case 4:
		target = &pbgear.BagsList_Lighter_Address20{
			Address20: To_BagsList_Address20(fields),
		}
	}
}

func To_BagsList_PutInFrontOfCall(in any) *pbgear.BagsList_PutInFrontOfCall {
	out := &pbgear.BagsList_PutInFrontOfCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Lighter
	out.Lighter = To_BagsList_PutInFrontOfCall_lighter(fields[0].Value)

	return out //from message
}

func To_BagsList_PutInFrontOfCall_lighter(in any) *pbgear.BagsList_Lighter {
	return nil //simple field
}

func To_BagsList_PutInFrontOfOtherCall(in any) *pbgear.BagsList_PutInFrontOfOtherCall {
	out := &pbgear.BagsList_PutInFrontOfOtherCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Heavier
	out.Heavier = To_BagsList_PutInFrontOfOtherCall_heavier(fields[0].Value)
	// field Lighter
	out.Lighter = To_BagsList_PutInFrontOfOtherCall_lighter(fields[1].Value)

	return out //from message
}

func To_BagsList_PutInFrontOfOtherCall_heavier(in any) *pbgear.BagsList_Heavier {
	return nil //simple field
}
func To_BagsList_PutInFrontOfOtherCall_lighter(in any) *pbgear.BagsList_Lighter {
	return nil //simple field
}

func To_BagsList_Raw(in any) *pbgear.BagsList_Raw {
	out := &pbgear.BagsList_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_BagsList_RebagCall(in any) *pbgear.BagsList_RebagCall {
	out := &pbgear.BagsList_RebagCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Dislocated
	out.Dislocated = To_BagsList_RebagCall_dislocated(fields[0].Value)

	return out //from message
}

func To_BagsList_RebagCall_dislocated(in any) *pbgear.BagsList_Dislocated {
	return nil //simple field
}

func To_BagsList_TupleNull(in any) *pbgear.BagsList_TupleNull {
	out := &pbgear.BagsList_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_BagsList_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_BagsList_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_BalancesPallet(in any) *pbgear.BalancesPallet {
	out := &pbgear.BalancesPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_BalancesPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_BalancesPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BalancesPallet_TransferAllowDeathCall{
			TransferAllowDeathCall: To_Balances_TransferAllowDeathCall(fields),
		}
	case 1:
		target = &pbgear.BalancesPallet_ForceTransferCall{
			ForceTransferCall: To_Balances_ForceTransferCall(fields),
		}
	case 2:
		target = &pbgear.BalancesPallet_TransferKeepAliveCall{
			TransferKeepAliveCall: To_Balances_TransferKeepAliveCall(fields),
		}
	case 3:
		target = &pbgear.BalancesPallet_TransferAllCall{
			TransferAllCall: To_Balances_TransferAllCall(fields),
		}
	case 4:
		target = &pbgear.BalancesPallet_ForceUnreserveCall{
			ForceUnreserveCall: To_Balances_ForceUnreserveCall(fields),
		}
	case 5:
		target = &pbgear.BalancesPallet_UpgradeAccountsCall{
			UpgradeAccountsCall: To_Balances_UpgradeAccountsCall(fields),
		}
	case 6:
		target = &pbgear.BalancesPallet_ForceSetBalanceCall{
			ForceSetBalanceCall: To_Balances_ForceSetBalanceCall(fields),
		}
	}
}

func To_Balances_Address20(in any) *pbgear.Balances_Address20 {
	out := &pbgear.Balances_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Balances_Address32(in any) *pbgear.Balances_Address32 {
	out := &pbgear.Balances_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Balances_Dest(in any) *pbgear.Balances_Dest {
	out := &pbgear.Balances_Dest{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Balances_Dest_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Balances_Dest_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Balances_Dest_Id{
			Id: To_Balances_Id(fields),
		}
	case 1:
		target = &pbgear.Balances_Dest_Index{
			Index: To_Balances_Index(fields),
		}
	case 2:
		target = &pbgear.Balances_Dest_Raw{
			Raw: To_Balances_Raw(fields),
		}
	case 3:
		target = &pbgear.Balances_Dest_Address32{
			Address32: To_Balances_Address32(fields),
		}
	case 4:
		target = &pbgear.Balances_Dest_Address20{
			Address20: To_Balances_Address20(fields),
		}
	}
}

func To_Balances_ForceSetBalanceCall(in any) *pbgear.Balances_ForceSetBalanceCall {
	out := &pbgear.Balances_ForceSetBalanceCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_Balances_ForceSetBalanceCall_who(fields[0].Value)
	// primitive field NewFree
	out.NewFree = To_string(fields[1].Value)

	return out //from message
}

func To_Balances_ForceSetBalanceCall_who(in any) *pbgear.Balances_Who {
	return nil //simple field
}

func To_Balances_ForceTransferCall(in any) *pbgear.Balances_ForceTransferCall {
	out := &pbgear.Balances_ForceTransferCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Source
	out.Source = To_Balances_ForceTransferCall_source(fields[0].Value)
	// field Dest
	out.Dest = To_Balances_ForceTransferCall_dest(fields[1].Value)
	// primitive field Value
	out.Value = To_string(fields[2].Value)

	return out //from message
}

func To_Balances_ForceTransferCall_source(in any) *pbgear.Balances_Source {
	return nil //simple field
}
func To_Balances_ForceTransferCall_dest(in any) *pbgear.Balances_Dest {
	return nil //simple field
}

func To_Balances_ForceUnreserveCall(in any) *pbgear.Balances_ForceUnreserveCall {
	out := &pbgear.Balances_ForceUnreserveCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_Balances_ForceUnreserveCall_who(fields[0].Value)
	// primitive field Amount
	out.Amount = To_string(fields[1].Value)

	return out //from message
}

func To_Balances_ForceUnreserveCall_who(in any) *pbgear.Balances_Who {
	return nil //simple field
}

func To_Balances_Id(in any) *pbgear.Balances_Id {
	out := &pbgear.Balances_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Balances_Id_value0(fields[0].Value)

	return out //from message
}

func To_Balances_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Balances_Index(in any) *pbgear.Balances_Index {
	out := &pbgear.Balances_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Balances_Index_value0(fields[0].Value)

	return out //from message
}

func To_Balances_Index_value0(in any) *pbgear.Balances_TupleNull {
	return nil //simple field
}

func To_Balances_Raw(in any) *pbgear.Balances_Raw {
	out := &pbgear.Balances_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Balances_Source(in any) *pbgear.Balances_Source {
	out := &pbgear.Balances_Source{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Balances_Source_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Balances_Source_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Balances_Source_Id{
			Id: To_Balances_Id(fields),
		}
	case 1:
		target = &pbgear.Balances_Source_Index{
			Index: To_Balances_Index(fields),
		}
	case 2:
		target = &pbgear.Balances_Source_Raw{
			Raw: To_Balances_Raw(fields),
		}
	case 3:
		target = &pbgear.Balances_Source_Address32{
			Address32: To_Balances_Address32(fields),
		}
	case 4:
		target = &pbgear.Balances_Source_Address20{
			Address20: To_Balances_Address20(fields),
		}
	}
}

func To_Balances_TransferAllCall(in any) *pbgear.Balances_TransferAllCall {
	out := &pbgear.Balances_TransferAllCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Dest
	out.Dest = To_Balances_TransferAllCall_dest(fields[0].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[1].Value)

	return out //from message
}

func To_Balances_TransferAllCall_dest(in any) *pbgear.Balances_Dest {
	return nil //simple field
}

func To_Balances_TransferAllowDeathCall(in any) *pbgear.Balances_TransferAllowDeathCall {
	out := &pbgear.Balances_TransferAllowDeathCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Dest
	out.Dest = To_Balances_TransferAllowDeathCall_dest(fields[0].Value)
	// primitive field Value
	out.Value = To_string(fields[1].Value)

	return out //from message
}

func To_Balances_TransferAllowDeathCall_dest(in any) *pbgear.Balances_Dest {
	return nil //simple field
}

func To_Balances_TransferKeepAliveCall(in any) *pbgear.Balances_TransferKeepAliveCall {
	out := &pbgear.Balances_TransferKeepAliveCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Dest
	out.Dest = To_Balances_TransferKeepAliveCall_dest(fields[0].Value)
	// primitive field Value
	out.Value = To_string(fields[1].Value)

	return out //from message
}

func To_Balances_TransferKeepAliveCall_dest(in any) *pbgear.Balances_Dest {
	return nil //simple field
}

func To_Balances_TupleNull(in any) *pbgear.Balances_TupleNull {
	out := &pbgear.Balances_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Balances_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Balances_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Balances_UpgradeAccountsCall(in any) *pbgear.Balances_UpgradeAccountsCall {
	out := &pbgear.Balances_UpgradeAccountsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Who
	out.Who = To_Repeated_Balances_UpgradeAccountsCall_who(fields[0].Value)

	return out //from message
}

func To_Repeated_Balances_UpgradeAccountsCall_who(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Balances_Who(in any) *pbgear.Balances_Who {
	out := &pbgear.Balances_Who{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Balances_Who_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Balances_Who_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Balances_Who_Id{
			Id: To_Balances_Id(fields),
		}
	case 1:
		target = &pbgear.Balances_Who_Index{
			Index: To_Balances_Index(fields),
		}
	case 2:
		target = &pbgear.Balances_Who_Raw{
			Raw: To_Balances_Raw(fields),
		}
	case 3:
		target = &pbgear.Balances_Who_Address32{
			Address32: To_Balances_Address32(fields),
		}
	case 4:
		target = &pbgear.Balances_Who_Address20{
			Address20: To_Balances_Address20(fields),
		}
	}
}

func To_BoundedCollectionsBoundedVecBoundedVec(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	out := &pbgear.BoundedCollectionsBoundedVecBoundedVec{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_BountiesPallet(in any) *pbgear.BountiesPallet {
	out := &pbgear.BountiesPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_BountiesPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_BountiesPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.BountiesPallet_ProposeBountyCall{
			ProposeBountyCall: To_Bounties_ProposeBountyCall(fields),
		}
	case 1:
		target = &pbgear.BountiesPallet_ApproveBountyCall{
			ApproveBountyCall: To_Bounties_ApproveBountyCall(fields),
		}
	case 2:
		target = &pbgear.BountiesPallet_ProposeCuratorCall{
			ProposeCuratorCall: To_Bounties_ProposeCuratorCall(fields),
		}
	case 3:
		target = &pbgear.BountiesPallet_UnassignCuratorCall{
			UnassignCuratorCall: To_Bounties_UnassignCuratorCall(fields),
		}
	case 4:
		target = &pbgear.BountiesPallet_AcceptCuratorCall{
			AcceptCuratorCall: To_Bounties_AcceptCuratorCall(fields),
		}
	case 5:
		target = &pbgear.BountiesPallet_AwardBountyCall{
			AwardBountyCall: To_Bounties_AwardBountyCall(fields),
		}
	case 6:
		target = &pbgear.BountiesPallet_ClaimBountyCall{
			ClaimBountyCall: To_Bounties_ClaimBountyCall(fields),
		}
	case 7:
		target = &pbgear.BountiesPallet_CloseBountyCall{
			CloseBountyCall: To_Bounties_CloseBountyCall(fields),
		}
	case 8:
		target = &pbgear.BountiesPallet_ExtendBountyExpiryCall{
			ExtendBountyExpiryCall: To_Bounties_ExtendBountyExpiryCall(fields),
		}
	}
}

func To_Bounties_AcceptCuratorCall(in any) *pbgear.Bounties_AcceptCuratorCall {
	out := &pbgear.Bounties_AcceptCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Bounties_Address20(in any) *pbgear.Bounties_Address20 {
	out := &pbgear.Bounties_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Bounties_Address32(in any) *pbgear.Bounties_Address32 {
	out := &pbgear.Bounties_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Bounties_ApproveBountyCall(in any) *pbgear.Bounties_ApproveBountyCall {
	out := &pbgear.Bounties_ApproveBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Bounties_AwardBountyCall(in any) *pbgear.Bounties_AwardBountyCall {
	out := &pbgear.Bounties_AwardBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)
	// field Beneficiary
	out.Beneficiary = To_Bounties_AwardBountyCall_beneficiary(fields[1].Value)

	return out //from message
}

func To_Bounties_AwardBountyCall_beneficiary(in any) *pbgear.Bounties_Beneficiary {
	return nil //simple field
}

func To_Bounties_Beneficiary(in any) *pbgear.Bounties_Beneficiary {
	out := &pbgear.Bounties_Beneficiary{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Bounties_Beneficiary_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Bounties_Beneficiary_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Bounties_Beneficiary_Id{
			Id: To_Bounties_Id(fields),
		}
	case 1:
		target = &pbgear.Bounties_Beneficiary_Index{
			Index: To_Bounties_Index(fields),
		}
	case 2:
		target = &pbgear.Bounties_Beneficiary_Raw{
			Raw: To_Bounties_Raw(fields),
		}
	case 3:
		target = &pbgear.Bounties_Beneficiary_Address32{
			Address32: To_Bounties_Address32(fields),
		}
	case 4:
		target = &pbgear.Bounties_Beneficiary_Address20{
			Address20: To_Bounties_Address20(fields),
		}
	}
}

func To_Bounties_ClaimBountyCall(in any) *pbgear.Bounties_ClaimBountyCall {
	out := &pbgear.Bounties_ClaimBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Bounties_CloseBountyCall(in any) *pbgear.Bounties_CloseBountyCall {
	out := &pbgear.Bounties_CloseBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Bounties_Curator(in any) *pbgear.Bounties_Curator {
	out := &pbgear.Bounties_Curator{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Bounties_Curator_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Bounties_Curator_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Bounties_Curator_Id{
			Id: To_Bounties_Id(fields),
		}
	case 1:
		target = &pbgear.Bounties_Curator_Index{
			Index: To_Bounties_Index(fields),
		}
	case 2:
		target = &pbgear.Bounties_Curator_Raw{
			Raw: To_Bounties_Raw(fields),
		}
	case 3:
		target = &pbgear.Bounties_Curator_Address32{
			Address32: To_Bounties_Address32(fields),
		}
	case 4:
		target = &pbgear.Bounties_Curator_Address20{
			Address20: To_Bounties_Address20(fields),
		}
	}
}

func To_Bounties_ExtendBountyExpiryCall(in any) *pbgear.Bounties_ExtendBountyExpiryCall {
	out := &pbgear.Bounties_ExtendBountyExpiryCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)
	// primitive field Remark
	out.Remark = To_bytes(fields[1].Value)

	return out //from message
}

func To_Bounties_Id(in any) *pbgear.Bounties_Id {
	out := &pbgear.Bounties_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Bounties_Id_value0(fields[0].Value)

	return out //from message
}

func To_Bounties_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Bounties_Index(in any) *pbgear.Bounties_Index {
	out := &pbgear.Bounties_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Bounties_Index_value0(fields[0].Value)

	return out //from message
}

func To_Bounties_Index_value0(in any) *pbgear.Bounties_TupleNull {
	return nil //simple field
}

func To_Bounties_ProposeBountyCall(in any) *pbgear.Bounties_ProposeBountyCall {
	out := &pbgear.Bounties_ProposeBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)
	// primitive field Description
	out.Description = To_bytes(fields[1].Value)

	return out //from message
}

func To_Bounties_ProposeCuratorCall(in any) *pbgear.Bounties_ProposeCuratorCall {
	out := &pbgear.Bounties_ProposeCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)
	// field Curator
	out.Curator = To_Bounties_ProposeCuratorCall_curator(fields[1].Value)
	// primitive field Fee
	out.Fee = To_string(fields[2].Value)

	return out //from message
}

func To_Bounties_ProposeCuratorCall_curator(in any) *pbgear.Bounties_Curator {
	return nil //simple field
}

func To_Bounties_Raw(in any) *pbgear.Bounties_Raw {
	out := &pbgear.Bounties_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Bounties_TupleNull(in any) *pbgear.Bounties_TupleNull {
	out := &pbgear.Bounties_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Bounties_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Bounties_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Bounties_UnassignCuratorCall(in any) *pbgear.Bounties_UnassignCuratorCall {
	out := &pbgear.Bounties_UnassignCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BountyId
	out.BountyId = To_uint32(fields[0].Value)

	return out //from message
}

func To_ChildBountiesPallet(in any) *pbgear.ChildBountiesPallet {
	out := &pbgear.ChildBountiesPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ChildBountiesPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ChildBountiesPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ChildBountiesPallet_AddChildBountyCall{
			AddChildBountyCall: To_ChildBounties_AddChildBountyCall(fields),
		}
	case 1:
		target = &pbgear.ChildBountiesPallet_ProposeCuratorCall{
			ProposeCuratorCall: To_ChildBounties_ProposeCuratorCall(fields),
		}
	case 2:
		target = &pbgear.ChildBountiesPallet_AcceptCuratorCall{
			AcceptCuratorCall: To_ChildBounties_AcceptCuratorCall(fields),
		}
	case 3:
		target = &pbgear.ChildBountiesPallet_UnassignCuratorCall{
			UnassignCuratorCall: To_ChildBounties_UnassignCuratorCall(fields),
		}
	case 4:
		target = &pbgear.ChildBountiesPallet_AwardChildBountyCall{
			AwardChildBountyCall: To_ChildBounties_AwardChildBountyCall(fields),
		}
	case 5:
		target = &pbgear.ChildBountiesPallet_ClaimChildBountyCall{
			ClaimChildBountyCall: To_ChildBounties_ClaimChildBountyCall(fields),
		}
	case 6:
		target = &pbgear.ChildBountiesPallet_CloseChildBountyCall{
			CloseChildBountyCall: To_ChildBounties_CloseChildBountyCall(fields),
		}
	}
}

func To_ChildBounties_AcceptCuratorCall(in any) *pbgear.ChildBounties_AcceptCuratorCall {
	out := &pbgear.ChildBounties_AcceptCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)

	return out //from message
}

func To_ChildBounties_AddChildBountyCall(in any) *pbgear.ChildBounties_AddChildBountyCall {
	out := &pbgear.ChildBounties_AddChildBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field Value
	out.Value = To_string(fields[1].Value)
	// primitive field Description
	out.Description = To_bytes(fields[2].Value)

	return out //from message
}

func To_ChildBounties_Address20(in any) *pbgear.ChildBounties_Address20 {
	out := &pbgear.ChildBounties_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ChildBounties_Address32(in any) *pbgear.ChildBounties_Address32 {
	out := &pbgear.ChildBounties_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ChildBounties_AwardChildBountyCall(in any) *pbgear.ChildBounties_AwardChildBountyCall {
	out := &pbgear.ChildBounties_AwardChildBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)
	// field Beneficiary
	out.Beneficiary = To_ChildBounties_AwardChildBountyCall_beneficiary(fields[2].Value)

	return out //from message
}

func To_ChildBounties_AwardChildBountyCall_beneficiary(in any) *pbgear.ChildBounties_Beneficiary {
	return nil //simple field
}

func To_ChildBounties_Beneficiary(in any) *pbgear.ChildBounties_Beneficiary {
	out := &pbgear.ChildBounties_Beneficiary{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ChildBounties_Beneficiary_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ChildBounties_Beneficiary_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ChildBounties_Beneficiary_Id{
			Id: To_ChildBounties_Id(fields),
		}
	case 1:
		target = &pbgear.ChildBounties_Beneficiary_Index{
			Index: To_ChildBounties_Index(fields),
		}
	case 2:
		target = &pbgear.ChildBounties_Beneficiary_Raw{
			Raw: To_ChildBounties_Raw(fields),
		}
	case 3:
		target = &pbgear.ChildBounties_Beneficiary_Address32{
			Address32: To_ChildBounties_Address32(fields),
		}
	case 4:
		target = &pbgear.ChildBounties_Beneficiary_Address20{
			Address20: To_ChildBounties_Address20(fields),
		}
	}
}

func To_ChildBounties_ClaimChildBountyCall(in any) *pbgear.ChildBounties_ClaimChildBountyCall {
	out := &pbgear.ChildBounties_ClaimChildBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)

	return out //from message
}

func To_ChildBounties_CloseChildBountyCall(in any) *pbgear.ChildBounties_CloseChildBountyCall {
	out := &pbgear.ChildBounties_CloseChildBountyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)

	return out //from message
}

func To_ChildBounties_Curator(in any) *pbgear.ChildBounties_Curator {
	out := &pbgear.ChildBounties_Curator{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ChildBounties_Curator_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ChildBounties_Curator_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ChildBounties_Curator_Id{
			Id: To_ChildBounties_Id(fields),
		}
	case 1:
		target = &pbgear.ChildBounties_Curator_Index{
			Index: To_ChildBounties_Index(fields),
		}
	case 2:
		target = &pbgear.ChildBounties_Curator_Raw{
			Raw: To_ChildBounties_Raw(fields),
		}
	case 3:
		target = &pbgear.ChildBounties_Curator_Address32{
			Address32: To_ChildBounties_Address32(fields),
		}
	case 4:
		target = &pbgear.ChildBounties_Curator_Address20{
			Address20: To_ChildBounties_Address20(fields),
		}
	}
}

func To_ChildBounties_Id(in any) *pbgear.ChildBounties_Id {
	out := &pbgear.ChildBounties_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ChildBounties_Id_value0(fields[0].Value)

	return out //from message
}

func To_ChildBounties_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_ChildBounties_Index(in any) *pbgear.ChildBounties_Index {
	out := &pbgear.ChildBounties_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ChildBounties_Index_value0(fields[0].Value)

	return out //from message
}

func To_ChildBounties_Index_value0(in any) *pbgear.ChildBounties_TupleNull {
	return nil //simple field
}

func To_ChildBounties_ProposeCuratorCall(in any) *pbgear.ChildBounties_ProposeCuratorCall {
	out := &pbgear.ChildBounties_ProposeCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)
	// field Curator
	out.Curator = To_ChildBounties_ProposeCuratorCall_curator(fields[2].Value)
	// primitive field Fee
	out.Fee = To_string(fields[3].Value)

	return out //from message
}

func To_ChildBounties_ProposeCuratorCall_curator(in any) *pbgear.ChildBounties_Curator {
	return nil //simple field
}

func To_ChildBounties_Raw(in any) *pbgear.ChildBounties_Raw {
	out := &pbgear.ChildBounties_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ChildBounties_TupleNull(in any) *pbgear.ChildBounties_TupleNull {
	out := &pbgear.ChildBounties_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_ChildBounties_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_ChildBounties_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_ChildBounties_UnassignCuratorCall(in any) *pbgear.ChildBounties_UnassignCuratorCall {
	out := &pbgear.ChildBounties_UnassignCuratorCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(fields[0].Value)
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(fields[1].Value)

	return out //from message
}

func To_ConvictionVotingPallet(in any) *pbgear.ConvictionVotingPallet {
	out := &pbgear.ConvictionVotingPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ConvictionVotingPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ConvictionVotingPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ConvictionVotingPallet_VoteCall{
			VoteCall: To_ConvictionVoting_VoteCall(fields),
		}
	case 1:
		target = &pbgear.ConvictionVotingPallet_DelegateCall{
			DelegateCall: To_ConvictionVoting_DelegateCall(fields),
		}
	case 2:
		target = &pbgear.ConvictionVotingPallet_UndelegateCall{
			UndelegateCall: To_ConvictionVoting_UndelegateCall(fields),
		}
	case 3:
		target = &pbgear.ConvictionVotingPallet_UnlockCall{
			UnlockCall: To_ConvictionVoting_UnlockCall(fields),
		}
	case 4:
		target = &pbgear.ConvictionVotingPallet_RemoveVoteCall{
			RemoveVoteCall: To_ConvictionVoting_RemoveVoteCall(fields),
		}
	case 5:
		target = &pbgear.ConvictionVotingPallet_RemoveOtherVoteCall{
			RemoveOtherVoteCall: To_ConvictionVoting_RemoveOtherVoteCall(fields),
		}
	}
}

func To_ConvictionVoting_Address20(in any) *pbgear.ConvictionVoting_Address20 {
	out := &pbgear.ConvictionVoting_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_Address32(in any) *pbgear.ConvictionVoting_Address32 {
	out := &pbgear.ConvictionVoting_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_Conviction(in any) *pbgear.ConvictionVoting_Conviction {
	out := &pbgear.ConvictionVoting_Conviction{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ConvictionVoting_Conviction_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ConvictionVoting_Conviction_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ConvictionVoting_Conviction_None{
			None: To_ConvictionVoting_None(fields),
		}
	case 1:
		target = &pbgear.ConvictionVoting_Conviction_Locked1X{
			Locked1X: To_ConvictionVoting_Locked1X(fields),
		}
	case 2:
		target = &pbgear.ConvictionVoting_Conviction_Locked2X{
			Locked2X: To_ConvictionVoting_Locked2X(fields),
		}
	case 3:
		target = &pbgear.ConvictionVoting_Conviction_Locked3X{
			Locked3X: To_ConvictionVoting_Locked3X(fields),
		}
	case 4:
		target = &pbgear.ConvictionVoting_Conviction_Locked4X{
			Locked4X: To_ConvictionVoting_Locked4X(fields),
		}
	case 5:
		target = &pbgear.ConvictionVoting_Conviction_Locked5X{
			Locked5X: To_ConvictionVoting_Locked5X(fields),
		}
	case 6:
		target = &pbgear.ConvictionVoting_Conviction_Locked6X{
			Locked6X: To_ConvictionVoting_Locked6X(fields),
		}
	}
}

func To_ConvictionVoting_DelegateCall(in any) *pbgear.ConvictionVoting_DelegateCall {
	out := &pbgear.ConvictionVoting_DelegateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Class
	out.Class = To_uint32(fields[0].Value)
	// field To
	out.To = To_ConvictionVoting_DelegateCall_to(fields[1].Value)
	// field Conviction
	out.Conviction = To_ConvictionVoting_DelegateCall_conviction(fields[2].Value)
	// primitive field Balance
	out.Balance = To_string(fields[3].Value)

	return out //from message
}

func To_ConvictionVoting_DelegateCall_to(in any) *pbgear.ConvictionVoting_To {
	return nil //simple field
}
func To_ConvictionVoting_DelegateCall_conviction(in any) *pbgear.ConvictionVoting_Conviction {
	return nil //simple field
}

func To_ConvictionVoting_Id(in any) *pbgear.ConvictionVoting_Id {
	out := &pbgear.ConvictionVoting_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ConvictionVoting_Id_value0(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_ConvictionVoting_Index(in any) *pbgear.ConvictionVoting_Index {
	out := &pbgear.ConvictionVoting_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ConvictionVoting_Index_value0(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_Index_value0(in any) *pbgear.ConvictionVoting_TupleNull {
	return nil //simple field
}

func To_ConvictionVoting_Locked1X(in any) *pbgear.ConvictionVoting_Locked1X {
	out := &pbgear.ConvictionVoting_Locked1X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_Locked2X(in any) *pbgear.ConvictionVoting_Locked2X {
	out := &pbgear.ConvictionVoting_Locked2X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_Locked3X(in any) *pbgear.ConvictionVoting_Locked3X {
	out := &pbgear.ConvictionVoting_Locked3X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_Locked4X(in any) *pbgear.ConvictionVoting_Locked4X {
	out := &pbgear.ConvictionVoting_Locked4X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_Locked5X(in any) *pbgear.ConvictionVoting_Locked5X {
	out := &pbgear.ConvictionVoting_Locked5X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_Locked6X(in any) *pbgear.ConvictionVoting_Locked6X {
	out := &pbgear.ConvictionVoting_Locked6X{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_None(in any) *pbgear.ConvictionVoting_None {
	out := &pbgear.ConvictionVoting_None{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_ConvictionVoting_PalletConvictionVotingVoteVote(in any) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
	out := &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_Raw(in any) *pbgear.ConvictionVoting_Raw {
	out := &pbgear.ConvictionVoting_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_RemoveOtherVoteCall(in any) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
	out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Target
	out.Target = To_ConvictionVoting_RemoveOtherVoteCall_target(fields[0].Value)
	// primitive field Class
	out.Class = To_uint32(fields[1].Value)
	// primitive field Index
	out.Index = To_uint32(fields[2].Value)

	return out //from message
}

func To_ConvictionVoting_RemoveOtherVoteCall_target(in any) *pbgear.ConvictionVoting_Target {
	return nil //simple field
}

func To_ConvictionVoting_RemoveVoteCall(in any) *pbgear.ConvictionVoting_RemoveVoteCall {
	out := &pbgear.ConvictionVoting_RemoveVoteCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Class
	out.Class = To_Optional_uint32(fields[0].Value)
	// primitive field Index
	out.Index = To_uint32(fields[1].Value)

	return out //from message
}

func To_ConvictionVoting_Split(in any) *pbgear.ConvictionVoting_Split {
	out := &pbgear.ConvictionVoting_Split{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Aye
	out.Aye = To_string(fields[0].Value)
	// primitive field Nay
	out.Nay = To_string(fields[1].Value)

	return out //from message
}

func To_ConvictionVoting_SplitAbstain(in any) *pbgear.ConvictionVoting_SplitAbstain {
	out := &pbgear.ConvictionVoting_SplitAbstain{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Aye
	out.Aye = To_string(fields[0].Value)
	// primitive field Nay
	out.Nay = To_string(fields[1].Value)
	// primitive field Abstain
	out.Abstain = To_string(fields[2].Value)

	return out //from message
}

func To_ConvictionVoting_Standard(in any) *pbgear.ConvictionVoting_Standard {
	out := &pbgear.ConvictionVoting_Standard{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Vote
	out.Vote = To_ConvictionVoting_Standard_vote(fields[0].Value)
	// primitive field Balance
	out.Balance = To_string(fields[1].Value)

	return out //from message
}

func To_ConvictionVoting_Standard_vote(in any) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
	return nil //simple field
}

func To_ConvictionVoting_Target(in any) *pbgear.ConvictionVoting_Target {
	out := &pbgear.ConvictionVoting_Target{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ConvictionVoting_Target_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ConvictionVoting_Target_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ConvictionVoting_Target_Id{
			Id: To_ConvictionVoting_Id(fields),
		}
	case 1:
		target = &pbgear.ConvictionVoting_Target_Index{
			Index: To_ConvictionVoting_Index(fields),
		}
	case 2:
		target = &pbgear.ConvictionVoting_Target_Raw{
			Raw: To_ConvictionVoting_Raw(fields),
		}
	case 3:
		target = &pbgear.ConvictionVoting_Target_Address32{
			Address32: To_ConvictionVoting_Address32(fields),
		}
	case 4:
		target = &pbgear.ConvictionVoting_Target_Address20{
			Address20: To_ConvictionVoting_Address20(fields),
		}
	}
}

func To_ConvictionVoting_To(in any) *pbgear.ConvictionVoting_To {
	out := &pbgear.ConvictionVoting_To{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ConvictionVoting_To_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ConvictionVoting_To_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ConvictionVoting_To_Id{
			Id: To_ConvictionVoting_Id(fields),
		}
	case 1:
		target = &pbgear.ConvictionVoting_To_Index{
			Index: To_ConvictionVoting_Index(fields),
		}
	case 2:
		target = &pbgear.ConvictionVoting_To_Raw{
			Raw: To_ConvictionVoting_Raw(fields),
		}
	case 3:
		target = &pbgear.ConvictionVoting_To_Address32{
			Address32: To_ConvictionVoting_Address32(fields),
		}
	case 4:
		target = &pbgear.ConvictionVoting_To_Address20{
			Address20: To_ConvictionVoting_Address20(fields),
		}
	}
}

func To_ConvictionVoting_TupleNull(in any) *pbgear.ConvictionVoting_TupleNull {
	out := &pbgear.ConvictionVoting_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_ConvictionVoting_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_ConvictionVoting_UndelegateCall(in any) *pbgear.ConvictionVoting_UndelegateCall {
	out := &pbgear.ConvictionVoting_UndelegateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Class
	out.Class = To_uint32(fields[0].Value)

	return out //from message
}

func To_ConvictionVoting_UnlockCall(in any) *pbgear.ConvictionVoting_UnlockCall {
	out := &pbgear.ConvictionVoting_UnlockCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Class
	out.Class = To_uint32(fields[0].Value)
	// field Target
	out.Target = To_ConvictionVoting_UnlockCall_target(fields[1].Value)

	return out //from message
}

func To_ConvictionVoting_UnlockCall_target(in any) *pbgear.ConvictionVoting_Target {
	return nil //simple field
}

func To_ConvictionVoting_Vote(in any) *pbgear.ConvictionVoting_Vote {
	out := &pbgear.ConvictionVoting_Vote{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_ConvictionVoting_Vote_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_ConvictionVoting_Vote_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ConvictionVoting_Vote_Standard{
			Standard: To_ConvictionVoting_Standard(fields),
		}
	case 1:
		target = &pbgear.ConvictionVoting_Vote_Split{
			Split: To_ConvictionVoting_Split(fields),
		}
	case 2:
		target = &pbgear.ConvictionVoting_Vote_SplitAbstain{
			SplitAbstain: To_ConvictionVoting_SplitAbstain(fields),
		}
	}
}

func To_ConvictionVoting_VoteCall(in any) *pbgear.ConvictionVoting_VoteCall {
	out := &pbgear.ConvictionVoting_VoteCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PollIndex
	out.PollIndex = To_uint32(fields[0].Value)
	// field Vote
	out.Vote = To_ConvictionVoting_VoteCall_vote(fields[1].Value)

	return out //from message
}

func To_ConvictionVoting_VoteCall_vote(in any) *pbgear.ConvictionVoting_Vote {
	return nil //simple field
}

func To_ElectionProviderMultiPhasePallet(in any) *pbgear.ElectionProviderMultiPhasePallet {
	out := &pbgear.ElectionProviderMultiPhasePallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ElectionProviderMultiPhasePallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ElectionProviderMultiPhasePallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ElectionProviderMultiPhasePallet_SubmitUnsignedCall{
			SubmitUnsignedCall: To_ElectionProviderMultiPhase_SubmitUnsignedCall(fields),
		}
	case 1:
		target = &pbgear.ElectionProviderMultiPhasePallet_SetMinimumUntrustedScoreCall{
			SetMinimumUntrustedScoreCall: To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(fields),
		}
	case 2:
		target = &pbgear.ElectionProviderMultiPhasePallet_SetEmergencyElectionResultCall{
			SetEmergencyElectionResultCall: To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(fields),
		}
	case 3:
		target = &pbgear.ElectionProviderMultiPhasePallet_SubmitCall{
			SubmitCall: To_ElectionProviderMultiPhase_SubmitCall(fields),
		}
	case 4:
		target = &pbgear.ElectionProviderMultiPhasePallet_GovernanceFallbackCall{
			GovernanceFallbackCall: To_ElectionProviderMultiPhase_GovernanceFallbackCall(fields),
		}
	}
}

func To_ElectionProviderMultiPhase_GovernanceFallbackCall(in any) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
	out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field MaybeMaxVoters
	out.MaybeMaxVoters = To_Optional_uint32(fields[0].Value)
	// primitive field MaybeMaxTargets
	out.MaybeMaxTargets = To_Optional_uint32(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Solution
	out.Solution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution_solution(fields[0].Value)
	// field Score
	out.Score = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution_score(fields[1].Value)
	// primitive field Round
	out.Round = To_uint32(fields[2].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution_solution(in any) *pbgear.VaraRuntimeNposSolution16 {
	return nil //simple field
}
func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution_score(in any) *pbgear.SpNposElectionsElectionScore {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Voters
	out.Voters = To_uint32(fields[0].Value)
	// primitive field Targets
	out.Targets = To_uint32(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(in any) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
	out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Supports
	out.Supports = To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(fields[0].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(in any) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
	out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// optional field MaybeNextScore
	out.MaybeNextScore = To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(fields[0].Value)

	return out //from message
}

func To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(in any) *pbgear.SpNposElectionsElectionScore {
	return nil //funcForOptionalField
}

func To_ElectionProviderMultiPhase_SubmitCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field RawSolution
	out.RawSolution = To_ElectionProviderMultiPhase_SubmitCall_raw_solution(fields[0].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_SubmitCall_raw_solution(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_SubmitUnsignedCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field RawSolution
	out.RawSolution = To_ElectionProviderMultiPhase_SubmitUnsignedCall_raw_solution(fields[0].Value)
	// field Witness
	out.Witness = To_ElectionProviderMultiPhase_SubmitUnsignedCall_witness(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_SubmitUnsignedCall_raw_solution(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	return nil //simple field
}
func To_ElectionProviderMultiPhase_SubmitUnsignedCall_witness(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport_value1(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}
func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport_value1(in any) *pbgear.SpNposElectionsSupport {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String_value0(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_string(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// field Value1
	out.Value1 = To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16_value1(fields[1].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16_value1(in any) *pbgear.SpArithmeticPerThingsPerU16 {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// field Value1
	out.Value1 = To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32_value1(fields[1].Value)
	// primitive field Value2
	out.Value2 = To_uint32(fields[2].Value)

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32_value1(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	return nil //simple field
}

func To_ElectionProviderMultiPhase_TupleUint32Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32Uint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_uint32(fields[1].Value)

	return out //from message
}

func To_FellowshipCollectivePallet(in any) *pbgear.FellowshipCollectivePallet {
	out := &pbgear.FellowshipCollectivePallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_FellowshipCollectivePallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipCollectivePallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipCollectivePallet_AddMemberCall{
			AddMemberCall: To_FellowshipCollective_AddMemberCall(fields),
		}
	case 1:
		target = &pbgear.FellowshipCollectivePallet_PromoteMemberCall{
			PromoteMemberCall: To_FellowshipCollective_PromoteMemberCall(fields),
		}
	case 2:
		target = &pbgear.FellowshipCollectivePallet_DemoteMemberCall{
			DemoteMemberCall: To_FellowshipCollective_DemoteMemberCall(fields),
		}
	case 3:
		target = &pbgear.FellowshipCollectivePallet_RemoveMemberCall{
			RemoveMemberCall: To_FellowshipCollective_RemoveMemberCall(fields),
		}
	case 4:
		target = &pbgear.FellowshipCollectivePallet_VoteCall{
			VoteCall: To_FellowshipCollective_VoteCall(fields),
		}
	case 5:
		target = &pbgear.FellowshipCollectivePallet_CleanupPollCall{
			CleanupPollCall: To_FellowshipCollective_CleanupPollCall(fields),
		}
	}
}

func To_FellowshipCollective_AddMemberCall(in any) *pbgear.FellowshipCollective_AddMemberCall {
	out := &pbgear.FellowshipCollective_AddMemberCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_FellowshipCollective_AddMemberCall_who(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_AddMemberCall_who(in any) *pbgear.FellowshipCollective_Who {
	return nil //simple field
}

func To_FellowshipCollective_Address20(in any) *pbgear.FellowshipCollective_Address20 {
	out := &pbgear.FellowshipCollective_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_Address32(in any) *pbgear.FellowshipCollective_Address32 {
	out := &pbgear.FellowshipCollective_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_CleanupPollCall(in any) *pbgear.FellowshipCollective_CleanupPollCall {
	out := &pbgear.FellowshipCollective_CleanupPollCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PollIndex
	out.PollIndex = To_uint32(fields[0].Value)
	// primitive field Max
	out.Max = To_uint32(fields[1].Value)

	return out //from message
}

func To_FellowshipCollective_DemoteMemberCall(in any) *pbgear.FellowshipCollective_DemoteMemberCall {
	out := &pbgear.FellowshipCollective_DemoteMemberCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_FellowshipCollective_DemoteMemberCall_who(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_DemoteMemberCall_who(in any) *pbgear.FellowshipCollective_Who {
	return nil //simple field
}

func To_FellowshipCollective_Id(in any) *pbgear.FellowshipCollective_Id {
	out := &pbgear.FellowshipCollective_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipCollective_Id_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_FellowshipCollective_Index(in any) *pbgear.FellowshipCollective_Index {
	out := &pbgear.FellowshipCollective_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipCollective_Index_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_Index_value0(in any) *pbgear.FellowshipCollective_TupleNull {
	return nil //simple field
}

func To_FellowshipCollective_PromoteMemberCall(in any) *pbgear.FellowshipCollective_PromoteMemberCall {
	out := &pbgear.FellowshipCollective_PromoteMemberCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_FellowshipCollective_PromoteMemberCall_who(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_PromoteMemberCall_who(in any) *pbgear.FellowshipCollective_Who {
	return nil //simple field
}

func To_FellowshipCollective_Raw(in any) *pbgear.FellowshipCollective_Raw {
	out := &pbgear.FellowshipCollective_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_RemoveMemberCall(in any) *pbgear.FellowshipCollective_RemoveMemberCall {
	out := &pbgear.FellowshipCollective_RemoveMemberCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Who
	out.Who = To_FellowshipCollective_RemoveMemberCall_who(fields[0].Value)
	// primitive field MinRank
	out.MinRank = To_uint32(fields[1].Value)

	return out //from message
}

func To_FellowshipCollective_RemoveMemberCall_who(in any) *pbgear.FellowshipCollective_Who {
	return nil //simple field
}

func To_FellowshipCollective_TupleNull(in any) *pbgear.FellowshipCollective_TupleNull {
	out := &pbgear.FellowshipCollective_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_FellowshipCollective_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_FellowshipCollective_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_FellowshipCollective_VoteCall(in any) *pbgear.FellowshipCollective_VoteCall {
	out := &pbgear.FellowshipCollective_VoteCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Poll
	out.Poll = To_uint32(fields[0].Value)
	// primitive field Aye
	out.Aye = To_bool(fields[1].Value)

	return out //from message
}

func To_FellowshipCollective_Who(in any) *pbgear.FellowshipCollective_Who {
	out := &pbgear.FellowshipCollective_Who{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_FellowshipCollective_Who_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipCollective_Who_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipCollective_Who_Id{
			Id: To_FellowshipCollective_Id(fields),
		}
	case 1:
		target = &pbgear.FellowshipCollective_Who_Index{
			Index: To_FellowshipCollective_Index(fields),
		}
	case 2:
		target = &pbgear.FellowshipCollective_Who_Raw{
			Raw: To_FellowshipCollective_Raw(fields),
		}
	case 3:
		target = &pbgear.FellowshipCollective_Who_Address32{
			Address32: To_FellowshipCollective_Address32(fields),
		}
	case 4:
		target = &pbgear.FellowshipCollective_Who_Address20{
			Address20: To_FellowshipCollective_Address20(fields),
		}
	}
}

func To_FellowshipReferendaPallet(in any) *pbgear.FellowshipReferendaPallet {
	out := &pbgear.FellowshipReferendaPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_FellowshipReferendaPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipReferendaPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipReferendaPallet_SubmitCall{
			SubmitCall: To_FellowshipReferenda_SubmitCall(fields),
		}
	case 1:
		target = &pbgear.FellowshipReferendaPallet_PlaceDecisionDepositCall{
			PlaceDecisionDepositCall: To_FellowshipReferenda_PlaceDecisionDepositCall(fields),
		}
	case 2:
		target = &pbgear.FellowshipReferendaPallet_RefundDecisionDepositCall{
			RefundDecisionDepositCall: To_FellowshipReferenda_RefundDecisionDepositCall(fields),
		}
	case 3:
		target = &pbgear.FellowshipReferendaPallet_CancelCall{
			CancelCall: To_FellowshipReferenda_CancelCall(fields),
		}
	case 4:
		target = &pbgear.FellowshipReferendaPallet_KillCall{
			KillCall: To_FellowshipReferenda_KillCall(fields),
		}
	case 5:
		target = &pbgear.FellowshipReferendaPallet_NudgeReferendumCall{
			NudgeReferendumCall: To_FellowshipReferenda_NudgeReferendumCall(fields),
		}
	case 6:
		target = &pbgear.FellowshipReferendaPallet_OneFewerDecidingCall{
			OneFewerDecidingCall: To_FellowshipReferenda_OneFewerDecidingCall(fields),
		}
	case 7:
		target = &pbgear.FellowshipReferendaPallet_RefundSubmissionDepositCall{
			RefundSubmissionDepositCall: To_FellowshipReferenda_RefundSubmissionDepositCall(fields),
		}
	case 8:
		target = &pbgear.FellowshipReferendaPallet_SetMetadataCall{
			SetMetadataCall: To_FellowshipReferenda_SetMetadataCall(fields),
		}
	}
}

func To_FellowshipReferenda_After(in any) *pbgear.FellowshipReferenda_After {
	out := &pbgear.FellowshipReferenda_After{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_At(in any) *pbgear.FellowshipReferenda_At {
	out := &pbgear.FellowshipReferenda_At{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_CancelCall(in any) *pbgear.FellowshipReferenda_CancelCall {
	out := &pbgear.FellowshipReferenda_CancelCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_EnactmentMoment(in any) *pbgear.FellowshipReferenda_EnactmentMoment {
	out := &pbgear.FellowshipReferenda_EnactmentMoment{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_FellowshipReferenda_EnactmentMoment_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipReferenda_EnactmentMoment_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipReferenda_EnactmentMoment_At{
			At: To_FellowshipReferenda_At(fields),
		}
	case 1:
		target = &pbgear.FellowshipReferenda_EnactmentMoment_After{
			After: To_FellowshipReferenda_After(fields),
		}
	}
}

func To_FellowshipReferenda_Inline(in any) *pbgear.FellowshipReferenda_Inline {
	out := &pbgear.FellowshipReferenda_Inline{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipReferenda_Inline_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Inline_value0(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	return nil //simple field
}

func To_FellowshipReferenda_KillCall(in any) *pbgear.FellowshipReferenda_KillCall {
	out := &pbgear.FellowshipReferenda_KillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Legacy(in any) *pbgear.FellowshipReferenda_Legacy {
	out := &pbgear.FellowshipReferenda_Legacy{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_FellowshipReferenda_Legacy_hash(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Legacy_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_FellowshipReferenda_Lookup(in any) *pbgear.FellowshipReferenda_Lookup {
	out := &pbgear.FellowshipReferenda_Lookup{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_FellowshipReferenda_Lookup_hash(fields[0].Value)
	// primitive field Len
	out.Len = To_uint32(fields[1].Value)

	return out //from message
}

func To_FellowshipReferenda_Lookup_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_FellowshipReferenda_NudgeReferendumCall(in any) *pbgear.FellowshipReferenda_NudgeReferendumCall {
	out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_OneFewerDecidingCall(in any) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
	out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Track
	out.Track = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Origins(in any) *pbgear.FellowshipReferenda_Origins {
	out := &pbgear.FellowshipReferenda_Origins{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipReferenda_Origins_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Origins_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_FellowshipReferenda_PlaceDecisionDepositCall(in any) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Proposal(in any) *pbgear.FellowshipReferenda_Proposal {
	out := &pbgear.FellowshipReferenda_Proposal{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_FellowshipReferenda_Proposal_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipReferenda_Proposal_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipReferenda_Proposal_Legacy{
			Legacy: To_FellowshipReferenda_Legacy(fields),
		}
	case 1:
		target = &pbgear.FellowshipReferenda_Proposal_Inline{
			Inline: To_FellowshipReferenda_Inline(fields),
		}
	case 2:
		target = &pbgear.FellowshipReferenda_Proposal_Lookup{
			Lookup: To_FellowshipReferenda_Lookup(fields),
		}
	}
}

func To_FellowshipReferenda_ProposalOrigin(in any) *pbgear.FellowshipReferenda_ProposalOrigin {
	out := &pbgear.FellowshipReferenda_ProposalOrigin{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_FellowshipReferenda_ProposalOrigin_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_FellowshipReferenda_ProposalOrigin_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.FellowshipReferenda_ProposalOrigin_System{
			System: To_FellowshipReferenda_System(fields),
		}
	case 20:
		target = &pbgear.FellowshipReferenda_ProposalOrigin_Origins{
			Origins: To_FellowshipReferenda_Origins(fields),
		}
	case 2:
		target = &pbgear.FellowshipReferenda_ProposalOrigin_Void{
			Void: To_FellowshipReferenda_Void(fields),
		}
	}
}

func To_FellowshipReferenda_RefundDecisionDepositCall(in any) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_RefundSubmissionDepositCall(in any) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_SetMetadataCall(in any) *pbgear.FellowshipReferenda_SetMetadataCall {
	out := &pbgear.FellowshipReferenda_SetMetadataCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// optional field MaybeHash
	out.MaybeHash = To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(fields[1].Value)

	return out //from message
}

func To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //funcForOptionalField
}

func To_FellowshipReferenda_SubmitCall(in any) *pbgear.FellowshipReferenda_SubmitCall {
	out := &pbgear.FellowshipReferenda_SubmitCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ProposalOrigin
	out.ProposalOrigin = To_FellowshipReferenda_SubmitCall_proposal_origin(fields[0].Value)
	// field Proposal
	out.Proposal = To_FellowshipReferenda_SubmitCall_proposal(fields[1].Value)
	// field EnactmentMoment
	out.EnactmentMoment = To_FellowshipReferenda_SubmitCall_enactment_moment(fields[2].Value)

	return out //from message
}

func To_FellowshipReferenda_SubmitCall_proposal_origin(in any) *pbgear.FellowshipReferenda_ProposalOrigin {
	return nil //simple field
}
func To_FellowshipReferenda_SubmitCall_proposal(in any) *pbgear.FellowshipReferenda_Proposal {
	return nil //simple field
}
func To_FellowshipReferenda_SubmitCall_enactment_moment(in any) *pbgear.FellowshipReferenda_EnactmentMoment {
	return nil //simple field
}

func To_FellowshipReferenda_System(in any) *pbgear.FellowshipReferenda_System {
	out := &pbgear.FellowshipReferenda_System{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipReferenda_System_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_System_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_FellowshipReferenda_Void(in any) *pbgear.FellowshipReferenda_Void {
	out := &pbgear.FellowshipReferenda_Void{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_FellowshipReferenda_Void_value0(fields[0].Value)

	return out //from message
}

func To_FellowshipReferenda_Void_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_FinalityGrandpaEquivocation(in any) *pbgear.FinalityGrandpaEquivocation {
	out := &pbgear.FinalityGrandpaEquivocation{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field RoundNumber
	out.RoundNumber = To_uint64(fields[0].Value)
	// field Identity
	out.Identity = To_FinalityGrandpaEquivocation_identity(fields[1].Value)
	// field First
	out.First = To_FinalityGrandpaEquivocation_first(fields[2].Value)
	// field Second
	out.Second = To_FinalityGrandpaEquivocation_second(fields[3].Value)

	return out //from message
}

func To_FinalityGrandpaEquivocation_identity(in any) *pbgear.SpConsensusGrandpaAppPublic {
	return nil //simple field
}
func To_FinalityGrandpaEquivocation_first(in any) *pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	return nil //simple field
}
func To_FinalityGrandpaEquivocation_second(in any) *pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	return nil //simple field
}

func To_FinalityGrandpaPrecommit(in any) *pbgear.FinalityGrandpaPrecommit {
	out := &pbgear.FinalityGrandpaPrecommit{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field TargetHash
	out.TargetHash = To_FinalityGrandpaPrecommit_target_hash(fields[0].Value)
	// primitive field TargetNumber
	out.TargetNumber = To_uint32(fields[1].Value)

	return out //from message
}

func To_FinalityGrandpaPrecommit_target_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_FinalityGrandpaPrevote(in any) *pbgear.FinalityGrandpaPrevote {
	out := &pbgear.FinalityGrandpaPrevote{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field TargetHash
	out.TargetHash = To_FinalityGrandpaPrevote_target_hash(fields[0].Value)
	// primitive field TargetNumber
	out.TargetNumber = To_uint32(fields[1].Value)

	return out //from message
}

func To_FinalityGrandpaPrevote_target_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_GearBankPallet(in any) *pbgear.GearBankPallet {
	out := &pbgear.GearBankPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearBuiltinPallet(in any) *pbgear.GearBuiltinPallet {
	out := &pbgear.GearBuiltinPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearGasPallet(in any) *pbgear.GearGasPallet {
	out := &pbgear.GearGasPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearMessengerPallet(in any) *pbgear.GearMessengerPallet {
	out := &pbgear.GearMessengerPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearPallet(in any) *pbgear.GearPallet {
	out := &pbgear.GearPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_GearPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_GearPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.GearPallet_UploadCodeCall{
			UploadCodeCall: To_Gear_UploadCodeCall(fields),
		}
	case 1:
		target = &pbgear.GearPallet_UploadProgramCall{
			UploadProgramCall: To_Gear_UploadProgramCall(fields),
		}
	case 2:
		target = &pbgear.GearPallet_CreateProgramCall{
			CreateProgramCall: To_Gear_CreateProgramCall(fields),
		}
	case 3:
		target = &pbgear.GearPallet_SendMessageCall{
			SendMessageCall: To_Gear_SendMessageCall(fields),
		}
	case 4:
		target = &pbgear.GearPallet_SendReplyCall{
			SendReplyCall: To_Gear_SendReplyCall(fields),
		}
	case 5:
		target = &pbgear.GearPallet_ClaimValueCall{
			ClaimValueCall: To_Gear_ClaimValueCall(fields),
		}
	case 6:
		target = &pbgear.GearPallet_RunCall{
			RunCall: To_Gear_RunCall(fields),
		}
	case 7:
		target = &pbgear.GearPallet_SetExecuteInherentCall{
			SetExecuteInherentCall: To_Gear_SetExecuteInherentCall(fields),
		}
	}
}

func To_GearPaymentPallet(in any) *pbgear.GearPaymentPallet {
	out := &pbgear.GearPaymentPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearProgramPallet(in any) *pbgear.GearProgramPallet {
	out := &pbgear.GearProgramPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearSchedulerPallet(in any) *pbgear.GearSchedulerPallet {
	out := &pbgear.GearSchedulerPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearVoucherPallet(in any) *pbgear.GearVoucherPallet {
	out := &pbgear.GearVoucherPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_GearVoucherPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_GearVoucherPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.GearVoucherPallet_IssueCall{
			IssueCall: To_GearVoucher_IssueCall(fields),
		}
	case 1:
		target = &pbgear.GearVoucherPallet_CallCall{
			CallCall: To_GearVoucher_CallCall(fields),
		}
	case 2:
		target = &pbgear.GearVoucherPallet_RevokeCall{
			RevokeCall: To_GearVoucher_RevokeCall(fields),
		}
	case 3:
		target = &pbgear.GearVoucherPallet_UpdateCall{
			UpdateCall: To_GearVoucher_UpdateCall(fields),
		}
	case 4:
		target = &pbgear.GearVoucherPallet_CallDeprecatedCall{
			CallDeprecatedCall: To_GearVoucher_CallDeprecatedCall(fields),
		}
	case 5:
		target = &pbgear.GearVoucherPallet_DeclineCall{
			DeclineCall: To_GearVoucher_DeclineCall(fields),
		}
	}
}

func To_GearVoucher_AppendPrograms(in any) *pbgear.GearVoucher_AppendPrograms {
	out := &pbgear.GearVoucher_AppendPrograms{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_GearVoucher_AppendPrograms_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_GearVoucher_AppendPrograms_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.GearVoucher_AppendPrograms_None{
			None: To_GearVoucher_None(fields),
		}
	case 1:
		target = &pbgear.GearVoucher_AppendPrograms_Some{
			Some: To_GearVoucher_Some(fields),
		}
	}
}

func To_GearVoucher_Call(in any) *pbgear.GearVoucher_Call {
	out := &pbgear.GearVoucher_Call{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_GearVoucher_Call_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_GearVoucher_Call_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.GearVoucher_Call_SendMessage{
			SendMessage: To_GearVoucher_SendMessage(fields),
		}
	case 1:
		target = &pbgear.GearVoucher_Call_SendReply{
			SendReply: To_GearVoucher_SendReply(fields),
		}
	case 2:
		target = &pbgear.GearVoucher_Call_UploadCode{
			UploadCode: To_GearVoucher_UploadCode(fields),
		}
	case 3:
		target = &pbgear.GearVoucher_Call_DeclineVoucher{
			DeclineVoucher: To_GearVoucher_DeclineVoucher(fields),
		}
	}
}

func To_GearVoucher_CallCall(in any) *pbgear.GearVoucher_CallCall {
	out := &pbgear.GearVoucher_CallCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field VoucherId
	out.VoucherId = To_GearVoucher_CallCall_voucher_id(fields[0].Value)
	// field Call
	out.Call = To_GearVoucher_CallCall_call(fields[1].Value)

	return out //from message
}

func To_GearVoucher_CallCall_voucher_id(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return nil //simple field
}
func To_GearVoucher_CallCall_call(in any) *pbgear.GearVoucher_Call {
	return nil //simple field
}

func To_GearVoucher_CallDeprecatedCall(in any) *pbgear.GearVoucher_CallDeprecatedCall {
	out := &pbgear.GearVoucher_CallDeprecatedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Call
	out.Call = To_GearVoucher_CallDeprecatedCall_call(fields[0].Value)

	return out //from message
}

func To_GearVoucher_CallDeprecatedCall_call(in any) *pbgear.GearVoucher_Call {
	return nil //simple field
}

func To_GearVoucher_DeclineCall(in any) *pbgear.GearVoucher_DeclineCall {
	out := &pbgear.GearVoucher_DeclineCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field VoucherId
	out.VoucherId = To_GearVoucher_DeclineCall_voucher_id(fields[0].Value)

	return out //from message
}

func To_GearVoucher_DeclineCall_voucher_id(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return nil //simple field
}

func To_GearVoucher_DeclineVoucher(in any) *pbgear.GearVoucher_DeclineVoucher {
	out := &pbgear.GearVoucher_DeclineVoucher{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearVoucher_IssueCall(in any) *pbgear.GearVoucher_IssueCall {
	out := &pbgear.GearVoucher_IssueCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Spender
	out.Spender = To_GearVoucher_IssueCall_spender(fields[0].Value)
	// primitive field Balance
	out.Balance = To_string(fields[1].Value)
	// optional field Programs
	out.Programs = To_Optional_GearVoucher_IssueCall_programs(fields[2].Value)
	// primitive field CodeUploading
	out.CodeUploading = To_bool(fields[3].Value)
	// primitive field Duration
	out.Duration = To_uint32(fields[4].Value)

	return out //from message
}

func To_GearVoucher_IssueCall_spender(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}
func To_Optional_GearVoucher_IssueCall_programs(in any) *pbgear.BTreeSet {
	return nil //funcForOptionalField
}

func To_GearVoucher_None(in any) *pbgear.GearVoucher_None {
	out := &pbgear.GearVoucher_None{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_GearVoucher_PalletGearVoucherInternalVoucherId(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_GearVoucher_RevokeCall(in any) *pbgear.GearVoucher_RevokeCall {
	out := &pbgear.GearVoucher_RevokeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Spender
	out.Spender = To_GearVoucher_RevokeCall_spender(fields[0].Value)
	// field VoucherId
	out.VoucherId = To_GearVoucher_RevokeCall_voucher_id(fields[1].Value)

	return out //from message
}

func To_GearVoucher_RevokeCall_spender(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}
func To_GearVoucher_RevokeCall_voucher_id(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return nil //simple field
}

func To_GearVoucher_SendMessage(in any) *pbgear.GearVoucher_SendMessage {
	out := &pbgear.GearVoucher_SendMessage{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Destination
	out.Destination = To_GearVoucher_SendMessage_destination(fields[0].Value)
	// primitive field Payload
	out.Payload = To_bytes(fields[1].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[2].Value)
	// primitive field Value
	out.Value = To_string(fields[3].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[4].Value)

	return out //from message
}

func To_GearVoucher_SendMessage_destination(in any) *pbgear.GprimitivesActorId {
	return nil //simple field
}

func To_GearVoucher_SendReply(in any) *pbgear.GearVoucher_SendReply {
	out := &pbgear.GearVoucher_SendReply{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ReplyToId
	out.ReplyToId = To_GearVoucher_SendReply_reply_to_id(fields[0].Value)
	// primitive field Payload
	out.Payload = To_bytes(fields[1].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[2].Value)
	// primitive field Value
	out.Value = To_string(fields[3].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[4].Value)

	return out //from message
}

func To_GearVoucher_SendReply_reply_to_id(in any) *pbgear.GprimitivesMessageId {
	return nil //simple field
}

func To_GearVoucher_Some(in any) *pbgear.GearVoucher_Some {
	out := &pbgear.GearVoucher_Some{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_GearVoucher_Some_value0(fields[0].Value)

	return out //from message
}

func To_GearVoucher_Some_value0(in any) *pbgear.BTreeSet {
	return nil //simple field
}

func To_GearVoucher_UpdateCall(in any) *pbgear.GearVoucher_UpdateCall {
	out := &pbgear.GearVoucher_UpdateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Spender
	out.Spender = To_GearVoucher_UpdateCall_spender(fields[0].Value)
	// field VoucherId
	out.VoucherId = To_GearVoucher_UpdateCall_voucher_id(fields[1].Value)
	// optional field MoveOwnership
	out.MoveOwnership = To_Optional_GearVoucher_UpdateCall_move_ownership(fields[2].Value)
	// primitive field BalanceTopUp
	out.BalanceTopUp = To_Optional_string(fields[3].Value)
	// optional field AppendPrograms
	out.AppendPrograms = To_Optional_GearVoucher_UpdateCall_append_programs(fields[4].Value)
	// primitive field CodeUploading
	out.CodeUploading = To_Optional_bool(fields[5].Value)
	// primitive field ProlongDuration
	out.ProlongDuration = To_Optional_uint32(fields[6].Value)

	return out //from message
}

func To_GearVoucher_UpdateCall_spender(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}
func To_GearVoucher_UpdateCall_voucher_id(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return nil //simple field
}
func To_Optional_GearVoucher_UpdateCall_move_ownership(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //funcForOptionalField
}
func To_Optional_GearVoucher_UpdateCall_append_programs(in any) *pbgear.GearVoucher_AppendPrograms {
	return nil //funcForOptionalField
}

func To_GearVoucher_UploadCode(in any) *pbgear.GearVoucher_UploadCode {
	out := &pbgear.GearVoucher_UploadCode{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Code
	out.Code = To_bytes(fields[0].Value)

	return out //from message
}

func To_Gear_ClaimValueCall(in any) *pbgear.Gear_ClaimValueCall {
	out := &pbgear.Gear_ClaimValueCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MessageId
	out.MessageId = To_Gear_ClaimValueCall_message_id(fields[0].Value)

	return out //from message
}

func To_Gear_ClaimValueCall_message_id(in any) *pbgear.GprimitivesMessageId {
	return nil //simple field
}

func To_Gear_CreateProgramCall(in any) *pbgear.Gear_CreateProgramCall {
	out := &pbgear.Gear_CreateProgramCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field CodeId
	out.CodeId = To_Gear_CreateProgramCall_code_id(fields[0].Value)
	// primitive field Salt
	out.Salt = To_bytes(fields[1].Value)
	// primitive field InitPayload
	out.InitPayload = To_bytes(fields[2].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[3].Value)
	// primitive field Value
	out.Value = To_string(fields[4].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[5].Value)

	return out //from message
}

func To_Gear_CreateProgramCall_code_id(in any) *pbgear.GprimitivesCodeId {
	return nil //simple field
}

func To_Gear_RunCall(in any) *pbgear.Gear_RunCall {
	out := &pbgear.Gear_RunCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field MaxGas
	out.MaxGas = To_Optional_uint64(fields[0].Value)

	return out //from message
}

func To_Gear_SendMessageCall(in any) *pbgear.Gear_SendMessageCall {
	out := &pbgear.Gear_SendMessageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Destination
	out.Destination = To_Gear_SendMessageCall_destination(fields[0].Value)
	// primitive field Payload
	out.Payload = To_bytes(fields[1].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[2].Value)
	// primitive field Value
	out.Value = To_string(fields[3].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[4].Value)

	return out //from message
}

func To_Gear_SendMessageCall_destination(in any) *pbgear.GprimitivesActorId {
	return nil //simple field
}

func To_Gear_SendReplyCall(in any) *pbgear.Gear_SendReplyCall {
	out := &pbgear.Gear_SendReplyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ReplyToId
	out.ReplyToId = To_Gear_SendReplyCall_reply_to_id(fields[0].Value)
	// primitive field Payload
	out.Payload = To_bytes(fields[1].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[2].Value)
	// primitive field Value
	out.Value = To_string(fields[3].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[4].Value)

	return out //from message
}

func To_Gear_SendReplyCall_reply_to_id(in any) *pbgear.GprimitivesMessageId {
	return nil //simple field
}

func To_Gear_SetExecuteInherentCall(in any) *pbgear.Gear_SetExecuteInherentCall {
	out := &pbgear.Gear_SetExecuteInherentCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_bool(fields[0].Value)

	return out //from message
}

func To_Gear_UploadCodeCall(in any) *pbgear.Gear_UploadCodeCall {
	out := &pbgear.Gear_UploadCodeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Code
	out.Code = To_bytes(fields[0].Value)

	return out //from message
}

func To_Gear_UploadProgramCall(in any) *pbgear.Gear_UploadProgramCall {
	out := &pbgear.Gear_UploadProgramCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Code
	out.Code = To_bytes(fields[0].Value)
	// primitive field Salt
	out.Salt = To_bytes(fields[1].Value)
	// primitive field InitPayload
	out.InitPayload = To_bytes(fields[2].Value)
	// primitive field GasLimit
	out.GasLimit = To_uint64(fields[3].Value)
	// primitive field Value
	out.Value = To_string(fields[4].Value)
	// primitive field KeepAlive
	out.KeepAlive = To_bool(fields[5].Value)

	return out //from message
}

func To_GprimitivesActorId(in any) *pbgear.GprimitivesActorId {
	out := &pbgear.GprimitivesActorId{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_GprimitivesCodeId(in any) *pbgear.GprimitivesCodeId {
	out := &pbgear.GprimitivesCodeId{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_GprimitivesMessageId(in any) *pbgear.GprimitivesMessageId {
	out := &pbgear.GprimitivesMessageId{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_GrandpaPallet(in any) *pbgear.GrandpaPallet {
	out := &pbgear.GrandpaPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_GrandpaPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_GrandpaPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.GrandpaPallet_ReportEquivocationCall{
			ReportEquivocationCall: To_Grandpa_ReportEquivocationCall(fields),
		}
	case 1:
		target = &pbgear.GrandpaPallet_ReportEquivocationUnsignedCall{
			ReportEquivocationUnsignedCall: To_Grandpa_ReportEquivocationUnsignedCall(fields),
		}
	case 2:
		target = &pbgear.GrandpaPallet_NoteStalledCall{
			NoteStalledCall: To_Grandpa_NoteStalledCall(fields),
		}
	}
}

func To_Grandpa_Equivocation(in any) *pbgear.Grandpa_Equivocation {
	out := &pbgear.Grandpa_Equivocation{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Grandpa_Equivocation_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Grandpa_Equivocation_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Grandpa_Equivocation_Prevote{
			Prevote: To_Grandpa_Prevote(fields),
		}
	case 1:
		target = &pbgear.Grandpa_Equivocation_Precommit{
			Precommit: To_Grandpa_Precommit(fields),
		}
	}
}

func To_Grandpa_GrandpaTrieNodesList(in any) *pbgear.Grandpa_GrandpaTrieNodesList {
	out := &pbgear.Grandpa_GrandpaTrieNodesList{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field TrieNodes
	out.TrieNodes = To_bytes(fields[0].Value)

	return out //from message
}

func To_Grandpa_NoteStalledCall(in any) *pbgear.Grandpa_NoteStalledCall {
	out := &pbgear.Grandpa_NoteStalledCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Delay
	out.Delay = To_uint32(fields[0].Value)
	// primitive field BestFinalizedBlockNumber
	out.BestFinalizedBlockNumber = To_uint32(fields[1].Value)

	return out //from message
}

func To_Grandpa_Precommit(in any) *pbgear.Grandpa_Precommit {
	out := &pbgear.Grandpa_Precommit{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Grandpa_Precommit_value0(fields[0].Value)

	return out //from message
}

func To_Grandpa_Precommit_value0(in any) *pbgear.FinalityGrandpaEquivocation {
	return nil //simple field
}

func To_Grandpa_Prevote(in any) *pbgear.Grandpa_Prevote {
	out := &pbgear.Grandpa_Prevote{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Grandpa_Prevote_value0(fields[0].Value)

	return out //from message
}

func To_Grandpa_Prevote_value0(in any) *pbgear.FinalityGrandpaEquivocation {
	return nil //simple field
}

func To_Grandpa_ReportEquivocationCall(in any) *pbgear.Grandpa_ReportEquivocationCall {
	out := &pbgear.Grandpa_ReportEquivocationCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field EquivocationProof
	out.EquivocationProof = To_Grandpa_ReportEquivocationCall_equivocation_proof(fields[0].Value)
	// field KeyOwnerProof
	out.KeyOwnerProof = To_Grandpa_ReportEquivocationCall_key_owner_proof(fields[1].Value)

	return out //from message
}

func To_Grandpa_ReportEquivocationCall_equivocation_proof(in any) *pbgear.SpConsensusGrandpaEquivocationProof {
	return nil //simple field
}
func To_Grandpa_ReportEquivocationCall_key_owner_proof(in any) *pbgear.SpSessionMembershipProof {
	return nil //simple field
}

func To_Grandpa_ReportEquivocationUnsignedCall(in any) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
	out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field EquivocationProof
	out.EquivocationProof = To_Grandpa_ReportEquivocationUnsignedCall_equivocation_proof(fields[0].Value)
	// field KeyOwnerProof
	out.KeyOwnerProof = To_Grandpa_ReportEquivocationUnsignedCall_key_owner_proof(fields[1].Value)

	return out //from message
}

func To_Grandpa_ReportEquivocationUnsignedCall_equivocation_proof(in any) *pbgear.SpConsensusGrandpaEquivocationProof {
	return nil //simple field
}
func To_Grandpa_ReportEquivocationUnsignedCall_key_owner_proof(in any) *pbgear.SpSessionMembershipProof {
	return nil //simple field
}

func To_HistoricalPallet(in any) *pbgear.HistoricalPallet {
	out := &pbgear.HistoricalPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_IdentityPallet(in any) *pbgear.IdentityPallet {
	out := &pbgear.IdentityPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_IdentityPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_IdentityPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.IdentityPallet_AddRegistrarCall{
			AddRegistrarCall: To_Identity_AddRegistrarCall(fields),
		}
	case 1:
		target = &pbgear.IdentityPallet_SetIdentityCall{
			SetIdentityCall: To_Identity_SetIdentityCall(fields),
		}
	case 2:
		target = &pbgear.IdentityPallet_SetSubsCall{
			SetSubsCall: To_Identity_SetSubsCall(fields),
		}
	case 3:
		target = &pbgear.IdentityPallet_ClearIdentityCall{
			ClearIdentityCall: To_Identity_ClearIdentityCall(fields),
		}
	case 4:
		target = &pbgear.IdentityPallet_RequestJudgementCall{
			RequestJudgementCall: To_Identity_RequestJudgementCall(fields),
		}
	case 5:
		target = &pbgear.IdentityPallet_CancelRequestCall{
			CancelRequestCall: To_Identity_CancelRequestCall(fields),
		}
	case 6:
		target = &pbgear.IdentityPallet_SetFeeCall{
			SetFeeCall: To_Identity_SetFeeCall(fields),
		}
	case 7:
		target = &pbgear.IdentityPallet_SetAccountIdCall{
			SetAccountIdCall: To_Identity_SetAccountIdCall(fields),
		}
	case 8:
		target = &pbgear.IdentityPallet_SetFieldsCall{
			SetFieldsCall: To_Identity_SetFieldsCall(fields),
		}
	case 9:
		target = &pbgear.IdentityPallet_ProvideJudgementCall{
			ProvideJudgementCall: To_Identity_ProvideJudgementCall(fields),
		}
	case 10:
		target = &pbgear.IdentityPallet_KillIdentityCall{
			KillIdentityCall: To_Identity_KillIdentityCall(fields),
		}
	case 11:
		target = &pbgear.IdentityPallet_AddSubCall{
			AddSubCall: To_Identity_AddSubCall(fields),
		}
	case 12:
		target = &pbgear.IdentityPallet_RenameSubCall{
			RenameSubCall: To_Identity_RenameSubCall(fields),
		}
	case 13:
		target = &pbgear.IdentityPallet_RemoveSubCall{
			RemoveSubCall: To_Identity_RemoveSubCall(fields),
		}
	case 14:
		target = &pbgear.IdentityPallet_QuitSubCall{
			QuitSubCall: To_Identity_QuitSubCall(fields),
		}
	}
}

func To_Identity_Account(in any) *pbgear.Identity_Account {
	out := &pbgear.Identity_Account{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Account_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Account_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Account_Id{
			Id: To_Identity_Id(fields),
		}
	case 1:
		target = &pbgear.Identity_Account_Index{
			Index: To_Identity_Index(fields),
		}
	case 2:
		target = &pbgear.Identity_Account_Raw{
			Raw: To_Identity_Raw(fields),
		}
	case 3:
		target = &pbgear.Identity_Account_Address32{
			Address32: To_Identity_Address32(fields),
		}
	case 4:
		target = &pbgear.Identity_Account_Address20{
			Address20: To_Identity_Address20(fields),
		}
	}
}

func To_Identity_AddRegistrarCall(in any) *pbgear.Identity_AddRegistrarCall {
	out := &pbgear.Identity_AddRegistrarCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Account
	out.Account = To_Identity_AddRegistrarCall_account(fields[0].Value)

	return out //from message
}

func To_Identity_AddRegistrarCall_account(in any) *pbgear.Identity_Account {
	return nil //simple field
}

func To_Identity_AddSubCall(in any) *pbgear.Identity_AddSubCall {
	out := &pbgear.Identity_AddSubCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Sub
	out.Sub = To_Identity_AddSubCall_sub(fields[0].Value)
	// field Data
	out.Data = To_Identity_AddSubCall_data(fields[1].Value)

	return out //from message
}

func To_Identity_AddSubCall_sub(in any) *pbgear.Identity_Sub {
	return nil //simple field
}
func To_Identity_AddSubCall_data(in any) *pbgear.Identity_Data {
	return nil //simple field
}

func To_Identity_Address20(in any) *pbgear.Identity_Address20 {
	out := &pbgear.Identity_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Address32(in any) *pbgear.Identity_Address32 {
	out := &pbgear.Identity_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_BlakeTwo256(in any) *pbgear.Identity_BlakeTwo256 {
	out := &pbgear.Identity_BlakeTwo256{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_CancelRequestCall(in any) *pbgear.Identity_CancelRequestCall {
	out := &pbgear.Identity_CancelRequestCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field RegIndex
	out.RegIndex = To_uint32(fields[0].Value)

	return out //from message
}

func To_Identity_ClearIdentityCall(in any) *pbgear.Identity_ClearIdentityCall {
	out := &pbgear.Identity_ClearIdentityCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_Data(in any) *pbgear.Identity_Data {
	out := &pbgear.Identity_Data{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Data_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Data_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Data_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Data_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Data_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Data_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Data_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Data_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Data_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Data_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Data_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Data_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Data_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Data_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Data_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Data_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Data_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Data_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Data_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Data_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Data_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Data_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Data_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Data_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Data_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Data_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Data_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Data_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Data_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Data_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Data_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Data_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Data_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Data_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Data_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Data_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Data_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Data_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Data_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Data_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Display(in any) *pbgear.Identity_Display {
	out := &pbgear.Identity_Display{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Display_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Display_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Display_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Display_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Display_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Display_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Display_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Display_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Display_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Display_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Display_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Display_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Display_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Display_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Display_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Display_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Display_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Display_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Display_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Display_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Display_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Display_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Display_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Display_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Display_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Display_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Display_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Display_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Display_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Display_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Display_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Display_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Display_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Display_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Display_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Display_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Display_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Display_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Display_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Display_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Email(in any) *pbgear.Identity_Email {
	out := &pbgear.Identity_Email{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Email_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Email_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Email_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Email_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Email_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Email_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Email_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Email_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Email_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Email_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Email_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Email_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Email_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Email_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Email_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Email_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Email_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Email_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Email_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Email_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Email_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Email_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Email_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Email_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Email_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Email_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Email_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Email_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Email_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Email_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Email_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Email_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Email_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Email_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Email_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Email_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Email_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Email_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Email_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Email_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Erroneous(in any) *pbgear.Identity_Erroneous {
	out := &pbgear.Identity_Erroneous{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_FeePaid(in any) *pbgear.Identity_FeePaid {
	out := &pbgear.Identity_FeePaid{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_string(fields[0].Value)

	return out //from message
}

func To_Identity_Id(in any) *pbgear.Identity_Id {
	out := &pbgear.Identity_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Identity_Id_value0(fields[0].Value)

	return out //from message
}

func To_Identity_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Identity_Image(in any) *pbgear.Identity_Image {
	out := &pbgear.Identity_Image{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Image_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Image_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Image_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Image_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Image_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Image_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Image_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Image_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Image_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Image_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Image_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Image_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Image_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Image_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Image_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Image_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Image_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Image_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Image_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Image_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Image_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Image_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Image_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Image_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Image_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Image_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Image_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Image_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Image_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Image_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Image_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Image_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Image_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Image_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Image_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Image_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Image_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Image_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Image_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Image_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Index(in any) *pbgear.Identity_Index {
	out := &pbgear.Identity_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Identity_Index_value0(fields[0].Value)

	return out //from message
}

func To_Identity_Index_value0(in any) *pbgear.Identity_TupleNull {
	return nil //simple field
}

func To_Identity_Judgement(in any) *pbgear.Identity_Judgement {
	out := &pbgear.Identity_Judgement{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Judgement_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Judgement_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Judgement_Unknown{
			Unknown: To_Identity_Unknown(fields),
		}
	case 1:
		target = &pbgear.Identity_Judgement_FeePaid{
			FeePaid: To_Identity_FeePaid(fields),
		}
	case 2:
		target = &pbgear.Identity_Judgement_Reasonable{
			Reasonable: To_Identity_Reasonable(fields),
		}
	case 3:
		target = &pbgear.Identity_Judgement_KnownGood{
			KnownGood: To_Identity_KnownGood(fields),
		}
	case 4:
		target = &pbgear.Identity_Judgement_OutOfDate{
			OutOfDate: To_Identity_OutOfDate(fields),
		}
	case 5:
		target = &pbgear.Identity_Judgement_LowQuality{
			LowQuality: To_Identity_LowQuality(fields),
		}
	case 6:
		target = &pbgear.Identity_Judgement_Erroneous{
			Erroneous: To_Identity_Erroneous(fields),
		}
	}
}

func To_Identity_Keccak256(in any) *pbgear.Identity_Keccak256 {
	out := &pbgear.Identity_Keccak256{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_KillIdentityCall(in any) *pbgear.Identity_KillIdentityCall {
	out := &pbgear.Identity_KillIdentityCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Target
	out.Target = To_Identity_KillIdentityCall_target(fields[0].Value)

	return out //from message
}

func To_Identity_KillIdentityCall_target(in any) *pbgear.Identity_Target {
	return nil //simple field
}

func To_Identity_KnownGood(in any) *pbgear.Identity_KnownGood {
	out := &pbgear.Identity_KnownGood{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_Legal(in any) *pbgear.Identity_Legal {
	out := &pbgear.Identity_Legal{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Legal_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Legal_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Legal_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Legal_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Legal_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Legal_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Legal_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Legal_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Legal_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Legal_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Legal_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Legal_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Legal_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Legal_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Legal_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Legal_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Legal_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Legal_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Legal_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Legal_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Legal_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Legal_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Legal_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Legal_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Legal_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Legal_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Legal_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Legal_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Legal_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Legal_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Legal_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Legal_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Legal_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Legal_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Legal_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Legal_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Legal_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Legal_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Legal_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Legal_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_LowQuality(in any) *pbgear.Identity_LowQuality {
	out := &pbgear.Identity_LowQuality{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_New(in any) *pbgear.Identity_New {
	out := &pbgear.Identity_New{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_New_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_New_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_New_Id{
			Id: To_Identity_Id(fields),
		}
	case 1:
		target = &pbgear.Identity_New_Index{
			Index: To_Identity_Index(fields),
		}
	case 2:
		target = &pbgear.Identity_New_Raw{
			Raw: To_Identity_Raw(fields),
		}
	case 3:
		target = &pbgear.Identity_New_Address32{
			Address32: To_Identity_Address32(fields),
		}
	case 4:
		target = &pbgear.Identity_New_Address20{
			Address20: To_Identity_Address20(fields),
		}
	}
}

func To_Identity_None(in any) *pbgear.Identity_None {
	out := &pbgear.Identity_None{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_OutOfDate(in any) *pbgear.Identity_OutOfDate {
	out := &pbgear.Identity_OutOfDate{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_PalletIdentitySimpleIdentityInfo(in any) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Additional
	out.Additional = To_Identity_PalletIdentitySimpleIdentityInfo_additional(fields[0].Value)
	// field Display
	out.Display = To_Identity_PalletIdentitySimpleIdentityInfo_display(fields[1].Value)
	// field Legal
	out.Legal = To_Identity_PalletIdentitySimpleIdentityInfo_legal(fields[2].Value)
	// field Web
	out.Web = To_Identity_PalletIdentitySimpleIdentityInfo_web(fields[3].Value)
	// field Riot
	out.Riot = To_Identity_PalletIdentitySimpleIdentityInfo_riot(fields[4].Value)
	// field Email
	out.Email = To_Identity_PalletIdentitySimpleIdentityInfo_email(fields[5].Value)
	// primitive field PgpFingerprint
	out.PgpFingerprint = To_bytes(fields[6].Value)
	// field Image
	out.Image = To_Identity_PalletIdentitySimpleIdentityInfo_image(fields[7].Value)
	// field Twitter
	out.Twitter = To_Identity_PalletIdentitySimpleIdentityInfo_twitter(fields[8].Value)

	return out //from message
}

func To_Identity_PalletIdentitySimpleIdentityInfo_additional(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_display(in any) *pbgear.Identity_Display {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_legal(in any) *pbgear.Identity_Legal {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_web(in any) *pbgear.Identity_Web {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_riot(in any) *pbgear.Identity_Riot {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_email(in any) *pbgear.Identity_Email {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_image(in any) *pbgear.Identity_Image {
	return nil //simple field
}
func To_Identity_PalletIdentitySimpleIdentityInfo_twitter(in any) *pbgear.Identity_Twitter {
	return nil //simple field
}

func To_Identity_PalletIdentityTypesBitFlags(in any) *pbgear.Identity_PalletIdentityTypesBitFlags {
	out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint64(fields[0].Value)

	return out //from message
}

func To_Identity_ProvideJudgementCall(in any) *pbgear.Identity_ProvideJudgementCall {
	out := &pbgear.Identity_ProvideJudgementCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field RegIndex
	out.RegIndex = To_uint32(fields[0].Value)
	// field Target
	out.Target = To_Identity_ProvideJudgementCall_target(fields[1].Value)
	// field Judgement
	out.Judgement = To_Identity_ProvideJudgementCall_judgement(fields[2].Value)
	// field Identity
	out.Identity = To_Identity_ProvideJudgementCall_identity(fields[3].Value)

	return out //from message
}

func To_Identity_ProvideJudgementCall_target(in any) *pbgear.Identity_Target {
	return nil //simple field
}
func To_Identity_ProvideJudgementCall_judgement(in any) *pbgear.Identity_Judgement {
	return nil //simple field
}
func To_Identity_ProvideJudgementCall_identity(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Identity_QuitSubCall(in any) *pbgear.Identity_QuitSubCall {
	out := &pbgear.Identity_QuitSubCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_Raw(in any) *pbgear.Identity_Raw {
	out := &pbgear.Identity_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw0(in any) *pbgear.Identity_Raw0 {
	out := &pbgear.Identity_Raw0{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw1(in any) *pbgear.Identity_Raw1 {
	out := &pbgear.Identity_Raw1{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw10(in any) *pbgear.Identity_Raw10 {
	out := &pbgear.Identity_Raw10{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw11(in any) *pbgear.Identity_Raw11 {
	out := &pbgear.Identity_Raw11{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw12(in any) *pbgear.Identity_Raw12 {
	out := &pbgear.Identity_Raw12{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw13(in any) *pbgear.Identity_Raw13 {
	out := &pbgear.Identity_Raw13{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw14(in any) *pbgear.Identity_Raw14 {
	out := &pbgear.Identity_Raw14{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw15(in any) *pbgear.Identity_Raw15 {
	out := &pbgear.Identity_Raw15{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw16(in any) *pbgear.Identity_Raw16 {
	out := &pbgear.Identity_Raw16{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw17(in any) *pbgear.Identity_Raw17 {
	out := &pbgear.Identity_Raw17{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw18(in any) *pbgear.Identity_Raw18 {
	out := &pbgear.Identity_Raw18{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw19(in any) *pbgear.Identity_Raw19 {
	out := &pbgear.Identity_Raw19{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw2(in any) *pbgear.Identity_Raw2 {
	out := &pbgear.Identity_Raw2{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw20(in any) *pbgear.Identity_Raw20 {
	out := &pbgear.Identity_Raw20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw21(in any) *pbgear.Identity_Raw21 {
	out := &pbgear.Identity_Raw21{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw22(in any) *pbgear.Identity_Raw22 {
	out := &pbgear.Identity_Raw22{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw23(in any) *pbgear.Identity_Raw23 {
	out := &pbgear.Identity_Raw23{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw24(in any) *pbgear.Identity_Raw24 {
	out := &pbgear.Identity_Raw24{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw25(in any) *pbgear.Identity_Raw25 {
	out := &pbgear.Identity_Raw25{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw26(in any) *pbgear.Identity_Raw26 {
	out := &pbgear.Identity_Raw26{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw27(in any) *pbgear.Identity_Raw27 {
	out := &pbgear.Identity_Raw27{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw28(in any) *pbgear.Identity_Raw28 {
	out := &pbgear.Identity_Raw28{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw29(in any) *pbgear.Identity_Raw29 {
	out := &pbgear.Identity_Raw29{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw3(in any) *pbgear.Identity_Raw3 {
	out := &pbgear.Identity_Raw3{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw30(in any) *pbgear.Identity_Raw30 {
	out := &pbgear.Identity_Raw30{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw31(in any) *pbgear.Identity_Raw31 {
	out := &pbgear.Identity_Raw31{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw32(in any) *pbgear.Identity_Raw32 {
	out := &pbgear.Identity_Raw32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw4(in any) *pbgear.Identity_Raw4 {
	out := &pbgear.Identity_Raw4{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw5(in any) *pbgear.Identity_Raw5 {
	out := &pbgear.Identity_Raw5{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw6(in any) *pbgear.Identity_Raw6 {
	out := &pbgear.Identity_Raw6{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw7(in any) *pbgear.Identity_Raw7 {
	out := &pbgear.Identity_Raw7{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw8(in any) *pbgear.Identity_Raw8 {
	out := &pbgear.Identity_Raw8{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Raw9(in any) *pbgear.Identity_Raw9 {
	out := &pbgear.Identity_Raw9{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Reasonable(in any) *pbgear.Identity_Reasonable {
	out := &pbgear.Identity_Reasonable{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_RemoveSubCall(in any) *pbgear.Identity_RemoveSubCall {
	out := &pbgear.Identity_RemoveSubCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Sub
	out.Sub = To_Identity_RemoveSubCall_sub(fields[0].Value)

	return out //from message
}

func To_Identity_RemoveSubCall_sub(in any) *pbgear.Identity_Sub {
	return nil //simple field
}

func To_Identity_RenameSubCall(in any) *pbgear.Identity_RenameSubCall {
	out := &pbgear.Identity_RenameSubCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Sub
	out.Sub = To_Identity_RenameSubCall_sub(fields[0].Value)
	// field Data
	out.Data = To_Identity_RenameSubCall_data(fields[1].Value)

	return out //from message
}

func To_Identity_RenameSubCall_sub(in any) *pbgear.Identity_Sub {
	return nil //simple field
}
func To_Identity_RenameSubCall_data(in any) *pbgear.Identity_Data {
	return nil //simple field
}

func To_Identity_RequestJudgementCall(in any) *pbgear.Identity_RequestJudgementCall {
	out := &pbgear.Identity_RequestJudgementCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field RegIndex
	out.RegIndex = To_uint32(fields[0].Value)
	// primitive field MaxFee
	out.MaxFee = To_string(fields[1].Value)

	return out //from message
}

func To_Identity_Riot(in any) *pbgear.Identity_Riot {
	out := &pbgear.Identity_Riot{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Riot_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Riot_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Riot_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Riot_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Riot_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Riot_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Riot_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Riot_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Riot_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Riot_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Riot_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Riot_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Riot_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Riot_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Riot_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Riot_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Riot_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Riot_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Riot_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Riot_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Riot_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Riot_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Riot_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Riot_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Riot_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Riot_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Riot_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Riot_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Riot_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Riot_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Riot_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Riot_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Riot_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Riot_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Riot_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Riot_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Riot_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Riot_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Riot_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Riot_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_SetAccountIdCall(in any) *pbgear.Identity_SetAccountIdCall {
	out := &pbgear.Identity_SetAccountIdCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// field New
	out.New = To_Identity_SetAccountIdCall_new(fields[1].Value)

	return out //from message
}

func To_Identity_SetAccountIdCall_new(in any) *pbgear.Identity_New {
	return nil //simple field
}

func To_Identity_SetFeeCall(in any) *pbgear.Identity_SetFeeCall {
	out := &pbgear.Identity_SetFeeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// primitive field Fee
	out.Fee = To_string(fields[1].Value)

	return out //from message
}

func To_Identity_SetFieldsCall(in any) *pbgear.Identity_SetFieldsCall {
	out := &pbgear.Identity_SetFieldsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// field Fields
	out.Fields = To_Identity_SetFieldsCall_fields(fields[1].Value)

	return out //from message
}

func To_Identity_SetFieldsCall_fields(in any) *pbgear.Identity_PalletIdentityTypesBitFlags {
	return nil //simple field
}

func To_Identity_SetIdentityCall(in any) *pbgear.Identity_SetIdentityCall {
	out := &pbgear.Identity_SetIdentityCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Info
	out.Info = To_Identity_SetIdentityCall_info(fields[0].Value)

	return out //from message
}

func To_Identity_SetIdentityCall_info(in any) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	return nil //simple field
}

func To_Identity_SetSubsCall(in any) *pbgear.Identity_SetSubsCall {
	out := &pbgear.Identity_SetSubsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Subs
	out.Subs = To_Repeated_Identity_SetSubsCall_subs(fields[0].Value)

	return out //from message
}

func To_Repeated_Identity_SetSubsCall_subs(in any) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	items := in.([]interface{})

	var out []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData
	for _, item := range items {
		o := To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Identity_Sha256(in any) *pbgear.Identity_Sha256 {
	out := &pbgear.Identity_Sha256{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_ShaThree256(in any) *pbgear.Identity_ShaThree256 {
	out := &pbgear.Identity_ShaThree256{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Identity_Sub(in any) *pbgear.Identity_Sub {
	out := &pbgear.Identity_Sub{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Sub_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Sub_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Sub_Id{
			Id: To_Identity_Id(fields),
		}
	case 1:
		target = &pbgear.Identity_Sub_Index{
			Index: To_Identity_Index(fields),
		}
	case 2:
		target = &pbgear.Identity_Sub_Raw{
			Raw: To_Identity_Raw(fields),
		}
	case 3:
		target = &pbgear.Identity_Sub_Address32{
			Address32: To_Identity_Address32(fields),
		}
	case 4:
		target = &pbgear.Identity_Sub_Address20{
			Address20: To_Identity_Address20(fields),
		}
	}
}

func To_Identity_Target(in any) *pbgear.Identity_Target {
	out := &pbgear.Identity_Target{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Target_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Target_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Target_Id{
			Id: To_Identity_Id(fields),
		}
	case 1:
		target = &pbgear.Identity_Target_Index{
			Index: To_Identity_Index(fields),
		}
	case 2:
		target = &pbgear.Identity_Target_Raw{
			Raw: To_Identity_Raw(fields),
		}
	case 3:
		target = &pbgear.Identity_Target_Address32{
			Address32: To_Identity_Address32(fields),
		}
	case 4:
		target = &pbgear.Identity_Target_Address20{
			Address20: To_Identity_Address20(fields),
		}
	}
}

func To_Identity_TupleNull(in any) *pbgear.Identity_TupleNull {
	out := &pbgear.Identity_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Identity_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Identity_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(in any) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1(fields[1].Value)

	return out //from message
}

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0(in any) *pbgear.Identity_Value0 {
	return nil //simple field
}
func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1(in any) *pbgear.Identity_Value1 {
	return nil //simple field
}

func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(in any) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1(fields[1].Value)

	return out //from message
}

func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}
func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1(in any) *pbgear.Identity_Value1 {
	return nil //simple field
}

func To_Identity_Twitter(in any) *pbgear.Identity_Twitter {
	out := &pbgear.Identity_Twitter{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Twitter_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Twitter_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Twitter_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Twitter_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Twitter_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Twitter_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Twitter_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Twitter_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Twitter_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Twitter_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Twitter_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Twitter_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Twitter_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Twitter_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Twitter_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Twitter_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Twitter_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Twitter_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Twitter_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Twitter_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Twitter_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Twitter_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Twitter_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Twitter_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Twitter_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Twitter_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Twitter_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Twitter_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Twitter_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Twitter_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Twitter_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Twitter_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Twitter_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Twitter_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Twitter_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Twitter_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Twitter_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Twitter_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Twitter_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Twitter_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Unknown(in any) *pbgear.Identity_Unknown {
	out := &pbgear.Identity_Unknown{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Identity_Value0(in any) *pbgear.Identity_Value0 {
	out := &pbgear.Identity_Value0{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Value0_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Value0_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Value0_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Value0_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Value0_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Value0_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Value0_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Value0_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Value0_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Value0_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Value0_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Value0_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Value0_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Value0_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Value0_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Value0_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Value0_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Value0_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Value0_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Value0_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Value0_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Value0_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Value0_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Value0_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Value0_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Value0_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Value0_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Value0_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Value0_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Value0_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Value0_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Value0_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Value0_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Value0_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Value0_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Value0_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Value0_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Value0_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Value0_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Value0_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Value1(in any) *pbgear.Identity_Value1 {
	out := &pbgear.Identity_Value1{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Value1_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Value1_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Value1_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Value1_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Value1_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Value1_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Value1_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Value1_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Value1_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Value1_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Value1_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Value1_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Value1_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Value1_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Value1_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Value1_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Value1_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Value1_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Value1_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Value1_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Value1_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Value1_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Value1_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Value1_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Value1_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Value1_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Value1_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Value1_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Value1_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Value1_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Value1_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Value1_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Value1_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Value1_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Value1_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Value1_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Value1_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Value1_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Value1_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Value1_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_Identity_Web(in any) *pbgear.Identity_Web {
	out := &pbgear.Identity_Web{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Identity_Web_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Identity_Web_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Identity_Web_None{
			None: To_Identity_None(fields),
		}
	case 1:
		target = &pbgear.Identity_Web_Raw0{
			Raw0: To_Identity_Raw0(fields),
		}
	case 2:
		target = &pbgear.Identity_Web_Raw1{
			Raw1: To_Identity_Raw1(fields),
		}
	case 3:
		target = &pbgear.Identity_Web_Raw2{
			Raw2: To_Identity_Raw2(fields),
		}
	case 4:
		target = &pbgear.Identity_Web_Raw3{
			Raw3: To_Identity_Raw3(fields),
		}
	case 5:
		target = &pbgear.Identity_Web_Raw4{
			Raw4: To_Identity_Raw4(fields),
		}
	case 6:
		target = &pbgear.Identity_Web_Raw5{
			Raw5: To_Identity_Raw5(fields),
		}
	case 7:
		target = &pbgear.Identity_Web_Raw6{
			Raw6: To_Identity_Raw6(fields),
		}
	case 8:
		target = &pbgear.Identity_Web_Raw7{
			Raw7: To_Identity_Raw7(fields),
		}
	case 9:
		target = &pbgear.Identity_Web_Raw8{
			Raw8: To_Identity_Raw8(fields),
		}
	case 10:
		target = &pbgear.Identity_Web_Raw9{
			Raw9: To_Identity_Raw9(fields),
		}
	case 11:
		target = &pbgear.Identity_Web_Raw10{
			Raw10: To_Identity_Raw10(fields),
		}
	case 12:
		target = &pbgear.Identity_Web_Raw11{
			Raw11: To_Identity_Raw11(fields),
		}
	case 13:
		target = &pbgear.Identity_Web_Raw12{
			Raw12: To_Identity_Raw12(fields),
		}
	case 14:
		target = &pbgear.Identity_Web_Raw13{
			Raw13: To_Identity_Raw13(fields),
		}
	case 15:
		target = &pbgear.Identity_Web_Raw14{
			Raw14: To_Identity_Raw14(fields),
		}
	case 16:
		target = &pbgear.Identity_Web_Raw15{
			Raw15: To_Identity_Raw15(fields),
		}
	case 17:
		target = &pbgear.Identity_Web_Raw16{
			Raw16: To_Identity_Raw16(fields),
		}
	case 18:
		target = &pbgear.Identity_Web_Raw17{
			Raw17: To_Identity_Raw17(fields),
		}
	case 19:
		target = &pbgear.Identity_Web_Raw18{
			Raw18: To_Identity_Raw18(fields),
		}
	case 20:
		target = &pbgear.Identity_Web_Raw19{
			Raw19: To_Identity_Raw19(fields),
		}
	case 21:
		target = &pbgear.Identity_Web_Raw20{
			Raw20: To_Identity_Raw20(fields),
		}
	case 22:
		target = &pbgear.Identity_Web_Raw21{
			Raw21: To_Identity_Raw21(fields),
		}
	case 23:
		target = &pbgear.Identity_Web_Raw22{
			Raw22: To_Identity_Raw22(fields),
		}
	case 24:
		target = &pbgear.Identity_Web_Raw23{
			Raw23: To_Identity_Raw23(fields),
		}
	case 25:
		target = &pbgear.Identity_Web_Raw24{
			Raw24: To_Identity_Raw24(fields),
		}
	case 26:
		target = &pbgear.Identity_Web_Raw25{
			Raw25: To_Identity_Raw25(fields),
		}
	case 27:
		target = &pbgear.Identity_Web_Raw26{
			Raw26: To_Identity_Raw26(fields),
		}
	case 28:
		target = &pbgear.Identity_Web_Raw27{
			Raw27: To_Identity_Raw27(fields),
		}
	case 29:
		target = &pbgear.Identity_Web_Raw28{
			Raw28: To_Identity_Raw28(fields),
		}
	case 30:
		target = &pbgear.Identity_Web_Raw29{
			Raw29: To_Identity_Raw29(fields),
		}
	case 31:
		target = &pbgear.Identity_Web_Raw30{
			Raw30: To_Identity_Raw30(fields),
		}
	case 32:
		target = &pbgear.Identity_Web_Raw31{
			Raw31: To_Identity_Raw31(fields),
		}
	case 33:
		target = &pbgear.Identity_Web_Raw32{
			Raw32: To_Identity_Raw32(fields),
		}
	case 34:
		target = &pbgear.Identity_Web_BlakeTwo256{
			BlakeTwo256: To_Identity_BlakeTwo256(fields),
		}
	case 35:
		target = &pbgear.Identity_Web_Sha256{
			Sha256: To_Identity_Sha256(fields),
		}
	case 36:
		target = &pbgear.Identity_Web_Keccak256{
			Keccak256: To_Identity_Keccak256(fields),
		}
	case 37:
		target = &pbgear.Identity_Web_ShaThree256{
			ShaThree256: To_Identity_ShaThree256(fields),
		}
	}
}

func To_ImOnlinePallet(in any) *pbgear.ImOnlinePallet {
	out := &pbgear.ImOnlinePallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ImOnlinePallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ImOnlinePallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ImOnlinePallet_HeartbeatCall{
			HeartbeatCall: To_ImOnline_HeartbeatCall(fields),
		}
	}
}

func To_ImOnline_HeartbeatCall(in any) *pbgear.ImOnline_HeartbeatCall {
	out := &pbgear.ImOnline_HeartbeatCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Heartbeat
	out.Heartbeat = To_ImOnline_HeartbeatCall_heartbeat(fields[0].Value)
	// field Signature
	out.Signature = To_ImOnline_HeartbeatCall_signature(fields[1].Value)

	return out //from message
}

func To_ImOnline_HeartbeatCall_heartbeat(in any) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	return nil //simple field
}
func To_ImOnline_HeartbeatCall_signature(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	return nil //simple field
}

func To_ImOnline_PalletImOnlineHeartbeat(in any) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field BlockNumber
	out.BlockNumber = To_uint32(fields[0].Value)
	// primitive field SessionIndex
	out.SessionIndex = To_uint32(fields[1].Value)
	// primitive field AuthorityIndex
	out.AuthorityIndex = To_uint32(fields[2].Value)
	// primitive field ValidatorsLen
	out.ValidatorsLen = To_uint32(fields[3].Value)

	return out //from message
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Public(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public {
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ImOnline_PalletImOnlineSr25519AppSr25519Public_value0(fields[0].Value)

	return out //from message
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Public_value0(in any) *pbgear.SpCoreSr25519Public {
	return nil //simple field
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_ImOnline_PalletImOnlineSr25519AppSr25519Signature_value0(fields[0].Value)

	return out //from message
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature_value0(in any) *pbgear.SpCoreSr25519Signature {
	return nil //simple field
}

func To_MultisigPallet(in any) *pbgear.MultisigPallet {
	out := &pbgear.MultisigPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_MultisigPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_MultisigPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.MultisigPallet_AsMultiThreshold1Call{
			AsMultiThreshold1Call: To_Multisig_AsMultiThreshold1Call(fields),
		}
	case 1:
		target = &pbgear.MultisigPallet_AsMultiCall{
			AsMultiCall: To_Multisig_AsMultiCall(fields),
		}
	case 2:
		target = &pbgear.MultisigPallet_ApproveAsMultiCall{
			ApproveAsMultiCall: To_Multisig_ApproveAsMultiCall(fields),
		}
	case 3:
		target = &pbgear.MultisigPallet_CancelAsMultiCall{
			CancelAsMultiCall: To_Multisig_CancelAsMultiCall(fields),
		}
	}
}

func To_Multisig_ApproveAsMultiCall(in any) *pbgear.Multisig_ApproveAsMultiCall {
	out := &pbgear.Multisig_ApproveAsMultiCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Threshold
	out.Threshold = To_uint32(fields[0].Value)
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(fields[1].Value)
	// optional field MaybeTimepoint
	out.MaybeTimepoint = To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(fields[2].Value)
	// primitive field CallHash
	out.CallHash = To_bytes(fields[3].Value)
	// field MaxWeight
	out.MaxWeight = To_Multisig_ApproveAsMultiCall_max_weight(fields[4].Value)

	return out //from message
}

func To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	return nil //funcForOptionalField
}
func To_Multisig_ApproveAsMultiCall_max_weight(in any) *pbgear.SpWeightsWeightV2Weight {
	return nil //simple field
}

func To_Multisig_AsMultiCall(in any) *pbgear.Multisig_AsMultiCall {
	out := &pbgear.Multisig_AsMultiCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Threshold
	out.Threshold = To_uint32(fields[0].Value)
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_AsMultiCall_other_signatories(fields[1].Value)
	// optional field MaybeTimepoint
	out.MaybeTimepoint = To_Optional_Multisig_AsMultiCall_maybe_timepoint(fields[2].Value)
	// oneOf field Call
	Set_OneOf_Multisig_AsMultiCall_call(out.Call, fields[3].Value)
	// field MaxWeight
	out.MaxWeight = To_Multisig_AsMultiCall_max_weight(fields[4].Value)

	return out //from message
}

func To_Repeated_Multisig_AsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Optional_Multisig_AsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	return nil //funcForOptionalField
}
func Set_OneOf_Multisig_AsMultiCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Multisig_AsMultiCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Multisig_AsMultiCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Multisig_AsMultiCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Multisig_AsMultiCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Multisig_AsMultiCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Multisig_AsMultiCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Multisig_AsMultiCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Multisig_AsMultiCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Multisig_AsMultiCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Multisig_AsMultiCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Multisig_AsMultiCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Multisig_AsMultiCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Multisig_AsMultiCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Multisig_AsMultiCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Multisig_AsMultiCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Multisig_AsMultiCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Multisig_AsMultiCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Multisig_AsMultiCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Multisig_AsMultiCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Multisig_AsMultiCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Multisig_AsMultiCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Multisig_AsMultiCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Multisig_AsMultiCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Multisig_AsMultiCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Multisig_AsMultiCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Multisig_AsMultiCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Multisig_AsMultiCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Multisig_AsMultiCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Multisig_AsMultiCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}
func To_Multisig_AsMultiCall_max_weight(in any) *pbgear.SpWeightsWeightV2Weight {
	return nil //simple field
}

func To_Multisig_AsMultiThreshold1Call(in any) *pbgear.Multisig_AsMultiThreshold1Call {
	out := &pbgear.Multisig_AsMultiThreshold1Call{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(fields[0].Value)
	// oneOf field Call
	Set_OneOf_Multisig_AsMultiThreshold1Call_call(out.Call, fields[1].Value)

	return out //from message
}

func To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func Set_OneOf_Multisig_AsMultiThreshold1Call_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Multisig_AsMultiThreshold1Call_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Multisig_AsMultiThreshold1Call_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Multisig_AsMultiThreshold1Call_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Multisig_AsMultiThreshold1Call_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Multisig_AsMultiThreshold1Call_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Multisig_AsMultiThreshold1Call_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Multisig_AsMultiThreshold1Call_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Multisig_AsMultiThreshold1Call_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Multisig_AsMultiThreshold1Call_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Multisig_AsMultiThreshold1Call_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Multisig_AsMultiThreshold1Call_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Multisig_CancelAsMultiCall(in any) *pbgear.Multisig_CancelAsMultiCall {
	out := &pbgear.Multisig_CancelAsMultiCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Threshold
	out.Threshold = To_uint32(fields[0].Value)
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_CancelAsMultiCall_other_signatories(fields[1].Value)
	// field Timepoint
	out.Timepoint = To_Multisig_CancelAsMultiCall_timepoint(fields[2].Value)
	// primitive field CallHash
	out.CallHash = To_bytes(fields[3].Value)

	return out //from message
}

func To_Repeated_Multisig_CancelAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Multisig_CancelAsMultiCall_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	return nil //simple field
}

func To_Multisig_PalletMultisigTimepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	out := &pbgear.Multisig_PalletMultisigTimepoint{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Height
	out.Height = To_uint32(fields[0].Value)
	// primitive field Index
	out.Index = To_uint32(fields[1].Value)

	return out //from message
}

func To_NominationPoolsPallet(in any) *pbgear.NominationPoolsPallet {
	out := &pbgear.NominationPoolsPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_NominationPoolsPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPoolsPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPoolsPallet_JoinCall{
			JoinCall: To_NominationPools_JoinCall(fields),
		}
	case 1:
		target = &pbgear.NominationPoolsPallet_BondExtraCall{
			BondExtraCall: To_NominationPools_BondExtraCall(fields),
		}
	case 2:
		target = &pbgear.NominationPoolsPallet_ClaimPayoutCall{
			ClaimPayoutCall: To_NominationPools_ClaimPayoutCall(fields),
		}
	case 3:
		target = &pbgear.NominationPoolsPallet_UnbondCall{
			UnbondCall: To_NominationPools_UnbondCall(fields),
		}
	case 4:
		target = &pbgear.NominationPoolsPallet_PoolWithdrawUnbondedCall{
			PoolWithdrawUnbondedCall: To_NominationPools_PoolWithdrawUnbondedCall(fields),
		}
	case 5:
		target = &pbgear.NominationPoolsPallet_WithdrawUnbondedCall{
			WithdrawUnbondedCall: To_NominationPools_WithdrawUnbondedCall(fields),
		}
	case 6:
		target = &pbgear.NominationPoolsPallet_CreateCall{
			CreateCall: To_NominationPools_CreateCall(fields),
		}
	case 7:
		target = &pbgear.NominationPoolsPallet_CreateWithPoolIdCall{
			CreateWithPoolIdCall: To_NominationPools_CreateWithPoolIdCall(fields),
		}
	case 8:
		target = &pbgear.NominationPoolsPallet_NominateCall{
			NominateCall: To_NominationPools_NominateCall(fields),
		}
	case 9:
		target = &pbgear.NominationPoolsPallet_SetStateCall{
			SetStateCall: To_NominationPools_SetStateCall(fields),
		}
	case 10:
		target = &pbgear.NominationPoolsPallet_SetMetadataCall{
			SetMetadataCall: To_NominationPools_SetMetadataCall(fields),
		}
	case 11:
		target = &pbgear.NominationPoolsPallet_SetConfigsCall{
			SetConfigsCall: To_NominationPools_SetConfigsCall(fields),
		}
	case 12:
		target = &pbgear.NominationPoolsPallet_UpdateRolesCall{
			UpdateRolesCall: To_NominationPools_UpdateRolesCall(fields),
		}
	case 13:
		target = &pbgear.NominationPoolsPallet_ChillCall{
			ChillCall: To_NominationPools_ChillCall(fields),
		}
	case 14:
		target = &pbgear.NominationPoolsPallet_BondExtraOtherCall{
			BondExtraOtherCall: To_NominationPools_BondExtraOtherCall(fields),
		}
	case 15:
		target = &pbgear.NominationPoolsPallet_SetClaimPermissionCall{
			SetClaimPermissionCall: To_NominationPools_SetClaimPermissionCall(fields),
		}
	case 16:
		target = &pbgear.NominationPoolsPallet_ClaimPayoutOtherCall{
			ClaimPayoutOtherCall: To_NominationPools_ClaimPayoutOtherCall(fields),
		}
	case 17:
		target = &pbgear.NominationPoolsPallet_SetCommissionCall{
			SetCommissionCall: To_NominationPools_SetCommissionCall(fields),
		}
	case 18:
		target = &pbgear.NominationPoolsPallet_SetCommissionMaxCall{
			SetCommissionMaxCall: To_NominationPools_SetCommissionMaxCall(fields),
		}
	case 19:
		target = &pbgear.NominationPoolsPallet_SetCommissionChangeRateCall{
			SetCommissionChangeRateCall: To_NominationPools_SetCommissionChangeRateCall(fields),
		}
	case 20:
		target = &pbgear.NominationPoolsPallet_ClaimCommissionCall{
			ClaimCommissionCall: To_NominationPools_ClaimCommissionCall(fields),
		}
	case 21:
		target = &pbgear.NominationPoolsPallet_AdjustPoolDepositCall{
			AdjustPoolDepositCall: To_NominationPools_AdjustPoolDepositCall(fields),
		}
	}
}

func To_NominationPools_Address20(in any) *pbgear.NominationPools_Address20 {
	out := &pbgear.NominationPools_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_NominationPools_Address32(in any) *pbgear.NominationPools_Address32 {
	out := &pbgear.NominationPools_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_NominationPools_AdjustPoolDepositCall(in any) *pbgear.NominationPools_AdjustPoolDepositCall {
	out := &pbgear.NominationPools_AdjustPoolDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)

	return out //from message
}

func To_NominationPools_Blocked(in any) *pbgear.NominationPools_Blocked {
	out := &pbgear.NominationPools_Blocked{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_BondExtraCall(in any) *pbgear.NominationPools_BondExtraCall {
	out := &pbgear.NominationPools_BondExtraCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Extra
	out.Extra = To_NominationPools_BondExtraCall_extra(fields[0].Value)

	return out //from message
}

func To_NominationPools_BondExtraCall_extra(in any) *pbgear.NominationPools_Extra {
	return nil //simple field
}

func To_NominationPools_BondExtraOtherCall(in any) *pbgear.NominationPools_BondExtraOtherCall {
	out := &pbgear.NominationPools_BondExtraOtherCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Member
	out.Member = To_NominationPools_BondExtraOtherCall_member(fields[0].Value)
	// field Extra
	out.Extra = To_NominationPools_BondExtraOtherCall_extra(fields[1].Value)

	return out //from message
}

func To_NominationPools_BondExtraOtherCall_member(in any) *pbgear.NominationPools_Member {
	return nil //simple field
}
func To_NominationPools_BondExtraOtherCall_extra(in any) *pbgear.NominationPools_Extra {
	return nil //simple field
}

func To_NominationPools_Bouncer(in any) *pbgear.NominationPools_Bouncer {
	out := &pbgear.NominationPools_Bouncer{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Bouncer_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Bouncer_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Bouncer_Id{
			Id: To_NominationPools_Id(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Bouncer_Index{
			Index: To_NominationPools_Index(fields),
		}
	case 2:
		target = &pbgear.NominationPools_Bouncer_Raw{
			Raw: To_NominationPools_Raw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_Bouncer_Address32{
			Address32: To_NominationPools_Address32(fields),
		}
	case 4:
		target = &pbgear.NominationPools_Bouncer_Address20{
			Address20: To_NominationPools_Address20(fields),
		}
	}
}

func To_NominationPools_ChillCall(in any) *pbgear.NominationPools_ChillCall {
	out := &pbgear.NominationPools_ChillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)

	return out //from message
}

func To_NominationPools_ClaimCommissionCall(in any) *pbgear.NominationPools_ClaimCommissionCall {
	out := &pbgear.NominationPools_ClaimCommissionCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)

	return out //from message
}

func To_NominationPools_ClaimPayoutCall(in any) *pbgear.NominationPools_ClaimPayoutCall {
	out := &pbgear.NominationPools_ClaimPayoutCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_ClaimPayoutOtherCall(in any) *pbgear.NominationPools_ClaimPayoutOtherCall {
	out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Other
	out.Other = To_NominationPools_ClaimPayoutOtherCall_other(fields[0].Value)

	return out //from message
}

func To_NominationPools_ClaimPayoutOtherCall_other(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_NominationPools_CreateCall(in any) *pbgear.NominationPools_CreateCall {
	out := &pbgear.NominationPools_CreateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Amount
	out.Amount = To_string(fields[0].Value)
	// field Root
	out.Root = To_NominationPools_CreateCall_root(fields[1].Value)
	// field Nominator
	out.Nominator = To_NominationPools_CreateCall_nominator(fields[2].Value)
	// field Bouncer
	out.Bouncer = To_NominationPools_CreateCall_bouncer(fields[3].Value)

	return out //from message
}

func To_NominationPools_CreateCall_root(in any) *pbgear.NominationPools_Root {
	return nil //simple field
}
func To_NominationPools_CreateCall_nominator(in any) *pbgear.NominationPools_Nominator {
	return nil //simple field
}
func To_NominationPools_CreateCall_bouncer(in any) *pbgear.NominationPools_Bouncer {
	return nil //simple field
}

func To_NominationPools_CreateWithPoolIdCall(in any) *pbgear.NominationPools_CreateWithPoolIdCall {
	out := &pbgear.NominationPools_CreateWithPoolIdCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Amount
	out.Amount = To_string(fields[0].Value)
	// field Root
	out.Root = To_NominationPools_CreateWithPoolIdCall_root(fields[1].Value)
	// field Nominator
	out.Nominator = To_NominationPools_CreateWithPoolIdCall_nominator(fields[2].Value)
	// field Bouncer
	out.Bouncer = To_NominationPools_CreateWithPoolIdCall_bouncer(fields[3].Value)
	// primitive field PoolId
	out.PoolId = To_uint32(fields[4].Value)

	return out //from message
}

func To_NominationPools_CreateWithPoolIdCall_root(in any) *pbgear.NominationPools_Root {
	return nil //simple field
}
func To_NominationPools_CreateWithPoolIdCall_nominator(in any) *pbgear.NominationPools_Nominator {
	return nil //simple field
}
func To_NominationPools_CreateWithPoolIdCall_bouncer(in any) *pbgear.NominationPools_Bouncer {
	return nil //simple field
}

func To_NominationPools_Destroying(in any) *pbgear.NominationPools_Destroying {
	out := &pbgear.NominationPools_Destroying{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_Extra(in any) *pbgear.NominationPools_Extra {
	out := &pbgear.NominationPools_Extra{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Extra_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Extra_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Extra_FreeBalance{
			FreeBalance: To_NominationPools_FreeBalance(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Extra_Rewards{
			Rewards: To_NominationPools_Rewards(fields),
		}
	}
}

func To_NominationPools_FreeBalance(in any) *pbgear.NominationPools_FreeBalance {
	out := &pbgear.NominationPools_FreeBalance{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_string(fields[0].Value)

	return out //from message
}

func To_NominationPools_GlobalMaxCommission(in any) *pbgear.NominationPools_GlobalMaxCommission {
	out := &pbgear.NominationPools_GlobalMaxCommission{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_GlobalMaxCommission_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_GlobalMaxCommission_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_GlobalMaxCommission_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_GlobalMaxCommission_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_GlobalMaxCommission_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_Id(in any) *pbgear.NominationPools_Id {
	out := &pbgear.NominationPools_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_NominationPools_Id_value0(fields[0].Value)

	return out //from message
}

func To_NominationPools_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_NominationPools_Index(in any) *pbgear.NominationPools_Index {
	out := &pbgear.NominationPools_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_NominationPools_Index_value0(fields[0].Value)

	return out //from message
}

func To_NominationPools_Index_value0(in any) *pbgear.NominationPools_TupleNull {
	return nil //simple field
}

func To_NominationPools_JoinCall(in any) *pbgear.NominationPools_JoinCall {
	out := &pbgear.NominationPools_JoinCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Amount
	out.Amount = To_string(fields[0].Value)
	// primitive field PoolId
	out.PoolId = To_uint32(fields[1].Value)

	return out //from message
}

func To_NominationPools_MaxMembers(in any) *pbgear.NominationPools_MaxMembers {
	out := &pbgear.NominationPools_MaxMembers{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MaxMembers_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MaxMembers_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MaxMembers_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MaxMembers_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MaxMembers_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_MaxMembersPerPool(in any) *pbgear.NominationPools_MaxMembersPerPool {
	out := &pbgear.NominationPools_MaxMembersPerPool{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MaxMembersPerPool_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MaxMembersPerPool_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MaxMembersPerPool_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MaxMembersPerPool_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MaxMembersPerPool_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_MaxPools(in any) *pbgear.NominationPools_MaxPools {
	out := &pbgear.NominationPools_MaxPools{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MaxPools_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MaxPools_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MaxPools_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MaxPools_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MaxPools_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_Member(in any) *pbgear.NominationPools_Member {
	out := &pbgear.NominationPools_Member{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Member_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Member_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Member_Id{
			Id: To_NominationPools_Id(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Member_Index{
			Index: To_NominationPools_Index(fields),
		}
	case 2:
		target = &pbgear.NominationPools_Member_Raw{
			Raw: To_NominationPools_Raw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_Member_Address32{
			Address32: To_NominationPools_Address32(fields),
		}
	case 4:
		target = &pbgear.NominationPools_Member_Address20{
			Address20: To_NominationPools_Address20(fields),
		}
	}
}

func To_NominationPools_MemberAccount(in any) *pbgear.NominationPools_MemberAccount {
	out := &pbgear.NominationPools_MemberAccount{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MemberAccount_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MemberAccount_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MemberAccount_Id{
			Id: To_NominationPools_Id(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MemberAccount_Index{
			Index: To_NominationPools_Index(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MemberAccount_Raw{
			Raw: To_NominationPools_Raw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_MemberAccount_Address32{
			Address32: To_NominationPools_Address32(fields),
		}
	case 4:
		target = &pbgear.NominationPools_MemberAccount_Address20{
			Address20: To_NominationPools_Address20(fields),
		}
	}
}

func To_NominationPools_MinCreateBond(in any) *pbgear.NominationPools_MinCreateBond {
	out := &pbgear.NominationPools_MinCreateBond{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MinCreateBond_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MinCreateBond_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MinCreateBond_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MinCreateBond_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MinCreateBond_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_MinJoinBond(in any) *pbgear.NominationPools_MinJoinBond {
	out := &pbgear.NominationPools_MinJoinBond{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_MinJoinBond_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_MinJoinBond_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_MinJoinBond_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_MinJoinBond_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_MinJoinBond_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_NewBouncer(in any) *pbgear.NominationPools_NewBouncer {
	out := &pbgear.NominationPools_NewBouncer{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_NewBouncer_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_NewBouncer_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_NewBouncer_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_NewBouncer_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_NewBouncer_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_NewNominator(in any) *pbgear.NominationPools_NewNominator {
	out := &pbgear.NominationPools_NewNominator{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_NewNominator_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_NewNominator_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_NewNominator_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_NewNominator_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_NewNominator_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_NewRoot(in any) *pbgear.NominationPools_NewRoot {
	out := &pbgear.NominationPools_NewRoot{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_NewRoot_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_NewRoot_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_NewRoot_Noop{
			Noop: To_NominationPools_Noop(fields),
		}
	case 1:
		target = &pbgear.NominationPools_NewRoot_Set{
			Set: To_NominationPools_Set(fields),
		}
	case 2:
		target = &pbgear.NominationPools_NewRoot_Remove{
			Remove: To_NominationPools_Remove(fields),
		}
	}
}

func To_NominationPools_NominateCall(in any) *pbgear.NominationPools_NominateCall {
	out := &pbgear.NominationPools_NominateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// repeated field Validators
	out.Validators = To_Repeated_NominationPools_NominateCall_validators(fields[1].Value)

	return out //from message
}

func To_Repeated_NominationPools_NominateCall_validators(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_NominationPools_Nominator(in any) *pbgear.NominationPools_Nominator {
	out := &pbgear.NominationPools_Nominator{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Nominator_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Nominator_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Nominator_Id{
			Id: To_NominationPools_Id(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Nominator_Index{
			Index: To_NominationPools_Index(fields),
		}
	case 2:
		target = &pbgear.NominationPools_Nominator_Raw{
			Raw: To_NominationPools_Raw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_Nominator_Address32{
			Address32: To_NominationPools_Address32(fields),
		}
	case 4:
		target = &pbgear.NominationPools_Nominator_Address20{
			Address20: To_NominationPools_Address20(fields),
		}
	}
}

func To_NominationPools_Noop(in any) *pbgear.NominationPools_Noop {
	out := &pbgear.NominationPools_Noop{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_Open(in any) *pbgear.NominationPools_Open {
	out := &pbgear.NominationPools_Open{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_PalletNominationPoolsCommissionChangeRate(in any) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MaxIncrease
	out.MaxIncrease = To_NominationPools_PalletNominationPoolsCommissionChangeRate_max_increase(fields[0].Value)
	// primitive field MinDelay
	out.MinDelay = To_uint32(fields[1].Value)

	return out //from message
}

func To_NominationPools_PalletNominationPoolsCommissionChangeRate_max_increase(in any) *pbgear.SpArithmeticPerThingsPerbill {
	return nil //simple field
}

func To_NominationPools_Permission(in any) *pbgear.NominationPools_Permission {
	out := &pbgear.NominationPools_Permission{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Permission_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Permission_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Permission_Permissioned{
			Permissioned: To_NominationPools_Permissioned(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Permission_PermissionlessCompound{
			PermissionlessCompound: To_NominationPools_PermissionlessCompound(fields),
		}
	case 2:
		target = &pbgear.NominationPools_Permission_PermissionlessWithdraw{
			PermissionlessWithdraw: To_NominationPools_PermissionlessWithdraw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_Permission_PermissionlessAll{
			PermissionlessAll: To_NominationPools_PermissionlessAll(fields),
		}
	}
}

func To_NominationPools_Permissioned(in any) *pbgear.NominationPools_Permissioned {
	out := &pbgear.NominationPools_Permissioned{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_PermissionlessAll(in any) *pbgear.NominationPools_PermissionlessAll {
	out := &pbgear.NominationPools_PermissionlessAll{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_PermissionlessCompound(in any) *pbgear.NominationPools_PermissionlessCompound {
	out := &pbgear.NominationPools_PermissionlessCompound{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_PermissionlessWithdraw(in any) *pbgear.NominationPools_PermissionlessWithdraw {
	out := &pbgear.NominationPools_PermissionlessWithdraw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_PoolWithdrawUnbondedCall(in any) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
	out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(fields[1].Value)

	return out //from message
}

func To_NominationPools_Raw(in any) *pbgear.NominationPools_Raw {
	out := &pbgear.NominationPools_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_NominationPools_Remove(in any) *pbgear.NominationPools_Remove {
	out := &pbgear.NominationPools_Remove{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_Rewards(in any) *pbgear.NominationPools_Rewards {
	out := &pbgear.NominationPools_Rewards{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_NominationPools_Root(in any) *pbgear.NominationPools_Root {
	out := &pbgear.NominationPools_Root{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_Root_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_Root_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_Root_Id{
			Id: To_NominationPools_Id(fields),
		}
	case 1:
		target = &pbgear.NominationPools_Root_Index{
			Index: To_NominationPools_Index(fields),
		}
	case 2:
		target = &pbgear.NominationPools_Root_Raw{
			Raw: To_NominationPools_Raw(fields),
		}
	case 3:
		target = &pbgear.NominationPools_Root_Address32{
			Address32: To_NominationPools_Address32(fields),
		}
	case 4:
		target = &pbgear.NominationPools_Root_Address20{
			Address20: To_NominationPools_Address20(fields),
		}
	}
}

func To_NominationPools_Set(in any) *pbgear.NominationPools_Set {
	out := &pbgear.NominationPools_Set{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_string(fields[0].Value)

	return out //from message
}

func To_NominationPools_SetClaimPermissionCall(in any) *pbgear.NominationPools_SetClaimPermissionCall {
	out := &pbgear.NominationPools_SetClaimPermissionCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Permission
	out.Permission = To_NominationPools_SetClaimPermissionCall_permission(fields[0].Value)

	return out //from message
}

func To_NominationPools_SetClaimPermissionCall_permission(in any) *pbgear.NominationPools_Permission {
	return nil //simple field
}

func To_NominationPools_SetCommissionCall(in any) *pbgear.NominationPools_SetCommissionCall {
	out := &pbgear.NominationPools_SetCommissionCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// optional field NewCommission
	out.NewCommission = To_Optional_NominationPools_SetCommissionCall_new_commission(fields[1].Value)

	return out //from message
}

func To_Optional_NominationPools_SetCommissionCall_new_commission(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	return nil //funcForOptionalField
}

func To_NominationPools_SetCommissionChangeRateCall(in any) *pbgear.NominationPools_SetCommissionChangeRateCall {
	out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// field ChangeRate
	out.ChangeRate = To_NominationPools_SetCommissionChangeRateCall_change_rate(fields[1].Value)

	return out //from message
}

func To_NominationPools_SetCommissionChangeRateCall_change_rate(in any) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	return nil //simple field
}

func To_NominationPools_SetCommissionMaxCall(in any) *pbgear.NominationPools_SetCommissionMaxCall {
	out := &pbgear.NominationPools_SetCommissionMaxCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// field MaxCommission
	out.MaxCommission = To_NominationPools_SetCommissionMaxCall_max_commission(fields[1].Value)

	return out //from message
}

func To_NominationPools_SetCommissionMaxCall_max_commission(in any) *pbgear.SpArithmeticPerThingsPerbill {
	return nil //simple field
}

func To_NominationPools_SetConfigsCall(in any) *pbgear.NominationPools_SetConfigsCall {
	out := &pbgear.NominationPools_SetConfigsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MinJoinBond
	out.MinJoinBond = To_NominationPools_SetConfigsCall_min_join_bond(fields[0].Value)
	// field MinCreateBond
	out.MinCreateBond = To_NominationPools_SetConfigsCall_min_create_bond(fields[1].Value)
	// field MaxPools
	out.MaxPools = To_NominationPools_SetConfigsCall_max_pools(fields[2].Value)
	// field MaxMembers
	out.MaxMembers = To_NominationPools_SetConfigsCall_max_members(fields[3].Value)
	// field MaxMembersPerPool
	out.MaxMembersPerPool = To_NominationPools_SetConfigsCall_max_members_per_pool(fields[4].Value)
	// field GlobalMaxCommission
	out.GlobalMaxCommission = To_NominationPools_SetConfigsCall_global_max_commission(fields[5].Value)

	return out //from message
}

func To_NominationPools_SetConfigsCall_min_join_bond(in any) *pbgear.NominationPools_MinJoinBond {
	return nil //simple field
}
func To_NominationPools_SetConfigsCall_min_create_bond(in any) *pbgear.NominationPools_MinCreateBond {
	return nil //simple field
}
func To_NominationPools_SetConfigsCall_max_pools(in any) *pbgear.NominationPools_MaxPools {
	return nil //simple field
}
func To_NominationPools_SetConfigsCall_max_members(in any) *pbgear.NominationPools_MaxMembers {
	return nil //simple field
}
func To_NominationPools_SetConfigsCall_max_members_per_pool(in any) *pbgear.NominationPools_MaxMembersPerPool {
	return nil //simple field
}
func To_NominationPools_SetConfigsCall_global_max_commission(in any) *pbgear.NominationPools_GlobalMaxCommission {
	return nil //simple field
}

func To_NominationPools_SetMetadataCall(in any) *pbgear.NominationPools_SetMetadataCall {
	out := &pbgear.NominationPools_SetMetadataCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// primitive field Metadata
	out.Metadata = To_bytes(fields[1].Value)

	return out //from message
}

func To_NominationPools_SetStateCall(in any) *pbgear.NominationPools_SetStateCall {
	out := &pbgear.NominationPools_SetStateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// field State
	out.State = To_NominationPools_SetStateCall_state(fields[1].Value)

	return out //from message
}

func To_NominationPools_SetStateCall_state(in any) *pbgear.NominationPools_State {
	return nil //simple field
}

func To_NominationPools_State(in any) *pbgear.NominationPools_State {
	out := &pbgear.NominationPools_State{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_NominationPools_State_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_NominationPools_State_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.NominationPools_State_Open{
			Open: To_NominationPools_Open(fields),
		}
	case 1:
		target = &pbgear.NominationPools_State_Blocked{
			Blocked: To_NominationPools_Blocked(fields),
		}
	case 2:
		target = &pbgear.NominationPools_State_Destroying{
			Destroying: To_NominationPools_Destroying(fields),
		}
	}
}

func To_NominationPools_TupleNull(in any) *pbgear.NominationPools_TupleNull {
	out := &pbgear.NominationPools_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_NominationPools_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_NominationPools_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32_value1(fields[1].Value)

	return out //from message
}

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32_value0(in any) *pbgear.SpArithmeticPerThingsPerbill {
	return nil //simple field
}
func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32_value1(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_NominationPools_UnbondCall(in any) *pbgear.NominationPools_UnbondCall {
	out := &pbgear.NominationPools_UnbondCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MemberAccount
	out.MemberAccount = To_NominationPools_UnbondCall_member_account(fields[0].Value)
	// primitive field UnbondingPoints
	out.UnbondingPoints = To_string(fields[1].Value)

	return out //from message
}

func To_NominationPools_UnbondCall_member_account(in any) *pbgear.NominationPools_MemberAccount {
	return nil //simple field
}

func To_NominationPools_UpdateRolesCall(in any) *pbgear.NominationPools_UpdateRolesCall {
	out := &pbgear.NominationPools_UpdateRolesCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field PoolId
	out.PoolId = To_uint32(fields[0].Value)
	// field NewRoot
	out.NewRoot = To_NominationPools_UpdateRolesCall_new_root(fields[1].Value)
	// field NewNominator
	out.NewNominator = To_NominationPools_UpdateRolesCall_new_nominator(fields[2].Value)
	// field NewBouncer
	out.NewBouncer = To_NominationPools_UpdateRolesCall_new_bouncer(fields[3].Value)

	return out //from message
}

func To_NominationPools_UpdateRolesCall_new_root(in any) *pbgear.NominationPools_NewRoot {
	return nil //simple field
}
func To_NominationPools_UpdateRolesCall_new_nominator(in any) *pbgear.NominationPools_NewNominator {
	return nil //simple field
}
func To_NominationPools_UpdateRolesCall_new_bouncer(in any) *pbgear.NominationPools_NewBouncer {
	return nil //simple field
}

func To_NominationPools_WithdrawUnbondedCall(in any) *pbgear.NominationPools_WithdrawUnbondedCall {
	out := &pbgear.NominationPools_WithdrawUnbondedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MemberAccount
	out.MemberAccount = To_NominationPools_WithdrawUnbondedCall_member_account(fields[0].Value)
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(fields[1].Value)

	return out //from message
}

func To_NominationPools_WithdrawUnbondedCall_member_account(in any) *pbgear.NominationPools_MemberAccount {
	return nil //simple field
}

func To_None(in any) *pbgear.None {
	out := &pbgear.None{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_OffencesPallet(in any) *pbgear.OffencesPallet {
	out := &pbgear.OffencesPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_OriginsPallet(in any) *pbgear.OriginsPallet {
	out := &pbgear.OriginsPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_PreimagePallet(in any) *pbgear.PreimagePallet {
	out := &pbgear.PreimagePallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_PreimagePallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_PreimagePallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.PreimagePallet_NotePreimageCall{
			NotePreimageCall: To_Preimage_NotePreimageCall(fields),
		}
	case 1:
		target = &pbgear.PreimagePallet_UnnotePreimageCall{
			UnnotePreimageCall: To_Preimage_UnnotePreimageCall(fields),
		}
	case 2:
		target = &pbgear.PreimagePallet_RequestPreimageCall{
			RequestPreimageCall: To_Preimage_RequestPreimageCall(fields),
		}
	case 3:
		target = &pbgear.PreimagePallet_UnrequestPreimageCall{
			UnrequestPreimageCall: To_Preimage_UnrequestPreimageCall(fields),
		}
	case 4:
		target = &pbgear.PreimagePallet_EnsureUpdatedCall{
			EnsureUpdatedCall: To_Preimage_EnsureUpdatedCall(fields),
		}
	}
}

func To_Preimage_EnsureUpdatedCall(in any) *pbgear.Preimage_EnsureUpdatedCall {
	out := &pbgear.Preimage_EnsureUpdatedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Hashes
	out.Hashes = To_Repeated_Preimage_EnsureUpdatedCall_hashes(fields[0].Value)

	return out //from message
}

func To_Repeated_Preimage_EnsureUpdatedCall_hashes(in any) []*pbgear.PrimitiveTypesH256 {
	items := in.([]interface{})

	var out []*pbgear.PrimitiveTypesH256
	for _, item := range items {
		o := To_PrimitiveTypesH256(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Preimage_NotePreimageCall(in any) *pbgear.Preimage_NotePreimageCall {
	out := &pbgear.Preimage_NotePreimageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Bytes
	out.Bytes = To_bytes(fields[0].Value)

	return out //from message
}

func To_Preimage_RequestPreimageCall(in any) *pbgear.Preimage_RequestPreimageCall {
	out := &pbgear.Preimage_RequestPreimageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_Preimage_RequestPreimageCall_hash(fields[0].Value)

	return out //from message
}

func To_Preimage_RequestPreimageCall_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Preimage_UnnotePreimageCall(in any) *pbgear.Preimage_UnnotePreimageCall {
	out := &pbgear.Preimage_UnnotePreimageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_Preimage_UnnotePreimageCall_hash(fields[0].Value)

	return out //from message
}

func To_Preimage_UnnotePreimageCall_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Preimage_UnrequestPreimageCall(in any) *pbgear.Preimage_UnrequestPreimageCall {
	out := &pbgear.Preimage_UnrequestPreimageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_Preimage_UnrequestPreimageCall_hash(fields[0].Value)

	return out //from message
}

func To_Preimage_UnrequestPreimageCall_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_PrimaryAndSecondaryPlainSlots(in any) *pbgear.PrimaryAndSecondaryPlainSlots {
	out := &pbgear.PrimaryAndSecondaryPlainSlots{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_PrimaryAndSecondaryVrfSlots(in any) *pbgear.PrimaryAndSecondaryVrfSlots {
	out := &pbgear.PrimaryAndSecondaryVrfSlots{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_PrimarySlots(in any) *pbgear.PrimarySlots {
	out := &pbgear.PrimarySlots{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_PrimitiveTypesH256(in any) *pbgear.PrimitiveTypesH256 {
	out := &pbgear.PrimitiveTypesH256{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_ProxyPallet(in any) *pbgear.ProxyPallet {
	out := &pbgear.ProxyPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ProxyPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ProxyPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ProxyPallet_ProxyCall{
			ProxyCall: To_Proxy_ProxyCall(fields),
		}
	case 1:
		target = &pbgear.ProxyPallet_AddProxyCall{
			AddProxyCall: To_Proxy_AddProxyCall(fields),
		}
	case 2:
		target = &pbgear.ProxyPallet_RemoveProxyCall{
			RemoveProxyCall: To_Proxy_RemoveProxyCall(fields),
		}
	case 3:
		target = &pbgear.ProxyPallet_RemoveProxiesCall{
			RemoveProxiesCall: To_Proxy_RemoveProxiesCall(fields),
		}
	case 4:
		target = &pbgear.ProxyPallet_CreatePureCall{
			CreatePureCall: To_Proxy_CreatePureCall(fields),
		}
	case 5:
		target = &pbgear.ProxyPallet_KillPureCall{
			KillPureCall: To_Proxy_KillPureCall(fields),
		}
	case 6:
		target = &pbgear.ProxyPallet_AnnounceCall{
			AnnounceCall: To_Proxy_AnnounceCall(fields),
		}
	case 7:
		target = &pbgear.ProxyPallet_RemoveAnnouncementCall{
			RemoveAnnouncementCall: To_Proxy_RemoveAnnouncementCall(fields),
		}
	case 8:
		target = &pbgear.ProxyPallet_RejectAnnouncementCall{
			RejectAnnouncementCall: To_Proxy_RejectAnnouncementCall(fields),
		}
	case 9:
		target = &pbgear.ProxyPallet_ProxyAnnouncedCall{
			ProxyAnnouncedCall: To_Proxy_ProxyAnnouncedCall(fields),
		}
	}
}

func To_Proxy_AddProxyCall(in any) *pbgear.Proxy_AddProxyCall {
	out := &pbgear.Proxy_AddProxyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Delegate
	out.Delegate = To_Proxy_AddProxyCall_delegate(fields[0].Value)
	// field ProxyType
	out.ProxyType = To_Proxy_AddProxyCall_proxy_type(fields[1].Value)
	// primitive field Delay
	out.Delay = To_uint32(fields[2].Value)

	return out //from message
}

func To_Proxy_AddProxyCall_delegate(in any) *pbgear.Proxy_Delegate {
	return nil //simple field
}
func To_Proxy_AddProxyCall_proxy_type(in any) *pbgear.Proxy_ProxyType {
	return nil //simple field
}

func To_Proxy_Address20(in any) *pbgear.Proxy_Address20 {
	out := &pbgear.Proxy_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Proxy_Address32(in any) *pbgear.Proxy_Address32 {
	out := &pbgear.Proxy_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Proxy_AnnounceCall(in any) *pbgear.Proxy_AnnounceCall {
	out := &pbgear.Proxy_AnnounceCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Real
	out.Real = To_Proxy_AnnounceCall_real(fields[0].Value)
	// field CallHash
	out.CallHash = To_Proxy_AnnounceCall_call_hash(fields[1].Value)

	return out //from message
}

func To_Proxy_AnnounceCall_real(in any) *pbgear.Proxy_Real {
	return nil //simple field
}
func To_Proxy_AnnounceCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Proxy_Any(in any) *pbgear.Proxy_Any {
	out := &pbgear.Proxy_Any{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_CancelProxy(in any) *pbgear.Proxy_CancelProxy {
	out := &pbgear.Proxy_CancelProxy{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_CreatePureCall(in any) *pbgear.Proxy_CreatePureCall {
	out := &pbgear.Proxy_CreatePureCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ProxyType
	out.ProxyType = To_Proxy_CreatePureCall_proxy_type(fields[0].Value)
	// primitive field Delay
	out.Delay = To_uint32(fields[1].Value)
	// primitive field Index
	out.Index = To_uint32(fields[2].Value)

	return out //from message
}

func To_Proxy_CreatePureCall_proxy_type(in any) *pbgear.Proxy_ProxyType {
	return nil //simple field
}

func To_Proxy_Delegate(in any) *pbgear.Proxy_Delegate {
	out := &pbgear.Proxy_Delegate{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Proxy_Delegate_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Proxy_Delegate_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_Delegate_Id{
			Id: To_Proxy_Id(fields),
		}
	case 1:
		target = &pbgear.Proxy_Delegate_Index{
			Index: To_Proxy_Index(fields),
		}
	case 2:
		target = &pbgear.Proxy_Delegate_Raw{
			Raw: To_Proxy_Raw(fields),
		}
	case 3:
		target = &pbgear.Proxy_Delegate_Address32{
			Address32: To_Proxy_Address32(fields),
		}
	case 4:
		target = &pbgear.Proxy_Delegate_Address20{
			Address20: To_Proxy_Address20(fields),
		}
	}
}

func To_Proxy_ForceProxyType(in any) *pbgear.Proxy_ForceProxyType {
	out := &pbgear.Proxy_ForceProxyType{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Proxy_ForceProxyType_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Proxy_ForceProxyType_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_ForceProxyType_Any{
			Any: To_Proxy_Any(fields),
		}
	case 1:
		target = &pbgear.Proxy_ForceProxyType_NonTransfer{
			NonTransfer: To_Proxy_NonTransfer(fields),
		}
	case 2:
		target = &pbgear.Proxy_ForceProxyType_Governance{
			Governance: To_Proxy_Governance(fields),
		}
	case 3:
		target = &pbgear.Proxy_ForceProxyType_Staking{
			Staking: To_Proxy_Staking(fields),
		}
	case 4:
		target = &pbgear.Proxy_ForceProxyType_IdentityJudgement{
			IdentityJudgement: To_Proxy_IdentityJudgement(fields),
		}
	case 5:
		target = &pbgear.Proxy_ForceProxyType_CancelProxy{
			CancelProxy: To_Proxy_CancelProxy(fields),
		}
	}
}

func To_Proxy_Governance(in any) *pbgear.Proxy_Governance {
	out := &pbgear.Proxy_Governance{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_Id(in any) *pbgear.Proxy_Id {
	out := &pbgear.Proxy_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Proxy_Id_value0(fields[0].Value)

	return out //from message
}

func To_Proxy_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Proxy_IdentityJudgement(in any) *pbgear.Proxy_IdentityJudgement {
	out := &pbgear.Proxy_IdentityJudgement{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_Index(in any) *pbgear.Proxy_Index {
	out := &pbgear.Proxy_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Proxy_Index_value0(fields[0].Value)

	return out //from message
}

func To_Proxy_Index_value0(in any) *pbgear.Proxy_TupleNull {
	return nil //simple field
}

func To_Proxy_KillPureCall(in any) *pbgear.Proxy_KillPureCall {
	out := &pbgear.Proxy_KillPureCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Spawner
	out.Spawner = To_Proxy_KillPureCall_spawner(fields[0].Value)
	// field ProxyType
	out.ProxyType = To_Proxy_KillPureCall_proxy_type(fields[1].Value)
	// primitive field Index
	out.Index = To_uint32(fields[2].Value)
	// primitive field Height
	out.Height = To_uint32(fields[3].Value)
	// primitive field ExtIndex
	out.ExtIndex = To_uint32(fields[4].Value)

	return out //from message
}

func To_Proxy_KillPureCall_spawner(in any) *pbgear.Proxy_Spawner {
	return nil //simple field
}
func To_Proxy_KillPureCall_proxy_type(in any) *pbgear.Proxy_ProxyType {
	return nil //simple field
}

func To_Proxy_NonTransfer(in any) *pbgear.Proxy_NonTransfer {
	out := &pbgear.Proxy_NonTransfer{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_ProxyAnnouncedCall(in any) *pbgear.Proxy_ProxyAnnouncedCall {
	out := &pbgear.Proxy_ProxyAnnouncedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Delegate
	out.Delegate = To_Proxy_ProxyAnnouncedCall_delegate(fields[0].Value)
	// field Real
	out.Real = To_Proxy_ProxyAnnouncedCall_real(fields[1].Value)
	// optional field ForceProxyType
	out.ForceProxyType = To_Optional_Proxy_ProxyAnnouncedCall_force_proxy_type(fields[2].Value)
	// oneOf field Call
	Set_OneOf_Proxy_ProxyAnnouncedCall_call(out.Call, fields[3].Value)

	return out //from message
}

func To_Proxy_ProxyAnnouncedCall_delegate(in any) *pbgear.Proxy_Delegate {
	return nil //simple field
}
func To_Proxy_ProxyAnnouncedCall_real(in any) *pbgear.Proxy_Real {
	return nil //simple field
}
func To_Optional_Proxy_ProxyAnnouncedCall_force_proxy_type(in any) *pbgear.Proxy_ForceProxyType {
	return nil //funcForOptionalField
}
func Set_OneOf_Proxy_ProxyAnnouncedCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_ProxyAnnouncedCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Proxy_ProxyAnnouncedCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Proxy_ProxyAnnouncedCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Proxy_ProxyAnnouncedCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Proxy_ProxyAnnouncedCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Proxy_ProxyAnnouncedCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Proxy_ProxyAnnouncedCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Proxy_ProxyAnnouncedCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Proxy_ProxyAnnouncedCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Proxy_ProxyAnnouncedCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Proxy_ProxyAnnouncedCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Proxy_ProxyCall(in any) *pbgear.Proxy_ProxyCall {
	out := &pbgear.Proxy_ProxyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Real
	out.Real = To_Proxy_ProxyCall_real(fields[0].Value)
	// optional field ForceProxyType
	out.ForceProxyType = To_Optional_Proxy_ProxyCall_force_proxy_type(fields[1].Value)
	// oneOf field Call
	Set_OneOf_Proxy_ProxyCall_call(out.Call, fields[2].Value)

	return out //from message
}

func To_Proxy_ProxyCall_real(in any) *pbgear.Proxy_Real {
	return nil //simple field
}
func To_Optional_Proxy_ProxyCall_force_proxy_type(in any) *pbgear.Proxy_ForceProxyType {
	return nil //funcForOptionalField
}
func Set_OneOf_Proxy_ProxyCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_ProxyCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Proxy_ProxyCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Proxy_ProxyCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Proxy_ProxyCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Proxy_ProxyCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Proxy_ProxyCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Proxy_ProxyCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Proxy_ProxyCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Proxy_ProxyCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Proxy_ProxyCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Proxy_ProxyCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Proxy_ProxyCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Proxy_ProxyCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Proxy_ProxyCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Proxy_ProxyCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Proxy_ProxyCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Proxy_ProxyCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Proxy_ProxyCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Proxy_ProxyCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Proxy_ProxyCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Proxy_ProxyCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Proxy_ProxyCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Proxy_ProxyCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Proxy_ProxyCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Proxy_ProxyCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Proxy_ProxyCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Proxy_ProxyCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Proxy_ProxyCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Proxy_ProxyCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Proxy_ProxyType(in any) *pbgear.Proxy_ProxyType {
	out := &pbgear.Proxy_ProxyType{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Proxy_ProxyType_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Proxy_ProxyType_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_ProxyType_Any{
			Any: To_Proxy_Any(fields),
		}
	case 1:
		target = &pbgear.Proxy_ProxyType_NonTransfer{
			NonTransfer: To_Proxy_NonTransfer(fields),
		}
	case 2:
		target = &pbgear.Proxy_ProxyType_Governance{
			Governance: To_Proxy_Governance(fields),
		}
	case 3:
		target = &pbgear.Proxy_ProxyType_Staking{
			Staking: To_Proxy_Staking(fields),
		}
	case 4:
		target = &pbgear.Proxy_ProxyType_IdentityJudgement{
			IdentityJudgement: To_Proxy_IdentityJudgement(fields),
		}
	case 5:
		target = &pbgear.Proxy_ProxyType_CancelProxy{
			CancelProxy: To_Proxy_CancelProxy(fields),
		}
	}
}

func To_Proxy_Raw(in any) *pbgear.Proxy_Raw {
	out := &pbgear.Proxy_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Proxy_Real(in any) *pbgear.Proxy_Real {
	out := &pbgear.Proxy_Real{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Proxy_Real_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Proxy_Real_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_Real_Id{
			Id: To_Proxy_Id(fields),
		}
	case 1:
		target = &pbgear.Proxy_Real_Index{
			Index: To_Proxy_Index(fields),
		}
	case 2:
		target = &pbgear.Proxy_Real_Raw{
			Raw: To_Proxy_Raw(fields),
		}
	case 3:
		target = &pbgear.Proxy_Real_Address32{
			Address32: To_Proxy_Address32(fields),
		}
	case 4:
		target = &pbgear.Proxy_Real_Address20{
			Address20: To_Proxy_Address20(fields),
		}
	}
}

func To_Proxy_RejectAnnouncementCall(in any) *pbgear.Proxy_RejectAnnouncementCall {
	out := &pbgear.Proxy_RejectAnnouncementCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Delegate
	out.Delegate = To_Proxy_RejectAnnouncementCall_delegate(fields[0].Value)
	// field CallHash
	out.CallHash = To_Proxy_RejectAnnouncementCall_call_hash(fields[1].Value)

	return out //from message
}

func To_Proxy_RejectAnnouncementCall_delegate(in any) *pbgear.Proxy_Delegate {
	return nil //simple field
}
func To_Proxy_RejectAnnouncementCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Proxy_RemoveAnnouncementCall(in any) *pbgear.Proxy_RemoveAnnouncementCall {
	out := &pbgear.Proxy_RemoveAnnouncementCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Real
	out.Real = To_Proxy_RemoveAnnouncementCall_real(fields[0].Value)
	// field CallHash
	out.CallHash = To_Proxy_RemoveAnnouncementCall_call_hash(fields[1].Value)

	return out //from message
}

func To_Proxy_RemoveAnnouncementCall_real(in any) *pbgear.Proxy_Real {
	return nil //simple field
}
func To_Proxy_RemoveAnnouncementCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Proxy_RemoveProxiesCall(in any) *pbgear.Proxy_RemoveProxiesCall {
	out := &pbgear.Proxy_RemoveProxiesCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_RemoveProxyCall(in any) *pbgear.Proxy_RemoveProxyCall {
	out := &pbgear.Proxy_RemoveProxyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Delegate
	out.Delegate = To_Proxy_RemoveProxyCall_delegate(fields[0].Value)
	// field ProxyType
	out.ProxyType = To_Proxy_RemoveProxyCall_proxy_type(fields[1].Value)
	// primitive field Delay
	out.Delay = To_uint32(fields[2].Value)

	return out //from message
}

func To_Proxy_RemoveProxyCall_delegate(in any) *pbgear.Proxy_Delegate {
	return nil //simple field
}
func To_Proxy_RemoveProxyCall_proxy_type(in any) *pbgear.Proxy_ProxyType {
	return nil //simple field
}

func To_Proxy_Spawner(in any) *pbgear.Proxy_Spawner {
	out := &pbgear.Proxy_Spawner{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Proxy_Spawner_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Proxy_Spawner_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Proxy_Spawner_Id{
			Id: To_Proxy_Id(fields),
		}
	case 1:
		target = &pbgear.Proxy_Spawner_Index{
			Index: To_Proxy_Index(fields),
		}
	case 2:
		target = &pbgear.Proxy_Spawner_Raw{
			Raw: To_Proxy_Raw(fields),
		}
	case 3:
		target = &pbgear.Proxy_Spawner_Address32{
			Address32: To_Proxy_Address32(fields),
		}
	case 4:
		target = &pbgear.Proxy_Spawner_Address20{
			Address20: To_Proxy_Address20(fields),
		}
	}
}

func To_Proxy_Staking(in any) *pbgear.Proxy_Staking {
	out := &pbgear.Proxy_Staking{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Proxy_TupleNull(in any) *pbgear.Proxy_TupleNull {
	out := &pbgear.Proxy_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Proxy_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Proxy_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_ReferendaPallet(in any) *pbgear.ReferendaPallet {
	out := &pbgear.ReferendaPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_ReferendaPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_ReferendaPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.ReferendaPallet_SubmitCall{
			SubmitCall: To_Referenda_SubmitCall(fields),
		}
	case 1:
		target = &pbgear.ReferendaPallet_PlaceDecisionDepositCall{
			PlaceDecisionDepositCall: To_Referenda_PlaceDecisionDepositCall(fields),
		}
	case 2:
		target = &pbgear.ReferendaPallet_RefundDecisionDepositCall{
			RefundDecisionDepositCall: To_Referenda_RefundDecisionDepositCall(fields),
		}
	case 3:
		target = &pbgear.ReferendaPallet_CancelCall{
			CancelCall: To_Referenda_CancelCall(fields),
		}
	case 4:
		target = &pbgear.ReferendaPallet_KillCall{
			KillCall: To_Referenda_KillCall(fields),
		}
	case 5:
		target = &pbgear.ReferendaPallet_NudgeReferendumCall{
			NudgeReferendumCall: To_Referenda_NudgeReferendumCall(fields),
		}
	case 6:
		target = &pbgear.ReferendaPallet_OneFewerDecidingCall{
			OneFewerDecidingCall: To_Referenda_OneFewerDecidingCall(fields),
		}
	case 7:
		target = &pbgear.ReferendaPallet_RefundSubmissionDepositCall{
			RefundSubmissionDepositCall: To_Referenda_RefundSubmissionDepositCall(fields),
		}
	case 8:
		target = &pbgear.ReferendaPallet_SetMetadataCall{
			SetMetadataCall: To_Referenda_SetMetadataCall(fields),
		}
	}
}

func To_Referenda_After(in any) *pbgear.Referenda_After {
	out := &pbgear.Referenda_After{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_At(in any) *pbgear.Referenda_At {
	out := &pbgear.Referenda_At{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_CancelCall(in any) *pbgear.Referenda_CancelCall {
	out := &pbgear.Referenda_CancelCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_EnactmentMoment(in any) *pbgear.Referenda_EnactmentMoment {
	out := &pbgear.Referenda_EnactmentMoment{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Referenda_EnactmentMoment_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Referenda_EnactmentMoment_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Referenda_EnactmentMoment_At{
			At: To_Referenda_At(fields),
		}
	case 1:
		target = &pbgear.Referenda_EnactmentMoment_After{
			After: To_Referenda_After(fields),
		}
	}
}

func To_Referenda_Inline(in any) *pbgear.Referenda_Inline {
	out := &pbgear.Referenda_Inline{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Referenda_Inline_value0(fields[0].Value)

	return out //from message
}

func To_Referenda_Inline_value0(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	return nil //simple field
}

func To_Referenda_KillCall(in any) *pbgear.Referenda_KillCall {
	out := &pbgear.Referenda_KillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_Legacy(in any) *pbgear.Referenda_Legacy {
	out := &pbgear.Referenda_Legacy{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_Referenda_Legacy_hash(fields[0].Value)

	return out //from message
}

func To_Referenda_Legacy_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Referenda_Lookup(in any) *pbgear.Referenda_Lookup {
	out := &pbgear.Referenda_Lookup{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Hash
	out.Hash = To_Referenda_Lookup_hash(fields[0].Value)
	// primitive field Len
	out.Len = To_uint32(fields[1].Value)

	return out //from message
}

func To_Referenda_Lookup_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Referenda_NudgeReferendumCall(in any) *pbgear.Referenda_NudgeReferendumCall {
	out := &pbgear.Referenda_NudgeReferendumCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_OneFewerDecidingCall(in any) *pbgear.Referenda_OneFewerDecidingCall {
	out := &pbgear.Referenda_OneFewerDecidingCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Track
	out.Track = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_Origins(in any) *pbgear.Referenda_Origins {
	out := &pbgear.Referenda_Origins{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Referenda_Origins_value0(fields[0].Value)

	return out //from message
}

func To_Referenda_Origins_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Referenda_PlaceDecisionDepositCall(in any) *pbgear.Referenda_PlaceDecisionDepositCall {
	out := &pbgear.Referenda_PlaceDecisionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_Proposal(in any) *pbgear.Referenda_Proposal {
	out := &pbgear.Referenda_Proposal{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Referenda_Proposal_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Referenda_Proposal_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Referenda_Proposal_Legacy{
			Legacy: To_Referenda_Legacy(fields),
		}
	case 1:
		target = &pbgear.Referenda_Proposal_Inline{
			Inline: To_Referenda_Inline(fields),
		}
	case 2:
		target = &pbgear.Referenda_Proposal_Lookup{
			Lookup: To_Referenda_Lookup(fields),
		}
	}
}

func To_Referenda_ProposalOrigin(in any) *pbgear.Referenda_ProposalOrigin {
	out := &pbgear.Referenda_ProposalOrigin{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Referenda_ProposalOrigin_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Referenda_ProposalOrigin_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Referenda_ProposalOrigin_System{
			System: To_Referenda_System(fields),
		}
	case 20:
		target = &pbgear.Referenda_ProposalOrigin_Origins{
			Origins: To_Referenda_Origins(fields),
		}
	case 2:
		target = &pbgear.Referenda_ProposalOrigin_Void{
			Void: To_Referenda_Void(fields),
		}
	}
}

func To_Referenda_RefundDecisionDepositCall(in any) *pbgear.Referenda_RefundDecisionDepositCall {
	out := &pbgear.Referenda_RefundDecisionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_RefundSubmissionDepositCall(in any) *pbgear.Referenda_RefundSubmissionDepositCall {
	out := &pbgear.Referenda_RefundSubmissionDepositCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Referenda_SetMetadataCall(in any) *pbgear.Referenda_SetMetadataCall {
	out := &pbgear.Referenda_SetMetadataCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// optional field MaybeHash
	out.MaybeHash = To_Optional_Referenda_SetMetadataCall_maybe_hash(fields[1].Value)

	return out //from message
}

func To_Optional_Referenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //funcForOptionalField
}

func To_Referenda_SubmitCall(in any) *pbgear.Referenda_SubmitCall {
	out := &pbgear.Referenda_SubmitCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ProposalOrigin
	out.ProposalOrigin = To_Referenda_SubmitCall_proposal_origin(fields[0].Value)
	// field Proposal
	out.Proposal = To_Referenda_SubmitCall_proposal(fields[1].Value)
	// field EnactmentMoment
	out.EnactmentMoment = To_Referenda_SubmitCall_enactment_moment(fields[2].Value)

	return out //from message
}

func To_Referenda_SubmitCall_proposal_origin(in any) *pbgear.Referenda_ProposalOrigin {
	return nil //simple field
}
func To_Referenda_SubmitCall_proposal(in any) *pbgear.Referenda_Proposal {
	return nil //simple field
}
func To_Referenda_SubmitCall_enactment_moment(in any) *pbgear.Referenda_EnactmentMoment {
	return nil //simple field
}

func To_Referenda_System(in any) *pbgear.Referenda_System {
	out := &pbgear.Referenda_System{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Referenda_System_value0(fields[0].Value)

	return out //from message
}

func To_Referenda_System_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Referenda_Void(in any) *pbgear.Referenda_Void {
	out := &pbgear.Referenda_Void{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Referenda_Void_value0(fields[0].Value)

	return out //from message
}

func To_Referenda_Void_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Root(in any) *pbgear.Root {
	out := &pbgear.Root{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_SchedulerPallet(in any) *pbgear.SchedulerPallet {
	out := &pbgear.SchedulerPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_SchedulerPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_SchedulerPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.SchedulerPallet_ScheduleCall{
			ScheduleCall: To_Scheduler_ScheduleCall(fields),
		}
	case 1:
		target = &pbgear.SchedulerPallet_CancelCall{
			CancelCall: To_Scheduler_CancelCall(fields),
		}
	case 2:
		target = &pbgear.SchedulerPallet_ScheduleNamedCall{
			ScheduleNamedCall: To_Scheduler_ScheduleNamedCall(fields),
		}
	case 3:
		target = &pbgear.SchedulerPallet_CancelNamedCall{
			CancelNamedCall: To_Scheduler_CancelNamedCall(fields),
		}
	case 4:
		target = &pbgear.SchedulerPallet_ScheduleAfterCall{
			ScheduleAfterCall: To_Scheduler_ScheduleAfterCall(fields),
		}
	case 5:
		target = &pbgear.SchedulerPallet_ScheduleNamedAfterCall{
			ScheduleNamedAfterCall: To_Scheduler_ScheduleNamedAfterCall(fields),
		}
	}
}

func To_Scheduler_CancelCall(in any) *pbgear.Scheduler_CancelCall {
	out := &pbgear.Scheduler_CancelCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field When
	out.When = To_uint32(fields[0].Value)
	// primitive field Index
	out.Index = To_uint32(fields[1].Value)

	return out //from message
}

func To_Scheduler_CancelNamedCall(in any) *pbgear.Scheduler_CancelNamedCall {
	out := &pbgear.Scheduler_CancelNamedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Id
	out.Id = To_bytes(fields[0].Value)

	return out //from message
}

func To_Scheduler_ScheduleAfterCall(in any) *pbgear.Scheduler_ScheduleAfterCall {
	out := &pbgear.Scheduler_ScheduleAfterCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field After
	out.After = To_uint32(fields[0].Value)
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(fields[1].Value)
	// primitive field Priority
	out.Priority = To_uint32(fields[2].Value)
	// oneOf field Call
	Set_OneOf_Scheduler_ScheduleAfterCall_call(out.Call, fields[3].Value)

	return out //from message
}

func To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	return nil //funcForOptionalField
}
func Set_OneOf_Scheduler_ScheduleAfterCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Scheduler_ScheduleAfterCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Scheduler_ScheduleAfterCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Scheduler_ScheduleAfterCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Scheduler_ScheduleAfterCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Scheduler_ScheduleAfterCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Scheduler_ScheduleAfterCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Scheduler_ScheduleAfterCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Scheduler_ScheduleAfterCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Scheduler_ScheduleAfterCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Scheduler_ScheduleAfterCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Scheduler_ScheduleAfterCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Scheduler_ScheduleAfterCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Scheduler_ScheduleAfterCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Scheduler_ScheduleAfterCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Scheduler_ScheduleAfterCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Scheduler_ScheduleAfterCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Scheduler_ScheduleAfterCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Scheduler_ScheduleAfterCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Scheduler_ScheduleAfterCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Scheduler_ScheduleAfterCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Scheduler_ScheduleAfterCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Scheduler_ScheduleAfterCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Scheduler_ScheduleAfterCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Scheduler_ScheduleAfterCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Scheduler_ScheduleAfterCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Scheduler_ScheduleAfterCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Scheduler_ScheduleAfterCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Scheduler_ScheduleAfterCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Scheduler_ScheduleAfterCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Scheduler_ScheduleCall(in any) *pbgear.Scheduler_ScheduleCall {
	out := &pbgear.Scheduler_ScheduleCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field When
	out.When = To_uint32(fields[0].Value)
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleCall_maybe_periodic(fields[1].Value)
	// primitive field Priority
	out.Priority = To_uint32(fields[2].Value)
	// oneOf field Call
	Set_OneOf_Scheduler_ScheduleCall_call(out.Call, fields[3].Value)

	return out //from message
}

func To_Optional_Scheduler_ScheduleCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	return nil //funcForOptionalField
}
func Set_OneOf_Scheduler_ScheduleCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Scheduler_ScheduleCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Scheduler_ScheduleCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Scheduler_ScheduleCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Scheduler_ScheduleCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Scheduler_ScheduleCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Scheduler_ScheduleCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Scheduler_ScheduleCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Scheduler_ScheduleCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Scheduler_ScheduleCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Scheduler_ScheduleCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Scheduler_ScheduleCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Scheduler_ScheduleCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Scheduler_ScheduleCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Scheduler_ScheduleCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Scheduler_ScheduleCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Scheduler_ScheduleCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Scheduler_ScheduleCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Scheduler_ScheduleCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Scheduler_ScheduleCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Scheduler_ScheduleCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Scheduler_ScheduleCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Scheduler_ScheduleCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Scheduler_ScheduleCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Scheduler_ScheduleCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Scheduler_ScheduleCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Scheduler_ScheduleCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Scheduler_ScheduleCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Scheduler_ScheduleCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Scheduler_ScheduleCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Scheduler_ScheduleNamedAfterCall(in any) *pbgear.Scheduler_ScheduleNamedAfterCall {
	out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Id
	out.Id = To_bytes(fields[0].Value)
	// primitive field After
	out.After = To_uint32(fields[1].Value)
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(fields[2].Value)
	// primitive field Priority
	out.Priority = To_uint32(fields[3].Value)
	// oneOf field Call
	Set_OneOf_Scheduler_ScheduleNamedAfterCall_call(out.Call, fields[4].Value)

	return out //from message
}

func To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	return nil //funcForOptionalField
}
func Set_OneOf_Scheduler_ScheduleNamedAfterCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Scheduler_ScheduleNamedAfterCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Scheduler_ScheduleNamedCall(in any) *pbgear.Scheduler_ScheduleNamedCall {
	out := &pbgear.Scheduler_ScheduleNamedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Id
	out.Id = To_bytes(fields[0].Value)
	// primitive field When
	out.When = To_uint32(fields[1].Value)
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(fields[2].Value)
	// primitive field Priority
	out.Priority = To_uint32(fields[3].Value)
	// oneOf field Call
	Set_OneOf_Scheduler_ScheduleNamedCall_call(out.Call, fields[4].Value)

	return out //from message
}

func To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	return nil //funcForOptionalField
}
func Set_OneOf_Scheduler_ScheduleNamedCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Scheduler_ScheduleNamedCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Scheduler_ScheduleNamedCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Scheduler_ScheduleNamedCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Scheduler_ScheduleNamedCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Scheduler_ScheduleNamedCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Scheduler_ScheduleNamedCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Scheduler_ScheduleNamedCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Scheduler_ScheduleNamedCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Scheduler_ScheduleNamedCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Scheduler_ScheduleNamedCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Scheduler_ScheduleNamedCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Scheduler_ScheduleNamedCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Scheduler_ScheduleNamedCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Scheduler_ScheduleNamedCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Scheduler_ScheduleNamedCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Scheduler_ScheduleNamedCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Scheduler_ScheduleNamedCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Scheduler_ScheduleNamedCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Scheduler_ScheduleNamedCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Scheduler_ScheduleNamedCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Scheduler_ScheduleNamedCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Scheduler_ScheduleNamedCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Scheduler_ScheduleNamedCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Scheduler_ScheduleNamedCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Scheduler_ScheduleNamedCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Scheduler_ScheduleNamedCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Scheduler_ScheduleNamedCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Scheduler_ScheduleNamedCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Scheduler_ScheduleNamedCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Scheduler_TupleUint32Uint32(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	out := &pbgear.Scheduler_TupleUint32Uint32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_uint32(fields[1].Value)

	return out //from message
}

func To_SessionPallet(in any) *pbgear.SessionPallet {
	out := &pbgear.SessionPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_SessionPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_SessionPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.SessionPallet_SetKeysCall{
			SetKeysCall: To_Session_SetKeysCall(fields),
		}
	case 1:
		target = &pbgear.SessionPallet_PurgeKeysCall{
			PurgeKeysCall: To_Session_PurgeKeysCall(fields),
		}
	}
}

func To_Session_PurgeKeysCall(in any) *pbgear.Session_PurgeKeysCall {
	out := &pbgear.Session_PurgeKeysCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Session_SetKeysCall(in any) *pbgear.Session_SetKeysCall {
	out := &pbgear.Session_SetKeysCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Keys
	out.Keys = To_Session_SetKeysCall_keys(fields[0].Value)
	// primitive field Proof
	out.Proof = To_bytes(fields[1].Value)

	return out //from message
}

func To_Session_SetKeysCall_keys(in any) *pbgear.VaraRuntimeSessionKeys {
	return nil //simple field
}

func To_Signed(in any) *pbgear.Signed {
	out := &pbgear.Signed{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Signed_value0(fields[0].Value)

	return out //from message
}

func To_Signed_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_SpArithmeticPerThingsPerU16(in any) *pbgear.SpArithmeticPerThingsPerU16 {
	out := &pbgear.SpArithmeticPerThingsPerU16{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_SpArithmeticPerThingsPerbill(in any) *pbgear.SpArithmeticPerThingsPerbill {
	out := &pbgear.SpArithmeticPerThingsPerbill{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_SpArithmeticPerThingsPercent(in any) *pbgear.SpArithmeticPerThingsPercent {
	out := &pbgear.SpArithmeticPerThingsPercent{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint32(fields[0].Value)

	return out //from message
}

func To_SpAuthorityDiscoveryAppPublic(in any) *pbgear.SpAuthorityDiscoveryAppPublic {
	out := &pbgear.SpAuthorityDiscoveryAppPublic{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_SpAuthorityDiscoveryAppPublic_value0(fields[0].Value)

	return out //from message
}

func To_SpAuthorityDiscoveryAppPublic_value0(in any) *pbgear.SpCoreSr25519Public {
	return nil //simple field
}

func To_SpConsensusBabeAppPublic(in any) *pbgear.SpConsensusBabeAppPublic {
	out := &pbgear.SpConsensusBabeAppPublic{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_SpConsensusBabeAppPublic_value0(fields[0].Value)

	return out //from message
}

func To_SpConsensusBabeAppPublic_value0(in any) *pbgear.SpCoreSr25519Public {
	return nil //simple field
}

func To_SpConsensusGrandpaAppPublic(in any) *pbgear.SpConsensusGrandpaAppPublic {
	out := &pbgear.SpConsensusGrandpaAppPublic{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_SpConsensusGrandpaAppPublic_value0(fields[0].Value)

	return out //from message
}

func To_SpConsensusGrandpaAppPublic_value0(in any) *pbgear.SpCoreEd25519Public {
	return nil //simple field
}

func To_SpConsensusGrandpaAppSignature(in any) *pbgear.SpConsensusGrandpaAppSignature {
	out := &pbgear.SpConsensusGrandpaAppSignature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_SpConsensusGrandpaAppSignature_value0(fields[0].Value)

	return out //from message
}

func To_SpConsensusGrandpaAppSignature_value0(in any) *pbgear.SpCoreEd25519Signature {
	return nil //simple field
}

func To_SpConsensusGrandpaEquivocationProof(in any) *pbgear.SpConsensusGrandpaEquivocationProof {
	out := &pbgear.SpConsensusGrandpaEquivocationProof{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field SetId
	out.SetId = To_uint64(fields[0].Value)
	// field Equivocation
	out.Equivocation = To_SpConsensusGrandpaEquivocationProof_equivocation(fields[1].Value)

	return out //from message
}

func To_SpConsensusGrandpaEquivocationProof_equivocation(in any) *pbgear.Grandpa_Equivocation {
	return nil //simple field
}

func To_SpConsensusSlotsEquivocationProof(in any) *pbgear.SpConsensusSlotsEquivocationProof {
	out := &pbgear.SpConsensusSlotsEquivocationProof{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Offender
	out.Offender = To_SpConsensusSlotsEquivocationProof_offender(fields[0].Value)
	// field Slot
	out.Slot = To_SpConsensusSlotsEquivocationProof_slot(fields[1].Value)
	// field FirstHeader
	out.FirstHeader = To_SpConsensusSlotsEquivocationProof_first_header(fields[2].Value)
	// field SecondHeader
	out.SecondHeader = To_SpConsensusSlotsEquivocationProof_second_header(fields[3].Value)

	return out //from message
}

func To_SpConsensusSlotsEquivocationProof_offender(in any) *pbgear.SpConsensusBabeAppPublic {
	return nil //simple field
}
func To_SpConsensusSlotsEquivocationProof_slot(in any) *pbgear.SpConsensusSlotsSlot {
	return nil //simple field
}
func To_SpConsensusSlotsEquivocationProof_first_header(in any) *pbgear.SpRuntimeGenericHeaderHeader {
	return nil //simple field
}
func To_SpConsensusSlotsEquivocationProof_second_header(in any) *pbgear.SpRuntimeGenericHeaderHeader {
	return nil //simple field
}

func To_SpConsensusSlotsSlot(in any) *pbgear.SpConsensusSlotsSlot {
	out := &pbgear.SpConsensusSlotsSlot{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_uint64(fields[0].Value)

	return out //from message
}

func To_SpCoreCryptoAccountId32(in any) *pbgear.SpCoreCryptoAccountId32 {
	out := &pbgear.SpCoreCryptoAccountId32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_SpCoreEd25519Public(in any) *pbgear.SpCoreEd25519Public {
	out := &pbgear.SpCoreEd25519Public{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_SpCoreEd25519Signature(in any) *pbgear.SpCoreEd25519Signature {
	out := &pbgear.SpCoreEd25519Signature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_SpCoreSr25519Public(in any) *pbgear.SpCoreSr25519Public {
	out := &pbgear.SpCoreSr25519Public{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_SpCoreSr25519Signature(in any) *pbgear.SpCoreSr25519Signature {
	out := &pbgear.SpCoreSr25519Signature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_SpNposElectionsElectionScore(in any) *pbgear.SpNposElectionsElectionScore {
	out := &pbgear.SpNposElectionsElectionScore{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field MinimalStake
	out.MinimalStake = To_string(fields[0].Value)
	// primitive field SumStake
	out.SumStake = To_string(fields[1].Value)
	// primitive field SumStakeSquared
	out.SumStakeSquared = To_string(fields[2].Value)

	return out //from message
}

func To_SpNposElectionsSupport(in any) *pbgear.SpNposElectionsSupport {
	out := &pbgear.SpNposElectionsSupport{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Total
	out.Total = To_string(fields[0].Value)
	// repeated field Voters
	out.Voters = To_Repeated_SpNposElectionsSupport_voters(fields[1].Value)

	return out //from message
}

func To_Repeated_SpNposElectionsSupport_voters(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_SpRuntimeGenericDigestDigest(in any) *pbgear.SpRuntimeGenericDigestDigest {
	out := &pbgear.SpRuntimeGenericDigestDigest{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Logs
	out.Logs = To_Repeated_SpRuntimeGenericDigestDigest_logs(fields[0].Value)

	return out //from message
}

func To_Repeated_SpRuntimeGenericDigestDigest_logs(in any) []*pbgear.SpRuntimeGenericDigestDigestItem {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeGenericDigestDigestItem
	for _, item := range items {
		o := To_SpRuntimeGenericDigestDigestItem(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_SpRuntimeGenericDigestDigestItem(in any) *pbgear.SpRuntimeGenericDigestDigestItem {
	out := &pbgear.SpRuntimeGenericDigestDigestItem{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Logs
	out.Logs = To_SpRuntimeGenericDigestDigestItem_logs(fields[0].Value)

	return out //from message
}

func To_SpRuntimeGenericDigestDigestItem_logs(in any) *pbgear.Babe_Logs {
	return nil //simple field
}

func To_SpRuntimeGenericHeaderHeader(in any) *pbgear.SpRuntimeGenericHeaderHeader {
	out := &pbgear.SpRuntimeGenericHeaderHeader{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ParentHash
	out.ParentHash = To_SpRuntimeGenericHeaderHeader_parent_hash(fields[0].Value)
	// primitive field Number
	out.Number = To_uint32(fields[1].Value)
	// field StateRoot
	out.StateRoot = To_SpRuntimeGenericHeaderHeader_state_root(fields[2].Value)
	// field ExtrinsicsRoot
	out.ExtrinsicsRoot = To_SpRuntimeGenericHeaderHeader_extrinsics_root(fields[3].Value)
	// field Digest
	out.Digest = To_SpRuntimeGenericHeaderHeader_digest(fields[4].Value)

	return out //from message
}

func To_SpRuntimeGenericHeaderHeader_parent_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}
func To_SpRuntimeGenericHeaderHeader_state_root(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}
func To_SpRuntimeGenericHeaderHeader_extrinsics_root(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}
func To_SpRuntimeGenericHeaderHeader_digest(in any) *pbgear.SpRuntimeGenericDigestDigest {
	return nil //simple field
}

func To_SpRuntimeMultiaddressMultiAddress(in any) *pbgear.SpRuntimeMultiaddressMultiAddress {
	out := &pbgear.SpRuntimeMultiaddressMultiAddress{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Targets
	out.Targets = To_SpRuntimeMultiaddressMultiAddress_targets(fields[0].Value)

	return out //from message
}

func To_SpRuntimeMultiaddressMultiAddress_targets(in any) *pbgear.Staking_Targets {
	return nil //simple field
}

func To_SpSessionMembershipProof(in any) *pbgear.SpSessionMembershipProof {
	out := &pbgear.SpSessionMembershipProof{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Session
	out.Session = To_uint32(fields[0].Value)
	// repeated field TrieNodes
	out.TrieNodes = To_Repeated_SpSessionMembershipProof_trie_nodes(fields[1].Value)
	// primitive field ValidatorCount
	out.ValidatorCount = To_uint32(fields[2].Value)

	return out //from message
}

func To_Repeated_SpSessionMembershipProof_trie_nodes(in any) []*pbgear.Babe_BabeTrieNodesList {
	items := in.([]interface{})

	var out []*pbgear.Babe_BabeTrieNodesList
	for _, item := range items {
		o := To_Babe_BabeTrieNodesList(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_SpWeightsWeightV2Weight(in any) *pbgear.SpWeightsWeightV2Weight {
	out := &pbgear.SpWeightsWeightV2Weight{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field RefTime
	out.RefTime = To_uint64(fields[0].Value)
	// primitive field ProofSize
	out.ProofSize = To_uint64(fields[1].Value)

	return out //from message
}

func To_StakingPallet(in any) *pbgear.StakingPallet {
	out := &pbgear.StakingPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_StakingPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_StakingPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.StakingPallet_BondCall{
			BondCall: To_Staking_BondCall(fields),
		}
	case 1:
		target = &pbgear.StakingPallet_BondExtraCall{
			BondExtraCall: To_Staking_BondExtraCall(fields),
		}
	case 2:
		target = &pbgear.StakingPallet_UnbondCall{
			UnbondCall: To_Staking_UnbondCall(fields),
		}
	case 3:
		target = &pbgear.StakingPallet_WithdrawUnbondedCall{
			WithdrawUnbondedCall: To_Staking_WithdrawUnbondedCall(fields),
		}
	case 4:
		target = &pbgear.StakingPallet_ValidateCall{
			ValidateCall: To_Staking_ValidateCall(fields),
		}
	case 5:
		target = &pbgear.StakingPallet_NominateCall{
			NominateCall: To_Staking_NominateCall(fields),
		}
	case 6:
		target = &pbgear.StakingPallet_ChillCall{
			ChillCall: To_Staking_ChillCall(fields),
		}
	case 7:
		target = &pbgear.StakingPallet_SetPayeeCall{
			SetPayeeCall: To_Staking_SetPayeeCall(fields),
		}
	case 8:
		target = &pbgear.StakingPallet_SetControllerCall{
			SetControllerCall: To_Staking_SetControllerCall(fields),
		}
	case 9:
		target = &pbgear.StakingPallet_SetValidatorCountCall{
			SetValidatorCountCall: To_Staking_SetValidatorCountCall(fields),
		}
	case 10:
		target = &pbgear.StakingPallet_IncreaseValidatorCountCall{
			IncreaseValidatorCountCall: To_Staking_IncreaseValidatorCountCall(fields),
		}
	case 11:
		target = &pbgear.StakingPallet_ScaleValidatorCountCall{
			ScaleValidatorCountCall: To_Staking_ScaleValidatorCountCall(fields),
		}
	case 12:
		target = &pbgear.StakingPallet_ForceNoErasCall{
			ForceNoErasCall: To_Staking_ForceNoErasCall(fields),
		}
	case 13:
		target = &pbgear.StakingPallet_ForceNewEraCall{
			ForceNewEraCall: To_Staking_ForceNewEraCall(fields),
		}
	case 14:
		target = &pbgear.StakingPallet_SetInvulnerablesCall{
			SetInvulnerablesCall: To_Staking_SetInvulnerablesCall(fields),
		}
	case 15:
		target = &pbgear.StakingPallet_ForceUnstakeCall{
			ForceUnstakeCall: To_Staking_ForceUnstakeCall(fields),
		}
	case 16:
		target = &pbgear.StakingPallet_ForceNewEraAlwaysCall{
			ForceNewEraAlwaysCall: To_Staking_ForceNewEraAlwaysCall(fields),
		}
	case 17:
		target = &pbgear.StakingPallet_CancelDeferredSlashCall{
			CancelDeferredSlashCall: To_Staking_CancelDeferredSlashCall(fields),
		}
	case 18:
		target = &pbgear.StakingPallet_PayoutStakersCall{
			PayoutStakersCall: To_Staking_PayoutStakersCall(fields),
		}
	case 19:
		target = &pbgear.StakingPallet_RebondCall{
			RebondCall: To_Staking_RebondCall(fields),
		}
	case 20:
		target = &pbgear.StakingPallet_ReapStashCall{
			ReapStashCall: To_Staking_ReapStashCall(fields),
		}
	case 21:
		target = &pbgear.StakingPallet_KickCall{
			KickCall: To_Staking_KickCall(fields),
		}
	case 22:
		target = &pbgear.StakingPallet_SetStakingConfigsCall{
			SetStakingConfigsCall: To_Staking_SetStakingConfigsCall(fields),
		}
	case 23:
		target = &pbgear.StakingPallet_ChillOtherCall{
			ChillOtherCall: To_Staking_ChillOtherCall(fields),
		}
	case 24:
		target = &pbgear.StakingPallet_ForceApplyMinCommissionCall{
			ForceApplyMinCommissionCall: To_Staking_ForceApplyMinCommissionCall(fields),
		}
	case 25:
		target = &pbgear.StakingPallet_SetMinCommissionCall{
			SetMinCommissionCall: To_Staking_SetMinCommissionCall(fields),
		}
	}
}

func To_StakingRewardsPallet(in any) *pbgear.StakingRewardsPallet {
	out := &pbgear.StakingRewardsPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_StakingRewardsPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_StakingRewardsPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.StakingRewardsPallet_RefillCall{
			RefillCall: To_StakingRewards_RefillCall(fields),
		}
	case 1:
		target = &pbgear.StakingRewardsPallet_ForceRefillCall{
			ForceRefillCall: To_StakingRewards_ForceRefillCall(fields),
		}
	case 2:
		target = &pbgear.StakingRewardsPallet_WithdrawCall{
			WithdrawCall: To_StakingRewards_WithdrawCall(fields),
		}
	case 3:
		target = &pbgear.StakingRewardsPallet_AlignSupplyCall{
			AlignSupplyCall: To_StakingRewards_AlignSupplyCall(fields),
		}
	}
}

func To_StakingRewards_Address20(in any) *pbgear.StakingRewards_Address20 {
	out := &pbgear.StakingRewards_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_StakingRewards_Address32(in any) *pbgear.StakingRewards_Address32 {
	out := &pbgear.StakingRewards_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_StakingRewards_AlignSupplyCall(in any) *pbgear.StakingRewards_AlignSupplyCall {
	out := &pbgear.StakingRewards_AlignSupplyCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Target
	out.Target = To_string(fields[0].Value)

	return out //from message
}

func To_StakingRewards_ForceRefillCall(in any) *pbgear.StakingRewards_ForceRefillCall {
	out := &pbgear.StakingRewards_ForceRefillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field From
	out.From = To_StakingRewards_ForceRefillCall_from(fields[0].Value)
	// primitive field Value
	out.Value = To_string(fields[1].Value)

	return out //from message
}

func To_StakingRewards_ForceRefillCall_from(in any) *pbgear.StakingRewards_From {
	return nil //simple field
}

func To_StakingRewards_From(in any) *pbgear.StakingRewards_From {
	out := &pbgear.StakingRewards_From{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_StakingRewards_From_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_StakingRewards_From_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.StakingRewards_From_Id{
			Id: To_StakingRewards_Id(fields),
		}
	case 1:
		target = &pbgear.StakingRewards_From_Index{
			Index: To_StakingRewards_Index(fields),
		}
	case 2:
		target = &pbgear.StakingRewards_From_Raw{
			Raw: To_StakingRewards_Raw(fields),
		}
	case 3:
		target = &pbgear.StakingRewards_From_Address32{
			Address32: To_StakingRewards_Address32(fields),
		}
	case 4:
		target = &pbgear.StakingRewards_From_Address20{
			Address20: To_StakingRewards_Address20(fields),
		}
	}
}

func To_StakingRewards_Id(in any) *pbgear.StakingRewards_Id {
	out := &pbgear.StakingRewards_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_StakingRewards_Id_value0(fields[0].Value)

	return out //from message
}

func To_StakingRewards_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_StakingRewards_Index(in any) *pbgear.StakingRewards_Index {
	out := &pbgear.StakingRewards_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_StakingRewards_Index_value0(fields[0].Value)

	return out //from message
}

func To_StakingRewards_Index_value0(in any) *pbgear.StakingRewards_TupleNull {
	return nil //simple field
}

func To_StakingRewards_Raw(in any) *pbgear.StakingRewards_Raw {
	out := &pbgear.StakingRewards_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_StakingRewards_RefillCall(in any) *pbgear.StakingRewards_RefillCall {
	out := &pbgear.StakingRewards_RefillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)

	return out //from message
}

func To_StakingRewards_To(in any) *pbgear.StakingRewards_To {
	out := &pbgear.StakingRewards_To{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_StakingRewards_To_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_StakingRewards_To_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.StakingRewards_To_Id{
			Id: To_StakingRewards_Id(fields),
		}
	case 1:
		target = &pbgear.StakingRewards_To_Index{
			Index: To_StakingRewards_Index(fields),
		}
	case 2:
		target = &pbgear.StakingRewards_To_Raw{
			Raw: To_StakingRewards_Raw(fields),
		}
	case 3:
		target = &pbgear.StakingRewards_To_Address32{
			Address32: To_StakingRewards_Address32(fields),
		}
	case 4:
		target = &pbgear.StakingRewards_To_Address20{
			Address20: To_StakingRewards_Address20(fields),
		}
	}
}

func To_StakingRewards_TupleNull(in any) *pbgear.StakingRewards_TupleNull {
	out := &pbgear.StakingRewards_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_StakingRewards_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_StakingRewards_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_StakingRewards_WithdrawCall(in any) *pbgear.StakingRewards_WithdrawCall {
	out := &pbgear.StakingRewards_WithdrawCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field To
	out.To = To_StakingRewards_WithdrawCall_to(fields[0].Value)
	// primitive field Value
	out.Value = To_string(fields[1].Value)

	return out //from message
}

func To_StakingRewards_WithdrawCall_to(in any) *pbgear.StakingRewards_To {
	return nil //simple field
}

func To_Staking_Account(in any) *pbgear.Staking_Account {
	out := &pbgear.Staking_Account{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Staking_Account_value0(fields[0].Value)

	return out //from message
}

func To_Staking_Account_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_Address20(in any) *pbgear.Staking_Address20 {
	out := &pbgear.Staking_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Staking_Address32(in any) *pbgear.Staking_Address32 {
	out := &pbgear.Staking_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Staking_BondCall(in any) *pbgear.Staking_BondCall {
	out := &pbgear.Staking_BondCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)
	// field Payee
	out.Payee = To_Staking_BondCall_payee(fields[1].Value)

	return out //from message
}

func To_Staking_BondCall_payee(in any) *pbgear.Staking_Payee {
	return nil //simple field
}

func To_Staking_BondExtraCall(in any) *pbgear.Staking_BondExtraCall {
	out := &pbgear.Staking_BondExtraCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field MaxAdditional
	out.MaxAdditional = To_string(fields[0].Value)

	return out //from message
}

func To_Staking_CancelDeferredSlashCall(in any) *pbgear.Staking_CancelDeferredSlashCall {
	out := &pbgear.Staking_CancelDeferredSlashCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Era
	out.Era = To_uint32(fields[0].Value)
	// primitive field SlashIndices
	out.SlashIndices = To_Repeated_uint32(fields[1].Value)

	return out //from message
}

func To_Staking_ChillCall(in any) *pbgear.Staking_ChillCall {
	out := &pbgear.Staking_ChillCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ChillOtherCall(in any) *pbgear.Staking_ChillOtherCall {
	out := &pbgear.Staking_ChillOtherCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Controller
	out.Controller = To_Staking_ChillOtherCall_controller(fields[0].Value)

	return out //from message
}

func To_Staking_ChillOtherCall_controller(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_ChillThreshold(in any) *pbgear.Staking_ChillThreshold {
	out := &pbgear.Staking_ChillThreshold{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_ChillThreshold_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_ChillThreshold_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_ChillThreshold_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_ChillThreshold_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_ChillThreshold_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_Controller(in any) *pbgear.Staking_Controller {
	out := &pbgear.Staking_Controller{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ForceApplyMinCommissionCall(in any) *pbgear.Staking_ForceApplyMinCommissionCall {
	out := &pbgear.Staking_ForceApplyMinCommissionCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ValidatorStash
	out.ValidatorStash = To_Staking_ForceApplyMinCommissionCall_validator_stash(fields[0].Value)

	return out //from message
}

func To_Staking_ForceApplyMinCommissionCall_validator_stash(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_ForceNewEraAlwaysCall(in any) *pbgear.Staking_ForceNewEraAlwaysCall {
	out := &pbgear.Staking_ForceNewEraAlwaysCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ForceNewEraCall(in any) *pbgear.Staking_ForceNewEraCall {
	out := &pbgear.Staking_ForceNewEraCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ForceNoErasCall(in any) *pbgear.Staking_ForceNoErasCall {
	out := &pbgear.Staking_ForceNoErasCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ForceUnstakeCall(in any) *pbgear.Staking_ForceUnstakeCall {
	out := &pbgear.Staking_ForceUnstakeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Stash
	out.Stash = To_Staking_ForceUnstakeCall_stash(fields[0].Value)
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(fields[1].Value)

	return out //from message
}

func To_Staking_ForceUnstakeCall_stash(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_Id(in any) *pbgear.Staking_Id {
	out := &pbgear.Staking_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Staking_Id_value0(fields[0].Value)

	return out //from message
}

func To_Staking_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_IncreaseValidatorCountCall(in any) *pbgear.Staking_IncreaseValidatorCountCall {
	out := &pbgear.Staking_IncreaseValidatorCountCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Additional
	out.Additional = To_uint32(fields[0].Value)

	return out //from message
}

func To_Staking_Index(in any) *pbgear.Staking_Index {
	out := &pbgear.Staking_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Staking_Index_value0(fields[0].Value)

	return out //from message
}

func To_Staking_Index_value0(in any) *pbgear.Staking_TupleNull {
	return nil //simple field
}

func To_Staking_KickCall(in any) *pbgear.Staking_KickCall {
	out := &pbgear.Staking_KickCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Who
	out.Who = To_Repeated_Staking_KickCall_who(fields[0].Value)

	return out //from message
}

func To_Repeated_Staking_KickCall_who(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeMultiaddressMultiAddress
	for _, item := range items {
		o := To_SpRuntimeMultiaddressMultiAddress(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Staking_MaxNominatorCount(in any) *pbgear.Staking_MaxNominatorCount {
	out := &pbgear.Staking_MaxNominatorCount{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_MaxNominatorCount_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_MaxNominatorCount_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_MaxNominatorCount_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_MaxNominatorCount_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_MaxNominatorCount_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_MaxValidatorCount(in any) *pbgear.Staking_MaxValidatorCount {
	out := &pbgear.Staking_MaxValidatorCount{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_MaxValidatorCount_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_MaxValidatorCount_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_MaxValidatorCount_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_MaxValidatorCount_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_MaxValidatorCount_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_MinCommission(in any) *pbgear.Staking_MinCommission {
	out := &pbgear.Staking_MinCommission{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_MinCommission_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_MinCommission_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_MinCommission_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_MinCommission_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_MinCommission_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_MinNominatorBond(in any) *pbgear.Staking_MinNominatorBond {
	out := &pbgear.Staking_MinNominatorBond{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_MinNominatorBond_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_MinNominatorBond_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_MinNominatorBond_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_MinNominatorBond_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_MinNominatorBond_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_MinValidatorBond(in any) *pbgear.Staking_MinValidatorBond {
	out := &pbgear.Staking_MinValidatorBond{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_MinValidatorBond_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_MinValidatorBond_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_MinValidatorBond_Noop{
			Noop: To_Staking_Noop(fields),
		}
	case 1:
		target = &pbgear.Staking_MinValidatorBond_Set{
			Set: To_Staking_Set(fields),
		}
	case 2:
		target = &pbgear.Staking_MinValidatorBond_Remove{
			Remove: To_Staking_Remove(fields),
		}
	}
}

func To_Staking_NominateCall(in any) *pbgear.Staking_NominateCall {
	out := &pbgear.Staking_NominateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Targets
	out.Targets = To_Repeated_Staking_NominateCall_targets(fields[0].Value)

	return out //from message
}

func To_Repeated_Staking_NominateCall_targets(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeMultiaddressMultiAddress
	for _, item := range items {
		o := To_SpRuntimeMultiaddressMultiAddress(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Staking_None(in any) *pbgear.Staking_None {
	out := &pbgear.Staking_None{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_Noop(in any) *pbgear.Staking_Noop {
	out := &pbgear.Staking_Noop{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_PalletStakingValidatorPrefs(in any) *pbgear.Staking_PalletStakingValidatorPrefs {
	out := &pbgear.Staking_PalletStakingValidatorPrefs{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Commission
	out.Commission = To_Staking_PalletStakingValidatorPrefs_commission(fields[0].Value)
	// primitive field Blocked
	out.Blocked = To_bool(fields[1].Value)

	return out //from message
}

func To_Staking_PalletStakingValidatorPrefs_commission(in any) *pbgear.SpArithmeticPerThingsPerbill {
	return nil //simple field
}

func To_Staking_Payee(in any) *pbgear.Staking_Payee {
	out := &pbgear.Staking_Payee{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_Payee_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_Payee_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_Payee_Staked{
			Staked: To_Staking_Staked(fields),
		}
	case 1:
		target = &pbgear.Staking_Payee_Stash{
			Stash: To_Staking_Stash(fields),
		}
	case 2:
		target = &pbgear.Staking_Payee_Controller{
			Controller: To_Staking_Controller(fields),
		}
	case 3:
		target = &pbgear.Staking_Payee_Account{
			Account: To_Staking_Account(fields),
		}
	case 4:
		target = &pbgear.Staking_Payee_None{
			None: To_Staking_None(fields),
		}
	}
}

func To_Staking_PayoutStakersCall(in any) *pbgear.Staking_PayoutStakersCall {
	out := &pbgear.Staking_PayoutStakersCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field ValidatorStash
	out.ValidatorStash = To_Staking_PayoutStakersCall_validator_stash(fields[0].Value)
	// primitive field Era
	out.Era = To_uint32(fields[1].Value)

	return out //from message
}

func To_Staking_PayoutStakersCall_validator_stash(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_Raw(in any) *pbgear.Staking_Raw {
	out := &pbgear.Staking_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Staking_ReapStashCall(in any) *pbgear.Staking_ReapStashCall {
	out := &pbgear.Staking_ReapStashCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Stash
	out.Stash = To_Staking_ReapStashCall_stash(fields[0].Value)
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(fields[1].Value)

	return out //from message
}

func To_Staking_ReapStashCall_stash(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Staking_RebondCall(in any) *pbgear.Staking_RebondCall {
	out := &pbgear.Staking_RebondCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)

	return out //from message
}

func To_Staking_Remove(in any) *pbgear.Staking_Remove {
	out := &pbgear.Staking_Remove{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_ScaleValidatorCountCall(in any) *pbgear.Staking_ScaleValidatorCountCall {
	out := &pbgear.Staking_ScaleValidatorCountCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Factor
	out.Factor = To_Staking_ScaleValidatorCountCall_factor(fields[0].Value)

	return out //from message
}

func To_Staking_ScaleValidatorCountCall_factor(in any) *pbgear.SpArithmeticPerThingsPercent {
	return nil //simple field
}

func To_Staking_Set(in any) *pbgear.Staking_Set {
	out := &pbgear.Staking_Set{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_string(fields[0].Value)

	return out //from message
}

func To_Staking_SetControllerCall(in any) *pbgear.Staking_SetControllerCall {
	out := &pbgear.Staking_SetControllerCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_SetInvulnerablesCall(in any) *pbgear.Staking_SetInvulnerablesCall {
	out := &pbgear.Staking_SetInvulnerablesCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Invulnerables
	out.Invulnerables = To_Repeated_Staking_SetInvulnerablesCall_invulnerables(fields[0].Value)

	return out //from message
}

func To_Repeated_Staking_SetInvulnerablesCall_invulnerables(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		o := To_SpCoreCryptoAccountId32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Staking_SetMinCommissionCall(in any) *pbgear.Staking_SetMinCommissionCall {
	out := &pbgear.Staking_SetMinCommissionCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field New
	out.New = To_Staking_SetMinCommissionCall_new(fields[0].Value)

	return out //from message
}

func To_Staking_SetMinCommissionCall_new(in any) *pbgear.SpArithmeticPerThingsPerbill {
	return nil //simple field
}

func To_Staking_SetPayeeCall(in any) *pbgear.Staking_SetPayeeCall {
	out := &pbgear.Staking_SetPayeeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Payee
	out.Payee = To_Staking_SetPayeeCall_payee(fields[0].Value)

	return out //from message
}

func To_Staking_SetPayeeCall_payee(in any) *pbgear.Staking_Payee {
	return nil //simple field
}

func To_Staking_SetStakingConfigsCall(in any) *pbgear.Staking_SetStakingConfigsCall {
	out := &pbgear.Staking_SetStakingConfigsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field MinNominatorBond
	out.MinNominatorBond = To_Staking_SetStakingConfigsCall_min_nominator_bond(fields[0].Value)
	// field MinValidatorBond
	out.MinValidatorBond = To_Staking_SetStakingConfigsCall_min_validator_bond(fields[1].Value)
	// field MaxNominatorCount
	out.MaxNominatorCount = To_Staking_SetStakingConfigsCall_max_nominator_count(fields[2].Value)
	// field MaxValidatorCount
	out.MaxValidatorCount = To_Staking_SetStakingConfigsCall_max_validator_count(fields[3].Value)
	// field ChillThreshold
	out.ChillThreshold = To_Staking_SetStakingConfigsCall_chill_threshold(fields[4].Value)
	// field MinCommission
	out.MinCommission = To_Staking_SetStakingConfigsCall_min_commission(fields[5].Value)

	return out //from message
}

func To_Staking_SetStakingConfigsCall_min_nominator_bond(in any) *pbgear.Staking_MinNominatorBond {
	return nil //simple field
}
func To_Staking_SetStakingConfigsCall_min_validator_bond(in any) *pbgear.Staking_MinValidatorBond {
	return nil //simple field
}
func To_Staking_SetStakingConfigsCall_max_nominator_count(in any) *pbgear.Staking_MaxNominatorCount {
	return nil //simple field
}
func To_Staking_SetStakingConfigsCall_max_validator_count(in any) *pbgear.Staking_MaxValidatorCount {
	return nil //simple field
}
func To_Staking_SetStakingConfigsCall_chill_threshold(in any) *pbgear.Staking_ChillThreshold {
	return nil //simple field
}
func To_Staking_SetStakingConfigsCall_min_commission(in any) *pbgear.Staking_MinCommission {
	return nil //simple field
}

func To_Staking_SetValidatorCountCall(in any) *pbgear.Staking_SetValidatorCountCall {
	out := &pbgear.Staking_SetValidatorCountCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field New
	out.New = To_uint32(fields[0].Value)

	return out //from message
}

func To_Staking_Staked(in any) *pbgear.Staking_Staked {
	out := &pbgear.Staking_Staked{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_Stash(in any) *pbgear.Staking_Stash {
	out := &pbgear.Staking_Stash{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Staking_Targets(in any) *pbgear.Staking_Targets {
	out := &pbgear.Staking_Targets{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_Targets_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_Targets_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_Targets_Id{
			Id: To_Staking_Id(fields),
		}
	case 1:
		target = &pbgear.Staking_Targets_Index{
			Index: To_Staking_Index(fields),
		}
	case 2:
		target = &pbgear.Staking_Targets_Raw{
			Raw: To_Staking_Raw(fields),
		}
	case 3:
		target = &pbgear.Staking_Targets_Address32{
			Address32: To_Staking_Address32(fields),
		}
	case 4:
		target = &pbgear.Staking_Targets_Address20{
			Address20: To_Staking_Address20(fields),
		}
	}
}

func To_Staking_TupleNull(in any) *pbgear.Staking_TupleNull {
	out := &pbgear.Staking_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Staking_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Staking_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Staking_UnbondCall(in any) *pbgear.Staking_UnbondCall {
	out := &pbgear.Staking_UnbondCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)

	return out //from message
}

func To_Staking_ValidateCall(in any) *pbgear.Staking_ValidateCall {
	out := &pbgear.Staking_ValidateCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Prefs
	out.Prefs = To_Staking_ValidateCall_prefs(fields[0].Value)

	return out //from message
}

func To_Staking_ValidateCall_prefs(in any) *pbgear.Staking_PalletStakingValidatorPrefs {
	return nil //simple field
}

func To_Staking_Who(in any) *pbgear.Staking_Who {
	out := &pbgear.Staking_Who{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Staking_Who_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Staking_Who_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Staking_Who_Id{
			Id: To_Staking_Id(fields),
		}
	case 1:
		target = &pbgear.Staking_Who_Index{
			Index: To_Staking_Index(fields),
		}
	case 2:
		target = &pbgear.Staking_Who_Raw{
			Raw: To_Staking_Raw(fields),
		}
	case 3:
		target = &pbgear.Staking_Who_Address32{
			Address32: To_Staking_Address32(fields),
		}
	case 4:
		target = &pbgear.Staking_Who_Address20{
			Address20: To_Staking_Address20(fields),
		}
	}
}

func To_Staking_WithdrawUnbondedCall(in any) *pbgear.Staking_WithdrawUnbondedCall {
	out := &pbgear.Staking_WithdrawUnbondedCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(fields[0].Value)

	return out //from message
}

func To_SystemPallet(in any) *pbgear.SystemPallet {
	out := &pbgear.SystemPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_SystemPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_SystemPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.SystemPallet_RemarkCall{
			RemarkCall: To_System_RemarkCall(fields),
		}
	case 1:
		target = &pbgear.SystemPallet_SetHeapPagesCall{
			SetHeapPagesCall: To_System_SetHeapPagesCall(fields),
		}
	case 2:
		target = &pbgear.SystemPallet_SetCodeCall{
			SetCodeCall: To_System_SetCodeCall(fields),
		}
	case 3:
		target = &pbgear.SystemPallet_SetCodeWithoutChecksCall{
			SetCodeWithoutChecksCall: To_System_SetCodeWithoutChecksCall(fields),
		}
	case 4:
		target = &pbgear.SystemPallet_SetStorageCall{
			SetStorageCall: To_System_SetStorageCall(fields),
		}
	case 5:
		target = &pbgear.SystemPallet_KillStorageCall{
			KillStorageCall: To_System_KillStorageCall(fields),
		}
	case 6:
		target = &pbgear.SystemPallet_KillPrefixCall{
			KillPrefixCall: To_System_KillPrefixCall(fields),
		}
	case 7:
		target = &pbgear.SystemPallet_RemarkWithEventCall{
			RemarkWithEventCall: To_System_RemarkWithEventCall(fields),
		}
	}
}

func To_System_KillPrefixCall(in any) *pbgear.System_KillPrefixCall {
	out := &pbgear.System_KillPrefixCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Prefix
	out.Prefix = To_bytes(fields[0].Value)
	// primitive field Subkeys
	out.Subkeys = To_uint32(fields[1].Value)

	return out //from message
}

func To_System_KillStorageCall(in any) *pbgear.System_KillStorageCall {
	out := &pbgear.System_KillStorageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Keys
	out.Keys = To_Repeated_System_KillStorageCall_keys(fields[0].Value)

	return out //from message
}

func To_Repeated_System_KillStorageCall_keys(in any) []*pbgear.System_SystemKeysList {
	items := in.([]interface{})

	var out []*pbgear.System_SystemKeysList
	for _, item := range items {
		o := To_System_SystemKeysList(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_System_RemarkCall(in any) *pbgear.System_RemarkCall {
	out := &pbgear.System_RemarkCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Remark
	out.Remark = To_bytes(fields[0].Value)

	return out //from message
}

func To_System_RemarkWithEventCall(in any) *pbgear.System_RemarkWithEventCall {
	out := &pbgear.System_RemarkWithEventCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Remark
	out.Remark = To_bytes(fields[0].Value)

	return out //from message
}

func To_System_SetCodeCall(in any) *pbgear.System_SetCodeCall {
	out := &pbgear.System_SetCodeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Code
	out.Code = To_bytes(fields[0].Value)

	return out //from message
}

func To_System_SetCodeWithoutChecksCall(in any) *pbgear.System_SetCodeWithoutChecksCall {
	out := &pbgear.System_SetCodeWithoutChecksCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Code
	out.Code = To_bytes(fields[0].Value)

	return out //from message
}

func To_System_SetHeapPagesCall(in any) *pbgear.System_SetHeapPagesCall {
	out := &pbgear.System_SetHeapPagesCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Pages
	out.Pages = To_uint64(fields[0].Value)

	return out //from message
}

func To_System_SetStorageCall(in any) *pbgear.System_SetStorageCall {
	out := &pbgear.System_SetStorageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Items
	out.Items = To_Repeated_System_SetStorageCall_items(fields[0].Value)

	return out //from message
}

func To_Repeated_System_SetStorageCall_items(in any) []*pbgear.System_TupleSystemItemsListSystemItemsList {
	items := in.([]interface{})

	var out []*pbgear.System_TupleSystemItemsListSystemItemsList
	for _, item := range items {
		o := To_System_TupleSystemItemsListSystemItemsList(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_System_SystemKeysList(in any) *pbgear.System_SystemKeysList {
	out := &pbgear.System_SystemKeysList{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Keys
	out.Keys = To_bytes(fields[0].Value)

	return out //from message
}

func To_System_TupleSystemItemsListSystemItemsList(in any) *pbgear.System_TupleSystemItemsListSystemItemsList {
	out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)
	// primitive field Value1
	out.Value1 = To_bytes(fields[1].Value)

	return out //from message
}

func To_TimestampPallet(in any) *pbgear.TimestampPallet {
	out := &pbgear.TimestampPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_TimestampPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_TimestampPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.TimestampPallet_SetCall{
			SetCall: To_Timestamp_SetCall(fields),
		}
	}
}

func To_Timestamp_SetCall(in any) *pbgear.Timestamp_SetCall {
	out := &pbgear.Timestamp_SetCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Now
	out.Now = To_uint64(fields[0].Value)

	return out //from message
}

func To_TransactionPaymentPallet(in any) *pbgear.TransactionPaymentPallet {
	out := &pbgear.TransactionPaymentPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_TreasuryPallet(in any) *pbgear.TreasuryPallet {
	out := &pbgear.TreasuryPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_TreasuryPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_TreasuryPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.TreasuryPallet_ProposeSpendCall{
			ProposeSpendCall: To_Treasury_ProposeSpendCall(fields),
		}
	case 1:
		target = &pbgear.TreasuryPallet_RejectProposalCall{
			RejectProposalCall: To_Treasury_RejectProposalCall(fields),
		}
	case 2:
		target = &pbgear.TreasuryPallet_ApproveProposalCall{
			ApproveProposalCall: To_Treasury_ApproveProposalCall(fields),
		}
	case 3:
		target = &pbgear.TreasuryPallet_SpendLocalCall{
			SpendLocalCall: To_Treasury_SpendLocalCall(fields),
		}
	case 4:
		target = &pbgear.TreasuryPallet_RemoveApprovalCall{
			RemoveApprovalCall: To_Treasury_RemoveApprovalCall(fields),
		}
	case 5:
		target = &pbgear.TreasuryPallet_SpendCall{
			SpendCall: To_Treasury_SpendCall(fields),
		}
	case 6:
		target = &pbgear.TreasuryPallet_PayoutCall{
			PayoutCall: To_Treasury_PayoutCall(fields),
		}
	case 7:
		target = &pbgear.TreasuryPallet_CheckStatusCall{
			CheckStatusCall: To_Treasury_CheckStatusCall(fields),
		}
	case 8:
		target = &pbgear.TreasuryPallet_VoidSpendCall{
			VoidSpendCall: To_Treasury_VoidSpendCall(fields),
		}
	}
}

func To_Treasury_Address20(in any) *pbgear.Treasury_Address20 {
	out := &pbgear.Treasury_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Treasury_Address32(in any) *pbgear.Treasury_Address32 {
	out := &pbgear.Treasury_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Treasury_ApproveProposalCall(in any) *pbgear.Treasury_ApproveProposalCall {
	out := &pbgear.Treasury_ApproveProposalCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ProposalId
	out.ProposalId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Treasury_Beneficiary(in any) *pbgear.Treasury_Beneficiary {
	out := &pbgear.Treasury_Beneficiary{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Treasury_Beneficiary_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Treasury_Beneficiary_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Treasury_Beneficiary_Id{
			Id: To_Treasury_Id(fields),
		}
	case 1:
		target = &pbgear.Treasury_Beneficiary_Index{
			Index: To_Treasury_Index(fields),
		}
	case 2:
		target = &pbgear.Treasury_Beneficiary_Raw{
			Raw: To_Treasury_Raw(fields),
		}
	case 3:
		target = &pbgear.Treasury_Beneficiary_Address32{
			Address32: To_Treasury_Address32(fields),
		}
	case 4:
		target = &pbgear.Treasury_Beneficiary_Address20{
			Address20: To_Treasury_Address20(fields),
		}
	}
}

func To_Treasury_CheckStatusCall(in any) *pbgear.Treasury_CheckStatusCall {
	out := &pbgear.Treasury_CheckStatusCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Treasury_Id(in any) *pbgear.Treasury_Id {
	out := &pbgear.Treasury_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Treasury_Id_value0(fields[0].Value)

	return out //from message
}

func To_Treasury_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Treasury_Index(in any) *pbgear.Treasury_Index {
	out := &pbgear.Treasury_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Treasury_Index_value0(fields[0].Value)

	return out //from message
}

func To_Treasury_Index_value0(in any) *pbgear.Treasury_TupleNull {
	return nil //simple field
}

func To_Treasury_PayoutCall(in any) *pbgear.Treasury_PayoutCall {
	out := &pbgear.Treasury_PayoutCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_Treasury_ProposeSpendCall(in any) *pbgear.Treasury_ProposeSpendCall {
	out := &pbgear.Treasury_ProposeSpendCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value
	out.Value = To_string(fields[0].Value)
	// field Beneficiary
	out.Beneficiary = To_Treasury_ProposeSpendCall_beneficiary(fields[1].Value)

	return out //from message
}

func To_Treasury_ProposeSpendCall_beneficiary(in any) *pbgear.Treasury_Beneficiary {
	return nil //simple field
}

func To_Treasury_Raw(in any) *pbgear.Treasury_Raw {
	out := &pbgear.Treasury_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Treasury_RejectProposalCall(in any) *pbgear.Treasury_RejectProposalCall {
	out := &pbgear.Treasury_RejectProposalCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ProposalId
	out.ProposalId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Treasury_RemoveApprovalCall(in any) *pbgear.Treasury_RemoveApprovalCall {
	out := &pbgear.Treasury_RemoveApprovalCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field ProposalId
	out.ProposalId = To_uint32(fields[0].Value)

	return out //from message
}

func To_Treasury_SpendCall(in any) *pbgear.Treasury_SpendCall {
	out := &pbgear.Treasury_SpendCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field AssetKind
	out.AssetKind = To_Treasury_SpendCall_asset_kind(fields[0].Value)
	// primitive field Amount
	out.Amount = To_string(fields[1].Value)
	// field Beneficiary
	out.Beneficiary = To_Treasury_SpendCall_beneficiary(fields[2].Value)
	// primitive field ValidFrom
	out.ValidFrom = To_Optional_uint32(fields[3].Value)

	return out //from message
}

func To_Treasury_SpendCall_asset_kind(in any) *pbgear.Treasury_TupleNull {
	return nil //simple field
}
func To_Treasury_SpendCall_beneficiary(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Treasury_SpendLocalCall(in any) *pbgear.Treasury_SpendLocalCall {
	out := &pbgear.Treasury_SpendLocalCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Amount
	out.Amount = To_string(fields[0].Value)
	// field Beneficiary
	out.Beneficiary = To_Treasury_SpendLocalCall_beneficiary(fields[1].Value)

	return out //from message
}

func To_Treasury_SpendLocalCall_beneficiary(in any) *pbgear.Treasury_Beneficiary {
	return nil //simple field
}

func To_Treasury_TupleNull(in any) *pbgear.Treasury_TupleNull {
	out := &pbgear.Treasury_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Treasury_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Treasury_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Treasury_VoidSpendCall(in any) *pbgear.Treasury_VoidSpendCall {
	out := &pbgear.Treasury_VoidSpendCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)

	return out //from message
}

func To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature(in any) *pbgear.TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature {
	out := &pbgear.TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature_value1(fields[1].Value)

	return out //from message
}

func To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature_value0(in any) *pbgear.FinalityGrandpaPrecommit {
	return nil //simple field
}
func To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature_value1(in any) *pbgear.SpConsensusGrandpaAppSignature {
	return nil //simple field
}

func To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(in any) *pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	out := &pbgear.TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature_value0(fields[0].Value)
	// field Value1
	out.Value1 = To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature_value1(fields[1].Value)

	return out //from message
}

func To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature_value0(in any) *pbgear.FinalityGrandpaPrevote {
	return nil //simple field
}
func To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature_value1(in any) *pbgear.SpConsensusGrandpaAppSignature {
	return nil //simple field
}

func To_TupleNull(in any) *pbgear.TupleNull {
	out := &pbgear.TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_UtilityPallet(in any) *pbgear.UtilityPallet {
	out := &pbgear.UtilityPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_UtilityPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_UtilityPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.UtilityPallet_BatchCall{
			BatchCall: To_Utility_BatchCall(fields),
		}
	case 1:
		target = &pbgear.UtilityPallet_AsDerivativeCall{
			AsDerivativeCall: To_Utility_AsDerivativeCall(fields),
		}
	case 2:
		target = &pbgear.UtilityPallet_BatchAllCall{
			BatchAllCall: To_Utility_BatchAllCall(fields),
		}
	case 3:
		target = &pbgear.UtilityPallet_DispatchAsCall{
			DispatchAsCall: To_Utility_DispatchAsCall(fields),
		}
	case 4:
		target = &pbgear.UtilityPallet_ForceBatchCall{
			ForceBatchCall: To_Utility_ForceBatchCall(fields),
		}
	case 5:
		target = &pbgear.UtilityPallet_WithWeightCall{
			WithWeightCall: To_Utility_WithWeightCall(fields),
		}
	}
}

func To_Utility_AsDerivativeCall(in any) *pbgear.Utility_AsDerivativeCall {
	out := &pbgear.Utility_AsDerivativeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Index
	out.Index = To_uint32(fields[0].Value)
	// oneOf field Call
	Set_OneOf_Utility_AsDerivativeCall_call(out.Call, fields[1].Value)

	return out //from message
}

func Set_OneOf_Utility_AsDerivativeCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Utility_AsDerivativeCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Utility_AsDerivativeCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Utility_AsDerivativeCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Utility_AsDerivativeCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Utility_AsDerivativeCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Utility_AsDerivativeCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Utility_AsDerivativeCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Utility_AsDerivativeCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Utility_AsDerivativeCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Utility_AsDerivativeCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Utility_AsDerivativeCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Utility_AsDerivativeCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Utility_AsDerivativeCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Utility_AsDerivativeCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Utility_AsDerivativeCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Utility_AsDerivativeCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Utility_AsDerivativeCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Utility_AsDerivativeCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Utility_AsDerivativeCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Utility_AsDerivativeCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Utility_AsDerivativeCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Utility_AsDerivativeCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Utility_AsDerivativeCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Utility_AsDerivativeCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Utility_AsDerivativeCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Utility_AsDerivativeCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Utility_AsDerivativeCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Utility_AsDerivativeCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Utility_AsDerivativeCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Utility_AsOrigin(in any) *pbgear.Utility_AsOrigin {
	out := &pbgear.Utility_AsOrigin{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Utility_AsOrigin_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Utility_AsOrigin_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Utility_AsOrigin_System{
			System: To_Utility_System(fields),
		}
	case 20:
		target = &pbgear.Utility_AsOrigin_Origins{
			Origins: To_Utility_Origins(fields),
		}
	case 2:
		target = &pbgear.Utility_AsOrigin_Void{
			Void: To_Utility_Void(fields),
		}
	}
}

func To_Utility_BatchAllCall(in any) *pbgear.Utility_BatchAllCall {
	out := &pbgear.Utility_BatchAllCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Calls
	out.Calls = To_Repeated_Utility_BatchAllCall_calls(fields[0].Value)

	return out //from message
}

func To_Repeated_Utility_BatchAllCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		o := To_VaraRuntimeRuntimeCall(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Utility_BatchCall(in any) *pbgear.Utility_BatchCall {
	out := &pbgear.Utility_BatchCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Calls
	out.Calls = To_Repeated_Utility_BatchCall_calls(fields[0].Value)

	return out //from message
}

func To_Repeated_Utility_BatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		o := To_VaraRuntimeRuntimeCall(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Utility_DispatchAsCall(in any) *pbgear.Utility_DispatchAsCall {
	out := &pbgear.Utility_DispatchAsCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field AsOrigin
	out.AsOrigin = To_Utility_DispatchAsCall_as_origin(fields[0].Value)
	// oneOf field Call
	Set_OneOf_Utility_DispatchAsCall_call(out.Call, fields[1].Value)

	return out //from message
}

func To_Utility_DispatchAsCall_as_origin(in any) *pbgear.Utility_AsOrigin {
	return nil //simple field
}
func Set_OneOf_Utility_DispatchAsCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Utility_DispatchAsCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Utility_DispatchAsCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Utility_DispatchAsCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Utility_DispatchAsCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Utility_DispatchAsCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Utility_DispatchAsCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Utility_DispatchAsCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Utility_DispatchAsCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Utility_DispatchAsCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Utility_DispatchAsCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Utility_DispatchAsCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Utility_DispatchAsCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Utility_DispatchAsCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Utility_DispatchAsCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Utility_DispatchAsCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Utility_DispatchAsCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Utility_DispatchAsCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Utility_DispatchAsCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Utility_DispatchAsCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Utility_DispatchAsCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Utility_DispatchAsCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Utility_DispatchAsCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Utility_DispatchAsCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Utility_DispatchAsCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Utility_DispatchAsCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Utility_DispatchAsCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Utility_DispatchAsCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Utility_DispatchAsCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Utility_DispatchAsCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Utility_ForceBatchCall(in any) *pbgear.Utility_ForceBatchCall {
	out := &pbgear.Utility_ForceBatchCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Calls
	out.Calls = To_Repeated_Utility_ForceBatchCall_calls(fields[0].Value)

	return out //from message
}

func To_Repeated_Utility_ForceBatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		o := To_VaraRuntimeRuntimeCall(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Utility_Origins(in any) *pbgear.Utility_Origins {
	out := &pbgear.Utility_Origins{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Utility_Origins_value0(fields[0].Value)

	return out //from message
}

func To_Utility_Origins_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Utility_System(in any) *pbgear.Utility_System {
	out := &pbgear.Utility_System{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Utility_System_value0(fields[0].Value)

	return out //from message
}

func To_Utility_System_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Utility_Void(in any) *pbgear.Utility_Void {
	out := &pbgear.Utility_Void{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Utility_Void_value0(fields[0].Value)

	return out //from message
}

func To_Utility_Void_value0(in any) *pbgear.Value0 {
	return nil //simple field
}

func To_Utility_WithWeightCall(in any) *pbgear.Utility_WithWeightCall {
	out := &pbgear.Utility_WithWeightCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_Utility_WithWeightCall_call(out.Call, fields[0].Value)
	// field Weight
	out.Weight = To_Utility_WithWeightCall_weight(fields[1].Value)

	return out //from message
}

func Set_OneOf_Utility_WithWeightCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Utility_WithWeightCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Utility_WithWeightCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Utility_WithWeightCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Utility_WithWeightCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Utility_WithWeightCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Utility_WithWeightCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Utility_WithWeightCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Utility_WithWeightCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Utility_WithWeightCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Utility_WithWeightCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Utility_WithWeightCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Utility_WithWeightCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Utility_WithWeightCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Utility_WithWeightCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Utility_WithWeightCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Utility_WithWeightCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Utility_WithWeightCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Utility_WithWeightCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Utility_WithWeightCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Utility_WithWeightCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Utility_WithWeightCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Utility_WithWeightCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Utility_WithWeightCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Utility_WithWeightCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Utility_WithWeightCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Utility_WithWeightCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Utility_WithWeightCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Utility_WithWeightCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Utility_WithWeightCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}
func To_Utility_WithWeightCall_weight(in any) *pbgear.SpWeightsWeightV2Weight {
	return nil //simple field
}

func To_Value0(in any) *pbgear.Value0 {
	out := &pbgear.Value0{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Value0_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Value0_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Value0_Root{
			Root: To_Root(fields),
		}
	case 1:
		target = &pbgear.Value0_Signed{
			Signed: To_Signed(fields),
		}
	case 2:
		target = &pbgear.Value0_None{
			None: To_None(fields),
		}
	}
}

func To_VaraRuntimeNposSolution16(in any) *pbgear.VaraRuntimeNposSolution16 {
	out := &pbgear.VaraRuntimeNposSolution16{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// repeated field Votes1
	out.Votes1 = To_Repeated_VaraRuntimeNposSolution16_votes1(fields[0].Value)
	// repeated field Votes2
	out.Votes2 = To_Repeated_VaraRuntimeNposSolution16_votes2(fields[1].Value)
	// repeated field Votes3
	out.Votes3 = To_Repeated_VaraRuntimeNposSolution16_votes3(fields[2].Value)
	// repeated field Votes4
	out.Votes4 = To_Repeated_VaraRuntimeNposSolution16_votes4(fields[3].Value)
	// repeated field Votes5
	out.Votes5 = To_Repeated_VaraRuntimeNposSolution16_votes5(fields[4].Value)
	// repeated field Votes6
	out.Votes6 = To_Repeated_VaraRuntimeNposSolution16_votes6(fields[5].Value)
	// repeated field Votes7
	out.Votes7 = To_Repeated_VaraRuntimeNposSolution16_votes7(fields[6].Value)
	// repeated field Votes8
	out.Votes8 = To_Repeated_VaraRuntimeNposSolution16_votes8(fields[7].Value)
	// repeated field Votes9
	out.Votes9 = To_Repeated_VaraRuntimeNposSolution16_votes9(fields[8].Value)
	// repeated field Votes10
	out.Votes10 = To_Repeated_VaraRuntimeNposSolution16_votes10(fields[9].Value)
	// repeated field Votes11
	out.Votes11 = To_Repeated_VaraRuntimeNposSolution16_votes11(fields[10].Value)
	// repeated field Votes12
	out.Votes12 = To_Repeated_VaraRuntimeNposSolution16_votes12(fields[11].Value)
	// repeated field Votes13
	out.Votes13 = To_Repeated_VaraRuntimeNposSolution16_votes13(fields[12].Value)
	// repeated field Votes14
	out.Votes14 = To_Repeated_VaraRuntimeNposSolution16_votes14(fields[13].Value)
	// repeated field Votes15
	out.Votes15 = To_Repeated_VaraRuntimeNposSolution16_votes15(fields[14].Value)
	// repeated field Votes16
	out.Votes16 = To_Repeated_VaraRuntimeNposSolution16_votes16(fields[15].Value)

	return out //from message
}

func To_Repeated_VaraRuntimeNposSolution16_votes1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32Uint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes2(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes3(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes4(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes5(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes6(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes7(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes8(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes9(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes10(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes11(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes12(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes13(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes14(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes15(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes16(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32
	for _, item := range items {
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(item)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_VaraRuntimeRuntimeCall(in any) *pbgear.VaraRuntimeRuntimeCall {
	out := &pbgear.VaraRuntimeRuntimeCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Calls
	Set_OneOf_VaraRuntimeRuntimeCall_calls(out.Calls, fields[0].Value)

	return out //from message
}

func Set_OneOf_VaraRuntimeRuntimeCall_calls(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.VaraRuntimeRuntimeCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.VaraRuntimeRuntimeCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.VaraRuntimeRuntimeCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.VaraRuntimeRuntimeCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.VaraRuntimeRuntimeCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.VaraRuntimeRuntimeCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.VaraRuntimeRuntimeCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.VaraRuntimeRuntimeCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.VaraRuntimeRuntimeCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.VaraRuntimeRuntimeCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.VaraRuntimeRuntimeCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.VaraRuntimeRuntimeCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.VaraRuntimeRuntimeCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.VaraRuntimeRuntimeCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.VaraRuntimeRuntimeCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.VaraRuntimeRuntimeCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.VaraRuntimeRuntimeCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.VaraRuntimeRuntimeCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.VaraRuntimeRuntimeCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.VaraRuntimeRuntimeCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.VaraRuntimeRuntimeCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.VaraRuntimeRuntimeCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.VaraRuntimeRuntimeCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.VaraRuntimeRuntimeCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.VaraRuntimeRuntimeCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.VaraRuntimeRuntimeCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.VaraRuntimeRuntimeCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.VaraRuntimeRuntimeCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.VaraRuntimeRuntimeCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_VaraRuntimeSessionKeys(in any) *pbgear.VaraRuntimeSessionKeys {
	out := &pbgear.VaraRuntimeSessionKeys{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Babe
	out.Babe = To_VaraRuntimeSessionKeys_babe(fields[0].Value)
	// field Grandpa
	out.Grandpa = To_VaraRuntimeSessionKeys_grandpa(fields[1].Value)
	// field ImOnline
	out.ImOnline = To_VaraRuntimeSessionKeys_im_online(fields[2].Value)
	// field AuthorityDiscovery
	out.AuthorityDiscovery = To_VaraRuntimeSessionKeys_authority_discovery(fields[3].Value)

	return out //from message
}

func To_VaraRuntimeSessionKeys_babe(in any) *pbgear.SpConsensusBabeAppPublic {
	return nil //simple field
}
func To_VaraRuntimeSessionKeys_grandpa(in any) *pbgear.SpConsensusGrandpaAppPublic {
	return nil //simple field
}
func To_VaraRuntimeSessionKeys_im_online(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public {
	return nil //simple field
}
func To_VaraRuntimeSessionKeys_authority_discovery(in any) *pbgear.SpAuthorityDiscoveryAppPublic {
	return nil //simple field
}

func To_VestingPallet(in any) *pbgear.VestingPallet {
	out := &pbgear.VestingPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_VestingPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_VestingPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.VestingPallet_VestCall{
			VestCall: To_Vesting_VestCall(fields),
		}
	case 1:
		target = &pbgear.VestingPallet_VestOtherCall{
			VestOtherCall: To_Vesting_VestOtherCall(fields),
		}
	case 2:
		target = &pbgear.VestingPallet_VestedTransferCall{
			VestedTransferCall: To_Vesting_VestedTransferCall(fields),
		}
	case 3:
		target = &pbgear.VestingPallet_ForceVestedTransferCall{
			ForceVestedTransferCall: To_Vesting_ForceVestedTransferCall(fields),
		}
	case 4:
		target = &pbgear.VestingPallet_MergeSchedulesCall{
			MergeSchedulesCall: To_Vesting_MergeSchedulesCall(fields),
		}
	}
}

func To_Vesting_Address20(in any) *pbgear.Vesting_Address20 {
	out := &pbgear.Vesting_Address20{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Vesting_Address32(in any) *pbgear.Vesting_Address32 {
	out := &pbgear.Vesting_Address32{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Vesting_ForceVestedTransferCall(in any) *pbgear.Vesting_ForceVestedTransferCall {
	out := &pbgear.Vesting_ForceVestedTransferCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Source
	out.Source = To_Vesting_ForceVestedTransferCall_source(fields[0].Value)
	// field Target
	out.Target = To_Vesting_ForceVestedTransferCall_target(fields[1].Value)
	// field Schedule
	out.Schedule = To_Vesting_ForceVestedTransferCall_schedule(fields[2].Value)

	return out //from message
}

func To_Vesting_ForceVestedTransferCall_source(in any) *pbgear.Vesting_Source {
	return nil //simple field
}
func To_Vesting_ForceVestedTransferCall_target(in any) *pbgear.Vesting_Target {
	return nil //simple field
}
func To_Vesting_ForceVestedTransferCall_schedule(in any) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	return nil //simple field
}

func To_Vesting_Id(in any) *pbgear.Vesting_Id {
	out := &pbgear.Vesting_Id{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Vesting_Id_value0(fields[0].Value)

	return out //from message
}

func To_Vesting_Id_value0(in any) *pbgear.SpCoreCryptoAccountId32 {
	return nil //simple field
}

func To_Vesting_Index(in any) *pbgear.Vesting_Index {
	out := &pbgear.Vesting_Index{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value0
	out.Value0 = To_Vesting_Index_value0(fields[0].Value)

	return out //from message
}

func To_Vesting_Index_value0(in any) *pbgear.Vesting_TupleNull {
	return nil //simple field
}

func To_Vesting_MergeSchedulesCall(in any) *pbgear.Vesting_MergeSchedulesCall {
	out := &pbgear.Vesting_MergeSchedulesCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Schedule1Index
	out.Schedule1Index = To_uint32(fields[0].Value)
	// primitive field Schedule2Index
	out.Schedule2Index = To_uint32(fields[1].Value)

	return out //from message
}

func To_Vesting_PalletVestingVestingInfoVestingInfo(in any) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Locked
	out.Locked = To_string(fields[0].Value)
	// primitive field PerBlock
	out.PerBlock = To_string(fields[1].Value)
	// primitive field StartingBlock
	out.StartingBlock = To_uint32(fields[2].Value)

	return out //from message
}

func To_Vesting_Raw(in any) *pbgear.Vesting_Raw {
	out := &pbgear.Vesting_Raw{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// primitive field Value0
	out.Value0 = To_bytes(fields[0].Value)

	return out //from message
}

func To_Vesting_Source(in any) *pbgear.Vesting_Source {
	out := &pbgear.Vesting_Source{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Vesting_Source_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Vesting_Source_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Vesting_Source_Id{
			Id: To_Vesting_Id(fields),
		}
	case 1:
		target = &pbgear.Vesting_Source_Index{
			Index: To_Vesting_Index(fields),
		}
	case 2:
		target = &pbgear.Vesting_Source_Raw{
			Raw: To_Vesting_Raw(fields),
		}
	case 3:
		target = &pbgear.Vesting_Source_Address32{
			Address32: To_Vesting_Address32(fields),
		}
	case 4:
		target = &pbgear.Vesting_Source_Address20{
			Address20: To_Vesting_Address20(fields),
		}
	}
}

func To_Vesting_Target(in any) *pbgear.Vesting_Target {
	out := &pbgear.Vesting_Target{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Value
	Set_OneOf_Vesting_Target_value(out.Value, fields[0].Value)

	return out //from message
}

func Set_OneOf_Vesting_Target_value(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Vesting_Target_Id{
			Id: To_Vesting_Id(fields),
		}
	case 1:
		target = &pbgear.Vesting_Target_Index{
			Index: To_Vesting_Index(fields),
		}
	case 2:
		target = &pbgear.Vesting_Target_Raw{
			Raw: To_Vesting_Raw(fields),
		}
	case 3:
		target = &pbgear.Vesting_Target_Address32{
			Address32: To_Vesting_Address32(fields),
		}
	case 4:
		target = &pbgear.Vesting_Target_Address20{
			Address20: To_Vesting_Address20(fields),
		}
	}
}

func To_Vesting_TupleNull(in any) *pbgear.Vesting_TupleNull {
	out := &pbgear.Vesting_TupleNull{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Value
	out.Value = To_Vesting_TupleNull_value(fields[0].Value)

	return out //from message
}

func To_Vesting_TupleNull_value(in any) *pbgear.TupleNull {
	return nil //simple field
}

func To_Vesting_VestCall(in any) *pbgear.Vesting_VestCall {
	out := &pbgear.Vesting_VestCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	return out //from message
}

func To_Vesting_VestOtherCall(in any) *pbgear.Vesting_VestOtherCall {
	out := &pbgear.Vesting_VestOtherCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Target
	out.Target = To_Vesting_VestOtherCall_target(fields[0].Value)

	return out //from message
}

func To_Vesting_VestOtherCall_target(in any) *pbgear.Vesting_Target {
	return nil //simple field
}

func To_Vesting_VestedTransferCall(in any) *pbgear.Vesting_VestedTransferCall {
	out := &pbgear.Vesting_VestedTransferCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field Target
	out.Target = To_Vesting_VestedTransferCall_target(fields[0].Value)
	// field Schedule
	out.Schedule = To_Vesting_VestedTransferCall_schedule(fields[1].Value)

	return out //from message
}

func To_Vesting_VestedTransferCall_target(in any) *pbgear.Vesting_Target {
	return nil //simple field
}
func To_Vesting_VestedTransferCall_schedule(in any) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	return nil //simple field
}

func To_WhitelistPallet(in any) *pbgear.WhitelistPallet {
	out := &pbgear.WhitelistPallet{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_WhitelistPallet_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_WhitelistPallet_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.WhitelistPallet_WhitelistCallCall{
			WhitelistCallCall: To_Whitelist_WhitelistCallCall(fields),
		}
	case 1:
		target = &pbgear.WhitelistPallet_RemoveWhitelistedCallCall{
			RemoveWhitelistedCallCall: To_Whitelist_RemoveWhitelistedCallCall(fields),
		}
	case 2:
		target = &pbgear.WhitelistPallet_DispatchWhitelistedCallCall{
			DispatchWhitelistedCallCall: To_Whitelist_DispatchWhitelistedCallCall(fields),
		}
	case 3:
		target = &pbgear.WhitelistPallet_DispatchWhitelistedCallWithPreimageCall{
			DispatchWhitelistedCallWithPreimageCall: To_Whitelist_DispatchWhitelistedCallWithPreimageCall(fields),
		}
	}
}

func To_Whitelist_DispatchWhitelistedCallCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field CallHash
	out.CallHash = To_Whitelist_DispatchWhitelistedCallCall_call_hash(fields[0].Value)
	// primitive field CallEncodedLen
	out.CallEncodedLen = To_uint32(fields[1].Value)
	// field CallWeightWitness
	out.CallWeightWitness = To_Whitelist_DispatchWhitelistedCallCall_call_weight_witness(fields[2].Value)

	return out //from message
}

func To_Whitelist_DispatchWhitelistedCallCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}
func To_Whitelist_DispatchWhitelistedCallCall_call_weight_witness(in any) *pbgear.SpWeightsWeightV2Weight {
	return nil //simple field
}

func To_Whitelist_DispatchWhitelistedCallWithPreimageCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// oneOf field Call
	Set_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call(out.Call, fields[0].Value)

	return out //from message
}

func Set_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call(target any, in any) {
	variantIn := in.(*registry.VariantWTF)
	variantFields := variantIn.Value.(registry.DecodedFields)
	fields := []*registry.DecodedField(variantFields)

	switch variantIn.VariantByte {
	case 0:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_System{
			System: To_SystemPallet(fields),
		}
	case 1:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Timestamp{
			Timestamp: To_TimestampPallet(fields),
		}
	case 3:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Babe{
			Babe: To_BabePallet(fields),
		}
	case 4:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Grandpa{
			Grandpa: To_GrandpaPallet(fields),
		}
	case 5:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Balances{
			Balances: To_BalancesPallet(fields),
		}
	case 10:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Vesting{
			Vesting: To_VestingPallet(fields),
		}
	case 11:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_BagsList{
			BagsList: To_BagsListPallet(fields),
		}
	case 12:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ImOnline{
			ImOnline: To_ImOnlinePallet(fields),
		}
	case 13:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Staking{
			Staking: To_StakingPallet(fields),
		}
	case 7:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Session{
			Session: To_SessionPallet(fields),
		}
	case 14:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Treasury{
			Treasury: To_TreasuryPallet(fields),
		}
	case 8:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Utility{
			Utility: To_UtilityPallet(fields),
		}
	case 16:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ConvictionVoting{
			ConvictionVoting: To_ConvictionVotingPallet(fields),
		}
	case 17:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Referenda{
			Referenda: To_ReferendaPallet(fields),
		}
	case 18:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipCollective{
			FellowshipCollective: To_FellowshipCollectivePallet(fields),
		}
	case 19:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_FellowshipReferenda{
			FellowshipReferenda: To_FellowshipReferendaPallet(fields),
		}
	case 21:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Whitelist{
			Whitelist: To_WhitelistPallet(fields),
		}
	case 22:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Scheduler{
			Scheduler: To_SchedulerPallet(fields),
		}
	case 23:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Preimage{
			Preimage: To_PreimagePallet(fields),
		}
	case 24:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Identity{
			Identity: To_IdentityPallet(fields),
		}
	case 25:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Proxy{
			Proxy: To_ProxyPallet(fields),
		}
	case 26:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Multisig{
			Multisig: To_MultisigPallet(fields),
		}
	case 27:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ElectionProviderMultiPhase{
			ElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(fields),
		}
	case 29:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Bounties{
			Bounties: To_BountiesPallet(fields),
		}
	case 30:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_ChildBounties{
			ChildBounties: To_ChildBountiesPallet(fields),
		}
	case 31:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_NominationPools{
			NominationPools: To_NominationPoolsPallet(fields),
		}
	case 104:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_Gear{
			Gear: To_GearPallet(fields),
		}
	case 106:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_StakingRewards{
			StakingRewards: To_StakingRewardsPallet(fields),
		}
	case 107:
		target = &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_GearVoucher{
			GearVoucher: To_GearVoucherPallet(fields),
		}
	}
}

func To_Whitelist_RemoveWhitelistedCallCall(in any) *pbgear.Whitelist_RemoveWhitelistedCallCall {
	out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field CallHash
	out.CallHash = To_Whitelist_RemoveWhitelistedCallCall_call_hash(fields[0].Value)

	return out //from message
}

func To_Whitelist_RemoveWhitelistedCallCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_Whitelist_WhitelistCallCall(in any) *pbgear.Whitelist_WhitelistCallCall {
	out := &pbgear.Whitelist_WhitelistCallCall{}
	ff := in.(registry.DecodedFields)
	fields := []*registry.DecodedField(ff)
	_ = fields

	// field CallHash
	out.CallHash = To_Whitelist_WhitelistCallCall_call_hash(fields[0].Value)

	return out //from message
}

func To_Whitelist_WhitelistCallCall_call_hash(in any) *pbgear.PrimitiveTypesH256 {
	return nil //simple field
}

func To_bytes(a any) []byte {
	in := a.([]interface{})
	var out []byte
	for _, f := range in {
		u := uint8(f.(types.U8))
		out = append(out, byte(u))
	}
	return out
}

func To_string(i any) string {
	switch v := i.(type) {
	case types.Text:
		return string(v)
	case types.U128:
		return v.String()
	case types.U256:
		return v.String()
	case types.I128:
		return v.String()
	case types.I256:
		return v.String()
	default:
		panic("Unknown type")
	}
}

func To_uint32(i any) uint32 {
	switch v := i.(type) {
	case types.UCompact:
		return uint32(v.Int64())
	case types.U8:
		return uint32(v)
	case types.U16:
		return uint32(v)
	case types.U32:
		return uint32(v)
	default:
		panic("Unknown type")
	}
}

func To_Optional_uint32(i any) *uint32 {
	o := To_uint32(i)
	return &o
}

func To_uint64(a any) uint64 {
	switch v := a.(type) {
	case types.UCompact:
		return uint64(v.Int64())
	case types.U8:
		return uint64(v)
	case types.U16:
		return uint64(v)
	case types.U32:
		return uint64(v)
	case types.U64:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	default:
		panic("Unknown type")
	}
}

func To_Optional_uint64(i any) *uint64 {
	o := To_uint64(i)
	return &o
}

func To_Optional_string(a any) *string {
	s := To_string(a)
	return &s
}

func To_bool(b any) bool {
	return b.(bool)
}

func To_Optional_bool(b any) *bool {
	o := To_bool(b)
	return &o
}

func To_Repeated_uint32(a any) []uint32 {
	panic("To_repeated_uint32")
}

var FuncMap map[string]reflect.Value

func init() {
	FuncMap = make(map[string]reflect.Value)
	FuncMap["To_AllowedSlots"] = reflect.ValueOf(To_AllowedSlots)
	FuncMap["To_AuthorityDiscoveryPallet"] = reflect.ValueOf(To_AuthorityDiscoveryPallet)
	FuncMap["To_AuthorshipPallet"] = reflect.ValueOf(To_AuthorshipPallet)
	FuncMap["To_BTreeSet"] = reflect.ValueOf(To_BTreeSet)
	FuncMap["To_BabePallet"] = reflect.ValueOf(To_BabePallet)
	FuncMap["To_Babe_BabeTrieNodesList"] = reflect.ValueOf(To_Babe_BabeTrieNodesList)
	FuncMap["To_Babe_Config"] = reflect.ValueOf(To_Babe_Config)
	FuncMap["To_Babe_Consensus"] = reflect.ValueOf(To_Babe_Consensus)
	FuncMap["To_Babe_Logs"] = reflect.ValueOf(To_Babe_Logs)
	FuncMap["To_Babe_Other"] = reflect.ValueOf(To_Babe_Other)
	FuncMap["To_Babe_PlanConfigChangeCall"] = reflect.ValueOf(To_Babe_PlanConfigChangeCall)
	FuncMap["To_Babe_PreRuntime"] = reflect.ValueOf(To_Babe_PreRuntime)
	FuncMap["To_Babe_ReportEquivocationCall"] = reflect.ValueOf(To_Babe_ReportEquivocationCall)
	FuncMap["To_Babe_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Babe_ReportEquivocationUnsignedCall)
	FuncMap["To_Babe_RuntimeEnvironmentUpdated"] = reflect.ValueOf(To_Babe_RuntimeEnvironmentUpdated)
	FuncMap["To_Babe_Seal"] = reflect.ValueOf(To_Babe_Seal)
	FuncMap["To_Babe_TupleUint64Uint64"] = reflect.ValueOf(To_Babe_TupleUint64Uint64)
	FuncMap["To_Babe_V1"] = reflect.ValueOf(To_Babe_V1)
	FuncMap["To_BagsListPallet"] = reflect.ValueOf(To_BagsListPallet)
	FuncMap["To_BagsList_Address20"] = reflect.ValueOf(To_BagsList_Address20)
	FuncMap["To_BagsList_Address32"] = reflect.ValueOf(To_BagsList_Address32)
	FuncMap["To_BagsList_Dislocated"] = reflect.ValueOf(To_BagsList_Dislocated)
	FuncMap["To_BagsList_Heavier"] = reflect.ValueOf(To_BagsList_Heavier)
	FuncMap["To_BagsList_Id"] = reflect.ValueOf(To_BagsList_Id)
	FuncMap["To_BagsList_Index"] = reflect.ValueOf(To_BagsList_Index)
	FuncMap["To_BagsList_Lighter"] = reflect.ValueOf(To_BagsList_Lighter)
	FuncMap["To_BagsList_PutInFrontOfCall"] = reflect.ValueOf(To_BagsList_PutInFrontOfCall)
	FuncMap["To_BagsList_PutInFrontOfOtherCall"] = reflect.ValueOf(To_BagsList_PutInFrontOfOtherCall)
	FuncMap["To_BagsList_Raw"] = reflect.ValueOf(To_BagsList_Raw)
	FuncMap["To_BagsList_RebagCall"] = reflect.ValueOf(To_BagsList_RebagCall)
	FuncMap["To_BagsList_TupleNull"] = reflect.ValueOf(To_BagsList_TupleNull)
	FuncMap["To_BalancesPallet"] = reflect.ValueOf(To_BalancesPallet)
	FuncMap["To_Balances_Address20"] = reflect.ValueOf(To_Balances_Address20)
	FuncMap["To_Balances_Address32"] = reflect.ValueOf(To_Balances_Address32)
	FuncMap["To_Balances_Dest"] = reflect.ValueOf(To_Balances_Dest)
	FuncMap["To_Balances_ForceSetBalanceCall"] = reflect.ValueOf(To_Balances_ForceSetBalanceCall)
	FuncMap["To_Balances_ForceTransferCall"] = reflect.ValueOf(To_Balances_ForceTransferCall)
	FuncMap["To_Balances_ForceUnreserveCall"] = reflect.ValueOf(To_Balances_ForceUnreserveCall)
	FuncMap["To_Balances_Id"] = reflect.ValueOf(To_Balances_Id)
	FuncMap["To_Balances_Index"] = reflect.ValueOf(To_Balances_Index)
	FuncMap["To_Balances_Raw"] = reflect.ValueOf(To_Balances_Raw)
	FuncMap["To_Balances_Source"] = reflect.ValueOf(To_Balances_Source)
	FuncMap["To_Balances_TransferAllCall"] = reflect.ValueOf(To_Balances_TransferAllCall)
	FuncMap["To_Balances_TransferAllowDeathCall"] = reflect.ValueOf(To_Balances_TransferAllowDeathCall)
	FuncMap["To_Balances_TransferKeepAliveCall"] = reflect.ValueOf(To_Balances_TransferKeepAliveCall)
	FuncMap["To_Balances_TupleNull"] = reflect.ValueOf(To_Balances_TupleNull)
	FuncMap["To_Balances_UpgradeAccountsCall"] = reflect.ValueOf(To_Balances_UpgradeAccountsCall)
	FuncMap["To_Balances_Who"] = reflect.ValueOf(To_Balances_Who)
	FuncMap["To_BoundedCollectionsBoundedVecBoundedVec"] = reflect.ValueOf(To_BoundedCollectionsBoundedVecBoundedVec)
	FuncMap["To_BountiesPallet"] = reflect.ValueOf(To_BountiesPallet)
	FuncMap["To_Bounties_AcceptCuratorCall"] = reflect.ValueOf(To_Bounties_AcceptCuratorCall)
	FuncMap["To_Bounties_Address20"] = reflect.ValueOf(To_Bounties_Address20)
	FuncMap["To_Bounties_Address32"] = reflect.ValueOf(To_Bounties_Address32)
	FuncMap["To_Bounties_ApproveBountyCall"] = reflect.ValueOf(To_Bounties_ApproveBountyCall)
	FuncMap["To_Bounties_AwardBountyCall"] = reflect.ValueOf(To_Bounties_AwardBountyCall)
	FuncMap["To_Bounties_Beneficiary"] = reflect.ValueOf(To_Bounties_Beneficiary)
	FuncMap["To_Bounties_ClaimBountyCall"] = reflect.ValueOf(To_Bounties_ClaimBountyCall)
	FuncMap["To_Bounties_CloseBountyCall"] = reflect.ValueOf(To_Bounties_CloseBountyCall)
	FuncMap["To_Bounties_Curator"] = reflect.ValueOf(To_Bounties_Curator)
	FuncMap["To_Bounties_ExtendBountyExpiryCall"] = reflect.ValueOf(To_Bounties_ExtendBountyExpiryCall)
	FuncMap["To_Bounties_Id"] = reflect.ValueOf(To_Bounties_Id)
	FuncMap["To_Bounties_Index"] = reflect.ValueOf(To_Bounties_Index)
	FuncMap["To_Bounties_ProposeBountyCall"] = reflect.ValueOf(To_Bounties_ProposeBountyCall)
	FuncMap["To_Bounties_ProposeCuratorCall"] = reflect.ValueOf(To_Bounties_ProposeCuratorCall)
	FuncMap["To_Bounties_Raw"] = reflect.ValueOf(To_Bounties_Raw)
	FuncMap["To_Bounties_TupleNull"] = reflect.ValueOf(To_Bounties_TupleNull)
	FuncMap["To_Bounties_UnassignCuratorCall"] = reflect.ValueOf(To_Bounties_UnassignCuratorCall)
	FuncMap["To_ChildBountiesPallet"] = reflect.ValueOf(To_ChildBountiesPallet)
	FuncMap["To_ChildBounties_AcceptCuratorCall"] = reflect.ValueOf(To_ChildBounties_AcceptCuratorCall)
	FuncMap["To_ChildBounties_AddChildBountyCall"] = reflect.ValueOf(To_ChildBounties_AddChildBountyCall)
	FuncMap["To_ChildBounties_Address20"] = reflect.ValueOf(To_ChildBounties_Address20)
	FuncMap["To_ChildBounties_Address32"] = reflect.ValueOf(To_ChildBounties_Address32)
	FuncMap["To_ChildBounties_AwardChildBountyCall"] = reflect.ValueOf(To_ChildBounties_AwardChildBountyCall)
	FuncMap["To_ChildBounties_Beneficiary"] = reflect.ValueOf(To_ChildBounties_Beneficiary)
	FuncMap["To_ChildBounties_ClaimChildBountyCall"] = reflect.ValueOf(To_ChildBounties_ClaimChildBountyCall)
	FuncMap["To_ChildBounties_CloseChildBountyCall"] = reflect.ValueOf(To_ChildBounties_CloseChildBountyCall)
	FuncMap["To_ChildBounties_Curator"] = reflect.ValueOf(To_ChildBounties_Curator)
	FuncMap["To_ChildBounties_Id"] = reflect.ValueOf(To_ChildBounties_Id)
	FuncMap["To_ChildBounties_Index"] = reflect.ValueOf(To_ChildBounties_Index)
	FuncMap["To_ChildBounties_ProposeCuratorCall"] = reflect.ValueOf(To_ChildBounties_ProposeCuratorCall)
	FuncMap["To_ChildBounties_Raw"] = reflect.ValueOf(To_ChildBounties_Raw)
	FuncMap["To_ChildBounties_TupleNull"] = reflect.ValueOf(To_ChildBounties_TupleNull)
	FuncMap["To_ChildBounties_UnassignCuratorCall"] = reflect.ValueOf(To_ChildBounties_UnassignCuratorCall)
	FuncMap["To_ConvictionVotingPallet"] = reflect.ValueOf(To_ConvictionVotingPallet)
	FuncMap["To_ConvictionVoting_Address20"] = reflect.ValueOf(To_ConvictionVoting_Address20)
	FuncMap["To_ConvictionVoting_Address32"] = reflect.ValueOf(To_ConvictionVoting_Address32)
	FuncMap["To_ConvictionVoting_Conviction"] = reflect.ValueOf(To_ConvictionVoting_Conviction)
	FuncMap["To_ConvictionVoting_DelegateCall"] = reflect.ValueOf(To_ConvictionVoting_DelegateCall)
	FuncMap["To_ConvictionVoting_Id"] = reflect.ValueOf(To_ConvictionVoting_Id)
	FuncMap["To_ConvictionVoting_Index"] = reflect.ValueOf(To_ConvictionVoting_Index)
	FuncMap["To_ConvictionVoting_Locked1X"] = reflect.ValueOf(To_ConvictionVoting_Locked1X)
	FuncMap["To_ConvictionVoting_Locked2X"] = reflect.ValueOf(To_ConvictionVoting_Locked2X)
	FuncMap["To_ConvictionVoting_Locked3X"] = reflect.ValueOf(To_ConvictionVoting_Locked3X)
	FuncMap["To_ConvictionVoting_Locked4X"] = reflect.ValueOf(To_ConvictionVoting_Locked4X)
	FuncMap["To_ConvictionVoting_Locked5X"] = reflect.ValueOf(To_ConvictionVoting_Locked5X)
	FuncMap["To_ConvictionVoting_Locked6X"] = reflect.ValueOf(To_ConvictionVoting_Locked6X)
	FuncMap["To_ConvictionVoting_None"] = reflect.ValueOf(To_ConvictionVoting_None)
	FuncMap["To_ConvictionVoting_PalletConvictionVotingVoteVote"] = reflect.ValueOf(To_ConvictionVoting_PalletConvictionVotingVoteVote)
	FuncMap["To_ConvictionVoting_Raw"] = reflect.ValueOf(To_ConvictionVoting_Raw)
	FuncMap["To_ConvictionVoting_RemoveOtherVoteCall"] = reflect.ValueOf(To_ConvictionVoting_RemoveOtherVoteCall)
	FuncMap["To_ConvictionVoting_RemoveVoteCall"] = reflect.ValueOf(To_ConvictionVoting_RemoveVoteCall)
	FuncMap["To_ConvictionVoting_Split"] = reflect.ValueOf(To_ConvictionVoting_Split)
	FuncMap["To_ConvictionVoting_SplitAbstain"] = reflect.ValueOf(To_ConvictionVoting_SplitAbstain)
	FuncMap["To_ConvictionVoting_Standard"] = reflect.ValueOf(To_ConvictionVoting_Standard)
	FuncMap["To_ConvictionVoting_Target"] = reflect.ValueOf(To_ConvictionVoting_Target)
	FuncMap["To_ConvictionVoting_To"] = reflect.ValueOf(To_ConvictionVoting_To)
	FuncMap["To_ConvictionVoting_TupleNull"] = reflect.ValueOf(To_ConvictionVoting_TupleNull)
	FuncMap["To_ConvictionVoting_UndelegateCall"] = reflect.ValueOf(To_ConvictionVoting_UndelegateCall)
	FuncMap["To_ConvictionVoting_UnlockCall"] = reflect.ValueOf(To_ConvictionVoting_UnlockCall)
	FuncMap["To_ConvictionVoting_Vote"] = reflect.ValueOf(To_ConvictionVoting_Vote)
	FuncMap["To_ConvictionVoting_VoteCall"] = reflect.ValueOf(To_ConvictionVoting_VoteCall)
	FuncMap["To_ElectionProviderMultiPhasePallet"] = reflect.ValueOf(To_ElectionProviderMultiPhasePallet)
	FuncMap["To_ElectionProviderMultiPhase_GovernanceFallbackCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_GovernanceFallbackCall)
	FuncMap["To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution"] = reflect.ValueOf(To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution)
	FuncMap["To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize"] = reflect.ValueOf(To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize)
	FuncMap["To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall)
	FuncMap["To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall)
	FuncMap["To_ElectionProviderMultiPhase_SubmitCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SubmitCall)
	FuncMap["To_ElectionProviderMultiPhase_SubmitUnsignedCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SubmitUnsignedCall)
	FuncMap["To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport)
	FuncMap["To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32)
	FuncMap["To_ElectionProviderMultiPhase_TupleUint32Uint32"] = reflect.ValueOf(To_ElectionProviderMultiPhase_TupleUint32Uint32)
	FuncMap["To_FellowshipCollectivePallet"] = reflect.ValueOf(To_FellowshipCollectivePallet)
	FuncMap["To_FellowshipCollective_AddMemberCall"] = reflect.ValueOf(To_FellowshipCollective_AddMemberCall)
	FuncMap["To_FellowshipCollective_Address20"] = reflect.ValueOf(To_FellowshipCollective_Address20)
	FuncMap["To_FellowshipCollective_Address32"] = reflect.ValueOf(To_FellowshipCollective_Address32)
	FuncMap["To_FellowshipCollective_CleanupPollCall"] = reflect.ValueOf(To_FellowshipCollective_CleanupPollCall)
	FuncMap["To_FellowshipCollective_DemoteMemberCall"] = reflect.ValueOf(To_FellowshipCollective_DemoteMemberCall)
	FuncMap["To_FellowshipCollective_Id"] = reflect.ValueOf(To_FellowshipCollective_Id)
	FuncMap["To_FellowshipCollective_Index"] = reflect.ValueOf(To_FellowshipCollective_Index)
	FuncMap["To_FellowshipCollective_PromoteMemberCall"] = reflect.ValueOf(To_FellowshipCollective_PromoteMemberCall)
	FuncMap["To_FellowshipCollective_Raw"] = reflect.ValueOf(To_FellowshipCollective_Raw)
	FuncMap["To_FellowshipCollective_RemoveMemberCall"] = reflect.ValueOf(To_FellowshipCollective_RemoveMemberCall)
	FuncMap["To_FellowshipCollective_TupleNull"] = reflect.ValueOf(To_FellowshipCollective_TupleNull)
	FuncMap["To_FellowshipCollective_VoteCall"] = reflect.ValueOf(To_FellowshipCollective_VoteCall)
	FuncMap["To_FellowshipCollective_Who"] = reflect.ValueOf(To_FellowshipCollective_Who)
	FuncMap["To_FellowshipReferendaPallet"] = reflect.ValueOf(To_FellowshipReferendaPallet)
	FuncMap["To_FellowshipReferenda_After"] = reflect.ValueOf(To_FellowshipReferenda_After)
	FuncMap["To_FellowshipReferenda_At"] = reflect.ValueOf(To_FellowshipReferenda_At)
	FuncMap["To_FellowshipReferenda_CancelCall"] = reflect.ValueOf(To_FellowshipReferenda_CancelCall)
	FuncMap["To_FellowshipReferenda_EnactmentMoment"] = reflect.ValueOf(To_FellowshipReferenda_EnactmentMoment)
	FuncMap["To_FellowshipReferenda_Inline"] = reflect.ValueOf(To_FellowshipReferenda_Inline)
	FuncMap["To_FellowshipReferenda_KillCall"] = reflect.ValueOf(To_FellowshipReferenda_KillCall)
	FuncMap["To_FellowshipReferenda_Legacy"] = reflect.ValueOf(To_FellowshipReferenda_Legacy)
	FuncMap["To_FellowshipReferenda_Lookup"] = reflect.ValueOf(To_FellowshipReferenda_Lookup)
	FuncMap["To_FellowshipReferenda_NudgeReferendumCall"] = reflect.ValueOf(To_FellowshipReferenda_NudgeReferendumCall)
	FuncMap["To_FellowshipReferenda_OneFewerDecidingCall"] = reflect.ValueOf(To_FellowshipReferenda_OneFewerDecidingCall)
	FuncMap["To_FellowshipReferenda_Origins"] = reflect.ValueOf(To_FellowshipReferenda_Origins)
	FuncMap["To_FellowshipReferenda_PlaceDecisionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_PlaceDecisionDepositCall)
	FuncMap["To_FellowshipReferenda_Proposal"] = reflect.ValueOf(To_FellowshipReferenda_Proposal)
	FuncMap["To_FellowshipReferenda_ProposalOrigin"] = reflect.ValueOf(To_FellowshipReferenda_ProposalOrigin)
	FuncMap["To_FellowshipReferenda_RefundDecisionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_RefundDecisionDepositCall)
	FuncMap["To_FellowshipReferenda_RefundSubmissionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_RefundSubmissionDepositCall)
	FuncMap["To_FellowshipReferenda_SetMetadataCall"] = reflect.ValueOf(To_FellowshipReferenda_SetMetadataCall)
	FuncMap["To_FellowshipReferenda_SubmitCall"] = reflect.ValueOf(To_FellowshipReferenda_SubmitCall)
	FuncMap["To_FellowshipReferenda_System"] = reflect.ValueOf(To_FellowshipReferenda_System)
	FuncMap["To_FellowshipReferenda_Void"] = reflect.ValueOf(To_FellowshipReferenda_Void)
	FuncMap["To_FinalityGrandpaEquivocation"] = reflect.ValueOf(To_FinalityGrandpaEquivocation)
	FuncMap["To_FinalityGrandpaPrecommit"] = reflect.ValueOf(To_FinalityGrandpaPrecommit)
	FuncMap["To_FinalityGrandpaPrevote"] = reflect.ValueOf(To_FinalityGrandpaPrevote)
	FuncMap["To_GearBankPallet"] = reflect.ValueOf(To_GearBankPallet)
	FuncMap["To_GearBuiltinPallet"] = reflect.ValueOf(To_GearBuiltinPallet)
	FuncMap["To_GearGasPallet"] = reflect.ValueOf(To_GearGasPallet)
	FuncMap["To_GearMessengerPallet"] = reflect.ValueOf(To_GearMessengerPallet)
	FuncMap["To_GearPallet"] = reflect.ValueOf(To_GearPallet)
	FuncMap["To_GearPaymentPallet"] = reflect.ValueOf(To_GearPaymentPallet)
	FuncMap["To_GearProgramPallet"] = reflect.ValueOf(To_GearProgramPallet)
	FuncMap["To_GearSchedulerPallet"] = reflect.ValueOf(To_GearSchedulerPallet)
	FuncMap["To_GearVoucherPallet"] = reflect.ValueOf(To_GearVoucherPallet)
	FuncMap["To_GearVoucher_AppendPrograms"] = reflect.ValueOf(To_GearVoucher_AppendPrograms)
	FuncMap["To_GearVoucher_Call"] = reflect.ValueOf(To_GearVoucher_Call)
	FuncMap["To_GearVoucher_CallCall"] = reflect.ValueOf(To_GearVoucher_CallCall)
	FuncMap["To_GearVoucher_CallDeprecatedCall"] = reflect.ValueOf(To_GearVoucher_CallDeprecatedCall)
	FuncMap["To_GearVoucher_DeclineCall"] = reflect.ValueOf(To_GearVoucher_DeclineCall)
	FuncMap["To_GearVoucher_DeclineVoucher"] = reflect.ValueOf(To_GearVoucher_DeclineVoucher)
	FuncMap["To_GearVoucher_IssueCall"] = reflect.ValueOf(To_GearVoucher_IssueCall)
	FuncMap["To_GearVoucher_None"] = reflect.ValueOf(To_GearVoucher_None)
	FuncMap["To_GearVoucher_PalletGearVoucherInternalVoucherId"] = reflect.ValueOf(To_GearVoucher_PalletGearVoucherInternalVoucherId)
	FuncMap["To_GearVoucher_RevokeCall"] = reflect.ValueOf(To_GearVoucher_RevokeCall)
	FuncMap["To_GearVoucher_SendMessage"] = reflect.ValueOf(To_GearVoucher_SendMessage)
	FuncMap["To_GearVoucher_SendReply"] = reflect.ValueOf(To_GearVoucher_SendReply)
	FuncMap["To_GearVoucher_Some"] = reflect.ValueOf(To_GearVoucher_Some)
	FuncMap["To_GearVoucher_UpdateCall"] = reflect.ValueOf(To_GearVoucher_UpdateCall)
	FuncMap["To_GearVoucher_UploadCode"] = reflect.ValueOf(To_GearVoucher_UploadCode)
	FuncMap["To_Gear_ClaimValueCall"] = reflect.ValueOf(To_Gear_ClaimValueCall)
	FuncMap["To_Gear_CreateProgramCall"] = reflect.ValueOf(To_Gear_CreateProgramCall)
	FuncMap["To_Gear_RunCall"] = reflect.ValueOf(To_Gear_RunCall)
	FuncMap["To_Gear_SendMessageCall"] = reflect.ValueOf(To_Gear_SendMessageCall)
	FuncMap["To_Gear_SendReplyCall"] = reflect.ValueOf(To_Gear_SendReplyCall)
	FuncMap["To_Gear_SetExecuteInherentCall"] = reflect.ValueOf(To_Gear_SetExecuteInherentCall)
	FuncMap["To_Gear_UploadCodeCall"] = reflect.ValueOf(To_Gear_UploadCodeCall)
	FuncMap["To_Gear_UploadProgramCall"] = reflect.ValueOf(To_Gear_UploadProgramCall)
	FuncMap["To_GprimitivesActorId"] = reflect.ValueOf(To_GprimitivesActorId)
	FuncMap["To_GprimitivesCodeId"] = reflect.ValueOf(To_GprimitivesCodeId)
	FuncMap["To_GprimitivesMessageId"] = reflect.ValueOf(To_GprimitivesMessageId)
	FuncMap["To_GrandpaPallet"] = reflect.ValueOf(To_GrandpaPallet)
	FuncMap["To_Grandpa_Equivocation"] = reflect.ValueOf(To_Grandpa_Equivocation)
	FuncMap["To_Grandpa_GrandpaTrieNodesList"] = reflect.ValueOf(To_Grandpa_GrandpaTrieNodesList)
	FuncMap["To_Grandpa_NoteStalledCall"] = reflect.ValueOf(To_Grandpa_NoteStalledCall)
	FuncMap["To_Grandpa_Precommit"] = reflect.ValueOf(To_Grandpa_Precommit)
	FuncMap["To_Grandpa_Prevote"] = reflect.ValueOf(To_Grandpa_Prevote)
	FuncMap["To_Grandpa_ReportEquivocationCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationCall)
	FuncMap["To_Grandpa_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationUnsignedCall)
	FuncMap["To_HistoricalPallet"] = reflect.ValueOf(To_HistoricalPallet)
	FuncMap["To_IdentityPallet"] = reflect.ValueOf(To_IdentityPallet)
	FuncMap["To_Identity_Account"] = reflect.ValueOf(To_Identity_Account)
	FuncMap["To_Identity_AddRegistrarCall"] = reflect.ValueOf(To_Identity_AddRegistrarCall)
	FuncMap["To_Identity_AddSubCall"] = reflect.ValueOf(To_Identity_AddSubCall)
	FuncMap["To_Identity_Address20"] = reflect.ValueOf(To_Identity_Address20)
	FuncMap["To_Identity_Address32"] = reflect.ValueOf(To_Identity_Address32)
	FuncMap["To_Identity_BlakeTwo256"] = reflect.ValueOf(To_Identity_BlakeTwo256)
	FuncMap["To_Identity_CancelRequestCall"] = reflect.ValueOf(To_Identity_CancelRequestCall)
	FuncMap["To_Identity_ClearIdentityCall"] = reflect.ValueOf(To_Identity_ClearIdentityCall)
	FuncMap["To_Identity_Data"] = reflect.ValueOf(To_Identity_Data)
	FuncMap["To_Identity_Display"] = reflect.ValueOf(To_Identity_Display)
	FuncMap["To_Identity_Email"] = reflect.ValueOf(To_Identity_Email)
	FuncMap["To_Identity_Erroneous"] = reflect.ValueOf(To_Identity_Erroneous)
	FuncMap["To_Identity_FeePaid"] = reflect.ValueOf(To_Identity_FeePaid)
	FuncMap["To_Identity_Id"] = reflect.ValueOf(To_Identity_Id)
	FuncMap["To_Identity_Image"] = reflect.ValueOf(To_Identity_Image)
	FuncMap["To_Identity_Index"] = reflect.ValueOf(To_Identity_Index)
	FuncMap["To_Identity_Judgement"] = reflect.ValueOf(To_Identity_Judgement)
	FuncMap["To_Identity_Keccak256"] = reflect.ValueOf(To_Identity_Keccak256)
	FuncMap["To_Identity_KillIdentityCall"] = reflect.ValueOf(To_Identity_KillIdentityCall)
	FuncMap["To_Identity_KnownGood"] = reflect.ValueOf(To_Identity_KnownGood)
	FuncMap["To_Identity_Legal"] = reflect.ValueOf(To_Identity_Legal)
	FuncMap["To_Identity_LowQuality"] = reflect.ValueOf(To_Identity_LowQuality)
	FuncMap["To_Identity_New"] = reflect.ValueOf(To_Identity_New)
	FuncMap["To_Identity_None"] = reflect.ValueOf(To_Identity_None)
	FuncMap["To_Identity_OutOfDate"] = reflect.ValueOf(To_Identity_OutOfDate)
	FuncMap["To_Identity_PalletIdentitySimpleIdentityInfo"] = reflect.ValueOf(To_Identity_PalletIdentitySimpleIdentityInfo)
	FuncMap["To_Identity_PalletIdentityTypesBitFlags"] = reflect.ValueOf(To_Identity_PalletIdentityTypesBitFlags)
	FuncMap["To_Identity_ProvideJudgementCall"] = reflect.ValueOf(To_Identity_ProvideJudgementCall)
	FuncMap["To_Identity_QuitSubCall"] = reflect.ValueOf(To_Identity_QuitSubCall)
	FuncMap["To_Identity_Raw"] = reflect.ValueOf(To_Identity_Raw)
	FuncMap["To_Identity_Raw0"] = reflect.ValueOf(To_Identity_Raw0)
	FuncMap["To_Identity_Raw1"] = reflect.ValueOf(To_Identity_Raw1)
	FuncMap["To_Identity_Raw10"] = reflect.ValueOf(To_Identity_Raw10)
	FuncMap["To_Identity_Raw11"] = reflect.ValueOf(To_Identity_Raw11)
	FuncMap["To_Identity_Raw12"] = reflect.ValueOf(To_Identity_Raw12)
	FuncMap["To_Identity_Raw13"] = reflect.ValueOf(To_Identity_Raw13)
	FuncMap["To_Identity_Raw14"] = reflect.ValueOf(To_Identity_Raw14)
	FuncMap["To_Identity_Raw15"] = reflect.ValueOf(To_Identity_Raw15)
	FuncMap["To_Identity_Raw16"] = reflect.ValueOf(To_Identity_Raw16)
	FuncMap["To_Identity_Raw17"] = reflect.ValueOf(To_Identity_Raw17)
	FuncMap["To_Identity_Raw18"] = reflect.ValueOf(To_Identity_Raw18)
	FuncMap["To_Identity_Raw19"] = reflect.ValueOf(To_Identity_Raw19)
	FuncMap["To_Identity_Raw2"] = reflect.ValueOf(To_Identity_Raw2)
	FuncMap["To_Identity_Raw20"] = reflect.ValueOf(To_Identity_Raw20)
	FuncMap["To_Identity_Raw21"] = reflect.ValueOf(To_Identity_Raw21)
	FuncMap["To_Identity_Raw22"] = reflect.ValueOf(To_Identity_Raw22)
	FuncMap["To_Identity_Raw23"] = reflect.ValueOf(To_Identity_Raw23)
	FuncMap["To_Identity_Raw24"] = reflect.ValueOf(To_Identity_Raw24)
	FuncMap["To_Identity_Raw25"] = reflect.ValueOf(To_Identity_Raw25)
	FuncMap["To_Identity_Raw26"] = reflect.ValueOf(To_Identity_Raw26)
	FuncMap["To_Identity_Raw27"] = reflect.ValueOf(To_Identity_Raw27)
	FuncMap["To_Identity_Raw28"] = reflect.ValueOf(To_Identity_Raw28)
	FuncMap["To_Identity_Raw29"] = reflect.ValueOf(To_Identity_Raw29)
	FuncMap["To_Identity_Raw3"] = reflect.ValueOf(To_Identity_Raw3)
	FuncMap["To_Identity_Raw30"] = reflect.ValueOf(To_Identity_Raw30)
	FuncMap["To_Identity_Raw31"] = reflect.ValueOf(To_Identity_Raw31)
	FuncMap["To_Identity_Raw32"] = reflect.ValueOf(To_Identity_Raw32)
	FuncMap["To_Identity_Raw4"] = reflect.ValueOf(To_Identity_Raw4)
	FuncMap["To_Identity_Raw5"] = reflect.ValueOf(To_Identity_Raw5)
	FuncMap["To_Identity_Raw6"] = reflect.ValueOf(To_Identity_Raw6)
	FuncMap["To_Identity_Raw7"] = reflect.ValueOf(To_Identity_Raw7)
	FuncMap["To_Identity_Raw8"] = reflect.ValueOf(To_Identity_Raw8)
	FuncMap["To_Identity_Raw9"] = reflect.ValueOf(To_Identity_Raw9)
	FuncMap["To_Identity_Reasonable"] = reflect.ValueOf(To_Identity_Reasonable)
	FuncMap["To_Identity_RemoveSubCall"] = reflect.ValueOf(To_Identity_RemoveSubCall)
	FuncMap["To_Identity_RenameSubCall"] = reflect.ValueOf(To_Identity_RenameSubCall)
	FuncMap["To_Identity_RequestJudgementCall"] = reflect.ValueOf(To_Identity_RequestJudgementCall)
	FuncMap["To_Identity_Riot"] = reflect.ValueOf(To_Identity_Riot)
	FuncMap["To_Identity_SetAccountIdCall"] = reflect.ValueOf(To_Identity_SetAccountIdCall)
	FuncMap["To_Identity_SetFeeCall"] = reflect.ValueOf(To_Identity_SetFeeCall)
	FuncMap["To_Identity_SetFieldsCall"] = reflect.ValueOf(To_Identity_SetFieldsCall)
	FuncMap["To_Identity_SetIdentityCall"] = reflect.ValueOf(To_Identity_SetIdentityCall)
	FuncMap["To_Identity_SetSubsCall"] = reflect.ValueOf(To_Identity_SetSubsCall)
	FuncMap["To_Identity_Sha256"] = reflect.ValueOf(To_Identity_Sha256)
	FuncMap["To_Identity_ShaThree256"] = reflect.ValueOf(To_Identity_ShaThree256)
	FuncMap["To_Identity_Sub"] = reflect.ValueOf(To_Identity_Sub)
	FuncMap["To_Identity_Target"] = reflect.ValueOf(To_Identity_Target)
	FuncMap["To_Identity_TupleNull"] = reflect.ValueOf(To_Identity_TupleNull)
	FuncMap["To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData"] = reflect.ValueOf(To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData)
	FuncMap["To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData"] = reflect.ValueOf(To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData)
	FuncMap["To_Identity_Twitter"] = reflect.ValueOf(To_Identity_Twitter)
	FuncMap["To_Identity_Unknown"] = reflect.ValueOf(To_Identity_Unknown)
	FuncMap["To_Identity_Value0"] = reflect.ValueOf(To_Identity_Value0)
	FuncMap["To_Identity_Value1"] = reflect.ValueOf(To_Identity_Value1)
	FuncMap["To_Identity_Web"] = reflect.ValueOf(To_Identity_Web)
	FuncMap["To_ImOnlinePallet"] = reflect.ValueOf(To_ImOnlinePallet)
	FuncMap["To_ImOnline_HeartbeatCall"] = reflect.ValueOf(To_ImOnline_HeartbeatCall)
	FuncMap["To_ImOnline_PalletImOnlineHeartbeat"] = reflect.ValueOf(To_ImOnline_PalletImOnlineHeartbeat)
	FuncMap["To_ImOnline_PalletImOnlineSr25519AppSr25519Public"] = reflect.ValueOf(To_ImOnline_PalletImOnlineSr25519AppSr25519Public)
	FuncMap["To_ImOnline_PalletImOnlineSr25519AppSr25519Signature"] = reflect.ValueOf(To_ImOnline_PalletImOnlineSr25519AppSr25519Signature)
	FuncMap["To_MultisigPallet"] = reflect.ValueOf(To_MultisigPallet)
	FuncMap["To_Multisig_ApproveAsMultiCall"] = reflect.ValueOf(To_Multisig_ApproveAsMultiCall)
	FuncMap["To_Multisig_AsMultiCall"] = reflect.ValueOf(To_Multisig_AsMultiCall)
	FuncMap["To_Multisig_AsMultiThreshold1Call"] = reflect.ValueOf(To_Multisig_AsMultiThreshold1Call)
	FuncMap["To_Multisig_CancelAsMultiCall"] = reflect.ValueOf(To_Multisig_CancelAsMultiCall)
	FuncMap["To_Multisig_PalletMultisigTimepoint"] = reflect.ValueOf(To_Multisig_PalletMultisigTimepoint)
	FuncMap["To_NominationPoolsPallet"] = reflect.ValueOf(To_NominationPoolsPallet)
	FuncMap["To_NominationPools_Address20"] = reflect.ValueOf(To_NominationPools_Address20)
	FuncMap["To_NominationPools_Address32"] = reflect.ValueOf(To_NominationPools_Address32)
	FuncMap["To_NominationPools_AdjustPoolDepositCall"] = reflect.ValueOf(To_NominationPools_AdjustPoolDepositCall)
	FuncMap["To_NominationPools_Blocked"] = reflect.ValueOf(To_NominationPools_Blocked)
	FuncMap["To_NominationPools_BondExtraCall"] = reflect.ValueOf(To_NominationPools_BondExtraCall)
	FuncMap["To_NominationPools_BondExtraOtherCall"] = reflect.ValueOf(To_NominationPools_BondExtraOtherCall)
	FuncMap["To_NominationPools_Bouncer"] = reflect.ValueOf(To_NominationPools_Bouncer)
	FuncMap["To_NominationPools_ChillCall"] = reflect.ValueOf(To_NominationPools_ChillCall)
	FuncMap["To_NominationPools_ClaimCommissionCall"] = reflect.ValueOf(To_NominationPools_ClaimCommissionCall)
	FuncMap["To_NominationPools_ClaimPayoutCall"] = reflect.ValueOf(To_NominationPools_ClaimPayoutCall)
	FuncMap["To_NominationPools_ClaimPayoutOtherCall"] = reflect.ValueOf(To_NominationPools_ClaimPayoutOtherCall)
	FuncMap["To_NominationPools_CreateCall"] = reflect.ValueOf(To_NominationPools_CreateCall)
	FuncMap["To_NominationPools_CreateWithPoolIdCall"] = reflect.ValueOf(To_NominationPools_CreateWithPoolIdCall)
	FuncMap["To_NominationPools_Destroying"] = reflect.ValueOf(To_NominationPools_Destroying)
	FuncMap["To_NominationPools_Extra"] = reflect.ValueOf(To_NominationPools_Extra)
	FuncMap["To_NominationPools_FreeBalance"] = reflect.ValueOf(To_NominationPools_FreeBalance)
	FuncMap["To_NominationPools_GlobalMaxCommission"] = reflect.ValueOf(To_NominationPools_GlobalMaxCommission)
	FuncMap["To_NominationPools_Id"] = reflect.ValueOf(To_NominationPools_Id)
	FuncMap["To_NominationPools_Index"] = reflect.ValueOf(To_NominationPools_Index)
	FuncMap["To_NominationPools_JoinCall"] = reflect.ValueOf(To_NominationPools_JoinCall)
	FuncMap["To_NominationPools_MaxMembers"] = reflect.ValueOf(To_NominationPools_MaxMembers)
	FuncMap["To_NominationPools_MaxMembersPerPool"] = reflect.ValueOf(To_NominationPools_MaxMembersPerPool)
	FuncMap["To_NominationPools_MaxPools"] = reflect.ValueOf(To_NominationPools_MaxPools)
	FuncMap["To_NominationPools_Member"] = reflect.ValueOf(To_NominationPools_Member)
	FuncMap["To_NominationPools_MemberAccount"] = reflect.ValueOf(To_NominationPools_MemberAccount)
	FuncMap["To_NominationPools_MinCreateBond"] = reflect.ValueOf(To_NominationPools_MinCreateBond)
	FuncMap["To_NominationPools_MinJoinBond"] = reflect.ValueOf(To_NominationPools_MinJoinBond)
	FuncMap["To_NominationPools_NewBouncer"] = reflect.ValueOf(To_NominationPools_NewBouncer)
	FuncMap["To_NominationPools_NewNominator"] = reflect.ValueOf(To_NominationPools_NewNominator)
	FuncMap["To_NominationPools_NewRoot"] = reflect.ValueOf(To_NominationPools_NewRoot)
	FuncMap["To_NominationPools_NominateCall"] = reflect.ValueOf(To_NominationPools_NominateCall)
	FuncMap["To_NominationPools_Nominator"] = reflect.ValueOf(To_NominationPools_Nominator)
	FuncMap["To_NominationPools_Noop"] = reflect.ValueOf(To_NominationPools_Noop)
	FuncMap["To_NominationPools_Open"] = reflect.ValueOf(To_NominationPools_Open)
	FuncMap["To_NominationPools_PalletNominationPoolsCommissionChangeRate"] = reflect.ValueOf(To_NominationPools_PalletNominationPoolsCommissionChangeRate)
	FuncMap["To_NominationPools_Permission"] = reflect.ValueOf(To_NominationPools_Permission)
	FuncMap["To_NominationPools_Permissioned"] = reflect.ValueOf(To_NominationPools_Permissioned)
	FuncMap["To_NominationPools_PermissionlessAll"] = reflect.ValueOf(To_NominationPools_PermissionlessAll)
	FuncMap["To_NominationPools_PermissionlessCompound"] = reflect.ValueOf(To_NominationPools_PermissionlessCompound)
	FuncMap["To_NominationPools_PermissionlessWithdraw"] = reflect.ValueOf(To_NominationPools_PermissionlessWithdraw)
	FuncMap["To_NominationPools_PoolWithdrawUnbondedCall"] = reflect.ValueOf(To_NominationPools_PoolWithdrawUnbondedCall)
	FuncMap["To_NominationPools_Raw"] = reflect.ValueOf(To_NominationPools_Raw)
	FuncMap["To_NominationPools_Remove"] = reflect.ValueOf(To_NominationPools_Remove)
	FuncMap["To_NominationPools_Rewards"] = reflect.ValueOf(To_NominationPools_Rewards)
	FuncMap["To_NominationPools_Root"] = reflect.ValueOf(To_NominationPools_Root)
	FuncMap["To_NominationPools_Set"] = reflect.ValueOf(To_NominationPools_Set)
	FuncMap["To_NominationPools_SetClaimPermissionCall"] = reflect.ValueOf(To_NominationPools_SetClaimPermissionCall)
	FuncMap["To_NominationPools_SetCommissionCall"] = reflect.ValueOf(To_NominationPools_SetCommissionCall)
	FuncMap["To_NominationPools_SetCommissionChangeRateCall"] = reflect.ValueOf(To_NominationPools_SetCommissionChangeRateCall)
	FuncMap["To_NominationPools_SetCommissionMaxCall"] = reflect.ValueOf(To_NominationPools_SetCommissionMaxCall)
	FuncMap["To_NominationPools_SetConfigsCall"] = reflect.ValueOf(To_NominationPools_SetConfigsCall)
	FuncMap["To_NominationPools_SetMetadataCall"] = reflect.ValueOf(To_NominationPools_SetMetadataCall)
	FuncMap["To_NominationPools_SetStateCall"] = reflect.ValueOf(To_NominationPools_SetStateCall)
	FuncMap["To_NominationPools_State"] = reflect.ValueOf(To_NominationPools_State)
	FuncMap["To_NominationPools_TupleNull"] = reflect.ValueOf(To_NominationPools_TupleNull)
	FuncMap["To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32"] = reflect.ValueOf(To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32)
	FuncMap["To_NominationPools_UnbondCall"] = reflect.ValueOf(To_NominationPools_UnbondCall)
	FuncMap["To_NominationPools_UpdateRolesCall"] = reflect.ValueOf(To_NominationPools_UpdateRolesCall)
	FuncMap["To_NominationPools_WithdrawUnbondedCall"] = reflect.ValueOf(To_NominationPools_WithdrawUnbondedCall)
	FuncMap["To_None"] = reflect.ValueOf(To_None)
	FuncMap["To_OffencesPallet"] = reflect.ValueOf(To_OffencesPallet)
	FuncMap["To_OriginsPallet"] = reflect.ValueOf(To_OriginsPallet)
	FuncMap["To_PreimagePallet"] = reflect.ValueOf(To_PreimagePallet)
	FuncMap["To_Preimage_EnsureUpdatedCall"] = reflect.ValueOf(To_Preimage_EnsureUpdatedCall)
	FuncMap["To_Preimage_NotePreimageCall"] = reflect.ValueOf(To_Preimage_NotePreimageCall)
	FuncMap["To_Preimage_RequestPreimageCall"] = reflect.ValueOf(To_Preimage_RequestPreimageCall)
	FuncMap["To_Preimage_UnnotePreimageCall"] = reflect.ValueOf(To_Preimage_UnnotePreimageCall)
	FuncMap["To_Preimage_UnrequestPreimageCall"] = reflect.ValueOf(To_Preimage_UnrequestPreimageCall)
	FuncMap["To_PrimaryAndSecondaryPlainSlots"] = reflect.ValueOf(To_PrimaryAndSecondaryPlainSlots)
	FuncMap["To_PrimaryAndSecondaryVrfSlots"] = reflect.ValueOf(To_PrimaryAndSecondaryVrfSlots)
	FuncMap["To_PrimarySlots"] = reflect.ValueOf(To_PrimarySlots)
	FuncMap["To_PrimitiveTypesH256"] = reflect.ValueOf(To_PrimitiveTypesH256)
	FuncMap["To_ProxyPallet"] = reflect.ValueOf(To_ProxyPallet)
	FuncMap["To_Proxy_AddProxyCall"] = reflect.ValueOf(To_Proxy_AddProxyCall)
	FuncMap["To_Proxy_Address20"] = reflect.ValueOf(To_Proxy_Address20)
	FuncMap["To_Proxy_Address32"] = reflect.ValueOf(To_Proxy_Address32)
	FuncMap["To_Proxy_AnnounceCall"] = reflect.ValueOf(To_Proxy_AnnounceCall)
	FuncMap["To_Proxy_Any"] = reflect.ValueOf(To_Proxy_Any)
	FuncMap["To_Proxy_CancelProxy"] = reflect.ValueOf(To_Proxy_CancelProxy)
	FuncMap["To_Proxy_CreatePureCall"] = reflect.ValueOf(To_Proxy_CreatePureCall)
	FuncMap["To_Proxy_Delegate"] = reflect.ValueOf(To_Proxy_Delegate)
	FuncMap["To_Proxy_ForceProxyType"] = reflect.ValueOf(To_Proxy_ForceProxyType)
	FuncMap["To_Proxy_Governance"] = reflect.ValueOf(To_Proxy_Governance)
	FuncMap["To_Proxy_Id"] = reflect.ValueOf(To_Proxy_Id)
	FuncMap["To_Proxy_IdentityJudgement"] = reflect.ValueOf(To_Proxy_IdentityJudgement)
	FuncMap["To_Proxy_Index"] = reflect.ValueOf(To_Proxy_Index)
	FuncMap["To_Proxy_KillPureCall"] = reflect.ValueOf(To_Proxy_KillPureCall)
	FuncMap["To_Proxy_NonTransfer"] = reflect.ValueOf(To_Proxy_NonTransfer)
	FuncMap["To_Proxy_ProxyAnnouncedCall"] = reflect.ValueOf(To_Proxy_ProxyAnnouncedCall)
	FuncMap["To_Proxy_ProxyCall"] = reflect.ValueOf(To_Proxy_ProxyCall)
	FuncMap["To_Proxy_ProxyType"] = reflect.ValueOf(To_Proxy_ProxyType)
	FuncMap["To_Proxy_Raw"] = reflect.ValueOf(To_Proxy_Raw)
	FuncMap["To_Proxy_Real"] = reflect.ValueOf(To_Proxy_Real)
	FuncMap["To_Proxy_RejectAnnouncementCall"] = reflect.ValueOf(To_Proxy_RejectAnnouncementCall)
	FuncMap["To_Proxy_RemoveAnnouncementCall"] = reflect.ValueOf(To_Proxy_RemoveAnnouncementCall)
	FuncMap["To_Proxy_RemoveProxiesCall"] = reflect.ValueOf(To_Proxy_RemoveProxiesCall)
	FuncMap["To_Proxy_RemoveProxyCall"] = reflect.ValueOf(To_Proxy_RemoveProxyCall)
	FuncMap["To_Proxy_Spawner"] = reflect.ValueOf(To_Proxy_Spawner)
	FuncMap["To_Proxy_Staking"] = reflect.ValueOf(To_Proxy_Staking)
	FuncMap["To_Proxy_TupleNull"] = reflect.ValueOf(To_Proxy_TupleNull)
	FuncMap["To_ReferendaPallet"] = reflect.ValueOf(To_ReferendaPallet)
	FuncMap["To_Referenda_After"] = reflect.ValueOf(To_Referenda_After)
	FuncMap["To_Referenda_At"] = reflect.ValueOf(To_Referenda_At)
	FuncMap["To_Referenda_CancelCall"] = reflect.ValueOf(To_Referenda_CancelCall)
	FuncMap["To_Referenda_EnactmentMoment"] = reflect.ValueOf(To_Referenda_EnactmentMoment)
	FuncMap["To_Referenda_Inline"] = reflect.ValueOf(To_Referenda_Inline)
	FuncMap["To_Referenda_KillCall"] = reflect.ValueOf(To_Referenda_KillCall)
	FuncMap["To_Referenda_Legacy"] = reflect.ValueOf(To_Referenda_Legacy)
	FuncMap["To_Referenda_Lookup"] = reflect.ValueOf(To_Referenda_Lookup)
	FuncMap["To_Referenda_NudgeReferendumCall"] = reflect.ValueOf(To_Referenda_NudgeReferendumCall)
	FuncMap["To_Referenda_OneFewerDecidingCall"] = reflect.ValueOf(To_Referenda_OneFewerDecidingCall)
	FuncMap["To_Referenda_Origins"] = reflect.ValueOf(To_Referenda_Origins)
	FuncMap["To_Referenda_PlaceDecisionDepositCall"] = reflect.ValueOf(To_Referenda_PlaceDecisionDepositCall)
	FuncMap["To_Referenda_Proposal"] = reflect.ValueOf(To_Referenda_Proposal)
	FuncMap["To_Referenda_ProposalOrigin"] = reflect.ValueOf(To_Referenda_ProposalOrigin)
	FuncMap["To_Referenda_RefundDecisionDepositCall"] = reflect.ValueOf(To_Referenda_RefundDecisionDepositCall)
	FuncMap["To_Referenda_RefundSubmissionDepositCall"] = reflect.ValueOf(To_Referenda_RefundSubmissionDepositCall)
	FuncMap["To_Referenda_SetMetadataCall"] = reflect.ValueOf(To_Referenda_SetMetadataCall)
	FuncMap["To_Referenda_SubmitCall"] = reflect.ValueOf(To_Referenda_SubmitCall)
	FuncMap["To_Referenda_System"] = reflect.ValueOf(To_Referenda_System)
	FuncMap["To_Referenda_Void"] = reflect.ValueOf(To_Referenda_Void)
	FuncMap["To_Root"] = reflect.ValueOf(To_Root)
	FuncMap["To_SchedulerPallet"] = reflect.ValueOf(To_SchedulerPallet)
	FuncMap["To_Scheduler_CancelCall"] = reflect.ValueOf(To_Scheduler_CancelCall)
	FuncMap["To_Scheduler_CancelNamedCall"] = reflect.ValueOf(To_Scheduler_CancelNamedCall)
	FuncMap["To_Scheduler_ScheduleAfterCall"] = reflect.ValueOf(To_Scheduler_ScheduleAfterCall)
	FuncMap["To_Scheduler_ScheduleCall"] = reflect.ValueOf(To_Scheduler_ScheduleCall)
	FuncMap["To_Scheduler_ScheduleNamedAfterCall"] = reflect.ValueOf(To_Scheduler_ScheduleNamedAfterCall)
	FuncMap["To_Scheduler_ScheduleNamedCall"] = reflect.ValueOf(To_Scheduler_ScheduleNamedCall)
	FuncMap["To_Scheduler_TupleUint32Uint32"] = reflect.ValueOf(To_Scheduler_TupleUint32Uint32)
	FuncMap["To_SessionPallet"] = reflect.ValueOf(To_SessionPallet)
	FuncMap["To_Session_PurgeKeysCall"] = reflect.ValueOf(To_Session_PurgeKeysCall)
	FuncMap["To_Session_SetKeysCall"] = reflect.ValueOf(To_Session_SetKeysCall)
	FuncMap["To_Signed"] = reflect.ValueOf(To_Signed)
	FuncMap["To_SpArithmeticPerThingsPerU16"] = reflect.ValueOf(To_SpArithmeticPerThingsPerU16)
	FuncMap["To_SpArithmeticPerThingsPerbill"] = reflect.ValueOf(To_SpArithmeticPerThingsPerbill)
	FuncMap["To_SpArithmeticPerThingsPercent"] = reflect.ValueOf(To_SpArithmeticPerThingsPercent)
	FuncMap["To_SpAuthorityDiscoveryAppPublic"] = reflect.ValueOf(To_SpAuthorityDiscoveryAppPublic)
	FuncMap["To_SpConsensusBabeAppPublic"] = reflect.ValueOf(To_SpConsensusBabeAppPublic)
	FuncMap["To_SpConsensusGrandpaAppPublic"] = reflect.ValueOf(To_SpConsensusGrandpaAppPublic)
	FuncMap["To_SpConsensusGrandpaAppSignature"] = reflect.ValueOf(To_SpConsensusGrandpaAppSignature)
	FuncMap["To_SpConsensusGrandpaEquivocationProof"] = reflect.ValueOf(To_SpConsensusGrandpaEquivocationProof)
	FuncMap["To_SpConsensusSlotsEquivocationProof"] = reflect.ValueOf(To_SpConsensusSlotsEquivocationProof)
	FuncMap["To_SpConsensusSlotsSlot"] = reflect.ValueOf(To_SpConsensusSlotsSlot)
	FuncMap["To_SpCoreCryptoAccountId32"] = reflect.ValueOf(To_SpCoreCryptoAccountId32)
	FuncMap["To_SpCoreEd25519Public"] = reflect.ValueOf(To_SpCoreEd25519Public)
	FuncMap["To_SpCoreEd25519Signature"] = reflect.ValueOf(To_SpCoreEd25519Signature)
	FuncMap["To_SpCoreSr25519Public"] = reflect.ValueOf(To_SpCoreSr25519Public)
	FuncMap["To_SpCoreSr25519Signature"] = reflect.ValueOf(To_SpCoreSr25519Signature)
	FuncMap["To_SpNposElectionsElectionScore"] = reflect.ValueOf(To_SpNposElectionsElectionScore)
	FuncMap["To_SpNposElectionsSupport"] = reflect.ValueOf(To_SpNposElectionsSupport)
	FuncMap["To_SpRuntimeGenericDigestDigest"] = reflect.ValueOf(To_SpRuntimeGenericDigestDigest)
	FuncMap["To_SpRuntimeGenericDigestDigestItem"] = reflect.ValueOf(To_SpRuntimeGenericDigestDigestItem)
	FuncMap["To_SpRuntimeGenericHeaderHeader"] = reflect.ValueOf(To_SpRuntimeGenericHeaderHeader)
	FuncMap["To_SpRuntimeMultiaddressMultiAddress"] = reflect.ValueOf(To_SpRuntimeMultiaddressMultiAddress)
	FuncMap["To_SpSessionMembershipProof"] = reflect.ValueOf(To_SpSessionMembershipProof)
	FuncMap["To_SpWeightsWeightV2Weight"] = reflect.ValueOf(To_SpWeightsWeightV2Weight)
	FuncMap["To_StakingPallet"] = reflect.ValueOf(To_StakingPallet)
	FuncMap["To_StakingRewardsPallet"] = reflect.ValueOf(To_StakingRewardsPallet)
	FuncMap["To_StakingRewards_Address20"] = reflect.ValueOf(To_StakingRewards_Address20)
	FuncMap["To_StakingRewards_Address32"] = reflect.ValueOf(To_StakingRewards_Address32)
	FuncMap["To_StakingRewards_AlignSupplyCall"] = reflect.ValueOf(To_StakingRewards_AlignSupplyCall)
	FuncMap["To_StakingRewards_ForceRefillCall"] = reflect.ValueOf(To_StakingRewards_ForceRefillCall)
	FuncMap["To_StakingRewards_From"] = reflect.ValueOf(To_StakingRewards_From)
	FuncMap["To_StakingRewards_Id"] = reflect.ValueOf(To_StakingRewards_Id)
	FuncMap["To_StakingRewards_Index"] = reflect.ValueOf(To_StakingRewards_Index)
	FuncMap["To_StakingRewards_Raw"] = reflect.ValueOf(To_StakingRewards_Raw)
	FuncMap["To_StakingRewards_RefillCall"] = reflect.ValueOf(To_StakingRewards_RefillCall)
	FuncMap["To_StakingRewards_To"] = reflect.ValueOf(To_StakingRewards_To)
	FuncMap["To_StakingRewards_TupleNull"] = reflect.ValueOf(To_StakingRewards_TupleNull)
	FuncMap["To_StakingRewards_WithdrawCall"] = reflect.ValueOf(To_StakingRewards_WithdrawCall)
	FuncMap["To_Staking_Account"] = reflect.ValueOf(To_Staking_Account)
	FuncMap["To_Staking_Address20"] = reflect.ValueOf(To_Staking_Address20)
	FuncMap["To_Staking_Address32"] = reflect.ValueOf(To_Staking_Address32)
	FuncMap["To_Staking_BondCall"] = reflect.ValueOf(To_Staking_BondCall)
	FuncMap["To_Staking_BondExtraCall"] = reflect.ValueOf(To_Staking_BondExtraCall)
	FuncMap["To_Staking_CancelDeferredSlashCall"] = reflect.ValueOf(To_Staking_CancelDeferredSlashCall)
	FuncMap["To_Staking_ChillCall"] = reflect.ValueOf(To_Staking_ChillCall)
	FuncMap["To_Staking_ChillOtherCall"] = reflect.ValueOf(To_Staking_ChillOtherCall)
	FuncMap["To_Staking_ChillThreshold"] = reflect.ValueOf(To_Staking_ChillThreshold)
	FuncMap["To_Staking_Controller"] = reflect.ValueOf(To_Staking_Controller)
	FuncMap["To_Staking_ForceApplyMinCommissionCall"] = reflect.ValueOf(To_Staking_ForceApplyMinCommissionCall)
	FuncMap["To_Staking_ForceNewEraAlwaysCall"] = reflect.ValueOf(To_Staking_ForceNewEraAlwaysCall)
	FuncMap["To_Staking_ForceNewEraCall"] = reflect.ValueOf(To_Staking_ForceNewEraCall)
	FuncMap["To_Staking_ForceNoErasCall"] = reflect.ValueOf(To_Staking_ForceNoErasCall)
	FuncMap["To_Staking_ForceUnstakeCall"] = reflect.ValueOf(To_Staking_ForceUnstakeCall)
	FuncMap["To_Staking_Id"] = reflect.ValueOf(To_Staking_Id)
	FuncMap["To_Staking_IncreaseValidatorCountCall"] = reflect.ValueOf(To_Staking_IncreaseValidatorCountCall)
	FuncMap["To_Staking_Index"] = reflect.ValueOf(To_Staking_Index)
	FuncMap["To_Staking_KickCall"] = reflect.ValueOf(To_Staking_KickCall)
	FuncMap["To_Staking_MaxNominatorCount"] = reflect.ValueOf(To_Staking_MaxNominatorCount)
	FuncMap["To_Staking_MaxValidatorCount"] = reflect.ValueOf(To_Staking_MaxValidatorCount)
	FuncMap["To_Staking_MinCommission"] = reflect.ValueOf(To_Staking_MinCommission)
	FuncMap["To_Staking_MinNominatorBond"] = reflect.ValueOf(To_Staking_MinNominatorBond)
	FuncMap["To_Staking_MinValidatorBond"] = reflect.ValueOf(To_Staking_MinValidatorBond)
	FuncMap["To_Staking_NominateCall"] = reflect.ValueOf(To_Staking_NominateCall)
	FuncMap["To_Staking_None"] = reflect.ValueOf(To_Staking_None)
	FuncMap["To_Staking_Noop"] = reflect.ValueOf(To_Staking_Noop)
	FuncMap["To_Staking_PalletStakingValidatorPrefs"] = reflect.ValueOf(To_Staking_PalletStakingValidatorPrefs)
	FuncMap["To_Staking_Payee"] = reflect.ValueOf(To_Staking_Payee)
	FuncMap["To_Staking_PayoutStakersCall"] = reflect.ValueOf(To_Staking_PayoutStakersCall)
	FuncMap["To_Staking_Raw"] = reflect.ValueOf(To_Staking_Raw)
	FuncMap["To_Staking_ReapStashCall"] = reflect.ValueOf(To_Staking_ReapStashCall)
	FuncMap["To_Staking_RebondCall"] = reflect.ValueOf(To_Staking_RebondCall)
	FuncMap["To_Staking_Remove"] = reflect.ValueOf(To_Staking_Remove)
	FuncMap["To_Staking_ScaleValidatorCountCall"] = reflect.ValueOf(To_Staking_ScaleValidatorCountCall)
	FuncMap["To_Staking_Set"] = reflect.ValueOf(To_Staking_Set)
	FuncMap["To_Staking_SetControllerCall"] = reflect.ValueOf(To_Staking_SetControllerCall)
	FuncMap["To_Staking_SetInvulnerablesCall"] = reflect.ValueOf(To_Staking_SetInvulnerablesCall)
	FuncMap["To_Staking_SetMinCommissionCall"] = reflect.ValueOf(To_Staking_SetMinCommissionCall)
	FuncMap["To_Staking_SetPayeeCall"] = reflect.ValueOf(To_Staking_SetPayeeCall)
	FuncMap["To_Staking_SetStakingConfigsCall"] = reflect.ValueOf(To_Staking_SetStakingConfigsCall)
	FuncMap["To_Staking_SetValidatorCountCall"] = reflect.ValueOf(To_Staking_SetValidatorCountCall)
	FuncMap["To_Staking_Staked"] = reflect.ValueOf(To_Staking_Staked)
	FuncMap["To_Staking_Stash"] = reflect.ValueOf(To_Staking_Stash)
	FuncMap["To_Staking_Targets"] = reflect.ValueOf(To_Staking_Targets)
	FuncMap["To_Staking_TupleNull"] = reflect.ValueOf(To_Staking_TupleNull)
	FuncMap["To_Staking_UnbondCall"] = reflect.ValueOf(To_Staking_UnbondCall)
	FuncMap["To_Staking_ValidateCall"] = reflect.ValueOf(To_Staking_ValidateCall)
	FuncMap["To_Staking_Who"] = reflect.ValueOf(To_Staking_Who)
	FuncMap["To_Staking_WithdrawUnbondedCall"] = reflect.ValueOf(To_Staking_WithdrawUnbondedCall)
	FuncMap["To_SystemPallet"] = reflect.ValueOf(To_SystemPallet)
	FuncMap["To_System_KillPrefixCall"] = reflect.ValueOf(To_System_KillPrefixCall)
	FuncMap["To_System_KillStorageCall"] = reflect.ValueOf(To_System_KillStorageCall)
	FuncMap["To_System_RemarkCall"] = reflect.ValueOf(To_System_RemarkCall)
	FuncMap["To_System_RemarkWithEventCall"] = reflect.ValueOf(To_System_RemarkWithEventCall)
	FuncMap["To_System_SetCodeCall"] = reflect.ValueOf(To_System_SetCodeCall)
	FuncMap["To_System_SetCodeWithoutChecksCall"] = reflect.ValueOf(To_System_SetCodeWithoutChecksCall)
	FuncMap["To_System_SetHeapPagesCall"] = reflect.ValueOf(To_System_SetHeapPagesCall)
	FuncMap["To_System_SetStorageCall"] = reflect.ValueOf(To_System_SetStorageCall)
	FuncMap["To_System_SystemKeysList"] = reflect.ValueOf(To_System_SystemKeysList)
	FuncMap["To_System_TupleSystemItemsListSystemItemsList"] = reflect.ValueOf(To_System_TupleSystemItemsListSystemItemsList)
	FuncMap["To_TimestampPallet"] = reflect.ValueOf(To_TimestampPallet)
	FuncMap["To_Timestamp_SetCall"] = reflect.ValueOf(To_Timestamp_SetCall)
	FuncMap["To_TransactionPaymentPallet"] = reflect.ValueOf(To_TransactionPaymentPallet)
	FuncMap["To_TreasuryPallet"] = reflect.ValueOf(To_TreasuryPallet)
	FuncMap["To_Treasury_Address20"] = reflect.ValueOf(To_Treasury_Address20)
	FuncMap["To_Treasury_Address32"] = reflect.ValueOf(To_Treasury_Address32)
	FuncMap["To_Treasury_ApproveProposalCall"] = reflect.ValueOf(To_Treasury_ApproveProposalCall)
	FuncMap["To_Treasury_Beneficiary"] = reflect.ValueOf(To_Treasury_Beneficiary)
	FuncMap["To_Treasury_CheckStatusCall"] = reflect.ValueOf(To_Treasury_CheckStatusCall)
	FuncMap["To_Treasury_Id"] = reflect.ValueOf(To_Treasury_Id)
	FuncMap["To_Treasury_Index"] = reflect.ValueOf(To_Treasury_Index)
	FuncMap["To_Treasury_PayoutCall"] = reflect.ValueOf(To_Treasury_PayoutCall)
	FuncMap["To_Treasury_ProposeSpendCall"] = reflect.ValueOf(To_Treasury_ProposeSpendCall)
	FuncMap["To_Treasury_Raw"] = reflect.ValueOf(To_Treasury_Raw)
	FuncMap["To_Treasury_RejectProposalCall"] = reflect.ValueOf(To_Treasury_RejectProposalCall)
	FuncMap["To_Treasury_RemoveApprovalCall"] = reflect.ValueOf(To_Treasury_RemoveApprovalCall)
	FuncMap["To_Treasury_SpendCall"] = reflect.ValueOf(To_Treasury_SpendCall)
	FuncMap["To_Treasury_SpendLocalCall"] = reflect.ValueOf(To_Treasury_SpendLocalCall)
	FuncMap["To_Treasury_TupleNull"] = reflect.ValueOf(To_Treasury_TupleNull)
	FuncMap["To_Treasury_VoidSpendCall"] = reflect.ValueOf(To_Treasury_VoidSpendCall)
	FuncMap["To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature"] = reflect.ValueOf(To_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature)
	FuncMap["To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature"] = reflect.ValueOf(To_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature)
	FuncMap["To_TupleNull"] = reflect.ValueOf(To_TupleNull)
	FuncMap["To_UtilityPallet"] = reflect.ValueOf(To_UtilityPallet)
	FuncMap["To_Utility_AsDerivativeCall"] = reflect.ValueOf(To_Utility_AsDerivativeCall)
	FuncMap["To_Utility_AsOrigin"] = reflect.ValueOf(To_Utility_AsOrigin)
	FuncMap["To_Utility_BatchAllCall"] = reflect.ValueOf(To_Utility_BatchAllCall)
	FuncMap["To_Utility_BatchCall"] = reflect.ValueOf(To_Utility_BatchCall)
	FuncMap["To_Utility_DispatchAsCall"] = reflect.ValueOf(To_Utility_DispatchAsCall)
	FuncMap["To_Utility_ForceBatchCall"] = reflect.ValueOf(To_Utility_ForceBatchCall)
	FuncMap["To_Utility_Origins"] = reflect.ValueOf(To_Utility_Origins)
	FuncMap["To_Utility_System"] = reflect.ValueOf(To_Utility_System)
	FuncMap["To_Utility_Void"] = reflect.ValueOf(To_Utility_Void)
	FuncMap["To_Utility_WithWeightCall"] = reflect.ValueOf(To_Utility_WithWeightCall)
	FuncMap["To_Value0"] = reflect.ValueOf(To_Value0)
	FuncMap["To_VaraRuntimeNposSolution16"] = reflect.ValueOf(To_VaraRuntimeNposSolution16)
	FuncMap["To_VaraRuntimeRuntimeCall"] = reflect.ValueOf(To_VaraRuntimeRuntimeCall)
	FuncMap["To_VaraRuntimeSessionKeys"] = reflect.ValueOf(To_VaraRuntimeSessionKeys)
	FuncMap["To_VestingPallet"] = reflect.ValueOf(To_VestingPallet)
	FuncMap["To_Vesting_Address20"] = reflect.ValueOf(To_Vesting_Address20)
	FuncMap["To_Vesting_Address32"] = reflect.ValueOf(To_Vesting_Address32)
	FuncMap["To_Vesting_ForceVestedTransferCall"] = reflect.ValueOf(To_Vesting_ForceVestedTransferCall)
	FuncMap["To_Vesting_Id"] = reflect.ValueOf(To_Vesting_Id)
	FuncMap["To_Vesting_Index"] = reflect.ValueOf(To_Vesting_Index)
	FuncMap["To_Vesting_MergeSchedulesCall"] = reflect.ValueOf(To_Vesting_MergeSchedulesCall)
	FuncMap["To_Vesting_PalletVestingVestingInfoVestingInfo"] = reflect.ValueOf(To_Vesting_PalletVestingVestingInfoVestingInfo)
	FuncMap["To_Vesting_Raw"] = reflect.ValueOf(To_Vesting_Raw)
	FuncMap["To_Vesting_Source"] = reflect.ValueOf(To_Vesting_Source)
	FuncMap["To_Vesting_Target"] = reflect.ValueOf(To_Vesting_Target)
	FuncMap["To_Vesting_TupleNull"] = reflect.ValueOf(To_Vesting_TupleNull)
	FuncMap["To_Vesting_VestCall"] = reflect.ValueOf(To_Vesting_VestCall)
	FuncMap["To_Vesting_VestOtherCall"] = reflect.ValueOf(To_Vesting_VestOtherCall)
	FuncMap["To_Vesting_VestedTransferCall"] = reflect.ValueOf(To_Vesting_VestedTransferCall)
	FuncMap["To_WhitelistPallet"] = reflect.ValueOf(To_WhitelistPallet)
	FuncMap["To_Whitelist_DispatchWhitelistedCallCall"] = reflect.ValueOf(To_Whitelist_DispatchWhitelistedCallCall)
	FuncMap["To_Whitelist_DispatchWhitelistedCallWithPreimageCall"] = reflect.ValueOf(To_Whitelist_DispatchWhitelistedCallWithPreimageCall)
	FuncMap["To_Whitelist_RemoveWhitelistedCallCall"] = reflect.ValueOf(To_Whitelist_RemoveWhitelistedCallCall)
	FuncMap["To_Whitelist_WhitelistCallCall"] = reflect.ValueOf(To_Whitelist_WhitelistCallCall)
}
