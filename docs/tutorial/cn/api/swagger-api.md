---
title: API服务
---

#### 开通自己节点的对外api服务，需要修改数据目录（默认的数据目录 `~/.plugchain`）里的 config/app.toml 的 `enable`,`swagger` ,修改为 `true`
```
###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]

# Enable defines if the API server should be enabled.
enable = true

# Swagger defines if swagger documentation should automatically be registered.
swagger = true
```