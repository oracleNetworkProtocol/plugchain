(window.webpackJsonp=window.webpackJsonp||[]).push([[26],{347:function(t,a,s){"use strict";s.r(a);var n=s(15),e=Object(n.a)({},(function(){var t=this,a=t._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"bank"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#bank"}},[t._v("#")]),t._v(" Bank")]),t._v(" "),a("p",[t._v("Bank module allows you to manage assets in your local accounts.")]),t._v(" "),a("h2",{attrs:{id:"available-commands"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#available-commands"}},[t._v("#")]),t._v(" Available Commands")]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("Name")]),t._v(" "),a("th",[t._v("Description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[a("a",{attrs:{href:"#plugchaind-query-bank-balances"}},[t._v("balances")])]),t._v(" "),a("td",[t._v("Query for account balances by address")])]),t._v(" "),a("tr",[a("td",[a("a",{attrs:{href:"#plugchaind-query-bank-total"}},[t._v("total")])]),t._v(" "),a("td",[t._v("Query the total supply of coins of the chain")])]),t._v(" "),a("tr",[a("td",[a("a",{attrs:{href:"#plugchaind-tx-bank-send"}},[t._v("send")])]),t._v(" "),a("td",[t._v("Create and/or sign and broadcast a MsgSend transaction")])])])]),t._v(" "),a("h2",{attrs:{id:"plugchaind-query-bank-balances"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-bank-balances"}},[t._v("#")]),t._v(" plugchaind query bank balances")]),t._v(" "),a("p",[t._v("Query the total balance of an account or of a specific denomination.")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind query bank balances "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("address"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])])]),a("p",[a("strong",[t._v("Flags:")])]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("Name, shorthand")]),t._v(" "),a("th",[t._v("Type")]),t._v(" "),a("th",[t._v("Required")]),t._v(" "),a("th",[t._v("Default")]),t._v(" "),a("th",[t._v("Description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v("-h, --help")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("Help for coin-type")])]),t._v(" "),a("tr",[a("td",[t._v("--denom")]),t._v(" "),a("td",[t._v("string")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("The specific balance denomination to query for")])]),t._v(" "),a("tr",[a("td",[t._v("--count-total")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("Count total number of records in all balances to query for")])])])]),t._v(" "),a("h3",{attrs:{id:"plugchaind-query-bank-total"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-bank-total"}},[t._v("#")]),t._v(" plugchaind query bank total")]),t._v(" "),a("p",[t._v("Query total supply of coins that are held by accounts in the chain.")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind query bank total "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])])]),a("p",[a("strong",[t._v("Flags:")])]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("Name, shorthand")]),t._v(" "),a("th",[t._v("Type")]),t._v(" "),a("th",[t._v("Required")]),t._v(" "),a("th",[t._v("Default")]),t._v(" "),a("th",[t._v("Description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v("-h, --help")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("Help for coin-type")])]),t._v(" "),a("tr",[a("td",[t._v("--denom")]),t._v(" "),a("td",[t._v("string")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("The specific balance denomination to query for")])])])]),t._v(" "),a("h2",{attrs:{id:"plugchaind-tx-bank-send"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-bank-send"}},[t._v("#")]),t._v(" plugchaind tx bank send")]),t._v(" "),a("p",[t._v("Sending tokens to another address, this command includes "),a("code",[t._v("generate")]),t._v(", "),a("code",[t._v("sign")]),t._v(" and "),a("code",[t._v("broadcast")]),t._v(" steps.")]),t._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("plugchaind tx bank send "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("from_key_or_address"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("to_address"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("amount"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])])]),a("p",[a("strong",[t._v("Flags:")])]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("Name, shorthand")]),t._v(" "),a("th",[t._v("Type")]),t._v(" "),a("th",[t._v("Required")]),t._v(" "),a("th",[t._v("Default")]),t._v(" "),a("th",[t._v("Description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v("-h, --help")]),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td"),t._v(" "),a("td",[t._v("Help for balances")])])])])])}),[],!1,null,null,null);a.default=e.exports}}]);