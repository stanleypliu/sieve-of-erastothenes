require_relative '../lib/sieve'
require 'benchmark/ips'

newTest = Sieve.new(100)

Benchmark.ips do |benchmark|
  benchmark.report("default") { newTest.primes }
end