### 钱包

地址例子：
- Address (Bech32): gx17jhdv3wj0a82h2xt007tag53farz5gvvzw2lnn
- Address (Hex): F4AED645D27F4EABA8CB7BFCBEA2914F462A218C
- Public Key (Bech32): gxpub1addwnpepq05dh673npd5pe2sz457qv2ypcmmelc8xpmv8zespmah02aaxtlpy8xlv5c

#### 创建钱包地址
```shell
plugchaind keys add wallet1   
```
**[!warning]**:
```
Enter keyring passphrase:123123132
Re-enter keyring passphrase:123123123
```
*初次输入的密码 `123123123` 为钱包密码，需要保存好，以下操作使用的密码都是此密码，如忘记此密码，想操作地址，必须把地址的助记词导入到另一个节点或者浏览器使用*

返回信息
```golang
- name: wallet1
  type: local
  address: gx17jhdv3wj0a82h2xt007tag53farz5gvvzw2lnn
  pubkey: gxpub1addwnpepq05dh673npd5pe2sz457qv2ypcmmelc8xpmv8zespmah02aaxtlpy8xlv5c
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

nut blame match license torch away uncover pair nose diagram pepper digital chef pattern traffic garden coral impact wall december renew desert little under
```
**[!danger]**:

`nut blame match license torch away uncover pair nose diagram pepper digital chef pattern traffic garden coral impact wall december renew desert little under` *是地址的助记词， 极为重要，需要手动保存起来*

#### 查看钱包所有地址
```shell
plugchaind keys list  
```
输出内容如下：
```
- name: ds
  type: local
  address: gx1yh9a0trjulf92m9nr4g90c9fwq3kcwnkwjdv9u
  pubkey: gxpub1addwnpepqt7llu0440cytsulyzqgjxjuhtjxk5grwuqraekvrp2p7ud3rxgmsnhezn5
  mnemonic: ""
  threshold: 0
  pubkeys: []
- name: newwallet1
  type: local
  address: gx1k7fzatpnurg9hen3ymug6f7cv2t3dy74ahtuxa
  pubkey: gxpub1addwnpepqw2hz7dvuq5r94w8xezfg65efnhmanklhspv6078qm0sux3z7uh6yktyn2n
  mnemonic: ""
  threshold: 0
  pubkeys: []
```

#### 查看固定地址信息
```shell
plugchaind keys show newwallet1
```
#### 查看地址信息

```shell
plugchaind q account gx17jhdv3wj0a82h2xt007tag53farz5gvvzw2lnn
```

如果返回信息如下：
1. Error: post failed: Post "http://localhost:26657": dial tcp [::1]:26657: connect: connection refused `需要启动node节点服务`
2. Error: rpc error: code = NotFound desc = rpc error: code = NotFound desc = account gx17jhdv3wj0a82h2xt007tag53farz5gvvzw2lnn not found: key not found `说明地址不存在或者地址没有币`

正常返回信息：
```golang
'@type': /cosmos.auth.v1beta1.BaseAccount
account_number: "10"
address: gx17jhdv3wj0a82h2xt007tag53farz5gvvzw2lnn
pub_key: null
sequence: "0"
```
#### 导出地址私钥 
```shell
plugchaind keys export wallet1
```
返回信息：
```
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 4C80A5888778604BA31C06609A0F7BD3
type: secp256k1

E9rkBYnTPYTkzu+iI3/7b7Z6fDuAIoprLF2t824HaXnse67yJXYl+C/nW8d9246I
jNhi7C0vFPnKt3lbxD6ghjAKYgg52u65Eu2FZRk=
=kmrX
-----END TENDERMINT PRIVATE KEY-----
```
#### 导入地址私钥
需要把导出的私钥信息保存到一个文件，例如 wallet1.txt
```shell
plugchaind keys import newmallet ./wallet1.txt
```
导入成功之后，没有任何返回信息，确认是否导入成功，使用 `plugchaind keys show newwallet`

#### 导入地址助记词
```shell
plugchaind keys add newwallet1 --recover
```
命令执行之后，会出现交互内容，要求输入地址的助记词，如下：
```
> Enter your bip39 mnemonic
```
然后输入助记词
```
nut blame match license torch away uncover pair nose diagram pepper digital chef pattern traffic garden coral impact wall december renew desert little under
```
如果节点存在这个地址，输入密码之后会报错，如下：
```
Enter keyring passphrase:
Error: public key already exist in keybase
```
如果节点没有当前地址，导入成功之后，信息如下：
```
plugchaind keys add newwallet1 --recover 
> Enter your bip39 mnemonic
wage sort make special maximum soup wheel high myth tongue skirt half dove connect ready camera giraffe surround dove member harsh inner rural coach
Enter keyring passphrase:

- name: newwallet1
  type: local
  address: gx1k7fzatpnurg9hen3ymug6f7cv2t3dy74ahtuxa
  pubkey: gxpub1addwnpepqw2hz7dvuq5r94w8xezfg65efnhmanklhspv6078qm0sux3z7uh6yktyn2n
  mnemonic: ""
  threshold: 0
  pubkeys: []

```