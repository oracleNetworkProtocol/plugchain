(window.webpackJsonp=window.webpackJsonp||[]).push([[132],{453:function(e,s,a){"use strict";a.r(s);var v=a(15),_=Object(v.a)({},(function(){var e=this,s=e._self._c;return s("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[s("h1",{attrs:{id:"gas-and-fees"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#gas-and-fees"}},[e._v("#")]),e._v(" Gas and Fees")]),e._v(" "),s("p",[e._v("了解 PVM 和 Cosmos 中“Gas”和“Fees”之间的区别。")]),e._v(" "),s("h2",{attrs:{id:"介绍-gas-和-fees"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#介绍-gas-和-fees"}},[e._v("#")]),e._v(" 介绍 "),s("code",[e._v("Gas")]),e._v(" 和 "),s("code",[e._v("Fees")])]),e._v(" "),s("p",[e._v("在 Cosmos SDK 中，"),s("code",[e._v("gas")]),e._v(" 是一个特殊的单位，用于跟踪执行期间的资源消耗。 "),s("code",[e._v("gas")]),e._v(" 通常在对存储进行读写时消耗，但如果需要进行昂贵的计算，也可以消耗它。 它有两个主要目的：")]),e._v(" "),s("ul",[s("li",[e._v("确保区块不会消耗太多资源并将最终确定。")]),e._v(" "),s("li",[e._v("防止最终用户的垃圾邮件和滥用。 为此，在 "),s("code",[e._v("Msg")]),e._v(" 执行期间消耗的 "),s("code",[e._v("gas")]),e._v(" 通常会被定价，从而产生 "),s("code",[e._v("fee")]),e._v("（"),s("code",[e._v("fees = gas * gas-prices")]),e._v("）。 "),s("code",[e._v("fees")]),e._v("一般必须由执行"),s("code",[e._v("Msg")]),e._v("的发送者支付。 请注意，Cosmos SDK 默认不强制执行“gas”定价，因为可能有其他方法可以防止垃圾邮件（例如带宽方案）。 尽管如此，大多数应用程序将实施“费用”机制来防止垃圾邮件。")])]),e._v(" "),s("h2",{attrs:{id:"cosmos-sdk-fees"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#cosmos-sdk-fees"}},[e._v("#")]),e._v(" Cosmos SDK "),s("code",[e._v("Fees")])]),e._v(" "),s("ul",[s("li",[e._v("验证交易提供的 gas 价格是否高于当地的 "),s("code",[e._v("min-gas-prices")]),e._v("。 "),s("code",[e._v("min-gas-prices")]),e._v(" 是每个全节点的本地参数，在 "),s("code",[e._v("CheckTx")]),e._v(" 期间用于丢弃不提供最低费用的交易。 这确保了内存池不会被垃圾事务发送到垃圾邮件中。")]),e._v(" "),s("li",[e._v("公式: "),s("code",[e._v("fees = gas * gas-price")])]),e._v(" "),s("li",[e._v("例如: "),s("code",[e._v("gas=200000")]),e._v("（默认）、"),s("code",[e._v("gas-price=0.0001")]),e._v("、"),s("code",[e._v("fees=20uplugcn")])])]),e._v(" "),s("h2",{attrs:{id:"pvm-fees"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#pvm-fees"}},[e._v("#")]),e._v(" PVM "),s("code",[e._v("Fees")])]),e._v(" "),s("p",[e._v("PVM 和 Cosmos 状态转换之间的主要区别在于，PVM 对每个 OPCODE 使用 "),s("a",{attrs:{href:"https://github.com/ethereum/go-ethereum/blob/master/params/protocol_params.go",target:"_blank",rel:"noopener noreferrer"}},[e._v("gas table"),s("OutboundLink")],1),e._v("，而 Cosmos 使用“GasConfig”，通过设置访问数据库的单位和每字节成本来为每个 CRUD 操作收取 gas。")]),e._v(" "),s("p",[e._v("PVM和EVM的Gas使用计算方式相同， 但PVM "),s("code",[e._v("gas")]),e._v("使用 与 EVM "),s("code",[e._v("gas")]),e._v("使用流程相同，当gas足够支付交易时，多余的gas会返回。")]),e._v(" "),s("ul",[s("li",[e._v("PVM交易首先要满足"),s("code",[e._v("fees")]),e._v("大于Cosmos SDK"),s("code",[e._v("要求的最低手续费（目前最低手续费为0.0001*200000=20uplugcn）")])]),e._v(" "),s("li",[e._v("交易提供的gas应大于或等于根据"),s("a",{attrs:{href:"https://github.com/ethereum/go-ethereum/blob/master/params",target:"_blank",rel:"noopener noreferrer"}},[e._v("gas table"),s("OutboundLink")],1),e._v("交易字节计算的gas总和")]),e._v(" "),s("li",[s("code",[e._v("gasPrice")]),e._v(" 最低为 "),s("code",[e._v("7")])]),e._v(" "),s("li",[e._v("公式: "),s("code",[e._v("fees = gas * gasPrice / 1000000")]),e._v(" (提案"),s("a",{attrs:{href:"https://www.plugchain.network/v2/communityDetail?id=10",target:"_blank",rel:"noopener noreferrer"}},[e._v("v1.7.0"),s("OutboundLink")],1),e._v(" 升级之后的算法)")])])])}),[],!1,null,null,null);s.default=_.exports}}]);