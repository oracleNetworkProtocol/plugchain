---
order: 13
---

# 获取提案列表

- 获取所有的提案信息

### 请求方式
`GET`

### URL
`/server/v1/gov/proposals`

### 请求示例

```
/server/v1/gov/proposals
```


#### 请求参数

- 无

### 返回值
- `code int64 `: 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `proposals array` : 验证者数组
    - `proposal_id string `: 验证者地址
    - `content object`: 提案详细内容
    - `status string`: 提案状态
    - `final_tally_result object`: 最终投票结果
    - `submit_time string `: 提交时间
    - `deposit_end_time string` : 存款结束时间
    - `total_deposit object`: 存款量
    - `voting_start_time string`: 投票开始时间
    - `voting_end_time string`: 投票结束时间

#### 返回示例
```json5
{
  "proposals": [
    {
      "proposal_id": "11",
      "content": {
        "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
        "title": "Slashing Param Change",
        "description": "Adjust the slash_fraction_downtime from 0.001 to 0.0001",
        "changes": [
          {
            "subspace": "slashing",
            "key": "SlashFractionDowntime",
            "value": "\"0.000100000000000000\""
          }
        ]
      },
      "status": "PROPOSAL_STATUS_VOTING_PERIOD",
      "final_tally_result": {
        "yes": "0",
        "abstain": "0",
        "no": "0",
        "no_with_veto": "0"
      },
      "submit_time": "2023-02-16T02:41:22.830539778Z",
      "deposit_end_time": "2023-02-18T02:41:22.830539778Z",
      "total_deposit": [
        {
          "denom": "uplugcn",
          "amount": "10000000"
        }
      ],
      "voting_start_time": "2023-02-16T02:41:22.830539778Z",
      "voting_end_time": "2023-02-21T02:41:22.830539778Z"
    },
    {
      "proposal_id": "10",
      "content": {
        "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
        "title": "v1.7",
        "description": "This on-chain upgrade governance proposal is to adopt plugchaind v1.7+ which includes a number of updates and fixes. Adjust the proposal fundraising time to 2 days and the voting time to 5 days. Rename plugcn to pc,dump cosmos-sdk v0.45+ ethermint v0.18+ tendermint v0.34+ ibc-go v3+ . del nft module and add authz module ...",
        "plan": {
          "name": "v1.7",
          "time": "0001-01-01T00:00:00Z",
          "height": "5633000",
          "info": "{\"binaries\":{\"linux/amd64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Linux_x86_64.tar.gz\",\"linux/arm64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Linux_arm64.tar.gz\",\"darwin/amd64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Darwin_x86_64.tar.gz\"}}",
          "upgraded_client_state": null
        }
      },
      "status": "PROPOSAL_STATUS_PASSED",
      "final_tally_result": {
        "yes": "4033515558113959",
        "abstain": "0",
        "no": "0",
        "no_with_veto": "0"
      },
      "submit_time": "2022-08-16T07:55:21.890959838Z",
      "deposit_end_time": "2022-08-19T07:55:21.890959838Z",
      "total_deposit": [
        {
          "denom": "uplugcn",
          "amount": "1000000000"
        }
      ],
      "voting_start_time": "2022-08-16T07:55:21.890959838Z",
      "voting_end_time": "2022-08-27T07:55:21.890959838Z"
    },
    //....
    //....
    //....
  ],
  "pagination": {
    "next_key": null,
    "total": "11"
  }
}
```