package domain

type eventSerializer struct {}

func NewEventSerializer() *eventSerializer {
	return &eventSerializer{}
}