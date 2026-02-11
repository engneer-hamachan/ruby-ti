# Method Signatures
## array_remove_at(Array<untyped>, untyped) -> Array<untyped>
- file: theme/editor_app.rb:320
- document: remove item from array at index
- call points:
  - theme/editor_app.rb:644
  - theme/editor_app.rb:252

## calculate_indent_decrease(Array<String>, Union<Integer untyped>) -> Union<Integer Float untyped>
- file: theme/editor_app.rb:125
- document: calculate indent decrease for tokens
- call points:
  - theme/editor_app.rb:379
  - theme/editor_app.rb:396

## calculate_indent_increase(Array<String>, untyped) -> untyped
- file: theme/editor_app.rb:141
- document: calculate indent increase for tokens
- call points:
  - theme/editor_app.rb:382
  - theme/editor_app.rb:399

## calculate_line_number_space(Union<untyped Integer>) -> Integer
- file: theme/editor_app.rb:116
- document: calculate left space for line number
- call points:
  - theme/editor_app.rb:962
  - theme/editor_app.rb:984

## draw_code_with_highlight(M5GFX, String, Integer, Integer) -> Array<String>
- file: theme/editor_app.rb:747
- document: draw code with syntax highlighting
- call points:
  - theme/editor_app.rb:969
  - theme/editor_app.rb:973
  - theme/editor_app.rb:989

## draw_completion(M5GFX, String) -> Union<NilClass Array<untyped> untyped>
- file: theme/editor_app.rb:843
- document: search and draw completion candidates
- call points:
  - theme/editor_app.rb:980

## draw_static_ui(M5GFX) -> NilClass
- file: theme/editor_app.rb:824
- document: draw static UI elements
- call points:
  - theme/editor_app.rb:1025

## execute_and_get_result(Sandbox, String) -> Hash
- file: theme/editor_app.rb:670
- document: execute code and update result
- call points:
  - theme/editor_app.rb:1111
  - theme/editor_app.rb:426

## get_input() -> String
- file: hardware_adapters/adv_input.rb:124
- document: read keyboard input
- call points:
  - theme/editor_app.rb:1098

## handle_backspace(String, Array<untyped>, String, Integer, Integer, Integer) -> Hash
- file: theme/editor_app.rb:249
- document: handle backspace key
- call points:
  - theme/editor_app.rb:1217

## handle_ctrl_input(String, String, String, Array<untyped>, Integer, Union<untyped Integer>, Integer, Union<untyped Integer>, String) -> Hash
- file: theme/editor_app.rb:528
- document: handle ctrl key commands
- call points:
  - theme/editor_app.rb:1294

## handle_fn_navigation(untyped, Integer, Union<NilClass String>, Array<Unknown>, String, Integer, Union<Integer Array<Symbol>>, Union<Integer untyped>, String, String) -> Hash
- file: theme/editor_app.rb:566
- document: handle function key navigation
- call points:
  - theme/editor_app.rb:1256

## handle_normal_char_input(String, String, Array<untyped>, String, Integer) -> Hash
- file: theme/editor_app.rb:510
- document: handle normal character input
- call points:
  - theme/editor_app.rb:1315

## handle_return_key(String, String, Array<untyped>, Integer, Integer, Integer, String, Bool, Sandbox, M5GFX, Integer, Integer, Integer) -> Hash
- file: theme/editor_app.rb:366
- document: handle return key press
- call points:
  - theme/editor_app.rb:1155

## handle_shift_input(String, Union<String Array<Symbol>>, Array<Unknown>, String, Union<Integer untyped>) -> Hash
- file: theme/editor_app.rb:484
- document: handle shift key input
- call points:
  - theme/editor_app.rb:1237

## handle_tab_key(String) -> Hash
- file: theme/editor_app.rb:239
- document: handle tab key for completion
- call points:
  - theme/editor_app.rb:1101

## hash_array_insert(Array<untyped>, Union<String untyped>, Union<untyped Integer>) -> Array<untyped String>
- file: theme/editor_app.rb:331
- document: insert hash item into array at position
- call points:
  - theme/editor_app.rb:386

