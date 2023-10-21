package main

import (
	"fmt"
	"sync"
)

func main() {
	myChann := make(chan int, 4)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(myChann)

	go func(ch chan int) {
		ch <- 10
		ch <- 20
		close(ch)
		ch <- 30
		ch <- 40
		wg.Done()
	}(myChann)

	wg.Wait()
}
