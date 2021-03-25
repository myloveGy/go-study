package http

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// 成功
	str, err := Get("http://local.verystar.net/data/index?name=123")
	assert.NoError(t, err)
	data := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal([]byte(str), &data))
	fmt.Println(data)

	// 错误
	_, err = Get("http://localhost:90023")
	assert.Error(t, err)
}

func TestPost(t *testing.T) {
	str, err := Post("http://local.verystar.net/data/index?name=123")
	assert.NoError(t, err)
	fmt.Println(str)
	_, err = Post("http://localhost:90023")
	assert.Error(t, err)
}

func TestGetUrl(t *testing.T) {
	data := url.Values{}
	data.Add("username", "jinxing.liu")
	assert.Equal(t, "http://localhost?u=123&username=jinxing.liu", GetUrl("http://localhost?u=123", data))
	assert.Equal(t, "http://localhost?u=123&username=jinxing.liu", GetUrl("http://localhost?u=123&", data))
	assert.Equal(t, "http://localhost?username=jinxing.liu", GetUrl("http://localhost", data))
	assert.Equal(t, "http://localhost?username=jinxing.liu", GetUrl("http://localhost?", data))
}
