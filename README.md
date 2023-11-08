# OpenTDP Blackbox 集群服务

这将注册一个 Blackbox 节点到 OpenTDP Blackbox 服务，运行前请确认知晓自己在做什么。

## 注册须知

- 目前仅用于测试，请不要在生产环境中使用
- 你的节点将分享给其他注册节点的用户
- 你将可以使用其他用户注册的节点

## 如何加入节点

选择一种注册方法，修改变量后运行响应的代码，即可注册你的节点到 OpenTDP Blackbox 服务。

### 环境变量列表

切勿在变量中添加更多的`;`，否则可能导致无法正常注册。

- **NODE_NAME**：以自己英文名开头，仅支持英文小写、短横线、数字（不超过20byte）
- **NODE_OWNER**：所有者昵称（utf-8编码，不超过30byte）
- **NODE_REGION**：国家或地区/省份/城市（utf-8编码，不超过30byte）
- **NODE_ISP**：运营商/云厂商（utf-8编码，不超过30byte）
- **NODE_BANNER**：自定义说明（utf-8编码，不超过30byte）

### 使用二进制文件

[前往版本发布页](https://github.com/opentdp/blackbox/releases)下载合适的二进制文件。修改环境变量后运行响应的代码。

```shell
NODE_NAME=your-node-name
NODE_OWNER=your-nickname
NODE_REGION=your-city
NODE_ISP=your-isp
NODE_BANNER=your-banner

tdp-blackbox
```

### 使用Docker运行

参数 `--publish 9115:9115` 并不是必须的，这决于该节点是否仍需要为其它 Prometheus 提供服务，如果不需要请删除它。

```shell
docker run -d \
    --name blackbox \
    --restart always \
    --cap-add=NET_ADMIN \
    --publish 9115:9115 \
    --env "NODE_NAME=your-node-name" \
    --env "NODE_OWNER=your-nickname" \
    --env "NODE_REGION=your-city" \
    --env "NODE_ISP=your-isp" \
    --env "NODE_BANNER=your-banner" \
    rehiy/blackbox

# 查看注册日志
docker logs -f blackbox
```
