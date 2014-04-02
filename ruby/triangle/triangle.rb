class Triangle

  attr_reader :a, :b, :c

  def initialize(a, b, c)
    @a = a
    @b = b
    @c = c
    raise TriangleError if invalid_triangle?
  end

  def kind
    types.find { |type, definition| definition }.first
  end

  private

  def types
    {
      :equilateral => all_sides_equal?,
      :isosceles => two_sides_equal?,
      :scalene => true
    }
  end

  def all_sides_equal?
    a == b && a == c && b == c
  end

  def two_sides_equal?
    a == b || a == c || b == c
  end

  def invalid_triangle?
    [a, b, c].any? { |size| size <= 0 }
    a + b <= c || a + c <= b || b + c <= a
  end

end

class TriangleError < ArgumentError; end
