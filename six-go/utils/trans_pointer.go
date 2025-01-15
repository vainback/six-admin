package utils

func Pointer[T any](v T) *T {
	return &v
}

func PointerTrueOrNil[T bool](v T) *T {
	if v {
		return &v
	}
	return nil
}
