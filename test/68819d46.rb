class Hoge
end

class Fuga
end

array = []

array << Hoge.new
array << Fuga.new

dbtp array

array.each do |x, y|
  dbtp x
  dbtp y
end
