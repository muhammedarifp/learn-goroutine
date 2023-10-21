package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Genarate rand using math package

	// val := rand.Intn(5)
	// fmt.Println(val)

	// Use crypto package

	val, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(val.Add(val, big.NewInt(1)))
}
