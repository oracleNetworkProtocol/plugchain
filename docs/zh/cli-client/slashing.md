# Slashing

Slashing 模块可以解禁被监禁的验证人

## 可用命令

| 名称                                                | 描述                         |
| --------------------------------------------------- | ---------------------------- |
| [unjail](#plugchaind-tx-slashing-unjail)                  | 解禁被监禁的验证人           |
| [params](#plugchaind-query-slashing-params)               | 查询当前`Slashing`的参数信息 |
| [signing-info](#plugchaind-query-slashing-signing-info)   | 查询验证人的签名信息         |
| [signing-infos](#plugchaind-query-slashing-signing-infos) | 查询所有验证人的签名信息     |

## plugchaind tx slashing unjail

解禁被监禁的验证人。

```bash
plugchaind tx slashing unjail [flags]
```

## plugchaind query slashing params

查询当前`Slashing`的参数信息。

```bash
plugchaind query slashing params [flags]
```

## plugchaind query slashing signing-info

查询验证人的签名信息。

```bash
plugchaind query slashing signing-info [validator-conspub] [flags]
```

## plugchaind query slashing signing-infos

查询所有验证人的签名信息。

```bash
plugchaind query slashing signing-infos [flags]
```
