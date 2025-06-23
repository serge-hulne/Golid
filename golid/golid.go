// golid.go
// A minimal reactive UI toolkit for Go+WASM using gomponents

package golid

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/google/uuid"

	"maragu.dev/gomponents"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ---------------------------
// 🔧 Core Types & Interfaces
// ---------------------------

type JsCallback func(this js.Value, args []js.Value) interface{}

var doc = js.Global().Get("document")

// ------------------------------------
// 📦 Reactive Signals (State Handling)
// ------------------------------------

type effect struct {
	fn   func()
	deps map[any]struct{}
}

var currentEffect *effect

type hasWatchers interface {
	removeWatcher(*effect)
}

type Signal[T any] struct {
	value    T
	watchers map[*effect]struct{}
}

func NewSignal[T any](initial T) *Signal[T] {
	return &Signal[T]{
		value:    initial,
		watchers: make(map[*effect]struct{}),
	}
}

func (s *Signal[T]) Get() T {
	if currentEffect != nil {
		s.watchers[currentEffect] = struct{}{}
		currentEffect.deps[s] = struct{}{}
	}
	return s.value
}

func (s *Signal[T]) Set(val T) {
	s.value = val
	for e := range s.watchers {
		go runEffect(e)
	}
}

func (s *Signal[T]) removeWatcher(e *effect) {
	delete(s.watchers, e)
}

func Watch(fn func()) {
	e := &effect{
		fn:   fn,
		deps: make(map[any]struct{}),
	}
	runEffect(e)
}

func runEffect(e *effect) {
	for dep := range e.deps {
		if s, ok := dep.(hasWatchers); ok {
			s.removeWatcher(e)
		}
	}
	e.deps = make(map[any]struct{})
	currentEffect = e
	e.fn()
	currentEffect = nil
}

// ------------------------
// 🖼  Reactive DOM Binding
// ------------------------

func Bind(fn func() Node) Node {
	id := GenID()
	placeholder := Span(Attr("id", id))

	var check js.Func
	check = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		elem := NodeFromID(id)
		if elem.Truthy() {
			Watch(func() {
				html := RenderHTML(Div(Attr("id", id), fn()))
				elem := NodeFromID(id)
				if elem.Truthy() {
					elem.Set("outerHTML", html)
				}
			})
			return nil
		}
		js.Global().Call("setTimeout", check, 10)
		return nil
	})
	js.Global().Call("setTimeout", check, 10)

	return placeholder
}

func BindText(fn func() string) Node {
	id := GenID()
	span := Span(Attr("id", id), Text(fn()))

	Watch(func() {
		elem := NodeFromID(id)
		if elem.Truthy() {
			elem.Set("textContent", fn())
		}
	})

	return span
}

// -------------------------
// 🧩 Event Binding Helpers
// -------------------------

func OnClick(f func()) Node {
	id := GenID()
	go func() {
		js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			elem := NodeFromID(id)
			if elem.Truthy() {
				elem.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
					f()
					return nil
				}))
			}
			return nil
		}), 0)
	}()

	return Attr("id", id)
}

// ------------------
// 🧱 DOM Utilities
// ------------------

func GenID() string {
	return uuid.NewString()
}

func Append(html string, Element js.Value) {
	Element.Call("insertAdjacentHTML", "beforeend", html)
}

func NodeFromID(id string) js.Value {
	return doc.Call("getElementById", id)
}

func BodyElement() js.Value {
	return doc.Get("body")
}

// ----------------------
// 🧪 Rendering Utilities
// ----------------------

func RenderHTML(n gomponents.Node) string {
	var b strings.Builder
	err := n.Render(&b)
	if err != nil {
		return "<div>render error</div>"
	}
	return b.String()
}

func Render(n Node) {
	Append(RenderHTML(n), BodyElement())
}

// ------------------
// 🛠 Callback Helper
// ------------------

func Callback(f func()) JsCallback {
	return func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
}

// --------------
// 🧭 App Entrypoint
// --------------

func Run() {
	select {}
}

// ------------------
// 🪵 Debugging
// ------------------

func Log(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}

func Logf(format string, args ...interface{}) {
	js.Global().Get("console").Call("log", fmt.Sprintf(format, args...))
}

// ------------------
// Lists (Foreach())
// -------------------

func ForEach[T any](items []T, render func(T) Node) Node {
	var children []Node
	for _, item := range items {
		children = append(children, render(item))
	}
	return Group(children)
}

func ForEachSignal[T any](sig *Signal[[]T], render func(T) Node) Node {
	return Bind(func() Node {
		items := sig.Get()
		var children []Node
		for _, item := range items {
			children = append(children, render(item))
		}
		return Group(children)
	})
}

// -----------
// text inputs
// -----------

func OnInput(handler func(string)) Node {
	id := GenID()
	go func() {
		js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			elem := NodeFromID(id)
			if elem.Truthy() {
				elem.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
					value := this.Get("value").String()
					handler(value)
					return nil
				}))
			}
			return nil
		}), 0)
	}()
	return Attr("id", id)
}

func BindInput(sig *Signal[string], placeholder string) Node {
	id := GenID()

	// Add listener for user input
	go func() {
		js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			elem := NodeFromID(id)
			if elem.Truthy() {
				elem.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
					val := elem.Get("value").String()
					if val != sig.Get() {
						sig.Set(val)
					}
					return nil
				}))
			}
			return nil
		}), 0)
	}()

	// When the signal changes, update DOM input's `.value`
	Watch(func() {
		elem := NodeFromID(id)
		if elem.Truthy() {
			signalVal := sig.Get()
			elemVal := elem.Get("value").String()
			if elemVal != signalVal {
				elem.Set("value", signalVal)
			}
		}
	})

	// Return static node that will stay in the DOM
	return Input(
		Attr("id", id),
		Type("text"),
		Placeholder(placeholder),
		Value(sig.Get()), // initial value
	)
}
