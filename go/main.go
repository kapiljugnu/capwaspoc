package main

import (
	"syscall/js"

	"github.com/a-h/templ"
)

func main() {
	c := make(chan struct{}, 0)

	// Define a simple component
	helloWorld := templ.ComponentFunc(func() templ.Node {
		return templ.Element("h1", nil, templ.Text("Hello, World!"))
	})

	// Render the component
	js.Global().Get("document").Call("getElementById", "app").Set("innerHTML", templ.RenderToString(helloWorld()))

	<-c
}
