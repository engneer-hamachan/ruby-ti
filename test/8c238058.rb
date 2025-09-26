obj = Object.new

def obj.singleton_method
  "singleton"
end

obj.sleep 1

class Hoge
  def test
  end
end

h = Hoge.new
h.test

