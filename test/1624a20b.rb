def test
  a = nil

  x =
    if a.is_a?(NilClass)
      1
    else
      a
    end
  x
end

dbtp test
