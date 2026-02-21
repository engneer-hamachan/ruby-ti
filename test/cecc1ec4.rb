a = {}
z = {}

a[:b] = 1
a[:b] = '1'

arr = []

arr << a

z[:b] = 2
arr << z

y = {}

y[:b] = []
arr << y

arr.each do |v, idx|
  dbtp v[:b]
#  v[:c] + '1'
end
