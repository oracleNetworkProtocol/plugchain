---
order: 1
---

# Introduction

## plugchain Network

The *plugchain* network is an internet of blockchains intended to provide a technology foundation that facilitates construction of distributed business applications.

The *plugchain* network is part of the larger Cosmos network -- all zones in the network would be able to interact with any other zone in the Cosmos network over the standard IBC protocol.  By introducing a layer of service semantics into the network, we are going to provide an innovative solution that enables a whole new set of business scenarios, which would result in an increase in scale and diversity of the Cosmos network.

## plugchain 

At the "center" of the *plugchain* network is a blockchain known as the *PLUGChain Hub*, which is a Proof-of-Stake (PoS) blockchain built on Cosmos SDK and Tendermint.  It will be the first regional hub that connects to the Cosmos Hub as one of its zones.  The *plugchain* hub is equipped with a service protocol that coordinates on-chain transaction processing with off-chain data processing and business logic execution.  We have also enhanced the IBC protocol to facilitate cross-chain invocation of those off-chain services, if so desired.

The service protocol and enhanced IBC protocol could eventually be contributed back into Cosmos SDK, allowing SDK users to develop blockchains that are compatible with the *plugchain* network.

Client-facing, programming language specific SDKs will also be available to make it easy to provide and consume off-chain services within the *plugchain* network.

## plugchain Tokens

The *plugchain* hub has its own native token known as *plug*.  It is designed to serve three purposes in the network.

* **Staking.**  Similar to the ATOM token in the Cosmos Hub, the *plugchain* token will be used as a staking token to secure the PoS blockchain.

* **Transaction Fee.**  The *plugchain* token will also be used to pay fees for all transactions in the *plugchain* network.

* **Service Fee.**  It is required that service providers in the *plugchain* network charge service fees denominated in the *plugchain* token.

It is intended that the *plugchain* network will eventually support all whitelisted fee tokens from the Cosmos network, which can be used to pay the transaction fees and service fees.
