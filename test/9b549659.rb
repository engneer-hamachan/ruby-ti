def test
  h = {}
  h[:a] = 1
  h[:b] = 2
  h[:b] = 2
  h[:b] = '2'


  h
end

a = test

dbtp a[:b]
