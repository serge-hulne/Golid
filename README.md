# Golid

**Golid** is a simple, solid, Go-native frontend framework for WebAssembly applications.  
It focuses on clarity, modularity, and reactivity — without the complexity of heavy Virtual DOM systems.

A minimal, Go-native frontend framework with signals, components, and WebAssembly — no Node.js, no npm, no JSX, no bundlers.


## ✨ What is Golid?

**Golid** (short for Go + Solid) is a lightweight frontend framework written entirely in Go, compiled to WebAssembly. It’s inspired by frameworks like Solid.js, but built for Go developers who want simplicity, control, and zero JS toolchain pain.

With Golid, you can build reactive web apps using:
- ✅ Pure Go
- ✅ Signals and components
- ✅ Tiny `.wasm` bundles (TinyGo optional)
- ✅ No Node.js, no npm, no React, no JSX, no bundlers

---

## 🚀 Quick Start

1. Clone the repo:
   ```bash
   git clone https://github.com/serge-hulne/Golid.git
   cd Golid
   ```

2. Copy the Go WASM runtime:
    ```
    go env GOROOT
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
	```

3. Build the webassembly binary:
    ```
	make wasm
	```

4.  Start a local server:
    ```
	python3 -m http.server 8090
	```

5. Watch the app in the browse:
    ```
	http://localhost:8090
	```

## 💡 Example: Counter Component

```
type Counter struct {
    DivID        string
    Counter      *golid.Signal[int]
    PlusButtonID string
    MinButtonID  string
}

func New() *Counter {
    return &Counter{
        DivID:        golid.ID(),
        Counter:      golid.NewSignal(0),
        PlusButtonID: golid.ID(),
        MinButtonID:  golid.ID(),
    }
}

func (c *Counter) Render() Node {
    return Div(
        Div(ID(c.DivID), Text(fmt.Sprintf("Count = %d", c.Counter.Get()))),
        Button(ID(c.PlusButtonID), Text("+")),
        Button(ID(c.MinButtonID), Text("-")),
    )
}
    
```


## ❌ What Golid Does Not Require

- No Node.js
- No npm or yarn
- No Parcel, Webpack, Vite, or other bundlers
- No React, Vue, Svelte, Solid.js, or JSX
- No go:generate or code generation

- Just:
✅ Go

## 🛣 Roadmap

- [] Add routing system
- [] Add built-in UI components (e.g., Toggle, Input, Form)
- [] Provide example apps and templates
- []  Optional CSS helper system


## 📜 License

Golid is open source under the GNU General Public License v3.

