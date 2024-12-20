package singleton

import "sync"

type (
	Singleton[T any] struct {
		once  sync.Once
		value T
	}
)

// GetOrInit returns the value of the singleton if it has been initialized, otherwise it initializes the singleton with the given function and returns the value.
func (s *Singleton[T]) GetOrInit(initFunc func() T) T {
	s.once.Do(func() {
		s.value = initFunc()
	})
	return s.value
}
