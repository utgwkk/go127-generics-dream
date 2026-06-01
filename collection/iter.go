package collection

import (
	"iter"
)

type Iter[T any] iter.Seq[T]

func (it Iter[T]) Map[U any](f func(x T) U) Iter[U] {
	return func(yield func(U) bool) {
		for x := range it {
			if !yield(f(x)) {
				return
			}
		}
	}
}

func (it Iter[T]) Filter(f func(x T) bool) Iter[T] {
	return func(yield func(T) bool) {
		for x := range it {
			if !f(x) {
				continue
			}
			if !yield(x) {
				return
			}
		}
	}
}

func (it Iter[T]) Unwrap() iter.Seq[T] {
	return iter.Seq[T](it)
}
