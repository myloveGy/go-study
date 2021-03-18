package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func Write(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write([]byte(content)); err != nil {
		return err
	}

	if _, err = f.Write([]byte{'\n'}); err != nil {
		return err
	}

	return nil
}

func WriteString(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.WriteString(f, content+"\n")
	return err
}

func WriteByIoutil(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), os.ModePerm)
}

func WriteByBufio(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	bufioWrite := bufio.NewWriter(f)
	_, err = bufioWrite.WriteString(content)
	if err != nil {
		return err
	}

	return bufioWrite.Flush()
}
