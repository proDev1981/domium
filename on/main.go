package on

import (
	"dom/html"
)

type EventListener struct {
	Type     string
	CallBack func()
}

func (h *EventListener) Add(parent *html.Element) {}
func (h *EventListener) IsEventListener() bool {
	return true
}

func Click(callBack func()) *EventListener {
	return &EventListener{"click", callBack}
}
