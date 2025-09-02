class Hoge
  def test x
    x
  end
end

class Fuga < Hoge
end


hoge = Hoge.new
hoge.test 1
hoge.test '1'

fuga = Fuga.new
fuga.test []

dbtp hoge.test 1
