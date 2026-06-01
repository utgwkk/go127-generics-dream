package option

type Option[T any] struct {
	value T
	ok    bool
}

func Some[T any](value T) *Option[T] {
	return &Option[T]{
		value: value,
		ok: true,
	}
}

func None[T any]() *Option[T] {
	return &Option[T]{
		ok: false,
	}
}

func (o *Option[T]) Map[U any](mapper func(T) U) *Option[U] {
	if o.ok {
		return Some(mapper(o.value))
	}
	return None[U]()
}

func (o *Option[T]) OrElse(fallback T) *Option[T] {
	if o.ok {
		return o
	}
	return Some(fallback)
}

func (o *Option[T]) Unwrap() (T, bool) {
	return o.value, o.ok
}

