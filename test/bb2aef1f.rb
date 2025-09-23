class Hoge
end

class Fuga
end

array = []

array << 1
array << Fuga.new

dbtp array

array.each do |x|
  dbtp x
end
