def test(x, **kwargs, &block)
  dbtp x
  dbtp kwargs
end

test 1, a: 1, b: 2
