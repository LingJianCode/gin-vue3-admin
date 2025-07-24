# 函数和类型注释

1. 安装Goomment插件

2. Ctrl+Shift+p打开user的setting.json文件
setting.json文件中添加如下两行
```
"functionTemplate": "// ${func_name} \n//\t@param ${param_type} \n//\t@return ${return_type} \n//  @author ${git_name} \n//  @date ${date}",
"typeTemplate": "// ${type_name} \n//  @author ${git_name} \n//  @date ${date}"
```

3. 使用
control + command + / (For windows: control + alt + /)

# 问题

## cannot use myvalidator.ValidateDateFormat (value of type func(fl "github.com/go-playground/validator".FieldLevel) bool) as "github.com/go-playground/validator/v10".Func value in argument to validate.RegisterValidationcompilerIncompatibleAssign

validate.RegisterValidation 方法期望的参数类型是 github.com/go-playground/validator/v10.Func，但你传入的 myvalidator.ValidateDateFormat 函数类型为 func(fl "github.com/go-playground/validator".FieldLevel) bool。这可能是因为包版本不一致或者导入路径不同所导致的。

## 未使用依赖注入

### 1. 并发安全的本质取决于服务是否有状态

无论是通过 package 级别全局变量（如 var UserServiceApp = new(SysUserService) ）还是在路由初始化时创建单例（ sysUserService := new(service.SysUserService) ），只要 `SysUserService` 保持 无状态设计 （当前实现中确实如此），两者在并发场景下表现完全一致—— 均不会导致数据竞争 。

### 2. 差异体现在工程化而非运行时
两种方式的核心区别在于 代码组织和可维护性 ：

| 维度|package全局变量|初始化单例注入|
| ---- | --------------- | -------------- |
| 依赖可见性 |隐式依赖，调用方无需声明|显式依赖，通过构造函数注入|
| 测试便利性 |无法替换为mock，测试耦合度高|可注入mock服务，支持隔离测试|
| 扩展性 |全局唯一实例，无法差异化配置|可根据需求创建多实例，支持定制|


### 3. 关键结论
- 无状态服务 ：两种方式在并发安全上等价，但后者更符合工程最佳实践
- 有状态服务 ：无论哪种方式 都会导致并发问题 ，需通过互斥锁或每个请求创建新实例解决
- 推荐实践 ：即使服务无状态，仍建议采用 依赖注入 模式，为未来功能扩展（如添加缓存、多租户支持）预留灵活性