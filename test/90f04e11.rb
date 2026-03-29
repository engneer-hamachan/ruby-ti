
dbtp Dir.open('aaa')
dbtp Dir.open('aaa') do
end


# ti-doc: bbb
def test x
  1
end

dbtp Dir.glob 'aaa', 1, base: '1'
dbtp Dir.glob 'aaa', base: '11111'
dbtp Dir.glob 'aaa', 1
dbtp Dir.glob 'aaa'

dbtp Dir.glob 'aaa', base: 1
dbtp Dir.glob 'aaa', '1'
