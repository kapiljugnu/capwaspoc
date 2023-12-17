package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
)

// "context"
// "os"

func main() {

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := args[0].String()
		var component templ.Component

		switch path {
		case "home":
			component = templates.Hello("Home")
		case "about":
			component = templates.Hello("About")
		}

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
