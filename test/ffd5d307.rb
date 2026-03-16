a = [
  false || 0,
]

dbtp a

a = false || 1, 2, 3

dbtp a

def test
  return 1, 2, 3
end

dbtp test

dbtp [true || 1, 2, 3]

def test2
  return [1, 2, 3], 1, 2
end

dbtp test2
