package main

import (
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	chat := tview.NewForm().
		AddTextView("Chat", "Messages", 0, 0, true, true).
		AddTextArea("Message", "", 0, 0, 0, nil)

	flex := tview.NewFlex()

	flex.AddItem(chat, tview.DefaultFormFieldWidth, 1, true)
	flex.SetBorder(true)
	flex.SetTitle("go-gpn-tron")

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
