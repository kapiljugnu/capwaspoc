package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
)

var FUEL_LOAD = ""
var TAKEOFF_MIS = "takeoff-mins"
var ALTERNATE = "alternate"

type Menu struct {
	Component templ.Component
}

var AppMenus = map[string]templ.Component{
	"fuel-load":    templates.FuelLoad(),
	"takeoff-mins": templates.Takeoffmins(),
	"alternate":    templates.Alternate(),
	// Menu{Item: "Repair", Key: "repair"},
	// Menu{Item: "Holding speed", Key: "holding-speed"},
}

func main() {

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := args[0].String()
		path = strings.ToLower(path)

		component := AppMenus[path]

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
