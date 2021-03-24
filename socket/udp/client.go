package udp

import (
	"fmt"
	"net"
)

func Client(address string) error {
	// 解析地址
	addr, err := ParseToUDPAddr(address)
	if err != nil {
		return err
	}

	// 建立连接
	client, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("UDP Client Dial error:", err)
		return err
	}

	defer client.Close()

	if _, err := client.Write([]byte("Hello UDP Server")); err != nil {
		fmt.Println("UDP Client write error:", err)
		return err
	}

	data := make([]byte, 1024)
	n, remoteAddr, err := client.ReadFromUDP(data)
	if err != nil {
		fmt.Println("UDP Client Read error:", err)
		return err
	}

	fmt.Printf("UDP Client Read: %s addr: %v count:%v\n", string(data[:n]), remoteAddr, n)
	return nil
}
