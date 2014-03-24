class Fixnum

  def to_roman
    result = ""
    if self == 9
      result << "IX"
    elsif (5..8).include?(self)
      result << "V" + "I" * (self - 5)
    elsif self == 4
      result << "IV"
    else
      result << "I" * self
    end
    result
  end

end
