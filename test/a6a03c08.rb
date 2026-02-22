def test x, y, z
  dbtp z
end

a = true
A = 1
B = 2

test(2, 2, a ? A : B)

