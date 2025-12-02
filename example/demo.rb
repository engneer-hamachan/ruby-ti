u = 
  if true
    1
  elsif true
    nil
  else
    1, 2, 3
  end

u + 1

if u.is_a?(Integer)
  u + 1
end

p u 


module Hoge
  CONST = '1'

  class Fuga
    def test x 
      {a: x, b: '1'}
    end
  end
end

Hoge::CONST + '1'
Hoge::CONST + 1


f = Hoge::Fuga.new
# f = Fuga.new

hash = f.test

p f.test(1)[:a]
