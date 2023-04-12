package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	amountOfTasks := 100
	c := make(chan (int), amountOfTasks)
	wg := &sync.WaitGroup{}

	// Create 10 workers
	createWorkers(c, wg, 10)

	publishTasks(c, amountOfTasks)

	fmt.Println("wait until tasks are processed")
	wg.Wait()
	fmt.Println("tasks done, exiting program")
}

func publishTasks(c chan<- (int), amountOfTasks int) {
	// Publish tasks to channel
	for i := 0; i < amountOfTasks; i++ {
		c <- i
	}
	close(c)
}

func createWorkers(c <-chan (int), wg *sync.WaitGroup, amount int) {
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go worker(c, wg, i+1)
	}
}

func worker(c <-chan (int), wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		val, ok := <-c
		if !ok {
			return
		}
		now := time.Now()
		longRunningTask()
		fmt.Printf("worker %d processed long running task %d in %s\n", i, val, time.Since(now))
	}
}

func longRunningTask() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(2000) * int(time.Millisecond)))
}
