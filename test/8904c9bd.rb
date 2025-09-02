def test
  yield 1
end

test do |x, y|
  dbtp x
end
