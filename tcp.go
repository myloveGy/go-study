package main

import (
	"fmt"
	"study/socket"
)

func main() {
	err := socket.ClientTCP("127.0.0.1:9099")
	fmt.Println("client tcp error:", err)
}
