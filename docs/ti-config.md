# .ti-config Configuration Guide

The `.ti-config` directory contains JSON configuration files that define type signatures for Ruby built-in classes. This system allows you to customize Ruby-TI for your specific mruby environment without modifying the Go source code.

## Setup

1. Copy the `.ti-config` directory to your project root:
   ```bash
   cd your-ruby-project
   cp -r /path/to/ruby-ti/.ti-config .
   ```

2. Edit or add JSON files in `.ti-config/` to match your mruby environment

## Basic Structure

Each JSON file defines a Ruby class with its methods and constants:

```json
{
  "frame": "Builtin",
  "class": "ClassName",
  "extends": ["ParentClass"],
  "constants": [],
  "instance_methods": [],
  "class_methods": []
}
```

### Fields

- **`frame`**: Frame name (usually `"Builtin"`)
- **`class`**: Ruby class name
- **`extends`**: Array of parent class names (optional)
- **`constants`**: Array of constant definitions (optional)
- **`instance_methods`**: Array of instance method definitions (optional)
- **`class_methods`**: Array of class method definitions (optional)

## Method Definition

```json
{
  "name": "method_name",
  "arguments": [
    {
      "type": ["String"],
      "key": "keyword_name",
      "is_asterisk": false
    }
  ],
  "return_type": {
    "type": ["String", "Nil"],
    "is_conditional": false,
    "is_destructive": false
  },
  "block_parameters": ["String"]
}
```

### Method Fields

- **`name`**: Method name
- **`arguments`**: Array of argument definitions (optional)
  - `type`: Array of type names (multiple types create a union type)
  - `key`: Keyword argument name (optional)
  - `is_asterisk`: Variable-length argument flag (optional)
- **`return_type`**: Return type definition
  - `type`: Array of return type names (multiple types create a union type)
  - `is_conditional`: Conditional return flag (optional)
  - `is_destructive`: Destructive operation flag (optional)
- **`block_parameters`**: Array of block parameter types (optional)

## Constant Definition

```json
{
  "name": "CONSTANT_NAME",
  "return_type": {
    "type": ["Int"]
  }
}
```

## Supported Types

### Basic Types
- `"Nil"` - Nil value
- `"Symbol"` - Symbol
- `"Bool"` - Boolean (true/false)
- `"Block"` - Block object
- `"Range"` - Range object
- `"Untyped"` - Any type (equivalent to `any`)

### String Types
- `"String"` - String
- `"DefaultString"` - Default argument string
- `"OptionalString"` - Nil or String

### Numeric Types
- `"Int"` - Integer
- `"DefaultInt"` - Default argument integer
- `"OptionalInt"` - Nil or Integer
- `"Float"` - Float
- `"DefaultFloat"` - Default argument float
- `"OptionalFloat"` - Nil or Float
- `"Number"` - Integer or Float

### Collection Types
- `"Array"` - Array
- `"Hash"` - Hash
- `"StringArray"` - Array of strings
- `"IntArray"` - Array of integers
- `"FloatArray"` - Array of floats
- `"KeyArray"` - Array of hash keys
- `"KeyValueArray"` - Array of hash values

### Advanced Types
- `"Self"` - Instance object
- `"Unify"` - Unified union type for Hash/Array/Union
- `"OptionalUnify"` - Nil or Unify
- `"BlockResultArray"` - Array of block results
- `"SelfArray"` - Convert instance to array
- `"Argument"` - Return argument as-is
- `"UnifyArgument"` - Return unified argument

## Example: GPIO Class for PicoRuby

```json
{
  "frame": "Builtin",
  "class": "GPIO",
  "instance_methods": [
    {
      "name": "high?",
      "arguments": [],
      "return_type": {"type": ["Bool"]}
    },
    {
      "name": "write",
      "arguments": [{"type": ["Int"]}],
      "return_type": {"type": ["Int"]}
    }
  ],
  "class_methods": [
    {
      "name": "new",
      "arguments": [
        {"type": ["Int", "String", "Symbol"]},
        {"type": ["Int"]},
        {"type": ["DefaultInt"]}
      ],
      "return_type": {"type": ["GPIO"]}
    },
    {
      "name": "read_at",
      "arguments": [{"type": ["Int", "String", "Symbol"]}],
      "return_type": {"type": ["Int"]}
    }
  ],
  "constants": [
    {
      "name": "OUT",
      "return_type": {"type": ["Int"]}
    }
  ]
}
```

## Tips

1. **Union Types**: Specify multiple types in the `type` array to create union types:
   ```json
   "type": ["String", "Int", "Nil"]
   ```

2. **Keyword Arguments**: Use the `key` field for keyword arguments:
   ```json
   {
     "type": ["String"],
     "key": "name"
   }
   ```

3. **Variable-Length Arguments**: Use `is_asterisk` for splat arguments:
   ```json
   {
     "type": ["String"],
     "is_asterisk": true
   }
   ```

4. **Method Chaining**: Use `"Self"` as return type for methods that return the object itself:
   ```json
   "return_type": {"type": ["Self"]}
   ```

5. **Default Arguments**: Use `"Default*"` types for optional arguments with defaults:
   ```json
   {"type": ["DefaultInt"]}
   ```
