module Hoge
  def test x:
    x.to_s
  end

  class << self
  end
end

class Fuga
  include Hoge
end


f = Fuga.new
dbtp f.test x: 1
