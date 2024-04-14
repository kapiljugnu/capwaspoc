package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"encoding/json"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/hashicorp/go-memdb"
)

var schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"menu": &memdb.TableSchema{
			Name: "menu",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Item"},
				},
			},
		},
	},
}

func insert_menus(db *memdb.MemDB) {
	txn := db.Txn(true)

	// Insert some people
	menus := []*templates.Menu{
		{Item: "About"},
		{Item: "Home"},
		{Item: "Login"},
	}
	for _, m := range menus {
		if err := txn.Insert("menu", m); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()
}

func read_menus(db *memdb.MemDB) []templates.Menu {
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("menu", "id")
	if err != nil {
		panic(err)
	}

	menus := []templates.Menu{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		m := obj.(*templates.Menu)
		menus = append(menus, *m)
	}

	return menus

}

func main() {

	// // connect db
	// db, err := memdb.NewMemDB(schema)
	// if err != nil {
	// 	panic(err)
	// }

	// // insert menu
	// insert_menus(db)
	// // read
	// menus := read_menus(db)

	// fmt.Println(menus)

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := args[0].String()
		path = strings.ToLower(path)
		var component templ.Component

		switch path {
		case "home":
			home := templates.Hello("Home")
			component = templates.Layout("Home", home)
		case "about":
			about := templates.Hello("About")
			component = templates.Layout("About", about)
		case "login":
			login := templates.Login()
			component = templates.Layout("Login", login)
		case "loggedin":
			data := []byte(args[1].String())
			var user templates.User
			err := json.Unmarshal(data, &user)
			if err != nil {
				component = templates.Layout("Error", templates.SomethingWrong())
			}
			loggedin := templates.LoggedIn(user)
			component = templates.Layout("Logged In", loggedin)
		case "loginfail":
			loginfail := templates.LoginFail()
			component = templates.Layout("Login Fail", loginfail)
		}

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
