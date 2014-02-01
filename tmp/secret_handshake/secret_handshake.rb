require 'pry'
class SecretHandshake

  attr_reader :number

  def initialize(number)
    @number = number
  end

  def commands
    return [] unless number.class == Fixnum
    binary = number.to_s(2)
    negative = is_marker?
    result = binary.reverse.chars.map.with_index do |char, i|
      if char == "1" && i == 0
        number_commands[i+1]
      elsif char == "1"
        number_commands[2**i]
      end
    end.flatten.compact
    if negative
      result = result.reverse
    end
    return result
  end

  def is_marker?
    true unless binary.length < 5
  end

  def number_commands
    {
      1 => ["wink"],
      2 => ["double blink"],
      4 => ["close your eyes"],
      8 => ["jump"]
    }
  end

end
