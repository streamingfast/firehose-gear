package gen_types

import (
    "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
    pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1"
)

func to_Preimage_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
    return &pbgear.Preimage_PrimitiveTypesH256{
                Hash : []uint32(field.Value),
    }
}
            
                   
func to_Multisig_CompactUint64salut(fields []*registry.DecodedField) *pbgear.Multisig_CompactUint64 {
    return &pbgear.Multisig_CompactUint64{
                Value : uint64(field.Value),
    }
}
            
                   
func to_NominationPools_MaxMemberssalut(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembers {
    return &pbgear.NominationPools_MaxMembers{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}
            
                func to_oneof_NominationPools_Value(in *registry.DecodedField) *pbgear.NominationPools_Value {
                    return nil
                }   
func to_GearVoucher_UploadCodesalut(fields []*registry.DecodedField) *pbgear.GearVoucher_UploadCode {
    return &pbgear.GearVoucher_UploadCode{
                Code : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_Votesalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Vote {
    return &pbgear.ConvictionVoting_Vote{
                Value: to_oneof_ConvictionVoting_Value(field.Value),
    }
}
            
                func to_oneof_ConvictionVoting_Value(in *registry.DecodedField) *pbgear.ConvictionVoting_Value {
                    return nil
                }   
func to_FellowshipReferenda_RefundSubmissionDepositCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
    return &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{
                Index : uint32(field.Value),
    }
}
            
                   
func to_Identity_Displaysalut(fields []*registry.DecodedField) *pbgear.Identity_Display {
    return &pbgear.Identity_Display{
                Value: to_oneof_Identity_Value(field.Value),
    }
}
            
                func to_oneof_Identity_Value(in *registry.DecodedField) *pbgear.Identity_Value {
                    return nil
                }   
func to_Scheduler_ScheduleAfterCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleAfterCall {
    return &pbgear.Scheduler_ScheduleAfterCall{
                After : uint32(field.Value),
                MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(field.Value),
                Priority : uint32(field.Value),
                Call: to_oneof_Scheduler_Call(field.Value),
    }
}
            
                  
            
                func to_optional_Scheduler_TupleUint32Uint32(in *registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
                    return &pbgear.Scheduler_TupleUint32Uint32{
                        // TODO fill all fields
                    }
                }   
            
                func to_oneof_Scheduler_Call(in *registry.DecodedField) *pbgear.Scheduler_Call {
                    return nil
                }   
func to_Identity_Address32salut(fields []*registry.DecodedField) *pbgear.Identity_Address32 {
    return &pbgear.Identity_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_ProxyAnnouncedCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_ProxyAnnouncedCall {
    return &pbgear.Proxy_ProxyAnnouncedCall{
                Delegate: to_Proxy_Delegate(field.Value),
                Real: to_Proxy_Real(field.Value),
                ForceProxyType: to_optional_Proxy_ForceProxyType(field.Value),
                Call: to_oneof_Proxy_Call(field.Value),
    }
}
            
                func to_Proxy_Delegate(in *registry.DecodedField) *pbgear.Proxy_Delegate {
                    return &pbgear.Proxy_Delegate{
                        // TODO fill all fields
                    }
                }  
            
                func to_Proxy_Real(in *registry.DecodedField) *pbgear.Proxy_Real {
                    return &pbgear.Proxy_Real{
                        // TODO fill all fields
                    }
                }  
            
                func to_optional_Proxy_ForceProxyType(in *registry.DecodedField) *pbgear.Proxy_ForceProxyType {
                    return &pbgear.Proxy_ForceProxyType{
                        // TODO fill all fields
                    }
                }  
            
                func to_oneof_Proxy_Call(in *registry.DecodedField) *pbgear.Proxy_Call {
                    return nil
                }   
func to_GearVoucher_GprimitivesActorIdsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
    return &pbgear.GearVoucher_GprimitivesActorId{
                Programs: to_GearVoucher_GprimitivesActorId(field.Value),
    }
}
            
                func to_GearVoucher_GprimitivesActorId(in *registry.DecodedField) *pbgear.GearVoucher_GprimitivesActorId {
                    return &pbgear.GearVoucher_GprimitivesActorId{
                        // TODO fill all fields
                    }
                }   
func to_Vesting_Address32salut(fields []*registry.DecodedField) *pbgear.Vesting_Address32 {
    return &pbgear.Vesting_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Referenda_CancelCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_CancelCall {
    return &pbgear.Referenda_CancelCall{
                Index : uint32(field.Value),
    }
}
            
                   
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}
            
                func to_ElectionProviderMultiPhase_CompactUint32(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
                    return &pbgear.ElectionProviderMultiPhase_CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{
                        // TODO fill all fields
                    }
                }    
func to_NominationPools_Address20salut(fields []*registry.DecodedField) *pbgear.NominationPools_Address20 {
    return &pbgear.NominationPools_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Session_SpConsensusGrandpaAppPublicsalut(fields []*registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
    return &pbgear.Session_SpConsensusGrandpaAppPublic{
                Grandpa: to_Session_SpCoreEd25519Public(field.Value),
    }
}
            
                func to_Session_SpCoreEd25519Public(in *registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
                    return &pbgear.Session_SpCoreEd25519Public{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipCollective_DemoteMemberCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_DemoteMemberCall {
    return &pbgear.FellowshipCollective_DemoteMemberCall{
                Who: to_FellowshipCollective_Who(field.Value),
    }
}
            
                func to_FellowshipCollective_Who(in *registry.DecodedField) *pbgear.FellowshipCollective_Who {
                    return &pbgear.FellowshipCollective_Who{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Indexsalut(fields []*registry.DecodedField) *pbgear.Identity_Index {
    return &pbgear.Identity_Index{
                Value0: to_Identity_CompactTupleNull(field.Value),
    }
}
            
                func to_Identity_CompactTupleNull(in *registry.DecodedField) *pbgear.Identity_CompactTupleNull {
                    return &pbgear.Identity_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_RemoveProxyCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxyCall {
    return &pbgear.Proxy_RemoveProxyCall{
                Delegate: to_Proxy_Delegate(field.Value),
                ProxyType: to_Proxy_ProxyType(field.Value),
                Delay : uint32(field.Value),
    }
} 
            
                func to_Proxy_ProxyType(in *registry.DecodedField) *pbgear.Proxy_ProxyType {
                    return &pbgear.Proxy_ProxyType{
                        // TODO fill all fields
                    }
                }  
            
                   
func to_GearVoucher_BTreeSetsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
    return &pbgear.GearVoucher_BTreeSet{
                Programs: to_repeated_GearVoucher_GprimitivesActorId(field.Value),
    }
}  
func to_BagsList_RebagCallsalut(fields []*registry.DecodedField) *pbgear.BagsList_RebagCall {
    return &pbgear.BagsList_RebagCall{
                Dislocated: to_BagsList_Dislocated(field.Value),
    }
}
            
                func to_BagsList_Dislocated(in *registry.DecodedField) *pbgear.BagsList_Dislocated {
                    return &pbgear.BagsList_Dislocated{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_ProposalOriginsalut(fields []*registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
    return &pbgear.Referenda_ProposalOrigin{
                Value: to_oneof_Referenda_Value(field.Value),
    }
}
            
                func to_oneof_Referenda_Value(in *registry.DecodedField) *pbgear.Referenda_Value {
                    return nil
                }   
func to_ChildBounties_ClaimChildBountyCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_ClaimChildBountyCall {
    return &pbgear.ChildBounties_ClaimChildBountyCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
    }
}
            
                func to_ChildBounties_CompactUint32(in *registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
                    return &pbgear.ChildBounties_CompactUint32{
                        // TODO fill all fields
                    }
                }    
func to_Grandpa_SpConsensusGrandpaAppSignaturesalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
    return &pbgear.Grandpa_SpConsensusGrandpaAppSignature{
                Value1: to_Grandpa_SpCoreEd25519Signature(field.Value),
    }
}
            
                func to_Grandpa_SpCoreEd25519Signature(in *registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Signature {
                    return &pbgear.Grandpa_SpCoreEd25519Signature{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipReferenda_Originssalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Origins {
    return &pbgear.FellowshipReferenda_Origins{
                Value0: to_FellowshipReferenda_Value0(field.Value),
    }
}
            
                func to_FellowshipReferenda_Value0(in *registry.DecodedField) *pbgear.FellowshipReferenda_Value0 {
                    return &pbgear.FellowshipReferenda_Value0{
                        // TODO fill all fields
                    }
                }   
func to_Staking_Targetssalut(fields []*registry.DecodedField) *pbgear.Staking_Targets {
    return &pbgear.Staking_Targets{
                Value: to_oneof_Staking_Value(field.Value),
    }
}
            
                func to_oneof_Staking_Value(in *registry.DecodedField) *pbgear.Staking_Value {
                    return nil
                }   
func to_Session_PurgeKeysCallsalut(fields []*registry.DecodedField) *pbgear.Session_PurgeKeysCall {
    return &pbgear.Session_PurgeKeysCall{
    }
} 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_NominationPools_Indexsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Index {
    return &pbgear.NominationPools_Index{
                Value0: to_NominationPools_CompactTupleNull(field.Value),
    }
}
            
                func to_NominationPools_CompactTupleNull(in *registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
                    return &pbgear.NominationPools_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Babe_BabeTrieNodesListsalut(fields []*registry.DecodedField) *pbgear.Babe_BabeTrieNodesList {
    return &pbgear.Babe_BabeTrieNodesList{
                TrieNodes : []uint32(field.Value),
    }
}  
func to_Grandpa_NoteStalledCallsalut(fields []*registry.DecodedField) *pbgear.Grandpa_NoteStalledCall {
    return &pbgear.Grandpa_NoteStalledCall{
                Delay : uint32(field.Value),
                BestFinalizedBlockNumber : uint32(field.Value),
    }
}
            
                    
func to_Referenda_Proposalsalut(fields []*registry.DecodedField) *pbgear.Referenda_Proposal {
    return &pbgear.Referenda_Proposal{
                Value: to_oneof_Referenda_Value(field.Value),
    }
}  
func to_Identity_Raw6salut(fields []*registry.DecodedField) *pbgear.Identity_Raw6 {
    return &pbgear.Identity_Raw6{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_CreateCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_CreateCall {
    return &pbgear.NominationPools_CreateCall{
                Amount: to_NominationPools_CompactString(field.Value),
                Root: to_NominationPools_Root(field.Value),
                Nominator: to_NominationPools_Nominator(field.Value),
                Bouncer: to_NominationPools_Bouncer(field.Value),
    }
}
            
                func to_NominationPools_CompactString(in *registry.DecodedField) *pbgear.NominationPools_CompactString {
                    return &pbgear.NominationPools_CompactString{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_Root(in *registry.DecodedField) *pbgear.NominationPools_Root {
                    return &pbgear.NominationPools_Root{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_Nominator(in *registry.DecodedField) *pbgear.NominationPools_Nominator {
                    return &pbgear.NominationPools_Nominator{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_Bouncer(in *registry.DecodedField) *pbgear.NominationPools_Bouncer {
                    return &pbgear.NominationPools_Bouncer{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_ChillCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_ChillCall {
    return &pbgear.NominationPools_ChillCall{
                PoolId : uint32(field.Value),
    }
}
            
                   
func to_BagsList_Rawsalut(fields []*registry.DecodedField) *pbgear.BagsList_Raw {
    return &pbgear.BagsList_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Utility_Value0salut(fields []*registry.DecodedField) *pbgear.Utility_Value0 {
    return &pbgear.Utility_Value0{
                Value: to_oneof_Utility_Value(field.Value),
    }
}
            
                func to_oneof_Utility_Value(in *registry.DecodedField) *pbgear.Utility_Value {
                    return nil
                }   
func to_FellowshipCollective_Rawsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Raw {
    return &pbgear.FellowshipCollective_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Bounties_Address20salut(fields []*registry.DecodedField) *pbgear.Bounties_Address20 {
    return &pbgear.Bounties_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Bounties_AwardBountyCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_AwardBountyCall {
    return &pbgear.Bounties_AwardBountyCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
                Beneficiary: to_Bounties_Beneficiary(field.Value),
    }
}
            
                func to_Bounties_CompactUint32(in *registry.DecodedField) *pbgear.Bounties_CompactUint32 {
                    return &pbgear.Bounties_CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_Bounties_Beneficiary(in *registry.DecodedField) *pbgear.Bounties_Beneficiary {
                    return &pbgear.Bounties_Beneficiary{
                        // TODO fill all fields
                    }
                }   
func to_Staking_Indexsalut(fields []*registry.DecodedField) *pbgear.Staking_Index {
    return &pbgear.Staking_Index{
                Value0: to_Staking_CompactTupleNull(field.Value),
    }
}
            
                func to_Staking_CompactTupleNull(in *registry.DecodedField) *pbgear.Staking_CompactTupleNull {
                    return &pbgear.Staking_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipCollective_TupleNullsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
    return &pbgear.FellowshipCollective_TupleNull{
    }
} 
func to_Staking_NominateCallsalut(fields []*registry.DecodedField) *pbgear.Staking_NominateCall {
    return &pbgear.Staking_NominateCall{
                Targets: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(field.Value),
    }
}
            
                func to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(in *registry.DecodedField) []*pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
                    return []*pbgear.Staking_SpRuntimeMultiaddressMultiAddress{
                        // TODO fill all fields
                    }
                }   
func to_Utility_AsDerivativeCallsalut(fields []*registry.DecodedField) *pbgear.Utility_AsDerivativeCall {
    return &pbgear.Utility_AsDerivativeCall{
                Index : uint32(field.Value),
                Call: to_oneof_Utility_Call(field.Value),
    }
}
            
                  
            
                func to_oneof_Utility_Call(in *registry.DecodedField) *pbgear.Utility_Call {
                    return nil
                }   
func to_ConvictionVoting_Address32salut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address32 {
    return &pbgear.ConvictionVoting_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw23salut(fields []*registry.DecodedField) *pbgear.Identity_Raw23 {
    return &pbgear.Identity_Raw23{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Identity_CompactUint32 {
    return &pbgear.Identity_CompactUint32{
                Value : uint32(field.Value),
    }
}
            
                   
func to_Babe_PreRuntimesalut(fields []*registry.DecodedField) *pbgear.Babe_PreRuntime {
    return &pbgear.Babe_PreRuntime{
                Value0 : []uint32(field.Value),
                Value1 : []uint32(field.Value),
    }
}   
func to_Staking_ValidateCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ValidateCall {
    return &pbgear.Staking_ValidateCall{
                Prefs: to_Staking_PalletStakingValidatorPrefs(field.Value),
    }
}
            
                func to_Staking_PalletStakingValidatorPrefs(in *registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
                    return &pbgear.Staking_PalletStakingValidatorPrefs{
                        // TODO fill all fields
                    }
                }   
func to_Utility_Rootsalut(fields []*registry.DecodedField) *pbgear.Utility_Root {
    return &pbgear.Utility_Root{
    }
} 
func to_Identity_PalletIdentityTypesBitFlagssalut(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
    return &pbgear.Identity_PalletIdentityTypesBitFlags{
                Fields : uint64(field.Value),
    }
}
            
                   
func to_Bounties_Idsalut(fields []*registry.DecodedField) *pbgear.Bounties_Id {
    return &pbgear.Bounties_Id{
                Value0: to_Bounties_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Bounties_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Bounties_SpCoreCryptoAccountId32 {
                    return &pbgear.Bounties_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Grandpa_SpCoreEd25519Publicsalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Public {
    return &pbgear.Grandpa_SpCoreEd25519Public{
                Identity : []uint32(field.Value),
    }
}  
func to_Staking_SetPayeeCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetPayeeCall {
    return &pbgear.Staking_SetPayeeCall{
                Payee: to_Staking_Payee(field.Value),
    }
}
            
                func to_Staking_Payee(in *registry.DecodedField) *pbgear.Staking_Payee {
                    return &pbgear.Staking_Payee{
                        // TODO fill all fields
                    }
                }   
func to_Identity_SetAccountIdCallsalut(fields []*registry.DecodedField) *pbgear.Identity_SetAccountIdCall {
    return &pbgear.Identity_SetAccountIdCall{
                Index: to_Identity_CompactUint32(field.Value),
                New: to_Identity_New(field.Value),
    }
}
            
                func to_Identity_CompactUint32(in *registry.DecodedField) *pbgear.Identity_CompactUint32 {
                    return &pbgear.Identity_CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_New(in *registry.DecodedField) *pbgear.Identity_New {
                    return &pbgear.Identity_New{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.ChildBounties_SpCoreCryptoAccountId32 {
    return &pbgear.ChildBounties_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_GearVoucher_IssueCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_IssueCall {
    return &pbgear.GearVoucher_IssueCall{
                Spender: to_GearVoucher_SpCoreCryptoAccountId32(field.Value),
                Balance : string(field.Value),
                Programs: to_optional_GearVoucher_BTreeSet(field.Value),
                CodeUploading : bool(field.Value),
                Duration : uint32(field.Value),
    }
}
            
                func to_GearVoucher_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
                    return &pbgear.GearVoucher_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }  
            
                  
            
                func to_optional_GearVoucher_BTreeSet(in *registry.DecodedField) *pbgear.GearVoucher_BTreeSet {
                    return &pbgear.GearVoucher_BTreeSet{
                        // TODO fill all fields
                    }
                }  
            
                  
            
                   
func to_Grandpa_SpCoreEd25519Signaturesalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Signature {
    return &pbgear.Grandpa_SpCoreEd25519Signature{
                Value1 : []uint32(field.Value),
    }
}  
func to_Session_VaraRuntimeSessionKeyssalut(fields []*registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
    return &pbgear.Session_VaraRuntimeSessionKeys{
                Babe: to_Session_SpConsensusBabeAppPublic(field.Value),
                Grandpa: to_Session_SpConsensusGrandpaAppPublic(field.Value),
                ImOnline: to_Session_PalletImOnlineSr25519AppSr25519Public(field.Value),
                AuthorityDiscovery: to_Session_SpAuthorityDiscoveryAppPublic(field.Value),
    }
}
            
                func to_Session_SpConsensusBabeAppPublic(in *registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
                    return &pbgear.Session_SpConsensusBabeAppPublic{
                        // TODO fill all fields
                    }
                }  
            
                func to_Session_SpConsensusGrandpaAppPublic(in *registry.DecodedField) *pbgear.Session_SpConsensusGrandpaAppPublic {
                    return &pbgear.Session_SpConsensusGrandpaAppPublic{
                        // TODO fill all fields
                    }
                }  
            
                func to_Session_PalletImOnlineSr25519AppSr25519Public(in *registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
                    return &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{
                        // TODO fill all fields
                    }
                }  
            
                func to_Session_SpAuthorityDiscoveryAppPublic(in *registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
                    return &pbgear.Session_SpAuthorityDiscoveryAppPublic{
                        // TODO fill all fields
                    }
                }   
func to_ElectionProviderMultiPhase_SpNposElectionsElectionScoresalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
    return &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{
                MinimalStake : string(field.Value),
                SumStake : string(field.Value),
                SumStakeSquared : string(field.Value),
    }
}
            
                     
func to_FellowshipReferenda_Proposalsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
    return &pbgear.FellowshipReferenda_Proposal{
                Value: to_oneof_FellowshipReferenda_Value(field.Value),
    }
}
            
                func to_oneof_FellowshipReferenda_Value(in *registry.DecodedField) *pbgear.FellowshipReferenda_Value {
                    return nil
                }   
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_Whitelist_WhitelistCallCallsalut(fields []*registry.DecodedField) *pbgear.Whitelist_WhitelistCallCall {
    return &pbgear.Whitelist_WhitelistCallCall{
                CallHash: to_Whitelist_PrimitiveTypesH256(field.Value),
    }
}
            
                func to_Whitelist_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
                    return &pbgear.Whitelist_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }   
func to_Scheduler_ScheduleNamedCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedCall {
    return &pbgear.Scheduler_ScheduleNamedCall{
                Id : []uint32(field.Value),
                When : uint32(field.Value),
                MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(field.Value),
                Priority : uint32(field.Value),
                Call: to_oneof_Scheduler_Call(field.Value),
    }
}      
func to_Identity_CompactStringsalut(fields []*registry.DecodedField) *pbgear.Identity_CompactString {
    return &pbgear.Identity_CompactString{
                Value : string(field.Value),
    }
}
            
                   
func to_Proxy_Anysalut(fields []*registry.DecodedField) *pbgear.Proxy_Any {
    return &pbgear.Proxy_Any{
    }
} 
func to_Proxy_Governancesalut(fields []*registry.DecodedField) *pbgear.Proxy_Governance {
    return &pbgear.Proxy_Governance{
    }
} 
func to_Balances_ForceTransferCallsalut(fields []*registry.DecodedField) *pbgear.Balances_ForceTransferCall {
    return &pbgear.Balances_ForceTransferCall{
                Source: to_Balances_Source(field.Value),
                Dest: to_Balances_Dest(field.Value),
                Value: to_Balances_CompactString(field.Value),
    }
}
            
                func to_Balances_Source(in *registry.DecodedField) *pbgear.Balances_Source {
                    return &pbgear.Balances_Source{
                        // TODO fill all fields
                    }
                }  
            
                func to_Balances_Dest(in *registry.DecodedField) *pbgear.Balances_Dest {
                    return &pbgear.Balances_Dest{
                        // TODO fill all fields
                    }
                }  
            
                func to_Balances_CompactString(in *registry.DecodedField) *pbgear.Balances_CompactString {
                    return &pbgear.Balances_CompactString{
                        // TODO fill all fields
                    }
                }   
func to_Staking_Nonesalut(fields []*registry.DecodedField) *pbgear.Staking_None {
    return &pbgear.Staking_None{
    }
} 
func to_Utility_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Utility_SpCoreCryptoAccountId32 {
    return &pbgear.Utility_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_CreateWithPoolIdCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_CreateWithPoolIdCall {
    return &pbgear.NominationPools_CreateWithPoolIdCall{
                Amount: to_NominationPools_CompactString(field.Value),
                Root: to_NominationPools_Root(field.Value),
                Nominator: to_NominationPools_Nominator(field.Value),
                Bouncer: to_NominationPools_Bouncer(field.Value),
                PoolId : uint32(field.Value),
    }
}      
func to_StakingRewards_Fromsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_From {
    return &pbgear.StakingRewards_From{
                Value: to_oneof_StakingRewards_Value(field.Value),
    }
}
            
                func to_oneof_StakingRewards_Value(in *registry.DecodedField) *pbgear.StakingRewards_Value {
                    return nil
                }   
func to_Timestamp_CompactUint64salut(fields []*registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
    return &pbgear.Timestamp_CompactUint64{
                Value : uint64(field.Value),
    }
}
            
                   
func to_Grandpa_SpSessionMembershipProofsalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
    return &pbgear.Grandpa_SpSessionMembershipProof{
                Session : uint32(field.Value),
                TrieNodes: to_repeated_Grandpa_GrandpaTrieNodesList(field.Value),
                ValidatorCount : uint32(field.Value),
    }
} 
            
                func to_repeated_Grandpa_GrandpaTrieNodesList(in *registry.DecodedField) []*pbgear.Grandpa_GrandpaTrieNodesList {
                    return []*pbgear.Grandpa_GrandpaTrieNodesList{
                        // TODO fill all fields
                    }
                }    
func to_Referenda_RefundDecisionDepositCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_RefundDecisionDepositCall {
    return &pbgear.Referenda_RefundDecisionDepositCall{
                Index : uint32(field.Value),
    }
}  
func to_Identity_Address20salut(fields []*registry.DecodedField) *pbgear.Identity_Address20 {
    return &pbgear.Identity_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Gear_SetExecuteInherentCallsalut(fields []*registry.DecodedField) *pbgear.Gear_SetExecuteInherentCall {
    return &pbgear.Gear_SetExecuteInherentCall{
                Value : bool(field.Value),
    }
}
            
                   
func to_Vesting_PalletVestingVestingInfoVestingInfosalut(fields []*registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
    return &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{
                Locked : string(field.Value),
                PerBlock : string(field.Value),
                StartingBlock : uint32(field.Value),
    }
}
            
                   
            
                   
func to_Treasury_Indexsalut(fields []*registry.DecodedField) *pbgear.Treasury_Index {
    return &pbgear.Treasury_Index{
                Value0: to_Treasury_CompactTupleNull(field.Value),
    }
}
            
                func to_Treasury_CompactTupleNull(in *registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
                    return &pbgear.Treasury_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw2salut(fields []*registry.DecodedField) *pbgear.Identity_Raw2 {
    return &pbgear.Identity_Raw2{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw5salut(fields []*registry.DecodedField) *pbgear.Identity_Raw5 {
    return &pbgear.Identity_Raw5{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_MaxPoolssalut(fields []*registry.DecodedField) *pbgear.NominationPools_MaxPools {
    return &pbgear.NominationPools_MaxPools{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_BondExtraOtherCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraOtherCall {
    return &pbgear.NominationPools_BondExtraOtherCall{
                Member: to_NominationPools_Member(field.Value),
                Extra: to_NominationPools_Extra(field.Value),
    }
}
            
                func to_NominationPools_Member(in *registry.DecodedField) *pbgear.NominationPools_Member {
                    return &pbgear.NominationPools_Member{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_Extra(in *registry.DecodedField) *pbgear.NominationPools_Extra {
                    return &pbgear.NominationPools_Extra{
                        // TODO fill all fields
                    }
                }   
func to_StakingRewards_Tosalut(fields []*registry.DecodedField) *pbgear.StakingRewards_To {
    return &pbgear.StakingRewards_To{
                Value: to_oneof_StakingRewards_Value(field.Value),
    }
}  
func to_GearVoucher_CallCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_CallCall {
    return &pbgear.GearVoucher_CallCall{
                VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(field.Value),
                Call: to_GearVoucher_Call(field.Value),
    }
}
            
                func to_GearVoucher_PalletGearVoucherInternalVoucherId(in *registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
                    return &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{
                        // TODO fill all fields
                    }
                }  
            
                func to_GearVoucher_Call(in *registry.DecodedField) *pbgear.GearVoucher_Call {
                    return &pbgear.GearVoucher_Call{
                        // TODO fill all fields
                    }
                }   
func to_Staking_SetMinCommissionCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetMinCommissionCall {
    return &pbgear.Staking_SetMinCommissionCall{
                New: to_Staking_SpArithmeticPerThingsPerbill(field.Value),
    }
}
            
                func to_Staking_SpArithmeticPerThingsPerbill(in *registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
                    return &pbgear.Staking_SpArithmeticPerThingsPerbill{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipCollective_PromoteMemberCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_PromoteMemberCall {
    return &pbgear.FellowshipCollective_PromoteMemberCall{
                Who: to_FellowshipCollective_Who(field.Value),
    }
}  
func to_GearVoucher_DeclineCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineCall {
    return &pbgear.GearVoucher_DeclineCall{
                VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(field.Value),
    }
}  
func to_Identity_Datasalut(fields []*registry.DecodedField) *pbgear.Identity_Data {
    return &pbgear.Identity_Data{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Proxy_AnnounceCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_AnnounceCall {
    return &pbgear.Proxy_AnnounceCall{
                Real: to_Proxy_Real(field.Value),
                CallHash: to_Proxy_PrimitiveTypesH256(field.Value),
    }
} 
            
                func to_Proxy_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
                    return &pbgear.Proxy_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_GlobalMaxCommissionsalut(fields []*registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
    return &pbgear.NominationPools_GlobalMaxCommission{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Staking_Address20salut(fields []*registry.DecodedField) *pbgear.Staking_Address20 {
    return &pbgear.Staking_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_Locked6Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked6X {
    return &pbgear.ConvictionVoting_Locked6X{
    }
} 
func to_Utility_Voidsalut(fields []*registry.DecodedField) *pbgear.Utility_Void {
    return &pbgear.Utility_Void{
                Value0: to_Utility_Value0(field.Value),
    }
}
            
                func to_Utility_Value0(in *registry.DecodedField) *pbgear.Utility_Value0 {
                    return &pbgear.Utility_Value0{
                        // TODO fill all fields
                    }
                }   
func to_ConvictionVoting_Convictionsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
    return &pbgear.ConvictionVoting_Conviction{
                Value: to_oneof_ConvictionVoting_Value(field.Value),
    }
}  
func to_Identity_Raw0salut(fields []*registry.DecodedField) *pbgear.Identity_Raw0 {
    return &pbgear.Identity_Raw0{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesDatasalut(fields []*registry.DecodedField) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
    return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{
                Value0: to_Identity_SpCoreCryptoAccountId32(field.Value),
                Value1: to_Identity_Value1(field.Value),
    }
}
            
                func to_Identity_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
                    return &pbgear.Identity_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Value1(in *registry.DecodedField) *pbgear.Identity_Value1 {
                    return &pbgear.Identity_Value1{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Judgementsalut(fields []*registry.DecodedField) *pbgear.Identity_Judgement {
    return &pbgear.Identity_Judgement{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Bounties_CompactStringsalut(fields []*registry.DecodedField) *pbgear.Bounties_CompactString {
    return &pbgear.Bounties_CompactString{
                Value : string(field.Value),
    }
}
            
                   
func to_Grandpa_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Grandpa_PrimitiveTypesH256 {
    return &pbgear.Grandpa_PrimitiveTypesH256{
                TargetHash : []uint32(field.Value),
    }
}  
func to_Staking_Controllersalut(fields []*registry.DecodedField) *pbgear.Staking_Controller {
    return &pbgear.Staking_Controller{
    }
} 
func to_GearVoucher_GprimitivesMessageIdsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_GprimitivesMessageId {
    return &pbgear.GearVoucher_GprimitivesMessageId{
                ReplyToId : []uint32(field.Value),
    }
}  
func to_Staking_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Staking_CompactTupleNull {
    return &pbgear.Staking_CompactTupleNull{
                Value: to_Staking_TupleNull(field.Value),
    }
}
            
                func to_Staking_TupleNull(in *registry.DecodedField) *pbgear.Staking_TupleNull {
                    return &pbgear.Staking_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_Originssalut(fields []*registry.DecodedField) *pbgear.Referenda_Origins {
    return &pbgear.Referenda_Origins{
                Value0: to_Referenda_Value0(field.Value),
    }
}
            
                func to_Referenda_Value0(in *registry.DecodedField) *pbgear.Referenda_Value0 {
                    return &pbgear.Referenda_Value0{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipReferenda_Rootsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Root {
    return &pbgear.FellowshipReferenda_Root{
    }
} 
func to_ChildBounties_Rawsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_Raw {
    return &pbgear.ChildBounties_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Balances_ForceSetBalanceCallsalut(fields []*registry.DecodedField) *pbgear.Balances_ForceSetBalanceCall {
    return &pbgear.Balances_ForceSetBalanceCall{
                Who: to_Balances_Who(field.Value),
                NewFree: to_Balances_CompactString(field.Value),
    }
}
            
                func to_Balances_Who(in *registry.DecodedField) *pbgear.Balances_Who {
                    return &pbgear.Balances_Who{
                        // TODO fill all fields
                    }
                }    
func to_Treasury_VoidSpendCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_VoidSpendCall {
    return &pbgear.Treasury_VoidSpendCall{
                Index : uint32(field.Value),
    }
}
            
                   
func to_FellowshipReferenda_KillCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_KillCall {
    return &pbgear.FellowshipReferenda_KillCall{
                Index : uint32(field.Value),
    }
}  
func to_NominationPools_NominateCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_NominateCall {
    return &pbgear.NominationPools_NominateCall{
                PoolId : uint32(field.Value),
                Validators: to_repeated_NominationPools_SpCoreCryptoAccountId32(field.Value),
    }
} 
            
                func to_repeated_NominationPools_SpCoreCryptoAccountId32(in *registry.DecodedField) []*pbgear.NominationPools_SpCoreCryptoAccountId32 {
                    return []*pbgear.NominationPools_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Gear_GprimitivesActorIdsalut(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
    return &pbgear.Gear_GprimitivesActorId{
                Destination : []uint32(field.Value),
    }
}  
func to_System_TupleSystemItemsListSystemItemsListsalut(fields []*registry.DecodedField) *pbgear.System_TupleSystemItemsListSystemItemsList {
    return &pbgear.System_TupleSystemItemsListSystemItemsList{
                Value0 : []uint32(field.Value),
                Value1 : []uint32(field.Value),
    }
}   
func to_ConvictionVoting_Address20salut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Address20 {
    return &pbgear.ConvictionVoting_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_JoinCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_JoinCall {
    return &pbgear.NominationPools_JoinCall{
                Amount: to_NominationPools_CompactString(field.Value),
                PoolId : uint32(field.Value),
    }
}   
func to_Staking_SetControllerCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetControllerCall {
    return &pbgear.Staking_SetControllerCall{
    }
} 
func to_Referenda_EnactmentMomentsalut(fields []*registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
    return &pbgear.Referenda_EnactmentMoment{
                Value: to_oneof_Referenda_Value(field.Value),
    }
}  
func to_Identity_ProvideJudgementCallsalut(fields []*registry.DecodedField) *pbgear.Identity_ProvideJudgementCall {
    return &pbgear.Identity_ProvideJudgementCall{
                RegIndex: to_Identity_CompactUint32(field.Value),
                Target: to_Identity_Target(field.Value),
                Judgement: to_Identity_Judgement(field.Value),
                Identity: to_Identity_PrimitiveTypesH256(field.Value),
    }
} 
            
                func to_Identity_Target(in *registry.DecodedField) *pbgear.Identity_Target {
                    return &pbgear.Identity_Target{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Judgement(in *registry.DecodedField) *pbgear.Identity_Judgement {
                    return &pbgear.Identity_Judgement{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
                    return &pbgear.Identity_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }   
func to_Treasury_Beneficiarysalut(fields []*registry.DecodedField) *pbgear.Treasury_Beneficiary {
    return &pbgear.Treasury_Beneficiary{
                Value: to_oneof_Treasury_Value(field.Value),
    }
}
            
                func to_oneof_Treasury_Value(in *registry.DecodedField) *pbgear.Treasury_Value {
                    return nil
                }   
func to_Identity_Raw29salut(fields []*registry.DecodedField) *pbgear.Identity_Raw29 {
    return &pbgear.Identity_Raw29{
                Value0 : []uint32(field.Value),
    }
}  
func to_Referenda_Systemsalut(fields []*registry.DecodedField) *pbgear.Referenda_System {
    return &pbgear.Referenda_System{
                Value0: to_Referenda_Value0(field.Value),
    }
}  
func to_FellowshipReferenda_Signedsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Signed {
    return &pbgear.FellowshipReferenda_Signed{
                Value0: to_FellowshipReferenda_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_FellowshipReferenda_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.FellowshipReferenda_SpCoreCryptoAccountId32 {
                    return &pbgear.FellowshipReferenda_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_RemoveAnnouncementCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_RemoveAnnouncementCall {
    return &pbgear.Proxy_RemoveAnnouncementCall{
                Real: to_Proxy_Real(field.Value),
                CallHash: to_Proxy_PrimitiveTypesH256(field.Value),
    }
}   
func to_NominationPools_MinCreateBondsalut(fields []*registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
    return &pbgear.NominationPools_MinCreateBond{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Vesting_ForceVestedTransferCallsalut(fields []*registry.DecodedField) *pbgear.Vesting_ForceVestedTransferCall {
    return &pbgear.Vesting_ForceVestedTransferCall{
                Source: to_Vesting_Source(field.Value),
                Target: to_Vesting_Target(field.Value),
                Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(field.Value),
    }
}
            
                func to_Vesting_Source(in *registry.DecodedField) *pbgear.Vesting_Source {
                    return &pbgear.Vesting_Source{
                        // TODO fill all fields
                    }
                }  
            
                func to_Vesting_Target(in *registry.DecodedField) *pbgear.Vesting_Target {
                    return &pbgear.Vesting_Target{
                        // TODO fill all fields
                    }
                }  
            
                func to_Vesting_PalletVestingVestingInfoVestingInfo(in *registry.DecodedField) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
                    return &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{
                        // TODO fill all fields
                    }
                }   
func to_Session_SpCoreEd25519Publicsalut(fields []*registry.DecodedField) *pbgear.Session_SpCoreEd25519Public {
    return &pbgear.Session_SpCoreEd25519Public{
                Grandpa : []uint32(field.Value),
    }
}  
func to_Multisig_ApproveAsMultiCallsalut(fields []*registry.DecodedField) *pbgear.Multisig_ApproveAsMultiCall {
    return &pbgear.Multisig_ApproveAsMultiCall{
                Threshold : uint32(field.Value),
                OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(field.Value),
                MaybeTimepoint: to_optional_Multisig_PalletMultisigTimepoint(field.Value),
                CallHash : []uint32(field.Value),
                MaxWeight: to_Multisig_SpWeightsWeightV2Weight(field.Value),
    }
}
            
                  
            
                func to_repeated_Multisig_SpCoreCryptoAccountId32(in *registry.DecodedField) []*pbgear.Multisig_SpCoreCryptoAccountId32 {
                    return []*pbgear.Multisig_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }  
            
                func to_optional_Multisig_PalletMultisigTimepoint(in *registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
                    return &pbgear.Multisig_PalletMultisigTimepoint{
                        // TODO fill all fields
                    }
                }   
            
                func to_Multisig_SpWeightsWeightV2Weight(in *registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
                    return &pbgear.Multisig_SpWeightsWeightV2Weight{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_Value0salut(fields []*registry.DecodedField) *pbgear.Referenda_Value0 {
    return &pbgear.Referenda_Value0{
                Value: to_oneof_Referenda_Value(field.Value),
    }
}  
func to_Multisig_AsMultiCallsalut(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiCall {
    return &pbgear.Multisig_AsMultiCall{
                Threshold : uint32(field.Value),
                OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(field.Value),
                MaybeTimepoint: to_optional_Multisig_PalletMultisigTimepoint(field.Value),
                Call: to_oneof_Multisig_Call(field.Value),
                MaxWeight: to_Multisig_SpWeightsWeightV2Weight(field.Value),
    }
}   
            
                func to_oneof_Multisig_Call(in *registry.DecodedField) *pbgear.Multisig_Call {
                    return nil
                }    
func to_NominationPools_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.NominationPools_SpCoreCryptoAccountId32 {
    return &pbgear.NominationPools_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_GearVoucher_AppendProgramssalut(fields []*registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
    return &pbgear.GearVoucher_AppendPrograms{
                Value: to_oneof_GearVoucher_Value(field.Value),
    }
}
            
                func to_oneof_GearVoucher_Value(in *registry.DecodedField) *pbgear.GearVoucher_Value {
                    return nil
                }   
func to_ConvictionVoting_Idsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Id {
    return &pbgear.ConvictionVoting_Id{
                Value0: to_ConvictionVoting_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_ConvictionVoting_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.ConvictionVoting_SpCoreCryptoAccountId32 {
                    return &pbgear.ConvictionVoting_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_Address20salut(fields []*registry.DecodedField) *pbgear.ChildBounties_Address20 {
    return &pbgear.ChildBounties_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_OutOfDatesalut(fields []*registry.DecodedField) *pbgear.Identity_OutOfDate {
    return &pbgear.Identity_OutOfDate{
    }
} 
func to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
    return &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{
                Value: to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(field.Value),
    }
}
            
                func to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
                    return &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{
                        // TODO fill all fields
                    }
                }   
func to_Vesting_VestedTransferCallsalut(fields []*registry.DecodedField) *pbgear.Vesting_VestedTransferCall {
    return &pbgear.Vesting_VestedTransferCall{
                Target: to_Vesting_Target(field.Value),
                Schedule: to_Vesting_PalletVestingVestingInfoVestingInfo(field.Value),
    }
}   
func to_Referenda_NudgeReferendumCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_NudgeReferendumCall {
    return &pbgear.Referenda_NudgeReferendumCall{
                Index : uint32(field.Value),
    }
}  
func to_Identity_Raw27salut(fields []*registry.DecodedField) *pbgear.Identity_Raw27 {
    return &pbgear.Identity_Raw27{
                Value0 : []uint32(field.Value),
    }
}  
func to_GearVoucher_RevokeCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_RevokeCall {
    return &pbgear.GearVoucher_RevokeCall{
                Spender: to_GearVoucher_SpCoreCryptoAccountId32(field.Value),
                VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(field.Value),
    }
}   
func to_Staking_Payeesalut(fields []*registry.DecodedField) *pbgear.Staking_Payee {
    return &pbgear.Staking_Payee{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_ConvictionVoting_VoteCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_VoteCall {
    return &pbgear.ConvictionVoting_VoteCall{
                PollIndex: to_ConvictionVoting_CompactUint32(field.Value),
                Vote: to_ConvictionVoting_Vote(field.Value),
    }
}
            
                func to_ConvictionVoting_CompactUint32(in *registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
                    return &pbgear.ConvictionVoting_CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_ConvictionVoting_Vote(in *registry.DecodedField) *pbgear.ConvictionVoting_Vote {
                    return &pbgear.ConvictionVoting_Vote{
                        // TODO fill all fields
                    }
                }   
func to_Timestamp_SetCallsalut(fields []*registry.DecodedField) *pbgear.Timestamp_SetCall {
    return &pbgear.Timestamp_SetCall{
                Now: to_Timestamp_CompactUint64(field.Value),
    }
}
            
                func to_Timestamp_CompactUint64(in *registry.DecodedField) *pbgear.Timestamp_CompactUint64 {
                    return &pbgear.Timestamp_CompactUint64{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Proxy_TupleNull {
    return &pbgear.Proxy_TupleNull{
    }
} 
func to_Referenda_RefundSubmissionDepositCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_RefundSubmissionDepositCall {
    return &pbgear.Referenda_RefundSubmissionDepositCall{
                Index : uint32(field.Value),
    }
}  
func to_Babe_PrimaryAndSecondaryPlainSlotssalut(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
    return &pbgear.Babe_PrimaryAndSecondaryPlainSlots{
    }
} 
func to_Treasury_RemoveApprovalCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_RemoveApprovalCall {
    return &pbgear.Treasury_RemoveApprovalCall{
                ProposalId: to_Treasury_CompactUint32(field.Value),
    }
}
            
                func to_Treasury_CompactUint32(in *registry.DecodedField) *pbgear.Treasury_CompactUint32 {
                    return &pbgear.Treasury_CompactUint32{
                        // TODO fill all fields
                    }
                }   
func to_GearVoucher_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.GearVoucher_SpCoreCryptoAccountId32 {
    return &pbgear.GearVoucher_SpCoreCryptoAccountId32{
                Spender : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_TupleNullsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
    return &pbgear.ConvictionVoting_TupleNull{
    }
} 
func to_Gear_ClaimValueCallsalut(fields []*registry.DecodedField) *pbgear.Gear_ClaimValueCall {
    return &pbgear.Gear_ClaimValueCall{
                MessageId: to_Gear_GprimitivesMessageId(field.Value),
    }
}
            
                func to_Gear_GprimitivesMessageId(in *registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
                    return &pbgear.Gear_GprimitivesMessageId{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Nonesalut(fields []*registry.DecodedField) *pbgear.Identity_None {
    return &pbgear.Identity_None{
    }
} 
func to_Identity_AddSubCallsalut(fields []*registry.DecodedField) *pbgear.Identity_AddSubCall {
    return &pbgear.Identity_AddSubCall{
                Sub: to_Identity_Sub(field.Value),
                Data: to_Identity_Data(field.Value),
    }
}
            
                func to_Identity_Sub(in *registry.DecodedField) *pbgear.Identity_Sub {
                    return &pbgear.Identity_Sub{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Data(in *registry.DecodedField) *pbgear.Identity_Data {
                    return &pbgear.Identity_Data{
                        // TODO fill all fields
                    }
                }   
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_NominationPools_ClaimPayoutOtherCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutOtherCall {
    return &pbgear.NominationPools_ClaimPayoutOtherCall{
                Other: to_NominationPools_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Staking_ForceNewEraAlwaysCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraAlwaysCall {
    return &pbgear.Staking_ForceNewEraAlwaysCall{
    }
} 
func to_Identity_Rawsalut(fields []*registry.DecodedField) *pbgear.Identity_Raw {
    return &pbgear.Identity_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw7salut(fields []*registry.DecodedField) *pbgear.Identity_Raw7 {
    return &pbgear.Identity_Raw7{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_PermissionlessCompoundsalut(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessCompound {
    return &pbgear.NominationPools_PermissionlessCompound{
    }
} 
func to_Babe_Sealsalut(fields []*registry.DecodedField) *pbgear.Babe_Seal {
    return &pbgear.Babe_Seal{
                Value0 : []uint32(field.Value),
                Value1 : []uint32(field.Value),
    }
}   
func to_ConvictionVoting_Locked2Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked2X {
    return &pbgear.ConvictionVoting_Locked2X{
    }
} 
func to_Staking_MinNominatorBondsalut(fields []*registry.DecodedField) *pbgear.Staking_MinNominatorBond {
    return &pbgear.Staking_MinNominatorBond{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_Identity_Websalut(fields []*registry.DecodedField) *pbgear.Identity_Web {
    return &pbgear.Identity_Web{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Grandpa_SpConsensusGrandpaEquivocationProofsalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
    return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{
                SetId : uint64(field.Value),
                Equivocation: to_Grandpa_Equivocation(field.Value),
    }
}
            
                  
            
                func to_Grandpa_Equivocation(in *registry.DecodedField) *pbgear.Grandpa_Equivocation {
                    return &pbgear.Grandpa_Equivocation{
                        // TODO fill all fields
                    }
                }   
func to_Scheduler_ScheduleNamedAfterCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleNamedAfterCall {
    return &pbgear.Scheduler_ScheduleNamedAfterCall{
                Id : []uint32(field.Value),
                After : uint32(field.Value),
                MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(field.Value),
                Priority : uint32(field.Value),
                Call: to_oneof_Scheduler_Call(field.Value),
    }
}      
func to_FellowshipCollective_Whosalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Who {
    return &pbgear.FellowshipCollective_Who{
                Value: to_oneof_FellowshipCollective_Value(field.Value),
    }
}
            
                func to_oneof_FellowshipCollective_Value(in *registry.DecodedField) *pbgear.FellowshipCollective_Value {
                    return nil
                }   
func to_Whitelist_DispatchWhitelistedCallWithPreimageCallsalut(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
    return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{
                Call: to_oneof_Whitelist_Call(field.Value),
    }
}
            
                func to_oneof_Whitelist_Call(in *registry.DecodedField) *pbgear.Whitelist_Call {
                    return nil
                }   
func to_NominationPools_MinJoinBondsalut(fields []*registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
    return &pbgear.NominationPools_MinJoinBond{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Gear_SendMessageCallsalut(fields []*registry.DecodedField) *pbgear.Gear_SendMessageCall {
    return &pbgear.Gear_SendMessageCall{
                Destination: to_Gear_GprimitivesActorId(field.Value),
                Payload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}
            
                func to_Gear_GprimitivesActorId(in *registry.DecodedField) *pbgear.Gear_GprimitivesActorId {
                    return &pbgear.Gear_GprimitivesActorId{
                        // TODO fill all fields
                    }
                }   
            
                  
            
                    
func to_Utility_BatchAllCallsalut(fields []*registry.DecodedField) *pbgear.Utility_BatchAllCall {
    return &pbgear.Utility_BatchAllCall{
                Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(field.Value),
    }
}
            
                func to_repeated_Utility_VaraRuntimeRuntimeCall(in *registry.DecodedField) []*pbgear.Utility_VaraRuntimeRuntimeCall {
                    return []*pbgear.Utility_VaraRuntimeRuntimeCall{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_KillCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_KillCall {
    return &pbgear.Referenda_KillCall{
                Index : uint32(field.Value),
    }
}  
func to_BagsList_Indexsalut(fields []*registry.DecodedField) *pbgear.BagsList_Index {
    return &pbgear.BagsList_Index{
                Value0: to_BagsList_CompactTupleNull(field.Value),
    }
}
            
                func to_BagsList_CompactTupleNull(in *registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
                    return &pbgear.BagsList_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Identity_ClearIdentityCallsalut(fields []*registry.DecodedField) *pbgear.Identity_ClearIdentityCall {
    return &pbgear.Identity_ClearIdentityCall{
    }
} 
func to_ElectionProviderMultiPhase_SubmitCallsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitCall {
    return &pbgear.ElectionProviderMultiPhase_SubmitCall{
                RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(field.Value),
    }
}
            
                func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
                    return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{
                        // TODO fill all fields
                    }
                }   
func to_Balances_Address20salut(fields []*registry.DecodedField) *pbgear.Balances_Address20 {
    return &pbgear.Balances_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_Stakingsalut(fields []*registry.DecodedField) *pbgear.Proxy_Staking {
    return &pbgear.Proxy_Staking{
    }
} 
func to_Staking_PayoutStakersCallsalut(fields []*registry.DecodedField) *pbgear.Staking_PayoutStakersCall {
    return &pbgear.Staking_PayoutStakersCall{
                ValidatorStash: to_Staking_SpCoreCryptoAccountId32(field.Value),
                Era : uint32(field.Value),
    }
}
            
                func to_Staking_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
                    return &pbgear.Staking_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }  
            
                   
func to_Identity_Raw30salut(fields []*registry.DecodedField) *pbgear.Identity_Raw30 {
    return &pbgear.Identity_Raw30{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_SetIdentityCallsalut(fields []*registry.DecodedField) *pbgear.Identity_SetIdentityCall {
    return &pbgear.Identity_SetIdentityCall{
                Info: to_Identity_PalletIdentitySimpleIdentityInfo(field.Value),
    }
}
            
                func to_Identity_PalletIdentitySimpleIdentityInfo(in *registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
                    return &pbgear.Identity_PalletIdentitySimpleIdentityInfo{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_Permissionedsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Permissioned {
    return &pbgear.NominationPools_Permissioned{
    }
} 
func to_Gear_RunCallsalut(fields []*registry.DecodedField) *pbgear.Gear_RunCall {
    return &pbgear.Gear_RunCall{
                MaxGas : uint64(field.Value),
    }
}  
func to_System_SystemKeysListsalut(fields []*registry.DecodedField) *pbgear.System_SystemKeysList {
    return &pbgear.System_SystemKeysList{
                Keys : []uint32(field.Value),
    }
}  
func to_Vesting_Rawsalut(fields []*registry.DecodedField) *pbgear.Vesting_Raw {
    return &pbgear.Vesting_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Treasury_ProposeSpendCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_ProposeSpendCall {
    return &pbgear.Treasury_ProposeSpendCall{
                Value: to_Treasury_CompactString(field.Value),
                Beneficiary: to_Treasury_Beneficiary(field.Value),
    }
}
            
                func to_Treasury_CompactString(in *registry.DecodedField) *pbgear.Treasury_CompactString {
                    return &pbgear.Treasury_CompactString{
                        // TODO fill all fields
                    }
                }  
            
                func to_Treasury_Beneficiary(in *registry.DecodedField) *pbgear.Treasury_Beneficiary {
                    return &pbgear.Treasury_Beneficiary{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipReferenda_Aftersalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_After {
    return &pbgear.FellowshipReferenda_After{
                Value0 : uint32(field.Value),
    }
}  
func to_NominationPools_TupleNullsalut(fields []*registry.DecodedField) *pbgear.NominationPools_TupleNull {
    return &pbgear.NominationPools_TupleNull{
    }
} 
func to_NominationPools_SetCommissionChangeRateCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionChangeRateCall {
    return &pbgear.NominationPools_SetCommissionChangeRateCall{
                PoolId : uint32(field.Value),
                ChangeRate: to_NominationPools_PalletNominationPoolsCommissionChangeRate(field.Value),
    }
} 
            
                func to_NominationPools_PalletNominationPoolsCommissionChangeRate(in *registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
                    return &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{
                        // TODO fill all fields
                    }
                }   
func to_Staking_Stakedsalut(fields []*registry.DecodedField) *pbgear.Staking_Staked {
    return &pbgear.Staking_Staked{
    }
} 
func to_Staking_PalletStakingValidatorPrefssalut(fields []*registry.DecodedField) *pbgear.Staking_PalletStakingValidatorPrefs {
    return &pbgear.Staking_PalletStakingValidatorPrefs{
                Commission: to_Staking_CompactSpArithmeticPerThingsPerbill(field.Value),
                Blocked : bool(field.Value),
    }
}
            
                func to_Staking_CompactSpArithmeticPerThingsPerbill(in *registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
                    return &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{
                        // TODO fill all fields
                    }
                }  
            
                   
func to_FellowshipReferenda_PlaceDecisionDepositCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
    return &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{
                Index : uint32(field.Value),
    }
}  
func to_Identity_Raw22salut(fields []*registry.DecodedField) *pbgear.Identity_Raw22 {
    return &pbgear.Identity_Raw22{
                Value0 : []uint32(field.Value),
    }
}  
func to_Bounties_AcceptCuratorCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_AcceptCuratorCall {
    return &pbgear.Bounties_AcceptCuratorCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
    }
}  
func to_NominationPools_Rawsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Raw {
    return &pbgear.NominationPools_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_Noopsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Noop {
    return &pbgear.NominationPools_Noop{
    }
} 
func to_Treasury_SpendLocalCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_SpendLocalCall {
    return &pbgear.Treasury_SpendLocalCall{
                Amount: to_Treasury_CompactString(field.Value),
                Beneficiary: to_Treasury_Beneficiary(field.Value),
    }
}   
func to_Treasury_PayoutCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_PayoutCall {
    return &pbgear.Treasury_PayoutCall{
                Index : uint32(field.Value),
    }
}  
func to_Proxy_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Proxy_PrimitiveTypesH256 {
    return &pbgear.Proxy_PrimitiveTypesH256{
                CallHash : []uint32(field.Value),
    }
}  
func to_ChildBounties_AwardChildBountyCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_AwardChildBountyCall {
    return &pbgear.ChildBounties_AwardChildBountyCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
                Beneficiary: to_ChildBounties_Beneficiary(field.Value),
    }
}  
            
                func to_ChildBounties_Beneficiary(in *registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
                    return &pbgear.ChildBounties_Beneficiary{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
    return &pbgear.Referenda_PrimitiveTypesH256{
                Hash : []uint32(field.Value),
    }
}  
func to_Identity_Raw1salut(fields []*registry.DecodedField) *pbgear.Identity_Raw1 {
    return &pbgear.Identity_Raw1{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Proxy_SpCoreCryptoAccountId32 {
    return &pbgear.Proxy_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_WithdrawUnbondedCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_WithdrawUnbondedCall {
    return &pbgear.NominationPools_WithdrawUnbondedCall{
                MemberAccount: to_NominationPools_MemberAccount(field.Value),
                NumSlashingSpans : uint32(field.Value),
    }
}
            
                func to_NominationPools_MemberAccount(in *registry.DecodedField) *pbgear.NominationPools_MemberAccount {
                    return &pbgear.NominationPools_MemberAccount{
                        // TODO fill all fields
                    }
                }    
func to_ConvictionVoting_Indexsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Index {
    return &pbgear.ConvictionVoting_Index{
                Value0: to_ConvictionVoting_CompactTupleNull(field.Value),
    }
}
            
                func to_ConvictionVoting_CompactTupleNull(in *registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
                    return &pbgear.ConvictionVoting_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_SetMetadataCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_SetMetadataCall {
    return &pbgear.Referenda_SetMetadataCall{
                Index : uint32(field.Value),
                MaybeHash: to_optional_Referenda_PrimitiveTypesH256(field.Value),
    }
} 
            
                func to_optional_Referenda_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Referenda_PrimitiveTypesH256 {
                    return &pbgear.Referenda_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipCollective_AddMemberCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_AddMemberCall {
    return &pbgear.FellowshipCollective_AddMemberCall{
                Who: to_FellowshipCollective_Who(field.Value),
    }
}  
func to_Identity_ShaThree256salut(fields []*registry.DecodedField) *pbgear.Identity_ShaThree256 {
    return &pbgear.Identity_ShaThree256{
                Value0 : []uint32(field.Value),
    }
}  
func to_Vesting_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Vesting_SpCoreCryptoAccountId32 {
    return &pbgear.Vesting_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Utility_VaraRuntimeRuntimeCallsalut(fields []*registry.DecodedField) *pbgear.Utility_VaraRuntimeRuntimeCall {
    return &pbgear.Utility_VaraRuntimeRuntimeCall{
                Calls: to_oneof_Utility_Calls(field.Value),
    }
}
            
                func to_oneof_Utility_Calls(in *registry.DecodedField) *pbgear.Utility_Calls {
                    return nil
                }   
func to_FellowshipReferenda_Nonesalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_None {
    return &pbgear.FellowshipReferenda_None{
    }
} 
func to_Identity_Raw19salut(fields []*registry.DecodedField) *pbgear.Identity_Raw19 {
    return &pbgear.Identity_Raw19{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_BoundedCollectionsBoundedVecBoundedVecsalut(fields []*registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
    return &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{
                Additional: to_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(field.Value),
    }
}
            
                func to_repeated_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(in *registry.DecodedField) []*pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
                    return []*pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{
                        // TODO fill all fields
                    }
                }   
func to_Identity_KillIdentityCallsalut(fields []*registry.DecodedField) *pbgear.Identity_KillIdentityCall {
    return &pbgear.Identity_KillIdentityCall{
                Target: to_Identity_Target(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_ChildBounties_CompactUint32salut(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactUint32 {
    return &pbgear.ChildBounties_CompactUint32{
                Value : uint32(field.Value),
    }
}
            
                   
func to_Utility_Signedsalut(fields []*registry.DecodedField) *pbgear.Utility_Signed {
    return &pbgear.Utility_Signed{
                Value0: to_Utility_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Utility_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Utility_SpCoreCryptoAccountId32 {
                    return &pbgear.Utility_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_ConvictionVoting_Locked3Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked3X {
    return &pbgear.ConvictionVoting_Locked3X{
    }
} 
func to_NominationPools_PalletNominationPoolsCommissionChangeRatesalut(fields []*registry.DecodedField) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
    return &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{
                MaxIncrease: to_NominationPools_SpArithmeticPerThingsPerbill(field.Value),
                MinDelay : uint32(field.Value),
    }
}
            
                func to_NominationPools_SpArithmeticPerThingsPerbill(in *registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
                    return &pbgear.NominationPools_SpArithmeticPerThingsPerbill{
                        // TODO fill all fields
                    }
                }    
func to_Whitelist_SpWeightsWeightV2Weightsalut(fields []*registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
    return &pbgear.Whitelist_SpWeightsWeightV2Weight{
                RefTime: to_Whitelist_CompactUint64(field.Value),
                ProofSize: to_Whitelist_CompactUint64(field.Value),
    }
}
            
                func to_Whitelist_CompactUint64(in *registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
                    return &pbgear.Whitelist_CompactUint64{
                        // TODO fill all fields
                    }
                }    
func to_Proxy_RemoveProxiesCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_RemoveProxiesCall {
    return &pbgear.Proxy_RemoveProxiesCall{
    }
} 
func to_NominationPools_Bouncersalut(fields []*registry.DecodedField) *pbgear.NominationPools_Bouncer {
    return &pbgear.NominationPools_Bouncer{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_NewNominatorsalut(fields []*registry.DecodedField) *pbgear.NominationPools_NewNominator {
    return &pbgear.NominationPools_NewNominator{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Staking_ReapStashCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ReapStashCall {
    return &pbgear.Staking_ReapStashCall{
                Stash: to_Staking_SpCoreCryptoAccountId32(field.Value),
                NumSlashingSpans : uint32(field.Value),
    }
}   
func to_FellowshipReferenda_Systemsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_System {
    return &pbgear.FellowshipReferenda_System{
                Value0: to_FellowshipReferenda_Value0(field.Value),
    }
}  
func to_Utility_CompactUint64salut(fields []*registry.DecodedField) *pbgear.Utility_CompactUint64 {
    return &pbgear.Utility_CompactUint64{
                Value : uint64(field.Value),
    }
}
            
                   
func to_FellowshipReferenda_SubmitCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SubmitCall {
    return &pbgear.FellowshipReferenda_SubmitCall{
                ProposalOrigin: to_FellowshipReferenda_ProposalOrigin(field.Value),
                Proposal: to_FellowshipReferenda_Proposal(field.Value),
                EnactmentMoment: to_FellowshipReferenda_EnactmentMoment(field.Value),
    }
}
            
                func to_FellowshipReferenda_ProposalOrigin(in *registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
                    return &pbgear.FellowshipReferenda_ProposalOrigin{
                        // TODO fill all fields
                    }
                }  
            
                func to_FellowshipReferenda_Proposal(in *registry.DecodedField) *pbgear.FellowshipReferenda_Proposal {
                    return &pbgear.FellowshipReferenda_Proposal{
                        // TODO fill all fields
                    }
                }  
            
                func to_FellowshipReferenda_EnactmentMoment(in *registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
                    return &pbgear.FellowshipReferenda_EnactmentMoment{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw3salut(fields []*registry.DecodedField) *pbgear.Identity_Raw3 {
    return &pbgear.Identity_Raw3{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw25salut(fields []*registry.DecodedField) *pbgear.Identity_Raw25 {
    return &pbgear.Identity_Raw25{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_Realsalut(fields []*registry.DecodedField) *pbgear.Proxy_Real {
    return &pbgear.Proxy_Real{
                Value: to_oneof_Proxy_Value(field.Value),
    }
}
            
                func to_oneof_Proxy_Value(in *registry.DecodedField) *pbgear.Proxy_Value {
                    return nil
                }   
func to_Vesting_MergeSchedulesCallsalut(fields []*registry.DecodedField) *pbgear.Vesting_MergeSchedulesCall {
    return &pbgear.Vesting_MergeSchedulesCall{
                Schedule1Index : uint32(field.Value),
                Schedule2Index : uint32(field.Value),
    }
}   
func to_Treasury_ApproveProposalCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_ApproveProposalCall {
    return &pbgear.Treasury_ApproveProposalCall{
                ProposalId: to_Treasury_CompactUint32(field.Value),
    }
}  
func to_Identity_AddRegistrarCallsalut(fields []*registry.DecodedField) *pbgear.Identity_AddRegistrarCall {
    return &pbgear.Identity_AddRegistrarCall{
                Account: to_Identity_Account(field.Value),
    }
}
            
                func to_Identity_Account(in *registry.DecodedField) *pbgear.Identity_Account {
                    return &pbgear.Identity_Account{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_Indexsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_Index {
    return &pbgear.ChildBounties_Index{
                Value0: to_ChildBounties_CompactTupleNull(field.Value),
    }
}
            
                func to_ChildBounties_CompactTupleNull(in *registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
                    return &pbgear.ChildBounties_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_Removesalut(fields []*registry.DecodedField) *pbgear.NominationPools_Remove {
    return &pbgear.NominationPools_Remove{
    }
} 
func to_FellowshipCollective_RemoveMemberCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_RemoveMemberCall {
    return &pbgear.FellowshipCollective_RemoveMemberCall{
                Who: to_FellowshipCollective_Who(field.Value),
                MinRank : uint32(field.Value),
    }
} 
            
                   
func to_FellowshipReferenda_NudgeReferendumCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_NudgeReferendumCall {
    return &pbgear.FellowshipReferenda_NudgeReferendumCall{
                Index : uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_NominationPools_Permissionsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Permission {
    return &pbgear.NominationPools_Permission{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Babe_SpRuntimeGenericDigestDigestsalut(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
    return &pbgear.Babe_SpRuntimeGenericDigestDigest{
                Logs: to_repeated_Babe_SpRuntimeGenericDigestDigestItem(field.Value),
    }
}
            
                func to_repeated_Babe_SpRuntimeGenericDigestDigestItem(in *registry.DecodedField) []*pbgear.Babe_SpRuntimeGenericDigestDigestItem {
                    return []*pbgear.Babe_SpRuntimeGenericDigestDigestItem{
                        // TODO fill all fields
                    }
                }   
func to_Preimage_UnrequestPreimageCallsalut(fields []*registry.DecodedField) *pbgear.Preimage_UnrequestPreimageCall {
    return &pbgear.Preimage_UnrequestPreimageCall{
                Hash: to_Preimage_PrimitiveTypesH256(field.Value),
    }
}
            
                func to_Preimage_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Preimage_PrimitiveTypesH256 {
                    return &pbgear.Preimage_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw11salut(fields []*registry.DecodedField) *pbgear.Identity_Raw11 {
    return &pbgear.Identity_Raw11{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw32salut(fields []*registry.DecodedField) *pbgear.Identity_Raw32 {
    return &pbgear.Identity_Raw32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Multisig_AsMultiThreshold1Callsalut(fields []*registry.DecodedField) *pbgear.Multisig_AsMultiThreshold1Call {
    return &pbgear.Multisig_AsMultiThreshold1Call{
                OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(field.Value),
                Call: to_oneof_Multisig_Call(field.Value),
    }
}   
func to_NominationPools_ClaimCommissionCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimCommissionCall {
    return &pbgear.NominationPools_ClaimCommissionCall{
                PoolId : uint32(field.Value),
    }
}  
func to_Babe_SpSessionMembershipProofsalut(fields []*registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
    return &pbgear.Babe_SpSessionMembershipProof{
                Session : uint32(field.Value),
                TrieNodes: to_repeated_Babe_BabeTrieNodesList(field.Value),
                ValidatorCount : uint32(field.Value),
    }
}
            
                  
            
                func to_repeated_Babe_BabeTrieNodesList(in *registry.DecodedField) []*pbgear.Babe_BabeTrieNodesList {
                    return []*pbgear.Babe_BabeTrieNodesList{
                        // TODO fill all fields
                    }
                }    
func to_Staking_SpArithmeticPerThingsPerbillsalut(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPerbill {
    return &pbgear.Staking_SpArithmeticPerThingsPerbill{
                Value : uint32(field.Value),
    }
}  
func to_ChildBounties_CloseChildBountyCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_CloseChildBountyCall {
    return &pbgear.ChildBounties_CloseChildBountyCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
    }
}   
func to_NominationPools_Statesalut(fields []*registry.DecodedField) *pbgear.NominationPools_State {
    return &pbgear.NominationPools_State{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_ConvictionVoting_RemoveVoteCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveVoteCall {
    return &pbgear.ConvictionVoting_RemoveVoteCall{
                Class : uint32(field.Value),
                Index : uint32(field.Value),
    }
}
            
                func to_optional_ConvictionVoting_uint32(in *registry.DecodedField) *uint32 {
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
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_Session_SpConsensusBabeAppPublicsalut(fields []*registry.DecodedField) *pbgear.Session_SpConsensusBabeAppPublic {
    return &pbgear.Session_SpConsensusBabeAppPublic{
                Babe: to_Session_SpCoreSr25519Public(field.Value),
    }
}
            
                func to_Session_SpCoreSr25519Public(in *registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
                    return &pbgear.Session_SpCoreSr25519Public{
                        // TODO fill all fields
                    }
                }   
func to_Utility_Originssalut(fields []*registry.DecodedField) *pbgear.Utility_Origins {
    return &pbgear.Utility_Origins{
                Value0: to_Utility_Value0(field.Value),
    }
}  
func to_FellowshipCollective_Idsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Id {
    return &pbgear.FellowshipCollective_Id{
                Value0: to_FellowshipCollective_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_FellowshipCollective_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.FellowshipCollective_SpCoreCryptoAccountId32 {
                    return &pbgear.FellowshipCollective_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Bounties_Address32salut(fields []*registry.DecodedField) *pbgear.Bounties_Address32 {
    return &pbgear.Bounties_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_Extrasalut(fields []*registry.DecodedField) *pbgear.NominationPools_Extra {
    return &pbgear.NominationPools_Extra{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_Setsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Set {
    return &pbgear.NominationPools_Set{
                Value0: to_NominationPools_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Babe_RuntimeEnvironmentUpdatedsalut(fields []*registry.DecodedField) *pbgear.Babe_RuntimeEnvironmentUpdated {
    return &pbgear.Babe_RuntimeEnvironmentUpdated{
    }
} 
func to_Grandpa_Equivocationsalut(fields []*registry.DecodedField) *pbgear.Grandpa_Equivocation {
    return &pbgear.Grandpa_Equivocation{
                Value: to_oneof_Grandpa_Value(field.Value),
    }
}
            
                func to_oneof_Grandpa_Value(in *registry.DecodedField) *pbgear.Grandpa_Value {
                    return nil
                }   
func to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVecsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec {
    return &pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_NonTransfersalut(fields []*registry.DecodedField) *pbgear.Proxy_NonTransfer {
    return &pbgear.Proxy_NonTransfer{
    }
} 
func to_ChildBounties_TupleNullsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_TupleNull {
    return &pbgear.ChildBounties_TupleNull{
    }
} 
func to_NominationPools_SetCommissionCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionCall {
    return &pbgear.NominationPools_SetCommissionCall{
                PoolId : uint32(field.Value),
                NewCommission: to_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(field.Value),
    }
} 
            
                func to_optional_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
                    return &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Babe_Configsalut(fields []*registry.DecodedField) *pbgear.Babe_Config {
    return &pbgear.Babe_Config{
                Value: to_oneof_Babe_Value(field.Value),
    }
}
            
                func to_oneof_Babe_Value(in *registry.DecodedField) *pbgear.Babe_Value {
                    return nil
                }   
func to_Staking_WithdrawUnbondedCallsalut(fields []*registry.DecodedField) *pbgear.Staking_WithdrawUnbondedCall {
    return &pbgear.Staking_WithdrawUnbondedCall{
                NumSlashingSpans : uint32(field.Value),
    }
}  
func to_Identity_BlakeTwo256salut(fields []*registry.DecodedField) *pbgear.Identity_BlakeTwo256 {
    return &pbgear.Identity_BlakeTwo256{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Twittersalut(fields []*registry.DecodedField) *pbgear.Identity_Twitter {
    return &pbgear.Identity_Twitter{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Identity_Erroneoussalut(fields []*registry.DecodedField) *pbgear.Identity_Erroneous {
    return &pbgear.Identity_Erroneous{
    }
} 
func to_Identity_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Identity_PrimitiveTypesH256 {
    return &pbgear.Identity_PrimitiveTypesH256{
                Identity : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32Stringsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
    return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{
                Value0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(field.Value),
                Value1 : string(field.Value),
    }
}
            
                func to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
                    return &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }    
func to_Bounties_ProposeBountyCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_ProposeBountyCall {
    return &pbgear.Bounties_ProposeBountyCall{
                Value: to_Bounties_CompactString(field.Value),
                Description : []uint32(field.Value),
    }
}
            
                func to_Bounties_CompactString(in *registry.DecodedField) *pbgear.Bounties_CompactString {
                    return &pbgear.Bounties_CompactString{
                        // TODO fill all fields
                    }
                }    
func to_System_RemarkCallsalut(fields []*registry.DecodedField) *pbgear.System_RemarkCall {
    return &pbgear.System_RemarkCall{
                Remark : []uint32(field.Value),
    }
}  
func to_Grandpa_ReportEquivocationCallsalut(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationCall {
    return &pbgear.Grandpa_ReportEquivocationCall{
                EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(field.Value),
                KeyOwnerProof: to_Grandpa_SpSessionMembershipProof(field.Value),
    }
}
            
                func to_Grandpa_SpConsensusGrandpaEquivocationProof(in *registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
                    return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{
                        // TODO fill all fields
                    }
                }  
            
                func to_Grandpa_SpSessionMembershipProof(in *registry.DecodedField) *pbgear.Grandpa_SpSessionMembershipProof {
                    return &pbgear.Grandpa_SpSessionMembershipProof{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw31salut(fields []*registry.DecodedField) *pbgear.Identity_Raw31 {
    return &pbgear.Identity_Raw31{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_NominationPools_SetClaimPermissionCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetClaimPermissionCall {
    return &pbgear.NominationPools_SetClaimPermissionCall{
                Permission: to_NominationPools_Permission(field.Value),
    }
}
            
                func to_NominationPools_Permission(in *registry.DecodedField) *pbgear.NominationPools_Permission {
                    return &pbgear.NominationPools_Permission{
                        // TODO fill all fields
                    }
                }   
func to_BagsList_PutInFrontOfCallsalut(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfCall {
    return &pbgear.BagsList_PutInFrontOfCall{
                Lighter: to_BagsList_Lighter(field.Value),
    }
}
            
                func to_BagsList_Lighter(in *registry.DecodedField) *pbgear.BagsList_Lighter {
                    return &pbgear.BagsList_Lighter{
                        // TODO fill all fields
                    }
                }   
func to_Staking_ChillThresholdsalut(fields []*registry.DecodedField) *pbgear.Staking_ChillThreshold {
    return &pbgear.Staking_ChillThreshold{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_ConvictionVoting_Targetsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Target {
    return &pbgear.ConvictionVoting_Target{
                Value: to_oneof_ConvictionVoting_Value(field.Value),
    }
}  
func to_Referenda_Voidsalut(fields []*registry.DecodedField) *pbgear.Referenda_Void {
    return &pbgear.Referenda_Void{
                Value0: to_Referenda_Value0(field.Value),
    }
}  
func to_Referenda_Atsalut(fields []*registry.DecodedField) *pbgear.Referenda_At {
    return &pbgear.Referenda_At{
                Value0 : uint32(field.Value),
    }
}  
func to_Proxy_RejectAnnouncementCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_RejectAnnouncementCall {
    return &pbgear.Proxy_RejectAnnouncementCall{
                Delegate: to_Proxy_Delegate(field.Value),
                CallHash: to_Proxy_PrimitiveTypesH256(field.Value),
    }
}   
func to_Bounties_Curatorsalut(fields []*registry.DecodedField) *pbgear.Bounties_Curator {
    return &pbgear.Bounties_Curator{
                Value: to_oneof_Bounties_Value(field.Value),
    }
}
            
                func to_oneof_Bounties_Value(in *registry.DecodedField) *pbgear.Bounties_Value {
                    return nil
                }   
func to_Gear_UploadCodeCallsalut(fields []*registry.DecodedField) *pbgear.Gear_UploadCodeCall {
    return &pbgear.Gear_UploadCodeCall{
                Code : []uint32(field.Value),
    }
}  
func to_Treasury_Address32salut(fields []*registry.DecodedField) *pbgear.Treasury_Address32 {
    return &pbgear.Treasury_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Utility_BatchCallsalut(fields []*registry.DecodedField) *pbgear.Utility_BatchCall {
    return &pbgear.Utility_BatchCall{
                Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(field.Value),
    }
}  
func to_StakingRewards_TupleNullsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_TupleNull {
    return &pbgear.StakingRewards_TupleNull{
    }
} 
func to_Staking_IncreaseValidatorCountCallsalut(fields []*registry.DecodedField) *pbgear.Staking_IncreaseValidatorCountCall {
    return &pbgear.Staking_IncreaseValidatorCountCall{
                Additional: to_Staking_CompactUint32(field.Value),
    }
}
            
                func to_Staking_CompactUint32(in *registry.DecodedField) *pbgear.Staking_CompactUint32 {
                    return &pbgear.Staking_CompactUint32{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipCollective_VoteCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_VoteCall {
    return &pbgear.FellowshipCollective_VoteCall{
                Poll : uint32(field.Value),
                Aye : bool(field.Value),
    }
} 
            
                   
func to_NominationPools_UnbondCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_UnbondCall {
    return &pbgear.NominationPools_UnbondCall{
                MemberAccount: to_NominationPools_MemberAccount(field.Value),
                UnbondingPoints: to_NominationPools_CompactString(field.Value),
    }
}   
func to_StakingRewards_WithdrawCallsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_WithdrawCall {
    return &pbgear.StakingRewards_WithdrawCall{
                To: to_StakingRewards_To(field.Value),
                Value : string(field.Value),
    }
}
            
                func to_StakingRewards_To(in *registry.DecodedField) *pbgear.StakingRewards_To {
                    return &pbgear.StakingRewards_To{
                        // TODO fill all fields
                    }
                }  
            
                   
func to_Babe_PlanConfigChangeCallsalut(fields []*registry.DecodedField) *pbgear.Babe_PlanConfigChangeCall {
    return &pbgear.Babe_PlanConfigChangeCall{
                Config: to_Babe_Config(field.Value),
    }
}
            
                func to_Babe_Config(in *registry.DecodedField) *pbgear.Babe_Config {
                    return &pbgear.Babe_Config{
                        // TODO fill all fields
                    }
                }   
func to_Vesting_Indexsalut(fields []*registry.DecodedField) *pbgear.Vesting_Index {
    return &pbgear.Vesting_Index{
                Value0: to_Vesting_CompactTupleNull(field.Value),
    }
}
            
                func to_Vesting_CompactTupleNull(in *registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
                    return &pbgear.Vesting_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Session_SpAuthorityDiscoveryAppPublicsalut(fields []*registry.DecodedField) *pbgear.Session_SpAuthorityDiscoveryAppPublic {
    return &pbgear.Session_SpAuthorityDiscoveryAppPublic{
                AuthorityDiscovery: to_Session_SpCoreSr25519Public(field.Value),
    }
}  
func to_Utility_Nonesalut(fields []*registry.DecodedField) *pbgear.Utility_None {
    return &pbgear.Utility_None{
    }
} 
func to_Identity_LowQualitysalut(fields []*registry.DecodedField) *pbgear.Identity_LowQuality {
    return &pbgear.Identity_LowQuality{
    }
} 
func to_Babe_ReportEquivocationCallsalut(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationCall {
    return &pbgear.Babe_ReportEquivocationCall{
                EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(field.Value),
                KeyOwnerProof: to_Babe_SpSessionMembershipProof(field.Value),
    }
}
            
                func to_Babe_SpConsensusSlotsEquivocationProof(in *registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
                    return &pbgear.Babe_SpConsensusSlotsEquivocationProof{
                        // TODO fill all fields
                    }
                }  
            
                func to_Babe_SpSessionMembershipProof(in *registry.DecodedField) *pbgear.Babe_SpSessionMembershipProof {
                    return &pbgear.Babe_SpSessionMembershipProof{
                        // TODO fill all fields
                    }
                }   
func to_Balances_Sourcesalut(fields []*registry.DecodedField) *pbgear.Balances_Source {
    return &pbgear.Balances_Source{
                Value: to_oneof_Balances_Value(field.Value),
    }
}
            
                func to_oneof_Balances_Value(in *registry.DecodedField) *pbgear.Balances_Value {
                    return nil
                }   
func to_FellowshipReferenda_ProposalOriginsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_ProposalOrigin {
    return &pbgear.FellowshipReferenda_ProposalOrigin{
                Value: to_oneof_FellowshipReferenda_Value(field.Value),
    }
}  
func to_Identity_Idsalut(fields []*registry.DecodedField) *pbgear.Identity_Id {
    return &pbgear.Identity_Id{
                Value0: to_Identity_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Multisig_CancelAsMultiCallsalut(fields []*registry.DecodedField) *pbgear.Multisig_CancelAsMultiCall {
    return &pbgear.Multisig_CancelAsMultiCall{
                Threshold : uint32(field.Value),
                OtherSignatories: to_repeated_Multisig_SpCoreCryptoAccountId32(field.Value),
                Timepoint: to_Multisig_PalletMultisigTimepoint(field.Value),
                CallHash : []uint32(field.Value),
    }
}     
func to_Gear_GprimitivesCodeIdsalut(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
    return &pbgear.Gear_GprimitivesCodeId{
                CodeId : []uint32(field.Value),
    }
}  
func to_GearVoucher_UpdateCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_UpdateCall {
    return &pbgear.GearVoucher_UpdateCall{
                Spender: to_GearVoucher_SpCoreCryptoAccountId32(field.Value),
                VoucherId: to_GearVoucher_PalletGearVoucherInternalVoucherId(field.Value),
                MoveOwnership: to_optional_GearVoucher_SpCoreCryptoAccountId32(field.Value),
                BalanceTopUp : string(field.Value),
                AppendPrograms: to_optional_GearVoucher_AppendPrograms(field.Value),
                CodeUploading : bool(field.Value),
                ProlongDuration : uint32(field.Value),
    }
}    
            
                func to_optional_GearVoucher_AppendPrograms(in *registry.DecodedField) *pbgear.GearVoucher_AppendPrograms {
                    return &pbgear.GearVoucher_AppendPrograms{
                        // TODO fill all fields
                    }
                }     
func to_BagsList_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.BagsList_CompactTupleNull {
    return &pbgear.BagsList_CompactTupleNull{
                Value: to_BagsList_TupleNull(field.Value),
    }
}
            
                func to_BagsList_TupleNull(in *registry.DecodedField) *pbgear.BagsList_TupleNull {
                    return &pbgear.BagsList_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Treasury_CheckStatusCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_CheckStatusCall {
    return &pbgear.Treasury_CheckStatusCall{
                Index : uint32(field.Value),
    }
}  
func to_Whitelist_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Whitelist_PrimitiveTypesH256 {
    return &pbgear.Whitelist_PrimitiveTypesH256{
                CallHash : []uint32(field.Value),
    }
}  
func to_Whitelist_DispatchWhitelistedCallCallsalut(fields []*registry.DecodedField) *pbgear.Whitelist_DispatchWhitelistedCallCall {
    return &pbgear.Whitelist_DispatchWhitelistedCallCall{
                CallHash: to_Whitelist_PrimitiveTypesH256(field.Value),
                CallEncodedLen : uint32(field.Value),
                CallWeightWitness: to_Whitelist_SpWeightsWeightV2Weight(field.Value),
    }
} 
            
                  
            
                func to_Whitelist_SpWeightsWeightV2Weight(in *registry.DecodedField) *pbgear.Whitelist_SpWeightsWeightV2Weight {
                    return &pbgear.Whitelist_SpWeightsWeightV2Weight{
                        // TODO fill all fields
                    }
                }   
func to_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCallsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
    return &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{
                MaybeNextScore: to_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(field.Value),
    }
}
            
                func to_optional_ElectionProviderMultiPhase_SpNposElectionsElectionScore(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore {
                    return &pbgear.ElectionProviderMultiPhase_SpNposElectionsElectionScore{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_SetStateCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetStateCall {
    return &pbgear.NominationPools_SetStateCall{
                PoolId : uint32(field.Value),
                State: to_NominationPools_State(field.Value),
    }
} 
            
                func to_NominationPools_State(in *registry.DecodedField) *pbgear.NominationPools_State {
                    return &pbgear.NominationPools_State{
                        // TODO fill all fields
                    }
                }   
func to_Babe_SpRuntimeGenericDigestDigestItemsalut(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigestItem {
    return &pbgear.Babe_SpRuntimeGenericDigestDigestItem{
                Logs: to_Babe_Logs(field.Value),
    }
}
            
                func to_Babe_Logs(in *registry.DecodedField) *pbgear.Babe_Logs {
                    return &pbgear.Babe_Logs{
                        // TODO fill all fields
                    }
                }   
func to_Babe_ReportEquivocationUnsignedCallsalut(fields []*registry.DecodedField) *pbgear.Babe_ReportEquivocationUnsignedCall {
    return &pbgear.Babe_ReportEquivocationUnsignedCall{
                EquivocationProof: to_Babe_SpConsensusSlotsEquivocationProof(field.Value),
                KeyOwnerProof: to_Babe_SpSessionMembershipProof(field.Value),
    }
}   
func to_Babe_V1salut(fields []*registry.DecodedField) *pbgear.Babe_V1 {
    return &pbgear.Babe_V1{
                C: to_Babe_TupleUint64Uint64(field.Value),
                AllowedSlots: to_Babe_AllowedSlots(field.Value),
    }
}
            
                func to_Babe_TupleUint64Uint64(in *registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
                    return &pbgear.Babe_TupleUint64Uint64{
                        // TODO fill all fields
                    }
                }  
            
                func to_Babe_AllowedSlots(in *registry.DecodedField) *pbgear.Babe_AllowedSlots {
                    return &pbgear.Babe_AllowedSlots{
                        // TODO fill all fields
                    }
                }   
func to_Identity_RemoveSubCallsalut(fields []*registry.DecodedField) *pbgear.Identity_RemoveSubCall {
    return &pbgear.Identity_RemoveSubCall{
                Sub: to_Identity_Sub(field.Value),
    }
}  
func to_FellowshipCollective_Address20salut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address20 {
    return &pbgear.FellowshipCollective_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_SetSubsCallsalut(fields []*registry.DecodedField) *pbgear.Identity_SetSubsCall {
    return &pbgear.Identity_SetSubsCall{
                Subs: to_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(field.Value),
    }
}
            
                func to_repeated_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(in *registry.DecodedField) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
                    return []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_IdentityJudgementsalut(fields []*registry.DecodedField) *pbgear.Proxy_IdentityJudgement {
    return &pbgear.Proxy_IdentityJudgement{
    }
} 
func to_Balances_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Balances_CompactTupleNull {
    return &pbgear.Balances_CompactTupleNull{
                Value: to_Balances_TupleNull(field.Value),
    }
}
            
                func to_Balances_TupleNull(in *registry.DecodedField) *pbgear.Balances_TupleNull {
                    return &pbgear.Balances_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Balances_UpgradeAccountsCallsalut(fields []*registry.DecodedField) *pbgear.Balances_UpgradeAccountsCall {
    return &pbgear.Balances_UpgradeAccountsCall{
                Who: to_repeated_Balances_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_repeated_Balances_SpCoreCryptoAccountId32(in *registry.DecodedField) []*pbgear.Balances_SpCoreCryptoAccountId32 {
                    return []*pbgear.Balances_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Vesting_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Vesting_TupleNull {
    return &pbgear.Vesting_TupleNull{
    }
} 
func to_Identity_Raw15salut(fields []*registry.DecodedField) *pbgear.Identity_Raw15 {
    return &pbgear.Identity_Raw15{
                Value0 : []uint32(field.Value),
    }
}  
func to_Balances_ForceUnreserveCallsalut(fields []*registry.DecodedField) *pbgear.Balances_ForceUnreserveCall {
    return &pbgear.Balances_ForceUnreserveCall{
                Who: to_Balances_Who(field.Value),
                Amount : string(field.Value),
    }
} 
            
                   
func to_Vesting_Address20salut(fields []*registry.DecodedField) *pbgear.Vesting_Address20 {
    return &pbgear.Vesting_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_Session_SpCoreSr25519Publicsalut(fields []*registry.DecodedField) *pbgear.Session_SpCoreSr25519Public {
    return &pbgear.Session_SpCoreSr25519Public{
                Babe : []uint32(field.Value),
    }
}  
func to_Identity_RenameSubCallsalut(fields []*registry.DecodedField) *pbgear.Identity_RenameSubCall {
    return &pbgear.Identity_RenameSubCall{
                Sub: to_Identity_Sub(field.Value),
                Data: to_Identity_Data(field.Value),
    }
}   
func to_Multisig_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Multisig_SpCoreCryptoAccountId32 {
    return &pbgear.Multisig_SpCoreCryptoAccountId32{
                OtherSignatories: to_Multisig_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Balances_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Balances_TupleNull {
    return &pbgear.Balances_TupleNull{
    }
} 
func to_Staking_Noopsalut(fields []*registry.DecodedField) *pbgear.Staking_Noop {
    return &pbgear.Staking_Noop{
    }
} 
func to_Proxy_ProxyCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_ProxyCall {
    return &pbgear.Proxy_ProxyCall{
                Real: to_Proxy_Real(field.Value),
                ForceProxyType: to_optional_Proxy_ForceProxyType(field.Value),
                Call: to_oneof_Proxy_Call(field.Value),
    }
}    
func to_Bounties_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Bounties_SpCoreCryptoAccountId32 {
    return &pbgear.Bounties_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_Splitsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Split {
    return &pbgear.ConvictionVoting_Split{
                Aye : string(field.Value),
                Nay : string(field.Value),
    }
}
            
                    
func to_FellowshipReferenda_Lookupsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Lookup {
    return &pbgear.FellowshipReferenda_Lookup{
                Hash: to_FellowshipReferenda_PrimitiveTypesH256(field.Value),
                Len : uint32(field.Value),
    }
}
            
                func to_FellowshipReferenda_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
                    return &pbgear.FellowshipReferenda_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }    
func to_BagsList_Idsalut(fields []*registry.DecodedField) *pbgear.BagsList_Id {
    return &pbgear.BagsList_Id{
                Value0: to_BagsList_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_BagsList_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.BagsList_SpCoreCryptoAccountId32 {
                    return &pbgear.BagsList_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_ConvictionVoting_Locked5Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked5X {
    return &pbgear.ConvictionVoting_Locked5X{
    }
} 
func to_FellowshipReferenda_Atsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_At {
    return &pbgear.FellowshipReferenda_At{
                Value0 : uint32(field.Value),
    }
}  
func to_Identity_Newsalut(fields []*registry.DecodedField) *pbgear.Identity_New {
    return &pbgear.Identity_New{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_NominationPools_Idsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Id {
    return &pbgear.NominationPools_Id{
                Value0: to_NominationPools_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Babe_Othersalut(fields []*registry.DecodedField) *pbgear.Babe_Other {
    return &pbgear.Babe_Other{
                Value0 : []uint32(field.Value),
    }
}  
func to_Babe_SpRuntimeGenericHeaderHeadersalut(fields []*registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
    return &pbgear.Babe_SpRuntimeGenericHeaderHeader{
                ParentHash: to_Babe_PrimitiveTypesH256(field.Value),
                Number: to_Babe_CompactUint32(field.Value),
                StateRoot: to_Babe_PrimitiveTypesH256(field.Value),
                ExtrinsicsRoot: to_Babe_PrimitiveTypesH256(field.Value),
                Digest: to_Babe_SpRuntimeGenericDigestDigest(field.Value),
    }
}
            
                func to_Babe_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
                    return &pbgear.Babe_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }  
            
                func to_Babe_CompactUint32(in *registry.DecodedField) *pbgear.Babe_CompactUint32 {
                    return &pbgear.Babe_CompactUint32{
                        // TODO fill all fields
                    }
                }    
            
                func to_Babe_SpRuntimeGenericDigestDigest(in *registry.DecodedField) *pbgear.Babe_SpRuntimeGenericDigestDigest {
                    return &pbgear.Babe_SpRuntimeGenericDigestDigest{
                        // TODO fill all fields
                    }
                }   
func to_Staking_CancelDeferredSlashCallsalut(fields []*registry.DecodedField) *pbgear.Staking_CancelDeferredSlashCall {
    return &pbgear.Staking_CancelDeferredSlashCall{
                Era : uint32(field.Value),
                SlashIndices : []uint32(field.Value),
    }
}   
func to_Staking_ForceApplyMinCommissionCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ForceApplyMinCommissionCall {
    return &pbgear.Staking_ForceApplyMinCommissionCall{
                ValidatorStash: to_Staking_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Treasury_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
    return &pbgear.Treasury_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Referenda_Signedsalut(fields []*registry.DecodedField) *pbgear.Referenda_Signed {
    return &pbgear.Referenda_Signed{
                Value0: to_Referenda_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Referenda_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Referenda_SpCoreCryptoAccountId32 {
                    return &pbgear.Referenda_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Whitelist_CompactUint64salut(fields []*registry.DecodedField) *pbgear.Whitelist_CompactUint64 {
    return &pbgear.Whitelist_CompactUint64{
                Value : uint64(field.Value),
    }
}
            
                   
func to_Identity_KnownGoodsalut(fields []*registry.DecodedField) *pbgear.Identity_KnownGood {
    return &pbgear.Identity_KnownGood{
    }
} 
func to_Balances_Destsalut(fields []*registry.DecodedField) *pbgear.Balances_Dest {
    return &pbgear.Balances_Dest{
                Value: to_oneof_Balances_Value(field.Value),
    }
}  
func to_Staking_Stashsalut(fields []*registry.DecodedField) *pbgear.Staking_Stash {
    return &pbgear.Staking_Stash{
    }
} 
func to_Proxy_Indexsalut(fields []*registry.DecodedField) *pbgear.Proxy_Index {
    return &pbgear.Proxy_Index{
                Value0: to_Proxy_CompactTupleNull(field.Value),
    }
}
            
                func to_Proxy_CompactTupleNull(in *registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
                    return &pbgear.Proxy_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_UpdateRolesCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_UpdateRolesCall {
    return &pbgear.NominationPools_UpdateRolesCall{
                PoolId : uint32(field.Value),
                NewRoot: to_NominationPools_NewRoot(field.Value),
                NewNominator: to_NominationPools_NewNominator(field.Value),
                NewBouncer: to_NominationPools_NewBouncer(field.Value),
    }
} 
            
                func to_NominationPools_NewRoot(in *registry.DecodedField) *pbgear.NominationPools_NewRoot {
                    return &pbgear.NominationPools_NewRoot{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_NewNominator(in *registry.DecodedField) *pbgear.NominationPools_NewNominator {
                    return &pbgear.NominationPools_NewNominator{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_NewBouncer(in *registry.DecodedField) *pbgear.NominationPools_NewBouncer {
                    return &pbgear.NominationPools_NewBouncer{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_CreatePureCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_CreatePureCall {
    return &pbgear.Proxy_CreatePureCall{
                ProxyType: to_Proxy_ProxyType(field.Value),
                Delay : uint32(field.Value),
                Index : uint32(field.Value),
    }
}    
func to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(field.Value),
    }
} 
            
                func to_ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16 {
                    return &pbgear.ElectionProviderMultiPhase_CompactSpArithmeticPerThingsPerU16{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactTupleNull {
    return &pbgear.ChildBounties_CompactTupleNull{
                Value: to_ChildBounties_TupleNull(field.Value),
    }
}
            
                func to_ChildBounties_TupleNull(in *registry.DecodedField) *pbgear.ChildBounties_TupleNull {
                    return &pbgear.ChildBounties_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_GearVoucher_Callsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_Call {
    return &pbgear.GearVoucher_Call{
                Value: to_oneof_GearVoucher_Value(field.Value),
    }
}  
func to_Babe_SpConsensusBabeAppPublicsalut(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
    return &pbgear.Babe_SpConsensusBabeAppPublic{
                Offender: to_Babe_SpCoreSr25519Public(field.Value),
    }
}
            
                func to_Babe_SpCoreSr25519Public(in *registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
                    return &pbgear.Babe_SpCoreSr25519Public{
                        // TODO fill all fields
                    }
                }   
func to_ConvictionVoting_Rawsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Raw {
    return &pbgear.ConvictionVoting_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
    return &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{
                Votes1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(field.Value),
                Votes2: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(field.Value),
                Votes3: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(field.Value),
                Votes4: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(field.Value),
                Votes5: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(field.Value),
                Votes6: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(field.Value),
                Votes7: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(field.Value),
                Votes8: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(field.Value),
                Votes9: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(field.Value),
                Votes10: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(field.Value),
                Votes11: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(field.Value),
                Votes12: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(field.Value),
                Votes13: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(field.Value),
                Votes14: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(field.Value),
                Votes15: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(field.Value),
                Votes16: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(field.Value),
    }
}
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes3ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes5ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes6ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes7ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes9ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes10ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes12ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes13ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{
                        // TODO fill all fields
                    }
                }  
            
                func to_repeated_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32 {
                    return []*pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes16ListCompactUint32{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.NominationPools_CompactTupleNull {
    return &pbgear.NominationPools_CompactTupleNull{
                Value: to_NominationPools_TupleNull(field.Value),
    }
}
            
                func to_NominationPools_TupleNull(in *registry.DecodedField) *pbgear.NominationPools_TupleNull {
                    return &pbgear.NominationPools_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Vesting_Sourcesalut(fields []*registry.DecodedField) *pbgear.Vesting_Source {
    return &pbgear.Vesting_Source{
                Value: to_oneof_Vesting_Value(field.Value),
    }
}
            
                func to_oneof_Vesting_Value(in *registry.DecodedField) *pbgear.Vesting_Value {
                    return nil
                }   
func to_Referenda_PlaceDecisionDepositCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_PlaceDecisionDepositCall {
    return &pbgear.Referenda_PlaceDecisionDepositCall{
                Index : uint32(field.Value),
    }
}  
func to_GearVoucher_DeclineVouchersalut(fields []*registry.DecodedField) *pbgear.GearVoucher_DeclineVoucher {
    return &pbgear.GearVoucher_DeclineVoucher{
    }
} 
func to_NominationPools_ClaimPayoutCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_ClaimPayoutCall {
    return &pbgear.NominationPools_ClaimPayoutCall{
    }
} 
func to_NominationPools_NewBouncersalut(fields []*registry.DecodedField) *pbgear.NominationPools_NewBouncer {
    return &pbgear.NominationPools_NewBouncer{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Treasury_Rawsalut(fields []*registry.DecodedField) *pbgear.Treasury_Raw {
    return &pbgear.Treasury_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_DelegateCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_DelegateCall {
    return &pbgear.ConvictionVoting_DelegateCall{
                Class : uint32(field.Value),
                To: to_ConvictionVoting_To(field.Value),
                Conviction: to_ConvictionVoting_Conviction(field.Value),
                Balance : string(field.Value),
    }
} 
            
                func to_ConvictionVoting_To(in *registry.DecodedField) *pbgear.ConvictionVoting_To {
                    return &pbgear.ConvictionVoting_To{
                        // TODO fill all fields
                    }
                }  
            
                func to_ConvictionVoting_Conviction(in *registry.DecodedField) *pbgear.ConvictionVoting_Conviction {
                    return &pbgear.ConvictionVoting_Conviction{
                        // TODO fill all fields
                    }
                }    
func to_FellowshipReferenda_Value0salut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Value0 {
    return &pbgear.FellowshipReferenda_Value0{
                Value: to_oneof_FellowshipReferenda_Value(field.Value),
    }
}  
func to_FellowshipReferenda_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_PrimitiveTypesH256 {
    return &pbgear.FellowshipReferenda_PrimitiveTypesH256{
                Hash : []uint32(field.Value),
    }
}  
func to_Identity_Raw21salut(fields []*registry.DecodedField) *pbgear.Identity_Raw21 {
    return &pbgear.Identity_Raw21{
                Value0 : []uint32(field.Value),
    }
}  
func to_ChildBounties_Beneficiarysalut(fields []*registry.DecodedField) *pbgear.ChildBounties_Beneficiary {
    return &pbgear.ChildBounties_Beneficiary{
                Value: to_oneof_ChildBounties_Value(field.Value),
    }
}
            
                func to_oneof_ChildBounties_Value(in *registry.DecodedField) *pbgear.ChildBounties_Value {
                    return nil
                }   
func to_Grandpa_GrandpaTrieNodesListsalut(fields []*registry.DecodedField) *pbgear.Grandpa_GrandpaTrieNodesList {
    return &pbgear.Grandpa_GrandpaTrieNodesList{
                TrieNodes : []uint32(field.Value),
    }
}  
func to_Vesting_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Vesting_CompactTupleNull {
    return &pbgear.Vesting_CompactTupleNull{
                Value: to_Vesting_TupleNull(field.Value),
    }
}
            
                func to_Vesting_TupleNull(in *registry.DecodedField) *pbgear.Vesting_TupleNull {
                    return &pbgear.Vesting_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_StakingRewards_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.StakingRewards_SpCoreCryptoAccountId32 {
    return &pbgear.StakingRewards_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_GearVoucher_SendReplysalut(fields []*registry.DecodedField) *pbgear.GearVoucher_SendReply {
    return &pbgear.GearVoucher_SendReply{
                ReplyToId: to_GearVoucher_GprimitivesMessageId(field.Value),
                Payload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}
            
                func to_GearVoucher_GprimitivesMessageId(in *registry.DecodedField) *pbgear.GearVoucher_GprimitivesMessageId {
                    return &pbgear.GearVoucher_GprimitivesMessageId{
                        // TODO fill all fields
                    }
                }   
            
                     
func to_Staking_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Staking_TupleNull {
    return &pbgear.Staking_TupleNull{
    }
} 
func to_Identity_Raw12salut(fields []*registry.DecodedField) *pbgear.Identity_Raw12 {
    return &pbgear.Identity_Raw12{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Proxy_CompactTupleNull {
    return &pbgear.Proxy_CompactTupleNull{
                Value: to_Proxy_TupleNull(field.Value),
    }
}
            
                func to_Proxy_TupleNull(in *registry.DecodedField) *pbgear.Proxy_TupleNull {
                    return &pbgear.Proxy_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_SpArithmeticPerThingsPerbillsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SpArithmeticPerThingsPerbill {
    return &pbgear.NominationPools_SpArithmeticPerThingsPerbill{
                Value0 : uint32(field.Value),
    }
}  
func to_Grandpa_SpConsensusGrandpaAppPublicsalut(fields []*registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
    return &pbgear.Grandpa_SpConsensusGrandpaAppPublic{
                Identity: to_Grandpa_SpCoreEd25519Public(field.Value),
    }
}
            
                func to_Grandpa_SpCoreEd25519Public(in *registry.DecodedField) *pbgear.Grandpa_SpCoreEd25519Public {
                    return &pbgear.Grandpa_SpCoreEd25519Public{
                        // TODO fill all fields
                    }
                }   
func to_Staking_UnbondCallsalut(fields []*registry.DecodedField) *pbgear.Staking_UnbondCall {
    return &pbgear.Staking_UnbondCall{
                Value: to_Staking_CompactString(field.Value),
    }
}
            
                func to_Staking_CompactString(in *registry.DecodedField) *pbgear.Staking_CompactString {
                    return &pbgear.Staking_CompactString{
                        // TODO fill all fields
                    }
                }   
func to_Staking_SpRuntimeMultiaddressMultiAddresssalut(fields []*registry.DecodedField) *pbgear.Staking_SpRuntimeMultiaddressMultiAddress {
    return &pbgear.Staking_SpRuntimeMultiaddressMultiAddress{
                Targets: to_Staking_Targets(field.Value),
    }
}
            
                func to_Staking_Targets(in *registry.DecodedField) *pbgear.Staking_Targets {
                    return &pbgear.Staking_Targets{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Subsalut(fields []*registry.DecodedField) *pbgear.Identity_Sub {
    return &pbgear.Identity_Sub{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Identity_Raw13salut(fields []*registry.DecodedField) *pbgear.Identity_Raw13 {
    return &pbgear.Identity_Raw13{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw16salut(fields []*registry.DecodedField) *pbgear.Identity_Raw16 {
    return &pbgear.Identity_Raw16{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw20salut(fields []*registry.DecodedField) *pbgear.Identity_Raw20 {
    return &pbgear.Identity_Raw20{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_SubmitUnsignedCallsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
    return &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{
                RawSolution: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(field.Value),
                Witness: to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(field.Value),
    }
} 
            
                func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
                    return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_NewRootsalut(fields []*registry.DecodedField) *pbgear.NominationPools_NewRoot {
    return &pbgear.NominationPools_NewRoot{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_StakingRewards_RefillCallsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_RefillCall {
    return &pbgear.StakingRewards_RefillCall{
                Value : string(field.Value),
    }
}  
func to_Treasury_Idsalut(fields []*registry.DecodedField) *pbgear.Treasury_Id {
    return &pbgear.Treasury_Id{
                Value0: to_Treasury_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Treasury_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Treasury_SpCoreCryptoAccountId32 {
                    return &pbgear.Treasury_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Referenda_SpCoreCryptoAccountId32 {
    return &pbgear.Referenda_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32 {
    return &pbgear.ElectionProviderMultiPhase_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_BagsList_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.BagsList_SpCoreCryptoAccountId32 {
    return &pbgear.BagsList_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw28salut(fields []*registry.DecodedField) *pbgear.Identity_Raw28 {
    return &pbgear.Identity_Raw28{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_UnlockCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UnlockCall {
    return &pbgear.ConvictionVoting_UnlockCall{
                Class : uint32(field.Value),
                Target: to_ConvictionVoting_Target(field.Value),
    }
} 
            
                func to_ConvictionVoting_Target(in *registry.DecodedField) *pbgear.ConvictionVoting_Target {
                    return &pbgear.ConvictionVoting_Target{
                        // TODO fill all fields
                    }
                }   
func to_Bounties_Rawsalut(fields []*registry.DecodedField) *pbgear.Bounties_Raw {
    return &pbgear.Bounties_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_ChildBounties_Idsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_Id {
    return &pbgear.ChildBounties_Id{
                Value0: to_ChildBounties_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_ChildBounties_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.ChildBounties_SpCoreCryptoAccountId32 {
                    return &pbgear.ChildBounties_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_PermissionlessAllsalut(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessAll {
    return &pbgear.NominationPools_PermissionlessAll{
    }
} 
func to_Staking_BondCallsalut(fields []*registry.DecodedField) *pbgear.Staking_BondCall {
    return &pbgear.Staking_BondCall{
                Value: to_Staking_CompactString(field.Value),
                Payee: to_Staking_Payee(field.Value),
    }
}   
func to_Staking_ChillCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ChillCall {
    return &pbgear.Staking_ChillCall{
    }
} 
func to_Identity_Raw8salut(fields []*registry.DecodedField) *pbgear.Identity_Raw8 {
    return &pbgear.Identity_Raw8{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw17salut(fields []*registry.DecodedField) *pbgear.Identity_Raw17 {
    return &pbgear.Identity_Raw17{
                Value0 : []uint32(field.Value),
    }
}  
func to_Bounties_ClaimBountyCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_ClaimBountyCall {
    return &pbgear.Bounties_ClaimBountyCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
    }
}  
func to_ChildBounties_ProposeCuratorCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_ProposeCuratorCall {
    return &pbgear.ChildBounties_ProposeCuratorCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
                Curator: to_ChildBounties_Curator(field.Value),
                Fee: to_ChildBounties_CompactString(field.Value),
    }
}  
            
                func to_ChildBounties_Curator(in *registry.DecodedField) *pbgear.ChildBounties_Curator {
                    return &pbgear.ChildBounties_Curator{
                        // TODO fill all fields
                    }
                }  
            
                func to_ChildBounties_CompactString(in *registry.DecodedField) *pbgear.ChildBounties_CompactString {
                    return &pbgear.ChildBounties_CompactString{
                        // TODO fill all fields
                    }
                }   
func to_GearVoucher_Somesalut(fields []*registry.DecodedField) *pbgear.GearVoucher_Some {
    return &pbgear.GearVoucher_Some{
                Value0: to_GearVoucher_BTreeSet(field.Value),
    }
}  
func to_Balances_TransferKeepAliveCallsalut(fields []*registry.DecodedField) *pbgear.Balances_TransferKeepAliveCall {
    return &pbgear.Balances_TransferKeepAliveCall{
                Dest: to_Balances_Dest(field.Value),
                Value: to_Balances_CompactString(field.Value),
    }
}   
func to_Staking_ForceUnstakeCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ForceUnstakeCall {
    return &pbgear.Staking_ForceUnstakeCall{
                Stash: to_Staking_SpCoreCryptoAccountId32(field.Value),
                NumSlashingSpans : uint32(field.Value),
    }
}   
func to_FellowshipCollective_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
    return &pbgear.FellowshipCollective_CompactTupleNull{
                Value: to_FellowshipCollective_TupleNull(field.Value),
    }
}
            
                func to_FellowshipCollective_TupleNull(in *registry.DecodedField) *pbgear.FellowshipCollective_TupleNull {
                    return &pbgear.FellowshipCollective_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Proxy_CancelProxysalut(fields []*registry.DecodedField) *pbgear.Proxy_CancelProxy {
    return &pbgear.Proxy_CancelProxy{
    }
} 
func to_Bounties_UnassignCuratorCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_UnassignCuratorCall {
    return &pbgear.Bounties_UnassignCuratorCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
    }
}  
func to_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    return &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{
                Value0: to_NominationPools_SpArithmeticPerThingsPerbill(field.Value),
                Value1: to_NominationPools_SpCoreCryptoAccountId32(field.Value),
    }
}   
func to_Balances_Idsalut(fields []*registry.DecodedField) *pbgear.Balances_Id {
    return &pbgear.Balances_Id{
                Value0: to_Balances_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Staking_ForceNewEraCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ForceNewEraCall {
    return &pbgear.Staking_ForceNewEraCall{
    }
} 
func to_ChildBounties_AddChildBountyCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_AddChildBountyCall {
    return &pbgear.ChildBounties_AddChildBountyCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                Value: to_ChildBounties_CompactString(field.Value),
                Description : []uint32(field.Value),
    }
}    
func to_Referenda_OneFewerDecidingCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_OneFewerDecidingCall {
    return &pbgear.Referenda_OneFewerDecidingCall{
                Track : uint32(field.Value),
    }
}  
func to_Whitelist_RemoveWhitelistedCallCallsalut(fields []*registry.DecodedField) *pbgear.Whitelist_RemoveWhitelistedCallCall {
    return &pbgear.Whitelist_RemoveWhitelistedCallCall{
                CallHash: to_Whitelist_PrimitiveTypesH256(field.Value),
    }
}  
func to_Vesting_VestOtherCallsalut(fields []*registry.DecodedField) *pbgear.Vesting_VestOtherCall {
    return &pbgear.Vesting_VestOtherCall{
                Target: to_Vesting_Target(field.Value),
    }
}  
func to_ImOnline_HeartbeatCallsalut(fields []*registry.DecodedField) *pbgear.ImOnline_HeartbeatCall {
    return &pbgear.ImOnline_HeartbeatCall{
                Heartbeat: to_ImOnline_PalletImOnlineHeartbeat(field.Value),
                Signature: to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(field.Value),
    }
}
            
                func to_ImOnline_PalletImOnlineHeartbeat(in *registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
                    return &pbgear.ImOnline_PalletImOnlineHeartbeat{
                        // TODO fill all fields
                    }
                }  
            
                func to_ImOnline_PalletImOnlineSr25519AppSr25519Signature(in *registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
                    return &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{
                        // TODO fill all fields
                    }
                }   
func to_Staking_BondExtraCallsalut(fields []*registry.DecodedField) *pbgear.Staking_BondExtraCall {
    return &pbgear.Staking_BondExtraCall{
                MaxAdditional: to_Staking_CompactString(field.Value),
    }
}  
func to_ConvictionVoting_Nonesalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_None {
    return &pbgear.ConvictionVoting_None{
    }
} 
func to_Referenda_Lookupsalut(fields []*registry.DecodedField) *pbgear.Referenda_Lookup {
    return &pbgear.Referenda_Lookup{
                Hash: to_Referenda_PrimitiveTypesH256(field.Value),
                Len : uint32(field.Value),
    }
}   
func to_Identity_SetFeeCallsalut(fields []*registry.DecodedField) *pbgear.Identity_SetFeeCall {
    return &pbgear.Identity_SetFeeCall{
                Index: to_Identity_CompactUint32(field.Value),
                Fee: to_Identity_CompactString(field.Value),
    }
} 
            
                func to_Identity_CompactString(in *registry.DecodedField) *pbgear.Identity_CompactString {
                    return &pbgear.Identity_CompactString{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_SetMetadataCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetMetadataCall {
    return &pbgear.NominationPools_SetMetadataCall{
                PoolId : uint32(field.Value),
                Metadata : []uint32(field.Value),
    }
}   
func to_StakingRewards_Idsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_Id {
    return &pbgear.StakingRewards_Id{
                Value0: to_StakingRewards_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_StakingRewards_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.StakingRewards_SpCoreCryptoAccountId32 {
                    return &pbgear.StakingRewards_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignaturesalut(fields []*registry.DecodedField) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
    return &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{
                Value0: to_Grandpa_FinalityGrandpaPrevote(field.Value),
                Value1: to_Grandpa_SpConsensusGrandpaAppSignature(field.Value),
    }
}
            
                func to_Grandpa_FinalityGrandpaPrevote(in *registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaPrevote {
                    return &pbgear.Grandpa_FinalityGrandpaPrevote{
                        // TODO fill all fields
                    }
                }  
            
                func to_Grandpa_SpConsensusGrandpaAppSignature(in *registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
                    return &pbgear.Grandpa_SpConsensusGrandpaAppSignature{
                        // TODO fill all fields
                    }
                }   
func to_BagsList_Address32salut(fields []*registry.DecodedField) *pbgear.BagsList_Address32 {
    return &pbgear.BagsList_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Identity_TupleNull {
    return &pbgear.Identity_TupleNull{
    }
} 
func to_Identity_Raw18salut(fields []*registry.DecodedField) *pbgear.Identity_Raw18 {
    return &pbgear.Identity_Raw18{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_ForceProxyTypesalut(fields []*registry.DecodedField) *pbgear.Proxy_ForceProxyType {
    return &pbgear.Proxy_ForceProxyType{
                Value: to_oneof_Proxy_Value(field.Value),
    }
}  
func to_NominationPools_FreeBalancesalut(fields []*registry.DecodedField) *pbgear.NominationPools_FreeBalance {
    return &pbgear.NominationPools_FreeBalance{
                Value0 : string(field.Value),
    }
}
            
                   
func to_Babe_Consensussalut(fields []*registry.DecodedField) *pbgear.Babe_Consensus {
    return &pbgear.Babe_Consensus{
                Value0 : []uint32(field.Value),
                Value1 : []uint32(field.Value),
    }
}   
func to_Grandpa_FinalityGrandpaPrevotesalut(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaPrevote {
    return &pbgear.Grandpa_FinalityGrandpaPrevote{
                TargetHash: to_Grandpa_PrimitiveTypesH256(field.Value),
                TargetNumber : uint32(field.Value),
    }
}
            
                func to_Grandpa_PrimitiveTypesH256(in *registry.DecodedField) *pbgear.Grandpa_PrimitiveTypesH256 {
                    return &pbgear.Grandpa_PrimitiveTypesH256{
                        // TODO fill all fields
                    }
                }    
func to_FellowshipReferenda_SetMetadataCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SetMetadataCall {
    return &pbgear.FellowshipReferenda_SetMetadataCall{
                Index : uint32(field.Value),
                MaybeHash: to_optional_FellowshipReferenda_PrimitiveTypesH256(field.Value),
    }
}   
func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSizesalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
    return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{
                Voters: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Targets: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}   
func to_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesDatasalut(fields []*registry.DecodedField) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
    return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{
                Value0: to_Identity_Value0(field.Value),
                Value1: to_Identity_Value1(field.Value),
    }
}
            
                func to_Identity_Value0(in *registry.DecodedField) *pbgear.Identity_Value0 {
                    return &pbgear.Identity_Value0{
                        // TODO fill all fields
                    }
                }    
func to_Identity_FeePaidsalut(fields []*registry.DecodedField) *pbgear.Identity_FeePaid {
    return &pbgear.Identity_FeePaid{
                Value0 : string(field.Value),
    }
}  
func to_Proxy_Address32salut(fields []*registry.DecodedField) *pbgear.Proxy_Address32 {
    return &pbgear.Proxy_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_Blockedsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Blocked {
    return &pbgear.NominationPools_Blocked{
    }
} 
func to_Babe_SpCoreSr25519Publicsalut(fields []*registry.DecodedField) *pbgear.Babe_SpCoreSr25519Public {
    return &pbgear.Babe_SpCoreSr25519Public{
                Offender : []uint32(field.Value),
    }
}  
func to_Referenda_Legacysalut(fields []*registry.DecodedField) *pbgear.Referenda_Legacy {
    return &pbgear.Referenda_Legacy{
                Hash: to_Referenda_PrimitiveTypesH256(field.Value),
    }
}  
func to_Preimage_UnnotePreimageCallsalut(fields []*registry.DecodedField) *pbgear.Preimage_UnnotePreimageCall {
    return &pbgear.Preimage_UnnotePreimageCall{
                Hash: to_Preimage_PrimitiveTypesH256(field.Value),
    }
}  
func to_Identity_Imagesalut(fields []*registry.DecodedField) *pbgear.Identity_Image {
    return &pbgear.Identity_Image{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Staking_Idsalut(fields []*registry.DecodedField) *pbgear.Staking_Id {
    return &pbgear.Staking_Id{
                Value0: to_Staking_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Session_PalletImOnlineSr25519AppSr25519Publicsalut(fields []*registry.DecodedField) *pbgear.Session_PalletImOnlineSr25519AppSr25519Public {
    return &pbgear.Session_PalletImOnlineSr25519AppSr25519Public{
                ImOnline: to_Session_SpCoreSr25519Public(field.Value),
    }
}  
func to_ConvictionVoting_Tosalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_To {
    return &pbgear.ConvictionVoting_To{
                Value: to_oneof_ConvictionVoting_Value(field.Value),
    }
}  
func to_Bounties_Beneficiarysalut(fields []*registry.DecodedField) *pbgear.Bounties_Beneficiary {
    return &pbgear.Bounties_Beneficiary{
                Value: to_oneof_Bounties_Value(field.Value),
    }
}  
func to_System_KillStorageCallsalut(fields []*registry.DecodedField) *pbgear.System_KillStorageCall {
    return &pbgear.System_KillStorageCall{
                Keys: to_repeated_System_SystemKeysList(field.Value),
    }
}
            
                func to_repeated_System_SystemKeysList(in *registry.DecodedField) []*pbgear.System_SystemKeysList {
                    return []*pbgear.System_SystemKeysList{
                        // TODO fill all fields
                    }
                }   
func to_Staking_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Staking_CompactUint32 {
    return &pbgear.Staking_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_Identity_Raw4salut(fields []*registry.DecodedField) *pbgear.Identity_Raw4 {
    return &pbgear.Identity_Raw4{
                Value0 : []uint32(field.Value),
    }
}  
func to_ChildBounties_UnassignCuratorCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_UnassignCuratorCall {
    return &pbgear.ChildBounties_UnassignCuratorCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
    }
}   
func to_Staking_RebondCallsalut(fields []*registry.DecodedField) *pbgear.Staking_RebondCall {
    return &pbgear.Staking_RebondCall{
                Value: to_Staking_CompactString(field.Value),
    }
}  
func to_Gear_UploadProgramCallsalut(fields []*registry.DecodedField) *pbgear.Gear_UploadProgramCall {
    return &pbgear.Gear_UploadProgramCall{
                Code : []uint32(field.Value),
                Salt : []uint32(field.Value),
                InitPayload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}       
func to_Staking_CompactSpArithmeticPerThingsPerbillsalut(fields []*registry.DecodedField) *pbgear.Staking_CompactSpArithmeticPerThingsPerbill {
    return &pbgear.Staking_CompactSpArithmeticPerThingsPerbill{
                Value: to_Staking_SpArithmeticPerThingsPerbill(field.Value),
    }
}  
func to_Bounties_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Bounties_CompactUint32 {
    return &pbgear.Bounties_CompactUint32{
                Value : uint32(field.Value),
    }
}
            
                   
func to_FellowshipCollective_Address32salut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Address32 {
    return &pbgear.FellowshipCollective_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Scheduler_CancelCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_CancelCall {
    return &pbgear.Scheduler_CancelCall{
                When : uint32(field.Value),
                Index : uint32(field.Value),
    }
}   
func to_NominationPools_Membersalut(fields []*registry.DecodedField) *pbgear.NominationPools_Member {
    return &pbgear.NominationPools_Member{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_System_SetStorageCallsalut(fields []*registry.DecodedField) *pbgear.System_SetStorageCall {
    return &pbgear.System_SetStorageCall{
                Items: to_repeated_System_TupleSystemItemsListSystemItemsList(field.Value),
    }
}
            
                func to_repeated_System_TupleSystemItemsListSystemItemsList(in *registry.DecodedField) []*pbgear.System_TupleSystemItemsListSystemItemsList {
                    return []*pbgear.System_TupleSystemItemsListSystemItemsList{
                        // TODO fill all fields
                    }
                }   
func to_BagsList_Dislocatedsalut(fields []*registry.DecodedField) *pbgear.BagsList_Dislocated {
    return &pbgear.BagsList_Dislocated{
                Value: to_oneof_BagsList_Value(field.Value),
    }
}
            
                func to_oneof_BagsList_Value(in *registry.DecodedField) *pbgear.BagsList_Value {
                    return nil
                }   
func to_Identity_Emailsalut(fields []*registry.DecodedField) *pbgear.Identity_Email {
    return &pbgear.Identity_Email{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Proxy_KillPureCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_KillPureCall {
    return &pbgear.Proxy_KillPureCall{
                Spawner: to_Proxy_Spawner(field.Value),
                ProxyType: to_Proxy_ProxyType(field.Value),
                Index : uint32(field.Value),
                Height: to_Proxy_CompactUint32(field.Value),
                ExtIndex: to_Proxy_CompactUint32(field.Value),
    }
}
            
                func to_Proxy_Spawner(in *registry.DecodedField) *pbgear.Proxy_Spawner {
                    return &pbgear.Proxy_Spawner{
                        // TODO fill all fields
                    }
                }    
            
                func to_Proxy_CompactUint32(in *registry.DecodedField) *pbgear.Proxy_CompactUint32 {
                    return &pbgear.Proxy_CompactUint32{
                        // TODO fill all fields
                    }
                }    
func to_Multisig_SpWeightsWeightV2Weightsalut(fields []*registry.DecodedField) *pbgear.Multisig_SpWeightsWeightV2Weight {
    return &pbgear.Multisig_SpWeightsWeightV2Weight{
                RefTime: to_Multisig_CompactUint64(field.Value),
                ProofSize: to_Multisig_CompactUint64(field.Value),
    }
}
            
                func to_Multisig_CompactUint64(in *registry.DecodedField) *pbgear.Multisig_CompactUint64 {
                    return &pbgear.Multisig_CompactUint64{
                        // TODO fill all fields
                    }
                }    
func to_Bounties_CloseBountyCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_CloseBountyCall {
    return &pbgear.Bounties_CloseBountyCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
    }
}  
func to_NominationPools_Address32salut(fields []*registry.DecodedField) *pbgear.NominationPools_Address32 {
    return &pbgear.NominationPools_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_FellowshipReferenda_Legacysalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Legacy {
    return &pbgear.FellowshipReferenda_Legacy{
                Hash: to_FellowshipReferenda_PrimitiveTypesH256(field.Value),
    }
}  
func to_Identity_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Identity_CompactTupleNull {
    return &pbgear.Identity_CompactTupleNull{
                Value: to_Identity_TupleNull(field.Value),
    }
}
            
                func to_Identity_TupleNull(in *registry.DecodedField) *pbgear.Identity_TupleNull {
                    return &pbgear.Identity_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_ElectionProviderMultiPhase_SpNposElectionsSupportsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
    return &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{
                Total : string(field.Value),
                Voters: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(field.Value),
    }
} 
            
                func to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
                    return []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{
                        // TODO fill all fields
                    }
                }   
func to_StakingRewards_Address20salut(fields []*registry.DecodedField) *pbgear.StakingRewards_Address20 {
    return &pbgear.StakingRewards_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_FellowshipCollective_CleanupPollCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_CleanupPollCall {
    return &pbgear.FellowshipCollective_CleanupPollCall{
                PollIndex : uint32(field.Value),
                Max : uint32(field.Value),
    }
}   
func to_FellowshipReferenda_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_SpCoreCryptoAccountId32 {
    return &pbgear.FellowshipReferenda_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Proxy_AddProxyCallsalut(fields []*registry.DecodedField) *pbgear.Proxy_AddProxyCall {
    return &pbgear.Proxy_AddProxyCall{
                Delegate: to_Proxy_Delegate(field.Value),
                ProxyType: to_Proxy_ProxyType(field.Value),
                Delay : uint32(field.Value),
    }
}    
func to_GearVoucher_PalletGearVoucherInternalVoucherIdsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
    return &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{
                VoucherId : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactTupleNull {
    return &pbgear.ConvictionVoting_CompactTupleNull{
                Value: to_ConvictionVoting_TupleNull(field.Value),
    }
}
            
                func to_ConvictionVoting_TupleNull(in *registry.DecodedField) *pbgear.ConvictionVoting_TupleNull {
                    return &pbgear.ConvictionVoting_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_FellowshipReferenda_CancelCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_CancelCall {
    return &pbgear.FellowshipReferenda_CancelCall{
                Index : uint32(field.Value),
    }
}  
func to_NominationPools_Nominatorsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Nominator {
    return &pbgear.NominationPools_Nominator{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_SetConfigsCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetConfigsCall {
    return &pbgear.NominationPools_SetConfigsCall{
                MinJoinBond: to_NominationPools_MinJoinBond(field.Value),
                MinCreateBond: to_NominationPools_MinCreateBond(field.Value),
                MaxPools: to_NominationPools_MaxPools(field.Value),
                MaxMembers: to_NominationPools_MaxMembers(field.Value),
                MaxMembersPerPool: to_NominationPools_MaxMembersPerPool(field.Value),
                GlobalMaxCommission: to_NominationPools_GlobalMaxCommission(field.Value),
    }
}
            
                func to_NominationPools_MinJoinBond(in *registry.DecodedField) *pbgear.NominationPools_MinJoinBond {
                    return &pbgear.NominationPools_MinJoinBond{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_MinCreateBond(in *registry.DecodedField) *pbgear.NominationPools_MinCreateBond {
                    return &pbgear.NominationPools_MinCreateBond{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_MaxPools(in *registry.DecodedField) *pbgear.NominationPools_MaxPools {
                    return &pbgear.NominationPools_MaxPools{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_MaxMembers(in *registry.DecodedField) *pbgear.NominationPools_MaxMembers {
                    return &pbgear.NominationPools_MaxMembers{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_MaxMembersPerPool(in *registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
                    return &pbgear.NominationPools_MaxMembersPerPool{
                        // TODO fill all fields
                    }
                }  
            
                func to_NominationPools_GlobalMaxCommission(in *registry.DecodedField) *pbgear.NominationPools_GlobalMaxCommission {
                    return &pbgear.NominationPools_GlobalMaxCommission{
                        // TODO fill all fields
                    }
                }   
func to_Gear_SendReplyCallsalut(fields []*registry.DecodedField) *pbgear.Gear_SendReplyCall {
    return &pbgear.Gear_SendReplyCall{
                ReplyToId: to_Gear_GprimitivesMessageId(field.Value),
                Payload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}      
func to_StakingRewards_ForceRefillCallsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_ForceRefillCall {
    return &pbgear.StakingRewards_ForceRefillCall{
                From: to_StakingRewards_From(field.Value),
                Value : string(field.Value),
    }
}
            
                func to_StakingRewards_From(in *registry.DecodedField) *pbgear.StakingRewards_From {
                    return &pbgear.StakingRewards_From{
                        // TODO fill all fields
                    }
                }    
func to_Balances_TransferAllCallsalut(fields []*registry.DecodedField) *pbgear.Balances_TransferAllCall {
    return &pbgear.Balances_TransferAllCall{
                Dest: to_Balances_Dest(field.Value),
                KeepAlive : bool(field.Value),
    }
} 
            
                   
func to_Treasury_Address20salut(fields []*registry.DecodedField) *pbgear.Treasury_Address20 {
    return &pbgear.Treasury_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_PoolWithdrawUnbondedCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
    return &pbgear.NominationPools_PoolWithdrawUnbondedCall{
                PoolId : uint32(field.Value),
                NumSlashingSpans : uint32(field.Value),
    }
}   
func to_Gear_GprimitivesMessageIdsalut(fields []*registry.DecodedField) *pbgear.Gear_GprimitivesMessageId {
    return &pbgear.Gear_GprimitivesMessageId{
                ReplyToId : []uint32(field.Value),
    }
}  
func to_Referenda_Inlinesalut(fields []*registry.DecodedField) *pbgear.Referenda_Inline {
    return &pbgear.Referenda_Inline{
                Value0: to_Referenda_BoundedCollectionsBoundedVecBoundedVec(field.Value),
    }
}
            
                func to_Referenda_BoundedCollectionsBoundedVecBoundedVec(in *registry.DecodedField) *pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec {
                    return &pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec{
                        // TODO fill all fields
                    }
                }   
func to_Bounties_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
    return &pbgear.Bounties_CompactTupleNull{
                Value: to_Bounties_TupleNull(field.Value),
    }
}
            
                func to_Bounties_TupleNull(in *registry.DecodedField) *pbgear.Bounties_TupleNull {
                    return &pbgear.Bounties_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_BagsList_TupleNullsalut(fields []*registry.DecodedField) *pbgear.BagsList_TupleNull {
    return &pbgear.BagsList_TupleNull{
    }
} 
func to_BagsList_PutInFrontOfOtherCallsalut(fields []*registry.DecodedField) *pbgear.BagsList_PutInFrontOfOtherCall {
    return &pbgear.BagsList_PutInFrontOfOtherCall{
                Heavier: to_BagsList_Heavier(field.Value),
                Lighter: to_BagsList_Lighter(field.Value),
    }
}
            
                func to_BagsList_Heavier(in *registry.DecodedField) *pbgear.BagsList_Heavier {
                    return &pbgear.BagsList_Heavier{
                        // TODO fill all fields
                    }
                }    
func to_Staking_SetValidatorCountCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetValidatorCountCall {
    return &pbgear.Staking_SetValidatorCountCall{
                New: to_Staking_CompactUint32(field.Value),
    }
}  
func to_Staking_MinValidatorBondsalut(fields []*registry.DecodedField) *pbgear.Staking_MinValidatorBond {
    return &pbgear.Staking_MinValidatorBond{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_Staking_MaxValidatorCountsalut(fields []*registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
    return &pbgear.Staking_MaxValidatorCount{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_Identity_Value1salut(fields []*registry.DecodedField) *pbgear.Identity_Value1 {
    return &pbgear.Identity_Value1{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Babe_TupleUint64Uint64salut(fields []*registry.DecodedField) *pbgear.Babe_TupleUint64Uint64 {
    return &pbgear.Babe_TupleUint64Uint64{
                Value0 : uint64(field.Value),
                Value1 : uint64(field.Value),
    }
}
            
                    
func to_Balances_Rawsalut(fields []*registry.DecodedField) *pbgear.Balances_Raw {
    return &pbgear.Balances_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_SetFieldsCallsalut(fields []*registry.DecodedField) *pbgear.Identity_SetFieldsCall {
    return &pbgear.Identity_SetFieldsCall{
                Index: to_Identity_CompactUint32(field.Value),
                Fields: to_Identity_PalletIdentityTypesBitFlags(field.Value),
    }
} 
            
                func to_Identity_PalletIdentityTypesBitFlags(in *registry.DecodedField) *pbgear.Identity_PalletIdentityTypesBitFlags {
                    return &pbgear.Identity_PalletIdentityTypesBitFlags{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_AcceptCuratorCallsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_AcceptCuratorCall {
    return &pbgear.ChildBounties_AcceptCuratorCall{
                ParentBountyId: to_ChildBounties_CompactUint32(field.Value),
                ChildBountyId: to_ChildBounties_CompactUint32(field.Value),
    }
}   
func to_Staking_Accountsalut(fields []*registry.DecodedField) *pbgear.Staking_Account {
    return &pbgear.Staking_Account{
                Value0: to_Staking_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Utility_AsOriginsalut(fields []*registry.DecodedField) *pbgear.Utility_AsOrigin {
    return &pbgear.Utility_AsOrigin{
                Value: to_oneof_Utility_Value(field.Value),
    }
}  
func to_ConvictionVoting_PalletConvictionVotingVoteVotesalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
    return &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{
                Vote : uint32(field.Value),
    }
}  
func to_Bounties_ExtendBountyExpiryCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_ExtendBountyExpiryCall {
    return &pbgear.Bounties_ExtendBountyExpiryCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
                Remark : []uint32(field.Value),
    }
}   
func to_Babe_SpConsensusSlotsSlotsalut(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
    return &pbgear.Babe_SpConsensusSlotsSlot{
                Slot : uint64(field.Value),
    }
}  
func to_Balances_Whosalut(fields []*registry.DecodedField) *pbgear.Balances_Who {
    return &pbgear.Balances_Who{
                Value: to_oneof_Balances_Value(field.Value),
    }
}  
func to_Identity_Reasonablesalut(fields []*registry.DecodedField) *pbgear.Identity_Reasonable {
    return &pbgear.Identity_Reasonable{
    }
} 
func to_Gear_CreateProgramCallsalut(fields []*registry.DecodedField) *pbgear.Gear_CreateProgramCall {
    return &pbgear.Gear_CreateProgramCall{
                CodeId: to_Gear_GprimitivesCodeId(field.Value),
                Salt : []uint32(field.Value),
                InitPayload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}
            
                func to_Gear_GprimitivesCodeId(in *registry.DecodedField) *pbgear.Gear_GprimitivesCodeId {
                    return &pbgear.Gear_GprimitivesCodeId{
                        // TODO fill all fields
                    }
                }        
func to_BagsList_Address20salut(fields []*registry.DecodedField) *pbgear.BagsList_Address20 {
    return &pbgear.BagsList_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_ImOnline_SpCoreSr25519Signaturesalut(fields []*registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
    return &pbgear.ImOnline_SpCoreSr25519Signature{
                Signature : []uint32(field.Value),
    }
}  
func to_Bounties_Indexsalut(fields []*registry.DecodedField) *pbgear.Bounties_Index {
    return &pbgear.Bounties_Index{
                Value0: to_Bounties_CompactTupleNull(field.Value),
    }
}
            
                func to_Bounties_CompactTupleNull(in *registry.DecodedField) *pbgear.Bounties_CompactTupleNull {
                    return &pbgear.Bounties_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_AdjustPoolDepositCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_AdjustPoolDepositCall {
    return &pbgear.NominationPools_AdjustPoolDepositCall{
                PoolId : uint32(field.Value),
    }
}  
func to_StakingRewards_Address32salut(fields []*registry.DecodedField) *pbgear.StakingRewards_Address32 {
    return &pbgear.StakingRewards_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_FellowshipReferenda_EnactmentMomentsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_EnactmentMoment {
    return &pbgear.FellowshipReferenda_EnactmentMoment{
                Value: to_oneof_FellowshipReferenda_Value(field.Value),
    }
}  
func to_Proxy_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Proxy_CompactUint32 {
    return &pbgear.Proxy_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_Balances_CompactStringsalut(fields []*registry.DecodedField) *pbgear.Balances_CompactString {
    return &pbgear.Balances_CompactString{
                Value : string(field.Value),
    }
}  
func to_Staking_Setsalut(fields []*registry.DecodedField) *pbgear.Staking_Set {
    return &pbgear.Staking_Set{
                Value0: to_Staking_SpArithmeticPerThingsPerbill(field.Value),
    }
}  
func to_Babe_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Babe_CompactUint32 {
    return &pbgear.Babe_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_Balances_Indexsalut(fields []*registry.DecodedField) *pbgear.Balances_Index {
    return &pbgear.Balances_Index{
                Value0: to_Balances_CompactTupleNull(field.Value),
    }
}
            
                func to_Balances_CompactTupleNull(in *registry.DecodedField) *pbgear.Balances_CompactTupleNull {
                    return &pbgear.Balances_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_ChildBounties_CompactStringsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_CompactString {
    return &pbgear.ChildBounties_CompactString{
                Value : string(field.Value),
    }
}
            
                   
func to_GearVoucher_SendMessagesalut(fields []*registry.DecodedField) *pbgear.GearVoucher_SendMessage {
    return &pbgear.GearVoucher_SendMessage{
                Destination: to_GearVoucher_GprimitivesActorId(field.Value),
                Payload : []uint32(field.Value),
                GasLimit : uint64(field.Value),
                Value : string(field.Value),
                KeepAlive : bool(field.Value),
    }
}      
func to_Staking_Removesalut(fields []*registry.DecodedField) *pbgear.Staking_Remove {
    return &pbgear.Staking_Remove{
    }
} 
func to_Identity_Keccak256salut(fields []*registry.DecodedField) *pbgear.Identity_Keccak256 {
    return &pbgear.Identity_Keccak256{
                Value0 : []uint32(field.Value),
    }
}  
func to_Treasury_RejectProposalCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_RejectProposalCall {
    return &pbgear.Treasury_RejectProposalCall{
                ProposalId: to_Treasury_CompactUint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes8ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_StakingRewards_Rawsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_Raw {
    return &pbgear.StakingRewards_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Grandpa_FinalityGrandpaEquivocationsalut(fields []*registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaEquivocation {
    return &pbgear.Grandpa_FinalityGrandpaEquivocation{
                RoundNumber : uint64(field.Value),
                Identity: to_Grandpa_SpConsensusGrandpaAppPublic(field.Value),
                First: to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(field.Value),
                Second: to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(field.Value),
    }
} 
            
                func to_Grandpa_SpConsensusGrandpaAppPublic(in *registry.DecodedField) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
                    return &pbgear.Grandpa_SpConsensusGrandpaAppPublic{
                        // TODO fill all fields
                    }
                }  
            
                func to_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(in *registry.DecodedField) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
                    return &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{
                        // TODO fill all fields
                    }
                }    
func to_Vesting_Targetsalut(fields []*registry.DecodedField) *pbgear.Vesting_Target {
    return &pbgear.Vesting_Target{
                Value: to_oneof_Vesting_Value(field.Value),
    }
}  
func to_ConvictionVoting_CompactUint32salut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_CompactUint32 {
    return &pbgear.ConvictionVoting_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_FellowshipCollective_Indexsalut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_Index {
    return &pbgear.FellowshipCollective_Index{
                Value0: to_FellowshipCollective_CompactTupleNull(field.Value),
    }
}
            
                func to_FellowshipCollective_CompactTupleNull(in *registry.DecodedField) *pbgear.FellowshipCollective_CompactTupleNull {
                    return &pbgear.FellowshipCollective_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Preimage_NotePreimageCallsalut(fields []*registry.DecodedField) *pbgear.Preimage_NotePreimageCall {
    return &pbgear.Preimage_NotePreimageCall{
                Bytes : []uint32(field.Value),
    }
}  
func to_Multisig_PalletMultisigTimepointsalut(fields []*registry.DecodedField) *pbgear.Multisig_PalletMultisigTimepoint {
    return &pbgear.Multisig_PalletMultisigTimepoint{
                Height : uint32(field.Value),
                Index : uint32(field.Value),
    }
}   
func to_ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16 {
    return &pbgear.ElectionProviderMultiPhase_SpArithmeticPerThingsPerU16{
                Value : uint32(field.Value),
    }
}
            
                   
func to_NominationPools_Rewardssalut(fields []*registry.DecodedField) *pbgear.NominationPools_Rewards {
    return &pbgear.NominationPools_Rewards{
    }
} 
func to_Staking_MaxNominatorCountsalut(fields []*registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
    return &pbgear.Staking_MaxNominatorCount{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_Treasury_CompactStringsalut(fields []*registry.DecodedField) *pbgear.Treasury_CompactString {
    return &pbgear.Treasury_CompactString{
                Value : string(field.Value),
    }
}
            
                   
func to_Staking_ChillOtherCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ChillOtherCall {
    return &pbgear.Staking_ChillOtherCall{
                Controller: to_Staking_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Utility_Systemsalut(fields []*registry.DecodedField) *pbgear.Utility_System {
    return &pbgear.Utility_System{
                Value0: to_Utility_Value0(field.Value),
    }
}  
func to_ConvictionVoting_Locked1Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked1X {
    return &pbgear.ConvictionVoting_Locked1X{
    }
} 
func to_Identity_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Identity_SpCoreCryptoAccountId32 {
    return &pbgear.Identity_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_GovernanceFallbackCallsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
    return &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{
                MaybeMaxVoters : uint32(field.Value),
                MaybeMaxTargets : uint32(field.Value),
    }
}   
func to_ChildBounties_Address32salut(fields []*registry.DecodedField) *pbgear.ChildBounties_Address32 {
    return &pbgear.ChildBounties_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Vesting_VestCallsalut(fields []*registry.DecodedField) *pbgear.Vesting_VestCall {
    return &pbgear.Vesting_VestCall{
    }
} 
func to_Staking_KickCallsalut(fields []*registry.DecodedField) *pbgear.Staking_KickCall {
    return &pbgear.Staking_KickCall{
                Who: to_repeated_Staking_SpRuntimeMultiaddressMultiAddress(field.Value),
    }
}  
func to_Treasury_CompactUint32salut(fields []*registry.DecodedField) *pbgear.Treasury_CompactUint32 {
    return &pbgear.Treasury_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_Treasury_SpendCallsalut(fields []*registry.DecodedField) *pbgear.Treasury_SpendCall {
    return &pbgear.Treasury_SpendCall{
                AssetKind: to_Treasury_TupleNull(field.Value),
                Amount: to_Treasury_CompactString(field.Value),
                Beneficiary: to_Treasury_SpCoreCryptoAccountId32(field.Value),
                ValidFrom : uint32(field.Value),
    }
}
            
                func to_Treasury_TupleNull(in *registry.DecodedField) *pbgear.Treasury_TupleNull {
                    return &pbgear.Treasury_TupleNull{
                        // TODO fill all fields
                    }
                }      
func to_Proxy_Spawnersalut(fields []*registry.DecodedField) *pbgear.Proxy_Spawner {
    return &pbgear.Proxy_Spawner{
                Value: to_oneof_Proxy_Value(field.Value),
    }
}  
func to_Staking_SetInvulnerablesCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetInvulnerablesCall {
    return &pbgear.Staking_SetInvulnerablesCall{
                Invulnerables: to_repeated_Staking_SpCoreCryptoAccountId32(field.Value),
    }
}  
func to_Treasury_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.Treasury_CompactTupleNull {
    return &pbgear.Treasury_CompactTupleNull{
                Value: to_Treasury_TupleNull(field.Value),
    }
}  
func to_NominationPools_BondExtraCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_BondExtraCall {
    return &pbgear.NominationPools_BondExtraCall{
                Extra: to_NominationPools_Extra(field.Value),
    }
}  
func to_NominationPools_MemberAccountsalut(fields []*registry.DecodedField) *pbgear.NominationPools_MemberAccount {
    return &pbgear.NominationPools_MemberAccount{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_SetCommissionMaxCallsalut(fields []*registry.DecodedField) *pbgear.NominationPools_SetCommissionMaxCall {
    return &pbgear.NominationPools_SetCommissionMaxCall{
                PoolId : uint32(field.Value),
                MaxCommission: to_NominationPools_SpArithmeticPerThingsPerbill(field.Value),
    }
}   
func to_Identity_CancelRequestCallsalut(fields []*registry.DecodedField) *pbgear.Identity_CancelRequestCall {
    return &pbgear.Identity_CancelRequestCall{
                RegIndex : uint32(field.Value),
    }
}  
func to_Bounties_ApproveBountyCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_ApproveBountyCall {
    return &pbgear.Bounties_ApproveBountyCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
    }
}  
func to_FellowshipReferenda_RefundDecisionDepositCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
    return &pbgear.FellowshipReferenda_RefundDecisionDepositCall{
                Index : uint32(field.Value),
    }
}  
func to_FellowshipReferenda_OneFewerDecidingCallsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
    return &pbgear.FellowshipReferenda_OneFewerDecidingCall{
                Track : uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolutionsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
    return &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{
                Solution: to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(field.Value),
                Score: to_ElectionProviderMultiPhase_SpNposElectionsElectionScore(field.Value),
                Round : uint32(field.Value),
    }
}
            
                func to_ElectionProviderMultiPhase_VaraRuntimeNposSolution16(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16 {
                    return &pbgear.ElectionProviderMultiPhase_VaraRuntimeNposSolution16{
                        // TODO fill all fields
                    }
                }     
func to_System_SetCodeWithoutChecksCallsalut(fields []*registry.DecodedField) *pbgear.System_SetCodeWithoutChecksCall {
    return &pbgear.System_SetCodeWithoutChecksCall{
                Code : []uint32(field.Value),
    }
}  
func to_FellowshipCollective_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.FellowshipCollective_SpCoreCryptoAccountId32 {
    return &pbgear.FellowshipCollective_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_SplitAbstainsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SplitAbstain {
    return &pbgear.ConvictionVoting_SplitAbstain{
                Aye : string(field.Value),
                Nay : string(field.Value),
                Abstain : string(field.Value),
    }
}    
func to_Identity_Accountsalut(fields []*registry.DecodedField) *pbgear.Identity_Account {
    return &pbgear.Identity_Account{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Identity_Raw14salut(fields []*registry.DecodedField) *pbgear.Identity_Raw14 {
    return &pbgear.Identity_Raw14{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Raw24salut(fields []*registry.DecodedField) *pbgear.Identity_Raw24 {
    return &pbgear.Identity_Raw24{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Value0salut(fields []*registry.DecodedField) *pbgear.Identity_Value0 {
    return &pbgear.Identity_Value0{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Identity_QuitSubCallsalut(fields []*registry.DecodedField) *pbgear.Identity_QuitSubCall {
    return &pbgear.Identity_QuitSubCall{
    }
} 
func to_Babe_SpConsensusSlotsEquivocationProofsalut(fields []*registry.DecodedField) *pbgear.Babe_SpConsensusSlotsEquivocationProof {
    return &pbgear.Babe_SpConsensusSlotsEquivocationProof{
                Offender: to_Babe_SpConsensusBabeAppPublic(field.Value),
                Slot: to_Babe_SpConsensusSlotsSlot(field.Value),
                FirstHeader: to_Babe_SpRuntimeGenericHeaderHeader(field.Value),
                SecondHeader: to_Babe_SpRuntimeGenericHeaderHeader(field.Value),
    }
}
            
                func to_Babe_SpConsensusBabeAppPublic(in *registry.DecodedField) *pbgear.Babe_SpConsensusBabeAppPublic {
                    return &pbgear.Babe_SpConsensusBabeAppPublic{
                        // TODO fill all fields
                    }
                }  
            
                func to_Babe_SpConsensusSlotsSlot(in *registry.DecodedField) *pbgear.Babe_SpConsensusSlotsSlot {
                    return &pbgear.Babe_SpConsensusSlotsSlot{
                        // TODO fill all fields
                    }
                }  
            
                func to_Babe_SpRuntimeGenericHeaderHeader(in *registry.DecodedField) *pbgear.Babe_SpRuntimeGenericHeaderHeader {
                    return &pbgear.Babe_SpRuntimeGenericHeaderHeader{
                        // TODO fill all fields
                    }
                }    
func to_Utility_DispatchAsCallsalut(fields []*registry.DecodedField) *pbgear.Utility_DispatchAsCall {
    return &pbgear.Utility_DispatchAsCall{
                AsOrigin: to_Utility_AsOrigin(field.Value),
                Call: to_oneof_Utility_Call(field.Value),
    }
}
            
                func to_Utility_AsOrigin(in *registry.DecodedField) *pbgear.Utility_AsOrigin {
                    return &pbgear.Utility_AsOrigin{
                        // TODO fill all fields
                    }
                }    
func to_Proxy_ProxyTypesalut(fields []*registry.DecodedField) *pbgear.Proxy_ProxyType {
    return &pbgear.Proxy_ProxyType{
                Value: to_oneof_Proxy_Value(field.Value),
    }
}  
func to_GearVoucher_CallDeprecatedCallsalut(fields []*registry.DecodedField) *pbgear.GearVoucher_CallDeprecatedCall {
    return &pbgear.GearVoucher_CallDeprecatedCall{
                Call: to_GearVoucher_Call(field.Value),
    }
}  
func to_Scheduler_ScheduleCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_ScheduleCall {
    return &pbgear.Scheduler_ScheduleCall{
                When : uint32(field.Value),
                MaybePeriodic: to_optional_Scheduler_TupleUint32Uint32(field.Value),
                Priority : uint32(field.Value),
                Call: to_oneof_Scheduler_Call(field.Value),
    }
}     
func to_System_RemarkWithEventCallsalut(fields []*registry.DecodedField) *pbgear.System_RemarkWithEventCall {
    return &pbgear.System_RemarkWithEventCall{
                Remark : []uint32(field.Value),
    }
}  
func to_Utility_SpWeightsWeightV2Weightsalut(fields []*registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
    return &pbgear.Utility_SpWeightsWeightV2Weight{
                RefTime: to_Utility_CompactUint64(field.Value),
                ProofSize: to_Utility_CompactUint64(field.Value),
    }
}
            
                func to_Utility_CompactUint64(in *registry.DecodedField) *pbgear.Utility_CompactUint64 {
                    return &pbgear.Utility_CompactUint64{
                        // TODO fill all fields
                    }
                }    
func to_Referenda_Rootsalut(fields []*registry.DecodedField) *pbgear.Referenda_Root {
    return &pbgear.Referenda_Root{
    }
} 
func to_Babe_Logssalut(fields []*registry.DecodedField) *pbgear.Babe_Logs {
    return &pbgear.Babe_Logs{
                Value: to_oneof_Babe_Value(field.Value),
    }
}  
func to_Staking_ForceNoErasCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ForceNoErasCall {
    return &pbgear.Staking_ForceNoErasCall{
    }
} 
func to_Vesting_Idsalut(fields []*registry.DecodedField) *pbgear.Vesting_Id {
    return &pbgear.Vesting_Id{
                Value0: to_Vesting_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Vesting_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Vesting_SpCoreCryptoAccountId32 {
                    return &pbgear.Vesting_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw10salut(fields []*registry.DecodedField) *pbgear.Identity_Raw10 {
    return &pbgear.Identity_Raw10{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Targetsalut(fields []*registry.DecodedField) *pbgear.Identity_Target {
    return &pbgear.Identity_Target{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Babe_PrimitiveTypesH256salut(fields []*registry.DecodedField) *pbgear.Babe_PrimitiveTypesH256 {
    return &pbgear.Babe_PrimitiveTypesH256{
                ParentHash : []uint32(field.Value),
    }
}  
func to_Babe_PrimarySlotssalut(fields []*registry.DecodedField) *pbgear.Babe_PrimarySlots {
    return &pbgear.Babe_PrimarySlots{
    }
} 
func to_ImOnline_PalletImOnlineSr25519AppSr25519Signaturesalut(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
    return &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{
                Signature: to_ImOnline_SpCoreSr25519Signature(field.Value),
    }
}
            
                func to_ImOnline_SpCoreSr25519Signature(in *registry.DecodedField) *pbgear.ImOnline_SpCoreSr25519Signature {
                    return &pbgear.ImOnline_SpCoreSr25519Signature{
                        // TODO fill all fields
                    }
                }   
func to_Treasury_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Treasury_TupleNull {
    return &pbgear.Treasury_TupleNull{
    }
} 
func to_Utility_WithWeightCallsalut(fields []*registry.DecodedField) *pbgear.Utility_WithWeightCall {
    return &pbgear.Utility_WithWeightCall{
                Call: to_oneof_Utility_Call(field.Value),
                Weight: to_Utility_SpWeightsWeightV2Weight(field.Value),
    }
} 
            
                func to_Utility_SpWeightsWeightV2Weight(in *registry.DecodedField) *pbgear.Utility_SpWeightsWeightV2Weight {
                    return &pbgear.Utility_SpWeightsWeightV2Weight{
                        // TODO fill all fields
                    }
                }   
func to_Referenda_Nonesalut(fields []*registry.DecodedField) *pbgear.Referenda_None {
    return &pbgear.Referenda_None{
    }
} 
func to_FellowshipReferenda_Voidsalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Void {
    return &pbgear.FellowshipReferenda_Void{
                Value0: to_FellowshipReferenda_Value0(field.Value),
    }
}  
func to_Grandpa_Precommitsalut(fields []*registry.DecodedField) *pbgear.Grandpa_Precommit {
    return &pbgear.Grandpa_Precommit{
                Value0: to_Grandpa_FinalityGrandpaEquivocation(field.Value),
    }
}
            
                func to_Grandpa_FinalityGrandpaEquivocation(in *registry.DecodedField) *pbgear.Grandpa_FinalityGrandpaEquivocation {
                    return &pbgear.Grandpa_FinalityGrandpaEquivocation{
                        // TODO fill all fields
                    }
                }   
func to_Balances_TransferAllowDeathCallsalut(fields []*registry.DecodedField) *pbgear.Balances_TransferAllowDeathCall {
    return &pbgear.Balances_TransferAllowDeathCall{
                Dest: to_Balances_Dest(field.Value),
                Value: to_Balances_CompactString(field.Value),
    }
}   
func to_ConvictionVoting_Locked4Xsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Locked4X {
    return &pbgear.ConvictionVoting_Locked4X{
    }
} 
func to_Referenda_BoundedCollectionsBoundedVecBoundedVecsalut(fields []*registry.DecodedField) *pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec {
    return &pbgear.Referenda_BoundedCollectionsBoundedVecBoundedVec{
                Value0 : []uint32(field.Value),
    }
}  
func to_Scheduler_CancelNamedCallsalut(fields []*registry.DecodedField) *pbgear.Scheduler_CancelNamedCall {
    return &pbgear.Scheduler_CancelNamedCall{
                Id : []uint32(field.Value),
    }
}  
func to_Preimage_RequestPreimageCallsalut(fields []*registry.DecodedField) *pbgear.Preimage_RequestPreimageCall {
    return &pbgear.Preimage_RequestPreimageCall{
                Hash: to_Preimage_PrimitiveTypesH256(field.Value),
    }
}  
func to_Preimage_EnsureUpdatedCallsalut(fields []*registry.DecodedField) *pbgear.Preimage_EnsureUpdatedCall {
    return &pbgear.Preimage_EnsureUpdatedCall{
                Hashes: to_repeated_Preimage_PrimitiveTypesH256(field.Value),
    }
}  
func to_Identity_Legalsalut(fields []*registry.DecodedField) *pbgear.Identity_Legal {
    return &pbgear.Identity_Legal{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Balances_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Balances_SpCoreCryptoAccountId32 {
    return &pbgear.Balances_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Utility_ForceBatchCallsalut(fields []*registry.DecodedField) *pbgear.Utility_ForceBatchCall {
    return &pbgear.Utility_ForceBatchCall{
                Calls: to_repeated_Utility_VaraRuntimeRuntimeCall(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes11ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_ChildBounties_Curatorsalut(fields []*registry.DecodedField) *pbgear.ChildBounties_Curator {
    return &pbgear.ChildBounties_Curator{
                Value: to_oneof_ChildBounties_Value(field.Value),
    }
}  
func to_Proxy_Rawsalut(fields []*registry.DecodedField) *pbgear.Proxy_Raw {
    return &pbgear.Proxy_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32CompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32CompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}   
func to_Session_SetKeysCallsalut(fields []*registry.DecodedField) *pbgear.Session_SetKeysCall {
    return &pbgear.Session_SetKeysCall{
                Keys: to_Session_VaraRuntimeSessionKeys(field.Value),
                Proof : []uint32(field.Value),
    }
}
            
                func to_Session_VaraRuntimeSessionKeys(in *registry.DecodedField) *pbgear.Session_VaraRuntimeSessionKeys {
                    return &pbgear.Session_VaraRuntimeSessionKeys{
                        // TODO fill all fields
                    }
                }    
func to_Identity_RequestJudgementCallsalut(fields []*registry.DecodedField) *pbgear.Identity_RequestJudgementCall {
    return &pbgear.Identity_RequestJudgementCall{
                RegIndex: to_Identity_CompactUint32(field.Value),
                MaxFee: to_Identity_CompactString(field.Value),
    }
}   
func to_ImOnline_PalletImOnlineHeartbeatsalut(fields []*registry.DecodedField) *pbgear.ImOnline_PalletImOnlineHeartbeat {
    return &pbgear.ImOnline_PalletImOnlineHeartbeat{
                BlockNumber : uint32(field.Value),
                SessionIndex : uint32(field.Value),
                AuthorityIndex : uint32(field.Value),
                ValidatorsLen : uint32(field.Value),
    }
}
            
                      
func to_Bounties_ProposeCuratorCallsalut(fields []*registry.DecodedField) *pbgear.Bounties_ProposeCuratorCall {
    return &pbgear.Bounties_ProposeCuratorCall{
                BountyId: to_Bounties_CompactUint32(field.Value),
                Curator: to_Bounties_Curator(field.Value),
                Fee: to_Bounties_CompactString(field.Value),
    }
} 
            
                func to_Bounties_Curator(in *registry.DecodedField) *pbgear.Bounties_Curator {
                    return &pbgear.Bounties_Curator{
                        // TODO fill all fields
                    }
                }    
func to_StakingRewards_CompactTupleNullsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
    return &pbgear.StakingRewards_CompactTupleNull{
                Value: to_StakingRewards_TupleNull(field.Value),
    }
}
            
                func to_StakingRewards_TupleNull(in *registry.DecodedField) *pbgear.StakingRewards_TupleNull {
                    return &pbgear.StakingRewards_TupleNull{
                        // TODO fill all fields
                    }
                }   
func to_Babe_PrimaryAndSecondaryVRFSlotssalut(fields []*registry.DecodedField) *pbgear.Babe_PrimaryAndSecondaryVRFSlots {
    return &pbgear.Babe_PrimaryAndSecondaryVRFSlots{
    }
} 
func to_Grandpa_Prevotesalut(fields []*registry.DecodedField) *pbgear.Grandpa_Prevote {
    return &pbgear.Grandpa_Prevote{
                Value0: to_Grandpa_FinalityGrandpaEquivocation(field.Value),
    }
}  
func to_Proxy_Idsalut(fields []*registry.DecodedField) *pbgear.Proxy_Id {
    return &pbgear.Proxy_Id{
                Value0: to_Proxy_SpCoreCryptoAccountId32(field.Value),
    }
}
            
                func to_Proxy_SpCoreCryptoAccountId32(in *registry.DecodedField) *pbgear.Proxy_SpCoreCryptoAccountId32 {
                    return &pbgear.Proxy_SpCoreCryptoAccountId32{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_MaxMembersPerPoolsalut(fields []*registry.DecodedField) *pbgear.NominationPools_MaxMembersPerPool {
    return &pbgear.NominationPools_MaxMembersPerPool{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_Referenda_SubmitCallsalut(fields []*registry.DecodedField) *pbgear.Referenda_SubmitCall {
    return &pbgear.Referenda_SubmitCall{
                ProposalOrigin: to_Referenda_ProposalOrigin(field.Value),
                Proposal: to_Referenda_Proposal(field.Value),
                EnactmentMoment: to_Referenda_EnactmentMoment(field.Value),
    }
}
            
                func to_Referenda_ProposalOrigin(in *registry.DecodedField) *pbgear.Referenda_ProposalOrigin {
                    return &pbgear.Referenda_ProposalOrigin{
                        // TODO fill all fields
                    }
                }  
            
                func to_Referenda_Proposal(in *registry.DecodedField) *pbgear.Referenda_Proposal {
                    return &pbgear.Referenda_Proposal{
                        // TODO fill all fields
                    }
                }  
            
                func to_Referenda_EnactmentMoment(in *registry.DecodedField) *pbgear.Referenda_EnactmentMoment {
                    return &pbgear.Referenda_EnactmentMoment{
                        // TODO fill all fields
                    }
                }   
func to_Scheduler_TupleUint32Uint32salut(fields []*registry.DecodedField) *pbgear.Scheduler_TupleUint32Uint32 {
    return &pbgear.Scheduler_TupleUint32Uint32{
                Value0 : uint32(field.Value),
                Value1 : uint32(field.Value),
    }
}   
func to_Staking_Rawsalut(fields []*registry.DecodedField) *pbgear.Staking_Raw {
    return &pbgear.Staking_Raw{
                Value0 : []uint32(field.Value),
    }
}  
func to_Staking_MinCommissionsalut(fields []*registry.DecodedField) *pbgear.Staking_MinCommission {
    return &pbgear.Staking_MinCommission{
                Value: to_oneof_Staking_Value(field.Value),
    }
}  
func to_Referenda_Aftersalut(fields []*registry.DecodedField) *pbgear.Referenda_After {
    return &pbgear.Referenda_After{
                Value0 : uint32(field.Value),
    }
}  
func to_Identity_PalletIdentitySimpleIdentityInfosalut(fields []*registry.DecodedField) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
    return &pbgear.Identity_PalletIdentitySimpleIdentityInfo{
                Additional: to_Identity_BoundedCollectionsBoundedVecBoundedVec(field.Value),
                Display: to_Identity_Display(field.Value),
                Legal: to_Identity_Legal(field.Value),
                Web: to_Identity_Web(field.Value),
                Riot: to_Identity_Riot(field.Value),
                Email: to_Identity_Email(field.Value),
                PgpFingerprint : []uint32(field.Value),
                Image: to_Identity_Image(field.Value),
                Twitter: to_Identity_Twitter(field.Value),
    }
}
            
                func to_Identity_BoundedCollectionsBoundedVecBoundedVec(in *registry.DecodedField) *pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec {
                    return &pbgear.Identity_BoundedCollectionsBoundedVecBoundedVec{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Display(in *registry.DecodedField) *pbgear.Identity_Display {
                    return &pbgear.Identity_Display{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Legal(in *registry.DecodedField) *pbgear.Identity_Legal {
                    return &pbgear.Identity_Legal{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Web(in *registry.DecodedField) *pbgear.Identity_Web {
                    return &pbgear.Identity_Web{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Riot(in *registry.DecodedField) *pbgear.Identity_Riot {
                    return &pbgear.Identity_Riot{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Email(in *registry.DecodedField) *pbgear.Identity_Email {
                    return &pbgear.Identity_Email{
                        // TODO fill all fields
                    }
                }   
            
                func to_Identity_Image(in *registry.DecodedField) *pbgear.Identity_Image {
                    return &pbgear.Identity_Image{
                        // TODO fill all fields
                    }
                }  
            
                func to_Identity_Twitter(in *registry.DecodedField) *pbgear.Identity_Twitter {
                    return &pbgear.Identity_Twitter{
                        // TODO fill all fields
                    }
                }   
func to_ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32TupleCompactUint32CompactSpArithmeticPerThingsPerU16CompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes14ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_System_SetCodeCallsalut(fields []*registry.DecodedField) *pbgear.System_SetCodeCall {
    return &pbgear.System_SetCodeCall{
                Code : []uint32(field.Value),
    }
}  
func to_Balances_Address32salut(fields []*registry.DecodedField) *pbgear.Balances_Address32 {
    return &pbgear.Balances_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_StakingRewards_Indexsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_Index {
    return &pbgear.StakingRewards_Index{
                Value0: to_StakingRewards_CompactTupleNull(field.Value),
    }
}
            
                func to_StakingRewards_CompactTupleNull(in *registry.DecodedField) *pbgear.StakingRewards_CompactTupleNull {
                    return &pbgear.StakingRewards_CompactTupleNull{
                        // TODO fill all fields
                    }
                }   
func to_StakingRewards_AlignSupplyCallsalut(fields []*registry.DecodedField) *pbgear.StakingRewards_AlignSupplyCall {
    return &pbgear.StakingRewards_AlignSupplyCall{
                Target : string(field.Value),
    }
}  
func to_NominationPools_Rootsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Root {
    return &pbgear.NominationPools_Root{
                Value: to_oneof_NominationPools_Value(field.Value),
    }
}  
func to_NominationPools_Destroyingsalut(fields []*registry.DecodedField) *pbgear.NominationPools_Destroying {
    return &pbgear.NominationPools_Destroying{
    }
} 
func to_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupportsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
    return &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{
                Value0: to_ElectionProviderMultiPhase_SpCoreCryptoAccountId32(field.Value),
                Value1: to_ElectionProviderMultiPhase_SpNposElectionsSupport(field.Value),
    }
} 
            
                func to_ElectionProviderMultiPhase_SpNposElectionsSupport(in *registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport {
                    return &pbgear.ElectionProviderMultiPhase_SpNposElectionsSupport{
                        // TODO fill all fields
                    }
                }   
func to_Babe_AllowedSlotssalut(fields []*registry.DecodedField) *pbgear.Babe_AllowedSlots {
    return &pbgear.Babe_AllowedSlots{
                Value: to_oneof_Babe_Value(field.Value),
    }
}  
func to_Proxy_Delegatesalut(fields []*registry.DecodedField) *pbgear.Proxy_Delegate {
    return &pbgear.Proxy_Delegate{
                Value: to_oneof_Proxy_Value(field.Value),
    }
}  
func to_Proxy_Address20salut(fields []*registry.DecodedField) *pbgear.Proxy_Address20 {
    return &pbgear.Proxy_Address20{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_CompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_CompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_CompactUint32{
                Value : uint32(field.Value),
    }
}  
func to_NominationPools_CompactStringsalut(fields []*registry.DecodedField) *pbgear.NominationPools_CompactString {
    return &pbgear.NominationPools_CompactString{
                Value : string(field.Value),
    }
}  
func to_BagsList_Lightersalut(fields []*registry.DecodedField) *pbgear.BagsList_Lighter {
    return &pbgear.BagsList_Lighter{
                Value: to_oneof_BagsList_Value(field.Value),
    }
}  
func to_ConvictionVoting_Standardsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_Standard {
    return &pbgear.ConvictionVoting_Standard{
                Vote: to_ConvictionVoting_PalletConvictionVotingVoteVote(field.Value),
                Balance : string(field.Value),
    }
}
            
                func to_ConvictionVoting_PalletConvictionVotingVoteVote(in *registry.DecodedField) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
                    return &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{
                        // TODO fill all fields
                    }
                }    
func to_Staking_Address32salut(fields []*registry.DecodedField) *pbgear.Staking_Address32 {
    return &pbgear.Staking_Address32{
                Value0 : []uint32(field.Value),
    }
}  
func to_FellowshipReferenda_Inlinesalut(fields []*registry.DecodedField) *pbgear.FellowshipReferenda_Inline {
    return &pbgear.FellowshipReferenda_Inline{
                Value0: to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(field.Value),
    }
}
            
                func to_FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec(in *registry.DecodedField) *pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec {
                    return &pbgear.FellowshipReferenda_BoundedCollectionsBoundedVecBoundedVec{
                        // TODO fill all fields
                    }
                }   
func to_Identity_Raw26salut(fields []*registry.DecodedField) *pbgear.Identity_Raw26 {
    return &pbgear.Identity_Raw26{
                Value0 : []uint32(field.Value),
    }
}  
func to_Identity_Riotsalut(fields []*registry.DecodedField) *pbgear.Identity_Riot {
    return &pbgear.Identity_Riot{
                Value: to_oneof_Identity_Value(field.Value),
    }
}  
func to_Identity_Unknownsalut(fields []*registry.DecodedField) *pbgear.Identity_Unknown {
    return &pbgear.Identity_Unknown{
    }
} 
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes15ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_System_KillPrefixCallsalut(fields []*registry.DecodedField) *pbgear.System_KillPrefixCall {
    return &pbgear.System_KillPrefixCall{
                Prefix : []uint32(field.Value),
                Subkeys : uint32(field.Value),
    }
} 
            
                   
func to_Grandpa_ReportEquivocationUnsignedCallsalut(fields []*registry.DecodedField) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
    return &pbgear.Grandpa_ReportEquivocationUnsignedCall{
                EquivocationProof: to_Grandpa_SpConsensusGrandpaEquivocationProof(field.Value),
                KeyOwnerProof: to_Grandpa_SpSessionMembershipProof(field.Value),
    }
}   
func to_ElectionProviderMultiPhase_SetEmergencyElectionResultCallsalut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
    return &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{
                Supports: to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(field.Value),
    }
}
            
                func to_repeated_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(in *registry.DecodedField) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
                    return []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{
                        // TODO fill all fields
                    }
                }   
func to_NominationPools_Opensalut(fields []*registry.DecodedField) *pbgear.NominationPools_Open {
    return &pbgear.NominationPools_Open{
    }
} 
func to_GearVoucher_Nonesalut(fields []*registry.DecodedField) *pbgear.GearVoucher_None {
    return &pbgear.GearVoucher_None{
    }
} 
func to_System_SetHeapPagesCallsalut(fields []*registry.DecodedField) *pbgear.System_SetHeapPagesCall {
    return &pbgear.System_SetHeapPagesCall{
                Pages : uint64(field.Value),
    }
}
            
                   
func to_BagsList_Heaviersalut(fields []*registry.DecodedField) *pbgear.BagsList_Heavier {
    return &pbgear.BagsList_Heavier{
                Value: to_oneof_BagsList_Value(field.Value),
    }
}  
func to_ConvictionVoting_UndelegateCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_UndelegateCall {
    return &pbgear.ConvictionVoting_UndelegateCall{
                Class : uint32(field.Value),
    }
}  
func to_Identity_Raw9salut(fields []*registry.DecodedField) *pbgear.Identity_Raw9 {
    return &pbgear.Identity_Raw9{
                Value0 : []uint32(field.Value),
    }
}  
func to_ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32salut(fields []*registry.DecodedField) *pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32 {
    return &pbgear.ElectionProviderMultiPhase_TupleCompactUint32ElectionProviderMultiPhaseVotes4ListCompactUint32{
                Value0: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
                Value1: to_repeated_ElectionProviderMultiPhase_TupleCompactUint32CompactSpArithmeticPerThingsPerU16(field.Value),
                Value2: to_ElectionProviderMultiPhase_CompactUint32(field.Value),
    }
}    
func to_Bounties_TupleNullsalut(fields []*registry.DecodedField) *pbgear.Bounties_TupleNull {
    return &pbgear.Bounties_TupleNull{
    }
} 
func to_Staking_CompactStringsalut(fields []*registry.DecodedField) *pbgear.Staking_CompactString {
    return &pbgear.Staking_CompactString{
                Value : string(field.Value),
    }
}
            
                   
func to_Staking_ScaleValidatorCountCallsalut(fields []*registry.DecodedField) *pbgear.Staking_ScaleValidatorCountCall {
    return &pbgear.Staking_ScaleValidatorCountCall{
                Factor: to_Staking_SpArithmeticPerThingsPercent(field.Value),
    }
}
            
                func to_Staking_SpArithmeticPerThingsPercent(in *registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
                    return &pbgear.Staking_SpArithmeticPerThingsPercent{
                        // TODO fill all fields
                    }
                }   
func to_Staking_SetStakingConfigsCallsalut(fields []*registry.DecodedField) *pbgear.Staking_SetStakingConfigsCall {
    return &pbgear.Staking_SetStakingConfigsCall{
                MinNominatorBond: to_Staking_MinNominatorBond(field.Value),
                MinValidatorBond: to_Staking_MinValidatorBond(field.Value),
                MaxNominatorCount: to_Staking_MaxNominatorCount(field.Value),
                MaxValidatorCount: to_Staking_MaxValidatorCount(field.Value),
                ChillThreshold: to_Staking_ChillThreshold(field.Value),
                MinCommission: to_Staking_MinCommission(field.Value),
    }
}
            
                func to_Staking_MinNominatorBond(in *registry.DecodedField) *pbgear.Staking_MinNominatorBond {
                    return &pbgear.Staking_MinNominatorBond{
                        // TODO fill all fields
                    }
                }  
            
                func to_Staking_MinValidatorBond(in *registry.DecodedField) *pbgear.Staking_MinValidatorBond {
                    return &pbgear.Staking_MinValidatorBond{
                        // TODO fill all fields
                    }
                }  
            
                func to_Staking_MaxNominatorCount(in *registry.DecodedField) *pbgear.Staking_MaxNominatorCount {
                    return &pbgear.Staking_MaxNominatorCount{
                        // TODO fill all fields
                    }
                }  
            
                func to_Staking_MaxValidatorCount(in *registry.DecodedField) *pbgear.Staking_MaxValidatorCount {
                    return &pbgear.Staking_MaxValidatorCount{
                        // TODO fill all fields
                    }
                }  
            
                func to_Staking_ChillThreshold(in *registry.DecodedField) *pbgear.Staking_ChillThreshold {
                    return &pbgear.Staking_ChillThreshold{
                        // TODO fill all fields
                    }
                }  
            
                func to_Staking_MinCommission(in *registry.DecodedField) *pbgear.Staking_MinCommission {
                    return &pbgear.Staking_MinCommission{
                        // TODO fill all fields
                    }
                }   
func to_ConvictionVoting_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_SpCoreCryptoAccountId32 {
    return &pbgear.ConvictionVoting_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_ConvictionVoting_RemoveOtherVoteCallsalut(fields []*registry.DecodedField) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
    return &pbgear.ConvictionVoting_RemoveOtherVoteCall{
                Target: to_ConvictionVoting_Target(field.Value),
                Class : uint32(field.Value),
                Index : uint32(field.Value),
    }
}    
func to_Identity_Sha256salut(fields []*registry.DecodedField) *pbgear.Identity_Sha256 {
    return &pbgear.Identity_Sha256{
                Value0 : []uint32(field.Value),
    }
}  
func to_NominationPools_PermissionlessWithdrawsalut(fields []*registry.DecodedField) *pbgear.NominationPools_PermissionlessWithdraw {
    return &pbgear.NominationPools_PermissionlessWithdraw{
    }
} 
func to_Staking_SpCoreCryptoAccountId32salut(fields []*registry.DecodedField) *pbgear.Staking_SpCoreCryptoAccountId32 {
    return &pbgear.Staking_SpCoreCryptoAccountId32{
                Value0 : []uint32(field.Value),
    }
}  
func to_Staking_SpArithmeticPerThingsPercentsalut(fields []*registry.DecodedField) *pbgear.Staking_SpArithmeticPerThingsPercent {
    return &pbgear.Staking_SpArithmeticPerThingsPercent{
                Factor : uint32(field.Value),
    }
}   
