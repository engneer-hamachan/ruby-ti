# Method Signatures
## array_remove_at(Array<untyped>, Integer) -> Array<untyped>
- file: theme/editor_app.rb:368
- document: remove item from array at index
- call points:
  - theme/editor_app.rb:311

## calculate_indent_decrease(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:192
- document: calculate indent decrease for tokens
- call points:
  - theme/editor_app.rb:415
  - theme/editor_app.rb:432

## calculate_indent_increase(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:208
- document: calculate indent increase for tokens
- call points:
  - theme/editor_app.rb:418
  - theme/editor_app.rb:435

## calculate_line_number_space(untyped) -> Integer
- file: theme/editor_app.rb:183
- document: calculate left space for line number
- call points:
  - theme/editor_app.rb:1102
  - theme/editor_app.rb:1130

## calculate_max_history_lines(Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1034
- document: calculate max history lines to display
- call points:
  - theme/editor_app.rb:1042
  - theme/editor_app.rb:1084

## calculate_scroll_offset(Array<untyped>, Integer, Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:1040
- document: calculate scroll offset to keep cursor visible
- call points:
  - theme/editor_app.rb:1086

## draw_code_with_highlight(M5GFX, String, Integer, Integer) -> Array<String>
- file: theme/editor_app.rb:879
- document: draw code with syntax highlighting
- call points:
  - theme/editor_app.rb:1109
  - theme/editor_app.rb:1118
  - theme/editor_app.rb:1135

## draw_completion(M5GFX, String) -> NilClass
- file: theme/editor_app.rb:970
- document: search and draw completion candidates
- call points:
  - theme/editor_app.rb:1125

## draw_static_ui(M5GFX) -> NilClass
- file: theme/editor_app.rb:934
- document: draw static UI elements
- call points:
  - theme/editor_app.rb:1178

## execute_and_get_result(Sandbox, untyped) -> ExecutionResult
- file: theme/editor_app.rb:741
- document: execute code and update result
- call points:
  - theme/editor_app.rb:464
  - theme/editor_app.rb:768

## execute_code_and_update_state(EditorState, Sandbox, M5GFX, Integer) -> Integer
- file: theme/editor_app.rb:767
- document: execute state.execute_code via sandbox, store result in state.res/display_res, update completion dict on success, then reset state; always resets on both success and error
- call points:
  - theme/editor_app.rb:1270

## find_completion_candidates(String) -> Array<untyped>
- file: theme/editor_app.rb:953
- document: find completion candidates for a given code string
- call points:
  - theme/editor_app.rb:980

## get_input() -> Array<String Bool>
- file: hardware_adapters/adv_input.rb:124
- document: read keyboard input (returns [key_name, is_pressed])
- call points:
  - theme/editor_app.rb:1247

## get_result_color(Class) -> Integer
- file: theme/editor_app.rb:807
- document: get color code for result display based on result type
- call points:
  - theme/editor_app.rb:1228

## get_token_color(String, Bool, Array<String>) -> Integer
- file: theme/editor_app.rb:861
- document: get syntax highlighting color for a token
- call points:
  - theme/editor_app.rb:926

## handle_backspace(EditorState) -> Bool
- file: theme/editor_app.rb:306
- document: handle backspace key (returns true if cursor moved)
- call points:
  - theme/editor_app.rb:1319

## handle_completion_navigation(String) -> NavigationResult
- file: theme/editor_app.rb:619
- document: handle completion menu navigation
- call points:
  - theme/editor_app.rb:719
  - theme/editor_app.rb:728

## handle_ctrl_input(EditorState, Union<String Bool>) -> Array<Bool>
- file: theme/editor_app.rb:545
- document: handle ctrl key commands (returns [handled, is_need_redraw_input])
- call points:
  - theme/editor_app.rb:1364

## handle_cursor_move(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:590
- document: handle cursor movement within current code line
- call points:
  - theme/editor_app.rb:713

## handle_fn_navigation(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:708
- document: handle function key navigation
- call points:
  - theme/editor_app.rb:1345

## handle_horizontal_scroll(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:565
- document: handle horizontal scroll for result display
- call points:
  - theme/editor_app.rb:715

