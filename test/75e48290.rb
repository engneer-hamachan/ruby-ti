def test(x, **kwargs, &block)
  dbtp x
  dbtp kwargs
end

test 1, 2
