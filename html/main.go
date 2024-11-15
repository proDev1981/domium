package html

import (
	"fmt"
	"strings"
)

type Element struct {
	ID         string
	ClassName  string
	InnerHTML  string
	Style      map[string]string // To store CSS styles as key-value pairs.
	Attributes map[string]string // To store element attributes.
	Children   []*Element        // Child elements of this element.
}

func (c *Element) Add(parent *Element) {}

// NewElement creates a new HTMLElement with the initial values.
func NewElement(id, className, innerHTML string) *Element {
	return &Element{
		ID:         id,
		ClassName:  className,
		InnerHTML:  innerHTML,
		Style:      make(map[string]string),
		Attributes: make(map[string]string),
		Children:   []*Element{},
	}
}

// AddChild adds a child element to the HTMLElement.
func (e *Element) AddChild(child *Element) {
	e.Children = append(e.Children, child)
}

// RemoveChild removes a child element from the HTMLElement.
func (e *Element) RemoveChild(child *Element) {
	for i, c := range e.Children {
		if c == child {
			e.Children = append(e.Children[:i], e.Children[i+1:]...)
			break
		}
	}
}

// SetAttribute sets an attribute on the element.
func (e *Element) SetAttribute(name, value string) {
	e.Attributes[name] = value
}

// GetAttribute retrieves an attribute from the element.
func (e *Element) GetAttribute(name string) string {
	if value, exists := e.Attributes[name]; exists {
		return value
	}
	return ""
}

// SetStyle sets a CSS style property on the element.
func (e *Element) SetStyle(prop, value string) {
	e.Style[prop] = value
}

// String provides the string representation (HTML) of the element.
func (e *Element) String() string {
	var sb strings.Builder

	sb.WriteString("<div")
	if e.ID != "" {
		sb.WriteString(fmt.Sprintf(` id="%s"`, e.ID))
	}
	if e.ClassName != "" {
		sb.WriteString(fmt.Sprintf(` class="%s"`, e.ClassName))
	}
	for key, value := range e.Attributes {
		sb.WriteString(fmt.Sprintf(` %s="%s"`, key, value))
	}
	sb.WriteString(">")

	sb.WriteString(e.InnerHTML)

	for _, child := range e.Children {
		sb.WriteString(child.String())
	}

	sb.WriteString("</div>")

	return sb.String()
}
