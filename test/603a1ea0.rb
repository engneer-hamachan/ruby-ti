class Hoge
  attr_reader :hoge, fuga

  def test
    test2
  end

  private

  def test2
    "1"
  end
end

h = Hoge.new

h.test + 1
