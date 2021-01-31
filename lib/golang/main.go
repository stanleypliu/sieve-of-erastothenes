package main

import (
	"fmt"
  "github.com/stanleypliu/sieve-of-eratosthenes/lib/golang/sieve_standard"
	"github.com/stanleypliu/sieve-of-eratosthenes/lib/golang/sieve"
)

func main() {
	fmt.Printf(`Select an Implementation to run:
	A - The official Golang sieve 
	B - My own implementation of the sieve
	Press any other key to quit.
	Note that both implementations will only output prime numbers up to 100`)

	var selection string

	fmt.Printf("\n")
	fmt.Scanln(&selection)

	if selection == "A" {
		fmt.Println(sieve_standard.PrimeSieve(25))
	} else if selection == "B" {
		fmt.Println(sieve.EratosthenesSieve(100))
	} 
	
	return
}
