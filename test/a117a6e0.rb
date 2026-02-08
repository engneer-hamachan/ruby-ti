def test
  a = 
    if true
      1
    else
      nil
    end

  b = '1'

  unless a == 2 && a.is_a?(NilClass)
    dbtp a
    p b
  else
    dbtp a
  end
end

