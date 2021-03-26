package util

import (
	"strings"
)

func Split(str, sep string) []string {

	if str == "" {
		return []string{""}
	}

	if sep == "" {
		return []string{str}
	}

	result := make([]string, 0)
	len := len(sep)

	for {
		index := strings.Index(str, sep)
		if index < 0 {
			break
		}

		result = append(result, str[0:index])
		str = str[index+len:]
	}

	result = append(result, str)
	return result
}

func Split2(str, sep string) []string {

	if str == "" {
		return []string{""}
	}

	if sep == "" {
		return []string{str}
	}

	sepLength := len(sep)
	searchLength := len(str)

	// 如果分隔字段大于查找字段、或者查找字段长度等于分隔字段长度，但不相等，那么直接返回查找字段
	if sepLength > searchLength || (searchLength == sepLength && str != sep) {
		return []string{str}
	}

	result := make([]string, 0)
	index := 0
	for i := 0; i < searchLength; i++ {
		// 一段一段匹配
		tmpIndex := i + sepLength
		if tmpIndex <= searchLength && str[i:tmpIndex] == sep {
			result = append(result, str[index:i])
			i += sepLength
			index = i
		}
	}

	if index <= searchLength {
		result = append(result, str[index:])
	}

	return result
}

func Snake(name string) string {
	array := []byte(name)
	bytes := make([]byte, 0)
	for k, v := range array {
		// 大写字母
		if 65 <= v && v <= 90 {
			if k != 0 {
				bytes = append(bytes, '_')
			}

			bytes = append(bytes, v+32)
		} else {
			bytes = append(bytes, v)
		}
	}

	return string(bytes)
}
