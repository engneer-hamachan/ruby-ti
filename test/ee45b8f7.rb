def test
  x = '1'.downcase!

  unless x.is_a?(Nil)
    x + '1'
  end

  unless x.is_a?(String)
    x + '1'
  end

  unless x.is_a?(Nil)
    x + '1'
  end
end
