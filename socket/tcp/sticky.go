package tcp

import (
	"fmt"
	"net"
	"study/socket"
)

func Sticky(address, writeType string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {

		message := fmt.Sprintf(`hello world! i = %d`, i)

		if writeType == "sticky" {
			data, err := socket.Encode(message)
			if err != nil {
				fmt.Println("encode message error:", err)
			}

			message = string(data)
		}

		if _, err := conn.Write([]byte(message)); err != nil {
			return err
		}
	}

	return nil
}
