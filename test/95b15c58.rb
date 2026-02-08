def test
  a = 
    if true
      1
    else
      nil
    end

  b = '1'

  if a == 1 || a.is_a?(Integer) || a.is_a?(NilClass)
    p a
    p b
  else
    dbtp a
  end
end

