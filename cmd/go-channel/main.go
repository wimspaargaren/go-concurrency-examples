package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c1, done1 := createProducer(time.Second)
	c2, done2 := createProducer(4 * time.Second)

	go func() {
		for {
			select {
			case msg1, ok := <-c1:
				if ok {
					fmt.Println("received from producer 1, msg:", msg1)
				}
			case msg2, ok := <-c2:
				if ok {
					fmt.Println("received from producer 2, msg:", msg2)
				}
			}
		}
	}()

	fmt.Println("waiting for producers to shutdown")
	<-done1
	<-done2

	fmt.Println("gracefully exit the program")
}

func createProducer(timeout time.Duration) (<-chan int, <-chan (struct{})) {
	producerChannel := make(chan int)
	doneChannel := make(chan struct{})
	go produce(producerChannel, timeout, doneChannel)
	// Return as receive only channels
	return producerChannel, doneChannel
}

func produce(c chan int, timeout time.Duration, done chan (struct{})) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	i := 1
	for {
		select {
		case <-sigs:
			fmt.Printf("shutting down producer with timeout: %v\n", timeout)
			close(c)
			done <- struct{}{}
			return
		default:
			c <- i
			i++
			time.Sleep(timeout)
		}
	}
}
