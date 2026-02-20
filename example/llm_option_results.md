# Method Signatures
## array_remove_at(Array<untyped>, Integer) -> Array<untyped>
- file: theme/editor_app.rb:402
- document: remove item from array at index
- total call points: 1
- call points:
  - theme/editor_app.rb:345

## calculate_indent_decrease(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:226
- document: returns indent_ct-1 if the first matching token is 'end'/'else'/'elsif'/'when', otherwise returns indent_ct unchanged
- total call points: 2
- call points:
  - theme/editor_app.rb:449
  - theme/editor_app.rb:466

## calculate_indent_increase(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:242
- document: returns indent_ct+1 if first token is a block-opening keyword (class/module/def/if/unless/elsif/else/do/case/when/while/until/for) or any token is 'do', otherwise returns indent_ct unchanged
- total call points: 2
- call points:
  - theme/editor_app.rb:452
  - theme/editor_app.rb:469

## calculate_line_number_space(untyped) -> Integer
- file: theme/editor_app.rb:217
- document: returns 1 space for line numbers >9, 2 spaces for single-digit numbers — right-aligns line numbers in the gutter
- total call points: 2
- call points:
  - theme/editor_app.rb:1189
  - theme/editor_app.rb:1217

## calculate_max_history_lines(Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1121
- document: returns max_visible_lines-1 when cursor is on new bottom line (cursor_row_index==-1) or temp_new_line_code is non-empty (reserves a row for the current input line); otherwise returns max_visible_lines
- total call points: 2
- call points:
  - theme/editor_app.rb:1129
  - theme/editor_app.rb:1171

## calculate_scroll_offset(Array<untyped>, Integer, Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1127
- document: when editing existing line, adjusts scroll_offset so cursor_row_index is within the visible window; when on new bottom line, scrolls to show the last max_history_lines; returns new scroll_offset
- total call points: 1
- call points:
  - theme/editor_app.rb:1173

## draw_code_with_highlight(M5GFX, String, Integer, Integer) -> Array<String>
- file: theme/editor_app.rb:966
- document: tokenizes code_str and draws each token at (x_pos, y_pos) with per-token syntax color; tracks 'def' context to colorize the following method name as COLOR_METHOD; x_pos advances by token.length*6 per token
- total call points: 3
- call points:
  - theme/editor_app.rb:1196
  - theme/editor_app.rb:1205
  - theme/editor_app.rb:1222

## draw_completion(M5GFX, String) -> NilClass
- file: theme/editor_app.rb:1057
- document: finds completion candidates for the last token of current_code (max 5), draws a bordered popup box with the selected candidate highlighted, sets $completion_chars to the suffix to append on Tab; clears $completion_chars if no match
- total call points: 1
- call points:
  - theme/editor_app.rb:1212

## draw_static_ui(M5GFX) -> NilClass
- file: theme/editor_app.rb:1021
- document: draws fixed header border with filename '/home/geek/picoruby/calc.rb', the '=>' result label, and footer separator lines — called once at startup, not redrawn in the main loop
- total call points: 1
- call points:
  - theme/editor_app.rb:1265

## execute_and_get_result(Sandbox, untyped) -> ExecutionResult
- file: theme/editor_app.rb:823
- document: wraps execute_code in '_ = (...)' and compiles/runs it in sandbox; returns ExecutionResult with res=result or error message, display_res=res.to_s, error=true on compile or runtime failure
- total call points: 2
- call points:
  - theme/editor_app.rb:499
  - theme/editor_app.rb:850

## execute_code_and_update_state(EditorState, Sandbox, M5GFX, Integer) -> Integer
- file: theme/editor_app.rb:849
- document: execute state.execute_code via sandbox, store result in state.res/display_res, update completion dict on success, then reset state; always resets on both success and error
- total call points: 1
- call points:
  - theme/editor_app.rb:1346

## find_completion_candidates(String) -> Array<untyped>
- file: theme/editor_app.rb:1040
- document: returns all $dict keys that start with target_code and are longer than it; returns empty array if target_code is empty
- total call points: 1
- call points:
  - theme/editor_app.rb:1067

## get_input() -> Array<String Bool>
- file: hardware_adapters/adv_input.rb:124
- document: read keyboard input (returns [key_name, is_pressed])
- total call points: 1
- call points:
  - theme/editor_app.rb:1334

## get_result_color(Class) -> Integer
- file: theme/editor_app.rb:894
- document: returns COLOR_NUMBER for Integer/Float, COLOR_STRING for String, COLOR_SYMBOL for NilClass/TrueClass/FalseClass, COLOR_WHITE for all other types
- total call points: 1
- call points:
  - theme/editor_app.rb:1315

## get_token_color(String, Bool, Array<String>) -> Integer
- file: theme/editor_app.rb:948
- document: returns syntax highlight color for token — strings→COLOR_STRING, symbols/nil/true/false→COLOR_SYMBOL, @/@@ /$vars→COLOR_VARIABLE, numbers→COLOR_NUMBER, keywords→COLOR_KEYWORD, Constants→COLOR_CONSTANT, def-context name→COLOR_METHOD, else COLOR_WHITE
- total call points: 1
- call points:
  - theme/editor_app.rb:1013

