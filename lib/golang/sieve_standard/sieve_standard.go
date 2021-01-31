// Package sieve_standard employs concurrency to find the desired number of primes
// - taken directly from https://golang.org/doc/play/sieve.go
package sieve_standard

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generateIntegers(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// PrimeSieve: Daisy-chain Filter processes.
func PrimeSieve(n int) []int {
	ch := make(chan int)
	go generateIntegers(ch)

	primes := make([]int, n)

	for i := 0; i < n; i++ {
		prime := <-ch
		primes[i] = prime
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	return primes
}
