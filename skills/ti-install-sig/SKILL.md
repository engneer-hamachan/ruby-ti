---
name: ti-install-sig
description: Load project-specific PicoRuby class and method signatures into context using the ti type checker
---

You load the classes and method signatures relevant to the current task into your context using the `ti` type checker.

## Why This Is Needed

In PicoRuby development, the available classes and methods vary per project. This skill lets you discover and internalize the project's type landscape before implementing anything, so you can work with accurate type knowledge.

## Usage

The user will invoke this skill with a prompt describing the task: `ti-install-sig {prompt}`

## Available Commands

### Code understanding (current file)
- `ti filename.rb --llm-nav` - List classes and top-level methods that have callers or callees in the code
- `ti filename.rb --llm-nav --target=Name` - Display detailed signatures, callers/callees for a specific class or method in the code

### API reference (available features in the type system)
- `ti filename.rb --llm-class` - List all classes available in the project
- `ti filename.rb --llm-define --class=ClassName` - Display method signatures and type info for a class

**IMPORTANT**: Always run these commands as-is. Never pipe through `head`, `tail`, or any other truncation tool. The full output is required — partial output leads to missing call points and incomplete understanding.

## Workflow

### Step 1: Review the prompt

Understand what the user is trying to implement and what kinds of types might be involved.

### Step 2: Browse the code structure

Run `ti {target_file.rb} --llm-nav` first. It lists classes and top-level methods that are actively used in the code. Identify which ones are relevant to the task.

### Step 3: Drill into relevant targets

For each relevant class or method identified in Step 2, run:

```bash
ti {target_file.rb} --llm-nav --target=Name
```

This gives you method signatures, callers/callees, file locations, and `document:` fields. Use this to understand what types are already in use and what additional classes you need to load.

### Step 4: Load API reference for relevant classes

1. Run `ti {target_file.rb} --llm-class` to see all available classes in the project.
2. Select relevant classes — from the class list, pick the classes that are likely needed for the task. Consider both what you saw in `--llm-nav` output and what the prompt requires.
3. Run `ti {target_file.rb} --llm-define --class={class}` for each selected class to retrieve detailed method signatures and type information.

Be selective: only pick classes that are clearly relevant to the task. Loading too many is noisy; loading too few means missing needed context.

### Step 5: Report completion

Once all selected classes have been loaded, inform the user that the signatures are ready and summarize what was loaded.

## Notes

- If the target file is not specified in the prompt, ask the user which file to use as the ti target.
