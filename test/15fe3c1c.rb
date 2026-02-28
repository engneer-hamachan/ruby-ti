def test x
  a = 1

  x.times {|c| y = 1 if c == 1}

  a
end

dbtp test 1
