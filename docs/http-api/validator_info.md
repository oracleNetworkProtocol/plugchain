---
order: 29
---

# Obtain verifier information through address

- Obtain verifier by address

### Request method
`GET`

### URL
`/server/v1/validator/info`

### Request example

```
/server/v1/validator/info?addr=gxvaloper1rkecwy9pjsfd058w0pwa2perquc8xe638h4u5p
```


#### Request parameters

- Not

### Return value
- `code int64 `: 0 or field does not exist is success, others are failures
- `message string` : Response information
- `validators array` : Verifier array
    - `operator_address`: Verifier address
    - `consensus_pubkey`: Consensus public key
    - `jailed`: Incarceration status
    - `status`: Verifier status
    - `tokens`: token
    - `delegator_shares`: Pledged shares
    - `description`: Verifier description
    - `unbonding_height`: Lock block height
    - `unbonding_time`: Lock time
    - `commission`: Consensus information
    - `min_self_delegation`: Minimum pledge amount, unit: uplugcn

#### Return to example
```json5
{
  "validator": {
    "operator_address": "gxvaloper1rkecwy9pjsfd058w0pwa2perquc8xe638h4u5p",
    "consensus_pubkey": {
      "@type": "/cosmos.crypto.ed25519.PubKey",
      "key": "i4unhp2x1nry+iBrSBa4eag4e9rOU//MA4vcwHVJws0="
    },
    "jailed": false,
    "status": "BOND_STATUS_BONDED",
    "tokens": "80216413085275",
    "delegator_shares": "80296709794859.478689896810257297",
    "description": {
      "moniker": "Plug Storm Foundation-1",
      "identity": "",
      "website": "",
      "security_contact": "",
      "details": ""
    },
    "unbonding_height": "3288539",
    "unbonding_time": "2022-04-19T08:50:41.926091423Z",
    "commission": {
      "commission_rates": {
        "rate": "0.100000000000000000",
        "max_rate": "0.200000000000000000",
        "max_change_rate": "0.010000000000000000"
      },
      "update_time": "2021-08-19T08:56:38.706170801Z"
    },
    "min_self_delegation": "1000000"
  }
}
```