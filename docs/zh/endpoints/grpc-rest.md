---
order: 3
---

# gRPC Gateway JSON REST


## Query API 激活方式和配置

所有路由均通过 `~/.plugchain/config/app.toml` 中以下字段下配置：

- `api.enable = true|false` 字段定义是否启用 REST 服务，默认为 `false`。
- `api.address = {string}` 字段定义服务器应绑定到的地址（实际为端口，因为主机应保持在 `0.0.0.0`）。 默认为 `tcp://0.0.0.0:1317`。
- 在 `~/.plugchain/config/app.toml` 中定义了一些其他 API 配置选项并附有注释，请直接参考该文件。

### gRPC-gateway REST 路由

如果由于各种原因而不能使用 gRPC（例如，您正在构建 Web 应用，且浏览器不支持 gRPC 依赖的 HTTP2），则 PLUGChain Hub 会通过 gRPC-gateway 提供 REST 路由。

[gRPC-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) 是将 gRPC 端点公开为 REST 端点的工具。对于 Protobuf 服务中定义的每个 RPC 端点，SDK提供了 REST 等效项。

对于应用程序开发者，需要将 gRPC-gateway REST 路由连接到 REST 服务器，通过在 ModuleManager 上调用 `RegisterGRPCGatewayRoutes` 方法完成。

### Swagger

