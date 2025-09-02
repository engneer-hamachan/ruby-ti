def test
  x = '1'.downcase!
  b = []

  case 1
  when nil
    x
  when 1
    x
  when b
    x
  else
    x
  end
end

dbtp test
