def test(x, y, *z, a: nil)
  dbtp x
  dbtp y
  dbtp z
  dbtp a
end

test 1, '1', a: 1
