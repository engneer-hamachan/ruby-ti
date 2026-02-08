def test
  a = 
    if true
      1
    else
      nil
    end

  b = '1'

  if a.is_a?(String) || a.is_a?(Integer) || a.is_a?(NilClass)
    p a
    p b
  else
    dbtp a
  end
end

