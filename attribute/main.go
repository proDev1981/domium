package attribute

import (
	"dom/html"
)

type Attribute struct {
	Type  string
	Value string
}

func (a *Attribute) Add(parent *html.Element) {}
func (a *Attribute) IsAttrOrStyle() bool {
	return true
}
func (a *Attribute) IsEventListener() bool {
	return false
}

func Set(typ, value string) *Attribute {
	return &Attribute{typ, value}
}

func Class(value string) *Attribute {
	return Set("class", value)
}
func Id(value string) *Attribute {
	return Set("id", value)
}
func Name(value string) *Attribute {
	return Set("name", value)
}
func Value(value string) *Attribute {
	return Set("value", value)
}
