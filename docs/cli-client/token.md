# Token

The Token module is used to manage your [PRC-10](../concepts/token.md#prc-10) assets issued on Plug Chain Hub.

## Available Commands

| Name                                       | Description                                |
| ------------------------------------------ | ------------------------------------------ |
| [issue](#plugchaind-tx-token-issue)              | Issue a new token                          |
| [edit](#plugchaind-tx-token-edit)                | Edit an existing token                     |
| [transfer](#plugchaind-tx-token-transfer)        | Transfer the ownership of a token          |
| [mint](#plugchaind-tx-token-mint)                | Mint tokens to a specified address         |
| [burn](#plugchaind-tx-token-burn)                | Burn some tokens                           |
| [token](#plugchaind-query-token-token)           | Query a token by symbol                    |
| [tokens](#plugchaind-query-token-tokens)         | Query tokens by owner                      |
| [fee](#plugchaind-query-token-fee)               | Query the token related fees               |
| [params](#plugchaind-query-token-params)         | Query the token related params             |

## plugchaind tx token issue

Issue a new token

```bash
plugchaind tx token issue [flags]
```

**Flags:**

| Name, shorthand  | Type    | Required | Default       | Description                                                                                                                    |
| ---------------- | ------- | -------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| --name           | string  | Yes      |               | Name of the newly issued token, limited to 32 unicode characters, e.g. "Plug Chain Network"                                          |
| --symbol         | string  | Yes      |               | The symbol of the token, length between 3 and 8, alphanumeric characters beginning with alpha, case insensitive                |
| --initial-supply | uint64  | Yes      |               | The initial supply of this token                                    |
| --max-supply     | uint64  |          | 100000000000 | The hard cap of this token, total supply can not exceed max supply. The amount before boosting should not exceed 100 billion. |
| --min-unit       | string  |          |               | The alias of minimum (Can be the same as `--symbol`) ,length between 3 and 8, alphanumeric characters beginning with alpha, case insensitive uint                                                                                                      |
| --scale          | uint8   | Yes      |               | A token can have a maximum of 8 digits of decimal                                                                             |
| --mintable       | boolean |          | false         | Whether this token could be minted(increased) after the initial issuing                                                        |

### Issue a token

```bash
plugchaind tx token issue \
    --name="Test Token" \
    --symbol="token" \
    --min-unit="token" \
    --scale=6 \
    --initial-supply=1000000000 \
    --max-supply=100000000000 \
    --mintable=true \
    --from=<key-name> \
    --chain-id=<chain-id> \
    --fees=<fee>
```

### Send tokens

You can send any tokens you have just like [sending plug](./bank.md#plugchaind-tx-bank-send)

#### Send tokens

```bash
plugchaind tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

## plugchaind tx token edit

Edit an existing token

```bash
plugchaind tx token edit [symbol] [flags]
```

**Flags:**

| Name         | Type   | Required | Default | Description                                       |
| ------------ | ------ | -------- | ------- | ------------------------------------------------- |
| --name       | string |          |         | The token name, e.g. Plug Chain Network                 |
| --max-supply | uint64 |          | 0       | The max supply of the token                       |

`max-supply` should not be less than the current total supply

### Edit Token

```bash
plugchaind tx token edit <symbol> --name="Cat Token" --max-supply=100000000000 --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token transfer

Transfer the ownership of a token

```bash
plugchaind tx token transfer [symbol] [flags]
```

**Flags:**

| Name | Type   | Required | Default | Description           |
| ---- | ------ | -------- | ------- | --------------------- |
| --to | string | Yes      |         | The new owner address |

### Transfer Token Owner

```bash
plugchaind tx token transfer <symbol> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token mint

Mint tokens to a specified address

```bash
plugchaind tx token mint [symbol] [flags]
```

**Flags:**

| Name     | Type   | Required | Default | Description                                                             |
| -------- | ------ | -------- | ------- | ----------------------------------------------------------------------- |
| --to     | string |          |         | Address to which the token will be minted, default to the owner address |
| --amount | uint64 | Yes      | 0       | Amount of the tokens to be minted                                       |

### Mint Token

```bash
plugchaind tx token mint <symbol> --amount=<amount> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token burn

Burn some tokens

```bash
plugchaind tx token burn [symbol] [flags]
```

**Flags:**

| Name     | Type   | Required | Default | Description                   |
| -------- | ------ | -------- | ------- | ----------------------------- |
| --amount | uint64 | Yes      | 0       | Amount of the tokens to burnt |

### Burn Token

```bash
plugchaind tx token burn <symbol> --amount=<amount> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind query token token

Query a token by symbol

```bash
plugchaind query token token [denom] [flags]
```

### Query a token

```bash
plugchaind query token token <denom>
```

## plugchaind query token tokens

Query tokens by the owner which is optional

```bash
plugchaind query token tokens [owner] [flags]
```

### Query all tokens

```bash
plugchaind query token tokens
```

### Query tokens with the specified owner

```bash
plugchaind query token tokens <owner>
```

## plugchaind query token fee

Query the token related fees, including token issuance and minting

```bash
plugchaind query token fee [symbol] [flags]
```

### Query fees of issuing and minting a token

```bash
plugchaind query token fee kitty
```

## plugchaind query token params

Query token module params

```bash
plugchaind query token params [flags]
```

### Query token module params

```bash
plugchaind query token params
```