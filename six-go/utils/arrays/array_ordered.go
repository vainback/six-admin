package arrays

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"slices"
)

type ArrayOrdered[T constraints.Ordered] []T

// NewOrdered 传入的参数必须是同一种基础类型 int int64 float64 string 都不能互相传入进来
func NewOrdered[T constraints.Ordered](arr ...T) ArrayOrdered[T] {
	return arr
}

func Dereference[T constraints.Ordered](v *[]T) ArrayOrdered[T] {
	if v == nil {
		return nil
	}
	return *v
}

func (a ArrayOrdered[T]) Len() int {
	return len(a)
}

func (a ArrayOrdered[T]) Index(x T) int {
	for i, y := range a {
		if a.EqualItem(x, y) {
			return i
		}
	}
	return -1
}

func (a ArrayOrdered[T]) Exist(x T) bool {
	for _, y := range a {
		if a.EqualItem(x, y) {
			return true
		}
	}
	return false
}

func (a ArrayOrdered[T]) Equal(y ArrayOrdered[T]) bool {
	l := a.Len()
	if l != y.Len() {
		return false
	}

	for i := 0; i < l; i++ {
		if !a.EqualItem(a[i], y[i]) {
			return false
		}
	}
	return true
}

func (a ArrayOrdered[T]) EqualItemIndex(index int, y T) bool {
	return a[index] == y
}

func (a ArrayOrdered[T]) EqualItem(x, y T) bool {
	return x == y
}

func (a ArrayOrdered[T]) Merge(x ArrayOrdered[T]) ArrayOrdered[T] {
	return append(a, x...)
}

// Unique 不保留原数组顺序
func (a ArrayOrdered[T]) Unique() ArrayOrdered[T] {
	m := make(map[T]bool)
	for _, x := range a {
		if _, ok := m[x]; !ok {
			m[x] = true
		}
	}
	var y ArrayOrdered[T]
	for k := range m {
		y = append(y, k)
	}
	return y
}

// UniqueOrdered 保留原数组顺序
func (a ArrayOrdered[T]) UniqueOrdered() ArrayOrdered[T] {
	var y = make(ArrayOrdered[T], 0)
	for _, x := range a {
		if !y.Exist(x) {
			y = append(y, x)
		}
	}
	return y
}

func (a ArrayOrdered[T]) Remove(index int) ArrayOrdered[T] {
	return slices.Delete(a, index, index+1)
}

func (a ArrayOrdered[T]) RemoveValue(value T) ArrayOrdered[T] {
	return slices.DeleteFunc(a, func(v T) bool {
		return v == value
	})
}

func (a ArrayOrdered[T]) Push(v T) ArrayOrdered[T] {
	return append(a, v)
}

func (a ArrayOrdered[T]) Replace(oldVal, newVal T) ArrayOrdered[T] {
	if a.Len() > 0 {
		for i, x := range a {
			if x == oldVal {
				a[i] = newVal
			}
		}
	}
	return a
}

func (a ArrayOrdered[T]) ToString() ArrayOrdered[string] {
	var y = make([]string, a.Len())
	for i, v := range a {
		y[i] = fmt.Sprint(v)
	}
	return y
}

