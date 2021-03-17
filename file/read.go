package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(source string) string {
	file, err := os.Open(source)
	if err != nil {
		return ""
	}
	defer file.Close()

	str, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}

	return string(str)
}

func ReadLine(source string) ([]string, error) {
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	list := make([]string, 0)
	buf := bufio.NewReader(f)
	for {
		// 可以使用 buf.ReadLine()
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			list = append(list, line)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return list, err
}

func ReadFileByIo(source string) ([]byte, error) {
	return ioutil.ReadFile(source)
}
