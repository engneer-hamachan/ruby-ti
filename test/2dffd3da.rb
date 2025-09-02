class Hoge
  def test
    a = {a: '1', b: 1}
    a[:b] = '1'

    a
  end
end

h = Hoge.new

h.test[:b] = 11111

dbtp h.test[:b]
