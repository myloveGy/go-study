package channel

import (
	"fmt"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 3; i++ {
		group.Add(1)
		go Add(&group)
	}

	group.Wait()
	fmt.Println("x = ", x)
}

func TestAddLock(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 3; i++ {
		group.Add(1)
		go AddLock(&group)
	}

	group.Wait()
	fmt.Println("x = ", x)
}
