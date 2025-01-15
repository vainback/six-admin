package crons

import "errors"

var (
	JobNotExists = errors.New("JobNotExists")
)

// Start 当开启定时任务时调用
func Start(id int64, name, spec string) error {
	if j, ok := Jobs[name]; ok {
		return NewCron().AddJob(id, spec, j)
	}
	return JobNotExists
}

// Stop 当删除或停止定时任务后 调用
func Stop(id int64) error {
	return NewCron().RemoveJob(id)
}
