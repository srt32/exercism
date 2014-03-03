class Gigasecond

  attr_reader :birthday

  def initialize(birthday)
    @birthday = birthday
  end

  def date
    birthday + (10**9/(60 * 60 * 24))
  end

end
