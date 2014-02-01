class Fizzbuzz < Struct.new(:max)
  MUTATIONS = {
    'fizzbuzz' => ->(i) {i % 3 == 0 && i % 5 == 0},
    'fizz' => ->(i) {i % 3 == 0},
    'buzz' => ->(i) {i % 5 == 0}
  }

  def fizzle
    1.upto(max).map do |i|
      mutations.find(-> {[i]}) do |k, v|
        v.call(i)
      end[0]
    end
  end

  private

  def mutations
    MUTATIONS
  end
end
