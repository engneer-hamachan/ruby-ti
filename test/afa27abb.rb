class Hoge
  attr_accessor :xxx
  def test2
    self.xxx = 1
  end

  def test
    1
  end
end

h = Hoge.new

dbtp h.test2
dbtp h.xxx
