class Hoge
  attr_reader :fuga

  def test
    @fuga = 1
  end
end

h = Hoge.new

h.fuga + "1"
