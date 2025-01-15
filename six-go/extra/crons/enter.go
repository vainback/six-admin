package crons

import (
	"log"
	"six-go/database/db"
	"six-go/database/models"
)

func Init() {
	var list []models.CronJob
	if err := db.DB().Where("status = ?", true).Find(&list).Error; err != nil {
		log.Println(err)
		return
	}
	for _, v := range list {
		if cr, ok := Jobs[v.Name]; ok {
			if err := NewCron().AddJob(v.Id, v.Times, cr); err != nil {
				log.Println(err)
			}
		}
	}
	NewCron().Start()
	return
}
