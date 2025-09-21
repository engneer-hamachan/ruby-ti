def hoge(&block)
  block.call(100) if true
end

dbtp hoge { |e| e + 1 } # => 101
