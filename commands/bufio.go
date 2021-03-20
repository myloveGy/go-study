package commands

import (
	"bufio"
	"fmt"
	"os"
)

func UseSanc() {
	fmt.Println("请输入内容")
	var s string
	n, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("出现错误:", err.Error())
	}

	fmt.Println("您输入的内容为:", s, n)
}

func UseBufio() {
	fmt.Println("请输入内容")
	reader := bufio.NewReader(os.Stdout)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取出现错误:", err)
	}

	fmt.Println("输入内容为:", s)
}
