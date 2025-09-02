class Hoge
  def test x
    x
  end
end

class Fuga < Hoge
end

f = Fuga.new

'1' + f.test 1
