(window.webpackJsonp=window.webpackJsonp||[]).push([[51],{372:function(t,a,e){"use strict";e.r(a);var s=e(15),r=Object(s.a)({},(function(){var t=this,a=t._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"set-client-configuration"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#set-client-configuration"}},[t._v("#")]),t._v(" Set client configuration")]),t._v(" "),a("h2",{attrs:{id:"introduction"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#introduction"}},[t._v("#")]),t._v(" Introduction")]),t._v(" "),a("p",[t._v("After the plugchain is upgraded to v1.0+, the terminal operation can use the "),a("code",[t._v("config")]),t._v(" command to configure the client configuration of the current data directory, or directly modify the "),a("code",[t._v("$HOME/.plugchain/config/client.toml")]),t._v(" file to meet some parameters of the operation command set up. There is no need to specify parameters for each operation command.")]),t._v(" "),a("h2",{attrs:{id:"usage"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#usage"}},[t._v("#")]),t._v(" usage")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v(" plugchaind config "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("key"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("value"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])])]),a("table",[a("thead",[a("tr",[a("th",[t._v("name, shorthand")]),t._v(" "),a("th",[t._v("type")]),t._v(" "),a("th",[t._v("required")]),t._v(" "),a("th",[t._v("default value")]),t._v(" "),a("th",[t._v("description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v("--home")]),t._v(" "),a("td",[t._v("string")]),t._v(" "),a("td"),t._v(" "),a("td",[t._v("$HOME/.plugchain")]),t._v(" "),a("td",[t._v("Specifies the directory to store configuration and blockchain data")])])])]),t._v(" "),a("ul",[a("li",[t._v("Set the main chain ID of the current directory")])]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v(" plugchaind config chain-id "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"plugchain_520-1"')]),t._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("path-to-your-home"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v("\n")])])]),a("ul",[a("li",[t._v("Query the main chain ID of the current directory")])]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v(" plugchaind config chain-id "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("<")]),t._v("path-to-your-home"),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">")]),t._v("\n")])])]),a("ul",[a("li",[t._v("The fields of operability are as follows:")])]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# The network chain ID")]),t._v("\nchain-id "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"plugchain_520-1"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)")]),t._v("\nkeyring-backend "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"os"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# CLI output format (text|json)")]),t._v("\noutput "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"text"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# <host>:<port> to Tendermint RPC interface for this chain")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("node")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"tcp://localhost:26657"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Transaction broadcasting mode (sync|async|block)")]),t._v("\nbroadcast-mode "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"sync"')]),t._v("\n")])])]),a("ul",[a("li",[t._v("After setting, when the operation command involves parameters such as "),a("code",[t._v("chain-id")]),t._v(", "),a("code",[t._v("keyring-backend")]),t._v(", etc., no additional specification is required. Take the data directory as "),a("code",[t._v("node")]),t._v(" as an example as follows:")])]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind tx bank send adr adr1 20000uplugcn "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--fees")]),t._v(" 200uplugcn "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[t._v("--home")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("node")]),t._v("\n")])])])])}),[],!1,null,null,null);a.default=r.exports}}]);