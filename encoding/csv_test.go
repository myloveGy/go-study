package encoding

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteCsv(t *testing.T) {
	err := WriteCsv("./user.csv", []*User{
		{UserId: 1, UserName: "jinxing.liu", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(12 * time.Hour)},
		{UserId: 2, UserName: "jinxing.liu", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(12 * time.Hour)},
	})

	assert.NoError(t, err)
	os.Remove("./user.csv")
}

func TestReadCsv(t *testing.T) {
	err := WriteCsv("./user.csv", []*User{
		{UserId: 1, UserName: "jinxing.liu", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(12 * time.Hour)},
		{UserId: 2, UserName: "jinxing.liu", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now().Add(12 * time.Hour)},
	})

	assert.NoError(t, err)

	users, err := ReadCsv("./user.csv")
	assert.NoError(t, err)
	for k, v := range users {
		if k == 0 {
			assert.Equal(t, 1, v.UserId)
		}

		fmt.Printf("key = %d: %#v\n", k, v)
	}

	os.Remove("./user.csv")
}
