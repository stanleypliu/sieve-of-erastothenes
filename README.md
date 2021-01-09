# Sieve Of Eratosthenes 

Sieve of Eratosthenes implemented in Ruby, Go and Elixir as a monorepo, 
currently only benchmarked with the Benchmark gem.

Algorithm logic taken from [Wikipedia](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Pseudocode)

Further inspiration taken from [this gist](https://gist.github.com/loganhasson/8937903) 
for Ruby's inbuilt `EratosthenesGenerator` which exists in the Prime module, as well
as a more compact version of the algorithm in the aforementioned link.

## Requirements

- Ruby (v2.6.5)
- Go (v1.15)

`$ gem install benchmark-ips`

## Instructions

- The Ruby implementation can be run simply using `ruby spec/sieve_benchmark.rb`,
which, by default, outputs a comparison of the performance of the `primes` method 
versus `Prime::EratosthenesGenerator` as per the example shown in 
[benchmark/ips](https://github.com/evanphx/benchmark-ips#synopsis). 