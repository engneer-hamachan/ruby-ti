def test(*x, y, z, a: nil)
  dbtp x
  dbtp a
end

test 1, 2, '1', 1, a: 1
