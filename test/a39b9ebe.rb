class Hoge
  def test2
    test
  end

  def test4
    test
  end

  protected

  def test
    1
  end
end

class Fuga < Hoge
  def test3
    test
  end
end

class Piyo
  def test3
    test
  end
end

h = Hoge.new
f = Fuga.new
piyo = Piyo.new


p h.test4
p f.test3
p f.test3

p f.test

# obj = Object.new
# def obj.singleton_method
#   "singleton"
# end
