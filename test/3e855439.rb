module Hoge
  module Fuga
    module Piyo
      class Piya
        def test
          1
        end
      end
    end

    class Piya
      def test
        '1'
      end
    end
  end
end

piya = Hoge::Fuga::Piyo::Piya.new

y = Hoge::Fuga::Piya.new

piya.test + y.test