[Swagger](https://swagger.io/)（或 OpenAPIv2 ）规范文件在 API 服务器上的 `/swagger` 路径下。 Swagger 是一个开放的规范，描述服务器服务的 API 端点，包括描述、输入参数、返回类型以及有关每个端点的更多信息。

可以通过 `~/.plugchain/config/app.toml` 中的 `api.swagger` 字段配置启用 `/swagger` 端点，默认为 `false`。

对于应用程序开发人员，您可能希望基于自定义模块生成自己的 Swagger 定义。可以从 PLUGChain Hub 的 [Swagger 生成脚本](https://github.com/oracleNetworkProtocol/plugchain/blob/main/testnet/scripts/protoc-swagger-gen.sh)开始。

## API 端点

**PLUGChain Hub API 端点**

| API Endpoints                                                                                                                               | Description                                              | Legacy REST Endpoint                                                              |
| :------------------------------------------------------------------------------------------------------------------------------------------ | :------------------------------------------------------- | :-------------------------------------------------------------------------------- |
| `GET` `/cosmos/auth/v1beta1/accounts/{address}`                                                                                             | 返回账户信息                                             | `GET` `/auth/accounts/{address}`                                                  |
| `GET` `/cosmos/auth/v1beta1/params`                                                                                                         | 查询所有参数                                             |                                                                                   |
| `GET` `/cosmos/bank/v1beta1/balances/{address}`                                                                                             | 查询某个账户的所有 token                                 | `GET` `/bank/balances/{address}`                                                  |
| `GET` `/cosmos/bank/v1beta1/balances/{address}/{denom}`                                                                                     | 查询一个账户的单种 token 余额                            |                                                                                   |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata`                                                                                                | 查询客户端所有已注册 token 的元数据                      |                                                                                   |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata/{denom}`                                                                                        | 查询客户端一种 token 的元数据                            |                                                                                   |
| `GET` `/cosmos/bank/v1beta1/params`                                                                                                         | 查询 bank 模块的参数                                     |                                                                                   |
| `GET` `/cosmos/bank/v1beta1/supply`                                                                                                         | 查询所有 token 的总发行量                                | `GET` `/bank/total`                                                               |
| `GET` `/cosmos/bank/v1beta1/supply/{denom}`                                                                                                 | 查询一种 token 的总发行量                                | `GET` `/bank/total/{denom}`                                                       |
| `GET` `/cosmos/distribution/v1beta1/community_pool`                                                                                         | 查询社区池中的 token                                     | `GET` `/distribution/community_pool`                                              |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards`                                                                 | 查询委托人在每个验证人处累计的总奖励                     | `GET` `/distribution/delegators/{delegatorAddr}/rewards`                          |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards/{validator_address}`                                             | 查询委托人在一个验证人处累积的奖励                       | `GET` `/distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}`          |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/validators`                                                              | 查询一个委托人所有委托的验证人                           |                                                                                   |
| `GET` `/cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address`                                                        | 查询一个委托人的提款地址                                 | `GET` `/distribution/delegators/{delegatorAddr}/withdraw_address`                 |
| `GET` `/cosmos/distribution/v1beta1/params`                                                                                                 | 查询 distribution 模块参数                               | `GET` `/distribution/parameters`                                                  |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/commission`                                                              | 查询一个验证人的累积佣金                                 |                                                                                   |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/outstanding_rewards`                                                     | 查询一个验证人的奖励                                     | `GET` `/distribution/validators/{validatorAddr}/outstanding_rewards`              |
| `GET` `/cosmos/distribution/v1beta1/validators/{validator_address}/slashes`                                                                 | 查询一个验证人的惩罚事件                                 |                                                                                   |
| `GET` `/cosmos/evidence/v1beta1/evidence`                                                                                                   | 查询所有 evidence                                        |                                                                                   |
| `GET` `/cosmos/evidence/v1beta1/evidence/{evidence_hash}`                                                                                   | 通过哈希查询 evidence                                    |                                                                                   |
| `GET` `/cosmos/gov/v1beta1/params/{params_type}`                                                                                            | 查询 gov 模块参数                                        | `GET` `/gov/parameters/{params_type}`                                             |
| `GET` `/cosmos/gov/v1beta1/proposals`                                                                                                       | 查询指定状态的所有提议                                   | `GET` `/gov/proposals`                                                            |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}`                                                                                         | 通过 ID 查询提议                                         | `GET` `/gov/proposals/{proposal-id}`                                              |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits`                                                                                | 查询某个提议的所有抵押                                   | `GET` `/gov/proposals/{proposal-id}/deposits`                                     |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}`                                                                    | 查询一个提议中一个抵押者的抵押信息                       | `GET` `/gov/proposals/{proposal-id}/deposits/{depositor}`                         |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/tally`                                                                                   | 查询一个提议的投票统计                                   | `GET` `/gov/proposals/{proposal-id}/tally`                                        |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes`                                                                                   | 查询一个提议的所有投票                                   | `GET` `/gov/proposals/{proposal-id}/votes`                                        |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}`                                                                           | 查询一个提议中某个投票者的投票信息                       | `GET` `/gov/proposals/{proposal-id}/votes/{voter}`                                |
| `GET` `/cosmos/params/v1beta1/params`                                                                                                       | 通过 subspace 和 key 查询一个模块的指定参数              |                                                                                   |
| `GET` `/cosmos/slashing/v1beta1/params`                                                                                                     | 查询 slashing 模块参数                                   | `GET` `/slashing/parameters`                                                      |
| `GET` `/cosmos/slashing/v1beta1/signing_infos`                                                                                              | 查询所有验证人的签名信息                                 | `GET` `/slashing/signing_infos`                                                   |
| `GET` `/cosmos/slashing/v1beta1/signing_infos/{cons_address}`                                                                               | 查询一个地址的签名信息                                   |                                                                                   |
| `GET` `/cosmos/staking/v1beta1/delegations/{delegator_addr}`                                                                                | 查询一个委托人所有的委托信息                             | `GET` `/staking/delegators/{delegatorAddr}/delegations`                           |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations`                                                                   | 查询一个地址的重委托信息                                 | `GET` `/staking/redelegations`                                                    |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations`                                                           | 查询给定委托人的所有解委托信息                           | `GET` `/staking/delegators/{delegatorAddr}/unbonding_delegations`                 |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators`                                                                      | 查询指定委托人的所有验证人信息                           | `GET` `/staking/delegators/{delegatorAddr}/validators`                            |
| `GET` `/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}`                                                     | 查询指定验证人和委托人对的验证人信息                     | `GET` `/staking/delegators/{delegatorAddr}/validators/{validatorAddr}`            |
| `GET` `/cosmos/staking/v1beta1/historical_info/{height}`                                                                                    | 查询指定高度的历史信息                                   |                                                                                   |
| `GET` `/cosmos/staking/v1beta1/params`                                                                                                      | 查询 staking 模块参数                                    | `GET` `/staking/parameters`                                                       |
| `GET` `/cosmos/staking/v1beta1/pool`                                                                                                        | 查询池子信息                                             | `GET` `/staking/pool`                                                             |
| `GET` `/cosmos/staking/v1beta1/validators`                                                                                                  | 查询匹配指定状态的所有验证人                             | `GET` `/staking/validators`                                                       |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}`                                                                                 | 通过验证人地址查询验证人信息                             | `GET` `/staking/validators/{validatorAddr}`                                       |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations`                                                                     | 查询一个验证人的委托信息                                 | `GET` `/staking/validators/{validatorAddr}/delegations`                           |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}`                                                    | 查询指定验证人和委托人之间的委托信息                     | `GET` `/staking/delegators/{delegatorAddr}/delegations/{validatorAddr}`           |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation`                               | 查询指定验证人和委托人之间的解委托信息                   | `GET` `/staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}` |
| `GET` `/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations`                                                           | 查询一个验证人的解委托信息                               | `GET` `/staking/validators/{validatorAddr}/unbonding_delegations`                 |
| `GET` `/cosmos/upgrade/v1beta1/applied_plan/{name}`                                                                                         | 通过名称查询已应用的升级计划                             |                                                                                   |
| `GET` `/cosmos/upgrade/v1beta1/current_plan`                                                                                                | 查询当前升级计划                                         |                                                                                   |
| `GET` `/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}`                                                                      | 查询共识状态，该状态将用作此链的下一版本的受信任内核     |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels`                                                                                                  | 查询所有 IBC channel                                     |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}`                                                                     | 查询一个 IBC channel                                     |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/client_state`                                                        | 查询与提供的 channel ID 关联的 channel 的客户端状态      |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/consensus_state/revision/{revision_number}/height/{revision_height}` | 查询与提供的 channel ID 关联的 channel 的共识状态        |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/next_sequence`                                                       | 返回给定 channel 的下一个接收 sequence                   |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_acknowledgements`                                             | 返回与 channel 关联的所有 packet 确认                    |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_acks/{sequence}`                                              | 查询已存储 packet 的 acknowledgement hash                |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_commitments`                                                  | 返回与 channel 关联的所有 packet commitment hash         |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_ack_sequences}/unreceived_acks`           | 返回与 channel 和 sequences 关联的所有未接收的 IBC 确认  |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_commitment_sequences}/unreceived_packets` | 返回与 channel 和 sequence 关联的所有未接收的 IBC 数据包 |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_commitments/{sequence}`                                       | 查询已存储 packet 的 commitment hash                     |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/channels/{channel_id}/ports/{port_id}/packet_receipts/{sequence}`                                          | 查询链上是否收到给定的 packet sequence                   |                                                                                   |
| `GET` `/ibc/core/channel/v1beta1/connections/{connection}/channels`                                                                         | 查询 connection 关联的所有 channel                       |                                                                                   |
| `GET` `/ibc/client/v1beta1/params`                                                                                                          | 查询 ibc client 模块参数                                 |                                                                                   |
| `GET` `/ibc/core/client/v1beta1/client_states`                                                                                              | 查询一条链的所有 IBC 轻客户端                            |                                                                                   |
| `GET` `/ibc/core/client/v1beta1/client_states/{client_id}`                                                                                  | 查询一个 IBC 轻客户端                                    |                                                                                   |
| `GET` `/ibc/core/client/v1beta1/consensus_states/{client_id}`                                                                               | 查询与给定客户端相关联的所有共识状态                     |                                                                                   |
| `GET` `/ibc/core/client/v1beta1/consensus_states/{client_id}/revision/{revision_number}/height/{revision_height}`                           | 查询与给定高度的客户端状态相关联的共识状态               |                                                                                   |
| `GET` `/ibc/core/connection/v1beta1/client_connections/{client_id}`                                                                         | 查询与客户端状态关联的连接路径                           |                                                                                   |
| `GET` `/ibc/core/connection/v1beta1/connections`                                                                                            | 查询所有的 IBC connection                                |                                                                                   |
| `GET` `/ibc/core/connection/v1beta1/connections/{connection_id}`                                                                            | 查询一个 IBC connection                                  |                                                                                   |
| `GET` `/ibc/core/connection/v1beta1/connections/{connection_id}/client_state`                                                               | 查询与 connection 关联的客户端状态                       |                                                                                   |
| `GET` `/ibc/core/connection/v1beta1/connections/{connection_id}/consensus_state/revision/{revision_number}/height/{revision_height}`        | 查询与 connection 关联的共识状态                         |                                                                                   |
| `GET` `/ibc/applications/transfer/v1beta1/denom_traces`                                                                                     | 查询所有 IBC token 的 denom 的追踪信息                   |                                                                                   |
| `GET` `/ibc/applications/transfer/v1beta1/denom_traces/{hash}`                                                                              | 查询一个 IBC token 的 denom 追踪信息                     |                                                                                   |
| `GET` `/ibc/applications/transfer/v1beta1/params`                                                                                           | 查询 ibc-transfer 模块参数                               

