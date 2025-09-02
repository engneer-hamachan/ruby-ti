a = '1', '2', '3'

b = 
  a.max(1) { |x, y, z|
    x + '1'
    y + '1'
  }

dbtp b
