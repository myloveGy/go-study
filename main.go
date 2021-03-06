package main

import (
	"flag"
	"fmt"
	"study/commands"
	"study/http"
	"study/socket/tcp"
	"study/socket/udp"
)

const (
	Command     = "command"        // 命令行
	TcpServer   = "tcp-server"     // tcp服务
	TcpClient   = "tcp-client"     // tcp连接
	TcpSticky   = "tcp-sticky"     // tcp 黏包测试
	TcpAddress  = "127.0.0.1:9099" // tcp 连接地址
	UdpAddress  = "127.0.0.1:9098" // udp 连接地址
	UdpServer   = "udp-server"     // upd 服务
	UdpClient   = "udp-client"     // upd 连接
	HttpServer  = "http-server"    // http 服务
	HttpAddress = ":8081"          // http 地址
)

func main() {
	fmt.Println("Hello Word! Golang")
	var name, args, port string
	flag.StringVar(&name, "name", "", "请输入需要执行的命令")
	flag.StringVar(&args, "args", "", "请输入执行命令的参数")
	flag.StringVar(&port, "port", "", "请输入HTTP监听端口")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("请输入需要执行的命令 -name=[xxx] -args=[xxx]")
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
	case TcpServer:
		// TCP监听
		err := tcp.Listen(TcpAddress, args)
		fmt.Println("Tcp Server error:", err)
	case TcpClient:
		err := tcp.Client(TcpAddress)
		fmt.Println("Tcp Client error:", err)
	case TcpSticky:
		err := tcp.Sticky(TcpAddress, args)
		fmt.Println("Tcp Client Sticky error:", err)
	case UdpServer:
		err := udp.Listen(UdpAddress)
		fmt.Println("Udp Server error:", err)
	case UdpClient:
		err := udp.Client(UdpAddress)
		fmt.Println("Udp Client error:", err)
	case HttpServer:
		if port == "" {
			port = HttpAddress
		}

		err := http.Listen(port)
		fmt.Println("Http Server error:", err)
	}
}
