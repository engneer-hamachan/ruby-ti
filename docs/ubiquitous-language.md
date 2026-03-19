# Ubiquitous Language: Ruby-TI

A static type analyzer for MRuby that performs type inference and checking without requiring type annotations.

---

## Core Type System

| Term | Definition |
|------|-----------|
| **T** | The universal type representation. Every inferred type is a `T` with a type category (`tType`), an object class name (`objectClass`), optional union variants, method definition args, block parameters, and metadata flags such as `IsStatic`, `IsDestructive`, `IsConditionalReturn`, and `IsBuiltinAsterisk`. |
| **tType** | An integer constant on `T` that classifies the kind of type: object, method, array, hash, union, identifier, unknown, etc. Used by predicate methods like `IsArrayType`, `IsHashType`, `IsUnionType` to branch evaluation logic. |
| **objectClass** | A string field on `T` naming the Ruby class (e.g. `"String"`, `"Integer"`, `"Array"`) that this type represents. |
| **Variant** | An element in the `variants` slice of a `T`. Arrays store their element type as a variant; hashes store key and value types; unions store each possible type. Methods like `AppendArrayVariant`, `AppendHashVariant`, and `AppendVariant` add variants, while `UnifyVariants` merges them. |
| **Union Type** | A `T` whose `variants` slice contains multiple possible types. Created by `MakeUnion` or `MakeUnifiedT`. Represented in output as `Type1 \| Type2`. |
| **Power** | A numeric priority (`int8`) assigned to identifiers and operators via `GetPowerByString`. Determines parsing precedence during expression evaluation --- higher power binds tighter. |
| **TypeSpec** | An alias type used to define type-category constants consumed by `T.tType`. |

---

## Type Storage

| Term | Definition |
|------|-----------|
| **TFrame** | A global `map[FrameKey]*T` storing all variable types discovered during analysis. Every variable, method, and constant type is written to and read from TFrame using composite keys. Functions like `SetValueT`, `GetValueT`, `SetMethodT`, `GetMethodT` operate on this map. |
| **FrameKey** | The composite key for TFrame lookups: `{frame, targetClass, targetMethod, targetVariable, isPrivate, isStatic}`. Different key-builder functions (`valueTFrameKey`, `methodTFrameKey`, `classMethodTFrameKey`, `constTFrameKey`) construct keys for different storage categories. |
| **Frame** | A string identifying the lexical scope of a type entry. For top-level code it is the file name; for nested classes it is computed by `CalculateFrame(frame, class)` which concatenates parent frame and class name. |
| **TSignatures** | A global map storing method signatures indexed by a method-signature string key. Each entry is a `Sig` containing method name, detail string, class, frame, visibility, file location, and documentation. Used for LSP completions, hover info, and define-info output. |
| **Sig** | A single method signature record in TSignatures. Fields: `Method` (name), `Detail` (signature string), `Frame`, `Class`, `IsStatic`, `IsPrivate`, `FileName`, `Row`, `Document`. |
| **ArgumentSnapShot** | A mechanism (`SnapShotArgumentTypes` / `RestoreArgumentTypes`) that temporarily saves and restores variable types in TFrame before and after evaluating a method body. Allows argument types inferred from call sites to propagate across analysis rounds. |
| **DefinedClass** | A record of `{frame, class}` tracking which classes have been defined during analysis. Stored in a global registry and used to resolve class references and inheritance. |

---

## Analysis Pipeline

