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

Hoge::Piya.test x: 1

piya = Hoge::Piya.new
piya.test x: '1'
