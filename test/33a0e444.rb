class Hoge
  class Piya
    def test x:
      dbtp x
      '1' + x
      x = 1
      dbtp x
    end

    class << self
      def test x:
        dbtp x
        1 + x
        x = '1'
        dbtp x
      end
    end
  end
end

h = Hoge::Piya.new
h.test x: '1'

Hoge::Piya.test x: 1
