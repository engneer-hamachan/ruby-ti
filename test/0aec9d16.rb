def test
  a = 1, 2, 3

  b = 
    a.collect! do |x|
      x.to_s
    end

  a
end

dbtp test
