class Hoge
  attr_accessor :aaa

  def test= x
    '1'
    dbtp x
  end

  def [] x
    dbtp x
  end

  def []= x
    dbtp x
  end
end
 
h = Hoge.new
h[0]
p h.test = 1

p h.aaa = '1'

dbtp h.aaa
