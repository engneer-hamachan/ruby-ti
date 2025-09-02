class Hoge
  def test x, *y, z
    x[0] + 1
  end
end

h = Hoge.new

h.test 1, 2, 3, 4
