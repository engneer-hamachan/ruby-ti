class Hoge
  class Piya
    def test(x:, y:)
      dbtp x
      dbtp y
    end

    class << self
      def test(x:, y:)
        dbtp x
        dbtp y
      end
    end
  end
end

class Fuga < Hoge::Piya
end

piyo = Fuga.new
piyo.test x: nil, y: nil

Fuga.test x: '1', y: 2
Fuga.test x: 1, y: '2'
