package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("some go routine with a long running task")
		time.Sleep(time.Second)
		fmt.Println("goroutine done")
	}(wg)
	fmt.Println("waiting for goroutine")
	wg.Wait()
	fmt.Println("goroutine done, exiting the program")
}
