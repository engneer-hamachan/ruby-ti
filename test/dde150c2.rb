def test
  x = '1'.downcase! # => Union<String | NilClass>

  if !x.is_a?(NilClass)
    x + 1
  end
end
