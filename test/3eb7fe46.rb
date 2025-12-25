

class Hoge
  def self.test x
    []
  end

  def test
    '1'
  end
end

class Piyo < Hoge
  def piyo
    test
  end

  def self.piyo
    test 1
  end
end

f = Piyo.new
dbtp f.piyo
dbtp Piyo.piyo



