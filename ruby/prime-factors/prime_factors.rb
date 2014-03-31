class PrimeFactors

  def self.for(number, divisor=2, factors=[])
    return factors if number == 1

    if number % divisor == 0
      factors.push(divisor)
      number /= divisor
      self.for(number, divisor, factors)
    else
      divisor += 1
      self.for(number, divisor, factors)
    end
  end

end
