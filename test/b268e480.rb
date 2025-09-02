class Hoge
  attr_reader :fuga

  def test
    @fuga = 1
  end

  def test2
    @fuga
  end
end

h = Hoge.new

h.test2 + "1"
