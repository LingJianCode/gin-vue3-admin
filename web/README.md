## 项目简介

基于 Vue3、Vite、JavaScript 和 Element-Plus 搭建的极简开箱即用企业级后台管理前端模板。在[vue3-element-admin-js](https://gitee.com/youlaiorg/vue3-element-admin-js)的基础上修改而来。
在开发中由于URI冲突问题，对字典项路由地址进行了修改。增加了api权限分配功能。

## 项目启动

- **环境准备**

| 环境类型       | 版本要求                     | 下载链接                     |
|----------------|-----------------------------|-----------------------------|
| **开发工具**   | Visual Studio Code (最新版) | [官方下载](https://code.visualstudio.com/Download) |
| **运行环境**   | Node.js 18.x (推荐18.16.1)  | [中文镜像](https://npmmirror.com/mirrors/node/v18.16.1/) |
> ⚠️ 注意：Node.js 20.6.0版本存在兼容性问题，请勿使用


- **快速开始**

```bash
# 克隆代码
git clone https://gitee.com/LingJianCode/my-ops-vue.git

# 切换目录
cd my-ops-vue

# 安装 pnpm
npm install pnpm -g

# 设置镜像源(可忽略)
pnpm config set registry https://registry.npmmirror.com

# 安装依赖
pnpm install

# 启动运行
pnpm run dev
```


## 项目部署

执行 `pnpm run build` 命令后，项目将被打包并生成 `dist` 目录。接下来，将 `dist` 目录下的文件上传到服务器 `/usr/share/nginx/html` 目录下，并配置 Nginx 进行反向代理。

```bash
pnpm run build
```

以下是 Nginx 的配置示例：

```nginx
server {
    listen      80;
    server_name localhost;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    # 反向代理配置
    location /prod-api/ {
        # 请将 backend 替换为您的后端 API 地址，并注意保留后面的斜杠 /
        proxy_pass http://backend/;
    }
}
```

## 本地Mock

项目同时支持在线和本地 Mock 接口，默认使用线上接口，如需替换为 Mock 接口，修改文件 `.env.development` 的 `VITE_MOCK_DEV_SERVER` 为  `true` **即可**。
