package main

import (
	"fmt"
	"sync"
)

type mutexObject struct {
	sync.Mutex

	Accessed int
}

func (m *mutexObject) incrementAccess() {
	defer m.Unlock()
	m.Lock()
	m.Accessed++
}

func main() {
	mo := &mutexObject{}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go accessObject(mo, wg, "1")
	go accessObject(mo, wg, "2")

	wg.Wait()
	fmt.Printf("mutex object accessed %d times\n", mo.Accessed)
}

func accessObject(mo *mutexObject, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("thread %s access: %d\n", name, i+1)
		mo.incrementAccess()
	}
}