## handle_normal_char_input(EditorState, Union<String Bool>) -> Union<NilClass String>
- file: theme/editor_app.rb:525
- document: handle normal character input
- call points:
  - theme/editor_app.rb:1376

## handle_return_key(EditorState, Bool, Sandbox, M5GFX, Integer) -> ReturnKeyResult
- file: theme/editor_app.rb:405
- document: handle return key — normal return (is_shift=false) adds new line with auto-indent; shift+return executes full code via sandbox and resets state; returns ReturnKeyResult with should_continue=true on normal return
- call points:
  - theme/editor_app.rb:1277

## handle_shift_input(EditorState, Union<String Bool>) -> Bool
- file: theme/editor_app.rb:497
- document: handle shift key input (returns true if handled)
- call points:
  - theme/editor_app.rb:1331

## handle_tab_key(EditorState) -> Union<NilClass Integer>
- file: theme/editor_app.rb:298
- document: handle tab key for completion
- call points:
  - theme/editor_app.rb:1261

## hash_array_insert(untyped, Hash, untyped) -> Array<untyped Hash>
- file: theme/editor_app.rb:379
- document: insert hash item into array at position
- call points:
  - theme/editor_app.rb:422

## init_keyboard() -> Bool
- file: hardware_adapters/adv_input.rb:73
- document: initialize keyboard
- call points:
  - theme/editor_app.rb:1172

## is_at_last_line?(EditorState) -> Bool
- file: theme/editor_app.rb:235
- document: check if cursor is at the last line
- call points:
  - theme/editor_app.rb:1266

## is_number?(String) -> Bool
- file: theme/editor_app.rb:796
- document: check if a string is a number
- call points:
  - theme/editor_app.rb:868

## load_constants() -> Array<Symbol>
- file: theme/editor_app.rb:169
- document: load Ruby constants into completion dictionary
- call points:
  - theme/editor_app.rb:276
  - theme/editor_app.rb:1194

## move_cursor_down(EditorState) -> NavigationResult
- file: theme/editor_app.rb:676
- document: move cursor down to next line — at last confirmed line with non-empty code, moves to new line at bottom (cursor_row_index=-1); does nothing if code is empty at last line; returns NavigationResult
- call points:
  - theme/editor_app.rb:733

## move_cursor_up(EditorState) -> NavigationResult
- file: theme/editor_app.rb:643
- document: move cursor up to previous line — if on new line (cursor_row_index==-1) with non-empty code, saves it as a confirmed line first; sets cursor_col_index=-1 (end of line); returns NavigationResult
- call points:
  - theme/editor_app.rb:724

## process_definition_tokens(Array<String>, Array<String>, Array<String>, Array<String>) -> Array<String>
- file: theme/editor_app.rb:246
- document: process definition tokens for completion dictionary
- call points:
  - theme/editor_app.rb:288

## rebuild_execute_code(untyped) -> String
- file: theme/editor_app.rb:394
- document: rebuild execute_code from code_lines
- call points:
  - theme/editor_app.rb:331
  - theme/editor_app.rb:360
  - theme/editor_app.rb:424
  - theme/editor_app.rb:434
  - theme/editor_app.rb:515
  - theme/editor_app.rb:540
  - theme/editor_app.rb:650

## redraw_code_area(M5GFX, Array<untyped>, Integer, Integer, String, untyped, Integer, Integer, String, untyped) -> Integer
- file: theme/editor_app.rb:1064
- document: redraw code area with scroll offset
- call points:
  - theme/editor_app.rb:444
  - theme/editor_app.rb:478
  - theme/editor_app.rb:780
  - theme/editor_app.rb:1179
  - theme/editor_app.rb:1202

## EditorState.reset() -> Integer
- file: theme/editor_app.rb:83
- document: reset editor state to initial values
- call points:
  - theme/editor_app.rb:241

## reset_editor_state(EditorState) -> Integer
- file: theme/editor_app.rb:240
- document: reset editor state to initial values and reset completion selection index to 0
- call points:
  - theme/editor_app.rb:470
  - theme/editor_app.rb:476
  - theme/editor_app.rb:550
  - theme/editor_app.rb:773
  - theme/editor_app.rb:778

