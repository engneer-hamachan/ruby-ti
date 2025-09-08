pr =
  Proc.new do |x|
    dbtp x
    x + 1
  end

dbtp pr.call 1, 2, 3
