
dbtp Test.test 1, '1'

y = 
  Test.test(1, '1') do |x|
    dbtp x
  end

dbtp y

