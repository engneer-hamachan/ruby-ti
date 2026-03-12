a = []
a << {b: 1, y: '2'}
a << {b: 1, y: '2', z: true}


x = a[1]
p x[:b]
p x[:y]
dbtp x[:z]


# a.each do |x|
#   p x[:b]
#   p x[:y]
#   p x[:z]
# end
