# テストが通らなくなったら削除
class Hoge
  def test2
    1
  end
end


h = Hoge.new

h.test2 + "1"
