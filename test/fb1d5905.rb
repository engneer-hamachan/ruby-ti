class Hoge
  def test
    test2
  end

  def test2
    1
  end
end

def test3 x
end

h = Hoge.new

test3 h.test + '1'
