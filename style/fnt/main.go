package fnt

import (
	. "dom/style"
	"fmt"
)

func Size(value int) *Style {
	return &Style{Attribute: "tw", Value: fmt.Sprintf("text-size-%d", value)}
}

var Black *Style = &Style{Attribute: "tw", Value: "text-black"}
var White *Style = &Style{Attribute: "tw", Value: "text-white"}
var Transparent *Style = &Style{Attribute: "tw", Value: "text-transparent"}