| Term | Definition |
|------|-----------|
| **Round** | One of three sequential analysis passes returned by `GetRounds()`: **collect**, **inference**, and **check**. Each round re-evaluates the entire file (including preloaded files) to progressively build and verify type information. |
| **Collect Round** | The first round. Collects method definitions, class structures, and method signatures into TFrame and TSignatures. Does not validate types. Checked via `Context.IsDefineRound()`. |
| **Inference Round** | The second round. Infers types based on usage patterns, propagates argument types via ArgumentSnapShot, and refines TFrame entries. |
| **Check Round** | The third and final round. Validates type consistency and reports type errors. Checked via `Context.IsCheckRound()`. |
| **Context** | The analysis context threaded through all evaluation functions. Tracks the current `frame`, `class`, `method`, analysis `round`, and state flags: `IsPrivate`, `IsProtected`, `IsBind`, `isMultiValue`, `IsCallArg`, `IsDefineArg`, `IsArrayCollect`, `IsDefineStatic`. Methods like `SetClass`, `SetMethod`, `StartPrivate`/`EndPrivate` manage scope transitions. |
| **Preload** | Files listed in `.ti-loader.json` that are analyzed before the main target file in each round. Establishes shared type definitions (e.g. common classes) that the target file depends on. Managed by `TiLoader` and `GetPreloadFiles()`. |

---

## Lexing and Parsing

| Term | Definition |
|------|-----------|
| **Lexer** | Tokenizes Ruby source code character by character. Reads from a `LexerReader`, produces a token (`rune`) and a value (`any`) per `Advance()` call. Handles identifiers, digits, strings, operators, comments, and whitespace tracking (`IsSpace`, `IsSpacePrev`). |
| **LexerReader** | A character-stream reader wrapping a `[]rune` buffer with `Read()` / `Unread()` and a `history` buffer for multi-character lookback during digit parsing. |
| **Parser** | Consumes tokens from the Lexer and drives AST evaluation. Tracks current and last tokens (`CurrentT`, `LastT`), the last evaluated type result (`lastEvaluatedT`), method return types (`lastReturnT`), errors, define-info output, LSP target row, and parsing-expression mode. Key methods: `Read`, `ReadWithCheck`, `ReadAhead`, `ReadTwice`, `Unget`, `Skip`, `Fatal`. |
| **Parsing Expression** | A parser mode flag (`isParsingExpression`) that indicates the parser is inside a sub-expression (e.g. method arguments, ternary branches). Affects how end-of-expression is determined. Toggled by `StartParsingExpression` / `EndParsingExpression`. |
| **Intern** | A string-interning function used by the Lexer to deduplicate identifier and operator strings, reducing memory allocation during tokenization. |

---

## Evaluation