## init_keyboard() -> Bool
- file: hardware_adapters/adv_input.rb:73
- document: initialize keyboard
- call points:
  - theme/editor_app.rb:1019

## is_number?(String) -> Bool
- file: theme/editor_app.rb:696
- document: check if a string is a number
- call points:
  - theme/editor_app.rb:798

## load_constants() -> Array<Symbol>
- file: theme/editor_app.rb:102
- document: load Ruby constants into completion dictionary
- call points:
  - theme/editor_app.rb:225
  - theme/editor_app.rb:1039

## rebuild_execute_code(Array<Unknown>) -> String
- file: theme/editor_app.rb:355
- document: rebuild execute_code from code_lines
- call points:
  - theme/editor_app.rb:388
  - theme/editor_app.rb:398
  - theme/editor_app.rb:495
  - theme/editor_app.rb:517
  - theme/editor_app.rb:602
  - theme/editor_app.rb:645
  - theme/editor_app.rb:271
  - theme/editor_app.rb:304

## redraw_code_area(M5GFX, Array<untyped>, Integer, Integer, String, Union<Integer untyped>, Integer, untyped) -> Integer
- file: theme/editor_app.rb:911
- document: redraw code area with scroll offset
- call points:
  - theme/editor_app.rb:1026
  - theme/editor_app.rb:1047
  - theme/editor_app.rb:1137
  - theme/editor_app.rb:408
  - theme/editor_app.rb:454

## reset_editor_state(Union<String untyped>, String, Array<untyped>, Union<Integer untyped>, Union<untyped Integer>, Integer, Union<untyped Integer>) -> Hash
- file: theme/editor_app.rb:168
- document: reset editor state
- call points:
  - theme/editor_app.rb:432
  - theme/editor_app.rb:445
  - theme/editor_app.rb:533
  - theme/editor_app.rb:1116
  - theme/editor_app.rb:1128

## tokenize(Union<String untyped>) -> Array<String>
- file: theme/editor_app.rb:706
- document: split code into tokens
- call points:
  - theme/editor_app.rb:782
  - theme/editor_app.rb:849
  - theme/editor_app.rb:996
  - theme/editor_app.rb:192
  - theme/editor_app.rb:375

## update_completion_dict(Array<untyped>) -> Union<Bool NilClass>
- file: theme/editor_app.rb:190
- document: update completion dictionary from code lines
- call points:
  - theme/editor_app.rb:442
  - theme/editor_app.rb:1125

# Special Code Comments
## theme/editor_app.rb:61
- comment: ADC object for reading battery voltage
```
bat_adc = ADC.new(10)
```

## theme/editor_app.rb:64
- comment: sandbox for executing user-entered mruby code
```
sandbox = Sandbox.new ''
```

## theme/editor_app.rb:67
- comment: flag to prevent duplicate key processing in same frame
```
is_input = false
```

## theme/editor_app.rb:86
- comment: accumulated source code string from all confirmed lines
```
execute_code = ''
```

## theme/editor_app.rb:88
- comment: -1 means new line at bottom, >= 0 means editing confirmed line at that index
```
cursor_row_index = -1
```

## theme/editor_app.rb:91
- comment: temporarily stores new line text when cursor moves to a confirmed line
```
temp_new_line_code = ''
```

## theme/editor_app.rb:93
- comment: stores the currently selected completion candidate string
```
$completion_chars = nil
```

## theme/editor_app.rb:95
- comment: list of current completion candidate strings
```
$completion_candidates = []
```

## theme/editor_app.rb:97
- comment: index of currently highlighted completion candidate
```
$completion_selected_index = 0
```

## theme/editor_app.rb:99
- comment: dictionary of known identifiers for autocompletion
```
$dict = {}
```

## theme/editor_app.rb:1107
- comment: execute code when return is pressed on the last line with empty input
```
  if key_input == 'ret' && (cursor_row_index == -1 || cursor_row_index == code_lines.length - 1) && code == '' && execute_code != '' && !is_input
```

