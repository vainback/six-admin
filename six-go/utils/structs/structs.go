package structs

import (
	"reflect"
)

// todo json tag 为 - 的字段，不会被设置到Map集合之中, json tag omitempty等标记 不会生效

// SliceMap must have json tag
func SliceMap[T any](stSlice []T) []map[string]any {
	length := len(stSlice)
	arrs := make([]map[string]any, length)
	if length > 0 {
		for i := 0; i < length; i++ {
			arrs[i] = Map(stSlice[i])
		}
	}
	return arrs
}

// Map must have json tag
func Map(st any) map[string]any {
	vof := reflect.ValueOf(st)
	tof := reflect.TypeOf(st)
	m := make(map[string]any)
	reflects(tof, vof, &m)
	return m
}

func reflects(tof reflect.Type, vof reflect.Value, w *map[string]any) {
	m := *w
	for i := 0; i < tof.NumField(); i++ {
		if tof.Field(i).Anonymous {
			reflects(tof.Field(i).Type, vof.Field(i), &m)
			continue
		}
		key := tof.Field(i).Tag.Get("json")
		if key != "" && key != "-" {
			value := vof.Field(i).Interface()
			m[key] = value
		}
	}
	w = &m
}
