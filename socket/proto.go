package socket

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	writer := new(bytes.Buffer)

	// 写入消息头
	if err := binary.Write(writer, binary.LittleEndian, length); err != nil {
		return nil, err
	}

	// 写入消息体
	if err := binary.Write(writer, binary.LittleEndian, []byte(message)); err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return "", err
	}

	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	if err := binary.Read(lengthBuff, binary.LittleEndian, &length); err != nil {
		return "", err
	}

	// buffered 返回缓存中现在有可读的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", nil
	}

	// 读取真正的消息信息
	pack := make([]byte, int(4+length))
	if _, err := reader.Read(pack); err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}
