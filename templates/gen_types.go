package gen_types

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func to_Identity_SetFieldsCall0000000(fields []*registry.DecodedField) *pbgear.Identity_SetFieldsCall {
	return &pbgear.Identity_SetFieldsCall{
		Index:  to_Identity_CompactUint32(fields[0].Value),
		Fields: to_Identity_PalletIdentityTypesBitFlags(fields[1].Value),
	}
}

func to_Balances_UpgradeAccountsCall0000000(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
	return &pbgear.Balances_UpgradeAccountsCall{
		Who: to_repeated_Balances_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Utility_BatchAllCall0000000(fields []*registry.DecodedField) *pbgear.Utility_BatchAllCall {
	return &pbgear.Utility_BatchAllCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value),
	}
}

func to_Referenda_Proposal0000000(fields []*registry.DecodedField) *pbgear.Referenda_Proposal {
	return &pbgear.Referenda_Proposal{
		Value: to_oneof_Referenda_Value(fields[0].Value),
	}
}

func to_oneof_Referenda_Value(in *registry.DecodedField) *pbgear.Referenda_Value {
	return nil
}

func to_Identity_Raw160000000(fields []*registry.DecodedField) *pbgear.Identity_Raw16 {
	return &pbgear.Identity_Raw16{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw300000000(fields []*registry.DecodedField) *pbgear.Identity_Raw30 {
	return &pbgear.Identity_Raw30{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_StakingRewards_To0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_To {
	return &pbgear.StakingRewards_To{
		Value: to_oneof_StakingRewards_Value(fields[0].Value),
	}
}

func to_oneof_StakingRewards_Value(in *registry.DecodedField) *pbgear.StakingRewards_Value {
	return nil
}

func to_Identity_Raw270000000(fields []*registry.DecodedField) *pbgear.Identity_Raw27 {
	return &pbgear.Identity_Raw27{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_New0000000(fields []*registry.DecodedField) *pbgear.Identity_New {
	return &pbgear.Identity_New{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_oneof_Identity_Value(in *registry.DecodedField) *pbgear.Identity_Value {
	return nil
}

func to_Proxy_NonTransfer0000000(fields []*registry.DecodedField) *pbgear.Proxy_NonTransfer {
	return &pbgear.Proxy_NonTransfer{}
}
func to_ChildBounties_Address320000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Address32 {
	return &pbgear.ChildBounties_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_Root0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Root {
	return &pbgear.NominationPools_Root{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_oneof_NominationPools_Value(in *registry.DecodedField) *pbgear.NominationPools_Value {
	return nil
}

func to_Babe_PrimaryAndSecondaryPlainSlots0000000(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
	return &pbgear.Babe_PrimaryAndSecondaryPlainSlots{}
}
func to_Staking_Index0000000(fields []*registry.DecodedField) *pbgear.Staking_Index {
	return &pbgear.Staking_Index{
		Value_0: to_Staking_CompactTupleNull(fields[0].Value),
	}
}

func to_Gear_GprimitivesMessageId0000000(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
	return &pbgear.Gear_GprimitivesMessageId{
		ReplyToId: []uint32(fields[0].Value),
	}
}

func to_Proxy_RemoveAnnouncementCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_RemoveAnnouncementCall {
	return &pbgear.Proxy_RemoveAnnouncementCall{
		Real:     to_Proxy_Real(fields[0].Value),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value),
	}
}

func to_NominationPools_Index0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Index {
	return &pbgear.NominationPools_Index{
		Value_0: to_NominationPools_CompactTupleNull(fields[0].Value),
	}
}

func to_Grandpa_SpCoreEd25519Signature0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Signature {
	return &pbgear.Grandpa_SpCoreEd25519Signature{
		Value_1: []uint32(fields[0].Value),
	}
}

func to_Grandpa_FinalityGrandpaEquivocation0000000(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaEquivocation {
	return &pbgear.Grandpa_FinalityGrandpaEquivocation{
		RoundNumber: uint64(fields[0].Value),
		Identity:    to_Grandpa_SpConsensusGrandpaAppPublic(fields[1].Value),
		First:       to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[2].Value),
		Second:      to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(fields[3].Value),
	}
}

func to_Treasury_Beneficiary0000000(fields []*registry.DecodedField) *pbgear.Treasury_Beneficiary {
	return &pbgear.Treasury_Beneficiary{
		Value: to_oneof_Treasury_Value(fields[0].Value),
	}
}

func to_oneof_Treasury_Value(in *registry.DecodedField) *pbgear.Treasury_Value {
	return nil
}

func to_ConvictionVoting_RemoveOtherVoteCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
	return &pbgear.ConvictionVoting_RemoveOtherVoteCall{
		Target: to_ConvictionVoting_Target(fields[0].Value),
		Class:  uint32(fields[1].Value),
		Index:  uint32(fields[2].Value),
	}
}

func to_Proxy_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
	return &pbgear.Proxy_PrimitiveTypesH256{
		CallHash: []uint32(fields[0].Value),
	}
}

func to_Proxy_ForceProxyType0000000(fields []*registry.DecodedField) *pbgear.Proxy_ForceProxyType {
	return &pbgear.Proxy_ForceProxyType{
		Value: to_oneof_Proxy_Value(fields[0].Value),
	}
}

func to_oneof_Proxy_Value(in *registry.DecodedField) *pbgear.Proxy_Value {
	return nil
}

func to_Bounties_AcceptCuratorCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_AcceptCuratorCall {
	return &pbgear.Bounties_AcceptCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
	}
}

func to_ChildBounties_ProposeCuratorCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_ProposeCuratorCall {
	return &pbgear.ChildBounties_ProposeCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
		Curator:        to_ChildBounties_Curator(fields[2].Value),
		Fee:            to_ChildBounties_CompactString(fields[3].Value),
	}
}

func to_BagsList_Dislocated0000000(fields []*registry.DecodedField) *pbgear.BagsList_Dislocated {
	return &pbgear.BagsList_Dislocated{
		Value: to_oneof_BagsList_Value(fields[0].Value),
	}
}

func to_oneof_BagsList_Value(in *registry.DecodedField) *pbgear.BagsList_Value {
	return nil
}

func to_FellowshipReferenda_Value00000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Value0 {
	return &pbgear.FellowshipReferenda_Value0{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value),
	}
}

func to_oneof_FellowshipReferenda_Value(in *registry.DecodedField) *pbgear.FellowshipReferenda_Value {
	return nil
}

func to_StakingRewards_Id0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_Id {
	return &pbgear.StakingRewards_Id{
		Value_0: to_StakingRewards_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_System_KillStorageCall0000000(fields []*registry.DecodedField) *pbgear.System_KillStorageCall {
	return &pbgear.System_KillStorageCall{
		Keys: to_repeated_System_SystemKeysList(fields[0].Value),
	}
}

func to_Vesting_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
	return &pbgear.Vesting_CompactTupleNull{
		Value: to_Vesting_TupleNull(fields[0].Value),
	}
}

func to_Staking_SetStakingConfigsCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetStakingConfigsCall {
	return &pbgear.Staking_SetStakingConfigsCall{
		MinNominatorBond:  to_Staking_MinNominatorBond(fields[0].Value),
		MinValidatorBond:  to_Staking_MinValidatorBond(fields[1].Value),
		MaxNominatorCount: to_Staking_MaxNominatorCount(fields[2].Value),
		MaxValidatorCount: to_Staking_MaxValidatorCount(fields[3].Value),
		ChillThreshold:    to_Staking_ChillThreshold(fields[4].Value),
		MinCommission:     to_Staking_MinCommission(fields[5].Value),
	}
}

func to_Referenda_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Referenda_SpCoreCryptoAccountId32 {
	return &pbgear.Referenda_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw320000000(fields []*registry.DecodedField) *pbgear.Identity_Raw32 {
	return &pbgear.Identity_Raw32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ChildBounties_Id0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Id {
	return &pbgear.ChildBounties_Id{
		Value_0: to_ChildBounties_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_NominationPools_GlobalMaxCommission0000000(fields []*registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
	return &pbgear.NominationPools_GlobalMaxCommission{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_StakingRewards_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.StakingRewards_SpCoreCryptoAccountId32 {
	return &pbgear.StakingRewards_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Timestamp_CompactUint640000000(fields []*registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
	return &pbgear.Timestamp_CompactUint64{
		Value: uint64(fields[0].Value),
	}
}

func to_ConvictionVoting_RemoveVoteCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveVoteCall {
	return &pbgear.ConvictionVoting_RemoveVoteCall{
		Class: uint32(fields[0].Value),
		Index: uint32(fields[1].Value),
	}
}

func to_Referenda_Value00000000(fields []*registry.DecodedField) *pbgear.Referenda_Value0 {
	return &pbgear.Referenda_Value0{
		Value: to_oneof_Referenda_Value(fields[0].Value),
	}
}

func to_Identity_Raw100000000(fields []*registry.DecodedField) *pbgear.Identity_Raw10 {
	return &pbgear.Identity_Raw10{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
	return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{
		Value_0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value),
		Value_1: string(fields[1].Value),
	}
}

func to_NominationPools_MaxMembersPerPool0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
	return &pbgear.NominationPools_MaxMembersPerPool{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_GearVoucher_DeclineVoucher0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineVoucher {
	return &pbgear.GearVoucher_DeclineVoucher{}
}
func to_Babe_Other0000000(fields []*registry.DecodedField) *pbgear.Babe_Other {
	return &pbgear.Babe_Other{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Balances_Source0000000(fields []*registry.DecodedField) *pbgear.Balances_Source {
	return &pbgear.Balances_Source{
		Value: to_oneof_Balances_Value(fields[0].Value),
	}
}

func to_oneof_Balances_Value(in *registry.DecodedField) *pbgear.Balances_Value {
	return nil
}

func to_Staking_Account0000000(fields []*registry.DecodedField) *pbgear.Staking_Account {
	return &pbgear.Staking_Account{
		Value_0: to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Identity_Raw80000000(fields []*registry.DecodedField) *pbgear.Identity_Raw8 {
	return &pbgear.Identity_Raw8{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_Address320000000(fields []*registry.DecodedField) *pbgear.Proxy_Address32 {
	return &pbgear.Proxy_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_PalletStakingValidatorPrefs0000000(fields []*registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
	return &pbgear.Staking_PalletStakingValidatorPrefs{
		Commission: to_Staking_CompactSpArithmeticPerThingsPerbill(fields[0].Value),
		Blocked:    bool(fields[1].Value),
	}
}

func to_FellowshipReferenda_Legacy0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Legacy {
	return &pbgear.FellowshipReferenda_Legacy{
		Hash: to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Identity_SetFeeCall0000000(fields []*registry.DecodedField) *pbgear.Identity_SetFeeCall {
	return &pbgear.Identity_SetFeeCall{
		Index: to_Identity_CompactUint32(fields[0].Value),
		Fee:   to_Identity_CompactString(fields[1].Value),
	}
}

func to_ChildBounties_Raw0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Raw {
	return &pbgear.ChildBounties_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_StakingRewards_Index0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_Index {
	return &pbgear.StakingRewards_Index{
		Value_0: to_StakingRewards_CompactTupleNull(fields[0].Value),
	}
}

func to_NominationPools_JoinCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_JoinCall {
	return &pbgear.NominationPools_JoinCall{
		Amount: to_NominationPools_CompactString(fields[0].Value),
		PoolId: uint32(fields[1].Value),
	}
}

func to_NominationPools_Permission0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Permission {
	return &pbgear.NominationPools_Permission{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_ConvictionVoting_UndelegateCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UndelegateCall {
	return &pbgear.ConvictionVoting_UndelegateCall{
		Class: uint32(fields[0].Value),
	}
}

func to_Whitelist_DispatchWhitelistedCallWithPreimageCall0000000(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
	return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{
		Call: to_oneof_Whitelist_Call(fields[0].Value),
	}
}

func to_oneof_Whitelist_Call(in *registry.DecodedField) *pbgear.Whitelist_Call {
	return nil
}

func to_Identity_AddRegistrarCall0000000(fields []*registry.DecodedField) *pbgear.Identity_AddRegistrarCall {
	return &pbgear.Identity_AddRegistrarCall{
		Account: to_Identity_Account(fields[0].Value),
	}
}

func to_Identity_ClearIdentityCall0000000(fields []*registry.DecodedField) *pbgear.Identity_ClearIdentityCall {
	return &pbgear.Identity_ClearIdentityCall{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_Proxy_Real0000000(fields []*registry.DecodedField) *pbgear.Proxy_Real {
	return &pbgear.Proxy_Real{
		Value: to_oneof_Proxy_Value(fields[0].Value),
	}
}

func to_Bounties_Id0000000(fields []*registry.DecodedField) *pbgear.Bounties_Id {
	return &pbgear.Bounties_Id{
		Value_0: to_Bounties_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_ChildBounties_AwardChildBountyCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_AwardChildBountyCall {
	return &pbgear.ChildBounties_AwardChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
		Beneficiary:    to_ChildBounties_Beneficiary(fields[2].Value),
	}
}

func to_Babe_SpRuntimeGenericDigestDigest0000000(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
	return &pbgear.Babe_SpRuntimeGenericDigestDigest{
		Logs: to_repeated_Babe_SpRuntimeGenericDigestDigestItem(fields[0].Value),
	}
}

func to_Staking_UnbondCall0000000(fields []*registry.DecodedField) *pbgear.Staking_UnbondCall {
	return &pbgear.Staking_UnbondCall{
		Value: to_Staking_CompactString(fields[0].Value),
	}
}

func to_FellowshipCollective_PromoteMemberCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_PromoteMemberCall {
	return &pbgear.FellowshipCollective_PromoteMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value),
	}
}

func to_Scheduler_TupleUint32Uint320000000(fields []*registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
	return &pbgear.Scheduler_TupleUint32Uint32{
		Value_0: uint32(fields[0].Value),
		Value_1: uint32(fields[1].Value),
	}
}

func to_Identity_Id0000000(fields []*registry.DecodedField) *pbgear.Identity_Id {
	return &pbgear.Identity_Id{
		Value_0: to_Identity_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Gear_SendReplyCall0000000(fields []*registry.DecodedField) *pbgear.Gear_SendReplyCall {
	return &pbgear.Gear_SendReplyCall{
		ReplyToId: to_Gear_GprimitivesMessageId(fields[0].Value),
		Payload:   []uint32(fields[1].Value),
		GasLimit:  uint64(fields[2].Value),
		Value:     string(fields[3].Value),
		KeepAlive: bool(fields[4].Value),
	}
}

func to_Treasury_RemoveApprovalCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_RemoveApprovalCall {
	return &pbgear.Treasury_RemoveApprovalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value),
	}
}

func to_FellowshipCollective_RemoveMemberCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_RemoveMemberCall {
	return &pbgear.FellowshipCollective_RemoveMemberCall{
		Who:     to_FellowshipCollective_Who(fields[0].Value),
		MinRank: uint32(fields[1].Value),
	}
}

func to_FellowshipReferenda_RefundDecisionDepositCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
	return &pbgear.FellowshipReferenda_RefundDecisionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Bounties_CloseBountyCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_CloseBountyCall {
	return &pbgear.Bounties_CloseBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
	}
}

func to_StakingRewards_AlignSupplyCall0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_AlignSupplyCall {
	return &pbgear.StakingRewards_AlignSupplyCall{
		Target: string(fields[0].Value),
	}
}

func to_ImOnline_HeartbeatCall0000000(fields []*registry.DecodedField) *pbgear.ImOnline_HeartbeatCall {
	return &pbgear.ImOnline_HeartbeatCall{
		Heartbeat: to_ImOnline_PalletImOnlineHeartbeat(fields[0].Value),
		Signature: to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(fields[1].Value),
	}
}

func to_Staking_MinNominatorBond0000000(fields []*registry.DecodedField) *pbgear.Staking_MinNominatorBond {
	return &pbgear.Staking_MinNominatorBond{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_oneof_Staking_Value(in *registry.DecodedField) *pbgear.Staking_Value {
	return nil
}

func to_Treasury_CheckStatusCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_CheckStatusCall {
	return &pbgear.Treasury_CheckStatusCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_Sha2560000000(fields []*registry.DecodedField) *pbgear.Identity_Sha256 {
	return &pbgear.Identity_Sha256{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_Id0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Id {
	return &pbgear.NominationPools_Id{
		Value_0: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_System_SetHeapPagesCall0000000(fields []*registry.DecodedField) *pbgear.System_SetHeapPagesCall {
	return &pbgear.System_SetHeapPagesCall{
		Pages: uint64(fields[0].Value),
	}
}

func to_Referenda_None0000000(fields []*registry.DecodedField) *pbgear.Referenda_None {
	return &pbgear.Referenda_None{}
}
func to_FellowshipReferenda_SubmitCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SubmitCall {
	return &pbgear.FellowshipReferenda_SubmitCall{
		ProposalOrigin:  to_FellowshipReferenda_ProposalOrigin(fields[0].Value),
		Proposal:        to_FellowshipReferenda_Proposal(fields[1].Value),
		EnactmentMoment: to_FellowshipReferenda_EnactmentMoment(fields[2].Value),
	}
}

func to_Bounties_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Bounties_SpCoreCryptoAccountId32 {
	return &pbgear.Bounties_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_SetStateCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetStateCall {
	return &pbgear.NominationPools_SetStateCall{
		PoolId: uint32(fields[0].Value),
		State:  to_NominationPools_State(fields[1].Value),
	}
}

func to_Bounties_Index0000000(fields []*registry.DecodedField) *pbgear.Bounties_Index {
	return &pbgear.Bounties_Index{
		Value_0: to_Bounties_CompactTupleNull(fields[0].Value),
	}
}

func to_NominationPools_FreeBalance0000000(fields []*registry.DecodedField) *pbgear.NominationPools_FreeBalance {
	return &pbgear.NominationPools_FreeBalance{
		Value_0: string(fields[0].Value),
	}
}

func to_Babe_Consensus0000000(fields []*registry.DecodedField) *pbgear.Babe_Consensus {
	return &pbgear.Babe_Consensus{
		Value_0: []uint32(fields[0].Value),
		Value_1: []uint32(fields[1].Value),
	}
}

func to_Staking_ChillThreshold0000000(fields []*registry.DecodedField) *pbgear.Staking_ChillThreshold {
	return &pbgear.Staking_ChillThreshold{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_ConvictionVoting_To0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_To {
	return &pbgear.ConvictionVoting_To{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value),
	}
}

func to_oneof_ConvictionVoting_Value(in *registry.DecodedField) *pbgear.ConvictionVoting_Value {
	return nil
}

func to_Referenda_RefundDecisionDepositCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_RefundDecisionDepositCall {
	return &pbgear.Referenda_RefundDecisionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_Raw250000000(fields []*registry.DecodedField) *pbgear.Identity_Raw25 {
	return &pbgear.Identity_Raw25{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Session_PalletImOnlineSr25519AppSr25519Public0000000(fields []*registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
	return &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{
		ImOnline: to_Session_SpCoreSr25519Public(fields[0].Value),
	}
}

func to_ConvictionVoting_Locked1X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked1X {
	return &pbgear.ConvictionVoting_Locked1X{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_StakingRewards_WithdrawCall0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_WithdrawCall {
	return &pbgear.StakingRewards_WithdrawCall{
		To:    to_StakingRewards_To(fields[0].Value),
		Value: string(fields[1].Value),
	}
}

func to_GearVoucher_Some0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_Some {
	return &pbgear.GearVoucher_Some{
		Value_0: to_GearVoucher_BTreeSet(fields[0].Value),
	}
}

func to_FellowshipReferenda_Origins0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Origins {
	return &pbgear.FellowshipReferenda_Origins{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value),
	}
}

func to_Proxy_CancelProxy0000000(fields []*registry.DecodedField) *pbgear.Proxy_CancelProxy {
	return &pbgear.Proxy_CancelProxy{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ChildBounties_CloseChildBountyCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_CloseChildBountyCall {
	return &pbgear.ChildBounties_CloseChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
	}
}

func to_ElectionProviderMultiPhase_SubmitCall0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitCall {
	return &pbgear.ElectionProviderMultiPhase_SubmitCall{
		RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value),
	}
}

func to_NominationPools_ClaimPayoutOtherCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutOtherCall {
	return &pbgear.NominationPools_ClaimPayoutOtherCall{
		Other: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Babe_PrimarySlots0000000(fields []*registry.DecodedField) *pbgear.Babe_PrimarySlots {
	return &pbgear.Babe_PrimarySlots{}
}
func to_Vesting_ForceVestedTransferCall0000000(fields []*registry.DecodedField) *pbgear.Vesting_ForceVestedTransferCall {
	return &pbgear.Vesting_ForceVestedTransferCall{
		Source:   to_Vesting_Source(fields[0].Value),
		Target:   to_Vesting_Target(fields[1].Value),
		Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(fields[2].Value),
	}
}

func to_BagsList_RebagCall0000000(fields []*registry.DecodedField) *pbgear.BagsList_RebagCall {
	return &pbgear.BagsList_RebagCall{
		Dislocated: to_BagsList_Dislocated(fields[0].Value),
	}
}

func to_Staking_IncreaseValidatorCountCall0000000(fields []*registry.DecodedField) *pbgear.Staking_IncreaseValidatorCountCall {
	return &pbgear.Staking_IncreaseValidatorCountCall{
		Additional: to_Staking_CompactUint32(fields[0].Value),
	}
}

func to_Identity_Sub0000000(fields []*registry.DecodedField) *pbgear.Identity_Sub {
	return &pbgear.Identity_Sub{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_System_TupleSystemItemsListSystemItemsList0000000(fields []*registry.DecodedField) *pbgear.System_TupleSystemItemsListSystemItemsList {
	return &pbgear.System_TupleSystemItemsListSystemItemsList{
		Value_0: []uint32(fields[0].Value),
		Value_1: []uint32(fields[1].Value),
	}
}

func to_Vesting_PalletVestingVestingInfoVestingInfo0000000(fields []*registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
	return &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{
		Locked:        string(fields[0].Value),
		PerBlock:      string(fields[1].Value),
		StartingBlock: uint32(fields[2].Value),
	}
}

func to_Identity_Erroneous0000000(fields []*registry.DecodedField) *pbgear.Identity_Erroneous {
	return &pbgear.Identity_Erroneous{}
}
func to_NominationPools_SetCommissionChangeRateCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionChangeRateCall {
	return &pbgear.NominationPools_SetCommissionChangeRateCall{
		PoolId:     uint32(fields[0].Value),
		ChangeRate: to_NominationPools_PalletNominationPoolsCommissionChangeRate(fields[1].Value),
	}
}

func to_Treasury_Address200000000(fields []*registry.DecodedField) *pbgear.Treasury_Address20 {
	return &pbgear.Treasury_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_Index0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Index {
	return &pbgear.FellowshipCollective_Index{
		Value_0: to_FellowshipCollective_CompactTupleNull(fields[0].Value),
	}
}

func to_Grandpa_SpConsensusGrandpaEquivocationProof0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
	return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{
		SetId:        uint64(fields[0].Value),
		Equivocation: to_Grandpa_Equivocation(fields[1].Value),
	}
}

func to_Scheduler_CancelCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_CancelCall {
	return &pbgear.Scheduler_CancelCall{
		When:  uint32(fields[0].Value),
		Index: uint32(fields[1].Value),
	}
}

func to_Proxy_ProxyCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_ProxyCall {
	return &pbgear.Proxy_ProxyCall{
		Real:           to_Proxy_Real(fields[0].Value),
		ForceProxyType: to_optional_Proxy_ForceProxyType(fields[1].Value),
		Call:           to_oneof_Proxy_Call(fields[2].Value),
	}
}

func to_oneof_Proxy_Call(in *registry.DecodedField) *pbgear.Proxy_Call {
	return nil
}

func to_ChildBounties_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.ChildBounties_SpCoreCryptoAccountId32 {
	return &pbgear.ChildBounties_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_NewRoot0000000(fields []*registry.DecodedField) *pbgear.NominationPools_NewRoot {
	return &pbgear.NominationPools_NewRoot{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_ClaimCommissionCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimCommissionCall {
	return &pbgear.NominationPools_ClaimCommissionCall{
		PoolId: uint32(fields[0].Value),
	}
}

func to_Balances_Raw0000000(fields []*registry.DecodedField) *pbgear.Balances_Raw {
	return &pbgear.Balances_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Vesting_Index0000000(fields []*registry.DecodedField) *pbgear.Vesting_Index {
	return &pbgear.Vesting_Index{
		Value_0: to_Vesting_CompactTupleNull(fields[0].Value),
	}
}

func to_BagsList_Index0000000(fields []*registry.DecodedField) *pbgear.BagsList_Index {
	return &pbgear.BagsList_Index{
		Value_0: to_BagsList_CompactTupleNull(fields[0].Value),
	}
}

func to_BagsList_Lighter0000000(fields []*registry.DecodedField) *pbgear.BagsList_Lighter {
	return &pbgear.BagsList_Lighter{
		Value: to_oneof_BagsList_Value(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU160000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{
		Value: to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(fields[0].Value),
	}
}

func to_Balances_Who0000000(fields []*registry.DecodedField) *pbgear.Balances_Who {
	return &pbgear.Balances_Who{
		Value: to_oneof_Balances_Value(fields[0].Value),
	}
}

func to_ConvictionVoting_Split0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Split {
	return &pbgear.ConvictionVoting_Split{
		Aye: string(fields[0].Value),
		Nay: string(fields[1].Value),
	}
}

func to_Gear_CreateProgramCall0000000(fields []*registry.DecodedField) *pbgear.Gear_CreateProgramCall {
	return &pbgear.Gear_CreateProgramCall{
		CodeId:      to_Gear_GprimitivesCodeId(fields[0].Value),
		Salt:        []uint32(fields[1].Value),
		InitPayload: []uint32(fields[2].Value),
		GasLimit:    uint64(fields[3].Value),
		Value:       string(fields[4].Value),
		KeepAlive:   bool(fields[5].Value),
	}
}

func to_Referenda_OneFewerDecidingCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_OneFewerDecidingCall {
	return &pbgear.Referenda_OneFewerDecidingCall{
		Track: uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_TupleNull0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
	return &pbgear.FellowshipCollective_TupleNull{}
}
func to_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData0000000(fields []*registry.DecodedField) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
	return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{
		Value_0: to_Identity_Value0(fields[0].Value),
		Value_1: to_Identity_Value1(fields[1].Value),
	}
}

func to_Preimage_EnsureUpdatedCall0000000(fields []*registry.DecodedField) *pbgear.Preimage_EnsureUpdatedCall {
	return &pbgear.Preimage_EnsureUpdatedCall{
		Hashes: to_repeated_Preimage_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Identity_Raw290000000(fields []*registry.DecodedField) *pbgear.Identity_Raw29 {
	return &pbgear.Identity_Raw29{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ChildBounties_Address200000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Address20 {
	return &pbgear.ChildBounties_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_System_SetStorageCall0000000(fields []*registry.DecodedField) *pbgear.System_SetStorageCall {
	return &pbgear.System_SetStorageCall{
		Items: to_repeated_System_TupleSystemItemsListSystemItemsList(fields[0].Value),
	}
}

func to_Grandpa_SpCoreEd25519Public0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Public {
	return &pbgear.Grandpa_SpCoreEd25519Public{
		Identity: []uint32(fields[0].Value),
	}
}

func to_Grandpa_SpConsensusGrandpaAppSignature0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
	return &pbgear.Grandpa_SpConsensusGrandpaAppSignature{
		Value_1: to_Grandpa_SpCoreEd25519Signature(fields[0].Value),
	}
}

func to_Staking_MinValidatorBond0000000(fields []*registry.DecodedField) *pbgear.Staking_MinValidatorBond {
	return &pbgear.Staking_MinValidatorBond{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Utility_Value00000000(fields []*registry.DecodedField) *pbgear.Utility_Value0 {
	return &pbgear.Utility_Value0{
		Value: to_oneof_Utility_Value(fields[0].Value),
	}
}

func to_oneof_Utility_Value(in *registry.DecodedField) *pbgear.Utility_Value {
	return nil
}

func to_NominationPools_PermissionlessWithdraw0000000(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessWithdraw {
	return &pbgear.NominationPools_PermissionlessWithdraw{}
}
func to_Utility_System0000000(fields []*registry.DecodedField) *pbgear.Utility_System {
	return &pbgear.Utility_System{
		Value_0: to_Utility_Value0(fields[0].Value),
	}
}

func to_NominationPools_SpArithmeticPerThingsPerbill0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
	return &pbgear.NominationPools_SpArithmeticPerThingsPerbill{
		Value_0: uint32(fields[0].Value),
	}
}

func to_NominationPools_AdjustPoolDepositCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_AdjustPoolDepositCall {
	return &pbgear.NominationPools_AdjustPoolDepositCall{
		PoolId: uint32(fields[0].Value),
	}
}

func to_Balances_CompactString0000000(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
	return &pbgear.Balances_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Vesting_VestOtherCall0000000(fields []*registry.DecodedField) *pbgear.Vesting_VestOtherCall {
	return &pbgear.Vesting_VestOtherCall{
		Target: to_Vesting_Target(fields[0].Value),
	}
}

func to_ImOnline_PalletImOnlineSr25519AppSr25519Signature0000000(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
	return &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{
		Signature: to_ImOnline_SpCoreSr25519Signature(fields[0].Value),
	}
}

func to_FellowshipReferenda_Lookup0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Lookup {
	return &pbgear.FellowshipReferenda_Lookup{
		Hash: to_FellowshipReferenda_PrimitiveTypesH256(fields[0].Value),
		Len:  uint32(fields[1].Value),
	}
}

func to_Identity_BlakeTwo2560000000(fields []*registry.DecodedField) *pbgear.Identity_BlakeTwo256 {
	return &pbgear.Identity_BlakeTwo256{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_CleanupPollCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CleanupPollCall {
	return &pbgear.FellowshipCollective_CleanupPollCall{
		PollIndex: uint32(fields[0].Value),
		Max:       uint32(fields[1].Value),
	}
}

func to_FellowshipReferenda_Proposal0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
	return &pbgear.FellowshipReferenda_Proposal{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value),
	}
}

func to_Proxy_KillPureCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_KillPureCall {
	return &pbgear.Proxy_KillPureCall{
		Spawner:   to_Proxy_Spawner(fields[0].Value),
		ProxyType: to_Proxy_ProxyType(fields[1].Value),
		Index:     uint32(fields[2].Value),
		Height:    to_Proxy_CompactUint32(fields[3].Value),
		ExtIndex:  to_Proxy_CompactUint32(fields[4].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_StakingRewards_Raw0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_Raw {
	return &pbgear.StakingRewards_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_GearVoucher_PalletGearVoucherInternalVoucherId0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
	return &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{
		VoucherId: []uint32(fields[0].Value),
	}
}

func to_Referenda_SubmitCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_SubmitCall {
	return &pbgear.Referenda_SubmitCall{
		ProposalOrigin:  to_Referenda_ProposalOrigin(fields[0].Value),
		Proposal:        to_Referenda_Proposal(fields[1].Value),
		EnactmentMoment: to_Referenda_EnactmentMoment(fields[2].Value),
	}
}

func to_FellowshipCollective_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
	return &pbgear.FellowshipCollective_CompactTupleNull{
		Value: to_FellowshipCollective_TupleNull(fields[0].Value),
	}
}

func to_Scheduler_ScheduleAfterCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleAfterCall {
	return &pbgear.Scheduler_ScheduleAfterCall{
		After:         uint32(fields[0].Value),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[1].Value),
		Priority:      uint32(fields[2].Value),
		Call:          to_oneof_Scheduler_Call(fields[3].Value),
	}
}

func to_oneof_Scheduler_Call(in *registry.DecodedField) *pbgear.Scheduler_Call {
	return nil
}

func to_Identity_Raw260000000(fields []*registry.DecodedField) *pbgear.Identity_Raw26 {
	return &pbgear.Identity_Raw26{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_BondExtraCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraCall {
	return &pbgear.NominationPools_BondExtraCall{
		Extra: to_NominationPools_Extra(fields[0].Value),
	}
}

func to_Proxy_Governance0000000(fields []*registry.DecodedField) *pbgear.Proxy_Governance {
	return &pbgear.Proxy_Governance{}
}
func to_NominationPools_MinJoinBond0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
	return &pbgear.NominationPools_MinJoinBond{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
	return &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{
		Value_0: to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value),
		Value_1: to_NominationPools_SpCoreCryptoAccountId32(fields[1].Value),
	}
}

func to_ConvictionVoting_Conviction0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
	return &pbgear.ConvictionVoting_Conviction{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_GovernanceFallbackCall0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
	return &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{
		MaybeMaxVoters:  uint32(fields[0].Value),
		MaybeMaxTargets: uint32(fields[1].Value),
	}
}

func to_optional_ElectionProviderMultiPhase_uint32(in *registry.DecodedField) *uint32 {
	var out *uint32

	if v, ok := in.Value.(uint32); ok {
		if v != 0 {
			panic("expected none")
		}
	} else {
		out = in.Value.(*uint32)
	}
	return out
}

func to_ChildBounties_ClaimChildBountyCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_ClaimChildBountyCall {
	return &pbgear.ChildBounties_ClaimChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
	}
}

func to_NominationPools_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
	return &pbgear.NominationPools_CompactTupleNull{
		Value: to_NominationPools_TupleNull(fields[0].Value),
	}
}

func to_NominationPools_Noop0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Noop {
	return &pbgear.NominationPools_Noop{}
}
func to_Identity_Address200000000(fields []*registry.DecodedField) *pbgear.Identity_Address20 {
	return &pbgear.Identity_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Bounties_UnassignCuratorCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_UnassignCuratorCall {
	return &pbgear.Bounties_UnassignCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
	}
}

func to_NominationPools_Destroying0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Destroying {
	return &pbgear.NominationPools_Destroying{}
}
func to_BagsList_TupleNull0000000(fields []*registry.DecodedField) *pbgear.BagsList_TupleNull {
	return &pbgear.BagsList_TupleNull{}
}
func to_BagsList_Heavier0000000(fields []*registry.DecodedField) *pbgear.BagsList_Heavier {
	return &pbgear.BagsList_Heavier{
		Value: to_oneof_BagsList_Value(fields[0].Value),
	}
}

func to_Staking_PayoutStakersCall0000000(fields []*registry.DecodedField) *pbgear.Staking_PayoutStakersCall {
	return &pbgear.Staking_PayoutStakersCall{
		ValidatorStash: to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
		Era:            uint32(fields[1].Value),
	}
}

func to_ConvictionVoting_Id0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Id {
	return &pbgear.ConvictionVoting_Id{
		Value_0: to_ConvictionVoting_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Identity_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Identity_CompactTupleNull {
	return &pbgear.Identity_CompactTupleNull{
		Value: to_Identity_TupleNull(fields[0].Value),
	}
}

func to_NominationPools_SetConfigsCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetConfigsCall {
	return &pbgear.NominationPools_SetConfigsCall{
		MinJoinBond:         to_NominationPools_MinJoinBond(fields[0].Value),
		MinCreateBond:       to_NominationPools_MinCreateBond(fields[1].Value),
		MaxPools:            to_NominationPools_MaxPools(fields[2].Value),
		MaxMembers:          to_NominationPools_MaxMembers(fields[3].Value),
		MaxMembersPerPool:   to_NominationPools_MaxMembersPerPool(fields[4].Value),
		GlobalMaxCommission: to_NominationPools_GlobalMaxCommission(fields[5].Value),
	}
}

func to_Staking_ForceNoErasCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ForceNoErasCall {
	return &pbgear.Staking_ForceNoErasCall{}
}
func to_Proxy_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
	return &pbgear.Proxy_CompactTupleNull{
		Value: to_Proxy_TupleNull(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU160000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{
		Value: uint32(fields[0].Value),
	}
}

func to_NominationPools_PermissionlessAll0000000(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessAll {
	return &pbgear.NominationPools_PermissionlessAll{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
	return &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{
		MaybeNextScore: to_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[0].Value),
	}
}

func to_Treasury_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
	return &pbgear.Treasury_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Utility_SpWeightsWeightV2Weight0000000(fields []*registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
	return &pbgear.Utility_SpWeightsWeightV2Weight{
		RefTime:   to_Utility_CompactUint64(fields[0].Value),
		ProofSize: to_Utility_CompactUint64(fields[1].Value),
	}
}

func to_ConvictionVoting_SplitAbstain0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SplitAbstain {
	return &pbgear.ConvictionVoting_SplitAbstain{
		Aye:     string(fields[0].Value),
		Nay:     string(fields[1].Value),
		Abstain: string(fields[2].Value),
	}
}

func to_ConvictionVoting_Address320000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address32 {
	return &pbgear.ConvictionVoting_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_SetAccountIdCall0000000(fields []*registry.DecodedField) *pbgear.Identity_SetAccountIdCall {
	return &pbgear.Identity_SetAccountIdCall{
		Index: to_Identity_CompactUint32(fields[0].Value),
		New:   to_Identity_New(fields[1].Value),
	}
}

func to_Staking_SetValidatorCountCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetValidatorCountCall {
	return &pbgear.Staking_SetValidatorCountCall{
		New: to_Staking_CompactUint32(fields[0].Value),
	}
}

func to_Treasury_SpendCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_SpendCall {
	return &pbgear.Treasury_SpendCall{
		AssetKind:   to_Treasury_TupleNull(fields[0].Value),
		Amount:      to_Treasury_CompactString(fields[1].Value),
		Beneficiary: to_Treasury_SpCoreCryptoAccountId32(fields[2].Value),
		ValidFrom:   uint32(fields[3].Value),
	}
}

func to_Identity_FeePaid0000000(fields []*registry.DecodedField) *pbgear.Identity_FeePaid {
	return &pbgear.Identity_FeePaid{
		Value_0: string(fields[0].Value),
	}
}

func to_Grandpa_Equivocation0000000(fields []*registry.DecodedField) *pbgear.Grandpa_Equivocation {
	return &pbgear.Grandpa_Equivocation{
		Value: to_oneof_Grandpa_Value(fields[0].Value),
	}
}

func to_oneof_Grandpa_Value(in *registry.DecodedField) *pbgear.Grandpa_Value {
	return nil
}

func to_Referenda_SetMetadataCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_SetMetadataCall {
	return &pbgear.Referenda_SetMetadataCall{
		Index:     uint32(fields[0].Value),
		MaybeHash: to_optional_Referenda_PrimitiveTypesH256(fields[1].Value),
	}
}

func to_Identity_Raw70000000(fields []*registry.DecodedField) *pbgear.Identity_Raw7 {
	return &pbgear.Identity_Raw7{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ChildBounties_AddChildBountyCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_AddChildBountyCall {
	return &pbgear.ChildBounties_AddChildBountyCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		Value:          to_ChildBounties_CompactString(fields[1].Value),
		Description:    []uint32(fields[2].Value),
	}
}

func to_Babe_AllowedSlots0000000(fields []*registry.DecodedField) *pbgear.Babe_AllowedSlots {
	return &pbgear.Babe_AllowedSlots{
		Value: to_oneof_Babe_Value(fields[0].Value),
	}
}

func to_oneof_Babe_Value(in *registry.DecodedField) *pbgear.Babe_Value {
	return nil
}

func to_Staking_MaxNominatorCount0000000(fields []*registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
	return &pbgear.Staking_MaxNominatorCount{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Utility_AsDerivativeCall0000000(fields []*registry.DecodedField) *pbgear.Utility_AsDerivativeCall {
	return &pbgear.Utility_AsDerivativeCall{
		Index: uint32(fields[0].Value),
		Call:  to_oneof_Utility_Call(fields[1].Value),
	}
}

func to_oneof_Utility_Call(in *registry.DecodedField) *pbgear.Utility_Call {
	return nil
}

func to_Identity_Raw10000000(fields []*registry.DecodedField) *pbgear.Identity_Raw1 {
	return &pbgear.Identity_Raw1{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_StakingRewards_TupleNull0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_TupleNull {
	return &pbgear.StakingRewards_TupleNull{}
}
func to_Bounties_Raw0000000(fields []*registry.DecodedField) *pbgear.Bounties_Raw {
	return &pbgear.Bounties_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Babe_SpRuntimeGenericHeaderHeader0000000(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
	return &pbgear.Babe_SpRuntimeGenericHeaderHeader{
		ParentHash:     to_Babe_PrimitiveTypesH256(fields[0].Value),
		Number:         to_Babe_CompactUint32(fields[1].Value),
		StateRoot:      to_Babe_PrimitiveTypesH256(fields[2].Value),
		ExtrinsicsRoot: to_Babe_PrimitiveTypesH256(fields[3].Value),
		Digest:         to_Babe_SpRuntimeGenericDigestDigest(fields[4].Value),
	}
}

func to_Vesting_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Vesting_SpCoreCryptoAccountId32 {
	return &pbgear.Vesting_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Referenda_PlaceDecisionDepositCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_PlaceDecisionDepositCall {
	return &pbgear.Referenda_PlaceDecisionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_Raw30000000(fields []*registry.DecodedField) *pbgear.Identity_Raw3 {
	return &pbgear.Identity_Raw3{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw220000000(fields []*registry.DecodedField) *pbgear.Identity_Raw22 {
	return &pbgear.Identity_Raw22{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Treasury_PayoutCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_PayoutCall {
	return &pbgear.Treasury_PayoutCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Preimage_NotePreimageCall0000000(fields []*registry.DecodedField) *pbgear.Preimage_NotePreimageCall {
	return &pbgear.Preimage_NotePreimageCall{
		Bytes: []uint32(fields[0].Value),
	}
}

func to_Identity_Address320000000(fields []*registry.DecodedField) *pbgear.Identity_Address32 {
	return &pbgear.Identity_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_GearVoucher_GprimitivesActorId0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
	return &pbgear.GearVoucher_GprimitivesActorId{
		Programs: to_GearVoucher_GprimitivesActorId(fields[0].Value),
	}
}

func to_BagsList_PutInFrontOfOtherCall0000000(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfOtherCall {
	return &pbgear.BagsList_PutInFrontOfOtherCall{
		Heavier: to_BagsList_Heavier(fields[0].Value),
		Lighter: to_BagsList_Lighter(fields[1].Value),
	}
}

func to_FellowshipReferenda_CancelCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_CancelCall {
	return &pbgear.FellowshipReferenda_CancelCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_RenameSubCall0000000(fields []*registry.DecodedField) *pbgear.Identity_RenameSubCall {
	return &pbgear.Identity_RenameSubCall{
		Sub:  to_Identity_Sub(fields[0].Value),
		Data: to_Identity_Data(fields[1].Value),
	}
}

func to_NominationPools_NewBouncer0000000(fields []*registry.DecodedField) *pbgear.NominationPools_NewBouncer {
	return &pbgear.NominationPools_NewBouncer{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Gear_SetExecuteInherentCall0000000(fields []*registry.DecodedField) *pbgear.Gear_SetExecuteInherentCall {
	return &pbgear.Gear_SetExecuteInherentCall{
		Value: bool(fields[0].Value),
	}
}

func to_Babe_SpCoreSr25519Public0000000(fields []*registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
	return &pbgear.Babe_SpCoreSr25519Public{
		Offender: []uint32(fields[0].Value),
	}
}

func to_Babe_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Babe_CompactUint32 {
	return &pbgear.Babe_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_Multisig_PalletMultisigTimepoint0000000(fields []*registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
	return &pbgear.Multisig_PalletMultisigTimepoint{
		Height: uint32(fields[0].Value),
		Index:  uint32(fields[1].Value),
	}
}

func to_Gear_UploadProgramCall0000000(fields []*registry.DecodedField) *pbgear.Gear_UploadProgramCall {
	return &pbgear.Gear_UploadProgramCall{
		Code:        []uint32(fields[0].Value),
		Salt:        []uint32(fields[1].Value),
		InitPayload: []uint32(fields[2].Value),
		GasLimit:    uint64(fields[3].Value),
		Value:       string(fields[4].Value),
		KeepAlive:   bool(fields[5].Value),
	}
}

func to_GearVoucher_Call0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_Call {
	return &pbgear.GearVoucher_Call{
		Value: to_oneof_GearVoucher_Value(fields[0].Value),
	}
}

func to_oneof_GearVoucher_Value(in *registry.DecodedField) *pbgear.GearVoucher_Value {
	return nil
}

func to_Proxy_IdentityJudgement0000000(fields []*registry.DecodedField) *pbgear.Proxy_IdentityJudgement {
	return &pbgear.Proxy_IdentityJudgement{}
}
func to_Multisig_CancelAsMultiCall0000000(fields []*registry.DecodedField) *pbgear.Multisig_CancelAsMultiCall {
	return &pbgear.Multisig_CancelAsMultiCall{
		Threshold:        uint32(fields[0].Value),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value),
		Timepoint:        to_Multisig_PalletMultisigTimepoint(fields[2].Value),
		CallHash:         []uint32(fields[3].Value),
	}
}

func to_GearVoucher_SendReply0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_SendReply {
	return &pbgear.GearVoucher_SendReply{
		ReplyToId: to_GearVoucher_GprimitivesMessageId(fields[0].Value),
		Payload:   []uint32(fields[1].Value),
		GasLimit:  uint64(fields[2].Value),
		Value:     string(fields[3].Value),
		KeepAlive: bool(fields[4].Value),
	}
}

func to_Babe_Config0000000(fields []*registry.DecodedField) *pbgear.Babe_Config {
	return &pbgear.Babe_Config{
		Value: to_oneof_Babe_Value(fields[0].Value),
	}
}

func to_FellowshipReferenda_None0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_None {
	return &pbgear.FellowshipReferenda_None{}
}
func to_Whitelist_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
	return &pbgear.Whitelist_PrimitiveTypesH256{
		CallHash: []uint32(fields[0].Value),
	}
}

func to_Identity_CompactString0000000(fields []*registry.DecodedField) *pbgear.Identity_CompactString {
	return &pbgear.Identity_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Identity_AddSubCall0000000(fields []*registry.DecodedField) *pbgear.Identity_AddSubCall {
	return &pbgear.Identity_AddSubCall{
		Sub:  to_Identity_Sub(fields[0].Value),
		Data: to_Identity_Data(fields[1].Value),
	}
}

func to_Preimage_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
	return &pbgear.Preimage_PrimitiveTypesH256{
		Hash: []uint32(fields[0].Value),
	}
}

func to_Identity_RemoveSubCall0000000(fields []*registry.DecodedField) *pbgear.Identity_RemoveSubCall {
	return &pbgear.Identity_RemoveSubCall{
		Sub: to_Identity_Sub(fields[0].Value),
	}
}

func to_Identity_QuitSubCall0000000(fields []*registry.DecodedField) *pbgear.Identity_QuitSubCall {
	return &pbgear.Identity_QuitSubCall{}
}
func to_Babe_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
	return &pbgear.Babe_PrimitiveTypesH256{
		ParentHash: []uint32(fields[0].Value),
	}
}

func to_Grandpa_SpConsensusGrandpaAppPublic0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
	return &pbgear.Grandpa_SpConsensusGrandpaAppPublic{
		Identity: to_Grandpa_SpCoreEd25519Public(fields[0].Value),
	}
}

func to_Vesting_Source0000000(fields []*registry.DecodedField) *pbgear.Vesting_Source {
	return &pbgear.Vesting_Source{
		Value: to_oneof_Vesting_Value(fields[0].Value),
	}
}

func to_oneof_Vesting_Value(in *registry.DecodedField) *pbgear.Vesting_Value {
	return nil
}

func to_Staking_Id0000000(fields []*registry.DecodedField) *pbgear.Staking_Id {
	return &pbgear.Staking_Id{
		Value_0: to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Treasury_SpendLocalCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_SpendLocalCall {
	return &pbgear.Treasury_SpendLocalCall{
		Amount:      to_Treasury_CompactString(fields[0].Value),
		Beneficiary: to_Treasury_Beneficiary(fields[1].Value),
	}
}

func to_Treasury_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
	return &pbgear.Treasury_CompactTupleNull{
		Value: to_Treasury_TupleNull(fields[0].Value),
	}
}

func to_Referenda_CancelCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_CancelCall {
	return &pbgear.Referenda_CancelCall{
		Index: uint32(fields[0].Value),
	}
}

func to_ChildBounties_Curator0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Curator {
	return &pbgear.ChildBounties_Curator{
		Value: to_oneof_ChildBounties_Value(fields[0].Value),
	}
}

func to_oneof_ChildBounties_Value(in *registry.DecodedField) *pbgear.ChildBounties_Value {
	return nil
}

func to_NominationPools_Set0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Set {
	return &pbgear.NominationPools_Set{
		Value_0: to_NominationPools_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Babe_BabeTrieNodesList0000000(fields []*registry.DecodedField) *pbgear.Babe_BabeTrieNodesList {
	return &pbgear.Babe_BabeTrieNodesList{
		TrieNodes: []uint32(fields[0].Value),
	}
}

func to_Referenda_Origins0000000(fields []*registry.DecodedField) *pbgear.Referenda_Origins {
	return &pbgear.Referenda_Origins{
		Value_0: to_Referenda_Value0(fields[0].Value),
	}
}

func to_Multisig_SpWeightsWeightV2Weight0000000(fields []*registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
	return &pbgear.Multisig_SpWeightsWeightV2Weight{
		RefTime:   to_Multisig_CompactUint64(fields[0].Value),
		ProofSize: to_Multisig_CompactUint64(fields[1].Value),
	}
}

func to_Bounties_CompactString0000000(fields []*registry.DecodedField) *pbgear.Bounties_CompactString {
	return &pbgear.Bounties_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Gear_UploadCodeCall0000000(fields []*registry.DecodedField) *pbgear.Gear_UploadCodeCall {
	return &pbgear.Gear_UploadCodeCall{
		Code: []uint32(fields[0].Value),
	}
}

func to_Babe_PlanConfigChangeCall0000000(fields []*registry.DecodedField) *pbgear.Babe_PlanConfigChangeCall {
	return &pbgear.Babe_PlanConfigChangeCall{
		Config: to_Babe_Config(fields[0].Value),
	}
}

func to_Treasury_Raw0000000(fields []*registry.DecodedField) *pbgear.Treasury_Raw {
	return &pbgear.Treasury_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_ProvideJudgementCall0000000(fields []*registry.DecodedField) *pbgear.Identity_ProvideJudgementCall {
	return &pbgear.Identity_ProvideJudgementCall{
		RegIndex:  to_Identity_CompactUint32(fields[0].Value),
		Target:    to_Identity_Target(fields[1].Value),
		Judgement: to_Identity_Judgement(fields[2].Value),
		Identity:  to_Identity_PrimitiveTypesH256(fields[3].Value),
	}
}

func to_Bounties_AwardBountyCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_AwardBountyCall {
	return &pbgear.Bounties_AwardBountyCall{
		BountyId:    to_Bounties_CompactUint32(fields[0].Value),
		Beneficiary: to_Bounties_Beneficiary(fields[1].Value),
	}
}

func to_ChildBounties_TupleNull0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_TupleNull {
	return &pbgear.ChildBounties_TupleNull{}
}
func to_NominationPools_MaxMembers0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembers {
	return &pbgear.NominationPools_MaxMembers{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_BondExtraOtherCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraOtherCall {
	return &pbgear.NominationPools_BondExtraOtherCall{
		Member: to_NominationPools_Member(fields[0].Value),
		Extra:  to_NominationPools_Extra(fields[1].Value),
	}
}

func to_Staking_SetInvulnerablesCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetInvulnerablesCall {
	return &pbgear.Staking_SetInvulnerablesCall{
		Invulnerables: to_repeated_Staking_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_ConvictionVoting_Vote0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Vote {
	return &pbgear.ConvictionVoting_Vote{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value),
	}
}

func to_ConvictionVoting_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SpCoreCryptoAccountId32 {
	return &pbgear.ConvictionVoting_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ChildBounties_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
	return &pbgear.ChildBounties_CompactTupleNull{
		Value: to_ChildBounties_TupleNull(fields[0].Value),
	}
}

func to_NominationPools_NominateCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_NominateCall {
	return &pbgear.NominationPools_NominateCall{
		PoolId:     uint32(fields[0].Value),
		Validators: to_repeated_NominationPools_SpCoreCryptoAccountId32(fields[1].Value),
	}
}

func to_Bounties_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Bounties_TupleNull {
	return &pbgear.Bounties_TupleNull{}
}
func to_ChildBounties_CompactString0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactString {
	return &pbgear.ChildBounties_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_NominationPools_Address200000000(fields []*registry.DecodedField) *pbgear.NominationPools_Address20 {
	return &pbgear.NominationPools_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_CompactString0000000(fields []*registry.DecodedField) *pbgear.Staking_CompactString {
	return &pbgear.Staking_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Utility_None0000000(fields []*registry.DecodedField) *pbgear.Utility_None {
	return &pbgear.Utility_None{}
}
func to_FellowshipCollective_Address320000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address32 {
	return &pbgear.FellowshipCollective_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_VaraRuntimeNposSolution160000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
	return &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{
		Votes_1:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(fields[0].Value),
		Votes_2:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(fields[1].Value),
		Votes_3:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(fields[2].Value),
		Votes_4:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(fields[3].Value),
		Votes_5:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(fields[4].Value),
		Votes_6:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(fields[5].Value),
		Votes_7:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(fields[6].Value),
		Votes_8:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(fields[7].Value),
		Votes_9:  to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(fields[8].Value),
		Votes_10: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(fields[9].Value),
		Votes_11: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(fields[10].Value),
		Votes_12: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(fields[11].Value),
		Votes_13: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(fields[12].Value),
		Votes_14: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(fields[13].Value),
		Votes_15: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(fields[14].Value),
		Votes_16: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(fields[15].Value),
	}
}

func to_FellowshipReferenda_Void0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Void {
	return &pbgear.FellowshipReferenda_Void{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value),
	}
}

func to_FellowshipReferenda_KillCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_KillCall {
	return &pbgear.FellowshipReferenda_KillCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_ShaThree2560000000(fields []*registry.DecodedField) *pbgear.Identity_ShaThree256 {
	return &pbgear.Identity_ShaThree256{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_BagsList_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
	return &pbgear.BagsList_CompactTupleNull{
		Value: to_BagsList_TupleNull(fields[0].Value),
	}
}

func to_BagsList_Raw0000000(fields []*registry.DecodedField) *pbgear.BagsList_Raw {
	return &pbgear.BagsList_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_SetControllerCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetControllerCall {
	return &pbgear.Staking_SetControllerCall{}
}
func to_Staking_KickCall0000000(fields []*registry.DecodedField) *pbgear.Staking_KickCall {
	return &pbgear.Staking_KickCall{
		Who: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value),
	}
}

func to_FellowshipCollective_Raw0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Raw {
	return &pbgear.FellowshipCollective_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_CreatePureCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_CreatePureCall {
	return &pbgear.Proxy_CreatePureCall{
		ProxyType: to_Proxy_ProxyType(fields[0].Value),
		Delay:     uint32(fields[1].Value),
		Index:     uint32(fields[2].Value),
	}
}

func to_NominationPools_PoolWithdrawUnbondedCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
	return &pbgear.NominationPools_PoolWithdrawUnbondedCall{
		PoolId:           uint32(fields[0].Value),
		NumSlashingSpans: uint32(fields[1].Value),
	}
}

func to_GearVoucher_BTreeSet0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
	return &pbgear.GearVoucher_BTreeSet{
		Programs: to_repeated_GearVoucher_GprimitivesActorId(fields[0].Value),
	}
}

func to_Bounties_ClaimBountyCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_ClaimBountyCall {
	return &pbgear.Bounties_ClaimBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
	}
}

func to_GearVoucher_DeclineCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineCall {
	return &pbgear.GearVoucher_DeclineCall{
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value),
	}
}

func to_Grandpa_ReportEquivocationUnsignedCall0000000(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
	return &pbgear.Grandpa_ReportEquivocationUnsignedCall{
		EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value),
		KeyOwnerProof:     to_Grandpa_SpSessionMembershipProof(fields[1].Value),
	}
}

func to_Session_SetKeysCall0000000(fields []*registry.DecodedField) *pbgear.Session_SetKeysCall {
	return &pbgear.Session_SetKeysCall{
		Keys:  to_Session_VaraRuntimeSessionKeys(fields[0].Value),
		Proof: []uint32(fields[1].Value),
	}
}

func to_Grandpa_GrandpaTrieNodesList0000000(fields []*registry.DecodedField) *pbgear.Grandpa_GrandpaTrieNodesList {
	return &pbgear.Grandpa_GrandpaTrieNodesList{
		TrieNodes: []uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_Id0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Id {
	return &pbgear.FellowshipCollective_Id{
		Value_0: to_FellowshipCollective_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_FellowshipReferenda_Root0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Root {
	return &pbgear.FellowshipReferenda_Root{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_Balances_Id0000000(fields []*registry.DecodedField) *pbgear.Balances_Id {
	return &pbgear.Balances_Id{
		Value_0: to_Balances_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Staking_WithdrawUnbondedCall0000000(fields []*registry.DecodedField) *pbgear.Staking_WithdrawUnbondedCall {
	return &pbgear.Staking_WithdrawUnbondedCall{
		NumSlashingSpans: uint32(fields[0].Value),
	}
}

func to_NominationPools_Address320000000(fields []*registry.DecodedField) *pbgear.NominationPools_Address32 {
	return &pbgear.NominationPools_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw40000000(fields []*registry.DecodedField) *pbgear.Identity_Raw4 {
	return &pbgear.Identity_Raw4{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_StakingRewards_Address320000000(fields []*registry.DecodedField) *pbgear.StakingRewards_Address32 {
	return &pbgear.StakingRewards_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Balances_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
	return &pbgear.Balances_TupleNull{}
}
func to_Utility_Origins0000000(fields []*registry.DecodedField) *pbgear.Utility_Origins {
	return &pbgear.Utility_Origins{
		Value_0: to_Utility_Value0(fields[0].Value),
	}
}

func to_Utility_DispatchAsCall0000000(fields []*registry.DecodedField) *pbgear.Utility_DispatchAsCall {
	return &pbgear.Utility_DispatchAsCall{
		AsOrigin: to_Utility_AsOrigin(fields[0].Value),
		Call:     to_oneof_Utility_Call(fields[1].Value),
	}
}

func to_ConvictionVoting_VoteCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_VoteCall {
	return &pbgear.ConvictionVoting_VoteCall{
		PollIndex: to_ConvictionVoting_CompactUint32(fields[0].Value),
		Vote:      to_ConvictionVoting_Vote(fields[1].Value),
	}
}

func to_Identity_Raw00000000(fields []*registry.DecodedField) *pbgear.Identity_Raw0 {
	return &pbgear.Identity_Raw0{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_SpNposElectionsSupport0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
	return &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{
		Total:  string(fields[0].Value),
		Voters: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(fields[1].Value),
	}
}

func to_NominationPools_MinCreateBond0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
	return &pbgear.NominationPools_MinCreateBond{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Balances_ForceUnreserveCall0000000(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
	return &pbgear.Balances_ForceUnreserveCall{
		Who:    to_Balances_Who(fields[0].Value),
		Amount: string(fields[1].Value),
	}
}

func to_Whitelist_SpWeightsWeightV2Weight0000000(fields []*registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
	return &pbgear.Whitelist_SpWeightsWeightV2Weight{
		RefTime:   to_Whitelist_CompactUint64(fields[0].Value),
		ProofSize: to_Whitelist_CompactUint64(fields[1].Value),
	}
}

func to_Identity_Raw20000000(fields []*registry.DecodedField) *pbgear.Identity_Raw2 {
	return &pbgear.Identity_Raw2{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw140000000(fields []*registry.DecodedField) *pbgear.Identity_Raw14 {
	return &pbgear.Identity_Raw14{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw310000000(fields []*registry.DecodedField) *pbgear.Identity_Raw31 {
	return &pbgear.Identity_Raw31{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Utility_AsOrigin0000000(fields []*registry.DecodedField) *pbgear.Utility_AsOrigin {
	return &pbgear.Utility_AsOrigin{
		Value: to_oneof_Utility_Value(fields[0].Value),
	}
}

func to_FellowshipReferenda_System0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_System {
	return &pbgear.FellowshipReferenda_System{
		Value_0: to_FellowshipReferenda_Value0(fields[0].Value),
	}
}

func to_Proxy_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Proxy_SpCoreCryptoAccountId32 {
	return &pbgear.Proxy_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Bounties_ProposeCuratorCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_ProposeCuratorCall {
	return &pbgear.Bounties_ProposeCuratorCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
		Curator:  to_Bounties_Curator(fields[1].Value),
		Fee:      to_Bounties_CompactString(fields[2].Value),
	}
}

func to_Identity_Raw170000000(fields []*registry.DecodedField) *pbgear.Identity_Raw17 {
	return &pbgear.Identity_Raw17{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Identity_CompactUint32 {
	return &pbgear.Identity_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_NominationPools_CreateWithPoolIdCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_CreateWithPoolIdCall {
	return &pbgear.NominationPools_CreateWithPoolIdCall{
		Amount:    to_NominationPools_CompactString(fields[0].Value),
		Root:      to_NominationPools_Root(fields[1].Value),
		Nominator: to_NominationPools_Nominator(fields[2].Value),
		Bouncer:   to_NominationPools_Bouncer(fields[3].Value),
		PoolId:    uint32(fields[4].Value),
	}
}

func to_BagsList_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.BagsList_SpCoreCryptoAccountId32 {
	return &pbgear.BagsList_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_None0000000(fields []*registry.DecodedField) *pbgear.Staking_None {
	return &pbgear.Staking_None{}
}
func to_Staking_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Staking_CompactTupleNull {
	return &pbgear.Staking_CompactTupleNull{
		Value: to_Staking_TupleNull(fields[0].Value),
	}
}

func to_Referenda_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
	return &pbgear.Referenda_PrimitiveTypesH256{
		Hash: []uint32(fields[0].Value),
	}
}

func to_FellowshipReferenda_NudgeReferendumCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_NudgeReferendumCall {
	return &pbgear.FellowshipReferenda_NudgeReferendumCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Staking_Controller0000000(fields []*registry.DecodedField) *pbgear.Staking_Controller {
	return &pbgear.Staking_Controller{}
}
func to_Proxy_Raw0000000(fields []*registry.DecodedField) *pbgear.Proxy_Raw {
	return &pbgear.Proxy_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_Address320000000(fields []*registry.DecodedField) *pbgear.Staking_Address32 {
	return &pbgear.Staking_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_LowQuality0000000(fields []*registry.DecodedField) *pbgear.Identity_LowQuality {
	return &pbgear.Identity_LowQuality{}
}
func to_StakingRewards_From0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_From {
	return &pbgear.StakingRewards_From{
		Value: to_oneof_StakingRewards_Value(fields[0].Value),
	}
}

func to_GearVoucher_IssueCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_IssueCall {
	return &pbgear.GearVoucher_IssueCall{
		Spender:       to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value),
		Balance:       string(fields[1].Value),
		Programs:      to_optional_GearVoucher_BTreeSet(fields[2].Value),
		CodeUploading: bool(fields[3].Value),
		Duration:      uint32(fields[4].Value),
	}
}

func to_System_KillPrefixCall0000000(fields []*registry.DecodedField) *pbgear.System_KillPrefixCall {
	return &pbgear.System_KillPrefixCall{
		Prefix:  []uint32(fields[0].Value),
		Subkeys: uint32(fields[1].Value),
	}
}

func to_Vesting_VestCall0000000(fields []*registry.DecodedField) *pbgear.Vesting_VestCall {
	return &pbgear.Vesting_VestCall{}
}
func to_Staking_ForceNewEraCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraCall {
	return &pbgear.Staking_ForceNewEraCall{}
}
func to_Identity_PalletIdentitySimpleIdentityInfo0000000(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
	return &pbgear.Identity_PalletIdentitySimpleIdentityInfo{
		Additional:     to_Identity_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value),
		Display:        to_Identity_Display(fields[1].Value),
		Legal:          to_Identity_Legal(fields[2].Value),
		Web:            to_Identity_Web(fields[3].Value),
		Riot:           to_Identity_Riot(fields[4].Value),
		Email:          to_Identity_Email(fields[5].Value),
		PgpFingerprint: []uint32(fields[6].Value),
		Image:          to_Identity_Image(fields[7].Value),
		Twitter:        to_Identity_Twitter(fields[8].Value),
	}
}

func to_System_SystemKeysList0000000(fields []*registry.DecodedField) *pbgear.System_SystemKeysList {
	return &pbgear.System_SystemKeysList{
		Keys: []uint32(fields[0].Value),
	}
}

func to_BagsList_Id0000000(fields []*registry.DecodedField) *pbgear.BagsList_Id {
	return &pbgear.BagsList_Id{
		Value_0: to_BagsList_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Utility_Signed0000000(fields []*registry.DecodedField) *pbgear.Utility_Signed {
	return &pbgear.Utility_Signed{
		Value_0: to_Utility_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Identity_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Identity_TupleNull {
	return &pbgear.Identity_TupleNull{}
}
func to_Identity_Display0000000(fields []*registry.DecodedField) *pbgear.Identity_Display {
	return &pbgear.Identity_Display{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Proxy_ProxyAnnouncedCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_ProxyAnnouncedCall {
	return &pbgear.Proxy_ProxyAnnouncedCall{
		Delegate:       to_Proxy_Delegate(fields[0].Value),
		Real:           to_Proxy_Real(fields[1].Value),
		ForceProxyType: to_optional_Proxy_ForceProxyType(fields[2].Value),
		Call:           to_oneof_Proxy_Call(fields[3].Value),
	}
}

func to_Babe_Logs0000000(fields []*registry.DecodedField) *pbgear.Babe_Logs {
	return &pbgear.Babe_Logs{
		Value: to_oneof_Babe_Value(fields[0].Value),
	}
}

func to_Vesting_VestedTransferCall0000000(fields []*registry.DecodedField) *pbgear.Vesting_VestedTransferCall {
	return &pbgear.Vesting_VestedTransferCall{
		Target:   to_Vesting_Target(fields[0].Value),
		Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(fields[1].Value),
	}
}

func to_ConvictionVoting_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
	return &pbgear.ConvictionVoting_CompactTupleNull{
		Value: to_ConvictionVoting_TupleNull(fields[0].Value),
	}
}

func to_ConvictionVoting_Locked5X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked5X {
	return &pbgear.ConvictionVoting_Locked5X{}
}
func to_Referenda_NudgeReferendumCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_NudgeReferendumCall {
	return &pbgear.Referenda_NudgeReferendumCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Bounties_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Bounties_CompactUint32 {
	return &pbgear.Bounties_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_NominationPools_SetMetadataCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetMetadataCall {
	return &pbgear.NominationPools_SetMetadataCall{
		PoolId:   uint32(fields[0].Value),
		Metadata: []uint32(fields[1].Value),
	}
}

func to_Babe_PrimaryAndSecondaryVRFSlots0000000(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryVRFSlots {
	return &pbgear.Babe_PrimaryAndSecondaryVRFSlots{}
}
func to_Balances_ForceTransferCall0000000(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
	return &pbgear.Balances_ForceTransferCall{
		Source: to_Balances_Source(fields[0].Value),
		Dest:   to_Balances_Dest(fields[1].Value),
		Value:  to_Balances_CompactString(fields[2].Value),
	}
}

func to_Staking_Set0000000(fields []*registry.DecodedField) *pbgear.Staking_Set {
	return &pbgear.Staking_Set{
		Value_0: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value),
	}
}

func to_Identity_Index0000000(fields []*registry.DecodedField) *pbgear.Identity_Index {
	return &pbgear.Identity_Index{
		Value_0: to_Identity_CompactTupleNull(fields[0].Value),
	}
}

func to_Bounties_ProposeBountyCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_ProposeBountyCall {
	return &pbgear.Bounties_ProposeBountyCall{
		Value:       to_Bounties_CompactString(fields[0].Value),
		Description: []uint32(fields[1].Value),
	}
}

func to_Staking_SetPayeeCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetPayeeCall {
	return &pbgear.Staking_SetPayeeCall{
		Payee: to_Staking_Payee(fields[0].Value),
	}
}

func to_Referenda_KillCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_KillCall {
	return &pbgear.Referenda_KillCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
	return &pbgear.Identity_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
	return &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_MinCommission0000000(fields []*registry.DecodedField) *pbgear.Staking_MinCommission {
	return &pbgear.Staking_MinCommission{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Treasury_Id0000000(fields []*registry.DecodedField) *pbgear.Treasury_Id {
	return &pbgear.Treasury_Id{
		Value_0: to_Treasury_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_ConvictionVoting_UnlockCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UnlockCall {
	return &pbgear.ConvictionVoting_UnlockCall{
		Class:  uint32(fields[0].Value),
		Target: to_ConvictionVoting_Target(fields[1].Value),
	}
}

func to_Identity_Web0000000(fields []*registry.DecodedField) *pbgear.Identity_Web {
	return &pbgear.Identity_Web{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Proxy_Index0000000(fields []*registry.DecodedField) *pbgear.Proxy_Index {
	return &pbgear.Proxy_Index{
		Value_0: to_Proxy_CompactTupleNull(fields[0].Value),
	}
}

func to_Bounties_Address320000000(fields []*registry.DecodedField) *pbgear.Bounties_Address32 {
	return &pbgear.Bounties_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_Nominator0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Nominator {
	return &pbgear.NominationPools_Nominator{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Babe_ReportEquivocationCall0000000(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationCall {
	return &pbgear.Babe_ReportEquivocationCall{
		EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value),
		KeyOwnerProof:     to_Babe_SpSessionMembershipProof(fields[1].Value),
	}
}

func to_Staking_Payee0000000(fields []*registry.DecodedField) *pbgear.Staking_Payee {
	return &pbgear.Staking_Payee{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Utility_VaraRuntimeRuntimeCall0000000(fields []*registry.DecodedField) *pbgear.Utility_VaraRuntimeRuntimeCall {
	return &pbgear.Utility_VaraRuntimeRuntimeCall{
		Calls: to_oneof_Utility_Calls(fields[0].Value),
	}
}

func to_oneof_Utility_Calls(in *registry.DecodedField) *pbgear.Utility_Calls {
	return nil
}

func to_Identity_None0000000(fields []*registry.DecodedField) *pbgear.Identity_None {
	return &pbgear.Identity_None{}
}
func to_Identity_Raw130000000(fields []*registry.DecodedField) *pbgear.Identity_Raw13 {
	return &pbgear.Identity_Raw13{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Vesting_Target0000000(fields []*registry.DecodedField) *pbgear.Vesting_Target {
	return &pbgear.Vesting_Target{
		Value: to_oneof_Vesting_Value(fields[0].Value),
	}
}

func to_Staking_Remove0000000(fields []*registry.DecodedField) *pbgear.Staking_Remove {
	return &pbgear.Staking_Remove{}
}
func to_Identity_Raw240000000(fields []*registry.DecodedField) *pbgear.Identity_Raw24 {
	return &pbgear.Identity_Raw24{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_WithdrawUnbondedCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_WithdrawUnbondedCall {
	return &pbgear.NominationPools_WithdrawUnbondedCall{
		MemberAccount:    to_NominationPools_MemberAccount(fields[0].Value),
		NumSlashingSpans: uint32(fields[1].Value),
	}
}

func to_ImOnline_PalletImOnlineHeartbeat0000000(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
	return &pbgear.ImOnline_PalletImOnlineHeartbeat{
		BlockNumber:    uint32(fields[0].Value),
		SessionIndex:   uint32(fields[1].Value),
		AuthorityIndex: uint32(fields[2].Value),
		ValidatorsLen:  uint32(fields[3].Value),
	}
}

func to_Proxy_Staking0000000(fields []*registry.DecodedField) *pbgear.Proxy_Staking {
	return &pbgear.Proxy_Staking{}
}
func to_Multisig_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Multisig_SpCoreCryptoAccountId32 {
	return &pbgear.Multisig_SpCoreCryptoAccountId32{
		OtherSignatories: to_Multisig_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_NominationPools_Raw0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Raw {
	return &pbgear.NominationPools_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Referenda_ProposalOrigin0000000(fields []*registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
	return &pbgear.Referenda_ProposalOrigin{
		Value: to_oneof_Referenda_Value(fields[0].Value),
	}
}

func to_Identity_Raw90000000(fields []*registry.DecodedField) *pbgear.Identity_Raw9 {
	return &pbgear.Identity_Raw9{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_CancelRequestCall0000000(fields []*registry.DecodedField) *pbgear.Identity_CancelRequestCall {
	return &pbgear.Identity_CancelRequestCall{
		RegIndex: uint32(fields[0].Value),
	}
}

func to_Staking_Stash0000000(fields []*registry.DecodedField) *pbgear.Staking_Stash {
	return &pbgear.Staking_Stash{}
}
func to_Utility_BatchCall0000000(fields []*registry.DecodedField) *pbgear.Utility_BatchCall {
	return &pbgear.Utility_BatchCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value),
	}
}

func to_Identity_Raw280000000(fields []*registry.DecodedField) *pbgear.Identity_Raw28 {
	return &pbgear.Identity_Raw28{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Multisig_AsMultiCall0000000(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiCall {
	return &pbgear.Multisig_AsMultiCall{
		Threshold:        uint32(fields[0].Value),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value),
		MaybeTimepoint:   to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value),
		Call:             to_oneof_Multisig_Call(fields[3].Value),
		MaxWeight:        to_Multisig_SpWeightsWeightV2Weight(fields[4].Value),
	}
}

func to_oneof_Multisig_Call(in *registry.DecodedField) *pbgear.Multisig_Call {
	return nil
}

func to_ElectionProviderMultiPhase_CompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_Treasury_ProposeSpendCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_ProposeSpendCall {
	return &pbgear.Treasury_ProposeSpendCall{
		Value:       to_Treasury_CompactString(fields[0].Value),
		Beneficiary: to_Treasury_Beneficiary(fields[1].Value),
	}
}

func to_ConvictionVoting_Standard0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Standard {
	return &pbgear.ConvictionVoting_Standard{
		Vote:    to_ConvictionVoting_PalletConvictionVotingVoteVote(fields[0].Value),
		Balance: string(fields[1].Value),
	}
}

func to_ConvictionVoting_Raw0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Raw {
	return &pbgear.ConvictionVoting_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Referenda_BoundedCollectionsBoundedVecBoundedVec0000000(fields []*registry.DecodedField) *pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_Who0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Who {
	return &pbgear.FellowshipCollective_Who{
		Value: to_oneof_FellowshipCollective_Value(fields[0].Value),
	}
}

func to_oneof_FellowshipCollective_Value(in *registry.DecodedField) *pbgear.FellowshipCollective_Value {
	return nil
}

func to_System_SetCodeCall0000000(fields []*registry.DecodedField) *pbgear.System_SetCodeCall {
	return &pbgear.System_SetCodeCall{
		Code: []uint32(fields[0].Value),
	}
}

func to_ChildBounties_AcceptCuratorCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_AcceptCuratorCall {
	return &pbgear.ChildBounties_AcceptCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
	}
}

func to_NominationPools_Blocked0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Blocked {
	return &pbgear.NominationPools_Blocked{}
}
func to_FellowshipReferenda_ProposalOrigin0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
	return &pbgear.FellowshipReferenda_ProposalOrigin{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value),
	}
}

func to_Identity_Value10000000(fields []*registry.DecodedField) *pbgear.Identity_Value1 {
	return &pbgear.Identity_Value1{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Identity_PalletIdentityTypesBitFlags0000000(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
	return &pbgear.Identity_PalletIdentityTypesBitFlags{
		Fields: uint64(fields[0].Value),
	}
}

func to_Timestamp_SetCall0000000(fields []*registry.DecodedField) *pbgear.Timestamp_SetCall {
	return &pbgear.Timestamp_SetCall{
		Now: to_Timestamp_CompactUint64(fields[0].Value),
	}
}

func to_Babe_SpConsensusBabeAppPublic0000000(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
	return &pbgear.Babe_SpConsensusBabeAppPublic{
		Offender: to_Babe_SpCoreSr25519Public(fields[0].Value),
	}
}

func to_Session_SpConsensusBabeAppPublic0000000(fields []*registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
	return &pbgear.Session_SpConsensusBabeAppPublic{
		Babe: to_Session_SpCoreSr25519Public(fields[0].Value),
	}
}

func to_FellowshipCollective_Address200000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address20 {
	return &pbgear.FellowshipCollective_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_RejectAnnouncementCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_RejectAnnouncementCall {
	return &pbgear.Proxy_RejectAnnouncementCall{
		Delegate: to_Proxy_Delegate(fields[0].Value),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value),
	}
}

func to_NominationPools_Extra0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Extra {
	return &pbgear.NominationPools_Extra{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_Bouncer0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Bouncer {
	return &pbgear.NominationPools_Bouncer{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_SetCommissionCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionCall {
	return &pbgear.NominationPools_SetCommissionCall{
		PoolId:        uint32(fields[0].Value),
		NewCommission: to_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(fields[1].Value),
	}
}

func to_Session_SpCoreEd25519Public0000000(fields []*registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
	return &pbgear.Session_SpCoreEd25519Public{
		Grandpa: []uint32(fields[0].Value),
	}
}

func to_FellowshipReferenda_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
	return &pbgear.FellowshipReferenda_PrimitiveTypesH256{
		Hash: []uint32(fields[0].Value),
	}
}

func to_Proxy_AddProxyCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_AddProxyCall {
	return &pbgear.Proxy_AddProxyCall{
		Delegate:  to_Proxy_Delegate(fields[0].Value),
		ProxyType: to_Proxy_ProxyType(fields[1].Value),
		Delay:     uint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32CompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value),
	}
}

func to_ChildBounties_UnassignCuratorCall0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_UnassignCuratorCall {
	return &pbgear.ChildBounties_UnassignCuratorCall{
		ParentBountyId: to_ChildBounties_CompactUint32(fields[0].Value),
		ChildBountyId:  to_ChildBounties_CompactUint32(fields[1].Value),
	}
}

func to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw150000000(fields []*registry.DecodedField) *pbgear.Identity_Raw15 {
	return &pbgear.Identity_Raw15{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw210000000(fields []*registry.DecodedField) *pbgear.Identity_Raw21 {
	return &pbgear.Identity_Raw21{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_Spawner0000000(fields []*registry.DecodedField) *pbgear.Proxy_Spawner {
	return &pbgear.Proxy_Spawner{
		Value: to_oneof_Proxy_Value(fields[0].Value),
	}
}

func to_Multisig_ApproveAsMultiCall0000000(fields []*registry.DecodedField) *pbgear.Multisig_ApproveAsMultiCall {
	return &pbgear.Multisig_ApproveAsMultiCall{
		Threshold:        uint32(fields[0].Value),
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[1].Value),
		MaybeTimepoint:   to_optional_Multisig_PalletMultisigTimepoint(fields[2].Value),
		CallHash:         []uint32(fields[3].Value),
		MaxWeight:        to_Multisig_SpWeightsWeightV2Weight(fields[4].Value),
	}
}

func to_Staking_CancelDeferredSlashCall0000000(fields []*registry.DecodedField) *pbgear.Staking_CancelDeferredSlashCall {
	return &pbgear.Staking_CancelDeferredSlashCall{
		Era:          uint32(fields[0].Value),
		SlashIndices: []uint32(fields[1].Value),
	}
}

func to_ConvictionVoting_Locked4X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked4X {
	return &pbgear.ConvictionVoting_Locked4X{}
}
func to_Identity_Raw230000000(fields []*registry.DecodedField) *pbgear.Identity_Raw23 {
	return &pbgear.Identity_Raw23{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Bounties_ApproveBountyCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_ApproveBountyCall {
	return &pbgear.Bounties_ApproveBountyCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
	}
}

func to_Gear_ClaimValueCall0000000(fields []*registry.DecodedField) *pbgear.Gear_ClaimValueCall {
	return &pbgear.Gear_ClaimValueCall{
		MessageId: to_Gear_GprimitivesMessageId(fields[0].Value),
	}
}

func to_StakingRewards_RefillCall0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_RefillCall {
	return &pbgear.StakingRewards_RefillCall{
		Value: string(fields[0].Value),
	}
}

func to_Vesting_Address200000000(fields []*registry.DecodedField) *pbgear.Vesting_Address20 {
	return &pbgear.Vesting_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_NominateCall0000000(fields []*registry.DecodedField) *pbgear.Staking_NominateCall {
	return &pbgear.Staking_NominateCall{
		Targets: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(fields[0].Value),
	}
}

func to_Staking_ChillCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ChillCall {
	return &pbgear.Staking_ChillCall{}
}
func to_Identity_Judgement0000000(fields []*registry.DecodedField) *pbgear.Identity_Judgement {
	return &pbgear.Identity_Judgement{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_NominationPools_MaxPools0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MaxPools {
	return &pbgear.NominationPools_MaxPools{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_NominationPools_Permissioned0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Permissioned {
	return &pbgear.NominationPools_Permissioned{}
}
func to_Grandpa_Precommit0000000(fields []*registry.DecodedField) *pbgear.Grandpa_Precommit {
	return &pbgear.Grandpa_Precommit{
		Value_0: to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value),
	}
}

func to_Staking_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Staking_CompactUint32 {
	return &pbgear.Staking_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_ConvictionVoting_TupleNull0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
	return &pbgear.ConvictionVoting_TupleNull{}
}
func to_Identity_Account0000000(fields []*registry.DecodedField) *pbgear.Identity_Account {
	return &pbgear.Identity_Account{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Identity_Data0000000(fields []*registry.DecodedField) *pbgear.Identity_Data {
	return &pbgear.Identity_Data{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Staking_BondExtraCall0000000(fields []*registry.DecodedField) *pbgear.Staking_BondExtraCall {
	return &pbgear.Staking_BondExtraCall{
		MaxAdditional: to_Staking_CompactString(fields[0].Value),
	}
}

func to_Staking_Address200000000(fields []*registry.DecodedField) *pbgear.Staking_Address20 {
	return &pbgear.Staking_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ConvictionVoting_PalletConvictionVotingVoteVote0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
	return &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{
		Vote: uint32(fields[0].Value),
	}
}

func to_Identity_RequestJudgementCall0000000(fields []*registry.DecodedField) *pbgear.Identity_RequestJudgementCall {
	return &pbgear.Identity_RequestJudgementCall{
		RegIndex: to_Identity_CompactUint32(fields[0].Value),
		MaxFee:   to_Identity_CompactString(fields[1].Value),
	}
}

func to_Proxy_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Proxy_CompactUint32 {
	return &pbgear.Proxy_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_Bounties_Beneficiary0000000(fields []*registry.DecodedField) *pbgear.Bounties_Beneficiary {
	return &pbgear.Bounties_Beneficiary{
		Value: to_oneof_Bounties_Value(fields[0].Value),
	}
}

func to_oneof_Bounties_Value(in *registry.DecodedField) *pbgear.Bounties_Value {
	return nil
}

func to_Balances_Address200000000(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
	return &pbgear.Balances_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Vesting_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Vesting_TupleNull {
	return &pbgear.Vesting_TupleNull{}
}
func to_Treasury_Address320000000(fields []*registry.DecodedField) *pbgear.Treasury_Address32 {
	return &pbgear.Treasury_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Referenda_Inline0000000(fields []*registry.DecodedField) *pbgear.Referenda_Inline {
	return &pbgear.Referenda_Inline{
		Value_0: to_Referenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value),
	}
}

func to_FellowshipReferenda_SetMetadataCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SetMetadataCall {
	return &pbgear.FellowshipReferenda_SetMetadataCall{
		Index:     uint32(fields[0].Value),
		MaybeHash: to_optional_FellowshipReferenda_PrimitiveTypesH256(fields[1].Value),
	}
}

func to_NominationPools_CreateCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_CreateCall {
	return &pbgear.NominationPools_CreateCall{
		Amount:    to_NominationPools_CompactString(fields[0].Value),
		Root:      to_NominationPools_Root(fields[1].Value),
		Nominator: to_NominationPools_Nominator(fields[2].Value),
		Bouncer:   to_NominationPools_Bouncer(fields[3].Value),
	}
}

func to_StakingRewards_Address200000000(fields []*registry.DecodedField) *pbgear.StakingRewards_Address20 {
	return &pbgear.StakingRewards_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Referenda_System0000000(fields []*registry.DecodedField) *pbgear.Referenda_System {
	return &pbgear.Referenda_System{
		Value_0: to_Referenda_Value0(fields[0].Value),
	}
}

func to_FellowshipReferenda_OneFewerDecidingCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
	return &pbgear.FellowshipReferenda_OneFewerDecidingCall{
		Track: uint32(fields[0].Value),
	}
}

func to_Scheduler_CancelNamedCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_CancelNamedCall {
	return &pbgear.Scheduler_CancelNamedCall{
		Id: []uint32(fields[0].Value),
	}
}

func to_Babe_ReportEquivocationUnsignedCall0000000(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationUnsignedCall {
	return &pbgear.Babe_ReportEquivocationUnsignedCall{
		EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(fields[0].Value),
		KeyOwnerProof:     to_Babe_SpSessionMembershipProof(fields[1].Value),
	}
}

func to_Vesting_Address320000000(fields []*registry.DecodedField) *pbgear.Vesting_Address32 {
	return &pbgear.Vesting_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Session_SpConsensusGrandpaAppPublic0000000(fields []*registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
	return &pbgear.Session_SpConsensusGrandpaAppPublic{
		Grandpa: to_Session_SpCoreEd25519Public(fields[0].Value),
	}
}

func to_ConvictionVoting_Index0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Index {
	return &pbgear.ConvictionVoting_Index{
		Value_0: to_ConvictionVoting_CompactTupleNull(fields[0].Value),
	}
}

func to_Referenda_Root0000000(fields []*registry.DecodedField) *pbgear.Referenda_Root {
	return &pbgear.Referenda_Root{}
}
func to_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData0000000(fields []*registry.DecodedField) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
	return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{
		Value_0: to_Identity_SpCoreCryptoAccountId32(fields[0].Value),
		Value_1: to_Identity_Value1(fields[1].Value),
	}
}

func to_Utility_Void0000000(fields []*registry.DecodedField) *pbgear.Utility_Void {
	return &pbgear.Utility_Void{
		Value_0: to_Utility_Value0(fields[0].Value),
	}
}

func to_ConvictionVoting_DelegateCall0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_DelegateCall {
	return &pbgear.ConvictionVoting_DelegateCall{
		Class:      uint32(fields[0].Value),
		To:         to_ConvictionVoting_To(fields[1].Value),
		Conviction: to_ConvictionVoting_Conviction(fields[2].Value),
		Balance:    string(fields[3].Value),
	}
}

func to_Referenda_At0000000(fields []*registry.DecodedField) *pbgear.Referenda_At {
	return &pbgear.Referenda_At{
		Value_0: uint32(fields[0].Value),
	}
}

func to_System_SetCodeWithoutChecksCall0000000(fields []*registry.DecodedField) *pbgear.System_SetCodeWithoutChecksCall {
	return &pbgear.System_SetCodeWithoutChecksCall{
		Code: []uint32(fields[0].Value),
	}
}

func to_Balances_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
	return &pbgear.Balances_CompactTupleNull{
		Value: to_Balances_TupleNull(fields[0].Value),
	}
}

func to_ImOnline_SpCoreSr25519Signature0000000(fields []*registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
	return &pbgear.ImOnline_SpCoreSr25519Signature{
		Signature: []uint32(fields[0].Value),
	}
}

func to_Staking_Noop0000000(fields []*registry.DecodedField) *pbgear.Staking_Noop {
	return &pbgear.Staking_Noop{}
}
func to_Utility_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Utility_SpCoreCryptoAccountId32 {
	return &pbgear.Utility_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw0000000(fields []*registry.DecodedField) *pbgear.Identity_Raw {
	return &pbgear.Identity_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_Address200000000(fields []*registry.DecodedField) *pbgear.Proxy_Address20 {
	return &pbgear.Proxy_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_Vesting_Id0000000(fields []*registry.DecodedField) *pbgear.Vesting_Id {
	return &pbgear.Vesting_Id{
		Value_0: to_Vesting_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Staking_ForceApplyMinCommissionCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ForceApplyMinCommissionCall {
	return &pbgear.Staking_ForceApplyMinCommissionCall{
		ValidatorStash: to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Treasury_RejectProposalCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_RejectProposalCall {
	return &pbgear.Treasury_RejectProposalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value),
	}
}

func to_Identity_Unknown0000000(fields []*registry.DecodedField) *pbgear.Identity_Unknown {
	return &pbgear.Identity_Unknown{}
}
func to_Proxy_ProxyType0000000(fields []*registry.DecodedField) *pbgear.Proxy_ProxyType {
	return &pbgear.Proxy_ProxyType{
		Value: to_oneof_Proxy_Value(fields[0].Value),
	}
}

func to_System_RemarkWithEventCall0000000(fields []*registry.DecodedField) *pbgear.System_RemarkWithEventCall {
	return &pbgear.System_RemarkWithEventCall{
		Remark: []uint32(fields[0].Value),
	}
}

func to_Session_PurgeKeysCall0000000(fields []*registry.DecodedField) *pbgear.Session_PurgeKeysCall {
	return &pbgear.Session_PurgeKeysCall{}
}
func to_Treasury_Index0000000(fields []*registry.DecodedField) *pbgear.Treasury_Index {
	return &pbgear.Treasury_Index{
		Value_0: to_Treasury_CompactTupleNull(fields[0].Value),
	}
}

func to_FellowshipReferenda_RefundSubmissionDepositCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
	return &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Bounties_ExtendBountyExpiryCall0000000(fields []*registry.DecodedField) *pbgear.Bounties_ExtendBountyExpiryCall {
	return &pbgear.Bounties_ExtendBountyExpiryCall{
		BountyId: to_Bounties_CompactUint32(fields[0].Value),
		Remark:   []uint32(fields[1].Value),
	}
}

func to_Identity_SetIdentityCall0000000(fields []*registry.DecodedField) *pbgear.Identity_SetIdentityCall {
	return &pbgear.Identity_SetIdentityCall{
		Info: to_Identity_PalletIdentitySimpleIdentityInfo(fields[0].Value),
	}
}

func to_Babe_SpConsensusSlotsSlot0000000(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
	return &pbgear.Babe_SpConsensusSlotsSlot{
		Slot: uint64(fields[0].Value),
	}
}

func to_Staking_ValidateCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ValidateCall {
	return &pbgear.Staking_ValidateCall{
		Prefs: to_Staking_PalletStakingValidatorPrefs(fields[0].Value),
	}
}

func to_Staking_SpArithmeticPerThingsPercent0000000(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
	return &pbgear.Staking_SpArithmeticPerThingsPercent{
		Factor: uint32(fields[0].Value),
	}
}

func to_Session_SpCoreSr25519Public0000000(fields []*registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
	return &pbgear.Session_SpCoreSr25519Public{
		Babe: []uint32(fields[0].Value),
	}
}

func to_Identity_Image0000000(fields []*registry.DecodedField) *pbgear.Identity_Image {
	return &pbgear.Identity_Image{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Staking_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
	return &pbgear.Staking_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_ConvictionVoting_Target0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Target {
	return &pbgear.ConvictionVoting_Target{
		Value: to_oneof_ConvictionVoting_Value(fields[0].Value),
	}
}

func to_FellowshipReferenda_At0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_At {
	return &pbgear.FellowshipReferenda_At{
		Value_0: uint32(fields[0].Value),
	}
}

func to_Staking_SetMinCommissionCall0000000(fields []*registry.DecodedField) *pbgear.Staking_SetMinCommissionCall {
	return &pbgear.Staking_SetMinCommissionCall{
		New: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value),
	}
}

func to_Referenda_Legacy0000000(fields []*registry.DecodedField) *pbgear.Referenda_Legacy {
	return &pbgear.Referenda_Legacy{
		Hash: to_Referenda_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Whitelist_WhitelistCallCall0000000(fields []*registry.DecodedField) *pbgear.Whitelist_WhitelistCallCall {
	return &pbgear.Whitelist_WhitelistCallCall{
		CallHash: to_Whitelist_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Scheduler_ScheduleNamedAfterCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedAfterCall {
	return &pbgear.Scheduler_ScheduleNamedAfterCall{
		Id:            []uint32(fields[0].Value),
		After:         uint32(fields[1].Value),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[2].Value),
		Priority:      uint32(fields[3].Value),
		Call:          to_oneof_Scheduler_Call(fields[4].Value),
	}
}

func to_Whitelist_RemoveWhitelistedCallCall0000000(fields []*registry.DecodedField) *pbgear.Whitelist_RemoveWhitelistedCallCall {
	return &pbgear.Whitelist_RemoveWhitelistedCallCall{
		CallHash: to_Whitelist_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Identity_Raw110000000(fields []*registry.DecodedField) *pbgear.Identity_Raw11 {
	return &pbgear.Identity_Raw11{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_ClaimPayoutCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutCall {
	return &pbgear.NominationPools_ClaimPayoutCall{}
}
func to_GearVoucher_SendMessage0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_SendMessage {
	return &pbgear.GearVoucher_SendMessage{
		Destination: to_GearVoucher_GprimitivesActorId(fields[0].Value),
		Payload:     []uint32(fields[1].Value),
		GasLimit:    uint64(fields[2].Value),
		Value:       string(fields[3].Value),
		KeepAlive:   bool(fields[4].Value),
	}
}

func to_GearVoucher_RevokeCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_RevokeCall {
	return &pbgear.GearVoucher_RevokeCall{
		Spender:   to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value),
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value),
	}
}

func to_GearVoucher_AppendPrograms0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
	return &pbgear.GearVoucher_AppendPrograms{
		Value: to_oneof_GearVoucher_Value(fields[0].Value),
	}
}

func to_Staking_SpRuntimeMultiaddressMultiAddress0000000(fields []*registry.DecodedField) *pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
	return &pbgear.Staking_SpRuntimeMultiaddressMultiAddress{
		Targets: to_Staking_Targets(fields[0].Value),
	}
}

func to_FellowshipCollective_VoteCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_VoteCall {
	return &pbgear.FellowshipCollective_VoteCall{
		Poll: uint32(fields[0].Value),
		Aye:  bool(fields[1].Value),
	}
}

func to_Identity_KnownGood0000000(fields []*registry.DecodedField) *pbgear.Identity_KnownGood {
	return &pbgear.Identity_KnownGood{}
}
func to_ElectionProviderMultiPhase_SetEmergencyElectionResultCall0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
	return &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{
		Supports: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(fields[0].Value),
	}
}

func to_Bounties_Curator0000000(fields []*registry.DecodedField) *pbgear.Bounties_Curator {
	return &pbgear.Bounties_Curator{
		Value: to_oneof_Bounties_Value(fields[0].Value),
	}
}

func to_FellowshipReferenda_After0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_After {
	return &pbgear.FellowshipReferenda_After{
		Value_0: uint32(fields[0].Value),
	}
}

func to_Identity_Raw50000000(fields []*registry.DecodedField) *pbgear.Identity_Raw5 {
	return &pbgear.Identity_Raw5{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Bounties_Address200000000(fields []*registry.DecodedField) *pbgear.Bounties_Address20 {
	return &pbgear.Bounties_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Gear_SendMessageCall0000000(fields []*registry.DecodedField) *pbgear.Gear_SendMessageCall {
	return &pbgear.Gear_SendMessageCall{
		Destination: to_Gear_GprimitivesActorId(fields[0].Value),
		Payload:     []uint32(fields[1].Value),
		GasLimit:    uint64(fields[2].Value),
		Value:       string(fields[3].Value),
		KeepAlive:   bool(fields[4].Value),
	}
}

func to_Balances_TransferAllowDeathCall0000000(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
	return &pbgear.Balances_TransferAllowDeathCall{
		Dest:  to_Balances_Dest(fields[0].Value),
		Value: to_Balances_CompactString(fields[1].Value),
	}
}

func to_Staking_RebondCall0000000(fields []*registry.DecodedField) *pbgear.Staking_RebondCall {
	return &pbgear.Staking_RebondCall{
		Value: to_Staking_CompactString(fields[0].Value),
	}
}

func to_Treasury_ApproveProposalCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_ApproveProposalCall {
	return &pbgear.Treasury_ApproveProposalCall{
		ProposalId: to_Treasury_CompactUint32(fields[0].Value),
	}
}

func to_Gear_GprimitivesActorId0000000(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
	return &pbgear.Gear_GprimitivesActorId{
		Destination: []uint32(fields[0].Value),
	}
}

func to_GearVoucher_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
	return &pbgear.GearVoucher_SpCoreCryptoAccountId32{
		Spender: []uint32(fields[0].Value),
	}
}

func to_Staking_BondCall0000000(fields []*registry.DecodedField) *pbgear.Staking_BondCall {
	return &pbgear.Staking_BondCall{
		Value: to_Staking_CompactString(fields[0].Value),
		Payee: to_Staking_Payee(fields[1].Value),
	}
}

func to_Identity_Value00000000(fields []*registry.DecodedField) *pbgear.Identity_Value0 {
	return &pbgear.Identity_Value0{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Identity_Legal0000000(fields []*registry.DecodedField) *pbgear.Identity_Legal {
	return &pbgear.Identity_Legal{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Bounties_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
	return &pbgear.Bounties_CompactTupleNull{
		Value: to_Bounties_TupleNull(fields[0].Value),
	}
}

func to_NominationPools_NewNominator0000000(fields []*registry.DecodedField) *pbgear.NominationPools_NewNominator {
	return &pbgear.NominationPools_NewNominator{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Referenda_RefundSubmissionDepositCall0000000(fields []*registry.DecodedField) *pbgear.Referenda_RefundSubmissionDepositCall {
	return &pbgear.Referenda_RefundSubmissionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Identity_Riot0000000(fields []*registry.DecodedField) *pbgear.Identity_Riot {
	return &pbgear.Identity_Riot{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_NominationPools_Member0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Member {
	return &pbgear.NominationPools_Member{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Babe_Seal0000000(fields []*registry.DecodedField) *pbgear.Babe_Seal {
	return &pbgear.Babe_Seal{
		Value_0: []uint32(fields[0].Value),
		Value_1: []uint32(fields[1].Value),
	}
}

func to_Grandpa_ReportEquivocationCall0000000(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationCall {
	return &pbgear.Grandpa_ReportEquivocationCall{
		EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(fields[0].Value),
		KeyOwnerProof:     to_Grandpa_SpSessionMembershipProof(fields[1].Value),
	}
}

func to_FellowshipCollective_DemoteMemberCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_DemoteMemberCall {
	return &pbgear.FellowshipCollective_DemoteMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value),
	}
}

func to_Identity_Keccak2560000000(fields []*registry.DecodedField) *pbgear.Identity_Keccak256 {
	return &pbgear.Identity_Keccak256{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_PalletNominationPoolsCommissionChangeRate0000000(fields []*registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
	return &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{
		MaxIncrease: to_NominationPools_SpArithmeticPerThingsPerbill(fields[0].Value),
		MinDelay:    uint32(fields[1].Value),
	}
}

func to_Staking_ForceUnstakeCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ForceUnstakeCall {
	return &pbgear.Staking_ForceUnstakeCall{
		Stash:            to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
		NumSlashingSpans: uint32(fields[1].Value),
	}
}

func to_ConvictionVoting_Locked3X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked3X {
	return &pbgear.ConvictionVoting_Locked3X{}
}
func to_ChildBounties_CompactUint320000000(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
	return &pbgear.ChildBounties_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_GearVoucher_CallCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_CallCall {
	return &pbgear.GearVoucher_CallCall{
		VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[0].Value),
		Call:      to_GearVoucher_Call(fields[1].Value),
	}
}

func to_GearVoucher_None0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_None {
	return &pbgear.GearVoucher_None{}
}
func to_Staking_SpArithmeticPerThingsPerbill0000000(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
	return &pbgear.Staking_SpArithmeticPerThingsPerbill{
		Value: uint32(fields[0].Value),
	}
}

func to_Staking_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Staking_TupleNull {
	return &pbgear.Staking_TupleNull{}
}
func to_Identity_Email0000000(fields []*registry.DecodedField) *pbgear.Identity_Email {
	return &pbgear.Identity_Email{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Multisig_AsMultiThreshold1Call0000000(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiThreshold1Call {
	return &pbgear.Multisig_AsMultiThreshold1Call{
		OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(fields[0].Value),
		Call:             to_oneof_Multisig_Call(fields[1].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_BagsList_Address200000000(fields []*registry.DecodedField) *pbgear.BagsList_Address20 {
	return &pbgear.BagsList_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw180000000(fields []*registry.DecodedField) *pbgear.Identity_Raw18 {
	return &pbgear.Identity_Raw18{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Raw200000000(fields []*registry.DecodedField) *pbgear.Identity_Raw20 {
	return &pbgear.Identity_Raw20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_TupleNull0000000(fields []*registry.DecodedField) *pbgear.NominationPools_TupleNull {
	return &pbgear.NominationPools_TupleNull{}
}
func to_Treasury_CompactUint320000000(fields []*registry.DecodedField) *pbgear.Treasury_CompactUint32 {
	return &pbgear.Treasury_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_FellowshipCollective_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_SpCoreCryptoAccountId32 {
	return &pbgear.FellowshipCollective_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Whitelist_CompactUint640000000(fields []*registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
	return &pbgear.Whitelist_CompactUint64{
		Value: uint64(fields[0].Value),
	}
}

func to_Whitelist_DispatchWhitelistedCallCall0000000(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallCall {
	return &pbgear.Whitelist_DispatchWhitelistedCallCall{
		CallHash:          to_Whitelist_PrimitiveTypesH256(fields[0].Value),
		CallEncodedLen:    uint32(fields[1].Value),
		CallWeightWitness: to_Whitelist_SpWeightsWeightV2Weight(fields[2].Value),
	}
}

func to_Preimage_UnnotePreimageCall0000000(fields []*registry.DecodedField) *pbgear.Preimage_UnnotePreimageCall {
	return &pbgear.Preimage_UnnotePreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Staking_CompactSpArithmeticPerThingsPerbill0000000(fields []*registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
	return &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{
		Value: to_Staking_SpArithmeticPerThingsPerbill(fields[0].Value),
	}
}

func to_NominationPools_UnbondCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_UnbondCall {
	return &pbgear.NominationPools_UnbondCall{
		MemberAccount:   to_NominationPools_MemberAccount(fields[0].Value),
		UnbondingPoints: to_NominationPools_CompactString(fields[1].Value),
	}
}

func to_NominationPools_PermissionlessCompound0000000(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessCompound {
	return &pbgear.NominationPools_PermissionlessCompound{}
}
func to_Staking_Raw0000000(fields []*registry.DecodedField) *pbgear.Staking_Raw {
	return &pbgear.Staking_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Staking_ChillOtherCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ChillOtherCall {
	return &pbgear.Staking_ChillOtherCall{
		Controller: to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Proxy_RemoveProxiesCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxiesCall {
	return &pbgear.Proxy_RemoveProxiesCall{}
}
func to_Scheduler_ScheduleNamedCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedCall {
	return &pbgear.Scheduler_ScheduleNamedCall{
		Id:            []uint32(fields[0].Value),
		When:          uint32(fields[1].Value),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[2].Value),
		Priority:      uint32(fields[3].Value),
		Call:          to_oneof_Scheduler_Call(fields[4].Value),
	}
}

func to_ChildBounties_Beneficiary0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
	return &pbgear.ChildBounties_Beneficiary{
		Value: to_oneof_ChildBounties_Value(fields[0].Value),
	}
}

func to_System_RemarkCall0000000(fields []*registry.DecodedField) *pbgear.System_RemarkCall {
	return &pbgear.System_RemarkCall{
		Remark: []uint32(fields[0].Value),
	}
}

func to_Babe_RuntimeEnvironmentUpdated0000000(fields []*registry.DecodedField) *pbgear.Babe_RuntimeEnvironmentUpdated {
	return &pbgear.Babe_RuntimeEnvironmentUpdated{}
}
func to_Grandpa_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Grandpa_PrimitiveTypesH256 {
	return &pbgear.Grandpa_PrimitiveTypesH256{
		TargetHash: []uint32(fields[0].Value),
	}
}

func to_Balances_TransferAllCall0000000(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
	return &pbgear.Balances_TransferAllCall{
		Dest:      to_Balances_Dest(fields[0].Value),
		KeepAlive: bool(fields[1].Value),
	}
}

func to_Staking_ScaleValidatorCountCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ScaleValidatorCountCall {
	return &pbgear.Staking_ScaleValidatorCountCall{
		Factor: to_Staking_SpArithmeticPerThingsPercent(fields[0].Value),
	}
}

func to_Babe_TupleUint64Uint640000000(fields []*registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
	return &pbgear.Babe_TupleUint64Uint64{
		Value_0: uint64(fields[0].Value),
		Value_1: uint64(fields[1].Value),
	}
}

func to_Vesting_Raw0000000(fields []*registry.DecodedField) *pbgear.Vesting_Raw {
	return &pbgear.Vesting_Raw{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Twitter0000000(fields []*registry.DecodedField) *pbgear.Identity_Twitter {
	return &pbgear.Identity_Twitter{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
	return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{
		Solution: to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(fields[0].Value),
		Score:    to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(fields[1].Value),
		Round:    uint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_SubmitUnsignedCall0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
	return &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{
		RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(fields[0].Value),
		Witness:     to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(fields[1].Value),
	}
}

func to_NominationPools_State0000000(fields []*registry.DecodedField) *pbgear.NominationPools_State {
	return &pbgear.NominationPools_State{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_GearVoucher_UpdateCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_UpdateCall {
	return &pbgear.GearVoucher_UpdateCall{
		Spender:         to_GearVoucher_SpCoreCryptoAccountId32(fields[0].Value),
		VoucherId:       to_GearVoucher_PalletGearVoucherInternalVoucherId(fields[1].Value),
		MoveOwnership:   to_optional_GearVoucher_SpCoreCryptoAccountId32(fields[2].Value),
		BalanceTopUp:    string(fields[3].Value),
		AppendPrograms:  to_optional_GearVoucher_AppendPrograms(fields[4].Value),
		CodeUploading:   bool(fields[5].Value),
		ProlongDuration: uint32(fields[6].Value),
	}
}

func to_Referenda_Lookup0000000(fields []*registry.DecodedField) *pbgear.Referenda_Lookup {
	return &pbgear.Referenda_Lookup{
		Hash: to_Referenda_PrimitiveTypesH256(fields[0].Value),
		Len:  uint32(fields[1].Value),
	}
}

func to_FellowshipReferenda_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SpCoreCryptoAccountId32 {
	return &pbgear.FellowshipReferenda_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_FellowshipReferenda_PlaceDecisionDepositCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
	return &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{
		Index: uint32(fields[0].Value),
	}
}

func to_Babe_SpSessionMembershipProof0000000(fields []*registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
	return &pbgear.Babe_SpSessionMembershipProof{
		Session:        uint32(fields[0].Value),
		TrieNodes:      to_repeated_Babe_BabeTrieNodesList(fields[1].Value),
		ValidatorCount: uint32(fields[2].Value),
	}
}

func to_BagsList_PutInFrontOfCall0000000(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfCall {
	return &pbgear.BagsList_PutInFrontOfCall{
		Lighter: to_BagsList_Lighter(fields[0].Value),
	}
}

func to_Treasury_CompactString0000000(fields []*registry.DecodedField) *pbgear.Treasury_CompactString {
	return &pbgear.Treasury_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Referenda_Signed0000000(fields []*registry.DecodedField) *pbgear.Referenda_Signed {
	return &pbgear.Referenda_Signed{
		Value_0: to_Referenda_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Referenda_Void0000000(fields []*registry.DecodedField) *pbgear.Referenda_Void {
	return &pbgear.Referenda_Void{
		Value_0: to_Referenda_Value0(fields[0].Value),
	}
}

func to_Multisig_CompactUint640000000(fields []*registry.DecodedField) *pbgear.Multisig_CompactUint64 {
	return &pbgear.Multisig_CompactUint64{
		Value: uint64(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
	return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{
		Voters:  to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Targets: to_ElectionProviderMultiPhase_CompactUint32(fields[1].Value),
	}
}

func to_NominationPools_UpdateRolesCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_UpdateRolesCall {
	return &pbgear.NominationPools_UpdateRolesCall{
		PoolId:       uint32(fields[0].Value),
		NewRoot:      to_NominationPools_NewRoot(fields[1].Value),
		NewNominator: to_NominationPools_NewNominator(fields[2].Value),
		NewBouncer:   to_NominationPools_NewBouncer(fields[3].Value),
	}
}

func to_Babe_V10000000(fields []*registry.DecodedField) *pbgear.Babe_V1 {
	return &pbgear.Babe_V1{
		C:            to_Babe_TupleUint64Uint64(fields[0].Value),
		AllowedSlots: to_Babe_AllowedSlots(fields[1].Value),
	}
}

func to_Referenda_EnactmentMoment0000000(fields []*registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
	return &pbgear.Referenda_EnactmentMoment{
		Value: to_oneof_Referenda_Value(fields[0].Value),
	}
}

func to_Preimage_RequestPreimageCall0000000(fields []*registry.DecodedField) *pbgear.Preimage_RequestPreimageCall {
	return &pbgear.Preimage_RequestPreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Proxy_RemoveProxyCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxyCall {
	return &pbgear.Proxy_RemoveProxyCall{
		Delegate:  to_Proxy_Delegate(fields[0].Value),
		ProxyType: to_Proxy_ProxyType(fields[1].Value),
		Delay:     uint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ConvictionVoting_None0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_None {
	return &pbgear.ConvictionVoting_None{}
}
func to_Referenda_After0000000(fields []*registry.DecodedField) *pbgear.Referenda_After {
	return &pbgear.Referenda_After{
		Value_0: uint32(fields[0].Value),
	}
}

func to_Proxy_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Proxy_TupleNull {
	return &pbgear.Proxy_TupleNull{}
}
func to_Grandpa_Prevote0000000(fields []*registry.DecodedField) *pbgear.Grandpa_Prevote {
	return &pbgear.Grandpa_Prevote{
		Value_0: to_Grandpa_FinalityGrandpaEquivocation(fields[0].Value),
	}
}

func to_Grandpa_SpSessionMembershipProof0000000(fields []*registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
	return &pbgear.Grandpa_SpSessionMembershipProof{
		Session:        uint32(fields[0].Value),
		TrieNodes:      to_repeated_Grandpa_GrandpaTrieNodesList(fields[1].Value),
		ValidatorCount: uint32(fields[2].Value),
	}
}

func to_Grandpa_NoteStalledCall0000000(fields []*registry.DecodedField) *pbgear.Grandpa_NoteStalledCall {
	return &pbgear.Grandpa_NoteStalledCall{
		Delay:                    uint32(fields[0].Value),
		BestFinalizedBlockNumber: uint32(fields[1].Value),
	}
}

func to_BagsList_Address320000000(fields []*registry.DecodedField) *pbgear.BagsList_Address32 {
	return &pbgear.BagsList_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Utility_Root0000000(fields []*registry.DecodedField) *pbgear.Utility_Root {
	return &pbgear.Utility_Root{}
}
func to_NominationPools_SetCommissionMaxCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionMaxCall {
	return &pbgear.NominationPools_SetCommissionMaxCall{
		PoolId:        uint32(fields[0].Value),
		MaxCommission: to_NominationPools_SpArithmeticPerThingsPerbill(fields[1].Value),
	}
}

func to_ConvictionVoting_Locked6X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked6X {
	return &pbgear.ConvictionVoting_Locked6X{}
}
func to_Identity_Raw190000000(fields []*registry.DecodedField) *pbgear.Identity_Raw19 {
	return &pbgear.Identity_Raw19{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_NominationPools_Remove0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Remove {
	return &pbgear.NominationPools_Remove{}
}
func to_Vesting_MergeSchedulesCall0000000(fields []*registry.DecodedField) *pbgear.Vesting_MergeSchedulesCall {
	return &pbgear.Vesting_MergeSchedulesCall{
		Schedule_1Index: uint32(fields[0].Value),
		Schedule_2Index: uint32(fields[1].Value),
	}
}

func to_Staking_ReapStashCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ReapStashCall {
	return &pbgear.Staking_ReapStashCall{
		Stash:            to_Staking_SpCoreCryptoAccountId32(fields[0].Value),
		NumSlashingSpans: uint32(fields[1].Value),
	}
}

func to_Session_SpAuthorityDiscoveryAppPublic0000000(fields []*registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
	return &pbgear.Session_SpAuthorityDiscoveryAppPublic{
		AuthorityDiscovery: to_Session_SpCoreSr25519Public(fields[0].Value),
	}
}

func to_Utility_CompactUint640000000(fields []*registry.DecodedField) *pbgear.Utility_CompactUint64 {
	return &pbgear.Utility_CompactUint64{
		Value: uint64(fields[0].Value),
	}
}

func to_ConvictionVoting_Locked2X0000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked2X {
	return &pbgear.ConvictionVoting_Locked2X{}
}
func to_Identity_SetSubsCall0000000(fields []*registry.DecodedField) *pbgear.Identity_SetSubsCall {
	return &pbgear.Identity_SetSubsCall{
		Subs: to_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_NominationPools_ChillCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_ChillCall {
	return &pbgear.NominationPools_ChillCall{
		PoolId: uint32(fields[0].Value),
	}
}

func to_Gear_RunCall0000000(fields []*registry.DecodedField) *pbgear.Gear_RunCall {
	return &pbgear.Gear_RunCall{
		MaxGas: uint64(fields[0].Value),
	}
}

func to_StakingRewards_CompactTupleNull0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
	return &pbgear.StakingRewards_CompactTupleNull{
		Value: to_StakingRewards_TupleNull(fields[0].Value),
	}
}

func to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature0000000(fields []*registry.DecodedField) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
	return &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{
		Value_0: to_Grandpa_FinalityGrandpaPrevote(fields[0].Value),
		Value_1: to_Grandpa_SpConsensusGrandpaAppSignature(fields[1].Value),
	}
}

func to_FellowshipReferenda_Inline0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Inline {
	return &pbgear.FellowshipReferenda_Inline{
		Value_0: to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(fields[0].Value),
	}
}

func to_Identity_Raw60000000(fields []*registry.DecodedField) *pbgear.Identity_Raw6 {
	return &pbgear.Identity_Raw6{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Balances_ForceSetBalanceCall0000000(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
	return &pbgear.Balances_ForceSetBalanceCall{
		Who:     to_Balances_Who(fields[0].Value),
		NewFree: to_Balances_CompactString(fields[1].Value),
	}
}

func to_Proxy_Id0000000(fields []*registry.DecodedField) *pbgear.Proxy_Id {
	return &pbgear.Proxy_Id{
		Value_0: to_Proxy_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_ChildBounties_Index0000000(fields []*registry.DecodedField) *pbgear.ChildBounties_Index {
	return &pbgear.ChildBounties_Index{
		Value_0: to_ChildBounties_CompactTupleNull(fields[0].Value),
	}
}

func to_NominationPools_Open0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Open {
	return &pbgear.NominationPools_Open{}
}
func to_Proxy_AnnounceCall0000000(fields []*registry.DecodedField) *pbgear.Proxy_AnnounceCall {
	return &pbgear.Proxy_AnnounceCall{
		Real:     to_Proxy_Real(fields[0].Value),
		CallHash: to_Proxy_PrimitiveTypesH256(fields[1].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_Staking_Targets0000000(fields []*registry.DecodedField) *pbgear.Staking_Targets {
	return &pbgear.Staking_Targets{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Staking_ForceNewEraAlwaysCall0000000(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraAlwaysCall {
	return &pbgear.Staking_ForceNewEraAlwaysCall{}
}
func to_Preimage_UnrequestPreimageCall0000000(fields []*registry.DecodedField) *pbgear.Preimage_UnrequestPreimageCall {
	return &pbgear.Preimage_UnrequestPreimageCall{
		Hash: to_Preimage_PrimitiveTypesH256(fields[0].Value),
	}
}

func to_Identity_Target0000000(fields []*registry.DecodedField) *pbgear.Identity_Target {
	return &pbgear.Identity_Target{
		Value: to_oneof_Identity_Value(fields[0].Value),
	}
}

func to_Identity_PrimitiveTypesH2560000000(fields []*registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
	return &pbgear.Identity_PrimitiveTypesH256{
		Identity: []uint32(fields[0].Value),
	}
}

func to_Identity_OutOfDate0000000(fields []*registry.DecodedField) *pbgear.Identity_OutOfDate {
	return &pbgear.Identity_OutOfDate{}
}
func to_GearVoucher_GprimitivesMessageId0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesMessageId {
	return &pbgear.GearVoucher_GprimitivesMessageId{
		ReplyToId: []uint32(fields[0].Value),
	}
}

func to_GearVoucher_CallDeprecatedCall0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_CallDeprecatedCall {
	return &pbgear.GearVoucher_CallDeprecatedCall{
		Call: to_GearVoucher_Call(fields[0].Value),
	}
}

func to_Balances_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
	return &pbgear.Balances_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Balances_Index0000000(fields []*registry.DecodedField) *pbgear.Balances_Index {
	return &pbgear.Balances_Index{
		Value_0: to_Balances_CompactTupleNull(fields[0].Value),
	}
}

func to_Balances_TransferKeepAliveCall0000000(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
	return &pbgear.Balances_TransferKeepAliveCall{
		Dest:  to_Balances_Dest(fields[0].Value),
		Value: to_Balances_CompactString(fields[1].Value),
	}
}

func to_Staking_MaxValidatorCount0000000(fields []*registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
	return &pbgear.Staking_MaxValidatorCount{
		Value: to_oneof_Staking_Value(fields[0].Value),
	}
}

func to_Treasury_VoidSpendCall0000000(fields []*registry.DecodedField) *pbgear.Treasury_VoidSpendCall {
	return &pbgear.Treasury_VoidSpendCall{
		Index: uint32(fields[0].Value),
	}
}

func to_ConvictionVoting_Address200000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address20 {
	return &pbgear.ConvictionVoting_Address20{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Identity_Reasonable0000000(fields []*registry.DecodedField) *pbgear.Identity_Reasonable {
	return &pbgear.Identity_Reasonable{}
}
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint320000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(fields[1].Value),
		Value_2: to_ElectionProviderMultiPhase_CompactUint32(fields[2].Value),
	}
}

func to_ElectionProviderMultiPhase_SpNposElectionsElectionScore0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
	return &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{
		MinimalStake:    string(fields[0].Value),
		SumStake:        string(fields[1].Value),
		SumStakeSquared: string(fields[2].Value),
	}
}

func to_NominationPools_SetClaimPermissionCall0000000(fields []*registry.DecodedField) *pbgear.NominationPools_SetClaimPermissionCall {
	return &pbgear.NominationPools_SetClaimPermissionCall{
		Permission: to_NominationPools_Permission(fields[0].Value),
	}
}

func to_Babe_SpRuntimeGenericDigestDigestItem0000000(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigestItem {
	return &pbgear.Babe_SpRuntimeGenericDigestDigestItem{
		Logs: to_Babe_Logs(fields[0].Value),
	}
}

func to_Utility_WithWeightCall0000000(fields []*registry.DecodedField) *pbgear.Utility_WithWeightCall {
	return &pbgear.Utility_WithWeightCall{
		Call:   to_oneof_Utility_Call(fields[0].Value),
		Weight: to_Utility_SpWeightsWeightV2Weight(fields[1].Value),
	}
}

func to_ConvictionVoting_CompactUint320000000(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
	return &pbgear.ConvictionVoting_CompactUint32{
		Value: uint32(fields[0].Value),
	}
}

func to_NominationPools_SpCoreCryptoAccountId320000000(fields []*registry.DecodedField) *pbgear.NominationPools_SpCoreCryptoAccountId32 {
	return &pbgear.NominationPools_SpCoreCryptoAccountId32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Grandpa_FinalityGrandpaPrevote0000000(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaPrevote {
	return &pbgear.Grandpa_FinalityGrandpaPrevote{
		TargetHash:   to_Grandpa_PrimitiveTypesH256(fields[0].Value),
		TargetNumber: uint32(fields[1].Value),
	}
}

func to_Treasury_TupleNull0000000(fields []*registry.DecodedField) *pbgear.Treasury_TupleNull {
	return &pbgear.Treasury_TupleNull{}
}
func to_FellowshipReferenda_Signed0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Signed {
	return &pbgear.FellowshipReferenda_Signed{
		Value_0: to_FellowshipReferenda_SpCoreCryptoAccountId32(fields[0].Value),
	}
}

func to_Identity_Raw120000000(fields []*registry.DecodedField) *pbgear.Identity_Raw12 {
	return &pbgear.Identity_Raw12{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Proxy_Any0000000(fields []*registry.DecodedField) *pbgear.Proxy_Any {
	return &pbgear.Proxy_Any{}
}
func to_Babe_SpConsensusSlotsEquivocationProof0000000(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
	return &pbgear.Babe_SpConsensusSlotsEquivocationProof{
		Offender:     to_Babe_SpConsensusBabeAppPublic(fields[0].Value),
		Slot:         to_Babe_SpConsensusSlotsSlot(fields[1].Value),
		FirstHeader:  to_Babe_SpRuntimeGenericHeaderHeader(fields[2].Value),
		SecondHeader: to_Babe_SpRuntimeGenericHeaderHeader(fields[3].Value),
	}
}

func to_Staking_Staked0000000(fields []*registry.DecodedField) *pbgear.Staking_Staked {
	return &pbgear.Staking_Staked{}
}
func to_Balances_Address320000000(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
	return &pbgear.Balances_Address32{
		Value_0: []uint32(fields[0].Value),
	}
}

func to_Utility_ForceBatchCall0000000(fields []*registry.DecodedField) *pbgear.Utility_ForceBatchCall {
	return &pbgear.Utility_ForceBatchCall{
		Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(fields[0].Value),
	}
}

func to_Identity_BoundedCollectionsBoundedVecBoundedVec0000000(fields []*registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
	return &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{
		Additional: to_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(fields[0].Value),
	}
}

func to_Gear_GprimitivesCodeId0000000(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
	return &pbgear.Gear_GprimitivesCodeId{
		CodeId: []uint32(fields[0].Value),
	}
}

func to_StakingRewards_ForceRefillCall0000000(fields []*registry.DecodedField) *pbgear.StakingRewards_ForceRefillCall {
	return &pbgear.StakingRewards_ForceRefillCall{
		From:  to_StakingRewards_From(fields[0].Value),
		Value: string(fields[1].Value),
	}
}

func to_Scheduler_ScheduleCall0000000(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleCall {
	return &pbgear.Scheduler_ScheduleCall{
		When:          uint32(fields[0].Value),
		MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(fields[1].Value),
		Priority:      uint32(fields[2].Value),
		Call:          to_oneof_Scheduler_Call(fields[3].Value),
	}
}

func to_Identity_KillIdentityCall0000000(fields []*registry.DecodedField) *pbgear.Identity_KillIdentityCall {
	return &pbgear.Identity_KillIdentityCall{
		Target: to_Identity_Target(fields[0].Value),
	}
}

func to_GearVoucher_UploadCode0000000(fields []*registry.DecodedField) *pbgear.GearVoucher_UploadCode {
	return &pbgear.GearVoucher_UploadCode{
		Code: []uint32(fields[0].Value),
	}
}

func to_Balances_Dest0000000(fields []*registry.DecodedField) *pbgear.Balances_Dest {
	return &pbgear.Balances_Dest{
		Value: to_oneof_Balances_Value(fields[0].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU160000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
	return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{
		Value_0: to_ElectionProviderMultiPhase_CompactUint32(fields[0].Value),
		Value_1: to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(fields[1].Value),
	}
}

func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport0000000(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
	return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{
		Value_0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(fields[0].Value),
		Value_1: to_ElectionProviderMultiPhase_SpNposElectionsSupport(fields[1].Value),
	}
}

func to_NominationPools_Rewards0000000(fields []*registry.DecodedField) *pbgear.NominationPools_Rewards {
	return &pbgear.NominationPools_Rewards{}
}
func to_NominationPools_MemberAccount0000000(fields []*registry.DecodedField) *pbgear.NominationPools_MemberAccount {
	return &pbgear.NominationPools_MemberAccount{
		Value: to_oneof_NominationPools_Value(fields[0].Value),
	}
}

func to_Babe_PreRuntime0000000(fields []*registry.DecodedField) *pbgear.Babe_PreRuntime {
	return &pbgear.Babe_PreRuntime{
		Value_0: []uint32(fields[0].Value),
		Value_1: []uint32(fields[1].Value),
	}
}

func to_FellowshipCollective_AddMemberCall0000000(fields []*registry.DecodedField) *pbgear.FellowshipCollective_AddMemberCall {
	return &pbgear.FellowshipCollective_AddMemberCall{
		Who: to_FellowshipCollective_Who(fields[0].Value),
	}
}

func to_FellowshipReferenda_EnactmentMoment0000000(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
	return &pbgear.FellowshipReferenda_EnactmentMoment{
		Value: to_oneof_FellowshipReferenda_Value(fields[0].Value),
	}
}

func to_Proxy_Delegate0000000(fields []*registry.DecodedField) *pbgear.Proxy_Delegate {
	return &pbgear.Proxy_Delegate{
		Value: to_oneof_Proxy_Value(fields[0].Value),
	}
}

func to_NominationPools_CompactString0000000(fields []*registry.DecodedField) *pbgear.NominationPools_CompactString {
	return &pbgear.NominationPools_CompactString{
		Value: string(fields[0].Value),
	}
}

func to_Session_VaraRuntimeSessionKeys0000000(fields []*registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
	return &pbgear.Session_VaraRuntimeSessionKeys{
		Babe:               to_Session_SpConsensusBabeAppPublic(fields[0].Value),
		Grandpa:            to_Session_SpConsensusGrandpaAppPublic(fields[1].Value),
		ImOnline:           to_Session_PalletImOnlineSr25519AppSr25519Public(fields[2].Value),
		AuthorityDiscovery: to_Session_SpAuthorityDiscoveryAppPublic(fields[3].Value),
	}
}
