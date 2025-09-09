def method_with_splat(*args)
  args.size
end

def method_with_block(&block)
  block.call if block
end

method_with_block do |x|
  dbtp x
end
