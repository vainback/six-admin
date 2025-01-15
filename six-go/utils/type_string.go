package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type XString string

func String(s string) XString {
	return XString(s)
}

func StringX(v any) XString {
	return XString(fmt.Sprint(v))
}

func Strings(ss ...string) XString {
	return XString(strings.Join(ss, ""))
}

func (s XString) Len() int {
	return len(s)
}

func (s XString) Size() int {
	return utf8.RuneCountInString(s.String())
}

func (s XString) IsEmpty() bool {
	return s.Len() == 0
}

func (s XString) Int() int {
	i, _ := strconv.Atoi(string(s))
	return i
}

func (s XString) Int64() int64 {
	return int64(s.Int())
}

func (s XString) Uint() uint {
	return uint(s.Int())
}

func (s XString) Uint64() uint64 {
	return uint64(s.Int())
}

func (s XString) Float() float64 {
	f, _ := strconv.ParseFloat(string(s), 64)
	return f
}

func (s XString) Bool() bool {
	b, _ := strconv.ParseBool(string(s))
	return b
}

func (s XString) String() string {
	return string(s)
}

func (s XString) Bytes() []byte {
	return []byte(s)
}

func (s XString) FirstUpper() XString {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return XString(runes)
}

func (s XString) FirstLower() XString {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}
	return XString(runes)
}

func (s XString) Strings(str ...string) XString {
	if len(str) == 0 {
		return s
	}
	return Strings(s.String(), Strings(str...).String())
}

func (s XString) StringsLeft(str ...string) XString {
	if len(str) == 0 {
		return s
	}
	return Strings(Strings(str...).String(), s.String())
}

type TrimType int

const (
	TrimLeft TrimType = iota
	TrimRight
)

func (s XString) Trim(cutest string, direction ...TrimType) string {
	if len(direction) == 1 {
		switch direction[0] {
		case TrimLeft:
			return strings.TrimLeft(string(s), cutest)
		case TrimRight:
			return strings.TrimRight(string(s), cutest)
		}
	}
	return strings.Trim(string(s), cutest)
}
