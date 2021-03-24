package main

import (
	"flag"
	"fmt"
	"study/commands"
	"study/socket"
)

const (
	Command    = "command"        // 命令行
	TcpService = "tcp-service"    // tcp服务
	TcpClient  = "tcp-client"     // tcp连接
	TcpSticky  = "tcp-sticky"     // tcp 黏包测试
	TcpAddress = "127.0.0.1:9099" // tcp 连接地址
)

func main() {
	fmt.Println("Hello Word! Golang")
	var name, args string
	flag.StringVar(&name, "name", "", "请输入需要执行的命令")
	flag.StringVar(&args, "args", "", "请输入执行命令的参数")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("请输入需要执行的命令 -name=")
		flag.PrintDefaults()
	}

	if name == "" {
		flag.Usage()
		return
	}

	switch name {
	case Command:
		// 读取输入内容
		commands.UseSanc()
		commands.UseBufio()
	case TcpService:
		// TCP监听
		err := socket.ListenTCP(TcpAddress, args)
		fmt.Println("error:", err)
	case TcpClient:
		err := socket.ClientTCP(TcpAddress)
		fmt.Println("client tcp error:", err)
	case TcpSticky:
		err := socket.StickyTCP(TcpAddress, args)
		fmt.Println("client sticky tcp error:", err)
	}
}
