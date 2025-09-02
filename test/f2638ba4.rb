def test x, *y, z 
  x + 1
  y[0] + 1
  z + 1
end

test 1, 2, 3, 4, 5
