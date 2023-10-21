// Simple goroutines example using time.Sleep
// It is a first step for goroutines journy

package main

import (
	"fmt"
	"time"
)

func main() {
	go PrintGreet("Hello")
	PrintGreet("World")
}

func PrintGreet(greet string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", greet)
		time.Sleep(time.Second * 1)
	}
}
