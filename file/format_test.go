package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFormat(t *testing.T) {
	assert.Equal(t, NewFormat(1), &Format{Size: 1})
}

func TestFormat(t *testing.T) {
	format := NewFormat(1025)
	assert.Equal(t, "1.00KB", format.Format())
	fmt.Println(format.Format())
	format = NewFormat(1024 * 5 * 1024)
	fmt.Println(format.Format())
	assert.Equal(t, "5.00MB", format.Format())
	format = NewFormat(1024 * 5 * 1024 * 1024)
	fmt.Println(format.Format())
	assert.Equal(t, "5.00GB", format.Format())
	format = NewFormat(1024 * 5 * 1024 * 1024 * 1024)
	fmt.Println(format.Format())
	assert.Equal(t, "5.00TB", format.Format())
	format = NewFormat(10)
	fmt.Println(format.Format())
	assert.Equal(t, "10.00Bytes", format.Format())
}