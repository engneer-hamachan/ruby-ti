def test(x: nil, y:)
  dbtp x
end

test(x: 1, y: 1)

def test(z, a:, b:, c:)
  a + b
  p c
  z + c
end

test(1, a: 1, c: '2', b: 2)
