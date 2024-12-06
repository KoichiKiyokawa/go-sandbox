package domain

// 同一パッケージ内からしか再代入できないようにする機構
type readonly[T any] struct {
	value T
}

func toReadonly[T any](v T) readonly[T] {
	return readonly[T]{v}
}

func (r readonly[T]) Value() T {
	return r.value
}
