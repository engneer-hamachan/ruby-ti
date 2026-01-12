tmp = 
  if true
    1
  else
    '1'
  end

case tmp
in String
  p tmp
in Integer
  p tmp
else
  p tmp
end

p tmp
