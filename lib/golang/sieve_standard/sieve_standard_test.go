package sieve_standard

import "testing"

func TestPrimeSieve(t *testing.T) {
	PrimeSieve(10_000)
}

func BenchmarkPrimeSieve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PrimeSieve(25)
	}
}

var result []int

func BenchmarkCompletePrimeSieve(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = PrimeSieve(25)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
