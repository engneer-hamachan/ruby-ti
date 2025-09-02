class Hoge
  def test
    if 1 == 1
      return 1
    else
      '1'
    end
  end
end

h = Hoge.new

h.test() + 1
