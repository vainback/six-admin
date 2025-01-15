package arrays

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

func Integer2Int64[T constraints.Integer](in ArrayOrdered[T]) ArrayOrdered[int64] {
	out := make(ArrayOrdered[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = int64(in[i])
	}
	return out
}

func Float2Int64[T constraints.Float](in ArrayOrdered[T]) ArrayOrdered[int64] {
	out := make(ArrayOrdered[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		v, _ := strconv.Atoi(fmt.Sprintf("%.0f", in[i]))
		out[i] = int64(v)
	}
	return out
}

func String2Int64(in ArrayOrdered[string]) ArrayOrdered[int64] {
	out := make(ArrayOrdered[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		v, _ := strconv.Atoi(in[i])
		out[i] = int64(v)
	}
	return out
}

func Integer2Int[T constraints.Integer](in ArrayOrdered[T]) ArrayOrdered[int] {
	out := make(ArrayOrdered[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = int(in[i])
	}
	return out
}

func Float2Int[T constraints.Float](in ArrayOrdered[T]) ArrayOrdered[int] {
	out := make(ArrayOrdered[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.Atoi(fmt.Sprintf("%.0f", in[i]))
	}
	return out
}

func String2Int(in ArrayOrdered[string]) ArrayOrdered[int] {
	out := make(ArrayOrdered[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.Atoi(in[i])
	}
	return out
}

func Number2String[T constraints.Ordered](in ArrayOrdered[T]) ArrayOrdered[string] {
	out := make(ArrayOrdered[string], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = fmt.Sprint(in[i])
	}
	return out
}

func Integer2Float64[T constraints.Integer](in ArrayOrdered[T]) ArrayOrdered[float64] {
	out := make(ArrayOrdered[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = float64(in[i])
	}
	return out
}

func Float2Float64[T constraints.Float](in ArrayOrdered[T]) ArrayOrdered[float64] {
	out := make(ArrayOrdered[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = float64(in[i])
	}
	return out
}

func String2Float64(in ArrayOrdered[string]) ArrayOrdered[float64] {
	out := make(ArrayOrdered[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.ParseFloat(in[i], 64)
	}
	return out
}
