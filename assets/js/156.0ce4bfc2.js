(window.webpackJsonp=window.webpackJsonp||[]).push([[156],{477:function(s,a,e){"use strict";e.r(a);var t=e(15),v=Object(t.a)({},(function(){var s=this,a=s._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[a("h1",{attrs:{id:"cosmosvisor"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#cosmosvisor"}},[s._v("#")]),s._v(" Cosmosvisor")]),s._v(" "),a("p",[a("code",[s._v("cosmovisor")]),s._v(" 是 Cosmos SDK 应用程序二进制文件的小型流程管理器，用于监控治理模块以获取传入的链升级提案。 如果它看到一个提案被批准，"),a("code",[s._v("cosmovisor")]),s._v(" 可以自动下载新的二进制文件，停止当前的二进制文件，从旧的二进制文件切换到新的二进制文件，最后用新的二进制文件重新启动节点。")]),s._v(" "),a("h2",{attrs:{id:"设计"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#设计"}},[s._v("#")]),s._v(" 设计")]),s._v(" "),a("p",[s._v("Cosmovisor 旨在用作"),a("code",[s._v("Cosmos SDK")]),s._v("应用程序的包装器：")]),s._v(" "),a("ul",[a("li",[s._v("它将参数传递给关联的应用程序（由 "),a("code",[s._v("DAEMON_NAME")]),s._v(" 环境变量配置）。\n运行 "),a("code",[s._v("cosmovisor run arg1 arg2 ....")]),s._v(" 将运行 "),a("code",[s._v("app arg1 arg2 ...")]),s._v("；")]),s._v(" "),a("li",[s._v("如果需要，它将通过重新启动和升级来管理应用程序；")]),s._v(" "),a("li",[s._v("它是使用环境变量配置的，而不是参数。")])]),s._v(" "),a("p",[a("em",[s._v("注意：如果应用程序的新版本未设置为运行就地存储迁移，则需要在使用新二进制文件重新启动“cosmovisor”之前手动运行迁移。 因此，我们建议应用程序采用就地存储迁移。")])]),s._v(" "),a("p",[a("em",[s._v("注意：如果验证者想要启用自动下载选项（"),a("a",{attrs:{href:"#auto-download"}},[s._v("我们不推荐")]),s._v("）")])]),s._v(" "),a("h2",{attrs:{id:"设置"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#设置"}},[s._v("#")]),s._v(" 设置")]),s._v(" "),a("h3",{attrs:{id:"安装"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#安装"}},[s._v("#")]),s._v(" 安装")]),s._v(" "),a("p",[s._v("要安装 "),a("code",[s._v("cosmovisor")]),s._v("，请运行以下命令：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" cosmovisor "),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("&&")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v("\n")])])]),a("p",[s._v("您可以运行 "),a("code",[s._v("cosmovisor --version")]),s._v(" 来检查 Cosmovisor 版本（仅适用于 Cosmovisor >=1.0.0）。")]),s._v(" "),a("h3",{attrs:{id:"命令行参数和环境变量"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#命令行参数和环境变量"}},[s._v("#")]),s._v(" 命令行参数和环境变量")]),s._v(" "),a("p",[s._v("传递给 "),a("code",[s._v("cosmovisor")]),s._v(" 的第一个参数是 "),a("code",[s._v("cosmovisor")]),s._v(" 采取的行动。 选项是：")]),s._v(" "),a("ul",[a("li",[a("code",[s._v("help")]),s._v(", "),a("code",[s._v("--help")]),s._v(", 或 "),a("code",[s._v("-h")]),s._v(" - 输出 "),a("code",[s._v("cosmovisor")]),s._v(" 帮助信息并检查您的 "),a("code",[s._v("cosmovisor")]),s._v(" 配置。")]),s._v(" "),a("li",[a("code",[s._v("run")]),s._v(" - 使用提供的其余参数运行配置的二进制文件。")]),s._v(" "),a("li",[a("code",[s._v("version")]),s._v(" 或 "),a("code",[s._v("--version")]),s._v(" - 输出 "),a("code",[s._v("cosmovisor")]),s._v(" 版本并使用 "),a("code",[s._v("version")]),s._v(" 参数运行二进制文件。")])]),s._v(" "),a("p",[s._v("传递给"),a("code",[s._v("cosmovisor run")]),s._v("的所有参数都将传递给应用程序二进制文件（作为子进程）。 "),a("code",[s._v("cosmovisor")]),s._v(" 将返回子进程的 "),a("code",[s._v("/dev/stdout")]),s._v(" 和 "),a("code",[s._v("/dev/stderr")]),s._v(" 作为它自己的。 出于这个原因，"),a("code",[s._v("cosmovisor run")]),s._v(" 不能接受除应用程序二进制文件可用的命令行参数之外的任何命令行参数。")]),s._v(" "),a("p",[s._v("*注意：不推荐使用没有操作参数之一的 "),a("code",[s._v("cosmovisor")]),s._v("。 为了向后兼容，如果第一个参数不是动作参数，则假定为"),a("code",[s._v("运行")]),s._v("。 但是，在未来的版本中可能会删除此回退，因此建议您始终提供 "),a("code",[s._v("run")]),s._v("。")]),s._v(" "),a("p",[a("code",[s._v("cosmovisor")]),s._v(" 从环境变量中读取其配置：")]),s._v(" "),a("ul",[a("li",[a("code",[s._v("DAEMON_HOME")]),s._v(" 是保存 "),a("code",[s._v("cosmovisor/")]),s._v(" 目录的位置，其中包含 genesis 二进制文件、升级二进制文件以及与每个二进制文件关联的任何其他辅助文件（例如 "),a("code",[s._v("$HOME/.plugchaind")]),s._v("、"),a("code",[s._v("$HOME/. regend")]),s._v("、"),a("code",[s._v("$HOME/.simd")]),s._v(" 等）。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_NAME")]),s._v(" 是二进制文件本身的名称（例如，"),a("code",[s._v("plugchaind")]),s._v("、"),a("code",[s._v("regend")]),s._v("、"),a("code",[s._v("simd")]),s._v(" 等）。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_ALLOW_DOWNLOAD_BINARIES")]),s._v("（"),a("em",[s._v("可选")]),s._v("），如果设置为 "),a("code",[s._v("true")]),s._v("，将启用新二进制文件的自动下载（出于安全原因，这适用于完整节点而不是验证器）。 默认情况下，"),a("code",[s._v("cosmovisor")]),s._v(" 不会自动下载新的二进制文件。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_RESTART_AFTER_UPGRADE")]),s._v(" ("),a("em",[s._v("optional")]),s._v(", default = "),a("code",[s._v("true")]),s._v(")，如果 "),a("code",[s._v("true")]),s._v("，在成功升级后使用相同的命令行参数和标志（但使用新的二进制文件）重新启动子进程。 否则（"),a("code",[s._v("false")]),s._v("），"),a("code",[s._v("cosmovisor")]),s._v("在升级后停止运行，需要系统管理员手动重启。 注意重启仅在升级后，发生错误后不会自动重启子进程。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_POLL_INTERVAL")]),s._v("是轮询升级计划文件的间隔长度。 该值可以是数字（以毫秒为单位）或持续时间（例如“1s”）。 默认值：1000 毫秒。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_BACKUP_DIR")]),s._v(" 选项设置自定义备份目录。 如果未设置，则使用 "),a("code",[s._v("DAEMON_HOME")]),s._v("。")]),s._v(" "),a("li",[a("code",[s._v("UNSAFE_SKIP_BACKUP")]),s._v("（默认为"),a("code",[s._v("false")]),s._v("），如果设置为"),a("code",[s._v("true")]),s._v("，直接升级而不执行备份。 否则（"),a("code",[s._v("false")]),s._v("，默认）在尝试升级之前备份数据。 默认值 false 很有用，建议在出现故障和需要回滚备份时使用。 我们建议使用默认备份选项"),a("code",[s._v("UNSAFE_SKIP_BACKUP=false")]),s._v("。")]),s._v(" "),a("li",[a("code",[s._v("DAEMON_PREUPGRADE_MAX_RETRIES")]),s._v("（默认为 "),a("code",[s._v("0")]),s._v("）。 退出状态为 31 后应用程序中调用 pre-upgrade 的最大次数。 在达到最大重试次数后，cosmovisor 升级失败。")])]),s._v(" "),a("h3",{attrs:{id:"文件夹布局"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#文件夹布局"}},[s._v("#")]),s._v(" 文件夹布局")]),s._v(" "),a("p",[a("code",[s._v("$DAEMON_HOME/cosmovisor")]),s._v(" 应该完全属于 "),a("code",[s._v("cosmovisor")]),s._v(" 及其控制的子进程。 文件夹内容组织如下：")]),s._v(" "),a("div",{staticClass:"language-text extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[s._v(".\n├── current -> genesis or upgrades/&lt;name>\n├── genesis\n│   └── bin\n│       └── $DAEMON_NAME\n└── upgrades\n    └── &lt;name>\n        ├── bin\n        │   └── $DAEMON_NAME\n        └── upgrade-info.json\n")])])]),a("p",[a("code",[s._v("cosmovisor/")]),s._v(" 目录包含应用程序每个版本的子目录（即 "),a("code",[s._v("genesis")]),s._v(" 或 "),a("code",[s._v("upgrades/<name>")]),s._v("）。 在每个子目录中是应用程序二进制文件（即"),a("code",[s._v("bin/$DAEMON_NAME")]),s._v("）和与每个二进制文件关联的任何附加辅助文件。 "),a("code",[s._v("current")]),s._v(" 是指向当前活动目录的符号链接（即 "),a("code",[s._v("genesis")]),s._v(" 或 "),a("code",[s._v("upgrades/<name>")]),s._v("）。 "),a("code",[s._v("upgrades/<name>")]),s._v(" 中的 "),a("code",[s._v("name")]),s._v(" 变量是升级模块计划中指定的升级的 URI 编码名称。")]),s._v(" "),a("p",[s._v("请注意，"),a("code",[s._v("$DAEMON_HOME/cosmovisor")]),s._v(" 仅存储 "),a("em",[s._v("应用程序二进制文件")]),s._v("。 "),a("code",[s._v("cosmovisor")]),s._v(" 二进制文件本身可以存储在任何典型位置（例如 "),a("code",[s._v("/usr/local/bin")]),s._v("）。 应用程序将继续将其数据存储在默认数据目录（例如 "),a("code",[s._v("$HOME/.plugchaind")]),s._v("）或使用 "),a("code",[s._v("--home")]),s._v(" 标志指定的数据目录中。 "),a("code",[s._v("$DAEMON_HOME")]),s._v(" 独立于数据目录，可以设置到任何位置。 如果将 "),a("code",[s._v("$DAEMON_HOME")]),s._v(" 设置为与数据目录相同的目录，最终将得到如下配置：")]),s._v(" "),a("div",{staticClass:"language-text extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[s._v(".plugchaind\n├── config\n├── data\n└── cosmovisor\n")])])]),a("h2",{attrs:{id:"用法"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#用法"}},[s._v("#")]),s._v(" 用法")]),s._v(" "),a("p",[s._v("系统管理员负责：")]),s._v(" "),a("ul",[a("li",[s._v("安装 "),a("code",[s._v("cosmovisor")]),s._v(" 二进制文件")]),s._v(" "),a("li",[s._v("配置主机的初始化系统（例如"),a("code",[s._v("systemd")]),s._v("、"),a("code",[s._v("launchd")]),s._v("等）")]),s._v(" "),a("li",[s._v("适当设置环境变量")]),s._v(" "),a("li",[s._v("手动创建 "),a("code",[s._v("genesis")]),s._v(" 文件夹")]),s._v(" "),a("li",[s._v("手动创建 "),a("code",[s._v("upgrades/<name>")]),s._v(" 文件夹")])]),s._v(" "),a("p",[a("code",[s._v("cosmovisor")]),s._v(" 会在第一次启动时（即不存在 "),a("code",[s._v("current")]),s._v(" 链接时）将 "),a("code",[s._v("current")]),s._v(" 链接设置为指向 "),a("code",[s._v("genesis")]),s._v("，然后在正确的时间点处理切换二进制文件，以便系统管理员可以提前几天做好准备 并在升级时放松。")]),s._v(" "),a("p",[s._v("为了支持可下载的二进制文件，每个升级二进制文件都需要打包并通过规范 URL 提供。 此外，包含 genesis 二进制文件和所有可用升级二进制文件的可以打包并提供，以便可以轻松下载从开始同步全节点所需的所有必要二进制文件。")]),s._v(" "),a("p",[a("code",[s._v("DAEMON")]),s._v(" 特定的代码和操作（例如，tendermint 配置、应用程序数据库、同步块等）都按预期工作。 应用程序二进制文件的指令（例如命令行标志和环境变量）也可以按预期工作。")]),s._v(" "),a("h3",{attrs:{id:"检测升级"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#检测升级"}},[s._v("#")]),s._v(" 检测升级")]),s._v(" "),a("p",[a("code",[s._v("cosmovisor")]),s._v(" 正在轮询 "),a("code",[s._v("$DAEMON_HOME/data/upgrade-info.json")]),s._v(" 文件以获取新的升级说明。 当检测到升级并且区块链达到升级高度时，该文件由 "),a("code",[s._v("BeginBlocker")]),s._v(" 中的 x/upgrade 模块创建。\n应用以下方法检测升级：")]),s._v(" "),a("ul",[a("li",[s._v("启动时，"),a("code",[s._v("cosmovisor")]),s._v(" 对当前正在运行的升级了解不多，除了二进制文件"),a("code",[s._v("current/bin/")]),s._v("。 它尝试读取 "),a("code",[s._v("current/update-info.json")]),s._v(" 文件以获取有关当前升级名称的信息。")]),s._v(" "),a("li",[s._v("如果 "),a("code",[s._v("cosmovisor/current/upgrade-info.json")]),s._v(" 和 "),a("code",[s._v("data/upgrade-info.json")]),s._v(" 都不存在，则 "),a("code",[s._v("cosmovisor")]),s._v(" 将等待 "),a("code",[s._v("data/upgrade-info.json")]),s._v(" 文件触发升级。")]),s._v(" "),a("li",[s._v("如果 "),a("code",[s._v("cosmovisor/current/upgrade-info.json")]),s._v(" 不存在但 "),a("code",[s._v("data/upgrade-info.json")]),s._v(" 存在，则 "),a("code",[s._v("cosmovisor")]),s._v(" 假定 "),a("code",[s._v("data/upgrade-info.json")]),s._v(" 中的任何内容都是有效的 升级请求。 在这种情况下，"),a("code",[s._v("cosmovisor")]),s._v(" 会立即尝试根据 "),a("code",[s._v("data/upgrade-info.json")]),s._v(" 中的 "),a("code",[s._v("name")]),s._v(" 属性进行升级。")]),s._v(" "),a("li",[s._v("否则，"),a("code",[s._v("cosmovisor")]),s._v(" 等待 "),a("code",[s._v("upgrade-info.json")]),s._v(" 中的更改。 一旦文件中记录了新的升级名称，"),a("code",[s._v("cosmovisor")]),s._v(" 将触发升级机制。")])]),s._v(" "),a("p",[s._v("当升级机制被触发时，"),a("code",[s._v("cosmovisor")]),s._v(" 将：")]),s._v(" "),a("ol",[a("li",[s._v("如果启用了 "),a("code",[s._v("DAEMON_ALLOW_DOWNLOAD_BINARIES")]),s._v("，首先将新的二进制文件自动下载到 "),a("code",[s._v("cosmovisor/<name>/bin")]),s._v("（其中 "),a("code",[s._v("<name>")]),s._v(" 是 "),a("code",[s._v("upgrade-info.json:name")]),s._v(" 属性）；")]),s._v(" "),a("li",[s._v("更新"),a("code",[s._v("current")]),s._v("符号链接指向新目录并将"),a("code",[s._v("data/upgrade-info.json")]),s._v("保存到"),a("code",[s._v("cosmovisor/current/upgrade-info.json")]),s._v("。")])]),s._v(" "),a("h3",{attrs:{id:"自动下载"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#自动下载"}},[s._v("#")]),s._v(" 自动下载")]),s._v(" "),a("p",[s._v("通常，"),a("code",[s._v("cosmovisor")]),s._v(" 要求系统管理员在升级之前将所有相关的二进制文件放在磁盘上。 然而，对于不需要这种控制并且想要自动设置的人（也许他们正在同步一个非验证全节点并且想要做很少的维护），还有另一种选择。")]),s._v(" "),a("p",[a("strong",[s._v("注意：我们不建议使用自动下载")]),s._v("，因为它不会提前验证二进制文件是否可用。 如果下载二进制文件有任何问题，cosmovisor 将停止并且不会重新启动应用程序（这可能导致链停止）。")]),s._v(" "),a("p",[s._v("如果 "),a("code",[s._v("DAEMON_ALLOW_DOWNLOAD_BINARIES")]),s._v(" 设置为 "),a("code",[s._v("true")]),s._v("，并且触发升级时找不到本地二进制文件，"),a("code",[s._v("cosmovisor")]),s._v(" 将尝试根据 "),a("code",[s._v("data 中的")]),s._v("info"),a("code",[s._v("属性中的说明自行下载和安装二进制文件 /upgrade-info.json")]),s._v(" 文件。 这些文件由 x/upgrade 模块构建，包含来自升级“计划”对象的数据。 "),a("code",[s._v("Plan")]),s._v(" 有一个 info 字段，该字段应具有以下两种有效格式之一来指定下载：")]),s._v(" "),a("ol",[a("li",[s._v("将 os/architecture -> 二进制 URI 映射以 JSON 格式存储在“binaries”键下的升级计划信息字段中。 例如：")])]),s._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"binaries"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"linux/amd64"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"')]),s._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])]),a("p",[s._v("您可以一次包含多个二进制文件，以确保多个环境将接收正确的二进制文件：")]),s._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"binaries"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"linux/amd64"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"linux/arm64"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n    "),a("span",{pre:!0,attrs:{class:"token property"}},[s._v('"darwin/amd64"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"https://example.com/plugchain.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"')]),s._v("\n    "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])]),a("p",[s._v("将其作为提案提交时，请确保没有空格。 使用 "),a("code",[s._v("plugchaind")]),s._v(" 的示例命令可能如下所示：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("plugchaind tx gov submit-proposal software-upgrade Vega "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--title")]),s._v(" Vega "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--deposit")]),s._v(" 100uatom "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n--upgrade-height "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("7368420")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n--upgrade-info "),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('\'{"binaries":{"linux/amd64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-linux-amd64","linux/arm64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-linux-arm64","darwin/amd64":"https://github.com/cosmos/plugchain/releases/download/v6.0.0-rc1/plugchaind-v6.0.0-rc1-darwin-amd64"}}\'')]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--description")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"upgrade to Vega"')]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--gas")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("400000")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--from")]),s._v(" user "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n--chain-id "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--home")]),s._v(" test/val2 "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--node")]),s._v(" tcp://localhost:36657 "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--yes")]),s._v("\n")])])]),a("ol",{attrs:{start:"2"}},[a("li",[a("p",[s._v("以上述格式存储包含所有信息的文件的链接（例如，如果您想指定大量二进制文件、更改日志信息等而不填满区块链）。 例如：")]),s._v(" "),a("div",{staticClass:"language-text extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[s._v("https://example.com/testnet-1001-info.json?checksum=sha256:deaaa99fda9407c4dbe1d04bd49bab0cc3c1dd76fa392cd55a9425be074af01e\n")])])])])]),s._v(" "),a("p",[s._v("当触发"),a("code",[s._v("cosmovisor")]),s._v("下载新的二进制文件时，"),a("code",[s._v("cosmovisor")]),s._v("会解析"),a("code",[s._v('"binaries"')]),s._v("字段，用[go-getter]下载新的二进制文件(https://github.com/hashicorp/go-getter) , 并在 "),a("code",[s._v("upgrades/<name>")]),s._v(" 文件夹中解压缩新的二进制文件，以便它可以像手动安装一样运行。")]),s._v(" "),a("p",[s._v("请注意，要使此机制提供强大的安全保证，所有 URL 都应包含 SHA 256/512 校验。 这确保了不会运行错误的二进制文件，即使有人入侵了服务器或劫持了 DNS。 "),a("code",[s._v("go-getter")]),s._v(" 将始终确保下载的文件与提供的校验和匹配。 "),a("code",[s._v("go-getter")]),s._v(" 还将处理解压缩档案到目录中（在这种情况下，下载链接应指向 "),a("code",[s._v("bin")]),s._v(" 目录中所有数据的 "),a("code",[s._v("zip")]),s._v(" 文件）。")]),s._v(" "),a("p",[s._v("要在 linux 上正确创建 sha256 校验和，您可以使用 "),a("code",[s._v("sha256sum")]),s._v(" 实用程序。 例如：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("sha256sum ./testdata/repo/zip_directory/autod.zip\n")])])]),a("p",[s._v("结果将如下所示：\n"),a("code",[s._v("29139e1381b8177aec909fab9a75d11381cab5adf7d3af0c05ff1c9c117743a7")]),s._v(".")]),s._v(" "),a("p",[s._v("如果您希望使用更长的哈希值，您也可以使用 "),a("code",[s._v("sha512sum")]),s._v("，如果您希望使用损坏的哈希值，则可以使用 "),a("code",[s._v("md5sum")]),s._v("。 无论您选择哪种方式，请确保在 URL 的校验和参数中正确设置哈希算法。")]),s._v(" "),a("h2",{attrs:{id:"示例-simapp-升级"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#示例-simapp-升级"}},[s._v("#")]),s._v(" 示例：SimApp 升级")]),s._v(" "),a("p",[s._v("以下说明使用 Cosmos SDK 源代码附带的模拟应用程序 ("),a("code",[s._v("simapp")]),s._v(") 演示了 "),a("code",[s._v("cosmovisor")]),s._v("。 以下命令将从 "),a("code",[s._v("cosmos-sdk")]),s._v(" 存储库中运行。")]),s._v(" "),a("h3",{attrs:{id:"链设置"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#链设置"}},[s._v("#")]),s._v(" 链设置")]),s._v(" "),a("p",[s._v("让我们使用 "),a("code",[s._v("v0.44")]),s._v(" 版本的 simapp（Cosmos SDK 演示应用）创建一个新链：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" checkout v0.44.6\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" build\n")])])]),a("p",[s._v("清理 "),a("code",[s._v("~/.simapp")]),s._v("（不要在生产环境中这样做）：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("./build/simd unsafe-reset-all\n")])])]),a("p",[s._v("设置应用配置：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("./build/simd config chain-id "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v("\n./build/simd config keyring-backend "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v("\n./build/simd config broadcast-mode block\n")])])]),a("p",[s._v("初始化节点并覆盖任何先前的创世文件（切勿在生产环境中这样做）：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("./build/simd init "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v(" --chain-id "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--overwrite")]),s._v("\n")])])]),a("p",[s._v("在 "),a("code",[s._v("~/.simapp/config/app.toml")]),s._v(" 中将最低 gas 价格设置为 "),a("code",[s._v("0stake")]),s._v("：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("minimum-gas-prices "),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('"0stake"')]),s._v("\n")])])]),a("p",[s._v("为了演示，将 "),a("code",[s._v("genesis.json")]),s._v(" 中的 "),a("code",[s._v("voting_period")]),s._v(" 修改为 20 秒 ("),a("code",[s._v("20s")]),s._v(") 的缩减时间：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("cat")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("<<<")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token variable"}},[a("span",{pre:!0,attrs:{class:"token variable"}},[s._v("$(")]),s._v("jq "),a("span",{pre:!0,attrs:{class:"token string"}},[s._v("'.app_state.gov.voting_params.voting_period = \"20s\"'")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token environment constant"}},[s._v("$HOME")]),s._v("/.simapp/config/genesis.json"),a("span",{pre:!0,attrs:{class:"token variable"}},[s._v(")")])]),s._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v(">")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token environment constant"}},[s._v("$HOME")]),s._v("/.simapp/config/genesis.json\n")])])]),a("p",[s._v("创建一个验证器，并设置创世交易：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("./build/simd keys "),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("add")]),s._v(" validator\n./build/simd add-genesis-account validator 1000000000stake --keyring-backend "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v("\n./build/simd gentx validator 1000000stake --chain-id "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("test")]),s._v("\n./build/simd collect-gentxs\n")])])]),a("h4",{attrs:{id:"准备-cosmovisor-并启动链"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#准备-cosmovisor-并启动链"}},[s._v("#")]),s._v(" 准备 Cosmovisor 并启动链")]),s._v(" "),a("p",[s._v("设置所需的环境变量:")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("export")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token assign-left variable"}},[s._v("DAEMON_NAME")]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("simd\n"),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("export")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token assign-left variable"}},[s._v("DAEMON_HOME")]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),a("span",{pre:!0,attrs:{class:"token environment constant"}},[s._v("$HOME")]),s._v("/.simapp\n")])])]),a("p",[s._v("设置可选环境变量以触发自动应用重启：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("export")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token assign-left variable"}},[s._v("DAEMON_RESTART_AFTER_UPGRADE")]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("true\n")])])]),a("p",[s._v("为 genesis 二进制文件创建文件夹并复制 "),a("code",[s._v("simd")]),s._v(" 二进制文件：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("mkdir")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token variable"}},[s._v("$DAEMON_HOME")]),s._v("/cosmovisor/genesis/bin\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("cp")]),s._v(" ./build/simd "),a("span",{pre:!0,attrs:{class:"token variable"}},[s._v("$DAEMON_HOME")]),s._v("/cosmovisor/genesis/bin\n")])])]),a("p",[s._v("现在您可以使用 simapp v0.44 运行 cosmovisor：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("cosmovisor run start\n")])])]),a("h4",{attrs:{id:"更新应用"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#更新应用"}},[s._v("#")]),s._v(" 更新应用")]),s._v(" "),a("p",[s._v("将应用程序更新到最新版本（例如 v0.45）。")]),s._v(" "),a("p",[s._v("接下来，我们可以添加一个迁移 - 使用 "),a("code",[s._v("x/upgrade")]),s._v(" "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/blob/master/docs/core/upgrade.md",target:"_blank",rel:"noopener noreferrer"}},[s._v("升级计划"),a("OutboundLink")],1),s._v(" 定义（你 如果您使用的是较旧的 Cosmos SDK 版本，则可能指的是过去的版本）。 在迁移中，我们可以进行任何确定性的状态更改。")]),s._v(" "),a("p",[s._v("构建新版本的 "),a("code",[s._v("simd")]),s._v(" 二进制文件：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" build\n")])])]),a("p",[s._v("为升级二进制文件创建文件夹并复制“simd”二进制文件：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("mkdir")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token variable"}},[s._v("$DAEMON_HOME")]),s._v("/cosmovisor/upgrades/test1/bin\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("cp")]),s._v(" ./build/simd "),a("span",{pre:!0,attrs:{class:"token variable"}},[s._v("$DAEMON_HOME")]),s._v("/cosmovisor/upgrades/test1/bin\n")])])]),a("p",[s._v("打开一个新的终端窗口并提交升级提案以及存款和投票（这些命令必须在 20 秒内运行）：")]),s._v(" "),a("div",{staticClass:"language-sh extra-class"},[a("pre",{pre:!0,attrs:{class:"language-sh"}},[a("code",[s._v("./build/simd tx gov submit-proposal software-upgrade test1 "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--title")]),s._v(" upgrade "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--description")]),s._v(" upgrade --upgrade-height "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("200")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--from")]),s._v(" validator "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--yes")]),s._v("\n./build/simd tx gov deposit "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v(" 10000000stake "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--from")]),s._v(" validator "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--yes")]),s._v("\n./build/simd tx gov vote "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("yes")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--from")]),s._v(" validator "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--yes")]),s._v("\n")])])]),a("p",[s._v("升级将在高度 200 处自动发生。注意：如果您的测试游戏需要更多时间，您可能需要在上面的代码段中更改升级高度。")])])}),[],!1,null,null,null);a.default=v.exports}}]);