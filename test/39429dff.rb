def test
  x = '1'.downcase!

  case x
  when nil
    x
  else
    dbtp x
  end
end

