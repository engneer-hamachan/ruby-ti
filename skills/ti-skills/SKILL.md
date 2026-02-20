---
name: ti-skills
description: Understand and refactor PicoRuby code using ti type checker
---

You are refactoring PicoRuby code using the ti type checker.

## Available Commands

- `ti filename.rb --llm --strict` - Display type information for the file
- `ti filename.rb --llm-error --strict` - Display type error information for the file

**IMPORTANT**: Always run these commands as-is. Never pipe through `head`, `tail`, or any other truncation tool. The full output is required — partial output leads to missing call points and incomplete understanding.

## `ti --llm --strict` is your primary source of truth

**Do not read source files directly.** Use `ti filename.rb --llm --strict` as your main tool for understanding code. It gives you everything you need:

- **Method signatures**: Parameter types and return types at a glance
- **`document:` field**: Intent and behavior of each method — rely on this to understand what a method does
- **Call points**: Every location a method is called from — use this to assess impact before changing anything
- **`untyped` warnings**: Parameters or returns marked `untyped` are ambiguous — treat carefully

### When you may read source code

Only read source when `ti --llm` output is insufficient for a **specific method**:

- The `document:` field is missing or too vague to act on
- You need to understand the exact logic of a single method (use call points to find the line, then read only that method)
- **Never read the entire file** — always target specific line ranges identified from `ti --llm` output

### Reading the output effectively

- `Union<A B>` means the value can be type A or B — handle both cases
- `untyped` means ti couldn't infer the type — these are often the riskiest spots
- `document:` lines describe intent — if they are too vague, consider improving them with ti-add-comments
- Call points list every location a method is used — check all of them when changing a signature

## Workflow

1. **Run `ti filename.rb --llm --strict`** — understand the full method landscape from signatures, documents, and call points
2. **Identify what to change** — use call points to find affected lines; read source only if `document:` is insufficient
3. **Make incremental changes** using the Edit tool
4. **Run `ti filename.rb --llm-error --strict`** after each change to verify no type errors
5. **If errors are found — stop and report** (see "Handling Type Errors" below); do not continue until the user responds
6. **Repeat steps 3–5** until all errors are resolved

## Handling Type Errors - Critical Rules

**NEVER delete code just because it causes type errors.** Instead, properly address the root cause:

1. **Understand the error first**:
   - Read the error message carefully to understand what ti expects
   - Check the line number and context
   - Identify why the type mismatch is occurring

2. **Preserve the original intent**:
   - The refactored code must maintain the same functionality
   - Don't sacrifice code quality just to pass type checking
   - If a helper function causes type errors, fix the function - don't remove it
   - Find alternative implementations that both satisfy ti and improve code quality

3. **Common type error solutions**:
   - For `Union` types containing `NilClass` (e.g., `Union<Integer NilClass>`): Use `is_a?` to narrow the type. ti recognizes `is_a?` checks and narrows the type within the branch:
     ```ruby
     x = arr[0] # Union<Integer NilClass>
     if x.is_a?(Integer)
       # ti narrows x to Integer here
       x + 1
     end
     ```
   - For array operations: Use explicit indexing instead of iterators if needed
   - For hash operations: Ensure consistent access patterns
   - For generic functions: Make them more specific if ti struggles with inference
   - Copy arrays/hashes before modifying to help ti track types

4. **When a type error occurs — stop and report**:
   - **Do not attempt to fix the error on your own.** Stop all work immediately and report to the user.
   - Wait for the user's instruction before proceeding

## Example Workflow

```bash
# Step 1: Check initial type information
ti main/main.rb --llm

# Step 2: Make a small refactoring change (use Edit tool)

# Step 3: Check for type errors
ti main/main.rb --llm-error --strict

# Step 4: If errors found — stop and report to user, wait for instruction

# Step 5: Once no errors, continue with next refactoring
```

Now, which file would you like to work with?
