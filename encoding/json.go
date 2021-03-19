package encoding

import (
	"bufio"
	"encoding/json"
	"os"
)

func WriteJson(filename string, users []*User) error {
	// 打开文件，以追加方式添加
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	write := bufio.NewWriter(f)
	defer write.Flush()

	for _, v := range users {
		jsonUser, err := json.Marshal(v)
		if err != nil {
			return err
		}

		if _, err := write.Write(jsonUser); err != nil {
			return err
		}

		if err := write.WriteByte('\n'); err != nil {
			return err
		}
	}

	return nil
}

func ReadJson(filename string) ([]*User, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	users := make([]*User, 0)
	read := bufio.NewScanner(f)
	for read.Scan() {
		user := &User{}
		if err := json.Unmarshal(read.Bytes(), user); err == nil {
			users = append(users, user)
		}
	}

	return users, nil
}
