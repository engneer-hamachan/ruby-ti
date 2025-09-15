class Hoge
  def test
  end
end

def hoge
  Hoge.new
end

class Fuga
  class Uu
    class Piyo
      def test
        '1'
      end
    end

    def inner
      Piyo.new
    end
  end
end

f = Fuga::Uu.new
h = f.inner
# h = hoge
dbtp h.test
