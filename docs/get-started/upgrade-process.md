---
order: 8
---


# Upgrade operation

This project elaborates on the upgrade plan and method for newly added community nodes and running community nodes
## Scheme 1 Upgrade to the latest height one by one

:::warning
Current plugchain version upgrade status v0.5.0 (762880)-> v0.6.0 -> +++
:::
1. Online node scheme
```shell
# After the upgrade community proposal takes effect, after the upgrade height is reached (762880), the node stops coming out, and the binary files need to be upgraded manually

# Pull v0.6.0 version code
git fetch origin v0.6.0

# Execute the compiled binary file
make install

# Print the version to confirm whether it is successful
plugchaind version

# After success, kill -9 pid to kill the original process of plugchaind, and then start its own node. When the verifier node upgrades more than 2/3, the node will generate a block
```
2. New node scheme
```shell
# According to the above prompt, the newly added node needs to use the v0.5.0 binary file to synchronize the block first to reach the upgrade height, and then use the v0.6.0 binary file to upgrade to synchronize the latest block

# Generate and compile the specified version of the binary
git checkout v0.5.0 && make install

# Print version view
plugchaind version

# Then generate your own data directory and synchronize the blocks
```


When the block reaches the upgrade height (762880)
```
ERR UPGRADE "x/token" NEEDED at height: 762880:
2:04AM ERR CONSENSUS FAILURE!!! err="UPGRADE \"x/token\" NEEDED at height: 762880: "
```

You need to upgrade with the upgraded version one by one. At present, it is v0.6.0. If the local code does not have tag v0.6.0, you need to execute `git fetch origin v0.6.0` to pull the remote warehouse code, and then execute

```
git fetch origin v0.6.0 && make install

# Use the v0.6.0 binary file to start the node for synchronization
```



## Scheme 2 Download and upgrade the current version of the block data, use the latest binary file plugchaind to start the node

This solution is used for newly added nodes to prevent multiple upgrades, the newly added nodes need to be upgraded step by step to reach the latest height

```shell
# Through the official version block data, download the version you need to the node server according to your needs

# Then according to the downloaded block data version, use the upgraded binary file of the next version to initialize the node, then decompress the data package to cover the data directory of the data directory, then configure the creation file and seed information, and start the node for synchronization area Piece

# For example, the official block data version is the block data before v0.6.0 upgrade
# 1. The version of the binary file that needs to be compiled is v0.6.0
# 2. Use the v0.6.0 version to initialize the node, configure the genesis file and seed information, then overwrite the data directory, and start the node.
```