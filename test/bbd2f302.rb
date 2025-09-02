class Hoge
  def test x
    x
  end
end

class Fuga < Hoge
end

class Piyo < Fuga
end

piyo = Piyo.new

dbtp piyo.test(1)
