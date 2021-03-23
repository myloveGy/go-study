package channel

import (
	"sync"
	"sync/atomic"
)

var x int = 0

var mutex sync.Mutex

var data = make(map[string]string)

func get(name string) string {
	return data[name]
}

func set(name, value string) {
	data[name] = value
}

var atomicInt64 int64

func AddAtomic() {
	atomic.AddInt64(&atomicInt64, 1)
}

func LockGet(name string) string {
	mutex.Lock()
	defer mutex.Unlock()
	return data[name]
}

func LockSet(name, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	data[name] = value
}

func Add(group *sync.WaitGroup) {

	defer group.Done()

	for i := 0; i <= 50000; i++ {
		x += i
	}
}

func AddLock(g *sync.WaitGroup) {

	defer g.Done()

	for i := 0; i <= 50000; i++ {
		mutex.Lock()
		x += i
		mutex.Unlock()
	}
}
