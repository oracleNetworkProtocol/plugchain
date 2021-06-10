### Delegate token

#### Query all verifiers of the node
```
plugchaind q staking validators [flags]
```
#### Screen out your favorite verifiers and perform delegation
```
plugchaind tx staking delegate onpvaloperxxxxxxxxxxxxxxxxxxxxx 20000onp --from onpxxxxxxxxxxxxxxxxxxxx --chain-id plugchain
```
#### Modify the verifier you delegate
```
plugchaind tx staking redelegate onpvaloperxxxxxxxxxxxxxxxxxxxx onpvaloperxxxxxxxxxxxxxxxxxxx 100000onp --from onpxxxxxxxxxxxxxxxx --chain-id plugchain
```