def test x
  nil
end

lines = ['','']
x = 1 > 0 ? lines.join("\n") + "\n" : test 1 + 1

dbtp x


