def with_block
  yield if block_given?
end

a = with_block do
    1
  end

p a

def with_block
  yield 2 if block_given?
end

b = Proc.new { p 3 }

a = with_block do |x|
    x
  end

p a

def with_block &block
  block.call
end

b = Proc.new { p 3 }

a = with_block { p 3 }


p a
