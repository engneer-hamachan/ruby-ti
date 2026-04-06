def test(x, y, *z, a: nil)
  dbtp z
  dbtp a
end

test 1, 2, '1', 1, a: 1
