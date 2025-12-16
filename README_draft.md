# Ruby-TI

<p align="center">
  <strong>Type Inference for MRuby without Type Annotations</strong>
</p>

<p align="center">
  <a href="#features-in-action">Features</a> •
  <a href="#quick-start">Quick Start</a> •
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

Ruby-TI is a **static type analyzer for MRuby** written in Go. It performs type inference and type checking on your Ruby code **without requiring any type annotations**. Just write Ruby as you always do, and Ruby-TI will analyze your code in real-time through your editor.

Experience the joy of Ruby's dynamic typing with the safety of static type checking!

<p align="center">
  <img src="image/ruby-ti-demo.gif" alt="Ruby-TI Demo" width="700"/>
</p>

---

## 🚀 Quick Start

### Prerequisites

- Go 1.24.5 or higher

```bash
go version
```

### Installation

#### 1. Install Ruby-TI

```bash
cd /path/to/your_directory
git clone https://github.com/engneer-hamachan/ruby-ti.git
cd ruby-ti
make install

# Add to your shell profile (e.g., ~/.bashrc, ~/.zshrc, ~/.bash_profile)
echo 'export PATH="$PATH:'$(pwd)'/bin"' >> ~/.bash_profile
source ~/.bash_profile
```

#### 2. Install LSP Server

```bash
cd /path/to/your_directory
git clone https://github.com/engneer-hamachan/ruby-ti-lsp.git
cd ruby-ti-lsp
make install

# Add to your shell profile (e.g., ~/.bashrc, ~/.zshrc, ~/.bash_profile)
echo 'export PATH="$PATH:'$(pwd)'/bin"' >> ~/.bash_profile
source ~/.bash_profile
```

#### 3. Configure Your Editor

##### Neovim

Use your preferred LSP plugin. Example configuration for coc.nvim (coc-settings.json):

```json
{
  "hover.target": "float",
  "codeLens.enable": true,
  "codeLens.separator": " #",
  "codeLens.position": "eol",
  "diagnostic.virtualText": true,
  "diagnostic.virtualTextCurrentLineOnly": false,
  "diagnostic.enableMessage": "never",
  "workspace.openResourceCommand": "edit",
  "suggest.floatConfig": {
    "border": true,
    "rounded": true
  },
  "hover.floatConfig": {
    "border": true,
    "rounded": true
  },
  "signature.floatConfig": {
    "border": true,
    "rounded": true
  },
  "diagnostic.floatConfig": {
    "border": true,
    "rounded": true
  },
  "languageserver": {
    "ruby-ti": {
      "command": "ti-lsp",
      "filetypes": ["ruby", "json"]
    }
  }
}
```

##### VSCode

Install the VSCode extension:

```bash
code --install-extension /path/to/ruby-ti-lsp/vscode/ruby-ti-lsp-0.1.0.vsix
```

#### 4. Project Setup

To use Ruby-TI in your Ruby project:

```bash
cd your-ruby-project
cp -r /path/to/ruby-ti/.ti-config .
```

The `.ti-config` directory contains type definitions for Ruby built-in classes. You can customize these to match your specific mruby environment by editing `.ti-config/*.json` files.

---

## ✨ Features in Action

### Diagnostics
Real-time type error detection

<p align="center">
  <img src="image/diagnostic.png" alt="Type diagnostics" width="700"/>
</p>

### Hover Information
Inspect types on hover

<p align="center">
  <img src="image/hover.png" alt="Hover type information" width="700"/>
</p>

### Auto-completion
Intelligent code suggestions

<p align="center">
  <img src="image/suggest.png" alt="Auto-completion" width="700"/>
</p>

### Code Actions
Quick fixes and refactoring

<p align="center">
  <img src="image/codeaction.png" alt="Code actions" width="700"/>
</p>

---

## 📚 Supported Classes

Ruby-TI is gradually expanding support centered around PicoRuby:

- **Array** - Array operations
- **Bool** - Boolean values (true/false)
- **Class** - Class objects
- **Enumerable** - Enumerable module
- **Float** - Floating-point arithmetic
- **GPIO** - GPIO control for PicoRuby
- **Hash** - Hash operations
- **Integer** - Integer arithmetic
- **Kernel** - Kernel module
- **Math** - Mathematical functions
- **Nil** - Nil type
- **Object** - Base class
- **Proc** - Proc/Lambda objects
- **Range** - Range objects
- **String** - String manipulation
- **Symbol** - Symbols

### Customize for Your Needs

The classes listed above are just the default configuration. **Ruby-TI aims to express the types that live in every Rubyist's heart.**

We encourage you to customize `.ti-config` to define types that make sense to you and create documentation that fits your workflow. Make Ruby-TI truly yours!

---

## 📖 Documentation
### 📘 User Guides
- **[.ti-config Configuration Guide](./docs/ti-config.md)** - How to customize type definitions for your mruby environment
---

## 🤝 Contributing

**We especially welcome issues!** While pull requests might be challenging at this stage due to the project's active development, we'd love to hear about bugs, feature requests, and any feedback you have.

Ruby-TI is being actively and rapidly improved by the author. Your issues help guide development and make Ruby-TI better for everyone!
---

## 📄 License

This project is licensed under the MIT License.

---
<p align="center">
  <a href="https://github.com/engneer-hamachan/ruby-ti">⭐ Star us on GitHub</a>
</p>
