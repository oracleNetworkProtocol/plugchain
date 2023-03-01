---
order: 10
---

# Get the latest block

### Request method
`GET`

### URL
`/server/v1/block/latest`

### Request example

```
/server/v1/block/latest
```


#### Request parameters
 
- Not
 
### Return value
- `jsonrpc string` : Jsonrpc version
- `id int` : id
- `result object` : Block hash information
- `error object` : Error object
    - `code int`: Error code
    - `message string`: Error information
    - `data string`: Error description
- `result object` : Block hash information
    - `block_id object` : Block hash information
    - `block object`: Block hash block height information
        - `header object` : Block header details
        - `data object` : Transaction information
        - `evidence object` : ""
        - `last_commit object` : Information of the previous block

#### Return to example
```json5
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "E37790A2903E1F488B49F1EE2C17B5FBEAFC6412034742AE2D3D7AC27F2322F4",
      "parts": {
        "total": 1,
        "hash": "1F5434FDD34EE09093FCBC0F837C0E72DA5F4380615A63866E75F747B1747092"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "plugchain_520-1",
        "height": "8251909",
        "time": "2023-02-16T06:18:32.774178625Z",
        "last_block_id": {
          "hash": "67338AB11AFC047F2A79F8BEFD6C8356A06D032683859D722889006B1C0D57A5",
          "parts": {
            "total": 1,
            "hash": "8F29D88783D409504D47C1C1DC161AAB26BDE08C89AFB7BA77DBFE261139DCEF"
          }
        },
        "last_commit_hash": "E3F9216273B8ADE958B1512C8C50330678451B4411AA00449C4DD14AEF4004F0",
        "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "validators_hash": "FE75E124EB85E465D2EAB9B31D1FC002680BDF6D4A1CA204EDDF89C5D15DF8A7",
        "next_validators_hash": "FE75E124EB85E465D2EAB9B31D1FC002680BDF6D4A1CA204EDDF89C5D15DF8A7",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "B03DF87630364405D92559A0AA872837775761EDE895541574A4337E7F9F0AF7",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "53343101192D407D28EEBB2A2CEF6D40BB2E9072"
      },
      "data": {
        "txs": []
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "8251908",
        "round": 0,
        "block_id": {
          "hash": "67338AB11AFC047F2A79F8BEFD6C8356A06D032683859D722889006B1C0D57A5",
          "parts": {
            "total": 1,
            "hash": "8F29D88783D409504D47C1C1DC161AAB26BDE08C89AFB7BA77DBFE261139DCEF"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "2F7006060FE93A15CC08EE82D1F60E072B7DC3E7",
            "timestamp": "2023-02-16T06:18:32.774178625Z",
            "signature": "wpwea8FJ+tRmoWMtfj3EdX2f4Qto8i12Lrl1dUbKF+5yYCLlPuTlhMB+vxcTWvy2HdqKGBqJxjW+FrAHdrv5BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F178D7DB105575C8CBB3B680AF221A08293CB202",
            "timestamp": "2023-02-16T06:18:33.410932761Z",
            "signature": "DjmIC/qghMawh+nVFAqSpkIMkgw6UZKifBQtigLPoQslDD9ROD19hNG/2H6zUYHv96OuwFZrE2y358KSkTzDBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A1FC51BBA3CE4B8A0C690B64D8A84F6382084F01",
            "timestamp": "2023-02-16T06:18:32.753171371Z",
            "signature": "jrR4mFw9oy8PW9XEqo1wj0sHtwQPxuDTuaUrYv9UiMpZjElflE6hDTvCin/E+VEqqwA6c8H8HHXnJ10/1hTuDw=="
          }
          //.....
          //.....
          //.....
        ]
      }
    }
  }
}
```