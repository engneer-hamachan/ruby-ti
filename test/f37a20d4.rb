class Hoge
  def test x
    x
  end
end

class Fuga < Hoge
end

class Piyo < Fuga
end

piyo1 = Piyo.new
piyo1.test '1'

piyo2 = Piyo.new
dbtp piyo2.test(1)
