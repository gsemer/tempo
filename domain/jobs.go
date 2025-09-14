package domain

import "github.com/google/uuid"

type Job interface {
	Process() error
	ID() uuid.UUID
	Type() string
}
