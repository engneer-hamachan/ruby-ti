---
name: ti-install-sig
description: Load project-specific MRuby class and method signatures into context using the ti type checker
---

You load the classes and method signatures relevant to the current task into your context using the `ti` type checker.

## Why This Is Needed

In MRuby development, the available classes and methods vary per project. This skill lets you discover and internalize the project's type landscape before implementing anything, so you can work with accurate type knowledge.

## Usage

The user will invoke this skill with a prompt describing the task: `ti-install-sig {prompt}`

## Workflow

1. **Review the prompt** — understand what the user is trying to implement and what kinds of types might be involved.

2. **Run `ti {target_file.rb} --llm-class`** — get the full list of classes available in the project.

3. **Select relevant classes** — from the class list, pick the classes that are likely needed for the task described in the prompt.

4. **Run `ti {target_file.rb} --llm-define --class={class}`** for each selected class — retrieve detailed method signatures and type information for that class.

5. **Report completion** — once all selected classes have been loaded, inform the user that the signatures are ready and summarize what was loaded.

## Notes

- Be selective at step 3: only pick classes that are clearly relevant to the task. Loading too many is noisy; loading too few means missing needed context.
- If the target file is not specified in the prompt, ask the user which file to use as the ti target.
