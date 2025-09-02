class Hoge
  def test(x)
    if 1 == 1
      return 1
    end

    '1'
  end
end

h = Hoge.new

h.test(1) + 1
