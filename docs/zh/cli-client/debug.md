# Debug

## 介绍

一个用于简单调试的工具。

## 可用命令

| 名称                               | 描述                                              |
| ---------------------------------- | ------------------------------------------------- |
| [addr](#plugchaind-debug-addr)           | 转换 hex 和 bech32 地址格式                       |
| [pubkey](#plugchaind-debug-pubkey)       | 解码 hex、base64 或 bech32 格式 ED25519 公钥      |
| [raw-bytes](#plugchaind-debug-raw-bytes) | 将原始字节输出（如[10 21 13 127]）转化为 hex 编码 |

### plugchaind debug addr

```bash
plugchaind debug addr gx1rulhmls7g9cjh239vnkjnw870t5urrutsfwvmc
```

returns

```bash
Address: [31 63 125 254 30 65 113 43 170 37 100 237 41 184 254 122 233 193 143 139]
Address (hex): 1F3F7DFE1E41712BAA2564ED29B8FE7AE9C18F8B
Bech32 Acc: gx1rulhmls7g9cjh239vnkjnw870t5urrutsfwvmc
Bech32 Val: gxvaloper1rulhmls7g9cjh239vnkjnw870t5urrut9cyrxl
```

### plugchaind debug pubkey

下面得到相同的结果：

```bash
plugchaind debug pubkey TZTQnfqOsi89SeoXVnIw+tnFJnr4X8qVC0U8AsEmFk4=
plugchaind debug pubkey 4D94D09DFA8EB22F3D49EA17567230FAD9C5267AF85FCA950B453C02C126164E
  ```

### plugchaind debug raw-bytes

将原始字节输出（如[10 21 13 127]）转化为hex编码

```bash
plugchaind debug raw-bytes <raw-bytes>
plugchaind debug raw-bytes "[10 21 13 127]"
```
