# Method Signatures
## array_remove_at(Array<untyped>, Integer) -> Array<untyped>
- file: theme/editor_app.rb:407
- document: remove item from array at index
- call points:
  - theme/editor_app.rb:350
    - method: handle_backspace
- total call points: 1

## calculate_indent_decrease(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:231
- document: returns indent_ct-1 if the first matching token is 'end'/'else'/'elsif'/'when', otherwise returns indent_ct unchanged
- call points:
  - theme/editor_app.rb:471
    - method: handle_return_key
- total call points: 1

## calculate_indent_increase(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:247
- document: returns indent_ct+1 if first token is a block-opening keyword (class/module/def/if/unless/elsif/else/do/case/when/while/until/for) or any token is 'do', otherwise returns indent_ct unchanged
- call points:
  - theme/editor_app.rb:457
    - method: handle_return_key
  - theme/editor_app.rb:474
    - method: handle_return_key
- total call points: 2

## calculate_line_number_space(untyped) -> Integer
- file: theme/editor_app.rb:222
- document: returns 1 space for line numbers >9, 2 spaces for single-digit numbers — right-aligns line numbers in the gutter
- call points:
  - theme/editor_app.rb:1232
    - method: redraw_code_area
  - theme/editor_app.rb:1260
    - method: redraw_code_area
- total call points: 2

## calculate_max_history_lines(Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1164
- document: returns max_visible_lines-1 when cursor is on new bottom line (cursor_row_index==-1) or temp_new_line_code is non-empty (reserves a row for the current input line); otherwise returns max_visible_lines
- call points:
  - theme/editor_app.rb:1172
    - method: calculate_scroll_offset
  - theme/editor_app.rb:1214
    - method: redraw_code_area
- total call points: 2

## calculate_scroll_offset(Array<untyped>, Integer, Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1170
- document: when editing existing line, adjusts scroll_offset so cursor_row_index is within the visible window; when on new bottom line, scrolls to show the last max_history_lines; returns new scroll_offset
- call points:
  - theme/editor_app.rb:1216
    - method: redraw_code_area
- total call points: 1

## draw_code_with_highlight(M5GFX, String, Integer, Integer) -> Array<String>
- file: theme/editor_app.rb:971
- document: tokenizes code_str and draws each token at (x_pos, y_pos) with per-token syntax color; tracks 'def' context to colorize the following method name as COLOR_METHOD; x_pos advances by token.length*6 per token
- call points:
  - theme/editor_app.rb:1239
    - method: redraw_code_area
  - theme/editor_app.rb:1248
    - method: redraw_code_area
  - theme/editor_app.rb:1265
    - method: redraw_code_area
- total call points: 3

## draw_completion(M5GFX, String) -> NilClass
- file: theme/editor_app.rb:1088
- document: finds completion candidates for the last token of current_code (max 5), draws a bordered popup box with the selected candidate highlighted, sets $completion_chars to the suffix to append on Tab; clears $completion_chars if no match
- call points:
  - theme/editor_app.rb:1255
    - method: redraw_code_area
- total call points: 1

## draw_static_ui(M5GFX) -> NilClass
- file: theme/editor_app.rb:1026
- document: draws fixed header border with filename '/home/geek/picoruby/calc.rb', the '=>' result label, and footer separator lines — called once at startup, not redrawn in the main loop
- call points:
  - theme/editor_app.rb:1308
    - method: top level
- total call points: 1

## execute_and_get_result(Sandbox, untyped) -> ExecutionResult
- file: theme/editor_app.rb:828
- document: wraps execute_code in '_ = (...)' and compiles/runs it in sandbox; returns ExecutionResult with res=result or error message, display_res=res.to_s, error=true on compile or runtime failure
- call points:
  - theme/editor_app.rb:504
    - method: handle_return_key
  - theme/editor_app.rb:855
    - method: execute_code_and_update_state
- total call points: 2

