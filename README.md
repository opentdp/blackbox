# OpenTDP Blackbox 集群服务

这将通过 Frp Client 注册一个 Blackbox 节点到 OpenTDP Blackbox 列表，运行前请确认知晓自己在做什么。

## 注册须知

- 目前仅用于测试，请不要在生产环境中使用
- ~~你的节点将分享给其他注册节点的用户~~
- ~~你将可以使用其他用户注册的节点~~

## 加入节点方法

修改环境变量后运行如下命令，注册你的节点到 OpenTDP Blackbox 列表。参数 `--publish 9115:9115` 并不是必须的，取决于该节点是否仍需要为其它 Prometheus 提供服务。

```shell
docker run -d \
    --name blackbox \
    --restart always \
    --publish 9115:9115 \
    --env "NODE_NAME=your-node-name" \
    --env "NODE_OWNER=your-name" \
    --env "NODE_REGION=your-city" \
    --env "NODE_ISP=your-isp" \
    --env "NODE_BANNER=your-banner" \
    rehiy/blackbox
# 查看日志
docker logs -f blackbox
```

## 环境变量说明

请不要在环境变量中添加任何标点符号或特殊字符，否则可能导致无法正常运行。

- NODE_NAME: 10-20符串，将用于注册子域名，支持数字、小写英文、短横线
- NODE_OWNER: 所有者昵称，支持数字、英文、中文（utf-8编码）
- NODE_REGION: 省份或城市，支持数字、英文、中文（utf-8编码）
- NODE_ISP: 运营商或云厂商，支持数字、英文、中文（utf-8编码）
- NODE_BANNER: 自定义说明，支持数字、英文、中文（utf-8编码）
