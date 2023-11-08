# OpenTDP Blackbox 公共接入服务

目前稳定性不足，故而只开放了部分能力的代码，仅供参考运作流程，或提交新的API能力。

## 未开源文件

- /api/join.php
- /vendor.php

## Nginx 示例配置

```nginx
server {

    listen 80;
    listen 443 ssl http2;

    server_name blackbox.opentdp.org;

    root  /var/www/blackbox;
    index index.php index.html;

    include http.d/server_fastcgi_php;

    if (!-e $request_filename) {
        rewrite ^/api/(\w+)/?$ /api/$1.php last;
        rewrite ^/(\w+)/?$ /app/$1.html last;
    }

}
```
