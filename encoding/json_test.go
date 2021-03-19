package encoding

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteJson(t *testing.T) {
	err := WriteJson("./user.json", []*User{
		{UserId: 1, UserName: "username", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
		{UserId: 2, UserName: "jinxing.liu", Status: 2, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
	})

	assert.NoError(t, err)
	os.Remove("./user.json")
}

func TestReadJson(t *testing.T) {
	err := WriteJson("./user.json", []*User{
		{UserId: 1, UserName: "username", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
		{UserId: 2, UserName: "jinxing.liu", Status: 2, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
	})

	assert.NoError(t, err)
	users, err := ReadJson("./user.json")
	assert.NoError(t, err)
	assert.Equal(t, true, len(users) > 0)
	for k, v := range users {
		fmt.Printf("%d: %#v\n", k, v)
	}

	os.Remove("./user.json")
}
