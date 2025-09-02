class Hoge
  def test
    a = []
    a << 1
    

    a
  end
end

h = Hoge.new

h.test[0] = '11111'

a = h.test[0]

a[0] = '1'

dbtp a[0]
