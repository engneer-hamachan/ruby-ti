---
name: ti-navi
description: Understand and work with mruby code using ti type checker
---

You are working with mruby code using the ti type checker.

## Commands

- `ti filename.rb --llm-nav` - List classes and top-level methods with callers/callees
- `ti filename.rb --llm-nav --target=Name` - Detailed signatures and call graph for a specific class or method
- `ti filename.rb --llm-error` - Display type errors
- `ti filename.rb --llm-class` - List all available classes
- `ti filename.rb --llm-define --class=ClassName` - Method signatures and type info for a class

**IMPORTANT**: Never pipe these commands through `head`, `tail`, or any truncation tool. Full output is required.

## Workflow

### 1. Browse the code structure

Run `ti filename.rb --llm-nav` first. Identify which classes and methods are relevant to your task.

### 2. Drill into relevant targets

For each relevant target, run `ti filename.rb --llm-nav --target=Name` to get signatures, callers/callees, and `document:` fields.

If `document:` is insufficient for a specific question you have, use the call points to read only the relevant source lines.

### 3. Make changes and verify

- Use the Edit tool for targeted changes
- Run `ti filename.rb --llm-error` after each change

### Type error hints
- `Union<A NilClass>`: use `is_a?` to narrow the type
