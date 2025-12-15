# Ruby-TI

<p align="center">
  <img src="logo.svg" alt="Ruby-TI Logo" width="400"/>
</p>

<p align="center">
  <strong>Type Inference for Ruby without Type Annotations</strong>
</p>

<p align="center">
  <a href="#quick-start">Quick Start</a> •
  <a href="#project-setup">Project Setup</a> •
  <a href="#supported-classes">Supported Classes</a> •
  <a href="#documentation">Documentation</a>
</p>

<p align="center">
  <img alt="Go Version" src="https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=flat&logo=go">
  <img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg">
  <img alt="mruby" src="https://img.shields.io/badge/mruby-compatible-red.svg">
</p>

---

## 💡 What is Ruby-TI?

Ruby-TI is a **static type analyzer for mruby** written in Go. It performs type inference and type checking on your Ruby code **without requiring any type annotations**. Just write Ruby as you always do, and Ruby-TI will analyze your code in real-time through your editor.

Experience the joy of Ruby's dynamic typing with the safety of static type checking!

<p align="center">
  <img src="image/sample2.png" alt="Ruby-TI in action" width="700"/>
</p>

### ✨ Why Ruby-TI?

I love Ruby. You write code, you run it immediately—no complex type systems to learn. Whether it's a morning inspiration or a quick experiment during a break, Ruby has always been there.

Ruby-TI makes these moments even better. **Write Ruby freely, get type safety automatically.** It's a fresh and fun development experience.

---

## 🚀 Quick Start

### Prerequisites

- Go 1.24.5 or higher

```bash
go version
```

### Installation

```bash
# Install Ruby-TI
git clone https://github.com/engneer-hamachan/ruby-ti.git
cd ruby-ti
make install
export PATH="$PATH:$(pwd)/bin"

# Install LSP Server
git clone https://github.com/engneer-hamachan/ruby-ti-lsp.git
cd ruby-ti-lsp
make install
```

---

## 📁 Project Setup

To use Ruby-TI in your Ruby project:

```bash
# Navigate to your Ruby project
cd your-ruby-project

# Copy the .ti-config directory
cp -r /path/to/ruby-ti/.ti-config .

# (Optional) Customize for your mruby environment
# Edit .ti-config/*.json files
```

The `.ti-config` directory contains type definitions for Ruby built-in classes. You can customize these to match your specific mruby environment.

---

## 📚 Supported Classes

Ruby-TI comes with built-in support for the following classes and modules (gradually expanding PicoRuby support):

### Core Classes
- 🔢 **Integer** - Integer arithmetic
- 🔢 **Float** - Floating-point arithmetic
- 📝 **String** - String manipulation
- 🔣 **Symbol** - Symbols
- ✅ **Bool** - Boolean values (true/false)
- ❌ **Nil** - Nil type

### Collections
- 📦 **Array** - Array operations
- 🗂️ **Hash** - Hash operations
- 📏 **Range** - Range objects

### Advanced Types
- 🔧 **Object** - Base class
- 🏗️ **Class** - Class objects
- 🎯 **Proc** - Proc/Lambda objects
- 🔄 **Enumerable** - Enumerable module
- 🧮 **Math** - Mathematical functions
- 🌐 **Kernel** - Kernel module

### PicoRuby Support
- ⚡ **GPIO** - GPIO control for PicoRuby
- 🌐 **JS** / **JSObject** - JavaScript interop for PicoRuby

> 💡 **Tip**: You can add custom classes by creating JSON files in `.ti-config/`. See [Builtin JSON Guide](./docs/builtin-json.md) for details.

---

## 📖 Documentation

### 📘 User Guides
- **[Builtin JSON Guide](./docs/builtin-json.md)** - How to add and customize type definitions
- **[Editor Setup Guide](https://github.com/engneer-hamachan/ruby-ti-lsp)** - VSCode/Vim/Neovim configuration

### 🔧 Developer Resources
- **[Developer Documentation](./CLAUDE.md)** - Architecture, development workflow, and contribution guidelines

### 🎬 Getting Started with Your Editor

Ruby-TI works best with LSP integration in your editor:

- **VSCode**: Full support with Code Lens, diagnostics, and auto-completion
- **Vim/Neovim**: Native LSP support with type annotations
- **Other LSP-compatible editors**: Should work out of the box

See the [ruby-ti-lsp repository](https://github.com/engneer-hamachan/ruby-ti-lsp) for detailed setup instructions.

---

## 🤝 Contributing

We welcome bug reports, feature requests, and pull requests! Feel free to open an issue or submit a PR.

### How to Contribute
1. 🐛 **Report bugs** - Found an issue? Let us know!
2. 💡 **Suggest features** - Have an idea? We'd love to hear it!
3. 🔧 **Submit PRs** - Contributions are always welcome!

---

## 📄 License

This project is licensed under the MIT License.

---

<p align="center">
  Made with ❤️ for the Ruby community
</p>

<p align="center">
  <a href="https://github.com/engneer-hamachan/ruby-ti">⭐ Star us on GitHub</a>
</p>