## tokenize(untyped) -> Array<String>
- file: theme/editor_app.rb:820
- document: split code into tokens
- call points:
  - theme/editor_app.rb:287
  - theme/editor_app.rb:411
  - theme/editor_app.rb:914
  - theme/editor_app.rb:976
  - theme/editor_app.rb:1149

## update_completion_dict(untyped) -> Union<Bool NilClass>
- file: theme/editor_app.rb:281
- document: update completion dictionary from code lines
- call points:
  - theme/editor_app.rb:473
  - theme/editor_app.rb:775

# Special Code Comments
## theme/editor_app.rb:61
- comment: holds all editor state variables in a single object to reduce Hash return values
```
class EditorState
```

## theme/editor_app.rb:96
- comment: holds fn navigation results for redraw flags
```
class NavigationResult
```

## theme/editor_app.rb:107
- comment: holds code execution results including error status
```
class ExecutionResult
```

## theme/editor_app.rb:119
- comment: holds return key handling results for control flow
```
class ReturnKeyResult
```

## theme/editor_app.rb:130
- comment: ADC object for reading battery voltage
```
bat_adc = ADC.new(10)
```

## theme/editor_app.rb:133
- comment: sandbox for executing user-entered mruby code
```
sandbox = Sandbox.new ''
```

## theme/editor_app.rb:136
- comment: flag to prevent duplicate key processing in same frame
```
is_input = false
```

## theme/editor_app.rb:138
- comment: true while shift key is held — causes characters to be looked up in SHIFT_TABLE
```
is_shift = false
```

## theme/editor_app.rb:140
- comment: true while fn key is held — causes keys to be looked up in FN_TABLE for navigation
```
is_fn = false
```

## theme/editor_app.rb:142
- comment: true while ctrl key is held — enables ctrl+c (cancel line) and ctrl+d (reset editor)
```
is_ctrl = false
```

## theme/editor_app.rb:144
- comment: set to true to force redraw of the code input area on the next loop iteration
```
is_need_redraw_input = false
```

## theme/editor_app.rb:146
- comment: set to true to force redraw of the result display area on the next loop iteration
```
is_need_redraw_result = false
```

## theme/editor_app.rb:148
- comment: previous code display string used to detect changes and trigger input area redraw
```
prev_code_display = ''
```

## theme/editor_app.rb:150
- comment: previous result string used to detect changes and trigger result area redraw
```
prev_res = ''
```

## theme/editor_app.rb:152
- comment: previous status bar string used to detect changes and trigger status bar redraw
```
prev_status = ''
```

## theme/editor_app.rb:154
- comment: maximum number of code lines visible in the code area at once (scroll window size)
```
max_visible_lines = 7
```

## theme/editor_app.rb:157
- comment: stores the currently selected completion candidate string
```
$completion_chars = nil
```

## theme/editor_app.rb:159
- comment: list of current completion candidate strings
```
$completion_candidates = []
```

## theme/editor_app.rb:161
- comment: index of currently highlighted completion candidate
```
$completion_selected_index = 0
```

## theme/editor_app.rb:163
- comment: dictionary of known identifiers for autocompletion
```
$dict = {}
```

## theme/editor_app.rb:166
- comment: main editor state instance holding all editor variables
```
editor_state = EditorState.new
```

## theme/editor_app.rb:1148
- comment: side effect — if any code_line starts with 'class', inject attr_reader/attr_accessor/initialize into the completion dict so class body methods appear as candidates
```
  code_lines.each do |line|
```

## theme/editor_app.rb:1266
- comment: execute code when return is pressed on the last line with empty input
```
  if key_input == 'ret' && key_pressed && is_at_last_line?(editor_state) && editor_state.code == '' && editor_state.execute_code != '' && !is_input
```

## theme/editor_app.rb:1274
- comment: normal return — add new line (code != '') or execute with shift+return; only fires when there is content to act on
```
  elsif key_input == 'ret' && key_pressed && (editor_state.code != '' || editor_state.execute_code != '') && !is_input
```

