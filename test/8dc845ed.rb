def test
  yield 1
end

def test2
  yield '1', 2
end

def test3
  test2 do |x, y|
    dbtp x
  end
end
