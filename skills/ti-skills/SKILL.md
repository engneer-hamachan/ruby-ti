---
name: ti-skills
description: Understand and work with PicoRuby code using ti type checker
---

You are working with PicoRuby code using the ti type checker.

## Available Commands

### Code understanding (current file)
- `ti filename.rb --llm-nav` - List classes and top-level methods that have callers or callees in the code
- `ti filename.rb --llm-nav --target=Name` - Display detailed signatures, callers/callees for a specific class or method in the code
- `ti filename.rb --llm-error` - Display type error information for the file

### API reference (available features in the type system)
- `ti filename.rb --llm-class` - List all classes available in the project
- `ti filename.rb --llm-define --class=ClassName` - Display method signatures and type info for a class

**IMPORTANT**: Always run these commands as-is. Never pipe through `head`, `tail`, or any other truncation tool. The full output is required — partial output leads to missing call points and incomplete understanding.

## Step 1: Browse the code structure

Run `ti filename.rb --llm-nav` first. It lists classes and top-level methods that are actively used in the code. Identify which ones are relevant to your task.

## Step 2: Get details for relevant targets

For each relevant class or method identified in Step 1, run:

```bash
ti filename.rb --llm-nav --target=Name
```

This gives you method signatures, callers/callees, file locations, and `document:` fields. Run this for each target you need to understand.

Do not fetch all details upfront — use `--llm-nav` to identify what to look at, then `--target` to drill in.

### When `document:` or comments are insufficient

If the `document:` field is vague, missing, or doesn't answer a **specific question you already have** — read the source. Use call points to find the exact line range, then read only that range.

Rules:
- Read source only when you have a concrete, specific question that the output didn't answer
- Read only the lines needed to answer that question
- Do not read "just in case" or to get a broader picture — form the question first, then read

## Step 3: Before making any change, verify your understanding

For every change you're about to make, write out:

1. **What values can each relevant variable take?**
2. **What are all the scenarios this code handles?** Trace through each with concrete values.
3. **Which scenario is broken, and why?**

Do not skip this. Concrete values prevent wrong assumptions.

## Step 4: Before removing or disabling code

**Never remove code just because it seems related to a complaint.**

First prove the code is the actual cause:
- Trace the specific scenario the user complained about
- Confirm this code triggers in that scenario
- Confirm no other scenario depends on this code correctly

If the complaint points to code X but the real cause is Y (something else), removing X makes things worse. Fix Y instead.

## Step 5: Make incremental changes

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
# 1. Browse the code structure
ti filename.rb --llm-nav

# 2. Drill into relevant classes/methods
ti filename.rb --llm-nav --target=RelevantClass
ti filename.rb --llm-nav --target=relevant_method

# 3. If document: is vague, read the relevant source lines
#    (use call points from output to find exact line numbers)

# 4. Trace the broken scenario with actual values

# 5. Identify root cause. Confirm it before touching anything.

# 6. Make the minimal change that fixes the root cause.

# 7. Verify no type errors
ti filename.rb --llm-error

# 8. Trace the scenario again with actual values to confirm the fix.
```
