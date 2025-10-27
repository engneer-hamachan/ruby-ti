class Hoge
  class Piya
    class << self
      def test x
        dbtp x
      end
    end
  end
end


Hoge::Piya.test 1

