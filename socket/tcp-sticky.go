package socket

import (
	"fmt"
	"net"
)

func StickyTCP(address, writeType string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {

		message := fmt.Sprintf(`hello world! i = %d`, i)

		if writeType == "sticky" {
			data, err := Encode(message)
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
