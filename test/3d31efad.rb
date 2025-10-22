class Hoge
  class Piya
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

Fuga.test x: '1', y: 2
