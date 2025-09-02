def test
  h = {}
  h[:a] = 1
  h[:b] = 2
  h[:b] = 2
  h[:b] = '2'

  h
end

a = test

a.each do |a, b|
  dbtp b
  p a
  p b
end
