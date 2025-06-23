package main

import (
	"app/golid"
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func main() {

	// ---Uncomment one of the following examples to run it ---

	app := CounterComponent()
	//app := List1()
	//app := List2()
	// app := TextCopyDemo()
	// app := TextCopyDemo1()
	//app := TextCopyDemo2()
	golid.Render(app)
	golid.Run()
}

func CounterComponent() Node {
	// Observable (represents the state of the app)
	count := golid.NewSignal(0)

	return Div(
		Style("border: 1px solid orange; padding: 10px; margin: 10px;"),

		// Bind text Element to the reactive count signal (observable)
		golid.Bind(func() Node {
			return Div(Text(fmt.Sprintf("Count = %d", count.Get())))
		}),

		// [+] Button element
		Button(
			Style("margin: 3px;"),
			Text("+"),
			golid.OnClick(func() {
				count.Set(count.Get() + 1)
			}),
		),

		// [-] Button element
		Button(
			Style("margin: 3px;"),
			Text("-"),
			golid.OnClick(func() {
				count.Set(count.Get() - 1)
			}),
		),
	)
}

func List1() Node {
	messages := []string{"Hello", "World", "Golid"}

	return Div(
		H3(Text("Messages")),
		golid.ForEach(messages, func(msg string) Node {
			return Li(Text(msg))
		}),
	)
}

func List2() Node {
	messages := []string{"Hello", "World", "Golid"}

	return Div(
		H3(Text("Messages")),
		Ul( // Wrap list items in a <ul>
			golid.ForEach(messages, func(msg string) Node {
				return Li(
					Text(msg),
					golid.OnClick(func() {
						golid.Log("Clicked on:", msg)
					}),
				)
			}),
		),
	)
}

func TextCopyDemo() Node {
	// Internal state
	inputValue := golid.NewSignal("")
	copiedText := golid.NewSignal("")

	return Div(
		H3(Text("Text Copier")),

		// Input field with OnInput binding
		Input(
			Type("text"),
			Placeholder("Type something..."),
			golid.OnInput(func(val string) {
				inputValue.Set(val)
			}),
		),

		// Button that copies current input value to output signal
		Button(
			Style("margin-left: 10px;"),
			Text("Copy"),
			golid.OnClick(func() {
				copiedText.Set(inputValue.Get())
			}),
		),

		// Output label
		Div(
			Style("margin-top: 10px; font-weight: bold;"),
			Text("Copied text: "),
			golid.Bind(func() Node {
				return Text(copiedText.Get())
			}),
		),
	)
}

func TextCopyDemo1() Node {
	inputValue := golid.NewSignal("")

	return Div(
		H3(Text("Live Mirror")),

		// Input field
		Input(
			Type("text"),
			Placeholder("Start typing..."),
			golid.OnInput(func(val string) {
				inputValue.Set(val)
			}),
		),

		// Live-updating label
		Div(
			Style("margin-top: 10px; font-weight: bold;"),
			Text("You typed: "),
			golid.Bind(func() Node {
				return Text(inputValue.Get())
			}),
		),
	)
}

func TextCopyDemo2() Node {
	shared := golid.NewSignal("")

	return Div(
		H3(Text("Twin Mirror Inputs")),

		// Input 1 (reactively updates on shared signal change)
		golid.Bind(func() Node {
			return Input(
				Type("text"),
				Placeholder("Input 1"),
				Style("margin-left: 10px;"),
				Value(shared.Get()),
				golid.OnInput(func(val string) {
					if val != shared.Get() {
						shared.Set(val)
					}
				}),
			)
		}),

		// Input 2 (also reacts to signal change)
		golid.Bind(func() Node {
			return Input(
				Type("text"),
				Placeholder("Input 2"),
				Style("margin-left: 10px;"),
				Value(shared.Get()),
				golid.OnInput(func(val string) {
					if val != shared.Get() {
						shared.Set(val)
					}
				}),
			)
		}),

		// Display current value
		Div(
			Style("margin-top: 10px; font-style: italic;"),
			Text("Shared value: "),
			golid.Bind(func() Node {
				return Text(shared.Get())
			}),
		),
	)
}
