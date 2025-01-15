package utils

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

type XInt int64

func Int[T constraints.Integer](i T) XInt {
	return XInt(i)
}

func (x XInt) Int() int {
	return int(x)
}

func (x XInt) Int64() int64 {
	return int64(x)
}

func (x XInt) String() string {
	return strconv.Itoa(x.Int())
}

func (x XInt) Float() float64 {
	return float64(x)
}

func (x XInt) Bool() bool {
	return x != 0
}

func (x XInt) IsZero() bool {
	return x == 0
}

func ToInt[T constraints.Integer](v any) T {
	switch v.(type) {
	case int:
		return T(v.(int))
	case int8:
		return T(v.(int8))
	case int16:
		return T(v.(int16))
	case int32:
		return T(v.(int32))
	case int64:
		return T(v.(int64))
	case uint:
		return T(v.(uint))
	case uint8:
		return T(v.(uint8))
	case uint16:
		return T(v.(uint16))
	case uint32:
		return T(v.(uint32))
	case uint64:
		return T(v.(uint64))
	case float32:
		return T(XFloat(v.(float32)).Int64())
	case float64:
		return T(XFloat(v.(float64)).Int64())
	case string:
		return T(XString(v.(string)).Int64())
	case bool:
		return T(XBool(v.(bool)).Int64())
	default:
		return 0
	}
}
