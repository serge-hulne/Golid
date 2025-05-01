package golid

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/google/uuid"

	"maragu.dev/gomponents"
)

type JsCallback func(this js.Value, args []js.Value) interface{}

var doc = js.Global().Get("document")

// --- Signal System ---
type Signal[T any] struct {
	value    T
	watchers []func(T)
}

func NewSignal[T any](initial T) *Signal[T] {
	return &Signal[T]{value: initial}
}

func (s *Signal[T]) Set(val T) {
	s.value = val
	for _, w := range s.watchers {
		w(val)
	}
}

func (s *Signal[T]) Get() T {
	return s.value
}

func (s *Signal[T]) Watch(f func(T)) {
	s.watchers = append(s.watchers, f)
}

// --- DOM Utilities ---
func ID() string {
	return uuid.NewString()
}

func Append(html string, Element js.Value) {
	Element.Call("insertAdjacentHTML", "beforeend", html)
}

func NodeFromID(id string) js.Value {
	return doc.Call("getElementById", id)
}

func On(eventType, id string, callback JsCallback) {
	NodeFromID(id).Call("addEventListener", eventType, js.FuncOf(callback))
}

func Body() js.Value {
	return doc.Get("body")
}

// --- Logging ---
func Log(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}

func Logf(format string, args ...interface{}) {
	js.Global().Get("console").Call("log", fmt.Sprintf(format, args...))
}

// --- Render HTMLGo component to string ---

func RenderHTML(n gomponents.Node) string {
	var b strings.Builder
	err := n.Render(&b)
	if err != nil {
		return "<div>render error</div>"
	}
	return b.String()
}

func Callback(f func()) JsCallback {
	return func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
}
