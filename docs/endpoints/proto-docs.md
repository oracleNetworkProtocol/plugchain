---
order: 6
---
# Protobuf Documentation
<!-- This file is auto-generated. Please do not modify it yourself. -->
<a name="top"></a>

## Table of Contents

- [cosmos/auth/v1beta1/auth.proto](#cosmos/auth/v1beta1/auth.proto)
    - [BaseAccount](#cosmos.auth.v1beta1.BaseAccount)
    - [ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount)
    - [Params](#cosmos.auth.v1beta1.Params)
  
- [cosmos/auth/v1beta1/genesis.proto](#cosmos/auth/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.auth.v1beta1.GenesisState)
  
- [cosmos/auth/v1beta1/query.proto](#cosmos/auth/v1beta1/query.proto)
    - [QueryAccountRequest](#cosmos.auth.v1beta1.QueryAccountRequest)
    - [QueryAccountResponse](#cosmos.auth.v1beta1.QueryAccountResponse)
    - [QueryAccountsRequest](#cosmos.auth.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#cosmos.auth.v1beta1.QueryAccountsResponse)
    - [QueryParamsRequest](#cosmos.auth.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.auth.v1beta1.QueryParamsResponse)
  
    - [Query](#cosmos.auth.v1beta1.Query)
  
- [cosmos/bank/v1beta1/authz.proto](#cosmos/bank/v1beta1/authz.proto)
    - [SendAuthorization](#cosmos.bank.v1beta1.SendAuthorization)
  
- [cosmos/bank/v1beta1/bank.proto](#cosmos/bank/v1beta1/bank.proto)
    - [DenomUnit](#cosmos.bank.v1beta1.DenomUnit)
    - [Input](#cosmos.bank.v1beta1.Input)
    - [Metadata](#cosmos.bank.v1beta1.Metadata)
    - [Output](#cosmos.bank.v1beta1.Output)
    - [Params](#cosmos.bank.v1beta1.Params)
    - [SendEnabled](#cosmos.bank.v1beta1.SendEnabled)
    - [Supply](#cosmos.bank.v1beta1.Supply)
  
- [cosmos/bank/v1beta1/genesis.proto](#cosmos/bank/v1beta1/genesis.proto)
    - [Balance](#cosmos.bank.v1beta1.Balance)
    - [GenesisState](#cosmos.bank.v1beta1.GenesisState)
  
- [cosmos/bank/v1beta1/query.proto](#cosmos/bank/v1beta1/query.proto)
    - [QueryAllBalancesRequest](#cosmos.bank.v1beta1.QueryAllBalancesRequest)
    - [QueryAllBalancesResponse](#cosmos.bank.v1beta1.QueryAllBalancesResponse)
    - [QueryBalanceRequest](#cosmos.bank.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#cosmos.bank.v1beta1.QueryBalanceResponse)
    - [QueryDenomMetadataRequest](#cosmos.bank.v1beta1.QueryDenomMetadataRequest)
    - [QueryDenomMetadataResponse](#cosmos.bank.v1beta1.QueryDenomMetadataResponse)
    - [QueryDenomsMetadataRequest](#cosmos.bank.v1beta1.QueryDenomsMetadataRequest)
    - [QueryDenomsMetadataResponse](#cosmos.bank.v1beta1.QueryDenomsMetadataResponse)
    - [QueryParamsRequest](#cosmos.bank.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.bank.v1beta1.QueryParamsResponse)
    - [QuerySupplyOfRequest](#cosmos.bank.v1beta1.QuerySupplyOfRequest)
    - [QuerySupplyOfResponse](#cosmos.bank.v1beta1.QuerySupplyOfResponse)
    - [QueryTotalSupplyRequest](#cosmos.bank.v1beta1.QueryTotalSupplyRequest)
    - [QueryTotalSupplyResponse](#cosmos.bank.v1beta1.QueryTotalSupplyResponse)
  
    - [Query](#cosmos.bank.v1beta1.Query)
  
- [cosmos/bank/v1beta1/tx.proto](#cosmos/bank/v1beta1/tx.proto)
    - [MsgMultiSend](#cosmos.bank.v1beta1.MsgMultiSend)
    - [MsgMultiSendResponse](#cosmos.bank.v1beta1.MsgMultiSendResponse)
    - [MsgSend](#cosmos.bank.v1beta1.MsgSend)
    - [MsgSendResponse](#cosmos.bank.v1beta1.MsgSendResponse)
  
    - [Msg](#cosmos.bank.v1beta1.Msg)
  
- [cosmos/crypto/ed25519/keys.proto](#cosmos/crypto/ed25519/keys.proto)
    - [PrivKey](#cosmos.crypto.ed25519.PrivKey)
    - [PubKey](#cosmos.crypto.ed25519.PubKey)
  
- [cosmos/crypto/multisig/keys.proto](#cosmos/crypto/multisig/keys.proto)
    - [LegacyAminoPubKey](#cosmos.crypto.multisig.LegacyAminoPubKey)
  
- [cosmos/crypto/multisig/v1beta1/multisig.proto](#cosmos/crypto/multisig/v1beta1/multisig.proto)
    - [CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray)
    - [MultiSignature](#cosmos.crypto.multisig.v1beta1.MultiSignature)
  
- [cosmos/crypto/secp256k1/keys.proto](#cosmos/crypto/secp256k1/keys.proto)
    - [PrivKey](#cosmos.crypto.secp256k1.PrivKey)
    - [PubKey](#cosmos.crypto.secp256k1.PubKey)
  
- [cosmos/crypto/secp256r1/keys.proto](#cosmos/crypto/secp256r1/keys.proto)
    - [PrivKey](#cosmos.crypto.secp256r1.PrivKey)
    - [PubKey](#cosmos.crypto.secp256r1.PubKey)
  
- [cosmos/gov/v1beta1/gov.proto](#cosmos/gov/v1beta1/gov.proto)
    - [Deposit](#cosmos.gov.v1beta1.Deposit)
    - [DepositParams](#cosmos.gov.v1beta1.DepositParams)
    - [Proposal](#cosmos.gov.v1beta1.Proposal)
    - [TallyParams](#cosmos.gov.v1beta1.TallyParams)
    - [TallyResult](#cosmos.gov.v1beta1.TallyResult)
    - [TextProposal](#cosmos.gov.v1beta1.TextProposal)
    - [Vote](#cosmos.gov.v1beta1.Vote)
    - [VotingParams](#cosmos.gov.v1beta1.VotingParams)
    - [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption)
  
    - [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus)
    - [VoteOption](#cosmos.gov.v1beta1.VoteOption)
  
- [cosmos/gov/v1beta1/genesis.proto](#cosmos/gov/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.gov.v1beta1.GenesisState)
  
- [cosmos/gov/v1beta1/query.proto](#cosmos/gov/v1beta1/query.proto)
    - [QueryDepositRequest](#cosmos.gov.v1beta1.QueryDepositRequest)
    - [QueryDepositResponse](#cosmos.gov.v1beta1.QueryDepositResponse)
    - [QueryDepositsRequest](#cosmos.gov.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#cosmos.gov.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#cosmos.gov.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.gov.v1beta1.QueryParamsResponse)
    - [QueryProposalRequest](#cosmos.gov.v1beta1.QueryProposalRequest)
    - [QueryProposalResponse](#cosmos.gov.v1beta1.QueryProposalResponse)
    - [QueryProposalsRequest](#cosmos.gov.v1beta1.QueryProposalsRequest)
    - [QueryProposalsResponse](#cosmos.gov.v1beta1.QueryProposalsResponse)
    - [QueryTallyResultRequest](#cosmos.gov.v1beta1.QueryTallyResultRequest)
    - [QueryTallyResultResponse](#cosmos.gov.v1beta1.QueryTallyResultResponse)
    - [QueryVoteRequest](#cosmos.gov.v1beta1.QueryVoteRequest)
    - [QueryVoteResponse](#cosmos.gov.v1beta1.QueryVoteResponse)
    - [QueryVotesRequest](#cosmos.gov.v1beta1.QueryVotesRequest)
    - [QueryVotesResponse](#cosmos.gov.v1beta1.QueryVotesResponse)
  
    - [Query](#cosmos.gov.v1beta1.Query)
  
- [cosmos/gov/v1beta1/tx.proto](#cosmos/gov/v1beta1/tx.proto)
    - [MsgDeposit](#cosmos.gov.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#cosmos.gov.v1beta1.MsgDepositResponse)
    - [MsgSubmitProposal](#cosmos.gov.v1beta1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#cosmos.gov.v1beta1.MsgSubmitProposalResponse)
    - [MsgVote](#cosmos.gov.v1beta1.MsgVote)
    - [MsgVoteResponse](#cosmos.gov.v1beta1.MsgVoteResponse)
    - [MsgVoteWeighted](#cosmos.gov.v1beta1.MsgVoteWeighted)
    - [MsgVoteWeightedResponse](#cosmos.gov.v1beta1.MsgVoteWeightedResponse)
  
    - [Msg](#cosmos.gov.v1beta1.Msg)
  
- [cosmos/mint/v1beta1/mint.proto](#cosmos/mint/v1beta1/mint.proto)
    - [Minter](#cosmos.mint.v1beta1.Minter)
    - [Params](#cosmos.mint.v1beta1.Params)
  
- [cosmos/mint/v1beta1/genesis.proto](#cosmos/mint/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.mint.v1beta1.GenesisState)
  
- [cosmos/mint/v1beta1/query.proto](#cosmos/mint/v1beta1/query.proto)
    - [QueryAnnualProvisionsRequest](#cosmos.mint.v1beta1.QueryAnnualProvisionsRequest)
    - [QueryAnnualProvisionsResponse](#cosmos.mint.v1beta1.QueryAnnualProvisionsResponse)
    - [QueryInflationRequest](#cosmos.mint.v1beta1.QueryInflationRequest)
    - [QueryInflationResponse](#cosmos.mint.v1beta1.QueryInflationResponse)
    - [QueryParamsRequest](#cosmos.mint.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.mint.v1beta1.QueryParamsResponse)
  
    - [Query](#cosmos.mint.v1beta1.Query)
  
- [cosmos/slashing/v1beta1/slashing.proto](#cosmos/slashing/v1beta1/slashing.proto)
    - [Params](#cosmos.slashing.v1beta1.Params)
    - [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo)
  
- [cosmos/slashing/v1beta1/genesis.proto](#cosmos/slashing/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.slashing.v1beta1.GenesisState)
    - [MissedBlock](#cosmos.slashing.v1beta1.MissedBlock)
    - [SigningInfo](#cosmos.slashing.v1beta1.SigningInfo)
    - [ValidatorMissedBlocks](#cosmos.slashing.v1beta1.ValidatorMissedBlocks)
  
- [cosmos/slashing/v1beta1/query.proto](#cosmos/slashing/v1beta1/query.proto)
    - [QueryParamsRequest](#cosmos.slashing.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.slashing.v1beta1.QueryParamsResponse)
    - [QuerySigningInfoRequest](#cosmos.slashing.v1beta1.QuerySigningInfoRequest)
    - [QuerySigningInfoResponse](#cosmos.slashing.v1beta1.QuerySigningInfoResponse)
    - [QuerySigningInfosRequest](#cosmos.slashing.v1beta1.QuerySigningInfosRequest)
    - [QuerySigningInfosResponse](#cosmos.slashing.v1beta1.QuerySigningInfosResponse)
  
    - [Query](#cosmos.slashing.v1beta1.Query)
  
- [cosmos/slashing/v1beta1/tx.proto](#cosmos/slashing/v1beta1/tx.proto)
    - [MsgUnjail](#cosmos.slashing.v1beta1.MsgUnjail)
    - [MsgUnjailResponse](#cosmos.slashing.v1beta1.MsgUnjailResponse)
  
    - [Msg](#cosmos.slashing.v1beta1.Msg)
  
- [cosmos/staking/v1beta1/authz.proto](#cosmos/staking/v1beta1/authz.proto)
    - [StakeAuthorization](#cosmos.staking.v1beta1.StakeAuthorization)
    - [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators)
  
    - [AuthorizationType](#cosmos.staking.v1beta1.AuthorizationType)
  
- [cosmos/staking/v1beta1/staking.proto](#cosmos/staking/v1beta1/staking.proto)
    - [Commission](#cosmos.staking.v1beta1.Commission)
    - [CommissionRates](#cosmos.staking.v1beta1.CommissionRates)
    - [DVPair](#cosmos.staking.v1beta1.DVPair)
    - [DVPairs](#cosmos.staking.v1beta1.DVPairs)
    - [DVVTriplet](#cosmos.staking.v1beta1.DVVTriplet)
    - [DVVTriplets](#cosmos.staking.v1beta1.DVVTriplets)
    - [Delegation](#cosmos.staking.v1beta1.Delegation)
    - [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse)
    - [Description](#cosmos.staking.v1beta1.Description)
    - [HistoricalInfo](#cosmos.staking.v1beta1.HistoricalInfo)
    - [Params](#cosmos.staking.v1beta1.Params)
    - [Pool](#cosmos.staking.v1beta1.Pool)
    - [Redelegation](#cosmos.staking.v1beta1.Redelegation)
    - [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry)
    - [RedelegationEntryResponse](#cosmos.staking.v1beta1.RedelegationEntryResponse)
    - [RedelegationResponse](#cosmos.staking.v1beta1.RedelegationResponse)
    - [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation)
    - [UnbondingDelegationEntry](#cosmos.staking.v1beta1.UnbondingDelegationEntry)
    - [ValAddresses](#cosmos.staking.v1beta1.ValAddresses)
    - [Validator](#cosmos.staking.v1beta1.Validator)
  
    - [BondStatus](#cosmos.staking.v1beta1.BondStatus)
  
- [cosmos/staking/v1beta1/genesis.proto](#cosmos/staking/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.staking.v1beta1.GenesisState)
    - [LastValidatorPower](#cosmos.staking.v1beta1.LastValidatorPower)
  
- [cosmos/staking/v1beta1/query.proto](#cosmos/staking/v1beta1/query.proto)
    - [QueryDelegationRequest](#cosmos.staking.v1beta1.QueryDelegationRequest)
    - [QueryDelegationResponse](#cosmos.staking.v1beta1.QueryDelegationResponse)
    - [QueryDelegatorDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest)
    - [QueryDelegatorDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse)
    - [QueryDelegatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest)
    - [QueryDelegatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse)
    - [QueryDelegatorValidatorRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorRequest)
    - [QueryDelegatorValidatorResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorResponse)
    - [QueryDelegatorValidatorsRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest)
    - [QueryDelegatorValidatorsResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse)
    - [QueryHistoricalInfoRequest](#cosmos.staking.v1beta1.QueryHistoricalInfoRequest)
    - [QueryHistoricalInfoResponse](#cosmos.staking.v1beta1.QueryHistoricalInfoResponse)
    - [QueryParamsRequest](#cosmos.staking.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.staking.v1beta1.QueryParamsResponse)
    - [QueryPoolRequest](#cosmos.staking.v1beta1.QueryPoolRequest)
    - [QueryPoolResponse](#cosmos.staking.v1beta1.QueryPoolResponse)
    - [QueryRedelegationsRequest](#cosmos.staking.v1beta1.QueryRedelegationsRequest)
    - [QueryRedelegationsResponse](#cosmos.staking.v1beta1.QueryRedelegationsResponse)
    - [QueryUnbondingDelegationRequest](#cosmos.staking.v1beta1.QueryUnbondingDelegationRequest)
    - [QueryUnbondingDelegationResponse](#cosmos.staking.v1beta1.QueryUnbondingDelegationResponse)
    - [QueryValidatorDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorDelegationsRequest)
    - [QueryValidatorDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorDelegationsResponse)
    - [QueryValidatorRequest](#cosmos.staking.v1beta1.QueryValidatorRequest)
    - [QueryValidatorResponse](#cosmos.staking.v1beta1.QueryValidatorResponse)
    - [QueryValidatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest)
    - [QueryValidatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse)
    - [QueryValidatorsRequest](#cosmos.staking.v1beta1.QueryValidatorsRequest)
    - [QueryValidatorsResponse](#cosmos.staking.v1beta1.QueryValidatorsResponse)
  
    - [Query](#cosmos.staking.v1beta1.Query)
  
- [cosmos/staking/v1beta1/tx.proto](#cosmos/staking/v1beta1/tx.proto)
    - [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate)
    - [MsgBeginRedelegateResponse](#cosmos.staking.v1beta1.MsgBeginRedelegateResponse)
    - [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator)
    - [MsgCreateValidatorResponse](#cosmos.staking.v1beta1.MsgCreateValidatorResponse)
    - [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate)
    - [MsgDelegateResponse](#cosmos.staking.v1beta1.MsgDelegateResponse)
    - [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator)
    - [MsgEditValidatorResponse](#cosmos.staking.v1beta1.MsgEditValidatorResponse)
    - [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate)
    - [MsgUndelegateResponse](#cosmos.staking.v1beta1.MsgUndelegateResponse)
  
    - [Msg](#cosmos.staking.v1beta1.Msg)
  
- [cosmos/tx/signing/v1beta1/signing.proto](#cosmos/tx/signing/v1beta1/signing.proto)
    - [SignatureDescriptor](#cosmos.tx.signing.v1beta1.SignatureDescriptor)
    - [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data)
    - [SignatureDescriptor.Data.Multi](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi)
    - [SignatureDescriptor.Data.Single](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single)
    - [SignatureDescriptors](#cosmos.tx.signing.v1beta1.SignatureDescriptors)
  
    - [SignMode](#cosmos.tx.signing.v1beta1.SignMode)
  
- [cosmos/tx/v1beta1/tx.proto](#cosmos/tx/v1beta1/tx.proto)
    - [AuthInfo](#cosmos.tx.v1beta1.AuthInfo)
    - [Fee](#cosmos.tx.v1beta1.Fee)
    - [ModeInfo](#cosmos.tx.v1beta1.ModeInfo)
    - [ModeInfo.Multi](#cosmos.tx.v1beta1.ModeInfo.Multi)
    - [ModeInfo.Single](#cosmos.tx.v1beta1.ModeInfo.Single)
    - [SignDoc](#cosmos.tx.v1beta1.SignDoc)
    - [SignerInfo](#cosmos.tx.v1beta1.SignerInfo)
    - [Tx](#cosmos.tx.v1beta1.Tx)
    - [TxBody](#cosmos.tx.v1beta1.TxBody)
    - [TxRaw](#cosmos.tx.v1beta1.TxRaw)
  
- [cosmos/tx/v1beta1/service.proto](#cosmos/tx/v1beta1/service.proto)
    - [BroadcastTxRequest](#cosmos.tx.v1beta1.BroadcastTxRequest)
    - [BroadcastTxResponse](#cosmos.tx.v1beta1.BroadcastTxResponse)
    - [GetTxRequest](#cosmos.tx.v1beta1.GetTxRequest)
    - [GetTxResponse](#cosmos.tx.v1beta1.GetTxResponse)
    - [GetTxsEventRequest](#cosmos.tx.v1beta1.GetTxsEventRequest)
    - [GetTxsEventResponse](#cosmos.tx.v1beta1.GetTxsEventResponse)
    - [SimulateRequest](#cosmos.tx.v1beta1.SimulateRequest)
    - [SimulateResponse](#cosmos.tx.v1beta1.SimulateResponse)
  
    - [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode)
    - [OrderBy](#cosmos.tx.v1beta1.OrderBy)
  
    - [Service](#cosmos.tx.v1beta1.Service)
  
- [ethermint/crypto/v1/ethsecp256k1/keys.proto](#ethermint/crypto/v1/ethsecp256k1/keys.proto)
    - [PrivKey](#ethermint.crypto.v1.ethsecp256k1.PrivKey)
    - [PubKey](#ethermint.crypto.v1.ethsecp256k1.PubKey)
  
- [ethermint/types/v1/account.proto](#ethermint/types/v1/account.proto)
    - [EthAccount](#ethermint.types.v1.EthAccount)
  
- [ethermint/types/v1/web3.proto](#ethermint/types/v1/web3.proto)
    - [ExtensionOptionsWeb3Tx](#ethermint.types.v1.ExtensionOptionsWeb3Tx)
  
- [prc10/prc10.proto](#prc10/prc10.proto)
    - [Params](#plugchain.prc10.Params)
    - [Token](#plugchain.prc10.Token)
  
- [prc10/genesis.proto](#prc10/genesis.proto)
    - [GenesisState](#plugchain.prc10.GenesisState)
  
- [prc10/query.proto](#prc10/query.proto)
    - [QueryFeesRequest](#plugchain.prc10.QueryFeesRequest)
    - [QueryFeesResponse](#plugchain.prc10.QueryFeesResponse)
    - [QueryParamsRequest](#plugchain.prc10.QueryParamsRequest)
    - [QueryParamsResponse](#plugchain.prc10.QueryParamsResponse)
    - [QueryTokenRequest](#plugchain.prc10.QueryTokenRequest)
    - [QueryTokenResponse](#plugchain.prc10.QueryTokenResponse)
    - [QueryTokensRequest](#plugchain.prc10.QueryTokensRequest)
    - [QueryTokensResponse](#plugchain.prc10.QueryTokensResponse)
  
    - [Query](#plugchain.prc10.Query)
  
- [prc10/tx.proto](#prc10/tx.proto)
    - [MsgBurnToken](#plugchain.prc10.MsgBurnToken)
    - [MsgBurnTokenResponse](#plugchain.prc10.MsgBurnTokenResponse)
    - [MsgEditToken](#plugchain.prc10.MsgEditToken)
    - [MsgEditTokenResponse](#plugchain.prc10.MsgEditTokenResponse)
    - [MsgIssueToken](#plugchain.prc10.MsgIssueToken)
    - [MsgIssueTokenResponse](#plugchain.prc10.MsgIssueTokenResponse)
    - [MsgMintToken](#plugchain.prc10.MsgMintToken)
    - [MsgMintTokenResponse](#plugchain.prc10.MsgMintTokenResponse)
    - [MsgTransferOwnerToken](#plugchain.prc10.MsgTransferOwnerToken)
    - [MsgTransferOwnerTokenResponse](#plugchain.prc10.MsgTransferOwnerTokenResponse)
  
    - [Msg](#plugchain.prc10.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cosmos/auth/v1beta1/auth.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/auth.proto



<a name="cosmos.auth.v1beta1.BaseAccount"></a>

### BaseAccount
BaseAccount defines a base account type. It contains all the necessary fields
for basic account functionality. Any custom account type should extend this
type for additional functionality (e.g. vesting).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `pub_key` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `account_number` | [uint64](#uint64) |  |  |
| `sequence` | [uint64](#uint64) |  |  |






<a name="cosmos.auth.v1beta1.ModuleAccount"></a>

### ModuleAccount
ModuleAccount defines an account for modules that holds coins on a pool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_account` | [BaseAccount](#cosmos.auth.v1beta1.BaseAccount) |  |  |
| `name` | [string](#string) |  |  |
| `permissions` | [string](#string) | repeated |  |






<a name="cosmos.auth.v1beta1.Params"></a>

### Params
Params defines the parameters for the auth module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_memo_characters` | [uint64](#uint64) |  |  |
| `tx_sig_limit` | [uint64](#uint64) |  |  |
| `tx_size_cost_per_byte` | [uint64](#uint64) |  |  |
| `sig_verify_cost_ed25519` | [uint64](#uint64) |  |  |
| `sig_verify_cost_secp256k1` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/auth/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/genesis.proto



<a name="cosmos.auth.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the auth module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.auth.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated | accounts are the accounts present at genesis. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/auth/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/query.proto



<a name="cosmos.auth.v1beta1.QueryAccountRequest"></a>

### QueryAccountRequest
QueryAccountRequest is the request type for the Query/Account RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address defines the address to query for. |






<a name="cosmos.auth.v1beta1.QueryAccountResponse"></a>

### QueryAccountResponse
QueryAccountResponse is the response type for the Query/Account RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account` | [google.protobuf.Any](#google.protobuf.Any) |  | account defines the account of the corresponding address. |






<a name="cosmos.auth.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest is the request type for the Query/Accounts RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.auth.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse is the response type for the Query/Accounts RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated | accounts are the existing accounts |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.auth.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="cosmos.auth.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.auth.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.auth.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Accounts` | [QueryAccountsRequest](#cosmos.auth.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#cosmos.auth.v1beta1.QueryAccountsResponse) | Accounts returns all the existing accounts

Since: cosmos-sdk 0.43 | GET|/cosmos/auth/v1beta1/accounts|
| `Account` | [QueryAccountRequest](#cosmos.auth.v1beta1.QueryAccountRequest) | [QueryAccountResponse](#cosmos.auth.v1beta1.QueryAccountResponse) | Account returns account details based on address. | GET|/cosmos/auth/v1beta1/accounts/{address}|
| `Params` | [QueryParamsRequest](#cosmos.auth.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.auth.v1beta1.QueryParamsResponse) | Params queries all parameters. | GET|/cosmos/auth/v1beta1/params|

 <!-- end services -->



<a name="cosmos/bank/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/authz.proto



<a name="cosmos.bank.v1beta1.SendAuthorization"></a>

### SendAuthorization
SendAuthorization allows the grantee to spend up to spend_limit coins from
the granter's account.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `spend_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/bank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/bank.proto



<a name="cosmos.bank.v1beta1.DenomUnit"></a>

### DenomUnit
DenomUnit represents a struct that describes a given
denomination unit of the basic token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom represents the string name of the given denom unit (e.g uatom). |
| `exponent` | [uint32](#uint32) |  | exponent represents power of 10 exponent that one must raise the base_denom to in order to equal the given DenomUnit's denom 1 denom = 1^exponent base_denom (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with exponent = 6, thus: 1 atom = 10^6 uatom). |
| `aliases` | [string](#string) | repeated | aliases is a list of string aliases for the given denom |






<a name="cosmos.bank.v1beta1.Input"></a>

### Input
Input models transaction input.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.Metadata"></a>

### Metadata
Metadata represents a struct that describes
a basic token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [string](#string) |  |  |
| `denom_units` | [DenomUnit](#cosmos.bank.v1beta1.DenomUnit) | repeated | denom_units represents the list of DenomUnit's for a given coin |
| `base` | [string](#string) |  | base represents the base denom (should be the DenomUnit with exponent = 0). |
| `display` | [string](#string) |  | display indicates the suggested denom that should be displayed in clients. |
| `name` | [string](#string) |  | name defines the name of the token (eg: Cosmos Atom)

Since: cosmos-sdk 0.43 |
| `symbol` | [string](#string) |  | symbol is the token symbol usually shown on exchanges (eg: ATOM). This can be the same as the display.

Since: cosmos-sdk 0.43 |






<a name="cosmos.bank.v1beta1.Output"></a>

### Output
Output models transaction outputs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.Params"></a>

### Params
Params defines the parameters for the bank module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `send_enabled` | [SendEnabled](#cosmos.bank.v1beta1.SendEnabled) | repeated |  |
| `default_send_enabled` | [bool](#bool) |  |  |






<a name="cosmos.bank.v1beta1.SendEnabled"></a>

### SendEnabled
SendEnabled maps coin denom to a send_enabled status (whether a denom is
sendable).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `enabled` | [bool](#bool) |  |  |






<a name="cosmos.bank.v1beta1.Supply"></a>

### Supply
Supply represents a struct that passively keeps track of the total supply
amounts in the network.
This message is deprecated now that supply is indexed by denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/genesis.proto



<a name="cosmos.bank.v1beta1.Balance"></a>

### Balance
Balance defines an account address and balance pair used in the bank module's
genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the balance holder. |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | coins defines the different coins this balance holds. |






<a name="cosmos.bank.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the bank module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.bank.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `balances` | [Balance](#cosmos.bank.v1beta1.Balance) | repeated | balances is an array containing the balances of all the accounts. |
| `supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | supply represents the total supply. If it is left empty, then supply will be calculated based on the provided balances. Otherwise, it will be used to validate that the sum of the balances equals this amount. |
| `denom_metadata` | [Metadata](#cosmos.bank.v1beta1.Metadata) | repeated | denom_metadata defines the metadata of the differents coins. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/query.proto



<a name="cosmos.bank.v1beta1.QueryAllBalancesRequest"></a>

### QueryAllBalancesRequest
QueryBalanceRequest is the request type for the Query/AllBalances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address to query balances for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QueryAllBalancesResponse"></a>

### QueryAllBalancesResponse
QueryAllBalancesResponse is the response type for the Query/AllBalances RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | balances is the balances of all the coins. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the request type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address to query balances for. |
| `denom` | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="cosmos.bank.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the response type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | balance is the balance of the coin. |






<a name="cosmos.bank.v1beta1.QueryDenomMetadataRequest"></a>

### QueryDenomMetadataRequest
QueryDenomMetadataRequest is the request type for the Query/DenomMetadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query the metadata for. |






<a name="cosmos.bank.v1beta1.QueryDenomMetadataResponse"></a>

### QueryDenomMetadataResponse
QueryDenomMetadataResponse is the response type for the Query/DenomMetadata RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [Metadata](#cosmos.bank.v1beta1.Metadata) |  | metadata describes and provides all the client information for the requested token. |






<a name="cosmos.bank.v1beta1.QueryDenomsMetadataRequest"></a>

### QueryDenomsMetadataRequest
QueryDenomsMetadataRequest is the request type for the Query/DenomsMetadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QueryDenomsMetadataResponse"></a>

### QueryDenomsMetadataResponse
QueryDenomsMetadataResponse is the response type for the Query/DenomsMetadata RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadatas` | [Metadata](#cosmos.bank.v1beta1.Metadata) | repeated | metadata provides the client information for all the registered tokens. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/bank parameters.






<a name="cosmos.bank.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/bank parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.bank.v1beta1.Params) |  |  |






<a name="cosmos.bank.v1beta1.QuerySupplyOfRequest"></a>

### QuerySupplyOfRequest
QuerySupplyOfRequest is the request type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="cosmos.bank.v1beta1.QuerySupplyOfResponse"></a>

### QuerySupplyOfResponse
QuerySupplyOfResponse is the response type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the supply of the coin. |






<a name="cosmos.bank.v1beta1.QueryTotalSupplyRequest"></a>

### QueryTotalSupplyRequest
QueryTotalSupplyRequest is the request type for the Query/TotalSupply RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request.

Since: cosmos-sdk 0.43 |






<a name="cosmos.bank.v1beta1.QueryTotalSupplyResponse"></a>

### QueryTotalSupplyResponse
QueryTotalSupplyResponse is the response type for the Query/TotalSupply RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | supply is the supply of the coins |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response.

Since: cosmos-sdk 0.43 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.bank.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balance` | [QueryBalanceRequest](#cosmos.bank.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#cosmos.bank.v1beta1.QueryBalanceResponse) | Balance queries the balance of a single coin for a single account. | GET|/cosmos/bank/v1beta1/balances/{address}/by_denom|
| `AllBalances` | [QueryAllBalancesRequest](#cosmos.bank.v1beta1.QueryAllBalancesRequest) | [QueryAllBalancesResponse](#cosmos.bank.v1beta1.QueryAllBalancesResponse) | AllBalances queries the balance of all coins for a single account. | GET|/cosmos/bank/v1beta1/balances/{address}|
| `TotalSupply` | [QueryTotalSupplyRequest](#cosmos.bank.v1beta1.QueryTotalSupplyRequest) | [QueryTotalSupplyResponse](#cosmos.bank.v1beta1.QueryTotalSupplyResponse) | TotalSupply queries the total supply of all coins. | GET|/cosmos/bank/v1beta1/supply|
| `SupplyOf` | [QuerySupplyOfRequest](#cosmos.bank.v1beta1.QuerySupplyOfRequest) | [QuerySupplyOfResponse](#cosmos.bank.v1beta1.QuerySupplyOfResponse) | SupplyOf queries the supply of a single coin. | GET|/cosmos/bank/v1beta1/supply/{denom}|
| `Params` | [QueryParamsRequest](#cosmos.bank.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.bank.v1beta1.QueryParamsResponse) | Params queries the parameters of x/bank module. | GET|/cosmos/bank/v1beta1/params|
| `DenomMetadata` | [QueryDenomMetadataRequest](#cosmos.bank.v1beta1.QueryDenomMetadataRequest) | [QueryDenomMetadataResponse](#cosmos.bank.v1beta1.QueryDenomMetadataResponse) | DenomsMetadata queries the client metadata of a given coin denomination. | GET|/cosmos/bank/v1beta1/denoms_metadata/{denom}|
| `DenomsMetadata` | [QueryDenomsMetadataRequest](#cosmos.bank.v1beta1.QueryDenomsMetadataRequest) | [QueryDenomsMetadataResponse](#cosmos.bank.v1beta1.QueryDenomsMetadataResponse) | DenomsMetadata queries the client metadata for all registered coin denominations. | GET|/cosmos/bank/v1beta1/denoms_metadata|

 <!-- end services -->



<a name="cosmos/bank/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/tx.proto



<a name="cosmos.bank.v1beta1.MsgMultiSend"></a>

### MsgMultiSend
MsgMultiSend represents an arbitrary multi-in, multi-out send message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inputs` | [Input](#cosmos.bank.v1beta1.Input) | repeated |  |
| `outputs` | [Output](#cosmos.bank.v1beta1.Output) | repeated |  |






<a name="cosmos.bank.v1beta1.MsgMultiSendResponse"></a>

### MsgMultiSendResponse
MsgMultiSendResponse defines the Msg/MultiSend response type.






<a name="cosmos.bank.v1beta1.MsgSend"></a>

### MsgSend
MsgSend represents a message to send coins from one account to another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse defines the Msg/Send response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.bank.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Send` | [MsgSend](#cosmos.bank.v1beta1.MsgSend) | [MsgSendResponse](#cosmos.bank.v1beta1.MsgSendResponse) | Send defines a method for sending coins from one account to another account. | |
| `MultiSend` | [MsgMultiSend](#cosmos.bank.v1beta1.MsgMultiSend) | [MsgMultiSendResponse](#cosmos.bank.v1beta1.MsgMultiSendResponse) | MultiSend defines a method for sending coins from some accounts to other accounts. | |

 <!-- end services -->



<a name="cosmos/crypto/ed25519/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/ed25519/keys.proto



<a name="cosmos.crypto.ed25519.PrivKey"></a>

### PrivKey
Deprecated: PrivKey defines a ed25519 private key.
NOTE: ed25519 keys must not be used in SDK apps except in a tendermint validator context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.ed25519.PubKey"></a>

### PubKey
PubKey is an ed25519 public key for handling Tendermint keys in SDK.
It's needed for Any serialization and SDK compatibility.
It must not be used in a non Tendermint key context because it doesn't implement
ADR-28. Nevertheless, you will like to use ed25519 in app user level
then you must create a new proto message and follow ADR-28 for Address construction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/multisig/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/multisig/keys.proto



<a name="cosmos.crypto.multisig.LegacyAminoPubKey"></a>

### LegacyAminoPubKey
LegacyAminoPubKey specifies a public key type
which nests multiple public keys and a threshold,
it uses legacy amino address rules.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `threshold` | [uint32](#uint32) |  |  |
| `public_keys` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/multisig/v1beta1/multisig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/multisig/v1beta1/multisig.proto



<a name="cosmos.crypto.multisig.v1beta1.CompactBitArray"></a>

### CompactBitArray
CompactBitArray is an implementation of a space efficient bit array.
This is used to ensure that the encoded data takes up a minimal amount of
space after proto encoding.
This is not thread safe, and is not intended for concurrent usage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `extra_bits_stored` | [uint32](#uint32) |  |  |
| `elems` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.multisig.v1beta1.MultiSignature"></a>

### MultiSignature
MultiSignature wraps the signatures from a multisig.LegacyAminoPubKey.
See cosmos.tx.v1betata1.ModeInfo.Multi for how to specify which signers
signed and with which modes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signatures` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/secp256k1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/secp256k1/keys.proto



<a name="cosmos.crypto.secp256k1.PrivKey"></a>

### PrivKey
PrivKey defines a secp256k1 private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.secp256k1.PubKey"></a>

### PubKey
PubKey defines a secp256k1 public key
Key is the compressed form of the pubkey. The first byte depends is a 0x02 byte
if the y-coordinate is the lexicographically largest of the two associated with
the x-coordinate. Otherwise the first byte is a 0x03.
This prefix is followed with the x-coordinate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/secp256r1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/secp256r1/keys.proto
Since: cosmos-sdk 0.43


<a name="cosmos.crypto.secp256r1.PrivKey"></a>

### PrivKey
PrivKey defines a secp256r1 ECDSA private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `secret` | [bytes](#bytes) |  | secret number serialized using big-endian encoding |






<a name="cosmos.crypto.secp256r1.PubKey"></a>

### PubKey
PubKey defines a secp256r1 ECDSA public key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | Point on secp256r1 curve in a compressed representation as specified in section 4.3.6 of ANSI X9.62: https://webstore.ansi.org/standards/ascx9/ansix9621998 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1beta1/gov.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/gov.proto



<a name="cosmos.gov.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount deposited by an account address to an active
proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1beta1.DepositParams"></a>

### DepositParams
DepositParams defines the params for deposits on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Minimum deposit for a proposal to enter voting period. |
| `max_deposit_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months. |






<a name="cosmos.gov.v1beta1.Proposal"></a>

### Proposal
Proposal defines the core field members of a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `status` | [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus) |  |  |
| `final_tally_result` | [TallyResult](#cosmos.gov.v1beta1.TallyResult) |  |  |
| `submit_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `deposit_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `total_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `voting_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `voting_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="cosmos.gov.v1beta1.TallyParams"></a>

### TallyParams
TallyParams defines the params for tallying votes on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `quorum` | [bytes](#bytes) |  | Minimum percentage of total stake needed to vote for a result to be considered valid. |
| `threshold` | [bytes](#bytes) |  | Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. |
| `veto_threshold` | [bytes](#bytes) |  | Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Default value: 1/3. |






<a name="cosmos.gov.v1beta1.TallyResult"></a>

### TallyResult
TallyResult defines a standard tally for a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `yes` | [string](#string) |  |  |
| `abstain` | [string](#string) |  |  |
| `no` | [string](#string) |  |  |
| `no_with_veto` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.TextProposal"></a>

### TextProposal
TextProposal defines a standard text proposal whose changes need to be
manually updated in case of approval.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.Vote"></a>

### Vote
Vote defines a vote on a governance proposal.
A Vote consists of a proposal ID, the voter, and the vote option.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  | **Deprecated.** Deprecated: Prefer to use `options` instead. This field is set in queries if and only if `len(options) == 1` and that option has weight 1. In all other cases, this field will default to VOTE_OPTION_UNSPECIFIED. |
| `options` | [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption) | repeated | Since: cosmos-sdk 0.43 |






<a name="cosmos.gov.v1beta1.VotingParams"></a>

### VotingParams
VotingParams defines the params for voting on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Length of the voting period. |






<a name="cosmos.gov.v1beta1.WeightedVoteOption"></a>

### WeightedVoteOption
WeightedVoteOption defines a unit of vote for vote split.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->


<a name="cosmos.gov.v1beta1.ProposalStatus"></a>

### ProposalStatus
ProposalStatus enumerates the valid statuses of a proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROPOSAL_STATUS_UNSPECIFIED | 0 | PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status. |
| PROPOSAL_STATUS_DEPOSIT_PERIOD | 1 | PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit period. |
| PROPOSAL_STATUS_VOTING_PERIOD | 2 | PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting period. |
| PROPOSAL_STATUS_PASSED | 3 | PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has passed. |
| PROPOSAL_STATUS_REJECTED | 4 | PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has been rejected. |
| PROPOSAL_STATUS_FAILED | 5 | PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has failed. |



<a name="cosmos.gov.v1beta1.VoteOption"></a>

### VoteOption
VoteOption enumerates the valid vote options for a given governance proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_OPTION_UNSPECIFIED | 0 | VOTE_OPTION_UNSPECIFIED defines a no-op vote option. |
| VOTE_OPTION_YES | 1 | VOTE_OPTION_YES defines a yes vote option. |
| VOTE_OPTION_ABSTAIN | 2 | VOTE_OPTION_ABSTAIN defines an abstain vote option. |
| VOTE_OPTION_NO | 3 | VOTE_OPTION_NO defines a no vote option. |
| VOTE_OPTION_NO_WITH_VETO | 4 | VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/genesis.proto



<a name="cosmos.gov.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the gov module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `starting_proposal_id` | [uint64](#uint64) |  | starting_proposal_id is the ID of the starting proposal. |
| `deposits` | [Deposit](#cosmos.gov.v1beta1.Deposit) | repeated | deposits defines all the deposits present at genesis. |
| `votes` | [Vote](#cosmos.gov.v1beta1.Vote) | repeated | votes defines all the votes present at genesis. |
| `proposals` | [Proposal](#cosmos.gov.v1beta1.Proposal) | repeated | proposals defines all the proposals present at genesis. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1beta1.DepositParams) |  | params defines all the paramaters of related to deposit. |
| `voting_params` | [VotingParams](#cosmos.gov.v1beta1.VotingParams) |  | params defines all the paramaters of related to voting. |
| `tally_params` | [TallyParams](#cosmos.gov.v1beta1.TallyParams) |  | params defines all the paramaters of related to tally. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/query.proto



<a name="cosmos.gov.v1beta1.QueryDepositRequest"></a>

### QueryDepositRequest
QueryDepositRequest is the request type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |






<a name="cosmos.gov.v1beta1.QueryDepositResponse"></a>

### QueryDepositResponse
QueryDepositResponse is the response type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit` | [Deposit](#cosmos.gov.v1beta1.Deposit) |  | deposit defines the requested deposit. |






<a name="cosmos.gov.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#cosmos.gov.v1beta1.Deposit) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params_type` | [string](#string) |  | params_type defines which parameters to query for, can be one of "voting", "tallying" or "deposit". |






<a name="cosmos.gov.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_params` | [VotingParams](#cosmos.gov.v1beta1.VotingParams) |  | voting_params defines the parameters related to voting. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1beta1.DepositParams) |  | deposit_params defines the parameters related to deposit. |
| `tally_params` | [TallyParams](#cosmos.gov.v1beta1.TallyParams) |  | tally_params defines the parameters related to tally. |






<a name="cosmos.gov.v1beta1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest is the request type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1beta1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse is the response type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal` | [Proposal](#cosmos.gov.v1beta1.Proposal) |  |  |






<a name="cosmos.gov.v1beta1.QueryProposalsRequest"></a>

### QueryProposalsRequest
QueryProposalsRequest is the request type for the Query/Proposals RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_status` | [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus) |  | proposal_status defines the status of the proposals. |
| `voter` | [string](#string) |  | voter defines the voter address for the proposals. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryProposalsResponse"></a>

### QueryProposalsResponse
QueryProposalsResponse is the response type for the Query/Proposals RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [Proposal](#cosmos.gov.v1beta1.Proposal) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1beta1.QueryTallyResultRequest"></a>

### QueryTallyResultRequest
QueryTallyResultRequest is the request type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1beta1.QueryTallyResultResponse"></a>

### QueryTallyResultResponse
QueryTallyResultResponse is the response type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tally` | [TallyResult](#cosmos.gov.v1beta1.TallyResult) |  | tally defines the requested tally. |






<a name="cosmos.gov.v1beta1.QueryVoteRequest"></a>

### QueryVoteRequest
QueryVoteRequest is the request type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `voter` | [string](#string) |  | voter defines the oter address for the proposals. |






<a name="cosmos.gov.v1beta1.QueryVoteResponse"></a>

### QueryVoteResponse
QueryVoteResponse is the response type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote` | [Vote](#cosmos.gov.v1beta1.Vote) |  | vote defined the queried vote. |






<a name="cosmos.gov.v1beta1.QueryVotesRequest"></a>

### QueryVotesRequest
QueryVotesRequest is the request type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryVotesResponse"></a>

### QueryVotesResponse
QueryVotesResponse is the response type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [Vote](#cosmos.gov.v1beta1.Vote) | repeated | votes defined the queried votes. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for gov module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Proposal` | [QueryProposalRequest](#cosmos.gov.v1beta1.QueryProposalRequest) | [QueryProposalResponse](#cosmos.gov.v1beta1.QueryProposalResponse) | Proposal queries proposal details based on ProposalID. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}|
| `Proposals` | [QueryProposalsRequest](#cosmos.gov.v1beta1.QueryProposalsRequest) | [QueryProposalsResponse](#cosmos.gov.v1beta1.QueryProposalsResponse) | Proposals queries all proposals based on given status. | GET|/cosmos/gov/v1beta1/proposals|
| `Vote` | [QueryVoteRequest](#cosmos.gov.v1beta1.QueryVoteRequest) | [QueryVoteResponse](#cosmos.gov.v1beta1.QueryVoteResponse) | Vote queries voted information based on proposalID, voterAddr. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}|
| `Votes` | [QueryVotesRequest](#cosmos.gov.v1beta1.QueryVotesRequest) | [QueryVotesResponse](#cosmos.gov.v1beta1.QueryVotesResponse) | Votes queries votes of a given proposal. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/votes|
| `Params` | [QueryParamsRequest](#cosmos.gov.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.gov.v1beta1.QueryParamsResponse) | Params queries all parameters of the gov module. | GET|/cosmos/gov/v1beta1/params/{params_type}|
| `Deposit` | [QueryDepositRequest](#cosmos.gov.v1beta1.QueryDepositRequest) | [QueryDepositResponse](#cosmos.gov.v1beta1.QueryDepositResponse) | Deposit queries single deposit information based proposalID, depositAddr. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}|
| `Deposits` | [QueryDepositsRequest](#cosmos.gov.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#cosmos.gov.v1beta1.QueryDepositsResponse) | Deposits queries all deposits of a single proposal. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits|
| `TallyResult` | [QueryTallyResultRequest](#cosmos.gov.v1beta1.QueryTallyResultRequest) | [QueryTallyResultResponse](#cosmos.gov.v1beta1.QueryTallyResultResponse) | TallyResult queries the tally of a proposal vote. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/tally|

 <!-- end services -->



<a name="cosmos/gov/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/tx.proto



<a name="cosmos.gov.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to submit a deposit to an existing proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="cosmos.gov.v1beta1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal defines an sdk.Msg type that supports submitting arbitrary
proposal Content.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `initial_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `proposer` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse defines the Msg/SubmitProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="cosmos.gov.v1beta1.MsgVote"></a>

### MsgVote
MsgVote defines a message to cast a vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  |  |






<a name="cosmos.gov.v1beta1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse defines the Msg/Vote response type.






<a name="cosmos.gov.v1beta1.MsgVoteWeighted"></a>

### MsgVoteWeighted
MsgVoteWeighted defines a message to cast a vote.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `options` | [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption) | repeated |  |






<a name="cosmos.gov.v1beta1.MsgVoteWeightedResponse"></a>

### MsgVoteWeightedResponse
MsgVoteWeightedResponse defines the Msg/VoteWeighted response type.

Since: cosmos-sdk 0.43





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitProposal` | [MsgSubmitProposal](#cosmos.gov.v1beta1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#cosmos.gov.v1beta1.MsgSubmitProposalResponse) | SubmitProposal defines a method to create new proposal given a content. | |
| `Vote` | [MsgVote](#cosmos.gov.v1beta1.MsgVote) | [MsgVoteResponse](#cosmos.gov.v1beta1.MsgVoteResponse) | Vote defines a method to add a vote on a specific proposal. | |
| `VoteWeighted` | [MsgVoteWeighted](#cosmos.gov.v1beta1.MsgVoteWeighted) | [MsgVoteWeightedResponse](#cosmos.gov.v1beta1.MsgVoteWeightedResponse) | VoteWeighted defines a method to add a weighted vote on a specific proposal.

Since: cosmos-sdk 0.43 | |
| `Deposit` | [MsgDeposit](#cosmos.gov.v1beta1.MsgDeposit) | [MsgDepositResponse](#cosmos.gov.v1beta1.MsgDepositResponse) | Deposit defines a method to add deposit on a specific proposal. | |

 <!-- end services -->



<a name="cosmos/mint/v1beta1/mint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/mint.proto



<a name="cosmos.mint.v1beta1.Minter"></a>

### Minter
Minter represents the minting state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [string](#string) |  | current annual inflation rate |
| `annual_provisions` | [string](#string) |  | current annual expected provisions |






<a name="cosmos.mint.v1beta1.Params"></a>

### Params
Params holds parameters for the mint module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  | type of coin to mint |
| `inflation_rate_change` | [string](#string) |  | maximum annual change in inflation rate |
| `inflation_max` | [string](#string) |  | maximum inflation rate |
| `inflation_min` | [string](#string) |  | minimum inflation rate |
| `goal_bonded` | [string](#string) |  | goal of percent bonded atoms |
| `blocks_per_year` | [uint64](#uint64) |  | expected blocks per year |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/mint/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/genesis.proto



<a name="cosmos.mint.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the mint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minter` | [Minter](#cosmos.mint.v1beta1.Minter) |  | minter is a space for holding current inflation information. |
| `params` | [Params](#cosmos.mint.v1beta1.Params) |  | params defines all the paramaters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/mint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/query.proto



<a name="cosmos.mint.v1beta1.QueryAnnualProvisionsRequest"></a>

### QueryAnnualProvisionsRequest
QueryAnnualProvisionsRequest is the request type for the
Query/AnnualProvisions RPC method.






<a name="cosmos.mint.v1beta1.QueryAnnualProvisionsResponse"></a>

### QueryAnnualProvisionsResponse
QueryAnnualProvisionsResponse is the response type for the
Query/AnnualProvisions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [bytes](#bytes) |  | annual_provisions is the current minting annual provisions value. |






<a name="cosmos.mint.v1beta1.QueryInflationRequest"></a>

### QueryInflationRequest
QueryInflationRequest is the request type for the Query/Inflation RPC method.






<a name="cosmos.mint.v1beta1.QueryInflationResponse"></a>

### QueryInflationResponse
QueryInflationResponse is the response type for the Query/Inflation RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [bytes](#bytes) |  | inflation is the current minting inflation value. |






<a name="cosmos.mint.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="cosmos.mint.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.mint.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.mint.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.mint.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.mint.v1beta1.QueryParamsResponse) | Params returns the total set of minting parameters. | GET|/cosmos/mint/v1beta1/params|
| `Inflation` | [QueryInflationRequest](#cosmos.mint.v1beta1.QueryInflationRequest) | [QueryInflationResponse](#cosmos.mint.v1beta1.QueryInflationResponse) | Inflation returns the current minting inflation value. | GET|/cosmos/mint/v1beta1/inflation|
| `AnnualProvisions` | [QueryAnnualProvisionsRequest](#cosmos.mint.v1beta1.QueryAnnualProvisionsRequest) | [QueryAnnualProvisionsResponse](#cosmos.mint.v1beta1.QueryAnnualProvisionsResponse) | AnnualProvisions current minting annual provisions value. | GET|/cosmos/mint/v1beta1/annual_provisions|

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/slashing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/slashing.proto



<a name="cosmos.slashing.v1beta1.Params"></a>

### Params
Params represents the parameters used for by the slashing module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signed_blocks_window` | [int64](#int64) |  |  |
| `min_signed_per_window` | [bytes](#bytes) |  |  |
| `downtime_jail_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `slash_fraction_double_sign` | [bytes](#bytes) |  |  |
| `slash_fraction_downtime` | [bytes](#bytes) |  |  |






<a name="cosmos.slashing.v1beta1.ValidatorSigningInfo"></a>

### ValidatorSigningInfo
ValidatorSigningInfo defines a validator's signing info for monitoring their
liveness activity.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `start_height` | [int64](#int64) |  | Height at which validator was first a candidate OR was unjailed |
| `index_offset` | [int64](#int64) |  | Index which is incremented each time the validator was a bonded in a block and may have signed a precommit or not. This in conjunction with the `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`. |
| `jailed_until` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp until which the validator is jailed due to liveness downtime. |
| `tombstoned` | [bool](#bool) |  | Whether or not a validator has been tombstoned (killed out of validator set). It is set once the validator commits an equivocation or for any other configured misbehiavor. |
| `missed_blocks_counter` | [int64](#int64) |  | A counter kept to avoid unnecessary array reads. Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/genesis.proto



<a name="cosmos.slashing.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the slashing module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.slashing.v1beta1.Params) |  | params defines all the paramaters of related to deposit. |
| `signing_infos` | [SigningInfo](#cosmos.slashing.v1beta1.SigningInfo) | repeated | signing_infos represents a map between validator addresses and their signing infos. |
| `missed_blocks` | [ValidatorMissedBlocks](#cosmos.slashing.v1beta1.ValidatorMissedBlocks) | repeated | missed_blocks represents a map between validator addresses and their missed blocks. |






<a name="cosmos.slashing.v1beta1.MissedBlock"></a>

### MissedBlock
MissedBlock contains height and missed status as boolean.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [int64](#int64) |  | index is the height at which the block was missed. |
| `missed` | [bool](#bool) |  | missed is the missed status. |






<a name="cosmos.slashing.v1beta1.SigningInfo"></a>

### SigningInfo
SigningInfo stores validator signing info of corresponding address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the validator address. |
| `validator_signing_info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) |  | validator_signing_info represents the signing info of this validator. |






<a name="cosmos.slashing.v1beta1.ValidatorMissedBlocks"></a>

### ValidatorMissedBlocks
ValidatorMissedBlocks contains array of missed blocks of corresponding
address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the validator address. |
| `missed_blocks` | [MissedBlock](#cosmos.slashing.v1beta1.MissedBlock) | repeated | missed_blocks is an array of missed blocks by the validator. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/query.proto



<a name="cosmos.slashing.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method






<a name="cosmos.slashing.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.slashing.v1beta1.Params) |  |  |






<a name="cosmos.slashing.v1beta1.QuerySigningInfoRequest"></a>

### QuerySigningInfoRequest
QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cons_address` | [string](#string) |  | cons_address is the address to query signing info of |






<a name="cosmos.slashing.v1beta1.QuerySigningInfoResponse"></a>

### QuerySigningInfoResponse
QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `val_signing_info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) |  | val_signing_info is the signing info of requested val cons address |






<a name="cosmos.slashing.v1beta1.QuerySigningInfosRequest"></a>

### QuerySigningInfosRequest
QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="cosmos.slashing.v1beta1.QuerySigningInfosResponse"></a>

### QuerySigningInfosResponse
QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) | repeated | info is the signing info of all validators |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.slashing.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.slashing.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.slashing.v1beta1.QueryParamsResponse) | Params queries the parameters of slashing module | GET|/cosmos/slashing/v1beta1/params|
| `SigningInfo` | [QuerySigningInfoRequest](#cosmos.slashing.v1beta1.QuerySigningInfoRequest) | [QuerySigningInfoResponse](#cosmos.slashing.v1beta1.QuerySigningInfoResponse) | SigningInfo queries the signing info of given cons address | GET|/cosmos/slashing/v1beta1/signing_infos/{cons_address}|
| `SigningInfos` | [QuerySigningInfosRequest](#cosmos.slashing.v1beta1.QuerySigningInfosRequest) | [QuerySigningInfosResponse](#cosmos.slashing.v1beta1.QuerySigningInfosResponse) | SigningInfos queries signing info of all validators | GET|/cosmos/slashing/v1beta1/signing_infos|

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/tx.proto



<a name="cosmos.slashing.v1beta1.MsgUnjail"></a>

### MsgUnjail
MsgUnjail defines the Msg/Unjail request type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  |  |






<a name="cosmos.slashing.v1beta1.MsgUnjailResponse"></a>

### MsgUnjailResponse
MsgUnjailResponse defines the Msg/Unjail response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.slashing.v1beta1.Msg"></a>

### Msg
Msg defines the slashing Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Unjail` | [MsgUnjail](#cosmos.slashing.v1beta1.MsgUnjail) | [MsgUnjailResponse](#cosmos.slashing.v1beta1.MsgUnjailResponse) | Unjail defines a method for unjailing a jailed validator, thus returning them into the bonded validator set, so they can begin receiving provisions and rewards again. | |

 <!-- end services -->



<a name="cosmos/staking/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/authz.proto



<a name="cosmos.staking.v1beta1.StakeAuthorization"></a>

### StakeAuthorization
StakeAuthorization defines authorization for delegate/undelegate/redelegate.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | max_tokens specifies the maximum amount of tokens can be delegate to a validator. If it is empty, there is no spend limit and any amount of coins can be delegated. |
| `allow_list` | [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators) |  | allow_list specifies list of validator addresses to whom grantee can delegate tokens on behalf of granter's account. |
| `deny_list` | [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators) |  | deny_list specifies list of validator addresses to whom grantee can not delegate tokens. |
| `authorization_type` | [AuthorizationType](#cosmos.staking.v1beta1.AuthorizationType) |  | authorization_type defines one of AuthorizationType. |






<a name="cosmos.staking.v1beta1.StakeAuthorization.Validators"></a>

### StakeAuthorization.Validators
Validators defines list of validator addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) | repeated |  |





 <!-- end messages -->


<a name="cosmos.staking.v1beta1.AuthorizationType"></a>

### AuthorizationType
AuthorizationType defines the type of staking module authorization type

Since: cosmos-sdk 0.43

| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHORIZATION_TYPE_UNSPECIFIED | 0 | AUTHORIZATION_TYPE_UNSPECIFIED specifies an unknown authorization type |
| AUTHORIZATION_TYPE_DELEGATE | 1 | AUTHORIZATION_TYPE_DELEGATE defines an authorization type for Msg/Delegate |
| AUTHORIZATION_TYPE_UNDELEGATE | 2 | AUTHORIZATION_TYPE_UNDELEGATE defines an authorization type for Msg/Undelegate |
| AUTHORIZATION_TYPE_REDELEGATE | 3 | AUTHORIZATION_TYPE_REDELEGATE defines an authorization type for Msg/BeginRedelegate |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/staking.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/staking.proto



<a name="cosmos.staking.v1beta1.Commission"></a>

### Commission
Commission defines commission parameters for a given validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commission_rates` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  | commission_rates defines the initial commission rates to be used for creating a validator. |
| `update_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | update_time is the last time the commission rate was changed. |






<a name="cosmos.staking.v1beta1.CommissionRates"></a>

### CommissionRates
CommissionRates defines the initial commission rates to be used for creating
a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rate` | [string](#string) |  | rate is the commission rate charged to delegators, as a fraction. |
| `max_rate` | [string](#string) |  | max_rate defines the maximum commission rate which validator can ever charge, as a fraction. |
| `max_change_rate` | [string](#string) |  | max_change_rate defines the maximum daily increase of the validator commission, as a fraction. |






<a name="cosmos.staking.v1beta1.DVPair"></a>

### DVPair
DVPair is struct that just has a delegator-validator pair with no other data.
It is intended to be used as a marshalable pointer. For example, a DVPair can
be used to construct the key to getting an UnbondingDelegation from state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.DVPairs"></a>

### DVPairs
DVPairs defines an array of DVPair objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pairs` | [DVPair](#cosmos.staking.v1beta1.DVPair) | repeated |  |






<a name="cosmos.staking.v1beta1.DVVTriplet"></a>

### DVVTriplet
DVVTriplet is struct that just has a delegator-validator-validator triplet
with no other data. It is intended to be used as a marshalable pointer. For
example, a DVVTriplet can be used to construct the key to getting a
Redelegation from state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.DVVTriplets"></a>

### DVVTriplets
DVVTriplets defines an array of DVVTriplet objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `triplets` | [DVVTriplet](#cosmos.staking.v1beta1.DVVTriplet) | repeated |  |






<a name="cosmos.staking.v1beta1.Delegation"></a>

### Delegation
Delegation represents the bond with tokens held by an account. It is
owned by one delegator, and is associated with the voting power of one
validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_address` | [string](#string) |  | validator_address is the bech32-encoded address of the validator. |
| `shares` | [string](#string) |  | shares define the delegation shares received. |






<a name="cosmos.staking.v1beta1.DelegationResponse"></a>

### DelegationResponse
DelegationResponse is equivalent to Delegation except that it contains a
balance in addition to shares which is more suitable for client responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation` | [Delegation](#cosmos.staking.v1beta1.Delegation) |  |  |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.Description"></a>

### Description
Description defines a validator description.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `moniker` | [string](#string) |  | moniker defines a human-readable name for the validator. |
| `identity` | [string](#string) |  | identity defines an optional identity signature (ex. UPort or Keybase). |
| `website` | [string](#string) |  | website defines an optional website link. |
| `security_contact` | [string](#string) |  | security_contact defines an optional email for security contact. |
| `details` | [string](#string) |  | details define other optional details. |






<a name="cosmos.staking.v1beta1.HistoricalInfo"></a>

### HistoricalInfo
HistoricalInfo contains header and validator information for a given block.
It is stored as part of staking module's state, which persists the `n` most
recent HistoricalInfo
(`n` is set by the staking module's `historical_entries` parameter).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header` | [tendermint.types.Header](#tendermint.types.Header) |  |  |
| `valset` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated |  |






<a name="cosmos.staking.v1beta1.Params"></a>

### Params
Params defines the parameters for the staking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_time` | [google.protobuf.Duration](#google.protobuf.Duration) |  | unbonding_time is the time duration of unbonding. |
| `max_validators` | [uint32](#uint32) |  | max_validators is the maximum number of validators. |
| `max_entries` | [uint32](#uint32) |  | max_entries is the max entries for either unbonding delegation or redelegation (per pair/trio). |
| `historical_entries` | [uint32](#uint32) |  | historical_entries is the number of historical entries to persist. |
| `bond_denom` | [string](#string) |  | bond_denom defines the bondable coin denomination. |






<a name="cosmos.staking.v1beta1.Pool"></a>

### Pool
Pool is used for tracking bonded and not-bonded token supply of the bond
denomination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `not_bonded_tokens` | [string](#string) |  |  |
| `bonded_tokens` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.Redelegation"></a>

### Redelegation
Redelegation contains the list of a particular delegator's redelegating bonds
from a particular source validator to a particular destination validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_src_address` | [string](#string) |  | validator_src_address is the validator redelegation source operator address. |
| `validator_dst_address` | [string](#string) |  | validator_dst_address is the validator redelegation destination operator address. |
| `entries` | [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry) | repeated | entries are the redelegation entries.

redelegation entries |






<a name="cosmos.staking.v1beta1.RedelegationEntry"></a>

### RedelegationEntry
RedelegationEntry defines a redelegation object with relevant metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creation_height` | [int64](#int64) |  | creation_height defines the height which the redelegation took place. |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | completion_time defines the unix time for redelegation completion. |
| `initial_balance` | [string](#string) |  | initial_balance defines the initial balance when redelegation started. |
| `shares_dst` | [string](#string) |  | shares_dst is the amount of destination-validator shares created by redelegation. |






<a name="cosmos.staking.v1beta1.RedelegationEntryResponse"></a>

### RedelegationEntryResponse
RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
contains a balance in addition to shares which is more suitable for client
responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation_entry` | [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry) |  |  |
| `balance` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.RedelegationResponse"></a>

### RedelegationResponse
RedelegationResponse is equivalent to a Redelegation except that its entries
contain a balance in addition to shares which is more suitable for client
responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation` | [Redelegation](#cosmos.staking.v1beta1.Redelegation) |  |  |
| `entries` | [RedelegationEntryResponse](#cosmos.staking.v1beta1.RedelegationEntryResponse) | repeated |  |






<a name="cosmos.staking.v1beta1.UnbondingDelegation"></a>

### UnbondingDelegation
UnbondingDelegation stores all of a single delegator's unbonding bonds
for a single validator in an time-ordered list.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_address` | [string](#string) |  | validator_address is the bech32-encoded address of the validator. |
| `entries` | [UnbondingDelegationEntry](#cosmos.staking.v1beta1.UnbondingDelegationEntry) | repeated | entries are the unbonding delegation entries.

unbonding delegation entries |






<a name="cosmos.staking.v1beta1.UnbondingDelegationEntry"></a>

### UnbondingDelegationEntry
UnbondingDelegationEntry defines an unbonding object with relevant metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creation_height` | [int64](#int64) |  | creation_height is the height which the unbonding took place. |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | completion_time is the unix time for unbonding completion. |
| `initial_balance` | [string](#string) |  | initial_balance defines the tokens initially scheduled to receive at completion. |
| `balance` | [string](#string) |  | balance defines the tokens to receive at completion. |






<a name="cosmos.staking.v1beta1.ValAddresses"></a>

### ValAddresses
ValAddresses defines a repeated set of validator addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |






<a name="cosmos.staking.v1beta1.Validator"></a>

### Validator
Validator defines a validator, together with the total amount of the
Validator's bond shares and their exchange rate to coins. Slashing results in
a decrease in the exchange rate, allowing correct calculation of future
undelegations without iterating over delegators. When coins are delegated to
this validator, the validator is credited with a delegation whose number of
bond shares is based on the amount of coins delegated divided by the current
exchange rate. Voting power can be calculated as total bonded shares
multiplied by exchange rate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  | operator_address defines the address of the validator's operator; bech encoded in JSON. |
| `consensus_pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  | consensus_pubkey is the consensus public key of the validator, as a Protobuf Any. |
| `jailed` | [bool](#bool) |  | jailed defined whether the validator has been jailed from bonded status or not. |
| `status` | [BondStatus](#cosmos.staking.v1beta1.BondStatus) |  | status is the validator status (bonded/unbonding/unbonded). |
| `tokens` | [string](#string) |  | tokens define the delegated tokens (incl. self-delegation). |
| `delegator_shares` | [string](#string) |  | delegator_shares defines total shares issued to a validator's delegators. |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  | description defines the description terms for the validator. |
| `unbonding_height` | [int64](#int64) |  | unbonding_height defines, if unbonding, the height at which this validator has begun unbonding. |
| `unbonding_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | unbonding_time defines, if unbonding, the min time for the validator to complete unbonding. |
| `commission` | [Commission](#cosmos.staking.v1beta1.Commission) |  | commission defines the commission parameters. |
| `min_self_delegation` | [string](#string) |  | min_self_delegation is the validator's self declared minimum self delegation. |





 <!-- end messages -->


<a name="cosmos.staking.v1beta1.BondStatus"></a>

### BondStatus
BondStatus is the status of a validator.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BOND_STATUS_UNSPECIFIED | 0 | UNSPECIFIED defines an invalid validator status. |
| BOND_STATUS_UNBONDED | 1 | UNBONDED defines a validator that is not bonded. |
| BOND_STATUS_UNBONDING | 2 | UNBONDING defines a validator that is unbonding. |
| BOND_STATUS_BONDED | 3 | BONDED defines a validator that is bonded. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/genesis.proto



<a name="cosmos.staking.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the staking module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.staking.v1beta1.Params) |  | params defines all the paramaters of related to deposit. |
| `last_total_power` | [bytes](#bytes) |  | last_total_power tracks the total amounts of bonded tokens recorded during the previous end block. |
| `last_validator_powers` | [LastValidatorPower](#cosmos.staking.v1beta1.LastValidatorPower) | repeated | last_validator_powers is a special index that provides a historical list of the last-block's bonded validators. |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | delegations defines the validator set at genesis. |
| `delegations` | [Delegation](#cosmos.staking.v1beta1.Delegation) | repeated | delegations defines the delegations active at genesis. |
| `unbonding_delegations` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated | unbonding_delegations defines the unbonding delegations active at genesis. |
| `redelegations` | [Redelegation](#cosmos.staking.v1beta1.Redelegation) | repeated | redelegations defines the redelegations active at genesis. |
| `exported` | [bool](#bool) |  |  |






<a name="cosmos.staking.v1beta1.LastValidatorPower"></a>

### LastValidatorPower
LastValidatorPower required for validator set update logic.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the validator. |
| `power` | [int64](#int64) |  | power defines the power of the validator. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/query.proto



<a name="cosmos.staking.v1beta1.QueryDelegationRequest"></a>

### QueryDelegationRequest
QueryDelegationRequest is request type for the Query/Delegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryDelegationResponse"></a>

### QueryDelegationResponse
QueryDelegationResponse is response type for the Query/Delegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_response` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) |  | delegation_responses defines the delegation info of a delegation. |






<a name="cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"></a>

### QueryDelegatorDelegationsRequest
QueryDelegatorDelegationsRequest is request type for the
Query/DelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"></a>

### QueryDelegatorDelegationsResponse
QueryDelegatorDelegationsResponse is response type for the
Query/DelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_responses` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) | repeated | delegation_responses defines all the delegations' info of a delegator. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"></a>

### QueryDelegatorUnbondingDelegationsRequest
QueryDelegatorUnbondingDelegationsRequest is request type for the
Query/DelegatorUnbondingDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"></a>

### QueryDelegatorUnbondingDelegationsResponse
QueryUnbondingDelegatorDelegationsResponse is response type for the
Query/UnbondingDelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_responses` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"></a>

### QueryDelegatorValidatorRequest
QueryDelegatorValidatorRequest is request type for the
Query/DelegatorValidator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"></a>

### QueryDelegatorValidatorResponse
QueryDelegatorValidatorResponse response type for the
Query/DelegatorValidator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [Validator](#cosmos.staking.v1beta1.Validator) |  | validator defines the the validator info. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"></a>

### QueryDelegatorValidatorsRequest
QueryDelegatorValidatorsRequest is request type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"></a>

### QueryDelegatorValidatorsResponse
QueryDelegatorValidatorsResponse is response type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | validators defines the the validators' info of a delegator. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryHistoricalInfoRequest"></a>

### QueryHistoricalInfoRequest
QueryHistoricalInfoRequest is request type for the Query/HistoricalInfo RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height defines at which height to query the historical info. |






<a name="cosmos.staking.v1beta1.QueryHistoricalInfoResponse"></a>

### QueryHistoricalInfoResponse
QueryHistoricalInfoResponse is response type for the Query/HistoricalInfo RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hist` | [HistoricalInfo](#cosmos.staking.v1beta1.HistoricalInfo) |  | hist defines the historical info at the given height. |






<a name="cosmos.staking.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="cosmos.staking.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.staking.v1beta1.Params) |  | params holds all the parameters of this module. |






<a name="cosmos.staking.v1beta1.QueryPoolRequest"></a>

### QueryPoolRequest
QueryPoolRequest is request type for the Query/Pool RPC method.






<a name="cosmos.staking.v1beta1.QueryPoolResponse"></a>

### QueryPoolResponse
QueryPoolResponse is response type for the Query/Pool RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [Pool](#cosmos.staking.v1beta1.Pool) |  | pool defines the pool info. |






<a name="cosmos.staking.v1beta1.QueryRedelegationsRequest"></a>

### QueryRedelegationsRequest
QueryRedelegationsRequest is request type for the Query/Redelegations RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `src_validator_addr` | [string](#string) |  | src_validator_addr defines the validator address to redelegate from. |
| `dst_validator_addr` | [string](#string) |  | dst_validator_addr defines the validator address to redelegate to. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryRedelegationsResponse"></a>

### QueryRedelegationsResponse
QueryRedelegationsResponse is response type for the Query/Redelegations RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation_responses` | [RedelegationResponse](#cosmos.staking.v1beta1.RedelegationResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"></a>

### QueryUnbondingDelegationRequest
QueryUnbondingDelegationRequest is request type for the
Query/UnbondingDelegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"></a>

### QueryUnbondingDelegationResponse
QueryDelegationResponse is response type for the Query/UnbondingDelegation
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbond` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) |  | unbond defines the unbonding information of a delegation. |






<a name="cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"></a>

### QueryValidatorDelegationsRequest
QueryValidatorDelegationsRequest is request type for the
Query/ValidatorDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"></a>

### QueryValidatorDelegationsResponse
QueryValidatorDelegationsResponse is response type for the
Query/ValidatorDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_responses` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryValidatorRequest"></a>

### QueryValidatorRequest
QueryValidatorRequest is response type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryValidatorResponse"></a>

### QueryValidatorResponse
QueryValidatorResponse is response type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [Validator](#cosmos.staking.v1beta1.Validator) |  | validator defines the the validator info. |






<a name="cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"></a>

### QueryValidatorUnbondingDelegationsRequest
QueryValidatorUnbondingDelegationsRequest is required type for the
Query/ValidatorUnbondingDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"></a>

### QueryValidatorUnbondingDelegationsResponse
QueryValidatorUnbondingDelegationsResponse is response type for the
Query/ValidatorUnbondingDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_responses` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryValidatorsRequest"></a>

### QueryValidatorsRequest
QueryValidatorsRequest is request type for Query/Validators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [string](#string) |  | status enables to query for validators matching a given status. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorsResponse"></a>

### QueryValidatorsResponse
QueryValidatorsResponse is response type for the Query/Validators RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | validators contains all the queried validators. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.staking.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Validators` | [QueryValidatorsRequest](#cosmos.staking.v1beta1.QueryValidatorsRequest) | [QueryValidatorsResponse](#cosmos.staking.v1beta1.QueryValidatorsResponse) | Validators queries all validators that match the given status. | GET|/cosmos/staking/v1beta1/validators|
| `Validator` | [QueryValidatorRequest](#cosmos.staking.v1beta1.QueryValidatorRequest) | [QueryValidatorResponse](#cosmos.staking.v1beta1.QueryValidatorResponse) | Validator queries validator info for given validator address. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}|
| `ValidatorDelegations` | [QueryValidatorDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorDelegationsRequest) | [QueryValidatorDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorDelegationsResponse) | ValidatorDelegations queries delegate info for given validator. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations|
| `ValidatorUnbondingDelegations` | [QueryValidatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest) | [QueryValidatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse) | ValidatorUnbondingDelegations queries unbonding delegations of a validator. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations|
| `Delegation` | [QueryDelegationRequest](#cosmos.staking.v1beta1.QueryDelegationRequest) | [QueryDelegationResponse](#cosmos.staking.v1beta1.QueryDelegationResponse) | Delegation queries delegate info for given validator delegator pair. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}|
| `UnbondingDelegation` | [QueryUnbondingDelegationRequest](#cosmos.staking.v1beta1.QueryUnbondingDelegationRequest) | [QueryUnbondingDelegationResponse](#cosmos.staking.v1beta1.QueryUnbondingDelegationResponse) | UnbondingDelegation queries unbonding info for given validator delegator pair. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation|
| `DelegatorDelegations` | [QueryDelegatorDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest) | [QueryDelegatorDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse) | DelegatorDelegations queries all delegations of a given delegator address. | GET|/cosmos/staking/v1beta1/delegations/{delegator_addr}|
| `DelegatorUnbondingDelegations` | [QueryDelegatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest) | [QueryDelegatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse) | DelegatorUnbondingDelegations queries all unbonding delegations of a given delegator address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations|
| `Redelegations` | [QueryRedelegationsRequest](#cosmos.staking.v1beta1.QueryRedelegationsRequest) | [QueryRedelegationsResponse](#cosmos.staking.v1beta1.QueryRedelegationsResponse) | Redelegations queries redelegations of given address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations|
| `DelegatorValidators` | [QueryDelegatorValidatorsRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest) | [QueryDelegatorValidatorsResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse) | DelegatorValidators queries all validators info for given delegator address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators|
| `DelegatorValidator` | [QueryDelegatorValidatorRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorRequest) | [QueryDelegatorValidatorResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorResponse) | DelegatorValidator queries validator info for given delegator validator pair. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}|
| `HistoricalInfo` | [QueryHistoricalInfoRequest](#cosmos.staking.v1beta1.QueryHistoricalInfoRequest) | [QueryHistoricalInfoResponse](#cosmos.staking.v1beta1.QueryHistoricalInfoResponse) | HistoricalInfo queries the historical info for given height. | GET|/cosmos/staking/v1beta1/historical_info/{height}|
| `Pool` | [QueryPoolRequest](#cosmos.staking.v1beta1.QueryPoolRequest) | [QueryPoolResponse](#cosmos.staking.v1beta1.QueryPoolResponse) | Pool queries the pool info. | GET|/cosmos/staking/v1beta1/pool|
| `Params` | [QueryParamsRequest](#cosmos.staking.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.staking.v1beta1.QueryParamsResponse) | Parameters queries the staking parameters. | GET|/cosmos/staking/v1beta1/params|

 <!-- end services -->



<a name="cosmos/staking/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/tx.proto



<a name="cosmos.staking.v1beta1.MsgBeginRedelegate"></a>

### MsgBeginRedelegate
MsgBeginRedelegate defines a SDK message for performing a redelegation
of coins from a delegator and source validator to a destination validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgBeginRedelegateResponse"></a>

### MsgBeginRedelegateResponse
MsgBeginRedelegateResponse defines the Msg/BeginRedelegate response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="cosmos.staking.v1beta1.MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator defines a SDK message for creating a new validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `commission` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  |  |
| `min_self_delegation` | [string](#string) |  |  |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgCreateValidatorResponse"></a>

### MsgCreateValidatorResponse
MsgCreateValidatorResponse defines the Msg/CreateValidator response type.






<a name="cosmos.staking.v1beta1.MsgDelegate"></a>

### MsgDelegate
MsgDelegate defines a SDK message for performing a delegation of coins
from a delegator to a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgDelegateResponse"></a>

### MsgDelegateResponse
MsgDelegateResponse defines the Msg/Delegate response type.






<a name="cosmos.staking.v1beta1.MsgEditValidator"></a>

### MsgEditValidator
MsgEditValidator defines a SDK message for editing an existing validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `validator_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  | We pass a reference to the new commission rate and min self delegation as it's not mandatory to update. If not updated, the deserialized rate will be zero with no way to distinguish if an update was intended. REF: #2373 |
| `min_self_delegation` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.MsgEditValidatorResponse"></a>

### MsgEditValidatorResponse
MsgEditValidatorResponse defines the Msg/EditValidator response type.






<a name="cosmos.staking.v1beta1.MsgUndelegate"></a>

### MsgUndelegate
MsgUndelegate defines a SDK message for performing an undelegation from a
delegate and a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgUndelegateResponse"></a>

### MsgUndelegateResponse
MsgUndelegateResponse defines the Msg/Undelegate response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.staking.v1beta1.Msg"></a>

### Msg
Msg defines the staking Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateValidator` | [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator) | [MsgCreateValidatorResponse](#cosmos.staking.v1beta1.MsgCreateValidatorResponse) | CreateValidator defines a method for creating a new validator. | |
| `EditValidator` | [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator) | [MsgEditValidatorResponse](#cosmos.staking.v1beta1.MsgEditValidatorResponse) | EditValidator defines a method for editing an existing validator. | |
| `Delegate` | [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate) | [MsgDelegateResponse](#cosmos.staking.v1beta1.MsgDelegateResponse) | Delegate defines a method for performing a delegation of coins from a delegator to a validator. | |
| `BeginRedelegate` | [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate) | [MsgBeginRedelegateResponse](#cosmos.staking.v1beta1.MsgBeginRedelegateResponse) | BeginRedelegate defines a method for performing a redelegation of coins from a delegator and source validator to a destination validator. | |
| `Undelegate` | [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate) | [MsgUndelegateResponse](#cosmos.staking.v1beta1.MsgUndelegateResponse) | Undelegate defines a method for performing an undelegation from a delegate and a validator. | |

 <!-- end services -->



<a name="cosmos/tx/signing/v1beta1/signing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/signing/v1beta1/signing.proto



<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor"></a>

### SignatureDescriptor
SignatureDescriptor is a convenience type which represents the full data for
a signature including the public key of the signer, signing modes and the
signature itself. It is primarily used for coordinating signatures between
clients.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer |
| `data` | [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data) |  |  |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"></a>

### SignatureDescriptor.Data
Data represents signature data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `single` | [SignatureDescriptor.Data.Single](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single) |  | single represents a single signer |
| `multi` | [SignatureDescriptor.Data.Multi](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi) |  | multi represents a multisig signer |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"></a>

### SignatureDescriptor.Data.Multi
Multi is the signature data for a multisig public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitarray` | [cosmos.crypto.multisig.v1beta1.CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray) |  | bitarray specifies which keys within the multisig are signing |
| `signatures` | [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data) | repeated | signatures is the signatures of the multi-signature |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"></a>

### SignatureDescriptor.Data.Single
Single is the signature data for a single signer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mode` | [SignMode](#cosmos.tx.signing.v1beta1.SignMode) |  | mode is the signing mode of the single signer |
| `signature` | [bytes](#bytes) |  | signature is the raw signature bytes |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptors"></a>

### SignatureDescriptors
SignatureDescriptors wraps multiple SignatureDescriptor's.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signatures` | [SignatureDescriptor](#cosmos.tx.signing.v1beta1.SignatureDescriptor) | repeated | signatures are the signature descriptors |





 <!-- end messages -->


<a name="cosmos.tx.signing.v1beta1.SignMode"></a>

### SignMode
SignMode represents a signing mode with its own security guarantees.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGN_MODE_UNSPECIFIED | 0 | SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be rejected |
| SIGN_MODE_DIRECT | 1 | SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is verified with raw bytes from Tx |
| SIGN_MODE_TEXTUAL | 2 | SIGN_MODE_TEXTUAL is a future signing mode that will verify some human-readable textual representation on top of the binary representation from SIGN_MODE_DIRECT |
| SIGN_MODE_LEGACY_AMINO_JSON | 127 | SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses Amino JSON and will be removed in the future |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/tx/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/tx.proto



<a name="cosmos.tx.v1beta1.AuthInfo"></a>

### AuthInfo
AuthInfo describes the fee and signer modes that are used to sign a
transaction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer_infos` | [SignerInfo](#cosmos.tx.v1beta1.SignerInfo) | repeated | signer_infos defines the signing modes for the required signers. The number and order of elements must match the required signers from TxBody's messages. The first element is the primary signer and the one which pays the fee. |
| `fee` | [Fee](#cosmos.tx.v1beta1.Fee) |  | Fee is the fee and gas limit for the transaction. The first signer is the primary signer and the one which pays the fee. The fee can be calculated based on the cost of evaluating the body and doing signature verification of the signers. This can be estimated via simulation. |






<a name="cosmos.tx.v1beta1.Fee"></a>

### Fee
Fee includes the amount of coins paid in fees and the maximum
gas to be used by the transaction. The ratio yields an effective "gasprice",
which must be above some miminum to be accepted into the mempool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount is the amount of coins to be paid as a fee |
| `gas_limit` | [uint64](#uint64) |  | gas_limit is the maximum gas that can be used in transaction processing before an out of gas error occurs |
| `payer` | [string](#string) |  | if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees. the payer must be a tx signer (and thus have signed this field in AuthInfo). setting this field does *not* change the ordering of required signers for the transaction. |
| `granter` | [string](#string) |  | if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does not support fee grants, this will fail |






<a name="cosmos.tx.v1beta1.ModeInfo"></a>

### ModeInfo
ModeInfo describes the signing mode of a single or nested multisig signer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `single` | [ModeInfo.Single](#cosmos.tx.v1beta1.ModeInfo.Single) |  | single represents a single signer |
| `multi` | [ModeInfo.Multi](#cosmos.tx.v1beta1.ModeInfo.Multi) |  | multi represents a nested multisig signer |






<a name="cosmos.tx.v1beta1.ModeInfo.Multi"></a>

### ModeInfo.Multi
Multi is the mode info for a multisig public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitarray` | [cosmos.crypto.multisig.v1beta1.CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray) |  | bitarray specifies which keys within the multisig are signing |
| `mode_infos` | [ModeInfo](#cosmos.tx.v1beta1.ModeInfo) | repeated | mode_infos is the corresponding modes of the signers of the multisig which could include nested multisig public keys |






<a name="cosmos.tx.v1beta1.ModeInfo.Single"></a>

### ModeInfo.Single
Single is the mode info for a single signer. It is structured as a message
to allow for additional fields such as locale for SIGN_MODE_TEXTUAL in the
future


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mode` | [cosmos.tx.signing.v1beta1.SignMode](#cosmos.tx.signing.v1beta1.SignMode) |  | mode is the signing mode of the single signer |






<a name="cosmos.tx.v1beta1.SignDoc"></a>

### SignDoc
SignDoc is the type used for generating sign bytes for SIGN_MODE_DIRECT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body_bytes` | [bytes](#bytes) |  | body_bytes is protobuf serialization of a TxBody that matches the representation in TxRaw. |
| `auth_info_bytes` | [bytes](#bytes) |  | auth_info_bytes is a protobuf serialization of an AuthInfo that matches the representation in TxRaw. |
| `chain_id` | [string](#string) |  | chain_id is the unique identifier of the chain this transaction targets. It prevents signed transactions from being used on another chain by an attacker |
| `account_number` | [uint64](#uint64) |  | account_number is the account number of the account in state |






<a name="cosmos.tx.v1beta1.SignerInfo"></a>

### SignerInfo
SignerInfo describes the public key and signing mode of a single top-level
signer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer. It is optional for accounts that already exist in state. If unset, the verifier can use the required \ signer address for this position and lookup the public key. |
| `mode_info` | [ModeInfo](#cosmos.tx.v1beta1.ModeInfo) |  | mode_info describes the signing mode of the signer and is a nested structure to support nested multisig pubkey's |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |






<a name="cosmos.tx.v1beta1.Tx"></a>

### Tx
Tx is the standard type used for broadcasting transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body` | [TxBody](#cosmos.tx.v1beta1.TxBody) |  | body is the processable content of the transaction |
| `auth_info` | [AuthInfo](#cosmos.tx.v1beta1.AuthInfo) |  | auth_info is the authorization related content of the transaction, specifically signers, signer modes and fee |
| `signatures` | [bytes](#bytes) | repeated | signatures is a list of signatures that matches the length and order of AuthInfo's signer_infos to allow connecting signature meta information like public key and signing mode by position. |






<a name="cosmos.tx.v1beta1.TxBody"></a>

### TxBody
TxBody is the body of a transaction that all signers sign over.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated | messages is a list of messages to be executed. The required signers of those messages define the number and order of elements in AuthInfo's signer_infos and Tx's signatures. Each required signer address is added to the list only the first time it occurs. By convention, the first required signer (usually from the first message) is referred to as the primary signer and pays the fee for the whole transaction. |
| `memo` | [string](#string) |  | memo is any arbitrary note/comment to be added to the transaction. WARNING: in clients, any publicly exposed text should not be called memo, but should be called `note` instead (see https://github.com/cosmos/cosmos-sdk/issues/9122). |
| `timeout_height` | [uint64](#uint64) |  | timeout is the block height after which this transaction will not be processed by the chain |
| `extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, the transaction will be rejected |
| `non_critical_extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, they will be ignored |






<a name="cosmos.tx.v1beta1.TxRaw"></a>

### TxRaw
TxRaw is a variant of Tx that pins the signer's exact binary representation
of body and auth_info. This is used for signing, broadcasting and
verification. The binary `serialize(tx: TxRaw)` is stored in Tendermint and
the hash `sha256(serialize(tx: TxRaw))` becomes the "txhash", commonly used
as the transaction ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body_bytes` | [bytes](#bytes) |  | body_bytes is a protobuf serialization of a TxBody that matches the representation in SignDoc. |
| `auth_info_bytes` | [bytes](#bytes) |  | auth_info_bytes is a protobuf serialization of an AuthInfo that matches the representation in SignDoc. |
| `signatures` | [bytes](#bytes) | repeated | signatures is a list of signatures that matches the length and order of AuthInfo's signer_infos to allow connecting signature meta information like public key and signing mode by position. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/tx/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/service.proto



<a name="cosmos.tx.v1beta1.BroadcastTxRequest"></a>

### BroadcastTxRequest
BroadcastTxRequest is the request type for the Service.BroadcastTxRequest
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_bytes` | [bytes](#bytes) |  | tx_bytes is the raw transaction. |
| `mode` | [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode) |  |  |






<a name="cosmos.tx.v1beta1.BroadcastTxResponse"></a>

### BroadcastTxResponse
BroadcastTxResponse is the response type for the
Service.BroadcastTx method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_response` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | tx_response is the queried TxResponses. |






<a name="cosmos.tx.v1beta1.GetTxRequest"></a>

### GetTxRequest
GetTxRequest is the request type for the Service.GetTx
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [string](#string) |  | hash is the tx hash to query, encoded as a hex string. |






<a name="cosmos.tx.v1beta1.GetTxResponse"></a>

### GetTxResponse
GetTxResponse is the response type for the Service.GetTx method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [Tx](#cosmos.tx.v1beta1.Tx) |  | tx is the queried transaction. |
| `tx_response` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | tx_response is the queried TxResponses. |






<a name="cosmos.tx.v1beta1.GetTxsEventRequest"></a>

### GetTxsEventRequest
GetTxsEventRequest is the request type for the Service.TxsByEvents
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `events` | [string](#string) | repeated | events is the list of transaction event type. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |
| `order_by` | [OrderBy](#cosmos.tx.v1beta1.OrderBy) |  |  |






<a name="cosmos.tx.v1beta1.GetTxsEventResponse"></a>

### GetTxsEventResponse
GetTxsEventResponse is the response type for the Service.TxsByEvents
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [Tx](#cosmos.tx.v1beta1.Tx) | repeated | txs is the list of queried transactions. |
| `tx_responses` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) | repeated | tx_responses is the list of queried TxResponses. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.tx.v1beta1.SimulateRequest"></a>

### SimulateRequest
SimulateRequest is the request type for the Service.Simulate
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [Tx](#cosmos.tx.v1beta1.Tx) |  | **Deprecated.** tx is the transaction to simulate. Deprecated. Send raw tx bytes instead. |
| `tx_bytes` | [bytes](#bytes) |  | tx_bytes is the raw transaction.

Since: cosmos-sdk 0.43 |






<a name="cosmos.tx.v1beta1.SimulateResponse"></a>

### SimulateResponse
SimulateResponse is the response type for the
Service.SimulateRPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_info` | [cosmos.base.abci.v1beta1.GasInfo](#cosmos.base.abci.v1beta1.GasInfo) |  | gas_info is the information about gas used in the simulation. |
| `result` | [cosmos.base.abci.v1beta1.Result](#cosmos.base.abci.v1beta1.Result) |  | result is the result of the simulation. |





 <!-- end messages -->


<a name="cosmos.tx.v1beta1.BroadcastMode"></a>

### BroadcastMode
BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BROADCAST_MODE_UNSPECIFIED | 0 | zero-value for mode ordering |
| BROADCAST_MODE_BLOCK | 1 | BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for the tx to be committed in a block. |
| BROADCAST_MODE_SYNC | 2 | BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for a CheckTx execution response only. |
| BROADCAST_MODE_ASYNC | 3 | BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns immediately. |



<a name="cosmos.tx.v1beta1.OrderBy"></a>

### OrderBy
OrderBy defines the sorting order

| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDER_BY_UNSPECIFIED | 0 | ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case. |
| ORDER_BY_ASC | 1 | ORDER_BY_ASC defines ascending order |
| ORDER_BY_DESC | 2 | ORDER_BY_DESC defines descending order |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.tx.v1beta1.Service"></a>

### Service
Service defines a gRPC service for interacting with transactions.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Simulate` | [SimulateRequest](#cosmos.tx.v1beta1.SimulateRequest) | [SimulateResponse](#cosmos.tx.v1beta1.SimulateResponse) | Simulate simulates executing a transaction for estimating gas usage. | POST|/cosmos/tx/v1beta1/simulate|
| `GetTx` | [GetTxRequest](#cosmos.tx.v1beta1.GetTxRequest) | [GetTxResponse](#cosmos.tx.v1beta1.GetTxResponse) | GetTx fetches a tx by hash. | GET|/cosmos/tx/v1beta1/txs/{hash}|
| `BroadcastTx` | [BroadcastTxRequest](#cosmos.tx.v1beta1.BroadcastTxRequest) | [BroadcastTxResponse](#cosmos.tx.v1beta1.BroadcastTxResponse) | BroadcastTx broadcast transaction. | POST|/cosmos/tx/v1beta1/txs|
| `GetTxsEvent` | [GetTxsEventRequest](#cosmos.tx.v1beta1.GetTxsEventRequest) | [GetTxsEventResponse](#cosmos.tx.v1beta1.GetTxsEventResponse) | GetTxsEvent fetches txs by event. | GET|/cosmos/tx/v1beta1/txs|

 <!-- end services -->



<a name="ethermint/crypto/v1/ethsecp256k1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ethermint/crypto/v1/ethsecp256k1/keys.proto



<a name="ethermint.crypto.v1.ethsecp256k1.PrivKey"></a>

### PrivKey
PrivKey defines a type alias for an ecdsa.PrivateKey that implements
Tendermint's PrivateKey interface.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="ethermint.crypto.v1.ethsecp256k1.PubKey"></a>

### PubKey
PubKey defines a type alias for an ecdsa.PublicKey that implements
Tendermint's PubKey interface. It represents the 33-byte compressed public
key format.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ethermint/types/v1/account.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ethermint/types/v1/account.proto



<a name="ethermint.types.v1.EthAccount"></a>

### EthAccount
EthAccount implements the authtypes.AccountI interface and embeds an
authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_account` | [cosmos.auth.v1beta1.BaseAccount](#cosmos.auth.v1beta1.BaseAccount) |  |  |
| `code_hash` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ethermint/types/v1/web3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ethermint/types/v1/web3.proto



<a name="ethermint.types.v1.ExtensionOptionsWeb3Tx"></a>

### ExtensionOptionsWeb3Tx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `typed_data_chain_id` | [uint64](#uint64) |  | typed data chain id used only in EIP712 Domain and should match Ethereum network ID in a Web3 provider (e.g. Metamask). |
| `fee_payer` | [string](#string) |  | fee payer is an account address for the fee payer. It will be validated during EIP712 signature checking. |
| `fee_payer_sig` | [bytes](#bytes) |  | fee payer sig is a signature data from the fee paying account, allows to perform fee delegation when using EIP712 Domain. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="prc10/prc10.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## prc10/prc10.proto



<a name="plugchain.prc10.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `issue_token_base_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `operate_token_fee_ratio` | [string](#string) |  |  |






<a name="plugchain.prc10.Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `scale` | [uint32](#uint32) |  |  |
| `min_unit` | [string](#string) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="prc10/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## prc10/genesis.proto



<a name="plugchain.prc10.GenesisState"></a>

### GenesisState
GenesisState defines the token module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#plugchain.prc10.Params) |  |  |
| `tokens` | [Token](#plugchain.prc10.Token) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="prc10/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## prc10/query.proto



<a name="plugchain.prc10.QueryFeesRequest"></a>

### QueryFeesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |






<a name="plugchain.prc10.QueryFeesResponse"></a>

### QueryFeesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exist` | [bool](#bool) |  |  |
| `issue_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `operate_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="plugchain.prc10.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="plugchain.prc10.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#plugchain.prc10.Params) |  |  |
| `res` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="plugchain.prc10.QueryTokenRequest"></a>

### QueryTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="plugchain.prc10.QueryTokenResponse"></a>

### QueryTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Token` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="plugchain.prc10.QueryTokensRequest"></a>

### QueryTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="plugchain.prc10.QueryTokensResponse"></a>

### QueryTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Tokens` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.prc10.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Token` | [QueryTokenRequest](#plugchain.prc10.QueryTokenRequest) | [QueryTokenResponse](#plugchain.prc10.QueryTokenResponse) |  | GET|/token/tokens/{denom}|
| `Fees` | [QueryFeesRequest](#plugchain.prc10.QueryFeesRequest) | [QueryFeesResponse](#plugchain.prc10.QueryFeesResponse) |  | GET|/token/fee/{symbol}|
| `Params` | [QueryParamsRequest](#plugchain.prc10.QueryParamsRequest) | [QueryParamsResponse](#plugchain.prc10.QueryParamsResponse) |  | GET|/token/params|
| `Tokens` | [QueryTokensRequest](#plugchain.prc10.QueryTokensRequest) | [QueryTokensResponse](#plugchain.prc10.QueryTokensResponse) |  | GET|/token/tokens|

 <!-- end services -->



<a name="prc10/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## prc10/tx.proto



<a name="plugchain.prc10.MsgBurnToken"></a>

### MsgBurnToken
MsgBurnToken defines an SDK message for burning some tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.prc10.MsgBurnTokenResponse"></a>

### MsgBurnTokenResponse







<a name="plugchain.prc10.MsgEditToken"></a>

### MsgEditToken
MsgEditToken defines an SDK message for editing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.prc10.MsgEditTokenResponse"></a>

### MsgEditTokenResponse







<a name="plugchain.prc10.MsgIssueToken"></a>

### MsgIssueToken
MsgIssueToken defines an SDK message for issuing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `scale` | [uint32](#uint32) |  |  |
| `min_unit` | [string](#string) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.prc10.MsgIssueTokenResponse"></a>

### MsgIssueTokenResponse







<a name="plugchain.prc10.MsgMintToken"></a>

### MsgMintToken
MsgMintTokenResponse defines the Msg/MintToken response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.prc10.MsgMintTokenResponse"></a>

### MsgMintTokenResponse







<a name="plugchain.prc10.MsgTransferOwnerToken"></a>

### MsgTransferOwnerToken
MsgTransferOwnerToken defines an SDK message for transferring the token owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |






<a name="plugchain.prc10.MsgTransferOwnerTokenResponse"></a>

### MsgTransferOwnerTokenResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.prc10.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueToken` | [MsgIssueToken](#plugchain.prc10.MsgIssueToken) | [MsgIssueTokenResponse](#plugchain.prc10.MsgIssueTokenResponse) |  | |
| `MintToken` | [MsgMintToken](#plugchain.prc10.MsgMintToken) | [MsgMintTokenResponse](#plugchain.prc10.MsgMintTokenResponse) |  | |
| `EditToken` | [MsgEditToken](#plugchain.prc10.MsgEditToken) | [MsgEditTokenResponse](#plugchain.prc10.MsgEditTokenResponse) |  | |
| `BurnToken` | [MsgBurnToken](#plugchain.prc10.MsgBurnToken) | [MsgBurnTokenResponse](#plugchain.prc10.MsgBurnTokenResponse) |  | |
| `TransferOwnerToken` | [MsgTransferOwnerToken](#plugchain.prc10.MsgTransferOwnerToken) | [MsgTransferOwnerTokenResponse](#plugchain.prc10.MsgTransferOwnerTokenResponse) |  | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

