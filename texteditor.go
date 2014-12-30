package atom

import "github.com/gopherjs/gopherjs/js"

type TextEditor struct {
	o js.Object
}

func (t *TextEditor) GetSelections() []Selection {
	jsSlections := t.o.Call("getSelections")
	l := jsSlections.Length()
	selections := make([]Selection, l)

	for i := 0; i < l; i++ {
		selections[i].o = jsSlections.Index(i)
	}

	return selections
}
