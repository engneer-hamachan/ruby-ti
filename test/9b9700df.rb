module Hoge
  CONST = '1'

  class Fuga
    def test
      '1'
    end
  end
end

module Piyo
  CONST = '1'

  class Fuga
    def test
      1
    end
  end
end

Hoge::CONST + '1'



f = Piyo::Fuga.new

f.test + 1
f.test2 + 1
