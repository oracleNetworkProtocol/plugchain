---
order: 1
---

# Hardhat：部署智能合约

了解如何使用Hardhat环境部署一个简单的基于坚固性的智能合约到Plug Chain

[Hardhat](https://hardhat.org/) 是一个灵活的开发环境，用于构建基于以太坊的智能合约。它的设计考虑了集成性和可扩展性

## 准备工作

- [安装](./../get-started/install.md) 
- [加入主网](./../get-started/mainnet.md) 

## 流程

在继续之前，您需要安装Node。js（我们将使用v16.x）和npm包管理器。你可以直接从 [Node.js](https://nodejs.org/en/download/) 或者在你的终端:


 Ubuntu

```bash
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -

sudo apt install -y nodejs
```

 MacOS

```bash
# 你可以用 homebrew (https://docs.brew.sh/Installation)
$ brew install node

# 或者使用 nvm (https://github.com/nvm-sh/nvm)
$ nvm install node
```


您可以通过查询每个软件包的版本来验证是否正确安装了所有软件包:

```bash
$ node -v
...

$ npm -v
...
```

> tip
如果您还没有安装，那么如果您计划在本地部署智能合约，您还需要安装Plug Chain。检查这个 [文档]](./../get-started/install.md) 详细说明请参阅.

## 创建 Hardhat 项目

若要创建新项目，请导航到项目目录并运行:

```bash
$ npx hardhat

888    888                      888 888               888
888    888                      888 888               888
888    888                      888 888               888
8888888888  8888b.  888d888 .d88888 88888b.   8888b.  888888
888    888     "88b 888P"  d88" 888 888 "88b     "88b 888
888    888 .d888888 888    888  888 888  888 .d888888 888
888    888 888  888 888    Y88b 888 888  888 888  888 Y88b.
888    888 "Y888888 888     "Y88888 888  888 "Y888888  "Y888

Welcome to Hardhat v2.0.8

? What do you want to do? …
❯ Create a sample project
  Create an empty hardhat.config.js
```

根据提示，应该在目录中创建一个新的项目结构。查阅 [Hardhat config page](https://hardhat.org/config/) 中指定的配置选项列表 `hardhat.config.js`. 最重要的是，您应该设置 `defaultNetwork` 入口指向您想要的JSON-RPC网络:

Local Node

```javascript
module.exports = {
  defaultNetwork: "local",
  networks: {
    hardhat: {
    },
    local: {
      url: "http://localhost:8545/",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}
```

Testnet

```javascript
module.exports = {
  defaultNetwork: "testnet",
  networks: {
    hardhat: {
    },
    testnet: {
      url: "https://evmos-archive-testnet.api.bdnodes.net:8545/",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}
```


为了确保您的目标是正确的网络，您可以从默认的网络提供商查询可用的帐户列表:

```bash
$ npx hardhat accounts
0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
0x70997970C51812dc3A010C7d01b50e0d17dc79C8
0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
0x90F79bf6EB2c4f870365E785982E1f101E93b906
...
```

## 部署智能合约

您将看到在下面已经提供了一个默认的智能合约，它是用Solidity编写的 `contracts/Greeter.sol`:

```javascript
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Greeter {
    string private greeting;

    constructor(string memory _greeting) {
        console.log("Deploying a Greeter with greeting:", _greeting);
        greeting = _greeting;
    }

    function greet() public view returns (string memory) {
        return greeting;
    }

    function setGreeting(string memory _greeting) public {
        console.log("Changing greeting from '%s' to '%s'", greeting, _greeting);
        greeting = _greeting;
    }
}
```

这个契约允许你设置和查询一个字符串`greeting`。Hardhat还提供了一个脚本来部署智能合约到目标网络;这可以通过以下命令调用，目标是您的默认网络:

```bash
npx hardhat run scripts/sample-script.js
```

Hardhat还允许您通过 `--network <your-network>` 标签:

Local Node

```bash
npx hardhat run --network {{ $themeConfig.project.rpc_url_local }} scripts/sample-script.js
```

Testnet

```bash
npx hardhat run --network {{ $themeConfig.project.rpc_url_testnet }} scripts/sample-script.js
```


最后，尝试运行Hardhat测试:

```bash
$ npx hardhat test
Compiling 1 file with 0.8.4
Compilation finished successfully


  Greeter
Deploying a Greeter with greeting: Hello, world!
Changing greeting from 'Hello, world!' to 'Hola, mundo!'
    ✓ Should return the new greeting once it's changed (803ms)


  1 passing (805ms)
```
