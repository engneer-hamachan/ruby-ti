a =
  if true
    1
  else
    {}
  end

test_cases = [
  ['1', a, []],
  [1, '1', 1, []],
]

test_cases.each do |a, b, c, d, e|
  p a
  p b
  p c
  p d
  p e
end

