x = Array.new

def hoge a
  def a.test
    p 1
  end

  dbtp a.test
end