## execute_code_and_update_state(EditorState, Sandbox, M5GFX, Integer) -> Integer
- file: theme/editor_app.rb:854
- document: execute state.execute_code via sandbox, store result in state.res/display_res, update completion dict on success, then reset state; always resets on both success and error
- call points:
  - theme/editor_app.rb:1389
    - method: top level
- total call points: 1

## find_completion_candidates(String) -> Array<untyped>
- file: theme/editor_app.rb:1045
- document: returns all $dict keys that start with target_code and are longer than it; returns empty array if target_code is empty
- call points:
  - theme/editor_app.rb:1105
    - method: draw_completion
- total call points: 1

## get_input() -> Array<String Bool>
- file: hardware_adapters/adv_input.rb:124
- document: read keyboard input (returns [key_name, is_pressed])
- call points:
  - theme/editor_app.rb:1377
    - method: top level
- total call points: 1

## get_result_color(Class) -> Integer
- file: theme/editor_app.rb:899
- document: returns COLOR_NUMBER for Integer/Float, COLOR_STRING for String, COLOR_SYMBOL for NilClass/TrueClass/FalseClass, COLOR_WHITE for all other types
- call points:
  - theme/editor_app.rb:1358
    - method: top level
- total call points: 1

## get_token_color(String, Bool, Array<String>) -> Integer
- file: theme/editor_app.rb:953
- document: returns syntax highlight color for token — strings→COLOR_STRING, symbols/nil/true/false→COLOR_SYMBOL, @/@@ /$vars→COLOR_VARIABLE, numbers→COLOR_NUMBER, keywords→COLOR_KEYWORD, Constants→COLOR_CONSTANT, def-context name→COLOR_METHOD, else COLOR_WHITE
- call points:
  - theme/editor_app.rb:1018
    - method: draw_code_with_highlight
- total call points: 1

## handle_backspace(EditorState) -> Bool
- file: theme/editor_app.rb:345
- document: deletes char before cursor; if code is empty on an existing line, deletes that line and moves cursor to previous line; if on new empty bottom line, moves cursor to last confirmed line; returns true if cursor changed lines
- call points:
  - theme/editor_app.rb:1439
    - method: top level
- total call points: 1

## handle_completion_navigation(String) -> NavigationResult
- file: theme/editor_app.rb:706
- document: cycles $completion_selected_index up or down through $completion_candidates with wrap-around; returns NavigationResult with is_need_redraw_input=true only if candidates exist
- call points:
  - theme/editor_app.rb:806
    - method: handle_fn_navigation
  - theme/editor_app.rb:815
    - method: handle_fn_navigation
- total call points: 2

## handle_ctrl_input(EditorState, String) -> Array<Bool>
- file: theme/editor_app.rb:597
- document: ctrl+d: full editor reset and history index reset; ctrl+c: clear current code line; ctrl+;: step backwards through $history; ctrl+.: step forwards through $history (clears editor past-end); returns [handled, is_need_redraw_input]
- call points:
  - theme/editor_app.rb:1484
    - method: top level
- total call points: 1

## handle_cursor_move(EditorState, String) -> NavigationResult
- file: theme/editor_app.rb:677
- document: moves cursor_col_index left (decrement) or right (increment) by one char; sets cursor_col_index=-1 when moving to end of line; returns NavigationResult with is_need_redraw_input=true if cursor moved
- call points:
  - theme/editor_app.rb:800
    - method: handle_fn_navigation
- total call points: 1

## handle_fn_navigation(EditorState, String) -> NavigationResult
- file: theme/editor_app.rb:795
- document: dispatches left/right to cursor_move (when code non-empty) or horizontal_scroll; dispatches up/down to completion_navigation first, falling back to move_cursor_up/down if no completion candidates active
- call points:
  - theme/editor_app.rb:1465
    - method: top level
- total call points: 1

## handle_horizontal_scroll(EditorState, String) -> NavigationResult
- file: theme/editor_app.rb:652
- document: scrolls result_display_offset by 37 chars in the given direction; right-scroll is bounded by display_res length; left-scroll clamps to 0; returns NavigationResult with is_need_redraw_result=true
- call points:
  - theme/editor_app.rb:802
    - method: handle_fn_navigation
