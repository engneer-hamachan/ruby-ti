def test
  x = '1'.downcase!

  case x
  when nil
    dbtp x
  else
    x
  end
end

