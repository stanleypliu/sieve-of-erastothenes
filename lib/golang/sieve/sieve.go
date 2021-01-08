// Package sieve contains a basic implementation of the algorithm (closely
// matching the algorithm in lib/ruby/sieve.rb) and is not intended to be
// highly performant!
package sieve

import (
	"math"
)

// EratosthenesSieve contains the main logic (iterating through the array of
// numbers up to the limit and checking for primality)
func EratosthenesSieve(limit int) []int {
	floatLimit := float64(limit)
	primes := make([]int, limit)

	for i := 0; i < limit; i++ {
		primes[i] = i
	}

	// Setting first two values to 0 as they aren't taken into account
	primes[0] = 0
	primes[1] = 0

	for i := range primes {
		if primes[i] == 0 {
			continue
		}
		if math.Pow(float64(primes[i]), 2) > floatLimit {
			break
		}

		for j := int(math.Pow(float64(primes[i]), 2)); j < limit; j += primes[i] {
			primes[j] = 0
		}

	}

	k := 0
	for _, n := range primes {
		if n != 0 {
			primes[k] = n
			k++
		}
	}

	primes = primes[:k]

	return primes
}
