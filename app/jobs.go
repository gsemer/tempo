package app

import "fmt"

type PrintJob struct{}

func (j *PrintJob) Process() error {
	fmt.Println("PrintJob")
	return nil
}
