class Hoge
  CONST = '1'

  class Fuga
    def test
      '1'
    end
  end
end

Hoge::CONST + '1'


f = Hoge::Fuga.new

dbtp f.test
