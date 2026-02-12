def array_remove_at(arr, index)
  dbtp index
  result = []
  arr.each_with_index do |item, i|
    if i != index
      result << item
    end
  end
  result
end

array_remove_at [], 1
