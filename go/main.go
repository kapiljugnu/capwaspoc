package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"strings"
	"syscall/js"
)

func main() {

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := args[0].String()
		path = strings.ToLower(path)

		menu := templates.AppMenus[path]
		component := templates.Layout(menu)

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
