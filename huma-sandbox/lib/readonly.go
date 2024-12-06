package lib

// 同一パッケージ内からしか再代入できないようにする機構
type Readonly[T any] struct {
	value T
}

func ToReadonly[T any](v T) Readonly[T] {
	return Readonly[T]{value: v}
}

func (r Readonly[T]) Value() T {
	return r.value
}
