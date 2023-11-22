(window.webpackJsonp=window.webpackJsonp||[]).push([[139],{460:function(t,a,s){"use strict";s.r(a);var e=s(15),r=Object(e.a)({},(function(){var t=this,a=t._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"导出区块状态"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#导出区块状态"}},[t._v("#")]),t._v(" 导出区块状态")]),t._v(" "),a("h2",{attrs:{id:"简介"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#简介"}},[t._v("#")]),t._v(" 简介")]),t._v(" "),a("p",[t._v("plugchaind可以导出区块链状态并将其输出为json格式的字符串，该字符串可用作新区块链的创世文件。")]),t._v(" "),a("p",[t._v("默认情况下，plugchaind存储每10,000个块和最后100个块的快照。可以从任何现有快照高度导出区块链状态。")]),t._v(" "),a("p",[t._v("如果要从不存在的快照高度导出状态，则需要首先将区块链状态"),a("RouterLink",{attrs:{to:"/zh/daemon/local-testnet.html#plugchaind-reset"}},[t._v("重置")]),t._v("到指定的高度。")],1),t._v(" "),a("h2",{attrs:{id:"用法"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#用法"}},[t._v("#")]),t._v(" 用法")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v(" plugchaind "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("export")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])])]),a("h2",{attrs:{id:"标识"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#标识"}},[t._v("#")]),t._v(" 标识")]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("名称，速记")]),t._v(" "),a("th",[t._v("类型")]),t._v(" "),a("th",[t._v("必须")]),t._v(" "),a("th",[t._v("默认值")]),t._v(" "),a("th",[t._v("描述")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v("--for-zero-height")]),t._v(" "),a("td",[t._v("bool")]),t._v(" "),a("td"),t._v(" "),a("td",[t._v("false")]),t._v(" "),a("td",[t._v("导出数据之前做一些清理性的工作，如果不想以导出的数据启动一条新链，可以不加这个标识")])]),t._v(" "),a("tr",[a("td",[t._v("--height")]),t._v(" "),a("td",[t._v("uint")]),t._v(" "),a("td"),t._v(" "),a("td",[t._v("0")]),t._v(" "),a("td",[t._v("从指定的高度导出，默认值为0表示导出当前高度状态")])]),t._v(" "),a("tr",[a("td",[t._v("--home")]),t._v(" "),a("td",[t._v("string")]),t._v(" "),a("td"),t._v(" "),a("td",[t._v("$HOME/.plugchain")]),t._v(" "),a("td",[t._v("指定存储配置和区块链数据的目录")])]),t._v(" "),a("tr",[a("td",[t._v("--output-file")]),t._v(" "),a("td",[t._v("string")]),t._v(" "),a("td"),t._v(" "),a("td",[t._v("genesis.json")]),t._v(" "),a("td",[t._v("存储导出状态的文件")])])])]),t._v(" "),a("h2",{attrs:{id:"示例"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#示例"}},[t._v("#")]),t._v(" 示例")]),t._v(" "),a("p",[t._v("导出当前的区块链状态")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v(" plugchaind "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("export")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("path-to-your-home"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v("\n")])])]),a("p",[t._v("从特定高度导出区块链状态，该高度必须是现有快照高度")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("export")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--height")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("10000")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("path-to-your-home"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v("\n")])])]),a("p",[t._v("如果想导出指定高度的区块链状态，并且以这个状态启动一条新链，可以尝试这个命令")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("export")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--height")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("10000")]),t._v(" --for-zero-height "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("path-to-your-home"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v("\n")])])])])}),[],!1,null,null,null);a.default=r.exports}}]);