- total call points: 1

## handle_normal_char_input(EditorState, String) -> Union<NilClass String>
- file: theme/editor_app.rb:571
- document: inserts key_input at cursor_col_index (or end of line if -1), advances cursor_col_index, resets $completion_selected_index; syncs code_lines[:text] and execute_code if editing existing line
- call points:
  - theme/editor_app.rb:1496
    - method: top level
- total call points: 1

## handle_return_key(EditorState, Bool, Sandbox, M5GFX, Integer) -> ReturnKeyResult
- file: theme/editor_app.rb:444
- document: handle return key — normal return (is_shift=false) adds new line with auto-indent; shift+return executes full code via sandbox and resets state; returns ReturnKeyResult with should_continue=true on normal return
- call points:
  - theme/editor_app.rb:1396
    - method: top level
- total call points: 1

## handle_shift_input(EditorState, String) -> Bool
- file: theme/editor_app.rb:537
- document: looks up key_input in SHIFT_TABLE; if found, inserts the shifted char at cursor position and advances cursor_col_index; returns true if key was in SHIFT_TABLE, false otherwise
- call points:
  - theme/editor_app.rb:1451
    - method: top level
- total call points: 1

## handle_tab_key(EditorState) -> Union<NilClass Integer>
- file: theme/editor_app.rb:337
- document: if $completion_chars is a non-nil string, appends it to state.code and resets $completion_selected_index to 0; does nothing if no completion candidate is active
- call points:
  - theme/editor_app.rb:1380
    - method: top level
- total call points: 1

## hash_array_insert(untyped, Hash, untyped) -> Array<untyped Hash>
- file: theme/editor_app.rb:418
- document: inserts item into arr after the element at index position-1 (1-based); position=1 inserts after the first element; returns new array with item inserted
- call points:
  - theme/editor_app.rb:461
    - method: handle_return_key
- total call points: 1

## init_keyboard() -> Bool
- file: hardware_adapters/adv_input.rb:73
- document: initialize keyboard
- call points:
  - theme/editor_app.rb:1302
    - method: top level
- total call points: 1

## is_at_last_line?(EditorState) -> Bool
- file: theme/editor_app.rb:274
- document: returns true if cursor_row_index==-1 (on new bottom line) or cursor_row_index==code_lines.length-1 (on last confirmed line)
- call points:
  - theme/editor_app.rb:1385
    - method: top level
  - theme/editor_app.rb:1385
    - method: top level
- total call points: 2

## is_number?(String) -> Bool
- file: theme/editor_app.rb:888
- document: returns true only if str is non-empty and every character is an ASCII digit '0'-'9'; does not accept floats, negatives, or hex
- call points:
  - theme/editor_app.rb:960
    - method: get_token_color
  - theme/editor_app.rb:960
    - method: get_token_color
- total call points: 2

## load_constants() -> Array<Symbol>
- file: theme/editor_app.rb:208
- document: populates $dict with all Object constants, excluding INTERNAL_CONSTANTS and any names containing 'Error'
- call points:
  - theme/editor_app.rb:315
    - method: process_definition_tokens
  - theme/editor_app.rb:1324
    - method: top level
- total call points: 2

## move_cursor_down(EditorState) -> NavigationResult
- file: theme/editor_app.rb:763
- document: move cursor down to next line — at last confirmed line with non-empty code, moves to new line at bottom (cursor_row_index=-1); does nothing if code is empty at last line; returns NavigationResult
- call points:
  - theme/editor_app.rb:820
    - method: handle_fn_navigation
- total call points: 1

## move_cursor_up(EditorState) -> NavigationResult
- file: theme/editor_app.rb:730
- document: move cursor up to previous line — if on new line (cursor_row_index==-1) with non-empty code, saves it as a confirmed line first; sets cursor_col_index=-1 (end of line); returns NavigationResult
- call points:
  - theme/editor_app.rb:811
    - method: handle_fn_navigation
