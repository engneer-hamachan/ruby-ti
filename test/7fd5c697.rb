def test x, y
  y = '1'
  Test.test x, y
end

def test(q, z, a:, b: nil, c:)
  a + b
  p c
  z + c
end

test(1, '1', a: 1, c: 1)
