---
order: 7
---

## Events

描述不同事件的值

## message.module
- bank
- distribution
- governance
- staking
- slashing
- token
- transfer
- evm
- liquidity
- ibc_client
- ibc_channel
- ibc_connection


## message.action
- /cosmos.bank.v1beta1.MsgSend
- /cosmos.bank.v1beta1.MsgMultiSend
- /cosmos.distribution.v1beta1.MsgFundCommunityPool
- /cosmos.distribution.v1beta1.MsgSetWithdrawAddress
- /cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward
- /cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission
- /cosmos.gov.v1beta1.MsgVote
<!-- - /cosmos.gov.v1beta1.MsgVoteWeighted
- /cosmos.gov.v1beta1.MsgSubmitProposal
- /cosmos.gov.v1beta1.MsgDeposit -->
- /cosmos.slashing.v1beta1.MsgUnjail
- /cosmos.staking.v1beta1.MsgBeginRedelegate
- /cosmos.staking.v1beta1.MsgCreateValidator
- /cosmos.staking.v1beta1.MsgDelegate
- /cosmos.staking.v1beta1.MsgEditValidator
- /cosmos.staking.v1beta1.MsgUndelegate
- /ethermint.evm.v1.MsgEthereumTx
- /ibc.applications.transfer.v1.MsgTransfer
- /ibc.core.client.v1.MsgCreateClient
- /ibc.core.channel.v1.MsgAcknowledgement
- /ibc.core.channel.v1.MsgChannelOpenAck
- /ibc.core.channel.v1.MsgChannelOpenInit
- /ibc.core.channel.v1.MsgRecvPacket
- /ibc.core.channel.v1.MsgTimeout
- /ibc.core.connection.v1.MsgConnectionOpenAck
- /ibc.core.connection.v1.MsgConnectionOpenInit
- /plugchain.prc10.MsgBurnToken
- /plugchain.prc10.MsgEditToken
- /plugchain.prc10.MsgIssueToken
- /plugchain.prc10.MsgMintToken
- /tendermint.liquidity.v1beta1.MsgCreatePool
- /tendermint.liquidity.v1beta1.MsgDepositWithinBatch
- /tendermint.liquidity.v1beta1.MsgSwapWithinBatch
- /tendermint.liquidity.v1beta1.MsgWithdrawWithinBatch