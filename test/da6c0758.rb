def test
  x = '1'.downcase!

  if !x.is_a?(Nil)
    x + '1'
  end
end
