class Hoge
  def test2
    test
  end

  def test4 x
    x.test
  end

  protected

  def test
    1
  end
end

class Fuga < Hoge
  def test3 x
    x.test
  end
end

class Piyo
  def test3 x
    x.test
  end
end

h = Hoge.new
f = Fuga.new
piyo = Piyo.new

piyo.test3 f

p h.test4 h
p f.test3 h
p f.test3 f

p f.test

# obj = Object.new
# def obj.singleton_method
#   "singleton"
# end
