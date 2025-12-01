## Intro
旨在打造一个适合 Emaldo Golang 服务端体质的，快速启动的项目脚手架。

## Quick Start
1. Clone 仓库
> git clone https://github.com/ykds/scaffold.git
> 
> cd scaffold && go mod tidy

1. 修改配置文件, 补充数据库连接信息
> cp config-sample.yaml config.yaml
```yaml
// Server 配置
server:
  debug: true         // 是否 debug 模式，目前只用于控制 gin 的 Mode
  port: ":8080"       // 监听端口
  read_timeout: 10    // http 读超时(s)
  write_timeout: 10   // http 写超时(s)

// 日志配置
logger:
  mode: file           // 打印模式。控制台: console  文件: file
  level: info          // 日志等级。debug/info/error/panic/fatal
  filename: "err.log"  // 日志文件名，可带上目录
  max_size: 10         // 单文件最大上限(MB)
  max_age: 3           // 文件保留时间(天)
  compress: false      // 是否压缩归档日志文件
  max_backups: 2       // 备份日志文件数

// Redis 配置
redis:
  host: "localhost"
  port: "6379"
  password: ""
  db: 0

// TDEngine 配置
taos:
  protocal: http        // 使用协议。tcp/http
  host: "localhost"
  port: "6041"
  username: ""
  password: ""
  db_name: ""
  max_open_conns: 100
  conn_max_life_time_second: 3600
  max_idle_conns: 10
  slow_sql_millis: 1000 // 慢日志阈值(ms)

// Mongodb 配置
mongo:
  hosts: "localhost:27003"
  username: ""
  password: ""
  db_name: ""
  repl_name: ""
```

## Run
1. 无 TDEngine 环境
> go run -tags=taosRestful main.go

2. 有 TDEngine 环境
> go run main.go

## Test
> curl http://localhost:8080/demo?name=emaldo


## Design
### 一、config
这个包主要放置配置相关，配置初始化的代码

### 二、errors
这个包用于实现内部业务错误，包含业务错误码与业务错误信息。

这个包封装了`errors`和`pkg/errors`一些常用的方法，避免业务代码中同时使用`errors`和`pkg/errors`。

`code.go`用于放置预先定义的业务错误，如下：
```
var (
	Success       = NewError(200, "Success")
	BadParameters = NewError(400, "Bad Parameters")
	Unauthorized  = NewError(401, "Unauthorized")
	InternalError = NewError(500, "Internal Error")
)

// xx业务: 1xxx
var (
    AErr = NewError(1000, "A")
    BErr = NewError(1001, "B")
)

// yy业务: 2xxx
var (
    CErr = NewError(2000, "C")
    DErr = NewError(2001, "D")
)
```
当发生对应业务错误时，需要返回这时定义的错误，便于快速定位错误。

`errors.go`实现了自定义`Error`，用于定义业务错误。

`Error`结构体在设计上，使用非导出字段，通过方法来获取值，避免在使用`Error`时修改其错误码与错误信息。

同一个错误码对应一个错误实例，不能重复定义。

### 三、 middleware
用于放置中间件，一个中间件一个文件。

### 四、response
http响应工具的封装。

当使用`Error`进行相应时，同时会将错误日志信息打印。

### 五、safego
`go`关键词的安全封装，避免在直接使用`go`时，发生`panic`导致程序退出。

`safego`内部做好了`recover`，可放心使用。

### 六、pkg
主要放置公共组件和工具，与任何业务模块都无直接相关。

### 七、internal
`internal`在这里的设计目的主要是为了将`repository`,`service`,`handler`放在里面。

因为在开发的过程中，我们大部分时间只关注这 3 个目录。

将这 3 个目录单独放在一个目录下，而不是在最顶层与其他目录混杂在一起，减少我们对其他目录的注意力。

#### handler
放置每个模块的控制层代码。

`handler.go`用于注册所有的`Handler`，新增加一个`Handler`时要注册到这里。

每个模块的`handler`目录都有两个基本文件，如`demo`模块，有`demo.go`和`router.go` 2个文件。

`demo.go`用于初始化控制层和写各个接口的代码，`router.go`用于注册这个`Handler`下的路由。

#### service
放置每个模块的服务层代码，这层的代码为主要业务逻辑。

`server.go`的作用同`handler.go`，用于注册所有`Service`。

当业务逻辑比较简单时，放在一个代码文件即可，如果业务代码比较多或者比较复杂，可以添加代码文件。

比如`demo`模块，当业务逻辑比较少和简单时，都写在这个文件即可；当需要拆分时，再在这下面根据场景进行文件拆分。

#### repository
放置每个模块的存储层代码。

同样的，查询简单时，代码放在一个代码文件即可，需要时也可以像`service`中说的一样处理。

### 代码生成
为了减少工作量和统一代码结构，提供了 `handler`、`service`、`repository` 3层代码的生成，目前只支持新增一个模块时生成代码。

代码生成的工具位于`cmd/codegen`下，生成方式:
> make gen module="" // 以驼峰格式填写新增的模块名

如：
> make gen module="Test"

默认会在`internal`下生成如下代码文件：
> handler/test/test.go
```
package test

import (
	"scaffold/internal/service"
)

type TestHandler struct {
	testSvc *service.TestService
}

func NewTestHandler(testSvc *service.TestService) *TestHandler {
	return &TestHandler{
		testSvc: testSvc,
	}
}

```

> handler/test/routergo
```
package test

import "github.com/gin-gonic/gin"

func (test *TestHandler) Name() string {
	return "test"
}

func (test *TestHandler) RegisterRouter(engine *gin.RouterGroup) {
	r := engine.Group("/test")
	{
		// define router here
	}
}
```

> service/test.go
```
package repository

import (
	"scaffold/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名
const testCol = "test"

// 表Model
type Test struct {
}

type TestRepository interface {
}

type testRepository struct {
	mgo *mongodb.Mongo
	col *mongo.Collection
}

func NewTestRepository(mgo *mongodb.Mongo) TestRepository {
	r := &testRepository{
		mgo: mgo,
		col: mgo.Database.Collection(testCol),
	}
	return r
}
```

> repository/test.go
```
package repository

import (
	"scaffold/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名
const testCol = "test"

// 表Model
type Test struct {
}

type TestRepository interface {
}

type testRepository struct {
	mgo *mongodb.Mongo
	col *mongo.Collection
}

func NewTestRepository(mgo *mongodb.Mongo) TestRepository {
	r := &testRepository{
		mgo: mgo,
		col: mgo.Database.Collection(testCol),
	}
	return r
}
```

生成后，只需要将`TestHandler`和`TestService`对应注册到`handler.go`和`service.go`即可。

接着只需要补充对应的业务代码即可。

实现存储层 --> 实现服务层 --> 实现控制层 --> 注册路由