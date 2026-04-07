---
name: ti-add-comments
description: Add ti-doc and ti-for-llm comments to mruby code for ti type checker integration
---

You add special comments to mruby code that are consumed by the `ti` type checker's `--llm-nav` output.

## Usage

The user will specify a file: `ti-add-comments {file name}.rb`

## Rule: All comments MUST be written in English

## Available Commands

### Code understanding (current file)
- `ti filename.rb --llm-nav` - List classes and top-level methods that have callers or callees in the code
- `ti filename.rb --llm-nav --target=Name` - Display detailed signatures, callers/callees for a specific class or method in the code

**IMPORTANT**: Always run these commands as-is. Never pipe through `head`, `tail`, or any other truncation tool. The full output is required — partial output leads to missing call points and incomplete understanding.

## Two Comment Types

There are two distinct comment types, each serving a different purpose in `ti --llm-nav` output:

### 1. `# ti-doc:` — Function documentation

- **Purpose**: Describes what a function does
- **Where it appears in `ti --llm-nav --target=Name`**: `[Method Signatures]` section as the `document:` field
- **Placement**: Always one line above the `def` keyword
- **Required for**: Every function definition (`def`)

Example:
```ruby
# ti-doc: calculate left space for line number
def calculate_line_number_space(line_number)
```

In `ti --llm-nav --target=calculate_line_number_space` output, this appears as:
```
## calculate_line_number_space(Union<untyped Integer>) -> Integer
- file: theme/editor_app.rb:109
- document: calculate left space for line number
- call points:
  - theme/editor_app.rb:955
```

### 2. `# ti-for-llm:` — Important code explanation

- **Purpose**: Explains significant non-function code (variables, conditional branches, complex logic)
- **Where it appears in `ti --llm-nav`**: `[Special Code Comments]` section as the `comment:` field
- **Placement**: Always one line above the target code
- **Required for**: Important variables, complex conditionals, non-obvious logic that LLMs need context to understand

Example:
```ruby
# ti-for-llm: stores the currently selected completion candidate
$completion_chars = nil
```

In `ti --llm-nav` output, this appears as:
```
[Special Code Comments]
## theme/editor_app.rb:90
- comment: stores the currently selected completion candidate
$completion_chars = nil
```

## Workflow

### Step 1: Browse the code structure

Run `ti {file}.rb --llm-nav` first. It lists classes and top-level methods. Identify which ones are missing `document:` fields or have vague ones.

### Step 2: Drill into each target

For each class or method that needs comments, run:

```bash
ti {file}.rb --llm-nav --target=Name
```

This shows you the method signatures, existing `document:` fields, and call points. Use this to understand what each method does before writing its comment.

Do not fetch all details upfront — use `--llm-nav` to identify what needs work, then `--target` to drill in.

### Step 3: Read source when needed

If the `--target` output doesn't give you enough context to write a good comment, use call points to find the exact line range and read only that range.

Rules:
- Read source only when you have a concrete, specific question that the output didn't answer
- Read only the lines needed to answer that question
- Do not read "just in case" or to get a broader picture — form the question first, then read

### Step 4: Add comments

1. **Add `# ti-doc:` to every function** that is missing one
   - Write a concise English description of what the function does
   - Place it on the line immediately above `def`
2. **Add `# ti-for-llm:` to important non-function code** such as:
   - Global/instance variables with special meaning
   - Complex conditional logic that requires explanation
   - Key business logic or state transitions
   - Non-obvious assignments or calculations
3. **Review all existing comments** — improve unclear or inaccurate ones

### Step 5: Verify

Run `ti {file}.rb --llm-nav` and drill into updated targets with `--target=Name` to confirm all comments appear correctly.

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

The `document:` field in `ti --llm-nav --target=Name` output is the primary way an LLM understands your code **without reading the source**. Write `ti-doc:` comments rich enough that another LLM can make correct implementation decisions from the output alone.

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
