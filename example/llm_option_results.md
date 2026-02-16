# Method Signatures
## array_remove_at(Array<untyped>, Integer) -> Array<untyped>
- file: theme/editor_app.rb:330
- document: remove item from array at index
- call points:
  - theme/editor_app.rb:288

## calculate_indent_decrease(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:177
- document: calculate indent decrease for tokens
- call points:
  - theme/editor_app.rb:377
  - theme/editor_app.rb:394

## calculate_indent_increase(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:193
- document: calculate indent increase for tokens
- call points:
  - theme/editor_app.rb:380
  - theme/editor_app.rb:397

## calculate_line_number_space(Integer) -> Integer
- file: theme/editor_app.rb:168
- document: calculate left space for line number
- call points:
  - theme/editor_app.rb:909
  - theme/editor_app.rb:932

## draw_code_with_highlight(M5GFX, String, Integer, Integer) -> Array<String>
- file: theme/editor_app.rb:678
- document: draw code with syntax highlighting
- call points:
  - theme/editor_app.rb:916
  - theme/editor_app.rb:920
  - theme/editor_app.rb:937

## draw_completion(M5GFX, String) -> NilClass
- file: theme/editor_app.rb:791
- document: search and draw completion candidates
- call points:
  - theme/editor_app.rb:927

## draw_static_ui(M5GFX) -> NilClass
- file: theme/editor_app.rb:755
- document: draw static UI elements
- call points:
  - theme/editor_app.rb:975

## execute_and_get_result(Sandbox, String) -> ExecutionResult
- file: theme/editor_app.rb:600
- document: execute code and update result
- call points:
  - theme/editor_app.rb:425
  - theme/editor_app.rb:1076

## find_completion_candidates(String) -> Array<untyped>
- file: theme/editor_app.rb:774
- document: find completion candidates for a given code string
- call points:
  - theme/editor_app.rb:801

## get_input() -> Array<String Bool>
- file: hardware_adapters/adv_input.rb:124
- document: read keyboard input (returns [key_name, is_pressed])
- call points:
  - theme/editor_app.rb:1050

## handle_backspace(EditorState) -> Bool
- file: theme/editor_app.rb:283
- document: handle backspace key (returns true if cursor moved)
- call points:
  - theme/editor_app.rb:1147

## handle_ctrl_input(EditorState, String) -> Array<Bool>
- file: theme/editor_app.rb:489
- document: handle ctrl key commands (returns [handled, is_need_redraw_input])
- call points:
  - theme/editor_app.rb:1192

## handle_fn_navigation(EditorState, untyped) -> NavigationResult
- file: theme/editor_app.rb:509
- document: handle function key navigation
- call points:
  - theme/editor_app.rb:1173

## handle_normal_char_input(EditorState, String) -> Union<NilClass String>
- file: theme/editor_app.rb:477
- document: handle normal character input
- call points:
  - theme/editor_app.rb:1204

## handle_return_key(EditorState, Bool, Sandbox, M5GFX, Integer) -> ReturnKeyResult
- file: theme/editor_app.rb:367
- document: handle return key press
- call points:
  - theme/editor_app.rb:1105

## handle_shift_input(EditorState, String) -> Bool
- file: theme/editor_app.rb:457
- document: handle shift key input (returns true if handled)
- call points:
  - theme/editor_app.rb:1159

## handle_tab_key(EditorState) -> Union<NilClass Integer>
- file: theme/editor_app.rb:275
- document: handle tab key for completion
- call points:
  - theme/editor_app.rb:1067

## hash_array_insert(Array<untyped>, Hash, Integer) -> Array<untyped Hash>
- file: theme/editor_app.rb:341
- document: insert hash item into array at position
- call points:
  - theme/editor_app.rb:384

## init_keyboard() -> Bool
- file: hardware_adapters/adv_input.rb:73
- document: initialize keyboard
- call points:
  - theme/editor_app.rb:969

## is_number?(String) -> Bool
- file: theme/editor_app.rb:626
- document: check if a string is a number
- call points:
  - theme/editor_app.rb:729

