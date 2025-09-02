def test x, y, *z 
  x + 1
  y + 1
  z[0] + 1
end

test 1, 2, '3', 4, 5
