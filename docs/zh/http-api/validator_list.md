---
order: 30
---


# 验证者列表
 
- 获取所有验证者

### 请求方式
`GET`

### URL
`/server/v1/validator/validators`

### 请求示例

```
/server/v1/validator/validators
```


#### 请求参数
 
- 无

### 返回值
- `code int64 `: 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `validators array` : 验证者数组
    - `operator_address`: 验证者地址
    - `consensus_pubkey`: 共识公钥
    - `jailed`: 监禁状态
    - `status`: 验证者状态
    - `tokens`: token
    - `delegator_shares`: 质押股份
    - `description`: 验证者描述
    - `unbonding_height`: 锁定块高
    - `unbonding_time`: 锁定时间
    - `commission`: 共识信息
    - `min_self_delegation`: 最小质押量,单位uplugcn

#### 返回示例
```json5
{
  "validators": [
    {
      "operator_address": "gxvaloper1punlewa9vsaq69fwrn4ycfn3e2d7ejyvu5e0qp",
      "consensus_pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "M/foz1Xvig2pmkhLMusQfTSurR1m6+WF1z7kd2LvZKY="
      },
      "jailed": false,
      "status": "BOND_STATUS_BONDED",
      "tokens": "26388224777218",
      "delegator_shares": "26467548042011.338285901039781442",
      "description": {
        "moniker": "xmmj",
        "identity": "",
        "website": "",
        "security_contact": "",
        "details": ""
      },
      "unbonding_height": "5633909",
      "unbonding_time": "2022-09-20T07:47:40.444863950Z",
      "commission": {
        "commission_rates": {
          "rate": "0.100000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.010000000000000000"
        },
        "update_time": "2022-03-31T04:07:20.605045206Z"
      },
      "min_self_delegation": "1000000"
    },
    {
      "operator_address": "gxvaloper1zrvx4yqucs8atgmsj2elclcaz0qmkqepn586mh",
      "consensus_pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "MlZhQS06BPAArzuM8Ckk3pKF/jNTp+rbmMCsFPqXUu8="
      },
      "jailed": false,
      "status": "BOND_STATUS_BONDED",
      "tokens": "4579703601093",
      "delegator_shares": "4630593769070.622385612603990876",
      "description": {
        "moniker": "Technology-sharing",
        "identity": "",
        "website": "",
        "security_contact": "",
        "details": "Increase,staking,users"
      },
      "unbonding_height": "5621628",
      "unbonding_time": "2022-09-19T12:27:07.543031418Z",
      "commission": {
        "commission_rates": {
          "rate": "0.030000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.100000000000000000"
        },
        "update_time": "2022-04-10T05:31:43.424570202Z"
      },
      "min_self_delegation": "100000"
    },
    {
      "operator_address": "gxvaloper1z2sxux4lcg2rsx3j6xzkk0xmnzsqvyrpzdetym",
      "consensus_pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "mX3WFxMIatM55hLtSNmadjX2Kcke80AKcO+TYKxnF78="
      },
      "jailed": false,
      "status": "BOND_STATUS_BONDED",
      "tokens": "68792545103896",
      "delegator_shares": "69905806937386.088083291364299746",
      "description": {
        "moniker": "Hong Shan Fund",
        "identity": "",
        "website": "",
        "security_contact": "",
        "details": ""
      },
      "unbonding_height": "7717131",
      "unbonding_time": "2023-02-02T13:02:12.323757166Z",
      "commission": {
        "commission_rates": {
          "rate": "0.100000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.010000000000000000"
        },
        "update_time": "2021-08-21T03:31:23.378932988Z"
      },
      "min_self_delegation": "1000000"
    },
    {
      "operator_address": "gxvaloper1rysmrrz6z5lxgsnulw2u6vp6memmr6k3t2eu5d",
      "consensus_pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "wsSSyS6kwlzYFQ3mBofWS1jjGDOuO3j4yCgT0uvkJ4Y="
      },
      "jailed": false,
      "status": "BOND_STATUS_BONDED",
      "tokens": "195398477676210",
      "delegator_shares": "196182029483138.820814281658388466",
      "description": {
        "moniker": "GENESIS LABS",
        "identity": "",
        "website": "",
        "security_contact": "",
        "details": ""
      },
      "unbonding_height": "7150878",
      "unbonding_time": "2022-12-27T17:14:16.143139695Z",
      "commission": {
        "commission_rates": {
          "rate": "0.100000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.020000000000000000"
        },
        "update_time": "2022-06-14T01:24:42.668718950Z"
      },
      "min_self_delegation": "1000000"
    }
    //....
    //....
    //....
  ],
  "pagination": {
    "next_key": null,
    "total": "52"
  }
}
```