---
order: 1
---

# 链 ID

了解 Plug Chain 链 ID 格式 

## 主网链 ID


|姓名 |链ID |标识符 | EIP155 编号 |版本号 |
|-----|-------------|--------| ---|-----------------|
| Plug Chain Mainnet | `plugchain_520-1` | `plugchain` | `520` | `1` |
| Plug Chain Testnet | `plugchain_521-1` | `plugchain` | `521` | `1` |


## 链标识符

每个链都必须有一个唯一的标识符或“chain-id”。 Tendermint 要求每个应用程序
在 [genesis.json 字段](https://docs.tendermint.com/master/spec/core/genesis.html#genesis-fields) 中定义自己的 `chain-id`。然而，为了同时符合 EIP155 和 Cosmos 的链升级标准，兼容 Plug Chain 的链必须为其链标识符实现特殊的结构。

## 结构

Plug Chain 链 ID 包含 3 个主要组件

- **标识符**：定义应用程序名称的非结构化字符串。
- **EIP155 编号**：不可变 [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `CHAIN_ID` 定义了重放攻击保护编号。
- **版本号**：是链当前正在运行的版本号（始终为正）。
每次链升级或分叉时，这个数字**必须**递增，以避免网络或共识错误。

### 格式

在 genesis 中指定和 Plug Chain 兼容的链 ID 的格式如下：

```重击
{标识符}_{EIP155}-{版本}
```

下表提供了一个示例，其中第二行对应于第一行的升级：

|链ID |标识符 | EIP155 编号 |版本号 |
|----------------|------------|---------------|----------|
| `plugchain_520-1` | plugchain | 520 | 1 |
| `plugchain_520-2` | plugchain | 520 | 2 |
| `...` | ... | ... | ... |
| `plugchain_520-N` | plugchain | 520 | N |