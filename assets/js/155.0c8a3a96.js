(window.webpackJsonp=window.webpackJsonp||[]).push([[155],{476:function(a,t,s){"use strict";s.r(t);var e=s(15),r=Object(e.a)({},(function(){var a=this,t=a._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"软件升级"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#软件升级"}},[a._v("#")]),a._v(" 软件升级")]),a._v(" "),t("h2",{attrs:{id:"概念"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#概念"}},[a._v("#")]),a._v(" 概念")]),a._v(" "),t("h3",{attrs:{id:"计划"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#计划"}},[a._v("#")]),a._v(" 计划")]),a._v(" "),t("p",[a._v("升级模块定义一个 "),t("code",[a._v("Plan")]),a._v("（"),t("code",[a._v("计划")]),a._v("） 类型，该计划调度一个在线升级过程。"),t("code",[a._v("计划")]),a._v(" 可以安排在一个特定的区块高度或时间，但不能同时指定两者。一旦对一个（冻结的）候选发布版本以及一个合适的处理器达成一致，则一个 "),t("code",[a._v("计划")]),a._v(" 被创建，其 "),t("code",[a._v("Name")]),a._v(" 对应一个特定的处理器。通常，"),t("code",[a._v("计划")]),a._v(" 通过一个治理提议过程被创建，当投票通过时，该计划将被调度。一个计划的 "),t("code",[a._v("Info")]),a._v(" 可以包含关于这次升级的各种元数据，典型的，要包含一些上链的特定于应用的升级信息，诸如验证人能自动升级的 "),t("code",[a._v("Git")]),a._v(" 提交。")]),a._v(" "),t("h4",{attrs:{id:"边车进程"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#边车进程"}},[a._v("#")]),a._v(" 边车进程")]),a._v(" "),t("p",[a._v("如果一个运行应用程序二进制的运营者也运行一个边车进程来辅助自动下载和升级二进制，"),t("code",[a._v("Info")]),a._v(" 允许这个进程是无摩擦的。即，升级模块实现 "),t("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor#upgradeable-binary-specification",target:"_blank",rel:"noopener noreferrer"}},[a._v("cosmovisor 可升级的二进制规范"),t("OutboundLink")],1),a._v(" 指定的规范，并且 "),t("code",[a._v("cosmovisor")]),a._v(" 能可选地被用于为节点运营者完全自动化升级过程。通过用必要的信息填充 "),t("code",[a._v("Info")]),a._v(" 字段，二进制能够被自动下载。参考"),t("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor#auto-download",target:"_blank",rel:"noopener noreferrer"}},[a._v("这里"),t("OutboundLink")],1),a._v("。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" Plan "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("struct")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),a._v("\n    Name                "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("string")]),a._v("\n    Time                time"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("Time\n    Height              "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("int64")]),a._v("\n    Info                "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("string")]),a._v("\n    UpgradedClientState "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("*")]),a._v("types"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("Any\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h3",{attrs:{id:"处理器"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#处理器"}},[a._v("#")]),a._v(" 处理器")]),a._v(" "),t("p",[t("code",[a._v("升级")]),a._v(" 模块便于从主版本 X 升级到主版本 Y。为完成这个过程，节点运营者必须首先将当前的二进制升级到一个新二进制，该二进制有一个与新版本 Y 相对应的 "),t("code",[a._v("处理器")]),a._v("。假定这个版本已经过大部分社区成员的充分测试和批准。这个 "),t("code",[a._v("处理器")]),a._v(" 定义了在新二进制 Y 成功运行链之前需要完成的状态迁移。当然，这个 "),t("code",[a._v("处理器")]),a._v(" 是特定于应用而不是在模块基础上定义的。通过应用中的 "),t("code",[a._v("Keeper#SetUpgradeHandler")]),a._v(" 完成 "),t("code",[a._v("处理器")]),a._v(" 的注册。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" UpgradeHandler "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("func")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("(")]),a._v("Context"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(",")]),a._v(" Plan"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(")")]),a._v("\n")])])]),t("p",[a._v("在每个 "),t("code",[a._v("BeginBlock")]),a._v(" 执行期间，"),t("code",[a._v("升级")]),a._v(" 模块检查是否存在一个应该执行的 "),t("code",[a._v("计划")]),a._v("（被调度在 "),t("code",[a._v("BeginBlock")]),a._v(" 运行时的区块高度或时间）。如果存在，则执行对应的 "),t("code",[a._v("处理器")]),a._v("。如果这个计划期待被执行但没有注册相应的处理器，或者二进制升级过早时，节点将优雅地 "),t("code",[a._v("panic")]),a._v(" 并退出。")]),a._v(" "),t("h3",{attrs:{id:"存储加载器"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#存储加载器"}},[a._v("#")]),a._v(" 存储加载器")]),a._v(" "),t("p",[a._v("升级模块也有助于将存储迁移作为升级的一部分。"),t("code",[a._v("存储加载器")]),a._v(" 执行在新二进制成功运行链之前需要完成的迁移。"),t("code",[a._v("存储加载器")]),a._v(" 也是特定于应用而不是基于模块定义的。通过应用中的 "),t("code",[a._v("app#SetStoreLoader")]),a._v(" 完成 "),t("code",[a._v("存储加载器")]),a._v(" 的注册。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("func")]),a._v(" UpgradeStoreLoader "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("(")]),a._v("upgradeHeight "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("int64")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(",")]),a._v(" storeUpgrades "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("*")]),a._v("store"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("StoreUpgrades"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(")")]),a._v(" baseapp"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("StoreLoader\n")])])]),t("p",[a._v("如果存在一个计划的升级并且已到达升级高度，在 panic 之前旧的二进制将写 "),t("code",[a._v("升级信息（UpgradeInfo）")]),a._v(" 到磁盘。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" UpgradeInfo "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("struct")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),a._v("\n    Name    "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("string")]),a._v("\n    Height  "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("int64")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("这个信息对于确保 "),t("code",[a._v("存储升级")]),a._v(" 在正确的区块高度顺利执行以及确保升级符合预期至关重要。它消除了新二进制每次重启时多次执行 "),t("code",[a._v("存储升级")]),a._v(" 的机会。而且，如果在相同的高度存在多个升级计划，"),t("code",[a._v("Name")]),a._v(" 将确保这些 "),t("code",[a._v("存储升级")]),a._v(" 仅仅发生在计划的升级处理器中。")]),a._v(" "),t("p",[a._v("目前在升级过程中，涉及到的状态迁移方式主要支持以下三种："),t("code",[a._v("Renamed")]),a._v("、"),t("code",[a._v("Deleted")]),a._v("、"),t("code",[a._v("Added")])]),a._v(" "),t("h4",{attrs:{id:"renamed"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#renamed"}},[a._v("#")]),a._v(" Renamed")]),a._v(" "),t("p",[a._v("用户可以在升级过程中指定将oldKey(前缀)下的所有数据迁移到newKey(前缀)下存储。")]),a._v(" "),t("h4",{attrs:{id:"deleted"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#deleted"}},[a._v("#")]),a._v(" Deleted")]),a._v(" "),t("p",[a._v("用户可以在升级过程中删除指定key(前缀)下的所有数据。")]),a._v(" "),t("h4",{attrs:{id:"added"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#added"}},[a._v("#")]),a._v(" Added")]),a._v(" "),t("p",[a._v("用户可以在升级过程中以指定key为前缀申请一块新的存储区域。")]),a._v(" "),t("h3",{attrs:{id:"提议"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#提议"}},[a._v("#")]),a._v(" 提议")]),a._v(" "),t("p",[a._v("通常，一个 "),t("code",[a._v("计划")]),a._v(" 经由 "),t("code",[a._v("软件升级提议（SoftwareUpgradeProposal）")]),a._v(" 通过治理的方式被提出并提交，这个提议规定了标准的治理过程。如果这个提议被通过，目标为一个特定 "),t("code",[a._v("处理器")]),a._v(" 的 "),t("code",[a._v("计划")]),a._v(" 将被持久化并调度。通过在一个新的提议中更新计划时间（"),t("code",[a._v("Plan.Time")]),a._v("），升级可以被推迟或者加速。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" SoftwareUpgradeProposal "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("struct")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),a._v("\n    Title       "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("string")]),a._v("\n    Description "),t("span",{pre:!0,attrs:{class:"token builtin"}},[a._v("string")]),a._v("\n    Plan        Plan\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h4",{attrs:{id:"治理过程"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#治理过程"}},[a._v("#")]),a._v(" 治理过程")]),a._v(" "),t("p",[a._v("当一个升级提议被接受时，升级过程分为如下两个步骤。")]),a._v(" "),t("h5",{attrs:{id:"停止网络共识"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#停止网络共识"}},[a._v("#")]),a._v(" 停止网络共识")]),a._v(" "),t("p",[a._v("软件升级提议被接受后，系统将在指定高度的"),t("code",[a._v("BeginBlock")]),a._v("阶段执行升级前的准备，包括下载升级计划、暂停网络共识。")]),a._v(" "),t("h6",{attrs:{id:"下载升级计划"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#下载升级计划"}},[a._v("#")]),a._v(" 下载升级计划")]),a._v(" "),t("p",[a._v("为了能够顺利升级软件，必须在停止网络共识之前，先记录升级所需要的信息："),t("code",[a._v("计划名称")]),a._v("、"),t("code",[a._v("升级高度")]),a._v("。")]),a._v(" "),t("ul",[t("li",[t("code",[a._v("计划名称")]),a._v("：网络重启时，需要根据计划名称路由到对应的"),t("code",[a._v("UpgradeHandler")]),a._v("以及"),t("code",[a._v("UpgradeStoreLoader")]),a._v("。")]),a._v(" "),t("li",[t("code",[a._v("升级高度")]),a._v("：网络重启时，检查是否需要执行网络升级计划。")])]),a._v(" "),t("h6",{attrs:{id:"暂停网络共识"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#暂停网络共识"}},[a._v("#")]),a._v(" 暂停网络共识")]),a._v(" "),t("p",[a._v("软件升级提议被接受后，系统将在指定高度的"),t("code",[a._v("BeginBlock")]),a._v("阶段优雅的暂停网络共识。")]),a._v(" "),t("h5",{attrs:{id:"重新启动新软件"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#重新启动新软件"}},[a._v("#")]),a._v(" 重新启动新软件")]),a._v(" "),t("p",[a._v("用户替换软件为指定版本并重新启动网络，系统将检测是否包含"),t("code",[a._v("计划名称")]),a._v("所指定的"),t("code",[a._v("处理器")]),a._v("，如果包含，系统首先执行"),t("code",[a._v("处理器")]),a._v("程序，然后开始网络共识，如果不包含，系统报错并退出。")]),a._v(" "),t("h4",{attrs:{id:"取消升级提议"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#取消升级提议"}},[a._v("#")]),a._v(" 取消升级提议")]),a._v(" "),t("p",[a._v("升级提议可以被取消。存在一个 "),t("code",[a._v("取消软件升级 （CancelSoftwareUpgrade）")]),a._v("的提议类型，当该类型提议投票通过时，将移除当前正在进行的升级计划。当然，这需要在升级计划执行之前被投票通过并执行。")]),a._v(" "),t("p",[a._v("如果当前升级计划已经被执行，但是该升级计划存在问题，那么此时提出"),t("code",[a._v("取消软件升级")]),a._v("提议是无效的(因为网络已经停止共识)。这时还有另一方案来弥补这一过失，就是在重新启动网络时，使用"),t("code",[a._v("--unsafe-skip-upgrades")]),a._v("参数跳过指定的升级高度(并不是真的跳过该高度，而是跳过软件升级"),t("code",[a._v("处理器")]),a._v(")。当然这要求参与共识的2/3验证人都执行同样的操作，否则同样无法达成网络共识。")]),a._v(" "),t("h2",{attrs:{id:"升级流程"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#升级流程"}},[a._v("#")]),a._v(" 升级流程")]),a._v(" "),t("h3",{attrs:{id:"提交升级提案"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#提交升级提案"}},[a._v("#")]),a._v(" 提交升级提案")]),a._v(" "),t("p",[a._v("执行软件升级流程的第一步是由治理模块发起一个软件升级提案，该提案详细说明了升级时间或者升级高度，具体见上面"),t("a",{attrs:{href:"#%E6%A6%82%E5%BF%B5"}},[a._v("概念")]),a._v("。发起提案的命令行示例如下：")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx gov submit-proposal software-upgrade "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("plan-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--deposit")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("deposit"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --upgrade-time "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("upgrade-time"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--title")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("title"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --upgrade-info "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("upgrade-info"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--description")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("description"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-b")]),a._v(" block\n")])])]),t("h3",{attrs:{id:"为提案抵押、投票"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#为提案抵押、投票"}},[a._v("#")]),a._v(" 为提案抵押、投票")]),a._v(" "),t("p",[a._v("软件升级提案和其他普通提案的执行流程基本一致，都需要验证人、委托人为该提案发表意见，具体信息请参考"),t("RouterLink",{attrs:{to:"/zh/features/governance.html"}},[a._v("治理模块")]),a._v("。为提案抵押的命令行示例如下：")],1),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx gov deposit "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("proposal-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("deposit"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-b")]),a._v(" block "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-y")]),a._v("\n")])])]),t("p",[a._v("一旦抵押金额达到最小抵押金额，提案将进入投票期，验证人或者委托人需要对该提案发起投票，发起投票的命令行示例如下：")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx gov vote "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("proposal-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("option"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcnchaind "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-b")]),a._v(" block "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-y")]),a._v("\n")])])]),t("p",[a._v("当软件升级提案被通过后，升级模块会创建一项升级计划，在指定高度或者时间使所有节点停止网络共识，等待新的软件重启网络。")]),a._v(" "),t("h3",{attrs:{id:"重启网络"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#重启网络"}},[a._v("#")]),a._v(" 重启网络")]),a._v(" "),t("p",[a._v("当升级提案通过，网络到达指定升级区块高度或时间时，所有节点会停止出块。用户需要根据在第一步"),t("a",{attrs:{href:"#%E6%8F%90%E4%BA%A4%E5%8D%87%E7%BA%A7%E6%8F%90%E6%A1%88"}},[a._v("提交升级提案")]),a._v("中指明的新版本信息下载源码，编译新软件，具体参考"),t("RouterLink",{attrs:{to:"/zh/get-started/install.html"}},[a._v("安装")]),a._v("。新软件安装完成后，使用新的版本重启节点。当全网超过2/3的权重的验证人使用新版本启动节点后，区块链网络将重新达成新的共识，继续产生区块。")],1)])}),[],!1,null,null,null);t.default=r.exports}}]);