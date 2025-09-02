module Hoge
  attr_reader :piyo

  def test
    @piyo = 1
  end
end

class Fuga
  include Hoge
end

f = Fuga.new

f.test
dbtp f.piyo
