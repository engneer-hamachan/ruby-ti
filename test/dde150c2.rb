def test
  x = '1'.downcase! # => Union<String | Nil>

  if !x.is_a?(Nil)
    x + 1
  end
end
