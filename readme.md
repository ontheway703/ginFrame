## 目录
1. 项目结构
2. 运行 
3. 接口文档
4. 返回值
5. 


## 项目结构
| 目录         | 用途                 |
|------------|--------------------|
| constdef   | 常量                 |
| dao        | 数据访问层              |
| docs       | swagger文档              |
| gerror     | 自定义错误码（game error） |
| handler    | 控制层                |
| idl        | 接口定义IDL            |
| middleware | 中间件层               |
| model      | 数据模型               |
| service    | 服务逻辑层              |
| util       | 实用工具层              |

## 运行
```shell
go run main.go
```

## 接口文档
生成接口文档
```shell
make gen_docs
```
访问接口文档
```text
localhost:port/swagger/index.html
```

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