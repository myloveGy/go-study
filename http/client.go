package http

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	// 发送请求
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	// 读取内容信息
	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(str), nil
}
