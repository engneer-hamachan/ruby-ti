class Hoge
  def test x
    x
  end
end

class Fuga < Hoge
end

class Piyo < Fuga
end

hoge = Hoge.new
hoge.test '1'

fuga = Fuga.new
fuga.test 1

piyo = Piyo.new
dbtp piyo.test([])
