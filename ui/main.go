package ui

import (
	"dom/asset/check"
	"dom/html"
	"dom/signal"
	"dom/types"
	"fmt"
)

func Column(args ...types.Argument) *html.Element {
	return &html.Element{}
}
func Row(args ...types.Argument) *html.Element {
	return &html.Element{}
}
func Text[T string | *signal.Signal[string]](data T, args ...types.AS) *html.Element {
	c := check.Type(data)
	check.Case(c, func(str string) {
		fmt.Printf("is string %s", str)
	})
	check.Case(c, func(str *signal.Signal[string]) {
		fmt.Printf("is signal %s", str.String())
	})
	return &html.Element{}
}
func Button(child *html.Element, args ...types.ASE) *html.Element {
	return &html.Element{}
}
func TextField(args ...types.ASE) *html.Element {
	return &html.Element{}
}
func Input(typ string, args ...types.ASE) *html.Element {
	return &html.Element{}
}
