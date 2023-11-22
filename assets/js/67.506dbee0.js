(window.webpackJsonp=window.webpackJsonp||[]).push([[67],{388:function(e,a,t){"use strict";t.r(a);var s=t(15),r=Object(s.a)({},(function(){var e=this,a=e._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"software-upgrade"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#software-upgrade"}},[e._v("#")]),e._v(" Software Upgrade")]),e._v(" "),a("h2",{attrs:{id:"concepts"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#concepts"}},[e._v("#")]),e._v(" Concepts")]),e._v(" "),a("h3",{attrs:{id:"plan"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#plan"}},[e._v("#")]),e._v(" Plan")]),e._v(" "),a("p",[e._v("The "),a("code",[e._v("Upgrade")]),e._v(" module defines a "),a("code",[e._v("Plan")]),e._v(" type in which a live upgrade is scheduled to occur. A "),a("code",[e._v("Plan")]),e._v(" can be scheduled at a specific block height or time, but not both.\nA "),a("code",[e._v("Plan")]),e._v(" is created once a (frozen) release candidate along with an appropriate upgrade "),a("code",[e._v("Handler")]),e._v(" (see below) is agreed upon, where the "),a("code",[e._v("Name")]),e._v(" of a "),a("code",[e._v("Plan")]),e._v(" corresponds to a\nspecific "),a("code",[e._v("Handler")]),e._v(". Typically, a "),a("code",[e._v("Plan")]),e._v(" is created through a governance proposal process, where if voted upon and passed, will be scheduled. The "),a("code",[e._v("Info")]),e._v(" of a "),a("code",[e._v("Plan")]),e._v("\nmay contain various metadata about the upgrade, typically application specific upgrade info to be included on-chain such as a git commit that validators could automatically upgrade to.")]),e._v(" "),a("h4",{attrs:{id:"sidecar-process"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#sidecar-process"}},[e._v("#")]),e._v(" Sidecar Process")]),e._v(" "),a("p",[e._v("If an operator running the application binary also runs a sidecar process to assist in the automatic download and upgrade of the binary, "),a("code",[e._v("Info")]),e._v(" allows this process to be frictionless. That is, the upgrade module implements the specification specified by "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor#upgradeable-binary-specification",target:"_blank",rel:"noopener noreferrer"}},[e._v("cosmovisor upgradeable binary specification"),a("OutboundLink")],1),e._v(", and "),a("code",[e._v("cosmovisor")]),e._v(" can Site selection is used to fully automate the upgrade process for node operators. By filling the "),a("code",[e._v("Info")]),e._v(" field with the necessary information, the binary can be downloaded automatically. Refer to "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor#auto-download",target:"_blank",rel:"noopener noreferrer"}},[e._v("here"),a("OutboundLink")],1),e._v(".")]),e._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("type")]),e._v(" Plan "),a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("struct")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("{")]),e._v("\n    Name                "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("string")]),e._v("\n    Time                time"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(".")]),e._v("Time\n    Height              "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("int64")]),e._v("\n    Info                "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("string")]),e._v("\n    UpgradedClientState "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("*")]),e._v("types"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(".")]),e._v("Any\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("}")]),e._v("\n")])])]),a("h3",{attrs:{id:"handler"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#handler"}},[e._v("#")]),e._v(" Handler")]),e._v(" "),a("p",[e._v("The "),a("code",[e._v("Upgrade")]),e._v(" module facilitates upgrading from major version X to major version Y. To accomplish this, node operators must first upgrade their current binary to a new binary that has a corresponding "),a("code",[e._v("Handler")]),e._v(" for the new version Y. It is assumed that this version has fully been tested and approved by the community at large. This "),a("code",[e._v("Handler")]),e._v(" defines what state migrations need to occur before the new binary Y can successfully run the chain. Naturally, this "),a("code",[e._v("Handler")]),e._v(" is application specific and not defined on a per-module basis. Registering a "),a("code",[e._v("Handler")]),e._v(" is done via "),a("code",[e._v("Keeper#SetUpgradeHandler")]),e._v(" in the application.")]),e._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("type")]),e._v(" UpgradeHandler "),a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("func")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("(")]),e._v("Context"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(",")]),e._v(" Plan"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(")")]),e._v("\n")])])]),a("p",[e._v("During the execution of each "),a("code",[e._v("BeginBlock")]),e._v(", the "),a("code",[e._v("Upgrade")]),e._v(" module checks whether there is a "),a("code",[e._v("plan")]),e._v(" that should be executed (the block height or time when the "),a("code",[e._v("BeginBlock")]),e._v(" is scheduled to run). If it exists, execute the corresponding "),a("code",[e._v("processor")]),e._v(". If the plan is expected to be executed but the corresponding processor is not registered, or the binary upgrade is too early, the node will gracefully panic and exit.")]),e._v(" "),a("h3",{attrs:{id:"storeloader"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#storeloader"}},[e._v("#")]),e._v(" StoreLoader")]),e._v(" "),a("p",[e._v("The "),a("code",[e._v("Upgrade")]),e._v(" module also facilitates store migrations as part of the upgrade. The "),a("code",[e._v("StoreLoader")]),e._v(" sets the migrations that need to occur before the new binary can successfully run the chain. This "),a("code",[e._v("StoreLoader")]),e._v(" is also application specific and not defined on a per-module basis. Registering this "),a("code",[e._v("StoreLoader")]),e._v(" is done via "),a("code",[e._v("app#SetStoreLoader")]),e._v(" in the application.")]),e._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("func")]),e._v(" UpgradeStoreLoader "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("(")]),e._v("upgradeHeight "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("int64")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(",")]),e._v(" storeUpgrades "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("*")]),e._v("store"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(".")]),e._v("StoreUpgrades"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(")")]),e._v(" baseapp"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v(".")]),e._v("StoreLoader\n")])])]),a("p",[e._v("If there's a planned upgrade and the upgrade height is reached, the old binary writes "),a("code",[e._v("UpgradeInfo")]),e._v(" to the disk before panic'ing.")]),e._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("type")]),e._v(" UpgradeInfo "),a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("struct")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("{")]),e._v("\n    Name    "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("string")]),e._v("\n    Height  "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("int64")]),e._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("}")]),e._v("\n")])])]),a("p",[e._v("This information is critical to ensure the "),a("code",[e._v("StoreUpgrades")]),e._v(" happens smoothly at correct height and expected upgrade. It eliminates the chances for the new binary to execute "),a("code",[e._v("StoreUpgrades")]),e._v(" multiple times everytime on restart. Also if there are multiple upgrades planned on same height, the "),a("code",[e._v("Name")]),e._v(" will ensure these "),a("code",[e._v("StoreUpgrades")]),e._v(" takes place only in planned upgrade handler.")]),e._v(" "),a("p",[e._v("Currently in the upgrade process, the state transition methods involved mainly support the following three types: "),a("code",[e._v("Renamed")]),e._v(", "),a("code",[e._v("Deleted")]),e._v(", and "),a("code",[e._v("Added")])]),e._v(" "),a("h4",{attrs:{id:"renamed"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#renamed"}},[e._v("#")]),e._v(" Renamed")]),e._v(" "),a("p",[e._v("The user can specify to migrate all data under oldKey (prefix) to storage under newKey (prefix) during the upgrade process.")]),e._v(" "),a("h4",{attrs:{id:"deleted"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#deleted"}},[e._v("#")]),e._v(" Deleted")]),e._v(" "),a("p",[e._v("Users can delete all data under the specified key (prefix) during the upgrade process.")]),e._v(" "),a("h4",{attrs:{id:"added"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#added"}},[e._v("#")]),e._v(" Added")]),e._v(" "),a("p",[e._v("Users can apply for a new storage area with the specified key as the prefix during the upgrade process.")]),e._v(" "),a("h3",{attrs:{id:"proposal"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#proposal"}},[e._v("#")]),e._v(" Proposal")]),e._v(" "),a("p",[e._v("Typically, a "),a("code",[e._v("Plan")]),e._v(" is proposed and submitted through governance via a "),a("code",[e._v("SoftwareUpgradeProposal")]),e._v(". This proposal prescribes to the standard governance process. If the proposal passes,\nthe "),a("code",[e._v("Plan")]),e._v(", which targets a specific "),a("code",[e._v("Handler")]),e._v(", is persisted and scheduled. The upgrade can be delayed or hastened by updating the "),a("code",[e._v("Plan.Time")]),e._v(" in a new proposal.")]),e._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("type")]),e._v(" SoftwareUpgradeProposal "),a("span",{pre:!0,attrs:{class:"token keyword"}},[e._v("struct")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("{")]),e._v("\n    Title       "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("string")]),e._v("\n    Description "),a("span",{pre:!0,attrs:{class:"token builtin"}},[e._v("string")]),e._v("\n    Plan        Plan\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("}")]),e._v("\n")])])]),a("h4",{attrs:{id:"governance-process"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#governance-process"}},[e._v("#")]),e._v(" Governance Process")]),e._v(" "),a("p",[e._v("When a upgrade proposal is accepted, the upgrade process is devided into two steps.")]),e._v(" "),a("h5",{attrs:{id:"stop-network-consensus"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#stop-network-consensus"}},[e._v("#")]),e._v(" Stop network consensus")]),e._v(" "),a("p",[e._v("After the software upgrade proposal is accepted, the system will perform pre-upgrade preparations at the designated height of the BeginBlock stage, including downloading the upgrade plan and suspending network consensus.")]),e._v(" "),a("h6",{attrs:{id:"download-the-upgrade-plan"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#download-the-upgrade-plan"}},[e._v("#")]),e._v(" Download the upgrade plan")]),e._v(" "),a("p",[e._v("In order to be able to upgrade the software smoothly, the information required for the upgrade must be recorded before stopping the network consensus: "),a("code",[e._v("plan name")]),e._v(", "),a("code",[e._v("upgrade height")]),e._v(".")]),e._v(" "),a("p",[e._v("-"),a("code",[e._v("Plan name")]),e._v(": When the network restarts, it needs to be routed to the corresponding "),a("code",[e._v("UpgradeHandler")]),e._v(" and "),a("code",[e._v("UpgradeStoreLoader")]),e._v(" according to the plan name.\n-"),a("code",[e._v("Upgrade height")]),e._v(": When the network restarts, check whether the network upgrade plan is required.")]),e._v(" "),a("h6",{attrs:{id:"suspend-network-consensus"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#suspend-network-consensus"}},[e._v("#")]),e._v(" Suspend network consensus")]),e._v(" "),a("p",[e._v("After the software upgrade proposal is accepted, the system will gracefully suspend the network consensus at the specified height of the "),a("code",[e._v("BeginBlock")]),e._v(" stage.")]),e._v(" "),a("h5",{attrs:{id:"restart-the-new-software"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#restart-the-new-software"}},[e._v("#")]),e._v(" Restart the new software")]),e._v(" "),a("p",[e._v("The user replaces the software with the specified version and restarts the network. The system will detect whether the "),a("code",[e._v("processor")]),e._v(" specified by the "),a("code",[e._v("plan name")]),e._v(" is included. If it is included, the system first executes the "),a("code",[e._v("Handler")]),e._v(" program, and then starts the network consensus. If not, the system report an error and exit.")]),e._v(" "),a("h4",{attrs:{id:"cancelling-upgrade-proposals"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#cancelling-upgrade-proposals"}},[e._v("#")]),e._v(" Cancelling Upgrade Proposals")]),e._v(" "),a("p",[e._v("The upgrade proposal can be cancelled. There is a proposal type of "),a("code",[e._v("Cancel Software Upgrade (CancelSoftwareUpgrade)")]),e._v(". When this type of proposal is voted through, the currently ongoing upgrade plan will be removed. Of course, this needs to be voted and executed before the upgrade plan is executed.")]),e._v(" "),a("p",[e._v("If the current upgrade plan has been executed, but there are problems with the upgrade plan, then the proposal of "),a("code",[e._v("Cancel Software Upgrade")]),e._v(" is invalid (because the network has stopped consensus). At this time, there is another solution to make up for this mistake, which is to use the "),a("code",[e._v("--unsafe-skip-upgrades")]),e._v(" parameter to skip the specified upgrade height when restarting the network (not really skip the height, but jump via software upgrade "),a("code",[e._v("Handler")]),e._v("). Of course, this requires that 2/3 of the validators participating in the consensus perform the same operation, otherwise the network consensus cannot be reached.")]),e._v(" "),a("h2",{attrs:{id:"upgrade-process"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#upgrade-process"}},[e._v("#")]),e._v(" Upgrade process")]),e._v(" "),a("h3",{attrs:{id:"submit-an-upgrade-proposal"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#submit-an-upgrade-proposal"}},[e._v("#")]),e._v(" Submit an upgrade proposal")]),e._v(" "),a("p",[e._v("The first step in the implementation of the software upgrade process is to initiate a software upgrade proposal by the governance module. The proposal details the upgrade height or time. For details, see the above "),a("a",{attrs:{href:"#Concepts"}},[e._v("Concept")]),e._v(". An example of the command line to initiate a proposal is as follows:")]),e._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("plugchaind tx gov submit-proposal software-upgrade "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("plan-name"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--deposit")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("deposit"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  --upgrade-time "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("upgrade-time"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--title")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("title"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  --upgrade-info "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("upgrade-info"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--description")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("description"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v("  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--from")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("from"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  --chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--fees")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),e._v("20uplugcn "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-b")]),e._v(" block\n")])])]),a("h3",{attrs:{id:"deposit-and-vote-for-the-proposal"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#deposit-and-vote-for-the-proposal"}},[e._v("#")]),e._v(" Deposit and vote for the proposal")]),e._v(" "),a("p",[e._v("The execution process of the software upgrade proposal is basically the same as that of other ordinary proposals. Both validators and delegators are required to comment on the proposal. For specific information, please refer to "),a("RouterLink",{attrs:{to:"/features/governance.html"}},[e._v("governance module")]),e._v(". An example of the command line to deposit the proposal is as follows:")],1),e._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("plugchaind tx gov deposit "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("proposal-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("deposit"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--from")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("from"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" --chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--fees")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),e._v("20uplugcn "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-b")]),e._v(" block "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-y")]),e._v("\n")])])]),a("p",[e._v("Once the deposit amount reaches the minimum deposit amount, the proposal will enter the voting period, and the validator or delegator needs to vote on the proposal. An example of the command line to initiate a vote is as follows:")]),e._v(" "),a("div",{staticClass:"language-bash extra-class"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("plugchaind tx gov vote "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("proposal-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("option"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--from")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("from"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" --chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("chain-id"),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("--fees")]),a("span",{pre:!0,attrs:{class:"token operator"}},[e._v("=")]),e._v("20uplugcn "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-b")]),e._v(" block "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-y")]),e._v("\n")])])]),a("p",[e._v("When the software upgrade proposal is passed, the upgrade module will create an upgrade plan to stop all nodes from the network consensus at a specified height or time, and wait for the new software to restart the network.")]),e._v(" "),a("h3",{attrs:{id:"restart-the-network"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#restart-the-network"}},[e._v("#")]),e._v(" Restart the network")]),e._v(" "),a("p",[e._v("When the upgrade proposal is passed and the network reaches the specified upgrade block height or time, all nodes will stop producing blocks. and the user needs to download the source code and compile the new software according to the new version information specified in the first step "),a("a",{attrs:{href:"#submit-an-upgrade-proposal"}},[e._v("Submit Upgrade Proposal")]),e._v(", refer to "),a("RouterLink",{attrs:{to:"/get-started/install.html"}},[e._v("Install")]),e._v(". After the new software is installed, restart the node with the new version, and the node will execute the upgrade logic corresponding to the plan name. Once the voting power of the entire network exceeds 2/3 and restarts the network using the new version, the blockchain network will re-reach a new consensus and continue to produce new blocks.")],1)])}),[],!1,null,null,null);a.default=r.exports}}]);