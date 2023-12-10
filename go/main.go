package main

import (
	"context"
	"strings"
	"syscall/js"

	templates "boozedog/capwaspoc/templ"
)

// "context"
// "os"

func main() {

	c := make(chan struct{}, 0)
	b := new(strings.Builder)

	// Define a simple component
	helloWorld := templates.Hello("abc")
	helloWorld.Render(context.Background(), b)

	// Render the component
	js.Global().Get("document").Call("getElementById", "app").Set("innerHTML", b.String())

	<-c
}
