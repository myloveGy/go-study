package file

import (
	"fmt"
)

const (
	Bytes = "Bytes"
	KB    = "KB"
	MB    = "MB"
	GB    = "GB"
	TB    = "TB"
	PB    = "PB"
)

var Unit = []string{Bytes, KB, MB, GB, TB, PB}

type Format struct {
	Size int64 // 文件大小
}

func NewFormat(size int64) *Format {
	return &Format{Size: size}
}

func (f *Format) Format() string {
	i := 0
	size := float64(f.Size)
	for size > 1024 && i < len(Unit) {
		size = size / 1024
		i++
	}

	return fmt.Sprintf("%0.2f%s", size, Unit[i])
}
