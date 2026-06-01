package pipeline

type Pipeline[A, B any] func(A) B

func Apply[A, B any](x A, p Pipeline[A, B]) B {
	return p(x)
}

func New[B, A any](f func(A) B) Pipeline[A, B] {
	return f
}

func (p Pipeline[A, B]) Then[C any](q func(B) C) Pipeline[A, C] {
	return func(a A) C {
		return q(p(a))
	}
}
