package atom

import "github.com/gopherjs/gopherjs/js"

type Selection struct {
	o js.Object
}

func (s *Selection) GetText() string {
	return s.o.Call("getText").Str()
}

func (s *Selection) InsertText(text string) {
	s.o.Call("insertText", text)
}
