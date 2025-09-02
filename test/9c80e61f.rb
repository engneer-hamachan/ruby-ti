module Hoge
  class Piya
    def test x
      x
    end
  end
end

class Piya
  def test x
    x
  end
end

hp = Hoge::Piya.new
op = Piya.new

hp.test 1 + op.test '1'
