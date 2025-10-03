
module Piyo
  class Hoge
  end

  def test
    1
  end

  class Hoge::Fuga
  end
end

f = Piyo::Hoge::Fuga.new

p f.test
