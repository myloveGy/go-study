package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Tailf(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	// 设置缓存渠道，出现错误的异常的话，通知主程序退出
	close := make(chan int, 1)

	go func() {
		tmp, err := f.Stat()
		if err != nil {
			close <- 1
			fmt.Printf("错误: 1 => %s\n", err.Error())
			return
		}

		t := tmp.ModTime()
		size := tmp.Size() // 文件大小
		for {

			fileStat, err := os.Stat(filename)
			if err != nil {
				close <- 2
				fmt.Printf("错误: 2 => %s\n", err.Error())
				return
			}

			// 监听文件修改时间变动
			if fileStat.ModTime().After(t) {
				t = fileStat.ModTime()
				if size > fileStat.Size() {
					size = fileStat.Size()
					_, _ = f.Seek(0, os.SEEK_END)
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 设置到文件末尾
	_, _ = f.Seek(0, os.SEEK_END)

	read := bufio.NewReader(f)
	for {
		select {
		case i := <-close:
			fmt.Println("监控出现问题:", i)
			return
		default:
			str, err := read.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					time.Sleep(1 * time.Second)
					continue
				}

				fmt.Printf("%#v\n", err)
				break
			}

			fmt.Println(strings.TrimSpace(string(str)))
		}
	}
}
