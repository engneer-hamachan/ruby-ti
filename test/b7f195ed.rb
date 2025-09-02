class Hoge
  attr_accessor :a

  def test
    @a = 1
  end
end

h = Hoge.new

h.a = 2
