## Golang model 示例

前置要求：

1. Golang 版本大于 1.11
2. 开启 Go mod

步骤:
```
export GO111MODULE=on #开启 go mod
export GOPROXY=https://goproxy.cn,direct #设置国内镜像代理

mkidr go-module-example && cd go-module-example 
go mod init go-module-example #创建项目
go get github.com/gin-gonic/gin #安装第三方库
go run main.go #运行项目
```

相关文件

`go.mod`  #记录依赖库版本
`go.sum`  #记录依赖库版本和 hash 值

一般情况下这两个文件都要添加到 git 仓库


Go mod 相关命令：

go help mod 查看帮助
go mod init <项目模块名称>初始化模块，会在项目根目录下生成 go.mod 文件。
go mod tidy 根据 `go.mod` 文件来处理依赖关系。
go mod vendor 将依赖包复制到项目下的 vendor 目录。
go list -m all 显示依赖关系。
go list -m -json all 以 json 格式显示详细依赖关系。
go list -m -versions <path>显示包有哪些已发布版本
go mod download <path@version>下载依赖。参数<path@version>不是非必须的，path 是包的路径，version 是包的版本。
其它命令可以通过go help mod来查看。
