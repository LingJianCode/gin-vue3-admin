# 项目介绍

基于 GIN + GORM + Casbin 的后台框架。当时想使用[vue3-element-admin-js](https://gitee.com/youlaiorg/vue3-element-admin-js)前端模板，但是没找到go后端，故自己参考接口文档实现go后端。在开发中由于URI冲突问题，对字典项路由地址进行了修改。配套前端地址[my-ops-vue](https://gitee.com/LingJianCode/my-ops-vue)。目前仅支持PostgreSQL数据库。

# 项目二次开发


## 启动
```shell
go run main.go
```

## 导入sql文件
```shell
sudo -u postgres psql -d mypg -f ./sql/my-ops-admin-format-p.sql
```