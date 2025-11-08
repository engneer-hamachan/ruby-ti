def test cnt, &block
  cnt.times { block.call }
end

test 3 do |x|
  dbtp x
  '111'
end


def test cnt, &block
  cnt.times do
    block.call 
  end
end

test 3 do |x|
  dbtp x
  '111'
end

