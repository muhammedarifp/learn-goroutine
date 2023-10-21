package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}

	wg.Add(2)
	go PrintSomthing5times("Two", wg)
	go PrintSomthing5times("One", wg)

	wg.Wait()
}

func PrintSomthing5times(text any, wg *sync.WaitGroup) {
	str, status := text.(string)
	if !status {
		log.Fatal("convertion error")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}

	wg.Done()
}
