# 交叉编译
> 需要三个重要参数，使用`go env`即可查看环境变量
- GOOS:目标平台的操作系统（Windows、Linux、Mac）
- GOARCH:目标平台的体系架构（arm、x86）
- CGO_ENABLED:是否启用CGO，由于交叉编译不支持CGO所以要禁用

使用流程
```
# Windows下设置环境变量参数，只会在当前环境终端生效，目的是在Linux机器上跑Go程序
$Env:CGO_ENABLED=0;$Env:GOARCH="amd64";$Env:GOOS="linux"

# 编译，并且输出到APP文件
go build -o app.exe .

go build # 编译本目录，或者main.go文件
```


## 使用Go mod
一般的文件夹没有go.mod文件的话，可以使用`go mod init xxxx(mod名称)`来实现

## go mod tidy
依赖对其，删除没被使用的，下载引用但是本地没有的

```azure
module gomod

go 1.20
// 当前module依赖的包
require (
	//github.com/gin-gonic/gin v1.9.0
)

// 排除第三方包
exclude (

)

// 无法使用外网原始包，本地引用
// 当依赖包发生迁移，或者使用代理
// 使用本地包替换原始包
replace (
	//source latest => target latest
)

// 撤回，当版本发布出去，别人使用，但是原先版本撤回，retract是为了让别人无法引用
retract (

)
```

