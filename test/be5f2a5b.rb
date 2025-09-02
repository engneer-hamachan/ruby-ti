class Hoge
  def test(x)
    1 + x
  end
end

class Fuga
  def test(x)
    1
  end
end


def test(x)
  x.test('1')
end

h = Hoge.new
f = Fuga.new

test(f)
test(h)
