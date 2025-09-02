x = '1'.downcase!
b = []

y =
  case x
  when nil
    x
  when 1
    x
  when b
    x
  else
    x
  end

dbtp y
