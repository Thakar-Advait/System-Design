package main

import (
	"fmt"

	"example.com/abstractfactory/ui"
)

func main() {
	ui, err := ui.NewUIFactory("dark")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	btn := ui.CreateButton()
	txt := ui.CreateTextbox()

	fmt.Println(btn.Render())
	fmt.Println(txt.Render())
}
