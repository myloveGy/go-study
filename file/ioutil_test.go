package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	err := CopyFile("./format.go", "./format.txt")
	assert.NoError(t, err)
	os.Remove("./format.txt")
	err = CopyFile("./format.go1", "./format.txt")
	fmt.Println(err)
	assert.Error(t, err)
}

func TestCopyFileByIo(t *testing.T) {
	err := CopyFileByIo("./format.go", "./format.txt")
	assert.NoError(t, err)
	os.Remove("./format.txt")

	err = CopyFileByIo("./format.go1", "./format.txt")
	fmt.Println(err)
	assert.Error(t, err)
}
