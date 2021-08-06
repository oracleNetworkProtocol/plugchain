## 节点类型

### 全节点
全节点是参与网络但无助于保护网络的节点。完整节点可用于存储区块链的整个状态。Tendermint 有两种形式的状态。首先是区块链状态，这代表区块链的区块。其次是Application状态，它代表事务修改的状态。交易如何修改状态的知识不是由 Tendermint 掌握的，而是由 ABCI 边界另一侧的应用程序掌握的。

作为全节点运营商，您正在为网络提供服务，帮助其达成共识，并帮助其他人赶上当前区块。即使完整节点只能帮助网络达成共识，但保护您的节点免受对抗性参与者的侵害也很重要。如果可能，我们建议使用防火墙和代理。运行一个完整的节点可能很容易，但它因网络而异。在运行节点之前验证您的应用程序文档。
### 验证者节点
- 验证器是参与网络安全的节点。验证者在 Tendermint 中拥有相关的权力，这种权力可以代表权益证明系统中的权益、权威证明中的声誉或任何类型的可衡量单位。运行安全且一致的在线验证器对于网络健康至关重要。验证器必须安全且容错，建议使用 2 个或更多哨兵节点运行验证器。
- 验证器节点应该有 `per=false`，这样它就不会影响整个网络。持久对等点将成为您的哨兵节点。由于验证器没有试图隐藏它与谁通信，所以可以将私有对等点留空。设置无条件对等点对于验证器是可选的，因为它们没有完整的地址簿。

作为验证者，您有可能减轻体重，这是由应用程序定义的。如果验证者应该增加或减少权重，应用程序会通知 Tendermint。应用程序具有不同类型的恶意行为，导致验证者权力的削减。请检查您将运行的应用程序的文档以找到更多信息。
### 哨兵节点
哨兵节点几乎在所有方面都类似于完整节点。不同之处在于哨兵节点将拥有一个或多个私有节点。这些对等点可能是验证者或网络中的其他完整节点。哨兵节点旨在为您的验证器提供一层安全性，类似于防火墙与计算机的工作方式。
### 种子节点
种子节点为节点提供节点可以连接到的对等点列表。启动节点时，您必须至少提供一种类型的节点才能连接到所需的网络。通过提供种子节点，您将能够快速填充您的地址。种子节点不会保留为对等节点，但会在提供对等节点列表后与您的节点断开连接。

**注意** : 
如果需要共享节点以便其他人连接，则需要设置`external_address="节点公网IP"`

#### 验证者节点配置:
|  Config Option   | Setting  |
|  :----:  | :----:  |
| pex  | false |
| persistent_peers  | list of sentry nodes |
| private_peer_ids  | none |
| unconditional_peer_ids  | optionally sentry node IDs |
| addr_book_strict  | false |
| double_sign_check_height  | 10 |


#### 哨兵节点配置:
|  Config Option   | Setting  |
|  :----:  | :----:  |
| pex  | true |
| persistent_peers  | validator node, optionally other sentry nodes |
| private_peer_ids  | validator node ID |
| unconditional_peer_ids  | validator node ID, optionally sentry node IDs |
| addr_book_strict  | false |


#### 种子节点配置:
| Config Option | Setting |
| :----: | :----: |
| addr_book_strict  | true |
| pex  | true |
| seed_mode | true |

