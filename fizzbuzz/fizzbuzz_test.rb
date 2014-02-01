require 'minitest'
require 'minitest/autorun'
require 'minitest/pride'
require_relative 'fizzbuzz'

class FizzbuzzTest < Minitest::Test
  def test_it_accepts_a_max_point
    fizzbuzz = Fizzbuzz.new(1)
    assert_equal fizzbuzz.max, 1
  end

  def test_it_handles_upto_2
    fizzbuzz = Fizzbuzz.new(2)
    assert_equal fizzbuzz.fizzle, [1,2]
  end

  def test_it_returns_fizz_for_3
    fizzbuzz = Fizzbuzz.new(3)
    assert_equal fizzbuzz.fizzle, [1, 2, 'fizz']
  end

  def test_it_returns_fizz_for_5
    fizzbuzz = Fizzbuzz.new(5)
    assert_equal fizzbuzz.fizzle, [1, 2, 'fizz', 4, 'buzz']
  end

  def test_it_returns_fizz_for_10
    fizzbuzz = Fizzbuzz.new(10)
    assert_equal fizzbuzz.fizzle, [1, 2, 'fizz', 4, 'buzz', 'fizz', 7, 8, 'fizz', 'buzz']
  end

  def test_it_returns_fizz_for_15
    fizzbuzz = Fizzbuzz.new(15)
    assert_equal fizzbuzz.fizzle, [1, 2, 'fizz', 4, 'buzz', 'fizz', 7, 8, 'fizz', 'buzz', 11, 'fizz', 13, 14, 'fizzbuzz']
  end
end
