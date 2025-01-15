package v_errors

import (
	"errors"
	"fmt"
)

var (
	CountsZero = errors.New("ErrorCountsZero")
	Unique     = func(filedName string) error {
		return fmt.Errorf("%s已存在", filedName)
	}
	DictUnique = errors.New("该值已在同一个字典类型中存在")
)
