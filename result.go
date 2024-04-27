package berr

import (
	"log"
)

type result[T any] struct {
	value T
	err   error
}

func (res *result[T]) ValueOr(fallback T) T {
	if res.err != nil {
		return fallback
	}

	return res.value
}

func (res *result[T]) ValueOrPanic() T {
	if res.err != nil {
		log.Panicln(res.err)
	}

	return res.value
}

func (res *result[T]) ValueOrFatal() T {
	if res.err != nil {
		log.Fatalln(res.err)
	}

	return res.value
}

func Get[T any](value T, err error) *result[T] {
	return &result[T]{value: value, err: err}
}
