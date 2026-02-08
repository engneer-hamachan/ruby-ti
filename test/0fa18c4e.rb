def test
  a = 
    if true
      1
    else
      nil
    end

  if a == 1
    dbtp a
  else
    dbtp a
  end
end

