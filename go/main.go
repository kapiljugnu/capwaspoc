package main

import (
	"syscall/js"
)

func main() {
	// Register a JavaScript function that we can call from the web page
	js.Global().Set("generateHTML", js.FuncOf(generateHTML))

	// Keep the program running to handle JavaScript callbacks
	select {}
}

func generateHTML(this js.Value, p []js.Value) interface{} {
	// Get the document object from the web page
	document := js.Global().Get("document")

	// Create a new <p> element
	paragraph := document.Call("createElement", "p")
	paragraph.Set("textContent", "Hello, WebAssembly!")

	// Append the <p> element to the body of the document
	body := document.Get("body")
	body.Call("appendChild", paragraph)

	return nil
}
