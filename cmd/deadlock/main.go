package main

func main() {
	ch := make(chan int)
	ch <- 42
}
