class PrimeFactors

  def self.for(number)
    divisor = 2
    prime_factors = []
    while number > 1
      if number % divisor == 0
        prime_factors << divisor
        number = number / divisor
      else
        divisor += 1
      end
    end
    prime_factors
  end

end
