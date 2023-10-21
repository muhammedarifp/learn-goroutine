package main

import (
	"fmt"
	"sync"
)

func main() {

	score := []int{0}
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go ChangeScore(1, &score, wg)
	go ChangeScore(100, &score, wg)

	fmt.Println(score)
	wg.Wait()
}

func ChangeScore(score int, array *[]int, wg *sync.WaitGroup) {
	// mut := sync.Mutex{}
	// mut.Lock()
	*array = append(*array, score)
	// mut.Unlock()
	wg.Done()
}
