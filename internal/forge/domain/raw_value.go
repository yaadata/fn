package forgedomain

type Value[T any] struct {
	Formatted T
	Raw       any
}

func NewValue[T any](formatted T, raw any) Value[T] {
	return Value[T]{
		Formatted: formatted,
		Raw:       raw,
	}
}
