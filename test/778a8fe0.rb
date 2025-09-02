def test
  a = 1, 2, 3

  a.collect do |x|
    if true
      x.to_s
    else
      x
    end

    11 
  end
end

dbtp test
