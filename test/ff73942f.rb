class Hoge
  def test x
    '1'
  end
end

h = Hoge.new

a = {}
a[:b] = h

dbtp a[:b].test 1
