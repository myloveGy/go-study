package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	str := ReadFile("./format.go")
	fmt.Println(str)
	assert.NotEqual(t, "", str)
}

func TestReadLine(t *testing.T) {
	str, err := ReadLine("./format.go")
	for _, v := range str {
		fmt.Println(v)
	}
	assert.NoError(t, err)
}

func TestReadFileByIo(t *testing.T) {
	str, err := ReadFileByIo("./format.go")
	assert.NoError(t, err)
	fmt.Println(string(str))
}
