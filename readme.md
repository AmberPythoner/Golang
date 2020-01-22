### 初始化配置
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### 函数传参数
1、go语言中的参数都是 值传递

### go mode 使用
1、更新依赖
```
go get [@v...]
go mod tidy
```
2、将项目迁移到 go mod
```
go mod init names
go build ./...
```


### 知识点
1、nil是一个 指针类型

2、import 的时候 加上 包名的 上一级文件名
3、使用者来定义接口


### 性能测试
1、go test .
2、go tool cover -html c.out    查看覆盖率
3、go test -bench .   性能测试

### go rounine
1、go语言的 协程 是非抢占式多任务处理

### WaitGroup使用
```cassandraql
外层函数
var wg sync.WaitGroup

wg.Add(20)

wg.Wait()

子程序
wg.Done()
```
