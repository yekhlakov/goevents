package goevents

import (
	"testing"
	"time"
)

func TestEvents(t *testing.T) {

	result := ""

	// no handlers
	Post("ololo", "ururu")
	time.Sleep(50 * time.Millisecond)

	if len(result) > 0 {
		t.Errorf("result failed to initialize properly")
	}

	// one handler
	RegisterHandler("qweqwe", func(event Event) error {
		result += event.Name
		return nil
	})

	Post("qweqwe", 1234)
	time.Sleep(50 * time.Millisecond)
	if result != "qweqwe" {
		t.Errorf("failed to execute one handler (got %s)", result)
	}

	// Two handlers
	RegisterHandler("qweqwe", func(event Event) error {
		result += "!"
		return nil
	})

	result = ""
	Post("qweqwe", 1234)
	time.Sleep(50 * time.Millisecond)
	if result != "qweqwe!" {
		t.Errorf("failed to execute two handlers (got %s)", result)
	}

	// Nested event gets processed after this
	RegisterHandler("nest", func(event Event) error {
		Post("qweqwe", nil)
		result += "nest"
		return nil
	})

	result = ""
	Post("nest", 1234)
	time.Sleep(50 * time.Millisecond)
	if result != "nestqweqwe!" {
		t.Errorf("failed to handle nested event (got %s)", result)
	}

	result = ""
	PostEvent(Event{"nest", 1234})
	time.Sleep(50 * time.Millisecond)
	if result != "nestqweqwe!" {
		t.Errorf("failed to handle nested event (got %s)", result)
	}
}
