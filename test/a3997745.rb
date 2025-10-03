
module Piyo
  class Hoge
  end

  class Hoge::Fuga
    def test
      1
    end
  end
end

f = Piyo::Hoge::Fuga.new

dbtp f.test
