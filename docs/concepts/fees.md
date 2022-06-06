---
order: 6
---

# Gas and Fees

Learn the difference between "Gas" and "Fees" in PVM and Cosmos.

## Introducing `Gas` and `Fees`

In the Cosmos SDK, `gas` is a special unit used to track resource consumption during execution. `gas` is normally consumed when reading and writing to storage, but it can also be consumed if expensive computations are required. It has two main purposes:

* Make sure the block doesn't consume too many resources and will be finalized.
* Prevent spam and abuse by end users. To this end, `gas` consumed during the execution of `Msg` is usually priced, resulting in `fee` (`fees = gas * gas-prices`). `fees` must generally be paid by the sender executing the `Msg`. Note that the Cosmos SDK does not enforce "gas" pricing by default, as there may be other ways to prevent spam (such as bandwidth schemes). Nonetheless, most applications will implement a "fee" mechanism to prevent spam.


## Cosmos SDK `Fees`

* Verify that the gas price offered by the transaction is higher than the local `min-gas-prices`. `min-gas-prices` is a local parameter for each full node and is used during `CheckTx` to discard transactions that do not offer minimum fees. This ensures that the mempool is not spammed by spam transactions.
* EQ: `fees = gas * gas-price`
* For example: `gas=200000` (default), `gas-price=0.0001`, `fees=20uplugcn`


## PVM `Fees`


The main difference between PVM and Cosmos state transitions is that PVM uses a [gas table](https://github.com/ethereum/go-ethereum/blob/master/params/protocol_params.go) for each OPCODE, while Cosmos Use "GasConfig" to charge gas for each CRUD operation by setting the unit of access to the database and the cost per byte.

The gas usage of PVM and EVM is calculated in the same way, but the PVM `gas` usage is different from the EVM `gas` usage process. When the gas is enough to pay the transaction, the excess gas will be returned. PVM fee gas will not be returned.

* PVM transactions must first meet the minimum handling fee required by `fees` greater than Cosmos SDK` (currently the minimum handling fee is 0.0001*200000=20uplugcn)`
* The gas provided by the transaction should be greater than or equal to the gas sum calculated according to the [gas table](https://github.com/ethereum/go-ethereum/blob/master/params) transaction bytes /protocol_params.go)
* Minimum gasPrice is `7`
* EQ: fees = `math.Ceil( (gas*gasPrice) / 1000 )` (Proposal [v1.5.0](https://www.plugchain.network/v2/communityDetail?id=9) algorithm after upgrade)
* For example: `gas=654321`,`gasPrice=7`,`fees=4581uplugcn`