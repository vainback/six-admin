package crons

import "github.com/robfig/cron/v3"

// Jobs Key is unique name Value is FuncJob func.
var Jobs = map[string]cron.FuncJob{
	"example": ExampleJob,
}
