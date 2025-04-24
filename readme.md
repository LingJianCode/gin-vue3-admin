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