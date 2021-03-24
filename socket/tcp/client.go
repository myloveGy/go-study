package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ClientTCP(address string) error {

	// 建立连接
	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("dial tcp: %s\n", address)
		return err
	}

	defer connection.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入内容:")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read tcp error: %v\n", err)
			break
		}

		// 判断退出
		input = strings.TrimSpace(input)
		if input == "exit" || input == "q" {
			fmt.Println("程序退出了")
			break
		}

		fmt.Println("读到输入数据:", input)

		// 写入数据
		if _, err := connection.Write([]byte(input)); err != nil {
			fmt.Printf("写入数据失败: %v\n", err)
			break
		}

		// 读取数据
		buf := [1024]byte{}
		n, err := connection.Read(buf[:])
		if err != nil {
			fmt.Printf("读取数失败: %v\n", err)
			break
		}

		fmt.Printf("读到响应数据: %s\n", buf[:n])
	}

	return nil
}
