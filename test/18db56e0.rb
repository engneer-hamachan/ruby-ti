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
  def test(x)
    1
  end
end

def test(x)
  x.test(1)
end


h = Hoge.new
test(h)

piyo = Piyo.new
test(piyo) 

f = Fuga.new
test(f) + 1