## load_constants() -> Array<Symbol>
- file: theme/editor_app.rb:154
- document: load Ruby constants into completion dictionary
- call points:
  - theme/editor_app.rb:261
  - theme/editor_app.rb:990

## rebuild_execute_code(Array<untyped>) -> String
- file: theme/editor_app.rb:356
- document: rebuild execute_code from code_lines
- call points:
  - theme/editor_app.rb:307
  - theme/editor_app.rb:322
  - theme/editor_app.rb:386
  - theme/editor_app.rb:396
  - theme/editor_app.rb:467
  - theme/editor_app.rb:484
  - theme/editor_app.rb:544

## redraw_code_area(M5GFX, Array<untyped>, Integer, Integer, String, Union<untyped Integer>, Integer, Integer, String) -> Integer
- file: theme/editor_app.rb:855
- document: redraw code area with scroll offset
- call points:
  - theme/editor_app.rb:406
  - theme/editor_app.rb:439
  - theme/editor_app.rb:976
  - theme/editor_app.rb:998
  - theme/editor_app.rb:1088

## EditorState.reset() -> Integer
- file: theme/editor_app.rb:81
- document: reset editor state to initial values
- call points:
  - theme/editor_app.rb:221

## reset_editor_state(EditorState) -> Integer
- file: theme/editor_app.rb:220
- document: reset editor state
- call points:
  - theme/editor_app.rb:431
  - theme/editor_app.rb:437
  - theme/editor_app.rb:494
  - theme/editor_app.rb:1081
  - theme/editor_app.rb:1086

## tokenize(Union<untyped String>) -> Array<String>
- file: theme/editor_app.rb:637
- document: split code into tokens
- call points:
  - theme/editor_app.rb:228
  - theme/editor_app.rb:373
  - theme/editor_app.rb:713
  - theme/editor_app.rb:797
  - theme/editor_app.rb:946

## update_completion_dict(Union<untyped Array<untyped>>) -> Union<Bool NilClass>
- file: theme/editor_app.rb:226
- document: update completion dictionary from code lines
- call points:
  - theme/editor_app.rb:434
  - theme/editor_app.rb:1083

# Special Code Comments
## theme/editor_app.rb:61
- comment: holds all editor state variables in a single object to reduce Hash return values
```
class EditorState
```

## theme/editor_app.rb:93
- comment: holds fn navigation results for redraw flags
```
class NavigationResult
```

## theme/editor_app.rb:103
- comment: holds code execution results including error status
```
class ExecutionResult
```

## theme/editor_app.rb:114
- comment: holds return key handling results for control flow
```
class ReturnKeyResult
```

## theme/editor_app.rb:124
- comment: ADC object for reading battery voltage
```
bat_adc = ADC.new(10)
```

## theme/editor_app.rb:127
- comment: sandbox for executing user-entered mruby code
```
sandbox = Sandbox.new ''
```

## theme/editor_app.rb:130
- comment: flag to prevent duplicate key processing in same frame
```
is_input = false
```

## theme/editor_app.rb:142
- comment: stores the currently selected completion candidate string
```
$completion_chars = nil
```

## theme/editor_app.rb:144
- comment: list of current completion candidate strings
```
$completion_candidates = []
```

## theme/editor_app.rb:146
- comment: index of currently highlighted completion candidate
```
$completion_selected_index = 0
```

## theme/editor_app.rb:148
- comment: dictionary of known identifiers for autocompletion
```
$dict = {}
```

## theme/editor_app.rb:151
- comment: main editor state instance holding all editor variables
```
editor_state = EditorState.new
```

## theme/editor_app.rb:1072
- comment: execute code when return is pressed on the last line with empty input
```
  if key_input == 'ret' && key_pressed && (editor_state.cursor_row_index == -1 || editor_state.cursor_row_index == editor_state.code_lines.length - 1) && editor_state.code == '' && editor_state.execute_code != '' && !is_input
```

