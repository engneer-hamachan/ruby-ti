module Hoge
  class << self
    def test x
      x + 1
    end
  end
end

Hoge.test(1) + '1'
