package encoding

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func WriteCsv(filename string, users []*User) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	bufioWrite := csv.NewWriter(file)
	defer bufioWrite.Flush()
	if err := bufioWrite.Write([]string{"用户ID", "用户名称", "用户状态", "创建时间", "修改时间"}); err != nil {
		return err
	}

	for _, v := range users {
		if err := bufioWrite.Write([]string{
			strconv.Itoa(v.UserId),
			v.UserName,
			strconv.Itoa(v.Status),
			v.CreatedAt.Format("2006-01-02 15:04:05"),
			v.UpdatedAt.Format("2006-01-02 15:04:05"),
		}); err != nil {
			return err
		}
	}

	return nil
}

func ReadCsv(filename string) ([]*User, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data := make([]*User, 0)
	read := csv.NewReader(f)
	i := 0
	for {
		arr, err := read.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if i != 0 {
			user, err := NewUserToString(arr)
			if err != nil {
				return nil, err
			}

			data = append(data, user)
		}

		i++
	}

	return data, nil
}
