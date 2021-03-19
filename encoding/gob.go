package encoding

import (
	"encoding/gob"
	"errors"
	"os"
	"strconv"
	"time"
)

type User struct {
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserToString(str []string) (*User, error) {

	if len(str) != 5 {
		return nil, errors.New("输入数据格式错误")
	}

	var (
		u   = new(User)
		err error
	)

	u.UserId, err = strconv.Atoi(str[0])
	if err != nil {
		return nil, err
	}

	u.UserName = str[1]
	u.Status, err = strconv.Atoi(str[2])
	if err != nil {
		return nil, err
	}

	u.CreatedAt, err = time.Parse("2006-01-02 15:04:05", str[3])
	if err != nil {
		return nil, err
	}

	u.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", str[4])
	if err != nil {
		return nil, err
	}

	return u, nil
}

func WriteUser(filename string, users []*User) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := gob.NewEncoder(f)
	for _, v := range users {
		if err := writer.Encode(v); err != nil {
			return err
		}
	}

	return nil
}

func ReadUser(filename string) ([]*User, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data := make([]*User, 0)
	read := gob.NewDecoder(f)
	for {
		a := &User{}
		if err := read.Decode(a); err != nil {
			break
		}

		data = append(data, a)
	}

	return data, nil
}
