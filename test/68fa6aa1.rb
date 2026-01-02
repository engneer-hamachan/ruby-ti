def test
  x = '1'.downcase!
  if x.is_a?(String)
    x + '1'
  elsif x.is_a?(Integer)
    x + 1
  elsif x.is_a?(NilClass)
    x + 1
  end
end
