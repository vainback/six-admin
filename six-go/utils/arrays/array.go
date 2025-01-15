package arrays

import (
	"reflect"
)

type Array[T any] []T

func New[T any](arr []T) Array[T] {
	return arr
}

func (a Array[T]) Len() int {
	return len(a)
}

func (a Array[T]) Index(x T) int {
	for i, y := range a {
		if a.EqualItem(x, y) {
			return i
		}
	}
	return -1
}

func (a Array[T]) Exist(x T) bool {
	for _, y := range a {
		if a.EqualItem(x, y) {
			return true
		}
	}
	return false
}

func (a Array[T]) Equal(y Array[T]) bool {
	return reflect.DeepEqual(a, y)
}

func (a Array[T]) EqualItemIndex(index int, y T) bool {
	return reflect.DeepEqual(a[index], y)
}

func (a Array[T]) EqualItem(x, y T) bool {
	return reflect.DeepEqual(x, y)
}

func (a Array[T]) Merge(x Array[T]) Array[T] {
	return append(a, x...)
}

// Unique 数组去重, 如果数组元素的类型是 integer|float|string这些元素 请使用 ArrayOrdered.Unique 或 ArrayOrdered.UniqueOrdered
func (a Array[T]) Unique() Array[T] {
	var y = make(Array[T], 0)
	for _, x := range a {
		if !y.Exist(x) {
			y = append(y, x)
		}
	}
	return y
}
