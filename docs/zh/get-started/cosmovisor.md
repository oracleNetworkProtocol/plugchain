# Cosmosvisor

`cosmovisor` 是 Cosmos SDK 应用程序二进制文件的小型流程管理器，用于监控治理模块以获取传入的链升级提案。 如果它看到一个提案被批准，`cosmovisor` 可以自动下载新的二进制文件，停止当前的二进制文件，从旧的二进制文件切换到新的二进制文件，最后用新的二进制文件重新启动节点。

## 设计

Cosmovisor 旨在用作`Cosmos SDK`应用程序的包装器：

* 它将参数传递给关联的应用程序（由 `DAEMON_NAME` 环境变量配置）。
   运行 `cosmovisor run arg1 arg2 ....` 将运行 `app arg1 arg2 ...`；
* 如果需要，它将通过重新启动和升级来管理应用程序；
* 它是使用环境变量配置的，而不是参数。

*注意：如果应用程序的新版本未设置为运行就地存储迁移，则需要在使用新二进制文件重新启动“cosmovisor”之前手动运行迁移。 因此，我们建议应用程序采用就地存储迁移。*

*注意：如果验证者想要启用自动下载选项（[我们不推荐](#auto-download)）*


## 设置

### 安装

要安装 `cosmovisor`，请运行以下命令：

```sh
cd cosmovisor && make install
```

您可以运行 `cosmovisor --version` 来检查 Cosmovisor 版本（仅适用于 Cosmovisor >=1.0.0）。


### 命令行参数和环境变量

传递给 `cosmovisor` 的第一个参数是 `cosmovisor` 采取的行动。 选项是：

* `help`, `--help`, 或 `-h` - 输出 `cosmovisor` 帮助信息并检查您的 `cosmovisor` 配置。
* `run` - 使用提供的其余参数运行配置的二进制文件。
* `version` 或 `--version` - 输出 `cosmovisor` 版本并使用 `version` 参数运行二进制文件。

传递给`cosmovisor run`的所有参数都将传递给应用程序二进制文件（作为子进程）。 `cosmovisor` 将返回子进程的 `/dev/stdout` 和 `/dev/stderr` 作为它自己的。 出于这个原因，`cosmovisor run` 不能接受除应用程序二进制文件可用的命令行参数之外的任何命令行参数。

*注意：不推荐使用没有操作参数之一的 `cosmovisor`。 为了向后兼容，如果第一个参数不是动作参数，则假定为`运行`。 但是，在未来的版本中可能会删除此回退，因此建议您始终提供 `run`。

`cosmovisor` 从环境变量中读取其配置：

* `DAEMON_HOME` 是保存 `cosmovisor/` 目录的位置，其中包含 genesis 二进制文件、升级二进制文件以及与每个二进制文件关联的任何其他辅助文件（例如 `$HOME/.plugchaind`、`$HOME/. regend`、`$HOME/.simd` 等）。
* `DAEMON_NAME` 是二进制文件本身的名称（例如，`plugchaind`、`regend`、`simd` 等）。
* `DAEMON_ALLOW_DOWNLOAD_BINARIES`（*可选*），如果设置为 `true`，将启用新二进制文件的自动下载（出于安全原因，这适用于完整节点而不是验证器）。 默认情况下，`cosmovisor` 不会自动下载新的二进制文件。
* `DAEMON_RESTART_AFTER_UPGRADE` (*optional*, default = `true`)，如果 `true`，在成功升级后使用相同的命令行参数和标志（但使用新的二进制文件）重新启动子进程。 否则（`false`），`cosmovisor`在升级后停止运行，需要系统管理员手动重启。 注意重启仅在升级后，发生错误后不会自动重启子进程。
* `DAEMON_POLL_INTERVAL`是轮询升级计划文件的间隔长度。 该值可以是数字（以毫秒为单位）或持续时间（例如“1s”）。 默认值：1000 毫秒。
* `DAEMON_BACKUP_DIR` 选项设置自定义备份目录。 如果未设置，则使用 `DAEMON_HOME`。
* `UNSAFE_SKIP_BACKUP`（默认为`false`），如果设置为`true`，直接升级而不执行备份。 否则（`false`，默认）在尝试升级之前备份数据。 默认值 false 很有用，建议在出现故障和需要回滚备份时使用。 我们建议使用默认备份选项`UNSAFE_SKIP_BACKUP=false`。
* `DAEMON_PREUPGRADE_MAX_RETRIES`（默认为 `0`）。 退出状态为 31 后应用程序中调用 pre-upgrade 的最大次数。 在达到最大重试次数后，cosmovisor 升级失败。

### 文件夹布局

`$DAEMON_HOME/cosmovisor` 应该完全属于 `cosmovisor` 及其控制的子进程。 文件夹内容组织如下：

```text
.
├── current -> genesis or upgrades/<name>
├── genesis
│   └── bin
│       └── $DAEMON_NAME
└── upgrades
    └── <name>
        ├── bin
        │   └── $DAEMON_NAME
        └── upgrade-info.json
```

`cosmovisor/` 目录包含应用程序每个版本的子目录（即 `genesis` 或 `upgrades/<name>`）。 在每个子目录中是应用程序二进制文件（即`bin/$DAEMON_NAME`）和与每个二进制文件关联的任何附加辅助文件。 `current` 是指向当前活动目录的符号链接（即 `genesis` 或 `upgrades/<name>`）。 `upgrades/<name>` 中的 `name` 变量是升级模块计划中指定的升级的 URI 编码名称。

请注意，`$DAEMON_HOME/cosmovisor` 仅存储 *应用程序二进制文件*。 `cosmovisor` 二进制文件本身可以存储在任何典型位置（例如 `/usr/local/bin`）。 应用程序将继续将其数据存储在默认数据目录（例如 `$HOME/.plugchaind`）或使用 `--home` 标志指定的数据目录中。 `$DAEMON_HOME` 独立于数据目录，可以设置到任何位置。 如果将 `$DAEMON_HOME` 设置为与数据目录相同的目录，最终将得到如下配置：

```text
.plugchaind
├── config
├── data
└── cosmovisor
```

## 用法

系统管理员负责：

* 安装 `cosmovisor` 二进制文件
* 配置主机的初始化系统（例如`systemd`、`launchd`等）
* 适当设置环境变量
* 手动创建 `genesis` 文件夹
* 手动创建 `upgrades/<name>` 文件夹

`cosmovisor` 会在第一次启动时（即不存在 `current` 链接时）将 `current` 链接设置为指向 `genesis`，然后在正确的时间点处理切换二进制文件，以便系统管理员可以提前几天做好准备 并在升级时放松。

为了支持可下载的二进制文件，每个升级二进制文件都需要打包并通过规范 URL 提供。 此外，包含 genesis 二进制文件和所有可用升级二进制文件的可以打包并提供，以便可以轻松下载从开始同步全节点所需的所有必要二进制文件。

`DAEMON` 特定的代码和操作（例如，tendermint 配置、应用程序数据库、同步块等）都按预期工作。 应用程序二进制文件的指令（例如命令行标志和环境变量）也可以按预期工作。

### 检测升级

`cosmovisor` 正在轮询 `$DAEMON_HOME/data/upgrade-info.json` 文件以获取新的升级说明。 当检测到升级并且区块链达到升级高度时，该文件由 `BeginBlocker` 中的 x/upgrade 模块创建。
应用以下方法检测升级：

* 启动时，`cosmovisor` 对当前正在运行的升级了解不多，除了二进制文件`current/bin/`。 它尝试读取 `current/update-info.json` 文件以获取有关当前升级名称的信息。
* 如果 `cosmovisor/current/upgrade-info.json` 和 `data/upgrade-info.json` 都不存在，则 `cosmovisor` 将等待 `data/upgrade-info.json` 文件触发升级。
* 如果 `cosmovisor/current/upgrade-info.json` 不存在但 `data/upgrade-info.json` 存在，则 `cosmovisor` 假定 `data/upgrade-info.json` 中的任何内容都是有效的 升级请求。 在这种情况下，`cosmovisor` 会立即尝试根据 `data/upgrade-info.json` 中的 `name` 属性进行升级。
* 否则，`cosmovisor` 等待 `upgrade-info.json` 中的更改。 一旦文件中记录了新的升级名称，`cosmovisor` 将触发升级机制。

当升级机制被触发时，`cosmovisor` 将：

1. 如果启用了 `DAEMON_ALLOW_DOWNLOAD_BINARIES`，首先将新的二进制文件自动下载到 `cosmovisor/<name>/bin`（其中 `<name>` 是 `upgrade-info.json:name` 属性）；
2. 更新`current`符号链接指向新目录并将`data/upgrade-info.json`保存到`cosmovisor/current/upgrade-info.json`。

### 自动下载

通常，`cosmovisor` 要求系统管理员在升级之前将所有相关的二进制文件放在磁盘上。 然而，对于不需要这种控制并且想要自动设置的人（也许他们正在同步一个非验证全节点并且想要做很少的维护），还有另一种选择。

**注意：我们不建议使用自动下载**，因为它不会提前验证二进制文件是否可用。 如果下载二进制文件有任何问题，cosmovisor 将停止并且不会重新启动应用程序（这可能导致链停止）。

如果 `DAEMON_ALLOW_DOWNLOAD_BINARIES` 设置为 `true`，并且触发升级时找不到本地二进制文件，`cosmovisor` 将尝试根据 `data 中的 `info` 属性中的说明自行下载和安装二进制文件 /upgrade-info.json` 文件。 这些文件由 x/upgrade 模块构建，包含来自升级“计划”对象的数据。 `Plan` 有一个 info 字段，该字段应具有以下两种有效格式之一来指定下载：

1. 将 os/architecture -> 二进制 URI 映射以 JSON 格式存储在“binaries”键下的升级计划信息字段中。 例如：

```json
{
  "binaries": {
    "linux/amd64":"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"
  }
}
```

您可以一次包含多个二进制文件，以确保多个环境将接收正确的二进制文件：

```json
{
  "binaries": {
    "linux/amd64":"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f",
    "linux/arm64":"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f",
    "darwin/amd64":"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"
    }
}
```

  将其作为提案提交时，请确保没有空格。 使用 `plugchaind` 的示例命令可能如下所示：

```sh
plugchaind tx gov submit-proposal software-upgrade Vega \
--title Vega \
--deposit 100uatom \
--upgrade-height 7368420 \
--upgrade-info '{"binaries":{"linux/amd64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-linux-amd64","linux/arm64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-linux-arm64","darwin/amd64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-darwin-amd64"}}' \
--description "upgrade to Vega" \
--gas 400000 \
--from user \
--chain-id test \
--home test/val2 \
--node tcp://localhost:36657 \
--yes
```

2. 以上述格式存储包含所有信息的文件的链接（例如，如果您想指定大量二进制文件、更改日志信息等而不填满区块链）。 例如：

    ```text
    https://example.com/testnet-1001-info.json?checksum=sha256:deaaa99fda9407c4dbe1d04bd49bab0cc3c1dd76fa392cd55a9425be074af01e
    ```

当触发`cosmovisor`下载新的二进制文件时，`cosmovisor`会解析`"binaries"`字段，用[go-getter]下载新的二进制文件(https://github.com/hashicorp/go-getter) , 并在 `upgrades/<name>` 文件夹中解压缩新的二进制文件，以便它可以像手动安装一样运行。

请注意，要使此机制提供强大的安全保证，所有 URL 都应包含 SHA 256/512 校验。 这确保了不会运行错误的二进制文件，即使有人入侵了服务器或劫持了 DNS。 `go-getter` 将始终确保下载的文件与提供的校验和匹配。 `go-getter` 还将处理解压缩档案到目录中（在这种情况下，下载链接应指向 `bin` 目录中所有数据的 `zip` 文件）。

要在 linux 上正确创建 sha256 校验和，您可以使用 `sha256sum` 实用程序。 例如：

```sh
sha256sum ./testdata/repo/zip_directory/autod.zip
```

结果将如下所示：
`29139e1381b8177aec909fab9a75d11381cab5adf7d3af0c05ff1c9c117743a7`.

如果您希望使用更长的哈希值，您也可以使用 `sha512sum`，如果您希望使用损坏的哈希值，则可以使用 `md5sum`。 无论您选择哪种方式，请确保在 URL 的校验和参数中正确设置哈希算法。

## 示例：SimApp 升级

以下说明使用 Cosmos SDK 源代码附带的模拟应用程序 (`simapp`) 演示了 `cosmovisor`。 以下命令将从 `cosmos-sdk` 存储库中运行。


### 链设置

让我们使用 `v0.44` 版本的 simapp（Cosmos SDK 演示应用）创建一个新链：

```sh
git checkout v0.44.6
make build
```

清理 `~/.simapp`（不要在生产环境中这样做）：

```sh
./build/simd unsafe-reset-all
```

设置应用配置：

```sh
./build/simd config chain-id test
./build/simd config keyring-backend test
./build/simd config broadcast-mode block
```

初始化节点并覆盖任何先前的创世文件（切勿在生产环境中这样做）：
 
<!-- TODO: init 不会从配置中读取链 ID -->

```sh
./build/simd init test --chain-id test --overwrite
```

在 `~/.simapp/config/app.toml` 中将最低 gas 价格设置为 `0stake`：

```sh
minimum-gas-prices = "0stake"
```

为了演示，将 `genesis.json` 中的 `voting_period` 修改为 20 秒 (`20s`) 的缩减时间：

```sh
cat <<< $(jq '.app_state.gov.voting_params.voting_period = "20s"' $HOME/.simapp/config/genesis.json) > $HOME/.simapp/config/genesis.json
```

创建一个验证器，并设置创世交易：

<!-- TODO: add-genesis-account 不会从配置中读取 keyring-backend -->
<!-- TODO: gentx 没有从配置中读取链 ID -->

```sh
./build/simd keys add validator
./build/simd add-genesis-account validator 1000000000stake --keyring-backend test
./build/simd gentx validator 1000000stake --chain-id test
./build/simd collect-gentxs
```

#### 准备 Cosmovisor 并启动链

设置所需的环境变量:

```sh
export DAEMON_NAME=simd
export DAEMON_HOME=$HOME/.simapp
```

设置可选环境变量以触发自动应用重启：

```sh
export DAEMON_RESTART_AFTER_UPGRADE=true
```

为 genesis 二进制文件创建文件夹并复制 `simd` 二进制文件：

```sh
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp ./build/simd $DAEMON_HOME/cosmovisor/genesis/bin
```

现在您可以使用 simapp v0.44 运行 cosmovisor：


```sh
cosmovisor run start
```


#### 更新应用

将应用程序更新到最新版本（例如 v0.45）。

接下来，我们可以添加一个迁移 - 使用 `x/upgrade` [升级计划](https://github.com/cosmos/cosmos-sdk/blob/master/docs/core/upgrade.md) 定义（你 如果您使用的是较旧的 Cosmos SDK 版本，则可能指的是过去的版本）。 在迁移中，我们可以进行任何确定性的状态更改。

构建新版本的 `simd` 二进制文件：

```sh
make build
```

为升级二进制文件创建文件夹并复制“simd”二进制文件：

```sh
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/test1/bin
cp ./build/simd $DAEMON_HOME/cosmovisor/upgrades/test1/bin
```


打开一个新的终端窗口并提交升级提案以及存款和投票（这些命令必须在 20 秒内运行）：

```sh
./build/simd tx gov submit-proposal software-upgrade test1 --title upgrade --description upgrade --upgrade-height 200 --from validator --yes
./build/simd tx gov deposit 1 10000000stake --from validator --yes
./build/simd tx gov vote 1 yes --from validator --yes
```

升级将在高度 200 处自动发生。注意：如果您的测试游戏需要更多时间，您可能需要在上面的代码段中更改升级高度。