- total call points: 1

## process_definition_tokens(Array<String>, Array<String>, Array<String>, Array<String>) -> Array<String>
- file: theme/editor_app.rb:285
- document: scans tokens to add def/class/module names to $dict; adds 'new' for class definitions; adds attr_reader/attr_accessor-declared attribute names; calls load_constants when 'require' token is found
- call points:
  - theme/editor_app.rb:327
    - method: update_completion_dict
- total call points: 1

## rebuild_execute_code(untyped) -> String
- file: theme/editor_app.rb:433
- document: joins all code_lines[:text] entries with trailing newlines into a single string; used to reconstruct execute_code after lines are added/deleted/edited
- call points:
  - theme/editor_app.rb:370
    - method: handle_backspace
  - theme/editor_app.rb:399
    - method: handle_backspace
  - theme/editor_app.rb:463
    - method: handle_return_key
  - theme/editor_app.rb:473
    - method: handle_return_key
  - theme/editor_app.rb:561
    - method: handle_shift_input
  - theme/editor_app.rb:592
    - method: handle_normal_char_input
  - theme/editor_app.rb:620
    - method: handle_ctrl_input
  - theme/editor_app.rb:633
    - method: handle_ctrl_input
  - theme/editor_app.rb:737
    - method: move_cursor_up
- total call points: 9

## redraw_code_area(M5GFX, Array<untyped>, Integer, Integer, String, untyped, Integer, Integer, String, Integer) -> Integer
- file: theme/editor_app.rb:1194
- document: clears code area, recalculates scroll_offset, draws visible committed lines with syntax highlighting and cursor underline, draws completion overlay, draws current input line; returns new scroll_offset
- call points:
  - theme/editor_app.rb:484
    - method: handle_return_key
  - theme/editor_app.rb:518
    - method: handle_return_key
  - theme/editor_app.rb:872
    - method: execute_code_and_update_state
  - theme/editor_app.rb:1309
    - method: top level
  - theme/editor_app.rb:1332
    - method: top level
- total call points: 5

## EditorState.reset() -> Integer
- file: theme/editor_app.rb:106
- document: reset editor state to initial values
- call points:
  - theme/editor_app.rb:280
    - method: reset_editor_state
- total call points: 1

## reset_editor_state(EditorState) -> Integer
- file: theme/editor_app.rb:279
- document: reset editor state to initial values and reset completion selection index to 0
- call points:
  - theme/editor_app.rb:510
    - method: handle_return_key
  - theme/editor_app.rb:516
    - method: handle_return_key
  - theme/editor_app.rb:602
    - method: handle_ctrl_input
  - theme/editor_app.rb:865
    - method: execute_code_and_update_state
  - theme/editor_app.rb:870
    - method: execute_code_and_update_state
- total call points: 5

## tokenize(untyped) -> Array<String>
- file: theme/editor_app.rb:912
- document: splits code into tokens using space/punctuation as delimiters; quoted string literals (single or double) are kept as single tokens including the quotes; returns array of token strings
- call points:
  - theme/editor_app.rb:326
    - method: update_completion_dict
  - theme/editor_app.rb:450
    - method: handle_return_key
  - theme/editor_app.rb:1006
    - method: draw_code_with_highlight
  - theme/editor_app.rb:1092
    - method: draw_completion
  - theme/editor_app.rb:1279
    - method: redraw_code_area
- total call points: 5

## update_class_method_dict(Array<String>) -> Union<NilClass Array<Symbol>>
- file: theme/editor_app.rb:1062
- document: populates $class_method_dict with '.method_name' entries when tokens end with [ConstantName, '.'] or [ConstantName, '.', partial] pattern; clears $class_method_dict when pattern is not detected; filters out Object base methods to show only class-specific methods
- call points:
  - theme/editor_app.rb:1093
    - method: draw_completion
- total call points: 1

