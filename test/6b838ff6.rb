def test
  a, *b, c = 1, '2', '3', 4

  d = [*b, 2]

  dbtp d
end

test

