package utils

import (
	"math/rand/v2"
	"strings"
)

func RandMix(l int) string {
	if l <= 0 {
		return ""
	}
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	var str strings.Builder
	for i := 0; i < l; i++ {
		str.WriteString(chars[randInt(0, 61)])
	}
	return str.String()
}

func RandLetter(l int) string {
	if l <= 0 {
		return ""
	}
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
	}
	var str strings.Builder
	for i := 0; i < l; i++ {
		str.WriteString(chars[randInt(0, 51)])
	}
	return str.String()
}

func RandUpper(l int) string {
	if l <= 0 {
		return ""
	}
	chars := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
	}
	var str strings.Builder
	for i := 0; i < l; i++ {
		str.WriteString(chars[randInt(0, 25)])
	}
	return str.String()
}

func RandLower(l int) string {
	if l <= 0 {
		return ""
	}
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
	}
	var str strings.Builder
	for i := 0; i < l; i++ {
		str.WriteString(chars[randInt(0, 25)])
	}
	return str.String()
}

func RandNumber(l int) string {
	if l <= 0 {
		return ""
	}
	chars := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	var str strings.Builder
	for i := 0; i < l; i++ {
		str.WriteString(chars[randInt(0, 9)])
	}
	return str.String()
}

func randInt(minValue, maxValue int) int {
	return rand.IntN(maxValue-minValue) + minValue
}
