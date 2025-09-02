def test
  a = nil
  a&.abs

  a = '1'
  a&.abs
end
