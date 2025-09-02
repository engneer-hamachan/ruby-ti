a = '1', '2', '3'

b = 
  a.max do |x, y, z|
    x + '1'
    y + '1'
  end

dbtp b
