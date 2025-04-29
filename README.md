# Golid

**Golid** is a simple, solid, Go-native frontend framework for WebAssembly applications.  
It focuses on clarity, modularity, and reactivity — without the complexity of heavy Virtual DOM systems.

## Features

- ✨ Simple, modular Go components
- ⚡ Fine-grained reactivity with Signals
- 🚀 WebAssembly-first performance
- 🛠️ Minimal runtime, maximum transparency
- 🌐 Built entirely in Go

## Quick Example

```go
counter := counter.New()
counter.Mount(solidgo.Body())
