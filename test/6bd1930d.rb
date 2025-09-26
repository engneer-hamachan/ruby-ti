class Hoge
  def test2
    self.test
  end

  def test
    1
  end
end

h = Hoge.new

dbtp h.test2
