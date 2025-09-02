def test
  a = {a: 1, b: '1'}
  b = {a: 1}

  a.merge! b do |k, o, n|
    dbtp n
  end

  a
end

test



