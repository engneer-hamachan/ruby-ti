def test
  a = {a: 1, b: 2}
  b = {a: '1', b: '2'}

  a.merge(b) do |k, o, n| 
    n
  end
end

dbtp test 
