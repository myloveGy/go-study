package encoding

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteUser(t *testing.T) {
	err := WriteUser("./user.gob", []*User{
		{UserId: 1, UserName: "username", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
		{UserId: 2, UserName: "jinxing.liu", Status: 2, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
	})

	assert.NoError(t, err)
	os.Remove("./user.gob")
}

func TestReadUser(t *testing.T) {

	// 先写入
	err := WriteUser("./user.gob", []*User{
		{UserId: 1, UserName: "username", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
		{UserId: 2, UserName: "jinxing.liu", Status: 2, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(time.Hour * 12)},
	})

	assert.NoError(t, err)

	// 后读出
	data, err := ReadUser("./user.gob")
	assert.NoError(t, err)
	for _, v := range data {
		fmt.Printf("%#v\n%s \n", v, v.UserName)
	}

	os.Remove("./user.gob")
}
