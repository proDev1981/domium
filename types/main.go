package types

import (
	"dom/html"
)

type Argument interface {
	Add(*html.Element)
}

type ASE interface {
	IsEventListener() bool
}

type AS interface {
	IsAttrOrStyle() bool
}