func (a ArrayOrdered[T]) ToInt() ArrayOrdered[int] {
	var x any = a
	switch x.(type) {
	case ArrayOrdered[int]:
		return x.(ArrayOrdered[int])
	case ArrayOrdered[int8]:
		return Integer2Int(x.(ArrayOrdered[int8]))
	case ArrayOrdered[int16]:
		return Integer2Int(x.(ArrayOrdered[int16]))
	case ArrayOrdered[int32]:
		return Integer2Int(x.(ArrayOrdered[int32]))
	case ArrayOrdered[int64]:
		return Integer2Int(x.(ArrayOrdered[int64]))
	case ArrayOrdered[uint]:
		return Integer2Int(x.(ArrayOrdered[uint]))
	case ArrayOrdered[uint8]:
		return Integer2Int(x.(ArrayOrdered[uint8]))
	case ArrayOrdered[uint16]:
		return Integer2Int(x.(ArrayOrdered[uint16]))
	case ArrayOrdered[uint32]:
		return Integer2Int(x.(ArrayOrdered[uint32]))
	case ArrayOrdered[uint64]:
		return Integer2Int(x.(ArrayOrdered[uint64]))
	case ArrayOrdered[float32]:
		return Float2Int(x.(ArrayOrdered[float32]))
	case ArrayOrdered[float64]:
		return Float2Int(x.(ArrayOrdered[float64]))
	case ArrayOrdered[string]:
		return String2Int(x.(ArrayOrdered[string]))
	default:
		return nil
	}
}

func (a ArrayOrdered[T]) ToInt64() ArrayOrdered[int64] {
	var x any = a
	switch x.(type) {
	case ArrayOrdered[int]:
		return Integer2Int64(x.(ArrayOrdered[int]))
	case ArrayOrdered[int8]:
		return Integer2Int64(x.(ArrayOrdered[int8]))
	case ArrayOrdered[int16]:
		return Integer2Int64(x.(ArrayOrdered[int16]))
	case ArrayOrdered[int32]:
		return Integer2Int64(x.(ArrayOrdered[int32]))
	case ArrayOrdered[int64]:
		return x.(ArrayOrdered[int64])
	case ArrayOrdered[uint]:
		return Integer2Int64(x.(ArrayOrdered[uint]))
	case ArrayOrdered[uint8]:
		return Integer2Int64(x.(ArrayOrdered[uint8]))
	case ArrayOrdered[uint16]:
		return Integer2Int64(x.(ArrayOrdered[uint16]))
	case ArrayOrdered[uint32]:
		return Integer2Int64(x.(ArrayOrdered[uint32]))
	case ArrayOrdered[uint64]:
		return Integer2Int64(x.(ArrayOrdered[uint64]))
	case ArrayOrdered[float32]:
		return Float2Int64(x.(ArrayOrdered[float32]))
	case ArrayOrdered[float64]:
		return Float2Int64(x.(ArrayOrdered[float64]))
	case ArrayOrdered[string]:
		return String2Int64(x.(ArrayOrdered[string]))
	default:
		return nil
	}
}

func (a ArrayOrdered[T]) ToFloat64() ArrayOrdered[float64] {
	var x any = a
	switch x.(type) {
	case ArrayOrdered[int]:
		return Integer2Float64(x.(ArrayOrdered[int]))
	case ArrayOrdered[int8]:
		return Integer2Float64(x.(ArrayOrdered[int8]))
	case ArrayOrdered[int16]:
		return Integer2Float64(x.(ArrayOrdered[int16]))
	case ArrayOrdered[int32]:
		return Integer2Float64(x.(ArrayOrdered[int32]))
	case ArrayOrdered[int64]:
		return Integer2Float64(x.(ArrayOrdered[int64]))
	case ArrayOrdered[uint]:
		return Integer2Float64(x.(ArrayOrdered[uint]))
	case ArrayOrdered[uint8]:
		return Integer2Float64(x.(ArrayOrdered[uint8]))
	case ArrayOrdered[uint16]:
		return Integer2Float64(x.(ArrayOrdered[uint16]))
	case ArrayOrdered[uint32]:
		return Integer2Float64(x.(ArrayOrdered[uint32]))
	case ArrayOrdered[uint64]:
		return Integer2Float64(x.(ArrayOrdered[uint64]))
	case ArrayOrdered[float32]:
		return Float2Float64(x.(ArrayOrdered[float32]))
	case ArrayOrdered[float64]:
		return x.(ArrayOrdered[float64])
	case ArrayOrdered[string]:
		return String2Float64(x.(ArrayOrdered[string]))
	default:
		return nil
	}
}
