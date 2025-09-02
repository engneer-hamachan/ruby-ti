def test
  a = '1'.downcase!

  a&.chars&.uniq

  a&.chars.uniq
end

