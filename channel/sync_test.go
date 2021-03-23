package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddAtomic(t *testing.T) {

	// 使用原子性
	g := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		g.Add(1)
		go func() {
			defer g.Done()
			AddAtomic()
		}()
	}

	g.Wait()
	fmt.Println("atomic", atomicInt64)

	// 不使用原子性
	atomicInt64 = 0
	group := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			atomicInt64 += 1
		}()
	}

	group.Wait()
	fmt.Println("atomic", atomicInt64)
}

func TestSyncMap(t *testing.T) {
	g := sync.WaitGroup{}
	data := sync.Map{}
	for i := 0; i < 100; i++ {
		g.Add(1)
		go func(i int) {
			defer g.Done()
			name := fmt.Sprintf("name_%v", i)
			data.Store(name, i)
			value, ok := data.Load(name)

			fmt.Printf("i = %d name = %s value = %v ok = %v\n", i, name, value, ok)
		}(i)
	}

	g.Wait()
}

func TestMapLock(t *testing.T) {
	g := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		g.Add(1)
		go func(i int) {
			defer g.Done()
			LockSet(fmt.Sprintf("%v", i), "123")
			fmt.Println("i = ", i, "name = ", LockGet(fmt.Sprintf("%d", i)))
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	g.Wait()
}

func TestMap(t *testing.T) {
	g := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		g.Add(1)
		go func(i int) {
			defer g.Done()
			set(fmt.Sprintf("%v", i), "123")
			fmt.Println("name", get(fmt.Sprintf("%d", i)))
		}(i)
	}

	g.Wait()
}

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
