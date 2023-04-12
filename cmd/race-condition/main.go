package main

import (
	"fmt"
)

func main() {
	x := 42
	go func() {
		x = 21 * 2
	}()
	go func() {
		x = 84 / 2
	}()
	fmt.Println(x)
}
