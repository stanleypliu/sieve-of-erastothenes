# My own implementation of the Sieve of Eratosthenes algorithm in Ruby. 
#
# Doesn't perform optimally compared to the inbuilt method but useful for benchmarks.

require 'prime'

class Sieve
  def initialize(limit)
    @limit = limit
  end

  def built_in_method
    Prime::EratosthenesGenerator.new.take_while { |i| i <= @limit }
  end

  def primes
    test_range = range_to_test

    test_range.each do |n|
      next unless n
      break if n**2 > @limit

      (n**2).step(@limit, n) { |m| test_range[m] = nil}
    end

    test_range.compact
  end

  private

  def range_to_test
    # Initial range of numbers to test (from 2 up to the limit)
    range = (0..@limit).to_a
    range[0] = range[1] = nil
    range
  end
end

# Test
# puts Sieve.new(100).primes
# puts "=== Using the built-in Ruby method ==="
# puts Sieve.new(100).built_in_method
