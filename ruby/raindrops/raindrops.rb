class Raindrops

  def convert(number)
    result = droplet_noises.select { |_, verifier| verifier.call(number) }.keys.join
    result = number.to_s if result.empty?
    result
  end

  private

  def droplet_noises
    {
      "Pling" => ->(number) { number.divides_evenly_by?(3) },
      "Plang" => ->(number) { number.divides_evenly_by?(5) },
      "Plong" => ->(number) { number.divides_evenly_by?(7) }
    }
  end

end

class Fixnum

  def divides_evenly_by?(divisor)
    self % divisor == 0
  end

end
