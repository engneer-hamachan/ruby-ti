function! ruby_ti#ui#setup_highlights()
  if &term ==# 'tmux-256color'
    highlight RubyTiErrorFloat guibg=#000a1a guifg=#00ff88 ctermbg=0 ctermfg=108 cterm=bold gui=bold
    highlight RubyTiErrorFloatBorder guibg=#000a1a guifg=#ff0088 ctermbg=0 ctermfg=198 cterm=bold gui=bold
  else
    highlight RubyTiErrorFloat guibg=#000a1a guifg=#88cc88 ctermbg=0 ctermfg=Green cterm=bold gui=bold
    highlight RubyTiErrorFloatBorder guibg=#000a1a guifg=#cc8888 ctermbg=0 ctermfg=Red cterm=bold gui=bold
  endif

  highlight RubyTiWarning ctermfg=Gray guifg=LightGray cterm=italic gui=italic
  highlight RubyTiTypeInfo ctermfg=Gray guifg=LightGray cterm=italic gui=italic
  
  if exists('&signcolumn')
    setlocal signcolumn=yes:2
  endif
  let mark = ruby_ti#config#get('mark', '💎')
  execute 'sign define RubyTiError text=' . mark . ' texthl=RubyTiWarning'
endfunction

function! ruby_ti#ui#echo_warning(message)
  if empty(a:message)
    return
  endif
  
  echohl RubyTiWarning
  echo 'Error: ' . a:message
  echohl None
endfunction

function! ruby_ti#ui#highlight_error_line(line_number)
  if !ruby_ti#config#get('enable_line_highlighting', 1)
    return
  endif
  
  if a:line_number > 0
    execute 'match RubyTiMatch /\%' . a:line_number . 'l/'
  else
    execute 'match none'
  endif
endfunction

function! ruby_ti#ui#highlight_error_lines(line_numbers)
  if !ruby_ti#config#get('enable_line_highlighting', 1)
    return
  endif
  
  call ruby_ti#ui#clear_error_signs()
  
  if empty(a:line_numbers)
    execute 'match none'
    return
  endif
  
  let pattern_parts = []
  for line_num in a:line_numbers
    if line_num > 0
      call add(pattern_parts, '\%' . line_num . 'l')
      call ruby_ti#ui#place_error_sign(line_num)
    endif
  endfor
  
  if !empty(pattern_parts)
    let pattern = '/\(' . join(pattern_parts, '\|') . '\)/'
    execute 'match RubyTiMatch ' . pattern
  else
    execute 'match none'
  endif
endfunction

function! ruby_ti#ui#place_error_sign(line_number)
  if a:line_number > 0
    if exists('&signcolumn')
      setlocal signcolumn=yes:2
    endif
    execute 'sign place ' . a:line_number . ' line=' . a:line_number . ' name=RubyTiError priority=0 buffer=' . bufnr('%')
  endif
endfunction

function! ruby_ti#ui#clear_error_signs()
  execute 'sign unplace * buffer=' . bufnr('%')
endfunction

function! ruby_ti#ui#show_popup_if_needed()
  " Check if error info display is enabled
  if !ruby_ti#config#get('enable_error_info', 1)
    return
  endif

  let current_line = line('.')
  let current_file = expand('%@:p')
  let all_errors = ruby_ti#state#get_all_errors()

  let current_line_error = {}
  let error_index = -1
  for i in range(len(all_errors))
    if all_errors[i].line_number == current_line && all_errors[i].file_path == current_file
      let current_line_error = all_errors[i]
      let error_index = i
      break
    endif
  endfor

  if ruby_ti#state#is_popup_visible() && empty(current_line_error)
    call ruby_ti#ui#hide_popup()
    return
  endif

  if !empty(current_line_error)
    let current_error_line = ruby_ti#state#get_error_info('line_number')
    let current_error_file = ruby_ti#state#get_error_info('file_path')

    if current_line_error.line_number != current_error_line || current_line_error.file_path != current_error_file
      if ruby_ti#state#is_popup_visible()
        call ruby_ti#ui#hide_popup()
      endif

      call ruby_ti#state#set_error_info(current_line_error)
      if error_index >= 0
        call ruby_ti#state#set_current_error_index(error_index)
      endif
      call ruby_ti#ui#show_popup()
    elseif !ruby_ti#state#is_popup_visible()
      call ruby_ti#state#set_error_info(current_line_error)
      if error_index >= 0
        call ruby_ti#state#set_current_error_index(error_index)
      endif
      call ruby_ti#ui#show_popup()
    endif
  endif
endfunction

