class Hoge
  def self.str
    '1'
  end

  def self.int
    1
  end
end

def test
  true ? Hoge.str : Hoge.int
end

dbtp test
