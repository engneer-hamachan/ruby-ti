class Hoge
  def test2
    test
  end

  def test4
    test
  end

  protected

  def self.test
    1
  end
end

class Fuga < Hoge
end

p Fuga.test

# obj = Object.new
# def obj.singleton_method
#   "singleton"
# end
