package utils

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

// Ternary 三元运算 等价于 condition ? a : b
func Ternary[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

// Binary 二元运算符 等价于 b := a ?: b | a 不为nil 不为空 不为0值时 取a 否则取b
func Binary[T any](a T, b T) T {
	vof := reflect.ValueOf(a)
	if !vof.IsNil() || !vof.IsZero() {
		return a
	}

	return b
}

// BinaryOrdered  二元运算符 等价于 b := a ?: b | a 不为空字符串 不为0 不为false时取a 否则取b
func BinaryOrdered[T constraints.Ordered](a T, b T) T {
	x := StringX(a)
	if !x.IsEmpty() || x.Int() != 0 || x.Bool() {
		return a
	}
	return b
}
