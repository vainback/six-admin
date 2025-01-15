package dao

import (
	"log"
	"six-go/database/db"
	"six-go/database/models"
	"six-go/utils"
)

func AutoMigrate() {
	err := db.DB().AutoMigrate(
		&models.AuthRule{},
		&models.AuthRole{},
		&models.AuthRelation{},
		&models.AuthUser{},
		&models.AuthUserLoginStorage{},
		&models.CronJob{},
		&models.Dict{},
		&models.Files{},
		&models.Job{},
		&models.OperateLog{},
		&models.Organization{},
		&models.Tenant{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}

	// 写入Root用户
	db.DB().FirstOrCreate(&models.AuthUser{
		MODEL:          db.MODEL{Id: 1},
		TenantIdentify: models.TenantIdentify{TenantId: 1},
		Username:       "admin",
		Password:       utils.Password.Hash("123456"),
		Nickname:       "six-root",
		Status:         1,
		RoleIds:        []int64{-20240929},
		IsRoot:         true,
		OrganizationId: 1,
		JobId:          1,
	}, "id = ?", 1)
}
