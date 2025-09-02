def test
  x = '1'.downcase!

  case x
  when nil
    x
  else
    x + '1'
  end

  dbtp x
end

