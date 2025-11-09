def test cnt, &block
  cnt.times { |x| dbtp x }
end

test 3 do |x|
  dbtp x
  '111'
end


def test cnt, &block
  cnt.times do |x|
    dbtp x
    block.call 
  end
end

test 3 do |x|
  dbtp x
  '111'
end

