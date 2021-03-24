package main

import (
	"fmt"
	"study/socket"
)

func main() {

	fmt.Println("Hello Word! Golang")

	// 读取输入内容
	// commands.UseSanc()
	// commands.UseBufio()

	// TCP监听
	err := socket.ListenTCP("127.0.0.1:9099")
	fmt.Println("error:", err)
}
