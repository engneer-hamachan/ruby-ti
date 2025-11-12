def test(**kwargs)
  args = []

  kwargs.collect do |x|
    args << x
  end

  dbtp args
end

test a: 1, b: '2'
