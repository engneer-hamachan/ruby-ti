KEYS = {}
KEYS['a'] = 1

def test x
  hash = {}

  hash['x'] = 'a'

  if true
    key = 'x'

    return KEYS[key]
  end

  ''
end

dbtp test 1
