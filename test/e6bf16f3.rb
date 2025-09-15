class Hoge
  def test
  end
end

def hoge
  Hoge.new
end

class Piyo
  def test
    '1'
  end
end

class Fuga
  class Uu
    def inner
      Piyo.new
    end
  end
end

f = Fuga::Uu.new
h = f.inner
dbtp h.test
