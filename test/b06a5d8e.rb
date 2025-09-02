a = 1, 2

a.each do |x|
  x + 1
end

a = 1, '2'

a.each do |x|
  if x.is_a?(Integer)
    x + 1
  end
end

a.each do |x|
  x + 1
end

