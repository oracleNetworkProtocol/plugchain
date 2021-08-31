# Slashing

Slashing module can unjail validator previously jailed for downtime

## Available Commands

| Name                                                | Description                                     |
| --------------------------------------------------- | ----------------------------------------------- |
| [unjail](#plugchaind-tx-slashing-unjail)                  | Unjail validator previously jailed for downtime |
| [params](#plugchaind-query-slashing-params)               | Query the current slashing parameters           |
| [signing-info](#plugchaind-query-slashing-signing-info)   | Query a validator's signing information         |
| [signing-infos](#plugchaind-query-slashing-signing-infos) | Query signing information of all validators     |

## plugchaind tx slashing unjail

Unjail validator previously jailed for downtime.

```bash
plugchaind tx slashing unjail [flags]
```

## plugchaind query slashing params

Query the current slashing parameters.

```bash
plugchaind query slashing params  [flags]
```

## plugchaind query slashing signing-info

Query a validator's signing information.

```bash
plugchaind query slashing signing-info [validator-conspub] [flags]
```

## plugchaind query slashing signing-infos

Query signing information of all validators.

```bash
plugchaind query slashing signing-infos [flags]
```
