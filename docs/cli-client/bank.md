# Bank

Bank module allows you to manage assets in your local accounts.

## Available Commands

| Name                                  | Description                                            |
| ------------------------------------- | ------------------------------------------------------ |
| [balances](#plugchaind-query-bank-balances) | Query for account balances by address                  |
| [total](#plugchaind-query-bank-total)       | Query the total supply of coins of the chain           |
| [send](#plugchaind-tx-bank-send)            | Create and/or sign and broadcast a MsgSend transaction |

## plugchaind query bank balances

Query the total balance of an account or of a specific denomination.

```bash
plugchaind query bank balances [address] [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description                                                |
| --------------- | ------ | -------- | ------- | ---------------------------------------------------------- |
| -h, --help      |        |          |         | Help for coin-type                                         |
| --denom         | string |          |         | The specific balance denomination to query for             |
| --count-total   |        |          |         | Count total number of records in all balances to query for |

### plugchaind query bank total

Query total supply of coins that are held by accounts in the chain.

```bash
plugchaind query bank total [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description                                    |
| --------------- | ------ | -------- | ------- | ---------------------------------------------- |
| -h, --help      |        |          |         | Help for coin-type                             |
| --denom         | string |          |         | The specific balance denomination to query for |

## plugchaind tx bank send

Sending tokens to another address, this command includes `generate`, `sign` and `broadcast` steps.

```bash
plugchaind tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

**Flags:**

| Name, shorthand | Type | Required | Default | Description       |
| --------------- | ---- | -------- | ------- | ----------------- |
| -h, --help      |      |          |         | Help for balances |
