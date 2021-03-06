package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	assert.Equal(t, []string{"ab", "edb", ""}, Split("abcedbc", "c"))
	assert.Equal(t, []string{"jinxing.liu"}, Split("jinxing.liu", ""))
	assert.Equal(t, []string{"jinxing", "liu"}, Split("jinxing.liu", "."))
}

func TestSplit2(t *testing.T) {
	assert.Equal(t, []string{"ab", "edb", ""}, Split2("abcedbc", "c"))
	assert.Equal(t, []string{"jinxing.liu"}, Split2("jinxing.liu", ""))
	assert.Equal(t, []string{"jinxing", "liu"}, Split2("jinxing.liu", "."))
	assert.Equal(t, []string{"jinxing.liu"}, Split2("jinxing.liu", "abc"))
}

func TestSnake(t *testing.T) {
	assert.Equal(t, "user_id", Snake("UserId"))
	assert.Equal(t, "user_id", Snake("userId"))
	assert.Equal(t, "user_id", Snake("user_id"))
}
