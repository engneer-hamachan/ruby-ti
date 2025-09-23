function! ruby_ti#checker#run()
  call ruby_ti#ui#highlight_error_line(-1)
  call ruby_ti#ui#clear_error_signs()
  call ruby_ti#state#clear_error_info()
  
  let file_path = expand('%')
  if empty(file_path)
    return
  endif
  
  if !filereadable(file_path)
    call ruby_ti#ui#echo_warning('File is not readable: ' . file_path)
    return
  endif
  
  let command = ruby_ti#config#get('checker_command', 'ti')
  
  if !executable(command)
    call ruby_ti#ui#echo_warning('Type checker command not found: ' . command . '. Please install Ruby-TI or set g:ruby_ti_config.checker_command')
    return
  endif
  
  let job_options = {
    \ 'on_stdout': function('s:on_checker_stdout'),
    \ 'on_stderr': function('s:on_checker_stderr'),
    \ 'on_exit': function('s:on_checker_exit'),
    \ 'stdout_buffered': v:true,
    \ 'stderr_buffered': v:true,
    \ 'cwd': getcwd()
  \ }
  
  try
    let job_id = jobstart([command, file_path, '-i'], job_options)
    if job_id <= 0
      call ruby_ti#ui#echo_warning('Failed to start type checker: ' . command)
    endif
  catch
    call ruby_ti#ui#echo_warning('Error starting type checker: ' . v:exception)
  endtry
endfunction

function! s:on_checker_stdout(job_id, data, event)
  call s:on_checker_complete(a:job_id, a:data, 'stdout')
endfunction

function! s:on_checker_stderr(job_id, data, event)
  call s:on_checker_complete(a:job_id, a:data, 'stderr')
endfunction

function! s:on_checker_complete(job_id, data, event)
  let output = join(a:data, "\n")
  let current_file = expand('%@:p')

  if empty(output) || output ==# "\n"
    return
  endif

  let all_errors = s:parse_all_checker_errors(output)
  if empty(all_errors)
    return
  endif

  " Check if error info display is enabled
  if !ruby_ti#config#get('enable_error_info', 1)
    " If error info is disabled, only show virtual text (type info) but skip error processing
    call ruby_ti#ui#show_virtual_text()
    return
  endif

  call ruby_ti#state#set_all_errors(all_errors)
  call ruby_ti#state#set_error_info(all_errors[0])

  let error_info = all_errors[0]

  let status_msg = error_info.message . ' (' . error_info.filename . ')'
  if len(all_errors) > 1
    let status_msg .= ' [1/' . len(all_errors) . ' errors]'
  endif
  call ruby_ti#ui#show_status(status_msg)

  let current_file_errors = []
  for error in all_errors
    if current_file ==# error.file_path
      call add(current_file_errors, error.line_number)
    endif
  endfor

  if !empty(current_file_errors)
    call ruby_ti#ui#highlight_error_lines(current_file_errors)

    call ruby_ti#ui#show_popup_if_needed()
  endif

  " Display type info as virtual text for current file
  call ruby_ti#ui#show_virtual_text()
endfunction

function! s:on_checker_exit(job_id, exit_code, event)
  " If no error was found (or no valid error was parsed), clear highlights and status
  let all_errors = ruby_ti#state#get_all_errors()
  if empty(all_errors)
    call ruby_ti#ui#highlight_error_lines([])
    call ruby_ti#ui#hide_popup()
    call ruby_ti#ui#clear_status()
  endif
  
  " Always show virtual text (even if no errors, there might be type info)
  call ruby_ti#ui#show_virtual_text()
endfunction

function! s:parse_all_checker_errors(output)
  " Parse multiple lines of format: "file_path::line_number::error_message" or "@file_path::line_number::type_info"
  let lines = split(a:output, '\n')
  let errors = []
  let type_infos = []
  
  for line in lines
    let line = s:sanitize_string(line)
    if empty(line)
      continue
    endif
    
    " Check if this is a type info message (starts with @)
    let is_type_info = line[0] ==# '@'
    if is_type_info
      let line = line[1:]  " Remove @ prefix
    endif
    
    let parts = split(line, '::')
    if len(parts) < 3
      continue
    endif
    
    let file_path = s:sanitize_string(parts[0])
    " For type info messages, ensure file_path is absolute
    if is_type_info && !empty(file_path) && file_path[0] !=# '/'
      let file_path = fnamemodify(file_path, ':p')
    endif
    let line_number = s:parse_line_number(parts[1])
    let message = s:sanitize_string(parts[2])
    
    " Validate parsed data
    if empty(file_path) || line_number <= 0 || empty(message)
      continue
    endif
    
    if is_type_info
      " Store type info for virtual text display
      call add(type_infos, {
        \ 'file_path': file_path,
        \ 'line_number': line_number,
        \ 'type_info': message
      \ })
    else
      " Create filename display (basename + line number)
      let filename_display = fnamemodify(file_path, ':t') . ' line:' . line_number
      
      call add(errors, {
        \ 'file_path': file_path,
        \ 'line_number': line_number,
        \ 'message': message,
        \ 'filename': filename_display
      \ })
    endif
  endfor
  
  " Store type infos in state for virtual text display
  call ruby_ti#state#set_type_infos(type_infos)
  
  return errors
endfunction

function! s:parse_checker_output(output)
  " Legacy function for backward compatibility
  let all_errors = s:parse_all_checker_errors(a:output)
  if !empty(all_errors)
    return all_errors[0]
  else
    return {}
  endif
endfunction

function! s:sanitize_string(str)
  " Remove leading/trailing whitespace and newlines
  return substitute(substitute(a:str, '^\s*', '', ''), '\s*$', '', '')
endfunction

function! s:parse_line_number(str)
  let cleaned = s:sanitize_string(a:str)
  let number = str2nr(cleaned)
  return number > 0 ? number : -1
endfunction
