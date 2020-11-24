# 配置文件

## Examples
- [descriptor](../examples/descriptor/config.yaml): 
自动生成[descriptors](../descriptors)包的代码

- [blog](../examples/blog/config.yaml):
生成`go-micro`微服务`handler`的代码

## plugins
- type: `[]string`

声明全局启用的插件，全部job都能使用插件提供的函数与变量.

## data
- type: `map[string]interface{}`

声明全局的变量.
支持模板解析.
与`plugin`交互的数据也写在这里

**注意** `map`是无序的，所以不要相互引用



## jobs
### job.name
- type: `string`

### job.plugins
参考全局的[plugins](#plugins)


### job.data
参考全局的[data](#data)


### job.loop
### job.if
### job.out 
### job.templatePath
