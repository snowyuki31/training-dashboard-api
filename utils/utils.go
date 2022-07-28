package utils

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Chmax[T Number](x *T, a T) {
	if *x >= a {
		return
	}
	*x = a
}