## update_completion_dict(untyped) -> Union<Bool NilClass>
- file: theme/editor_app.rb:320
- document: scans all code_lines via process_definition_tokens to populate $dict with identifiers; afterwards removes temporary entries 'attr_reader', 'attr_accessor', 'initialize' from the dict
- call points:
  - theme/editor_app.rb:513
    - method: handle_return_key
  - theme/editor_app.rb:867
    - method: execute_code_and_update_state
- total call points: 2

---
# Special Code Comments
## theme/editor_app.rb:72
- comment: holds all editor state variables in a single object to reduce Hash return values
```
class EditorState
```

## theme/editor_app.rb:80
- comment: code being typed on the current (bottom) new line — empty string means no code yet
```
    @code = ''
```

## theme/editor_app.rb:82
- comment: full code string to execute, built by joining code_lines — represents all committed lines
```
    @execute_code = ''
```

## theme/editor_app.rb:84
- comment: array of committed code lines (not including the current new line being typed)
```
    @code_lines = []
```

## theme/editor_app.rb:86
- comment: current indentation level (number of indent units), used for auto-indent on new lines
```
    @indent_ct = 0
```

## theme/editor_app.rb:88
- comment: 1-based row number displayed on screen
```
    @current_row_number = 1
```

## theme/editor_app.rb:90
- comment: scroll offset for the result display panel
```
    @result_display_offset = 0
```

## theme/editor_app.rb:92
- comment: -1 means cursor is on the new bottom line; >=0 means editing an existing committed line at that index
```
    @cursor_row_index = -1
```

## theme/editor_app.rb:94
- comment: saves new line code when entering edit mode for an existing line, restored on exit
```
    @temp_new_line_code = ''
```

## theme/editor_app.rb:96
- comment: scroll offset for the code lines area to keep cursor visible
```
    @scroll_offset = 0
```

## theme/editor_app.rb:98
- comment: raw result string from the last code execution
```
    @res = ''
```

## theme/editor_app.rb:100
- comment: formatted/truncated result string for display in the result panel
```
    @display_res = ''
```

## theme/editor_app.rb:102
- comment: column position within the current code line — -1 means end of line (default), 0..n means specific position
```
    @cursor_col_index = -1
```

## theme/editor_app.rb:119
- comment: holds fn navigation results for redraw flags
```
class NavigationResult
```

## theme/editor_app.rb:125
- comment: true if result panel needs to be redrawn this frame
```
    @is_need_redraw_result = false
```

## theme/editor_app.rb:127
- comment: true if input/code area needs to be redrawn this frame
```
    @is_need_redraw_input = false
```

## theme/editor_app.rb:132
- comment: holds code execution results including error status
```
class ExecutionResult
```

## theme/editor_app.rb:138
- comment: raw result string from execution (inspect of return value)
```
    @res = ''
```

## theme/editor_app.rb:140
- comment: formatted result string for display (may be truncated or annotated)
```
    @display_res = ''
```

## theme/editor_app.rb:142
- comment: true if execution raised an exception or runtime error
```
    @error = false
```

## theme/editor_app.rb:147
- comment: holds return key handling results for control flow
```
class ReturnKeyResult
```

## theme/editor_app.rb:153
- comment: true if the main loop should `next` (skip remaining processing for this frame)
```
    @should_continue = false
```

## theme/editor_app.rb:155
- comment: true if an error occurred during return key handling (used to trigger error redraw)
```
    @error_occurred = false
```

## theme/editor_app.rb:160
- comment: ADC object for reading battery voltage
```
bat_adc = ADC.new(10)
```

## theme/editor_app.rb:163
- comment: sandbox for executing user-entered mruby code
```
sandbox = Sandbox.new ''
```

## theme/editor_app.rb:166
- comment: flag to prevent duplicate key processing in same frame
```
is_input = false
```

## theme/editor_app.rb:168
- comment: true while shift key is held — causes characters to be looked up in SHIFT_TABLE
```
is_shift = false
```

