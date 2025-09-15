class Hoge
  def test
  end
end

def hoge
  Hoge.new
end

class Fuga
  class Piyo
    def test
      '1'
    end
  end

  class Uu
    def inner
      Piyo.new
    end
  end
end

f = Fuga::Uu.new
h = f.inner
# h = hoge
dbtp h.test
