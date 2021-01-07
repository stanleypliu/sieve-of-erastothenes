require_relative '../lib/ruby/sieve'
require 'benchmark/ips'

new_test = Sieve.new(100)

Benchmark.ips do |benchmark|
  benchmark.report('default') { new_test.primes }
  benchmark.report('ruby inbuilt') { new_test.built_in_method }

  benchmark.compare!
end
