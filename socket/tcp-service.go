package socket

import (
	"bufio"
	"fmt"
	"net"
)

func process(connection net.Conn, readType string) {

	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		var (
			line string
			err  error
		)

		// 读取方式
		if readType == "sticky" {
			line, err = Decode(reader)
		} else {
			line, err = simpleRead(reader)
		}

		fmt.Printf("收到消息: %s\n", line)

		if err != nil {
			fmt.Printf("read error: %v\n", err)
			break
		}

		// 写入消息
		if _, err = connection.Write([]byte(line)); err != nil {
			fmt.Printf("write error: %v\n", err)
			continue
		}
	}
}

func simpleRead(reader *bufio.Reader) (string, error) {
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return "", err
	}

	return string(buf[:n]), nil
}

func ListenTCP(address, readerType string) error {
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

		go process(connection, readerType)
	}
}
