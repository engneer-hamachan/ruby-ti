class Hoge
  def test(x)
    "1" + x
  end

  def gethash
    {abc: 1}
  end
end


obj = Hoge.new
obj.test(1)
