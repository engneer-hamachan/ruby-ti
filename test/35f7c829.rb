
module Piyo
  class Hoge
    def test
      1
    end
  end

  class Hoge::Fuga
  end
end

f = Piyo::Hoge::Fuga.new

p f.test
