class Hoge
  def test
    1
  end
end

def test(x, y)
  "1" + x
end

h = Hoge.new

test(h.test, 2)
