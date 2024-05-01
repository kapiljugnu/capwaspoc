package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
)

func parseJson(arg string) (map[string]interface{}, error) {
	data := []byte(arg)
	var json_data interface{}
	err := json.Unmarshal(data, &json_data)
	if err != nil {
		return nil, err
	}
	return json_data.(map[string]interface{}), nil
}

func ui_response(title string, child templ.Component) string {
	b := new(strings.Builder)
	component := templates.Layout(title, child)
	component.Render(context.Background(), b)
	return b.String()
}

type Route struct {
	title     string
	component func(templates.JsonData) templ.Component
	data      templates.JsonData
}

var routes = map[string]Route{
	"home":      {title: "Home", component: templates.Hello, data: templates.JsonData{"name": "Hello"}},
	"about":     {title: "About", component: templates.Hello, data: templates.JsonData{"name": "About"}},
	"login":     {title: "Login", component: templates.Login},
	"loggedin":  {title: "Logged In", component: templates.LoggedIn},
	"loginfail": {title: "Login Fail", component: templates.LoginFail},
}

func main() {

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := strings.ToLower(args[0].String())
		json_data := args[1].String()

		var data = make(templates.JsonData)
		var err error
		if json_data != "<undefined>" {
			if data, err = parseJson(json_data); err != nil {
				fmt.Println(err)
				return ui_response("Error", templates.SomethingWrong(nil))
			}
		}

		instance := routes[path]
		if instance.data != nil {
			for key, value := range instance.data {
				data[key] = value
			}
		}
		return ui_response(instance.title, instance.component(data))
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
