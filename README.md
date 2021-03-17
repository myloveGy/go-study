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
或者
```shell
go test study/file -test.run -v TestFormat
```