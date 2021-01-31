# Sieve Of Eratosthenes 

Sieve of Eratosthenes implemented in Ruby and Go as a monorepo, 
currently only benchmarked with the Benchmark gem.

Algorithm logic taken from [Wikipedia](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Pseudocode)

Further inspiration taken from [this gist](https://gist.github.com/loganhasson/8937903) 
for Ruby's inbuilt `EratosthenesGenerator` which exists in the Prime module, as well
as a more compact version of the algorithm in the aforementioned link.

For Golang, I adapted the [Go Playground](https://play.golang.org/p/9U22NfrXeq) version of the 
Sieve (package `sieve_standard`), whilst also adapting the Ruby implementation into Go (package `sieve`). 

## Requirements

- Ruby (v2.6.5)
- Go (v1.15)

`$ gem install benchmark-ips`

## Instructions

- The Ruby implementation can be run simply using `ruby spec/sieve_benchmark.rb`,
which, by default, outputs a comparison of the performance of the `primes` method 
versus `Prime::EratosthenesGenerator` as per the example shown in 
[benchmark/ips](https://github.com/evanphx/benchmark-ips#synopsis). 
- The Golang implementation can be run via `go run lib/golang/main.go`,
which brings up a simple UI to allow selection of the desired sieve algorithm.

For tests: `cd lib/golang/sieve` or alternatively `cd lib/golang/sieve_standard`,
and run `go test .`. 

## TODO

- Table driven tests for `sieve_standard`
- Modify `EratosthenesSieve` in `sieve.go` to use concurrency. 
- Look into Ractor for the Ruby method (which requires an upgrade to Ruby 3.0.0)