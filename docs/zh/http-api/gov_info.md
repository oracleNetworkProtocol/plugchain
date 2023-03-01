---
order: 12
---

# 获取提案

- 通过提案ID获取提案信息

### 请求方式
`GET`

### URL
`/server/v1/gov/detail`

### 请求示例

```
/server/v1/gov/detail?proposal_id=11
```


#### 请求参数

- 无

### 返回值
- `code int64 `: 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `proposals array` : 验证者数组
  - `proposal_id string `: 验证者地址
  - `content object`: 提案详细内容
    - `@type`
    - `title`
    - `description`
    - `changes`
  - `status string`: 提案状态
  - `final_tally_result object`: 最终投票结果
    - `yes`:同意票数
    - `abstain`: 弃权票数
    - `no`: 反对票数
    - `no_with_veto`: 行使否决权的否定票数
  - `submit_time string `: 提交时间
  - `deposit_end_time string` : 存款结束时间
  - `total_deposit object`: 存款量
  - `voting_start_time string`: 投票开始时间
  - `voting_end_time string`: 投票结束时间

#### 返回示例
```json5
{
  "proposal": {
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
  }
}
```