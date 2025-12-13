class Input
  def initialize x, y
  end

  def test
    1
  end
end

class Hoge
  def initialize
    @fuga = Input.new 1, 2
  end
end

i = Hoge.new

p i.fuga.test
dbtp i.fuga.test

