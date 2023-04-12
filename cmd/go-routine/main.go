package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello from main thread")
	go func() {
		for {
			fmt.Println("hello from goroutine")
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("main thread sleeping for 10 seconds")
	time.Sleep(time.Second * 10)
	fmt.Println("shutting down...")
}
