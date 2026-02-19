class Hoge
  attr_accessor :fuga
end

class Fuga
  attr_accessor :name
end

h = Hoge.new
f = Fuga.new

h.fuga = f
h.fuga.name = '1'

a = 1 == h.fuga.name ? 1 : 1

dbtp a
