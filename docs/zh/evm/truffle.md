<!--
order: 3
-->

# Truffle: 部署智能合约

 学习如何使用Truffle环境部署一个简单的基于solid的智能合约到Plug Chain {synopsis}

## Pre-requisite Readings

- [安装](./../get-started/install.md) {prereq}
- [加入主网](./../get-started/mainnet.md) {prereq}

[Truffle](https://www.trufflesuite.com/truffle) 是用于部署和管理的开发框架
[Solidity](https://github.com/ethereum/solidity) 智能合约.

## 安装依赖包

首先，在您的机器上安装最新的Truffle版本.

```bash
yarn install truffle -g
```

::: tip
如果您还没有安装，那么如果您计划在本地部署智能合约，您还需要安装Plug Chain。检查这个 [文档](./../get-started/install.md) 详细说明请参阅.
:::

## 创建 Truffle 项目

在这一步中，我们将创建一个简单的智能合约。如果您已经有了自己的编译合约，可以跳过此步骤。

创建一个存放合约的新目录并初始化它:

```console
mkdir plugchain-truffle
cd plugchain-truffle
```

初始化Truffle:

```bash
truffle init
```

创建 `contracts/Counter.sol` 合约:

```javascript
pragma solidity >=0.7.0 <0.9.0;

contract Counter {
  uint256 counter = 0;

  function add() public {
    counter++;
  }

  function subtract() public {
    counter--;
  }

  function getCounter() public view returns (uint256) {
    return counter;
  }
}
```

使用 `compile` 命令:

```bash
truffle compile
```

创建 `test/counter_test.js` 使用Javascript包含以下测试 [Mocha](https://mochajs.org/):

```javascript
const Counter = artifacts.require("Counter")

contract('Counter', accounts => {
  const from = accounts[0]
  let counter

  before(async() => {
    counter = await Counter.new()
  })

  it('should add', async() => {
    await counter.add()
    let count = await counter.getCounter()
    assert(count == 1, `count was ${count}`)
  })
})
```

## Truffle configuration

打开`truffle-config.js`，取消`networks`中的`development`部分的注释:

```javascript
    development: {
      host: "127.0.0.1",     // Localhost (default: none)
      port: 8545,            // Standard Ethereum port (default: none)
      network_id: "*",       // Any network (default: none)
    },
```

这将允许您的合约连接到您的Plug Chain本地节点.

## 启动节点

::: tip
有关如何运行节点的详细信息，请参阅 [quickstart guide](./../get-started/mainnet.md).
:::

## 部署合约

在Truffle终端，迁移合约使用:

```bash
truffle migrate --network development
```

您应该在Plug Chain守护程序Terminal选项卡中看到每个事务的传入部署日志(一个用来部署 `Migrations.sol` 另一个用来部署 `Counter.sol`).

```bash
$ I[2020-07-15|17:35:59.934] Added good transaction                       module=mempool tx=22245B935689918D332F58E82690F02073F0453D54D5944B6D64AAF1F21974E2 res="&{CheckTx:log:\"[]\" gas_wanted:6721975 }" height=3 total=1
I[2020-07-15|17:36:02.065] Executed block                               module=state height=4 validTxs=1 invalidTxs=0
I[2020-07-15|17:36:02.068] Committed state                              module=state height=4 txs=1 appHash=76BA85365F10A59FE24ADCA87544191C2D72B9FB5630466C5B71E878F9C0A111
I[2020-07-15|17:36:02.981] Added good transaction                       module=mempool tx=84516B4588CBB21E6D562A6A295F1F8876076A0CFF2EF1B0EC670AD8D8BB5425 res="&{CheckTx:log:\"[]\" gas_wanted:6721975 }" height=4 total=1
```

## 运行 Truffle 测试

现在，您可以使用`test`命令使用插头链节点运行Truffle 测试:

```bash
$ truffle test --network development

Using network 'development'.


Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.



  Contract: Counter
    ✓ should add (5036ms)


  1 passing (10s)
```