## handle_backspace(EditorState) -> Bool
- file: theme/editor_app.rb:340
- document: deletes char before cursor; if code is empty on an existing line, deletes that line and moves cursor to previous line; if on new empty bottom line, moves cursor to last confirmed line; returns true if cursor changed lines
- total call points: 1
- call points:
  - theme/editor_app.rb:1395

## handle_completion_navigation(String) -> NavigationResult
- file: theme/editor_app.rb:701
- document: cycles $completion_selected_index up or down through $completion_candidates with wrap-around; returns NavigationResult with is_need_redraw_input=true only if candidates exist
- total call points: 2
- call points:
  - theme/editor_app.rb:801
  - theme/editor_app.rb:810

## handle_ctrl_input(EditorState, String) -> Array<Bool>
- file: theme/editor_app.rb:592
- document: ctrl+d: full editor reset and history index reset; ctrl+c: clear current code line; ctrl+;: step backwards through $history; ctrl+.: step forwards through $history (clears editor past-end); returns [handled, is_need_redraw_input]
- total call points: 1
- call points:
  - theme/editor_app.rb:1440

## handle_cursor_move(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:672
- document: moves cursor_col_index left (decrement) or right (increment) by one char; sets cursor_col_index=-1 when moving to end of line; returns NavigationResult with is_need_redraw_input=true if cursor moved
- total call points: 1
- call points:
  - theme/editor_app.rb:795

## handle_fn_navigation(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:790
- document: dispatches left/right to cursor_move (when code non-empty) or horizontal_scroll; dispatches up/down to completion_navigation first, falling back to move_cursor_up/down if no completion candidates active
- total call points: 1
- call points:
  - theme/editor_app.rb:1421

## handle_horizontal_scroll(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:647
- document: scrolls result_display_offset by 37 chars in the given direction; right-scroll is bounded by display_res length; left-scroll clamps to 0; returns NavigationResult with is_need_redraw_result=true
- total call points: 1
- call points:
  - theme/editor_app.rb:797

## handle_normal_char_input(EditorState, String) -> Union<NilClass String>
- file: theme/editor_app.rb:566
- document: inserts key_input at cursor_col_index (or end of line if -1), advances cursor_col_index, resets $completion_selected_index; syncs code_lines[:text] and execute_code if editing existing line
- total call points: 1
- call points:
  - theme/editor_app.rb:1452

## handle_return_key(EditorState, Bool, Sandbox, M5GFX, Integer) -> ReturnKeyResult
- file: theme/editor_app.rb:439
- document: handle return key — normal return (is_shift=false) adds new line with auto-indent; shift+return executes full code via sandbox and resets state; returns ReturnKeyResult with should_continue=true on normal return
- total call points: 1
- call points:
  - theme/editor_app.rb:1353

## handle_shift_input(EditorState, String) -> Bool
- file: theme/editor_app.rb:532
- document: looks up key_input in SHIFT_TABLE; if found, inserts the shifted char at cursor position and advances cursor_col_index; returns true if key was in SHIFT_TABLE, false otherwise
- total call points: 1
- call points:
  - theme/editor_app.rb:1407

## handle_tab_key(EditorState) -> Union<NilClass Integer>
- file: theme/editor_app.rb:332
- document: if $completion_chars is a non-nil string, appends it to state.code and resets $completion_selected_index to 0; does nothing if no completion candidate is active
- total call points: 1
- call points:
  - theme/editor_app.rb:1337

## hash_array_insert(untyped, Hash, untyped) -> Array<untyped Hash>
- file: theme/editor_app.rb:413
- document: inserts item into arr after the element at index position-1 (1-based); position=1 inserts after the first element; returns new array with item inserted
- total call points: 1
- call points:
  - theme/editor_app.rb:456

## init_keyboard() -> Bool
- file: hardware_adapters/adv_input.rb:73
- document: initialize keyboard
- total call points: 1
- call points:
  - theme/editor_app.rb:1259

## is_at_last_line?(EditorState) -> Bool
- file: theme/editor_app.rb:269
- document: returns true if cursor_row_index==-1 (on new bottom line) or cursor_row_index==code_lines.length-1 (on last confirmed line)
- total call points: 1
- call points:
  - theme/editor_app.rb:1342

## is_number?(String) -> Bool
- file: theme/editor_app.rb:883
- document: check if a string is a number
- total call points: 1
- call points:
  - theme/editor_app.rb:955

## load_constants() -> Array<Symbol>
- file: theme/editor_app.rb:203
- document: load Ruby constants into completion dictionary
- total call points: 2
- call points:
  - theme/editor_app.rb:310
  - theme/editor_app.rb:1281

## move_cursor_down(EditorState) -> NavigationResult
- file: theme/editor_app.rb:758
- document: move cursor down to next line — at last confirmed line with non-empty code, moves to new line at bottom (cursor_row_index=-1); does nothing if code is empty at last line; returns NavigationResult
- total call points: 1
- call points:
  - theme/editor_app.rb:815

