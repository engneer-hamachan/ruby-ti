def test
  a = 1, 2, 3

  a.collect do |x|
    if true
      x.to_s
    else
      x
    end
  end
end

dbtp test
