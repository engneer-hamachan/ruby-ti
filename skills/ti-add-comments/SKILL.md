---
name: ti-add-comments
description: Add ti-doc and ti-for-llm comments to PicoRuby code for ti type checker integration
---

You add special comments to PicoRuby code that are consumed by the `ti` type checker's `--llm` output.

## Usage

The user will specify a file: `ti-add-comments {file name}.rb`

## Rule: All comments MUST be written in English

## Two Comment Types

There are two distinct comment types, each serving a different purpose in `ti --llm` output:

### 1. `# ti-doc:` — Function documentation

- **Purpose**: Describes what a function does
- **Where it appears in `ti --llm`**: `[Method Signatures]` section as the `document:` field
- **Placement**: Always one line above the `def` keyword
- **Required for**: Every function definition (`def`)

Example:
```ruby
# ti-doc: calculate left space for line number
def calculate_line_number_space(line_number)
```

In `ti --llm` output, this appears as:
```
## calculate_line_number_space(Union<untyped Integer>) -> Integer
- file: theme/editor_app.rb:109
- document: calculate left space for line number
- call points:
  - theme/editor_app.rb:955
```

### 2. `# ti-for-llm:` — Important code explanation

- **Purpose**: Explains significant non-function code (variables, conditional branches, complex logic)
- **Where it appears in `ti --llm`**: `[Special Code Comments]` section as the `comment:` field
- **Placement**: Always one line above the target code
- **Required for**: Important variables, complex conditionals, non-obvious logic that LLMs need context to understand

Example:
```ruby
# ti-for-llm: stores the currently selected completion candidate
$completion_chars = nil
```

In `ti --llm` output, this appears as:
```
[Special Code Comments]
## theme/editor_app.rb:90
- comment: stores the currently selected completion candidate
$completion_chars = nil
```

## Workflow

1. **Run `ti {file}.rb --llm`** to see the current state of comments
2. **Read the source file** to understand the code
3. **Add `# ti-doc:` to every function** that is missing one
   - Write a concise English description of what the function does
   - Place it on the line immediately above `def`
4. **Add `# ti-for-llm:` to important non-function code** such as:
   - Global/instance variables with special meaning
   - Complex conditional logic that requires explanation
   - Key business logic or state transitions
   - Non-obvious assignments or calculations
5. **Review all existing comments** — improve unclear or inaccurate ones
6. **Run `ti {file}.rb --llm`** again to verify all comments appear correctly

## Guidelines for Writing Good Comments

- Keep comments concise but descriptive (typically 3-10 words for ti-doc)
- Focus on **what** and **why**, not **how**
- Do not repeat the function name in the ti-doc comment
- For ti-for-llm, explain the intent or significance of the code
- **The smell test**: If your comment could be guessed by just reading the method name, it is not good enough.
  The comment must add information that cannot be inferred from the name alone.
  - Bad: `# ti-doc: calculate indent decrease for tokens` (just a rephrasing of the name)
  - Good: `# ti-doc: returns indent_ct-1 if first token is 'end'/'else'/'elsif'/'when', otherwise returns indent_ct unchanged`

## Goal: LLMs should not need to read source code

The `document:` field in `ti --llm` output is the primary way an LLM understands your code **without reading the source**. Write `ti-doc:` comments rich enough that another LLM can make correct implementation decisions from the `ti --llm` output alone.

### What to include in `ti-doc:` when relevant

- **Return value semantics**: What does the return value mean? (e.g. "returns true if cursor moved to a different line")
- **Special parameter values**: Note sentinel values like `-1` (e.g. "direction: 'left' or 'right'")
- **Side effects**: What state does this function mutate? (e.g. "updates state.code and state.cursor_col_index")
- **Key constraints**: Preconditions or postconditions the caller must know (e.g. "only called when code is non-empty")

### Examples

Too vague (LLM still needs to read source):
```ruby
# ti-doc: handle backspace
def handle_backspace(state)
```

Rich enough to act on:
```ruby
# ti-doc: handle backspace key — deletes char at cursor position, or moves to previous line if code is empty; returns true if cursor moved to a different line
def handle_backspace(state)
```

Too vague:
```ruby
# ti-for-llm: cursor column index
@cursor_col_index = -1
```

Rich enough:
```ruby
# ti-for-llm: cursor column position within current code line — -1 means end of line (default), 0..n means specific position
@cursor_col_index = -1
```
