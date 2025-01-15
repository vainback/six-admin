package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(value string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))
}

func Md5Byte(value []byte) []byte {
	sum := md5.Sum(value)
	return sum[:]
}
