def test
  a = 1, 2, 3

  a.collect do |x|
    x.to_s
  end
end

dbtp test
