package udp

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Listen(address string) error {
	// 解析地址
	addr, err := ParseToUDPAddr(address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		data := make([]byte, 1024)
		n, address, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}

		// 读取数据
		message := string(data[:n])
		fmt.Printf("data: %v addr:%v count:%v\n", message, address, n)
		if _, err = conn.WriteToUDP(data[:n], address); err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}

func ParseToUDPAddr(address string) (*net.UDPAddr, error) {
	// 127.0.0.1:9098
	str := strings.Split(address, ":")
	if len(str) != 2 {
		return nil, errors.New("address error")
	}

	// 在次处理端口
	arr := strings.Split(str[0], ".")
	if len(arr) != 4 {
		return nil, errors.New("address ip error")
	}

	// 转换类型
	port, err := strconv.Atoi(str[1])
	if err != nil {
		return nil, fmt.Errorf("address port %s error: %s", str[1], err.Error())
	}

	var ip [4]byte
	for k, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("address ip:%s [%d]:[%v] error: %s", str[0], k, v, err.Error())
		}

		ip[k] = byte(i)
	}

	return &net.UDPAddr{
		IP:   net.IPv4(ip[0], ip[1], ip[2], ip[3]),
		Port: port,
	}, nil
}
