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

func login(ctx context.Context) (*supa.AuthenticatedDetails, error) {
	supabaseUrl := "url"
	supabaseKey := "key"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	return supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "email",
		Password: "password",
	})
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
		case "welcome page":
			var existing_name js.Value
			existing_name = js.Global().Get("localStorage").Call("getItem", "name")
			var child templ.Component
			if existing_name.IsNull() {
				child = templates.NameForm()
			} else {
				child = templates.Welcome(existing_name.String())
			}
			component = templates.Layout("Welcome pgae", child)
		case "save-name-form":
			name := args[1].String()
			js.Global().Get("localStorage").Call("setItem", "name", name)
			welcome := templates.Welcome(name)
			component = templates.Layout("Welcome Page", welcome)
		case "remove-name":
			js.Global().Get("localStorage").Call("removeItem", "name")
			name_form := templates.NameForm()
			component = templates.Layout("Welcome Page", name_form)
		case "login":
			login := templates.Login()
			component = templates.Layout("Login", login)
		case "login-progress":
			login_progress := templates.LoginProgress()
			component = templates.Layout("Login", login_progress)
		case "login-init":
			var child templ.Component
			ctx, cancel := context.WithCancel(context.Background())
			auth, err := login(ctx)
			if err != nil {
				cancel()
				child = templates.LoginFail()
			} else {
				child = templates.LoginDetails(auth.User.Email)
			}
			component = templates.Layout("Login", child)
		}

		b := new(strings.Builder)
		component.Render(context.Background(), b)
		return b.String()
	})

	js.Global().Set("go_wasm_handler", cb)

	<-c
}
