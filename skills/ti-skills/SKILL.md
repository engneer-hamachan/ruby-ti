---
name: ti-skills
description: Understand and refactor PicoRuby code using ti type checker
---

You are refactoring PicoRuby code using the ti type checker.

## Available Commands

- `ti filename.rb --llm` - Display type information for the file
- `ti filename.rb --llm-error` - Display type error information for the file

## Using `ti --llm` for Code Understanding (IMPORTANT)

**Before reading a file's source code**, always run `ti filename.rb --llm` first. This gives you:

- **Method signatures**: Parameter types and return types at a glance
- **Call points**: Where each method is called from — use this to understand impact/dependencies
- **`untyped` warnings**: Parameters or returns marked `untyped` indicate areas where types are ambiguous — read these carefully

### When to use `ti --llm`

- **Before editing a file**: Understand the full method landscape before touching anything
- **When investigating a bug**: Check call points to trace how data flows through the code
- **When assessing change impact**: Call points tell you exactly which lines will be affected
- **When exploring unfamiliar code**: Method signatures + documents give a quick overview without reading every line

### Reading the output effectively

- `Union<A B>` means the value can be type A or B — handle both cases
- `untyped` means ti couldn't infer the type — these are often the riskiest spots
- `document:` lines are inline comments from the code — use them to understand intent
- Call points list every location a method is used — check all of them when changing a signature

## Refactoring Workflow

1. **Start by checking type information**: Run `ti filename.rb --llm` to understand the current type state
2. **Make incremental changes**: Refactor the code in small steps
3. **Verify frequently**: After each change, run `ti filename.rb --llm-error` to check for type errors
4. **Prioritize ti compliance**: While the ti command is not perfect, prioritize making changes that result in no type errors from ti

## Handling Type Errors - Critical Rules

**NEVER delete code just because it causes type errors.** Instead, properly address the root cause:

1. **Understand the error first**:
   - Read the error message carefully to understand what ti expects
   - Check the line number and context
   - Identify why the type mismatch is occurring

2. **Adjust to ti's type inference constraints**:
   - ti may have limitations with certain patterns (e.g., `each` vs `while` loops)
   - If a modern Ruby idiom causes errors, try a more explicit approach
   - Initialize arrays/variables in ways that help ti infer types correctly
   - Example: If `arr.each` fails, try `while i < arr.length` with index access

3. **Preserve the original intent**:
   - The refactored code must maintain the same functionality
   - Don't sacrifice code quality just to pass type checking
   - If a helper function causes type errors, fix the function - don't remove it
   - Find alternative implementations that both satisfy ti and improve code quality

4. **Common type error solutions**:
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

5. **When stuck**:
   - Try breaking the operation into smaller, more explicit steps
   - Use intermediate variables with clear types
   - Consult the original working code for patterns that ti accepts
   - Remember: The goal is correct, maintainable code that passes type checking

## Example Workflow

```bash
# Step 1: Check initial type information
ti main/main.rb --llm

# Step 2: Make a small refactoring change (use Edit tool)

# Step 3: Check for type errors
ti main/main.rb --llm-error

# Step 4: If errors found, fix them and repeat step 3

# Step 5: Once no errors, continue with next refactoring
```

Now, which file would you like to work with?
