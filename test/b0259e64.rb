class Hoge
  class Fuga
    def x
      1
    end
  end
end

class Fuga
  def x
    '1'
  end
end

test = Test.hogefuga

dbtp test.x
