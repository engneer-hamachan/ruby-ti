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
