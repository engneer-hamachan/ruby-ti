# bind type
u = 
  if true
    1
  elsif true
    nil
  else
    1, 2, 3
  end

# no method error
u + 1

# narrowing
if u.is_a?(Integer)
  u + 1
end

# print type
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

# type missmatch error
Hoge::CONST + 1


# method call
f = Hoge::Fuga.new

# class not defined error
f = Fuga.new

# argument error
hash = f.test

# method call after ref
p f.test(1)[:a]
