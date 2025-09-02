x = '1'.downcase!

y =
  case x
  when nil
    x
  else
    x
  end

dbtp y
