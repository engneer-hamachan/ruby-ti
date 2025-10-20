class Hoge
  class Piya
    def test x
      1
    end

    class << self
      def test2
        1
      end
    end
  end
end

class Fuga < Hoge::Piya
end

class Piyo < Fuga
end

piyo = Piyo.new
dbtp piyo