**Tendermint API 端点**

| API 端点                                                       | 描述                        | Legacy REST 端点                |
| :------------------------------------------------------------- | :-------------------------- | :------------------------------ |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/latest`          | 返回最新高度的区块          | `GET` `/blocks/latest`          |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/{height}`        | 查询指定高度的区块          | `GET` `/blocks/{height}`        |
| `GET` `/cosmos/base/tendermint/v1beta1/node_info`              | 查询当前节点信息            | `GET` `/node_info`              |
| `GET` `/cosmos/base/tendermint/v1beta1/syncing`                | 查询节点同步信息            | `GET` `/syncing`                |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/latest`   | 查询当前区块验证人集合      | `GET` `/validatorsets/latest`   |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/{height}` | 查询指定高度验证人集合      | `GET` `/validatorsets/{height}` |
| `POST` `/cosmos/tx/v1beta1/simulate`                           | 模拟交易执行以预估 Gas 消耗 |                                 |
| `GET` `/cosmos/tx/v1beta1/txs`                                 | 通过事件筛选交易            | `GET` `/txs`                    |
| `POST` `/cosmos/tx/v1beta1/txs`                                | 广播交易                    | `POST` `/txs`                   |
| `GET` `/cosmos/tx/v1beta1/txs/{hash}`                          | 通过哈希查询交易            | `GET` `/txs/{hash}`             |

