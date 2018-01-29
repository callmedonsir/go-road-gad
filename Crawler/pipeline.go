package Crawler

type Pipeline interface {
	Module
	ItemProcessors() []ProcessItem

	Send(item Item) []error

	FailFast() bool

	SetFailFast(failFast bool)
}

type ProcessItem func(item Item)(result Item,err error)
