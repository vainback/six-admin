package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"six-go/config"
)

var db *gorm.DB

func DB() *gorm.DB {
	return connect()
}

func Reconnect() {
	connect(true)
}

func connect(reconnect ...bool) *gorm.DB {
	if len(reconnect) == 0 && db != nil {
		return db
	}
	mysqlDriver := mysql.New(mysql.Config{
		DSN:               config.DB().DSN(),
		DefaultStringSize: 255,
	})
	var err error
	db, err = gorm.Open(mysqlDriver, &gorm.Config{
		Logger: logger.Default.LogMode(config.DB().LogLevel()),
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	if sqlDB, e := db.DB(); e == nil {
		sqlDB.SetMaxOpenConns(config.DB().MaxOpenConns())
		sqlDB.SetMaxIdleConns(config.DB().MaxIdleConns())
		sqlDB.SetConnMaxLifetime(config.DB().ConnMaxLifetime())
	}
	return db
}
