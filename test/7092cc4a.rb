a = {}

a[:b] = 1
a[:c] = '2'

arr = []

arr << a

arr.each do |v, idx|
  v[:b] + 1
  v[:c] + '1'
end
