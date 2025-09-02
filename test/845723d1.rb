module Hoge
  module Fuga
    def test x
      @hoge = 1
      p x.abs
    end
  end
end


class Piyo
  include Hoge::Fuga
end

piyo = Piyo.new

dbtp piyo.hoge

piyo.test 1
