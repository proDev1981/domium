package signal

import "fmt"

type Signal[T any] struct {
	value T
}

// constructor
func New[T any](value T) *Signal[T] {
	return &Signal[T]{value}
}

/* METHODS */

func (s *Signal[T]) String() string {
	return fmt.Sprint(s.value)
}

func (s *Signal[T]) Get() T {
	return s.value
}

func (s *Signal[T]) Set(value T) {
	s.value = value
}
