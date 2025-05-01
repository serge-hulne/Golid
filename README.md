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
package main

import (
	c "app/counter"
	s "app/golid"
)

func main() {
	c1 := c.New()
	c1.Mount(s.Body())

	select {}
}
```