## move_cursor_up(EditorState) -> NavigationResult
- file: theme/editor_app.rb:725
- document: move cursor up to previous line — if on new line (cursor_row_index==-1) with non-empty code, saves it as a confirmed line first; sets cursor_col_index=-1 (end of line); returns NavigationResult
- total call points: 1
- call points:
  - theme/editor_app.rb:806

## process_definition_tokens(Array<String>, Array<String>, Array<String>, Array<String>) -> Array<String>
- file: theme/editor_app.rb:280
- document: scans tokens to add def/class/module names to $dict; adds 'new' for class definitions; adds attr_reader/attr_accessor-declared attribute names; calls load_constants when 'require' token is found
- total call points: 1
- call points:
  - theme/editor_app.rb:322

## rebuild_execute_code(untyped) -> String
- file: theme/editor_app.rb:428
- document: joins all code_lines[:text] entries with trailing newlines into a single string; used to reconstruct execute_code after lines are added/deleted/edited
- total call points: 9
- call points:
  - theme/editor_app.rb:365
  - theme/editor_app.rb:394
  - theme/editor_app.rb:458
  - theme/editor_app.rb:468
  - theme/editor_app.rb:556
  - theme/editor_app.rb:587
  - theme/editor_app.rb:615
  - theme/editor_app.rb:628
  - theme/editor_app.rb:732

## redraw_code_area(M5GFX, Array<untyped>, Integer, Integer, String, untyped, Integer, Integer, String, Integer) -> Integer
- file: theme/editor_app.rb:1151
- document: clears code area, recalculates scroll_offset, draws visible committed lines with syntax highlighting and cursor underline, draws completion overlay, draws current input line; returns new scroll_offset
- total call points: 5
- call points:
  - theme/editor_app.rb:479
  - theme/editor_app.rb:513
  - theme/editor_app.rb:867
  - theme/editor_app.rb:1266
  - theme/editor_app.rb:1289

## EditorState.reset() -> Integer
- file: theme/editor_app.rb:106
- document: reset editor state to initial values
- total call points: 1
- call points:
  - theme/editor_app.rb:275

## reset_editor_state(EditorState) -> Integer
- file: theme/editor_app.rb:274
- document: reset editor state to initial values and reset completion selection index to 0
- total call points: 5
- call points:
  - theme/editor_app.rb:505
  - theme/editor_app.rb:511
  - theme/editor_app.rb:597
  - theme/editor_app.rb:860
  - theme/editor_app.rb:865

## tokenize(untyped) -> Array<String>
- file: theme/editor_app.rb:907
- document: splits code into tokens using space/punctuation as delimiters; quoted string literals (single or double) are kept as single tokens including the quotes; returns array of token strings
- total call points: 5
- call points:
  - theme/editor_app.rb:321
  - theme/editor_app.rb:445
  - theme/editor_app.rb:1001
  - theme/editor_app.rb:1063
  - theme/editor_app.rb:1236

## update_completion_dict(untyped) -> Union<Bool NilClass>
- file: theme/editor_app.rb:315
- document: scans all code_lines via process_definition_tokens to populate $dict with identifiers; afterwards removes temporary entries 'attr_reader', 'attr_accessor', 'initialize' from the dict
- total call points: 2
- call points:
  - theme/editor_app.rb:508
  - theme/editor_app.rb:862

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

## theme/editor_app.rb:187
- comment: stores the currently selected completion candidate string
```
$completion_chars = nil
```

## theme/editor_app.rb:189
- comment: list of current completion candidate strings
```
$completion_candidates = []
```

## theme/editor_app.rb:191
- comment: index of currently highlighted completion candidate
```
$completion_selected_index = 0
```

## theme/editor_app.rb:193
- comment: dictionary of known identifiers for autocompletion
```
$dict = {}
```

## theme/editor_app.rb:195
- comment: history of executed code_lines arrays (each element is an array of line hashes)
```
$history = []
```

## theme/editor_app.rb:197
- comment: current position in history navigation (-1 means not navigating history)
```
$history_index = -1
```

## theme/editor_app.rb:200
- comment: main editor state instance holding all editor variables
```
editor_state = EditorState.new
```

## theme/editor_app.rb:1235
- comment: side effect — if any code_line starts with 'class', inject attr_reader/attr_accessor/initialize into the completion dict so class body methods appear as candidates
```
  code_lines.each do |line|
```

## theme/editor_app.rb:1342
- comment: execute code when return is pressed on the last line with empty input
```
  if key_input == 'ret' && key_pressed && is_at_last_line?(editor_state) && editor_state.code == '' && editor_state.execute_code != '' && !is_input
```

## theme/editor_app.rb:1350
- comment: normal return — add new line (code != '') or execute with shift+return; only fires when there is content to act on
```
  elsif key_input == 'ret' && key_pressed && (editor_state.code != '' || editor_state.execute_code != '') && !is_input
```

