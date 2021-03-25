# Golang 语言学习项目

各个目录为学习的🌰，然后有对应的测试

使用的 `go mod` 初始化项目依赖

```shell
go mod tidy
```

测试例子

```shell
go test -timeout 60s study/file -run ^TestReadFileByIo$ -v
```

测试覆盖率
```shell
go test -v -short -covermode=count -coverprofile=cover.out 
```

```shell
go tool cover -html=cover.out -o cover.html 
````