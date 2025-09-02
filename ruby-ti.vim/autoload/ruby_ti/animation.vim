function! ruby_ti#animation#start_typing(buffer_id, error_text, file_text, inner_width)
  if !ruby_ti#config#get('enable_animation', 1)
    call s:set_final_content(a:buffer_id, a:error_text, a:file_text, a:inner_width)
    return
  endif

  call ruby_ti#animation#stop()
  
  call ruby_ti#state#set_typing_state(0, -1)
  
  let speed = ruby_ti#config#get('animation_speed', 7)
  let timer_id = timer_start(speed, 
    \ function('s:typing_callback', [a:buffer_id, a:error_text, a:file_text, a:inner_width]), 
    \ {'repeat': -1})
  
  call ruby_ti#state#set_typing_state(0, timer_id)
endfunction

function! ruby_ti#animation#stop()
  let typing_state = ruby_ti#state#get_typing_state()
  if typing_state.timer_id != -1
    call timer_stop(typing_state.timer_id)
    call ruby_ti#state#reset_typing_state()
  endif
endfunction

function! s:typing_callback(buffer_id, error_text, file_text, inner_width, timer_id)
  if !s:is_buffer_valid(a:buffer_id)
    call ruby_ti#animation#stop()
    return
  endif
  
  let typing_state = ruby_ti#state#get_typing_state()
  let position = typing_state.position
  
  let max_length = max([len(a:error_text), len(a:file_text)])
  
  if position >= max_length
    call ruby_ti#animation#stop()
    return
  endif
  
  let error_typed = s:get_typed_text(a:error_text, position)
  let file_typed = s:get_typed_text(a:file_text, position)
  
  let config = ruby_ti#config#get('popup_style')
  let error_symbol = config.error_symbol
  let file_symbol = config.file_symbol
  
  let error_line = s:create_content_line(error_symbol, error_typed, a:inner_width)
  let file_line = s:create_content_line(file_symbol, file_typed, a:inner_width)
  
  try
    call nvim_buf_set_lines(a:buffer_id, 1, 2, v:true, [error_line])
    call nvim_buf_set_lines(a:buffer_id, 3, 4, v:true, [file_line])
  catch
    call ruby_ti#animation#stop()
    return
  endtry
  
  call ruby_ti#state#set_typing_state(position + 1, typing_state.timer_id)
endfunction

function! s:set_final_content(buffer_id, error_text, file_text, inner_width)
  if !s:is_buffer_valid(a:buffer_id)
    return
  endif
  
  let config = ruby_ti#config#get('popup_style')
  let error_symbol = config.error_symbol
  let file_symbol = config.file_symbol
  
  let error_line = s:create_content_line(error_symbol, a:error_text, a:inner_width)
  let file_line = s:create_content_line(file_symbol, a:file_text, a:inner_width)
  
  try
    call nvim_buf_set_lines(a:buffer_id, 1, 2, v:true, [error_line])
    call nvim_buf_set_lines(a:buffer_id, 3, 4, v:true, [file_line])
  catch
 
  endtry
endfunction

function! s:get_typed_text(text, position)
  return a:position < len(a:text) ? strpart(a:text, 0, a:position + 1) : a:text
endfunction

function! s:create_content_line(symbol, text, inner_width)
  let border_char = ruby_ti#config#get('popup_style').border_chars.vertical
  let symbol_text = a:symbol . ' '
  let padding_length = a:inner_width - len(symbol_text . a:text) - 0
  let padding = padding_length > 0 ? repeat(' ', padding_length) : ''
  return border_char . ' ' . symbol_text . a:text . padding . border_char
endfunction

function! s:is_buffer_valid(buffer_id)
  try
    return nvim_buf_is_valid(a:buffer_id)
  catch
    return 0
  endtry
endfunction
