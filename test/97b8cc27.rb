module Hoge
  def test
    @piyo
  end

  def test3 x
    @piyo = x
  end
end

class Fuga
  include Hoge
end

f = Fuga.new

f.test3 1000000

f.test + '1'
