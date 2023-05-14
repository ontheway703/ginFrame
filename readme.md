## 目录结构
| 目录         | 用途                 |
|------------|--------------------|
| constdef   | 常量                 |
| dao        | 数据访问层              |
| gerror     | 自定义错误码（game error） |
| handler    | 控制层                |
| idl        | 接口定义IDL            |
| middleware | 中间件层               |
| model      | 数据模型               |
| router     | 路由                 |
| service    | 服务逻辑层              |
| util       | 实用工具层              |

## 返回值
正常返回
```json
{
  "status_code":0,
  "status_msg":"",
  "data":"json object"
}
```
异常返回
```json
{
  "status_code":400,
  "status_msg":"xx错误",
  "data":""
}
```


## 其他
```
1. 不用IDL：引入外部依赖，增加复杂性
2. 统一处理Req：有利于确定参数，简化代码，增加统一打点
3. 统一处理Resp：简化代码，增加统一打点
```