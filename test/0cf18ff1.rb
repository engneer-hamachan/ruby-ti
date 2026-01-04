def get_completions(input)
  return [] if input == ''

  # get last token from input (simple word extraction)
  last_word = ''
  i = input.length - 1
 
  for num in [1, 2, 3]
    p num
    c = input[i]
    if true 
      last_word = c + last_word
      i = 1
    end
  end

  []
end
 
def draw_completions(completions)
  dbtp completions
end
   
completions2 = get_completions ''

draw_completions completions2

