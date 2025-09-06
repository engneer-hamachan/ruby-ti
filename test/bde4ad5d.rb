x = Array.new

def hoge a
  def a.test
    p 1
  end

  a.test
end

dbtp a.test
