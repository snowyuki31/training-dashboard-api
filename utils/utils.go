package utils

func Chmax[T int32 | float32](x *T, a T) {
	if *x >= a {
		return
	}
	*x = a
}
