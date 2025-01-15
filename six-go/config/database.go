package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	data map[string]any
}

func DB() *DatabaseConfig {
	if db, ok := _config["database"]; ok {
		return &DatabaseConfig{data: cast.ToStringMap(db)}
	}
	return &DatabaseConfig{data: make(map[string]any)}
}

func (d *DatabaseConfig) AutoMigrate() bool {
	maxOpen, ok := d.data["auto_migrate"]
	if !ok {
		return false
	}
	return cast.ToBool(maxOpen)
}

func (d *DatabaseConfig) LogLevel() logger.LogLevel {
	logLevel, ok := d.data["log_level"]
	if !ok {
		return 1
	}

	switch cast.ToString(logLevel) {
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Silent
	}
}

func (d *DatabaseConfig) MaxOpenConns() int {
	maxOpen, ok := d.data["max_open_conn"]
	if !ok {
		return 10
	}

	maxNum := cast.ToInt(maxOpen)
	if maxNum <= 0 {
		return 10
	}
	return maxNum
}

func (d *DatabaseConfig) MaxIdleConns() int {
	maxIdle, ok := d.data["max_idle_conn"]
	if !ok {
		return 10
	}
	maxNum := cast.ToInt(maxIdle)
	if maxNum <= 0 {
		return 10
	}
	return maxNum
}

func (d *DatabaseConfig) ConnMaxLifetime() time.Duration {
	connMaxLife, ok := d.data["conn_max_life_time"]
	if !ok {
		return 86400 * time.Second
	}
	connMaxLifeNum := cast.ToInt64(connMaxLife)
	if connMaxLifeNum <= 0 {
		return 86400 * time.Second
	}
	return time.Duration(connMaxLifeNum) * time.Second
}

func (d *DatabaseConfig) DSN() string {
	dsn := strings.TrimSpace(cast.ToString(d.data["dsn"]))
	if dsn == "" {
		log.Println("database dsn is empty. at config/database.go line 87")
		return ""
	}
	return dsn
}
