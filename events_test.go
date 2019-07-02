package goevents

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestEvents(t *testing.T) {

    result := ""

    // no handlers
    Post("ololo", "ururu")
    time.Sleep(50*time.Millisecond)
    assert.Equal(t, "", result)

    // one handler
    RegisterHandler("qweqwe", func(event Event) error {
        result += event.Name
        return nil
    })

    Post("qweqwe", 1234)
    time.Sleep(50*time.Millisecond)
    assert.Equal(t, "qweqwe", result)

    // Two handlers
    RegisterHandler("qweqwe", func(event Event) error {
        result += "!"
        return nil
    })

    result = ""
    Post("qweqwe", 1234)
    time.Sleep(50*time.Millisecond)
    assert.Equal(t, "qweqwe!", result)


    // Nested event gets processed after this
    RegisterHandler ("nest", func(event Event) error {

        Post ("qweqwe", nil)

        result += "nest";

        return nil
    })

    result = ""
    Post("nest", 1234)
    time.Sleep(50*time.Millisecond)
    assert.Equal(t, "nestqweqwe!", result)
}
