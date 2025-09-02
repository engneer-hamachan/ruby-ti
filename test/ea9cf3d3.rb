module Hoge
  class Fuga
    def test
      a = Array.new
      a.<< 1

      a
    end
  end
end

f = Hoge::Fuga.new

dbtp f.test[0]
