---
order: 8
---



# 磁盘使用优化

自定义配置设置以降低验证器节点的磁盘要求

区块链数据库往往会随着时间的推移而增长，这取决于例如 在块速度和交易金额。

可以进行很少的配置来减少所需的磁盘使用率相当显着。 其中一些更改全面生效仅当您进行配置并从 start 开始同步时
他们在使用中。

## 索引

如果您不需要从特定节点查询交易，您可以禁用索引。 在 `config.toml` 设置

```toml
indexer = "null"
```

如果您在已同步的节点上执行此操作，则不会清除收集到的索引自动，您需要手动删除它。 该索引位于在名为“data/tx_index.db/”的数据库目录下。

## 状态同步快照

我相信这是默认情况下在 Plugchain 上禁用的。 在 `app.toml` 设置

```toml
snapshot-interval = 0
```

请注意，如果在网络上启用了状态同步并且工作正常，它将允许一个人在几分钟内同步一个新节点。 但是这个节点不会有历史。

## 配置 修剪

默认情况下，每 500 个状态，最后 100 个状态被保留。 这从长远来看会消耗大量磁盘空间，并且可以通过以下方式进行优化
以下自定义配置：

```toml
pruning = "custom"
pruning-keep-recent = "100"
pruning-keep-every = "0"
pruning-interval = "10"
```

配置 `pruning-keep-recent = "0"` 可能听起来很诱人，但这如果 `plugchaind` 因任何原因被杀死，将面临数据库损坏的风险。
因此，建议保留少数最新状态。

## 日志记录

默认情况下，日志记录级别设置为`info`，这会产生很多日志。 此日志级别在启动时可能会很好，以查看
节点开始正确同步。 但是，在您看到同步后
顺利进行，您可以将日志级别降低到`warn`或者`error`。 在
`config.toml` 设置如下

```toml
log_level = "warn"
```

还要确保您的日志轮换配置正确。

## 结果

以下是 Plugchai 测试网两周后的磁盘使用情况。 默认
配置导致磁盘使用量为 90GB。

```bash
5.3G    ./state.db
70G     ./application.db
20K     ./snapshots/metadata.db
24K     ./snapshots
9.0G    ./blockstore.db
20K     ./evidence.db
1018M   ./cs.wal
4.7G    ./tx_index.db
90G     .
```

此优化配置已将磁盘使用量减少到 17 GB。

```bash
17G     .
1.1G    ./cs.wal
946M    ./application.db
20K     ./evidence.db
9.1G    ./blockstore.db
24K     ./snapshots
20K     ./snapshots/metadata.db
5.3G    ./state.db
```
