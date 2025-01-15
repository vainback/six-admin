package crons

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"six-go/extra/xredis"
	"sync"
)

type Cron struct {
	container sync.Map // 定时任务池
	cron      *cron.Cron
}

type job struct {
	id   cron.EntryID
	name string
	spec string
	job  cron.Job
}

var clock *Cron

func NewCron() *Cron {
	if clock == nil {
		clock = &Cron{
			cron: cron.New(cron.WithSeconds()),
		}
	}
	return clock
}

func (c *Cron) AddJob(ormId int64, spec string, cmd cron.Job) error {
	if err := c.RemoveJob(ormId); err != nil {
		return err
	}

	id, err := c.cron.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	return xredis.Set(c.storeKey(ormId), int64(id))
}

func (c *Cron) RemoveJob(ormId int64) error {
	id := xredis.GetInt(c.storeKey(ormId))
	if id > 0 {
		if err := xredis.Del(c.storeKey(ormId)); err != nil {
			log.Println("remove job err:", err)
		}
		c.cron.Remove(cron.EntryID(id))
		c.Restart()
	}
	return nil
}

func (c *Cron) Start() {
	c.cron.Start()
}

func (c *Cron) Stop() {
	c.cron.Stop()
}

func (c *Cron) Restart() {
	c.Stop()
	c.Start()
}

func (c *Cron) storeKey(ormId int64) string {
	return fmt.Sprintf("six-admin:cron:x-%d", ormId)
}
