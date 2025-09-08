pr = Proc.new {|x| dbtp x }

dbtp pr.call 1, 2, 3
