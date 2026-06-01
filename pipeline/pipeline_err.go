package pipeline

type PipelineErr[A, B any] func(A) (B, error)

func ApplyErr[A, B any](x A, p PipelineErr[A, B]) (B, error) {
	return p(x)
}

func NewErr[B, A any](f func(A) (B, error)) PipelineErr[A, B] {
	return f
}

func (p Pipeline[A, B]) Err() PipelineErr[A, B] {
	return func(a A) (B, error) {
		return p(a), nil
	}
}

func (p PipelineErr[A, B]) Then[C any](q func(B) (C, error)) PipelineErr[A, C] {
	return func(a A) (C, error) {
		b, err := p(a)
		if err != nil {
			var zero C
			return zero, err
		}
		return q(b)
	}
}
