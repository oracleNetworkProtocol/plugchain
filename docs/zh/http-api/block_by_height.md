---
order: 9
---

# 通过块高获取块信息

- 获取块基础信息

### 请求方式
`GET`

### URL
`/server/v1/block/by_height`

### 请求示例

```
/server/v1/block/by_height?height=8237741
```


#### 请求参数
- `height string 必填` 块高

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `block_id object` : 区块hash信息,需要base转码
- `block` : 区块详细信息
    - `header`: 区块hash块高等信息
    - `data` : 交易信息
    - `last_commit` : 上个区块的信息

#### 返回示例
```json
{
    "block_id": {
        "hash": "NEIesSzKhsoeanVNnuEDiXfLWM2cm/RxnFfMs1pOF48=",
        "part_set_header": {
            "total": 1,
            "hash": "P5zkV65bxHNTV4FtAlbsYC1IWchgVxc4AccsnFATFP4="
        }
    },
    "block": {
        "header": {
            "version": {
                "block": "11",
                "app": "0"
            },
            "chain_id": "plugchain_520-1",
            "height": "8237741",
            "time": "2023-02-15T08:08:04.931248134Z",
            "last_block_id": {
                "hash": "9yvoB+qPCWnXJ/9XMXJCqHn2GyVI29FXLFmu1QtHF5A=",
                "part_set_header": {
                    "total": 1,
                    "hash": "BZzucI/GJRHrHnnEwl8E8nrC7stNVeRNa34t2pp5kI0="
                }
            },
            "last_commit_hash": "/4/K/8vGLQCTTJY1poKUe+6a4wr9Zth+9otWU9s2V/o=",
            "data_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
            "validators_hash": "p62m6Y4UENCszcX4jJuhd1+RZH62tR0NgtJEkwBb26s=",
            "next_validators_hash": "p62m6Y4UENCszcX4jJuhd1+RZH62tR0NgtJEkwBb26s=",
            "consensus_hash": "BICRvH3cKD93v7+R1zxE2ljD34qcvIZ0Bdi389qtoi8=",
            "app_hash": "ITVYgt+c5udNNgVUxZT7po7aATxVO+QmOn1zsEclu+g=",
            "last_results_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
            "evidence_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
            "proposer_address": "8XjX2xBVdcjLs7aAryIaCCk8sgI="
        },
        "data": {
            "txs": []
        },
        "evidence": {
            "evidence": []
        },
        "last_commit": {
            "height": "8237740",
            "round": 0,
            "block_id": {
                "hash": "9yvoB+qPCWnXJ/9XMXJCqHn2GyVI29FXLFmu1QtHF5A=",
                "part_set_header": {
                    "total": 1,
                    "hash": "BZzucI/GJRHrHnnEwl8E8nrC7stNVeRNa34t2pp5kI0="
                }
            },
            "signatures": [
                {
                    "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
                    "validator_address": "L3AGBg/pOhXMCO6C0fYOByt9w+c=",
                    "timestamp": "2023-02-15T08:08:04.885207465Z",
                    "signature": "pVJFQYUQXJviLMRPJVYT+OrkaFOc62X6yupKJriOCNlFKUZ04C3iyBySGdONbvDejv1a/P8E6fAC6kSDw7uWDw=="
                },
                {
                    "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
                    "validator_address": "8XjX2xBVdcjLs7aAryIaCCk8sgI=",
                    "timestamp": "2023-02-15T08:08:05.009627894Z",
                    "signature": "XME0EKjEB5aUhFk3S9ut2S5y3BmtfWWssLH/x03/dFCxsrd85zb6BUokhrlAEDcJMCq4guWtT1eG3+CMxT3NBg=="
                },
                {
                    "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
                    "validator_address": "ofxRu6POS4oMaQtk2KhPY4IITwE=",
                    "timestamp": "2023-02-15T08:08:05.001029443Z",
                    "signature": "4YzIl7pd+WzCC5HHk4/705K+2ab4hHJtaGLPjEx0se5aNpLF5q+SRJnd/lv6j9u9zc1Q0EfbyXHrpgZWs0LDDQ=="
                },
                {
                    "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
                    "validator_address": "HuA9yQeDsNwhZ46sQp+1fsEhEb0=",
                    "timestamp": "2023-02-15T08:08:04.892850177Z",
                    "signature": "gYV4TJNBZVXRDWLOXCAEvyY+CVCA03TFpduY3WbushpwSsDuokC8vMpvNQEy232zHteG84R5FJofcTNj/41zCg=="
                }
            ]
        }
    }
}
```