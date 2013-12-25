require 'pry'
class Proverb

  attr_reader :args

  def initialize(*args)
    @args = *args
  end

  def to_s
    #args.map{|arg| phrase(arg, arg)}.join("\n")
    song = args.each_with_index.map do |arg, i|
      phrase(arg, args[i+1])
    end.join("\n")
    song << "\n"
    song << "And all for the want of a #{args[0]}"
    return song
  end

  def phrase(first, second)
    "For want of a #{first} the #{second} was lost."
  end

end
