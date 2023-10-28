# OpenTDP Blackbox

这将通过 Frp Client 注册一个 Blackbox 节点到 OpenTDP Blackbox 列表，运行前请确认知晓自己在做什么。

## 注册须知

- 目前仅用于测试，请不要在生产环境中使用
- ~~你的节点将分享给其他注册节点的用户~~
- ~~你将可以使用其他用户注册的节点~~

## 环境变量

请不要环境变量中添加任何标点符号或特殊字符，否则可能导致无法正常运行。

- NODE_NAME: 10-20符串，将用于注册子域名，支持英文、数字
- NODE_OWNER: 所有者昵称，支持英文、数字、中文
- NODE_REGION: 省份或城市，支持英文、数字、中文
- NODE_ISP: 运营商或云厂商，支持英文、数字、中文
- NODE_BANNER: 自定义说明，支持英文、数字、中文

## 快速启动

修改环境变量后运行如下命令，注册你的节点到 OpenTDP Blackbox 列表。

```shell
docker run --name blackbox -d \
    --restart=always \
    --env "NODE_NAME=your-node-name" \
    --env "NODE_OWNER=your-name" \
    --env "NODE_REGION=your-city" \
    --env "NODE_ISP=your-isp" \
    --env "NODE_BANNER=your-banner" \
    rehiy/blackbox

docker logs -f blackbox
```
