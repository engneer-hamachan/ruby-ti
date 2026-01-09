# Ruby-ti Design Philosophy

> **Types live in your heart** 💎

---

## 1. What is a Type?

Consider this simple Ruby code:

```ruby
array.max
```

In most type inference systems, the `max` method returns an `Optional` type (either `NilClass` or a collection item). This means the following code would raise a type error:

```ruby
array.max + 1
```

### The Traditional Solution

To avoid type errors, you'd need to write:

```ruby
int = array.max

if int.is_a?(Integer)
  int + 1
end
```

### The Question

But is this type inference system really user-friendly?

Consider this code:

```ruby
array << 1
array.max + 1
```

**Do you really want to add type checking to this code?**

---

## 2. In Ruby-ti, **You** Decide the Types

There's no single correct answer to the cases presented above. Whether you should add type checking depends on the context—it's **case by case**.

**Ruby-ti's core belief:** The person who decides types should be **the Rubyist, not the system**.

### Example: Two Different Mental Models

#### Config for Rubyists who believe "max always returns an element"

`.ti-config/array.json`:
```json
{
  "name": "max",
  "arguments": [
    {
      "type": ["DefaultInt"]
    }
  ],
  "block_parameters": ["Unify"],
  "return_type": {
    "type": ["Unify"]
  }
}
```

#### Config for Rubyists who believe "max may not always return an element"

`.ti-config/array.json`:
```json
{
  "name": "max",
  "arguments": [
    {
      "type": ["DefaultInt"]
    }
  ],
  "block_parameters": ["Unify"],
  "return_type": {
    "type": ["OptionalUnify"]
  }
}
```

---

## 3. Types Live in the Rubyist's Heart 💖

In ruby-ti, **the Rubyist decides types, not the system**. More specifically, it's **you**—the person working with the code right now.

### Our Philosophy

- 🎯 **Don't be bound by logical correctness**
- ✍️ **Rewrite type configs to match what feels right to you**
- 💡 **ruby-ti reflects your mental model and provides the best support possible**

### Our Goal

**Keep Ruby a language that's always fun to write** ✨

---

*ruby-ti is a type inference tool that respects your intuition and mental model, not a rigid system that enforces one "correct" way of thinking about types.*
