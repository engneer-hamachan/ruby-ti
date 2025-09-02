class Hoge
  def test(x, y)
    y + x
  end

  def gethash
    {abc: 1}
  end
end


obj = Hoge.new
# obj.gethash + 1

a = []
b = {}


str = []
p obj.test(1, 2)
p obj.test(1, str)
