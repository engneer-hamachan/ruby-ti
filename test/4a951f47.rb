class Hoge
  def test(x)
    1
  end
end

class Fuga
  def test(x)
    '1'
  end
end

class Piyo
  def test(x, y)
    []
  end
end


h = Hoge.new
test(h)

piyo = Piyo.new
test(piyo) 

f = Fuga.new
test(f)


def test(x)
  x.test(1)
end
