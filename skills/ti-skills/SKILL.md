---
name: ti-skills
description: Understand and work with PicoRuby code using ti type checker
---

You are working with PicoRuby code using the ti type checker.

## Available Commands

- `ti filename.rb --llm` - Display type information for the file
- `ti filename.rb --llm-error` - Display type error information for the file

**IMPORTANT**: Always run these commands as-is. Never pipe through `head`, `tail`, or any other truncation tool. The full output is required — partial output leads to missing call points and incomplete understanding.

## Step 1: Always start with `ti --llm`

Run `ti filename.rb --llm` first, every time. It gives you:

- **Method signatures**: Parameter types and return types
- **`document:` field**: Intent and behavior of each method
- **Call points**: Every location a method is called from
- **`ti-for-llm` comments**: Inline annotations explaining non-obvious values, invariants, and constraints

### When `document:` or comments are insufficient

If the `document:` field is vague, missing, or doesn't answer a **specific question you already have** — read the source. Use call points to find the exact line range, then read only that range.

Rules:
- Read source only when you have a concrete, specific question that `ti --llm` didn't answer
- Read only the lines needed to answer that question
- Do not read "just in case" or to get a broader picture — form the question first, then read

## Step 2: Before making any change, verify your understanding

For every change you're about to make, write out:

1. **What values can each relevant variable take?**
2. **What are all the scenarios this code handles?** Trace through each with concrete values.
3. **Which scenario is broken, and why?**

Do not skip this. Concrete values prevent wrong assumptions.

## Step 3: Before removing or disabling code

**Never remove code just because it seems related to a complaint.**

First prove the code is the actual cause:
- Trace the specific scenario the user complained about
- Confirm this code triggers in that scenario
- Confirm no other scenario depends on this code correctly

If the complaint points to code X but the real cause is Y (something else), removing X makes things worse. Fix Y instead.

## Step 4: Make incremental changes

- Use the Edit tool for targeted changes
- Run `ti filename.rb --llm-error` after each change
- If type errors appear — stop and report to the user; do not continue

## Handling Type Errors

**Never delete code just because it causes type errors.** Understand the root cause first.

- For `Union<A NilClass>`: use `is_a?` to narrow the type
- For `untyped`: these are the riskiest spots — read the source before touching them
- If you can't resolve an error — stop and report to the user

## Example: Bug Fix Workflow

```bash
# 1. Understand the code
ti filename.rb --llm

# 2. If document: is vague, read the relevant source lines
#    (use call points from ti output to find exact line numbers)

# 3. Trace the broken scenario with actual values

# 4. Identify root cause. Confirm it before touching anything.

# 5. Make the minimal change that fixes the root cause.

# 6. Verify no type errors
ti filename.rb --llm-error

# 7. Trace the scenario again with actual values to confirm the fix.
```
