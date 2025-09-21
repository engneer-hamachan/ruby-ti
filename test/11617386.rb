def with_block
  yield if block_given?
end

dbtp with_block