## 构造和签名交易

使用 REST 不能构造和签名交易，只能广播交易。您可以使用 [gRPC 客户端](grpc-client.md) 构造和签名交易。

## 广播交易

使用 gRPC-gateway REST 端点 `cosmos/tx/v1beta1/txs` 广播交易可以通过发送 POST 请求来完成，如下所示，其中 `NewTxBytes` 是已签名交易的 protobuf 编码的字节数组：

```go
import (
    "fmt"
    "encoding/base64"
)

func sendTx() error {
    //--剪断--
    
    
    // base64 对编码的 tx 字节进行编码
    // txBytes 通过 `encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())` 或者 `encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())` 获得

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

## 查询交易

使用 gRPC-gateway REST 端点查询事务可以通过发送 GET 请求来完成，示例如下所示：

- **Query tx by hash:** `/cosmos/tx/v1beta1/txs/{hash}`

    ```bash
    curl -X GET \
        -H "accept: application/json" \
        "http://localhost:1317/cosmos/tx/v1beta1/txs/{hash}"
    ```

- **Query tx by events:** `/cosmos/tx/v1beta1/txs`

    ``` bash
    curl -X GET \
        -H "accept: application/json" \
        "http://localhost:1317/cosmos/tx/v1beta1/txs?events={event_content}"
    ```
