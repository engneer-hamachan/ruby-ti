tmp = 
  if true
    1
  else
    '1'
  end

case tmp
in String
  dbtp tmp
in Integer
  dbtp tmp
else
  dbtp tmp
end

dbtp tmp
