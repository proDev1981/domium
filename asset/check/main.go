package check

type Checker[T any] struct {
	value T
}

// constructor
func Type[T any](value T) *Checker[T] {
	return &Checker[T]{value}
}

func Case[T any, V any](src *Checker[T], action func(V)) {
	if value, ok := any(src.value).(V); ok {
		action(value)
	}
}
