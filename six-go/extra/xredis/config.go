package xredis

import (
	"github.com/spf13/cast"
	"six-go/config"
)

const KeyPrefix = "six-admin"

type _config struct {
	data map[string]any
}

func cfg() *_config {
	if data := config.Get("redis"); data != nil {
		return &_config{data: cast.ToStringMap(data)}
	}
	return &_config{data: make(map[string]any)}
}

func (r *_config) addr() string {
	return cast.ToString(r.data["addr"])
}

func (r *_config) pass() string {
	return cast.ToString(r.data["addr"])
}

func (r *_config) db() int {
	return cast.ToInt(r.data["db"])
}
