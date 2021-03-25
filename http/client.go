package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func Post(url string) (string, error) {

	// 发送POST请求
	response, err := http.Post(url, "appliction/json;charset=utf-8", strings.NewReader(`{"username":"jinxing.liu"}`))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func GetUrl(api string, data url.Values) string {
	query := data.Encode()
	
	// 判断不存在? 那么直接返回: http://localhost?username=123
	if !strings.Contains(api, "?") {
		return api + "?" + query
	}

	// 判断最后一个字符串是哪个，如果是?、& 直接返回
	last := api[len(api)-1:]
	if last == "?" || last == "&" {
		return api + query
	}

	// 剩下 http://localhost?u=123 这种情况需要添加 & 
	return	api + "&" + query
}
