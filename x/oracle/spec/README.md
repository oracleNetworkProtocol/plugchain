<!--
order: 0
title: "Overview"
parent:
  title: "Oracle Module"
-->

# Oracle & Offchain Workers

## Abstract

This document specifies the Oracle module for the Cosmos SDK.

The Oracle module accepts and stores arbitrary `Claims` submitted by the chain validators. Multiple oracly types can run simultaniously.

The module includes helper methods to design off-chain workers that supply data to the Oracle module.

The Oracle modules assists with deciding how a consensus is reached on oracle results, however it is the responsibility of external modules to implement this logic as needed.

## Contents

1. **[Concepts](01_concepts.md)**
   - [Worker](01_concepts.md#Worker)
   - [Claim](01_concepts.md#Claim)
   - [Vote](01_concepts.md#Vote)
   - [Prevote](01_concepts.md#Prevote)
   - [Round](01_concepts.md#Round)
   - [PendingRounds](01_concepts.md#PendingRounds)
   - [Tallying Votes](01_concepts.md#Tallying-Votes)
   - [HouseKeeping](01_concepts.md#HouseKeeping)
   - [Rewards and Punishments](01_concepts.md#Rewards-and-Punishments)
   - [Validator Delegation](01_concepts.md#Validator-Delegation)

2) **[State](02_state.md)**
   - [Round](02_state.md#Round)
   - [PendingRounds](02_state.md#PendingRounds)
   - [Prevote](02_state.md#Prevote)
   - [FeedDelegateKey](02_state.md#FeedDelegateKey)
3) **[Messages](03_messages.md)**
   - [MsgDelegate](03_messages.md#MsgDelegate)
   - [MsgPrevote](03_messages.md#MsgPrevote)
   - [MsgVote](03_messages.md#MsgVote)
4) **[Events](04_events.md)**
   - [MsgDelegate](04_events.md#MsgDelegate)
   - [MsgPrevote](04_events.md#MsgPrevote)
   - [MsgVote](04_events.md#MsgVote)
5) **[Params](05_params.md)**
