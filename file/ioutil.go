package file

import (
	"io"
	"os"
)

func CopyFile(source, dest string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		if _, err = destFile.Write(buffer[:n]); err != nil {
			return err
		}

	}

	return nil
}

func CopyFileByIo(source, dest string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, file)
	return err
}
