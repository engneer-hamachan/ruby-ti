a = '1'.downcase!
b = true ? 1 : true

z =
  if true
    [1, b]
  else
    [{}, a, []]
  end

x, y, u = z

dbtp x
dbtp y
dbtp u
