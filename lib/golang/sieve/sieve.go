// Package sieve contains a basic implementation of the Sieve of Eratosthenes
// algorithm
package sieve

import (
	"math"
)

// setupSlice initialises a slice of numbers up to and including limit
// starting from 0
func setupSlice(limit int) []int {
	integers := make([]int, limit)

	// Keeping first two values as 0 as they aren't taken into account
	for i := 2; i < limit; i++ {
		integers[i] = i
	}

	return integers
}

// compact removes all the 0s and compacts the slice
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
	sqrtLimit := math.Sqrt(floatLimit)
	integers := setupSlice(limit)

	for _, n := range integers {
		if n == 0 {
			continue
		}

		floatN := float64(n)

		if floatN > sqrtLimit {
			break
		}

		squared := math.Pow(floatN, 2)

		for j := int(squared); j < limit; j += n {
			integers[j] = 0
		}
	}

	return compact(integers)
}
