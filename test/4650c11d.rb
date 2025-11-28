class Hoge
  def self.test2 x
    []
  end

  def test
    '1'
  end
end

class Fuga < Hoge
  def piyo
    self.test
  end

  def self.piyo
    self.test2 1 
  end
end

f = Fuga.new
a = []

dbtp f.piyo
dbtp Fuga.piyo
