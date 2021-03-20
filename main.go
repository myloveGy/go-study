package main

import (
	"fmt"
	"study/commands"
)

func main() {
	fmt.Println("Hello Word! Golang")
	var a string
	fmt.Scanln(&a)
	fmt.Println("您输入的为:", a)

	// 读取输入内容
	commands.UseSanc()
	commands.UseBufio()
}
