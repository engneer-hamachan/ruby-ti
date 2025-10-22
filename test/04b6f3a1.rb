class Hoge
  class Piya
    def test x:
      dbtp x
    end

    class << self
      def test x:
        dbtp x
      end
    end
  end
end

def test2 x:
end

class Fuga < Hoge::Piya
end

piyo = Fuga.new
piyo.test x: 1
Fuga.test x: '1'

a = Array.new
a << 1
x = a.count

