class Hoge
  attr_accessor :a, :b
end

h = Hoge.new

h.a, h.b = 1, '2'

h.a + 1
h.b + 1