| Term | Definition |
|------|-----------|
| **Evaluator** | The main evaluation engine. Its `Eval` method reads the next token from the Parser, dispatches to the appropriate `DynamicEvaluator` (Def, Class, Module, Do, Bind, IfUnless, etc.), and recursively evaluates sub-expressions. Also provides `EvalToZeroPower`, `EvalToTargetToken`, and `ContinuousEval` for controlled evaluation boundaries. |
| **DynamicEvaluator** | An interface with a single `Evaluation(e, p, ctx, t)` method. Each Ruby language construct (def, class, module, do/block, if/unless, case, while, bind, etc.) implements this interface as a separate struct. |
| **Eval (interface)** | An interface with a single `Eval(p, ctx, t)` method, implemented by `Evaluator`. Passed to `MethodEvaluator` and other components that need to trigger recursive evaluation without a direct `Evaluator` dependency. |
| **Def** | Evaluator for Ruby `def` method definitions. Handles method name extraction, static method detection (`self.method`), argument variable creation, default argument binding, keyword argument binding, body evaluation, return type inference, and method type registration in TFrame and TSignatures. Supports endless method definitions (`def foo = expr`). |
| **Class** | Evaluator for Ruby `class` definitions. Computes the next frame, handles inheritance (`<` superclass), processes `attr_reader`/`attr_accessor`/`include`/`private`/`protected`/`public` declarations, evaluates the class body, and registers a `new` class method that returns an instance of the class. |
| **Module** | Evaluator for Ruby `module` definitions. Similar to Class but without inheritance handling. Processes visibility declarations and body evaluation. |
| **Do** | Evaluator for Ruby blocks (`do...end` and `{...}`). Prepares block scope by snapshotting TFrame (`DeepCopyTFrame`), collects block variables (`|x, y|`), infers block parameter types from the calling method's signature, evaluates the block body, and restores the outer scope afterward via a restore function. |
| **Bind** | Evaluator for assignment operations (`=`). Handles scalar assignment (single variable), multiple assignment (`a, b = ...`), and destructuring. Manages read-only enforcement, default value tracking, and `@`-prefixed instance variable detection. |
| **IfUnless** | Evaluator for `if`/`unless`/`elsif`/`else` branches. Performs **type narrowing**: when a condition is `x.is_a?(Class)` or `x.nil?`, narrows the variable's type within the branch. Tracks `originalTs`, `narrowTs`, and `ifNarrowTs` to restore types after branches. Unifies branch result types with `MakeUnifiedT`. |
| **Case** | Evaluator for `case`/`when`/`in` pattern matching expressions. |
| **In** | Evaluator for pattern matching `in` clauses. Supports hash patterns, array patterns, parenthesized patterns, class patterns (deconstruct), range patterns, and pin operators. |
| **MethodEvaluator** | Dispatches method call evaluation. Takes the object type, method name, and parenthesization state, selects a `MethodEvaluateStrategy`, collects evaluated arguments, and delegates to the strategy. Handles error resolution for undefined methods. |
| **MethodEvaluateStrategy** | An interface with an `evaluate()` method. Each method call context has a corresponding strategy: `instanceMethodStrategy`, `classMethodStrategy`, `topLevelMethodStrategy`, `unionInstanceStrategy`, plus specialized strategies for array methods, hash methods, kernel methods (`yield`, `puts`/`p`), and object methods (`include`, `attr_reader`, `attr_accessor`, `class`, `raise`). |
| **Type Narrowing** | The process in `IfUnless` where a union type is refined within a conditional branch. For example, if `x` is `String | nil` and the condition is `if x`, then `x` is narrowed to `String` in the if-branch and `nil` in the else-branch. Narrowed types are stored in `narrowTs`/`ifNarrowTs` and applied via the `narrowing` method. |

---

## Built-in Type Configuration

| Term | Definition |
|------|-----------|
| **TiClassConfig** | A JSON-deserialized configuration for a single Ruby class. Fields: `Frame`, `Class`, `Extends` (parent classes), `Constants`, `ClassMethods`, `InstanceMethods`. Loaded from `.ti-config/*.json` files by `json_loader.go`. |
| **TiMethod** | A method definition within a `TiClassConfig`. Fields: `Name`, `Arguments` (list of `TiArgument`), `ReturnType` (`TiReturnType`), `Document` (shown in LSP hover). |
| **TiArgument** | A method parameter in a TiMethod. `Type` is a list of type strings (multiple means union); `Key` is the keyword argument name (empty for positional args). Special type strings: `"Self"` (returns the receiver), `"Untyped"` (any type), `"Unify"` (unified collection element type). |
| **TiReturnType** | The return type of a TiMethod. `Type` is a list of type strings (multiple means union). |
| **TiConstantType** | A constant definition within a TiClassConfig, representing class-level constants. |
| **TiLoader** | Configuration loaded from `.ti-loader.json`. Contains a `Preload` list of file paths to analyze before the target file. |
| **defineBuiltinMethod** | The processor that iterates over loaded `TiClassConfig` entries and registers each method's type in TFrame and TSignatures. Handles static methods separately via `defineBuiltinStaticMethod`. Converts JSON type strings to `T` objects via `ConvertToBuiltinT`. |

---

## LSP and Output

