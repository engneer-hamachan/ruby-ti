class Point
  def +(other)
    Point.new(@x + other.x, @y + other.y)
  end
end

p1 = Point.new(1, 2)
p2 = Point.new(3, 4)
p3 = p1 + p2

dbtp p3
