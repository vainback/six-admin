package db

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
}

type SoftDelete struct {
	DeleteTime gorm.DeletedAt `json:"delete_time" gorm:"index"`
}
