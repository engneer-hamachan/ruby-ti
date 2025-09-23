test_cases = [
  ['1', 2, []],
  ['1', 2, 3, 4],
]

test_cases.each do |a, b, c, d, e|
  dbtp a
  dbtp b
  dbtp c
  dbtp d
  dbtp e
end

