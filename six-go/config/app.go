package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

type AppConfig struct {
	data map[string]any
}

func APP() *AppConfig {
	if data, ok := _config["app"]; ok {
		return &AppConfig{data: cast.ToStringMap(data)}
	}
	return &AppConfig{data: make(map[string]any)}
}

func (a *AppConfig) ListenPort() string {
	port := strings.TrimSpace(cast.ToString(a.data["port"]))
	if port == "" {
		return ":8080"
	}
	return fmt.Sprintf(":%s", port)
}

func (a *AppConfig) Debug() bool {
	debug, ok := a.data["debug"]
	if !ok {
		return false
	}
	return cast.ToBool(debug)
}
