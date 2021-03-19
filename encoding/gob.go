package encoding

import (
	"encoding/gob"
	"os"
	"time"
)

type User struct {
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
