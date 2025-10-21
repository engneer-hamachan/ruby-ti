class Hoge
  class Piya
    def test x
      dbtp x
    end

    class << self
      def test
        1
      end
    end
  end
end

class Fuga < Hoge::Piya
end

piyo = Fuga.new
piyo.test 1
piyo.test '1'
