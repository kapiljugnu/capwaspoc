package main

import (
	templates "boozedog/capwaspoc/templ"
	"context"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/hashicorp/go-memdb"
	supa "github.com/nedpals/supabase-go"
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
		&templates.Menu{Item: "About"},
		&templates.Menu{Item: "Home"},
		&templates.Menu{Item: "Welcome Page"},
		&templates.Menu{Item: "Login"},
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

func login() (*supa.AuthenticatedDetails, error) {
	supabaseUrl := "https://ufqekjzxanxjlglrysbw.supabase.co/"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InVmcWVranp4YW54amxnbHJ5c2J3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDkwMzg3ODIsImV4cCI6MjAyNDYxNDc4Mn0.mHDWDGat47YLzV1Bx5ob4fs2YWPuIY8Afqhs5BEm7X8"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	ctx := context.Background()
	return supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "sup@booze.dog",
		Password: "sup",
	})
}

func main() {

	// connect db
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// insert menu
	insert_menus(db)
	// read
	menus := read_menus(db)

	// fmt.Println(menus)

	c := make(chan struct{}, 0)

	var cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		path := args[0].String()
		path = strings.ToLower(path)
		var component templ.Component

		switch path {
		case "sidemenu":
			component = templates.SideMenuRender(menus)
		case "home":
			component = templates.Hello("Home")
		case "about":
			component = templates.Hello("About")
		case "welcome page":
			var existing_name js.Value
			existing_name = js.Global().Get("localStorage").Call("getItem", "name")
			if existing_name.IsNull() {
				component = templates.NameForm()
			} else {
				component = templates.Welcome(existing_name.String())
			}
		case "save-name-form":
			name := args[1].String()
			js.Global().Get("localStorage").Call("setItem", "name", name)
			component = templates.Welcome(name)
		case "remove-name":
			js.Global().Get("localStorage").Call("removeItem", "name")
			component = templates.NameForm()
		case "login":
			component = templates.Login()
		case "login-progress":
			component = templates.LoginProgress()
		case "login-init":
			auth, err := login()
			if err != nil {
				component = templates.LoginFail()
			}
			component = templates.LoginDetails(auth.User.Email)
		}

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
