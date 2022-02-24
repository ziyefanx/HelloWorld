# 结合gorm+gin

提供四个接口：
1. 插入
2. 更改
3. 删除
4. 查询（ID+Name+Sex）


postman -> gin -> gorm -> DB

1. 定义每个接口的请求&返回&URL&Http文档
2. 实现并验证

查询接口
URL：localhost:8888/api/stduent/get
Method: Get
Req
```json
{
    "id":121212,
    "search":{
        "name":"xxx",
        "sex": 1
    }
}
```
Response
```json
{
    "code": 200,
    "message": "success",
    "data":[
        {
            "id" : 1,
            "name" : "xxx",
            "sex" : 1
        }
    ]
}
```