## theme/editor_app.rb:170
- comment: true while fn key is held — causes keys to be looked up in FN_TABLE for navigation
```
is_fn = false
```

## theme/editor_app.rb:172
- comment: true while ctrl key is held — enables ctrl+c (cancel line) and ctrl+d (reset editor)
```
is_ctrl = false
```

## theme/editor_app.rb:174
- comment: set to true to force redraw of the code input area on the next loop iteration
```
is_need_redraw_input = false
```

## theme/editor_app.rb:176
- comment: set to true to force redraw of the result display area on the next loop iteration
```
is_need_redraw_result = false
```

## theme/editor_app.rb:178
- comment: previous code display string used to detect changes and trigger input area redraw
```
prev_code_display = ''
```

## theme/editor_app.rb:180
- comment: previous result string used to detect changes and trigger result area redraw
```
prev_res = ''
```

## theme/editor_app.rb:182
- comment: previous status bar string used to detect changes and trigger status bar redraw
```
prev_status = ''
```

## theme/editor_app.rb:184
- comment: maximum number of code lines visible in the code area at once (scroll window size)
```
max_visible_lines = 7
```

## theme/editor_app.rb:186
- comment: unused variable, kept for historical reasons
```
cursor_pos = 0
```

## theme/editor_app.rb:188
- comment: stores the currently selected completion candidate string
```
$completion_chars = nil
```

## theme/editor_app.rb:190
- comment: list of current completion candidate strings
```
$completion_candidates = []
```

## theme/editor_app.rb:192
- comment: index of currently highlighted completion candidate
```
$completion_selected_index = 0
```

## theme/editor_app.rb:194
- comment: dictionary of known identifiers for autocompletion
```
$dict = {}
```

## theme/editor_app.rb:196
- comment: temporary dictionary for class method candidates (populated when ClassName. pattern detected)
```
$class_method_dict = {}
```

## theme/editor_app.rb:198
- comment: tracks which class name is currently providing dot-completion methods (nil if not active)
```
$class_method_prefix = nil
```

## theme/editor_app.rb:200
- comment: history of executed code_lines arrays (each element is an array of line hashes)
```
$history = []
```

## theme/editor_app.rb:202
- comment: current position in history navigation (-1 means not navigating history)
```
$history_index = -1
```

## theme/editor_app.rb:205
- comment: main editor state instance holding all editor variables
```
editor_state = EditorState.new
```

## theme/editor_app.rb:1101
- comment: when pattern is [ConstantName, '.', partial], prepend '.' to target_code_raw so it matches class method candidates stored as '.method_name' in $class_method_dict
```
  if tokens.length >= 3 && tokens[-2] == '.' && tokens[-3].length > 0 && tokens[-3][0] >= 'A' && tokens[-3][0] <= 'Z' && target_code_raw != '.'
```

## theme/editor_app.rb:1278
- comment: side effect — if any code_line starts with 'class', inject attr_reader/attr_accessor/initialize into the completion dict so class body methods appear as candidates
```
  code_lines.each do |line|
```

## theme/editor_app.rb:1385
- comment: execute code when return is pressed on the last line with empty input
```
  if key_input == 'ret' && key_pressed && is_at_last_line?(editor_state) && editor_state.code == '' && editor_state.execute_code != '' && !is_input
```

## theme/editor_app.rb:1393
- comment: normal return — add new line (code != '') or execute with shift+return; only fires when there is content to act on
```
  elsif key_input == 'ret' && key_pressed && (editor_state.code != '' || editor_state.execute_code != '') && !is_input
```

## theme/editor_app.rb:1412
- comment: main key dispatch block — handles all keys except 'ret' and tab; modifier keys (shift/fn/ctrl) update their flags then skip; regular keys are dispatched in priority order: del → shift-modified → fn-mapped → ctrl-modified → normal char
```
  if key_input != '' && key_input != 'ret' && !is_input
```

## theme/editor_app.rb:1502
- comment: reset is_input flag when no key is being held — allows the next key event to be processed
```
  if key_input == ''
```

