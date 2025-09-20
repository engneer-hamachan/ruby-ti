let s:state = {}

function! ruby_ti#state#init()
  let s:state = {
    \ 'popup_visible': 0,
    \ 'popup_window_id': -1,
    \ 'typing_position': 0,
    \ 'typing_timer_id': -1,
    \ 'error_info': {
    \   'message': '',
    \   'filename': '',
    \   'line_number': -1,
    \   'file_path': ''
    \ },
    \ 'all_errors': [],
    \ 'current_error_index': 0,
    \ 'match_id': -1,
    \ 'type_infos': [],
    \ 'virtual_text_ns': -1
  \ }
endfunction

function! ruby_ti#state#get(key, ...)
  let default = a:0 > 0 ? a:1 : v:null
  return get(s:state, a:key, default)
endfunction

function! ruby_ti#state#set(key, value)
  let s:state[a:key] = a:value
endfunction

function! ruby_ti#state#get_error_info(key, ...)
  let default = a:0 > 0 ? a:1 : ''
  return get(s:state['error_info'], a:key, default)
endfunction

function! ruby_ti#state#set_error_info(info_dict)
  call extend(s:state['error_info'], a:info_dict, 'force')
endfunction

function! ruby_ti#state#clear_error_info()
  let s:state['error_info'] = {
    \ 'message': '',
    \ 'filename': '',
    \ 'line_number': -1,
    \ 'file_path': ''
  \ }
  let s:state['all_errors'] = []
  let s:state['current_error_index'] = 0
  let s:state['type_infos'] = []
endfunction

function! ruby_ti#state#is_popup_visible()
  return get(s:state, 'popup_visible', 0)
endfunction

function! ruby_ti#state#get_popup_window_id()
  return get(s:state, 'popup_window_id', -1)
endfunction

function! ruby_ti#state#set_popup_window(window_id, visible)
  let s:state['popup_window_id'] = a:window_id
  let s:state['popup_visible'] = a:visible
endfunction

function! ruby_ti#state#get_typing_state()
  return {
    \ 'position': get(s:state, 'typing_position', 0),
    \ 'timer_id': get(s:state, 'typing_timer_id', -1)
  \ }
endfunction

function! ruby_ti#state#set_typing_state(position, timer_id)
  let s:state['typing_position'] = a:position
  let s:state['typing_timer_id'] = a:timer_id
endfunction

function! ruby_ti#state#reset_typing_state()
  let s:state['typing_position'] = 0
  let s:state['typing_timer_id'] = -1
endfunction

function! ruby_ti#state#set_all_errors(errors)
  let s:state['all_errors'] = a:errors
  let s:state['current_error_index'] = 0
endfunction

function! ruby_ti#state#get_all_errors()
  return get(s:state, 'all_errors', [])
endfunction

function! ruby_ti#state#get_current_error_index()
  return get(s:state, 'current_error_index', 0)
endfunction

function! ruby_ti#state#set_current_error_index(index)
  let s:state['current_error_index'] = a:index
endfunction

function! ruby_ti#state#cycle_to_next_error()
  let all_errors = ruby_ti#state#get_all_errors()
  if empty(all_errors)
    return
  endif
  
  let current_index = ruby_ti#state#get_current_error_index()
  let next_index = (current_index + 1) % len(all_errors)
  call ruby_ti#state#set_current_error_index(next_index)
  call ruby_ti#state#set_error_info(all_errors[next_index])
endfunction

function! ruby_ti#state#set_type_infos(type_infos)
  let s:state['type_infos'] = a:type_infos
endfunction

function! ruby_ti#state#get_type_infos()
  return get(s:state, 'type_infos', [])
endfunction

function! ruby_ti#state#get_virtual_text_ns()
  if s:state['virtual_text_ns'] == -1
    let s:state['virtual_text_ns'] = nvim_create_namespace('ruby_ti_virtual_text')
  endif
  return s:state['virtual_text_ns']
endfunction

function! ruby_ti#state#reset()
  execute 'match none'
  
  if get(s:state, 'popup_visible', 0) && get(s:state, 'popup_window_id', -1) != -1
    try
      call nvim_win_close(get(s:state, 'popup_window_id', -1), v:true)
    catch
    endtry
  endif
  
  if get(s:state, 'typing_timer_id', -1) != -1
    call timer_stop(get(s:state, 'typing_timer_id', -1))
  endif
  
  let s:state['popup_visible'] = 0
  let s:state['popup_window_id'] = -1
  call ruby_ti#state#reset_typing_state()
  call ruby_ti#state#clear_error_info()
  let s:state['type_infos'] = []
  
  " Clear virtual text
  if s:state['virtual_text_ns'] != -1
    call nvim_buf_clear_namespace(bufnr('%'), s:state['virtual_text_ns'], 0, -1)
  endif
endfunction
