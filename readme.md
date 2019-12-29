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

