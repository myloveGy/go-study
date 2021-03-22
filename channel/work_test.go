package channel

import (
	"fmt"
	"testing"

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

func TestDigitsSum(t *testing.T) {
	assert.Equal(t, int64(1), DigitsSum(1))
	assert.Equal(t, int64(3), DigitsSum(12))
	assert.Equal(t, int64(6), DigitsSum(123))
	assert.Equal(t, int64(10), DigitsSum(19))
}
