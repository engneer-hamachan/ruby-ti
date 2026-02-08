def test
  a = 
    if true
      1
    else
      nil
    end

  if a.nil?
    1
  elsif a < 0
    dbtp a
  end
end
