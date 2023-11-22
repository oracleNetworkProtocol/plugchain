(window.webpackJsonp=window.webpackJsonp||[]).push([[36],{357:function(a,t,e){"use strict";e.r(t);var s=e(15),r=Object(s.a)({},(function(){var a=this,t=a._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"staking"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#staking"}},[a._v("#")]),a._v(" Staking")]),a._v(" "),t("p",[a._v("Staking module provides a set of subcommands to query staking state and send staking transactions.")]),a._v(" "),t("h2",{attrs:{id:"available-commands"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#available-commands"}},[a._v("#")]),a._v(" Available Commands")]),a._v(" "),t("table",[t("thead",[t("tr",[t("th",[a._v("Name")]),a._v(" "),t("th",[a._v("Description")])])]),a._v(" "),t("tbody",[t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-validator"}},[a._v("validator")])]),a._v(" "),t("td",[a._v("Query a validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-validators"}},[a._v("validators")])]),a._v(" "),t("td",[a._v("Query for all validators")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-delegation"}},[a._v("delegation")])]),a._v(" "),t("td",[a._v("Query a delegation based on address and validator address")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-delegations"}},[a._v("delegations")])]),a._v(" "),t("td",[a._v("Query all delegations made from one delegator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-delegations-to"}},[a._v("delegations-to")])]),a._v(" "),t("td",[a._v("Query all delegations to one validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-unbonding-delegation"}},[a._v("unbonding-delegation")])]),a._v(" "),t("td",[a._v("Query an unbonding-delegation record based on delegator and validator address")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-unbonding-delegations"}},[a._v("unbonding-delegations")])]),a._v(" "),t("td",[a._v("Query all unbonding-delegations records for one delegator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-unbonding-delegations-from"}},[a._v("unbonding-delegations-from")])]),a._v(" "),t("td",[a._v("Query all unbonding delegatations from a validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-redelegations-from"}},[a._v("redelegations-from")])]),a._v(" "),t("td",[a._v("Query all outgoing redelegatations from a validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-redelegation"}},[a._v("redelegation")])]),a._v(" "),t("td",[a._v("Query a redelegation record based on delegator and a source and destination validator address")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-redelegations"}},[a._v("redelegations")])]),a._v(" "),t("td",[a._v("Query all redelegations records for one delegator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-pool"}},[a._v("pool")])]),a._v(" "),t("td",[a._v("Query the current staking pool values")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-params"}},[a._v("params")])]),a._v(" "),t("td",[a._v("Query the current staking parameters information")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-query-staking-historical-info"}},[a._v("historical-info")])]),a._v(" "),t("td",[a._v("Query historical info at given height")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-tx-staking-create-validator"}},[a._v("create-validator")])]),a._v(" "),t("td",[a._v("Create new validator initialized with a self-delegation to it")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-tx-staking-edit-validator"}},[a._v("edit-validator")])]),a._v(" "),t("td",[a._v("Edit existing validator account")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-tx-staking-delegate"}},[a._v("delegate")])]),a._v(" "),t("td",[a._v("Delegate liquid tokens to an validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-tx-staking-unbond"}},[a._v("unbond")])]),a._v(" "),t("td",[a._v("Unbond shares from a validator")])]),a._v(" "),t("tr",[t("td",[t("a",{attrs:{href:"#plugchaind-tx-staking-redelegate"}},[a._v("redelegate")])]),a._v(" "),t("td",[a._v("Redelegate illiquid tokens from one validator to another")])])])]),a._v(" "),t("h2",{attrs:{id:"plugchaind-query-staking-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-validator"}},[a._v("#")]),a._v(" plugchaind query staking validator")]),a._v(" "),t("h3",{attrs:{id:"query-a-validator-by-validator-address"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-a-validator-by-validator-address"}},[a._v("#")]),a._v(" Query a validator by validator address")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking validator "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-validators"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-validators"}},[a._v("#")]),a._v(" plugchaind query staking validators")]),a._v(" "),t("h3",{attrs:{id:"query-all-validators"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-validators"}},[a._v("#")]),a._v(" Query all validators")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking validators\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-delegation"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-delegation"}},[a._v("#")]),a._v(" plugchaind query staking delegation")]),a._v(" "),t("p",[a._v("Query a delegation based on delegator address and validator address.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegation "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("delegator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-a-delegation"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-a-delegation"}},[a._v("#")]),a._v(" Query a delegation")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegation "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("p",[a._v("Example Output:")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("Delegation:\n  Delegator:  gx13lcwnxpyn2ea3skzmek64vvnp97jsk8qrcezvm\n  Validator:  gxvaloper15grv3xg3ekxh9xrf79zd0w077krgv5xfzzunhs\n  Shares:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("1.0000000000000000000000000000")]),a._v("\n  Height:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("26")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-delegations"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-delegations"}},[a._v("#")]),a._v(" plugchaind query staking delegations")]),a._v(" "),t("p",[a._v("Query all delegations delegated from one delegator.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegations "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("delegator-address"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-all-delegations-of-a-delegator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-delegations-of-a-delegator"}},[a._v("#")]),a._v(" Query all delegations of a delegator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegations "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-delegations-to"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-delegations-to"}},[a._v("#")]),a._v(" plugchaind query staking delegations-to")]),a._v(" "),t("p",[a._v("Query all delegations to one validator.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegations-to "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-address"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-all-delegations-to-one-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-delegations-to-one-validator"}},[a._v("#")]),a._v(" Query all delegations to one validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking delegations-to "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("p",[a._v("Example Output:")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("Delegation:\n  Delegator:  gx13lcwnxpyn2ea3skzmek64vvnp97jsk8qrcezvm\n  Validator:  gxvaloper1yclscskdtqu9rgufgws293wxp3njsesxxlnhmh\n  Shares:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("100.0000000000000000000000000000")]),a._v("\n  Height:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0")]),a._v("\nDelegation:\n  Delegator:  gx1td4xnefkthfs6jg469x33shzf578fed6n7k7ua\n  Validator:  gxvaloper1yclscskdtqu9rgufgws293wxp3njsesxxlnhmh\n  Shares:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("1.0000000000000000000000000000")]),a._v("\n  Height:     "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("26")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-unbonding-delegation"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-unbonding-delegation"}},[a._v("#")]),a._v(" plugchaind query staking unbonding-delegation")]),a._v(" "),t("p",[a._v("Query an unbonding-delegation record based on delegator and validator address.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking unbonding-delegation "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("delegator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-an-unbonding-delegation-record"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-an-unbonding-delegation-record"}},[a._v("#")]),a._v(" Query an unbonding delegation record")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking unbonding-delegation "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-unbonding-delegations"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-unbonding-delegations"}},[a._v("#")]),a._v(" plugchaind query staking unbonding-delegations")]),a._v(" "),t("h3",{attrs:{id:"query-all-unbonding-delegations-records-of-a-delegator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-unbonding-delegations-records-of-a-delegator"}},[a._v("#")]),a._v(" Query all unbonding delegations records of a delegator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking unbonding-delegations "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-unbonding-delegations-from"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-unbonding-delegations-from"}},[a._v("#")]),a._v(" plugchaind query staking unbonding-delegations-from")]),a._v(" "),t("h3",{attrs:{id:"query-all-unbonding-delegations-from-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-unbonding-delegations-from-a-validator"}},[a._v("#")]),a._v(" Query all unbonding delegations from a validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking unbonding-delegations-from "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-redelegations-from"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-redelegations-from"}},[a._v("#")]),a._v(" plugchaind query staking redelegations-from")]),a._v(" "),t("p",[a._v("Query all outgoing redelegations of a validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking redelegations-from "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-address"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-all-outgoing-redelegatations-of-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-outgoing-redelegatations-of-a-validator"}},[a._v("#")]),a._v(" Query all outgoing redelegatations of a validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking redelegations-from "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-redelegation"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-redelegation"}},[a._v("#")]),a._v(" plugchaind query staking redelegation")]),a._v(" "),t("p",[a._v("Query a redelegation record based on delegator and source validator address and destination validator address.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking redelegation "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("delegator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("src-validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("dst-validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"query-a-redelegation-record"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-a-redelegation-record"}},[a._v("#")]),a._v(" Query a redelegation record")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking redelegation "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-redelegations"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-redelegations"}},[a._v("#")]),a._v(" plugchaind query staking redelegations")]),a._v(" "),t("h3",{attrs:{id:"query-all-redelegations-records-of-a-delegator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-all-redelegations-records-of-a-delegator"}},[a._v("#")]),a._v(" Query all redelegations records of a delegator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking redelegations "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gx"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-pool"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-pool"}},[a._v("#")]),a._v(" plugchaind query staking pool")]),a._v(" "),t("h3",{attrs:{id:"query-the-current-staking-pool-values"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-the-current-staking-pool-values"}},[a._v("#")]),a._v(" Query the current staking pool values")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking pool\n")])])]),t("p",[a._v("Example Output:")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("Pool:\n  Loose Tokens:   "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("1409493892.759816067399143966")]),a._v("\n  Bonded Tokens:  "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("590526409.65743521209068061")]),a._v("\n  Token Supply:   "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("2000020302.417251279489824576")]),a._v("\n  Bonded Ratio:   "),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.2952602076")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-params"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-params"}},[a._v("#")]),a._v(" plugchaind query staking params")]),a._v(" "),t("h3",{attrs:{id:"query-the-current-staking-parameters-information"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-the-current-staking-parameters-information"}},[a._v("#")]),a._v(" Query the current staking parameters information")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking params\n")])])]),t("h2",{attrs:{id:"plugchaind-query-staking-historical-info"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-query-staking-historical-info"}},[a._v("#")]),a._v(" plugchaind query staking historical-info")]),a._v(" "),t("h3",{attrs:{id:"query-historical-info-at-given-height"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-historical-info-at-given-height"}},[a._v("#")]),a._v(" Query historical info at given height")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind query staking historical-info "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("height"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h2",{attrs:{id:"plugchaind-tx-staking-create-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-staking-create-validator"}},[a._v("#")]),a._v(" plugchaind tx staking create-validator")]),a._v(" "),t("p",[a._v("Send a transaction to apply to be a validator and delegate a certain amount of plugchaind to it.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking create-validator "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("p",[t("strong",[a._v("Flags:")])]),a._v(" "),t("table",[t("thead",[t("tr",[t("th",[a._v("Name, shorthand")]),a._v(" "),t("th",[a._v("type")]),a._v(" "),t("th",[a._v("Required")]),a._v(" "),t("th",[a._v("Default")]),a._v(" "),t("th",[a._v("Description")])])]),a._v(" "),t("tbody",[t("tr",[t("td",[a._v("--amount")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td",[a._v("Yes")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Amount of coins to bond")])]),a._v(" "),t("tr",[t("td",[a._v("--commission-rate")]),a._v(" "),t("td",[a._v("float")]),a._v(" "),t("td",[a._v("Yes")]),a._v(" "),t("td",[a._v("0.0")]),a._v(" "),t("td",[a._v("The initial commission rate percentage")])]),a._v(" "),t("tr",[t("td",[a._v("--commission-max-rate")]),a._v(" "),t("td",[a._v("float")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("0.0")]),a._v(" "),t("td",[a._v("The maximum commission rate percentage")])]),a._v(" "),t("tr",[t("td",[a._v("--commission-max-change-rate")]),a._v(" "),t("td",[a._v("float")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("0.0")]),a._v(" "),t("td",[a._v("The maximum commission change rate percentage (per day)")])]),a._v(" "),t("tr",[t("td",[a._v("--min-self-delegation")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("The minimum self delegation required on the validator")])]),a._v(" "),t("tr",[t("td",[a._v("--details")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional details")])]),a._v(" "),t("tr",[t("td",[a._v("--genesis-format")]),a._v(" "),t("td",[a._v("bool")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("false")]),a._v(" "),t("td",[a._v("Export the transaction in gen-tx format; it implies --generate-only")])]),a._v(" "),t("tr",[t("td",[a._v("--identity")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional identity signature (ex. UPort or Keybase)")])]),a._v(" "),t("tr",[t("td",[a._v("--ip")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Node's public IP. It takes effect only when used in combination with")])]),a._v(" "),t("tr",[t("td",[a._v("--node-id")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("The node's ID")])]),a._v(" "),t("tr",[t("td",[a._v("--moniker")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td",[a._v("Yes")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Validator name")])]),a._v(" "),t("tr",[t("td",[a._v("--pubkey")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td",[a._v("Yes")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Go-Amino encoded hex PubKey of the validator. For Ed25519 the go-amino prepend hex is 1624de6220")])]),a._v(" "),t("tr",[t("td",[a._v("--website")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional website")])]),a._v(" "),t("tr",[t("td",[a._v("--security-contact")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("The validator's (optional) security contact email")])])])]),a._v(" "),t("h3",{attrs:{id:"create-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#create-a-validator"}},[a._v("#")]),a._v(" Create a validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking create-validator --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("plugchain_520-1 "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--pubkey")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("validator-pubKey"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.1")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--amount")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("100uplugcn "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--moniker")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("validator-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[a._v("Follow the "),t("RouterLink",{attrs:{to:"/get-started/mainnet.html#create-validator"}},[a._v("Mainnet")]),a._v(" instructions to learn more.")],1)]),a._v(" "),t("h2",{attrs:{id:"plugchaind-tx-staking-edit-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-staking-edit-validator"}},[a._v("#")]),a._v(" plugchaind tx staking edit-validator")]),a._v(" "),t("p",[a._v("Edit an existing validator's settings, such as commission rate, name, etc.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking edit-validator "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("p",[t("strong",[a._v("Flags:")])]),a._v(" "),t("table",[t("thead",[t("tr",[t("th",[a._v("Name, shorthand")]),a._v(" "),t("th",[a._v("type")]),a._v(" "),t("th",[a._v("Required")]),a._v(" "),t("th",[a._v("Default")]),a._v(" "),t("th",[a._v("Description")])])]),a._v(" "),t("tbody",[t("tr",[t("td",[a._v("--commission-rate")]),a._v(" "),t("td",[a._v("float")]),a._v(" "),t("td"),a._v(" "),t("td",[a._v("0.0")]),a._v(" "),t("td",[a._v("Commission rate percentage")])]),a._v(" "),t("tr",[t("td",[a._v("--moniker")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Validator name")])]),a._v(" "),t("tr",[t("td",[a._v("--identity")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional identity signature (ex. UPort or Keybase)")])]),a._v(" "),t("tr",[t("td",[a._v("--website")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional website")])]),a._v(" "),t("tr",[t("td",[a._v("--details")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("Optional details")])]),a._v(" "),t("tr",[t("td",[a._v("--security-contact")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("The validator's (optional) security contact email")])]),a._v(" "),t("tr",[t("td",[a._v("--min-self-delegation")]),a._v(" "),t("td",[a._v("string")]),a._v(" "),t("td"),a._v(" "),t("td"),a._v(" "),t("td",[a._v("The minimum self delegation required on the validator")])])])]),a._v(" "),t("h3",{attrs:{id:"edit-validator-information"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#edit-validator-information"}},[a._v("#")]),a._v(" Edit validator information")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking edit-validator "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("plugchain_520-1 "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn --commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.10")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--moniker")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("validator-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])])]),t("h3",{attrs:{id:"upload-validator-avatar"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#upload-validator-avatar"}},[a._v("#")]),a._v(" Upload validator avatar")]),a._v(" "),t("p",[a._v("Please refer to "),t("RouterLink",{attrs:{to:"/concepts/validator-faq.html#how-to-upload-my-validator-s-logo-to-the-explorers"}},[a._v("How to upload my validator's logo to the Explorers")])],1),a._v(" "),t("h2",{attrs:{id:"plugchaind-tx-staking-delegate"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-staking-delegate"}},[a._v("#")]),a._v(" plugchaind tx staking delegate")]),a._v(" "),t("p",[a._v("Delegate tokens to a validator.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking delegate "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("amount"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking delegate "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("amount"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("plugchain_520-1 "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn\n")])])]),t("h2",{attrs:{id:"plugchaind-tx-staking-unbond"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-staking-unbond"}},[a._v("#")]),a._v(" plugchaind tx staking unbond")]),a._v(" "),t("p",[a._v("Unbond tokens from a validator.")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking unbond "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("amount"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"unbond-some-tokens-from-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#unbond-some-tokens-from-a-validator"}},[a._v("#")]),a._v(" Unbond some tokens from a validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking unbond "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" 10uplugcn "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("plugchain_520-1 "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn\n")])])]),t("h2",{attrs:{id:"plugchaind-tx-staking-redelegate"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugchaind-tx-staking-redelegate"}},[a._v("#")]),a._v(" plugchaind tx staking redelegate")]),a._v(" "),t("p",[a._v("Transfer delegation from one validator to another.")]),a._v(" "),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[a._v("There is no "),t("code",[a._v("unbonding time")]),a._v(" during the redelegation, so you will not miss the rewards. But you can only redelegate once per validator, until a period (= "),t("code",[a._v("unbonding time")]),a._v(") exceed.")])]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking redelegate "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("src-validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("dst-validator-addr"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("amount"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("flags"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h3",{attrs:{id:"redelegate-some-tokens-to-another-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#redelegate-some-tokens-to-another-validator"}},[a._v("#")]),a._v(" Redelegate some tokens to another validator")]),a._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("plugchaind tx staking redelegate "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("gxvaloper"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("..")]),a._v("."),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" 10uplugcn --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("plugchain_520-1 "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--from")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key-name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--fees")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("20uplugcn\n")])])])])}),[],!1,null,null,null);t.default=r.exports}}]);