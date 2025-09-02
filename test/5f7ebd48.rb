class Hoge
  attr_reader :hog

  def test
    test2
  end

  private

  def test2
    "1"
  end
end

h = Hoge.new

h.hog

h.hoge
