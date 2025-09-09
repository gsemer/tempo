package domain

type Job interface {
	Process() error
}
