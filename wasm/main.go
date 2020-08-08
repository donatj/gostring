package main

import (
	"fmt"
	myjs "syscall/js"
)

var document = myjs.Global().Get("document")

func getElementByID(id string) myjs.Value {
	return document.Call("getElementById", id)
}

func renderEditor(parent myjs.Value) myjs.Value {
	editorMarkup := `
<div id="editor">
	<textarea id="input" style="" placeholder="input text"></textarea>
	<textarea id="preview" readonly placeholder="output go string"></textarea>
</div>
	`
	parent.Call("insertAdjacentHTML", "beforeend", editorMarkup)
	return getElementByID("editor")
}

func main() {
	quit := make(chan struct{}, 0)

	editor := renderEditor(document.Get("body"))
	preview := getElementByID("preview")

	input := getElementByID("input")

	// renderButton := getElementByID("render")
	input.Set("oninput", myjs.FuncOf(func(this myjs.Value, inputs []myjs.Value) interface{} {

		preview.Set("textContent", fmt.Sprintf("%#v\n", input.Get("value").String()))

		return nil
	}))

	preview.Set("onclick", myjs.FuncOf(func(this myjs.Value, inputs []myjs.Value) interface{} {
		preview.Call("select")

		return nil
	}))

	<-quit
	editor.Call("remove")
}
