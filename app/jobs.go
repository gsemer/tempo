package app

import (
	"fmt"

	"github.com/google/uuid"
)

type PrintJob struct {
	Id    uuid.UUID
	Type_ string
}

func (j *PrintJob) Process() error {
	fmt.Println("PrintJob")
	return nil
}

func (j *PrintJob) ID() uuid.UUID {
	return j.Id
}

func (j *PrintJob) Type() string {
	return j.Type_
}
