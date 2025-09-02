class Hoge
  def test
    yield 1, 2, '3'
  end
end

h = Hoge.new


h.test do |x, y, z, q|
  dbtp z
end
