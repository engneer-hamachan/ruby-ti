class Hoge
  attr_reader :a, :b

  def initialize(x:, y:)
    @a = x
    @b = y
  end
end

h = Hoge.new(x: 1, y: '2')
dbtp h.b
