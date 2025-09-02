class Hoge
  attr_reader :a

  def test
    self
  end
end

h = Hoge.new
p h.a
a = h.test

a + 1
