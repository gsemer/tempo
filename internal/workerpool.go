package internal

import (
	"fmt"
	"sync"
	"tempo/domain"
)

type WorkerPool struct {
	workers int
	jobs    chan domain.Job
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int, jobs chan domain.Job, wg *sync.WaitGroup) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		jobs:    jobs,
		wg:      wg,
	}
}

func (wp *WorkerPool) worker(i int) {
	for job := range wp.jobs {
		fmt.Printf("Worker %d processes the job %v\n", i, job.ID())
		err := job.Process()
		if err != nil {
			wp.Submit(job)
		} else {
			wp.wg.Done()
		}
	}
}

func (wp *WorkerPool) Start() {
	fmt.Println("Start worker pool")
	for i := 0; i < wp.workers; i++ {
		go wp.worker(i)
	}
}

func (wp *WorkerPool) Submit(job domain.Job) {
	fmt.Println("Submitting a job")
	wp.wg.Add(1)
	wp.jobs <- job
}

func (wp *WorkerPool) Shutdown() {
	fmt.Println("Shutting down worker pool")
	close(wp.jobs)
}

func (wp *WorkerPool) Wait() {
	fmt.Println("Waiting for all processes to finish their execution")
	wp.wg.Wait()
}
