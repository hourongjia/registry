package registry

type Event struct {
	Async     bool
	EventType EventType
	Data      EventData
}

type EventData struct {
	OldData string
	NewData string
}

type AsyncResultEvent struct {
	Code AsyncResultCode
	Data interface{}
}
