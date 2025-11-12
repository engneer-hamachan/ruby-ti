def test(array)
  args = []

  array.collect do |x|
    args << x
  end

  dbtp args
end

test [1, 2, 3]
