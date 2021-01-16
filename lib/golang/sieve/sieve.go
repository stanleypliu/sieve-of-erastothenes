// Package sieve contains a basic implementation of the Sieve of Eratosthenes
// algorithm
package sieve

import (
	"math"
)

// setupSlice initialises a slice of numbers up to and including limit
// starting from 0
func setupSlice(limit int) []int {
	primes := make([]int, limit)

	for i := 0; i < limit; i++ {
		primes[i] = i
	}

	// Setting first two values to 0 as they aren't taken into account
	primes[0] = 0
	primes[1] = 0

	return primes
}

// primes removes all the 0s and compacts the slice
func compact(primes []int) []int {
	k := 0
	for _, n := range primes {
		if n != 0 {
			primes[k] = n
			k++
		}
	}

	return primes[:k]
}

// EratosthenesSieve contains the main logic (iterating through the array of
// numbers and checking for primality)
func EratosthenesSieve(limit int) []int {
	floatLimit := float64(limit)
	primes := setupSlice(limit)

	for _, n := range primes {
		squaredPrime := math.Pow(float64(n), 2)

		if n == 0 {
			continue
		}

		if squaredPrime > floatLimit {
			break
		}

		for j := int(squaredPrime); j < limit; j += n {
			primes[j] = 0
		}
	}

	return compact(primes)
}
