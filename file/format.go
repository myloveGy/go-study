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
)

var MapUnit = map[int]string{
	0: Bytes,
	1: KB,
	2: MB,
	3: GB,
	4: TB,
}

type Format struct {
	Size int64 // 文件大小
}

func NewFormat(size int64) *Format {
	return &Format{Size: size}
}

func (f *Format) Format() string {
	i := 0;
	size := float64(f.Size);
	for  size > 1024 {
		size = size / 1024
		i ++
	}

	s := fmt.Sprintf("%0.2f", size)
	v, ok := MapUnit[i]
	if ok {
		return fmt.Sprintf("%s%s", s, v)
	}

	return fmt.Sprintf("%s", s)
}
