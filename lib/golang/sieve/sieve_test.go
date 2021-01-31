package sieve

import "testing"

var sieveTests = []struct {
	n        int   // input
	expected []int // expected result
}{
	{10, []int{2, 3, 5, 7}},
	{50, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
	{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
}

func TestSieve(t *testing.T) {
	for _, tt := range sieveTests {
		actual := EratosthenesSieve(tt.n)
		if len(actual) != len(tt.expected) {
			t.Errorf("EratosthenesSieve(%d): Wrong output slice length", tt.n)
		} else {
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("EratosthenesSieve(%d): expected %d, actual %d", tt.n, tt.expected, actual)
				}
			}
		}
	}
}

func BenchmarkEratosthenesSieve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EratosthenesSieve(100)
	}
}

var result []int

func BenchmarkCompleteEratosthenesSieve(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = EratosthenesSieve(100)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
