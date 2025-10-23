class Hoge
  class Piya
    def test2(x:, y:)
    end

    class << self
      def test(x:, y:)
        1 + x
      end
    end
  end
end

class Fuga < Hoge::Piya
end

piyo = Fuga.new
piyo.test2 x: nil, y: nil
Fuga.test x: '1', y: 1


def hoge(x)
  Test.test(x, y)
end

hoge '1'

