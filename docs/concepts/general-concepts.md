---
order: 1
---

# General Concepts

## PLUGChain Hub Node Types

### Full Node

A full-node is a program that fully validates transactions and blocks of a blockchain. It is distinct from a light-node that only processes block headers and a small subset of transactions. Running a full-node requires more resources than a light-node but is necessary in order to be a validator. In practice, running a full-node only implies running a non-compromised and up-to-date version of the software with low network latency and without downtime.

### Validator Node

The [PLUGChain Hub](../get-started/intro.md#plugchaind-hub) is based on [Cosmos SDK](https://cosmos.network/docs/intro/) and [Tendermint](https://tendermint.com/docs/introduction/what-is-tendermint.html), which relies on a set of validators to secure the network. The role of validators is to run a full-node and participate in consensus by broadcasting votes which contain cryptographic signatures signed by their private keys. Validators commit new blocks to the blockchain and receive revenue in exchange for their work. They must also participate in governance by voting on proposals. Validators are weighted according to their total stake.

### Validator Candidate Node

Only top 100 bonded full nodes can become validator nodes, the rest will become candidates. The situation will change as delegation amount changes.

## PLUGChain Hub User Types

### Validator Operator

A validator operator is the only one who can operate the Validator's informations or participate in governance as the validator.

### Delegator

Delegators are plugchaind holders who cannot, or do not want to run a validator themselves. plugchaind holders can delegate plugchaind to a validator and obtain a part of their revenue in exchange. They can earn as much as the validators and only need to pay some commission.

### Profiler

Profiler is a special type of user who can submit software upgrade/halt proposals

### Trustee

Trustee is a special type of user who will receive funds from CommunityTaxUsage proposals

## plugchaind Token

The PLUGChain Hub has its own native token known as *plug*.  It is designed to serve three purposes in the network.

- **Staking.** Similar to the ATOM token in the Cosmos Hub, the plugchaind token will be used as a staking token to secure the PoS blockchain.
- **Transaction Fee.** The plugchaind token will also be used to pay fees for all transactions in the plugchaind network.
- **Service Fee.** It is required that service providers in the plugchaind network charge service fees denominated in the plugchaind token.

It is intended that the plugchaind network will eventually support all whitelisted fee tokens from the Cosmos network, which can be used to pay the transaction fees and service fees.

## Staking Rewards

The validator and its delegators can share the following rewards by proportion:

- **Block Inflation**

  Block Inflation exists to incentivize plugchaind holders to stake. The more staked plugchaind tokens are, more secure the network become(Read more about [Staking](../features/staking.md)).

  Block Inflation will be [distributed every block](../features/mint.md). The inflation rate of PLUGChain has been changing, ranging from 7% to 20%

- **Block Proposer Reward**

  In PLUGChain Hub, the probability for being a proposer is proportional to the validator's bonded tokens. If one proposed block is finalized, the proposer gets extra rewards for it.

- **Fee**

    Various transactions in the plug network need to pay a certain fee. PLUGChain Hub sets the minimum unit price of Gas (ie min-gas-price) in the system.
  
   Handling fees can only be paid using plug at present. In the future, a whitelist of handling fees may be maintained through on-chain governance. Any token in the whitelist can be used to pay for handling fees.

   The fee income will be distributed according to the weight of the validator.

### Staking Rewards Calculation Formula

The following formulas are based on the current [PLUGChain Hub Mainnet Params](gov-params.md).

#### Annual Rewards (ignore proposer rewards and fees)

- **AnnualInflation =** `Base * InflationRate` (aka 15.989 billion * 13% = 207857 million plug)
- **ValidatorRewards =** `(AnnualInflation / BondedTokens) * (1 - CommunityTax) * (ValidatorSelfDelegation +  DelegatorsDelegation * ValidatorCommissionRate)`
- **DelegatorRewards =** `(AnnualInflation / BondedTokens) * (1 - CommunityTax) * DelegatorSelfDelegation * (1 - ValidatorCommissionRate)`

#### Block Rewards

- **BlockInflation =** `AnnualInflation / (365*24*60*12)` (aka 329.555  plug)
- **ProposerExtraRewards =** `(BaseProposerReward + BonusProposerReward * PrecommitPower/TotalVotingPower) * (BlockInflation + BlockCollectedFees)`
- **BlockRewards =** `(BlockInflation + BlockCollectedFees) * (1 - CommunityTax) - ProposerExtraRewards`
- **ValidatorTotalRewards =**
  - Non-Proposer: `(BlockRewards / BondedTokens) * ValidatorBondedTokens`
  - Proposer: `NonProposerValidatorTotalRewards + ProposerExtraRewards`
- **Commission =** `ValidatorTotalRewards * ValidatorCommissionRate`
- **ValidatorRewards =** `ValidatorTotalRewards * (ValidatorSelfDelegation / ValidatorBondedTokens) + Commission`
- **DelegatorRewards =** `(ValidatorTotalRewards - Commission) * (DelegatorSelfDelegation / ValidatorBondedTokens)`

## Validator Responsibilities

Validators have two main responsibilities:

- **Be able to constantly run a correct version of the software:** Validators need to make sure that their servers are always online and their private keys are not compromised.
- **Actively participate in governance:** Validators are required to vote on every proposal.

Additionally, validators are expected to be active members of the community. They should always be up-to-date with the current state of the ecosystem so that they can easily adapt to any change.

## Validator Risks

- **Unavailability**: Validators are expected to keep signing votes for making new blocks. If a validator's signature has not been included in more than 30% of the last 34,560 blocks (which amounts to approximately 48 hours, assuming an average block-generating time of 5 seconds), this validator will get jailed and removed from current validatorset for 10 minutes, and their bonded tokens will get slashed by 1%.
- **Double Sign**: If the protocol detects that a validator voted multiple different opinions about the same block (same height/round), or voted for different blocks at the same height/round, this validator will get jailed and removed from current validatorset for 2 days. Their bonded tokens will get slashed by 5%.
<!-- - **Censorship**: If the protocol detects that a proposer included invalid transactions in a block, this validator will get jailed and removed from current validatorset for 2 days. -->

All metrics mentioned can be adjusted by `parameter-change` proposals.
