package main

import (
	"fmt"

	"github.com/stanleypliu/sieve-of-eratosthenes/lib/golang/sieve"
)

func main() {
	newGenerator := sieve.EratosthenesSieve(100)
	fmt.Println(newGenerator)
}
