---
order: 3
---

# Restful API


## Query API activation method and configuration

All routes are configured via `~/.plugchain/config/app.toml` under the following fields:

- The `api.enable = true|false` field defines whether to enable the REST service, the default is `false`.
- The `api.address = {string}` field defines the address the server should bind to (actually the port, since the host should remain at `0.0.0.0`). Defaults to `tcp://0.0.0.0:1317`.
- Some other API configuration options are defined and commented in `~/.plugchain/config/app.toml`, please refer to that file directly.
```
################################################## #############################
### API Configuration ###
################################################## #############################

[api]

# Enable defines if the API server should be enabled.
enable = true

# Swagger defines if swagger documentation should automatically be registered.
swagger = true

# Address defines the API server to listen on.
address = "tcp://0.0.0.0:1317"
```
### gRPC-gateway REST routing

If gRPC cannot be used for various reasons (for example, you are building a web application and the browser does not support HTTP2, which gRPC depends on), Plug Chain Hub provides REST routing through gRPC-gateway.

[gRPC-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) is a tool to expose gRPC endpoints as REST endpoints. For every RPC endpoint defined in a Protobuf service, the SDK provides a REST equivalent.

For application developers, the gRPC-gateway REST route needs to be connected to the REST server by calling the `RegisterGRPCGatewayRoutes` method on the ModuleManager.

### Swagger

The [Swagger](https://swagger.io/) (or OpenAPIv2 ) specification file is under the `/swagger` path on the API server. Swagger is an open specification that describes API endpoints for server services, including descriptions, input parameters, return types, and more information about each endpoint.

The `/swagger` endpoint can be configured via the `api.swagger` field in `~/.plugchain/config/app.toml`, defaults to `false`.


For application developers, you may want to generate your own Swagger definitions based on custom modules. You can start from Plug Chain Hub's [Swagger Generation Script](https://github.com/oracleNetworkProtocol/plugchain/blob/main/scripts/protoc-swagger-gen.sh).

## API endpoints

**Part of Plug Chain Hub API endpoints**

| API Endpoints | Description |
| :----------------------------------------------------------------- | :---------------------------------------------------------- |
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; |
| `GET` `/cosmos/auth/v1beta1/accounts/{address}`                                                                           | return account information |
| `GET` `/cosmos/auth/v1beta1/params`                                                                           | Query all parameters |
| `GET` `/cosmos/bank/v1beta1/balances/{address}`                                                                           | Query all tokens of an account |
| `GET` `/cosmos/bank/v1beta1/balances/{address}/{denom}`                                                                           | Query the balance of a single token of an account |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata`                                                                          | Query the metadata of all registered tokens of the client |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata/{denom}`                                                                          | Query the metadata of a token on the client side |
| `GET` `/cosmos/bank/v1beta1/params`                                                                           | Query the parameters of the bank module |
| `GET` `/cosmos/bank/v1beta1/supply`                                                                           | Query the total supply of all tokens |
| `GET` `/cosmos/bank/v1beta1/supply/{denom}`                                                                           | Query the total issuance of a token |
| `GET` `/cosmos/distribution/v1beta1/community_pool`                                                                           | Query the tokens in the community pool |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards`                                                                           | Query the total rewards accumulated by delegators at each validator |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards/{validator_address}`                                                                           | Query the rewards accumulated by a delegator at a validator |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/validators`                                                                            | Query all delegated validations of a delegator people |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address`                                                                          | Query a delegator's withdrawal address |
| `GET` `/cosmos/distribution/v1beta1/params`                                                                           | Query distribution module parameters |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/commission`                                                                            | Query the accumulated commission of a validator |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/outstanding_rewards`                                                                           | Query a validator's reward |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/slashes`                                                                           | Query a validator's punishment event |
| `GET` `/cosmos/evidence/v1beta1/evidence`                                                                             | Query all evidence |
| `GET` `/cosmos/evidence/v1beta1/evidence/{evidence_hash}`                                                                             | Query evidence by hash |
| `GET` `/cosmos/gov/v1beta1/params/{params_type}`                                                                          | Query gov module parameters |
| `GET` `/cosmos/gov/v1beta1/proposals`                                                                             | Query all proposals in the specified state |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}`                                                                           | Query proposals by ID |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits`                                                                          | Query all mortgages for a proposal |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}`                                                                          | Query the mortgage information of a staker in a proposal |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/tally`                                                                             | Query voting statistics for a proposal |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes`                                                                             | Query all votes for a proposal |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}`                                                                             | Query the voting information of a voter in a proposal |
| `GET` `/cosmos/params/v1beta1/params`                                                                             | Query the specified parameters of a module by subspace and key |
| `GET` `/cosmos/slashing/v1beta1/params`                                                                           | Query slashing module parameters |
| `GET` `/cosmos/slashing/v1beta1/signing_infos`                                                                            | Query the signature information of all validators |
| `GET` `/cosmos/slashing/v1beta1/signing_infos/{cons_address}`                                                                             | Query the signature information of an address |
| `GET` `/cosmos/staking/v1beta1/delegations/{delegator_addr}`                                                                          | Query all delegation information of a delegator |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations`                                                                             | Query the redelegation information of an address |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations`                                                                             | Query all unbonding information for a given delegator |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators`                                                                            | Query all validator information of the specified delegator |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}`                                                        | Query the validator information of the specified validator and delegator pair |
| `GET` `/cosmos/staking/v1beta1/historical_info/{height}` | Query the historical information of the specified height |
| `GET` `/cosmos/staking/v1beta1/params` | Query staking module parameters |
| `GET` `/cosmos/staking/v1beta1/pool` | Query pool information |
| `GET` `/cosmos/staking/v1beta1/validators` | Query all validators matching the specified state |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}` | Query validator information by validator address |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations` | Query a validator's delegation information |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}` | Query the delegation information between the specified validator and the delegator |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation` | Query the unbonding information between the specified validator and the delegator |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations` | Query a validator's unbonding information |
| `GET` `/cosmos/upgrade/v1beta1/applied_plan/{name}` | Query the applied upgrade plan by name |
| `GET` `/cosmos/upgrade/v1beta1/current_plan` | Query the current upgrade plan |
| `GET` `/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}` | Query the consensus state, which will be used as the trusted kernel for the next version of this chain |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/latest` | Returns the latest block height |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/{height}` | Query blocks with specified height |
| `GET` `/cosmos/base/tendermint/v1beta1/node_info` | Query current node information |
| `GET` `/cosmos/base/tendermint/v1beta1/syncing` | Query node synchronization information |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/latest` | Query the current block validator set |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/{height}` | Query the set of validators with a specified height |
| `POST` `/cosmos/tx/v1beta1/simulate` | Simulate transaction execution to estimate gas consumption |
| `GET` `/cosmos/tx/v1beta1/txs` | Filter transactions by events |
| `POST` `/cosmos/tx/v1beta1/txs` | Broadcast transaction |
| `GET` `/cosmos/tx/v1beta1/txs/{hash}` | Query transactions by hash |


