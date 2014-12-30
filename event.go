package atom

import "github.com/gopherjs/gopherjs/js"

// Event doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-onWillThrowError
type Event struct {
	o             js.Object
	OriginalError js.Object
	Message       string
	URL           string
	Line          int
	Column        int
}

func newEventFromJsObject(jsEvent js.Object) *Event {
	event := &Event{o: jsEvent}
	event.OriginalError = event.o.Get("originalError")
	event.Message = event.o.Get("message").Str()
	event.URL = event.o.Get("url").Str()
	event.Line = event.o.Get("line").Int()
	event.Column = event.o.Get("column").Int()
	return event
}

// PreventDefault doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-onWillThrowError
func (e *Event) PreventDefault() {
	e.o.Call("preventDefault")
}
