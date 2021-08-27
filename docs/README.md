# PLUGChain documentation

## PLUGChain Introduction

The Internet has long become an indispensable tool in people's lives. Many people have long been accustomed to obtaining the knowledge they need from the Internet, whether it is accurate or false. And in most cases, we will not judge the authenticity of the knowledge that the Internet brings to us, nor will we actively sort out the messy knowledge instillation. Data transmission all over the world has been very convenient, and at the same time, false news has also maintained a very high transmission speed. Nowadays, people urgently need to obtain real data quickly in the chaotic Internet world, and the same is true in the world of blockchain.

The blockchain industry has undergone years of development. Nowadays, from the aspects of finance, product traceability, government affairs, people's livelihood to the construction of smart cities, the technology application scenarios are continuously deepening and diversified, and they are closely related to the public. However, all existing application chains have a fatal problem-they cannot communicate credibly with the real world. The Oracle Network Protocol oracle machine is the tool that connects the application chain with external systems and promotes the real implementation of blockchain technology application scenarios.

Under the multiple requirements of the floating mid-air of the blockchain, the data cross-chain interaction of the oracle, and the application of the physical business scenario of the public chain, the PLUG R&D team is committed to building a system that can be used in the blockchain and the blockchain, the blockchain and the real world. Between, the high-performance public chain that realizes non-inductive interaction, that is, PLUGChain.

PLUGChain is the name of the Cosmos SDK application of the Cosmos Hub. It has 1 main entry point:

-`plugchaind` daemon and command line interface (CLI). A full node running the plugchain application.

The plugchain is built on the Cosmos SDK using the following modules:

- x/auth: account and signature.
- x/bank: Token transfer.
- x/staking: Staking logic.
- x/mint: Inflation logic.
- x/distribution: Cost distribution logic.
- x/slashing: Slashing logic.
- x/gov: governance logic.
- ibc-go/modules: Communication between blockchains. Hosted in the cosmos/ibc-go repository.
- x/params: Process application-level parameters.
- x/token: token logic. In development
- x/dex: automated market maker. Hosted in tendermint/liquidity repository.