## query transaction

Querying a transaction using the gRPC-gateway REST endpoint can be done by sending a GET request, an example is shown below:

- **Query tx by hash:** `/cosmos/tx/v1beta1/txs/{hash}`

    ```bash
    curl -X GET \
        -H "accept: application/json" \
        "http://localhost:1317/cosmos/tx/v1beta1/txs/{hash}"
    ```

- **Query tx by events:** `/cosmos/tx/v1beta1/txs` [event format {eventType}.{eventAttribute}={value}](#event)
    
    ``` bash
    curl -X GET \
        -H "accept: application/json" \
        "http://localhost:1317/cosmos/tx/v1beta1/txs?events={eventType}.{eventAttribute}={value}&events={eventType}.{eventAttribute}={value}&pagination.limit=100&pagination.offset=0"
    ```
## event

Events are implemented in the Cosmos SDK as aliases of the ABCI `Event` type, and
Takes the form: `{eventType}.{attributeKey}={attributeValue}`.

An event contains:

* A "type" for advanced categorization of events; for example, the Cosmos SDK uses the `"message"` type to filter events by `Msg`.
* The `attributes` list is a key-value pair that provides more information about classified events. For example, for the `"message"` type, we can filter events using `message.action={some_action}`, `message.module={some_module}`, or `message.sender={some_sender} key-value pairs`.
* Pagination query requires parameters `pagination.limit` and `pagination.offset` to be used together. egg: Query 100 records, starting from 0: `pagination.limit=100&pagination.offset=0` Query the parameters of the second page `pagination.limit=100&pagination.offset=100` The maximum value of the pagination.limit parameter is `100`

::: warning
To parse property values ​​as strings, be sure to add `'` (single quotes) around each property value.
:::

### example

The following example shows how to query events using the Cosmos SDK.

| Event | Description |
| ------------------------- | ------------------------------- |
| `tx.height=23` | Query all transactions with height 23 |
| `tx.hash='DF9738772AAECE776187EFF106190FF169F00C725968A15D23FA1DC9B4A1B651'` | Query the specified hash transaction |
| `message.action='/cosmos.bank.v1beta1.MsgSend'` | Query all transactions containing x/bank `Send`. |
| `message.module='bank'` | Query all transactions that contain messages from the `x/bank` module. |








## Construct and sign the transaction

Transactions cannot be constructed and signed using REST, only transactions can be broadcast. You can use [gRPC client](grpc-client.md) to construct and sign transactions.

## Broadcast transaction

Broadcasting a transaction using the gRPC-gateway REST endpoint `cosmos/tx/v1beta1/txs` can be done by sending a POST request as follows, where `NewTxBytes` is a protobuf encoded byte array of the signed transaction:

```go
import (
    "fmt"
    "encoding/base64"
)

func sendTx() error {
    //--shear--
    
    
    // base64 encode the encoded tx bytes
    // txBytes are obtained via `encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())` or `encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())`

    NewTxBytes := base64.StdEncoding.EncodeToString(txBytes)
    
    fmt.Println(NewTxBytes)
}
```

```bash
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"tx_bytes":"{{NewTxBytes}}","mode":"BROADCAST_MODE_SYNC"}' \
    "localhost:1317/cosmos/tx/v1beta1/txs"
```