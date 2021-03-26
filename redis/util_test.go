package redis

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Username  string
	Age       int
	Status    int
	CreatedAt time.Time
}

func TestMapStringToStruct(t *testing.T) {
	
	m := map[string]string{
		"username":   "jinxing.liu",
		"age":        "1",
		"status":     "2",
		"created_at": "2021-03-26 11:03:30",
	}

	data := &User{}
	err := MapStringToStruct(m, data)
	fmt.Println(err, data)
}
