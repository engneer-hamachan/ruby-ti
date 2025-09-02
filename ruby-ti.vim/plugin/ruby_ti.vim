if exists('g:loaded_ruby_ti')
  finish
endif
let g:loaded_ruby_ti = 1

function! s:initialize_ruby_ti()
  try
    call ruby_ti#state#init()

    if exists('g:ruby_ti_config')
      call ruby_ti#config#update(g:ruby_ti_config)
    endif
    
    call ruby_ti#ui#setup_highlights()
    
    call s:setup_autocommands()
  catch
    echo 'Ruby-TI Error: Failed to initialize - ' . v:exception
  endtry
endfunction

autocmd VimEnter * call s:initialize_ruby_ti()

function! s:setup_autocommands()
  augroup RubyTi
    autocmd!
    autocmd BufRead *.* call ruby_ti#state#reset()
    autocmd BufWinEnter *.* call ruby_ti#state#reset()
    autocmd BufWritePost *.rb if ruby_ti#config#get('auto_run', 0) | call ruby_ti#checker#run() | endif
    autocmd BufReadPost *.rb if ruby_ti#config#get('auto_run', 0) | call ruby_ti#checker#run() | endif
    autocmd BufWinEnter *.rb if ruby_ti#config#get('auto_run', 0) | call ruby_ti#checker#run() | endif
    autocmd CursorMoved *.rb call ruby_ti#ui#show_popup_if_needed()
  augroup END
endfunction

command! RubyTiRun call ruby_ti#checker#run()
