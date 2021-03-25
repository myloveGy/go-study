package http

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	str, err := Get("http://localhost:8081?name=123")
	assert.NoError(t, err)
	data := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal([]byte(str), &data))
	fmt.Println(data)
}
