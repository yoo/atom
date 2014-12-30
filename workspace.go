package atom

import "github.com/gopherjs/gopherjs/js"

// https://atom.io/docs/api/v0.165.0/Workspace
type WorkspaceT struct {
	o js.Object
}

func (w *WorkspaceT) GetActiveTextEditor() *TextEditor {
	editor := new(TextEditor)
	editor.o = Workspace.o.Call("getActiveTextEditor")
	return editor
}
