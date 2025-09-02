def test x
  x.each do |k, v|
    dbtp v
  end
end

test({a: 1, b: '2'})
