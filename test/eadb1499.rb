a = 
  if true
    1
  else
    'a'
  end

b = 
  if true
    1
  else
    'a'
  end

if a != 'a' && b.is_a?(Integer)
  dbtp a
  dbtp b
end

dbtp a
dbtp b

