package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	err := Write("./format.txt", "my name liujinxing")
	assert.NoError(t, err)
	err = Write("./format.txt", "my name jinxing.liu")
	assert.NoError(t, err)
	os.Remove("./format.txt")
}

func TestWriteString(t *testing.T) {
	err := WriteString("./format.txt", "my name jinxing.liu")
	assert.NoError(t, err)
	assert.Equal(t, "my name jinxing.liu\n", ReadFile("./format.txt"))
	os.Remove("./format.txt")
}

func TestWriteByIoutil(t *testing.T) {
	err := WriteByIoutil("./format.txt", "username=jinxing.liu&data=1")
	assert.NoError(t, err)
	assert.Equal(t, "username=jinxing.liu&data=1", ReadFile("./format.txt"))
	fmt.Println(ReadFile("./format.txt"))
	os.Remove("./format.txt")
}

func TestWriteByBufio(t *testing.T) {
	err := WriteByBufio("./format.txt", "name=jinxing.liu&data=1")
	assert.NoError(t, err)
	assert.Equal(t, "name=jinxing.liu&data=1", ReadFile("./format.txt"))
	os.Remove("./format.txt")
}
