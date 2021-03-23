package channel

import "sync"

var x int = 0

var mutex sync.Mutex

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
