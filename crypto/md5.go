package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func FileMd5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()

	m := md5.New()
	str := make([]byte, 1024)
	for {
		n, err := f.Read(str)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		m.Write(str[:n])
	}

	return fmt.Sprintf("%x", m.Sum(nil)), nil

}
