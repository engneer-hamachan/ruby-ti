def test
  a = {}

  a[:x] =
    if true
      1
    else
      '1'
    end

  a
end


y = {}

y[:a] = 1
y[:b] = '1'

dbtp y[:a]
dbtp y[:b]
dbtp y[:c]

y['a'] = 1
y['b'] = '1'

dbtp y['a']
dbtp y['b']
dbtp y['c']


