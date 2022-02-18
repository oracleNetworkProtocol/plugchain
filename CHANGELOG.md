

# Changelog



## Released 

## v1.0.0 
The mainnet hard fork will be carried out on block height 300,000. The estimated date is March 5, 2022. The hard fork process is expected to take 24 hours

The Hardfork includes the following details:
- Change the from "chain-ID" to "plugchain_520-1"
- Change the Token name "PLUG" (the relevant coin data on the chain is the data with precision 0, and the coin name with precision 6 is used for wallet, browser, and external APP application display)
- Precision 0: UPLUGCN
- Precision 6: PLUGCN
- Modify the community proposal module (GOV) and shorten the "voting period" from 504 hours (21 days) to 288 hours (14 days) 
- Adjust the maximum number of verifiers to 50
- Online smart contract module functions for PVM (Plug Virtual Machine); token generation contract "PRC20" will be generated for usage
- Enable "UPLUGCN" to be burned.
- Overall adjustment of "X/Liquidity", "X/Token" and other module fee parameters for liquid providers.
- The wallet supports two key signature algorithms "Secp256k1" and "ETH_Secp256k1" from EVM (Ethereum Virtual Machine)

Attention for old on-chain records:
- Transaction records of "PLUG" will not be kep. Historical data will be uploaded to GitHub and open to the public.
- Current block height will be retained, but not for historical block heights.
- Retained historical DAO proposal information
- Retained the existing Liquidity data and existing Token data

## v0.7.3 - 2021-12-22
Stable version of existing functions in the warehouse

## v0.7.2 - 2021-12-1
* update token symbol len validate
* del keywords for plug
## v0.7.1 - 2021-10-18

- x/nft add query owner and nft finish
 
## v0.7.0 - 2021-10-12
- cosmos-sdk v0.42.9
- tendermint v0.34.12
- x/nft add functions: issueclass, issuenft, editnft, burnnft, transfernft, transferclass. Add query class, classes, NFT, collection, supply.

## v0.6.0 - 2021-09-29
- Mainnet upgrade 
- Bump x/token and x/nft
- Add [Liquidity](https://github.com/oracleNetworkProtocol/liquidity) module [v0.1.2](https://github.com/oracleNetworkProtocol/liquidity/tree/v0.1.2) 

## v0.5.0 - 2021-08-19
- cosmos-sdk v0.42.7
- tendermint v0.34.11
- Mainnet online version

## v0.2.0 - 2021-06-16
- cosmos-sdk v0.42.4
- tendermint v0.34.10
- Testnet online version