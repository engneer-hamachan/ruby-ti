module Hoge
  class Piya
    CONST = '1'
      
    def test
      '1'
    end
  end
end

module Hoge
  class Piyo
    CONST = '1'
      
    def test
      1
    end
  end
end

class Piya
  CONST = '1'
    
  def test
    '1'
  end
end

# p = Piya::CONST
# y = Hoge::Piya::CONST
# 
# dbtp y

pp = Hoge::Piya.new
ppp = Hoge::Piyo.new

pp.test + ppp.test
