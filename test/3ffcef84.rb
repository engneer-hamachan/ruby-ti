class Hoge
  attr_reader :a

  def test
    @a = 1
  end
end

h = Hoge.new

h.a = 2

