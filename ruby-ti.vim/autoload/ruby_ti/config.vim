let s:config = {
  \ 'animation_speed': 7,
  \ 'min_popup_width': 40,
  \ 'popup_offset_col': -2,
  \ 'popup_offset_row': 1,
  \ 'checker_command': 'ti',
  \ 'enable_animation': 1,
  \ 'enable_line_highlighting': 1,
  \ 'enable_type_display': 0,
  \ 'enable_error_info': 1,
  \ 'enable_error_virtual_text': 0,
  \ 'auto_run': 0,
  \ 'mark': '💎',
  \ 'popup_style': {
  \   'title': ' MRUBY TYPE ANALYSIS ',
  \   'footer': 'alerted by ruby-ti',
  \   'error_symbol': '▸',
  \   'file_symbol': '📁',
  \   'border_chars': {
  \     'top_left': '╔',
  \     'top_right': '╗',
  \     'bottom_left': '╚',
  \     'bottom_right': '╝',
  \     'horizontal': '═',
  \     'vertical': '║',
  \     'separator_left': '╠',
  \     'separator_right': '╣',
  \     'footer_left': '▷',
  \     'footer_right': '◁'
  \   }
  \ },
\ }

function! ruby_ti#config#get(key, ...)
  let default = a:0 > 0 ? a:1 : v:null
  return get(s:config, a:key, default)
endfunction

function! ruby_ti#config#set(key, value)
  let s:config[a:key] = a:value
endfunction

function! ruby_ti#config#update(config_dict)
  call extend(s:config, a:config_dict, 'force')
endfunction
