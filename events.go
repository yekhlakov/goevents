package goevents

// The event type accepted by event handlers
type Event struct {
	Name    string
	Payload interface{}
}

// The dispatcher instance
var dispatcher struct {
	Queue    chan Event
	Handlers map[string][]func(Event) error
}

// package initialization
func init() {

	dispatcher.Queue = make(chan Event, 10)
	dispatcher.Handlers = make(map[string][]func(Event) error)

	go func() {
		for {
			event := <-dispatcher.Queue

			for _, handler := range dispatcher.Handlers[event.Name] {
				_ = handler(event)
			}
		}
	}()
}

// Register an event handler
func RegisterHandler(name string, handler func(Event) error) {
	dispatcher.Handlers[name] = append(dispatcher.Handlers[name], handler)
}

// Post an event
func Post(name string, payload interface{}) {
	if dispatcher.Handlers[name] == nil {
		return
	}

	dispatcher.Queue <- Event{name, payload}
}

// Post an event (as an object)
func PostEvent(event Event) {
	if dispatcher.Handlers[event.Name] == nil {
		return
	}

	dispatcher.Queue <- event
}
