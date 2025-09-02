def test x
  return x if x > 10

  test x + 1
end

dbtp test 1