| Term | Definition |
|------|-----------|
| **ExecuteFlags** | Parsed command-line flags that determine the output mode: `IsDefineInfo`, `IsSuggest`, `IsHover`, `IsExtends`, `IsVersion`, `IsLlmInfo`, `IsLlmError`, `IsLlmDefine`, `IsLlmClass`, etc. Built by `BuildFlags()` from `os.Args`. |
| **DefineInfoArticle** | A snapshot of a method definition for LSP define-info output. Stores the Parser state, Context, method type (`MethodT`), and source row at the point of definition. |
| **SpecialCodeComment** | A comment annotation (`# @type: ...`) found in source code. Stores `FileName`, `Row`, and `Document`. Used by the LSP to provide additional hover information at specific lines. |
| **PrintSuggestionsForLsp** | Output function that emits method completion suggestions. Filters TSignatures by the target class and formats them for LSP auto-complete consumption. |
| **PrintHover** | Output function that produces hover information for a symbol at the LSP target row. Looks up matching signatures and special comments to compose the hover content. |
| **PrintErrors** | Output function that formats type-checking errors collected in `Parser.Errors` for display, including file name, row number, and error message. |
| **LspTargetRow** | A field on `Parser` set from `--suggest` / `--hover` / `--define-info` flags. Indicates the source line the LSP cursor is on, so that output functions can filter results to that location. |

---

## Scope and Visibility

| Term | Definition |
|------|-----------|
| **Private** | A visibility state tracked by `Context.IsPrivate`. When active, method definitions are stored with `isPrivate=true` in their FrameKey. Private methods are excluded from external method resolution in `GetMethodT`. Toggled by `StartPrivate`/`EndPrivate`. |
| **Protected** | A visibility state tracked by `Context.IsProtected`. Similar to private but with different access rules. Toggled by `StartProtected`/`EndProtected`. |
| **Static** | A flag (`IsStatic`, `IsDefineStatic`) indicating a class-level (as opposed to instance-level) method or variable. Static methods are defined with `self.method_name` syntax and stored/retrieved via `SetClassMethodT`/`GetClassMethodT`. |
| **Block Scope** | The lexical scope created when entering a `do...end` or `{...}` block. `Do.prepareBlockScope` snapshots the current TFrame, sets block parameter types, and returns a restore function that reinstates the outer scope's variable types after the block exits. |
| **RestoreVariable** | A record of `{id, t}` representing a variable's type before entering a block scope. Used by the restore function to undo block-local variable bindings. |

---

## Expression and Assignment Concepts

| Term | Definition |
|------|-----------|
| **Comma** | Evaluator for comma-separated expressions. Produces array types by collecting each comma-separated value as an array variant. |
| **SquareBracket** | Evaluator for `[]` reference operations. Dispatches to type-specific handlers: `arrayReferenceEvaluation`, `hashReferenceEvaluation`, `stringReferenceEvaluation`, `integerReferenceEvaluation`. Also handles array literal construction via `makeArray`. |
| **Range** | Evaluator for Ruby range expressions (`..` and `...`). |
| **Ternary** | Evaluator for ternary expressions (`condition ? a : b`). Unifies the types of both branches. |
| **Logical** | Evaluator for logical operators (`&&`, `\|\|`, `and`, `or`). |
| **Exclamation** | Evaluator for the `!` (not) prefix operator. |
| **Return** | Evaluator for explicit `return` statements. Appends the return value to the parser's `lastReturnT` list for method return type inference. |
| **Rescue** | Evaluator for `rescue` exception handling blocks. |
| **While** | Evaluator for `while`/`until` loop constructs. |
| **NameSpace** | Evaluator for `::` namespace resolution (e.g. `Foo::Bar`). Resolves nested class/module references by computing the frame with `CalculateFrame`. |
| **Self** | Evaluator for the `self` keyword. Returns the current class as an object type. |
| **Hash** | Evaluator for hash literal construction (`{key: value, ...}`). Builds hash types with key-value variant pairs. |
| **OpenParentheses** | Evaluator for grouped expressions in parentheses. |
| **Begin** | Evaluator for `begin...end` blocks. |
