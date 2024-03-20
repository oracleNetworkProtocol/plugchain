<!--
order: 1
-->

# Concepts

## Worker

A worker is an off-chain process that is triggered by a Tendermint event. The module includes a `start-worker` tx command that will listen for Tendermint events and trigger user-defined event handlers. Before use, the worker must be initialized with the necessary event handlers. This should be done in the `/cmd/<appName>d/cmd/root.go` file.

### WorkerHelper

The Oracle module provides some useful utility functions for running worker handlers and submitting results (see `/x/oracle/client/cli/workerHleper.go`).

## Claim

A Claim is an abstract type that represents the result of an off-chain worker process. All claims must have a `Type`. You must specify a list of `ClaimTypes` and their settings in the oracle module's `ClaimParams` parameter.

## Vote

A `Vote` is a `Claim` submitted by a validator. Votes are tallied using the validator power to determine consensus.

## Prevote

ClaimType parameters can be configured to require a `Prevote`. This is highly recomended because it reduces centralization and free-rider risk in the Oracle. Whith prevote enabled, before submitting the actual claim data, the validator must submit a hash of a random salt + cliam hash + validator address. After the prevote round is over, validators submit their `Vote` wich includes the actual `Claim`.

## Round

All Claims must belong to a Round. Round duration can be set via `ClaimParams`. It can be as short as 1 block, or longer.

### PendingRounds

When a new Vote is cast, the associate `Round` will be added to the pending que if its not already there.

### LatestFinalizedRound

Once a round is finalized (consensus is reached), `LatestFinalizedRound` is updated with the roundId. Claims for with a `RoundId` lower than the `LatestFinalizedRound` are not accepted.

## Tallying Votes

Apps relying on the Oracle will most likely want to tally the oracle votes. This can be done in some module's `EndBlock` method. The Oracle Module provides a helper method for tallying the votes that returns the `RoundResult` struct once the required vote threshold has been met.

## HouseKeeping

It is the responsibility of the app using the oracle module to remove stale oracle data.

- Once it is determined that a Claim has reached consensus, the app should let the Oracle know this by calling `FinalizeRound`. This will remove the `Round` from the pending que and updated `LatestFinalizedRound` and remove all votes and claims associated with the round.

## Rewards and Punishments

Rewards and Slashing logic is outside of the scope of the Oracle module. Modules relying on oracle data may implement their own logic. We reccomend taking a look at the (https://docs.terra.money/dev/spec-oracle.html#concepts)[Terra Oracle] and the (https://github.com/PeggyJV/sommelier/tree/main/x/oracle)[Sommelier Oracle] for examples.

## Validator Delegation

Validators are able to delegate oracle responsibilities to a cosmos address of their choice.
