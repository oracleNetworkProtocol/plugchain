(window.webpackJsonp=window.webpackJsonp||[]).push([[47],{368:function(t,e,v){"use strict";v.r(e);var a=v(15),_=Object(a.a)({},(function(){var t=this,e=t._self._c;return e("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[e("h1",{attrs:{id:"gov-parameters"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#gov-parameters"}},[t._v("#")]),t._v(" Gov Parameters")]),t._v(" "),e("p",[t._v("In Plug Chain Hub, there are some special parameters that can be modified through on-chain governance.\nAll the plugchaind holders may participate in the on-chain governance. If the community is not satisfied with certain modifiable parameters, it is available to submit a "),e("RouterLink",{attrs:{to:"/features/governance.html#usage-scenario-of-parameter-change"}},[t._v("parameter-change")]),t._v(" proposal, and the params will be changed online automatically when the proposal passes.")],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-auth"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-auth"}},[t._v("#")]),t._v(" Parameters in Auth")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("auth/MaxMemoCharacters")])]),t._v(" "),e("td",[t._v("Maximum number of characters in the memo field in a transaction")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("256")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("auth/TxSigLimit")])]),t._v(" "),e("td",[t._v("Maximum number of signatures per transaction")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("7")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("auth/TxSizeCostPerByte")])]),t._v(" "),e("td",[t._v("The amount of gas consumed per byte of the transaction")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("10")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("auth/SigVerifyCostED25519")])]),t._v(" "),e("td",[t._v("Gas spent on edd2519 algorithm signature verification")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("590")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("auth/SigVerifyCostSecp256k1")])]),t._v(" "),e("td",[t._v("Gas spent on secp256k1 algorithm signature verification")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("1000")])])])]),t._v(" "),e("h2",{attrs:{id:"parameters-in-bank"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-bank"}},[t._v("#")]),t._v(" Parameters in Bank")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("bank/SendEnabled")])]),t._v(" "),e("td",[t._v("Tokens that support transfer")]),t._v(" "),e("td"),t._v(" "),e("td",[t._v("[]")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("bank/DefaultSendEnabled")])]),t._v(" "),e("td",[t._v("Whether to enable the transfer function by default")]),t._v(" "),e("td",[t._v("{true,false}")]),t._v(" "),e("td",[t._v("true")])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/bank.html"}},[t._v("Bank")])],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-distribution"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-distribution"}},[t._v("#")]),t._v(" Parameters in Distribution")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("distribution/communitytax")])]),t._v(" "),e("td",[t._v("Fees charged for withdrawal")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.02")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("distribution/baseproposerreward")])]),t._v(" "),e("td",[t._v("The base reward rate of the block proposer")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.01")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("distribution/bonusproposerreward")])]),t._v(" "),e("td",[t._v("Reward rate for block proposers")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.04")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("distribution/withdrawaddrenabled")])]),t._v(" "),e("td",[t._v("Whether to support setting the withdrawal address")]),t._v(" "),e("td",[t._v("{true,false}")]),t._v(" "),e("td",[t._v("true")])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/distribution.html"}},[t._v("Distribution")])],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-gov"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-gov"}},[t._v("#")]),t._v(" Parameters in Gov")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("gov/depositparams")])]),t._v(" "),e("td",[t._v("Related parameters of the deposit mortgage phase")]),t._v(" "),e("td",[t._v("max_deposit_period:(0, 9223372036854775807]")]),t._v(" "),e("td",[t._v('{"min_deposit": [{"denom": "uplugcn", "amount": "10000000"}], "max_deposit_period": "604800s" }')])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("gov/votingparams")])]),t._v(" "),e("td",[t._v("Related parameters of the voting mortgage phase")]),t._v(" "),e("td",[t._v("voting_period:(0, 9223372036854775807]")]),t._v(" "),e("td",[t._v('{"voting_period": "1209600s"}')])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("gov/tallyparams")])]),t._v(" "),e("td",[t._v("Related parameters of the voting tally phase")]),t._v(" "),e("td",[t._v("quorum:[0,1]"),e("br"),t._v("threshold:(0,1]"),e("br"),t._v("veto_threshold:(0,1]")]),t._v(" "),e("td",[t._v('{"quorum":"0.334000000000000000","threshold": "0.500000000000000000","veto_threshold": "0.334000000000000000"}')])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/governance.html"}},[t._v("Governance")])],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-ibc"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-ibc"}},[t._v("#")]),t._v(" Parameters in IBC")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("ibc/AllowedClients")])]),t._v(" "),e("td",[t._v("Clients that support ibc")]),t._v(" "),e("td"),t._v(" "),e("td",[t._v('["06-solomachine","07-tendermint"]')])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("transfer/SendEnabled")])]),t._v(" "),e("td",[t._v("Whether to enable the transfer function")]),t._v(" "),e("td",[t._v("{true,false}")]),t._v(" "),e("td",[t._v("true")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("transfer/ReceiveEnabled")])]),t._v(" "),e("td",[t._v("Whether to enable the receive function")]),t._v(" "),e("td",[t._v("{true,false}")]),t._v(" "),e("td",[t._v("true")])])])]),t._v(" "),e("h2",{attrs:{id:"parameters-in-mint"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-mint"}},[t._v("#")]),t._v(" Parameters in Mint")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("mint/Inflation")])]),t._v(" "),e("td",[t._v("Token issuance frequency")]),t._v(" "),e("td",[t._v("[0, 0.2]")]),t._v(" "),e("td",[t._v("0.13")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("mint/MintDenom")])]),t._v(" "),e("td",[t._v("Denom of the token mintable")]),t._v(" "),e("td"),t._v(" "),e("td",[t._v("uplugcn")])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/mint.html"}},[t._v("Mint")])],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-slashing"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-slashing"}},[t._v("#")]),t._v(" Parameters in Slashing")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("slashing/DowntimeJailDuration")])]),t._v(" "),e("td",[t._v("Maximum downtime  (continuous)")]),t._v(" "),e("td",[t._v("(0, 9223372036854775807]")]),t._v(" "),e("td",[t._v("600s")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("slashing/MinSignedPerWindow")])]),t._v(" "),e("td",[t._v("Minimum signature rate in each window")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.5")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("slashing/SignedBlocksWindow")])]),t._v(" "),e("td",[t._v("Sliding window for downtime slashing")]),t._v(" "),e("td",[t._v("(0, 18446744073709551615]")]),t._v(" "),e("td",[t._v("100")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("slashing/SlashFractionDoubleSign")])]),t._v(" "),e("td",[t._v("Penalty coefficient for double sign")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.05")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("slashing/SlashFractionDowntime")])]),t._v(" "),e("td",[t._v("Penalty coefficient for downtime")]),t._v(" "),e("td",[t._v("[0, 1]")]),t._v(" "),e("td",[t._v("0.01")])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/slashing.html"}},[t._v("Slashing")])],1),t._v(" "),e("h2",{attrs:{id:"parameters-in-staking"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#parameters-in-staking"}},[t._v("#")]),t._v(" Parameters in Staking")]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",[t._v("key")]),t._v(" "),e("th",[t._v("Description")]),t._v(" "),e("th",[t._v("Range")]),t._v(" "),e("th",[t._v("Current")])])]),t._v(" "),e("tbody",[e("tr",[e("td",[e("code",[t._v("staking/UnbondingTime")])]),t._v(" "),e("td",[t._v("Mortgage redemption time")]),t._v(" "),e("td",[t._v("(0, 9223372036854775807]")]),t._v(" "),e("td",[t._v("1814400s")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("staking/MaxValidators")])]),t._v(" "),e("td",[t._v("Maximum number of validators")]),t._v(" "),e("td",[t._v("(0, 4294967295]")]),t._v(" "),e("td",[t._v("100")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("staking/MaxEntries")])]),t._v(" "),e("td",[t._v("The maximum number of unbinding/redelegation orders in progress")]),t._v(" "),e("td",[t._v("(0, 4294967295]")]),t._v(" "),e("td",[t._v("7")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("staking/BondDenom")])]),t._v(" "),e("td",[t._v("Bond denom")]),t._v(" "),e("td"),t._v(" "),e("td",[t._v("uplugcn")])]),t._v(" "),e("tr",[e("td",[e("code",[t._v("staking/HistoricalEntries")])]),t._v(" "),e("td",[t._v("Historical entries")]),t._v(" "),e("td",[t._v("[0, 4294967295]")]),t._v(" "),e("td",[t._v("10000")])])])]),t._v(" "),e("p",[t._v("Details in "),e("RouterLink",{attrs:{to:"/features/staking.html"}},[t._v("Staking")])],1)])}),[],!1,null,null,null);e.default=_.exports}}]);