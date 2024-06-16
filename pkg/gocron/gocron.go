package gocron

import (
	"time"

	"github.com/go-co-op/gocron"
)

// Schedule 定时任务 gocron
type Schedule struct {
	*gocron.Scheduler
}

func NewSchedule() *gocron.Scheduler {
	return gocron.NewScheduler(time.Local)
}
