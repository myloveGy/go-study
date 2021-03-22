package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChanel(t *testing.T) {

	ch := make(chan int)

	// ch <- 1 写入在消费之前的话，会造成阻塞
	go func() {
		fmt.Println(<-ch)
	}()

	ch <- 1

	fmt.Println("HaHa")
}

func TestChanelWaitGroup(t *testing.T) {

	var group sync.WaitGroup

	for i := 0; i < 10; i++ {
		group.Add(1)

		// 异步执行
		go func(num int) {
			defer group.Done()
			time.Sleep(time.Second)
			fmt.Println(num)
		}(i)
	}

	// 等待异步执行完成
	group.Wait()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "所有事情处理都完了")
}

func TestDigitsSum(t *testing.T) {
	assert.Equal(t, int64(1), DigitsSum(1))
	assert.Equal(t, int64(3), DigitsSum(12))
	assert.Equal(t, int64(6), DigitsSum(123))
	assert.Equal(t, int64(10), DigitsSum(19))
}

func TestWorker(t *testing.T) {
	job := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, job, results)
	}

	// 开启5个任务
	for i := 0; i < 5; i++ {
		job <- i
	}
	close(job)

	for i := 0; i < 5; i++ {
		num := <-results
		fmt.Println("num = ", num)
	}
	// for i := range results {
	// 	fmt.Println("num = ", i)
	// }
}
