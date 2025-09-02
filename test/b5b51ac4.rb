class Hoge
  attr_reader :a, :b

  def initialize(x, y)
    @a = x
    @b = y
  end
end

h = Hoge.new 1, 2
dbtp h.a
