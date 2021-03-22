package channel

import (
	"fmt"
	"time"
)

func DigitsSum(num int64) int64 {
	var sum int64 = 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return sum
}

func worker(id int, job <-chan int, result chan<- int) {
	for i := range job {
		fmt.Printf("worker: %d start job: %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker: %d end job: %d\n", id, i)
		result <- i * 2
	}
}
