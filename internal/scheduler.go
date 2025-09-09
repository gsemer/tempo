package internal

import (
	"fmt"
	"tempo/domain"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
	wp   *WorkerPool
}

func NewScheduler(cron *cron.Cron, wp *WorkerPool) *Scheduler {
	return &Scheduler{
		cron: cron,
		wp:   wp,
	}
}

func (s *Scheduler) PeriodicJob(crontab string, job domain.Job) (cron.EntryID, error) {
	return s.cron.AddFunc(
		crontab,
		func() {
			s.wp.Submit(job)
		},
	)
}

func (s *Scheduler) Start() {
	fmt.Println("Start the scheduler")
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	fmt.Println("Stop the scheduler")
	s.cron.Stop()
}
