def test
  yield '1', 1
end

def test2
  test do |x, y|
    dbtp y
  end
end

test2
