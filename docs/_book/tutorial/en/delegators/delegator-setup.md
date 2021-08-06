### Delegate token

#### Query all verifiers of the node
```
plugchaind q staking validators [flags]
```
#### Screen out your favorite verifiers and perform delegation
```
plugchaind tx staking delegate gxvaloperxxxxxxxxxxxxxxxxxxxxx 20000plug--from gxxxxxxxxxxxxxxxxxxxxx --chain-id plugchain
```
#### Modify the verifier you delegate
```
plugchaind tx staking redelegate gxvaloperxxxxxxxxxxxxxxxxxxxx gxvaloperxxxxxxxxxxxxxxxxxxx 100000plug--from gxxxxxxxxxxxxxxxxx --chain-id plugchain
```