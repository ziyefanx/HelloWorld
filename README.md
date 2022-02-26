# HelloWorld

```
.
├── README.md 
├── dal # 数据访问层，所有和数据交互的操作放到这里
│   └── db 数据库
│       ├── conn.go 连接数据库
│       └── student.go 对学生表的操作
├── go.mod
├── go.sum
├── grom.exe
├── handler # 网络接口层 DDD https://zhuanlan.zhihu.com/p/109114670
│   └── student # 每个领域对应一个子目录
│       ├── base.go # 该领域公用结构体
│       └── query.go # 具体的接口
├── main.go
├── model # 存放项目的公用结构体（包括数据模型，枚举）
│   └── student.go
├── service # 业务逻辑层，存放公共逻辑
└── util # 工具，存放一些业务无关的函数/逻辑
```