package main

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

func main() {
	ctx := context.Background()
	supabaseUrl := "https://ufqekjzxanxjlglrysbw.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InVmcWVranp4YW54amxnbHJ5c2J3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDkwMzg3ODIsImV4cCI6MjAyNDYxNDc4Mn0.mHDWDGat47YLzV1Bx5ob4fs2YWPuIY8Afqhs5BEm7X8"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	auth, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "kr_kapil@hotmail.com",
		Password: "kr_kapil",
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println(auth.User.Email)
	}

}
