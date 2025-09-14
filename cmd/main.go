package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"tempo/app"
	"tempo/domain"
	"tempo/internal"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

func main() {
	workers := 5
	jobs := make(chan domain.Job, 5)
	wg := sync.WaitGroup{}

	workerPool := internal.NewWorkerPool(workers, jobs, &wg)
	workerPool.Start()

	cr := cron.New(cron.WithSeconds())
	scheduler := internal.NewScheduler(cr, workerPool)

	_, err := scheduler.PeriodicJob("0 * * * * *", func() domain.Job {
		return &app.PrintJob{
			Id:    uuid.New(),
			Type_: "print_job",
		}
	})
	if err != nil {
		log.Fatal("failed to schedule job:", err)
	}

	scheduler.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	scheduler.Stop()
	workerPool.Shutdown()
	workerPool.Wait()
}
