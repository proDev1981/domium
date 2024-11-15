package main

import (
	attr "dom/attribute"
	"dom/on"
	"dom/signal"
	"dom/style/fnt"
	"dom/ui"
	"fmt"
)

func main() {

	state := signal.New("hola")

	ui.Column(
		ui.Button(
			ui.Text(state),
			on.Click(func() { fmt.Println("hola mundo") }),
		),
		ui.TextField(attr.Value("Hola"), fnt.Size(200)),
		ui.Input("number"),
	)

}
