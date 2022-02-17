---
order: 6
---
# Gov Parameters

In Plug Chain Hub, there are some special parameters that can be modified through on-chain governance.
All the plugchaind holders may participate in the on-chain governance. If the community is not satisfied with certain modifiable parameters, it is available to submit a [parameter-change](../features/governance.md#usage-scenario-of-parameter-change) proposal, and the params will be changed online automatically when the proposal passes.

## Parameters in Auth

| key                           | Description                                                     | Range                     | Current |
| ----------------------------- | --------------------------------------------------------------- | ------------------------- | ------- |
| `auth/MaxMemoCharacters`      | Maximum number of characters in the memo field in a transaction | (0, 18446744073709551615] | 256     |
| `auth/TxSigLimit`             | Maximum number of signatures per transaction                    | (0, 18446744073709551615] | 7       |
| `auth/TxSizeCostPerByte`      | The amount of gas consumed per byte of the transaction          | (0, 18446744073709551615] | 10      |
| `auth/SigVerifyCostED25519`   | Gas spent on edd2519 algorithm signature verification           | (0, 18446744073709551615] | 590     |
| `auth/SigVerifyCostSecp256k1` | Gas spent on secp256k1 algorithm signature verification         | (0, 18446744073709551615] | 1000    |

## Parameters in Bank

| key                       | Description                                        | Range        | Current |
| ------------------------- | -------------------------------------------------- | ------------ | ------- |
| `bank/SendEnabled`        | Tokens that support transfer                       |              | []      |
| `bank/DefaultSendEnabled` | Whether to enable the transfer function by default | {true,false} | true    |

Details in [Bank](../features/bank.md)

## Parameters in Distribution

| key                                | Description                                       | Range        | Current |
| ---------------------------------- | ------------------------------------------------- | ------------ | ------- |
| `distribution/communitytax`        | Fees charged for withdrawal                       | [0, 1]       | 0.02    |
| `distribution/baseproposerreward`  | The base reward rate of the block proposer        | [0, 1]       | 0.01    |
| `distribution/bonusproposerreward` | Reward rate for block proposers                   | [0, 1]       | 0.04    |
| `distribution/withdrawaddrenabled` | Whether to support setting the withdrawal address | {true,false} | true    |

Details in [Distribution](../features/distribution.md)

## Parameters in Gov

| key                 | Description                                      | Range                                                    | Current                                                                                                        |
| ------------------- | ------------------------------------------------ | -------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `gov/depositparams` | Related parameters of the deposit mortgage phase | max_deposit_period:(0, 9223372036854775807]              | {"min_deposit": [{"denom": "uplugcn", "amount": "10000000"}], "max_deposit_period": "604800s" }                 |
| `gov/votingparams`  | Related parameters of the voting mortgage phase  | voting_period:(0, 9223372036854775807]                   | {"voting_period": "1209600s"}                                                                                   |
| `gov/tallyparams`   | Related parameters of the voting tally phase     | quorum:[0,1]<br>threshold:(0,1]<br/>veto_threshold:(0,1] | {"quorum":"0.334000000000000000","threshold": "0.500000000000000000","veto_threshold": "0.334000000000000000"} |

Details in [Governance](../features/governance.md)

## Parameters in IBC

| key                       | Description                             | Range        | Current                            |
| ------------------------- | --------------------------------------- | ------------ | ---------------------------------- |
| `ibc/AllowedClients`      | Clients that support ibc                |              | ["06-solomachine","07-tendermint"] |
| `transfer/SendEnabled`    | Whether to enable the transfer function | {true,false} | true                              |
| `transfer/ReceiveEnabled` | Whether to enable the receive function  | {true,false} | true                              |

## Parameters in Mint

| key              | Description                 | Range    | Current |
| ---------------- | --------------------------- | -------- | ------- |
| `mint/Inflation` | Token issuance frequency    | [0, 0.2] | 0.13    |
| `mint/MintDenom` | Denom of the token mintable |          | uplugcn   |

Details in [Mint](../features/mint.md)


## Parameters in Slashing

| key                                | Description                           | Range                     | Current |
| ---------------------------------- | ------------------------------------- | ------------------------- | ------- |
| `slashing/DowntimeJailDuration`    | Maximum downtime  (continuous)        | (0, 9223372036854775807]  | 600s   |
| `slashing/MinSignedPerWindow`      | Minimum signature rate in each window | [0, 1]                    | 0.5     |
| `slashing/SignedBlocksWindow`      | Sliding window for downtime slashing  | (0, 18446744073709551615] | 100   |
| `slashing/SlashFractionDoubleSign` | Penalty coefficient for double sign   | [0, 1]                    | 0.05    |
| `slashing/SlashFractionDowntime`   | Penalty coefficient for downtime      | [0, 1]                    | 0.01 |

Details in [Slashing](../features/slashing.md)

## Parameters in Staking

| key                         | Description                                                     | Range                    | Current  |
| --------------------------- | --------------------------------------------------------------- | ------------------------ | -------- |
| `staking/UnbondingTime`     | Mortgage redemption time                                        | (0, 9223372036854775807] | 1814400s |
| `staking/MaxValidators`     | Maximum number of validators                                    | (0, 4294967295]          | 100      |
| `staking/MaxEntries`        | The maximum number of unbinding/redelegation orders in progress | (0, 4294967295]          | 7        |
| `staking/BondDenom`         | Bond denom                                                      |                          | uplugcn    |
| `staking/HistoricalEntries` | Historical entries                                              | [0, 4294967295]          | 10000    |

Details in [Staking](../features/staking.md)
