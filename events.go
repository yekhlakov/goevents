package goevents

import (
    "reflect"
)

type Event struct {
    Name    string
    Payload interface{}
}

type Dispatcher struct {
    isStarted bool
    Queue    chan Event
    Handlers map[string][]func(Event) error
}

func getType(x interface{}) string {
    return reflect.TypeOf(x).String()
}

var dispatcher = Dispatcher{
    isStarted: false,
    Queue:    make(chan Event, 10),
    Handlers: make(map[string][]func(Event) error),
}

func RegisterHandler(name string, handler func(Event) error) {

    dispatcher.Handlers[name] = append(dispatcher.Handlers[name], handler)

}

func Post(name string, payload interface{}) {
    if dispatcher.Handlers[name] == nil {
        return
    }

    if !dispatcher.isStarted {
        go execute()
    }

    dispatcher.Queue <- Event{name, payload}
}

func execute() []error {

    for {
        event := <-dispatcher.Queue

        for _, handler := range dispatcher.Handlers[event.Name] {
            _ = handler(event)
        }
    }
}
