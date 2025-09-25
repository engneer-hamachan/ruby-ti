def test
  array = [1, '2']

  array.each do |v|
    if v.is_a?(String)
      dbtp v
    else
      dbtp v
    end
  end
end
