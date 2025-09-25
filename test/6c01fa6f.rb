def test
  array = [1, '2', nil]

  array.each do |v|
    unless !v.is_a?(String)
      dbtp v
    else
      dbtp v
    end
  end
end
