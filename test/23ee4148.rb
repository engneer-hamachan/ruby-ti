class Custom
  def test x
    1
  end

  def self.test
    'a'
  end
end

class Hoge < Custom
  def fuga
    test 1
  end

  def self.piyo
    p test
  end
end

h = Hoge.new
dbtp h.fuga
dbtp Hoge.piyo
