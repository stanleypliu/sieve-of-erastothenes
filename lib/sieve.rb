require 'prime'

class Sieve
  def initialize(limit)
    @limit = limit
  end

  def built_in_method
    Prime::EratosthenesGenerator.new.take_while {|i| i <= @limit }
  end

  def primes
    test_range = range_to_test
    
    factors.each do |factor|
      next unless test_range[factor]
      
      non_prime_numbers(factor).each do |non_prime|
        test_range[non_prime] = nil
      end
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

  # Get a list of factors up to the rounded (to nearest lowest integer) square root of @limit
  def factors 
    factors = []
    x = 2

    until x > Math.sqrt(@limit).floor
      factors << x 
      x += 1
    end

    factors
  end

  def non_prime_numbers(factor)
    # Numbers which aren't prime
    non_prime_numbers = []
    y = factor**2
    multiplier = 1

    until y > @limit
      non_prime_numbers << y
      y = factor**2 + (factor * multiplier)
      multiplier += 1
    end

    non_prime_numbers
  end
end

# Test
puts Sieve.new(100).primes
puts "=== Using the built-in Ruby method ==="
puts Sieve.new(100).built_in_method