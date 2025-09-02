module Hoge
  module Fuga
    def test x
      p x.abs
    end
  end
end

module Fuga
  def test x
    x.abs
  end
end

class Piyo
  include Hoge::Fuga
end

piyo = Piyo.new
piyo.test 1
