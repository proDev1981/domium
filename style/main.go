package style

import (
	"dom/html"
)

type Style struct {
	Attribute string
	Value     string
}

func (s *Style) Add(parent *html.Element) {}
func (s *Style) IsAttrOrStyle() bool {
	return true
}
func (s *Style) IsEventListener() bool {
	return false
}
