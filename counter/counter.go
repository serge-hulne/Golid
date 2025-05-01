package counter

import (
	"fmt"
	"syscall/js"

	"app/golid"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		Style("border: 1px solid grey; padding: 10px; margin: 10px ;"),
		Div(
			ID(c.DivID),
			Text(fmt.Sprintf("Count = %d", c.Counter.Get())),
		),
		Button(
			ID(c.PlusButtonID),
			Text("+"),
		),
		Button(
			ID(c.MinButtonID),
			Text("-"),
		),
	)
}

func (c *Counter) Mount(target js.Value) {
	golid.Append(golid.RenderHTML(c.Render()), target)

	c.Counter.Watch(func(val int) {
		golid.NodeFromID(c.DivID).Set("innerHTML", fmt.Sprintf("Count = %d", val))
	})

	golid.On("click", c.PlusButtonID, golid.Callback(func() {
		c.Counter.Set(c.Counter.Get() + 1)
	}))

	golid.On("click", c.MinButtonID, golid.Callback(func() {
		c.Counter.Set(c.Counter.Get() - 1)
	}))
}
