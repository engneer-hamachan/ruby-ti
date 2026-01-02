def test
  x = '1'.downcase!

  if !x.is_a?(NilClass)
    x + '1'
  end
end