function! ruby_ti#ui#show_popup()
  let error_message = ruby_ti#state#get_error_info('message')
  let error_filename = ruby_ti#state#get_error_info('filename')
  let error_line = ruby_ti#state#get_error_info('line_number')
  let error_file = ruby_ti#state#get_error_info('file_path')
  
  if empty(error_message) || empty(error_filename) || error_line <= 0 || empty(error_file)
    return
  endif
  
  let all_errors = ruby_ti#state#get_all_errors()
  let current_index = ruby_ti#state#get_current_error_index()
  if len(all_errors) > 1
    let error_filename .= ' [' . (current_index + 1) . '/' . len(all_errors) . ']'
  endif
  
  call ruby_ti#animation#stop()
  
  try
    let buffer_id = nvim_create_buf(v:false, v:true)
    if buffer_id == -1
      call ruby_ti#ui#echo_warning('Failed to create popup buffer')
      return
    endif
  catch
    call ruby_ti#ui#echo_warning('nvim_create_buf failed: ' . v:exception)
    return
  endtry
  
  try
    let dimensions = s:calculate_popup_dimensions(error_message, error_filename)
  catch
    call ruby_ti#ui#echo_warning('Failed to calculate dimensions: ' . v:exception)
    return
  endtry
  
  try
    let frame_content = s:create_popup_frame(dimensions.width, dimensions.inner_width)
  catch
    call ruby_ti#ui#echo_warning('Failed to create frame: ' . v:exception)
    return
  endtry
  
  try
    call nvim_buf_set_lines(buffer_id, 0, -1, v:true, frame_content)
  catch
    call ruby_ti#ui#echo_warning('Failed to set popup content: ' . v:exception)
    return
  endtry
  
  let config = ruby_ti#config#get('popup_style')
  let popup_options = {
    \ 'relative': 'cursor',
    \ 'width': dimensions.width,
    \ 'height': 5,
    \ 'col': ruby_ti#config#get('popup_offset_col', -2),
    \ 'row': ruby_ti#config#get('popup_offset_row', 1),
    \ 'anchor': 'NW',
    \ 'style': 'minimal',
    \ 'border': 'none'
  \ }
  
  try
    let window_id = nvim_open_win(buffer_id, 0, popup_options)
    call nvim_win_set_option(window_id, 'winhl', 'Normal:RubyTiErrorFloat,FloatBorder:RubyTiErrorFloatBorder')
    call ruby_ti#state#set_popup_window(window_id, 1)
  catch
    call ruby_ti#state#set_popup_window(-1, 0)
    return
  endtry
  
  let clean_error = substitute(error_message, '^\s*', '', '')
  let clean_filename = substitute(error_filename, '^\s*', '', '')
  call ruby_ti#animation#start_typing(buffer_id, clean_error, clean_filename, dimensions.inner_width)
endfunction

function! ruby_ti#ui#hide_popup()
  call ruby_ti#animation#stop()
  
  let window_id = ruby_ti#state#get_popup_window_id()
  if window_id != -1
    try
      call nvim_win_close(window_id, v:true)
    catch
    endtry
  endif
  
  call ruby_ti#state#set_popup_window(-1, 0)
endfunction

function! s:calculate_popup_dimensions(error_text, file_text)
  let config = ruby_ti#config#get('popup_style')
  let min_width = ruby_ti#config#get('min_popup_width', 40)
  
  let error_width = len(config.error_symbol . ' ' . a:error_text) + 4
  let file_width = len(config.file_symbol . ' ' . a:file_text) + 4
  let title_width = len(config.title)
  let footer_width = len(config.footer)
  
  let content_width = max([error_width, file_width, title_width, footer_width])
  let popup_width = max([content_width, min_width])
  let inner_width = popup_width - 1
  
  return {
    \ 'width': popup_width,
    \ 'inner_width': inner_width
  \ }
endfunction

function! s:create_popup_frame(popup_width, inner_width)
  let config = ruby_ti#config#get('popup_style')
  let chars = config.border_chars
  
  if a:inner_width <= 0
    throw 'Invalid inner_width: ' . a:inner_width
  endif
  
  let title_padding_length = a:inner_width - len(config.title) - 1
  let title_padding = title_padding_length > 0 ? repeat(chars.horizontal, title_padding_length) : ''
  
  let header = chars.top_left . repeat(chars.horizontal, 6) . config.title . repeat(chars.horizontal, a:inner_width - 28) . chars.top_right
  
  let footer_content = chars.footer_left . ' ' . config.footer . ' ' . chars.footer_right
  let footer_padding_length = a:inner_width - 15 - len(config.footer)
  let footer_padding = footer_padding_length > 0 ? repeat(chars.horizontal, footer_padding_length) : ''
  let footer_line = chars.bottom_left . footer_padding . footer_content . repeat(chars.horizontal, 10) . chars.bottom_right
  
  let separator = chars.separator_left . repeat(chars.horizontal, a:inner_width - 1) . chars.separator_right
  
  let empty_error = chars.vertical . ' ' . config.error_symbol . ' ' . repeat(' ', a:inner_width - len(config.error_symbol . ' ')) . chars.vertical
  let empty_file = chars.vertical . ' ' . config.file_symbol . ' ' . repeat(' ', a:inner_width - len(config.file_symbol . ' ')) . chars.vertical
  
  return [header, empty_error, separator, empty_file, footer_line]
endfunction

function! ruby_ti#ui#show_status(message)
  if empty(a:message)
    return
  endif
  
  echohl RubyTiWarning
  echo 'Ruby-TI Error: ' . a:message
  echohl None
endfunction

function! ruby_ti#ui#clear_status()
  echo ''
endfunction

function! ruby_ti#ui#show_virtual_text()
  if !ruby_ti#config#get('enable_def_type_info', 1)
    return
  endif
  
  let current_file = expand('%:p')
  let type_infos = ruby_ti#state#get_type_infos()
  let ns = ruby_ti#state#get_virtual_text_ns()
  
  " Clear existing virtual text
  call nvim_buf_clear_namespace(bufnr('%'), ns, 0, -1)
  
  " Add virtual text for type info messages in current file
  for type_info in type_infos
    if type_info.file_path == current_file
      let line_idx = type_info.line_number - 1  " Convert to 0-based indexing
      if line_idx >= 0
        call nvim_buf_set_virtual_text(bufnr('%'), ns, line_idx, 
          \ [[' → ' . type_info.type_info, 'RubyTiTypeInfo']], {})
      endif
    endif
  endfor
endfunction

function! ruby_ti#ui#clear_virtual_text()
  let ns = ruby_ti#state#get_virtual_text_ns()
  call nvim_buf_clear_namespace(bufnr('%'), ns, 0, -1)
endfunction
