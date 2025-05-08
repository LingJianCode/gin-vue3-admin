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