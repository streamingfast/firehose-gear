package gen_types

import (
	"fmt"
	"reflect"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func To_AuthorityDiscoveryPallet(in any) *pbgear.AuthorityDiscoveryPallet {
	out := &pbgear.AuthorityDiscoveryPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_AuthorshipPallet(in any) *pbgear.AuthorshipPallet {
	out := &pbgear.AuthorshipPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_BTreeSet(in any) *pbgear.BTreeSet {
	out := &pbgear.BTreeSet{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Value0
	out.Value0 = To_Repeated_BTreeSet_value0(v.ValueAt(0))

	return out //from message
}

func To_Repeated_BTreeSet_value0(in any) []*pbgear.GprimitivesActorId {
	items := in.([]interface{})

	var out []*pbgear.GprimitivesActorId
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_GprimitivesActorId(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_BabePallet(in any) *pbgear.BabePallet {
	out := &pbgear.BabePallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_BabePallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BabePallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BabePallet_CallReportEquivocationCall{
			CallReportEquivocationCall: To_Babe_ReportEquivocationCall(param),
		}
	case 1:
		return &pbgear.BabePallet_CallReportEquivocationUnsignedCall{
			CallReportEquivocationUnsignedCall: To_Babe_ReportEquivocationUnsignedCall(param),
		}
	case 2:
		return &pbgear.BabePallet_CallPlanConfigChangeCall{
			CallPlanConfigChangeCall: To_Babe_PlanConfigChangeCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Babe_AllowedSlots(in any) *pbgear.Babe_AllowedSlots {
	out := &pbgear.Babe_AllowedSlots{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field AllowedSlots
	v0 := To_OneOf_Babe_AllowedSlots_allowed_slots(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("AllowedSlots").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Babe_AllowedSlots_allowed_slots(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimarySlots{
			AllowedSlotsPrimarySlots: To_Babe_PrimarySlots(param),
		}
	case 1:
		return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimaryAndSecondaryPlainSlots{
			AllowedSlotsPrimaryAndSecondaryPlainSlots: To_Babe_PrimaryAndSecondaryPlainSlots(param),
		}
	case 2:
		return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimaryAndSecondaryVrfSlots{
			AllowedSlotsPrimaryAndSecondaryVrfSlots: To_Babe_PrimaryAndSecondaryVrfSlots(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Babe_BabeTrieNodesList(in any) *pbgear.Babe_BabeTrieNodesList {
	out := &pbgear.Babe_BabeTrieNodesList{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field TrieNodes
	out.TrieNodes = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Babe_Config(in any) *pbgear.Babe_Config {
	out := &pbgear.Babe_Config{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Config
	v0 := To_OneOf_Babe_Config_config(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Config").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Babe_Config_config(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 1:
		return &pbgear.Babe_Config_ConfigV1{
			ConfigV1: To_Babe_V1(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Babe_Consensus(in any) *pbgear.Babe_Consensus {
	out := &pbgear.Babe_Consensus{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Babe_Logs(in any) *pbgear.Babe_Logs {
	out := &pbgear.Babe_Logs{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Logs
	v0 := To_OneOf_Babe_Logs_logs(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Logs").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Babe_Logs_logs(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 6:
		return &pbgear.Babe_Logs_LogsPreRuntime{
			LogsPreRuntime: To_Babe_PreRuntime(param),
		}
	case 4:
		return &pbgear.Babe_Logs_LogsConsensus{
			LogsConsensus: To_Babe_Consensus(param),
		}
	case 5:
		return &pbgear.Babe_Logs_LogsSeal{
			LogsSeal: To_Babe_Seal(param),
		}
	case 0:
		return &pbgear.Babe_Logs_LogsOther{
			LogsOther: To_Babe_Other(param),
		}
	case 8:
		return &pbgear.Babe_Logs_LogsRuntimeEnvironmentUpdated{
			LogsRuntimeEnvironmentUpdated: To_Babe_RuntimeEnvironmentUpdated(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Babe_Other(in any) *pbgear.Babe_Other {
	out := &pbgear.Babe_Other{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Babe_PlanConfigChangeCall(in any) *pbgear.Babe_PlanConfigChangeCall {
	out := &pbgear.Babe_PlanConfigChangeCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Config
	v0 := To_OneOf_Babe_PlanConfigChangeCall_config(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Config").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Babe_PlanConfigChangeCall_config(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 1:
		return &pbgear.Babe_PlanConfigChangeCall_ConfigV1{
			ConfigV1: To_Babe_V1(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Babe_PreRuntime(in any) *pbgear.Babe_PreRuntime {
	out := &pbgear.Babe_PreRuntime{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Babe_PrimaryAndSecondaryPlainSlots(in any) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
	out := &pbgear.Babe_PrimaryAndSecondaryPlainSlots{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Babe_PrimaryAndSecondaryVrfSlots(in any) *pbgear.Babe_PrimaryAndSecondaryVrfSlots {
	out := &pbgear.Babe_PrimaryAndSecondaryVrfSlots{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Babe_PrimarySlots(in any) *pbgear.Babe_PrimarySlots {
	out := &pbgear.Babe_PrimarySlots{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Babe_ReportEquivocationCall(in any) *pbgear.Babe_ReportEquivocationCall {
	out := &pbgear.Babe_ReportEquivocationCall{}
	v := in.(registry.Valuable)
	_ = v

	// field EquivocationProof To_SpConsensusSlotsEquivocationProof(w)
	out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(v.ValueAt(0))
	// field KeyOwnerProof To_SpSessionMembershipProof(w)
	out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

	return out //from message
}

func To_Babe_ReportEquivocationUnsignedCall(in any) *pbgear.Babe_ReportEquivocationUnsignedCall {
	out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
	v := in.(registry.Valuable)
	_ = v

	// field EquivocationProof To_SpConsensusSlotsEquivocationProof(w)
	out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(v.ValueAt(0))
	// field KeyOwnerProof To_SpSessionMembershipProof(w)
	out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

	return out //from message
}

func To_Babe_RuntimeEnvironmentUpdated(in any) *pbgear.Babe_RuntimeEnvironmentUpdated {
	out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Babe_Seal(in any) *pbgear.Babe_Seal {
	out := &pbgear.Babe_Seal{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Babe_SpConsensusBabeAppPublic(in any) *pbgear.Babe_SpConsensusBabeAppPublic {
	out := &pbgear.Babe_SpConsensusBabeAppPublic{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreSr25519Public(w)
	out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

	return out //from message
}

func To_Babe_TupleUint64Uint64(in any) *pbgear.Babe_TupleUint64Uint64 {
	out := &pbgear.Babe_TupleUint64Uint64{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint64(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_uint64(v.ValueAt(1))

	return out //from message
}

func To_Babe_V1(in any) *pbgear.Babe_V1 {
	out := &pbgear.Babe_V1{}
	v := in.(registry.Valuable)
	_ = v

	// field C To_Babe_TupleUint64Uint64(w)
	out.C = To_Babe_TupleUint64Uint64(v.ValueAt(0))
	// oneOf field AllowedSlots
	v1 := To_OneOf_Babe_V1_allowed_slots(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("AllowedSlots").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Babe_V1_allowed_slots(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Babe_V1_AllowedSlotsPrimarySlots{
			AllowedSlotsPrimarySlots: To_Babe_PrimarySlots(param),
		}
	case 1:
		return &pbgear.Babe_V1_AllowedSlotsPrimaryAndSecondaryPlainSlots{
			AllowedSlotsPrimaryAndSecondaryPlainSlots: To_Babe_PrimaryAndSecondaryPlainSlots(param),
		}
	case 2:
		return &pbgear.Babe_V1_AllowedSlotsPrimaryAndSecondaryVrfSlots{
			AllowedSlotsPrimaryAndSecondaryVrfSlots: To_Babe_PrimaryAndSecondaryVrfSlots(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsListPallet(in any) *pbgear.BagsListPallet {
	out := &pbgear.BagsListPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_BagsListPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsListPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsListPallet_CallRebagCall{
			CallRebagCall: To_BagsList_RebagCall(param),
		}
	case 1:
		return &pbgear.BagsListPallet_CallPutInFrontOfCall{
			CallPutInFrontOfCall: To_BagsList_PutInFrontOfCall(param),
		}
	case 2:
		return &pbgear.BagsListPallet_CallPutInFrontOfOtherCall{
			CallPutInFrontOfOtherCall: To_BagsList_PutInFrontOfOtherCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_Address20(in any) *pbgear.BagsList_Address20 {
	out := &pbgear.BagsList_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_BagsList_Address32(in any) *pbgear.BagsList_Address32 {
	out := &pbgear.BagsList_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_BagsList_Dislocated(in any) *pbgear.BagsList_Dislocated {
	out := &pbgear.BagsList_Dislocated{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dislocated
	v0 := To_OneOf_BagsList_Dislocated_dislocated(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dislocated").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsList_Dislocated_dislocated(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_Dislocated_DislocatedId{
			DislocatedId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_Dislocated_DislocatedIndex{
			DislocatedIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_Dislocated_DislocatedRaw{
			DislocatedRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_Dislocated_DislocatedAddress32{
			DislocatedAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_Dislocated_DislocatedAddress20{
			DislocatedAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_Heavier(in any) *pbgear.BagsList_Heavier {
	out := &pbgear.BagsList_Heavier{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Heavier
	v0 := To_OneOf_BagsList_Heavier_heavier(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Heavier").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsList_Heavier_heavier(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_Heavier_HeavierId{
			HeavierId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_Heavier_HeavierIndex{
			HeavierIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_Heavier_HeavierRaw{
			HeavierRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_Heavier_HeavierAddress32{
			HeavierAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_Heavier_HeavierAddress20{
			HeavierAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_Id(in any) *pbgear.BagsList_Id {
	out := &pbgear.BagsList_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_BagsList_Index(in any) *pbgear.BagsList_Index {
	out := &pbgear.BagsList_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_BagsList_TupleNull(w)
	out.Value0 = To_BagsList_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_BagsList_Lighter(in any) *pbgear.BagsList_Lighter {
	out := &pbgear.BagsList_Lighter{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Lighter
	v0 := To_OneOf_BagsList_Lighter_lighter(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsList_Lighter_lighter(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_Lighter_LighterId{
			LighterId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_Lighter_LighterIndex{
			LighterIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_Lighter_LighterRaw{
			LighterRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_Lighter_LighterAddress32{
			LighterAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_Lighter_LighterAddress20{
			LighterAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_PutInFrontOfCall(in any) *pbgear.BagsList_PutInFrontOfCall {
	out := &pbgear.BagsList_PutInFrontOfCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Lighter
	v0 := To_OneOf_BagsList_PutInFrontOfCall_lighter(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsList_PutInFrontOfCall_lighter(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_PutInFrontOfCall_LighterId{
			LighterId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_PutInFrontOfCall_LighterIndex{
			LighterIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_PutInFrontOfCall_LighterRaw{
			LighterRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_PutInFrontOfCall_LighterAddress32{
			LighterAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_PutInFrontOfCall_LighterAddress20{
			LighterAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_PutInFrontOfOtherCall(in any) *pbgear.BagsList_PutInFrontOfOtherCall {
	out := &pbgear.BagsList_PutInFrontOfOtherCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Heavier
	v0 := To_OneOf_BagsList_PutInFrontOfOtherCall_heavier(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Heavier").Set(reflect.ValueOf(v0))
	// oneOf field Lighter
	v1 := To_OneOf_BagsList_PutInFrontOfOtherCall_lighter(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_BagsList_PutInFrontOfOtherCall_heavier(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierId{
			HeavierId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierIndex{
			HeavierIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierRaw{
			HeavierRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierAddress32{
			HeavierAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierAddress20{
			HeavierAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_BagsList_PutInFrontOfOtherCall_lighter(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_PutInFrontOfOtherCall_LighterId{
			LighterId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_PutInFrontOfOtherCall_LighterIndex{
			LighterIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_PutInFrontOfOtherCall_LighterRaw{
			LighterRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_PutInFrontOfOtherCall_LighterAddress32{
			LighterAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_PutInFrontOfOtherCall_LighterAddress20{
			LighterAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_Raw(in any) *pbgear.BagsList_Raw {
	out := &pbgear.BagsList_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_BagsList_RebagCall(in any) *pbgear.BagsList_RebagCall {
	out := &pbgear.BagsList_RebagCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dislocated
	v0 := To_OneOf_BagsList_RebagCall_dislocated(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dislocated").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BagsList_RebagCall_dislocated(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BagsList_RebagCall_DislocatedId{
			DislocatedId: To_BagsList_Id(param),
		}
	case 1:
		return &pbgear.BagsList_RebagCall_DislocatedIndex{
			DislocatedIndex: To_BagsList_Index(param),
		}
	case 2:
		return &pbgear.BagsList_RebagCall_DislocatedRaw{
			DislocatedRaw: To_BagsList_Raw(param),
		}
	case 3:
		return &pbgear.BagsList_RebagCall_DislocatedAddress32{
			DislocatedAddress32: To_BagsList_Address32(param),
		}
	case 4:
		return &pbgear.BagsList_RebagCall_DislocatedAddress20{
			DislocatedAddress20: To_BagsList_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BagsList_TupleNull(in any) *pbgear.BagsList_TupleNull {
	out := &pbgear.BagsList_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_BalancesPallet(in any) *pbgear.BalancesPallet {
	out := &pbgear.BalancesPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_BalancesPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BalancesPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BalancesPallet_CallTransferAllowDeathCall{
			CallTransferAllowDeathCall: To_Balances_TransferAllowDeathCall(param),
		}
	case 1:
		return &pbgear.BalancesPallet_CallForceTransferCall{
			CallForceTransferCall: To_Balances_ForceTransferCall(param),
		}
	case 2:
		return &pbgear.BalancesPallet_CallTransferKeepAliveCall{
			CallTransferKeepAliveCall: To_Balances_TransferKeepAliveCall(param),
		}
	case 3:
		return &pbgear.BalancesPallet_CallTransferAllCall{
			CallTransferAllCall: To_Balances_TransferAllCall(param),
		}
	case 4:
		return &pbgear.BalancesPallet_CallForceUnreserveCall{
			CallForceUnreserveCall: To_Balances_ForceUnreserveCall(param),
		}
	case 5:
		return &pbgear.BalancesPallet_CallUpgradeAccountsCall{
			CallUpgradeAccountsCall: To_Balances_UpgradeAccountsCall(param),
		}
	case 6:
		return &pbgear.BalancesPallet_CallForceSetBalanceCall{
			CallForceSetBalanceCall: To_Balances_ForceSetBalanceCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_Address20(in any) *pbgear.Balances_Address20 {
	out := &pbgear.Balances_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Balances_Address32(in any) *pbgear.Balances_Address32 {
	out := &pbgear.Balances_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Balances_Dest(in any) *pbgear.Balances_Dest {
	out := &pbgear.Balances_Dest{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dest
	v0 := To_OneOf_Balances_Dest_dest(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Balances_Dest_dest(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_Dest_DestId{
			DestId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_Dest_DestIndex{
			DestIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_Dest_DestRaw{
			DestRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_Dest_DestAddress32{
			DestAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_Dest_DestAddress20{
			DestAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_ForceSetBalanceCall(in any) *pbgear.Balances_ForceSetBalanceCall {
	out := &pbgear.Balances_ForceSetBalanceCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_Balances_ForceSetBalanceCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
	// primitive field NewFree
	out.NewFree = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Balances_ForceSetBalanceCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_ForceSetBalanceCall_WhoId{
			WhoId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_ForceSetBalanceCall_WhoIndex{
			WhoIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_ForceSetBalanceCall_WhoRaw{
			WhoRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_ForceSetBalanceCall_WhoAddress32{
			WhoAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_ForceSetBalanceCall_WhoAddress20{
			WhoAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_ForceTransferCall(in any) *pbgear.Balances_ForceTransferCall {
	out := &pbgear.Balances_ForceTransferCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Source
	v0 := To_OneOf_Balances_ForceTransferCall_source(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))
	// oneOf field Dest
	v1 := To_OneOf_Balances_ForceTransferCall_dest(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v1))
	// primitive field Value
	out.Value = To_string(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Balances_ForceTransferCall_source(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_ForceTransferCall_SourceId{
			SourceId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_ForceTransferCall_SourceIndex{
			SourceIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_ForceTransferCall_SourceRaw{
			SourceRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_ForceTransferCall_SourceAddress32{
			SourceAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_ForceTransferCall_SourceAddress20{
			SourceAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Balances_ForceTransferCall_dest(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_ForceTransferCall_DestId{
			DestId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_ForceTransferCall_DestIndex{
			DestIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_ForceTransferCall_DestRaw{
			DestRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_ForceTransferCall_DestAddress32{
			DestAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_ForceTransferCall_DestAddress20{
			DestAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_ForceUnreserveCall(in any) *pbgear.Balances_ForceUnreserveCall {
	out := &pbgear.Balances_ForceUnreserveCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_Balances_ForceUnreserveCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
	// primitive field Amount
	out.Amount = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Balances_ForceUnreserveCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_ForceUnreserveCall_WhoId{
			WhoId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_ForceUnreserveCall_WhoIndex{
			WhoIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_ForceUnreserveCall_WhoRaw{
			WhoRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_ForceUnreserveCall_WhoAddress32{
			WhoAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_ForceUnreserveCall_WhoAddress20{
			WhoAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_Id(in any) *pbgear.Balances_Id {
	out := &pbgear.Balances_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Balances_Index(in any) *pbgear.Balances_Index {
	out := &pbgear.Balances_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Balances_TupleNull(w)
	out.Value0 = To_Balances_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Balances_Raw(in any) *pbgear.Balances_Raw {
	out := &pbgear.Balances_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Balances_Source(in any) *pbgear.Balances_Source {
	out := &pbgear.Balances_Source{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Source
	v0 := To_OneOf_Balances_Source_source(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Balances_Source_source(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_Source_SourceId{
			SourceId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_Source_SourceIndex{
			SourceIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_Source_SourceRaw{
			SourceRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_Source_SourceAddress32{
			SourceAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_Source_SourceAddress20{
			SourceAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_TransferAllCall(in any) *pbgear.Balances_TransferAllCall {
	out := &pbgear.Balances_TransferAllCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dest
	v0 := To_OneOf_Balances_TransferAllCall_dest(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Balances_TransferAllCall_dest(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_TransferAllCall_DestId{
			DestId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_TransferAllCall_DestIndex{
			DestIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_TransferAllCall_DestRaw{
			DestRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_TransferAllCall_DestAddress32{
			DestAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_TransferAllCall_DestAddress20{
			DestAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_TransferAllowDeathCall(in any) *pbgear.Balances_TransferAllowDeathCall {
	out := &pbgear.Balances_TransferAllowDeathCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dest
	v0 := To_OneOf_Balances_TransferAllowDeathCall_dest(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
	// primitive field Value
	out.Value = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Balances_TransferAllowDeathCall_dest(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_TransferAllowDeathCall_DestId{
			DestId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_TransferAllowDeathCall_DestIndex{
			DestIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_TransferAllowDeathCall_DestRaw{
			DestRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_TransferAllowDeathCall_DestAddress32{
			DestAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_TransferAllowDeathCall_DestAddress20{
			DestAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_TransferKeepAliveCall(in any) *pbgear.Balances_TransferKeepAliveCall {
	out := &pbgear.Balances_TransferKeepAliveCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Dest
	v0 := To_OneOf_Balances_TransferKeepAliveCall_dest(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
	// primitive field Value
	out.Value = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Balances_TransferKeepAliveCall_dest(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_TransferKeepAliveCall_DestId{
			DestId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_TransferKeepAliveCall_DestIndex{
			DestIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_TransferKeepAliveCall_DestRaw{
			DestRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_TransferKeepAliveCall_DestAddress32{
			DestAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_TransferKeepAliveCall_DestAddress20{
			DestAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Balances_TupleNull(in any) *pbgear.Balances_TupleNull {
	out := &pbgear.Balances_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Balances_UpgradeAccountsCall(in any) *pbgear.Balances_UpgradeAccountsCall {
	out := &pbgear.Balances_UpgradeAccountsCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Who
	out.Who = To_Repeated_Balances_UpgradeAccountsCall_who(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Balances_UpgradeAccountsCall_who(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Balances_Who(in any) *pbgear.Balances_Who {
	out := &pbgear.Balances_Who{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_Balances_Who_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Balances_Who_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Balances_Who_WhoId{
			WhoId: To_Balances_Id(param),
		}
	case 1:
		return &pbgear.Balances_Who_WhoIndex{
			WhoIndex: To_Balances_Index(param),
		}
	case 2:
		return &pbgear.Balances_Who_WhoRaw{
			WhoRaw: To_Balances_Raw(param),
		}
	case 3:
		return &pbgear.Balances_Who_WhoAddress32{
			WhoAddress32: To_Balances_Address32(param),
		}
	case 4:
		return &pbgear.Balances_Who_WhoAddress20{
			WhoAddress20: To_Balances_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_BigSpender(in any) *pbgear.BigSpender {
	out := &pbgear.BigSpender{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_BigTipper(in any) *pbgear.BigTipper {
	out := &pbgear.BigTipper{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_BoundedCollectionsBoundedVecBoundedVec(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
	out := &pbgear.BoundedCollectionsBoundedVecBoundedVec{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_BountiesPallet(in any) *pbgear.BountiesPallet {
	out := &pbgear.BountiesPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_BountiesPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_BountiesPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.BountiesPallet_CallProposeBountyCall{
			CallProposeBountyCall: To_Bounties_ProposeBountyCall(param),
		}
	case 1:
		return &pbgear.BountiesPallet_CallApproveBountyCall{
			CallApproveBountyCall: To_Bounties_ApproveBountyCall(param),
		}
	case 2:
		return &pbgear.BountiesPallet_CallProposeCuratorCall{
			CallProposeCuratorCall: To_Bounties_ProposeCuratorCall(param),
		}
	case 3:
		return &pbgear.BountiesPallet_CallUnassignCuratorCall{
			CallUnassignCuratorCall: To_Bounties_UnassignCuratorCall(param),
		}
	case 4:
		return &pbgear.BountiesPallet_CallAcceptCuratorCall{
			CallAcceptCuratorCall: To_Bounties_AcceptCuratorCall(param),
		}
	case 5:
		return &pbgear.BountiesPallet_CallAwardBountyCall{
			CallAwardBountyCall: To_Bounties_AwardBountyCall(param),
		}
	case 6:
		return &pbgear.BountiesPallet_CallClaimBountyCall{
			CallClaimBountyCall: To_Bounties_ClaimBountyCall(param),
		}
	case 7:
		return &pbgear.BountiesPallet_CallCloseBountyCall{
			CallCloseBountyCall: To_Bounties_CloseBountyCall(param),
		}
	case 8:
		return &pbgear.BountiesPallet_CallExtendBountyExpiryCall{
			CallExtendBountyExpiryCall: To_Bounties_ExtendBountyExpiryCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Bounties_AcceptCuratorCall(in any) *pbgear.Bounties_AcceptCuratorCall {
	out := &pbgear.Bounties_AcceptCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Bounties_Address20(in any) *pbgear.Bounties_Address20 {
	out := &pbgear.Bounties_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Bounties_Address32(in any) *pbgear.Bounties_Address32 {
	out := &pbgear.Bounties_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Bounties_ApproveBountyCall(in any) *pbgear.Bounties_ApproveBountyCall {
	out := &pbgear.Bounties_ApproveBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Bounties_AwardBountyCall(in any) *pbgear.Bounties_AwardBountyCall {
	out := &pbgear.Bounties_AwardBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))
	// oneOf field Beneficiary
	v1 := To_OneOf_Bounties_AwardBountyCall_beneficiary(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Bounties_AwardBountyCall_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Bounties_AwardBountyCall_BeneficiaryId{
			BeneficiaryId: To_Bounties_Id(param),
		}
	case 1:
		return &pbgear.Bounties_AwardBountyCall_BeneficiaryIndex{
			BeneficiaryIndex: To_Bounties_Index(param),
		}
	case 2:
		return &pbgear.Bounties_AwardBountyCall_BeneficiaryRaw{
			BeneficiaryRaw: To_Bounties_Raw(param),
		}
	case 3:
		return &pbgear.Bounties_AwardBountyCall_BeneficiaryAddress32{
			BeneficiaryAddress32: To_Bounties_Address32(param),
		}
	case 4:
		return &pbgear.Bounties_AwardBountyCall_BeneficiaryAddress20{
			BeneficiaryAddress20: To_Bounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Bounties_Beneficiary(in any) *pbgear.Bounties_Beneficiary {
	out := &pbgear.Bounties_Beneficiary{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Beneficiary
	v0 := To_OneOf_Bounties_Beneficiary_beneficiary(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Bounties_Beneficiary_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Bounties_Beneficiary_BeneficiaryId{
			BeneficiaryId: To_Bounties_Id(param),
		}
	case 1:
		return &pbgear.Bounties_Beneficiary_BeneficiaryIndex{
			BeneficiaryIndex: To_Bounties_Index(param),
		}
	case 2:
		return &pbgear.Bounties_Beneficiary_BeneficiaryRaw{
			BeneficiaryRaw: To_Bounties_Raw(param),
		}
	case 3:
		return &pbgear.Bounties_Beneficiary_BeneficiaryAddress32{
			BeneficiaryAddress32: To_Bounties_Address32(param),
		}
	case 4:
		return &pbgear.Bounties_Beneficiary_BeneficiaryAddress20{
			BeneficiaryAddress20: To_Bounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Bounties_ClaimBountyCall(in any) *pbgear.Bounties_ClaimBountyCall {
	out := &pbgear.Bounties_ClaimBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Bounties_CloseBountyCall(in any) *pbgear.Bounties_CloseBountyCall {
	out := &pbgear.Bounties_CloseBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Bounties_Curator(in any) *pbgear.Bounties_Curator {
	out := &pbgear.Bounties_Curator{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Curator
	v0 := To_OneOf_Bounties_Curator_curator(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Bounties_Curator_curator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Bounties_Curator_CuratorId{
			CuratorId: To_Bounties_Id(param),
		}
	case 1:
		return &pbgear.Bounties_Curator_CuratorIndex{
			CuratorIndex: To_Bounties_Index(param),
		}
	case 2:
		return &pbgear.Bounties_Curator_CuratorRaw{
			CuratorRaw: To_Bounties_Raw(param),
		}
	case 3:
		return &pbgear.Bounties_Curator_CuratorAddress32{
			CuratorAddress32: To_Bounties_Address32(param),
		}
	case 4:
		return &pbgear.Bounties_Curator_CuratorAddress20{
			CuratorAddress20: To_Bounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Bounties_ExtendBountyExpiryCall(in any) *pbgear.Bounties_ExtendBountyExpiryCall {
	out := &pbgear.Bounties_ExtendBountyExpiryCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))
	// primitive field Remark
	out.Remark = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Bounties_Id(in any) *pbgear.Bounties_Id {
	out := &pbgear.Bounties_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Bounties_Index(in any) *pbgear.Bounties_Index {
	out := &pbgear.Bounties_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Bounties_TupleNull(w)
	out.Value0 = To_Bounties_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Bounties_ProposeBountyCall(in any) *pbgear.Bounties_ProposeBountyCall {
	out := &pbgear.Bounties_ProposeBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))
	// primitive field Description
	out.Description = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Bounties_ProposeCuratorCall(in any) *pbgear.Bounties_ProposeCuratorCall {
	out := &pbgear.Bounties_ProposeCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))
	// oneOf field Curator
	v1 := To_OneOf_Bounties_ProposeCuratorCall_curator(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v1))
	// primitive field Fee
	out.Fee = To_string(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Bounties_ProposeCuratorCall_curator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Bounties_ProposeCuratorCall_CuratorId{
			CuratorId: To_Bounties_Id(param),
		}
	case 1:
		return &pbgear.Bounties_ProposeCuratorCall_CuratorIndex{
			CuratorIndex: To_Bounties_Index(param),
		}
	case 2:
		return &pbgear.Bounties_ProposeCuratorCall_CuratorRaw{
			CuratorRaw: To_Bounties_Raw(param),
		}
	case 3:
		return &pbgear.Bounties_ProposeCuratorCall_CuratorAddress32{
			CuratorAddress32: To_Bounties_Address32(param),
		}
	case 4:
		return &pbgear.Bounties_ProposeCuratorCall_CuratorAddress20{
			CuratorAddress20: To_Bounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Bounties_Raw(in any) *pbgear.Bounties_Raw {
	out := &pbgear.Bounties_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Bounties_TupleNull(in any) *pbgear.Bounties_TupleNull {
	out := &pbgear.Bounties_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Bounties_UnassignCuratorCall(in any) *pbgear.Bounties_UnassignCuratorCall {
	out := &pbgear.Bounties_UnassignCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BountyId
	out.BountyId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_ChildBountiesPallet(in any) *pbgear.ChildBountiesPallet {
	out := &pbgear.ChildBountiesPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ChildBountiesPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ChildBountiesPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ChildBountiesPallet_CallAddChildBountyCall{
			CallAddChildBountyCall: To_ChildBounties_AddChildBountyCall(param),
		}
	case 1:
		return &pbgear.ChildBountiesPallet_CallProposeCuratorCall{
			CallProposeCuratorCall: To_ChildBounties_ProposeCuratorCall(param),
		}
	case 2:
		return &pbgear.ChildBountiesPallet_CallAcceptCuratorCall{
			CallAcceptCuratorCall: To_ChildBounties_AcceptCuratorCall(param),
		}
	case 3:
		return &pbgear.ChildBountiesPallet_CallUnassignCuratorCall{
			CallUnassignCuratorCall: To_ChildBounties_UnassignCuratorCall(param),
		}
	case 4:
		return &pbgear.ChildBountiesPallet_CallAwardChildBountyCall{
			CallAwardChildBountyCall: To_ChildBounties_AwardChildBountyCall(param),
		}
	case 5:
		return &pbgear.ChildBountiesPallet_CallClaimChildBountyCall{
			CallClaimChildBountyCall: To_ChildBounties_ClaimChildBountyCall(param),
		}
	case 6:
		return &pbgear.ChildBountiesPallet_CallCloseChildBountyCall{
			CallCloseChildBountyCall: To_ChildBounties_CloseChildBountyCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ChildBounties_AcceptCuratorCall(in any) *pbgear.ChildBounties_AcceptCuratorCall {
	out := &pbgear.ChildBounties_AcceptCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ChildBounties_AddChildBountyCall(in any) *pbgear.ChildBounties_AddChildBountyCall {
	out := &pbgear.ChildBounties_AddChildBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field Value
	out.Value = To_string(v.ValueAt(1))
	// primitive field Description
	out.Description = To_bytes(v.ValueAt(2))

	return out //from message
}

func To_ChildBounties_Address20(in any) *pbgear.ChildBounties_Address20 {
	out := &pbgear.ChildBounties_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_Address32(in any) *pbgear.ChildBounties_Address32 {
	out := &pbgear.ChildBounties_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_AwardChildBountyCall(in any) *pbgear.ChildBounties_AwardChildBountyCall {
	out := &pbgear.ChildBounties_AwardChildBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))
	// oneOf field Beneficiary
	v2 := To_OneOf_ChildBounties_AwardChildBountyCall_beneficiary(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v2))

	return out //from message
}

func To_OneOf_ChildBounties_AwardChildBountyCall_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryId{
			BeneficiaryId: To_ChildBounties_Id(param),
		}
	case 1:
		return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryIndex{
			BeneficiaryIndex: To_ChildBounties_Index(param),
		}
	case 2:
		return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryRaw{
			BeneficiaryRaw: To_ChildBounties_Raw(param),
		}
	case 3:
		return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryAddress32{
			BeneficiaryAddress32: To_ChildBounties_Address32(param),
		}
	case 4:
		return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryAddress20{
			BeneficiaryAddress20: To_ChildBounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ChildBounties_Beneficiary(in any) *pbgear.ChildBounties_Beneficiary {
	out := &pbgear.ChildBounties_Beneficiary{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Beneficiary
	v0 := To_OneOf_ChildBounties_Beneficiary_beneficiary(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ChildBounties_Beneficiary_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ChildBounties_Beneficiary_BeneficiaryId{
			BeneficiaryId: To_ChildBounties_Id(param),
		}
	case 1:
		return &pbgear.ChildBounties_Beneficiary_BeneficiaryIndex{
			BeneficiaryIndex: To_ChildBounties_Index(param),
		}
	case 2:
		return &pbgear.ChildBounties_Beneficiary_BeneficiaryRaw{
			BeneficiaryRaw: To_ChildBounties_Raw(param),
		}
	case 3:
		return &pbgear.ChildBounties_Beneficiary_BeneficiaryAddress32{
			BeneficiaryAddress32: To_ChildBounties_Address32(param),
		}
	case 4:
		return &pbgear.ChildBounties_Beneficiary_BeneficiaryAddress20{
			BeneficiaryAddress20: To_ChildBounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ChildBounties_ClaimChildBountyCall(in any) *pbgear.ChildBounties_ClaimChildBountyCall {
	out := &pbgear.ChildBounties_ClaimChildBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ChildBounties_CloseChildBountyCall(in any) *pbgear.ChildBounties_CloseChildBountyCall {
	out := &pbgear.ChildBounties_CloseChildBountyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ChildBounties_Curator(in any) *pbgear.ChildBounties_Curator {
	out := &pbgear.ChildBounties_Curator{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Curator
	v0 := To_OneOf_ChildBounties_Curator_curator(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ChildBounties_Curator_curator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ChildBounties_Curator_CuratorId{
			CuratorId: To_ChildBounties_Id(param),
		}
	case 1:
		return &pbgear.ChildBounties_Curator_CuratorIndex{
			CuratorIndex: To_ChildBounties_Index(param),
		}
	case 2:
		return &pbgear.ChildBounties_Curator_CuratorRaw{
			CuratorRaw: To_ChildBounties_Raw(param),
		}
	case 3:
		return &pbgear.ChildBounties_Curator_CuratorAddress32{
			CuratorAddress32: To_ChildBounties_Address32(param),
		}
	case 4:
		return &pbgear.ChildBounties_Curator_CuratorAddress20{
			CuratorAddress20: To_ChildBounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ChildBounties_Id(in any) *pbgear.ChildBounties_Id {
	out := &pbgear.ChildBounties_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_Index(in any) *pbgear.ChildBounties_Index {
	out := &pbgear.ChildBounties_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_ChildBounties_TupleNull(w)
	out.Value0 = To_ChildBounties_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_ProposeCuratorCall(in any) *pbgear.ChildBounties_ProposeCuratorCall {
	out := &pbgear.ChildBounties_ProposeCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))
	// oneOf field Curator
	v2 := To_OneOf_ChildBounties_ProposeCuratorCall_curator(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v2))
	// primitive field Fee
	out.Fee = To_string(v.ValueAt(3))

	return out //from message
}

func To_OneOf_ChildBounties_ProposeCuratorCall_curator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ChildBounties_ProposeCuratorCall_CuratorId{
			CuratorId: To_ChildBounties_Id(param),
		}
	case 1:
		return &pbgear.ChildBounties_ProposeCuratorCall_CuratorIndex{
			CuratorIndex: To_ChildBounties_Index(param),
		}
	case 2:
		return &pbgear.ChildBounties_ProposeCuratorCall_CuratorRaw{
			CuratorRaw: To_ChildBounties_Raw(param),
		}
	case 3:
		return &pbgear.ChildBounties_ProposeCuratorCall_CuratorAddress32{
			CuratorAddress32: To_ChildBounties_Address32(param),
		}
	case 4:
		return &pbgear.ChildBounties_ProposeCuratorCall_CuratorAddress20{
			CuratorAddress20: To_ChildBounties_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ChildBounties_Raw(in any) *pbgear.ChildBounties_Raw {
	out := &pbgear.ChildBounties_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_TupleNull(in any) *pbgear.ChildBounties_TupleNull {
	out := &pbgear.ChildBounties_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_ChildBounties_UnassignCuratorCall(in any) *pbgear.ChildBounties_UnassignCuratorCall {
	out := &pbgear.ChildBounties_UnassignCuratorCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ParentBountyId
	out.ParentBountyId = To_uint32(v.ValueAt(0))
	// primitive field ChildBountyId
	out.ChildBountyId = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ConvictionVotingPallet(in any) *pbgear.ConvictionVotingPallet {
	out := &pbgear.ConvictionVotingPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ConvictionVotingPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ConvictionVotingPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVotingPallet_CallVoteCall{
			CallVoteCall: To_ConvictionVoting_VoteCall(param),
		}
	case 1:
		return &pbgear.ConvictionVotingPallet_CallDelegateCall{
			CallDelegateCall: To_ConvictionVoting_DelegateCall(param),
		}
	case 2:
		return &pbgear.ConvictionVotingPallet_CallUndelegateCall{
			CallUndelegateCall: To_ConvictionVoting_UndelegateCall(param),
		}
	case 3:
		return &pbgear.ConvictionVotingPallet_CallUnlockCall{
			CallUnlockCall: To_ConvictionVoting_UnlockCall(param),
		}
	case 4:
		return &pbgear.ConvictionVotingPallet_CallRemoveVoteCall{
			CallRemoveVoteCall: To_ConvictionVoting_RemoveVoteCall(param),
		}
	case 5:
		return &pbgear.ConvictionVotingPallet_CallRemoveOtherVoteCall{
			CallRemoveOtherVoteCall: To_ConvictionVoting_RemoveOtherVoteCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_Address20(in any) *pbgear.ConvictionVoting_Address20 {
	out := &pbgear.ConvictionVoting_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_Address32(in any) *pbgear.ConvictionVoting_Address32 {
	out := &pbgear.ConvictionVoting_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_Conviction(in any) *pbgear.ConvictionVoting_Conviction {
	out := &pbgear.ConvictionVoting_Conviction{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Conviction
	v0 := To_OneOf_ConvictionVoting_Conviction_conviction(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Conviction").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ConvictionVoting_Conviction_conviction(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_Conviction_ConvictionNone{
			ConvictionNone: To_ConvictionVoting_None(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked1X{
			ConvictionLocked1X: To_ConvictionVoting_Locked1X(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked2X{
			ConvictionLocked2X: To_ConvictionVoting_Locked2X(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked3X{
			ConvictionLocked3X: To_ConvictionVoting_Locked3X(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked4X{
			ConvictionLocked4X: To_ConvictionVoting_Locked4X(param),
		}
	case 5:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked5X{
			ConvictionLocked5X: To_ConvictionVoting_Locked5X(param),
		}
	case 6:
		return &pbgear.ConvictionVoting_Conviction_ConvictionLocked6X{
			ConvictionLocked6X: To_ConvictionVoting_Locked6X(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_DelegateCall(in any) *pbgear.ConvictionVoting_DelegateCall {
	out := &pbgear.ConvictionVoting_DelegateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Class
	out.Class = To_uint32(v.ValueAt(0))
	// oneOf field To
	v1 := To_OneOf_ConvictionVoting_DelegateCall_to(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v1))
	// oneOf field Conviction
	v2 := To_OneOf_ConvictionVoting_DelegateCall_conviction(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Conviction").Set(reflect.ValueOf(v2))
	// primitive field Balance
	out.Balance = To_string(v.ValueAt(3))

	return out //from message
}

func To_OneOf_ConvictionVoting_DelegateCall_to(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_DelegateCall_ToId{
			ToId: To_ConvictionVoting_Id(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_DelegateCall_ToIndex{
			ToIndex: To_ConvictionVoting_Index(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_DelegateCall_ToRaw{
			ToRaw: To_ConvictionVoting_Raw(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_DelegateCall_ToAddress32{
			ToAddress32: To_ConvictionVoting_Address32(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_DelegateCall_ToAddress20{
			ToAddress20: To_ConvictionVoting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_ConvictionVoting_DelegateCall_conviction(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionNone{
			ConvictionNone: To_ConvictionVoting_None(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked1X{
			ConvictionLocked1X: To_ConvictionVoting_Locked1X(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked2X{
			ConvictionLocked2X: To_ConvictionVoting_Locked2X(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked3X{
			ConvictionLocked3X: To_ConvictionVoting_Locked3X(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked4X{
			ConvictionLocked4X: To_ConvictionVoting_Locked4X(param),
		}
	case 5:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked5X{
			ConvictionLocked5X: To_ConvictionVoting_Locked5X(param),
		}
	case 6:
		return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked6X{
			ConvictionLocked6X: To_ConvictionVoting_Locked6X(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_Id(in any) *pbgear.ConvictionVoting_Id {
	out := &pbgear.ConvictionVoting_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_Index(in any) *pbgear.ConvictionVoting_Index {
	out := &pbgear.ConvictionVoting_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_ConvictionVoting_TupleNull(w)
	out.Value0 = To_ConvictionVoting_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_Locked1X(in any) *pbgear.ConvictionVoting_Locked1X {
	out := &pbgear.ConvictionVoting_Locked1X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_Locked2X(in any) *pbgear.ConvictionVoting_Locked2X {
	out := &pbgear.ConvictionVoting_Locked2X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_Locked3X(in any) *pbgear.ConvictionVoting_Locked3X {
	out := &pbgear.ConvictionVoting_Locked3X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_Locked4X(in any) *pbgear.ConvictionVoting_Locked4X {
	out := &pbgear.ConvictionVoting_Locked4X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_Locked5X(in any) *pbgear.ConvictionVoting_Locked5X {
	out := &pbgear.ConvictionVoting_Locked5X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_Locked6X(in any) *pbgear.ConvictionVoting_Locked6X {
	out := &pbgear.ConvictionVoting_Locked6X{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_None(in any) *pbgear.ConvictionVoting_None {
	out := &pbgear.ConvictionVoting_None{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ConvictionVoting_PalletConvictionVotingVoteVote(in any) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
	out := &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_Raw(in any) *pbgear.ConvictionVoting_Raw {
	out := &pbgear.ConvictionVoting_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_RemoveOtherVoteCall(in any) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
	out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_ConvictionVoting_RemoveOtherVoteCall_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))
	// primitive field Class
	out.Class = To_uint32(v.ValueAt(1))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_OneOf_ConvictionVoting_RemoveOtherVoteCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetId{
			TargetId: To_ConvictionVoting_Id(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetIndex{
			TargetIndex: To_ConvictionVoting_Index(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetRaw{
			TargetRaw: To_ConvictionVoting_Raw(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetAddress32{
			TargetAddress32: To_ConvictionVoting_Address32(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetAddress20{
			TargetAddress20: To_ConvictionVoting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_RemoveVoteCall(in any) *pbgear.ConvictionVoting_RemoveVoteCall {
	out := &pbgear.ConvictionVoting_RemoveVoteCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Class
	out.Class = To_Optional_uint32(v.ValueAt(0))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ConvictionVoting_Split(in any) *pbgear.ConvictionVoting_Split {
	out := &pbgear.ConvictionVoting_Split{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Aye
	out.Aye = To_string(v.ValueAt(0))
	// primitive field Nay
	out.Nay = To_string(v.ValueAt(1))

	return out //from message
}

func To_ConvictionVoting_SplitAbstain(in any) *pbgear.ConvictionVoting_SplitAbstain {
	out := &pbgear.ConvictionVoting_SplitAbstain{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Aye
	out.Aye = To_string(v.ValueAt(0))
	// primitive field Nay
	out.Nay = To_string(v.ValueAt(1))
	// primitive field Abstain
	out.Abstain = To_string(v.ValueAt(2))

	return out //from message
}

func To_ConvictionVoting_Standard(in any) *pbgear.ConvictionVoting_Standard {
	out := &pbgear.ConvictionVoting_Standard{}
	v := in.(registry.Valuable)
	_ = v

	// field Vote To_ConvictionVoting_PalletConvictionVotingVoteVote(w)
	out.Vote = To_ConvictionVoting_PalletConvictionVotingVoteVote(v.ValueAt(0))
	// primitive field Balance
	out.Balance = To_string(v.ValueAt(1))

	return out //from message
}

func To_ConvictionVoting_Target(in any) *pbgear.ConvictionVoting_Target {
	out := &pbgear.ConvictionVoting_Target{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_ConvictionVoting_Target_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ConvictionVoting_Target_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_Target_TargetId{
			TargetId: To_ConvictionVoting_Id(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_Target_TargetIndex{
			TargetIndex: To_ConvictionVoting_Index(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_Target_TargetRaw{
			TargetRaw: To_ConvictionVoting_Raw(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_Target_TargetAddress32{
			TargetAddress32: To_ConvictionVoting_Address32(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_Target_TargetAddress20{
			TargetAddress20: To_ConvictionVoting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_To(in any) *pbgear.ConvictionVoting_To {
	out := &pbgear.ConvictionVoting_To{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field To
	v0 := To_OneOf_ConvictionVoting_To_to(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ConvictionVoting_To_to(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_To_ToId{
			ToId: To_ConvictionVoting_Id(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_To_ToIndex{
			ToIndex: To_ConvictionVoting_Index(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_To_ToRaw{
			ToRaw: To_ConvictionVoting_Raw(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_To_ToAddress32{
			ToAddress32: To_ConvictionVoting_Address32(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_To_ToAddress20{
			ToAddress20: To_ConvictionVoting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_TupleNull(in any) *pbgear.ConvictionVoting_TupleNull {
	out := &pbgear.ConvictionVoting_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_UndelegateCall(in any) *pbgear.ConvictionVoting_UndelegateCall {
	out := &pbgear.ConvictionVoting_UndelegateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Class
	out.Class = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_ConvictionVoting_UnlockCall(in any) *pbgear.ConvictionVoting_UnlockCall {
	out := &pbgear.ConvictionVoting_UnlockCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Class
	out.Class = To_uint32(v.ValueAt(0))
	// oneOf field Target
	v1 := To_OneOf_ConvictionVoting_UnlockCall_target(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_ConvictionVoting_UnlockCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_UnlockCall_TargetId{
			TargetId: To_ConvictionVoting_Id(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_UnlockCall_TargetIndex{
			TargetIndex: To_ConvictionVoting_Index(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_UnlockCall_TargetRaw{
			TargetRaw: To_ConvictionVoting_Raw(param),
		}
	case 3:
		return &pbgear.ConvictionVoting_UnlockCall_TargetAddress32{
			TargetAddress32: To_ConvictionVoting_Address32(param),
		}
	case 4:
		return &pbgear.ConvictionVoting_UnlockCall_TargetAddress20{
			TargetAddress20: To_ConvictionVoting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_Vote(in any) *pbgear.ConvictionVoting_Vote {
	out := &pbgear.ConvictionVoting_Vote{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Vote
	v0 := To_OneOf_ConvictionVoting_Vote_vote(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ConvictionVoting_Vote_vote(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_Vote_VoteStandard{
			VoteStandard: To_ConvictionVoting_Standard(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_Vote_VoteSplit{
			VoteSplit: To_ConvictionVoting_Split(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_Vote_VoteSplitAbstain{
			VoteSplitAbstain: To_ConvictionVoting_SplitAbstain(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ConvictionVoting_VoteCall(in any) *pbgear.ConvictionVoting_VoteCall {
	out := &pbgear.ConvictionVoting_VoteCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PollIndex
	out.PollIndex = To_uint32(v.ValueAt(0))
	// oneOf field Vote
	v1 := To_OneOf_ConvictionVoting_VoteCall_vote(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_ConvictionVoting_VoteCall_vote(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ConvictionVoting_VoteCall_VoteStandard{
			VoteStandard: To_ConvictionVoting_Standard(param),
		}
	case 1:
		return &pbgear.ConvictionVoting_VoteCall_VoteSplit{
			VoteSplit: To_ConvictionVoting_Split(param),
		}
	case 2:
		return &pbgear.ConvictionVoting_VoteCall_VoteSplitAbstain{
			VoteSplitAbstain: To_ConvictionVoting_SplitAbstain(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ElectionProviderMultiPhasePallet(in any) *pbgear.ElectionProviderMultiPhasePallet {
	out := &pbgear.ElectionProviderMultiPhasePallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ElectionProviderMultiPhasePallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ElectionProviderMultiPhasePallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ElectionProviderMultiPhasePallet_CallSubmitUnsignedCall{
			CallSubmitUnsignedCall: To_ElectionProviderMultiPhase_SubmitUnsignedCall(param),
		}
	case 1:
		return &pbgear.ElectionProviderMultiPhasePallet_CallSetMinimumUntrustedScoreCall{
			CallSetMinimumUntrustedScoreCall: To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(param),
		}
	case 2:
		return &pbgear.ElectionProviderMultiPhasePallet_CallSetEmergencyElectionResultCall{
			CallSetEmergencyElectionResultCall: To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(param),
		}
	case 3:
		return &pbgear.ElectionProviderMultiPhasePallet_CallSubmitCall{
			CallSubmitCall: To_ElectionProviderMultiPhase_SubmitCall(param),
		}
	case 4:
		return &pbgear.ElectionProviderMultiPhasePallet_CallGovernanceFallbackCall{
			CallGovernanceFallbackCall: To_ElectionProviderMultiPhase_GovernanceFallbackCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ElectionProviderMultiPhase_GovernanceFallbackCall(in any) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
	out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field MaybeMaxVoters
	out.MaybeMaxVoters = To_Optional_uint32(v.ValueAt(0))
	// primitive field MaybeMaxTargets
	out.MaybeMaxTargets = To_Optional_uint32(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
	v := in.(registry.Valuable)
	_ = v

	// field Solution To_VaraRuntimeNposSolution16(w)
	out.Solution = To_VaraRuntimeNposSolution16(v.ValueAt(0))
	// field Score To_SpNposElectionsElectionScore(w)
	out.Score = To_SpNposElectionsElectionScore(v.ValueAt(1))
	// primitive field Round
	out.Round = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Voters
	out.Voters = To_uint32(v.ValueAt(0))
	// primitive field Targets
	out.Targets = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(in any) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
	out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Supports
	out.Supports = To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(v.ValueAt(0))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(in any) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
	out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
	v := in.(registry.Valuable)
	_ = v

	// optional field MaybeNextScore
	out.MaybeNextScore = To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(v.ValueAt(0))

	return out //from message
}

func To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(in any) *pbgear.SpNposElectionsElectionScore {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_ElectionProviderMultiPhase_SubmitCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
	v := in.(registry.Valuable)
	_ = v

	// field RawSolution To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(w)
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(v.ValueAt(0))

	return out //from message
}

func To_ElectionProviderMultiPhase_SubmitUnsignedCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
	out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
	v := in.(registry.Valuable)
	_ = v

	// field RawSolution To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(w)
	out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(v.ValueAt(0))
	// field Witness To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(w)
	out.Witness = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// field Value1 To_SpNposElectionsSupport(w)
	out.Value1 = To_SpNposElectionsSupport(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_string(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// repeated field Value1
	out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// field Value1 To_SpArithmeticPerThingsPerU16(w)
	out.Value1 = To_SpArithmeticPerThingsPerU16(v.ValueAt(1))

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// field Value1 To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
	out.Value1 = To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(v.ValueAt(1))
	// primitive field Value2
	out.Value2 = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_ElectionProviderMultiPhase_TupleUint32Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
	out := &pbgear.ElectionProviderMultiPhase_TupleUint32Uint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Fellows(in any) *pbgear.Fellows {
	out := &pbgear.Fellows{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship1Dan(in any) *pbgear.Fellowship1Dan {
	out := &pbgear.Fellowship1Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship2Dan(in any) *pbgear.Fellowship2Dan {
	out := &pbgear.Fellowship2Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship3Dan(in any) *pbgear.Fellowship3Dan {
	out := &pbgear.Fellowship3Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship4Dan(in any) *pbgear.Fellowship4Dan {
	out := &pbgear.Fellowship4Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship5Dan(in any) *pbgear.Fellowship5Dan {
	out := &pbgear.Fellowship5Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship6Dan(in any) *pbgear.Fellowship6Dan {
	out := &pbgear.Fellowship6Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship7Dan(in any) *pbgear.Fellowship7Dan {
	out := &pbgear.Fellowship7Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship8Dan(in any) *pbgear.Fellowship8Dan {
	out := &pbgear.Fellowship8Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Fellowship9Dan(in any) *pbgear.Fellowship9Dan {
	out := &pbgear.Fellowship9Dan{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_FellowshipAdmin(in any) *pbgear.FellowshipAdmin {
	out := &pbgear.FellowshipAdmin{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_FellowshipCollectivePallet(in any) *pbgear.FellowshipCollectivePallet {
	out := &pbgear.FellowshipCollectivePallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_FellowshipCollectivePallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipCollectivePallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollectivePallet_CallAddMemberCall{
			CallAddMemberCall: To_FellowshipCollective_AddMemberCall(param),
		}
	case 1:
		return &pbgear.FellowshipCollectivePallet_CallPromoteMemberCall{
			CallPromoteMemberCall: To_FellowshipCollective_PromoteMemberCall(param),
		}
	case 2:
		return &pbgear.FellowshipCollectivePallet_CallDemoteMemberCall{
			CallDemoteMemberCall: To_FellowshipCollective_DemoteMemberCall(param),
		}
	case 3:
		return &pbgear.FellowshipCollectivePallet_CallRemoveMemberCall{
			CallRemoveMemberCall: To_FellowshipCollective_RemoveMemberCall(param),
		}
	case 4:
		return &pbgear.FellowshipCollectivePallet_CallVoteCall{
			CallVoteCall: To_FellowshipCollective_VoteCall(param),
		}
	case 5:
		return &pbgear.FellowshipCollectivePallet_CallCleanupPollCall{
			CallCleanupPollCall: To_FellowshipCollective_CleanupPollCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipCollective_AddMemberCall(in any) *pbgear.FellowshipCollective_AddMemberCall {
	out := &pbgear.FellowshipCollective_AddMemberCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_FellowshipCollective_AddMemberCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipCollective_AddMemberCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollective_AddMemberCall_WhoId{
			WhoId: To_FellowshipCollective_Id(param),
		}
	case 1:
		return &pbgear.FellowshipCollective_AddMemberCall_WhoIndex{
			WhoIndex: To_FellowshipCollective_Index(param),
		}
	case 2:
		return &pbgear.FellowshipCollective_AddMemberCall_WhoRaw{
			WhoRaw: To_FellowshipCollective_Raw(param),
		}
	case 3:
		return &pbgear.FellowshipCollective_AddMemberCall_WhoAddress32{
			WhoAddress32: To_FellowshipCollective_Address32(param),
		}
	case 4:
		return &pbgear.FellowshipCollective_AddMemberCall_WhoAddress20{
			WhoAddress20: To_FellowshipCollective_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipCollective_Address20(in any) *pbgear.FellowshipCollective_Address20 {
	out := &pbgear.FellowshipCollective_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_Address32(in any) *pbgear.FellowshipCollective_Address32 {
	out := &pbgear.FellowshipCollective_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_CleanupPollCall(in any) *pbgear.FellowshipCollective_CleanupPollCall {
	out := &pbgear.FellowshipCollective_CleanupPollCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PollIndex
	out.PollIndex = To_uint32(v.ValueAt(0))
	// primitive field Max
	out.Max = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_FellowshipCollective_DemoteMemberCall(in any) *pbgear.FellowshipCollective_DemoteMemberCall {
	out := &pbgear.FellowshipCollective_DemoteMemberCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_FellowshipCollective_DemoteMemberCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipCollective_DemoteMemberCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollective_DemoteMemberCall_WhoId{
			WhoId: To_FellowshipCollective_Id(param),
		}
	case 1:
		return &pbgear.FellowshipCollective_DemoteMemberCall_WhoIndex{
			WhoIndex: To_FellowshipCollective_Index(param),
		}
	case 2:
		return &pbgear.FellowshipCollective_DemoteMemberCall_WhoRaw{
			WhoRaw: To_FellowshipCollective_Raw(param),
		}
	case 3:
		return &pbgear.FellowshipCollective_DemoteMemberCall_WhoAddress32{
			WhoAddress32: To_FellowshipCollective_Address32(param),
		}
	case 4:
		return &pbgear.FellowshipCollective_DemoteMemberCall_WhoAddress20{
			WhoAddress20: To_FellowshipCollective_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipCollective_Id(in any) *pbgear.FellowshipCollective_Id {
	out := &pbgear.FellowshipCollective_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_Index(in any) *pbgear.FellowshipCollective_Index {
	out := &pbgear.FellowshipCollective_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_FellowshipCollective_TupleNull(w)
	out.Value0 = To_FellowshipCollective_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_PromoteMemberCall(in any) *pbgear.FellowshipCollective_PromoteMemberCall {
	out := &pbgear.FellowshipCollective_PromoteMemberCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_FellowshipCollective_PromoteMemberCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipCollective_PromoteMemberCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollective_PromoteMemberCall_WhoId{
			WhoId: To_FellowshipCollective_Id(param),
		}
	case 1:
		return &pbgear.FellowshipCollective_PromoteMemberCall_WhoIndex{
			WhoIndex: To_FellowshipCollective_Index(param),
		}
	case 2:
		return &pbgear.FellowshipCollective_PromoteMemberCall_WhoRaw{
			WhoRaw: To_FellowshipCollective_Raw(param),
		}
	case 3:
		return &pbgear.FellowshipCollective_PromoteMemberCall_WhoAddress32{
			WhoAddress32: To_FellowshipCollective_Address32(param),
		}
	case 4:
		return &pbgear.FellowshipCollective_PromoteMemberCall_WhoAddress20{
			WhoAddress20: To_FellowshipCollective_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipCollective_Raw(in any) *pbgear.FellowshipCollective_Raw {
	out := &pbgear.FellowshipCollective_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_RemoveMemberCall(in any) *pbgear.FellowshipCollective_RemoveMemberCall {
	out := &pbgear.FellowshipCollective_RemoveMemberCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_FellowshipCollective_RemoveMemberCall_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
	// primitive field MinRank
	out.MinRank = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_OneOf_FellowshipCollective_RemoveMemberCall_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollective_RemoveMemberCall_WhoId{
			WhoId: To_FellowshipCollective_Id(param),
		}
	case 1:
		return &pbgear.FellowshipCollective_RemoveMemberCall_WhoIndex{
			WhoIndex: To_FellowshipCollective_Index(param),
		}
	case 2:
		return &pbgear.FellowshipCollective_RemoveMemberCall_WhoRaw{
			WhoRaw: To_FellowshipCollective_Raw(param),
		}
	case 3:
		return &pbgear.FellowshipCollective_RemoveMemberCall_WhoAddress32{
			WhoAddress32: To_FellowshipCollective_Address32(param),
		}
	case 4:
		return &pbgear.FellowshipCollective_RemoveMemberCall_WhoAddress20{
			WhoAddress20: To_FellowshipCollective_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipCollective_TupleNull(in any) *pbgear.FellowshipCollective_TupleNull {
	out := &pbgear.FellowshipCollective_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_FellowshipCollective_VoteCall(in any) *pbgear.FellowshipCollective_VoteCall {
	out := &pbgear.FellowshipCollective_VoteCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Poll
	out.Poll = To_uint32(v.ValueAt(0))
	// primitive field Aye
	out.Aye = To_bool(v.ValueAt(1))

	return out //from message
}

func To_FellowshipCollective_Who(in any) *pbgear.FellowshipCollective_Who {
	out := &pbgear.FellowshipCollective_Who{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_FellowshipCollective_Who_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipCollective_Who_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipCollective_Who_WhoId{
			WhoId: To_FellowshipCollective_Id(param),
		}
	case 1:
		return &pbgear.FellowshipCollective_Who_WhoIndex{
			WhoIndex: To_FellowshipCollective_Index(param),
		}
	case 2:
		return &pbgear.FellowshipCollective_Who_WhoRaw{
			WhoRaw: To_FellowshipCollective_Raw(param),
		}
	case 3:
		return &pbgear.FellowshipCollective_Who_WhoAddress32{
			WhoAddress32: To_FellowshipCollective_Address32(param),
		}
	case 4:
		return &pbgear.FellowshipCollective_Who_WhoAddress20{
			WhoAddress20: To_FellowshipCollective_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipExperts(in any) *pbgear.FellowshipExperts {
	out := &pbgear.FellowshipExperts{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_FellowshipInitiates(in any) *pbgear.FellowshipInitiates {
	out := &pbgear.FellowshipInitiates{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_FellowshipMasters(in any) *pbgear.FellowshipMasters {
	out := &pbgear.FellowshipMasters{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_FellowshipReferendaPallet(in any) *pbgear.FellowshipReferendaPallet {
	out := &pbgear.FellowshipReferendaPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_FellowshipReferendaPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferendaPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferendaPallet_CallSubmitCall{
			CallSubmitCall: To_FellowshipReferenda_SubmitCall(param),
		}
	case 1:
		return &pbgear.FellowshipReferendaPallet_CallPlaceDecisionDepositCall{
			CallPlaceDecisionDepositCall: To_FellowshipReferenda_PlaceDecisionDepositCall(param),
		}
	case 2:
		return &pbgear.FellowshipReferendaPallet_CallRefundDecisionDepositCall{
			CallRefundDecisionDepositCall: To_FellowshipReferenda_RefundDecisionDepositCall(param),
		}
	case 3:
		return &pbgear.FellowshipReferendaPallet_CallCancelCall{
			CallCancelCall: To_FellowshipReferenda_CancelCall(param),
		}
	case 4:
		return &pbgear.FellowshipReferendaPallet_CallKillCall{
			CallKillCall: To_FellowshipReferenda_KillCall(param),
		}
	case 5:
		return &pbgear.FellowshipReferendaPallet_CallNudgeReferendumCall{
			CallNudgeReferendumCall: To_FellowshipReferenda_NudgeReferendumCall(param),
		}
	case 6:
		return &pbgear.FellowshipReferendaPallet_CallOneFewerDecidingCall{
			CallOneFewerDecidingCall: To_FellowshipReferenda_OneFewerDecidingCall(param),
		}
	case 7:
		return &pbgear.FellowshipReferendaPallet_CallRefundSubmissionDepositCall{
			CallRefundSubmissionDepositCall: To_FellowshipReferenda_RefundSubmissionDepositCall(param),
		}
	case 8:
		return &pbgear.FellowshipReferendaPallet_CallSetMetadataCall{
			CallSetMetadataCall: To_FellowshipReferenda_SetMetadataCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_After(in any) *pbgear.FellowshipReferenda_After {
	out := &pbgear.FellowshipReferenda_After{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_At(in any) *pbgear.FellowshipReferenda_At {
	out := &pbgear.FellowshipReferenda_At{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_CancelCall(in any) *pbgear.FellowshipReferenda_CancelCall {
	out := &pbgear.FellowshipReferenda_CancelCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_EnactmentMoment(in any) *pbgear.FellowshipReferenda_EnactmentMoment {
	out := &pbgear.FellowshipReferenda_EnactmentMoment{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field EnactmentMoment
	v0 := To_OneOf_FellowshipReferenda_EnactmentMoment_enactment_moment(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_EnactmentMoment_enactment_moment(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_EnactmentMoment_EnactmentMomentAt{
			EnactmentMomentAt: To_FellowshipReferenda_At(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_EnactmentMoment_EnactmentMomentAfter{
			EnactmentMomentAfter: To_FellowshipReferenda_After(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_Inline(in any) *pbgear.FellowshipReferenda_Inline {
	out := &pbgear.FellowshipReferenda_Inline{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_BoundedCollectionsBoundedVecBoundedVec(w)
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_KillCall(in any) *pbgear.FellowshipReferenda_KillCall {
	out := &pbgear.FellowshipReferenda_KillCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_Legacy(in any) *pbgear.FellowshipReferenda_Legacy {
	out := &pbgear.FellowshipReferenda_Legacy{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_Lookup(in any) *pbgear.FellowshipReferenda_Lookup {
	out := &pbgear.FellowshipReferenda_Lookup{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field Len
	out.Len = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_FellowshipReferenda_NudgeReferendumCall(in any) *pbgear.FellowshipReferenda_NudgeReferendumCall {
	out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_OneFewerDecidingCall(in any) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
	out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Track
	out.Track = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_Origins(in any) *pbgear.FellowshipReferenda_Origins {
	out := &pbgear.FellowshipReferenda_Origins{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_FellowshipReferenda_Origins_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_Origins_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_Origins_Value0StakingAdmin{
			Value0StakingAdmin: To_StakingAdmin(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_Origins_Value0Treasurer{
			Value0Treasurer: To_Treasurer(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_Origins_Value0FellowshipAdmin{
			Value0FellowshipAdmin: To_FellowshipAdmin(param),
		}
	case 3:
		return &pbgear.FellowshipReferenda_Origins_Value0GeneralAdmin{
			Value0GeneralAdmin: To_GeneralAdmin(param),
		}
	case 4:
		return &pbgear.FellowshipReferenda_Origins_Value0ReferendumCanceller{
			Value0ReferendumCanceller: To_ReferendumCanceller(param),
		}
	case 5:
		return &pbgear.FellowshipReferenda_Origins_Value0ReferendumKiller{
			Value0ReferendumKiller: To_ReferendumKiller(param),
		}
	case 6:
		return &pbgear.FellowshipReferenda_Origins_Value0SmallTipper{
			Value0SmallTipper: To_SmallTipper(param),
		}
	case 7:
		return &pbgear.FellowshipReferenda_Origins_Value0BigTipper{
			Value0BigTipper: To_BigTipper(param),
		}
	case 8:
		return &pbgear.FellowshipReferenda_Origins_Value0SmallSpender{
			Value0SmallSpender: To_SmallSpender(param),
		}
	case 9:
		return &pbgear.FellowshipReferenda_Origins_Value0MediumSpender{
			Value0MediumSpender: To_MediumSpender(param),
		}
	case 10:
		return &pbgear.FellowshipReferenda_Origins_Value0BigSpender{
			Value0BigSpender: To_BigSpender(param),
		}
	case 11:
		return &pbgear.FellowshipReferenda_Origins_Value0WhitelistedCaller{
			Value0WhitelistedCaller: To_WhitelistedCaller(param),
		}
	case 12:
		return &pbgear.FellowshipReferenda_Origins_Value0FellowshipInitiates{
			Value0FellowshipInitiates: To_FellowshipInitiates(param),
		}
	case 13:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellows{
			Value0Fellows: To_Fellows(param),
		}
	case 14:
		return &pbgear.FellowshipReferenda_Origins_Value0FellowshipExperts{
			Value0FellowshipExperts: To_FellowshipExperts(param),
		}
	case 15:
		return &pbgear.FellowshipReferenda_Origins_Value0FellowshipMasters{
			Value0FellowshipMasters: To_FellowshipMasters(param),
		}
	case 16:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship1Dan{
			Value0Fellowship1Dan: To_Fellowship1Dan(param),
		}
	case 17:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship2Dan{
			Value0Fellowship2Dan: To_Fellowship2Dan(param),
		}
	case 18:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship3Dan{
			Value0Fellowship3Dan: To_Fellowship3Dan(param),
		}
	case 19:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship4Dan{
			Value0Fellowship4Dan: To_Fellowship4Dan(param),
		}
	case 20:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship5Dan{
			Value0Fellowship5Dan: To_Fellowship5Dan(param),
		}
	case 21:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship6Dan{
			Value0Fellowship6Dan: To_Fellowship6Dan(param),
		}
	case 22:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship7Dan{
			Value0Fellowship7Dan: To_Fellowship7Dan(param),
		}
	case 23:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship8Dan{
			Value0Fellowship8Dan: To_Fellowship8Dan(param),
		}
	case 24:
		return &pbgear.FellowshipReferenda_Origins_Value0Fellowship9Dan{
			Value0Fellowship9Dan: To_Fellowship9Dan(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_PlaceDecisionDepositCall(in any) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_Proposal(in any) *pbgear.FellowshipReferenda_Proposal {
	out := &pbgear.FellowshipReferenda_Proposal{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Proposal
	v0 := To_OneOf_FellowshipReferenda_Proposal_proposal(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_Proposal_proposal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_Proposal_ProposalLegacy{
			ProposalLegacy: To_FellowshipReferenda_Legacy(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_Proposal_ProposalInline{
			ProposalInline: To_FellowshipReferenda_Inline(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_Proposal_ProposalLookup{
			ProposalLookup: To_FellowshipReferenda_Lookup(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_ProposalOrigin(in any) *pbgear.FellowshipReferenda_ProposalOrigin {
	out := &pbgear.FellowshipReferenda_ProposalOrigin{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProposalOrigin
	v0 := To_OneOf_FellowshipReferenda_ProposalOrigin_proposal_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_ProposalOrigin_proposal_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginSystem{
			ProposalOriginSystem: To_FellowshipReferenda_System(param),
		}
	case 20:
		return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginOrigins{
			ProposalOriginOrigins: To_FellowshipReferenda_Origins(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginVoid{
			ProposalOriginVoid: To_FellowshipReferenda_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_RefundDecisionDepositCall(in any) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_RefundSubmissionDepositCall(in any) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
	out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_FellowshipReferenda_SetMetadataCall(in any) *pbgear.FellowshipReferenda_SetMetadataCall {
	out := &pbgear.FellowshipReferenda_SetMetadataCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// optional field MaybeHash
	out.MaybeHash = To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(v.ValueAt(1))

	return out //from message
}

func To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_FellowshipReferenda_SubmitCall(in any) *pbgear.FellowshipReferenda_SubmitCall {
	out := &pbgear.FellowshipReferenda_SubmitCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProposalOrigin
	v0 := To_OneOf_FellowshipReferenda_SubmitCall_proposal_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))
	// oneOf field Proposal
	v1 := To_OneOf_FellowshipReferenda_SubmitCall_proposal(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v1))
	// oneOf field EnactmentMoment
	v2 := To_OneOf_FellowshipReferenda_SubmitCall_enactment_moment(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v2))

	return out //from message
}

func To_OneOf_FellowshipReferenda_SubmitCall_proposal_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginSystem{
			ProposalOriginSystem: To_FellowshipReferenda_System(param),
		}
	case 20:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginOrigins{
			ProposalOriginOrigins: To_FellowshipReferenda_Origins(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginVoid{
			ProposalOriginVoid: To_FellowshipReferenda_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_FellowshipReferenda_SubmitCall_proposal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalLegacy{
			ProposalLegacy: To_FellowshipReferenda_Legacy(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalInline{
			ProposalInline: To_FellowshipReferenda_Inline(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_SubmitCall_ProposalLookup{
			ProposalLookup: To_FellowshipReferenda_Lookup(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_FellowshipReferenda_SubmitCall_enactment_moment(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_SubmitCall_EnactmentMomentAt{
			EnactmentMomentAt: To_FellowshipReferenda_At(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_SubmitCall_EnactmentMomentAfter{
			EnactmentMomentAfter: To_FellowshipReferenda_After(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_System(in any) *pbgear.FellowshipReferenda_System {
	out := &pbgear.FellowshipReferenda_System{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_FellowshipReferenda_System_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_System_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.FellowshipReferenda_System_Value0Root{
			Value0Root: To_Root(param),
		}
	case 1:
		return &pbgear.FellowshipReferenda_System_Value0Signed{
			Value0Signed: To_Signed(param),
		}
	case 2:
		return &pbgear.FellowshipReferenda_System_Value0None{
			Value0None: To_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_FellowshipReferenda_Void(in any) *pbgear.FellowshipReferenda_Void {
	out := &pbgear.FellowshipReferenda_Void{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_FellowshipReferenda_Void_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_FellowshipReferenda_Void_value0(in any) any {
	return nil
}
func To_GearBankPallet(in any) *pbgear.GearBankPallet {
	out := &pbgear.GearBankPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearBuiltinPallet(in any) *pbgear.GearBuiltinPallet {
	out := &pbgear.GearBuiltinPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearGasPallet(in any) *pbgear.GearGasPallet {
	out := &pbgear.GearGasPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearMessengerPallet(in any) *pbgear.GearMessengerPallet {
	out := &pbgear.GearMessengerPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearPallet(in any) *pbgear.GearPallet {
	out := &pbgear.GearPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_GearPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GearPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearPallet_CallUploadCodeCall{
			CallUploadCodeCall: To_Gear_UploadCodeCall(param),
		}
	case 1:
		return &pbgear.GearPallet_CallUploadProgramCall{
			CallUploadProgramCall: To_Gear_UploadProgramCall(param),
		}
	case 2:
		return &pbgear.GearPallet_CallCreateProgramCall{
			CallCreateProgramCall: To_Gear_CreateProgramCall(param),
		}
	case 3:
		return &pbgear.GearPallet_CallSendMessageCall{
			CallSendMessageCall: To_Gear_SendMessageCall(param),
		}
	case 4:
		return &pbgear.GearPallet_CallSendReplyCall{
			CallSendReplyCall: To_Gear_SendReplyCall(param),
		}
	case 5:
		return &pbgear.GearPallet_CallClaimValueCall{
			CallClaimValueCall: To_Gear_ClaimValueCall(param),
		}
	case 6:
		return &pbgear.GearPallet_CallRunCall{
			CallRunCall: To_Gear_RunCall(param),
		}
	case 7:
		return &pbgear.GearPallet_CallSetExecuteInherentCall{
			CallSetExecuteInherentCall: To_Gear_SetExecuteInherentCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearPaymentPallet(in any) *pbgear.GearPaymentPallet {
	out := &pbgear.GearPaymentPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearProgramPallet(in any) *pbgear.GearProgramPallet {
	out := &pbgear.GearProgramPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearSchedulerPallet(in any) *pbgear.GearSchedulerPallet {
	out := &pbgear.GearSchedulerPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearVoucherPallet(in any) *pbgear.GearVoucherPallet {
	out := &pbgear.GearVoucherPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_GearVoucherPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GearVoucherPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucherPallet_CallIssueCall{
			CallIssueCall: To_GearVoucher_IssueCall(param),
		}
	case 1:
		return &pbgear.GearVoucherPallet_CallCallCall{
			CallCallCall: To_GearVoucher_CallCall(param),
		}
	case 2:
		return &pbgear.GearVoucherPallet_CallRevokeCall{
			CallRevokeCall: To_GearVoucher_RevokeCall(param),
		}
	case 3:
		return &pbgear.GearVoucherPallet_CallUpdateCall{
			CallUpdateCall: To_GearVoucher_UpdateCall(param),
		}
	case 4:
		return &pbgear.GearVoucherPallet_CallCallDeprecatedCall{
			CallCallDeprecatedCall: To_GearVoucher_CallDeprecatedCall(param),
		}
	case 5:
		return &pbgear.GearVoucherPallet_CallDeclineCall{
			CallDeclineCall: To_GearVoucher_DeclineCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_AppendPrograms(in any) *pbgear.GearVoucher_AppendPrograms {
	out := &pbgear.GearVoucher_AppendPrograms{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field AppendPrograms
	v0 := To_OneOf_GearVoucher_AppendPrograms_append_programs(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("AppendPrograms").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GearVoucher_AppendPrograms_append_programs(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucher_AppendPrograms_AppendProgramsNone{
			AppendProgramsNone: To_GearVoucher_None(param),
		}
	case 1:
		return &pbgear.GearVoucher_AppendPrograms_AppendProgramsSome{
			AppendProgramsSome: To_GearVoucher_Some(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_Call(in any) *pbgear.GearVoucher_Call {
	out := &pbgear.GearVoucher_Call{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_GearVoucher_Call_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GearVoucher_Call_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucher_Call_CallSendMessage{
			CallSendMessage: To_GearVoucher_SendMessage(param),
		}
	case 1:
		return &pbgear.GearVoucher_Call_CallSendReply{
			CallSendReply: To_GearVoucher_SendReply(param),
		}
	case 2:
		return &pbgear.GearVoucher_Call_CallUploadCode{
			CallUploadCode: To_GearVoucher_UploadCode(param),
		}
	case 3:
		return &pbgear.GearVoucher_Call_CallDeclineVoucher{
			CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_CallCall(in any) *pbgear.GearVoucher_CallCall {
	out := &pbgear.GearVoucher_CallCall{}
	v := in.(registry.Valuable)
	_ = v

	// field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(0))
	// oneOf field Call
	v1 := To_OneOf_GearVoucher_CallCall_call(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_GearVoucher_CallCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucher_CallCall_CallSendMessage{
			CallSendMessage: To_GearVoucher_SendMessage(param),
		}
	case 1:
		return &pbgear.GearVoucher_CallCall_CallSendReply{
			CallSendReply: To_GearVoucher_SendReply(param),
		}
	case 2:
		return &pbgear.GearVoucher_CallCall_CallUploadCode{
			CallUploadCode: To_GearVoucher_UploadCode(param),
		}
	case 3:
		return &pbgear.GearVoucher_CallCall_CallDeclineVoucher{
			CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_CallDeprecatedCall(in any) *pbgear.GearVoucher_CallDeprecatedCall {
	out := &pbgear.GearVoucher_CallDeprecatedCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_GearVoucher_CallDeprecatedCall_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GearVoucher_CallDeprecatedCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucher_CallDeprecatedCall_CallSendMessage{
			CallSendMessage: To_GearVoucher_SendMessage(param),
		}
	case 1:
		return &pbgear.GearVoucher_CallDeprecatedCall_CallSendReply{
			CallSendReply: To_GearVoucher_SendReply(param),
		}
	case 2:
		return &pbgear.GearVoucher_CallDeprecatedCall_CallUploadCode{
			CallUploadCode: To_GearVoucher_UploadCode(param),
		}
	case 3:
		return &pbgear.GearVoucher_CallDeprecatedCall_CallDeclineVoucher{
			CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_DeclineCall(in any) *pbgear.GearVoucher_DeclineCall {
	out := &pbgear.GearVoucher_DeclineCall{}
	v := in.(registry.Valuable)
	_ = v

	// field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(0))

	return out //from message
}

func To_GearVoucher_DeclineVoucher(in any) *pbgear.GearVoucher_DeclineVoucher {
	out := &pbgear.GearVoucher_DeclineVoucher{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearVoucher_IssueCall(in any) *pbgear.GearVoucher_IssueCall {
	out := &pbgear.GearVoucher_IssueCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Spender To_SpCoreCryptoAccountId32(w)
	out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// primitive field Balance
	out.Balance = To_string(v.ValueAt(1))
	// optional field Programs
	out.Programs = To_Optional_GearVoucher_IssueCall_programs(v.ValueAt(2))
	// primitive field CodeUploading
	out.CodeUploading = To_bool(v.ValueAt(3))
	// primitive field Duration
	out.Duration = To_uint32(v.ValueAt(4))

	return out //from message
}

func To_Optional_GearVoucher_IssueCall_programs(in any) *pbgear.BTreeSet {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_GearVoucher_None(in any) *pbgear.GearVoucher_None {
	out := &pbgear.GearVoucher_None{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GearVoucher_PalletGearVoucherInternalVoucherId(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_GearVoucher_RevokeCall(in any) *pbgear.GearVoucher_RevokeCall {
	out := &pbgear.GearVoucher_RevokeCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Spender To_SpCoreCryptoAccountId32(w)
	out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))

	return out //from message
}

func To_GearVoucher_SendMessage(in any) *pbgear.GearVoucher_SendMessage {
	out := &pbgear.GearVoucher_SendMessage{}
	v := in.(registry.Valuable)
	_ = v

	// field Destination To_GprimitivesActorId(w)
	out.Destination = To_GprimitivesActorId(v.ValueAt(0))
	// primitive field Payload
	out.Payload = To_bytes(v.ValueAt(1))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(2))
	// primitive field Value
	out.Value = To_string(v.ValueAt(3))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(4))

	return out //from message
}

func To_GearVoucher_SendReply(in any) *pbgear.GearVoucher_SendReply {
	out := &pbgear.GearVoucher_SendReply{}
	v := in.(registry.Valuable)
	_ = v

	// field ReplyToId To_GprimitivesMessageId(w)
	out.ReplyToId = To_GprimitivesMessageId(v.ValueAt(0))
	// primitive field Payload
	out.Payload = To_bytes(v.ValueAt(1))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(2))
	// primitive field Value
	out.Value = To_string(v.ValueAt(3))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(4))

	return out //from message
}

func To_GearVoucher_Some(in any) *pbgear.GearVoucher_Some {
	out := &pbgear.GearVoucher_Some{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_BTreeSet(w)
	out.Value0 = To_BTreeSet(v.ValueAt(0))

	return out //from message
}

func To_GearVoucher_UpdateCall(in any) *pbgear.GearVoucher_UpdateCall {
	out := &pbgear.GearVoucher_UpdateCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Spender To_SpCoreCryptoAccountId32(w)
	out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
	out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))
	// optional field MoveOwnership
	out.MoveOwnership = To_Optional_GearVoucher_UpdateCall_move_ownership(v.ValueAt(2))
	// primitive field BalanceTopUp
	out.BalanceTopUp = To_Optional_string(v.ValueAt(3))
	// oneOf field AppendPrograms
	v4 := To_OneOf_GearVoucher_UpdateCall_append_programs(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("AppendPrograms").Set(reflect.ValueOf(v4))
	// primitive field CodeUploading
	out.CodeUploading = To_Optional_bool(v.ValueAt(5))
	// primitive field ProlongDuration
	out.ProlongDuration = To_Optional_uint32(v.ValueAt(6))

	return out //from message
}

func To_Optional_GearVoucher_UpdateCall_move_ownership(in any) *pbgear.SpCoreCryptoAccountId32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_GearVoucher_UpdateCall_append_programs(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GearVoucher_UpdateCall_AppendProgramsNone{
			AppendProgramsNone: To_GearVoucher_None(param),
		}
	case 1:
		return &pbgear.GearVoucher_UpdateCall_AppendProgramsSome{
			AppendProgramsSome: To_GearVoucher_Some(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_GearVoucher_UploadCode(in any) *pbgear.GearVoucher_UploadCode {
	out := &pbgear.GearVoucher_UploadCode{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Code
	out.Code = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Gear_ClaimValueCall(in any) *pbgear.Gear_ClaimValueCall {
	out := &pbgear.Gear_ClaimValueCall{}
	v := in.(registry.Valuable)
	_ = v

	// field MessageId To_GprimitivesMessageId(w)
	out.MessageId = To_GprimitivesMessageId(v.ValueAt(0))

	return out //from message
}

func To_Gear_CreateProgramCall(in any) *pbgear.Gear_CreateProgramCall {
	out := &pbgear.Gear_CreateProgramCall{}
	v := in.(registry.Valuable)
	_ = v

	// field CodeId To_GprimitivesCodeId(w)
	out.CodeId = To_GprimitivesCodeId(v.ValueAt(0))
	// primitive field Salt
	out.Salt = To_bytes(v.ValueAt(1))
	// primitive field InitPayload
	out.InitPayload = To_bytes(v.ValueAt(2))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(3))
	// primitive field Value
	out.Value = To_string(v.ValueAt(4))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(5))

	return out //from message
}

func To_Gear_RunCall(in any) *pbgear.Gear_RunCall {
	out := &pbgear.Gear_RunCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field MaxGas
	out.MaxGas = To_Optional_uint64(v.ValueAt(0))

	return out //from message
}

func To_Gear_SendMessageCall(in any) *pbgear.Gear_SendMessageCall {
	out := &pbgear.Gear_SendMessageCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Destination To_GprimitivesActorId(w)
	out.Destination = To_GprimitivesActorId(v.ValueAt(0))
	// primitive field Payload
	out.Payload = To_bytes(v.ValueAt(1))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(2))
	// primitive field Value
	out.Value = To_string(v.ValueAt(3))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(4))

	return out //from message
}

func To_Gear_SendReplyCall(in any) *pbgear.Gear_SendReplyCall {
	out := &pbgear.Gear_SendReplyCall{}
	v := in.(registry.Valuable)
	_ = v

	// field ReplyToId To_GprimitivesMessageId(w)
	out.ReplyToId = To_GprimitivesMessageId(v.ValueAt(0))
	// primitive field Payload
	out.Payload = To_bytes(v.ValueAt(1))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(2))
	// primitive field Value
	out.Value = To_string(v.ValueAt(3))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(4))

	return out //from message
}

func To_Gear_SetExecuteInherentCall(in any) *pbgear.Gear_SetExecuteInherentCall {
	out := &pbgear.Gear_SetExecuteInherentCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_bool(v.ValueAt(0))

	return out //from message
}

func To_Gear_UploadCodeCall(in any) *pbgear.Gear_UploadCodeCall {
	out := &pbgear.Gear_UploadCodeCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Code
	out.Code = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Gear_UploadProgramCall(in any) *pbgear.Gear_UploadProgramCall {
	out := &pbgear.Gear_UploadProgramCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Code
	out.Code = To_bytes(v.ValueAt(0))
	// primitive field Salt
	out.Salt = To_bytes(v.ValueAt(1))
	// primitive field InitPayload
	out.InitPayload = To_bytes(v.ValueAt(2))
	// primitive field GasLimit
	out.GasLimit = To_uint64(v.ValueAt(3))
	// primitive field Value
	out.Value = To_string(v.ValueAt(4))
	// primitive field KeepAlive
	out.KeepAlive = To_bool(v.ValueAt(5))

	return out //from message
}

func To_GeneralAdmin(in any) *pbgear.GeneralAdmin {
	out := &pbgear.GeneralAdmin{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_GprimitivesActorId(in any) *pbgear.GprimitivesActorId {
	out := &pbgear.GprimitivesActorId{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_GprimitivesCodeId(in any) *pbgear.GprimitivesCodeId {
	out := &pbgear.GprimitivesCodeId{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_GprimitivesMessageId(in any) *pbgear.GprimitivesMessageId {
	out := &pbgear.GprimitivesMessageId{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_GrandpaPallet(in any) *pbgear.GrandpaPallet {
	out := &pbgear.GrandpaPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_GrandpaPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_GrandpaPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.GrandpaPallet_CallReportEquivocationCall{
			CallReportEquivocationCall: To_Grandpa_ReportEquivocationCall(param),
		}
	case 1:
		return &pbgear.GrandpaPallet_CallReportEquivocationUnsignedCall{
			CallReportEquivocationUnsignedCall: To_Grandpa_ReportEquivocationUnsignedCall(param),
		}
	case 2:
		return &pbgear.GrandpaPallet_CallNoteStalledCall{
			CallNoteStalledCall: To_Grandpa_NoteStalledCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Grandpa_Equivocation(in any) *pbgear.Grandpa_Equivocation {
	out := &pbgear.Grandpa_Equivocation{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Equivocation
	v0 := To_OneOf_Grandpa_Equivocation_equivocation(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Equivocation").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Grandpa_Equivocation_equivocation(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Grandpa_Equivocation_EquivocationPrevote{
			EquivocationPrevote: To_Grandpa_Prevote(param),
		}
	case 1:
		return &pbgear.Grandpa_Equivocation_EquivocationPrecommit{
			EquivocationPrecommit: To_Grandpa_Precommit(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Grandpa_FinalityGrandpaEquivocation(in any) *pbgear.Grandpa_FinalityGrandpaEquivocation {
	out := &pbgear.Grandpa_FinalityGrandpaEquivocation{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field RoundNumber
	out.RoundNumber = To_uint64(v.ValueAt(0))
	// field Identity To_Grandpa_SpConsensusGrandpaAppPublic(w)
	out.Identity = To_Grandpa_SpConsensusGrandpaAppPublic(v.ValueAt(1))
	// field First To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(w)
	out.First = To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(v.ValueAt(2))
	// field Second To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(w)
	out.Second = To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(v.ValueAt(3))

	return out //from message
}

func To_Grandpa_FinalityGrandpaPrecommit(in any) *pbgear.Grandpa_FinalityGrandpaPrecommit {
	out := &pbgear.Grandpa_FinalityGrandpaPrecommit{}
	v := in.(registry.Valuable)
	_ = v

	// field TargetHash To_PrimitiveTypesH256(w)
	out.TargetHash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field TargetNumber
	out.TargetNumber = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_FinalityGrandpaPrevote(in any) *pbgear.Grandpa_FinalityGrandpaPrevote {
	out := &pbgear.Grandpa_FinalityGrandpaPrevote{}
	v := in.(registry.Valuable)
	_ = v

	// field TargetHash To_PrimitiveTypesH256(w)
	out.TargetHash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field TargetNumber
	out.TargetNumber = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_GrandpaTrieNodesList(in any) *pbgear.Grandpa_GrandpaTrieNodesList {
	out := &pbgear.Grandpa_GrandpaTrieNodesList{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field TrieNodes
	out.TrieNodes = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Grandpa_NoteStalledCall(in any) *pbgear.Grandpa_NoteStalledCall {
	out := &pbgear.Grandpa_NoteStalledCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Delay
	out.Delay = To_uint32(v.ValueAt(0))
	// primitive field BestFinalizedBlockNumber
	out.BestFinalizedBlockNumber = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_Precommit(in any) *pbgear.Grandpa_Precommit {
	out := &pbgear.Grandpa_Precommit{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Grandpa_FinalityGrandpaEquivocation(w)
	out.Value0 = To_Grandpa_FinalityGrandpaEquivocation(v.ValueAt(0))

	return out //from message
}

func To_Grandpa_Prevote(in any) *pbgear.Grandpa_Prevote {
	out := &pbgear.Grandpa_Prevote{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Grandpa_FinalityGrandpaEquivocation(w)
	out.Value0 = To_Grandpa_FinalityGrandpaEquivocation(v.ValueAt(0))

	return out //from message
}

func To_Grandpa_ReportEquivocationCall(in any) *pbgear.Grandpa_ReportEquivocationCall {
	out := &pbgear.Grandpa_ReportEquivocationCall{}
	v := in.(registry.Valuable)
	_ = v

	// field EquivocationProof To_Grandpa_SpConsensusGrandpaEquivocationProof(w)
	out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(v.ValueAt(0))
	// field KeyOwnerProof To_SpSessionMembershipProof(w)
	out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_ReportEquivocationUnsignedCall(in any) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
	out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
	v := in.(registry.Valuable)
	_ = v

	// field EquivocationProof To_Grandpa_SpConsensusGrandpaEquivocationProof(w)
	out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(v.ValueAt(0))
	// field KeyOwnerProof To_SpSessionMembershipProof(w)
	out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_SpConsensusGrandpaAppPublic(in any) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
	out := &pbgear.Grandpa_SpConsensusGrandpaAppPublic{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreEd25519Public(w)
	out.Value0 = To_SpCoreEd25519Public(v.ValueAt(0))

	return out //from message
}

func To_Grandpa_SpConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
	out := &pbgear.Grandpa_SpConsensusGrandpaAppSignature{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreEd25519Signature(w)
	out.Value0 = To_SpCoreEd25519Signature(v.ValueAt(0))

	return out //from message
}

func To_Grandpa_SpConsensusGrandpaEquivocationProof(in any) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
	out := &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field SetId
	out.SetId = To_uint64(v.ValueAt(0))
	// oneOf field Equivocation
	v1 := To_OneOf_Grandpa_SpConsensusGrandpaEquivocationProof_equivocation(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Equivocation").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Grandpa_SpConsensusGrandpaEquivocationProof_equivocation(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof_EquivocationPrevote{
			EquivocationPrevote: To_Grandpa_Prevote(param),
		}
	case 1:
		return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof_EquivocationPrecommit{
			EquivocationPrecommit: To_Grandpa_Precommit(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature {
	out := &pbgear.Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Grandpa_FinalityGrandpaPrecommit(w)
	out.Value0 = To_Grandpa_FinalityGrandpaPrecommit(v.ValueAt(0))
	// field Value1 To_Grandpa_SpConsensusGrandpaAppSignature(w)
	out.Value1 = To_Grandpa_SpConsensusGrandpaAppSignature(v.ValueAt(1))

	return out //from message
}

func To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	out := &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Grandpa_FinalityGrandpaPrevote(w)
	out.Value0 = To_Grandpa_FinalityGrandpaPrevote(v.ValueAt(0))
	// field Value1 To_Grandpa_SpConsensusGrandpaAppSignature(w)
	out.Value1 = To_Grandpa_SpConsensusGrandpaAppSignature(v.ValueAt(1))

	return out //from message
}

func To_HistoricalPallet(in any) *pbgear.HistoricalPallet {
	out := &pbgear.HistoricalPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_IdentityPallet(in any) *pbgear.IdentityPallet {
	out := &pbgear.IdentityPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_IdentityPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_IdentityPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.IdentityPallet_CallAddRegistrarCall{
			CallAddRegistrarCall: To_Identity_AddRegistrarCall(param),
		}
	case 1:
		return &pbgear.IdentityPallet_CallSetIdentityCall{
			CallSetIdentityCall: To_Identity_SetIdentityCall(param),
		}
	case 2:
		return &pbgear.IdentityPallet_CallSetSubsCall{
			CallSetSubsCall: To_Identity_SetSubsCall(param),
		}
	case 3:
		return &pbgear.IdentityPallet_CallClearIdentityCall{
			CallClearIdentityCall: To_Identity_ClearIdentityCall(param),
		}
	case 4:
		return &pbgear.IdentityPallet_CallRequestJudgementCall{
			CallRequestJudgementCall: To_Identity_RequestJudgementCall(param),
		}
	case 5:
		return &pbgear.IdentityPallet_CallCancelRequestCall{
			CallCancelRequestCall: To_Identity_CancelRequestCall(param),
		}
	case 6:
		return &pbgear.IdentityPallet_CallSetFeeCall{
			CallSetFeeCall: To_Identity_SetFeeCall(param),
		}
	case 7:
		return &pbgear.IdentityPallet_CallSetAccountIdCall{
			CallSetAccountIdCall: To_Identity_SetAccountIdCall(param),
		}
	case 8:
		return &pbgear.IdentityPallet_CallSetFieldsCall{
			CallSetFieldsCall: To_Identity_SetFieldsCall(param),
		}
	case 9:
		return &pbgear.IdentityPallet_CallProvideJudgementCall{
			CallProvideJudgementCall: To_Identity_ProvideJudgementCall(param),
		}
	case 10:
		return &pbgear.IdentityPallet_CallKillIdentityCall{
			CallKillIdentityCall: To_Identity_KillIdentityCall(param),
		}
	case 11:
		return &pbgear.IdentityPallet_CallAddSubCall{
			CallAddSubCall: To_Identity_AddSubCall(param),
		}
	case 12:
		return &pbgear.IdentityPallet_CallRenameSubCall{
			CallRenameSubCall: To_Identity_RenameSubCall(param),
		}
	case 13:
		return &pbgear.IdentityPallet_CallRemoveSubCall{
			CallRemoveSubCall: To_Identity_RemoveSubCall(param),
		}
	case 14:
		return &pbgear.IdentityPallet_CallQuitSubCall{
			CallQuitSubCall: To_Identity_QuitSubCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Account(in any) *pbgear.Identity_Account {
	out := &pbgear.Identity_Account{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Account
	v0 := To_OneOf_Identity_Account_account(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Account").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Account_account(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Account_AccountId{
			AccountId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_Account_AccountIndex{
			AccountIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_Account_AccountRaw{
			AccountRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_Account_AccountAddress32{
			AccountAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_Account_AccountAddress20{
			AccountAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_AddRegistrarCall(in any) *pbgear.Identity_AddRegistrarCall {
	out := &pbgear.Identity_AddRegistrarCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Account
	v0 := To_OneOf_Identity_AddRegistrarCall_account(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Account").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_AddRegistrarCall_account(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_AddRegistrarCall_AccountId{
			AccountId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_AddRegistrarCall_AccountIndex{
			AccountIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_AddRegistrarCall_AccountRaw{
			AccountRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_AddRegistrarCall_AccountAddress32{
			AccountAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_AddRegistrarCall_AccountAddress20{
			AccountAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_AddSubCall(in any) *pbgear.Identity_AddSubCall {
	out := &pbgear.Identity_AddSubCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Sub
	v0 := To_OneOf_Identity_AddSubCall_sub(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))
	// oneOf field Data
	v1 := To_OneOf_Identity_AddSubCall_data(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Identity_AddSubCall_sub(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_AddSubCall_SubId{
			SubId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_AddSubCall_SubIndex{
			SubIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_AddSubCall_SubRaw{
			SubRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_AddSubCall_SubAddress32{
			SubAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_AddSubCall_SubAddress20{
			SubAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_AddSubCall_data(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_AddSubCall_DataNone{
			DataNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_AddSubCall_DataRaw0{
			DataRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_AddSubCall_DataRaw1{
			DataRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_AddSubCall_DataRaw2{
			DataRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_AddSubCall_DataRaw3{
			DataRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_AddSubCall_DataRaw4{
			DataRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_AddSubCall_DataRaw5{
			DataRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_AddSubCall_DataRaw6{
			DataRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_AddSubCall_DataRaw7{
			DataRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_AddSubCall_DataRaw8{
			DataRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_AddSubCall_DataRaw9{
			DataRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_AddSubCall_DataRaw10{
			DataRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_AddSubCall_DataRaw11{
			DataRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_AddSubCall_DataRaw12{
			DataRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_AddSubCall_DataRaw13{
			DataRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_AddSubCall_DataRaw14{
			DataRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_AddSubCall_DataRaw15{
			DataRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_AddSubCall_DataRaw16{
			DataRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_AddSubCall_DataRaw17{
			DataRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_AddSubCall_DataRaw18{
			DataRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_AddSubCall_DataRaw19{
			DataRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_AddSubCall_DataRaw20{
			DataRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_AddSubCall_DataRaw21{
			DataRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_AddSubCall_DataRaw22{
			DataRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_AddSubCall_DataRaw23{
			DataRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_AddSubCall_DataRaw24{
			DataRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_AddSubCall_DataRaw25{
			DataRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_AddSubCall_DataRaw26{
			DataRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_AddSubCall_DataRaw27{
			DataRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_AddSubCall_DataRaw28{
			DataRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_AddSubCall_DataRaw29{
			DataRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_AddSubCall_DataRaw30{
			DataRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_AddSubCall_DataRaw31{
			DataRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_AddSubCall_DataRaw32{
			DataRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_AddSubCall_DataBlakeTwo256{
			DataBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_AddSubCall_DataSha256{
			DataSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_AddSubCall_DataKeccak256{
			DataKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_AddSubCall_DataShaThree256{
			DataShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Address20(in any) *pbgear.Identity_Address20 {
	out := &pbgear.Identity_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Address32(in any) *pbgear.Identity_Address32 {
	out := &pbgear.Identity_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_BlakeTwo256(in any) *pbgear.Identity_BlakeTwo256 {
	out := &pbgear.Identity_BlakeTwo256{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_CancelRequestCall(in any) *pbgear.Identity_CancelRequestCall {
	out := &pbgear.Identity_CancelRequestCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field RegIndex
	out.RegIndex = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Identity_ClearIdentityCall(in any) *pbgear.Identity_ClearIdentityCall {
	out := &pbgear.Identity_ClearIdentityCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_Data(in any) *pbgear.Identity_Data {
	out := &pbgear.Identity_Data{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Data
	v0 := To_OneOf_Identity_Data_data(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Data_data(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Data_DataNone{
			DataNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Data_DataRaw0{
			DataRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Data_DataRaw1{
			DataRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Data_DataRaw2{
			DataRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Data_DataRaw3{
			DataRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Data_DataRaw4{
			DataRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Data_DataRaw5{
			DataRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Data_DataRaw6{
			DataRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Data_DataRaw7{
			DataRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Data_DataRaw8{
			DataRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Data_DataRaw9{
			DataRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Data_DataRaw10{
			DataRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Data_DataRaw11{
			DataRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Data_DataRaw12{
			DataRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Data_DataRaw13{
			DataRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Data_DataRaw14{
			DataRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Data_DataRaw15{
			DataRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Data_DataRaw16{
			DataRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Data_DataRaw17{
			DataRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Data_DataRaw18{
			DataRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Data_DataRaw19{
			DataRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Data_DataRaw20{
			DataRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Data_DataRaw21{
			DataRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Data_DataRaw22{
			DataRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Data_DataRaw23{
			DataRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Data_DataRaw24{
			DataRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Data_DataRaw25{
			DataRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Data_DataRaw26{
			DataRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Data_DataRaw27{
			DataRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Data_DataRaw28{
			DataRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Data_DataRaw29{
			DataRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Data_DataRaw30{
			DataRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Data_DataRaw31{
			DataRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Data_DataRaw32{
			DataRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Data_DataBlakeTwo256{
			DataBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Data_DataSha256{
			DataSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Data_DataKeccak256{
			DataKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Data_DataShaThree256{
			DataShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Display(in any) *pbgear.Identity_Display {
	out := &pbgear.Identity_Display{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Display
	v0 := To_OneOf_Identity_Display_display(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Display").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Display_display(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Display_DisplayNone{
			DisplayNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Display_DisplayRaw0{
			DisplayRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Display_DisplayRaw1{
			DisplayRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Display_DisplayRaw2{
			DisplayRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Display_DisplayRaw3{
			DisplayRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Display_DisplayRaw4{
			DisplayRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Display_DisplayRaw5{
			DisplayRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Display_DisplayRaw6{
			DisplayRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Display_DisplayRaw7{
			DisplayRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Display_DisplayRaw8{
			DisplayRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Display_DisplayRaw9{
			DisplayRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Display_DisplayRaw10{
			DisplayRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Display_DisplayRaw11{
			DisplayRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Display_DisplayRaw12{
			DisplayRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Display_DisplayRaw13{
			DisplayRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Display_DisplayRaw14{
			DisplayRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Display_DisplayRaw15{
			DisplayRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Display_DisplayRaw16{
			DisplayRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Display_DisplayRaw17{
			DisplayRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Display_DisplayRaw18{
			DisplayRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Display_DisplayRaw19{
			DisplayRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Display_DisplayRaw20{
			DisplayRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Display_DisplayRaw21{
			DisplayRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Display_DisplayRaw22{
			DisplayRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Display_DisplayRaw23{
			DisplayRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Display_DisplayRaw24{
			DisplayRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Display_DisplayRaw25{
			DisplayRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Display_DisplayRaw26{
			DisplayRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Display_DisplayRaw27{
			DisplayRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Display_DisplayRaw28{
			DisplayRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Display_DisplayRaw29{
			DisplayRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Display_DisplayRaw30{
			DisplayRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Display_DisplayRaw31{
			DisplayRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Display_DisplayRaw32{
			DisplayRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Display_DisplayBlakeTwo256{
			DisplayBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Display_DisplaySha256{
			DisplaySha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Display_DisplayKeccak256{
			DisplayKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Display_DisplayShaThree256{
			DisplayShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Email(in any) *pbgear.Identity_Email {
	out := &pbgear.Identity_Email{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Email
	v0 := To_OneOf_Identity_Email_email(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Email").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Email_email(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Email_EmailNone{
			EmailNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Email_EmailRaw0{
			EmailRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Email_EmailRaw1{
			EmailRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Email_EmailRaw2{
			EmailRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Email_EmailRaw3{
			EmailRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Email_EmailRaw4{
			EmailRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Email_EmailRaw5{
			EmailRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Email_EmailRaw6{
			EmailRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Email_EmailRaw7{
			EmailRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Email_EmailRaw8{
			EmailRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Email_EmailRaw9{
			EmailRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Email_EmailRaw10{
			EmailRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Email_EmailRaw11{
			EmailRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Email_EmailRaw12{
			EmailRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Email_EmailRaw13{
			EmailRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Email_EmailRaw14{
			EmailRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Email_EmailRaw15{
			EmailRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Email_EmailRaw16{
			EmailRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Email_EmailRaw17{
			EmailRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Email_EmailRaw18{
			EmailRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Email_EmailRaw19{
			EmailRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Email_EmailRaw20{
			EmailRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Email_EmailRaw21{
			EmailRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Email_EmailRaw22{
			EmailRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Email_EmailRaw23{
			EmailRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Email_EmailRaw24{
			EmailRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Email_EmailRaw25{
			EmailRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Email_EmailRaw26{
			EmailRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Email_EmailRaw27{
			EmailRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Email_EmailRaw28{
			EmailRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Email_EmailRaw29{
			EmailRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Email_EmailRaw30{
			EmailRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Email_EmailRaw31{
			EmailRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Email_EmailRaw32{
			EmailRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Email_EmailBlakeTwo256{
			EmailBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Email_EmailSha256{
			EmailSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Email_EmailKeccak256{
			EmailKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Email_EmailShaThree256{
			EmailShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Erroneous(in any) *pbgear.Identity_Erroneous {
	out := &pbgear.Identity_Erroneous{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_FeePaid(in any) *pbgear.Identity_FeePaid {
	out := &pbgear.Identity_FeePaid{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_string(v.ValueAt(0))

	return out //from message
}

func To_Identity_Id(in any) *pbgear.Identity_Id {
	out := &pbgear.Identity_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Identity_Image(in any) *pbgear.Identity_Image {
	out := &pbgear.Identity_Image{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Image
	v0 := To_OneOf_Identity_Image_image(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Image").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Image_image(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Image_ImageNone{
			ImageNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Image_ImageRaw0{
			ImageRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Image_ImageRaw1{
			ImageRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Image_ImageRaw2{
			ImageRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Image_ImageRaw3{
			ImageRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Image_ImageRaw4{
			ImageRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Image_ImageRaw5{
			ImageRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Image_ImageRaw6{
			ImageRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Image_ImageRaw7{
			ImageRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Image_ImageRaw8{
			ImageRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Image_ImageRaw9{
			ImageRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Image_ImageRaw10{
			ImageRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Image_ImageRaw11{
			ImageRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Image_ImageRaw12{
			ImageRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Image_ImageRaw13{
			ImageRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Image_ImageRaw14{
			ImageRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Image_ImageRaw15{
			ImageRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Image_ImageRaw16{
			ImageRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Image_ImageRaw17{
			ImageRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Image_ImageRaw18{
			ImageRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Image_ImageRaw19{
			ImageRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Image_ImageRaw20{
			ImageRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Image_ImageRaw21{
			ImageRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Image_ImageRaw22{
			ImageRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Image_ImageRaw23{
			ImageRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Image_ImageRaw24{
			ImageRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Image_ImageRaw25{
			ImageRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Image_ImageRaw26{
			ImageRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Image_ImageRaw27{
			ImageRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Image_ImageRaw28{
			ImageRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Image_ImageRaw29{
			ImageRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Image_ImageRaw30{
			ImageRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Image_ImageRaw31{
			ImageRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Image_ImageRaw32{
			ImageRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Image_ImageBlakeTwo256{
			ImageBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Image_ImageSha256{
			ImageSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Image_ImageKeccak256{
			ImageKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Image_ImageShaThree256{
			ImageShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Index(in any) *pbgear.Identity_Index {
	out := &pbgear.Identity_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Identity_TupleNull(w)
	out.Value0 = To_Identity_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Identity_Judgement(in any) *pbgear.Identity_Judgement {
	out := &pbgear.Identity_Judgement{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Judgement
	v0 := To_OneOf_Identity_Judgement_judgement(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Judgement").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Judgement_judgement(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Judgement_JudgementUnknown{
			JudgementUnknown: To_Identity_Unknown(param),
		}
	case 1:
		return &pbgear.Identity_Judgement_JudgementFeePaid{
			JudgementFeePaid: To_Identity_FeePaid(param),
		}
	case 2:
		return &pbgear.Identity_Judgement_JudgementReasonable{
			JudgementReasonable: To_Identity_Reasonable(param),
		}
	case 3:
		return &pbgear.Identity_Judgement_JudgementKnownGood{
			JudgementKnownGood: To_Identity_KnownGood(param),
		}
	case 4:
		return &pbgear.Identity_Judgement_JudgementOutOfDate{
			JudgementOutOfDate: To_Identity_OutOfDate(param),
		}
	case 5:
		return &pbgear.Identity_Judgement_JudgementLowQuality{
			JudgementLowQuality: To_Identity_LowQuality(param),
		}
	case 6:
		return &pbgear.Identity_Judgement_JudgementErroneous{
			JudgementErroneous: To_Identity_Erroneous(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Keccak256(in any) *pbgear.Identity_Keccak256 {
	out := &pbgear.Identity_Keccak256{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_KillIdentityCall(in any) *pbgear.Identity_KillIdentityCall {
	out := &pbgear.Identity_KillIdentityCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_Identity_KillIdentityCall_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_KillIdentityCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_KillIdentityCall_TargetId{
			TargetId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_KillIdentityCall_TargetIndex{
			TargetIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_KillIdentityCall_TargetRaw{
			TargetRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_KillIdentityCall_TargetAddress32{
			TargetAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_KillIdentityCall_TargetAddress20{
			TargetAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_KnownGood(in any) *pbgear.Identity_KnownGood {
	out := &pbgear.Identity_KnownGood{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_Legal(in any) *pbgear.Identity_Legal {
	out := &pbgear.Identity_Legal{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Legal
	v0 := To_OneOf_Identity_Legal_legal(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Legal").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Legal_legal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Legal_LegalNone{
			LegalNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Legal_LegalRaw0{
			LegalRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Legal_LegalRaw1{
			LegalRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Legal_LegalRaw2{
			LegalRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Legal_LegalRaw3{
			LegalRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Legal_LegalRaw4{
			LegalRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Legal_LegalRaw5{
			LegalRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Legal_LegalRaw6{
			LegalRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Legal_LegalRaw7{
			LegalRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Legal_LegalRaw8{
			LegalRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Legal_LegalRaw9{
			LegalRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Legal_LegalRaw10{
			LegalRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Legal_LegalRaw11{
			LegalRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Legal_LegalRaw12{
			LegalRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Legal_LegalRaw13{
			LegalRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Legal_LegalRaw14{
			LegalRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Legal_LegalRaw15{
			LegalRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Legal_LegalRaw16{
			LegalRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Legal_LegalRaw17{
			LegalRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Legal_LegalRaw18{
			LegalRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Legal_LegalRaw19{
			LegalRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Legal_LegalRaw20{
			LegalRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Legal_LegalRaw21{
			LegalRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Legal_LegalRaw22{
			LegalRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Legal_LegalRaw23{
			LegalRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Legal_LegalRaw24{
			LegalRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Legal_LegalRaw25{
			LegalRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Legal_LegalRaw26{
			LegalRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Legal_LegalRaw27{
			LegalRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Legal_LegalRaw28{
			LegalRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Legal_LegalRaw29{
			LegalRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Legal_LegalRaw30{
			LegalRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Legal_LegalRaw31{
			LegalRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Legal_LegalRaw32{
			LegalRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Legal_LegalBlakeTwo256{
			LegalBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Legal_LegalSha256{
			LegalSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Legal_LegalKeccak256{
			LegalKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Legal_LegalShaThree256{
			LegalShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_LowQuality(in any) *pbgear.Identity_LowQuality {
	out := &pbgear.Identity_LowQuality{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_New(in any) *pbgear.Identity_New {
	out := &pbgear.Identity_New{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field New
	v0 := To_OneOf_Identity_New_new(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("New").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_New_new(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_New_NewId{
			NewId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_New_NewIndex{
			NewIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_New_NewRaw{
			NewRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_New_NewAddress32{
			NewAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_New_NewAddress20{
			NewAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_None(in any) *pbgear.Identity_None {
	out := &pbgear.Identity_None{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_OutOfDate(in any) *pbgear.Identity_OutOfDate {
	out := &pbgear.Identity_OutOfDate{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_PalletIdentitySimpleIdentityInfo(in any) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
	v := in.(registry.Valuable)
	_ = v

	// field Additional To_BoundedCollectionsBoundedVecBoundedVec(w)
	out.Additional = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))
	// oneOf field Display
	v1 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_display(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Display").Set(reflect.ValueOf(v1))
	// oneOf field Legal
	v2 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_legal(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Legal").Set(reflect.ValueOf(v2))
	// oneOf field Web
	v3 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_web(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Web").Set(reflect.ValueOf(v3))
	// oneOf field Riot
	v4 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_riot(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("Riot").Set(reflect.ValueOf(v4))
	// oneOf field Email
	v5 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_email(v.ValueAt(5))
	reflect.ValueOf(out).Elem().FieldByName("Email").Set(reflect.ValueOf(v5))
	// primitive field PgpFingerprint
	out.PgpFingerprint = To_bytes(v.ValueAt(6))
	// oneOf field Image
	v7 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_image(v.ValueAt(7))
	reflect.ValueOf(out).Elem().FieldByName("Image").Set(reflect.ValueOf(v7))
	// oneOf field Twitter
	v8 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_twitter(v.ValueAt(8))
	reflect.ValueOf(out).Elem().FieldByName("Twitter").Set(reflect.ValueOf(v8))

	return out //from message
}

func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_display(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayNone{
			DisplayNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw0{
			DisplayRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw1{
			DisplayRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw2{
			DisplayRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw3{
			DisplayRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw4{
			DisplayRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw5{
			DisplayRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw6{
			DisplayRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw7{
			DisplayRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw8{
			DisplayRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw9{
			DisplayRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw10{
			DisplayRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw11{
			DisplayRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw12{
			DisplayRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw13{
			DisplayRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw14{
			DisplayRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw15{
			DisplayRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw16{
			DisplayRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw17{
			DisplayRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw18{
			DisplayRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw19{
			DisplayRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw20{
			DisplayRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw21{
			DisplayRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw22{
			DisplayRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw23{
			DisplayRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw24{
			DisplayRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw25{
			DisplayRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw26{
			DisplayRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw27{
			DisplayRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw28{
			DisplayRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw29{
			DisplayRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw30{
			DisplayRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw31{
			DisplayRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw32{
			DisplayRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayBlakeTwo256{
			DisplayBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplaySha256{
			DisplaySha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayKeccak256{
			DisplayKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayShaThree256{
			DisplayShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_legal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalNone{
			LegalNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw0{
			LegalRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw1{
			LegalRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw2{
			LegalRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw3{
			LegalRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw4{
			LegalRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw5{
			LegalRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw6{
			LegalRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw7{
			LegalRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw8{
			LegalRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw9{
			LegalRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw10{
			LegalRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw11{
			LegalRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw12{
			LegalRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw13{
			LegalRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw14{
			LegalRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw15{
			LegalRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw16{
			LegalRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw17{
			LegalRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw18{
			LegalRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw19{
			LegalRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw20{
			LegalRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw21{
			LegalRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw22{
			LegalRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw23{
			LegalRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw24{
			LegalRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw25{
			LegalRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw26{
			LegalRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw27{
			LegalRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw28{
			LegalRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw29{
			LegalRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw30{
			LegalRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw31{
			LegalRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw32{
			LegalRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalBlakeTwo256{
			LegalBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalSha256{
			LegalSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalKeccak256{
			LegalKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalShaThree256{
			LegalShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_web(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebNone{
			WebNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw0{
			WebRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw1{
			WebRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw2{
			WebRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw3{
			WebRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw4{
			WebRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw5{
			WebRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw6{
			WebRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw7{
			WebRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw8{
			WebRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw9{
			WebRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw10{
			WebRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw11{
			WebRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw12{
			WebRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw13{
			WebRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw14{
			WebRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw15{
			WebRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw16{
			WebRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw17{
			WebRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw18{
			WebRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw19{
			WebRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw20{
			WebRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw21{
			WebRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw22{
			WebRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw23{
			WebRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw24{
			WebRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw25{
			WebRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw26{
			WebRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw27{
			WebRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw28{
			WebRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw29{
			WebRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw30{
			WebRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw31{
			WebRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw32{
			WebRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebBlakeTwo256{
			WebBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebSha256{
			WebSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebKeccak256{
			WebKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebShaThree256{
			WebShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_riot(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotNone{
			RiotNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw0{
			RiotRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw1{
			RiotRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw2{
			RiotRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw3{
			RiotRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw4{
			RiotRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw5{
			RiotRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw6{
			RiotRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw7{
			RiotRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw8{
			RiotRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw9{
			RiotRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw10{
			RiotRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw11{
			RiotRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw12{
			RiotRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw13{
			RiotRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw14{
			RiotRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw15{
			RiotRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw16{
			RiotRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw17{
			RiotRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw18{
			RiotRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw19{
			RiotRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw20{
			RiotRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw21{
			RiotRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw22{
			RiotRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw23{
			RiotRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw24{
			RiotRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw25{
			RiotRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw26{
			RiotRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw27{
			RiotRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw28{
			RiotRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw29{
			RiotRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw30{
			RiotRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw31{
			RiotRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw32{
			RiotRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotBlakeTwo256{
			RiotBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotSha256{
			RiotSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotKeccak256{
			RiotKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotShaThree256{
			RiotShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_email(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailNone{
			EmailNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw0{
			EmailRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw1{
			EmailRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw2{
			EmailRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw3{
			EmailRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw4{
			EmailRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw5{
			EmailRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw6{
			EmailRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw7{
			EmailRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw8{
			EmailRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw9{
			EmailRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw10{
			EmailRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw11{
			EmailRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw12{
			EmailRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw13{
			EmailRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw14{
			EmailRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw15{
			EmailRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw16{
			EmailRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw17{
			EmailRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw18{
			EmailRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw19{
			EmailRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw20{
			EmailRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw21{
			EmailRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw22{
			EmailRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw23{
			EmailRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw24{
			EmailRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw25{
			EmailRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw26{
			EmailRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw27{
			EmailRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw28{
			EmailRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw29{
			EmailRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw30{
			EmailRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw31{
			EmailRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw32{
			EmailRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailBlakeTwo256{
			EmailBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailSha256{
			EmailSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailKeccak256{
			EmailKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailShaThree256{
			EmailShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_image(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageNone{
			ImageNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw0{
			ImageRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw1{
			ImageRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw2{
			ImageRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw3{
			ImageRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw4{
			ImageRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw5{
			ImageRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw6{
			ImageRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw7{
			ImageRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw8{
			ImageRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw9{
			ImageRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw10{
			ImageRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw11{
			ImageRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw12{
			ImageRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw13{
			ImageRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw14{
			ImageRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw15{
			ImageRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw16{
			ImageRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw17{
			ImageRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw18{
			ImageRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw19{
			ImageRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw20{
			ImageRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw21{
			ImageRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw22{
			ImageRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw23{
			ImageRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw24{
			ImageRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw25{
			ImageRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw26{
			ImageRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw27{
			ImageRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw28{
			ImageRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw29{
			ImageRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw30{
			ImageRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw31{
			ImageRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw32{
			ImageRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageBlakeTwo256{
			ImageBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageSha256{
			ImageSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageKeccak256{
			ImageKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageShaThree256{
			ImageShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_twitter(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterNone{
			TwitterNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw0{
			TwitterRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw1{
			TwitterRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw2{
			TwitterRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw3{
			TwitterRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw4{
			TwitterRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw5{
			TwitterRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw6{
			TwitterRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw7{
			TwitterRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw8{
			TwitterRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw9{
			TwitterRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw10{
			TwitterRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw11{
			TwitterRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw12{
			TwitterRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw13{
			TwitterRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw14{
			TwitterRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw15{
			TwitterRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw16{
			TwitterRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw17{
			TwitterRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw18{
			TwitterRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw19{
			TwitterRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw20{
			TwitterRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw21{
			TwitterRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw22{
			TwitterRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw23{
			TwitterRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw24{
			TwitterRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw25{
			TwitterRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw26{
			TwitterRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw27{
			TwitterRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw28{
			TwitterRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw29{
			TwitterRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw30{
			TwitterRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw31{
			TwitterRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw32{
			TwitterRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterBlakeTwo256{
			TwitterBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterSha256{
			TwitterSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterKeccak256{
			TwitterKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterShaThree256{
			TwitterShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_PalletIdentityTypesBitFlags(in any) *pbgear.Identity_PalletIdentityTypesBitFlags {
	out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint64(v.ValueAt(0))

	return out //from message
}

func To_Identity_ProvideJudgementCall(in any) *pbgear.Identity_ProvideJudgementCall {
	out := &pbgear.Identity_ProvideJudgementCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field RegIndex
	out.RegIndex = To_uint32(v.ValueAt(0))
	// oneOf field Target
	v1 := To_OneOf_Identity_ProvideJudgementCall_target(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))
	// oneOf field Judgement
	v2 := To_OneOf_Identity_ProvideJudgementCall_judgement(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Judgement").Set(reflect.ValueOf(v2))
	// field Identity To_PrimitiveTypesH256(w)
	out.Identity = To_PrimitiveTypesH256(v.ValueAt(3))

	return out //from message
}

func To_OneOf_Identity_ProvideJudgementCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_ProvideJudgementCall_TargetId{
			TargetId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_ProvideJudgementCall_TargetIndex{
			TargetIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_ProvideJudgementCall_TargetRaw{
			TargetRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_ProvideJudgementCall_TargetAddress32{
			TargetAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_ProvideJudgementCall_TargetAddress20{
			TargetAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_ProvideJudgementCall_judgement(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_ProvideJudgementCall_JudgementUnknown{
			JudgementUnknown: To_Identity_Unknown(param),
		}
	case 1:
		return &pbgear.Identity_ProvideJudgementCall_JudgementFeePaid{
			JudgementFeePaid: To_Identity_FeePaid(param),
		}
	case 2:
		return &pbgear.Identity_ProvideJudgementCall_JudgementReasonable{
			JudgementReasonable: To_Identity_Reasonable(param),
		}
	case 3:
		return &pbgear.Identity_ProvideJudgementCall_JudgementKnownGood{
			JudgementKnownGood: To_Identity_KnownGood(param),
		}
	case 4:
		return &pbgear.Identity_ProvideJudgementCall_JudgementOutOfDate{
			JudgementOutOfDate: To_Identity_OutOfDate(param),
		}
	case 5:
		return &pbgear.Identity_ProvideJudgementCall_JudgementLowQuality{
			JudgementLowQuality: To_Identity_LowQuality(param),
		}
	case 6:
		return &pbgear.Identity_ProvideJudgementCall_JudgementErroneous{
			JudgementErroneous: To_Identity_Erroneous(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Identity_QuitSubCall(in any) *pbgear.Identity_QuitSubCall {
	out := &pbgear.Identity_QuitSubCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_Raw(in any) *pbgear.Identity_Raw {
	out := &pbgear.Identity_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw0(in any) *pbgear.Identity_Raw0 {
	out := &pbgear.Identity_Raw0{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw1(in any) *pbgear.Identity_Raw1 {
	out := &pbgear.Identity_Raw1{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw10(in any) *pbgear.Identity_Raw10 {
	out := &pbgear.Identity_Raw10{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw11(in any) *pbgear.Identity_Raw11 {
	out := &pbgear.Identity_Raw11{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw12(in any) *pbgear.Identity_Raw12 {
	out := &pbgear.Identity_Raw12{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw13(in any) *pbgear.Identity_Raw13 {
	out := &pbgear.Identity_Raw13{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw14(in any) *pbgear.Identity_Raw14 {
	out := &pbgear.Identity_Raw14{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw15(in any) *pbgear.Identity_Raw15 {
	out := &pbgear.Identity_Raw15{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw16(in any) *pbgear.Identity_Raw16 {
	out := &pbgear.Identity_Raw16{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw17(in any) *pbgear.Identity_Raw17 {
	out := &pbgear.Identity_Raw17{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw18(in any) *pbgear.Identity_Raw18 {
	out := &pbgear.Identity_Raw18{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw19(in any) *pbgear.Identity_Raw19 {
	out := &pbgear.Identity_Raw19{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw2(in any) *pbgear.Identity_Raw2 {
	out := &pbgear.Identity_Raw2{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw20(in any) *pbgear.Identity_Raw20 {
	out := &pbgear.Identity_Raw20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw21(in any) *pbgear.Identity_Raw21 {
	out := &pbgear.Identity_Raw21{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw22(in any) *pbgear.Identity_Raw22 {
	out := &pbgear.Identity_Raw22{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw23(in any) *pbgear.Identity_Raw23 {
	out := &pbgear.Identity_Raw23{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw24(in any) *pbgear.Identity_Raw24 {
	out := &pbgear.Identity_Raw24{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw25(in any) *pbgear.Identity_Raw25 {
	out := &pbgear.Identity_Raw25{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw26(in any) *pbgear.Identity_Raw26 {
	out := &pbgear.Identity_Raw26{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw27(in any) *pbgear.Identity_Raw27 {
	out := &pbgear.Identity_Raw27{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw28(in any) *pbgear.Identity_Raw28 {
	out := &pbgear.Identity_Raw28{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw29(in any) *pbgear.Identity_Raw29 {
	out := &pbgear.Identity_Raw29{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw3(in any) *pbgear.Identity_Raw3 {
	out := &pbgear.Identity_Raw3{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw30(in any) *pbgear.Identity_Raw30 {
	out := &pbgear.Identity_Raw30{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw31(in any) *pbgear.Identity_Raw31 {
	out := &pbgear.Identity_Raw31{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw32(in any) *pbgear.Identity_Raw32 {
	out := &pbgear.Identity_Raw32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw4(in any) *pbgear.Identity_Raw4 {
	out := &pbgear.Identity_Raw4{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw5(in any) *pbgear.Identity_Raw5 {
	out := &pbgear.Identity_Raw5{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw6(in any) *pbgear.Identity_Raw6 {
	out := &pbgear.Identity_Raw6{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw7(in any) *pbgear.Identity_Raw7 {
	out := &pbgear.Identity_Raw7{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw8(in any) *pbgear.Identity_Raw8 {
	out := &pbgear.Identity_Raw8{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Raw9(in any) *pbgear.Identity_Raw9 {
	out := &pbgear.Identity_Raw9{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Reasonable(in any) *pbgear.Identity_Reasonable {
	out := &pbgear.Identity_Reasonable{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_RemoveSubCall(in any) *pbgear.Identity_RemoveSubCall {
	out := &pbgear.Identity_RemoveSubCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Sub
	v0 := To_OneOf_Identity_RemoveSubCall_sub(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_RemoveSubCall_sub(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_RemoveSubCall_SubId{
			SubId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_RemoveSubCall_SubIndex{
			SubIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_RemoveSubCall_SubRaw{
			SubRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_RemoveSubCall_SubAddress32{
			SubAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_RemoveSubCall_SubAddress20{
			SubAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_RenameSubCall(in any) *pbgear.Identity_RenameSubCall {
	out := &pbgear.Identity_RenameSubCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Sub
	v0 := To_OneOf_Identity_RenameSubCall_sub(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))
	// oneOf field Data
	v1 := To_OneOf_Identity_RenameSubCall_data(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Identity_RenameSubCall_sub(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_RenameSubCall_SubId{
			SubId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_RenameSubCall_SubIndex{
			SubIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_RenameSubCall_SubRaw{
			SubRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_RenameSubCall_SubAddress32{
			SubAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_RenameSubCall_SubAddress20{
			SubAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_RenameSubCall_data(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_RenameSubCall_DataNone{
			DataNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_RenameSubCall_DataRaw0{
			DataRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_RenameSubCall_DataRaw1{
			DataRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_RenameSubCall_DataRaw2{
			DataRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_RenameSubCall_DataRaw3{
			DataRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_RenameSubCall_DataRaw4{
			DataRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_RenameSubCall_DataRaw5{
			DataRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_RenameSubCall_DataRaw6{
			DataRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_RenameSubCall_DataRaw7{
			DataRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_RenameSubCall_DataRaw8{
			DataRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_RenameSubCall_DataRaw9{
			DataRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_RenameSubCall_DataRaw10{
			DataRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_RenameSubCall_DataRaw11{
			DataRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_RenameSubCall_DataRaw12{
			DataRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_RenameSubCall_DataRaw13{
			DataRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_RenameSubCall_DataRaw14{
			DataRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_RenameSubCall_DataRaw15{
			DataRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_RenameSubCall_DataRaw16{
			DataRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_RenameSubCall_DataRaw17{
			DataRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_RenameSubCall_DataRaw18{
			DataRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_RenameSubCall_DataRaw19{
			DataRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_RenameSubCall_DataRaw20{
			DataRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_RenameSubCall_DataRaw21{
			DataRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_RenameSubCall_DataRaw22{
			DataRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_RenameSubCall_DataRaw23{
			DataRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_RenameSubCall_DataRaw24{
			DataRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_RenameSubCall_DataRaw25{
			DataRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_RenameSubCall_DataRaw26{
			DataRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_RenameSubCall_DataRaw27{
			DataRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_RenameSubCall_DataRaw28{
			DataRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_RenameSubCall_DataRaw29{
			DataRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_RenameSubCall_DataRaw30{
			DataRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_RenameSubCall_DataRaw31{
			DataRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_RenameSubCall_DataRaw32{
			DataRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_RenameSubCall_DataBlakeTwo256{
			DataBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_RenameSubCall_DataSha256{
			DataSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_RenameSubCall_DataKeccak256{
			DataKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_RenameSubCall_DataShaThree256{
			DataShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_RequestJudgementCall(in any) *pbgear.Identity_RequestJudgementCall {
	out := &pbgear.Identity_RequestJudgementCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field RegIndex
	out.RegIndex = To_uint32(v.ValueAt(0))
	// primitive field MaxFee
	out.MaxFee = To_string(v.ValueAt(1))

	return out //from message
}

func To_Identity_Riot(in any) *pbgear.Identity_Riot {
	out := &pbgear.Identity_Riot{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Riot
	v0 := To_OneOf_Identity_Riot_riot(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Riot").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Riot_riot(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Riot_RiotNone{
			RiotNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Riot_RiotRaw0{
			RiotRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Riot_RiotRaw1{
			RiotRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Riot_RiotRaw2{
			RiotRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Riot_RiotRaw3{
			RiotRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Riot_RiotRaw4{
			RiotRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Riot_RiotRaw5{
			RiotRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Riot_RiotRaw6{
			RiotRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Riot_RiotRaw7{
			RiotRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Riot_RiotRaw8{
			RiotRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Riot_RiotRaw9{
			RiotRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Riot_RiotRaw10{
			RiotRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Riot_RiotRaw11{
			RiotRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Riot_RiotRaw12{
			RiotRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Riot_RiotRaw13{
			RiotRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Riot_RiotRaw14{
			RiotRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Riot_RiotRaw15{
			RiotRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Riot_RiotRaw16{
			RiotRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Riot_RiotRaw17{
			RiotRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Riot_RiotRaw18{
			RiotRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Riot_RiotRaw19{
			RiotRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Riot_RiotRaw20{
			RiotRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Riot_RiotRaw21{
			RiotRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Riot_RiotRaw22{
			RiotRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Riot_RiotRaw23{
			RiotRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Riot_RiotRaw24{
			RiotRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Riot_RiotRaw25{
			RiotRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Riot_RiotRaw26{
			RiotRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Riot_RiotRaw27{
			RiotRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Riot_RiotRaw28{
			RiotRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Riot_RiotRaw29{
			RiotRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Riot_RiotRaw30{
			RiotRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Riot_RiotRaw31{
			RiotRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Riot_RiotRaw32{
			RiotRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Riot_RiotBlakeTwo256{
			RiotBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Riot_RiotSha256{
			RiotSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Riot_RiotKeccak256{
			RiotKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Riot_RiotShaThree256{
			RiotShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_SetAccountIdCall(in any) *pbgear.Identity_SetAccountIdCall {
	out := &pbgear.Identity_SetAccountIdCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// oneOf field New
	v1 := To_OneOf_Identity_SetAccountIdCall_new(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("New").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Identity_SetAccountIdCall_new(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_SetAccountIdCall_NewId{
			NewId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_SetAccountIdCall_NewIndex{
			NewIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_SetAccountIdCall_NewRaw{
			NewRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_SetAccountIdCall_NewAddress32{
			NewAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_SetAccountIdCall_NewAddress20{
			NewAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_SetFeeCall(in any) *pbgear.Identity_SetFeeCall {
	out := &pbgear.Identity_SetFeeCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// primitive field Fee
	out.Fee = To_string(v.ValueAt(1))

	return out //from message
}

func To_Identity_SetFieldsCall(in any) *pbgear.Identity_SetFieldsCall {
	out := &pbgear.Identity_SetFieldsCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// field Fields To_Identity_PalletIdentityTypesBitFlags(w)
	out.Fields = To_Identity_PalletIdentityTypesBitFlags(v.ValueAt(1))

	return out //from message
}

func To_Identity_SetIdentityCall(in any) *pbgear.Identity_SetIdentityCall {
	out := &pbgear.Identity_SetIdentityCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Info To_Identity_PalletIdentitySimpleIdentityInfo(w)
	out.Info = To_Identity_PalletIdentitySimpleIdentityInfo(v.ValueAt(0))

	return out //from message
}

func To_Identity_SetSubsCall(in any) *pbgear.Identity_SetSubsCall {
	out := &pbgear.Identity_SetSubsCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Subs
	out.Subs = To_Repeated_Identity_SetSubsCall_subs(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Identity_SetSubsCall_subs(in any) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	items := in.([]interface{})

	var out []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Identity_Sha256(in any) *pbgear.Identity_Sha256 {
	out := &pbgear.Identity_Sha256{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_ShaThree256(in any) *pbgear.Identity_ShaThree256 {
	out := &pbgear.Identity_ShaThree256{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Identity_Sub(in any) *pbgear.Identity_Sub {
	out := &pbgear.Identity_Sub{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Sub
	v0 := To_OneOf_Identity_Sub_sub(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Sub_sub(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Sub_SubId{
			SubId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_Sub_SubIndex{
			SubIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_Sub_SubRaw{
			SubRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_Sub_SubAddress32{
			SubAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_Sub_SubAddress20{
			SubAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Target(in any) *pbgear.Identity_Target {
	out := &pbgear.Identity_Target{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_Identity_Target_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Target_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Target_TargetId{
			TargetId: To_Identity_Id(param),
		}
	case 1:
		return &pbgear.Identity_Target_TargetIndex{
			TargetIndex: To_Identity_Index(param),
		}
	case 2:
		return &pbgear.Identity_Target_TargetRaw{
			TargetRaw: To_Identity_Raw(param),
		}
	case 3:
		return &pbgear.Identity_Target_TargetAddress32{
			TargetAddress32: To_Identity_Address32(param),
		}
	case 4:
		return &pbgear.Identity_Target_TargetAddress20{
			TargetAddress20: To_Identity_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_TupleNull(in any) *pbgear.Identity_TupleNull {
	out := &pbgear.Identity_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(in any) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))
	// oneOf field Value1
	v1 := To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0None{
			Value0None: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw0{
			Value0Raw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw1{
			Value0Raw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw2{
			Value0Raw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw3{
			Value0Raw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw4{
			Value0Raw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw5{
			Value0Raw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw6{
			Value0Raw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw7{
			Value0Raw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw8{
			Value0Raw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw9{
			Value0Raw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw10{
			Value0Raw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw11{
			Value0Raw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw12{
			Value0Raw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw13{
			Value0Raw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw14{
			Value0Raw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw15{
			Value0Raw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw16{
			Value0Raw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw17{
			Value0Raw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw18{
			Value0Raw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw19{
			Value0Raw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw20{
			Value0Raw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw21{
			Value0Raw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw22{
			Value0Raw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw23{
			Value0Raw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw24{
			Value0Raw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw25{
			Value0Raw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw26{
			Value0Raw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw27{
			Value0Raw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw28{
			Value0Raw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw29{
			Value0Raw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw30{
			Value0Raw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw31{
			Value0Raw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw32{
			Value0Raw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0BlakeTwo256{
			Value0BlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Sha256{
			Value0Sha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Keccak256{
			Value0Keccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0ShaThree256{
			Value0ShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1None{
			Value1None: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw0{
			Value1Raw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw1{
			Value1Raw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw2{
			Value1Raw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw3{
			Value1Raw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw4{
			Value1Raw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw5{
			Value1Raw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw6{
			Value1Raw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw7{
			Value1Raw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw8{
			Value1Raw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw9{
			Value1Raw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw10{
			Value1Raw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw11{
			Value1Raw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw12{
			Value1Raw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw13{
			Value1Raw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw14{
			Value1Raw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw15{
			Value1Raw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw16{
			Value1Raw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw17{
			Value1Raw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw18{
			Value1Raw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw19{
			Value1Raw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw20{
			Value1Raw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw21{
			Value1Raw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw22{
			Value1Raw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw23{
			Value1Raw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw24{
			Value1Raw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw25{
			Value1Raw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw26{
			Value1Raw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw27{
			Value1Raw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw28{
			Value1Raw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw29{
			Value1Raw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw30{
			Value1Raw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw31{
			Value1Raw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw32{
			Value1Raw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1BlakeTwo256{
			Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Sha256{
			Value1Sha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Keccak256{
			Value1Keccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1ShaThree256{
			Value1ShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(in any) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// oneOf field Value1
	v1 := To_OneOf_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1None{
			Value1None: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw0{
			Value1Raw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw1{
			Value1Raw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw2{
			Value1Raw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw3{
			Value1Raw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw4{
			Value1Raw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw5{
			Value1Raw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw6{
			Value1Raw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw7{
			Value1Raw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw8{
			Value1Raw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw9{
			Value1Raw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw10{
			Value1Raw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw11{
			Value1Raw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw12{
			Value1Raw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw13{
			Value1Raw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw14{
			Value1Raw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw15{
			Value1Raw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw16{
			Value1Raw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw17{
			Value1Raw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw18{
			Value1Raw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw19{
			Value1Raw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw20{
			Value1Raw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw21{
			Value1Raw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw22{
			Value1Raw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw23{
			Value1Raw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw24{
			Value1Raw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw25{
			Value1Raw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw26{
			Value1Raw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw27{
			Value1Raw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw28{
			Value1Raw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw29{
			Value1Raw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw30{
			Value1Raw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw31{
			Value1Raw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw32{
			Value1Raw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1BlakeTwo256{
			Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Sha256{
			Value1Sha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Keccak256{
			Value1Keccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1ShaThree256{
			Value1ShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Twitter(in any) *pbgear.Identity_Twitter {
	out := &pbgear.Identity_Twitter{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Twitter
	v0 := To_OneOf_Identity_Twitter_twitter(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Twitter").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Twitter_twitter(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Twitter_TwitterNone{
			TwitterNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Twitter_TwitterRaw0{
			TwitterRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Twitter_TwitterRaw1{
			TwitterRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Twitter_TwitterRaw2{
			TwitterRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Twitter_TwitterRaw3{
			TwitterRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Twitter_TwitterRaw4{
			TwitterRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Twitter_TwitterRaw5{
			TwitterRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Twitter_TwitterRaw6{
			TwitterRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Twitter_TwitterRaw7{
			TwitterRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Twitter_TwitterRaw8{
			TwitterRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Twitter_TwitterRaw9{
			TwitterRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Twitter_TwitterRaw10{
			TwitterRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Twitter_TwitterRaw11{
			TwitterRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Twitter_TwitterRaw12{
			TwitterRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Twitter_TwitterRaw13{
			TwitterRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Twitter_TwitterRaw14{
			TwitterRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Twitter_TwitterRaw15{
			TwitterRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Twitter_TwitterRaw16{
			TwitterRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Twitter_TwitterRaw17{
			TwitterRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Twitter_TwitterRaw18{
			TwitterRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Twitter_TwitterRaw19{
			TwitterRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Twitter_TwitterRaw20{
			TwitterRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Twitter_TwitterRaw21{
			TwitterRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Twitter_TwitterRaw22{
			TwitterRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Twitter_TwitterRaw23{
			TwitterRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Twitter_TwitterRaw24{
			TwitterRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Twitter_TwitterRaw25{
			TwitterRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Twitter_TwitterRaw26{
			TwitterRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Twitter_TwitterRaw27{
			TwitterRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Twitter_TwitterRaw28{
			TwitterRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Twitter_TwitterRaw29{
			TwitterRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Twitter_TwitterRaw30{
			TwitterRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Twitter_TwitterRaw31{
			TwitterRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Twitter_TwitterRaw32{
			TwitterRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Twitter_TwitterBlakeTwo256{
			TwitterBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Twitter_TwitterSha256{
			TwitterSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Twitter_TwitterKeccak256{
			TwitterKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Twitter_TwitterShaThree256{
			TwitterShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Unknown(in any) *pbgear.Identity_Unknown {
	out := &pbgear.Identity_Unknown{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Identity_Value0(in any) *pbgear.Identity_Value0 {
	out := &pbgear.Identity_Value0{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Identity_Value0_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Value0_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Value0_Value0None{
			Value0None: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Value0_Value0Raw0{
			Value0Raw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Value0_Value0Raw1{
			Value0Raw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Value0_Value0Raw2{
			Value0Raw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Value0_Value0Raw3{
			Value0Raw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Value0_Value0Raw4{
			Value0Raw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Value0_Value0Raw5{
			Value0Raw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Value0_Value0Raw6{
			Value0Raw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Value0_Value0Raw7{
			Value0Raw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Value0_Value0Raw8{
			Value0Raw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Value0_Value0Raw9{
			Value0Raw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Value0_Value0Raw10{
			Value0Raw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Value0_Value0Raw11{
			Value0Raw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Value0_Value0Raw12{
			Value0Raw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Value0_Value0Raw13{
			Value0Raw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Value0_Value0Raw14{
			Value0Raw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Value0_Value0Raw15{
			Value0Raw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Value0_Value0Raw16{
			Value0Raw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Value0_Value0Raw17{
			Value0Raw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Value0_Value0Raw18{
			Value0Raw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Value0_Value0Raw19{
			Value0Raw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Value0_Value0Raw20{
			Value0Raw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Value0_Value0Raw21{
			Value0Raw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Value0_Value0Raw22{
			Value0Raw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Value0_Value0Raw23{
			Value0Raw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Value0_Value0Raw24{
			Value0Raw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Value0_Value0Raw25{
			Value0Raw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Value0_Value0Raw26{
			Value0Raw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Value0_Value0Raw27{
			Value0Raw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Value0_Value0Raw28{
			Value0Raw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Value0_Value0Raw29{
			Value0Raw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Value0_Value0Raw30{
			Value0Raw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Value0_Value0Raw31{
			Value0Raw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Value0_Value0Raw32{
			Value0Raw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Value0_Value0BlakeTwo256{
			Value0BlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Value0_Value0Sha256{
			Value0Sha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Value0_Value0Keccak256{
			Value0Keccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Value0_Value0ShaThree256{
			Value0ShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Value1(in any) *pbgear.Identity_Value1 {
	out := &pbgear.Identity_Value1{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value1
	v0 := To_OneOf_Identity_Value1_value1(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Value1_value1(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Value1_Value1None{
			Value1None: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Value1_Value1Raw0{
			Value1Raw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Value1_Value1Raw1{
			Value1Raw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Value1_Value1Raw2{
			Value1Raw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Value1_Value1Raw3{
			Value1Raw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Value1_Value1Raw4{
			Value1Raw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Value1_Value1Raw5{
			Value1Raw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Value1_Value1Raw6{
			Value1Raw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Value1_Value1Raw7{
			Value1Raw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Value1_Value1Raw8{
			Value1Raw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Value1_Value1Raw9{
			Value1Raw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Value1_Value1Raw10{
			Value1Raw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Value1_Value1Raw11{
			Value1Raw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Value1_Value1Raw12{
			Value1Raw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Value1_Value1Raw13{
			Value1Raw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Value1_Value1Raw14{
			Value1Raw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Value1_Value1Raw15{
			Value1Raw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Value1_Value1Raw16{
			Value1Raw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Value1_Value1Raw17{
			Value1Raw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Value1_Value1Raw18{
			Value1Raw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Value1_Value1Raw19{
			Value1Raw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Value1_Value1Raw20{
			Value1Raw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Value1_Value1Raw21{
			Value1Raw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Value1_Value1Raw22{
			Value1Raw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Value1_Value1Raw23{
			Value1Raw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Value1_Value1Raw24{
			Value1Raw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Value1_Value1Raw25{
			Value1Raw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Value1_Value1Raw26{
			Value1Raw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Value1_Value1Raw27{
			Value1Raw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Value1_Value1Raw28{
			Value1Raw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Value1_Value1Raw29{
			Value1Raw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Value1_Value1Raw30{
			Value1Raw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Value1_Value1Raw31{
			Value1Raw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Value1_Value1Raw32{
			Value1Raw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Value1_Value1BlakeTwo256{
			Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Value1_Value1Sha256{
			Value1Sha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Value1_Value1Keccak256{
			Value1Keccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Value1_Value1ShaThree256{
			Value1ShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Identity_Web(in any) *pbgear.Identity_Web {
	out := &pbgear.Identity_Web{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Web
	v0 := To_OneOf_Identity_Web_web(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Web").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Identity_Web_web(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Identity_Web_WebNone{
			WebNone: To_Identity_None(param),
		}
	case 1:
		return &pbgear.Identity_Web_WebRaw0{
			WebRaw0: To_Identity_Raw0(param),
		}
	case 2:
		return &pbgear.Identity_Web_WebRaw1{
			WebRaw1: To_Identity_Raw1(param),
		}
	case 3:
		return &pbgear.Identity_Web_WebRaw2{
			WebRaw2: To_Identity_Raw2(param),
		}
	case 4:
		return &pbgear.Identity_Web_WebRaw3{
			WebRaw3: To_Identity_Raw3(param),
		}
	case 5:
		return &pbgear.Identity_Web_WebRaw4{
			WebRaw4: To_Identity_Raw4(param),
		}
	case 6:
		return &pbgear.Identity_Web_WebRaw5{
			WebRaw5: To_Identity_Raw5(param),
		}
	case 7:
		return &pbgear.Identity_Web_WebRaw6{
			WebRaw6: To_Identity_Raw6(param),
		}
	case 8:
		return &pbgear.Identity_Web_WebRaw7{
			WebRaw7: To_Identity_Raw7(param),
		}
	case 9:
		return &pbgear.Identity_Web_WebRaw8{
			WebRaw8: To_Identity_Raw8(param),
		}
	case 10:
		return &pbgear.Identity_Web_WebRaw9{
			WebRaw9: To_Identity_Raw9(param),
		}
	case 11:
		return &pbgear.Identity_Web_WebRaw10{
			WebRaw10: To_Identity_Raw10(param),
		}
	case 12:
		return &pbgear.Identity_Web_WebRaw11{
			WebRaw11: To_Identity_Raw11(param),
		}
	case 13:
		return &pbgear.Identity_Web_WebRaw12{
			WebRaw12: To_Identity_Raw12(param),
		}
	case 14:
		return &pbgear.Identity_Web_WebRaw13{
			WebRaw13: To_Identity_Raw13(param),
		}
	case 15:
		return &pbgear.Identity_Web_WebRaw14{
			WebRaw14: To_Identity_Raw14(param),
		}
	case 16:
		return &pbgear.Identity_Web_WebRaw15{
			WebRaw15: To_Identity_Raw15(param),
		}
	case 17:
		return &pbgear.Identity_Web_WebRaw16{
			WebRaw16: To_Identity_Raw16(param),
		}
	case 18:
		return &pbgear.Identity_Web_WebRaw17{
			WebRaw17: To_Identity_Raw17(param),
		}
	case 19:
		return &pbgear.Identity_Web_WebRaw18{
			WebRaw18: To_Identity_Raw18(param),
		}
	case 20:
		return &pbgear.Identity_Web_WebRaw19{
			WebRaw19: To_Identity_Raw19(param),
		}
	case 21:
		return &pbgear.Identity_Web_WebRaw20{
			WebRaw20: To_Identity_Raw20(param),
		}
	case 22:
		return &pbgear.Identity_Web_WebRaw21{
			WebRaw21: To_Identity_Raw21(param),
		}
	case 23:
		return &pbgear.Identity_Web_WebRaw22{
			WebRaw22: To_Identity_Raw22(param),
		}
	case 24:
		return &pbgear.Identity_Web_WebRaw23{
			WebRaw23: To_Identity_Raw23(param),
		}
	case 25:
		return &pbgear.Identity_Web_WebRaw24{
			WebRaw24: To_Identity_Raw24(param),
		}
	case 26:
		return &pbgear.Identity_Web_WebRaw25{
			WebRaw25: To_Identity_Raw25(param),
		}
	case 27:
		return &pbgear.Identity_Web_WebRaw26{
			WebRaw26: To_Identity_Raw26(param),
		}
	case 28:
		return &pbgear.Identity_Web_WebRaw27{
			WebRaw27: To_Identity_Raw27(param),
		}
	case 29:
		return &pbgear.Identity_Web_WebRaw28{
			WebRaw28: To_Identity_Raw28(param),
		}
	case 30:
		return &pbgear.Identity_Web_WebRaw29{
			WebRaw29: To_Identity_Raw29(param),
		}
	case 31:
		return &pbgear.Identity_Web_WebRaw30{
			WebRaw30: To_Identity_Raw30(param),
		}
	case 32:
		return &pbgear.Identity_Web_WebRaw31{
			WebRaw31: To_Identity_Raw31(param),
		}
	case 33:
		return &pbgear.Identity_Web_WebRaw32{
			WebRaw32: To_Identity_Raw32(param),
		}
	case 34:
		return &pbgear.Identity_Web_WebBlakeTwo256{
			WebBlakeTwo256: To_Identity_BlakeTwo256(param),
		}
	case 35:
		return &pbgear.Identity_Web_WebSha256{
			WebSha256: To_Identity_Sha256(param),
		}
	case 36:
		return &pbgear.Identity_Web_WebKeccak256{
			WebKeccak256: To_Identity_Keccak256(param),
		}
	case 37:
		return &pbgear.Identity_Web_WebShaThree256{
			WebShaThree256: To_Identity_ShaThree256(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ImOnlinePallet(in any) *pbgear.ImOnlinePallet {
	out := &pbgear.ImOnlinePallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ImOnlinePallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ImOnlinePallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ImOnlinePallet_CallHeartbeatCall{
			CallHeartbeatCall: To_ImOnline_HeartbeatCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_ImOnline_HeartbeatCall(in any) *pbgear.ImOnline_HeartbeatCall {
	out := &pbgear.ImOnline_HeartbeatCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Heartbeat To_ImOnline_PalletImOnlineHeartbeat(w)
	out.Heartbeat = To_ImOnline_PalletImOnlineHeartbeat(v.ValueAt(0))
	// field Signature To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(w)
	out.Signature = To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(v.ValueAt(1))

	return out //from message
}

func To_ImOnline_PalletImOnlineHeartbeat(in any) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field BlockNumber
	out.BlockNumber = To_uint32(v.ValueAt(0))
	// primitive field SessionIndex
	out.SessionIndex = To_uint32(v.ValueAt(1))
	// primitive field AuthorityIndex
	out.AuthorityIndex = To_uint32(v.ValueAt(2))
	// primitive field ValidatorsLen
	out.ValidatorsLen = To_uint32(v.ValueAt(3))

	return out //from message
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Public(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public {
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreSr25519Public(w)
	out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

	return out //from message
}

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreSr25519Signature(w)
	out.Value0 = To_SpCoreSr25519Signature(v.ValueAt(0))

	return out //from message
}

func To_MediumSpender(in any) *pbgear.MediumSpender {
	out := &pbgear.MediumSpender{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_MultisigPallet(in any) *pbgear.MultisigPallet {
	out := &pbgear.MultisigPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_MultisigPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_MultisigPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.MultisigPallet_CallAsMultiThreshold1Call{
			CallAsMultiThreshold1Call: To_Multisig_AsMultiThreshold1Call(param),
		}
	case 1:
		return &pbgear.MultisigPallet_CallAsMultiCall{
			CallAsMultiCall: To_Multisig_AsMultiCall(param),
		}
	case 2:
		return &pbgear.MultisigPallet_CallApproveAsMultiCall{
			CallApproveAsMultiCall: To_Multisig_ApproveAsMultiCall(param),
		}
	case 3:
		return &pbgear.MultisigPallet_CallCancelAsMultiCall{
			CallCancelAsMultiCall: To_Multisig_CancelAsMultiCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Multisig_ApproveAsMultiCall(in any) *pbgear.Multisig_ApproveAsMultiCall {
	out := &pbgear.Multisig_ApproveAsMultiCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Threshold
	out.Threshold = To_uint32(v.ValueAt(0))
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(v.ValueAt(1))
	// optional field MaybeTimepoint
	out.MaybeTimepoint = To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(v.ValueAt(2))
	// primitive field CallHash
	out.CallHash = To_bytes(v.ValueAt(3))
	// field MaxWeight To_SpWeightsWeightV2Weight(w)
	out.MaxWeight = To_SpWeightsWeightV2Weight(v.ValueAt(4))

	return out //from message
}

func To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	panic("Not implemented")
	return nil //funcForOptionalField
}

func To_Multisig_AsMultiCall(in any) *pbgear.Multisig_AsMultiCall {
	out := &pbgear.Multisig_AsMultiCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Threshold
	out.Threshold = To_uint32(v.ValueAt(0))
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_AsMultiCall_other_signatories(v.ValueAt(1))
	// optional field MaybeTimepoint
	out.MaybeTimepoint = To_Optional_Multisig_AsMultiCall_maybe_timepoint(v.ValueAt(2))
	// oneOf field Call
	v3 := To_OneOf_Multisig_AsMultiCall_call(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))
	// field MaxWeight To_SpWeightsWeightV2Weight(w)
	out.MaxWeight = To_SpWeightsWeightV2Weight(v.ValueAt(4))

	return out //from message
}

func To_Repeated_Multisig_AsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Optional_Multisig_AsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_Multisig_AsMultiCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Multisig_AsMultiCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Multisig_AsMultiCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Multisig_AsMultiCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Multisig_AsMultiCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Multisig_AsMultiCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Multisig_AsMultiCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Multisig_AsMultiCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Multisig_AsMultiCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Multisig_AsMultiCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Multisig_AsMultiCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Multisig_AsMultiCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Multisig_AsMultiCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Multisig_AsMultiCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Multisig_AsMultiCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Multisig_AsMultiCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Multisig_AsMultiCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Multisig_AsMultiCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Multisig_AsMultiCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Multisig_AsMultiCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Multisig_AsMultiCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Multisig_AsMultiCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Multisig_AsMultiCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Multisig_AsMultiCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Multisig_AsMultiCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Multisig_AsMultiCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Multisig_AsMultiCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Multisig_AsMultiCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Multisig_AsMultiCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Multisig_AsMultiCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Multisig_AsMultiThreshold1Call(in any) *pbgear.Multisig_AsMultiThreshold1Call {
	out := &pbgear.Multisig_AsMultiThreshold1Call{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(v.ValueAt(0))
	// oneOf field Call
	v1 := To_OneOf_Multisig_AsMultiThreshold1Call_call(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_OneOf_Multisig_AsMultiThreshold1Call_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Multisig_AsMultiThreshold1Call_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Multisig_CancelAsMultiCall(in any) *pbgear.Multisig_CancelAsMultiCall {
	out := &pbgear.Multisig_CancelAsMultiCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Threshold
	out.Threshold = To_uint32(v.ValueAt(0))
	// repeated field OtherSignatories
	out.OtherSignatories = To_Repeated_Multisig_CancelAsMultiCall_other_signatories(v.ValueAt(1))
	// field Timepoint To_Multisig_PalletMultisigTimepoint(w)
	out.Timepoint = To_Multisig_PalletMultisigTimepoint(v.ValueAt(2))
	// primitive field CallHash
	out.CallHash = To_bytes(v.ValueAt(3))

	return out //from message
}

func To_Repeated_Multisig_CancelAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Multisig_PalletMultisigTimepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
	out := &pbgear.Multisig_PalletMultisigTimepoint{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Height
	out.Height = To_uint32(v.ValueAt(0))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_NominationPoolsPallet(in any) *pbgear.NominationPoolsPallet {
	out := &pbgear.NominationPoolsPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_NominationPoolsPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPoolsPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPoolsPallet_CallJoinCall{
			CallJoinCall: To_NominationPools_JoinCall(param),
		}
	case 1:
		return &pbgear.NominationPoolsPallet_CallBondExtraCall{
			CallBondExtraCall: To_NominationPools_BondExtraCall(param),
		}
	case 2:
		return &pbgear.NominationPoolsPallet_CallClaimPayoutCall{
			CallClaimPayoutCall: To_NominationPools_ClaimPayoutCall(param),
		}
	case 3:
		return &pbgear.NominationPoolsPallet_CallUnbondCall{
			CallUnbondCall: To_NominationPools_UnbondCall(param),
		}
	case 4:
		return &pbgear.NominationPoolsPallet_CallPoolWithdrawUnbondedCall{
			CallPoolWithdrawUnbondedCall: To_NominationPools_PoolWithdrawUnbondedCall(param),
		}
	case 5:
		return &pbgear.NominationPoolsPallet_CallWithdrawUnbondedCall{
			CallWithdrawUnbondedCall: To_NominationPools_WithdrawUnbondedCall(param),
		}
	case 6:
		return &pbgear.NominationPoolsPallet_CallCreateCall{
			CallCreateCall: To_NominationPools_CreateCall(param),
		}
	case 7:
		return &pbgear.NominationPoolsPallet_CallCreateWithPoolIdCall{
			CallCreateWithPoolIdCall: To_NominationPools_CreateWithPoolIdCall(param),
		}
	case 8:
		return &pbgear.NominationPoolsPallet_CallNominateCall{
			CallNominateCall: To_NominationPools_NominateCall(param),
		}
	case 9:
		return &pbgear.NominationPoolsPallet_CallSetStateCall{
			CallSetStateCall: To_NominationPools_SetStateCall(param),
		}
	case 10:
		return &pbgear.NominationPoolsPallet_CallSetMetadataCall{
			CallSetMetadataCall: To_NominationPools_SetMetadataCall(param),
		}
	case 11:
		return &pbgear.NominationPoolsPallet_CallSetConfigsCall{
			CallSetConfigsCall: To_NominationPools_SetConfigsCall(param),
		}
	case 12:
		return &pbgear.NominationPoolsPallet_CallUpdateRolesCall{
			CallUpdateRolesCall: To_NominationPools_UpdateRolesCall(param),
		}
	case 13:
		return &pbgear.NominationPoolsPallet_CallChillCall{
			CallChillCall: To_NominationPools_ChillCall(param),
		}
	case 14:
		return &pbgear.NominationPoolsPallet_CallBondExtraOtherCall{
			CallBondExtraOtherCall: To_NominationPools_BondExtraOtherCall(param),
		}
	case 15:
		return &pbgear.NominationPoolsPallet_CallSetClaimPermissionCall{
			CallSetClaimPermissionCall: To_NominationPools_SetClaimPermissionCall(param),
		}
	case 16:
		return &pbgear.NominationPoolsPallet_CallClaimPayoutOtherCall{
			CallClaimPayoutOtherCall: To_NominationPools_ClaimPayoutOtherCall(param),
		}
	case 17:
		return &pbgear.NominationPoolsPallet_CallSetCommissionCall{
			CallSetCommissionCall: To_NominationPools_SetCommissionCall(param),
		}
	case 18:
		return &pbgear.NominationPoolsPallet_CallSetCommissionMaxCall{
			CallSetCommissionMaxCall: To_NominationPools_SetCommissionMaxCall(param),
		}
	case 19:
		return &pbgear.NominationPoolsPallet_CallSetCommissionChangeRateCall{
			CallSetCommissionChangeRateCall: To_NominationPools_SetCommissionChangeRateCall(param),
		}
	case 20:
		return &pbgear.NominationPoolsPallet_CallClaimCommissionCall{
			CallClaimCommissionCall: To_NominationPools_ClaimCommissionCall(param),
		}
	case 21:
		return &pbgear.NominationPoolsPallet_CallAdjustPoolDepositCall{
			CallAdjustPoolDepositCall: To_NominationPools_AdjustPoolDepositCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Address20(in any) *pbgear.NominationPools_Address20 {
	out := &pbgear.NominationPools_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_Address32(in any) *pbgear.NominationPools_Address32 {
	out := &pbgear.NominationPools_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_AdjustPoolDepositCall(in any) *pbgear.NominationPools_AdjustPoolDepositCall {
	out := &pbgear.NominationPools_AdjustPoolDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_Blocked(in any) *pbgear.NominationPools_Blocked {
	out := &pbgear.NominationPools_Blocked{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_BondExtraCall(in any) *pbgear.NominationPools_BondExtraCall {
	out := &pbgear.NominationPools_BondExtraCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Extra
	v0 := To_OneOf_NominationPools_BondExtraCall_extra(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_BondExtraCall_extra(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_BondExtraCall_ExtraFreeBalance{
			ExtraFreeBalance: To_NominationPools_FreeBalance(param),
		}
	case 1:
		return &pbgear.NominationPools_BondExtraCall_ExtraRewards{
			ExtraRewards: To_NominationPools_Rewards(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_BondExtraOtherCall(in any) *pbgear.NominationPools_BondExtraOtherCall {
	out := &pbgear.NominationPools_BondExtraOtherCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Member
	v0 := To_OneOf_NominationPools_BondExtraOtherCall_member(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Member").Set(reflect.ValueOf(v0))
	// oneOf field Extra
	v1 := To_OneOf_NominationPools_BondExtraOtherCall_extra(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_NominationPools_BondExtraOtherCall_member(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_BondExtraOtherCall_MemberId{
			MemberId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_BondExtraOtherCall_MemberIndex{
			MemberIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_BondExtraOtherCall_MemberRaw{
			MemberRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_BondExtraOtherCall_MemberAddress32{
			MemberAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_BondExtraOtherCall_MemberAddress20{
			MemberAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_BondExtraOtherCall_extra(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_BondExtraOtherCall_ExtraFreeBalance{
			ExtraFreeBalance: To_NominationPools_FreeBalance(param),
		}
	case 1:
		return &pbgear.NominationPools_BondExtraOtherCall_ExtraRewards{
			ExtraRewards: To_NominationPools_Rewards(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Bouncer(in any) *pbgear.NominationPools_Bouncer {
	out := &pbgear.NominationPools_Bouncer{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Bouncer
	v0 := To_OneOf_NominationPools_Bouncer_bouncer(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Bouncer_bouncer(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Bouncer_BouncerId{
			BouncerId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_Bouncer_BouncerIndex{
			BouncerIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_Bouncer_BouncerRaw{
			BouncerRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_Bouncer_BouncerAddress32{
			BouncerAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_Bouncer_BouncerAddress20{
			BouncerAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_ChillCall(in any) *pbgear.NominationPools_ChillCall {
	out := &pbgear.NominationPools_ChillCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_ClaimCommissionCall(in any) *pbgear.NominationPools_ClaimCommissionCall {
	out := &pbgear.NominationPools_ClaimCommissionCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_ClaimPayoutCall(in any) *pbgear.NominationPools_ClaimPayoutCall {
	out := &pbgear.NominationPools_ClaimPayoutCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_ClaimPayoutOtherCall(in any) *pbgear.NominationPools_ClaimPayoutOtherCall {
	out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Other To_SpCoreCryptoAccountId32(w)
	out.Other = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_CreateCall(in any) *pbgear.NominationPools_CreateCall {
	out := &pbgear.NominationPools_CreateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Amount
	out.Amount = To_string(v.ValueAt(0))
	// oneOf field Root
	v1 := To_OneOf_NominationPools_CreateCall_root(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v1))
	// oneOf field Nominator
	v2 := To_OneOf_NominationPools_CreateCall_nominator(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v2))
	// oneOf field Bouncer
	v3 := To_OneOf_NominationPools_CreateCall_bouncer(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v3))

	return out //from message
}

func To_OneOf_NominationPools_CreateCall_root(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateCall_RootId{
			RootId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateCall_RootIndex{
			RootIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateCall_RootRaw{
			RootRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateCall_RootAddress32{
			RootAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateCall_RootAddress20{
			RootAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_CreateCall_nominator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateCall_NominatorId{
			NominatorId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateCall_NominatorIndex{
			NominatorIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateCall_NominatorRaw{
			NominatorRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateCall_NominatorAddress32{
			NominatorAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateCall_NominatorAddress20{
			NominatorAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_CreateCall_bouncer(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateCall_BouncerId{
			BouncerId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateCall_BouncerIndex{
			BouncerIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateCall_BouncerRaw{
			BouncerRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateCall_BouncerAddress32{
			BouncerAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateCall_BouncerAddress20{
			BouncerAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_CreateWithPoolIdCall(in any) *pbgear.NominationPools_CreateWithPoolIdCall {
	out := &pbgear.NominationPools_CreateWithPoolIdCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Amount
	out.Amount = To_string(v.ValueAt(0))
	// oneOf field Root
	v1 := To_OneOf_NominationPools_CreateWithPoolIdCall_root(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v1))
	// oneOf field Nominator
	v2 := To_OneOf_NominationPools_CreateWithPoolIdCall_nominator(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v2))
	// oneOf field Bouncer
	v3 := To_OneOf_NominationPools_CreateWithPoolIdCall_bouncer(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v3))
	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(4))

	return out //from message
}

func To_OneOf_NominationPools_CreateWithPoolIdCall_root(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateWithPoolIdCall_RootId{
			RootId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateWithPoolIdCall_RootIndex{
			RootIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateWithPoolIdCall_RootRaw{
			RootRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateWithPoolIdCall_RootAddress32{
			RootAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateWithPoolIdCall_RootAddress20{
			RootAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_CreateWithPoolIdCall_nominator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorId{
			NominatorId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorIndex{
			NominatorIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorRaw{
			NominatorRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorAddress32{
			NominatorAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorAddress20{
			NominatorAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_CreateWithPoolIdCall_bouncer(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerId{
			BouncerId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerIndex{
			BouncerIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerRaw{
			BouncerRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerAddress32{
			BouncerAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerAddress20{
			BouncerAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Destroying(in any) *pbgear.NominationPools_Destroying {
	out := &pbgear.NominationPools_Destroying{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_Extra(in any) *pbgear.NominationPools_Extra {
	out := &pbgear.NominationPools_Extra{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Extra
	v0 := To_OneOf_NominationPools_Extra_extra(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Extra_extra(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Extra_ExtraFreeBalance{
			ExtraFreeBalance: To_NominationPools_FreeBalance(param),
		}
	case 1:
		return &pbgear.NominationPools_Extra_ExtraRewards{
			ExtraRewards: To_NominationPools_Rewards(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_FreeBalance(in any) *pbgear.NominationPools_FreeBalance {
	out := &pbgear.NominationPools_FreeBalance{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_string(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_GlobalMaxCommission(in any) *pbgear.NominationPools_GlobalMaxCommission {
	out := &pbgear.NominationPools_GlobalMaxCommission{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field GlobalMaxCommission
	v0 := To_OneOf_NominationPools_GlobalMaxCommission_global_max_commission(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("GlobalMaxCommission").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_GlobalMaxCommission_global_max_commission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionNoop{
			GlobalMaxCommissionNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionSet{
			GlobalMaxCommissionSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionRemove{
			GlobalMaxCommissionRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Id(in any) *pbgear.NominationPools_Id {
	out := &pbgear.NominationPools_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_Index(in any) *pbgear.NominationPools_Index {
	out := &pbgear.NominationPools_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_NominationPools_TupleNull(w)
	out.Value0 = To_NominationPools_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_JoinCall(in any) *pbgear.NominationPools_JoinCall {
	out := &pbgear.NominationPools_JoinCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Amount
	out.Amount = To_string(v.ValueAt(0))
	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_MaxMembers(in any) *pbgear.NominationPools_MaxMembers {
	out := &pbgear.NominationPools_MaxMembers{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MaxMembers
	v0 := To_OneOf_NominationPools_MaxMembers_max_members(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MaxMembers").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MaxMembers_max_members(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MaxMembers_MaxMembersNoop{
			MaxMembersNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_MaxMembers_MaxMembersSet{
			MaxMembersSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_MaxMembers_MaxMembersRemove{
			MaxMembersRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_MaxMembersPerPool(in any) *pbgear.NominationPools_MaxMembersPerPool {
	out := &pbgear.NominationPools_MaxMembersPerPool{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MaxMembersPerPool
	v0 := To_OneOf_NominationPools_MaxMembersPerPool_max_members_per_pool(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MaxMembersPerPool").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MaxMembersPerPool_max_members_per_pool(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolNoop{
			MaxMembersPerPoolNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolSet{
			MaxMembersPerPoolSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolRemove{
			MaxMembersPerPoolRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_MaxPools(in any) *pbgear.NominationPools_MaxPools {
	out := &pbgear.NominationPools_MaxPools{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MaxPools
	v0 := To_OneOf_NominationPools_MaxPools_max_pools(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MaxPools").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MaxPools_max_pools(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MaxPools_MaxPoolsNoop{
			MaxPoolsNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_MaxPools_MaxPoolsSet{
			MaxPoolsSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_MaxPools_MaxPoolsRemove{
			MaxPoolsRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Member(in any) *pbgear.NominationPools_Member {
	out := &pbgear.NominationPools_Member{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Member
	v0 := To_OneOf_NominationPools_Member_member(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Member").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Member_member(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Member_MemberId{
			MemberId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_Member_MemberIndex{
			MemberIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_Member_MemberRaw{
			MemberRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_Member_MemberAddress32{
			MemberAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_Member_MemberAddress20{
			MemberAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_MemberAccount(in any) *pbgear.NominationPools_MemberAccount {
	out := &pbgear.NominationPools_MemberAccount{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MemberAccount
	v0 := To_OneOf_NominationPools_MemberAccount_member_account(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MemberAccount_member_account(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MemberAccount_MemberAccountId{
			MemberAccountId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_MemberAccount_MemberAccountIndex{
			MemberAccountIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_MemberAccount_MemberAccountRaw{
			MemberAccountRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_MemberAccount_MemberAccountAddress32{
			MemberAccountAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_MemberAccount_MemberAccountAddress20{
			MemberAccountAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_MinCreateBond(in any) *pbgear.NominationPools_MinCreateBond {
	out := &pbgear.NominationPools_MinCreateBond{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinCreateBond
	v0 := To_OneOf_NominationPools_MinCreateBond_min_create_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinCreateBond").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MinCreateBond_min_create_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MinCreateBond_MinCreateBondNoop{
			MinCreateBondNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_MinCreateBond_MinCreateBondSet{
			MinCreateBondSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_MinCreateBond_MinCreateBondRemove{
			MinCreateBondRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_MinJoinBond(in any) *pbgear.NominationPools_MinJoinBond {
	out := &pbgear.NominationPools_MinJoinBond{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinJoinBond
	v0 := To_OneOf_NominationPools_MinJoinBond_min_join_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinJoinBond").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_MinJoinBond_min_join_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_MinJoinBond_MinJoinBondNoop{
			MinJoinBondNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_MinJoinBond_MinJoinBondSet{
			MinJoinBondSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_MinJoinBond_MinJoinBondRemove{
			MinJoinBondRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_NewBouncer(in any) *pbgear.NominationPools_NewBouncer {
	out := &pbgear.NominationPools_NewBouncer{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field NewBouncer
	v0 := To_OneOf_NominationPools_NewBouncer_new_bouncer(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("NewBouncer").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_NewBouncer_new_bouncer(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_NewBouncer_NewBouncerNoop{
			NewBouncerNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_NewBouncer_NewBouncerSet{
			NewBouncerSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_NewBouncer_NewBouncerRemove{
			NewBouncerRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_NewNominator(in any) *pbgear.NominationPools_NewNominator {
	out := &pbgear.NominationPools_NewNominator{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field NewNominator
	v0 := To_OneOf_NominationPools_NewNominator_new_nominator(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("NewNominator").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_NewNominator_new_nominator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_NewNominator_NewNominatorNoop{
			NewNominatorNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_NewNominator_NewNominatorSet{
			NewNominatorSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_NewNominator_NewNominatorRemove{
			NewNominatorRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_NewRoot(in any) *pbgear.NominationPools_NewRoot {
	out := &pbgear.NominationPools_NewRoot{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field NewRoot
	v0 := To_OneOf_NominationPools_NewRoot_new_root(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("NewRoot").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_NewRoot_new_root(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_NewRoot_NewRootNoop{
			NewRootNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_NewRoot_NewRootSet{
			NewRootSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_NewRoot_NewRootRemove{
			NewRootRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_NominateCall(in any) *pbgear.NominationPools_NominateCall {
	out := &pbgear.NominationPools_NominateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// repeated field Validators
	out.Validators = To_Repeated_NominationPools_NominateCall_validators(v.ValueAt(1))

	return out //from message
}

func To_Repeated_NominationPools_NominateCall_validators(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_NominationPools_Nominator(in any) *pbgear.NominationPools_Nominator {
	out := &pbgear.NominationPools_Nominator{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Nominator
	v0 := To_OneOf_NominationPools_Nominator_nominator(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Nominator_nominator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Nominator_NominatorId{
			NominatorId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_Nominator_NominatorIndex{
			NominatorIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_Nominator_NominatorRaw{
			NominatorRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_Nominator_NominatorAddress32{
			NominatorAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_Nominator_NominatorAddress20{
			NominatorAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Noop(in any) *pbgear.NominationPools_Noop {
	out := &pbgear.NominationPools_Noop{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_Open(in any) *pbgear.NominationPools_Open {
	out := &pbgear.NominationPools_Open{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_PalletNominationPoolsCommissionChangeRate(in any) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
	v := in.(registry.Valuable)
	_ = v

	// field MaxIncrease To_SpArithmeticPerThingsPerbill(w)
	out.MaxIncrease = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
	// primitive field MinDelay
	out.MinDelay = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_Permission(in any) *pbgear.NominationPools_Permission {
	out := &pbgear.NominationPools_Permission{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Permission
	v0 := To_OneOf_NominationPools_Permission_permission(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Permission").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Permission_permission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Permission_PermissionPermissioned{
			PermissionPermissioned: To_NominationPools_Permissioned(param),
		}
	case 1:
		return &pbgear.NominationPools_Permission_PermissionPermissionlessCompound{
			PermissionPermissionlessCompound: To_NominationPools_PermissionlessCompound(param),
		}
	case 2:
		return &pbgear.NominationPools_Permission_PermissionPermissionlessWithdraw{
			PermissionPermissionlessWithdraw: To_NominationPools_PermissionlessWithdraw(param),
		}
	case 3:
		return &pbgear.NominationPools_Permission_PermissionPermissionlessAll{
			PermissionPermissionlessAll: To_NominationPools_PermissionlessAll(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Permissioned(in any) *pbgear.NominationPools_Permissioned {
	out := &pbgear.NominationPools_Permissioned{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_PermissionlessAll(in any) *pbgear.NominationPools_PermissionlessAll {
	out := &pbgear.NominationPools_PermissionlessAll{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_PermissionlessCompound(in any) *pbgear.NominationPools_PermissionlessCompound {
	out := &pbgear.NominationPools_PermissionlessCompound{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_PermissionlessWithdraw(in any) *pbgear.NominationPools_PermissionlessWithdraw {
	out := &pbgear.NominationPools_PermissionlessWithdraw{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_PoolWithdrawUnbondedCall(in any) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
	out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_Raw(in any) *pbgear.NominationPools_Raw {
	out := &pbgear.NominationPools_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_Remove(in any) *pbgear.NominationPools_Remove {
	out := &pbgear.NominationPools_Remove{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_Rewards(in any) *pbgear.NominationPools_Rewards {
	out := &pbgear.NominationPools_Rewards{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_NominationPools_Root(in any) *pbgear.NominationPools_Root {
	out := &pbgear.NominationPools_Root{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Root
	v0 := To_OneOf_NominationPools_Root_root(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_Root_root(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_Root_RootId{
			RootId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_Root_RootIndex{
			RootIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_Root_RootRaw{
			RootRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_Root_RootAddress32{
			RootAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_Root_RootAddress20{
			RootAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_Set(in any) *pbgear.NominationPools_Set {
	out := &pbgear.NominationPools_Set{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_string(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_SetClaimPermissionCall(in any) *pbgear.NominationPools_SetClaimPermissionCall {
	out := &pbgear.NominationPools_SetClaimPermissionCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Permission
	v0 := To_OneOf_NominationPools_SetClaimPermissionCall_permission(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Permission").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_SetClaimPermissionCall_permission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissioned{
			PermissionPermissioned: To_NominationPools_Permissioned(param),
		}
	case 1:
		return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessCompound{
			PermissionPermissionlessCompound: To_NominationPools_PermissionlessCompound(param),
		}
	case 2:
		return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessWithdraw{
			PermissionPermissionlessWithdraw: To_NominationPools_PermissionlessWithdraw(param),
		}
	case 3:
		return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessAll{
			PermissionPermissionlessAll: To_NominationPools_PermissionlessAll(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_SetCommissionCall(in any) *pbgear.NominationPools_SetCommissionCall {
	out := &pbgear.NominationPools_SetCommissionCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// optional field NewCommission
	out.NewCommission = To_Optional_NominationPools_SetCommissionCall_new_commission(v.ValueAt(1))

	return out //from message
}

func To_Optional_NominationPools_SetCommissionCall_new_commission(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_NominationPools_SetCommissionChangeRateCall(in any) *pbgear.NominationPools_SetCommissionChangeRateCall {
	out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// field ChangeRate To_NominationPools_PalletNominationPoolsCommissionChangeRate(w)
	out.ChangeRate = To_NominationPools_PalletNominationPoolsCommissionChangeRate(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_SetCommissionMaxCall(in any) *pbgear.NominationPools_SetCommissionMaxCall {
	out := &pbgear.NominationPools_SetCommissionMaxCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// field MaxCommission To_SpArithmeticPerThingsPerbill(w)
	out.MaxCommission = To_SpArithmeticPerThingsPerbill(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_SetConfigsCall(in any) *pbgear.NominationPools_SetConfigsCall {
	out := &pbgear.NominationPools_SetConfigsCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinJoinBond
	v0 := To_OneOf_NominationPools_SetConfigsCall_min_join_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinJoinBond").Set(reflect.ValueOf(v0))
	// oneOf field MinCreateBond
	v1 := To_OneOf_NominationPools_SetConfigsCall_min_create_bond(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("MinCreateBond").Set(reflect.ValueOf(v1))
	// oneOf field MaxPools
	v2 := To_OneOf_NominationPools_SetConfigsCall_max_pools(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("MaxPools").Set(reflect.ValueOf(v2))
	// oneOf field MaxMembers
	v3 := To_OneOf_NominationPools_SetConfigsCall_max_members(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("MaxMembers").Set(reflect.ValueOf(v3))
	// oneOf field MaxMembersPerPool
	v4 := To_OneOf_NominationPools_SetConfigsCall_max_members_per_pool(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("MaxMembersPerPool").Set(reflect.ValueOf(v4))
	// oneOf field GlobalMaxCommission
	v5 := To_OneOf_NominationPools_SetConfigsCall_global_max_commission(v.ValueAt(5))
	reflect.ValueOf(out).Elem().FieldByName("GlobalMaxCommission").Set(reflect.ValueOf(v5))

	return out //from message
}

func To_OneOf_NominationPools_SetConfigsCall_min_join_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_MinJoinBondNoop{
			MinJoinBondNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_MinJoinBondSet{
			MinJoinBondSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_MinJoinBondRemove{
			MinJoinBondRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_SetConfigsCall_min_create_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_MinCreateBondNoop{
			MinCreateBondNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_MinCreateBondSet{
			MinCreateBondSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_MinCreateBondRemove{
			MinCreateBondRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_SetConfigsCall_max_pools(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_MaxPoolsNoop{
			MaxPoolsNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_MaxPoolsSet{
			MaxPoolsSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_MaxPoolsRemove{
			MaxPoolsRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_SetConfigsCall_max_members(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersNoop{
			MaxMembersNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersSet{
			MaxMembersSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersRemove{
			MaxMembersRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_SetConfigsCall_max_members_per_pool(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolNoop{
			MaxMembersPerPoolNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolSet{
			MaxMembersPerPoolSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolRemove{
			MaxMembersPerPoolRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_SetConfigsCall_global_max_commission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionNoop{
			GlobalMaxCommissionNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionSet{
			GlobalMaxCommissionSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionRemove{
			GlobalMaxCommissionRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_SetMetadataCall(in any) *pbgear.NominationPools_SetMetadataCall {
	out := &pbgear.NominationPools_SetMetadataCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// primitive field Metadata
	out.Metadata = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_SetStateCall(in any) *pbgear.NominationPools_SetStateCall {
	out := &pbgear.NominationPools_SetStateCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// oneOf field State
	v1 := To_OneOf_NominationPools_SetStateCall_state(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("State").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_NominationPools_SetStateCall_state(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_SetStateCall_StateOpen{
			StateOpen: To_NominationPools_Open(param),
		}
	case 1:
		return &pbgear.NominationPools_SetStateCall_StateBlocked{
			StateBlocked: To_NominationPools_Blocked(param),
		}
	case 2:
		return &pbgear.NominationPools_SetStateCall_StateDestroying{
			StateDestroying: To_NominationPools_Destroying(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_State(in any) *pbgear.NominationPools_State {
	out := &pbgear.NominationPools_State{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field State
	v0 := To_OneOf_NominationPools_State_state(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("State").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_NominationPools_State_state(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_State_StateOpen{
			StateOpen: To_NominationPools_Open(param),
		}
	case 1:
		return &pbgear.NominationPools_State_StateBlocked{
			StateBlocked: To_NominationPools_Blocked(param),
		}
	case 2:
		return &pbgear.NominationPools_State_StateDestroying{
			StateDestroying: To_NominationPools_Destroying(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_TupleNull(in any) *pbgear.NominationPools_TupleNull {
	out := &pbgear.NominationPools_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpArithmeticPerThingsPerbill(w)
	out.Value0 = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
	// field Value1 To_SpCoreCryptoAccountId32(w)
	out.Value1 = To_SpCoreCryptoAccountId32(v.ValueAt(1))

	return out //from message
}

func To_NominationPools_UnbondCall(in any) *pbgear.NominationPools_UnbondCall {
	out := &pbgear.NominationPools_UnbondCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MemberAccount
	v0 := To_OneOf_NominationPools_UnbondCall_member_account(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))
	// primitive field UnbondingPoints
	out.UnbondingPoints = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_NominationPools_UnbondCall_member_account(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_UnbondCall_MemberAccountId{
			MemberAccountId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_UnbondCall_MemberAccountIndex{
			MemberAccountIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_UnbondCall_MemberAccountRaw{
			MemberAccountRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_UnbondCall_MemberAccountAddress32{
			MemberAccountAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_UnbondCall_MemberAccountAddress20{
			MemberAccountAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_UpdateRolesCall(in any) *pbgear.NominationPools_UpdateRolesCall {
	out := &pbgear.NominationPools_UpdateRolesCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field PoolId
	out.PoolId = To_uint32(v.ValueAt(0))
	// oneOf field NewRoot
	v1 := To_OneOf_NominationPools_UpdateRolesCall_new_root(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("NewRoot").Set(reflect.ValueOf(v1))
	// oneOf field NewNominator
	v2 := To_OneOf_NominationPools_UpdateRolesCall_new_nominator(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("NewNominator").Set(reflect.ValueOf(v2))
	// oneOf field NewBouncer
	v3 := To_OneOf_NominationPools_UpdateRolesCall_new_bouncer(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("NewBouncer").Set(reflect.ValueOf(v3))

	return out //from message
}

func To_OneOf_NominationPools_UpdateRolesCall_new_root(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_UpdateRolesCall_NewRootNoop{
			NewRootNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_UpdateRolesCall_NewRootSet{
			NewRootSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_UpdateRolesCall_NewRootRemove{
			NewRootRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_UpdateRolesCall_new_nominator(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_UpdateRolesCall_NewNominatorNoop{
			NewNominatorNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_UpdateRolesCall_NewNominatorSet{
			NewNominatorSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_UpdateRolesCall_NewNominatorRemove{
			NewNominatorRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_NominationPools_UpdateRolesCall_new_bouncer(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_UpdateRolesCall_NewBouncerNoop{
			NewBouncerNoop: To_NominationPools_Noop(param),
		}
	case 1:
		return &pbgear.NominationPools_UpdateRolesCall_NewBouncerSet{
			NewBouncerSet: To_NominationPools_Set(param),
		}
	case 2:
		return &pbgear.NominationPools_UpdateRolesCall_NewBouncerRemove{
			NewBouncerRemove: To_NominationPools_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_NominationPools_WithdrawUnbondedCall(in any) *pbgear.NominationPools_WithdrawUnbondedCall {
	out := &pbgear.NominationPools_WithdrawUnbondedCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MemberAccount
	v0 := To_OneOf_NominationPools_WithdrawUnbondedCall_member_account(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_OneOf_NominationPools_WithdrawUnbondedCall_member_account(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountId{
			MemberAccountId: To_NominationPools_Id(param),
		}
	case 1:
		return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountIndex{
			MemberAccountIndex: To_NominationPools_Index(param),
		}
	case 2:
		return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountRaw{
			MemberAccountRaw: To_NominationPools_Raw(param),
		}
	case 3:
		return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountAddress32{
			MemberAccountAddress32: To_NominationPools_Address32(param),
		}
	case 4:
		return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountAddress20{
			MemberAccountAddress20: To_NominationPools_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_None(in any) *pbgear.None {
	out := &pbgear.None{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_OffencesPallet(in any) *pbgear.OffencesPallet {
	out := &pbgear.OffencesPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_OriginsPallet(in any) *pbgear.OriginsPallet {
	out := &pbgear.OriginsPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_PreimagePallet(in any) *pbgear.PreimagePallet {
	out := &pbgear.PreimagePallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_PreimagePallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_PreimagePallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.PreimagePallet_CallNotePreimageCall{
			CallNotePreimageCall: To_Preimage_NotePreimageCall(param),
		}
	case 1:
		return &pbgear.PreimagePallet_CallUnnotePreimageCall{
			CallUnnotePreimageCall: To_Preimage_UnnotePreimageCall(param),
		}
	case 2:
		return &pbgear.PreimagePallet_CallRequestPreimageCall{
			CallRequestPreimageCall: To_Preimage_RequestPreimageCall(param),
		}
	case 3:
		return &pbgear.PreimagePallet_CallUnrequestPreimageCall{
			CallUnrequestPreimageCall: To_Preimage_UnrequestPreimageCall(param),
		}
	case 4:
		return &pbgear.PreimagePallet_CallEnsureUpdatedCall{
			CallEnsureUpdatedCall: To_Preimage_EnsureUpdatedCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Preimage_EnsureUpdatedCall(in any) *pbgear.Preimage_EnsureUpdatedCall {
	out := &pbgear.Preimage_EnsureUpdatedCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Hashes
	out.Hashes = To_Repeated_Preimage_EnsureUpdatedCall_hashes(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Preimage_EnsureUpdatedCall_hashes(in any) []*pbgear.PrimitiveTypesH256 {
	items := in.([]interface{})

	var out []*pbgear.PrimitiveTypesH256
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_PrimitiveTypesH256(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Preimage_NotePreimageCall(in any) *pbgear.Preimage_NotePreimageCall {
	out := &pbgear.Preimage_NotePreimageCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Bytes
	out.Bytes = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Preimage_RequestPreimageCall(in any) *pbgear.Preimage_RequestPreimageCall {
	out := &pbgear.Preimage_RequestPreimageCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_Preimage_UnnotePreimageCall(in any) *pbgear.Preimage_UnnotePreimageCall {
	out := &pbgear.Preimage_UnnotePreimageCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_Preimage_UnrequestPreimageCall(in any) *pbgear.Preimage_UnrequestPreimageCall {
	out := &pbgear.Preimage_UnrequestPreimageCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_PrimitiveTypesH256(in any) *pbgear.PrimitiveTypesH256 {
	out := &pbgear.PrimitiveTypesH256{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_ProxyPallet(in any) *pbgear.ProxyPallet {
	out := &pbgear.ProxyPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ProxyPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ProxyPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ProxyPallet_CallProxyCall{
			CallProxyCall: To_Proxy_ProxyCall(param),
		}
	case 1:
		return &pbgear.ProxyPallet_CallAddProxyCall{
			CallAddProxyCall: To_Proxy_AddProxyCall(param),
		}
	case 2:
		return &pbgear.ProxyPallet_CallRemoveProxyCall{
			CallRemoveProxyCall: To_Proxy_RemoveProxyCall(param),
		}
	case 3:
		return &pbgear.ProxyPallet_CallRemoveProxiesCall{
			CallRemoveProxiesCall: To_Proxy_RemoveProxiesCall(param),
		}
	case 4:
		return &pbgear.ProxyPallet_CallCreatePureCall{
			CallCreatePureCall: To_Proxy_CreatePureCall(param),
		}
	case 5:
		return &pbgear.ProxyPallet_CallKillPureCall{
			CallKillPureCall: To_Proxy_KillPureCall(param),
		}
	case 6:
		return &pbgear.ProxyPallet_CallAnnounceCall{
			CallAnnounceCall: To_Proxy_AnnounceCall(param),
		}
	case 7:
		return &pbgear.ProxyPallet_CallRemoveAnnouncementCall{
			CallRemoveAnnouncementCall: To_Proxy_RemoveAnnouncementCall(param),
		}
	case 8:
		return &pbgear.ProxyPallet_CallRejectAnnouncementCall{
			CallRejectAnnouncementCall: To_Proxy_RejectAnnouncementCall(param),
		}
	case 9:
		return &pbgear.ProxyPallet_CallProxyAnnouncedCall{
			CallProxyAnnouncedCall: To_Proxy_ProxyAnnouncedCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_AddProxyCall(in any) *pbgear.Proxy_AddProxyCall {
	out := &pbgear.Proxy_AddProxyCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Delegate
	v0 := To_OneOf_Proxy_AddProxyCall_delegate(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
	// oneOf field ProxyType
	v1 := To_OneOf_Proxy_AddProxyCall_proxy_type(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
	// primitive field Delay
	out.Delay = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Proxy_AddProxyCall_delegate(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_AddProxyCall_DelegateId{
			DelegateId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_AddProxyCall_DelegateIndex{
			DelegateIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_AddProxyCall_DelegateRaw{
			DelegateRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_AddProxyCall_DelegateAddress32{
			DelegateAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_AddProxyCall_DelegateAddress20{
			DelegateAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_AddProxyCall_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeAny{
			ProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeNonTransfer{
			ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeGovernance{
			ProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeStaking{
			ProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeIdentityJudgement{
			ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_AddProxyCall_ProxyTypeCancelProxy{
			ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Address20(in any) *pbgear.Proxy_Address20 {
	out := &pbgear.Proxy_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Proxy_Address32(in any) *pbgear.Proxy_Address32 {
	out := &pbgear.Proxy_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Proxy_AnnounceCall(in any) *pbgear.Proxy_AnnounceCall {
	out := &pbgear.Proxy_AnnounceCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Real
	v0 := To_OneOf_Proxy_AnnounceCall_real(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Proxy_AnnounceCall_real(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_AnnounceCall_RealId{
			RealId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_AnnounceCall_RealIndex{
			RealIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_AnnounceCall_RealRaw{
			RealRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_AnnounceCall_RealAddress32{
			RealAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_AnnounceCall_RealAddress20{
			RealAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Proxy_Any(in any) *pbgear.Proxy_Any {
	out := &pbgear.Proxy_Any{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_CancelProxy(in any) *pbgear.Proxy_CancelProxy {
	out := &pbgear.Proxy_CancelProxy{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_CreatePureCall(in any) *pbgear.Proxy_CreatePureCall {
	out := &pbgear.Proxy_CreatePureCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProxyType
	v0 := To_OneOf_Proxy_CreatePureCall_proxy_type(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v0))
	// primitive field Delay
	out.Delay = To_uint32(v.ValueAt(1))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Proxy_CreatePureCall_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeAny{
			ProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeNonTransfer{
			ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeGovernance{
			ProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeStaking{
			ProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeIdentityJudgement{
			ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_CreatePureCall_ProxyTypeCancelProxy{
			ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Delegate(in any) *pbgear.Proxy_Delegate {
	out := &pbgear.Proxy_Delegate{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Delegate
	v0 := To_OneOf_Proxy_Delegate_delegate(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Proxy_Delegate_delegate(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_Delegate_DelegateId{
			DelegateId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_Delegate_DelegateIndex{
			DelegateIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_Delegate_DelegateRaw{
			DelegateRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_Delegate_DelegateAddress32{
			DelegateAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_Delegate_DelegateAddress20{
			DelegateAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_ForceProxyType(in any) *pbgear.Proxy_ForceProxyType {
	out := &pbgear.Proxy_ForceProxyType{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ForceProxyType
	v0 := To_OneOf_Proxy_ForceProxyType_force_proxy_type(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Proxy_ForceProxyType_force_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeAny{
			ForceProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeNonTransfer{
			ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeGovernance{
			ForceProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeStaking{
			ForceProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeIdentityJudgement{
			ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_ForceProxyType_ForceProxyTypeCancelProxy{
			ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Governance(in any) *pbgear.Proxy_Governance {
	out := &pbgear.Proxy_Governance{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_Id(in any) *pbgear.Proxy_Id {
	out := &pbgear.Proxy_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Proxy_IdentityJudgement(in any) *pbgear.Proxy_IdentityJudgement {
	out := &pbgear.Proxy_IdentityJudgement{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_Index(in any) *pbgear.Proxy_Index {
	out := &pbgear.Proxy_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Proxy_TupleNull(w)
	out.Value0 = To_Proxy_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Proxy_KillPureCall(in any) *pbgear.Proxy_KillPureCall {
	out := &pbgear.Proxy_KillPureCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Spawner
	v0 := To_OneOf_Proxy_KillPureCall_spawner(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Spawner").Set(reflect.ValueOf(v0))
	// oneOf field ProxyType
	v1 := To_OneOf_Proxy_KillPureCall_proxy_type(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(2))
	// primitive field Height
	out.Height = To_uint32(v.ValueAt(3))
	// primitive field ExtIndex
	out.ExtIndex = To_uint32(v.ValueAt(4))

	return out //from message
}

func To_OneOf_Proxy_KillPureCall_spawner(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_KillPureCall_SpawnerId{
			SpawnerId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_KillPureCall_SpawnerIndex{
			SpawnerIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_KillPureCall_SpawnerRaw{
			SpawnerRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_KillPureCall_SpawnerAddress32{
			SpawnerAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_KillPureCall_SpawnerAddress20{
			SpawnerAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_KillPureCall_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_KillPureCall_ProxyTypeAny{
			ProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_KillPureCall_ProxyTypeNonTransfer{
			ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_KillPureCall_ProxyTypeGovernance{
			ProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_KillPureCall_ProxyTypeStaking{
			ProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_KillPureCall_ProxyTypeIdentityJudgement{
			ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_KillPureCall_ProxyTypeCancelProxy{
			ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_NonTransfer(in any) *pbgear.Proxy_NonTransfer {
	out := &pbgear.Proxy_NonTransfer{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_ProxyAnnouncedCall(in any) *pbgear.Proxy_ProxyAnnouncedCall {
	out := &pbgear.Proxy_ProxyAnnouncedCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Delegate
	v0 := To_OneOf_Proxy_ProxyAnnouncedCall_delegate(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
	// oneOf field Real
	v1 := To_OneOf_Proxy_ProxyAnnouncedCall_real(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v1))
	// oneOf field ForceProxyType
	v2 := To_OneOf_Proxy_ProxyAnnouncedCall_force_proxy_type(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v2))
	// oneOf field Call
	v3 := To_OneOf_Proxy_ProxyAnnouncedCall_call(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

	return out //from message
}

func To_OneOf_Proxy_ProxyAnnouncedCall_delegate(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyAnnouncedCall_DelegateId{
			DelegateId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyAnnouncedCall_DelegateIndex{
			DelegateIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyAnnouncedCall_DelegateRaw{
			DelegateRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyAnnouncedCall_DelegateAddress32{
			DelegateAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyAnnouncedCall_DelegateAddress20{
			DelegateAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_ProxyAnnouncedCall_real(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyAnnouncedCall_RealId{
			RealId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyAnnouncedCall_RealIndex{
			RealIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyAnnouncedCall_RealRaw{
			RealRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyAnnouncedCall_RealAddress32{
			RealAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyAnnouncedCall_RealAddress20{
			RealAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_ProxyAnnouncedCall_force_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeAny{
			ForceProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeNonTransfer{
			ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeGovernance{
			ForceProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeStaking{
			ForceProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeIdentityJudgement{
			ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeCancelProxy{
			ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_ProxyAnnouncedCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Proxy_ProxyAnnouncedCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_ProxyCall(in any) *pbgear.Proxy_ProxyCall {
	out := &pbgear.Proxy_ProxyCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Real
	v0 := To_OneOf_Proxy_ProxyCall_real(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
	// oneOf field ForceProxyType
	v1 := To_OneOf_Proxy_ProxyCall_force_proxy_type(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v1))
	// oneOf field Call
	v2 := To_OneOf_Proxy_ProxyCall_call(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v2))

	return out //from message
}

func To_OneOf_Proxy_ProxyCall_real(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyCall_RealId{
			RealId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyCall_RealIndex{
			RealIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyCall_RealRaw{
			RealRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyCall_RealAddress32{
			RealAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyCall_RealAddress20{
			RealAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_ProxyCall_force_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeAny{
			ForceProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeNonTransfer{
			ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeGovernance{
			ForceProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeStaking{
			ForceProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeIdentityJudgement{
			ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_ProxyCall_ForceProxyTypeCancelProxy{
			ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_ProxyCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Proxy_ProxyCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Proxy_ProxyCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Proxy_ProxyCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Proxy_ProxyCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Proxy_ProxyCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Proxy_ProxyCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Proxy_ProxyCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Proxy_ProxyCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Proxy_ProxyCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Proxy_ProxyCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Proxy_ProxyCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Proxy_ProxyCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Proxy_ProxyCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Proxy_ProxyCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Proxy_ProxyCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Proxy_ProxyCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Proxy_ProxyCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Proxy_ProxyCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Proxy_ProxyCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Proxy_ProxyCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Proxy_ProxyCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Proxy_ProxyCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Proxy_ProxyCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Proxy_ProxyCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Proxy_ProxyCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_ProxyType(in any) *pbgear.Proxy_ProxyType {
	out := &pbgear.Proxy_ProxyType{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProxyType
	v0 := To_OneOf_Proxy_ProxyType_proxy_type(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Proxy_ProxyType_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_ProxyType_ProxyTypeAny{
			ProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_ProxyType_ProxyTypeNonTransfer{
			ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_ProxyType_ProxyTypeGovernance{
			ProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_ProxyType_ProxyTypeStaking{
			ProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_ProxyType_ProxyTypeIdentityJudgement{
			ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_ProxyType_ProxyTypeCancelProxy{
			ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Raw(in any) *pbgear.Proxy_Raw {
	out := &pbgear.Proxy_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Proxy_Real(in any) *pbgear.Proxy_Real {
	out := &pbgear.Proxy_Real{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Real
	v0 := To_OneOf_Proxy_Real_real(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Proxy_Real_real(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_Real_RealId{
			RealId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_Real_RealIndex{
			RealIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_Real_RealRaw{
			RealRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_Real_RealAddress32{
			RealAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_Real_RealAddress20{
			RealAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_RejectAnnouncementCall(in any) *pbgear.Proxy_RejectAnnouncementCall {
	out := &pbgear.Proxy_RejectAnnouncementCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Delegate
	v0 := To_OneOf_Proxy_RejectAnnouncementCall_delegate(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Proxy_RejectAnnouncementCall_delegate(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_RejectAnnouncementCall_DelegateId{
			DelegateId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_RejectAnnouncementCall_DelegateIndex{
			DelegateIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_RejectAnnouncementCall_DelegateRaw{
			DelegateRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_RejectAnnouncementCall_DelegateAddress32{
			DelegateAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_RejectAnnouncementCall_DelegateAddress20{
			DelegateAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Proxy_RemoveAnnouncementCall(in any) *pbgear.Proxy_RemoveAnnouncementCall {
	out := &pbgear.Proxy_RemoveAnnouncementCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Real
	v0 := To_OneOf_Proxy_RemoveAnnouncementCall_real(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Proxy_RemoveAnnouncementCall_real(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_RemoveAnnouncementCall_RealId{
			RealId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_RemoveAnnouncementCall_RealIndex{
			RealIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_RemoveAnnouncementCall_RealRaw{
			RealRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_RemoveAnnouncementCall_RealAddress32{
			RealAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_RemoveAnnouncementCall_RealAddress20{
			RealAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Proxy_RemoveProxiesCall(in any) *pbgear.Proxy_RemoveProxiesCall {
	out := &pbgear.Proxy_RemoveProxiesCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_RemoveProxyCall(in any) *pbgear.Proxy_RemoveProxyCall {
	out := &pbgear.Proxy_RemoveProxyCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Delegate
	v0 := To_OneOf_Proxy_RemoveProxyCall_delegate(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
	// oneOf field ProxyType
	v1 := To_OneOf_Proxy_RemoveProxyCall_proxy_type(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
	// primitive field Delay
	out.Delay = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Proxy_RemoveProxyCall_delegate(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_RemoveProxyCall_DelegateId{
			DelegateId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_RemoveProxyCall_DelegateIndex{
			DelegateIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_RemoveProxyCall_DelegateRaw{
			DelegateRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_RemoveProxyCall_DelegateAddress32{
			DelegateAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_RemoveProxyCall_DelegateAddress20{
			DelegateAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Proxy_RemoveProxyCall_proxy_type(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeAny{
			ProxyTypeAny: To_Proxy_Any(param),
		}
	case 1:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeNonTransfer{
			ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
		}
	case 2:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeGovernance{
			ProxyTypeGovernance: To_Proxy_Governance(param),
		}
	case 3:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeStaking{
			ProxyTypeStaking: To_Proxy_Staking(param),
		}
	case 4:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeIdentityJudgement{
			ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
		}
	case 5:
		return &pbgear.Proxy_RemoveProxyCall_ProxyTypeCancelProxy{
			ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Spawner(in any) *pbgear.Proxy_Spawner {
	out := &pbgear.Proxy_Spawner{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Spawner
	v0 := To_OneOf_Proxy_Spawner_spawner(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Spawner").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Proxy_Spawner_spawner(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Proxy_Spawner_SpawnerId{
			SpawnerId: To_Proxy_Id(param),
		}
	case 1:
		return &pbgear.Proxy_Spawner_SpawnerIndex{
			SpawnerIndex: To_Proxy_Index(param),
		}
	case 2:
		return &pbgear.Proxy_Spawner_SpawnerRaw{
			SpawnerRaw: To_Proxy_Raw(param),
		}
	case 3:
		return &pbgear.Proxy_Spawner_SpawnerAddress32{
			SpawnerAddress32: To_Proxy_Address32(param),
		}
	case 4:
		return &pbgear.Proxy_Spawner_SpawnerAddress20{
			SpawnerAddress20: To_Proxy_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Proxy_Staking(in any) *pbgear.Proxy_Staking {
	out := &pbgear.Proxy_Staking{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Proxy_TupleNull(in any) *pbgear.Proxy_TupleNull {
	out := &pbgear.Proxy_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_ReferendaPallet(in any) *pbgear.ReferendaPallet {
	out := &pbgear.ReferendaPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_ReferendaPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_ReferendaPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.ReferendaPallet_CallSubmitCall{
			CallSubmitCall: To_Referenda_SubmitCall(param),
		}
	case 1:
		return &pbgear.ReferendaPallet_CallPlaceDecisionDepositCall{
			CallPlaceDecisionDepositCall: To_Referenda_PlaceDecisionDepositCall(param),
		}
	case 2:
		return &pbgear.ReferendaPallet_CallRefundDecisionDepositCall{
			CallRefundDecisionDepositCall: To_Referenda_RefundDecisionDepositCall(param),
		}
	case 3:
		return &pbgear.ReferendaPallet_CallCancelCall{
			CallCancelCall: To_Referenda_CancelCall(param),
		}
	case 4:
		return &pbgear.ReferendaPallet_CallKillCall{
			CallKillCall: To_Referenda_KillCall(param),
		}
	case 5:
		return &pbgear.ReferendaPallet_CallNudgeReferendumCall{
			CallNudgeReferendumCall: To_Referenda_NudgeReferendumCall(param),
		}
	case 6:
		return &pbgear.ReferendaPallet_CallOneFewerDecidingCall{
			CallOneFewerDecidingCall: To_Referenda_OneFewerDecidingCall(param),
		}
	case 7:
		return &pbgear.ReferendaPallet_CallRefundSubmissionDepositCall{
			CallRefundSubmissionDepositCall: To_Referenda_RefundSubmissionDepositCall(param),
		}
	case 8:
		return &pbgear.ReferendaPallet_CallSetMetadataCall{
			CallSetMetadataCall: To_Referenda_SetMetadataCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_After(in any) *pbgear.Referenda_After {
	out := &pbgear.Referenda_After{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_At(in any) *pbgear.Referenda_At {
	out := &pbgear.Referenda_At{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_CancelCall(in any) *pbgear.Referenda_CancelCall {
	out := &pbgear.Referenda_CancelCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_EnactmentMoment(in any) *pbgear.Referenda_EnactmentMoment {
	out := &pbgear.Referenda_EnactmentMoment{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field EnactmentMoment
	v0 := To_OneOf_Referenda_EnactmentMoment_enactment_moment(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_EnactmentMoment_enactment_moment(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_EnactmentMoment_EnactmentMomentAt{
			EnactmentMomentAt: To_Referenda_At(param),
		}
	case 1:
		return &pbgear.Referenda_EnactmentMoment_EnactmentMomentAfter{
			EnactmentMomentAfter: To_Referenda_After(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_Inline(in any) *pbgear.Referenda_Inline {
	out := &pbgear.Referenda_Inline{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_BoundedCollectionsBoundedVecBoundedVec(w)
	out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))

	return out //from message
}

func To_Referenda_KillCall(in any) *pbgear.Referenda_KillCall {
	out := &pbgear.Referenda_KillCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_Legacy(in any) *pbgear.Referenda_Legacy {
	out := &pbgear.Referenda_Legacy{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_Referenda_Lookup(in any) *pbgear.Referenda_Lookup {
	out := &pbgear.Referenda_Lookup{}
	v := in.(registry.Valuable)
	_ = v

	// field Hash To_PrimitiveTypesH256(w)
	out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field Len
	out.Len = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Referenda_NudgeReferendumCall(in any) *pbgear.Referenda_NudgeReferendumCall {
	out := &pbgear.Referenda_NudgeReferendumCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_OneFewerDecidingCall(in any) *pbgear.Referenda_OneFewerDecidingCall {
	out := &pbgear.Referenda_OneFewerDecidingCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Track
	out.Track = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_Origins(in any) *pbgear.Referenda_Origins {
	out := &pbgear.Referenda_Origins{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Referenda_Origins_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_Origins_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_Origins_Value0StakingAdmin{
			Value0StakingAdmin: To_StakingAdmin(param),
		}
	case 1:
		return &pbgear.Referenda_Origins_Value0Treasurer{
			Value0Treasurer: To_Treasurer(param),
		}
	case 2:
		return &pbgear.Referenda_Origins_Value0FellowshipAdmin{
			Value0FellowshipAdmin: To_FellowshipAdmin(param),
		}
	case 3:
		return &pbgear.Referenda_Origins_Value0GeneralAdmin{
			Value0GeneralAdmin: To_GeneralAdmin(param),
		}
	case 4:
		return &pbgear.Referenda_Origins_Value0ReferendumCanceller{
			Value0ReferendumCanceller: To_ReferendumCanceller(param),
		}
	case 5:
		return &pbgear.Referenda_Origins_Value0ReferendumKiller{
			Value0ReferendumKiller: To_ReferendumKiller(param),
		}
	case 6:
		return &pbgear.Referenda_Origins_Value0SmallTipper{
			Value0SmallTipper: To_SmallTipper(param),
		}
	case 7:
		return &pbgear.Referenda_Origins_Value0BigTipper{
			Value0BigTipper: To_BigTipper(param),
		}
	case 8:
		return &pbgear.Referenda_Origins_Value0SmallSpender{
			Value0SmallSpender: To_SmallSpender(param),
		}
	case 9:
		return &pbgear.Referenda_Origins_Value0MediumSpender{
			Value0MediumSpender: To_MediumSpender(param),
		}
	case 10:
		return &pbgear.Referenda_Origins_Value0BigSpender{
			Value0BigSpender: To_BigSpender(param),
		}
	case 11:
		return &pbgear.Referenda_Origins_Value0WhitelistedCaller{
			Value0WhitelistedCaller: To_WhitelistedCaller(param),
		}
	case 12:
		return &pbgear.Referenda_Origins_Value0FellowshipInitiates{
			Value0FellowshipInitiates: To_FellowshipInitiates(param),
		}
	case 13:
		return &pbgear.Referenda_Origins_Value0Fellows{
			Value0Fellows: To_Fellows(param),
		}
	case 14:
		return &pbgear.Referenda_Origins_Value0FellowshipExperts{
			Value0FellowshipExperts: To_FellowshipExperts(param),
		}
	case 15:
		return &pbgear.Referenda_Origins_Value0FellowshipMasters{
			Value0FellowshipMasters: To_FellowshipMasters(param),
		}
	case 16:
		return &pbgear.Referenda_Origins_Value0Fellowship1Dan{
			Value0Fellowship1Dan: To_Fellowship1Dan(param),
		}
	case 17:
		return &pbgear.Referenda_Origins_Value0Fellowship2Dan{
			Value0Fellowship2Dan: To_Fellowship2Dan(param),
		}
	case 18:
		return &pbgear.Referenda_Origins_Value0Fellowship3Dan{
			Value0Fellowship3Dan: To_Fellowship3Dan(param),
		}
	case 19:
		return &pbgear.Referenda_Origins_Value0Fellowship4Dan{
			Value0Fellowship4Dan: To_Fellowship4Dan(param),
		}
	case 20:
		return &pbgear.Referenda_Origins_Value0Fellowship5Dan{
			Value0Fellowship5Dan: To_Fellowship5Dan(param),
		}
	case 21:
		return &pbgear.Referenda_Origins_Value0Fellowship6Dan{
			Value0Fellowship6Dan: To_Fellowship6Dan(param),
		}
	case 22:
		return &pbgear.Referenda_Origins_Value0Fellowship7Dan{
			Value0Fellowship7Dan: To_Fellowship7Dan(param),
		}
	case 23:
		return &pbgear.Referenda_Origins_Value0Fellowship8Dan{
			Value0Fellowship8Dan: To_Fellowship8Dan(param),
		}
	case 24:
		return &pbgear.Referenda_Origins_Value0Fellowship9Dan{
			Value0Fellowship9Dan: To_Fellowship9Dan(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_PlaceDecisionDepositCall(in any) *pbgear.Referenda_PlaceDecisionDepositCall {
	out := &pbgear.Referenda_PlaceDecisionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_Proposal(in any) *pbgear.Referenda_Proposal {
	out := &pbgear.Referenda_Proposal{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Proposal
	v0 := To_OneOf_Referenda_Proposal_proposal(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_Proposal_proposal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_Proposal_ProposalLegacy{
			ProposalLegacy: To_Referenda_Legacy(param),
		}
	case 1:
		return &pbgear.Referenda_Proposal_ProposalInline{
			ProposalInline: To_Referenda_Inline(param),
		}
	case 2:
		return &pbgear.Referenda_Proposal_ProposalLookup{
			ProposalLookup: To_Referenda_Lookup(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_ProposalOrigin(in any) *pbgear.Referenda_ProposalOrigin {
	out := &pbgear.Referenda_ProposalOrigin{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProposalOrigin
	v0 := To_OneOf_Referenda_ProposalOrigin_proposal_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_ProposalOrigin_proposal_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_ProposalOrigin_ProposalOriginSystem{
			ProposalOriginSystem: To_Referenda_System(param),
		}
	case 20:
		return &pbgear.Referenda_ProposalOrigin_ProposalOriginOrigins{
			ProposalOriginOrigins: To_Referenda_Origins(param),
		}
	case 2:
		return &pbgear.Referenda_ProposalOrigin_ProposalOriginVoid{
			ProposalOriginVoid: To_Referenda_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_RefundDecisionDepositCall(in any) *pbgear.Referenda_RefundDecisionDepositCall {
	out := &pbgear.Referenda_RefundDecisionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_RefundSubmissionDepositCall(in any) *pbgear.Referenda_RefundSubmissionDepositCall {
	out := &pbgear.Referenda_RefundSubmissionDepositCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Referenda_SetMetadataCall(in any) *pbgear.Referenda_SetMetadataCall {
	out := &pbgear.Referenda_SetMetadataCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// optional field MaybeHash
	out.MaybeHash = To_Optional_Referenda_SetMetadataCall_maybe_hash(v.ValueAt(1))

	return out //from message
}

func To_Optional_Referenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_Referenda_SubmitCall(in any) *pbgear.Referenda_SubmitCall {
	out := &pbgear.Referenda_SubmitCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ProposalOrigin
	v0 := To_OneOf_Referenda_SubmitCall_proposal_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))
	// oneOf field Proposal
	v1 := To_OneOf_Referenda_SubmitCall_proposal(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v1))
	// oneOf field EnactmentMoment
	v2 := To_OneOf_Referenda_SubmitCall_enactment_moment(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v2))

	return out //from message
}

func To_OneOf_Referenda_SubmitCall_proposal_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_SubmitCall_ProposalOriginSystem{
			ProposalOriginSystem: To_Referenda_System(param),
		}
	case 20:
		return &pbgear.Referenda_SubmitCall_ProposalOriginOrigins{
			ProposalOriginOrigins: To_Referenda_Origins(param),
		}
	case 2:
		return &pbgear.Referenda_SubmitCall_ProposalOriginVoid{
			ProposalOriginVoid: To_Referenda_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Referenda_SubmitCall_proposal(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_SubmitCall_ProposalLegacy{
			ProposalLegacy: To_Referenda_Legacy(param),
		}
	case 1:
		return &pbgear.Referenda_SubmitCall_ProposalInline{
			ProposalInline: To_Referenda_Inline(param),
		}
	case 2:
		return &pbgear.Referenda_SubmitCall_ProposalLookup{
			ProposalLookup: To_Referenda_Lookup(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Referenda_SubmitCall_enactment_moment(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_SubmitCall_EnactmentMomentAt{
			EnactmentMomentAt: To_Referenda_At(param),
		}
	case 1:
		return &pbgear.Referenda_SubmitCall_EnactmentMomentAfter{
			EnactmentMomentAfter: To_Referenda_After(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_System(in any) *pbgear.Referenda_System {
	out := &pbgear.Referenda_System{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Referenda_System_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_System_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Referenda_System_Value0Root{
			Value0Root: To_Root(param),
		}
	case 1:
		return &pbgear.Referenda_System_Value0Signed{
			Value0Signed: To_Signed(param),
		}
	case 2:
		return &pbgear.Referenda_System_Value0None{
			Value0None: To_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Referenda_Void(in any) *pbgear.Referenda_Void {
	out := &pbgear.Referenda_Void{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Referenda_Void_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Referenda_Void_value0(in any) any {
	return nil
}
func To_ReferendumCanceller(in any) *pbgear.ReferendumCanceller {
	out := &pbgear.ReferendumCanceller{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_ReferendumKiller(in any) *pbgear.ReferendumKiller {
	out := &pbgear.ReferendumKiller{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Root(in any) *pbgear.Root {
	out := &pbgear.Root{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_SchedulerPallet(in any) *pbgear.SchedulerPallet {
	out := &pbgear.SchedulerPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_SchedulerPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_SchedulerPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.SchedulerPallet_CallScheduleCall{
			CallScheduleCall: To_Scheduler_ScheduleCall(param),
		}
	case 1:
		return &pbgear.SchedulerPallet_CallCancelCall{
			CallCancelCall: To_Scheduler_CancelCall(param),
		}
	case 2:
		return &pbgear.SchedulerPallet_CallScheduleNamedCall{
			CallScheduleNamedCall: To_Scheduler_ScheduleNamedCall(param),
		}
	case 3:
		return &pbgear.SchedulerPallet_CallCancelNamedCall{
			CallCancelNamedCall: To_Scheduler_CancelNamedCall(param),
		}
	case 4:
		return &pbgear.SchedulerPallet_CallScheduleAfterCall{
			CallScheduleAfterCall: To_Scheduler_ScheduleAfterCall(param),
		}
	case 5:
		return &pbgear.SchedulerPallet_CallScheduleNamedAfterCall{
			CallScheduleNamedAfterCall: To_Scheduler_ScheduleNamedAfterCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Scheduler_CancelCall(in any) *pbgear.Scheduler_CancelCall {
	out := &pbgear.Scheduler_CancelCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field When
	out.When = To_uint32(v.ValueAt(0))
	// primitive field Index
	out.Index = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Scheduler_CancelNamedCall(in any) *pbgear.Scheduler_CancelNamedCall {
	out := &pbgear.Scheduler_CancelNamedCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Id
	out.Id = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Scheduler_ScheduleAfterCall(in any) *pbgear.Scheduler_ScheduleAfterCall {
	out := &pbgear.Scheduler_ScheduleAfterCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field After
	out.After = To_uint32(v.ValueAt(0))
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(v.ValueAt(1))
	// primitive field Priority
	out.Priority = To_uint32(v.ValueAt(2))
	// oneOf field Call
	v3 := To_OneOf_Scheduler_ScheduleAfterCall_call(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

	return out //from message
}

func To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleAfterCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Scheduler_ScheduleAfterCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Scheduler_ScheduleAfterCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Scheduler_ScheduleAfterCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Scheduler_ScheduleAfterCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Scheduler_ScheduleAfterCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Scheduler_ScheduleAfterCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Scheduler_ScheduleAfterCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Scheduler_ScheduleAfterCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Scheduler_ScheduleAfterCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Scheduler_ScheduleAfterCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Scheduler_ScheduleAfterCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Scheduler_ScheduleAfterCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Scheduler_ScheduleAfterCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Scheduler_ScheduleAfterCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Scheduler_ScheduleAfterCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Scheduler_ScheduleAfterCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Scheduler_ScheduleAfterCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Scheduler_ScheduleAfterCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Scheduler_ScheduleAfterCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Scheduler_ScheduleAfterCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Scheduler_ScheduleAfterCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Scheduler_ScheduleAfterCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Scheduler_ScheduleAfterCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Scheduler_ScheduleAfterCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Scheduler_ScheduleAfterCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Scheduler_ScheduleAfterCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Scheduler_ScheduleAfterCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Scheduler_ScheduleAfterCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Scheduler_ScheduleAfterCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Scheduler_ScheduleCall(in any) *pbgear.Scheduler_ScheduleCall {
	out := &pbgear.Scheduler_ScheduleCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field When
	out.When = To_uint32(v.ValueAt(0))
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleCall_maybe_periodic(v.ValueAt(1))
	// primitive field Priority
	out.Priority = To_uint32(v.ValueAt(2))
	// oneOf field Call
	v3 := To_OneOf_Scheduler_ScheduleCall_call(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

	return out //from message
}

func To_Optional_Scheduler_ScheduleCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Scheduler_ScheduleCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Scheduler_ScheduleCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Scheduler_ScheduleCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Scheduler_ScheduleCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Scheduler_ScheduleCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Scheduler_ScheduleCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Scheduler_ScheduleCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Scheduler_ScheduleCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Scheduler_ScheduleCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Scheduler_ScheduleCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Scheduler_ScheduleCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Scheduler_ScheduleCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Scheduler_ScheduleCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Scheduler_ScheduleCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Scheduler_ScheduleCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Scheduler_ScheduleCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Scheduler_ScheduleCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Scheduler_ScheduleCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Scheduler_ScheduleCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Scheduler_ScheduleCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Scheduler_ScheduleCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Scheduler_ScheduleCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Scheduler_ScheduleCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Scheduler_ScheduleCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Scheduler_ScheduleCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Scheduler_ScheduleCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Scheduler_ScheduleCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Scheduler_ScheduleCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Scheduler_ScheduleCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Scheduler_ScheduleNamedAfterCall(in any) *pbgear.Scheduler_ScheduleNamedAfterCall {
	out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Id
	out.Id = To_bytes(v.ValueAt(0))
	// primitive field After
	out.After = To_uint32(v.ValueAt(1))
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(v.ValueAt(2))
	// primitive field Priority
	out.Priority = To_uint32(v.ValueAt(3))
	// oneOf field Call
	v4 := To_OneOf_Scheduler_ScheduleNamedAfterCall_call(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v4))

	return out //from message
}

func To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleNamedAfterCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Scheduler_ScheduleNamedCall(in any) *pbgear.Scheduler_ScheduleNamedCall {
	out := &pbgear.Scheduler_ScheduleNamedCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Id
	out.Id = To_bytes(v.ValueAt(0))
	// primitive field When
	out.When = To_uint32(v.ValueAt(1))
	// optional field MaybePeriodic
	out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(v.ValueAt(2))
	// primitive field Priority
	out.Priority = To_uint32(v.ValueAt(3))
	// oneOf field Call
	v4 := To_OneOf_Scheduler_ScheduleNamedCall_call(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v4))

	return out //from message
}

func To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	panic("Not implemented")
	return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleNamedCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Scheduler_ScheduleNamedCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Scheduler_ScheduleNamedCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Scheduler_ScheduleNamedCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Scheduler_ScheduleNamedCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Scheduler_ScheduleNamedCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Scheduler_ScheduleNamedCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Scheduler_ScheduleNamedCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Scheduler_ScheduleNamedCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Scheduler_ScheduleNamedCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Scheduler_ScheduleNamedCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Scheduler_ScheduleNamedCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Scheduler_ScheduleNamedCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Scheduler_ScheduleNamedCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Scheduler_ScheduleNamedCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Scheduler_ScheduleNamedCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Scheduler_ScheduleNamedCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Scheduler_ScheduleNamedCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Scheduler_ScheduleNamedCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Scheduler_ScheduleNamedCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Scheduler_ScheduleNamedCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Scheduler_ScheduleNamedCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Scheduler_ScheduleNamedCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Scheduler_ScheduleNamedCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Scheduler_ScheduleNamedCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Scheduler_ScheduleNamedCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Scheduler_ScheduleNamedCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Scheduler_ScheduleNamedCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Scheduler_ScheduleNamedCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Scheduler_ScheduleNamedCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Scheduler_TupleUint32Uint32(in any) *pbgear.Scheduler_TupleUint32Uint32 {
	out := &pbgear.Scheduler_TupleUint32Uint32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_SessionPallet(in any) *pbgear.SessionPallet {
	out := &pbgear.SessionPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_SessionPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_SessionPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.SessionPallet_CallSetKeysCall{
			CallSetKeysCall: To_Session_SetKeysCall(param),
		}
	case 1:
		return &pbgear.SessionPallet_CallPurgeKeysCall{
			CallPurgeKeysCall: To_Session_PurgeKeysCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Session_PurgeKeysCall(in any) *pbgear.Session_PurgeKeysCall {
	out := &pbgear.Session_PurgeKeysCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Session_SetKeysCall(in any) *pbgear.Session_SetKeysCall {
	out := &pbgear.Session_SetKeysCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Keys To_VaraRuntimeSessionKeys(w)
	out.Keys = To_VaraRuntimeSessionKeys(v.ValueAt(0))
	// primitive field Proof
	out.Proof = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_Signed(in any) *pbgear.Signed {
	out := &pbgear.Signed{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_SmallSpender(in any) *pbgear.SmallSpender {
	out := &pbgear.SmallSpender{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_SmallTipper(in any) *pbgear.SmallTipper {
	out := &pbgear.SmallTipper{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_SpArithmeticPerThingsPerU16(in any) *pbgear.SpArithmeticPerThingsPerU16 {
	out := &pbgear.SpArithmeticPerThingsPerU16{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_SpArithmeticPerThingsPerbill(in any) *pbgear.SpArithmeticPerThingsPerbill {
	out := &pbgear.SpArithmeticPerThingsPerbill{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_SpArithmeticPerThingsPercent(in any) *pbgear.SpArithmeticPerThingsPercent {
	out := &pbgear.SpArithmeticPerThingsPercent{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_SpAuthorityDiscoveryAppPublic(in any) *pbgear.SpAuthorityDiscoveryAppPublic {
	out := &pbgear.SpAuthorityDiscoveryAppPublic{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreSr25519Public(w)
	out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

	return out //from message
}

func To_SpConsensusBabeAppPublic(in any) *pbgear.SpConsensusBabeAppPublic {
	out := &pbgear.SpConsensusBabeAppPublic{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreSr25519Public(w)
	out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

	return out //from message
}

func To_SpConsensusGrandpaAppPublic(in any) *pbgear.SpConsensusGrandpaAppPublic {
	out := &pbgear.SpConsensusGrandpaAppPublic{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreEd25519Public(w)
	out.Value0 = To_SpCoreEd25519Public(v.ValueAt(0))

	return out //from message
}

func To_SpConsensusSlotsEquivocationProof(in any) *pbgear.SpConsensusSlotsEquivocationProof {
	out := &pbgear.SpConsensusSlotsEquivocationProof{}
	v := in.(registry.Valuable)
	_ = v

	// field Offender To_Babe_SpConsensusBabeAppPublic(w)
	out.Offender = To_Babe_SpConsensusBabeAppPublic(v.ValueAt(0))
	// field Slot To_SpConsensusSlotsSlot(w)
	out.Slot = To_SpConsensusSlotsSlot(v.ValueAt(1))
	// field FirstHeader To_SpRuntimeGenericHeaderHeader(w)
	out.FirstHeader = To_SpRuntimeGenericHeaderHeader(v.ValueAt(2))
	// field SecondHeader To_SpRuntimeGenericHeaderHeader(w)
	out.SecondHeader = To_SpRuntimeGenericHeaderHeader(v.ValueAt(3))

	return out //from message
}

func To_SpConsensusSlotsSlot(in any) *pbgear.SpConsensusSlotsSlot {
	out := &pbgear.SpConsensusSlotsSlot{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_uint64(v.ValueAt(0))

	return out //from message
}

func To_SpCoreCryptoAccountId32(in any) *pbgear.SpCoreCryptoAccountId32 {
	out := &pbgear.SpCoreCryptoAccountId32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_SpCoreEd25519Public(in any) *pbgear.SpCoreEd25519Public {
	out := &pbgear.SpCoreEd25519Public{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_SpCoreEd25519Signature(in any) *pbgear.SpCoreEd25519Signature {
	out := &pbgear.SpCoreEd25519Signature{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_SpCoreSr25519Public(in any) *pbgear.SpCoreSr25519Public {
	out := &pbgear.SpCoreSr25519Public{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_SpCoreSr25519Signature(in any) *pbgear.SpCoreSr25519Signature {
	out := &pbgear.SpCoreSr25519Signature{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_SpNposElectionsElectionScore(in any) *pbgear.SpNposElectionsElectionScore {
	out := &pbgear.SpNposElectionsElectionScore{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field MinimalStake
	out.MinimalStake = To_string(v.ValueAt(0))
	// primitive field SumStake
	out.SumStake = To_string(v.ValueAt(1))
	// primitive field SumStakeSquared
	out.SumStakeSquared = To_string(v.ValueAt(2))

	return out //from message
}

func To_SpNposElectionsSupport(in any) *pbgear.SpNposElectionsSupport {
	out := &pbgear.SpNposElectionsSupport{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Total
	out.Total = To_string(v.ValueAt(0))
	// repeated field Voters
	out.Voters = To_Repeated_SpNposElectionsSupport_voters(v.ValueAt(1))

	return out //from message
}

func To_Repeated_SpNposElectionsSupport_voters(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_SpRuntimeGenericDigestDigest(in any) *pbgear.SpRuntimeGenericDigestDigest {
	out := &pbgear.SpRuntimeGenericDigestDigest{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Logs
	out.Logs = To_Repeated_SpRuntimeGenericDigestDigest_logs(v.ValueAt(0))

	return out //from message
}

func To_Repeated_SpRuntimeGenericDigestDigest_logs(in any) []*pbgear.SpRuntimeGenericDigestDigestItem {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeGenericDigestDigestItem
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpRuntimeGenericDigestDigestItem(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_SpRuntimeGenericDigestDigestItem(in any) *pbgear.SpRuntimeGenericDigestDigestItem {
	out := &pbgear.SpRuntimeGenericDigestDigestItem{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Logs
	v0 := To_OneOf_SpRuntimeGenericDigestDigestItem_logs(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Logs").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_SpRuntimeGenericDigestDigestItem_logs(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 6:
		return &pbgear.SpRuntimeGenericDigestDigestItem_LogsPreRuntime{
			LogsPreRuntime: To_Babe_PreRuntime(param),
		}
	case 4:
		return &pbgear.SpRuntimeGenericDigestDigestItem_LogsConsensus{
			LogsConsensus: To_Babe_Consensus(param),
		}
	case 5:
		return &pbgear.SpRuntimeGenericDigestDigestItem_LogsSeal{
			LogsSeal: To_Babe_Seal(param),
		}
	case 0:
		return &pbgear.SpRuntimeGenericDigestDigestItem_LogsOther{
			LogsOther: To_Babe_Other(param),
		}
	case 8:
		return &pbgear.SpRuntimeGenericDigestDigestItem_LogsRuntimeEnvironmentUpdated{
			LogsRuntimeEnvironmentUpdated: To_Babe_RuntimeEnvironmentUpdated(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_SpRuntimeGenericHeaderHeader(in any) *pbgear.SpRuntimeGenericHeaderHeader {
	out := &pbgear.SpRuntimeGenericHeaderHeader{}
	v := in.(registry.Valuable)
	_ = v

	// field ParentHash To_PrimitiveTypesH256(w)
	out.ParentHash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field Number
	out.Number = To_uint32(v.ValueAt(1))
	// field StateRoot To_PrimitiveTypesH256(w)
	out.StateRoot = To_PrimitiveTypesH256(v.ValueAt(2))
	// field ExtrinsicsRoot To_PrimitiveTypesH256(w)
	out.ExtrinsicsRoot = To_PrimitiveTypesH256(v.ValueAt(3))
	// field Digest To_SpRuntimeGenericDigestDigest(w)
	out.Digest = To_SpRuntimeGenericDigestDigest(v.ValueAt(4))

	return out //from message
}

func To_SpRuntimeMultiaddressMultiAddress(in any) *pbgear.SpRuntimeMultiaddressMultiAddress {
	out := &pbgear.SpRuntimeMultiaddressMultiAddress{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Targets
	v0 := To_OneOf_SpRuntimeMultiaddressMultiAddress_targets(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Targets").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_SpRuntimeMultiaddressMultiAddress_targets(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsId{
			TargetsId: To_Staking_Id(param),
		}
	case 1:
		return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsIndex{
			TargetsIndex: To_Staking_Index(param),
		}
	case 2:
		return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsRaw{
			TargetsRaw: To_Staking_Raw(param),
		}
	case 3:
		return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsAddress32{
			TargetsAddress32: To_Staking_Address32(param),
		}
	case 4:
		return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsAddress20{
			TargetsAddress20: To_Staking_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_SpSessionMembershipProof(in any) *pbgear.SpSessionMembershipProof {
	out := &pbgear.SpSessionMembershipProof{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Session
	out.Session = To_uint32(v.ValueAt(0))
	// repeated field TrieNodes
	out.TrieNodes = To_Repeated_SpSessionMembershipProof_trie_nodes(v.ValueAt(1))
	// primitive field ValidatorCount
	out.ValidatorCount = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Repeated_SpSessionMembershipProof_trie_nodes(in any) []*pbgear.Babe_BabeTrieNodesList {
	items := in.([]interface{})

	var out []*pbgear.Babe_BabeTrieNodesList
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_Babe_BabeTrieNodesList(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_SpWeightsWeightV2Weight(in any) *pbgear.SpWeightsWeightV2Weight {
	out := &pbgear.SpWeightsWeightV2Weight{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field RefTime
	out.RefTime = To_uint64(v.ValueAt(0))
	// primitive field ProofSize
	out.ProofSize = To_uint64(v.ValueAt(1))

	return out //from message
}

func To_StakingAdmin(in any) *pbgear.StakingAdmin {
	out := &pbgear.StakingAdmin{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_StakingPallet(in any) *pbgear.StakingPallet {
	out := &pbgear.StakingPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_StakingPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_StakingPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingPallet_CallBondCall{
			CallBondCall: To_Staking_BondCall(param),
		}
	case 1:
		return &pbgear.StakingPallet_CallBondExtraCall{
			CallBondExtraCall: To_Staking_BondExtraCall(param),
		}
	case 2:
		return &pbgear.StakingPallet_CallUnbondCall{
			CallUnbondCall: To_Staking_UnbondCall(param),
		}
	case 3:
		return &pbgear.StakingPallet_CallWithdrawUnbondedCall{
			CallWithdrawUnbondedCall: To_Staking_WithdrawUnbondedCall(param),
		}
	case 4:
		return &pbgear.StakingPallet_CallValidateCall{
			CallValidateCall: To_Staking_ValidateCall(param),
		}
	case 5:
		return &pbgear.StakingPallet_CallNominateCall{
			CallNominateCall: To_Staking_NominateCall(param),
		}
	case 6:
		return &pbgear.StakingPallet_CallChillCall{
			CallChillCall: To_Staking_ChillCall(param),
		}
	case 7:
		return &pbgear.StakingPallet_CallSetPayeeCall{
			CallSetPayeeCall: To_Staking_SetPayeeCall(param),
		}
	case 8:
		return &pbgear.StakingPallet_CallSetControllerCall{
			CallSetControllerCall: To_Staking_SetControllerCall(param),
		}
	case 9:
		return &pbgear.StakingPallet_CallSetValidatorCountCall{
			CallSetValidatorCountCall: To_Staking_SetValidatorCountCall(param),
		}
	case 10:
		return &pbgear.StakingPallet_CallIncreaseValidatorCountCall{
			CallIncreaseValidatorCountCall: To_Staking_IncreaseValidatorCountCall(param),
		}
	case 11:
		return &pbgear.StakingPallet_CallScaleValidatorCountCall{
			CallScaleValidatorCountCall: To_Staking_ScaleValidatorCountCall(param),
		}
	case 12:
		return &pbgear.StakingPallet_CallForceNoErasCall{
			CallForceNoErasCall: To_Staking_ForceNoErasCall(param),
		}
	case 13:
		return &pbgear.StakingPallet_CallForceNewEraCall{
			CallForceNewEraCall: To_Staking_ForceNewEraCall(param),
		}
	case 14:
		return &pbgear.StakingPallet_CallSetInvulnerablesCall{
			CallSetInvulnerablesCall: To_Staking_SetInvulnerablesCall(param),
		}
	case 15:
		return &pbgear.StakingPallet_CallForceUnstakeCall{
			CallForceUnstakeCall: To_Staking_ForceUnstakeCall(param),
		}
	case 16:
		return &pbgear.StakingPallet_CallForceNewEraAlwaysCall{
			CallForceNewEraAlwaysCall: To_Staking_ForceNewEraAlwaysCall(param),
		}
	case 17:
		return &pbgear.StakingPallet_CallCancelDeferredSlashCall{
			CallCancelDeferredSlashCall: To_Staking_CancelDeferredSlashCall(param),
		}
	case 18:
		return &pbgear.StakingPallet_CallPayoutStakersCall{
			CallPayoutStakersCall: To_Staking_PayoutStakersCall(param),
		}
	case 19:
		return &pbgear.StakingPallet_CallRebondCall{
			CallRebondCall: To_Staking_RebondCall(param),
		}
	case 20:
		return &pbgear.StakingPallet_CallReapStashCall{
			CallReapStashCall: To_Staking_ReapStashCall(param),
		}
	case 21:
		return &pbgear.StakingPallet_CallKickCall{
			CallKickCall: To_Staking_KickCall(param),
		}
	case 22:
		return &pbgear.StakingPallet_CallSetStakingConfigsCall{
			CallSetStakingConfigsCall: To_Staking_SetStakingConfigsCall(param),
		}
	case 23:
		return &pbgear.StakingPallet_CallChillOtherCall{
			CallChillOtherCall: To_Staking_ChillOtherCall(param),
		}
	case 24:
		return &pbgear.StakingPallet_CallForceApplyMinCommissionCall{
			CallForceApplyMinCommissionCall: To_Staking_ForceApplyMinCommissionCall(param),
		}
	case 25:
		return &pbgear.StakingPallet_CallSetMinCommissionCall{
			CallSetMinCommissionCall: To_Staking_SetMinCommissionCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_StakingRewardsPallet(in any) *pbgear.StakingRewardsPallet {
	out := &pbgear.StakingRewardsPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_StakingRewardsPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_StakingRewardsPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingRewardsPallet_CallRefillCall{
			CallRefillCall: To_StakingRewards_RefillCall(param),
		}
	case 1:
		return &pbgear.StakingRewardsPallet_CallForceRefillCall{
			CallForceRefillCall: To_StakingRewards_ForceRefillCall(param),
		}
	case 2:
		return &pbgear.StakingRewardsPallet_CallWithdrawCall{
			CallWithdrawCall: To_StakingRewards_WithdrawCall(param),
		}
	case 3:
		return &pbgear.StakingRewardsPallet_CallAlignSupplyCall{
			CallAlignSupplyCall: To_StakingRewards_AlignSupplyCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_StakingRewards_Address20(in any) *pbgear.StakingRewards_Address20 {
	out := &pbgear.StakingRewards_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_Address32(in any) *pbgear.StakingRewards_Address32 {
	out := &pbgear.StakingRewards_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_AlignSupplyCall(in any) *pbgear.StakingRewards_AlignSupplyCall {
	out := &pbgear.StakingRewards_AlignSupplyCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Target
	out.Target = To_string(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_ForceRefillCall(in any) *pbgear.StakingRewards_ForceRefillCall {
	out := &pbgear.StakingRewards_ForceRefillCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field From
	v0 := To_OneOf_StakingRewards_ForceRefillCall_from(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))
	// primitive field Value
	out.Value = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_StakingRewards_ForceRefillCall_from(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingRewards_ForceRefillCall_FromId{
			FromId: To_StakingRewards_Id(param),
		}
	case 1:
		return &pbgear.StakingRewards_ForceRefillCall_FromIndex{
			FromIndex: To_StakingRewards_Index(param),
		}
	case 2:
		return &pbgear.StakingRewards_ForceRefillCall_FromRaw{
			FromRaw: To_StakingRewards_Raw(param),
		}
	case 3:
		return &pbgear.StakingRewards_ForceRefillCall_FromAddress32{
			FromAddress32: To_StakingRewards_Address32(param),
		}
	case 4:
		return &pbgear.StakingRewards_ForceRefillCall_FromAddress20{
			FromAddress20: To_StakingRewards_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_StakingRewards_From(in any) *pbgear.StakingRewards_From {
	out := &pbgear.StakingRewards_From{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field From
	v0 := To_OneOf_StakingRewards_From_from(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_StakingRewards_From_from(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingRewards_From_FromId{
			FromId: To_StakingRewards_Id(param),
		}
	case 1:
		return &pbgear.StakingRewards_From_FromIndex{
			FromIndex: To_StakingRewards_Index(param),
		}
	case 2:
		return &pbgear.StakingRewards_From_FromRaw{
			FromRaw: To_StakingRewards_Raw(param),
		}
	case 3:
		return &pbgear.StakingRewards_From_FromAddress32{
			FromAddress32: To_StakingRewards_Address32(param),
		}
	case 4:
		return &pbgear.StakingRewards_From_FromAddress20{
			FromAddress20: To_StakingRewards_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_StakingRewards_Id(in any) *pbgear.StakingRewards_Id {
	out := &pbgear.StakingRewards_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_Index(in any) *pbgear.StakingRewards_Index {
	out := &pbgear.StakingRewards_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_StakingRewards_TupleNull(w)
	out.Value0 = To_StakingRewards_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_Raw(in any) *pbgear.StakingRewards_Raw {
	out := &pbgear.StakingRewards_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_RefillCall(in any) *pbgear.StakingRewards_RefillCall {
	out := &pbgear.StakingRewards_RefillCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_To(in any) *pbgear.StakingRewards_To {
	out := &pbgear.StakingRewards_To{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field To
	v0 := To_OneOf_StakingRewards_To_to(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_StakingRewards_To_to(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingRewards_To_ToId{
			ToId: To_StakingRewards_Id(param),
		}
	case 1:
		return &pbgear.StakingRewards_To_ToIndex{
			ToIndex: To_StakingRewards_Index(param),
		}
	case 2:
		return &pbgear.StakingRewards_To_ToRaw{
			ToRaw: To_StakingRewards_Raw(param),
		}
	case 3:
		return &pbgear.StakingRewards_To_ToAddress32{
			ToAddress32: To_StakingRewards_Address32(param),
		}
	case 4:
		return &pbgear.StakingRewards_To_ToAddress20{
			ToAddress20: To_StakingRewards_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_StakingRewards_TupleNull(in any) *pbgear.StakingRewards_TupleNull {
	out := &pbgear.StakingRewards_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_StakingRewards_WithdrawCall(in any) *pbgear.StakingRewards_WithdrawCall {
	out := &pbgear.StakingRewards_WithdrawCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field To
	v0 := To_OneOf_StakingRewards_WithdrawCall_to(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))
	// primitive field Value
	out.Value = To_string(v.ValueAt(1))

	return out //from message
}

func To_OneOf_StakingRewards_WithdrawCall_to(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.StakingRewards_WithdrawCall_ToId{
			ToId: To_StakingRewards_Id(param),
		}
	case 1:
		return &pbgear.StakingRewards_WithdrawCall_ToIndex{
			ToIndex: To_StakingRewards_Index(param),
		}
	case 2:
		return &pbgear.StakingRewards_WithdrawCall_ToRaw{
			ToRaw: To_StakingRewards_Raw(param),
		}
	case 3:
		return &pbgear.StakingRewards_WithdrawCall_ToAddress32{
			ToAddress32: To_StakingRewards_Address32(param),
		}
	case 4:
		return &pbgear.StakingRewards_WithdrawCall_ToAddress20{
			ToAddress20: To_StakingRewards_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_Account(in any) *pbgear.Staking_Account {
	out := &pbgear.Staking_Account{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Staking_Address20(in any) *pbgear.Staking_Address20 {
	out := &pbgear.Staking_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Staking_Address32(in any) *pbgear.Staking_Address32 {
	out := &pbgear.Staking_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Staking_BondCall(in any) *pbgear.Staking_BondCall {
	out := &pbgear.Staking_BondCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))
	// oneOf field Payee
	v1 := To_OneOf_Staking_BondCall_payee(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Staking_BondCall_payee(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_BondCall_PayeeStaked{
			PayeeStaked: To_Staking_Staked(param),
		}
	case 1:
		return &pbgear.Staking_BondCall_PayeeStash{
			PayeeStash: To_Staking_Stash(param),
		}
	case 2:
		return &pbgear.Staking_BondCall_PayeeController{
			PayeeController: To_Staking_Controller(param),
		}
	case 3:
		return &pbgear.Staking_BondCall_PayeeAccount{
			PayeeAccount: To_Staking_Account(param),
		}
	case 4:
		return &pbgear.Staking_BondCall_PayeeNone{
			PayeeNone: To_Staking_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_BondExtraCall(in any) *pbgear.Staking_BondExtraCall {
	out := &pbgear.Staking_BondExtraCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field MaxAdditional
	out.MaxAdditional = To_string(v.ValueAt(0))

	return out //from message
}

func To_Staking_CancelDeferredSlashCall(in any) *pbgear.Staking_CancelDeferredSlashCall {
	out := &pbgear.Staking_CancelDeferredSlashCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Era
	out.Era = To_uint32(v.ValueAt(0))
	// primitive field SlashIndices
	out.SlashIndices = To_Repeated_uint32(v.ValueAt(1))

	return out //from message
}

func To_Staking_ChillCall(in any) *pbgear.Staking_ChillCall {
	out := &pbgear.Staking_ChillCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ChillOtherCall(in any) *pbgear.Staking_ChillOtherCall {
	out := &pbgear.Staking_ChillOtherCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Controller To_SpCoreCryptoAccountId32(w)
	out.Controller = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Staking_ChillThreshold(in any) *pbgear.Staking_ChillThreshold {
	out := &pbgear.Staking_ChillThreshold{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field ChillThreshold
	v0 := To_OneOf_Staking_ChillThreshold_chill_threshold(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("ChillThreshold").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_ChillThreshold_chill_threshold(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_ChillThreshold_ChillThresholdNoop{
			ChillThresholdNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_ChillThreshold_ChillThresholdSet{
			ChillThresholdSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_ChillThreshold_ChillThresholdRemove{
			ChillThresholdRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_Controller(in any) *pbgear.Staking_Controller {
	out := &pbgear.Staking_Controller{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ForceApplyMinCommissionCall(in any) *pbgear.Staking_ForceApplyMinCommissionCall {
	out := &pbgear.Staking_ForceApplyMinCommissionCall{}
	v := in.(registry.Valuable)
	_ = v

	// field ValidatorStash To_SpCoreCryptoAccountId32(w)
	out.ValidatorStash = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Staking_ForceNewEraAlwaysCall(in any) *pbgear.Staking_ForceNewEraAlwaysCall {
	out := &pbgear.Staking_ForceNewEraAlwaysCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ForceNewEraCall(in any) *pbgear.Staking_ForceNewEraCall {
	out := &pbgear.Staking_ForceNewEraCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ForceNoErasCall(in any) *pbgear.Staking_ForceNoErasCall {
	out := &pbgear.Staking_ForceNoErasCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ForceUnstakeCall(in any) *pbgear.Staking_ForceUnstakeCall {
	out := &pbgear.Staking_ForceUnstakeCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Stash To_SpCoreCryptoAccountId32(w)
	out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Staking_Id(in any) *pbgear.Staking_Id {
	out := &pbgear.Staking_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Staking_IncreaseValidatorCountCall(in any) *pbgear.Staking_IncreaseValidatorCountCall {
	out := &pbgear.Staking_IncreaseValidatorCountCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Additional
	out.Additional = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Staking_Index(in any) *pbgear.Staking_Index {
	out := &pbgear.Staking_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Staking_TupleNull(w)
	out.Value0 = To_Staking_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Staking_KickCall(in any) *pbgear.Staking_KickCall {
	out := &pbgear.Staking_KickCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Who
	out.Who = To_Repeated_Staking_KickCall_who(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Staking_KickCall_who(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeMultiaddressMultiAddress
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpRuntimeMultiaddressMultiAddress(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Staking_MaxNominatorCount(in any) *pbgear.Staking_MaxNominatorCount {
	out := &pbgear.Staking_MaxNominatorCount{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MaxNominatorCount
	v0 := To_OneOf_Staking_MaxNominatorCount_max_nominator_count(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MaxNominatorCount").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_MaxNominatorCount_max_nominator_count(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountNoop{
			MaxNominatorCountNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountSet{
			MaxNominatorCountSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountRemove{
			MaxNominatorCountRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_MaxValidatorCount(in any) *pbgear.Staking_MaxValidatorCount {
	out := &pbgear.Staking_MaxValidatorCount{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MaxValidatorCount
	v0 := To_OneOf_Staking_MaxValidatorCount_max_validator_count(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MaxValidatorCount").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_MaxValidatorCount_max_validator_count(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountNoop{
			MaxValidatorCountNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountSet{
			MaxValidatorCountSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountRemove{
			MaxValidatorCountRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_MinCommission(in any) *pbgear.Staking_MinCommission {
	out := &pbgear.Staking_MinCommission{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinCommission
	v0 := To_OneOf_Staking_MinCommission_min_commission(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinCommission").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_MinCommission_min_commission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_MinCommission_MinCommissionNoop{
			MinCommissionNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_MinCommission_MinCommissionSet{
			MinCommissionSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_MinCommission_MinCommissionRemove{
			MinCommissionRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_MinNominatorBond(in any) *pbgear.Staking_MinNominatorBond {
	out := &pbgear.Staking_MinNominatorBond{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinNominatorBond
	v0 := To_OneOf_Staking_MinNominatorBond_min_nominator_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinNominatorBond").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_MinNominatorBond_min_nominator_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_MinNominatorBond_MinNominatorBondNoop{
			MinNominatorBondNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_MinNominatorBond_MinNominatorBondSet{
			MinNominatorBondSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_MinNominatorBond_MinNominatorBondRemove{
			MinNominatorBondRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_MinValidatorBond(in any) *pbgear.Staking_MinValidatorBond {
	out := &pbgear.Staking_MinValidatorBond{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinValidatorBond
	v0 := To_OneOf_Staking_MinValidatorBond_min_validator_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinValidatorBond").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_MinValidatorBond_min_validator_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_MinValidatorBond_MinValidatorBondNoop{
			MinValidatorBondNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_MinValidatorBond_MinValidatorBondSet{
			MinValidatorBondSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_MinValidatorBond_MinValidatorBondRemove{
			MinValidatorBondRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_NominateCall(in any) *pbgear.Staking_NominateCall {
	out := &pbgear.Staking_NominateCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Targets
	out.Targets = To_Repeated_Staking_NominateCall_targets(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Staking_NominateCall_targets(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
	items := in.([]interface{})

	var out []*pbgear.SpRuntimeMultiaddressMultiAddress
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpRuntimeMultiaddressMultiAddress(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Staking_None(in any) *pbgear.Staking_None {
	out := &pbgear.Staking_None{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_Noop(in any) *pbgear.Staking_Noop {
	out := &pbgear.Staking_Noop{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_PalletStakingValidatorPrefs(in any) *pbgear.Staking_PalletStakingValidatorPrefs {
	out := &pbgear.Staking_PalletStakingValidatorPrefs{}
	v := in.(registry.Valuable)
	_ = v

	// field Commission To_SpArithmeticPerThingsPerbill(w)
	out.Commission = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
	// primitive field Blocked
	out.Blocked = To_bool(v.ValueAt(1))

	return out //from message
}

func To_Staking_Payee(in any) *pbgear.Staking_Payee {
	out := &pbgear.Staking_Payee{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Payee
	v0 := To_OneOf_Staking_Payee_payee(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_Payee_payee(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_Payee_PayeeStaked{
			PayeeStaked: To_Staking_Staked(param),
		}
	case 1:
		return &pbgear.Staking_Payee_PayeeStash{
			PayeeStash: To_Staking_Stash(param),
		}
	case 2:
		return &pbgear.Staking_Payee_PayeeController{
			PayeeController: To_Staking_Controller(param),
		}
	case 3:
		return &pbgear.Staking_Payee_PayeeAccount{
			PayeeAccount: To_Staking_Account(param),
		}
	case 4:
		return &pbgear.Staking_Payee_PayeeNone{
			PayeeNone: To_Staking_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_PayoutStakersCall(in any) *pbgear.Staking_PayoutStakersCall {
	out := &pbgear.Staking_PayoutStakersCall{}
	v := in.(registry.Valuable)
	_ = v

	// field ValidatorStash To_SpCoreCryptoAccountId32(w)
	out.ValidatorStash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// primitive field Era
	out.Era = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Staking_Raw(in any) *pbgear.Staking_Raw {
	out := &pbgear.Staking_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Staking_ReapStashCall(in any) *pbgear.Staking_ReapStashCall {
	out := &pbgear.Staking_ReapStashCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Stash To_SpCoreCryptoAccountId32(w)
	out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Staking_RebondCall(in any) *pbgear.Staking_RebondCall {
	out := &pbgear.Staking_RebondCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))

	return out //from message
}

func To_Staking_Remove(in any) *pbgear.Staking_Remove {
	out := &pbgear.Staking_Remove{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_ScaleValidatorCountCall(in any) *pbgear.Staking_ScaleValidatorCountCall {
	out := &pbgear.Staking_ScaleValidatorCountCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Factor To_SpArithmeticPerThingsPercent(w)
	out.Factor = To_SpArithmeticPerThingsPercent(v.ValueAt(0))

	return out //from message
}

func To_Staking_Set(in any) *pbgear.Staking_Set {
	out := &pbgear.Staking_Set{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_string(v.ValueAt(0))

	return out //from message
}

func To_Staking_SetControllerCall(in any) *pbgear.Staking_SetControllerCall {
	out := &pbgear.Staking_SetControllerCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_SetInvulnerablesCall(in any) *pbgear.Staking_SetInvulnerablesCall {
	out := &pbgear.Staking_SetInvulnerablesCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Invulnerables
	out.Invulnerables = To_Repeated_Staking_SetInvulnerablesCall_invulnerables(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Staking_SetInvulnerablesCall_invulnerables(in any) []*pbgear.SpCoreCryptoAccountId32 {
	items := in.([]interface{})

	var out []*pbgear.SpCoreCryptoAccountId32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_SpCoreCryptoAccountId32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Staking_SetMinCommissionCall(in any) *pbgear.Staking_SetMinCommissionCall {
	out := &pbgear.Staking_SetMinCommissionCall{}
	v := in.(registry.Valuable)
	_ = v

	// field New To_SpArithmeticPerThingsPerbill(w)
	out.New = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))

	return out //from message
}

func To_Staking_SetPayeeCall(in any) *pbgear.Staking_SetPayeeCall {
	out := &pbgear.Staking_SetPayeeCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Payee
	v0 := To_OneOf_Staking_SetPayeeCall_payee(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_SetPayeeCall_payee(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetPayeeCall_PayeeStaked{
			PayeeStaked: To_Staking_Staked(param),
		}
	case 1:
		return &pbgear.Staking_SetPayeeCall_PayeeStash{
			PayeeStash: To_Staking_Stash(param),
		}
	case 2:
		return &pbgear.Staking_SetPayeeCall_PayeeController{
			PayeeController: To_Staking_Controller(param),
		}
	case 3:
		return &pbgear.Staking_SetPayeeCall_PayeeAccount{
			PayeeAccount: To_Staking_Account(param),
		}
	case 4:
		return &pbgear.Staking_SetPayeeCall_PayeeNone{
			PayeeNone: To_Staking_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_SetStakingConfigsCall(in any) *pbgear.Staking_SetStakingConfigsCall {
	out := &pbgear.Staking_SetStakingConfigsCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field MinNominatorBond
	v0 := To_OneOf_Staking_SetStakingConfigsCall_min_nominator_bond(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("MinNominatorBond").Set(reflect.ValueOf(v0))
	// oneOf field MinValidatorBond
	v1 := To_OneOf_Staking_SetStakingConfigsCall_min_validator_bond(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("MinValidatorBond").Set(reflect.ValueOf(v1))
	// oneOf field MaxNominatorCount
	v2 := To_OneOf_Staking_SetStakingConfigsCall_max_nominator_count(v.ValueAt(2))
	reflect.ValueOf(out).Elem().FieldByName("MaxNominatorCount").Set(reflect.ValueOf(v2))
	// oneOf field MaxValidatorCount
	v3 := To_OneOf_Staking_SetStakingConfigsCall_max_validator_count(v.ValueAt(3))
	reflect.ValueOf(out).Elem().FieldByName("MaxValidatorCount").Set(reflect.ValueOf(v3))
	// oneOf field ChillThreshold
	v4 := To_OneOf_Staking_SetStakingConfigsCall_chill_threshold(v.ValueAt(4))
	reflect.ValueOf(out).Elem().FieldByName("ChillThreshold").Set(reflect.ValueOf(v4))
	// oneOf field MinCommission
	v5 := To_OneOf_Staking_SetStakingConfigsCall_min_commission(v.ValueAt(5))
	reflect.ValueOf(out).Elem().FieldByName("MinCommission").Set(reflect.ValueOf(v5))

	return out //from message
}

func To_OneOf_Staking_SetStakingConfigsCall_min_nominator_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondNoop{
			MinNominatorBondNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondSet{
			MinNominatorBondSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondRemove{
			MinNominatorBondRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Staking_SetStakingConfigsCall_min_validator_bond(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondNoop{
			MinValidatorBondNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondSet{
			MinValidatorBondSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondRemove{
			MinValidatorBondRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Staking_SetStakingConfigsCall_max_nominator_count(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountNoop{
			MaxNominatorCountNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountSet{
			MaxNominatorCountSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountRemove{
			MaxNominatorCountRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Staking_SetStakingConfigsCall_max_validator_count(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountNoop{
			MaxValidatorCountNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountSet{
			MaxValidatorCountSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountRemove{
			MaxValidatorCountRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Staking_SetStakingConfigsCall_chill_threshold(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdNoop{
			ChillThresholdNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdSet{
			ChillThresholdSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdRemove{
			ChillThresholdRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Staking_SetStakingConfigsCall_min_commission(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_SetStakingConfigsCall_MinCommissionNoop{
			MinCommissionNoop: To_Staking_Noop(param),
		}
	case 1:
		return &pbgear.Staking_SetStakingConfigsCall_MinCommissionSet{
			MinCommissionSet: To_Staking_Set(param),
		}
	case 2:
		return &pbgear.Staking_SetStakingConfigsCall_MinCommissionRemove{
			MinCommissionRemove: To_Staking_Remove(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_SetValidatorCountCall(in any) *pbgear.Staking_SetValidatorCountCall {
	out := &pbgear.Staking_SetValidatorCountCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field New
	out.New = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Staking_Staked(in any) *pbgear.Staking_Staked {
	out := &pbgear.Staking_Staked{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_Stash(in any) *pbgear.Staking_Stash {
	out := &pbgear.Staking_Stash{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Staking_Targets(in any) *pbgear.Staking_Targets {
	out := &pbgear.Staking_Targets{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Targets
	v0 := To_OneOf_Staking_Targets_targets(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Targets").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_Targets_targets(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_Targets_TargetsId{
			TargetsId: To_Staking_Id(param),
		}
	case 1:
		return &pbgear.Staking_Targets_TargetsIndex{
			TargetsIndex: To_Staking_Index(param),
		}
	case 2:
		return &pbgear.Staking_Targets_TargetsRaw{
			TargetsRaw: To_Staking_Raw(param),
		}
	case 3:
		return &pbgear.Staking_Targets_TargetsAddress32{
			TargetsAddress32: To_Staking_Address32(param),
		}
	case 4:
		return &pbgear.Staking_Targets_TargetsAddress20{
			TargetsAddress20: To_Staking_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_TupleNull(in any) *pbgear.Staking_TupleNull {
	out := &pbgear.Staking_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Staking_UnbondCall(in any) *pbgear.Staking_UnbondCall {
	out := &pbgear.Staking_UnbondCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))

	return out //from message
}

func To_Staking_ValidateCall(in any) *pbgear.Staking_ValidateCall {
	out := &pbgear.Staking_ValidateCall{}
	v := in.(registry.Valuable)
	_ = v

	// field Prefs To_Staking_PalletStakingValidatorPrefs(w)
	out.Prefs = To_Staking_PalletStakingValidatorPrefs(v.ValueAt(0))

	return out //from message
}

func To_Staking_Who(in any) *pbgear.Staking_Who {
	out := &pbgear.Staking_Who{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Who
	v0 := To_OneOf_Staking_Who_who(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Staking_Who_who(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Staking_Who_WhoId{
			WhoId: To_Staking_Id(param),
		}
	case 1:
		return &pbgear.Staking_Who_WhoIndex{
			WhoIndex: To_Staking_Index(param),
		}
	case 2:
		return &pbgear.Staking_Who_WhoRaw{
			WhoRaw: To_Staking_Raw(param),
		}
	case 3:
		return &pbgear.Staking_Who_WhoAddress32{
			WhoAddress32: To_Staking_Address32(param),
		}
	case 4:
		return &pbgear.Staking_Who_WhoAddress20{
			WhoAddress20: To_Staking_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Staking_WithdrawUnbondedCall(in any) *pbgear.Staking_WithdrawUnbondedCall {
	out := &pbgear.Staking_WithdrawUnbondedCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field NumSlashingSpans
	out.NumSlashingSpans = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_SystemPallet(in any) *pbgear.SystemPallet {
	out := &pbgear.SystemPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_SystemPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_SystemPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.SystemPallet_CallRemarkCall{
			CallRemarkCall: To_System_RemarkCall(param),
		}
	case 1:
		return &pbgear.SystemPallet_CallSetHeapPagesCall{
			CallSetHeapPagesCall: To_System_SetHeapPagesCall(param),
		}
	case 2:
		return &pbgear.SystemPallet_CallSetCodeCall{
			CallSetCodeCall: To_System_SetCodeCall(param),
		}
	case 3:
		return &pbgear.SystemPallet_CallSetCodeWithoutChecksCall{
			CallSetCodeWithoutChecksCall: To_System_SetCodeWithoutChecksCall(param),
		}
	case 4:
		return &pbgear.SystemPallet_CallSetStorageCall{
			CallSetStorageCall: To_System_SetStorageCall(param),
		}
	case 5:
		return &pbgear.SystemPallet_CallKillStorageCall{
			CallKillStorageCall: To_System_KillStorageCall(param),
		}
	case 6:
		return &pbgear.SystemPallet_CallKillPrefixCall{
			CallKillPrefixCall: To_System_KillPrefixCall(param),
		}
	case 7:
		return &pbgear.SystemPallet_CallRemarkWithEventCall{
			CallRemarkWithEventCall: To_System_RemarkWithEventCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_System_KillPrefixCall(in any) *pbgear.System_KillPrefixCall {
	out := &pbgear.System_KillPrefixCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Prefix
	out.Prefix = To_bytes(v.ValueAt(0))
	// primitive field Subkeys
	out.Subkeys = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_System_KillStorageCall(in any) *pbgear.System_KillStorageCall {
	out := &pbgear.System_KillStorageCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Keys
	out.Keys = To_Repeated_System_KillStorageCall_keys(v.ValueAt(0))

	return out //from message
}

func To_Repeated_System_KillStorageCall_keys(in any) []*pbgear.System_SystemKeysList {
	items := in.([]interface{})

	var out []*pbgear.System_SystemKeysList
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_System_SystemKeysList(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_System_RemarkCall(in any) *pbgear.System_RemarkCall {
	out := &pbgear.System_RemarkCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Remark
	out.Remark = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_System_RemarkWithEventCall(in any) *pbgear.System_RemarkWithEventCall {
	out := &pbgear.System_RemarkWithEventCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Remark
	out.Remark = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_System_SetCodeCall(in any) *pbgear.System_SetCodeCall {
	out := &pbgear.System_SetCodeCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Code
	out.Code = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_System_SetCodeWithoutChecksCall(in any) *pbgear.System_SetCodeWithoutChecksCall {
	out := &pbgear.System_SetCodeWithoutChecksCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Code
	out.Code = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_System_SetHeapPagesCall(in any) *pbgear.System_SetHeapPagesCall {
	out := &pbgear.System_SetHeapPagesCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Pages
	out.Pages = To_uint64(v.ValueAt(0))

	return out //from message
}

func To_System_SetStorageCall(in any) *pbgear.System_SetStorageCall {
	out := &pbgear.System_SetStorageCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Items
	out.Items = To_Repeated_System_SetStorageCall_items(v.ValueAt(0))

	return out //from message
}

func To_Repeated_System_SetStorageCall_items(in any) []*pbgear.System_TupleSystemItemsListSystemItemsList {
	items := in.([]interface{})

	var out []*pbgear.System_TupleSystemItemsListSystemItemsList
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_System_TupleSystemItemsListSystemItemsList(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_System_SystemKeysList(in any) *pbgear.System_SystemKeysList {
	out := &pbgear.System_SystemKeysList{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Keys
	out.Keys = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_System_TupleSystemItemsListSystemItemsList(in any) *pbgear.System_TupleSystemItemsListSystemItemsList {
	out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))
	// primitive field Value1
	out.Value1 = To_bytes(v.ValueAt(1))

	return out //from message
}

func To_TimestampPallet(in any) *pbgear.TimestampPallet {
	out := &pbgear.TimestampPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_TimestampPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_TimestampPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.TimestampPallet_CallSetCall{
			CallSetCall: To_Timestamp_SetCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Timestamp_SetCall(in any) *pbgear.Timestamp_SetCall {
	out := &pbgear.Timestamp_SetCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Now
	out.Now = To_uint64(v.ValueAt(0))

	return out //from message
}

func To_TransactionPaymentPallet(in any) *pbgear.TransactionPaymentPallet {
	out := &pbgear.TransactionPaymentPallet{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Treasurer(in any) *pbgear.Treasurer {
	out := &pbgear.Treasurer{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_TreasuryPallet(in any) *pbgear.TreasuryPallet {
	out := &pbgear.TreasuryPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_TreasuryPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_TreasuryPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.TreasuryPallet_CallProposeSpendCall{
			CallProposeSpendCall: To_Treasury_ProposeSpendCall(param),
		}
	case 1:
		return &pbgear.TreasuryPallet_CallRejectProposalCall{
			CallRejectProposalCall: To_Treasury_RejectProposalCall(param),
		}
	case 2:
		return &pbgear.TreasuryPallet_CallApproveProposalCall{
			CallApproveProposalCall: To_Treasury_ApproveProposalCall(param),
		}
	case 3:
		return &pbgear.TreasuryPallet_CallSpendLocalCall{
			CallSpendLocalCall: To_Treasury_SpendLocalCall(param),
		}
	case 4:
		return &pbgear.TreasuryPallet_CallRemoveApprovalCall{
			CallRemoveApprovalCall: To_Treasury_RemoveApprovalCall(param),
		}
	case 5:
		return &pbgear.TreasuryPallet_CallSpendCall{
			CallSpendCall: To_Treasury_SpendCall(param),
		}
	case 6:
		return &pbgear.TreasuryPallet_CallPayoutCall{
			CallPayoutCall: To_Treasury_PayoutCall(param),
		}
	case 7:
		return &pbgear.TreasuryPallet_CallCheckStatusCall{
			CallCheckStatusCall: To_Treasury_CheckStatusCall(param),
		}
	case 8:
		return &pbgear.TreasuryPallet_CallVoidSpendCall{
			CallVoidSpendCall: To_Treasury_VoidSpendCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Treasury_Address20(in any) *pbgear.Treasury_Address20 {
	out := &pbgear.Treasury_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Treasury_Address32(in any) *pbgear.Treasury_Address32 {
	out := &pbgear.Treasury_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Treasury_ApproveProposalCall(in any) *pbgear.Treasury_ApproveProposalCall {
	out := &pbgear.Treasury_ApproveProposalCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ProposalId
	out.ProposalId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_Beneficiary(in any) *pbgear.Treasury_Beneficiary {
	out := &pbgear.Treasury_Beneficiary{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Beneficiary
	v0 := To_OneOf_Treasury_Beneficiary_beneficiary(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Treasury_Beneficiary_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Treasury_Beneficiary_BeneficiaryId{
			BeneficiaryId: To_Treasury_Id(param),
		}
	case 1:
		return &pbgear.Treasury_Beneficiary_BeneficiaryIndex{
			BeneficiaryIndex: To_Treasury_Index(param),
		}
	case 2:
		return &pbgear.Treasury_Beneficiary_BeneficiaryRaw{
			BeneficiaryRaw: To_Treasury_Raw(param),
		}
	case 3:
		return &pbgear.Treasury_Beneficiary_BeneficiaryAddress32{
			BeneficiaryAddress32: To_Treasury_Address32(param),
		}
	case 4:
		return &pbgear.Treasury_Beneficiary_BeneficiaryAddress20{
			BeneficiaryAddress20: To_Treasury_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Treasury_CheckStatusCall(in any) *pbgear.Treasury_CheckStatusCall {
	out := &pbgear.Treasury_CheckStatusCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_Id(in any) *pbgear.Treasury_Id {
	out := &pbgear.Treasury_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_Index(in any) *pbgear.Treasury_Index {
	out := &pbgear.Treasury_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Treasury_TupleNull(w)
	out.Value0 = To_Treasury_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Treasury_PayoutCall(in any) *pbgear.Treasury_PayoutCall {
	out := &pbgear.Treasury_PayoutCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_ProposeSpendCall(in any) *pbgear.Treasury_ProposeSpendCall {
	out := &pbgear.Treasury_ProposeSpendCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value
	out.Value = To_string(v.ValueAt(0))
	// oneOf field Beneficiary
	v1 := To_OneOf_Treasury_ProposeSpendCall_beneficiary(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Treasury_ProposeSpendCall_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Treasury_ProposeSpendCall_BeneficiaryId{
			BeneficiaryId: To_Treasury_Id(param),
		}
	case 1:
		return &pbgear.Treasury_ProposeSpendCall_BeneficiaryIndex{
			BeneficiaryIndex: To_Treasury_Index(param),
		}
	case 2:
		return &pbgear.Treasury_ProposeSpendCall_BeneficiaryRaw{
			BeneficiaryRaw: To_Treasury_Raw(param),
		}
	case 3:
		return &pbgear.Treasury_ProposeSpendCall_BeneficiaryAddress32{
			BeneficiaryAddress32: To_Treasury_Address32(param),
		}
	case 4:
		return &pbgear.Treasury_ProposeSpendCall_BeneficiaryAddress20{
			BeneficiaryAddress20: To_Treasury_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Treasury_Raw(in any) *pbgear.Treasury_Raw {
	out := &pbgear.Treasury_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Treasury_RejectProposalCall(in any) *pbgear.Treasury_RejectProposalCall {
	out := &pbgear.Treasury_RejectProposalCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ProposalId
	out.ProposalId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_RemoveApprovalCall(in any) *pbgear.Treasury_RemoveApprovalCall {
	out := &pbgear.Treasury_RemoveApprovalCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field ProposalId
	out.ProposalId = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_Treasury_SpendCall(in any) *pbgear.Treasury_SpendCall {
	out := &pbgear.Treasury_SpendCall{}
	v := in.(registry.Valuable)
	_ = v

	// field AssetKind To_Treasury_TupleNull(w)
	out.AssetKind = To_Treasury_TupleNull(v.ValueAt(0))
	// primitive field Amount
	out.Amount = To_string(v.ValueAt(1))
	// field Beneficiary To_SpCoreCryptoAccountId32(w)
	out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(2))
	// primitive field ValidFrom
	out.ValidFrom = To_Optional_uint32(v.ValueAt(3))

	return out //from message
}

func To_Treasury_SpendLocalCall(in any) *pbgear.Treasury_SpendLocalCall {
	out := &pbgear.Treasury_SpendLocalCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Amount
	out.Amount = To_string(v.ValueAt(0))
	// oneOf field Beneficiary
	v1 := To_OneOf_Treasury_SpendLocalCall_beneficiary(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Treasury_SpendLocalCall_beneficiary(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Treasury_SpendLocalCall_BeneficiaryId{
			BeneficiaryId: To_Treasury_Id(param),
		}
	case 1:
		return &pbgear.Treasury_SpendLocalCall_BeneficiaryIndex{
			BeneficiaryIndex: To_Treasury_Index(param),
		}
	case 2:
		return &pbgear.Treasury_SpendLocalCall_BeneficiaryRaw{
			BeneficiaryRaw: To_Treasury_Raw(param),
		}
	case 3:
		return &pbgear.Treasury_SpendLocalCall_BeneficiaryAddress32{
			BeneficiaryAddress32: To_Treasury_Address32(param),
		}
	case 4:
		return &pbgear.Treasury_SpendLocalCall_BeneficiaryAddress20{
			BeneficiaryAddress20: To_Treasury_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Treasury_TupleNull(in any) *pbgear.Treasury_TupleNull {
	out := &pbgear.Treasury_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Treasury_VoidSpendCall(in any) *pbgear.Treasury_VoidSpendCall {
	out := &pbgear.Treasury_VoidSpendCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))

	return out //from message
}

func To_TupleNull(in any) *pbgear.TupleNull {
	out := &pbgear.TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_UtilityPallet(in any) *pbgear.UtilityPallet {
	out := &pbgear.UtilityPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_UtilityPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_UtilityPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.UtilityPallet_CallBatchCall{
			CallBatchCall: To_Utility_BatchCall(param),
		}
	case 1:
		return &pbgear.UtilityPallet_CallAsDerivativeCall{
			CallAsDerivativeCall: To_Utility_AsDerivativeCall(param),
		}
	case 2:
		return &pbgear.UtilityPallet_CallBatchAllCall{
			CallBatchAllCall: To_Utility_BatchAllCall(param),
		}
	case 3:
		return &pbgear.UtilityPallet_CallDispatchAsCall{
			CallDispatchAsCall: To_Utility_DispatchAsCall(param),
		}
	case 4:
		return &pbgear.UtilityPallet_CallForceBatchCall{
			CallForceBatchCall: To_Utility_ForceBatchCall(param),
		}
	case 5:
		return &pbgear.UtilityPallet_CallWithWeightCall{
			CallWithWeightCall: To_Utility_WithWeightCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_AsDerivativeCall(in any) *pbgear.Utility_AsDerivativeCall {
	out := &pbgear.Utility_AsDerivativeCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Index
	out.Index = To_uint32(v.ValueAt(0))
	// oneOf field Call
	v1 := To_OneOf_Utility_AsDerivativeCall_call(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Utility_AsDerivativeCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_AsDerivativeCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Utility_AsDerivativeCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Utility_AsDerivativeCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Utility_AsDerivativeCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Utility_AsDerivativeCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Utility_AsDerivativeCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Utility_AsDerivativeCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Utility_AsDerivativeCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Utility_AsDerivativeCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Utility_AsDerivativeCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Utility_AsDerivativeCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Utility_AsDerivativeCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Utility_AsDerivativeCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Utility_AsDerivativeCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Utility_AsDerivativeCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Utility_AsDerivativeCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Utility_AsDerivativeCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Utility_AsDerivativeCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Utility_AsDerivativeCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Utility_AsDerivativeCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Utility_AsDerivativeCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Utility_AsDerivativeCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Utility_AsDerivativeCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Utility_AsDerivativeCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Utility_AsDerivativeCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Utility_AsDerivativeCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Utility_AsDerivativeCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Utility_AsDerivativeCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Utility_AsDerivativeCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_AsOrigin(in any) *pbgear.Utility_AsOrigin {
	out := &pbgear.Utility_AsOrigin{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field AsOrigin
	v0 := To_OneOf_Utility_AsOrigin_as_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("AsOrigin").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Utility_AsOrigin_as_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_AsOrigin_AsOriginSystem{
			AsOriginSystem: To_Utility_System(param),
		}
	case 20:
		return &pbgear.Utility_AsOrigin_AsOriginOrigins{
			AsOriginOrigins: To_Utility_Origins(param),
		}
	case 2:
		return &pbgear.Utility_AsOrigin_AsOriginVoid{
			AsOriginVoid: To_Utility_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_BatchAllCall(in any) *pbgear.Utility_BatchAllCall {
	out := &pbgear.Utility_BatchAllCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Calls
	out.Calls = To_Repeated_Utility_BatchAllCall_calls(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Utility_BatchAllCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_VaraRuntimeRuntimeCall(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Utility_BatchCall(in any) *pbgear.Utility_BatchCall {
	out := &pbgear.Utility_BatchCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Calls
	out.Calls = To_Repeated_Utility_BatchCall_calls(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Utility_BatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_VaraRuntimeRuntimeCall(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Utility_DispatchAsCall(in any) *pbgear.Utility_DispatchAsCall {
	out := &pbgear.Utility_DispatchAsCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field AsOrigin
	v0 := To_OneOf_Utility_DispatchAsCall_as_origin(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("AsOrigin").Set(reflect.ValueOf(v0))
	// oneOf field Call
	v1 := To_OneOf_Utility_DispatchAsCall_call(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

	return out //from message
}

func To_OneOf_Utility_DispatchAsCall_as_origin(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_DispatchAsCall_AsOriginSystem{
			AsOriginSystem: To_Utility_System(param),
		}
	case 20:
		return &pbgear.Utility_DispatchAsCall_AsOriginOrigins{
			AsOriginOrigins: To_Utility_Origins(param),
		}
	case 2:
		return &pbgear.Utility_DispatchAsCall_AsOriginVoid{
			AsOriginVoid: To_Utility_Void(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Utility_DispatchAsCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_DispatchAsCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Utility_DispatchAsCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Utility_DispatchAsCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Utility_DispatchAsCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Utility_DispatchAsCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Utility_DispatchAsCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Utility_DispatchAsCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Utility_DispatchAsCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Utility_DispatchAsCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Utility_DispatchAsCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Utility_DispatchAsCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Utility_DispatchAsCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Utility_DispatchAsCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Utility_DispatchAsCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Utility_DispatchAsCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Utility_DispatchAsCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Utility_DispatchAsCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Utility_DispatchAsCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Utility_DispatchAsCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Utility_DispatchAsCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Utility_DispatchAsCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Utility_DispatchAsCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Utility_DispatchAsCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Utility_DispatchAsCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Utility_DispatchAsCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Utility_DispatchAsCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Utility_DispatchAsCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Utility_DispatchAsCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Utility_DispatchAsCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_ForceBatchCall(in any) *pbgear.Utility_ForceBatchCall {
	out := &pbgear.Utility_ForceBatchCall{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Calls
	out.Calls = To_Repeated_Utility_ForceBatchCall_calls(v.ValueAt(0))

	return out //from message
}

func To_Repeated_Utility_ForceBatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
	items := in.([]interface{})

	var out []*pbgear.VaraRuntimeRuntimeCall
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_VaraRuntimeRuntimeCall(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_Utility_Origins(in any) *pbgear.Utility_Origins {
	out := &pbgear.Utility_Origins{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Utility_Origins_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Utility_Origins_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_Origins_Value0StakingAdmin{
			Value0StakingAdmin: To_StakingAdmin(param),
		}
	case 1:
		return &pbgear.Utility_Origins_Value0Treasurer{
			Value0Treasurer: To_Treasurer(param),
		}
	case 2:
		return &pbgear.Utility_Origins_Value0FellowshipAdmin{
			Value0FellowshipAdmin: To_FellowshipAdmin(param),
		}
	case 3:
		return &pbgear.Utility_Origins_Value0GeneralAdmin{
			Value0GeneralAdmin: To_GeneralAdmin(param),
		}
	case 4:
		return &pbgear.Utility_Origins_Value0ReferendumCanceller{
			Value0ReferendumCanceller: To_ReferendumCanceller(param),
		}
	case 5:
		return &pbgear.Utility_Origins_Value0ReferendumKiller{
			Value0ReferendumKiller: To_ReferendumKiller(param),
		}
	case 6:
		return &pbgear.Utility_Origins_Value0SmallTipper{
			Value0SmallTipper: To_SmallTipper(param),
		}
	case 7:
		return &pbgear.Utility_Origins_Value0BigTipper{
			Value0BigTipper: To_BigTipper(param),
		}
	case 8:
		return &pbgear.Utility_Origins_Value0SmallSpender{
			Value0SmallSpender: To_SmallSpender(param),
		}
	case 9:
		return &pbgear.Utility_Origins_Value0MediumSpender{
			Value0MediumSpender: To_MediumSpender(param),
		}
	case 10:
		return &pbgear.Utility_Origins_Value0BigSpender{
			Value0BigSpender: To_BigSpender(param),
		}
	case 11:
		return &pbgear.Utility_Origins_Value0WhitelistedCaller{
			Value0WhitelistedCaller: To_WhitelistedCaller(param),
		}
	case 12:
		return &pbgear.Utility_Origins_Value0FellowshipInitiates{
			Value0FellowshipInitiates: To_FellowshipInitiates(param),
		}
	case 13:
		return &pbgear.Utility_Origins_Value0Fellows{
			Value0Fellows: To_Fellows(param),
		}
	case 14:
		return &pbgear.Utility_Origins_Value0FellowshipExperts{
			Value0FellowshipExperts: To_FellowshipExperts(param),
		}
	case 15:
		return &pbgear.Utility_Origins_Value0FellowshipMasters{
			Value0FellowshipMasters: To_FellowshipMasters(param),
		}
	case 16:
		return &pbgear.Utility_Origins_Value0Fellowship1Dan{
			Value0Fellowship1Dan: To_Fellowship1Dan(param),
		}
	case 17:
		return &pbgear.Utility_Origins_Value0Fellowship2Dan{
			Value0Fellowship2Dan: To_Fellowship2Dan(param),
		}
	case 18:
		return &pbgear.Utility_Origins_Value0Fellowship3Dan{
			Value0Fellowship3Dan: To_Fellowship3Dan(param),
		}
	case 19:
		return &pbgear.Utility_Origins_Value0Fellowship4Dan{
			Value0Fellowship4Dan: To_Fellowship4Dan(param),
		}
	case 20:
		return &pbgear.Utility_Origins_Value0Fellowship5Dan{
			Value0Fellowship5Dan: To_Fellowship5Dan(param),
		}
	case 21:
		return &pbgear.Utility_Origins_Value0Fellowship6Dan{
			Value0Fellowship6Dan: To_Fellowship6Dan(param),
		}
	case 22:
		return &pbgear.Utility_Origins_Value0Fellowship7Dan{
			Value0Fellowship7Dan: To_Fellowship7Dan(param),
		}
	case 23:
		return &pbgear.Utility_Origins_Value0Fellowship8Dan{
			Value0Fellowship8Dan: To_Fellowship8Dan(param),
		}
	case 24:
		return &pbgear.Utility_Origins_Value0Fellowship9Dan{
			Value0Fellowship9Dan: To_Fellowship9Dan(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_System(in any) *pbgear.Utility_System {
	out := &pbgear.Utility_System{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Utility_System_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Utility_System_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_System_Value0Root{
			Value0Root: To_Root(param),
		}
	case 1:
		return &pbgear.Utility_System_Value0Signed{
			Value0Signed: To_Signed(param),
		}
	case 2:
		return &pbgear.Utility_System_Value0None{
			Value0None: To_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Utility_Void(in any) *pbgear.Utility_Void {
	out := &pbgear.Utility_Void{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Utility_Void_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Utility_Void_value0(in any) any {
	return nil
}
func To_Utility_WithWeightCall(in any) *pbgear.Utility_WithWeightCall {
	out := &pbgear.Utility_WithWeightCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_Utility_WithWeightCall_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))
	// field Weight To_SpWeightsWeightV2Weight(w)
	out.Weight = To_SpWeightsWeightV2Weight(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Utility_WithWeightCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Utility_WithWeightCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Utility_WithWeightCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Utility_WithWeightCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Utility_WithWeightCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Utility_WithWeightCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Utility_WithWeightCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Utility_WithWeightCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Utility_WithWeightCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Utility_WithWeightCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Utility_WithWeightCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Utility_WithWeightCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Utility_WithWeightCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Utility_WithWeightCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Utility_WithWeightCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Utility_WithWeightCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Utility_WithWeightCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Utility_WithWeightCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Utility_WithWeightCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Utility_WithWeightCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Utility_WithWeightCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Utility_WithWeightCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Utility_WithWeightCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Utility_WithWeightCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Utility_WithWeightCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Utility_WithWeightCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Utility_WithWeightCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Utility_WithWeightCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Utility_WithWeightCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Utility_WithWeightCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Value0(in any) *pbgear.Value0 {
	out := &pbgear.Value0{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Value0
	v0 := To_OneOf_Value0_value0(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Value0_value0(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Value0_Value0Root{
			Value0Root: To_Root(param),
		}
	case 1:
		return &pbgear.Value0_Value0Signed{
			Value0Signed: To_Signed(param),
		}
	case 2:
		return &pbgear.Value0_Value0None{
			Value0None: To_None(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_VaraRuntimeNposSolution16(in any) *pbgear.VaraRuntimeNposSolution16 {
	out := &pbgear.VaraRuntimeNposSolution16{}
	v := in.(registry.Valuable)
	_ = v

	// repeated field Votes1
	out.Votes1 = To_Repeated_VaraRuntimeNposSolution16_votes1(v.ValueAt(0))
	// repeated field Votes2
	out.Votes2 = To_Repeated_VaraRuntimeNposSolution16_votes2(v.ValueAt(1))
	// repeated field Votes3
	out.Votes3 = To_Repeated_VaraRuntimeNposSolution16_votes3(v.ValueAt(2))
	// repeated field Votes4
	out.Votes4 = To_Repeated_VaraRuntimeNposSolution16_votes4(v.ValueAt(3))
	// repeated field Votes5
	out.Votes5 = To_Repeated_VaraRuntimeNposSolution16_votes5(v.ValueAt(4))
	// repeated field Votes6
	out.Votes6 = To_Repeated_VaraRuntimeNposSolution16_votes6(v.ValueAt(5))
	// repeated field Votes7
	out.Votes7 = To_Repeated_VaraRuntimeNposSolution16_votes7(v.ValueAt(6))
	// repeated field Votes8
	out.Votes8 = To_Repeated_VaraRuntimeNposSolution16_votes8(v.ValueAt(7))
	// repeated field Votes9
	out.Votes9 = To_Repeated_VaraRuntimeNposSolution16_votes9(v.ValueAt(8))
	// repeated field Votes10
	out.Votes10 = To_Repeated_VaraRuntimeNposSolution16_votes10(v.ValueAt(9))
	// repeated field Votes11
	out.Votes11 = To_Repeated_VaraRuntimeNposSolution16_votes11(v.ValueAt(10))
	// repeated field Votes12
	out.Votes12 = To_Repeated_VaraRuntimeNposSolution16_votes12(v.ValueAt(11))
	// repeated field Votes13
	out.Votes13 = To_Repeated_VaraRuntimeNposSolution16_votes13(v.ValueAt(12))
	// repeated field Votes14
	out.Votes14 = To_Repeated_VaraRuntimeNposSolution16_votes14(v.ValueAt(13))
	// repeated field Votes15
	out.Votes15 = To_Repeated_VaraRuntimeNposSolution16_votes15(v.ValueAt(14))
	// repeated field Votes16
	out.Votes16 = To_Repeated_VaraRuntimeNposSolution16_votes16(v.ValueAt(15))

	return out //from message
}

func To_Repeated_VaraRuntimeNposSolution16_votes1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32Uint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes2(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes3(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes4(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes5(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes6(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes7(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes8(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes9(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes10(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes11(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes12(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes13(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes14(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes15(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes16(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
	items := in.([]interface{})

	var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32
	for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok { //ugly ack the will bite later
			w = &Wrap{Value: i}
		}
		o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(w)
		out = append(out, o)
	}
	return nil //funcForRepeatedField
}
func To_VaraRuntimeRuntimeCall(in any) *pbgear.VaraRuntimeRuntimeCall {
	out := &pbgear.VaraRuntimeRuntimeCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Calls
	v0 := To_OneOf_VaraRuntimeRuntimeCall_calls(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Calls").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_VaraRuntimeRuntimeCall_calls(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.VaraRuntimeRuntimeCall_CallsSystem{
			CallsSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.VaraRuntimeRuntimeCall_CallsTimestamp{
			CallsTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.VaraRuntimeRuntimeCall_CallsBabe{
			CallsBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.VaraRuntimeRuntimeCall_CallsGrandpa{
			CallsGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.VaraRuntimeRuntimeCall_CallsBalances{
			CallsBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.VaraRuntimeRuntimeCall_CallsVesting{
			CallsVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.VaraRuntimeRuntimeCall_CallsBagsList{
			CallsBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.VaraRuntimeRuntimeCall_CallsImOnline{
			CallsImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.VaraRuntimeRuntimeCall_CallsStaking{
			CallsStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.VaraRuntimeRuntimeCall_CallsSession{
			CallsSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.VaraRuntimeRuntimeCall_CallsTreasury{
			CallsTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.VaraRuntimeRuntimeCall_CallsUtility{
			CallsUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.VaraRuntimeRuntimeCall_CallsConvictionVoting{
			CallsConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.VaraRuntimeRuntimeCall_CallsReferenda{
			CallsReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.VaraRuntimeRuntimeCall_CallsFellowshipCollective{
			CallsFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.VaraRuntimeRuntimeCall_CallsFellowshipReferenda{
			CallsFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.VaraRuntimeRuntimeCall_CallsWhitelist{
			CallsWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.VaraRuntimeRuntimeCall_CallsScheduler{
			CallsScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.VaraRuntimeRuntimeCall_CallsPreimage{
			CallsPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.VaraRuntimeRuntimeCall_CallsIdentity{
			CallsIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.VaraRuntimeRuntimeCall_CallsProxy{
			CallsProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.VaraRuntimeRuntimeCall_CallsMultisig{
			CallsMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.VaraRuntimeRuntimeCall_CallsElectionProviderMultiPhase{
			CallsElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.VaraRuntimeRuntimeCall_CallsBounties{
			CallsBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.VaraRuntimeRuntimeCall_CallsChildBounties{
			CallsChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.VaraRuntimeRuntimeCall_CallsNominationPools{
			CallsNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.VaraRuntimeRuntimeCall_CallsGear{
			CallsGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.VaraRuntimeRuntimeCall_CallsStakingRewards{
			CallsStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.VaraRuntimeRuntimeCall_CallsGearVoucher{
			CallsGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_VaraRuntimeSessionKeys(in any) *pbgear.VaraRuntimeSessionKeys {
	out := &pbgear.VaraRuntimeSessionKeys{}
	v := in.(registry.Valuable)
	_ = v

	// field Babe To_SpConsensusBabeAppPublic(w)
	out.Babe = To_SpConsensusBabeAppPublic(v.ValueAt(0))
	// field Grandpa To_SpConsensusGrandpaAppPublic(w)
	out.Grandpa = To_SpConsensusGrandpaAppPublic(v.ValueAt(1))
	// field ImOnline To_ImOnline_PalletImOnlineSr25519AppSr25519Public(w)
	out.ImOnline = To_ImOnline_PalletImOnlineSr25519AppSr25519Public(v.ValueAt(2))
	// field AuthorityDiscovery To_SpAuthorityDiscoveryAppPublic(w)
	out.AuthorityDiscovery = To_SpAuthorityDiscoveryAppPublic(v.ValueAt(3))

	return out //from message
}

func To_VestingPallet(in any) *pbgear.VestingPallet {
	out := &pbgear.VestingPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_VestingPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_VestingPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.VestingPallet_CallVestCall{
			CallVestCall: To_Vesting_VestCall(param),
		}
	case 1:
		return &pbgear.VestingPallet_CallVestOtherCall{
			CallVestOtherCall: To_Vesting_VestOtherCall(param),
		}
	case 2:
		return &pbgear.VestingPallet_CallVestedTransferCall{
			CallVestedTransferCall: To_Vesting_VestedTransferCall(param),
		}
	case 3:
		return &pbgear.VestingPallet_CallForceVestedTransferCall{
			CallForceVestedTransferCall: To_Vesting_ForceVestedTransferCall(param),
		}
	case 4:
		return &pbgear.VestingPallet_CallMergeSchedulesCall{
			CallMergeSchedulesCall: To_Vesting_MergeSchedulesCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Vesting_Address20(in any) *pbgear.Vesting_Address20 {
	out := &pbgear.Vesting_Address20{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Vesting_Address32(in any) *pbgear.Vesting_Address32 {
	out := &pbgear.Vesting_Address32{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Vesting_ForceVestedTransferCall(in any) *pbgear.Vesting_ForceVestedTransferCall {
	out := &pbgear.Vesting_ForceVestedTransferCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Source
	v0 := To_OneOf_Vesting_ForceVestedTransferCall_source(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))
	// oneOf field Target
	v1 := To_OneOf_Vesting_ForceVestedTransferCall_target(v.ValueAt(1))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))
	// field Schedule To_Vesting_PalletVestingVestingInfoVestingInfo(w)
	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(v.ValueAt(2))

	return out //from message
}

func To_OneOf_Vesting_ForceVestedTransferCall_source(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_ForceVestedTransferCall_SourceId{
			SourceId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_ForceVestedTransferCall_SourceIndex{
			SourceIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_ForceVestedTransferCall_SourceRaw{
			SourceRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_ForceVestedTransferCall_SourceAddress32{
			SourceAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_ForceVestedTransferCall_SourceAddress20{
			SourceAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_OneOf_Vesting_ForceVestedTransferCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_ForceVestedTransferCall_TargetId{
			TargetId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_ForceVestedTransferCall_TargetIndex{
			TargetIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_ForceVestedTransferCall_TargetRaw{
			TargetRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_ForceVestedTransferCall_TargetAddress32{
			TargetAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_ForceVestedTransferCall_TargetAddress20{
			TargetAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_Vesting_Id(in any) *pbgear.Vesting_Id {
	out := &pbgear.Vesting_Id{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_SpCoreCryptoAccountId32(w)
	out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

	return out //from message
}

func To_Vesting_Index(in any) *pbgear.Vesting_Index {
	out := &pbgear.Vesting_Index{}
	v := in.(registry.Valuable)
	_ = v

	// field Value0 To_Vesting_TupleNull(w)
	out.Value0 = To_Vesting_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Vesting_MergeSchedulesCall(in any) *pbgear.Vesting_MergeSchedulesCall {
	out := &pbgear.Vesting_MergeSchedulesCall{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Schedule1Index
	out.Schedule1Index = To_uint32(v.ValueAt(0))
	// primitive field Schedule2Index
	out.Schedule2Index = To_uint32(v.ValueAt(1))

	return out //from message
}

func To_Vesting_PalletVestingVestingInfoVestingInfo(in any) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Locked
	out.Locked = To_string(v.ValueAt(0))
	// primitive field PerBlock
	out.PerBlock = To_string(v.ValueAt(1))
	// primitive field StartingBlock
	out.StartingBlock = To_uint32(v.ValueAt(2))

	return out //from message
}

func To_Vesting_Raw(in any) *pbgear.Vesting_Raw {
	out := &pbgear.Vesting_Raw{}
	v := in.(registry.Valuable)
	_ = v

	// primitive field Value0
	out.Value0 = To_bytes(v.ValueAt(0))

	return out //from message
}

func To_Vesting_Source(in any) *pbgear.Vesting_Source {
	out := &pbgear.Vesting_Source{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Source
	v0 := To_OneOf_Vesting_Source_source(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Vesting_Source_source(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_Source_SourceId{
			SourceId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_Source_SourceIndex{
			SourceIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_Source_SourceRaw{
			SourceRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_Source_SourceAddress32{
			SourceAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_Source_SourceAddress20{
			SourceAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Vesting_Target(in any) *pbgear.Vesting_Target {
	out := &pbgear.Vesting_Target{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_Vesting_Target_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Vesting_Target_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_Target_TargetId{
			TargetId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_Target_TargetIndex{
			TargetIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_Target_TargetRaw{
			TargetRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_Target_TargetAddress32{
			TargetAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_Target_TargetAddress20{
			TargetAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Vesting_TupleNull(in any) *pbgear.Vesting_TupleNull {
	out := &pbgear.Vesting_TupleNull{}
	v := in.(registry.Valuable)
	_ = v

	// field Value To_TupleNull(w)
	out.Value = To_TupleNull(v.ValueAt(0))

	return out //from message
}

func To_Vesting_VestCall(in any) *pbgear.Vesting_VestCall {
	out := &pbgear.Vesting_VestCall{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

func To_Vesting_VestOtherCall(in any) *pbgear.Vesting_VestOtherCall {
	out := &pbgear.Vesting_VestOtherCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_Vesting_VestOtherCall_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Vesting_VestOtherCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_VestOtherCall_TargetId{
			TargetId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_VestOtherCall_TargetIndex{
			TargetIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_VestOtherCall_TargetRaw{
			TargetRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_VestOtherCall_TargetAddress32{
			TargetAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_VestOtherCall_TargetAddress20{
			TargetAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Vesting_VestedTransferCall(in any) *pbgear.Vesting_VestedTransferCall {
	out := &pbgear.Vesting_VestedTransferCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Target
	v0 := To_OneOf_Vesting_VestedTransferCall_target(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))
	// field Schedule To_Vesting_PalletVestingVestingInfoVestingInfo(w)
	out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(v.ValueAt(1))

	return out //from message
}

func To_OneOf_Vesting_VestedTransferCall_target(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Vesting_VestedTransferCall_TargetId{
			TargetId: To_Vesting_Id(param),
		}
	case 1:
		return &pbgear.Vesting_VestedTransferCall_TargetIndex{
			TargetIndex: To_Vesting_Index(param),
		}
	case 2:
		return &pbgear.Vesting_VestedTransferCall_TargetRaw{
			TargetRaw: To_Vesting_Raw(param),
		}
	case 3:
		return &pbgear.Vesting_VestedTransferCall_TargetAddress32{
			TargetAddress32: To_Vesting_Address32(param),
		}
	case 4:
		return &pbgear.Vesting_VestedTransferCall_TargetAddress20{
			TargetAddress20: To_Vesting_Address20(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}

func To_WhitelistPallet(in any) *pbgear.WhitelistPallet {
	out := &pbgear.WhitelistPallet{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_WhitelistPallet_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_WhitelistPallet_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.WhitelistPallet_CallWhitelistCallCall{
			CallWhitelistCallCall: To_Whitelist_WhitelistCallCall(param),
		}
	case 1:
		return &pbgear.WhitelistPallet_CallRemoveWhitelistedCallCall{
			CallRemoveWhitelistedCallCall: To_Whitelist_RemoveWhitelistedCallCall(param),
		}
	case 2:
		return &pbgear.WhitelistPallet_CallDispatchWhitelistedCallCall{
			CallDispatchWhitelistedCallCall: To_Whitelist_DispatchWhitelistedCallCall(param),
		}
	case 3:
		return &pbgear.WhitelistPallet_CallDispatchWhitelistedCallWithPreimageCall{
			CallDispatchWhitelistedCallWithPreimageCall: To_Whitelist_DispatchWhitelistedCallWithPreimageCall(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Whitelist_DispatchWhitelistedCallCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
	v := in.(registry.Valuable)
	_ = v

	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))
	// primitive field CallEncodedLen
	out.CallEncodedLen = To_uint32(v.ValueAt(1))
	// field CallWeightWitness To_SpWeightsWeightV2Weight(w)
	out.CallWeightWitness = To_SpWeightsWeightV2Weight(v.ValueAt(2))

	return out //from message
}

func To_Whitelist_DispatchWhitelistedCallWithPreimageCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
	out := &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{}
	v := in.(registry.Valuable)
	_ = v

	// oneOf field Call
	v0 := To_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call(v.ValueAt(0))
	reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

	return out //from message
}

func To_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call(in any) any {
	variantIn := in.(*registry.VariantWTF)
	param := variantIn.Value
	switch variantIn.VariantByte {
	case 0:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallSystem{
			CallSystem: To_SystemPallet(param),
		}
	case 1:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallTimestamp{
			CallTimestamp: To_TimestampPallet(param),
		}
	case 3:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBabe{
			CallBabe: To_BabePallet(param),
		}
	case 4:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGrandpa{
			CallGrandpa: To_GrandpaPallet(param),
		}
	case 5:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBalances{
			CallBalances: To_BalancesPallet(param),
		}
	case 10:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallVesting{
			CallVesting: To_VestingPallet(param),
		}
	case 11:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBagsList{
			CallBagsList: To_BagsListPallet(param),
		}
	case 12:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallImOnline{
			CallImOnline: To_ImOnlinePallet(param),
		}
	case 13:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallStaking{
			CallStaking: To_StakingPallet(param),
		}
	case 7:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallSession{
			CallSession: To_SessionPallet(param),
		}
	case 14:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallTreasury{
			CallTreasury: To_TreasuryPallet(param),
		}
	case 8:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallUtility{
			CallUtility: To_UtilityPallet(param),
		}
	case 16:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallConvictionVoting{
			CallConvictionVoting: To_ConvictionVotingPallet(param),
		}
	case 17:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallReferenda{
			CallReferenda: To_ReferendaPallet(param),
		}
	case 18:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallFellowshipCollective{
			CallFellowshipCollective: To_FellowshipCollectivePallet(param),
		}
	case 19:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallFellowshipReferenda{
			CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
		}
	case 21:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallWhitelist{
			CallWhitelist: To_WhitelistPallet(param),
		}
	case 22:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallScheduler{
			CallScheduler: To_SchedulerPallet(param),
		}
	case 23:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallPreimage{
			CallPreimage: To_PreimagePallet(param),
		}
	case 24:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallIdentity{
			CallIdentity: To_IdentityPallet(param),
		}
	case 25:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallProxy{
			CallProxy: To_ProxyPallet(param),
		}
	case 26:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallMultisig{
			CallMultisig: To_MultisigPallet(param),
		}
	case 27:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallElectionProviderMultiPhase{
			CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
		}
	case 29:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBounties{
			CallBounties: To_BountiesPallet(param),
		}
	case 30:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallChildBounties{
			CallChildBounties: To_ChildBountiesPallet(param),
		}
	case 31:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallNominationPools{
			CallNominationPools: To_NominationPoolsPallet(param),
		}
	case 104:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGear{
			CallGear: To_GearPallet(param),
		}
	case 106:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallStakingRewards{
			CallStakingRewards: To_StakingRewardsPallet(param),
		}
	case 107:
		return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGearVoucher{
			CallGearVoucher: To_GearVoucherPallet(param),
		}
	default:
		panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
	}
}
func To_Whitelist_RemoveWhitelistedCallCall(in any) *pbgear.Whitelist_RemoveWhitelistedCallCall {
	out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
	v := in.(registry.Valuable)
	_ = v

	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_Whitelist_WhitelistCallCall(in any) *pbgear.Whitelist_WhitelistCallCall {
	out := &pbgear.Whitelist_WhitelistCallCall{}
	v := in.(registry.Valuable)
	_ = v

	// field CallHash To_PrimitiveTypesH256(w)
	out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

	return out //from message
}

func To_WhitelistedCaller(in any) *pbgear.WhitelistedCaller {
	out := &pbgear.WhitelistedCaller{}
	v := in.(registry.Valuable)
	_ = v

	return out //from message
}

type Wrap struct {
	Value registry.Valuable
}

func (w Wrap) ValueAt(index int) any {
	return w.Value
}

func To_bytes(a any) []byte {
	var items []any
	switch in := a.(type) {
	case registry.DecodedFields:
		fields := []*registry.DecodedField(in)
		items = fields[0].Value.([]any)
	case []any:
		items = in
	default:
		panic(fmt.Sprintf("Unsupported type %T", in))
	}
	var out []byte
	for _, i := range items {
		u := uint8(i.(types.U8))
		out = append(out, byte(u))
	}
	return out
}

func To_string(i any) string {
	switch v := i.(type) {
	case types.UCompact:
		return fmt.Sprintf("%d", v.Int64())
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
	FuncMap["To_AuthorityDiscoveryPallet"] = reflect.ValueOf(To_AuthorityDiscoveryPallet)
	FuncMap["To_AuthorshipPallet"] = reflect.ValueOf(To_AuthorshipPallet)
	FuncMap["To_BTreeSet"] = reflect.ValueOf(To_BTreeSet)
	FuncMap["To_BabePallet"] = reflect.ValueOf(To_BabePallet)
	FuncMap["To_Babe_AllowedSlots"] = reflect.ValueOf(To_Babe_AllowedSlots)
	FuncMap["To_Babe_BabeTrieNodesList"] = reflect.ValueOf(To_Babe_BabeTrieNodesList)
	FuncMap["To_Babe_Config"] = reflect.ValueOf(To_Babe_Config)
	FuncMap["To_Babe_Consensus"] = reflect.ValueOf(To_Babe_Consensus)
	FuncMap["To_Babe_Logs"] = reflect.ValueOf(To_Babe_Logs)
	FuncMap["To_Babe_Other"] = reflect.ValueOf(To_Babe_Other)
	FuncMap["To_Babe_PlanConfigChangeCall"] = reflect.ValueOf(To_Babe_PlanConfigChangeCall)
	FuncMap["To_Babe_PreRuntime"] = reflect.ValueOf(To_Babe_PreRuntime)
	FuncMap["To_Babe_PrimaryAndSecondaryPlainSlots"] = reflect.ValueOf(To_Babe_PrimaryAndSecondaryPlainSlots)
	FuncMap["To_Babe_PrimaryAndSecondaryVrfSlots"] = reflect.ValueOf(To_Babe_PrimaryAndSecondaryVrfSlots)
	FuncMap["To_Babe_PrimarySlots"] = reflect.ValueOf(To_Babe_PrimarySlots)
	FuncMap["To_Babe_ReportEquivocationCall"] = reflect.ValueOf(To_Babe_ReportEquivocationCall)
	FuncMap["To_Babe_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Babe_ReportEquivocationUnsignedCall)
	FuncMap["To_Babe_RuntimeEnvironmentUpdated"] = reflect.ValueOf(To_Babe_RuntimeEnvironmentUpdated)
	FuncMap["To_Babe_Seal"] = reflect.ValueOf(To_Babe_Seal)
	FuncMap["To_Babe_SpConsensusBabeAppPublic"] = reflect.ValueOf(To_Babe_SpConsensusBabeAppPublic)
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
	FuncMap["To_BigSpender"] = reflect.ValueOf(To_BigSpender)
	FuncMap["To_BigTipper"] = reflect.ValueOf(To_BigTipper)
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
	FuncMap["To_Fellows"] = reflect.ValueOf(To_Fellows)
	FuncMap["To_Fellowship1Dan"] = reflect.ValueOf(To_Fellowship1Dan)
	FuncMap["To_Fellowship2Dan"] = reflect.ValueOf(To_Fellowship2Dan)
	FuncMap["To_Fellowship3Dan"] = reflect.ValueOf(To_Fellowship3Dan)
	FuncMap["To_Fellowship4Dan"] = reflect.ValueOf(To_Fellowship4Dan)
	FuncMap["To_Fellowship5Dan"] = reflect.ValueOf(To_Fellowship5Dan)
	FuncMap["To_Fellowship6Dan"] = reflect.ValueOf(To_Fellowship6Dan)
	FuncMap["To_Fellowship7Dan"] = reflect.ValueOf(To_Fellowship7Dan)
	FuncMap["To_Fellowship8Dan"] = reflect.ValueOf(To_Fellowship8Dan)
	FuncMap["To_Fellowship9Dan"] = reflect.ValueOf(To_Fellowship9Dan)
	FuncMap["To_FellowshipAdmin"] = reflect.ValueOf(To_FellowshipAdmin)
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
	FuncMap["To_FellowshipExperts"] = reflect.ValueOf(To_FellowshipExperts)
	FuncMap["To_FellowshipInitiates"] = reflect.ValueOf(To_FellowshipInitiates)
	FuncMap["To_FellowshipMasters"] = reflect.ValueOf(To_FellowshipMasters)
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
	FuncMap["To_GeneralAdmin"] = reflect.ValueOf(To_GeneralAdmin)
	FuncMap["To_GprimitivesActorId"] = reflect.ValueOf(To_GprimitivesActorId)
	FuncMap["To_GprimitivesCodeId"] = reflect.ValueOf(To_GprimitivesCodeId)
	FuncMap["To_GprimitivesMessageId"] = reflect.ValueOf(To_GprimitivesMessageId)
	FuncMap["To_GrandpaPallet"] = reflect.ValueOf(To_GrandpaPallet)
	FuncMap["To_Grandpa_Equivocation"] = reflect.ValueOf(To_Grandpa_Equivocation)
	FuncMap["To_Grandpa_FinalityGrandpaEquivocation"] = reflect.ValueOf(To_Grandpa_FinalityGrandpaEquivocation)
	FuncMap["To_Grandpa_FinalityGrandpaPrecommit"] = reflect.ValueOf(To_Grandpa_FinalityGrandpaPrecommit)
	FuncMap["To_Grandpa_FinalityGrandpaPrevote"] = reflect.ValueOf(To_Grandpa_FinalityGrandpaPrevote)
	FuncMap["To_Grandpa_GrandpaTrieNodesList"] = reflect.ValueOf(To_Grandpa_GrandpaTrieNodesList)
	FuncMap["To_Grandpa_NoteStalledCall"] = reflect.ValueOf(To_Grandpa_NoteStalledCall)
	FuncMap["To_Grandpa_Precommit"] = reflect.ValueOf(To_Grandpa_Precommit)
	FuncMap["To_Grandpa_Prevote"] = reflect.ValueOf(To_Grandpa_Prevote)
	FuncMap["To_Grandpa_ReportEquivocationCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationCall)
	FuncMap["To_Grandpa_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationUnsignedCall)
	FuncMap["To_Grandpa_SpConsensusGrandpaAppPublic"] = reflect.ValueOf(To_Grandpa_SpConsensusGrandpaAppPublic)
	FuncMap["To_Grandpa_SpConsensusGrandpaAppSignature"] = reflect.ValueOf(To_Grandpa_SpConsensusGrandpaAppSignature)
	FuncMap["To_Grandpa_SpConsensusGrandpaEquivocationProof"] = reflect.ValueOf(To_Grandpa_SpConsensusGrandpaEquivocationProof)
	FuncMap["To_Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature"] = reflect.ValueOf(To_Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature)
	FuncMap["To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature"] = reflect.ValueOf(To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature)
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
	FuncMap["To_MediumSpender"] = reflect.ValueOf(To_MediumSpender)
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
	FuncMap["To_ReferendumCanceller"] = reflect.ValueOf(To_ReferendumCanceller)
	FuncMap["To_ReferendumKiller"] = reflect.ValueOf(To_ReferendumKiller)
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
	FuncMap["To_SmallSpender"] = reflect.ValueOf(To_SmallSpender)
	FuncMap["To_SmallTipper"] = reflect.ValueOf(To_SmallTipper)
	FuncMap["To_SpArithmeticPerThingsPerU16"] = reflect.ValueOf(To_SpArithmeticPerThingsPerU16)
	FuncMap["To_SpArithmeticPerThingsPerbill"] = reflect.ValueOf(To_SpArithmeticPerThingsPerbill)
	FuncMap["To_SpArithmeticPerThingsPercent"] = reflect.ValueOf(To_SpArithmeticPerThingsPercent)
	FuncMap["To_SpAuthorityDiscoveryAppPublic"] = reflect.ValueOf(To_SpAuthorityDiscoveryAppPublic)
	FuncMap["To_SpConsensusBabeAppPublic"] = reflect.ValueOf(To_SpConsensusBabeAppPublic)
	FuncMap["To_SpConsensusGrandpaAppPublic"] = reflect.ValueOf(To_SpConsensusGrandpaAppPublic)
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
	FuncMap["To_StakingAdmin"] = reflect.ValueOf(To_StakingAdmin)
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
	FuncMap["To_Treasurer"] = reflect.ValueOf(To_Treasurer)
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
	FuncMap["To_WhitelistedCaller"] = reflect.ValueOf(To_WhitelistedCaller)
}
