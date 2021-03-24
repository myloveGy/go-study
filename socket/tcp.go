package socket

import (
	"bufio"
	"fmt"
	"net"
)

func process(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			break
		}

		fmt.Printf("收到消息: %s\n", buf[:n])
		if _, err = connection.Write(buf[:n]); err != nil {
			fmt.Printf("write error: %v\n", err)
			break
		}
	}
}

func ListenTCP(address string) error {
	// 监听连接
	lister, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("tcp listen error: %v\n", err)
		return err
	}

	defer lister.Close()

	// 等待连接
	for {
		connection, err := lister.Accept()
		if err != nil {
			fmt.Printf("listen accept errror: %v\n", err)
			continue
		}

		go process(connection)
	}
}
