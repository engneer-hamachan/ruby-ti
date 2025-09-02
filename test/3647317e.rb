def test
  result = {}

  result[:a] = 
    if true
      1
    else
      '1'
    end

  result
end

dbtp test
