package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	str := Md5("123456")
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", str)
	fmt.Println(str)
}

func TestFileMd5(t *testing.T) {
	str, err := FileMd5("../file/format.go")
	assert.NoError(t, err)
	fmt.Println(str)
	assert.Equal(t, "862086ee99e3ba08b9f828437f20259b", str)
}
