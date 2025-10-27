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

class Fuga < Hoge::Piya
end

Hoge::Piya.test x: 1
Fuga.test x: '1'

piya = Hoge::Piya.new
piya.test x: '1'

f = Fuga.new
f.test x: 1
